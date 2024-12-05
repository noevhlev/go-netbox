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

// SiteStatusLabel the model 'SiteStatusLabel'
type SiteStatusLabel string

// List of Site_status_label
const (
	SITESTATUSLABEL_PLANNED         SiteStatusLabel = "Planned"
	SITESTATUSLABEL_STAGING         SiteStatusLabel = "Staging"
	SITESTATUSLABEL_ACTIVE          SiteStatusLabel = "Active"
	SITESTATUSLABEL_DECOMMISSIONING SiteStatusLabel = "Decommissioning"
	SITESTATUSLABEL_RETIRED         SiteStatusLabel = "Retired"
)

// All allowed values of SiteStatusLabel enum
var AllowedSiteStatusLabelEnumValues = []SiteStatusLabel{
	"Planned",
	"Staging",
	"Active",
	"Decommissioning",
	"Retired",
}

func (v *SiteStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteStatusLabel(value)
	for _, existing := range AllowedSiteStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteStatusLabel", value)
}

// NewSiteStatusLabelFromValue returns a pointer to a valid SiteStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteStatusLabelFromValue(v string) (*SiteStatusLabel, error) {
	ev := SiteStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteStatusLabel: valid values are %v", v, AllowedSiteStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteStatusLabel) IsValid() bool {
	for _, existing := range AllowedSiteStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Site_status_label value
func (v SiteStatusLabel) Ptr() *SiteStatusLabel {
	return &v
}

type NullableSiteStatusLabel struct {
	value *SiteStatusLabel
	isSet bool
}

func (v NullableSiteStatusLabel) Get() *SiteStatusLabel {
	return v.value
}

func (v *NullableSiteStatusLabel) Set(val *SiteStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteStatusLabel(val *SiteStatusLabel) *NullableSiteStatusLabel {
	return &NullableSiteStatusLabel{value: val, isSet: true}
}

func (v NullableSiteStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}