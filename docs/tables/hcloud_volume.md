# Table: hcloud_volume

List volumes for the Hetzner Cloud account.

## Examples

### List all volumes

```sql
select
  id,
  name,
  created
from
  hcloud_volume
order by
  name
```

### List volumes with location data

```sql
select
  v.name as volume_name,
  l.name as location_name
from
  hcloud_volume as v,
  hcloud_location as l
where
  v.location_id = l.id
```

### Largest volumes

```sql
select
  id,
  name,
  size
from
  hcloud_volume
order by
  size desc
limit 5
```
