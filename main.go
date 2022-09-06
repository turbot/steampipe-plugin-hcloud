package main

import (
	"github.com/turbot/steampipe-plugin-hcloud/hcloud"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hcloud.Plugin})
}
