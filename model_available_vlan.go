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

// checks if the AvailableVLAN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableVLAN{}

// AvailableVLAN Representation of a VLAN which does not exist in the database.
type AvailableVLAN struct {
	Vid                  int32                  `json:"vid"`
	Group                NullableBriefVLANGroup `json:"group"`
	AdditionalProperties map[string]interface{}
}

type _AvailableVLAN AvailableVLAN

// NewAvailableVLAN instantiates a new AvailableVLAN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableVLAN(vid int32, group NullableBriefVLANGroup) *AvailableVLAN {
	this := AvailableVLAN{}
	this.Vid = vid
	this.Group = group
	return &this
}

// NewAvailableVLANWithDefaults instantiates a new AvailableVLAN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableVLANWithDefaults() *AvailableVLAN {
	this := AvailableVLAN{}
	return &this
}

// GetVid returns the Vid field value
func (o *AvailableVLAN) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *AvailableVLAN) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *AvailableVLAN) SetVid(v int32) {
	o.Vid = v
}

// GetGroup returns the Group field value
// If the value is explicit nil, the zero value for BriefVLANGroup will be returned
func (o *AvailableVLAN) GetGroup() BriefVLANGroup {
	if o == nil || o.Group.Get() == nil {
		var ret BriefVLANGroup
		return ret
	}

	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailableVLAN) GetGroupOk() (*BriefVLANGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// SetGroup sets field value
func (o *AvailableVLAN) SetGroup(v BriefVLANGroup) {
	o.Group.Set(&v)
}

func (o AvailableVLAN) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableVLAN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vid"] = o.Vid
	toSerialize["group"] = o.Group.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvailableVLAN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vid",
		"group",
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

	varAvailableVLAN := _AvailableVLAN{}

	err = json.Unmarshal(data, &varAvailableVLAN)

	if err != nil {
		return err
	}

	*o = AvailableVLAN(varAvailableVLAN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vid")
		delete(additionalProperties, "group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvailableVLAN struct {
	value *AvailableVLAN
	isSet bool
}

func (v NullableAvailableVLAN) Get() *AvailableVLAN {
	return v.value
}

func (v *NullableAvailableVLAN) Set(val *AvailableVLAN) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableVLAN) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableVLAN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableVLAN(val *AvailableVLAN) *NullableAvailableVLAN {
	return &NullableAvailableVLAN{value: val, isSet: true}
}

func (v NullableAvailableVLAN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableVLAN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
