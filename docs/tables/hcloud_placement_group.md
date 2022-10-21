# Table: hcloud_placement_group

List placement groups for the Hetzner Cloud account.

## Examples

### List placement groups

```sql
select
  id,
  name,
  type,
  servers
from
  hcloud_placement_group
order by
  id;
```

### List placement groups with the label env=prod

```sql
select
  id,
  name,
  labels,
  servers
from
  hcloud_placement_group
where
  labels ->> 'env' = 'prod';
```

### Get server details for servers in placement groups

```sql
with placement_groups as (
  select
    p.name,
    p.type,
    server_id
  from
    hcloud_placement_group as p,
    jsonb_array_elements(servers) as server_id
)
select
  p.name as placement_group_name,
  p.type as placement_group_type,
  s.name as server_name,
  s.status as server_status,
  s.image_id as server_image_id,
  s.id as server_id
from
  placement_groups as p,
  hcloud_server as s
where
  s.id::text = p.server_id::text;
```
