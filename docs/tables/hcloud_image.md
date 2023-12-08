---
title: "Steampipe Table: hcloud_image - Query Hetzner Cloud Images using SQL"
description: "Allows users to query Images in Hetzner Cloud, specifically the image ID, type, status, and other related information, providing insights into the various system images and snapshots available in the cloud environment."
---

# Table: hcloud_image - Query Hetzner Cloud Images using SQL

Hetzner Cloud Images are system images and snapshots that can be used to create new servers. These images include pre-configured operating systems, applications, and configurations that are stored in Hetzner Cloud. They can be public (provided by Hetzner), private (custom images created by users), or snapshots created from existing servers.

## Table Usage Guide

The `hcloud_image` table provides insights into images within Hetzner Cloud. As a cloud administrator, explore image-specific details through this table, including image type, status, and associated metadata. Utilize it to uncover information about images, such as their description, source server, and the last time they were updated.

## Examples

### List all images
Explore all available images in your system to gain a better understanding of your resources. This can be useful for managing and organizing your digital assets.

```sql+postgres
select
  id,
  name,
  description
from
  hcloud_image
order by
  id;
```

```sql+sqlite
select
  id,
  name,
  description
from
  hcloud_image
order by
  id;
```

### Find all deprecated images
Uncover the details of outdated images within your system to better manage and update your resources for optimal performance.

```sql+postgres
select
  id,
  name,
  description,
  deprecated
from
  hcloud_image
where
  deprecated is not null;
```

```sql+sqlite
select
  id,
  name,
  description,
  deprecated
from
  hcloud_image
where
  deprecated is not null;
```