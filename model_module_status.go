/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the ModuleStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleStatus{}

// ModuleStatus struct for ModuleStatus
type ModuleStatus struct {
	Value                *ModuleStatusValue `json:"value,omitempty"`
	Label                *ModuleStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleStatus ModuleStatus

// NewModuleStatus instantiates a new ModuleStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleStatus() *ModuleStatus {
	this := ModuleStatus{}
	return &this
}

// NewModuleStatusWithDefaults instantiates a new ModuleStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleStatusWithDefaults() *ModuleStatus {
	this := ModuleStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ModuleStatus) GetValue() ModuleStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret ModuleStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatus) GetValueOk() (*ModuleStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ModuleStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ModuleStatusValue and assigns it to the Value field.
func (o *ModuleStatus) SetValue(v ModuleStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ModuleStatus) GetLabel() ModuleStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret ModuleStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatus) GetLabelOk() (*ModuleStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ModuleStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given ModuleStatusLabel and assigns it to the Label field.
func (o *ModuleStatus) SetLabel(v ModuleStatusLabel) {
	o.Label = &v
}

func (o ModuleStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleStatus) ToMap() (map[string]interface{}, error) {
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

func (o *ModuleStatus) UnmarshalJSON(data []byte) (err error) {
	varModuleStatus := _ModuleStatus{}

	err = json.Unmarshal(data, &varModuleStatus)

	if err != nil {
		return err
	}

	*o = ModuleStatus(varModuleStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleStatus struct {
	value *ModuleStatus
	isSet bool
}

func (v NullableModuleStatus) Get() *ModuleStatus {
	return v.value
}

func (v *NullableModuleStatus) Set(val *ModuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatus(val *ModuleStatus) *NullableModuleStatus {
	return &NullableModuleStatus{value: val, isSet: true}
}

func (v NullableModuleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
