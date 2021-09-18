package hcloud

import (
	"context"

	"github.com/hetznercloud/hcloud-go/hcloud"
	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableHcloudAction(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_action",
		Description: "Actions performed in the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"id", "status"}),
			Hydrate:    listAction,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAction,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the Action."},
			{Name: "command", Type: proto.ColumnType_STRING, Description: "Command executed in the Action."},
			{Name: "started", Type: proto.ColumnType_TIMESTAMP, Description: "Point in time when the Action was started."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the Action: success, running, error."},
			// Other columns
			{Name: "error", Type: proto.ColumnType_JSON, Description: "Error message for the Action if error occurred, otherwise null."},
			{Name: "finished", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Finished").NullIfZero(), Description: "Point in time when the Action was finished. Only set if the Action is finished otherwise null."},
			{Name: "progress", Type: proto.ColumnType_INT, Description: "Progress of Action in percent."},
			{Name: "resources", Type: proto.ColumnType_JSON, Description: "Resources the Action relates to."},
		},
	}
}

func listAction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_action.listAction", "connection_error", err)
		return nil, err
	}

	opts := hcloudgo.ActionListOpts{ListOpts: hcloudgo.ListOpts{Page: 1, PerPage: 50}}
	if d.KeyColumnQuals["id"] != nil {
		opts.ID = []int{int(d.KeyColumnQuals["id"].GetInt64Value())}
	}
	if d.KeyColumnQuals["status"] != nil {
		opts.Status = []hcloud.ActionStatus{hcloud.ActionStatus(d.KeyColumnQuals["status"].GetStringValue())}
	}

	for {
		items, resp, err := conn.Action.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("hcloud_action.listAction", "query_error", err)
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

func getAction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_action.getAction", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.Action
	var resp *hcloudgo.Response
	id := int(d.KeyColumnQuals["id"].GetInt64Value())
	item, resp, err = conn.Action.GetByID(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_action.getAction", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
