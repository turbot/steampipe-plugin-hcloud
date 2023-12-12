---
title: "Steampipe Table: hcloud_ssh_key - Query Hetzner Cloud SSH Keys using SQL"
description: "Allows users to query SSH Keys in Hetzner Cloud, specifically the details about the keys including their names, fingerprints, and associated metadata."
---

# Table: hcloud_ssh_key - Query Hetzner Cloud SSH Keys using SQL

Hetzner Cloud SSH Key is a resource in Hetzner Cloud that allows users to authenticate without using a password. SSH keys provide a more secure way of logging into a virtual private server with SSH than using a password alone. It is a highly recommended way of managing server authentication.

## Table Usage Guide

The `hcloud_ssh_key` table provides insights into SSH keys within Hetzner Cloud. As a system administrator, explore key-specific details through this table, including names, fingerprints, and associated metadata. Utilize it to uncover information about keys, such as their fingerprints, the creation and modification dates, and the verification of key details.

## Examples

### List all SSH keys
Explore all SSH keys within your system to gain a comprehensive overview of your security mechanisms. This can be particularly useful when auditing system security or troubleshooting access issues.

```sql+postgres
select
  *
from
  hcloud_ssh_key;
```

```sql+sqlite
select
  *
from
  hcloud_ssh_key;
```

### Oldest SSH keys
Explore which SSH keys have been in use the longest to assess potential security vulnerabilities and update protocols as necessary.

```sql+postgres
select
  name,
  fingerprint,
  created
from
  hcloud_ssh_key
order by
  created
limit 5;
```

```sql+sqlite
select
  name,
  fingerprint,
  created
from
  hcloud_ssh_key
order by
  created
limit 5;
```