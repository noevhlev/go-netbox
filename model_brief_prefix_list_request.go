/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the BriefPrefixListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefPrefixListRequest{}

// BriefPrefixListRequest Adds support for custom fields and tags.
type BriefPrefixListRequest struct {
	Name                 string `json:"name"`
	Description          string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _BriefPrefixListRequest BriefPrefixListRequest

// NewBriefPrefixListRequest instantiates a new BriefPrefixListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefPrefixListRequest(name string, description string) *BriefPrefixListRequest {
	this := BriefPrefixListRequest{}
	this.Name = name
	this.Description = description
	return &this
}

// NewBriefPrefixListRequestWithDefaults instantiates a new BriefPrefixListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefPrefixListRequestWithDefaults() *BriefPrefixListRequest {
	this := BriefPrefixListRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefPrefixListRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefPrefixListRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefPrefixListRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *BriefPrefixListRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BriefPrefixListRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BriefPrefixListRequest) SetDescription(v string) {
	o.Description = v
}

func (o BriefPrefixListRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefPrefixListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefPrefixListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
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

	varBriefPrefixListRequest := _BriefPrefixListRequest{}

	err = json.Unmarshal(data, &varBriefPrefixListRequest)

	if err != nil {
		return err
	}

	*o = BriefPrefixListRequest(varBriefPrefixListRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefPrefixListRequest struct {
	value *BriefPrefixListRequest
	isSet bool
}

func (v NullableBriefPrefixListRequest) Get() *BriefPrefixListRequest {
	return v.value
}

func (v *NullableBriefPrefixListRequest) Set(val *BriefPrefixListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefPrefixListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefPrefixListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefPrefixListRequest(val *BriefPrefixListRequest) *NullableBriefPrefixListRequest {
	return &NullableBriefPrefixListRequest{value: val, isSet: true}
}

func (v NullableBriefPrefixListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefPrefixListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}