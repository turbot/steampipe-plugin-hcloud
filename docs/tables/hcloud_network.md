# Table: hcloud_network

List networks for the Hetzner Cloud account.

## Examples

### List all networks

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
