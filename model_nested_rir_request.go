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

// checks if the NestedRIRRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedRIRRequest{}

// NestedRIRRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedRIRRequest struct {
	Name                 string `json:"name"`
	Slug                 string `json:"slug"`
	AdditionalProperties map[string]interface{}
}

type _NestedRIRRequest NestedRIRRequest

// NewNestedRIRRequest instantiates a new NestedRIRRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRIRRequest(name string, slug string) *NestedRIRRequest {
	this := NestedRIRRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedRIRRequestWithDefaults instantiates a new NestedRIRRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRIRRequestWithDefaults() *NestedRIRRequest {
	this := NestedRIRRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedRIRRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRIRRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRIRRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedRIRRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedRIRRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedRIRRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedRIRRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedRIRRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedRIRRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varNestedRIRRequest := _NestedRIRRequest{}

	err = json.Unmarshal(data, &varNestedRIRRequest)

	if err != nil {
		return err
	}

	*o = NestedRIRRequest(varNestedRIRRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedRIRRequest struct {
	value *NestedRIRRequest
	isSet bool
}

func (v NullableNestedRIRRequest) Get() *NestedRIRRequest {
	return v.value
}

func (v *NullableNestedRIRRequest) Set(val *NestedRIRRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRIRRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRIRRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRIRRequest(val *NestedRIRRequest) *NullableNestedRIRRequest {
	return &NullableNestedRIRRequest{value: val, isSet: true}
}

func (v NullableNestedRIRRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRIRRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
