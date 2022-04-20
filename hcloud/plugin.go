package hcloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-hcloud",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"hcloud_action":      tableHcloudAction(ctx),
			"hcloud_datacenter":  tableHcloudDataCenter(ctx),
			"hcloud_image":       tableHcloudImage(ctx),
			"hcloud_location":    tableHcloudLocation(ctx),
			"hcloud_network":     tableHcloudNetwork(ctx),
			"hcloud_server":      tableHcloudServer(ctx),
			"hcloud_server_type": tableHcloudServerType(ctx),
			"hcloud_ssh_key":     tableHcloudSSHKey(ctx),
			"hcloud_volume":      tableHcloudVolume(ctx),
		},
	}
	return p
}
