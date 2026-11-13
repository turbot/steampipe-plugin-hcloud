package hcloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type hcloudConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &hcloudConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hcloudConfig {
	if connection == nil || connection.GetConfig() == nil {
		return hcloudConfig{}
	}
	config, _ := connection.GetConfig().(hcloudConfig)
	return config
}
