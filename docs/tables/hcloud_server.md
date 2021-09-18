# Table: hcloud_server

List servers for the Hetzner Cloud account.

## Examples

### List all servers

```sql
select
  *
from
  hcloud_server
order by
  name
```

### Get server by name

```sql
select
  *
from
  hcloud_server
where
  name = 'ubuntu-2gb-hel1-1'
```

### List servers with IPs

```sql
select
  name,
  priv ->> 'IP' as private_ipv4,
  public_net -> 'IPv4' ->> 'IP' as public_ipv4,
  public_net -> 'IPv6' ->> 'IP' as public_ipv6
from
  hcloud_server as s,
  jsonb_array_elements(s.private_net) as priv
order by
  name
```

### List servers with server type information

```sql
select
  s.name as server_name,
  st.name as server_type_name,
  st.cores,
  st.cpu_type
from
  hcloud_server as s,
  hcloud_server_type as st
where
  s.server_type_id = st.id
```

### List all volumes for all servers

```sql
select
  s.name as server_name,
  v.name as volume_name,
  v.size
from
  hcloud_server as s,
  jsonb_array_elements(s.volume_ids) as sv,
  hcloud_volume as v
where
  sv::int = v.id
```
