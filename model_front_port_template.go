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
	"time"
)

// checks if the FrontPortTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrontPortTemplate{}

// FrontPortTemplate Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type FrontPortTemplate struct {
	Id         int32                   `json:"id"`
	Url        string                  `json:"url"`
	Display    string                  `json:"display"`
	DeviceType NullableBriefDeviceType `json:"device_type,omitempty"`
	ModuleType NullableBriefModuleType `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label                *string               `json:"label,omitempty"`
	Type                 FrontPortType         `json:"type"`
	Color                *string               `json:"color,omitempty"`
	RearPort             BriefRearPortTemplate `json:"rear_port"`
	RearPortPosition     *int32                `json:"rear_port_position,omitempty"`
	Description          *string               `json:"description,omitempty"`
	Created              NullableTime          `json:"created"`
	LastUpdated          NullableTime          `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _FrontPortTemplate FrontPortTemplate

// NewFrontPortTemplate instantiates a new FrontPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontPortTemplate(id int32, url string, display string, name string, type_ FrontPortType, rearPort BriefRearPortTemplate, created NullableTime, lastUpdated NullableTime) *FrontPortTemplate {
	this := FrontPortTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Type = type_
	this.RearPort = rearPort
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewFrontPortTemplateWithDefaults instantiates a new FrontPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontPortTemplateWithDefaults() *FrontPortTemplate {
	this := FrontPortTemplate{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// GetId returns the Id field value
func (o *FrontPortTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FrontPortTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *FrontPortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FrontPortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *FrontPortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *FrontPortTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FrontPortTemplate) GetDeviceType() BriefDeviceType {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BriefDeviceType
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplate) GetDeviceTypeOk() (*BriefDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBriefDeviceType and assigns it to the DeviceType field.
func (o *FrontPortTemplate) SetDeviceType(v BriefDeviceType) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *FrontPortTemplate) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *FrontPortTemplate) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FrontPortTemplate) GetModuleType() BriefModuleType {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BriefModuleType
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplate) GetModuleTypeOk() (*BriefModuleType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBriefModuleType and assigns it to the ModuleType field.
func (o *FrontPortTemplate) SetModuleType(v BriefModuleType) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *FrontPortTemplate) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *FrontPortTemplate) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *FrontPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrontPortTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FrontPortTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *FrontPortTemplate) GetType() FrontPortType {
	if o == nil {
		var ret FrontPortType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetTypeOk() (*FrontPortType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FrontPortTemplate) SetType(v FrontPortType) {
	o.Type = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *FrontPortTemplate) SetColor(v string) {
	o.Color = &v
}

// GetRearPort returns the RearPort field value
func (o *FrontPortTemplate) GetRearPort() BriefRearPortTemplate {
	if o == nil {
		var ret BriefRearPortTemplate
		return ret
	}

	return o.RearPort
}

// GetRearPortOk returns a tuple with the RearPort field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetRearPortOk() (*BriefRearPortTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RearPort, true
}

// SetRearPort sets field value
func (o *FrontPortTemplate) SetRearPort(v BriefRearPortTemplate) {
	o.RearPort = v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetRearPortPosition() int32 {
	if o == nil || IsNil(o.RearPortPosition) {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.RearPortPosition) {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasRearPortPosition() bool {
	if o != nil && !IsNil(o.RearPortPosition) {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *FrontPortTemplate) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FrontPortTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FrontPortTemplate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *FrontPortTemplate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FrontPortTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *FrontPortTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o FrontPortTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrontPortTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	toSerialize["rear_port"] = o.RearPort
	if !IsNil(o.RearPortPosition) {
		toSerialize["rear_port_position"] = o.RearPortPosition
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrontPortTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"type",
		"rear_port",
		"created",
		"last_updated",
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

	varFrontPortTemplate := _FrontPortTemplate{}

	err = json.Unmarshal(data, &varFrontPortTemplate)

	if err != nil {
		return err
	}

	*o = FrontPortTemplate(varFrontPortTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "color")
		delete(additionalProperties, "rear_port")
		delete(additionalProperties, "rear_port_position")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFrontPortTemplate struct {
	value *FrontPortTemplate
	isSet bool
}

func (v NullableFrontPortTemplate) Get() *FrontPortTemplate {
	return v.value
}

func (v *NullableFrontPortTemplate) Set(val *FrontPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPortTemplate(val *FrontPortTemplate) *NullableFrontPortTemplate {
	return &NullableFrontPortTemplate{value: val, isSet: true}
}

func (v NullableFrontPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
