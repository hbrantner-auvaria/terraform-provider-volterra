---

page_title: "Volterra: api_discovery"

description: "The api_discovery allows CRUD of Api Discovery resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_api_discovery
===============================

The Api Discovery allows CRUD of Api Discovery resource on Volterra SaaS

~> **Note:** Please refer to [Api Discovery API docs](https://docs.cloud.f5.com/docs-v2/api/api-discovery) to learn more

Example Usage
-------------

```hcl
resource "volterra_api_discovery" "example" {
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

`custom_auth_types` - (Optional) Select your custom authentication types to be detected in the API discovery. See [Custom Auth Types ](#custom-auth-types) below for details.

`user_defined_api_discovery_policy` - (Optional) Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.. See [User Defined Api Discovery Policy ](#user-defined-api-discovery-policy) below for details.

### Custom Auth Types

Select your custom authentication types to be detected in the API discovery.

`parameter_name` - (Required) The authentication parameter name. (`String`).

`parameter_type` - (Required) x-displayName: "Parameter Type" (`String`).

### User Defined Api Discovery Policy

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field..

###### One of the arguments from this list "exclusive, inclusive" must be set

`exclusive` - (Optional) Any traffic that does not match the specified rules will be ignored by default.. See [Default Behavior Choice Exclusive ](#default-behavior-choice-exclusive) below for details.

`inclusive` - (Optional) Any traffic that does not match the specified rules will be discovered by default. (`Bool`).

`discovery_rules` - (Optional) Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom; unmatched endpoints follow the default action.. See [User Defined Api Discovery Policy Discovery Rules ](#user-defined-api-discovery-policy-discovery-rules) below for details.

### Action Choice Archive

Endpoints are removed from active API Discovery and stored in the Archive under Non-API Rules for reference and auditing..

### Action Choice Ignore

Excluded endpoints are completely ignored by API Discovery and will not appear in discovery results, inventories, or insights..

### Criteria Http Header Criteria

x-displayName: "HTTP Header".

`field_name` - (Required) x-inlineHint: "e.g. content-type, user-id" (`String`).

`location` - (Optional) x-displayName: "Location" (`String`).

`match_type` - (Optional) x-displayName: "Match Type" (`String`).

`value` - (Required) x-displayName: "Value" (`String`).

### Default Behavior Choice Exclusive

Any traffic that does not match the specified rules will be ignored by default..

###### One of the arguments from this list "archive, ignore" must be set

`archive` - (Optional) Endpoints are removed from active API Discovery and stored in the Archive under Non-API Rules for reference and auditing. (`Bool`).

`ignore` - (Optional) Excluded endpoints are completely ignored by API Discovery and will not appear in discovery results, inventories, or insights. (`Bool`).

### Default Behavior Choice Inclusive

Any traffic that does not match the specified rules will be discovered by default..

### Discovery Rules Metadata

Standard object metadata including name, labels, and description.

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Discovery Rules Rule Properties

Configuration for rule type and matching criteria.

###### One of the arguments from this list "http_header_criteria, pattern" must be set

`http_header_criteria` - (Optional) x-displayName: "HTTP Header". See [Criteria Http Header Criteria ](#criteria-http-header-criteria) below for details.

`pattern` - (Optional) Patterns are matched against the request path to identify endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern are affected by the rule. (`String`).

###### One of the arguments from this list "exclusion, inclusion" must be set

`exclusion` - (Optional) Matching endpoints will be excluded from API discovery. See [Rule Type Choice Exclusion ](#rule-type-choice-exclusion) below for details.

`inclusion` - (Optional) Matching endpoints will be included in API discovery (`Bool`).

### Rule Type Choice Exclusion

Matching endpoints will be excluded from API discovery.

###### One of the arguments from this list "archive, ignore" must be set

`archive` - (Optional) Endpoints are removed from active API Discovery and stored in the Archive under Non-API Rules for reference and auditing. (`Bool`).

`ignore` - (Optional) Excluded endpoints are completely ignored by API Discovery and will not appear in discovery results, inventories, or insights. (`Bool`).

### Rule Type Choice Inclusion

Matching endpoints will be included in API discovery.

### User Defined Api Discovery Policy Discovery Rules

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom; unmatched endpoints follow the default action..

`labels` - (Optional) Map of string keys and values that can be used to organize and categorize the rule (`String`).

`metadata` - (Optional) Standard object metadata including name, labels, and description. See [Discovery Rules Metadata ](#discovery-rules-metadata) below for details.

`rule_properties` - (Optional) Configuration for rule type and matching criteria. See [Discovery Rules Rule Properties ](#discovery-rules-rule-properties) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured api_discovery.
