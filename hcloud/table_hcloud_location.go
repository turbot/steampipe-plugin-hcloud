package hcloud

import (
	"context"

	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableHcloudLocation(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_location",
		Description: "Locations available to the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			Hydrate: listLocation,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getLocation,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Location."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique identifier of the Location."},
			// Other columns
			{Name: "city", Type: proto.ColumnType_STRING, Description: "City the Location is closest to."},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "ISO 3166-1 alpha-2 code of the country the Location resides in."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the Location."},
			{Name: "latitude", Type: proto.ColumnType_DOUBLE, Description: "Latitude of the city closest to the Location."},
			{Name: "longitude", Type: proto.ColumnType_DOUBLE, Description: "Longitude of the city closest to the Location."},
			{Name: "network_zone", Type: proto.ColumnType_STRING, Description: "Name of network zone this Location resides in."},
		},
	}
}

func listLocation(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_location.listLocation", "connection_error", err)
		return nil, err
	}
	items, err := conn.Location.All(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_location.listLocation", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getLocation(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_location.getLocation", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.Location
	var resp *hcloudgo.Response
	if d.KeyColumnQuals["id"] != nil {
		id := int(d.KeyColumnQuals["id"].GetInt64Value())
		item, resp, err = conn.Location.GetByID(ctx, id)
	} else if d.KeyColumnQuals["name"] != nil {
		name := d.KeyColumnQuals["name"].GetStringValue()
		item, resp, err = conn.Location.GetByName(ctx, name)
	}
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_location.getLocation", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
