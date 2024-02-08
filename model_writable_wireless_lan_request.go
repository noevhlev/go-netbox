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

// checks if the WritableWirelessLANRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableWirelessLANRequest{}

// WritableWirelessLANRequest Adds support for custom fields and tags.
type WritableWirelessLANRequest struct {
	Ssid                 string                                   `json:"ssid"`
	Description          *string                                  `json:"description,omitempty"`
	Group                NullableInt32                            `json:"group,omitempty"`
	Status               *PatchedWritableWirelessLANRequestStatus `json:"status,omitempty"`
	Vlan                 NullableInt32                            `json:"vlan,omitempty"`
	Tenant               NullableInt32                            `json:"tenant,omitempty"`
	AuthType             *AuthenticationType1                     `json:"auth_type,omitempty"`
	AuthCipher           *AuthenticationCipher                    `json:"auth_cipher,omitempty"`
	AuthPsk              *string                                  `json:"auth_psk,omitempty"`
	Comments             *string                                  `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                       `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                   `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableWirelessLANRequest WritableWirelessLANRequest

// NewWritableWirelessLANRequest instantiates a new WritableWirelessLANRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableWirelessLANRequest(ssid string) *WritableWirelessLANRequest {
	this := WritableWirelessLANRequest{}
	this.Ssid = ssid
	return &this
}

// NewWritableWirelessLANRequestWithDefaults instantiates a new WritableWirelessLANRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableWirelessLANRequestWithDefaults() *WritableWirelessLANRequest {
	this := WritableWirelessLANRequest{}
	return &this
}

// GetSsid returns the Ssid field value
func (o *WritableWirelessLANRequest) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *WritableWirelessLANRequest) SetSsid(v string) {
	o.Ssid = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableWirelessLANRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableWirelessLANRequest) GetGroup() int32 {
	if o == nil || IsNil(o.Group.Get()) {
		var ret int32
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableWirelessLANRequest) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableInt32 and assigns it to the Group field.
func (o *WritableWirelessLANRequest) SetGroup(v int32) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *WritableWirelessLANRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *WritableWirelessLANRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetStatus() PatchedWritableWirelessLANRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableWirelessLANRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetStatusOk() (*PatchedWritableWirelessLANRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableWirelessLANRequestStatus and assigns it to the Status field.
func (o *WritableWirelessLANRequest) SetStatus(v PatchedWritableWirelessLANRequestStatus) {
	o.Status = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableWirelessLANRequest) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan.Get()) {
		var ret int32
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableWirelessLANRequest) GetVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableInt32 and assigns it to the Vlan field.
func (o *WritableWirelessLANRequest) SetVlan(v int32) {
	o.Vlan.Set(&v)
}

// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *WritableWirelessLANRequest) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *WritableWirelessLANRequest) UnsetVlan() {
	o.Vlan.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableWirelessLANRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableWirelessLANRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *WritableWirelessLANRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableWirelessLANRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableWirelessLANRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetAuthType() AuthenticationType1 {
	if o == nil || IsNil(o.AuthType) {
		var ret AuthenticationType1
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetAuthTypeOk() (*AuthenticationType1, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given AuthenticationType1 and assigns it to the AuthType field.
func (o *WritableWirelessLANRequest) SetAuthType(v AuthenticationType1) {
	o.AuthType = &v
}

// GetAuthCipher returns the AuthCipher field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetAuthCipher() AuthenticationCipher {
	if o == nil || IsNil(o.AuthCipher) {
		var ret AuthenticationCipher
		return ret
	}
	return *o.AuthCipher
}

// GetAuthCipherOk returns a tuple with the AuthCipher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetAuthCipherOk() (*AuthenticationCipher, bool) {
	if o == nil || IsNil(o.AuthCipher) {
		return nil, false
	}
	return o.AuthCipher, true
}

// HasAuthCipher returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasAuthCipher() bool {
	if o != nil && !IsNil(o.AuthCipher) {
		return true
	}

	return false
}

// SetAuthCipher gets a reference to the given AuthenticationCipher and assigns it to the AuthCipher field.
func (o *WritableWirelessLANRequest) SetAuthCipher(v AuthenticationCipher) {
	o.AuthCipher = &v
}

// GetAuthPsk returns the AuthPsk field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetAuthPsk() string {
	if o == nil || IsNil(o.AuthPsk) {
		var ret string
		return ret
	}
	return *o.AuthPsk
}

// GetAuthPskOk returns a tuple with the AuthPsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetAuthPskOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPsk) {
		return nil, false
	}
	return o.AuthPsk, true
}

// HasAuthPsk returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasAuthPsk() bool {
	if o != nil && !IsNil(o.AuthPsk) {
		return true
	}

	return false
}

// SetAuthPsk gets a reference to the given string and assigns it to the AuthPsk field.
func (o *WritableWirelessLANRequest) SetAuthPsk(v string) {
	o.AuthPsk = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableWirelessLANRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableWirelessLANRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableWirelessLANRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableWirelessLANRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableWirelessLANRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ssid"] = o.Ssid
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.AuthCipher) {
		toSerialize["auth_cipher"] = o.AuthCipher
	}
	if !IsNil(o.AuthPsk) {
		toSerialize["auth_psk"] = o.AuthPsk
	}
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

func (o *WritableWirelessLANRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ssid",
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

	varWritableWirelessLANRequest := _WritableWirelessLANRequest{}

	err = json.Unmarshal(data, &varWritableWirelessLANRequest)

	if err != nil {
		return err
	}

	*o = WritableWirelessLANRequest(varWritableWirelessLANRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "group")
		delete(additionalProperties, "status")
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "auth_cipher")
		delete(additionalProperties, "auth_psk")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableWirelessLANRequest struct {
	value *WritableWirelessLANRequest
	isSet bool
}

func (v NullableWritableWirelessLANRequest) Get() *WritableWirelessLANRequest {
	return v.value
}

func (v *NullableWritableWirelessLANRequest) Set(val *WritableWirelessLANRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableWirelessLANRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableWirelessLANRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableWirelessLANRequest(val *WritableWirelessLANRequest) *NullableWritableWirelessLANRequest {
	return &NullableWritableWirelessLANRequest{value: val, isSet: true}
}

func (v NullableWritableWirelessLANRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableWirelessLANRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
