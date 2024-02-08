/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the FHRPGroupAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FHRPGroupAssignmentRequest{}

// FHRPGroupAssignmentRequest Adds support for custom fields and tags.
type FHRPGroupAssignmentRequest struct {
	Group                NestedFHRPGroupRequest `json:"group"`
	InterfaceType        string                 `json:"interface_type"`
	InterfaceId          int64                  `json:"interface_id"`
	Priority             int32                  `json:"priority"`
	AdditionalProperties map[string]interface{}
}

type _FHRPGroupAssignmentRequest FHRPGroupAssignmentRequest

// NewFHRPGroupAssignmentRequest instantiates a new FHRPGroupAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFHRPGroupAssignmentRequest(group NestedFHRPGroupRequest, interfaceType string, interfaceId int64, priority int32) *FHRPGroupAssignmentRequest {
	this := FHRPGroupAssignmentRequest{}
	this.Group = group
	this.InterfaceType = interfaceType
	this.InterfaceId = interfaceId
	this.Priority = priority
	return &this
}

// NewFHRPGroupAssignmentRequestWithDefaults instantiates a new FHRPGroupAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFHRPGroupAssignmentRequestWithDefaults() *FHRPGroupAssignmentRequest {
	this := FHRPGroupAssignmentRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *FHRPGroupAssignmentRequest) GetGroup() NestedFHRPGroupRequest {
	if o == nil {
		var ret NestedFHRPGroupRequest
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignmentRequest) GetGroupOk() (*NestedFHRPGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *FHRPGroupAssignmentRequest) SetGroup(v NestedFHRPGroupRequest) {
	o.Group = v
}

// GetInterfaceType returns the InterfaceType field value
func (o *FHRPGroupAssignmentRequest) GetInterfaceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignmentRequest) GetInterfaceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceType, true
}

// SetInterfaceType sets field value
func (o *FHRPGroupAssignmentRequest) SetInterfaceType(v string) {
	o.InterfaceType = v
}

// GetInterfaceId returns the InterfaceId field value
func (o *FHRPGroupAssignmentRequest) GetInterfaceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignmentRequest) GetInterfaceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceId, true
}

// SetInterfaceId sets field value
func (o *FHRPGroupAssignmentRequest) SetInterfaceId(v int64) {
	o.InterfaceId = v
}

// GetPriority returns the Priority field value
func (o *FHRPGroupAssignmentRequest) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignmentRequest) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *FHRPGroupAssignmentRequest) SetPriority(v int32) {
	o.Priority = v
}

func (o FHRPGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FHRPGroupAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["interface_type"] = o.InterfaceType
	toSerialize["interface_id"] = o.InterfaceId
	toSerialize["priority"] = o.Priority

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FHRPGroupAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
		"interface_type",
		"interface_id",
		"priority",
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

	varFHRPGroupAssignmentRequest := _FHRPGroupAssignmentRequest{}

	err = json.Unmarshal(data, &varFHRPGroupAssignmentRequest)

	if err != nil {
		return err
	}

	*o = FHRPGroupAssignmentRequest(varFHRPGroupAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "interface_type")
		delete(additionalProperties, "interface_id")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFHRPGroupAssignmentRequest struct {
	value *FHRPGroupAssignmentRequest
	isSet bool
}

func (v NullableFHRPGroupAssignmentRequest) Get() *FHRPGroupAssignmentRequest {
	return v.value
}

func (v *NullableFHRPGroupAssignmentRequest) Set(val *FHRPGroupAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFHRPGroupAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFHRPGroupAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFHRPGroupAssignmentRequest(val *FHRPGroupAssignmentRequest) *NullableFHRPGroupAssignmentRequest {
	return &NullableFHRPGroupAssignmentRequest{value: val, isSet: true}
}

func (v NullableFHRPGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFHRPGroupAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
