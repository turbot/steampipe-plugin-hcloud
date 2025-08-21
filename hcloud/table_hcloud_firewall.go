package hcloud

import (
	"context"

	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableHcloudFirewall(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_firewall",
		Description: "Firewall resources in Hetzner Cloud.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"name"}),
			Hydrate:    listFirewall,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getFirewall,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Firewall."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the Firewall."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "User-defined labels."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Creation time of the Firewall."},
			{Name: "rules", Type: proto.ColumnType_JSON, Description: "Firewall rules."},
			{Name: "applied_to", Type: proto.ColumnType_JSON, Description: "Resources this firewall is applied to."},
		},
	}
}

func listFirewall(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_firewall.listFirewall", "connection_error", err)
		return nil, err
	}

	opts := hcloudgo.FirewallListOpts{ListOpts: hcloudgo.ListOpts{Page: 1, PerPage: 50}}

	if d.EqualsQuals["name"] != nil {
		opts.Name = d.EqualsQuals["name"].GetStringValue()
	}

	for {
		items, resp, err := conn.Firewall.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("hcloud_firewall.listFirewall", "query_error", err)
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

func getFirewall(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_firewall.getFirewall", "connection_error", err)
		return nil, err
	}

	id := int(d.EqualsQuals["id"].GetInt64Value())
	if id == 0 {
		return nil, nil
	}

	item, resp, err := conn.Firewall.GetByID(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_firewall.getFirewall", "query_error", err, "resp", resp)
		return nil, err
	}

	return item, nil
}
