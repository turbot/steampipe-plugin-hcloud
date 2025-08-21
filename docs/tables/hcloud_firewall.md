---
title: "Steampipe Table: hcloud_firewall - Query Hetzner Cloud Firewalls using SQL"
description: "Allows users to query Firewalls in Hetzner Cloud, specifically the firewall ID, name, and other related information."
---

# Table: hcloud_firewall - Query Hetzner Cloud Firewalls using SQL

Hetzner Cloud provides firewall management for securing access to cloud resources. Firewalls consist of configurable rules and can be applied to servers either directly or via label selectors. They help enforce network-level security policies to limit unauthorized access.

## Table Usage Guide

The `hcloud_firewall` table allows you to retrieve and explore information about firewalls configured in your Hetzner Cloud environment. As a system administrator or DevOps engineer, you can use this table to audit firewall usage, analyze rule configurations, and identify servers without firewall protection.

## Examples

### List all firewalls
Explore all available firewalls to understand what exists in your cloud account and how each is identified.

```sql+postgres
select
  *
from
  hcloud_firewall
order by
  name;
```

```sql+sqlite
select
  *
from
  hcloud_firewall
order by
  name;
```

### Get firewall by name
Retrieve the specific configuration of a named firewall to inspect its rules and applied targets.

```sql+postgres
select
  *
from
  hcloud_firewall
where
  name = 'firewall1;
```

```sql+sqlite
select
  *
from
  hcloud_firewall
where
  name = 'firewall1;
```

### List firewalls not applied to any resource
Find firewalls that are currently not used, so you can clean them up or reassign them to the appropriate servers.

```sql+postgres
select
  id,
  name,
  applied_to
from
  hcloud_firewall
where
  applied_to is null;
```

```sql+sqlite
select
  id,
  name,
  applied_to
from
  hcloud_firewall
where
  applied_to is null;
```