---
title: "Steampipe Table: hcloud_server_type - Query Hetzner Cloud Server Types using SQL"
description: "Allows users to query Server Types in Hetzner Cloud, specifically the details of each server type, providing insights into available server configurations and their characteristics."
---

# Table: hcloud_server_type - Query Hetzner Cloud Server Types using SQL

Hetzner Cloud Server Types are predefined configurations of servers available for use in Hetzner Cloud. Each server type comes with a specific set of resources, such as CPU cores, memory, and disk space. These server types are designed to accommodate a wide range of computing needs, from small-scale applications to large-scale enterprise projects.

## Table Usage Guide

The `hcloud_server_type` table provides insights into the server types within Hetzner Cloud. As a system administrator or a DevOps engineer, explore server type-specific details through this table, including CPU cores, memory, disk size, and more. Utilize it to uncover information about server types, such as those with specific resource configurations, to help in selecting the most suitable server type for your application or project.

## Examples

### List all server types
Discover the different server types available in your infrastructure, which can help you understand your resource distribution and inform decisions about capacity planning or resource allocation.

```sql+postgres
select
  name,
  description
from
  hcloud_server_type
order by
  name;
```

```sql+sqlite
select
  name,
  description
from
  hcloud_server_type
order by
  name;
```

### Get all server types with pricing
Explore various server types alongside their pricing details to make informed decisions about cost management and resource allocation. This can help in optimizing your spending by choosing the most cost-effective server type for your needs.

```sql+postgres
select
  st.name,
  st.description,
  (p -> 'Hourly' ->> 'Net')::float as hourly_net,
  (p -> 'Monthly' ->> 'Net')::float as monthly_net
from
  hcloud_server_type as st,
  jsonb_array_elements(st.prices) as p
order by
  hourly_net;
```

```sql+sqlite
select
  st.name,
  st.description,
  cast(json_extract(p.value, '$.Hourly.Net') as float) as hourly_net,
  cast(json_extract(p.value, '$.Monthly.Net') as float) as monthly_net
from
  hcloud_server_type as st,
  json_each(st.prices) as p
order by
  hourly_net;
```