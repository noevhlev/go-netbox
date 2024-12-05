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

// ExtrasCustomFieldsListFilterLogicParameter the model 'ExtrasCustomFieldsListFilterLogicParameter'
type ExtrasCustomFieldsListFilterLogicParameter string

// List of extras_custom_fields_list_filter_logic_parameter
const (
	EXTRASCUSTOMFIELDSLISTFILTERLOGICPARAMETER_DISABLED ExtrasCustomFieldsListFilterLogicParameter = "disabled"
	EXTRASCUSTOMFIELDSLISTFILTERLOGICPARAMETER_EXACT    ExtrasCustomFieldsListFilterLogicParameter = "exact"
	EXTRASCUSTOMFIELDSLISTFILTERLOGICPARAMETER_LOOSE    ExtrasCustomFieldsListFilterLogicParameter = "loose"
)

// All allowed values of ExtrasCustomFieldsListFilterLogicParameter enum
var AllowedExtrasCustomFieldsListFilterLogicParameterEnumValues = []ExtrasCustomFieldsListFilterLogicParameter{
	"disabled",
	"exact",
	"loose",
}

func (v *ExtrasCustomFieldsListFilterLogicParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExtrasCustomFieldsListFilterLogicParameter(value)
	for _, existing := range AllowedExtrasCustomFieldsListFilterLogicParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExtrasCustomFieldsListFilterLogicParameter", value)
}

// NewExtrasCustomFieldsListFilterLogicParameterFromValue returns a pointer to a valid ExtrasCustomFieldsListFilterLogicParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExtrasCustomFieldsListFilterLogicParameterFromValue(v string) (*ExtrasCustomFieldsListFilterLogicParameter, error) {
	ev := ExtrasCustomFieldsListFilterLogicParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExtrasCustomFieldsListFilterLogicParameter: valid values are %v", v, AllowedExtrasCustomFieldsListFilterLogicParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExtrasCustomFieldsListFilterLogicParameter) IsValid() bool {
	for _, existing := range AllowedExtrasCustomFieldsListFilterLogicParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to extras_custom_fields_list_filter_logic_parameter value
func (v ExtrasCustomFieldsListFilterLogicParameter) Ptr() *ExtrasCustomFieldsListFilterLogicParameter {
	return &v
}

type NullableExtrasCustomFieldsListFilterLogicParameter struct {
	value *ExtrasCustomFieldsListFilterLogicParameter
	isSet bool
}

func (v NullableExtrasCustomFieldsListFilterLogicParameter) Get() *ExtrasCustomFieldsListFilterLogicParameter {
	return v.value
}

func (v *NullableExtrasCustomFieldsListFilterLogicParameter) Set(val *ExtrasCustomFieldsListFilterLogicParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableExtrasCustomFieldsListFilterLogicParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableExtrasCustomFieldsListFilterLogicParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtrasCustomFieldsListFilterLogicParameter(val *ExtrasCustomFieldsListFilterLogicParameter) *NullableExtrasCustomFieldsListFilterLogicParameter {
	return &NullableExtrasCustomFieldsListFilterLogicParameter{value: val, isSet: true}
}

func (v NullableExtrasCustomFieldsListFilterLogicParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtrasCustomFieldsListFilterLogicParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}