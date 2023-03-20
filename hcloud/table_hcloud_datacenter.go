package hcloud

import (
	"context"

	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHcloudDataCenter(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_datacenter",
		Description: "Data centers available to the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			Hydrate: listDataCenter,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getDataCenter,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Data Center."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique identifier of the Data Center."},
			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the Data Center."},
			{Name: "location_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Location.ID"), Description: "Location of the Data Center."},
			{Name: "server_types_available", Type: proto.ColumnType_JSON, Transform: transform.FromField("ServerTypes.Available").Transform(serverTypeArrayToServerTypeIDArray), Description: "IDs of the server types available at the Data Center."},
			{Name: "server_types_supported", Type: proto.ColumnType_JSON, Transform: transform.FromField("ServerTypes.Supported").Transform(serverTypeArrayToServerTypeIDArray), Description: "IDs of the server types supported at the Data Center."},
		},
	}
}

func serverTypeArrayToServerTypeIDArray(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	var ids []int
	items := input.Value.([]*hcloudgo.ServerType)
	for _, i := range items {
		ids = append(ids, i.ID)
	}
	return ids, nil
}

func listDataCenter(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_datacenter.listDataCenter", "connection_error", err)
		return nil, err
	}
	items, err := conn.Datacenter.All(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_datacenter.listDataCenter", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getDataCenter(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_datacenter.getDataCenter", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.Datacenter
	var resp *hcloudgo.Response
	if d.EqualsQuals["id"] != nil {
		id := int(d.EqualsQuals["id"].GetInt64Value())
		item, resp, err = conn.Datacenter.GetByID(ctx, id)
	} else if d.EqualsQuals["name"] != nil {
		name := d.EqualsQuals["name"].GetStringValue()
		item, resp, err = conn.Datacenter.GetByName(ctx, name)
	}
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_datacenter.getDataCenter", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
