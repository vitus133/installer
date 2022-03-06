// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudAssetFolderFeed() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAssetFolderFeedCreate,
		Read:   resourceCloudAssetFolderFeedRead,
		Update: resourceCloudAssetFolderFeedUpdate,
		Delete: resourceCloudAssetFolderFeedDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudAssetFolderFeedImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"billing_project": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The project whose identity will be used when sending messages to the
destination pubsub topic. It also specifies the project for API 
enablement check, quota, and billing.`,
			},
			"feed_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.`,
			},
			"feed_output_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Output configuration for asset feed destination.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pubsub_destination": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Destination on Cloud Pubsub.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topic": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Destination on Cloud Pubsub topic.`,
									},
								},
							},
						},
					},
				},
			},
			"folder": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The folder this feed should be created in.`,
			},
			"asset_names": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of the full names of the assets to receive updates. You must specify either or both of 
assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"asset_types": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of types of the assets to receive updates. You must specify either or both of assetNames
and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
the feed. For example: "compute.googleapis.com/Disk"
See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
supported asset types.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"condition": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A condition which determines whether an asset update should be published. If specified, an asset
will be returned only when the expression evaluates to true. When set, expression field
must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
condition are optional.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Textual representation of an expression in Common Expression Language syntax.`,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Description of the expression. This is a longer text which describes the expression,
e.g. when hovered over it in a UI.`,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `String indicating the location of the expression for error reporting, e.g. a file 
name and a position in the file.`,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Title for the expression, i.e. a short string describing its purpose.
This can be used e.g. in UIs which allow to enter the expression.`,
						},
					},
				},
			},
			"content_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "ACCESS_POLICY", ""}, false),
				Description:  `Asset content type. If not specified, no content but the asset name and type will be returned. Possible values: ["CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "ACCESS_POLICY"]`,
			},
			"folder_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The ID of the folder where this feed has been created. Both [FOLDER_NUMBER]
and folders/[FOLDER_NUMBER] are accepted.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceCloudAssetFolderFeedCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetFolderFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_names"); !isEmptyValue(reflect.ValueOf(assetNamesProp)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetFolderFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_types"); !isEmptyValue(reflect.ValueOf(assetTypesProp)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetFolderFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_type"); !isEmptyValue(reflect.ValueOf(contentTypeProp)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetFolderFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("feed_output_config"); !isEmptyValue(reflect.ValueOf(feedOutputConfigProp)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetFolderFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !isEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	obj, err = resourceCloudAssetFolderFeedEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudAssetBasePath}}folders/{{folder_id}}/feeds?feedId={{feed_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FolderFeed: %#v", obj)
	billingProject := ""

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// Send the project ID in the X-Goog-User-Project header.
	origUserProjectOverride := config.UserProjectOverride
	config.UserProjectOverride = true
	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating FolderFeed: %s", err)
	}
	if err := d.Set("name", flattenCloudAssetFolderFeedName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Restore the original value of user_project_override.
	config.UserProjectOverride = origUserProjectOverride

	log.Printf("[DEBUG] Finished creating FolderFeed %q: %#v", d.Id(), res)

	return resourceCloudAssetFolderFeedRead(d, meta)
}

func resourceCloudAssetFolderFeedRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudAssetFolderFeed %q", d.Id()))
	}

	if err := d.Set("folder_id", flattenCloudAssetFolderFeedFolderId(res["folder_id"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderFeed: %s", err)
	}
	if err := d.Set("name", flattenCloudAssetFolderFeedName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderFeed: %s", err)
	}
	if err := d.Set("asset_names", flattenCloudAssetFolderFeedAssetNames(res["assetNames"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderFeed: %s", err)
	}
	if err := d.Set("asset_types", flattenCloudAssetFolderFeedAssetTypes(res["assetTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderFeed: %s", err)
	}
	if err := d.Set("content_type", flattenCloudAssetFolderFeedContentType(res["contentType"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderFeed: %s", err)
	}
	if err := d.Set("feed_output_config", flattenCloudAssetFolderFeedFeedOutputConfig(res["feedOutputConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderFeed: %s", err)
	}
	if err := d.Set("condition", flattenCloudAssetFolderFeedCondition(res["condition"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderFeed: %s", err)
	}

	return nil
}

func resourceCloudAssetFolderFeedUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetFolderFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_names"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetFolderFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_types"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetFolderFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetFolderFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("feed_output_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetFolderFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	obj, err = resourceCloudAssetFolderFeedEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FolderFeed %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("asset_names") {
		updateMask = append(updateMask, "assetNames")
	}

	if d.HasChange("asset_types") {
		updateMask = append(updateMask, "assetTypes")
	}

	if d.HasChange("content_type") {
		updateMask = append(updateMask, "contentType")
	}

	if d.HasChange("feed_output_config") {
		updateMask = append(updateMask, "feedOutputConfig")
	}

	if d.HasChange("condition") {
		updateMask = append(updateMask, "condition")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating FolderFeed %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating FolderFeed %q: %#v", d.Id(), res)
	}

	return resourceCloudAssetFolderFeedRead(d, meta)
}

func resourceCloudAssetFolderFeedDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	log.Printf("[DEBUG] Deleting FolderFeed %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "FolderFeed")
	}

	log.Printf("[DEBUG] Finished deleting FolderFeed %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudAssetFolderFeedImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := d.Set("name", d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func flattenCloudAssetFolderFeedFolderId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedAssetNames(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedAssetTypes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedContentType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedFeedOutputConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_destination"] =
		flattenCloudAssetFolderFeedFeedOutputConfigPubsubDestination(original["pubsubDestination"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetFolderFeedFeedOutputConfigPubsubDestination(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["topic"] =
		flattenCloudAssetFolderFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetFolderFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedCondition(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expression"] =
		flattenCloudAssetFolderFeedConditionExpression(original["expression"], d, config)
	transformed["title"] =
		flattenCloudAssetFolderFeedConditionTitle(original["title"], d, config)
	transformed["description"] =
		flattenCloudAssetFolderFeedConditionDescription(original["description"], d, config)
	transformed["location"] =
		flattenCloudAssetFolderFeedConditionLocation(original["location"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetFolderFeedConditionExpression(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedConditionTitle(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedConditionDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetFolderFeedConditionLocation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandCloudAssetFolderFeedAssetNames(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedAssetTypes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedContentType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedFeedOutputConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubDestination, err := expandCloudAssetFolderFeedFeedOutputConfigPubsubDestination(original["pubsub_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubDestination); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubDestination"] = transformedPubsubDestination
	}

	return transformed, nil
}

func expandCloudAssetFolderFeedFeedOutputConfigPubsubDestination(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopic, err := expandCloudAssetFolderFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandCloudAssetFolderFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedCondition(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandCloudAssetFolderFeedConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !isEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandCloudAssetFolderFeedConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !isEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandCloudAssetFolderFeedConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandCloudAssetFolderFeedConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandCloudAssetFolderFeedConditionExpression(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedConditionTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedConditionDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedConditionLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceCloudAssetFolderFeedEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Remove the "folders/" prefix from the folder ID
	if folder, ok := d.GetOkExists("folder"); ok {
		if err := d.Set("folder_id", strings.TrimPrefix(folder.(string), "folders/")); err != nil {
			return nil, fmt.Errorf("Error setting folder_id: %s", err)
		}
	}
	// The feed object must be under the "feed" attribute on the request.
	newObj := make(map[string]interface{})
	newObj["feed"] = obj
	return newObj, nil
}