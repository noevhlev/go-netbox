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

// PowerFeedPhaseLabel the model 'PowerFeedPhaseLabel'
type PowerFeedPhaseLabel string

// List of PowerFeed_phase_label
const (
	POWERFEEDPHASELABEL_SINGLE_PHASE PowerFeedPhaseLabel = "Single phase"
	POWERFEEDPHASELABEL_THREE_PHASE  PowerFeedPhaseLabel = "Three-phase"
)

// All allowed values of PowerFeedPhaseLabel enum
var AllowedPowerFeedPhaseLabelEnumValues = []PowerFeedPhaseLabel{
	"Single phase",
	"Three-phase",
}

func (v *PowerFeedPhaseLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerFeedPhaseLabel(value)
	for _, existing := range AllowedPowerFeedPhaseLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerFeedPhaseLabel", value)
}

// NewPowerFeedPhaseLabelFromValue returns a pointer to a valid PowerFeedPhaseLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerFeedPhaseLabelFromValue(v string) (*PowerFeedPhaseLabel, error) {
	ev := PowerFeedPhaseLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerFeedPhaseLabel: valid values are %v", v, AllowedPowerFeedPhaseLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerFeedPhaseLabel) IsValid() bool {
	for _, existing := range AllowedPowerFeedPhaseLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerFeed_phase_label value
func (v PowerFeedPhaseLabel) Ptr() *PowerFeedPhaseLabel {
	return &v
}

type NullablePowerFeedPhaseLabel struct {
	value *PowerFeedPhaseLabel
	isSet bool
}

func (v NullablePowerFeedPhaseLabel) Get() *PowerFeedPhaseLabel {
	return v.value
}

func (v *NullablePowerFeedPhaseLabel) Set(val *PowerFeedPhaseLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedPhaseLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedPhaseLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedPhaseLabel(val *PowerFeedPhaseLabel) *NullablePowerFeedPhaseLabel {
	return &NullablePowerFeedPhaseLabel{value: val, isSet: true}
}

func (v NullablePowerFeedPhaseLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedPhaseLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
