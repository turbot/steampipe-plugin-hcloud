---
title: "Steampipe Table: hcloud_datacenter - Query Hetzner Cloud Datacenters using SQL"
description: "Allows users to query Datacenters in Hetzner Cloud, specifically the location, server types, and availability, providing insights into datacenter distribution and potential capacity."
---

# Table: hcloud_datacenter - Query Hetzner Cloud Datacenters using SQL

Hetzner Cloud Datacenters are physical locations around the world where Hetzner Cloud servers reside. They are designed to be highly reliable, secure, and efficient. These datacenters contain the server types and their availability, which are crucial for planning and managing cloud resources.

## Table Usage Guide

The `hcloud_datacenter` table provides insights into datacenters within Hetzner Cloud. As a system administrator or DevOps engineer, explore datacenter-specific details through this table, including location, server types, and availability. Utilize it to uncover information about datacenters, such as their geographical distribution, the types of servers available, and their current availability status.

## Examples

### List all data centers
Explore the various data centers available, ordered by their names, to better manage resources and optimize performance. This can be particularly useful when planning the deployment of new services or assessing existing infrastructure.

```sql
select
  *
from
  hcloud_datacenter
order by
  name
```

### Get data center by name
Explore which specific data center is associated with a certain name, allowing you to quickly identify and access relevant data center information for further analysis or management tasks. This is particularly useful in environments with multiple data centers, where locating specific ones by name can streamline administrative tasks.

```sql
select
  *
from
  hcloud_datacenter
where
  name = 'hel1'
```

### Get all available server types for all data centers
Explore the variety of server types across all data centers. This is useful for understanding the diverse server options available for different data center deployments, aiding in strategic decision-making and resource allocation.

```sql
select
  dc.name,
  st.*
from
  hcloud_datacenter as dc,
  jsonb_array_elements(dc.server_types_available) as sta,
  hcloud_server_type as st
where
  sta::int = st.id
```