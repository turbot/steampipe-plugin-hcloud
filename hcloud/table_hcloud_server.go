package hcloud

import (
	"context"

	"github.com/hetznercloud/hcloud-go/hcloud"
	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableHcloudServer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_server",
		Description: "Servers in the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"name", "status"}),
			Hydrate:    listServer,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getServer,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Server."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the Server. Must be unique per Project."},
			// Other columns
			{Name: "backup_window", Type: proto.ColumnType_STRING, Description: "Time window (UTC) in which the backup will run, or null if the backups are not enabled."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Point in time when the Server was created."},
			{Name: "datacenter_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Datacenter.ID"), Description: "Datacenter this Server is located at."},
			{Name: "image_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Image.ID"), Description: "ID of the base image for the server."},
			{Name: "included_traffic", Type: proto.ColumnType_INT, Description: "Free Traffic for the current billing period in bytes."},
			{Name: "ingoing_traffic", Type: proto.ColumnType_INT, Description: "Inbound Traffic for the current billing period in bytes."},
			{Name: "iso", Type: proto.ColumnType_JSON, Transform: transform.FromField("ISO"), Description: "ISO Image that is attached to this Server. Null if no ISO is attached."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "User-defined labels (key-value pairs)."},
			{Name: "load_balancers", Type: proto.ColumnType_JSON, Description: "Array of load balancer IDs."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Description: "True if Server has been locked and is not available to user."},
			{Name: "outgoing_traffic", Type: proto.ColumnType_INT, Description: "Outbound Traffic for the current billing period in bytes."},
			{Name: "placement_group", Type: proto.ColumnType_JSON, Description: "Placement group for the server."},
			{Name: "primary_disk_size", Type: proto.ColumnType_INT, Description: "Size of the primary Disk."},
			{Name: "private_net", Type: proto.ColumnType_JSON, Description: "Private network information."},
			{Name: "protection", Type: proto.ColumnType_JSON, Description: "Protection configuration for the Server."},
			{Name: "public_net", Type: proto.ColumnType_JSON, Description: "Public network information. The Server's IPv4 address can be found in public_net->ipv4->ip."},
			{Name: "rescue_enabled", Type: proto.ColumnType_BOOL, Description: "True if rescue mode is enabled. Server will then boot into rescue system on next reboot."},
			{Name: "server_type_id", Type: proto.ColumnType_INT, Transform: transform.FromField("ServerType.ID"), Description: "Type of Server - determines how much ram, disk and cpu a Server has."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the server: running, initializing, starting, stopping, off, deleting, migrating, rebuilding, unknown."},
			{Name: "volume_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Volumes").Transform(volumeArrayToVolumeIDArray), Description: "IDs of Volumes assigned to this Server."},
		},
	}
}

func volumeArrayToVolumeIDArray(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	var ids []int
	// Avoid errors if the transform is run without data (shouldn't happen, but did in testing on 2021-10-06)
	if input.Value == nil {
		return ids, nil
	}
	items := input.Value.([]*hcloudgo.Volume)
	for _, i := range items {
		ids = append(ids, i.ID)
	}
	return ids, nil
}

func listServer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_server.listServer", "connection_error", err)
		return nil, err
	}

	opts := hcloudgo.ServerListOpts{ListOpts: hcloudgo.ListOpts{Page: 1, PerPage: 50}}

	if d.KeyColumnQuals["name"] != nil {
		opts.Name = d.KeyColumnQuals["name"].GetStringValue()
	}
	if d.KeyColumnQuals["status"] != nil {
		opts.Status = []hcloud.ServerStatus{hcloud.ServerStatus(d.KeyColumnQuals["status"].GetStringValue())}
	}

	for {
		items, resp, err := conn.Server.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("hcloud_server.listServer", "query_error", err)
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

func getServer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_server.getServer", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.Server
	var resp *hcloudgo.Response
	if d.KeyColumnQuals["id"] != nil {
		id := int(d.KeyColumnQuals["id"].GetInt64Value())
		item, resp, err = conn.Server.GetByID(ctx, id)
	} else if d.KeyColumnQuals["name"] != nil {
		name := d.KeyColumnQuals["name"].GetStringValue()
		item, resp, err = conn.Server.GetByName(ctx, name)
	}
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_server.getServer", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
