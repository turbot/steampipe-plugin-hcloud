---
title: "Steampipe Table: hcloud_volume - Query Hetzner Cloud Volumes using SQL"
description: "Allows users to query Volumes in Hetzner Cloud, specifically the details of each volume including size, location, and attached server."
---

# Table: hcloud_volume - Query Hetzner Cloud Volumes using SQL

A Volume in Hetzner Cloud is a block storage device that can be attached to a server in the same location. Volumes are independent resources that persist beyond the life of any server. They can be dynamically resized and moved between servers.

## Table Usage Guide

The `hcloud_volume` table provides insights into Volumes within Hetzner Cloud. As a Cloud engineer, explore volume-specific details through this table, including size, location, and attached server. Utilize it to uncover information about volumes, such as those with larger sizes, the servers they are attached to, and their geographical locations.

## Examples

### List all volumes
Explore all the storage volumes in your infrastructure, ordered by their names, to gain insights into their creation timeline and manage resources effectively. This can be useful in assessing the utilization of storage resources and planning for future capacity needs.

```sql
select
  id,
  name,
  created
from
  hcloud_volume
order by
  name
```

### List volumes with location data
Explore which volumes are associated with specific locations in your Hetzner Cloud environment. This can be useful in managing resources and understanding data distribution across different geographical areas.

```sql
select
  v.name as volume_name,
  l.name as location_name
from
  hcloud_volume as v,
  hcloud_location as l
where
  v.location_id = l.id
```

### Largest volumes
Discover the top five largest storage volumes in your cloud environment to better manage resource allocation and utilization. This can help in identifying potential areas for cost savings or performance improvements.

```sql
select
  id,
  name,
  size
from
  hcloud_volume
order by
  size desc
limit 5
```