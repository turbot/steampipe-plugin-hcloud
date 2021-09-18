# Table: hcloud_action

Query actions from your Hetzner Cloud account.

## Examples

### List create_server actions

```sql
select
  *
from
  hcloud_action
where
  command = 'create_server'
```

### Get a specific action

```sql
select
  *
from
  hcloud_action
where
  id = 271232672
```

### Actions in error

```sql
select
  *
from
  hcloud_action
where
  status = 'error'
```

### All actions related to a given server

```sql
select
  a.*
from
  hcloud_action as a,
  jsonb_array_elements(a.resources) as r
where
  r->>'Type' = 'server'
  and (r->'ID')::int = 14462596
```
