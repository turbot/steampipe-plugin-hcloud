## v0.5.0 [2023-04-07]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#15](https://github.com/turbot/steampipe-plugin-hcloud/pull/15))

## v0.4.0 [2022-10-20]

- New tables added
  - [hcloud_placement_group](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_placement_group) ([#13](https://github.com/turbot/steampipe-plugin-hcloud/pull/13)) (Thanks to [@hannesfant](https://github.com/hannesfant) for the new table!)

## v0.3.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#11](https://github.com/turbot/steampipe-plugin-hcloud/pull/11))
- Recompiled plugin with Go version `1.19`. ([#11](https://github.com/turbot/steampipe-plugin-hcloud/pull/11))

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#7](https://github.com/turbot/steampipe-plugin-hcloud/pull/7))

## v0.2.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#5](https://github.com/turbot/steampipe-plugin-hcloud/pull/5))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#4](https://github.com/turbot/steampipe-plugin-hcloud/pull/4))

## v0.1.0 [2021-12-16]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#2](https://github.com/turbot/steampipe-plugin-hcloud/pull/2))
- Recompiled plugin with Go version 1.17 ([#2](https://github.com/turbot/steampipe-plugin-hcloud/pull/2))

## v0.0.2 [2021-10-06]

_Bug fixes_

- Fix brand color in Steampipe Hub icons.

## v0.0.1 [2021-10-06]

_What's new?_

- New tables added
  - [hcloud_action](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_action)
  - [hcloud_datacenter](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_datacenter)
  - [hcloud_image](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_image)
  - [hcloud_location](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_location)
  - [hcloud_network](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_network)
  - [hcloud_server](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_server)
  - [hcloud_server_type](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_server_type)
  - [hcloud_ssh_key](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_ssh_key)
  - [hcloud_volume](https://hub.steampipe.io/plugins/turbot/hcloud/tables/hcloud_volume)
