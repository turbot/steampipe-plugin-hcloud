package hcloud

import (
	"context"

	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableHcloudNetwork(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_network",
		Description: "Data centers available to the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			Hydrate: listNetwork,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getNetwork,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Data Center."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique identifier of the Data Center."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Point in time when the Network was created."},
			{Name: "ip_range", Type: proto.ColumnType_JSON, Description: "IPv4 prefix of the whole Network."},
			{Name: "subnets", Type: proto.ColumnType_JSON, Description: "Array subnets allocated in this Network."},
			{Name: "routes", Type: proto.ColumnType_JSON, Description: "Array of routes set in this Network."},
			{Name: "server_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Servers").Transform(serverArrayToServerIDArray), Description: "Array of IDs of Servers attached to this Network."},
			{Name: "protection", Type: proto.ColumnType_JSON, Description: "Protection configuration for the Network."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "User-defined labels (key-value pairs)."},
			{Name: "raw", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: ""},
		},
	}
}

func serverArrayToServerIDArray(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	var ids []int
	items := input.Value.([]*hcloudgo.Server)
	for _, i := range items {
		ids = append(ids, i.ID)
	}
	return ids, nil
}

func listNetwork(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_network.listNetwork", "connection_error", err)
		return nil, err
	}
	items, err := conn.Network.All(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_network.listNetwork", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getNetwork(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_network.getNetwork", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.Network
	var resp *hcloudgo.Response
	if d.KeyColumnQuals["id"] != nil {
		id := int(d.KeyColumnQuals["id"].GetInt64Value())
		item, resp, err = conn.Network.GetByID(ctx, id)
	} else if d.KeyColumnQuals["name"] != nil {
		name := d.KeyColumnQuals["name"].GetStringValue()
		item, resp, err = conn.Network.GetByName(ctx, name)
	}
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_network.getNetwork", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
