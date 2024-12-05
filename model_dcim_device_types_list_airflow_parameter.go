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

// DcimDeviceTypesListAirflowParameter the model 'DcimDeviceTypesListAirflowParameter'
type DcimDeviceTypesListAirflowParameter string

// List of dcim_device_types_list_airflow_parameter
const (
	DCIMDEVICETYPESLISTAIRFLOWPARAMETER_FRONT_TO_REAR DcimDeviceTypesListAirflowParameter = "front-to-rear"
	DCIMDEVICETYPESLISTAIRFLOWPARAMETER_LEFT_TO_RIGHT DcimDeviceTypesListAirflowParameter = "left-to-right"
	DCIMDEVICETYPESLISTAIRFLOWPARAMETER_MIXED         DcimDeviceTypesListAirflowParameter = "mixed"
	DCIMDEVICETYPESLISTAIRFLOWPARAMETER_PASSIVE       DcimDeviceTypesListAirflowParameter = "passive"
	DCIMDEVICETYPESLISTAIRFLOWPARAMETER_REAR_TO_FRONT DcimDeviceTypesListAirflowParameter = "rear-to-front"
	DCIMDEVICETYPESLISTAIRFLOWPARAMETER_RIGHT_TO_LEFT DcimDeviceTypesListAirflowParameter = "right-to-left"
	DCIMDEVICETYPESLISTAIRFLOWPARAMETER_SIDE_TO_REAR  DcimDeviceTypesListAirflowParameter = "side-to-rear"
)

// All allowed values of DcimDeviceTypesListAirflowParameter enum
var AllowedDcimDeviceTypesListAirflowParameterEnumValues = []DcimDeviceTypesListAirflowParameter{
	"front-to-rear",
	"left-to-right",
	"mixed",
	"passive",
	"rear-to-front",
	"right-to-left",
	"side-to-rear",
}

func (v *DcimDeviceTypesListAirflowParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimDeviceTypesListAirflowParameter(value)
	for _, existing := range AllowedDcimDeviceTypesListAirflowParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimDeviceTypesListAirflowParameter", value)
}

// NewDcimDeviceTypesListAirflowParameterFromValue returns a pointer to a valid DcimDeviceTypesListAirflowParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimDeviceTypesListAirflowParameterFromValue(v string) (*DcimDeviceTypesListAirflowParameter, error) {
	ev := DcimDeviceTypesListAirflowParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimDeviceTypesListAirflowParameter: valid values are %v", v, AllowedDcimDeviceTypesListAirflowParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimDeviceTypesListAirflowParameter) IsValid() bool {
	for _, existing := range AllowedDcimDeviceTypesListAirflowParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_device_types_list_airflow_parameter value
func (v DcimDeviceTypesListAirflowParameter) Ptr() *DcimDeviceTypesListAirflowParameter {
	return &v
}

type NullableDcimDeviceTypesListAirflowParameter struct {
	value *DcimDeviceTypesListAirflowParameter
	isSet bool
}

func (v NullableDcimDeviceTypesListAirflowParameter) Get() *DcimDeviceTypesListAirflowParameter {
	return v.value
}

func (v *NullableDcimDeviceTypesListAirflowParameter) Set(val *DcimDeviceTypesListAirflowParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimDeviceTypesListAirflowParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimDeviceTypesListAirflowParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimDeviceTypesListAirflowParameter(val *DcimDeviceTypesListAirflowParameter) *NullableDcimDeviceTypesListAirflowParameter {
	return &NullableDcimDeviceTypesListAirflowParameter{value: val, isSet: true}
}

func (v NullableDcimDeviceTypesListAirflowParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimDeviceTypesListAirflowParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
