# Table: hcloud_server_type

List server types for the Hetzner Cloud account.

## Examples

### List all server types

```sql
select
  name,
  description
from
  hcloud_server_type
order by
  name
```

### Get all server types with pricing

```sql
select
  st.name,
  st.description,
  (p -> 'Hourly' ->> 'Net')::float as hourly_net,
  (p -> 'Monthly' ->> 'Net')::float as monthly_net
from
  hcloud_server_type as st,
  jsonb_array_elements(st.prices) as p
order by
  hourly_net
```
