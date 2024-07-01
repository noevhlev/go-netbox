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

// checks if the CableStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CableStatus{}

// CableStatus struct for CableStatus
type CableStatus struct {
	Value                *CableStatusValue `json:"value,omitempty"`
	Label                *CableStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CableStatus CableStatus

// NewCableStatus instantiates a new CableStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCableStatus() *CableStatus {
	this := CableStatus{}
	return &this
}

// NewCableStatusWithDefaults instantiates a new CableStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCableStatusWithDefaults() *CableStatus {
	this := CableStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CableStatus) GetValue() CableStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret CableStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableStatus) GetValueOk() (*CableStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CableStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CableStatusValue and assigns it to the Value field.
func (o *CableStatus) SetValue(v CableStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CableStatus) GetLabel() CableStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret CableStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableStatus) GetLabelOk() (*CableStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CableStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CableStatusLabel and assigns it to the Label field.
func (o *CableStatus) SetLabel(v CableStatusLabel) {
	o.Label = &v
}

func (o CableStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CableStatus) ToMap() (map[string]interface{}, error) {
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

func (o *CableStatus) UnmarshalJSON(data []byte) (err error) {
	varCableStatus := _CableStatus{}

	err = json.Unmarshal(data, &varCableStatus)

	if err != nil {
		return err
	}

	*o = CableStatus(varCableStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCableStatus struct {
	value *CableStatus
	isSet bool
}

func (v NullableCableStatus) Get() *CableStatus {
	return v.value
}

func (v *NullableCableStatus) Set(val *CableStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableStatus(val *CableStatus) *NullableCableStatus {
	return &NullableCableStatus{value: val, isSet: true}
}

func (v NullableCableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
