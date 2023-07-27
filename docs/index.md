---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/hcloud.svg"
brand_color: "#d50c2d"
display_name: "Hetzner Cloud"
short_name: "hcloud"
description: "Steampipe plugin to query servers, networks and more from Hetzner Cloud."
og_description: "Query Hetzner Cloud with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/hcloud-social-graphic.png"
---

# Hetzner Cloud + Steampipe

[Hetzner Cloud](https://www.hetzner.com/cloud) is a cloud hosting located in Germany.

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

## Multi-Account Connections

You may create multiple hcloud connections:

```hcl
connection "hcloud_dev" {
  plugin = "hcloud"
  token  = "RCgLZcAGBBGqkr8lTFXCfLhCjNNY6OJtzhg4ZsDRdMvmUwAeLssvLWyRRRdV8lkB"
}

connection "hcloud_qa" {
  plugin = "hcloud"
  token  = "RCgLZcAGBBGqkr8lTFXCfLhCjNNY6OJtzhg4ZsDRdMvmUwAeLzzrLWyRRRdV8lkB"
}

connection "hcloud_prod" {
  plugin = "hcloud"
  token  = "RCgLZcAGBBGqkr8lTFXCfLhCjNNY6OJtzhg4ZsDRdMvmUwAeLmmyLWyRRRdV8lkB"
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html). As such, you can use qualified table names to query a specific connection:

```sql
select * from hcloud_qa.hcloud_volume
```

You can create multi-account connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators). Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection.

```hcl
connection "hcloud_all" {
  plugin      = "hcloud"
  type        = "aggregator"
  connections = ["hcloud_dev", "hcloud_qa", "hcloud_prod"]
}
```

Querying tables from this connection will return results from the `hcloud_dev`, `hcloud_qa`, and `hcloud_prod` connections:

```sql
select * from hcloud_all.hcloud_volume
```

Alternatively, you can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/guides/search-path). It's a good idea to name your aggregator first alphabetically so that it is the first connection in the search path (i.e. `hcloud_all` comes before `hcloud_dev`):

```sql
select * from hcloud_volume
```

Steampipe supports the `*` wildcard in the connection names. For example, to aggregate all the hcloud plugin connections whose names begin with `hcloud_`:

```hcl
connection "hcloud_all" {
  type        = "aggregator"
  plugin      = "hcloud"
  connections = ["hcloud_*"]
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-hcloud
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
