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

// CustomFieldFilterLogicValue * `disabled` - Disabled * `loose` - Loose * `exact` - Exact
type CustomFieldFilterLogicValue string

// List of CustomField_filter_logic_value
const (
	CUSTOMFIELDFILTERLOGICVALUE_DISABLED CustomFieldFilterLogicValue = "disabled"
	CUSTOMFIELDFILTERLOGICVALUE_LOOSE    CustomFieldFilterLogicValue = "loose"
	CUSTOMFIELDFILTERLOGICVALUE_EXACT    CustomFieldFilterLogicValue = "exact"
)

// All allowed values of CustomFieldFilterLogicValue enum
var AllowedCustomFieldFilterLogicValueEnumValues = []CustomFieldFilterLogicValue{
	"disabled",
	"loose",
	"exact",
}

func (v *CustomFieldFilterLogicValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldFilterLogicValue(value)
	for _, existing := range AllowedCustomFieldFilterLogicValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldFilterLogicValue", value)
}

// NewCustomFieldFilterLogicValueFromValue returns a pointer to a valid CustomFieldFilterLogicValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldFilterLogicValueFromValue(v string) (*CustomFieldFilterLogicValue, error) {
	ev := CustomFieldFilterLogicValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldFilterLogicValue: valid values are %v", v, AllowedCustomFieldFilterLogicValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldFilterLogicValue) IsValid() bool {
	for _, existing := range AllowedCustomFieldFilterLogicValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomField_filter_logic_value value
func (v CustomFieldFilterLogicValue) Ptr() *CustomFieldFilterLogicValue {
	return &v
}

type NullableCustomFieldFilterLogicValue struct {
	value *CustomFieldFilterLogicValue
	isSet bool
}

func (v NullableCustomFieldFilterLogicValue) Get() *CustomFieldFilterLogicValue {
	return v.value
}

func (v *NullableCustomFieldFilterLogicValue) Set(val *CustomFieldFilterLogicValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldFilterLogicValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldFilterLogicValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldFilterLogicValue(val *CustomFieldFilterLogicValue) *NullableCustomFieldFilterLogicValue {
	return &NullableCustomFieldFilterLogicValue{value: val, isSet: true}
}

func (v NullableCustomFieldFilterLogicValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldFilterLogicValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
