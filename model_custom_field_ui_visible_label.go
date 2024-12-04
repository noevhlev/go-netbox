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

// CustomFieldUiVisibleLabel the model 'CustomFieldUiVisibleLabel'
type CustomFieldUiVisibleLabel string

// List of CustomField_ui_visible_label
const (
	CUSTOMFIELDUIVISIBLELABEL_ALWAYS CustomFieldUiVisibleLabel = "Always"
	CUSTOMFIELDUIVISIBLELABEL_IF_SET CustomFieldUiVisibleLabel = "If set"
	CUSTOMFIELDUIVISIBLELABEL_HIDDEN CustomFieldUiVisibleLabel = "Hidden"
)

// All allowed values of CustomFieldUiVisibleLabel enum
var AllowedCustomFieldUiVisibleLabelEnumValues = []CustomFieldUiVisibleLabel{
	"Always",
	"If set",
	"Hidden",
}

func (v *CustomFieldUiVisibleLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldUiVisibleLabel(value)
	for _, existing := range AllowedCustomFieldUiVisibleLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldUiVisibleLabel", value)
}

// NewCustomFieldUiVisibleLabelFromValue returns a pointer to a valid CustomFieldUiVisibleLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldUiVisibleLabelFromValue(v string) (*CustomFieldUiVisibleLabel, error) {
	ev := CustomFieldUiVisibleLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldUiVisibleLabel: valid values are %v", v, AllowedCustomFieldUiVisibleLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldUiVisibleLabel) IsValid() bool {
	for _, existing := range AllowedCustomFieldUiVisibleLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomField_ui_visible_label value
func (v CustomFieldUiVisibleLabel) Ptr() *CustomFieldUiVisibleLabel {
	return &v
}

type NullableCustomFieldUiVisibleLabel struct {
	value *CustomFieldUiVisibleLabel
	isSet bool
}

func (v NullableCustomFieldUiVisibleLabel) Get() *CustomFieldUiVisibleLabel {
	return v.value
}

func (v *NullableCustomFieldUiVisibleLabel) Set(val *CustomFieldUiVisibleLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldUiVisibleLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldUiVisibleLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldUiVisibleLabel(val *CustomFieldUiVisibleLabel) *NullableCustomFieldUiVisibleLabel {
	return &NullableCustomFieldUiVisibleLabel{value: val, isSet: true}
}

func (v NullableCustomFieldUiVisibleLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldUiVisibleLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
