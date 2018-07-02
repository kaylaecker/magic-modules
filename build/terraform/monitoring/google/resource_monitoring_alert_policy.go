// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMonitoringAlertPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringAlertPolicyCreate,
		Read: resourceMonitoringAlertPolicyRead,
				Delete: resourceMonitoringAlertPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringAlertPolicyImport,
		},


		Schema: map[string]*schema.Schema{
	"display_name": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
},
	"conditions": {
    Type: schema.TypeList,
    Optional: true,
  ForceNew: true,
	        Elem: &schema.Resource{
        Schema: map[string]*schema.Schema{
                      "name": {
    Type: schema.TypeString,
    Optional: true,
  ForceNew: true,
},
                      "display_name": {
    Type: schema.TypeString,
    Optional: true,
  ForceNew: true,
},
                      "condition_threshold": {
    Type: schema.TypeList,
    Optional: true,
  ForceNew: true,
  MaxItems: 1,
  Elem: &schema.Resource{
    Schema: map[string]*schema.Schema{
              "duration": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
},
              "filter": {
    Type: schema.TypeString,
    Optional: true,
  ForceNew: true,
},
          },
  },
},
                      "condition_absent": {
    Type: schema.TypeList,
    Optional: true,
  ForceNew: true,
  MaxItems: 1,
  Elem: &schema.Resource{
    Schema: map[string]*schema.Schema{
              "filter": {
    Type: schema.TypeString,
    Optional: true,
  ForceNew: true,
},
          },
  },
},
                  },
      },
    },
	"name": {
    Type: schema.TypeString,
    Computed: true,
},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMonitoringAlertPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	displayNameProp, err := expandMonitoringAlertPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	}
	conditionsProp, err := expandMonitoringAlertPolicyConditions(d.Get("conditions"), d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
	"displayName": displayNameProp,
	"conditions": conditionsProp,
	}

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/projects/{{project}}/alertPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AlertPolicy: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating AlertPolicy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)


	log.Printf("[DEBUG] Finished creating AlertPolicy %q: %#v", d.Id(), res)


	return resourceMonitoringAlertPolicyRead(d, meta)
}

func resourceMonitoringAlertPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

  	project, err := getProject(d, config)
	if err != nil {
		return err
  }

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/projects/{{project}}/alertPolicies/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringAlertPolicy %q", d.Id()))
	}





  if err := d.Set("name", flattenMonitoringAlertPolicyName(res["name"])); err != nil {
		return fmt.Errorf("Error reading AlertPolicy: %s", err)
	}
  if err := d.Set("display_name", flattenMonitoringAlertPolicyDisplayName(res["displayName"])); err != nil {
		return fmt.Errorf("Error reading AlertPolicy: %s", err)
	}
  if err := d.Set("conditions", flattenMonitoringAlertPolicyConditions(res["conditions"])); err != nil {
		return fmt.Errorf("Error reading AlertPolicy: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AlertPolicy: %s", err)
	}

	return nil
}


func resourceMonitoringAlertPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/projects/{{project}}/alertPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting AlertPolicy %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return handleNotFoundError(err, d, "AlertPolicy")
	}


	log.Printf("[DEBUG] Finished deleting AlertPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringAlertPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"/projects/(?P<project>[^/]+)/alertPolicies/(?P<name>[^/]+)","(?P<project>[^/]+)/(?P<name>[^/]+)","(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringAlertPolicyName(v interface{}) interface{} {
  return v
}

func flattenMonitoringAlertPolicyDisplayName(v interface{}) interface{} {
  return v
}

func flattenMonitoringAlertPolicyConditions(v interface{}) interface{} {
  l := v.([]interface{})
  transformed := make([]interface{}, 0, len(l))
  for _, raw := range l {
    original := raw.(map[string]interface{})
    transformed = append(transformed, map[string]interface{}{
          "name": flattenMonitoringAlertPolicyConditionsName(original["name"]),
          "display_name": flattenMonitoringAlertPolicyConditionsDisplayName(original["displayName"]),
          "condition_threshold": flattenMonitoringAlertPolicyConditionsConditionThreshold(original["conditionThreshold"]),
          "condition_absent": flattenMonitoringAlertPolicyConditionsConditionAbsent(original["conditionAbsent"]),
        })
  }
  return transformed
}
      func flattenMonitoringAlertPolicyConditionsName(v interface{}) interface{} {
  return v
}

      func flattenMonitoringAlertPolicyConditionsDisplayName(v interface{}) interface{} {
  return v
}

      func flattenMonitoringAlertPolicyConditionsConditionThreshold(v interface{}) interface{} {
  if v == nil {
    return nil
  }
  original := v.(map[string]interface{})
  transformed := make(map[string]interface{})
      transformed["filter"] =
    flattenMonitoringAlertPolicyConditionsConditionThresholdFilter(original["filter"])
      transformed["duration"] =
    flattenMonitoringAlertPolicyConditionsConditionThresholdDuration(original["duration"])
    return []interface{}{transformed}
}
      func flattenMonitoringAlertPolicyConditionsConditionThresholdFilter(v interface{}) interface{} {
  return v
}

      func flattenMonitoringAlertPolicyConditionsConditionThresholdDuration(v interface{}) interface{} {
  return v
}



      func flattenMonitoringAlertPolicyConditionsConditionAbsent(v interface{}) interface{} {
  if v == nil {
    return nil
  }
  original := v.(map[string]interface{})
  transformed := make(map[string]interface{})
      transformed["filter"] =
    flattenMonitoringAlertPolicyConditionsConditionAbsentFilter(original["filter"])
    return []interface{}{transformed}
}
      func flattenMonitoringAlertPolicyConditionsConditionAbsentFilter(v interface{}) interface{} {
  return v
}






func expandMonitoringAlertPolicyDisplayName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
    return v, nil
}


func expandMonitoringAlertPolicyConditions(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  l := v.([]interface{})
  req := make([]interface{}, 0, len(l))
  for _, raw := range l {
    original := raw.(map[string]interface{})
    transformed := make(map[string]interface{})

          transformedName, err := expandMonitoringAlertPolicyConditionsName(original["name"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["name"] = transformedName
          transformedDisplayName, err := expandMonitoringAlertPolicyConditionsDisplayName(original["display_name"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["displayName"] = transformedDisplayName
          transformedConditionThreshold, err := expandMonitoringAlertPolicyConditionsConditionThreshold(original["condition_threshold"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["conditionThreshold"] = transformedConditionThreshold
          transformedConditionAbsent, err := expandMonitoringAlertPolicyConditionsConditionAbsent(original["condition_absent"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["conditionAbsent"] = transformedConditionAbsent

    req = append(req, transformed)
  }
  return req, nil
}

      func expandMonitoringAlertPolicyConditionsName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
    return v, nil
}


      func expandMonitoringAlertPolicyConditionsDisplayName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
    return v, nil
}


      func expandMonitoringAlertPolicyConditionsConditionThreshold(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFilter, err := expandMonitoringAlertPolicyConditionsConditionThresholdFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["filter"] = transformedFilter
	return transformed, nil
}


func expandMonitoringAlertPolicyConditionsConditionThresholdFilter(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

      func expandMonitoringAlertPolicyConditionsConditionAbsent(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFilter, err := expandMonitoringAlertPolicyConditionsConditionAbsentFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["filter"] = transformedFilter
	return transformed, nil
}


func expandMonitoringAlertPolicyConditionsConditionAbsentFilter(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
