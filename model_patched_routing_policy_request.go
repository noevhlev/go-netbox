/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedRoutingPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedRoutingPolicyRequest{}

// PatchedRoutingPolicyRequest Adds support for custom fields and tags.
type PatchedRoutingPolicyRequest struct {
	Name                 *string                `json:"name,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedRoutingPolicyRequest PatchedRoutingPolicyRequest

// NewPatchedRoutingPolicyRequest instantiates a new PatchedRoutingPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRoutingPolicyRequest() *PatchedRoutingPolicyRequest {
	this := PatchedRoutingPolicyRequest{}
	return &this
}

// NewPatchedRoutingPolicyRequestWithDefaults instantiates a new PatchedRoutingPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRoutingPolicyRequestWithDefaults() *PatchedRoutingPolicyRequest {
	this := PatchedRoutingPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRoutingPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedRoutingPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedRoutingPolicyRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedRoutingPolicyRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedRoutingPolicyRequest) SetComments(v string) {
	o.Comments = &v
}

func (o PatchedRoutingPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedRoutingPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedRoutingPolicyRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedRoutingPolicyRequest := _PatchedRoutingPolicyRequest{}

	err = json.Unmarshal(data, &varPatchedRoutingPolicyRequest)

	if err != nil {
		return err
	}

	*o = PatchedRoutingPolicyRequest(varPatchedRoutingPolicyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedRoutingPolicyRequest struct {
	value *PatchedRoutingPolicyRequest
	isSet bool
}

func (v NullablePatchedRoutingPolicyRequest) Get() *PatchedRoutingPolicyRequest {
	return v.value
}

func (v *NullablePatchedRoutingPolicyRequest) Set(val *PatchedRoutingPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRoutingPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRoutingPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRoutingPolicyRequest(val *PatchedRoutingPolicyRequest) *NullablePatchedRoutingPolicyRequest {
	return &NullablePatchedRoutingPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedRoutingPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRoutingPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
