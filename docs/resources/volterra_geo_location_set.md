---

page_title: "Volterra: geo_location_set"

description: "The geo_location_set allows CRUD of Geo Location Set resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_geo_location_set
==================================

The Geo Location Set allows CRUD of Geo Location Set resource on Volterra SaaS

~> **Note:** Please refer to [Geo Location Set API docs](https://docs.cloud.f5.com/docs-v2/api/geo-location-set) to learn more

Example Usage
-------------

```hcl
resource "volterra_geo_location_set" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "custom_geo_location_selector global" must be set

  global = true
}

```

Argument Reference
------------------

### Metadata Argument Reference

`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).

`description` - (Optional) Human readable description for the object (`String`).

`disable` - (Optional) A value of true will administratively disable the object (`Bool`).

`labels` - (Optional) by selector expression (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

### Spec Argument Reference

###### One of the arguments from this list "custom_geo_location_selector, global" must be set

`custom_geo_location_selector` - (Optional) Select multiple geo locations. See [Location Choice Custom Geo Location Selector ](#location-choice-custom-geo-location-selector) below for details.

`global` - (Optional) Includes all geo locations (`Bool`).

### Location Choice Custom Geo Location Selector

Select multiple geo locations.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured geo_location_set.
