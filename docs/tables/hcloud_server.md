---
title: "Steampipe Table: hcloud_server - Query Hetzner Cloud Servers using SQL"
description: "Allows users to query Hetzner Cloud Servers, specifically the servers' details like ID, name, status, data center location, and more."
---

# Table: hcloud_server - Query Hetzner Cloud Servers using SQL

Hetzner Cloud is a cloud infrastructure solutions provider that offers scalable and cost-effective cloud servers. These servers are located in multiple data centers across the globe, providing high-performance and reliable cloud services. Hetzner Cloud servers are ideal for a wide range of applications, including web hosting, development environments, and high-traffic websites.

## Table Usage Guide

The `hcloud_server` table offers insights into the servers hosted on Hetzner Cloud. As a system administrator or a DevOps engineer, you can explore server-specific details through this table, including server status, location, type, and associated metadata. Use it to monitor server availability, manage server resources, and track the performance of your cloud infrastructure.

## Examples

### List all servers
Explore a comprehensive list of all servers in a structured order based on their names. This is useful for maintaining an organized inventory of servers and simplifying server management tasks.

```sql+postgres
select
  *
from
  hcloud_server
order by
  name;
```

```sql+sqlite
select
  *
from
  hcloud_server
order by
  name;
```

### Get server by name
Discover the segments that pertain to a specific server by its name. This is particularly useful in a scenario where you want to understand the configuration and details of a specific server within your infrastructure.

```sql+postgres
select
  *
from
  hcloud_server
where
  name = 'ubuntu-2gb-hel1-1';
```

```sql+sqlite
select
  *
from
  hcloud_server
where
  name = 'ubuntu-2gb-hel1-1';
```

### List servers with IPs
Explore which servers have assigned IP addresses, both private and public. This could be useful for network management and troubleshooting connectivity issues.

```sql+postgres
select
  name,
  priv ->> 'IP' as private_ipv4,
  public_net -> 'IPv4' ->> 'IP' as public_ipv4,
  public_net -> 'IPv6' ->> 'IP' as public_ipv6
from
  hcloud_server as s,
  jsonb_array_elements(s.private_net) as priv
order by
  name;
```

```sql+sqlite
select
  s.name,
  json_extract(priv.value, '$.IP') as private_ipv4,
  json_extract(pub.value, '$.IPv4.IP') as public_ipv4,
  json_extract(pub.value, '$.IPv6.IP') as public_ipv6
from
  hcloud_server as s,
  json_each(s.private_net) as priv,
  json_each(s.public_net) as pub
order by
  s.name;
```

### List servers with server type information
Explore which servers are associated with specific server types, including details like the number of cores and CPU type. This can be beneficial in managing resources and understanding hardware specifications across your server infrastructure.

```sql+postgres
select
  s.name as server_name,
  st.name as server_type_name,
  st.cores,
  st.cpu_type
from
  hcloud_server as s,
  hcloud_server_type as st
where
  s.server_type_id = st.id;;
```

```sql+sqlite
select
  s.name as server_name,
  st.name as server_type_name,
  st.cores,
  st.cpu_type
from
  hcloud_server as s
join
  hcloud_server_type as st
on
  s.server_type_id = st.id;
```

### List all volumes for all servers
Determine the areas in which server storage is being utilized by listing all volumes across all servers. This can help in assessing storage distribution and planning for capacity management.

```sql+postgres
select
  s.name as server_name,
  v.name as volume_name,
  v.size
from
  hcloud_server as s,
  jsonb_array_elements(s.volume_ids) as sv,
  hcloud_volume as v
where
  sv::int = v.id;
```

```sql+sqlite
select
  s.name as server_name,
  v.name as volume_name,
  v.size
from
  hcloud_server as s,
  json_each(s.volume_ids) as sv,
  hcloud_volume as v
where
  CAST(sv.value AS INTEGER) = v.id;
```