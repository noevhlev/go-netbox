/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the WritableVRFRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableVRFRequest{}

// WritableVRFRequest Adds support for custom fields and tags.
type WritableVRFRequest struct {
	Name string `json:"name"`
	// Unique route distinguisher (as defined in RFC 4364)
	Rd     NullableString `json:"rd,omitempty"`
	Tenant NullableInt32  `json:"tenant,omitempty"`
	// Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique        *bool                  `json:"enforce_unique,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	ImportTargets        []int32                `json:"import_targets,omitempty"`
	ExportTargets        []int32                `json:"export_targets,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableVRFRequest WritableVRFRequest

// NewWritableVRFRequest instantiates a new WritableVRFRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVRFRequest(name string) *WritableVRFRequest {
	this := WritableVRFRequest{}
	this.Name = name
	return &this
}

// NewWritableVRFRequestWithDefaults instantiates a new WritableVRFRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVRFRequestWithDefaults() *WritableVRFRequest {
	this := WritableVRFRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableVRFRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableVRFRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableVRFRequest) SetName(v string) {
	o.Name = v
}

// GetRd returns the Rd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVRFRequest) GetRd() string {
	if o == nil || IsNil(o.Rd.Get()) {
		var ret string
		return ret
	}
	return *o.Rd.Get()
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVRFRequest) GetRdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rd.Get(), o.Rd.IsSet()
}

// HasRd returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasRd() bool {
	if o != nil && o.Rd.IsSet() {
		return true
	}

	return false
}

// SetRd gets a reference to the given NullableString and assigns it to the Rd field.
func (o *WritableVRFRequest) SetRd(v string) {
	o.Rd.Set(&v)
}

// SetRdNil sets the value for Rd to be an explicit nil
func (o *WritableVRFRequest) SetRdNil() {
	o.Rd.Set(nil)
}

// UnsetRd ensures that no value is present for Rd, not even an explicit nil
func (o *WritableVRFRequest) UnsetRd() {
	o.Rd.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVRFRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVRFRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *WritableVRFRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableVRFRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableVRFRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetEnforceUnique returns the EnforceUnique field value if set, zero value otherwise.
func (o *WritableVRFRequest) GetEnforceUnique() bool {
	if o == nil || IsNil(o.EnforceUnique) {
		var ret bool
		return ret
	}
	return *o.EnforceUnique
}

// GetEnforceUniqueOk returns a tuple with the EnforceUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRFRequest) GetEnforceUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceUnique) {
		return nil, false
	}
	return o.EnforceUnique, true
}

// HasEnforceUnique returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasEnforceUnique() bool {
	if o != nil && !IsNil(o.EnforceUnique) {
		return true
	}

	return false
}

// SetEnforceUnique gets a reference to the given bool and assigns it to the EnforceUnique field.
func (o *WritableVRFRequest) SetEnforceUnique(v bool) {
	o.EnforceUnique = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableVRFRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRFRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableVRFRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableVRFRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRFRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableVRFRequest) SetComments(v string) {
	o.Comments = &v
}

// GetImportTargets returns the ImportTargets field value if set, zero value otherwise.
func (o *WritableVRFRequest) GetImportTargets() []int32 {
	if o == nil || IsNil(o.ImportTargets) {
		var ret []int32
		return ret
	}
	return o.ImportTargets
}

// GetImportTargetsOk returns a tuple with the ImportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRFRequest) GetImportTargetsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ImportTargets) {
		return nil, false
	}
	return o.ImportTargets, true
}

// HasImportTargets returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasImportTargets() bool {
	if o != nil && !IsNil(o.ImportTargets) {
		return true
	}

	return false
}

// SetImportTargets gets a reference to the given []int32 and assigns it to the ImportTargets field.
func (o *WritableVRFRequest) SetImportTargets(v []int32) {
	o.ImportTargets = v
}

// GetExportTargets returns the ExportTargets field value if set, zero value otherwise.
func (o *WritableVRFRequest) GetExportTargets() []int32 {
	if o == nil || IsNil(o.ExportTargets) {
		var ret []int32
		return ret
	}
	return o.ExportTargets
}

// GetExportTargetsOk returns a tuple with the ExportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRFRequest) GetExportTargetsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ExportTargets) {
		return nil, false
	}
	return o.ExportTargets, true
}

// HasExportTargets returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasExportTargets() bool {
	if o != nil && !IsNil(o.ExportTargets) {
		return true
	}

	return false
}

// SetExportTargets gets a reference to the given []int32 and assigns it to the ExportTargets field.
func (o *WritableVRFRequest) SetExportTargets(v []int32) {
	o.ExportTargets = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVRFRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRFRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableVRFRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVRFRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRFRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVRFRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableVRFRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableVRFRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableVRFRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Rd.IsSet() {
		toSerialize["rd"] = o.Rd.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.EnforceUnique) {
		toSerialize["enforce_unique"] = o.EnforceUnique
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.ImportTargets) {
		toSerialize["import_targets"] = o.ImportTargets
	}
	if !IsNil(o.ExportTargets) {
		toSerialize["export_targets"] = o.ExportTargets
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableVRFRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWritableVRFRequest := _WritableVRFRequest{}

	err = json.Unmarshal(data, &varWritableVRFRequest)

	if err != nil {
		return err
	}

	*o = WritableVRFRequest(varWritableVRFRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "rd")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "enforce_unique")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "import_targets")
		delete(additionalProperties, "export_targets")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableVRFRequest struct {
	value *WritableVRFRequest
	isSet bool
}

func (v NullableWritableVRFRequest) Get() *WritableVRFRequest {
	return v.value
}

func (v *NullableWritableVRFRequest) Set(val *WritableVRFRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVRFRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVRFRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVRFRequest(val *WritableVRFRequest) *NullableWritableVRFRequest {
	return &NullableWritableVRFRequest{value: val, isSet: true}
}

func (v NullableWritableVRFRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVRFRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
