---

page_title: "Volterra: discovery_cloud"

description: "The discovery_cloud allows CRUD of Discovery Cloud resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_discovery_cloud
=================================

The Discovery Cloud allows CRUD of Discovery Cloud resource on Volterra SaaS

~> **Note:** Please refer to [Discovery Cloud API docs](https://docs.cloud.f5.com/docs-v2/api/discovery-cloud) to learn more

Example Usage
-------------

```hcl
resource "volterra_discovery_cloud" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

###### One of the arguments from this list "aws_discovery" can be set

`aws_discovery` - (Optional) Discovery of AWS services. See [Discovery Choice Aws Discovery ](#discovery-choice-aws-discovery) below for details.

### Aws Discovery Vpc Configs

Configuration of Virtual Private Clouds (VPCs) to discover from.

`region` - (Required) Region of VPC. See [ref](#ref) below for details.

`service_types` - (Required) Types of services to discover. See [Vpc Configs Service Types ](#vpc-configs-service-types) below for details.

`vpc_id` - (Required) ID of VPC to discover from (`String`).(Deprecated)

`vpc_name` - (Required) Name of VPC to discover from (`String`).

`where` - (Optional) Where discovered service endpoints will be hosted. See [Vpc Configs Where ](#vpc-configs-where) below for details.

### Discovery Choice Aws Discovery

Discovery of AWS services.

###### One of the arguments from this list "cloud_user_account_ref" can be set

`cloud_user_account_ref` - (Required) Cloud user account to discover services from. See [ref](#ref) below for details.

`vpc_configs` - (Required) Configuration of Virtual Private Clouds (VPCs) to discover from. See [Aws Discovery Vpc Configs ](#aws-discovery-vpc-configs) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Vpc Configs Service Types

Types of services to discover.

`selector` - (Optional) Service tag selector for selecting services to discovery. Uses kubernetes style label expression for selections. (`String`).

`type` - (Required) x-required (`String`).

### Vpc Configs Where

Where discovered service endpoints will be hosted.

###### One of the arguments from this list "site, virtual_site" can be set

`site` - (Optional) Direct reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [ref](#ref) below for details.

Attribute Reference
-------------------

*   `id` - This is the id of the configured discovery_cloud.
