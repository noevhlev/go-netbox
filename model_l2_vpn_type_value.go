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

// L2VPNTypeValue * `vpws` - VPWS * `vpls` - VPLS * `vxlan` - VXLAN * `vxlan-evpn` - VXLAN-EVPN * `mpls-evpn` - MPLS EVPN * `pbb-evpn` - PBB EVPN * `epl` - EPL * `evpl` - EVPL * `ep-lan` - Ethernet Private LAN * `evp-lan` - Ethernet Virtual Private LAN * `ep-tree` - Ethernet Private Tree * `evp-tree` - Ethernet Virtual Private Tree
type L2VPNTypeValue string

// List of L2VPN_type_value
const (
	L2VPNTYPEVALUE_VPWS       L2VPNTypeValue = "vpws"
	L2VPNTYPEVALUE_VPLS       L2VPNTypeValue = "vpls"
	L2VPNTYPEVALUE_VXLAN      L2VPNTypeValue = "vxlan"
	L2VPNTYPEVALUE_VXLAN_EVPN L2VPNTypeValue = "vxlan-evpn"
	L2VPNTYPEVALUE_MPLS_EVPN  L2VPNTypeValue = "mpls-evpn"
	L2VPNTYPEVALUE_PBB_EVPN   L2VPNTypeValue = "pbb-evpn"
	L2VPNTYPEVALUE_EPL        L2VPNTypeValue = "epl"
	L2VPNTYPEVALUE_EVPL       L2VPNTypeValue = "evpl"
	L2VPNTYPEVALUE_EP_LAN     L2VPNTypeValue = "ep-lan"
	L2VPNTYPEVALUE_EVP_LAN    L2VPNTypeValue = "evp-lan"
	L2VPNTYPEVALUE_EP_TREE    L2VPNTypeValue = "ep-tree"
	L2VPNTYPEVALUE_EVP_TREE   L2VPNTypeValue = "evp-tree"
)

// All allowed values of L2VPNTypeValue enum
var AllowedL2VPNTypeValueEnumValues = []L2VPNTypeValue{
	"vpws",
	"vpls",
	"vxlan",
	"vxlan-evpn",
	"mpls-evpn",
	"pbb-evpn",
	"epl",
	"evpl",
	"ep-lan",
	"evp-lan",
	"ep-tree",
	"evp-tree",
}

func (v *L2VPNTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := L2VPNTypeValue(value)
	for _, existing := range AllowedL2VPNTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid L2VPNTypeValue", value)
}

// NewL2VPNTypeValueFromValue returns a pointer to a valid L2VPNTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewL2VPNTypeValueFromValue(v string) (*L2VPNTypeValue, error) {
	ev := L2VPNTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for L2VPNTypeValue: valid values are %v", v, AllowedL2VPNTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v L2VPNTypeValue) IsValid() bool {
	for _, existing := range AllowedL2VPNTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to L2VPN_type_value value
func (v L2VPNTypeValue) Ptr() *L2VPNTypeValue {
	return &v
}

type NullableL2VPNTypeValue struct {
	value *L2VPNTypeValue
	isSet bool
}

func (v NullableL2VPNTypeValue) Get() *L2VPNTypeValue {
	return v.value
}

func (v *NullableL2VPNTypeValue) Set(val *L2VPNTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableL2VPNTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableL2VPNTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2VPNTypeValue(val *L2VPNTypeValue) *NullableL2VPNTypeValue {
	return &NullableL2VPNTypeValue{value: val, isSet: true}
}

func (v NullableL2VPNTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2VPNTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
