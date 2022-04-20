package hcloud

import (
	"context"
	"os"

	hcloudgo "github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/pkg/errors"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*hcloudgo.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "hcloud"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*hcloudgo.Client), nil
	}

	// Default to the env var settings
	token := os.Getenv("HCLOUD_TOKEN")

	// Prefer config settings
	hcloudConfig := GetConfig(d.Connection)
	if &hcloudConfig != nil {
		if hcloudConfig.Token != nil {
			token = *hcloudConfig.Token
		}
	}

	// Error if the minimum config is not set
	if token == "" {
		return nil, errors.New("token must be configured")
	}

	conn := hcloudgo.NewClient(hcloudgo.WithToken(token))

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func isNotFoundError(err error) bool {
	return err.Error() == "[404] Not found"
}
