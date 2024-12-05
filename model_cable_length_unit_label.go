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

// CableLengthUnitLabel the model 'CableLengthUnitLabel'
type CableLengthUnitLabel string

// List of Cable_length_unit_label
const (
	CABLELENGTHUNITLABEL_KILOMETERS  CableLengthUnitLabel = "Kilometers"
	CABLELENGTHUNITLABEL_METERS      CableLengthUnitLabel = "Meters"
	CABLELENGTHUNITLABEL_CENTIMETERS CableLengthUnitLabel = "Centimeters"
	CABLELENGTHUNITLABEL_MILES       CableLengthUnitLabel = "Miles"
	CABLELENGTHUNITLABEL_FEET        CableLengthUnitLabel = "Feet"
	CABLELENGTHUNITLABEL_INCHES      CableLengthUnitLabel = "Inches"
)

// All allowed values of CableLengthUnitLabel enum
var AllowedCableLengthUnitLabelEnumValues = []CableLengthUnitLabel{
	"Kilometers",
	"Meters",
	"Centimeters",
	"Miles",
	"Feet",
	"Inches",
}

func (v *CableLengthUnitLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CableLengthUnitLabel(value)
	for _, existing := range AllowedCableLengthUnitLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CableLengthUnitLabel", value)
}

// NewCableLengthUnitLabelFromValue returns a pointer to a valid CableLengthUnitLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableLengthUnitLabelFromValue(v string) (*CableLengthUnitLabel, error) {
	ev := CableLengthUnitLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CableLengthUnitLabel: valid values are %v", v, AllowedCableLengthUnitLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableLengthUnitLabel) IsValid() bool {
	for _, existing := range AllowedCableLengthUnitLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cable_length_unit_label value
func (v CableLengthUnitLabel) Ptr() *CableLengthUnitLabel {
	return &v
}

type NullableCableLengthUnitLabel struct {
	value *CableLengthUnitLabel
	isSet bool
}

func (v NullableCableLengthUnitLabel) Get() *CableLengthUnitLabel {
	return v.value
}

func (v *NullableCableLengthUnitLabel) Set(val *CableLengthUnitLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCableLengthUnitLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCableLengthUnitLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableLengthUnitLabel(val *CableLengthUnitLabel) *NullableCableLengthUnitLabel {
	return &NullableCableLengthUnitLabel{value: val, isSet: true}
}

func (v NullableCableLengthUnitLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableLengthUnitLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
