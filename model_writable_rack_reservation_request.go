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

// checks if the WritableRackReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableRackReservationRequest{}

// WritableRackReservationRequest Adds support for custom fields and tags.
type WritableRackReservationRequest struct {
	Rack                 int32                  `json:"rack"`
	Units                []int32                `json:"units"`
	User                 int32                  `json:"user"`
	Tenant               NullableInt32          `json:"tenant,omitempty"`
	Description          string                 `json:"description"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableRackReservationRequest WritableRackReservationRequest

// NewWritableRackReservationRequest instantiates a new WritableRackReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableRackReservationRequest(rack int32, units []int32, user int32, description string) *WritableRackReservationRequest {
	this := WritableRackReservationRequest{}
	this.Rack = rack
	this.Units = units
	this.User = user
	this.Description = description
	return &this
}

// NewWritableRackReservationRequestWithDefaults instantiates a new WritableRackReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableRackReservationRequestWithDefaults() *WritableRackReservationRequest {
	this := WritableRackReservationRequest{}
	return &this
}

// GetRack returns the Rack field value
func (o *WritableRackReservationRequest) GetRack() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rack
}

// GetRackOk returns a tuple with the Rack field value
// and a boolean to check if the value has been set.
func (o *WritableRackReservationRequest) GetRackOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rack, true
}

// SetRack sets field value
func (o *WritableRackReservationRequest) SetRack(v int32) {
	o.Rack = v
}

// GetUnits returns the Units field value
func (o *WritableRackReservationRequest) GetUnits() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *WritableRackReservationRequest) GetUnitsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Units, true
}

// SetUnits sets field value
func (o *WritableRackReservationRequest) SetUnits(v []int32) {
	o.Units = v
}

// GetUser returns the User field value
func (o *WritableRackReservationRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *WritableRackReservationRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *WritableRackReservationRequest) SetUser(v int32) {
	o.User = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackReservationRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackReservationRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableRackReservationRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *WritableRackReservationRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableRackReservationRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableRackReservationRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value
func (o *WritableRackReservationRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WritableRackReservationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WritableRackReservationRequest) SetDescription(v string) {
	o.Description = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableRackReservationRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackReservationRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableRackReservationRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableRackReservationRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableRackReservationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackReservationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableRackReservationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableRackReservationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableRackReservationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackReservationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableRackReservationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableRackReservationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableRackReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableRackReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rack"] = o.Rack
	toSerialize["units"] = o.Units
	toSerialize["user"] = o.User
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableRackReservationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rack",
		"units",
		"user",
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

	varWritableRackReservationRequest := _WritableRackReservationRequest{}

	err = json.Unmarshal(data, &varWritableRackReservationRequest)

	if err != nil {
		return err
	}

	*o = WritableRackReservationRequest(varWritableRackReservationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rack")
		delete(additionalProperties, "units")
		delete(additionalProperties, "user")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableRackReservationRequest struct {
	value *WritableRackReservationRequest
	isSet bool
}

func (v NullableWritableRackReservationRequest) Get() *WritableRackReservationRequest {
	return v.value
}

func (v *NullableWritableRackReservationRequest) Set(val *WritableRackReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableRackReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableRackReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableRackReservationRequest(val *WritableRackReservationRequest) *NullableWritableRackReservationRequest {
	return &NullableWritableRackReservationRequest{value: val, isSet: true}
}

func (v NullableWritableRackReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableRackReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
