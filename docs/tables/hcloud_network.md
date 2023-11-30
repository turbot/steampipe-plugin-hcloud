---
title: "Steampipe Table: hcloud_network - Query Hetzner Cloud Networks using SQL"
description: "Allows users to query Networks in Hetzner Cloud, providing insights into the network infrastructure and configurations."
---

# Table: hcloud_network - Query Hetzner Cloud Networks using SQL

Hetzner Cloud Networks is a service within Hetzner Cloud that allows users to manage and configure their network infrastructure. It provides a centralized way to set up and manage networks for various resources, such as virtual machines and databases. Hetzner Cloud Networks helps users to maintain the health and performance of their network resources, and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `hcloud_network` table provides insights into networks within Hetzner Cloud. As a network administrator, explore network-specific details through this table, including network configurations, associated subnets, and IP ranges. Utilize it to uncover information about networks, such as their IP ranges, associated resources, and the overall structure of your network infrastructure.

## Examples

### List all networks
Explore all the networks available in your environment, ordered by their ID. This can assist in identifying and understanding the various networks in use, which is crucial for effective network management and security.

```sql
select
  id,
  name,
  description
from
  hcloud_network
order by
  id
```

### List all networks with the label env=prod
Explore which networks are labeled as 'production environment'. This is beneficial for managing and distinguishing your production systems from development or testing systems, ensuring the right resources are allocated and maintained.

```sql
select
  id,
  name,
  labels
from
  hcloud_network
where
  labels->>'env' = 'prod'
```

### List all servers in the network
Explore which servers are part of your network. This can help you manage and monitor your network infrastructure more effectively.

```sql
select
  n.name as network_name,
  s.name as server_name,
  s.server_type
from
  hcloud_network as n,
  jsonb_array_elements(n.server_ids) as sid,
  hcloud_server as s
where
  s.id = sid::int
```