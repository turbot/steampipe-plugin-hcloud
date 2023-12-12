package hcloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type hcloudConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &hcloudConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hcloudConfig {
	if connection == nil || connection.Config == nil {
		return hcloudConfig{}
	}
	config, _ := connection.Config.(hcloudConfig)
	return config
}
