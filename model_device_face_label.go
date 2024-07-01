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

// DeviceFaceLabel the model 'DeviceFaceLabel'
type DeviceFaceLabel string

// List of Device_face_label
const (
	DEVICEFACELABEL_FRONT DeviceFaceLabel = "Front"
	DEVICEFACELABEL_REAR  DeviceFaceLabel = "Rear"
)

// All allowed values of DeviceFaceLabel enum
var AllowedDeviceFaceLabelEnumValues = []DeviceFaceLabel{
	"Front",
	"Rear",
}

func (v *DeviceFaceLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceFaceLabel(value)
	for _, existing := range AllowedDeviceFaceLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceFaceLabel", value)
}

// NewDeviceFaceLabelFromValue returns a pointer to a valid DeviceFaceLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceFaceLabelFromValue(v string) (*DeviceFaceLabel, error) {
	ev := DeviceFaceLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceFaceLabel: valid values are %v", v, AllowedDeviceFaceLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceFaceLabel) IsValid() bool {
	for _, existing := range AllowedDeviceFaceLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Device_face_label value
func (v DeviceFaceLabel) Ptr() *DeviceFaceLabel {
	return &v
}

type NullableDeviceFaceLabel struct {
	value *DeviceFaceLabel
	isSet bool
}

func (v NullableDeviceFaceLabel) Get() *DeviceFaceLabel {
	return v.value
}

func (v *NullableDeviceFaceLabel) Set(val *DeviceFaceLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceFaceLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceFaceLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceFaceLabel(val *DeviceFaceLabel) *NullableDeviceFaceLabel {
	return &NullableDeviceFaceLabel{value: val, isSet: true}
}

func (v NullableDeviceFaceLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceFaceLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
