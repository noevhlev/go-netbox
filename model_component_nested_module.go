/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.8 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the ComponentNestedModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentNestedModule{}

// ComponentNestedModule Used by device component serializers.
type ComponentNestedModule struct {
	Id                   int32                 `json:"id"`
	Url                  string                `json:"url"`
	Display              string                `json:"display"`
	Device               int32                 `json:"device"`
	ModuleBay            ModuleNestedModuleBay `json:"module_bay"`
	AdditionalProperties map[string]interface{}
}

type _ComponentNestedModule ComponentNestedModule

// NewComponentNestedModule instantiates a new ComponentNestedModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentNestedModule(id int32, url string, display string, device int32, moduleBay ModuleNestedModuleBay) *ComponentNestedModule {
	this := ComponentNestedModule{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Device = device
	this.ModuleBay = moduleBay
	return &this
}

// NewComponentNestedModuleWithDefaults instantiates a new ComponentNestedModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentNestedModuleWithDefaults() *ComponentNestedModule {
	this := ComponentNestedModule{}
	return &this
}

// GetId returns the Id field value
func (o *ComponentNestedModule) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ComponentNestedModule) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ComponentNestedModule) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ComponentNestedModule) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ComponentNestedModule) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ComponentNestedModule) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ComponentNestedModule) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ComponentNestedModule) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ComponentNestedModule) SetDisplay(v string) {
	o.Display = v
}

// GetDevice returns the Device field value
func (o *ComponentNestedModule) GetDevice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *ComponentNestedModule) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *ComponentNestedModule) SetDevice(v int32) {
	o.Device = v
}

// GetModuleBay returns the ModuleBay field value
func (o *ComponentNestedModule) GetModuleBay() ModuleNestedModuleBay {
	if o == nil {
		var ret ModuleNestedModuleBay
		return ret
	}

	return o.ModuleBay
}

// GetModuleBayOk returns a tuple with the ModuleBay field value
// and a boolean to check if the value has been set.
func (o *ComponentNestedModule) GetModuleBayOk() (*ModuleNestedModuleBay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleBay, true
}

// SetModuleBay sets field value
func (o *ComponentNestedModule) SetModuleBay(v ModuleNestedModuleBay) {
	o.ModuleBay = v
}

func (o ComponentNestedModule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentNestedModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["device"] = o.Device
	// skip: module_bay is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ComponentNestedModule) UnmarshalJSON(bytes []byte) (err error) {
	varComponentNestedModule := _ComponentNestedModule{}

	if err = json.Unmarshal(bytes, &varComponentNestedModule); err == nil {
		*o = ComponentNestedModule(varComponentNestedModule)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module_bay")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComponentNestedModule struct {
	value *ComponentNestedModule
	isSet bool
}

func (v NullableComponentNestedModule) Get() *ComponentNestedModule {
	return v.value
}

func (v *NullableComponentNestedModule) Set(val *ComponentNestedModule) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentNestedModule) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentNestedModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentNestedModule(val *ComponentNestedModule) *NullableComponentNestedModule {
	return &NullableComponentNestedModule{value: val, isSet: true}
}

func (v NullableComponentNestedModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentNestedModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}