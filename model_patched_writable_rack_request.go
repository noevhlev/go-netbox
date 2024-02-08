/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableRackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableRackRequest{}

// PatchedWritableRackRequest Adds support for custom fields and tags.
type PatchedWritableRackRequest struct {
	Name       *string                           `json:"name,omitempty"`
	FacilityId NullableString                    `json:"facility_id,omitempty"`
	Site       *int32                            `json:"site,omitempty"`
	Location   NullableInt32                     `json:"location,omitempty"`
	Tenant     NullableInt32                     `json:"tenant,omitempty"`
	Status     *PatchedWritableRackRequestStatus `json:"status,omitempty"`
	// Functional role
	Role   NullableInt32 `json:"role,omitempty"`
	Serial *string       `json:"serial,omitempty"`
	// A unique tag used to identify this rack
	AssetTag NullableString                   `json:"asset_tag,omitempty"`
	Type     *PatchedWritableRackRequestType  `json:"type,omitempty"`
	Width    *PatchedWritableRackRequestWidth `json:"width,omitempty"`
	// Height in rack units
	UHeight *int32 `json:"u_height,omitempty"`
	// Starting unit for rack
	StartingUnit *int32          `json:"starting_unit,omitempty"`
	Weight       NullableFloat64 `json:"weight,omitempty"`
	// Maximum load capacity for the rack
	MaxWeight  NullableInt32              `json:"max_weight,omitempty"`
	WeightUnit *DeviceTypeWeightUnitValue `json:"weight_unit,omitempty"`
	// Units are numbered top-to-bottom
	DescUnits *bool `json:"desc_units,omitempty"`
	// Outer dimension of rack (width)
	OuterWidth NullableInt32 `json:"outer_width,omitempty"`
	// Outer dimension of rack (depth)
	OuterDepth NullableInt32                        `json:"outer_depth,omitempty"`
	OuterUnit  *PatchedWritableRackRequestOuterUnit `json:"outer_unit,omitempty"`
	// Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	MountingDepth        NullableInt32          `json:"mounting_depth,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableRackRequest PatchedWritableRackRequest

// NewPatchedWritableRackRequest instantiates a new PatchedWritableRackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableRackRequest() *PatchedWritableRackRequest {
	this := PatchedWritableRackRequest{}
	return &this
}

// NewPatchedWritableRackRequestWithDefaults instantiates a new PatchedWritableRackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableRackRequestWithDefaults() *PatchedWritableRackRequest {
	this := PatchedWritableRackRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableRackRequest) SetName(v string) {
	o.Name = &v
}

// GetFacilityId returns the FacilityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetFacilityId() string {
	if o == nil || IsNil(o.FacilityId.Get()) {
		var ret string
		return ret
	}
	return *o.FacilityId.Get()
}

// GetFacilityIdOk returns a tuple with the FacilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetFacilityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacilityId.Get(), o.FacilityId.IsSet()
}

// HasFacilityId returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasFacilityId() bool {
	if o != nil && o.FacilityId.IsSet() {
		return true
	}

	return false
}

// SetFacilityId gets a reference to the given NullableString and assigns it to the FacilityId field.
func (o *PatchedWritableRackRequest) SetFacilityId(v string) {
	o.FacilityId.Set(&v)
}

// SetFacilityIdNil sets the value for FacilityId to be an explicit nil
func (o *PatchedWritableRackRequest) SetFacilityIdNil() {
	o.FacilityId.Set(nil)
}

// UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetFacilityId() {
	o.FacilityId.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetSite() int32 {
	if o == nil || IsNil(o.Site) {
		var ret int32
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetSiteOk() (*int32, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given int32 and assigns it to the Site field.
func (o *PatchedWritableRackRequest) SetSite(v int32) {
	o.Site = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetLocation() int32 {
	if o == nil || IsNil(o.Location.Get()) {
		var ret int32
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetLocationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableInt32 and assigns it to the Location field.
func (o *PatchedWritableRackRequest) SetLocation(v int32) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *PatchedWritableRackRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *PatchedWritableRackRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableRackRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetStatus() PatchedWritableRackRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableRackRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetStatusOk() (*PatchedWritableRackRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableRackRequestStatus and assigns it to the Status field.
func (o *PatchedWritableRackRequest) SetStatus(v PatchedWritableRackRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetRole() int32 {
	if o == nil || IsNil(o.Role.Get()) {
		var ret int32
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableInt32 and assigns it to the Role field.
func (o *PatchedWritableRackRequest) SetRole(v int32) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedWritableRackRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetRole() {
	o.Role.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *PatchedWritableRackRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *PatchedWritableRackRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *PatchedWritableRackRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetType() PatchedWritableRackRequestType {
	if o == nil || IsNil(o.Type) {
		var ret PatchedWritableRackRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetTypeOk() (*PatchedWritableRackRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PatchedWritableRackRequestType and assigns it to the Type field.
func (o *PatchedWritableRackRequest) SetType(v PatchedWritableRackRequestType) {
	o.Type = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetWidth() PatchedWritableRackRequestWidth {
	if o == nil || IsNil(o.Width) {
		var ret PatchedWritableRackRequestWidth
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetWidthOk() (*PatchedWritableRackRequestWidth, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given PatchedWritableRackRequestWidth and assigns it to the Width field.
func (o *PatchedWritableRackRequest) SetWidth(v PatchedWritableRackRequestWidth) {
	o.Width = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *PatchedWritableRackRequest) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetStartingUnit returns the StartingUnit field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetStartingUnit() int32 {
	if o == nil || IsNil(o.StartingUnit) {
		var ret int32
		return ret
	}
	return *o.StartingUnit
}

// GetStartingUnitOk returns a tuple with the StartingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetStartingUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.StartingUnit) {
		return nil, false
	}
	return o.StartingUnit, true
}

// HasStartingUnit returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasStartingUnit() bool {
	if o != nil && !IsNil(o.StartingUnit) {
		return true
	}

	return false
}

// SetStartingUnit gets a reference to the given int32 and assigns it to the StartingUnit field.
func (o *PatchedWritableRackRequest) SetStartingUnit(v int32) {
	o.StartingUnit = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *PatchedWritableRackRequest) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *PatchedWritableRackRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetMaxWeight() int32 {
	if o == nil || IsNil(o.MaxWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxWeight.Get()
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetMaxWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxWeight.Get(), o.MaxWeight.IsSet()
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasMaxWeight() bool {
	if o != nil && o.MaxWeight.IsSet() {
		return true
	}

	return false
}

// SetMaxWeight gets a reference to the given NullableInt32 and assigns it to the MaxWeight field.
func (o *PatchedWritableRackRequest) SetMaxWeight(v int32) {
	o.MaxWeight.Set(&v)
}

// SetMaxWeightNil sets the value for MaxWeight to be an explicit nil
func (o *PatchedWritableRackRequest) SetMaxWeightNil() {
	o.MaxWeight.Set(nil)
}

// UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetMaxWeight() {
	o.MaxWeight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetWeightUnit() DeviceTypeWeightUnitValue {
	if o == nil || IsNil(o.WeightUnit) {
		var ret DeviceTypeWeightUnitValue
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetWeightUnitOk() (*DeviceTypeWeightUnitValue, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given DeviceTypeWeightUnitValue and assigns it to the WeightUnit field.
func (o *PatchedWritableRackRequest) SetWeightUnit(v DeviceTypeWeightUnitValue) {
	o.WeightUnit = &v
}

// GetDescUnits returns the DescUnits field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetDescUnits() bool {
	if o == nil || IsNil(o.DescUnits) {
		var ret bool
		return ret
	}
	return *o.DescUnits
}

// GetDescUnitsOk returns a tuple with the DescUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetDescUnitsOk() (*bool, bool) {
	if o == nil || IsNil(o.DescUnits) {
		return nil, false
	}
	return o.DescUnits, true
}

// HasDescUnits returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasDescUnits() bool {
	if o != nil && !IsNil(o.DescUnits) {
		return true
	}

	return false
}

// SetDescUnits gets a reference to the given bool and assigns it to the DescUnits field.
func (o *PatchedWritableRackRequest) SetDescUnits(v bool) {
	o.DescUnits = &v
}

// GetOuterWidth returns the OuterWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetOuterWidth() int32 {
	if o == nil || IsNil(o.OuterWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterWidth.Get()
}

// GetOuterWidthOk returns a tuple with the OuterWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetOuterWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterWidth.Get(), o.OuterWidth.IsSet()
}

// HasOuterWidth returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasOuterWidth() bool {
	if o != nil && o.OuterWidth.IsSet() {
		return true
	}

	return false
}

// SetOuterWidth gets a reference to the given NullableInt32 and assigns it to the OuterWidth field.
func (o *PatchedWritableRackRequest) SetOuterWidth(v int32) {
	o.OuterWidth.Set(&v)
}

// SetOuterWidthNil sets the value for OuterWidth to be an explicit nil
func (o *PatchedWritableRackRequest) SetOuterWidthNil() {
	o.OuterWidth.Set(nil)
}

// UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetOuterWidth() {
	o.OuterWidth.Unset()
}

// GetOuterDepth returns the OuterDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetOuterDepth() int32 {
	if o == nil || IsNil(o.OuterDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterDepth.Get()
}

// GetOuterDepthOk returns a tuple with the OuterDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetOuterDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterDepth.Get(), o.OuterDepth.IsSet()
}

// HasOuterDepth returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasOuterDepth() bool {
	if o != nil && o.OuterDepth.IsSet() {
		return true
	}

	return false
}

// SetOuterDepth gets a reference to the given NullableInt32 and assigns it to the OuterDepth field.
func (o *PatchedWritableRackRequest) SetOuterDepth(v int32) {
	o.OuterDepth.Set(&v)
}

// SetOuterDepthNil sets the value for OuterDepth to be an explicit nil
func (o *PatchedWritableRackRequest) SetOuterDepthNil() {
	o.OuterDepth.Set(nil)
}

// UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetOuterDepth() {
	o.OuterDepth.Unset()
}

// GetOuterUnit returns the OuterUnit field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetOuterUnit() PatchedWritableRackRequestOuterUnit {
	if o == nil || IsNil(o.OuterUnit) {
		var ret PatchedWritableRackRequestOuterUnit
		return ret
	}
	return *o.OuterUnit
}

// GetOuterUnitOk returns a tuple with the OuterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetOuterUnitOk() (*PatchedWritableRackRequestOuterUnit, bool) {
	if o == nil || IsNil(o.OuterUnit) {
		return nil, false
	}
	return o.OuterUnit, true
}

// HasOuterUnit returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasOuterUnit() bool {
	if o != nil && !IsNil(o.OuterUnit) {
		return true
	}

	return false
}

// SetOuterUnit gets a reference to the given PatchedWritableRackRequestOuterUnit and assigns it to the OuterUnit field.
func (o *PatchedWritableRackRequest) SetOuterUnit(v PatchedWritableRackRequestOuterUnit) {
	o.OuterUnit = &v
}

// GetMountingDepth returns the MountingDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackRequest) GetMountingDepth() int32 {
	if o == nil || IsNil(o.MountingDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.MountingDepth.Get()
}

// GetMountingDepthOk returns a tuple with the MountingDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackRequest) GetMountingDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountingDepth.Get(), o.MountingDepth.IsSet()
}

// HasMountingDepth returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasMountingDepth() bool {
	if o != nil && o.MountingDepth.IsSet() {
		return true
	}

	return false
}

// SetMountingDepth gets a reference to the given NullableInt32 and assigns it to the MountingDepth field.
func (o *PatchedWritableRackRequest) SetMountingDepth(v int32) {
	o.MountingDepth.Set(&v)
}

// SetMountingDepthNil sets the value for MountingDepth to be an explicit nil
func (o *PatchedWritableRackRequest) SetMountingDepthNil() {
	o.MountingDepth.Set(nil)
}

// UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
func (o *PatchedWritableRackRequest) UnsetMountingDepth() {
	o.MountingDepth.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableRackRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableRackRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableRackRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableRackRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableRackRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableRackRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableRackRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableRackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.FacilityId.IsSet() {
		toSerialize["facility_id"] = o.FacilityId.Get()
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.StartingUnit) {
		toSerialize["starting_unit"] = o.StartingUnit
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.MaxWeight.IsSet() {
		toSerialize["max_weight"] = o.MaxWeight.Get()
	}
	if !IsNil(o.WeightUnit) {
		toSerialize["weight_unit"] = o.WeightUnit
	}
	if !IsNil(o.DescUnits) {
		toSerialize["desc_units"] = o.DescUnits
	}
	if o.OuterWidth.IsSet() {
		toSerialize["outer_width"] = o.OuterWidth.Get()
	}
	if o.OuterDepth.IsSet() {
		toSerialize["outer_depth"] = o.OuterDepth.Get()
	}
	if !IsNil(o.OuterUnit) {
		toSerialize["outer_unit"] = o.OuterUnit
	}
	if o.MountingDepth.IsSet() {
		toSerialize["mounting_depth"] = o.MountingDepth.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableRackRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableRackRequest := _PatchedWritableRackRequest{}

	err = json.Unmarshal(data, &varPatchedWritableRackRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableRackRequest(varPatchedWritableRackRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "facility_id")
		delete(additionalProperties, "site")
		delete(additionalProperties, "location")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "type")
		delete(additionalProperties, "width")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "starting_unit")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "max_weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "desc_units")
		delete(additionalProperties, "outer_width")
		delete(additionalProperties, "outer_depth")
		delete(additionalProperties, "outer_unit")
		delete(additionalProperties, "mounting_depth")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableRackRequest struct {
	value *PatchedWritableRackRequest
	isSet bool
}

func (v NullablePatchedWritableRackRequest) Get() *PatchedWritableRackRequest {
	return v.value
}

func (v *NullablePatchedWritableRackRequest) Set(val *PatchedWritableRackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackRequest(val *PatchedWritableRackRequest) *NullablePatchedWritableRackRequest {
	return &NullablePatchedWritableRackRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableRackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
