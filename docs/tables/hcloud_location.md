# Table: hcloud_location

List locations for the Hetzner Cloud account.

## Examples

### List all locations

```sql
select
  *
from
  hcloud_location
order by
  name
```

### Get location by name

```sql
select
  *
from
  hcloud_location
where
  name = 'hel1'
```
