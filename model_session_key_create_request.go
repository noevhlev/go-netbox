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

// checks if the SessionKeyCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionKeyCreateRequest{}

// SessionKeyCreateRequest struct for SessionKeyCreateRequest
type SessionKeyCreateRequest struct {
	PreserveKey          *bool  `json:"preserve_key,omitempty"`
	PrivateKey           string `json:"private_key"`
	AdditionalProperties map[string]interface{}
}

type _SessionKeyCreateRequest SessionKeyCreateRequest

// NewSessionKeyCreateRequest instantiates a new SessionKeyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionKeyCreateRequest(privateKey string) *SessionKeyCreateRequest {
	this := SessionKeyCreateRequest{}
	var preserveKey bool = false
	this.PreserveKey = &preserveKey
	this.PrivateKey = privateKey
	return &this
}

// NewSessionKeyCreateRequestWithDefaults instantiates a new SessionKeyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionKeyCreateRequestWithDefaults() *SessionKeyCreateRequest {
	this := SessionKeyCreateRequest{}
	var preserveKey bool = false
	this.PreserveKey = &preserveKey
	return &this
}

// GetPreserveKey returns the PreserveKey field value if set, zero value otherwise.
func (o *SessionKeyCreateRequest) GetPreserveKey() bool {
	if o == nil || IsNil(o.PreserveKey) {
		var ret bool
		return ret
	}
	return *o.PreserveKey
}

// GetPreserveKeyOk returns a tuple with the PreserveKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionKeyCreateRequest) GetPreserveKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.PreserveKey) {
		return nil, false
	}
	return o.PreserveKey, true
}

// HasPreserveKey returns a boolean if a field has been set.
func (o *SessionKeyCreateRequest) HasPreserveKey() bool {
	if o != nil && !IsNil(o.PreserveKey) {
		return true
	}

	return false
}

// SetPreserveKey gets a reference to the given bool and assigns it to the PreserveKey field.
func (o *SessionKeyCreateRequest) SetPreserveKey(v bool) {
	o.PreserveKey = &v
}

// GetPrivateKey returns the PrivateKey field value
func (o *SessionKeyCreateRequest) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *SessionKeyCreateRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *SessionKeyCreateRequest) SetPrivateKey(v string) {
	o.PrivateKey = v
}

func (o SessionKeyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionKeyCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreserveKey) {
		toSerialize["preserve_key"] = o.PreserveKey
	}
	toSerialize["private_key"] = o.PrivateKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionKeyCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"private_key",
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

	varSessionKeyCreateRequest := _SessionKeyCreateRequest{}

	err = json.Unmarshal(data, &varSessionKeyCreateRequest)

	if err != nil {
		return err
	}

	*o = SessionKeyCreateRequest(varSessionKeyCreateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "preserve_key")
		delete(additionalProperties, "private_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionKeyCreateRequest struct {
	value *SessionKeyCreateRequest
	isSet bool
}

func (v NullableSessionKeyCreateRequest) Get() *SessionKeyCreateRequest {
	return v.value
}

func (v *NullableSessionKeyCreateRequest) Set(val *SessionKeyCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionKeyCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionKeyCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionKeyCreateRequest(val *SessionKeyCreateRequest) *NullableSessionKeyCreateRequest {
	return &NullableSessionKeyCreateRequest{value: val, isSet: true}
}

func (v NullableSessionKeyCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionKeyCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
