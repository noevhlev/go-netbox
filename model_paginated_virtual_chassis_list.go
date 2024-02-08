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

// checks if the PaginatedVirtualChassisList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedVirtualChassisList{}

// PaginatedVirtualChassisList struct for PaginatedVirtualChassisList
type PaginatedVirtualChassisList struct {
	Count                *int32           `json:"count,omitempty"`
	Next                 NullableString   `json:"next,omitempty"`
	Previous             NullableString   `json:"previous,omitempty"`
	Results              []VirtualChassis `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedVirtualChassisList PaginatedVirtualChassisList

// NewPaginatedVirtualChassisList instantiates a new PaginatedVirtualChassisList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedVirtualChassisList() *PaginatedVirtualChassisList {
	this := PaginatedVirtualChassisList{}
	return &this
}

// NewPaginatedVirtualChassisListWithDefaults instantiates a new PaginatedVirtualChassisList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedVirtualChassisListWithDefaults() *PaginatedVirtualChassisList {
	this := PaginatedVirtualChassisList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedVirtualChassisList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedVirtualChassisList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedVirtualChassisList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedVirtualChassisList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedVirtualChassisList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedVirtualChassisList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedVirtualChassisList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedVirtualChassisList) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedVirtualChassisList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedVirtualChassisList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedVirtualChassisList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedVirtualChassisList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedVirtualChassisList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedVirtualChassisList) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedVirtualChassisList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedVirtualChassisList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedVirtualChassisList) GetResults() []VirtualChassis {
	if o == nil || IsNil(o.Results) {
		var ret []VirtualChassis
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedVirtualChassisList) GetResultsOk() ([]VirtualChassis, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedVirtualChassisList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []VirtualChassis and assigns it to the Results field.
func (o *PaginatedVirtualChassisList) SetResults(v []VirtualChassis) {
	o.Results = v
}

func (o PaginatedVirtualChassisList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedVirtualChassisList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedVirtualChassisList) UnmarshalJSON(data []byte) (err error) {
	varPaginatedVirtualChassisList := _PaginatedVirtualChassisList{}

	err = json.Unmarshal(data, &varPaginatedVirtualChassisList)

	if err != nil {
		return err
	}

	*o = PaginatedVirtualChassisList(varPaginatedVirtualChassisList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedVirtualChassisList struct {
	value *PaginatedVirtualChassisList
	isSet bool
}

func (v NullablePaginatedVirtualChassisList) Get() *PaginatedVirtualChassisList {
	return v.value
}

func (v *NullablePaginatedVirtualChassisList) Set(val *PaginatedVirtualChassisList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedVirtualChassisList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedVirtualChassisList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedVirtualChassisList(val *PaginatedVirtualChassisList) *NullablePaginatedVirtualChassisList {
	return &NullablePaginatedVirtualChassisList{value: val, isSet: true}
}

func (v NullablePaginatedVirtualChassisList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedVirtualChassisList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
