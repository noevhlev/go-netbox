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

// checks if the NestedUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedUser{}

// NestedUser Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedUser struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username             string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _NestedUser NestedUser

// NewNestedUser instantiates a new NestedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedUser(id int32, url string, display string, username string) *NestedUser {
	this := NestedUser{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Username = username
	return &this
}

// NewNestedUserWithDefaults instantiates a new NestedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedUserWithDefaults() *NestedUser {
	this := NestedUser{}
	return &this
}

// GetId returns the Id field value
func (o *NestedUser) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedUser) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedUser) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedUser) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedUser) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedUser) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedUser) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedUser) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedUser) SetDisplay(v string) {
	o.Display = v
}

// GetUsername returns the Username field value
func (o *NestedUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *NestedUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *NestedUser) SetUsername(v string) {
	o.Username = v
}

func (o NestedUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"username",
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

	varNestedUser := _NestedUser{}

	err = json.Unmarshal(data, &varNestedUser)

	if err != nil {
		return err
	}

	*o = NestedUser(varNestedUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedUser struct {
	value *NestedUser
	isSet bool
}

func (v NullableNestedUser) Get() *NestedUser {
	return v.value
}

func (v *NullableNestedUser) Set(val *NestedUser) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedUser) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedUser(val *NestedUser) *NullableNestedUser {
	return &NullableNestedUser{value: val, isSet: true}
}

func (v NullableNestedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
