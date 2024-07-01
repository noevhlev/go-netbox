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

// checks if the NestedFHRPGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedFHRPGroupRequest{}

// NestedFHRPGroupRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedFHRPGroupRequest struct {
	Protocol             FHRPGroupProtocol `json:"protocol"`
	GroupId              int32             `json:"group_id"`
	AdditionalProperties map[string]interface{}
}

type _NestedFHRPGroupRequest NestedFHRPGroupRequest

// NewNestedFHRPGroupRequest instantiates a new NestedFHRPGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedFHRPGroupRequest(protocol FHRPGroupProtocol, groupId int32) *NestedFHRPGroupRequest {
	this := NestedFHRPGroupRequest{}
	this.Protocol = protocol
	this.GroupId = groupId
	return &this
}

// NewNestedFHRPGroupRequestWithDefaults instantiates a new NestedFHRPGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedFHRPGroupRequestWithDefaults() *NestedFHRPGroupRequest {
	this := NestedFHRPGroupRequest{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *NestedFHRPGroupRequest) GetProtocol() FHRPGroupProtocol {
	if o == nil {
		var ret FHRPGroupProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroupRequest) GetProtocolOk() (*FHRPGroupProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NestedFHRPGroupRequest) SetProtocol(v FHRPGroupProtocol) {
	o.Protocol = v
}

// GetGroupId returns the GroupId field value
func (o *NestedFHRPGroupRequest) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroupRequest) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *NestedFHRPGroupRequest) SetGroupId(v int32) {
	o.GroupId = v
}

func (o NestedFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedFHRPGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["group_id"] = o.GroupId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedFHRPGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"protocol",
		"group_id",
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

	varNestedFHRPGroupRequest := _NestedFHRPGroupRequest{}

	err = json.Unmarshal(data, &varNestedFHRPGroupRequest)

	if err != nil {
		return err
	}

	*o = NestedFHRPGroupRequest(varNestedFHRPGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "group_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedFHRPGroupRequest struct {
	value *NestedFHRPGroupRequest
	isSet bool
}

func (v NullableNestedFHRPGroupRequest) Get() *NestedFHRPGroupRequest {
	return v.value
}

func (v *NullableNestedFHRPGroupRequest) Set(val *NestedFHRPGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedFHRPGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedFHRPGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedFHRPGroupRequest(val *NestedFHRPGroupRequest) *NullableNestedFHRPGroupRequest {
	return &NullableNestedFHRPGroupRequest{value: val, isSet: true}
}

func (v NullableNestedFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedFHRPGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
