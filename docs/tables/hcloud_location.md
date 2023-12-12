---
title: "Steampipe Table: hcloud_location - Query Hetzner Cloud Locations using SQL"
description: "Allows users to query Locations in Hetzner Cloud, specifically the data center locations where resources can be deployed, providing insights into available regions and their respective details."
---

# Table: hcloud_location - Query Hetzner Cloud Locations using SQL

A Hetzner Cloud Location represents a data center where resources can be deployed. These locations are geographically distributed and are used to host the virtual machines (servers) and other resources provided by Hetzner Cloud. Each location has a unique name and provides different features and services.

## Table Usage Guide

The `hcloud_location` table provides insights into the data center locations within Hetzner Cloud. As a system administrator or a DevOps engineer, explore location-specific details through this table, including the name, description, and network zone of each location. Utilize it to uncover information about availability and features of each location, aiding in the strategic deployment of resources.

## Examples

### List all locations
Explore the various locations available in your Hetzner Cloud infrastructure, arranged in alphabetical order. This is particularly useful for understanding your geographical distribution and planning for data redundancy and latency optimization.

```sql+postgres
select
  *
from
  hcloud_location
order by
  name;
```

```sql+sqlite
select
  *
from
  hcloud_location
order by
  name;
```

### Get location by name
Discover the specifics of a particular location within the Hetzner Cloud service. This query can be used to gain insights into the settings and configuration of a chosen location, which is beneficial for planning and managing resources.

```sql+postgres
select
  *
from
  hcloud_location
where
  name = 'hel1';
```

```sql+sqlite
select
  *
from
  hcloud_location
where
  name = 'hel1';
```