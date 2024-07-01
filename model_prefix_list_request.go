/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the PrefixListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrefixListRequest{}

// PrefixListRequest Adds support for custom fields and tags.
type PrefixListRequest struct {
	Name                 string                         `json:"name"`
	Description          string                         `json:"description"`
	Family               PatchedPrefixListRequestFamily `json:"family"`
	Tags                 []NestedTagRequest             `json:"tags,omitempty"`
	CustomFields         map[string]interface{}         `json:"custom_fields,omitempty"`
	Comments             *string                        `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrefixListRequest PrefixListRequest

// NewPrefixListRequest instantiates a new PrefixListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrefixListRequest(name string, description string, family PatchedPrefixListRequestFamily) *PrefixListRequest {
	this := PrefixListRequest{}
	this.Name = name
	this.Description = description
	this.Family = family
	return &this
}

// NewPrefixListRequestWithDefaults instantiates a new PrefixListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrefixListRequestWithDefaults() *PrefixListRequest {
	this := PrefixListRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PrefixListRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrefixListRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrefixListRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *PrefixListRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PrefixListRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PrefixListRequest) SetDescription(v string) {
	o.Description = v
}

// GetFamily returns the Family field value
func (o *PrefixListRequest) GetFamily() PatchedPrefixListRequestFamily {
	if o == nil {
		var ret PatchedPrefixListRequestFamily
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *PrefixListRequest) GetFamilyOk() (*PatchedPrefixListRequestFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *PrefixListRequest) SetFamily(v PatchedPrefixListRequestFamily) {
	o.Family = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PrefixListRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixListRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PrefixListRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PrefixListRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PrefixListRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixListRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PrefixListRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PrefixListRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PrefixListRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixListRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PrefixListRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PrefixListRequest) SetComments(v string) {
	o.Comments = &v
}

func (o PrefixListRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrefixListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["family"] = o.Family
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

func (o *PrefixListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"family",
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

	varPrefixListRequest := _PrefixListRequest{}

	err = json.Unmarshal(data, &varPrefixListRequest)

	if err != nil {
		return err
	}

	*o = PrefixListRequest(varPrefixListRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "family")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrefixListRequest struct {
	value *PrefixListRequest
	isSet bool
}

func (v NullablePrefixListRequest) Get() *PrefixListRequest {
	return v.value
}

func (v *NullablePrefixListRequest) Set(val *PrefixListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePrefixListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePrefixListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrefixListRequest(val *PrefixListRequest) *NullablePrefixListRequest {
	return &NullablePrefixListRequest{value: val, isSet: true}
}

func (v NullablePrefixListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrefixListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
