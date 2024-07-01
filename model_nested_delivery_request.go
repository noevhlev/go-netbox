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

// checks if the NestedDeliveryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedDeliveryRequest{}

// NestedDeliveryRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedDeliveryRequest struct {
	Name string `json:"name"`
	// Date when this delivery was made
	Date                 NullableString `json:"date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedDeliveryRequest NestedDeliveryRequest

// NewNestedDeliveryRequest instantiates a new NestedDeliveryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedDeliveryRequest(name string) *NestedDeliveryRequest {
	this := NestedDeliveryRequest{}
	this.Name = name
	return &this
}

// NewNestedDeliveryRequestWithDefaults instantiates a new NestedDeliveryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedDeliveryRequestWithDefaults() *NestedDeliveryRequest {
	this := NestedDeliveryRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedDeliveryRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedDeliveryRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedDeliveryRequest) SetName(v string) {
	o.Name = v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedDeliveryRequest) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedDeliveryRequest) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *NestedDeliveryRequest) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *NestedDeliveryRequest) SetDate(v string) {
	o.Date.Set(&v)
}

// SetDateNil sets the value for Date to be an explicit nil
func (o *NestedDeliveryRequest) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *NestedDeliveryRequest) UnsetDate() {
	o.Date.Unset()
}

func (o NestedDeliveryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedDeliveryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedDeliveryRequest) UnmarshalJSON(data []byte) (err error) {
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

	varNestedDeliveryRequest := _NestedDeliveryRequest{}

	err = json.Unmarshal(data, &varNestedDeliveryRequest)

	if err != nil {
		return err
	}

	*o = NestedDeliveryRequest(varNestedDeliveryRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedDeliveryRequest struct {
	value *NestedDeliveryRequest
	isSet bool
}

func (v NullableNestedDeliveryRequest) Get() *NestedDeliveryRequest {
	return v.value
}

func (v *NullableNestedDeliveryRequest) Set(val *NestedDeliveryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedDeliveryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedDeliveryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedDeliveryRequest(val *NestedDeliveryRequest) *NullableNestedDeliveryRequest {
	return &NullableNestedDeliveryRequest{value: val, isSet: true}
}

func (v NullableNestedDeliveryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedDeliveryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
