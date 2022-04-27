package hcloud

import (
	"context"

	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableHcloudSSHKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hcloud_ssh_key",
		Description: "SSH keys in the Hetzner Cloud account.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"name", "fingerprint"}),
			Hydrate:    listSSHKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"fingerprint", "id", "name"}),
			Hydrate:    getSSHKey,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the SSH key."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the SSH key. Must be unique per project."},
			{Name: "fingerprint", Type: proto.ColumnType_STRING, Description: "Fingerprint of public key."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Point in time when the SSH key was created."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "User-defined labels (key-value pairs)."},
			{Name: "public_key", Type: proto.ColumnType_STRING, Description: "Public key."},
		},
	}
}

func listSSHKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_ssh_key.listSSHKey", "connection_error", err)
		return nil, err
	}

	opts := hcloudgo.SSHKeyListOpts{ListOpts: hcloudgo.ListOpts{Page: 1, PerPage: 50}}
	if d.KeyColumnQuals["name"] != nil {
		opts.Name = d.KeyColumnQuals["name"].GetStringValue()
	}
	if d.KeyColumnQuals["fingerprint"] != nil {
		opts.Fingerprint = d.KeyColumnQuals["fingerprint"].GetStringValue()
	}

	for {
		items, resp, err := conn.SSHKey.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("hcloud_ssh_key.listSSHKey", "query_error", err)
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

func getSSHKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_ssh_key.getSSHKey", "connection_error", err)
		return nil, err
	}
	var item *hcloudgo.SSHKey
	var resp *hcloudgo.Response
	id := int(d.KeyColumnQuals["id"].GetInt64Value())
	item, resp, err = conn.SSHKey.GetByID(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("hcloud_ssh_key.getSSHKey", "query_error", err, "resp", resp)
		return nil, err
	}
	return item, err
}
