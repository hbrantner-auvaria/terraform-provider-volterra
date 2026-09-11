//
// Copyright (c) 2026 F5 Inc. All rights reserved.
//

package driftdetection

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_alert_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/alert_policy"
)

func FlattenAPReceivers(x []*ves_io_schema.ObjectRefType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		if v == nil {
			continue
		}
		t := map[string]interface{}{
			"kind":      v.GetKind(),
			"name":      v.GetName(),
			"namespace": v.GetNamespace(),
			"tenant":    v.GetTenant(),
		}
		rslt = append(rslt, t)
	}
	return rslt
}

func FlattenAPLabelMatcher(x *ves_io_schema_alert_policy.LabelMatcher) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		v := map[string]interface{}{
			"exact_match": x.GetExactMatch(),
			"regex_match": x.GetRegexMatch(),
		}
		rslt = append(rslt, v)
	}
	return rslt
}

func FlattenAPAlertLabel(x map[string]*ves_io_schema_alert_policy.LabelMatcher) []interface{} {
	rslt := make([]interface{}, 0)
	for k, v := range x {
		entry := map[string]interface{}{
			"name":  k,
			"value": FlattenAPLabelMatcher(v),
		}
		rslt = append(rslt, entry)
	}
	return rslt
}

func FlattenAPCustomMatcher(x *ves_io_schema_alert_policy.CustomMatcher) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		v := map[string]interface{}{
			"alertlabel": FlattenAPAlertLabel(x.GetAlertlabel()),
			"alertname":  FlattenAPLabelMatcher(x.GetAlertname()),
			"group":      FlattenAPLabelMatcher(x.GetGroup()),
			"severity":   FlattenAPLabelMatcher(x.GetSeverity()),
		}
		rslt = append(rslt, v)
	}
	return rslt
}

func FlattenAPGroupMatcher(x *ves_io_schema_alert_policy.GroupMatcher) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		groups := make([]string, 0, len(x.GetGroups()))
		for _, g := range x.GetGroups() {
			groups = append(groups, g.String())
		}
		v := map[string]interface{}{
			"groups": groups,
		}
		rslt = append(rslt, v)
	}
	return rslt
}

func FlattenAPSeverityMatcher(x *ves_io_schema_alert_policy.SeverityMatcher) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		sev := make([]string, 0, len(x.GetSeverities()))
		for _, s := range x.GetSeverities() {
			sev = append(sev, s.String())
		}
		v := map[string]interface{}{
			"severities": sev,
		}
		rslt = append(rslt, v)
	}
	return rslt
}

func FlattenAPCustomGroupBy(x *ves_io_schema_alert_policy.CustomGroupBy) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		v := map[string]interface{}{
			"labels": x.GetLabels(),
		}
		rslt = append(rslt, v)
	}
	return rslt
}

func FlattenAPNotificationParameters(x *ves_io_schema_alert_policy.NotificationParameters) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		v := map[string]interface{}{
			"custom":          FlattenAPCustomGroupBy(x.GetCustom()),
			"default":         isEmpty(x.GetDefault()),
			"individual":      isEmpty(x.GetIndividual()),
			"ves_io_group":    isEmpty(x.GetVesIoGroup()),
			"group_interval":  x.GetGroupInterval(),
			"group_wait":      x.GetGroupWait(),
			"repeat_interval": x.GetRepeatInterval(),
		}
		rslt = append(rslt, v)
	}
	return rslt
}

func FlattenAPRoutes(x []*ves_io_schema_alert_policy.Route) []interface{} {
	rslt := make([]interface{}, 0)
	for _, r := range x {
		if r == nil {
			continue
		}
		// Route.Alertname lives inside the Matcher oneof. When a different
		// variant (any / custom / group / severity / alertname_regex) is
		// selected, Route.GetAlertname() returns a non-zero default
		// (SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN) instead of the zero value.
		// Only surface the value when the alertname variant is actually
		// selected, otherwise every plan would show a spurious drift of
		// alertname = "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN" -> null.
		alertname := ""
		if am, ok := r.GetMatcher().(*ves_io_schema_alert_policy.Route_Alertname); ok {
			alertname = am.Alertname.String()
		}
		v := map[string]interface{}{
			"dont_send":               isEmpty(r.GetDontSend()),
			"send":                    isEmpty(r.GetSend()),
			"alertname":               alertname,
			"alertname_regex":         r.GetAlertnameRegex(),
			"any":                     isEmpty(r.GetAny()),
			"custom":                  FlattenAPCustomMatcher(r.GetCustom()),
			"group":                   FlattenAPGroupMatcher(r.GetGroup()),
			"severity":                FlattenAPSeverityMatcher(r.GetSeverity()),
			"notification_parameters": FlattenAPNotificationParameters(r.GetNotificationParameters()),
		}
		rslt = append(rslt, v)
	}
	return rslt
}

func DriftDetectionSpec_AlertPolicy(d *schema.ResourceData, resp vesapi.GetObjectResponse) {
	spec := resp.GetObjSpec().(*ves_io_schema_alert_policy.SpecType)

	d.Set("notification_parameters", FlattenAPNotificationParameters(spec.GcSpec.GetNotificationParameters()))
	d.Set("receivers", FlattenAPReceivers(spec.GcSpec.GetReceivers()))
	d.Set("routes", FlattenAPRoutes(spec.GcSpec.GetRoutes()))
}
