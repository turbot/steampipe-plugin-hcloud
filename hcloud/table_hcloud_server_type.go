package hcloud

import (
	"context"

	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHcloudServerType(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_server_type",
		Description: "Server types in Hetzner Cloud.",
		List: &plugin.ListConfig{
			Hydrate: listServerType,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getServerType,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Server Type."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique identifier of the Server Type."},
			// Other columns
			{Name: "cores", Type: proto.ColumnType_INT, Transform: transform.FromField("Cores"), Description: "Number of cpu cores a Server of this type will have."},
			{Name: "cpu_type", Type: proto.ColumnType_STRING, Description: "Type of CPU: shared, dedicated."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the Server Type."},
			{Name: "disk", Type: proto.ColumnType_INT, Transform: transform.FromField("Disk"), Description: "Disk size a Server of this type will have in GB."},
			{Name: "memory", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Memory"), Description: "Memory a Server of this type will have in GB."},
			{Name: "prices", Type: proto.ColumnType_JSON, Transform: transform.FromField("Pricings"), Description: "Prices in different Locations."},
			{Name: "storage_type", Type: proto.ColumnType_STRING, Description: "Type of Server boot drive. local has higher speed. network has better availability."},
		},
	}
}

func listServerType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_server_type.listServerType", "connection_error", err)
		return nil, err
	}
	items, err := conn.ServerType.All(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_server_type.listServerType", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getServerType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_server_type.getServerType", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.ServerType
	var resp *hcloudgo.Response
	if d.EqualsQuals["id"] != nil {
		id := int(d.EqualsQuals["id"].GetInt64Value())
		item, resp, err = conn.ServerType.GetByID(ctx, id)
	} else if d.EqualsQuals["name"] != nil {
		name := d.EqualsQuals["name"].GetStringValue()
		item, resp, err = conn.ServerType.GetByName(ctx, name)
	}
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_server_type.getServerType", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
