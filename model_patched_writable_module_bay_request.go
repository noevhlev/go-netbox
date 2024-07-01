/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableModuleBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableModuleBayRequest{}

// PatchedWritableModuleBayRequest Adds support for custom fields and tags.
type PatchedWritableModuleBayRequest struct {
	Device          *int32  `json:"device,omitempty"`
	Name            *string `json:"name,omitempty"`
	InstalledModule *int32  `json:"installed_module,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// Identifier to reference when renaming installed components
	Position             *string                `json:"position,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableModuleBayRequest PatchedWritableModuleBayRequest

// NewPatchedWritableModuleBayRequest instantiates a new PatchedWritableModuleBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableModuleBayRequest() *PatchedWritableModuleBayRequest {
	this := PatchedWritableModuleBayRequest{}
	return &this
}

// NewPatchedWritableModuleBayRequestWithDefaults instantiates a new PatchedWritableModuleBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableModuleBayRequestWithDefaults() *PatchedWritableModuleBayRequest {
	this := PatchedWritableModuleBayRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayRequest) GetDevice() int32 {
	if o == nil || IsNil(o.Device) {
		var ret int32
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayRequest) GetDeviceOk() (*int32, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given int32 and assigns it to the Device field.
func (o *PatchedWritableModuleBayRequest) SetDevice(v int32) {
	o.Device = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableModuleBayRequest) SetName(v string) {
	o.Name = &v
}

// GetInstalledModule returns the InstalledModule field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayRequest) GetInstalledModule() int32 {
	if o == nil || IsNil(o.InstalledModule) {
		var ret int32
		return ret
	}
	return *o.InstalledModule
}

// GetInstalledModuleOk returns a tuple with the InstalledModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayRequest) GetInstalledModuleOk() (*int32, bool) {
	if o == nil || IsNil(o.InstalledModule) {
		return nil, false
	}
	return o.InstalledModule, true
}

// HasInstalledModule returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayRequest) HasInstalledModule() bool {
	if o != nil && !IsNil(o.InstalledModule) {
		return true
	}

	return false
}

// SetInstalledModule gets a reference to the given int32 and assigns it to the InstalledModule field.
func (o *PatchedWritableModuleBayRequest) SetInstalledModule(v int32) {
	o.InstalledModule = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableModuleBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayRequest) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayRequest) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *PatchedWritableModuleBayRequest) SetPosition(v string) {
	o.Position = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableModuleBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableModuleBayRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableModuleBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableModuleBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableModuleBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.InstalledModule) {
		toSerialize["installed_module"] = o.InstalledModule
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *PatchedWritableModuleBayRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableModuleBayRequest := _PatchedWritableModuleBayRequest{}

	err = json.Unmarshal(data, &varPatchedWritableModuleBayRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableModuleBayRequest(varPatchedWritableModuleBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "name")
		delete(additionalProperties, "installed_module")
		delete(additionalProperties, "label")
		delete(additionalProperties, "position")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableModuleBayRequest struct {
	value *PatchedWritableModuleBayRequest
	isSet bool
}

func (v NullablePatchedWritableModuleBayRequest) Get() *PatchedWritableModuleBayRequest {
	return v.value
}

func (v *NullablePatchedWritableModuleBayRequest) Set(val *PatchedWritableModuleBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableModuleBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableModuleBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableModuleBayRequest(val *PatchedWritableModuleBayRequest) *NullablePatchedWritableModuleBayRequest {
	return &NullablePatchedWritableModuleBayRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableModuleBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableModuleBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
