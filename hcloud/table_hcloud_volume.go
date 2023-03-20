package hcloud

import (
	"context"

	"github.com/hetznercloud/hcloud-go/hcloud"
	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHcloudVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_volume",
		Description: "Volumes in the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"name", "status"}),
			Hydrate:    listVolume,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getVolume,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Volume."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the Volume. Must be unique per Project."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Point in time when the Volume was created."},
			{Name: "format", Type: proto.ColumnType_STRING, Description: "Filesystem of the Volume if formatted on creation, null if not formatted on creation."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "User-defined labels (key-value pairs)."},
			{Name: "linux_device", Type: proto.ColumnType_STRING, Description: "Device path on the file system for the Volume."},
			{Name: "location_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Location.ID"), Description: "Location of the Volume. Volume can only be attached to Servers in the same Location."},
			{Name: "protection", Type: proto.ColumnType_JSON, Description: "Protection configuration for the Volume."},
			{Name: "server_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Server.ID"), Description: "ID of the Server the Volume is attached to, null if it is not attached at all."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Size in GB of the Volume."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the volume: creating, available."},
		},
	}
}

func listVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_volume.listVolume", "connection_error", err)
		return nil, err
	}

	opts := hcloudgo.VolumeListOpts{ListOpts: hcloudgo.ListOpts{Page: 1, PerPage: 50}}

	if d.EqualsQuals["name"] != nil {
		opts.Name = d.EqualsQuals["name"].GetStringValue()
	}
	if d.EqualsQuals["status"] != nil {
		opts.Status = []hcloud.VolumeStatus{hcloud.VolumeStatus(d.EqualsQuals["status"].GetStringValue())}
	}

	for {
		items, resp, err := conn.Volume.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("hcloud_volume.listVolume", "query_error", err)
			return nil, err
		}
		for _, i := range items {
			d.StreamListItem(ctx, i)
		}
		opts.ListOpts.Page++
		if resp.Meta.Pagination.Page >= resp.Meta.Pagination.LastPage {
			break
		}
	}

	return nil, nil
}

func getVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_volume.getVolume", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.Volume
	var resp *hcloudgo.Response
	if d.EqualsQuals["id"] != nil {
		id := int(d.EqualsQuals["id"].GetInt64Value())
		item, resp, err = conn.Volume.GetByID(ctx, id)
	} else if d.EqualsQuals["name"] != nil {
		name := d.EqualsQuals["name"].GetStringValue()
		item, resp, err = conn.Volume.GetByName(ctx, name)
	}
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_volume.getVolume", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
