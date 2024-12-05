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
	"time"
)

// checks if the BGPSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BGPSession{}

// BGPSession Adds support for custom fields and tags.
type BGPSession struct {
	Id                   int32                     `json:"id"`
	Url                  string                    `json:"url"`
	Tags                 []NestedTag               `json:"tags,omitempty"`
	CustomFields         map[string]interface{}    `json:"custom_fields,omitempty"`
	Display              string                    `json:"display"`
	Status               *BGPSessionStatus         `json:"status,omitempty"`
	Site                 NullableBriefSite         `json:"site,omitempty"`
	Tenant               NullableBriefTenant       `json:"tenant,omitempty"`
	Device               NullableBriefDevice       `json:"device,omitempty"`
	LocalAddress         BriefIPAddress            `json:"local_address"`
	RemoteAddress        BriefIPAddress            `json:"remote_address"`
	LocalAs              BriefASN                  `json:"local_as"`
	RemoteAs             BriefASN                  `json:"remote_as"`
	PeerGroup            NullableBriefBGPPeerGroup `json:"peer_group,omitempty"`
	ImportPolicies       []*RoutingPolicy          `json:"import_policies,omitempty"`
	ExportPolicies       []*RoutingPolicy          `json:"export_policies,omitempty"`
	PrefixListIn         NullableBriefPrefixList   `json:"prefix_list_in,omitempty"`
	PrefixListOut        NullableBriefPrefixList   `json:"prefix_list_out,omitempty"`
	Created              NullableTime              `json:"created"`
	LastUpdated          NullableTime              `json:"last_updated"`
	Name                 NullableString            `json:"name,omitempty"`
	Description          *string                   `json:"description,omitempty"`
	Comments             *string                   `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BGPSession BGPSession

// NewBGPSession instantiates a new BGPSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBGPSession(id int32, url string, display string, localAddress BriefIPAddress, remoteAddress BriefIPAddress, localAs BriefASN, remoteAs BriefASN, created NullableTime, lastUpdated NullableTime) *BGPSession {
	this := BGPSession{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.LocalAddress = localAddress
	this.RemoteAddress = remoteAddress
	this.LocalAs = localAs
	this.RemoteAs = remoteAs
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewBGPSessionWithDefaults instantiates a new BGPSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBGPSessionWithDefaults() *BGPSession {
	this := BGPSession{}
	return &this
}

// GetId returns the Id field value
func (o *BGPSession) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BGPSession) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BGPSession) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BGPSession) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BGPSession) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BGPSession) SetUrl(v string) {
	o.Url = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BGPSession) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPSession) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BGPSession) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *BGPSession) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BGPSession) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPSession) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BGPSession) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BGPSession) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *BGPSession) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BGPSession) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BGPSession) SetDisplay(v string) {
	o.Display = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BGPSession) GetStatus() BGPSessionStatus {
	if o == nil || IsNil(o.Status) {
		var ret BGPSessionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPSession) GetStatusOk() (*BGPSessionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BGPSession) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BGPSessionStatus and assigns it to the Status field.
func (o *BGPSession) SetStatus(v BGPSessionStatus) {
	o.Status = &v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BGPSession) GetSite() BriefSite {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSite
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetSiteOk() (*BriefSite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *BGPSession) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSite and assigns it to the Site field.
func (o *BGPSession) SetSite(v BriefSite) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *BGPSession) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *BGPSession) UnsetSite() {
	o.Site.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BGPSession) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *BGPSession) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *BGPSession) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *BGPSession) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *BGPSession) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BGPSession) GetDevice() BriefDevice {
	if o == nil || IsNil(o.Device.Get()) {
		var ret BriefDevice
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *BGPSession) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableBriefDevice and assigns it to the Device field.
func (o *BGPSession) SetDevice(v BriefDevice) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *BGPSession) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *BGPSession) UnsetDevice() {
	o.Device.Unset()
}

// GetLocalAddress returns the LocalAddress field value
func (o *BGPSession) GetLocalAddress() BriefIPAddress {
	if o == nil {
		var ret BriefIPAddress
		return ret
	}

	return o.LocalAddress
}

// GetLocalAddressOk returns a tuple with the LocalAddress field value
// and a boolean to check if the value has been set.
func (o *BGPSession) GetLocalAddressOk() (*BriefIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalAddress, true
}

// SetLocalAddress sets field value
func (o *BGPSession) SetLocalAddress(v BriefIPAddress) {
	o.LocalAddress = v
}

// GetRemoteAddress returns the RemoteAddress field value
func (o *BGPSession) GetRemoteAddress() BriefIPAddress {
	if o == nil {
		var ret BriefIPAddress
		return ret
	}

	return o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value
// and a boolean to check if the value has been set.
func (o *BGPSession) GetRemoteAddressOk() (*BriefIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAddress, true
}

// SetRemoteAddress sets field value
func (o *BGPSession) SetRemoteAddress(v BriefIPAddress) {
	o.RemoteAddress = v
}

// GetLocalAs returns the LocalAs field value
func (o *BGPSession) GetLocalAs() BriefASN {
	if o == nil {
		var ret BriefASN
		return ret
	}

	return o.LocalAs
}

// GetLocalAsOk returns a tuple with the LocalAs field value
// and a boolean to check if the value has been set.
func (o *BGPSession) GetLocalAsOk() (*BriefASN, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalAs, true
}

// SetLocalAs sets field value
func (o *BGPSession) SetLocalAs(v BriefASN) {
	o.LocalAs = v
}

// GetRemoteAs returns the RemoteAs field value
func (o *BGPSession) GetRemoteAs() BriefASN {
	if o == nil {
		var ret BriefASN
		return ret
	}

	return o.RemoteAs
}

// GetRemoteAsOk returns a tuple with the RemoteAs field value
// and a boolean to check if the value has been set.
func (o *BGPSession) GetRemoteAsOk() (*BriefASN, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAs, true
}

// SetRemoteAs sets field value
func (o *BGPSession) SetRemoteAs(v BriefASN) {
	o.RemoteAs = v
}

// GetPeerGroup returns the PeerGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BGPSession) GetPeerGroup() BriefBGPPeerGroup {
	if o == nil || IsNil(o.PeerGroup.Get()) {
		var ret BriefBGPPeerGroup
		return ret
	}
	return *o.PeerGroup.Get()
}

// GetPeerGroupOk returns a tuple with the PeerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetPeerGroupOk() (*BriefBGPPeerGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeerGroup.Get(), o.PeerGroup.IsSet()
}

// HasPeerGroup returns a boolean if a field has been set.
func (o *BGPSession) HasPeerGroup() bool {
	if o != nil && o.PeerGroup.IsSet() {
		return true
	}

	return false
}

// SetPeerGroup gets a reference to the given NullableBriefBGPPeerGroup and assigns it to the PeerGroup field.
func (o *BGPSession) SetPeerGroup(v BriefBGPPeerGroup) {
	o.PeerGroup.Set(&v)
}

// SetPeerGroupNil sets the value for PeerGroup to be an explicit nil
func (o *BGPSession) SetPeerGroupNil() {
	o.PeerGroup.Set(nil)
}

// UnsetPeerGroup ensures that no value is present for PeerGroup, not even an explicit nil
func (o *BGPSession) UnsetPeerGroup() {
	o.PeerGroup.Unset()
}

// GetImportPolicies returns the ImportPolicies field value if set, zero value otherwise.
func (o *BGPSession) GetImportPolicies() []*RoutingPolicy {
	if o == nil || IsNil(o.ImportPolicies) {
		var ret []*RoutingPolicy
		return ret
	}
	return o.ImportPolicies
}

// GetImportPoliciesOk returns a tuple with the ImportPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPSession) GetImportPoliciesOk() ([]*RoutingPolicy, bool) {
	if o == nil || IsNil(o.ImportPolicies) {
		return nil, false
	}
	return o.ImportPolicies, true
}

// HasImportPolicies returns a boolean if a field has been set.
func (o *BGPSession) HasImportPolicies() bool {
	if o != nil && !IsNil(o.ImportPolicies) {
		return true
	}

	return false
}

// SetImportPolicies gets a reference to the given []*RoutingPolicy and assigns it to the ImportPolicies field.
func (o *BGPSession) SetImportPolicies(v []*RoutingPolicy) {
	o.ImportPolicies = v
}

// GetExportPolicies returns the ExportPolicies field value if set, zero value otherwise.
func (o *BGPSession) GetExportPolicies() []*RoutingPolicy {
	if o == nil || IsNil(o.ExportPolicies) {
		var ret []*RoutingPolicy
		return ret
	}
	return o.ExportPolicies
}

// GetExportPoliciesOk returns a tuple with the ExportPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPSession) GetExportPoliciesOk() ([]*RoutingPolicy, bool) {
	if o == nil || IsNil(o.ExportPolicies) {
		return nil, false
	}
	return o.ExportPolicies, true
}

// HasExportPolicies returns a boolean if a field has been set.
func (o *BGPSession) HasExportPolicies() bool {
	if o != nil && !IsNil(o.ExportPolicies) {
		return true
	}

	return false
}

// SetExportPolicies gets a reference to the given []*RoutingPolicy and assigns it to the ExportPolicies field.
func (o *BGPSession) SetExportPolicies(v []*RoutingPolicy) {
	o.ExportPolicies = v
}

// GetPrefixListIn returns the PrefixListIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BGPSession) GetPrefixListIn() BriefPrefixList {
	if o == nil || IsNil(o.PrefixListIn.Get()) {
		var ret BriefPrefixList
		return ret
	}
	return *o.PrefixListIn.Get()
}

// GetPrefixListInOk returns a tuple with the PrefixListIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetPrefixListInOk() (*BriefPrefixList, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrefixListIn.Get(), o.PrefixListIn.IsSet()
}

// HasPrefixListIn returns a boolean if a field has been set.
func (o *BGPSession) HasPrefixListIn() bool {
	if o != nil && o.PrefixListIn.IsSet() {
		return true
	}

	return false
}

// SetPrefixListIn gets a reference to the given NullableBriefPrefixList and assigns it to the PrefixListIn field.
func (o *BGPSession) SetPrefixListIn(v BriefPrefixList) {
	o.PrefixListIn.Set(&v)
}

// SetPrefixListInNil sets the value for PrefixListIn to be an explicit nil
func (o *BGPSession) SetPrefixListInNil() {
	o.PrefixListIn.Set(nil)
}

// UnsetPrefixListIn ensures that no value is present for PrefixListIn, not even an explicit nil
func (o *BGPSession) UnsetPrefixListIn() {
	o.PrefixListIn.Unset()
}

// GetPrefixListOut returns the PrefixListOut field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BGPSession) GetPrefixListOut() BriefPrefixList {
	if o == nil || IsNil(o.PrefixListOut.Get()) {
		var ret BriefPrefixList
		return ret
	}
	return *o.PrefixListOut.Get()
}

// GetPrefixListOutOk returns a tuple with the PrefixListOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetPrefixListOutOk() (*BriefPrefixList, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrefixListOut.Get(), o.PrefixListOut.IsSet()
}

// HasPrefixListOut returns a boolean if a field has been set.
func (o *BGPSession) HasPrefixListOut() bool {
	if o != nil && o.PrefixListOut.IsSet() {
		return true
	}

	return false
}

// SetPrefixListOut gets a reference to the given NullableBriefPrefixList and assigns it to the PrefixListOut field.
func (o *BGPSession) SetPrefixListOut(v BriefPrefixList) {
	o.PrefixListOut.Set(&v)
}

// SetPrefixListOutNil sets the value for PrefixListOut to be an explicit nil
func (o *BGPSession) SetPrefixListOutNil() {
	o.PrefixListOut.Set(nil)
}

// UnsetPrefixListOut ensures that no value is present for PrefixListOut, not even an explicit nil
func (o *BGPSession) UnsetPrefixListOut() {
	o.PrefixListOut.Unset()
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BGPSession) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *BGPSession) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BGPSession) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *BGPSession) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BGPSession) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BGPSession) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BGPSession) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BGPSession) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *BGPSession) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BGPSession) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BGPSession) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPSession) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BGPSession) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BGPSession) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *BGPSession) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPSession) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *BGPSession) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *BGPSession) SetComments(v string) {
	o.Comments = &v
}

func (o BGPSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BGPSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["display"] = o.Display
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	toSerialize["local_address"] = o.LocalAddress
	toSerialize["remote_address"] = o.RemoteAddress
	toSerialize["local_as"] = o.LocalAs
	toSerialize["remote_as"] = o.RemoteAs
	if o.PeerGroup.IsSet() {
		toSerialize["peer_group"] = o.PeerGroup.Get()
	}
	if !IsNil(o.ImportPolicies) {
		toSerialize["import_policies"] = o.ImportPolicies
	}
	if !IsNil(o.ExportPolicies) {
		toSerialize["export_policies"] = o.ExportPolicies
	}
	if o.PrefixListIn.IsSet() {
		toSerialize["prefix_list_in"] = o.PrefixListIn.Get()
	}
	if o.PrefixListOut.IsSet() {
		toSerialize["prefix_list_out"] = o.PrefixListOut.Get()
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BGPSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"local_address",
		"remote_address",
		"local_as",
		"remote_as",
		"created",
		"last_updated",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBGPSession := _BGPSession{}

	err = json.Unmarshal(data, &varBGPSession)

	if err != nil {
		return err
	}

	*o = BGPSession(varBGPSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "display")
		delete(additionalProperties, "status")
		delete(additionalProperties, "site")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "device")
		delete(additionalProperties, "local_address")
		delete(additionalProperties, "remote_address")
		delete(additionalProperties, "local_as")
		delete(additionalProperties, "remote_as")
		delete(additionalProperties, "peer_group")
		delete(additionalProperties, "import_policies")
		delete(additionalProperties, "export_policies")
		delete(additionalProperties, "prefix_list_in")
		delete(additionalProperties, "prefix_list_out")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBGPSession struct {
	value *BGPSession
	isSet bool
}

func (v NullableBGPSession) Get() *BGPSession {
	return v.value
}

func (v *NullableBGPSession) Set(val *BGPSession) {
	v.value = val
	v.isSet = true
}

func (v NullableBGPSession) IsSet() bool {
	return v.isSet
}

func (v *NullableBGPSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBGPSession(val *BGPSession) *NullableBGPSession {
	return &NullableBGPSession{value: val, isSet: true}
}

func (v NullableBGPSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBGPSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
