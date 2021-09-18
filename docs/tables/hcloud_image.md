# Table: hcloud_image

List images for the Hetzner Cloud account.

## Examples

### List all images

```sql
select
  id,
  name,
  description
from
  hcloud_image
order by
  id
```

### Find all deprecated images

```sql
select
  id,
  name,
  description,
  deprecated
from
  hcloud_image
where
  deprecated is not null
```
