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

// AssetStatus Asset lifecycle status  * `stored` - Stored * `used` - Used * `retired` - Retired
type AssetStatus string

// List of Asset_status
const (
	ASSETSTATUS_STORED  AssetStatus = "stored"
	ASSETSTATUS_USED    AssetStatus = "used"
	ASSETSTATUS_RETIRED AssetStatus = "retired"
)

// All allowed values of AssetStatus enum
var AllowedAssetStatusEnumValues = []AssetStatus{
	"stored",
	"used",
	"retired",
}

func (v *AssetStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssetStatus(value)
	for _, existing := range AllowedAssetStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssetStatus", value)
}

// NewAssetStatusFromValue returns a pointer to a valid AssetStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetStatusFromValue(v string) (*AssetStatus, error) {
	ev := AssetStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetStatus: valid values are %v", v, AllowedAssetStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetStatus) IsValid() bool {
	for _, existing := range AllowedAssetStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Asset_status value
func (v AssetStatus) Ptr() *AssetStatus {
	return &v
}

type NullableAssetStatus struct {
	value *AssetStatus
	isSet bool
}

func (v NullableAssetStatus) Get() *AssetStatus {
	return v.value
}

func (v *NullableAssetStatus) Set(val *AssetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetStatus(val *AssetStatus) *NullableAssetStatus {
	return &NullableAssetStatus{value: val, isSet: true}
}

func (v NullableAssetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
