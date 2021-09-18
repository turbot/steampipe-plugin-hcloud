# Table: hcloud_datacenter

List data centers for the Hetzner Cloud account.

## Examples

### List all data centers

```sql
select
  *
from
  hcloud_datacenter
order by
  name
```

### Get data center by name

```sql
select
  *
from
  hcloud_datacenter
where
  name = 'hel1'
```

### Get all available server types for all data centers

```sql
select
  dc.name,
  st.*
from
  hcloud_datacenter as dc,
  jsonb_array_elements(dc.server_types_available) as sta,
  hcloud_server_type as st
where
  sta::int = st.id
```
