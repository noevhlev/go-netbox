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

// checks if the L2VPNType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L2VPNType{}

// L2VPNType struct for L2VPNType
type L2VPNType struct {
	Value                *L2VPNTypeValue `json:"value,omitempty"`
	Label                *L2VPNTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _L2VPNType L2VPNType

// NewL2VPNType instantiates a new L2VPNType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL2VPNType() *L2VPNType {
	this := L2VPNType{}
	return &this
}

// NewL2VPNTypeWithDefaults instantiates a new L2VPNType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL2VPNTypeWithDefaults() *L2VPNType {
	this := L2VPNType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *L2VPNType) GetValue() L2VPNTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret L2VPNTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNType) GetValueOk() (*L2VPNTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *L2VPNType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given L2VPNTypeValue and assigns it to the Value field.
func (o *L2VPNType) SetValue(v L2VPNTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *L2VPNType) GetLabel() L2VPNTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret L2VPNTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNType) GetLabelOk() (*L2VPNTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *L2VPNType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given L2VPNTypeLabel and assigns it to the Label field.
func (o *L2VPNType) SetLabel(v L2VPNTypeLabel) {
	o.Label = &v
}

func (o L2VPNType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L2VPNType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *L2VPNType) UnmarshalJSON(data []byte) (err error) {
	varL2VPNType := _L2VPNType{}

	err = json.Unmarshal(data, &varL2VPNType)

	if err != nil {
		return err
	}

	*o = L2VPNType(varL2VPNType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableL2VPNType struct {
	value *L2VPNType
	isSet bool
}

func (v NullableL2VPNType) Get() *L2VPNType {
	return v.value
}

func (v *NullableL2VPNType) Set(val *L2VPNType) {
	v.value = val
	v.isSet = true
}

func (v NullableL2VPNType) IsSet() bool {
	return v.isSet
}

func (v *NullableL2VPNType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2VPNType(val *L2VPNType) *NullableL2VPNType {
	return &NullableL2VPNType{value: val, isSet: true}
}

func (v NullableL2VPNType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2VPNType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
