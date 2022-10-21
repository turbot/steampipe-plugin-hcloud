# Table: hcloud_placement_group

List placement groups for the Hetzner Cloud account.

## Examples

### List all placement groups

```sql
select
  id,
  name,
  type
from
  hcloud_placement_group
order by
  id
```

### List all placement groups with the label env=prod

```sql
select
  id,
  name,
  labels
from
  hcloud_placement_group
where
  labels ->> 'env' = 'prod'
```

### Get the names of all servers within the given placement group
```sql
select hcloud_placement_group.name,servers,hcloud_server.name from hcloud_placement_group
CROSS JOIN LATERAL JSONB_ARRAY_ELEMENTS(hcloud_placement_group.servers) AS e(srv)
inner join hcloud_server on hcloud_server.id  = (e.srv)::text::int
```