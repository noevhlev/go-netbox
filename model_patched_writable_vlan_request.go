/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableVLANRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableVLANRequest{}

// PatchedWritableVLANRequest Adds support for custom fields and tags.
type PatchedWritableVLANRequest struct {
	Site  NullableBriefSiteRequest      `json:"site,omitempty"`
	Group NullableBriefVLANGroupRequest `json:"group,omitempty"`
	// Numeric VLAN ID (1-4094)
	Vid                  *int32                            `json:"vid,omitempty"`
	Name                 *string                           `json:"name,omitempty"`
	Tenant               NullableBriefTenantRequest        `json:"tenant,omitempty"`
	Status               *PatchedWritableVLANRequestStatus `json:"status,omitempty"`
	Role                 NullableBriefRoleRequest          `json:"role,omitempty"`
	Description          *string                           `json:"description,omitempty"`
	Comments             *string                           `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                `json:"tags,omitempty"`
	CustomFields         map[string]interface{}            `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableVLANRequest PatchedWritableVLANRequest

// NewPatchedWritableVLANRequest instantiates a new PatchedWritableVLANRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableVLANRequest() *PatchedWritableVLANRequest {
	this := PatchedWritableVLANRequest{}
	return &this
}

// NewPatchedWritableVLANRequestWithDefaults instantiates a new PatchedWritableVLANRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableVLANRequestWithDefaults() *PatchedWritableVLANRequest {
	this := PatchedWritableVLANRequest{}
	return &this
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVLANRequest) GetSite() BriefSiteRequest {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSiteRequest
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVLANRequest) GetSiteOk() (*BriefSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSiteRequest and assigns it to the Site field.
func (o *PatchedWritableVLANRequest) SetSite(v BriefSiteRequest) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *PatchedWritableVLANRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *PatchedWritableVLANRequest) UnsetSite() {
	o.Site.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVLANRequest) GetGroup() BriefVLANGroupRequest {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefVLANGroupRequest
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVLANRequest) GetGroupOk() (*BriefVLANGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefVLANGroupRequest and assigns it to the Group field.
func (o *PatchedWritableVLANRequest) SetGroup(v BriefVLANGroupRequest) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *PatchedWritableVLANRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *PatchedWritableVLANRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *PatchedWritableVLANRequest) GetVid() int32 {
	if o == nil || IsNil(o.Vid) {
		var ret int32
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVLANRequest) GetVidOk() (*int32, bool) {
	if o == nil || IsNil(o.Vid) {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasVid() bool {
	if o != nil && !IsNil(o.Vid) {
		return true
	}

	return false
}

// SetVid gets a reference to the given int32 and assigns it to the Vid field.
func (o *PatchedWritableVLANRequest) SetVid(v int32) {
	o.Vid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableVLANRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVLANRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableVLANRequest) SetName(v string) {
	o.Name = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVLANRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVLANRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *PatchedWritableVLANRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableVLANRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableVLANRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableVLANRequest) GetStatus() PatchedWritableVLANRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableVLANRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVLANRequest) GetStatusOk() (*PatchedWritableVLANRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableVLANRequestStatus and assigns it to the Status field.
func (o *PatchedWritableVLANRequest) SetStatus(v PatchedWritableVLANRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVLANRequest) GetRole() BriefRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVLANRequest) GetRoleOk() (*BriefRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefRoleRequest and assigns it to the Role field.
func (o *PatchedWritableVLANRequest) SetRole(v BriefRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedWritableVLANRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedWritableVLANRequest) UnsetRole() {
	o.Role.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableVLANRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVLANRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableVLANRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableVLANRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVLANRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableVLANRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableVLANRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVLANRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableVLANRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableVLANRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVLANRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableVLANRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableVLANRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableVLANRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableVLANRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Vid) {
		toSerialize["vid"] = o.Vid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

func (o *PatchedWritableVLANRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableVLANRequest := _PatchedWritableVLANRequest{}

	err = json.Unmarshal(data, &varPatchedWritableVLANRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableVLANRequest(varPatchedWritableVLANRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "site")
		delete(additionalProperties, "group")
		delete(additionalProperties, "vid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableVLANRequest struct {
	value *PatchedWritableVLANRequest
	isSet bool
}

func (v NullablePatchedWritableVLANRequest) Get() *PatchedWritableVLANRequest {
	return v.value
}

func (v *NullablePatchedWritableVLANRequest) Set(val *PatchedWritableVLANRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableVLANRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableVLANRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableVLANRequest(val *PatchedWritableVLANRequest) *NullablePatchedWritableVLANRequest {
	return &NullablePatchedWritableVLANRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableVLANRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableVLANRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
