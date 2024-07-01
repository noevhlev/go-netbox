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

// L2VPNTypeLabel the model 'L2VPNTypeLabel'
type L2VPNTypeLabel string

// List of L2VPN_type_label
const (
	L2VPNTYPELABEL_VPWS                          L2VPNTypeLabel = "VPWS"
	L2VPNTYPELABEL_VPLS                          L2VPNTypeLabel = "VPLS"
	L2VPNTYPELABEL_VXLAN                         L2VPNTypeLabel = "VXLAN"
	L2VPNTYPELABEL_VXLAN_EVPN                    L2VPNTypeLabel = "VXLAN-EVPN"
	L2VPNTYPELABEL_MPLS_EVPN                     L2VPNTypeLabel = "MPLS EVPN"
	L2VPNTYPELABEL_PBB_EVPN                      L2VPNTypeLabel = "PBB EVPN"
	L2VPNTYPELABEL_EPL                           L2VPNTypeLabel = "EPL"
	L2VPNTYPELABEL_EVPL                          L2VPNTypeLabel = "EVPL"
	L2VPNTYPELABEL_ETHERNET_PRIVATE_LAN          L2VPNTypeLabel = "Ethernet Private LAN"
	L2VPNTYPELABEL_ETHERNET_VIRTUAL_PRIVATE_LAN  L2VPNTypeLabel = "Ethernet Virtual Private LAN"
	L2VPNTYPELABEL_ETHERNET_PRIVATE_TREE         L2VPNTypeLabel = "Ethernet Private Tree"
	L2VPNTYPELABEL_ETHERNET_VIRTUAL_PRIVATE_TREE L2VPNTypeLabel = "Ethernet Virtual Private Tree"
)

// All allowed values of L2VPNTypeLabel enum
var AllowedL2VPNTypeLabelEnumValues = []L2VPNTypeLabel{
	"VPWS",
	"VPLS",
	"VXLAN",
	"VXLAN-EVPN",
	"MPLS EVPN",
	"PBB EVPN",
	"EPL",
	"EVPL",
	"Ethernet Private LAN",
	"Ethernet Virtual Private LAN",
	"Ethernet Private Tree",
	"Ethernet Virtual Private Tree",
}

func (v *L2VPNTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := L2VPNTypeLabel(value)
	for _, existing := range AllowedL2VPNTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid L2VPNTypeLabel", value)
}

// NewL2VPNTypeLabelFromValue returns a pointer to a valid L2VPNTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewL2VPNTypeLabelFromValue(v string) (*L2VPNTypeLabel, error) {
	ev := L2VPNTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for L2VPNTypeLabel: valid values are %v", v, AllowedL2VPNTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v L2VPNTypeLabel) IsValid() bool {
	for _, existing := range AllowedL2VPNTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to L2VPN_type_label value
func (v L2VPNTypeLabel) Ptr() *L2VPNTypeLabel {
	return &v
}

type NullableL2VPNTypeLabel struct {
	value *L2VPNTypeLabel
	isSet bool
}

func (v NullableL2VPNTypeLabel) Get() *L2VPNTypeLabel {
	return v.value
}

func (v *NullableL2VPNTypeLabel) Set(val *L2VPNTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableL2VPNTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableL2VPNTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2VPNTypeLabel(val *L2VPNTypeLabel) *NullableL2VPNTypeLabel {
	return &NullableL2VPNTypeLabel{value: val, isSet: true}
}

func (v NullableL2VPNTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2VPNTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
