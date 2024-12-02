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

// checks if the CommunityStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityStatus{}

// CommunityStatus struct for CommunityStatus
type CommunityStatus struct {
	Value                *CommunityStatusValue `json:"value,omitempty"`
	Label                *CommunityStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommunityStatus CommunityStatus

// NewCommunityStatus instantiates a new CommunityStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityStatus() *CommunityStatus {
	this := CommunityStatus{}
	return &this
}

// NewCommunityStatusWithDefaults instantiates a new CommunityStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityStatusWithDefaults() *CommunityStatus {
	this := CommunityStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CommunityStatus) GetValue() CommunityStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret CommunityStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityStatus) GetValueOk() (*CommunityStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CommunityStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CommunityStatusValue and assigns it to the Value field.
func (o *CommunityStatus) SetValue(v CommunityStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CommunityStatus) GetLabel() CommunityStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret CommunityStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityStatus) GetLabelOk() (*CommunityStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CommunityStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CommunityStatusLabel and assigns it to the Label field.
func (o *CommunityStatus) SetLabel(v CommunityStatusLabel) {
	o.Label = &v
}

func (o CommunityStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityStatus) ToMap() (map[string]interface{}, error) {
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

func (o *CommunityStatus) UnmarshalJSON(data []byte) (err error) {
	varCommunityStatus := _CommunityStatus{}

	err = json.Unmarshal(data, &varCommunityStatus)

	if err != nil {
		return err
	}

	*o = CommunityStatus(varCommunityStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommunityStatus struct {
	value *CommunityStatus
	isSet bool
}

func (v NullableCommunityStatus) Get() *CommunityStatus {
	return v.value
}

func (v *NullableCommunityStatus) Set(val *CommunityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityStatus(val *CommunityStatus) *NullableCommunityStatus {
	return &NullableCommunityStatus{value: val, isSet: true}
}

func (v NullableCommunityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
