package hcloud

import (
	"context"

	"github.com/hetznercloud/hcloud-go/hcloud"
	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHcloudImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_image",
		Description: "Images in the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"image_type", "name", "status"}),
			Hydrate:    listImage,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getImage,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Image."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique identifier of the Image. This value is only set for system Images."},
			// Other columns
			{Name: "bound_to", Type: proto.ColumnType_INT, Transform: transform.FromField("BoundTo.ID").NullIfEqual(0), Description: "ID of Server the Image is bound to. Only set for Images of type backup."},
			{Name: "build_id", Type: proto.ColumnType_STRING, Description: "Build ID of the Image."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Point in time when the Image was created."},
			{Name: "created_from", Type: proto.ColumnType_JSON, Description: "Information about the Server the Image was created from."},
			{Name: "deleted", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Deleted").NullIfZero(), Description: "Point in time when the Image was deleted."},
			{Name: "deprecated", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Deprecated").NullIfZero(), Description: "Point in time when the Image was deprecated."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the Image."},
			{Name: "disk_size", Type: proto.ColumnType_INT, Description: "Size of the disk contained in the Image in GB."},
			{Name: "image_size", Type: proto.ColumnType_INT, Description: "Size of the Image file in our storage in GB. For snapshot Images this is the value relevant for calculating costs for the Image."},
			{Name: "image_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "Type of the Image: system, app, snapshot, backup, temporary."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "User-defined labels (key-value pairs)."},
			{Name: "os_flavor", Type: proto.ColumnType_STRING, Description: "Flavor of operating system contained in the Image: ubuntu, centos, debian, fedora, unknown."},
			{Name: "os_version", Type: proto.ColumnType_STRING, Description: "Operating system version."},
			{Name: "protection", Type: proto.ColumnType_JSON, Description: "Protection configuration for the Resource."},
			{Name: "rapid_deploy", Type: proto.ColumnType_BOOL, Description: "Indicates that rapid deploy of the Image is available."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Whether the Image can be used or if it's still being created or unavailable: available, creating, unavailable."},
		},
	}
}

func listImage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_image.listImage", "connection_error", err)
		return nil, err
	}

	opts := hcloudgo.ImageListOpts{ListOpts: hcloudgo.ListOpts{Page: 1, PerPage: 50}, IncludeDeprecated: true}

	if d.EqualsQuals["image_type"] != nil {
		opts.Type = []hcloud.ImageType{hcloud.ImageType(d.EqualsQuals["image_type"].GetStringValue())}
	}
	if d.EqualsQuals["name"] != nil {
		opts.Name = d.EqualsQuals["name"].GetStringValue()
	}
	if d.EqualsQuals["status"] != nil {
		opts.Status = []hcloud.ImageStatus{hcloud.ImageStatus(d.EqualsQuals["status"].GetStringValue())}
	}

	for {
		items, resp, err := conn.Image.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("hcloud_image.listImage", "query_error", err)
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

func getImage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_image.getImage", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.Image
	var resp *hcloudgo.Response
	if d.EqualsQuals["id"] != nil {
		id := int(d.EqualsQuals["id"].GetInt64Value())
		item, resp, err = conn.Image.GetByID(ctx, id)
	} else if d.EqualsQuals["name"] != nil {
		name := d.EqualsQuals["name"].GetStringValue()
		item, resp, err = conn.Image.GetByName(ctx, name)
	}
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_image.getImage", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
