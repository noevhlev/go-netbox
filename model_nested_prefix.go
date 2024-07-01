/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the NestedPrefix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedPrefix{}

// NestedPrefix Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedPrefix struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Family  int32  `json:"family"`
	// IPv4 or IPv6 network with mask
	Prefix               string `json:"prefix"`
	Depth                int32  `json:"_depth"`
	AdditionalProperties map[string]interface{}
}

type _NestedPrefix NestedPrefix

// NewNestedPrefix instantiates a new NestedPrefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedPrefix(id int32, url string, display string, family int32, prefix string, depth int32) *NestedPrefix {
	this := NestedPrefix{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Family = family
	this.Prefix = prefix
	this.Depth = depth
	return &this
}

// NewNestedPrefixWithDefaults instantiates a new NestedPrefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedPrefixWithDefaults() *NestedPrefix {
	this := NestedPrefix{}
	return &this
}

// GetId returns the Id field value
func (o *NestedPrefix) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedPrefix) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedPrefix) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedPrefix) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedPrefix) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedPrefix) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedPrefix) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedPrefix) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedPrefix) SetDisplay(v string) {
	o.Display = v
}

// GetFamily returns the Family field value
func (o *NestedPrefix) GetFamily() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *NestedPrefix) GetFamilyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *NestedPrefix) SetFamily(v int32) {
	o.Family = v
}

// GetPrefix returns the Prefix field value
func (o *NestedPrefix) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *NestedPrefix) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *NestedPrefix) SetPrefix(v string) {
	o.Prefix = v
}

// GetDepth returns the Depth field value
func (o *NestedPrefix) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *NestedPrefix) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *NestedPrefix) SetDepth(v int32) {
	o.Depth = v
}

func (o NestedPrefix) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedPrefix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["family"] = o.Family
	toSerialize["prefix"] = o.Prefix
	toSerialize["_depth"] = o.Depth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedPrefix) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"family",
		"prefix",
		"_depth",
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

	varNestedPrefix := _NestedPrefix{}

	err = json.Unmarshal(data, &varNestedPrefix)

	if err != nil {
		return err
	}

	*o = NestedPrefix(varNestedPrefix)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "family")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "_depth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedPrefix struct {
	value *NestedPrefix
	isSet bool
}

func (v NullableNestedPrefix) Get() *NestedPrefix {
	return v.value
}

func (v *NullableNestedPrefix) Set(val *NestedPrefix) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedPrefix) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedPrefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedPrefix(val *NestedPrefix) *NullableNestedPrefix {
	return &NullableNestedPrefix{value: val, isSet: true}
}

func (v NullableNestedPrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedPrefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
