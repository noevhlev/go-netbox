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

// DeviceStatusLabel the model 'DeviceStatusLabel'
type DeviceStatusLabel string

// List of Device_status_label
const (
	DEVICESTATUSLABEL_OFFLINE         DeviceStatusLabel = "Offline"
	DEVICESTATUSLABEL_ACTIVE          DeviceStatusLabel = "Active"
	DEVICESTATUSLABEL_PLANNED         DeviceStatusLabel = "Planned"
	DEVICESTATUSLABEL_STAGED          DeviceStatusLabel = "Staged"
	DEVICESTATUSLABEL_FAILED          DeviceStatusLabel = "Failed"
	DEVICESTATUSLABEL_INVENTORY       DeviceStatusLabel = "Inventory"
	DEVICESTATUSLABEL_DECOMMISSIONING DeviceStatusLabel = "Decommissioning"
	DEVICESTATUSLABEL_MAINTENANCE     DeviceStatusLabel = "Maintenance"
	DEVICESTATUSLABEL_TESTING         DeviceStatusLabel = "Testing"
)

// All allowed values of DeviceStatusLabel enum
var AllowedDeviceStatusLabelEnumValues = []DeviceStatusLabel{
	"Offline",
	"Active",
	"Planned",
	"Staged",
	"Failed",
	"Inventory",
	"Decommissioning",
	"Maintenance",
	"Testing",
}

func (v *DeviceStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceStatusLabel(value)
	for _, existing := range AllowedDeviceStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceStatusLabel", value)
}

// NewDeviceStatusLabelFromValue returns a pointer to a valid DeviceStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceStatusLabelFromValue(v string) (*DeviceStatusLabel, error) {
	ev := DeviceStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceStatusLabel: valid values are %v", v, AllowedDeviceStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceStatusLabel) IsValid() bool {
	for _, existing := range AllowedDeviceStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Device_status_label value
func (v DeviceStatusLabel) Ptr() *DeviceStatusLabel {
	return &v
}

type NullableDeviceStatusLabel struct {
	value *DeviceStatusLabel
	isSet bool
}

func (v NullableDeviceStatusLabel) Get() *DeviceStatusLabel {
	return v.value
}

func (v *NullableDeviceStatusLabel) Set(val *DeviceStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceStatusLabel(val *DeviceStatusLabel) *NullableDeviceStatusLabel {
	return &NullableDeviceStatusLabel{value: val, isSet: true}
}

func (v NullableDeviceStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
