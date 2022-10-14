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
  labels->>'env' = 'prod'
```
