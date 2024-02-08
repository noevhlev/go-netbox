/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the Dashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dashboard{}

// Dashboard struct for Dashboard
type Dashboard struct {
	Layout               interface{} `json:"layout,omitempty"`
	Config               interface{} `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Dashboard Dashboard

// NewDashboard instantiates a new Dashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboard() *Dashboard {
	this := Dashboard{}
	return &this
}

// NewDashboardWithDefaults instantiates a new Dashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardWithDefaults() *Dashboard {
	this := Dashboard{}
	return &this
}

// GetLayout returns the Layout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetLayout() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dashboard) GetLayoutOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return &o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *Dashboard) HasLayout() bool {
	if o != nil && IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given interface{} and assigns it to the Layout field.
func (o *Dashboard) SetLayout(v interface{}) {
	o.Layout = v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dashboard) GetConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Dashboard) HasConfig() bool {
	if o != nil && IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *Dashboard) SetConfig(v interface{}) {
	o.Config = v
}

func (o Dashboard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Dashboard) UnmarshalJSON(data []byte) (err error) {
	varDashboard := _Dashboard{}

	err = json.Unmarshal(data, &varDashboard)

	if err != nil {
		return err
	}

	*o = Dashboard(varDashboard)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "layout")
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDashboard struct {
	value *Dashboard
	isSet bool
}

func (v NullableDashboard) Get() *Dashboard {
	return v.value
}

func (v *NullableDashboard) Set(val *Dashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboard(val *Dashboard) *NullableDashboard {
	return &NullableDashboard{value: val, isSet: true}
}

func (v NullableDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
