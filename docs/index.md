---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/hcloud.svg"
brand_color: "#00b050"
display_name: "Hetzner Cloud"
short_name: "hcloud"
description: "Steampipe plugin to query servers, networks and more from Hetzner Cloud."
og_description: "Query Hetzner Cloud with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/hcloud-social-graphic.png"
---

# Hetzner Cloud + Steampipe

[Hetzner Cloud](https://hcloud.com) is a cloud hosting located in Germany.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List servers in your Hetzner Cloud account:

```sql
select
  id,
  name,
  created
from
  hcloud_server
```

```
+----------+-------------------+---------------------+
| id       | name              | created             |
+----------+-------------------+---------------------+
| 14462596 | ubuntu-2gb-hel1-1 | 2021-09-18 01:53:27 |
+----------+-------------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/hcloud/tables)**

## Get started

### Install

Download and install the latest Hetzner Cloud plugin:

```bash
steampipe plugin install hcloud
```

### Configuration

Installing the latest hcloud plugin will create a config file (`~/.steampipe/config/hcloud.spc`) with a single connection named `hcloud`:

```hcl
connection "hcloud" {
  plugin = "hcloud"
  token  = "RCgLZcAGBBGqkr8lTFXCfLhCjLMM6OJtzhg4ZsDRdMvmUwAeLssvLWyCCTdV8lkB"
}
```

- `token` - API token from Hetzner Cloud.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-hcloud
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
