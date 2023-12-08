---
title: "Steampipe Table: hcloud_action - Query Hetzner Cloud Actions using SQL"
description: "Allows users to query Actions in Hetzner Cloud, specifically the status and progress of actions taken on resources, providing insights into resource management and potential issues."
---

# Table: hcloud_action - Query Hetzner Cloud Actions using SQL

Hetzner Cloud Actions are tasks performed on resources within the Hetzner Cloud platform. These actions include tasks such as creating, updating, or deleting resources like servers, volumes, and networks. Each action has a status and progress which can be tracked to understand the outcome of the task.

## Table Usage Guide

The `hcloud_action` table provides insights into actions taken within Hetzner Cloud. As a system administrator, explore action-specific details through this table, including statuses, progress, and associated metadata. Utilize it to monitor the progress and outcomes of tasks, such as resource creation, updates, and deletions, and to identify any potential issues that may arise during these processes.

## Examples

### List create_server actions
Explore instances where server creation actions have been initiated to gain insights into system activity and manage resources effectively. This can be particularly useful for tracking resource allocation and identifying areas for optimization.

```sql+postgres
select
  *
from
  hcloud_action
where
  command = 'create_server';
```

```sql+sqlite
select
  *
from
  hcloud_action
where
  command = 'create_server';
```

### Get a specific action
Analyze the settings to understand the specifics of a particular action in your Hetzner Cloud environment. This can be useful in identifying changes or actions that could impact system performance or security.

```sql+postgres
select
  *
from
  hcloud_action
where
  id = 271232672;
```

```sql+sqlite
select
  *
from
  hcloud_action
where
  id = 271232672;
```

### Actions in error
Discover instances where actions have failed to execute correctly, allowing for the identification and resolution of potential issues within your system. This is crucial for maintaining optimal system performance and stability.

```sql+postgres
select
  *
from
  hcloud_action
where
  status = 'error';
```

```sql+sqlite
select
  *
from
  hcloud_action
where
  status = 'error';
```

### All actions related to a given server
Identify instances where specific server-related actions have taken place. This can help in monitoring server activity, providing insights into system usage and behavior patterns.

```sql+postgres
select
  a.*
from
  hcloud_action as a,
  jsonb_array_elements(a.resources) as r
where
  r->>'Type' = 'server'
  and (r->'ID')::int = 14462596;
```

```sql+sqlite
select
  a.*
from
  hcloud_action as a,
  json_each(a.resources) as r
where
  json_extract(r.value, '$.Type') = 'server'
  and json_extract(r.value, '$.ID') = 14462596;
```