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

// checks if the BriefCommunityList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefCommunityList{}

// BriefCommunityList Adds support for custom fields and tags.
type BriefCommunityList struct {
	Id                   int32  `json:"id"`
	Url                  string `json:"url"`
	Name                 string `json:"name"`
	Display              string `json:"display"`
	Description          string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _BriefCommunityList BriefCommunityList

// NewBriefCommunityList instantiates a new BriefCommunityList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefCommunityList(id int32, url string, name string, display string, description string) *BriefCommunityList {
	this := BriefCommunityList{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Display = display
	this.Description = description
	return &this
}

// NewBriefCommunityListWithDefaults instantiates a new BriefCommunityList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefCommunityListWithDefaults() *BriefCommunityList {
	this := BriefCommunityList{}
	return &this
}

// GetId returns the Id field value
func (o *BriefCommunityList) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefCommunityList) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefCommunityList) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefCommunityList) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefCommunityList) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefCommunityList) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *BriefCommunityList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefCommunityList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefCommunityList) SetName(v string) {
	o.Name = v
}

// GetDisplay returns the Display field value
func (o *BriefCommunityList) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefCommunityList) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefCommunityList) SetDisplay(v string) {
	o.Display = v
}

// GetDescription returns the Description field value
func (o *BriefCommunityList) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BriefCommunityList) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BriefCommunityList) SetDescription(v string) {
	o.Description = v
}

func (o BriefCommunityList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefCommunityList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["name"] = o.Name
	toSerialize["display"] = o.Display
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefCommunityList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"name",
		"display",
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

	varBriefCommunityList := _BriefCommunityList{}

	err = json.Unmarshal(data, &varBriefCommunityList)

	if err != nil {
		return err
	}

	*o = BriefCommunityList(varBriefCommunityList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "display")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefCommunityList struct {
	value *BriefCommunityList
	isSet bool
}

func (v NullableBriefCommunityList) Get() *BriefCommunityList {
	return v.value
}

func (v *NullableBriefCommunityList) Set(val *BriefCommunityList) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCommunityList) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCommunityList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCommunityList(val *BriefCommunityList) *NullableBriefCommunityList {
	return &NullableBriefCommunityList{value: val, isSet: true}
}

func (v NullableBriefCommunityList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCommunityList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
