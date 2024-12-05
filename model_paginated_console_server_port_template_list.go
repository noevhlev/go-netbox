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

// checks if the PaginatedConsoleServerPortTemplateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedConsoleServerPortTemplateList{}

// PaginatedConsoleServerPortTemplateList struct for PaginatedConsoleServerPortTemplateList
type PaginatedConsoleServerPortTemplateList struct {
	Count                int32                       `json:"count"`
	Next                 NullableString              `json:"next,omitempty"`
	Previous             NullableString              `json:"previous,omitempty"`
	Results              []ConsoleServerPortTemplate `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedConsoleServerPortTemplateList PaginatedConsoleServerPortTemplateList

// NewPaginatedConsoleServerPortTemplateList instantiates a new PaginatedConsoleServerPortTemplateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedConsoleServerPortTemplateList(count int32, results []ConsoleServerPortTemplate) *PaginatedConsoleServerPortTemplateList {
	this := PaginatedConsoleServerPortTemplateList{}
	this.Count = count
	this.Results = results
	return &this
}

// NewPaginatedConsoleServerPortTemplateListWithDefaults instantiates a new PaginatedConsoleServerPortTemplateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedConsoleServerPortTemplateListWithDefaults() *PaginatedConsoleServerPortTemplateList {
	this := PaginatedConsoleServerPortTemplateList{}
	return &this
}

// GetCount returns the Count field value
func (o *PaginatedConsoleServerPortTemplateList) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedConsoleServerPortTemplateList) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedConsoleServerPortTemplateList) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedConsoleServerPortTemplateList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedConsoleServerPortTemplateList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedConsoleServerPortTemplateList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedConsoleServerPortTemplateList) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedConsoleServerPortTemplateList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedConsoleServerPortTemplateList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedConsoleServerPortTemplateList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedConsoleServerPortTemplateList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedConsoleServerPortTemplateList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedConsoleServerPortTemplateList) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedConsoleServerPortTemplateList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedConsoleServerPortTemplateList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PaginatedConsoleServerPortTemplateList) GetResults() []ConsoleServerPortTemplate {
	if o == nil {
		var ret []ConsoleServerPortTemplate
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedConsoleServerPortTemplateList) GetResultsOk() ([]ConsoleServerPortTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedConsoleServerPortTemplateList) SetResults(v []ConsoleServerPortTemplate) {
	o.Results = v
}

func (o PaginatedConsoleServerPortTemplateList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedConsoleServerPortTemplateList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedConsoleServerPortTemplateList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"results",
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

	varPaginatedConsoleServerPortTemplateList := _PaginatedConsoleServerPortTemplateList{}

	err = json.Unmarshal(data, &varPaginatedConsoleServerPortTemplateList)

	if err != nil {
		return err
	}

	*o = PaginatedConsoleServerPortTemplateList(varPaginatedConsoleServerPortTemplateList)

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

type NullablePaginatedConsoleServerPortTemplateList struct {
	value *PaginatedConsoleServerPortTemplateList
	isSet bool
}

func (v NullablePaginatedConsoleServerPortTemplateList) Get() *PaginatedConsoleServerPortTemplateList {
	return v.value
}

func (v *NullablePaginatedConsoleServerPortTemplateList) Set(val *PaginatedConsoleServerPortTemplateList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedConsoleServerPortTemplateList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedConsoleServerPortTemplateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedConsoleServerPortTemplateList(val *PaginatedConsoleServerPortTemplateList) *NullablePaginatedConsoleServerPortTemplateList {
	return &NullablePaginatedConsoleServerPortTemplateList{value: val, isSet: true}
}

func (v NullablePaginatedConsoleServerPortTemplateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedConsoleServerPortTemplateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
