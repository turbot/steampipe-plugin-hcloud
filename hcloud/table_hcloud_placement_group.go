package hcloud

import (
	"context"

	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableHcloudPlacementGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_placement_group",
		Description: "Placement groups available in the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			Hydrate: listPlacementGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getPlacementGroup,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the placement group."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique user-defined name of the placement group."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Point in time (in ISO-8601 format) when the placement group was created."},
			{Name: "servers", Type: proto.ColumnType_JSON, Description: "Array of IDs of servers that are part of this placement group."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of placement group."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "User-defined labels (key-value pairs)."},
		},
	}
}

func listPlacementGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_placement_group.listPlacementGroups", "connection_error", err)
		return nil, err
	}

	items, err := conn.PlacementGroup.All(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_placement_group.listPlacementGroups", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}

	return nil, nil
}

func getPlacementGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_placement_group.getPlacementGroup", "connection_error", err)
		return nil, err
	}

	var item *hcloudgo.PlacementGroup
	var resp *hcloudgo.Response

	if d.KeyColumnQuals["id"] != nil {
		id := int(d.KeyColumnQuals["id"].GetInt64Value())
		item, resp, err = conn.PlacementGroup.GetByID(ctx, id)
	} else if d.KeyColumnQuals["name"] != nil {
		name := d.KeyColumnQuals["name"].GetStringValue()
		item, resp, err = conn.PlacementGroup.GetByName(ctx, name)
	}

	if err != nil {
		plugin.Logger(ctx).Error("hcloud_network.getPlacementGroup", "query_error", err, "resp", resp)
		return nil, err
	}

	return item, err
}
