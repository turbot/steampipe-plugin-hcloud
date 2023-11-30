---
title: "Steampipe Table: hcloud_placement_group - Query Hetzner Cloud Placement Groups using SQL"
description: "Allows users to query Placement Groups in Hetzner Cloud, providing details about placement groups which are used to determine the physical location of servers."
---

# Table: hcloud_placement_group - Query Hetzner Cloud Placement Groups using SQL

A Placement Group in Hetzner Cloud is a resource used to determine the physical location of servers. It provides an overview of the server's placement group, including its type and servers. The placement group's type can be either "spread" or "cluster", and it is used to control the distribution of servers within the same data center.

## Table Usage Guide

The `hcloud_placement_group` table provides insights into Placement Groups within Hetzner Cloud. As a Cloud Engineer, you can explore placement group-specific details through this table, including their type and associated servers. Utilize it to manage the distribution of servers within the same data center, ensuring optimal performance and availability.

## Examples

### List placement groups
Explore the organization of your server infrastructure by identifying the groups in which your servers are placed. This can help to understand the distribution and categorization of your servers, thereby aiding in efficient resource management.

```sql
select
  id,
  name,
  type,
  servers
from
  hcloud_placement_group
order by
  id;
```

### List placement groups with the label env=prod
Discover the segments that have been labeled as 'production environment' within your placement groups. This is useful for understanding how resources are allocated and managed within your production environment.

```sql
select
  id,
  name,
  labels,
  servers
from
  hcloud_placement_group
where
  labels ->> 'env' = 'prod';
```

### Get server details for servers in placement groups
This query is useful for identifying the status and other details of servers located within specific placement groups. This can help in managing and monitoring server performance and resource allocation more effectively.

```sql
with placement_groups as (
  select
    p.name,
    p.type,
    server_id
  from
    hcloud_placement_group as p,
    jsonb_array_elements(servers) as server_id
)
select
  p.name as placement_group_name,
  p.type as placement_group_type,
  s.name as server_name,
  s.status as server_status,
  s.image_id as server_image_id,
  s.id as server_id
from
  placement_groups as p,
  hcloud_server as s
where
  s.id::text = p.server_id::text;
```