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

// LocationStatusLabel the model 'LocationStatusLabel'
type LocationStatusLabel string

// List of Location_status_label
const (
	LOCATIONSTATUSLABEL_AVAILABLE       LocationStatusLabel = "Available"
	LOCATIONSTATUSLABEL_RESERVED        LocationStatusLabel = "Reserved"
	LOCATIONSTATUSLABEL_ACTIVE          LocationStatusLabel = "Active"
	LOCATIONSTATUSLABEL_DECOMMISSIONING LocationStatusLabel = "Decommissioning"
	LOCATIONSTATUSLABEL_RETIRED         LocationStatusLabel = "Retired"
)

// All allowed values of LocationStatusLabel enum
var AllowedLocationStatusLabelEnumValues = []LocationStatusLabel{
	"Available",
	"Reserved",
	"Active",
	"Decommissioning",
	"Retired",
}

func (v *LocationStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationStatusLabel(value)
	for _, existing := range AllowedLocationStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationStatusLabel", value)
}

// NewLocationStatusLabelFromValue returns a pointer to a valid LocationStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationStatusLabelFromValue(v string) (*LocationStatusLabel, error) {
	ev := LocationStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationStatusLabel: valid values are %v", v, AllowedLocationStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationStatusLabel) IsValid() bool {
	for _, existing := range AllowedLocationStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Location_status_label value
func (v LocationStatusLabel) Ptr() *LocationStatusLabel {
	return &v
}

type NullableLocationStatusLabel struct {
	value *LocationStatusLabel
	isSet bool
}

func (v NullableLocationStatusLabel) Get() *LocationStatusLabel {
	return v.value
}

func (v *NullableLocationStatusLabel) Set(val *LocationStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationStatusLabel(val *LocationStatusLabel) *NullableLocationStatusLabel {
	return &NullableLocationStatusLabel{value: val, isSet: true}
}

func (v NullableLocationStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
