# PatchedWritableBGPSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**BGPSessionStatusValue**](BGPSessionStatusValue.md) |  | [optional] 
**Site** | Pointer to [**NullableBriefSiteRequest**](BriefSiteRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Device** | Pointer to [**NullableBriefDeviceRequest**](BriefDeviceRequest.md) |  | [optional] 
**LocalAddress** | Pointer to [**BriefIPAddressRequest**](BriefIPAddressRequest.md) |  | [optional] 
**RemoteAddress** | Pointer to [**BriefIPAddressRequest**](BriefIPAddressRequest.md) |  | [optional] 
**LocalAs** | Pointer to [**BriefASNRequest**](BriefASNRequest.md) |  | [optional] 
**RemoteAs** | Pointer to [**BriefASNRequest**](BriefASNRequest.md) |  | [optional] 
**PeerGroup** | Pointer to [**NullableBriefBGPPeerGroupRequest**](BriefBGPPeerGroupRequest.md) |  | [optional] 
**ImportPolicies** | Pointer to **[]int32** |  | [optional] 
**ExportPolicies** | Pointer to **[]int32** |  | [optional] 
**PrefixListIn** | Pointer to [**NullableBriefPrefixListRequest**](BriefPrefixListRequest.md) |  | [optional] 
**PrefixListOut** | Pointer to [**NullableBriefPrefixListRequest**](BriefPrefixListRequest.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedWritableBGPSessionRequest

`func NewPatchedWritableBGPSessionRequest() *PatchedWritableBGPSessionRequest`

NewPatchedWritableBGPSessionRequest instantiates a new PatchedWritableBGPSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableBGPSessionRequestWithDefaults

`func NewPatchedWritableBGPSessionRequestWithDefaults() *PatchedWritableBGPSessionRequest`

NewPatchedWritableBGPSessionRequestWithDefaults instantiates a new PatchedWritableBGPSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PatchedWritableBGPSessionRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableBGPSessionRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableBGPSessionRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableBGPSessionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableBGPSessionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableBGPSessionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableBGPSessionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableBGPSessionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableBGPSessionRequest) GetStatus() BGPSessionStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableBGPSessionRequest) GetStatusOk() (*BGPSessionStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableBGPSessionRequest) SetStatus(v BGPSessionStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableBGPSessionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSite

`func (o *PatchedWritableBGPSessionRequest) GetSite() BriefSiteRequest`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableBGPSessionRequest) GetSiteOk() (*BriefSiteRequest, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableBGPSessionRequest) SetSite(v BriefSiteRequest)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableBGPSessionRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *PatchedWritableBGPSessionRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *PatchedWritableBGPSessionRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetTenant

`func (o *PatchedWritableBGPSessionRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableBGPSessionRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableBGPSessionRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableBGPSessionRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableBGPSessionRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableBGPSessionRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDevice

`func (o *PatchedWritableBGPSessionRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableBGPSessionRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableBGPSessionRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableBGPSessionRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *PatchedWritableBGPSessionRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *PatchedWritableBGPSessionRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetLocalAddress

`func (o *PatchedWritableBGPSessionRequest) GetLocalAddress() BriefIPAddressRequest`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *PatchedWritableBGPSessionRequest) GetLocalAddressOk() (*BriefIPAddressRequest, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *PatchedWritableBGPSessionRequest) SetLocalAddress(v BriefIPAddressRequest)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *PatchedWritableBGPSessionRequest) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *PatchedWritableBGPSessionRequest) GetRemoteAddress() BriefIPAddressRequest`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *PatchedWritableBGPSessionRequest) GetRemoteAddressOk() (*BriefIPAddressRequest, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *PatchedWritableBGPSessionRequest) SetRemoteAddress(v BriefIPAddressRequest)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *PatchedWritableBGPSessionRequest) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetLocalAs

`func (o *PatchedWritableBGPSessionRequest) GetLocalAs() BriefASNRequest`

GetLocalAs returns the LocalAs field if non-nil, zero value otherwise.

### GetLocalAsOk

`func (o *PatchedWritableBGPSessionRequest) GetLocalAsOk() (*BriefASNRequest, bool)`

GetLocalAsOk returns a tuple with the LocalAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAs

`func (o *PatchedWritableBGPSessionRequest) SetLocalAs(v BriefASNRequest)`

SetLocalAs sets LocalAs field to given value.

### HasLocalAs

`func (o *PatchedWritableBGPSessionRequest) HasLocalAs() bool`

HasLocalAs returns a boolean if a field has been set.

### GetRemoteAs

`func (o *PatchedWritableBGPSessionRequest) GetRemoteAs() BriefASNRequest`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *PatchedWritableBGPSessionRequest) GetRemoteAsOk() (*BriefASNRequest, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *PatchedWritableBGPSessionRequest) SetRemoteAs(v BriefASNRequest)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *PatchedWritableBGPSessionRequest) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetPeerGroup

`func (o *PatchedWritableBGPSessionRequest) GetPeerGroup() BriefBGPPeerGroupRequest`

GetPeerGroup returns the PeerGroup field if non-nil, zero value otherwise.

### GetPeerGroupOk

`func (o *PatchedWritableBGPSessionRequest) GetPeerGroupOk() (*BriefBGPPeerGroupRequest, bool)`

GetPeerGroupOk returns a tuple with the PeerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGroup

`func (o *PatchedWritableBGPSessionRequest) SetPeerGroup(v BriefBGPPeerGroupRequest)`

SetPeerGroup sets PeerGroup field to given value.

### HasPeerGroup

`func (o *PatchedWritableBGPSessionRequest) HasPeerGroup() bool`

HasPeerGroup returns a boolean if a field has been set.

### SetPeerGroupNil

`func (o *PatchedWritableBGPSessionRequest) SetPeerGroupNil(b bool)`

 SetPeerGroupNil sets the value for PeerGroup to be an explicit nil

### UnsetPeerGroup
`func (o *PatchedWritableBGPSessionRequest) UnsetPeerGroup()`

UnsetPeerGroup ensures that no value is present for PeerGroup, not even an explicit nil
### GetImportPolicies

`func (o *PatchedWritableBGPSessionRequest) GetImportPolicies() []*int32`

GetImportPolicies returns the ImportPolicies field if non-nil, zero value otherwise.

### GetImportPoliciesOk

`func (o *PatchedWritableBGPSessionRequest) GetImportPoliciesOk() (*[]*int32, bool)`

GetImportPoliciesOk returns a tuple with the ImportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicies

`func (o *PatchedWritableBGPSessionRequest) SetImportPolicies(v []*int32)`

SetImportPolicies sets ImportPolicies field to given value.

### HasImportPolicies

`func (o *PatchedWritableBGPSessionRequest) HasImportPolicies() bool`

HasImportPolicies returns a boolean if a field has been set.

### GetExportPolicies

`func (o *PatchedWritableBGPSessionRequest) GetExportPolicies() []*int32`

GetExportPolicies returns the ExportPolicies field if non-nil, zero value otherwise.

### GetExportPoliciesOk

`func (o *PatchedWritableBGPSessionRequest) GetExportPoliciesOk() (*[]*int32, bool)`

GetExportPoliciesOk returns a tuple with the ExportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicies

`func (o *PatchedWritableBGPSessionRequest) SetExportPolicies(v []*int32)`

SetExportPolicies sets ExportPolicies field to given value.

### HasExportPolicies

`func (o *PatchedWritableBGPSessionRequest) HasExportPolicies() bool`

HasExportPolicies returns a boolean if a field has been set.

### GetPrefixListIn

`func (o *PatchedWritableBGPSessionRequest) GetPrefixListIn() BriefPrefixListRequest`

GetPrefixListIn returns the PrefixListIn field if non-nil, zero value otherwise.

### GetPrefixListInOk

`func (o *PatchedWritableBGPSessionRequest) GetPrefixListInOk() (*BriefPrefixListRequest, bool)`

GetPrefixListInOk returns a tuple with the PrefixListIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListIn

`func (o *PatchedWritableBGPSessionRequest) SetPrefixListIn(v BriefPrefixListRequest)`

SetPrefixListIn sets PrefixListIn field to given value.

### HasPrefixListIn

`func (o *PatchedWritableBGPSessionRequest) HasPrefixListIn() bool`

HasPrefixListIn returns a boolean if a field has been set.

### SetPrefixListInNil

`func (o *PatchedWritableBGPSessionRequest) SetPrefixListInNil(b bool)`

 SetPrefixListInNil sets the value for PrefixListIn to be an explicit nil

### UnsetPrefixListIn
`func (o *PatchedWritableBGPSessionRequest) UnsetPrefixListIn()`

UnsetPrefixListIn ensures that no value is present for PrefixListIn, not even an explicit nil
### GetPrefixListOut

`func (o *PatchedWritableBGPSessionRequest) GetPrefixListOut() BriefPrefixListRequest`

GetPrefixListOut returns the PrefixListOut field if non-nil, zero value otherwise.

### GetPrefixListOutOk

`func (o *PatchedWritableBGPSessionRequest) GetPrefixListOutOk() (*BriefPrefixListRequest, bool)`

GetPrefixListOutOk returns a tuple with the PrefixListOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListOut

`func (o *PatchedWritableBGPSessionRequest) SetPrefixListOut(v BriefPrefixListRequest)`

SetPrefixListOut sets PrefixListOut field to given value.

### HasPrefixListOut

`func (o *PatchedWritableBGPSessionRequest) HasPrefixListOut() bool`

HasPrefixListOut returns a boolean if a field has been set.

### SetPrefixListOutNil

`func (o *PatchedWritableBGPSessionRequest) SetPrefixListOutNil(b bool)`

 SetPrefixListOutNil sets the value for PrefixListOut to be an explicit nil

### UnsetPrefixListOut
`func (o *PatchedWritableBGPSessionRequest) UnsetPrefixListOut()`

UnsetPrefixListOut ensures that no value is present for PrefixListOut, not even an explicit nil
### GetName

`func (o *PatchedWritableBGPSessionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableBGPSessionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableBGPSessionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableBGPSessionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedWritableBGPSessionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedWritableBGPSessionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PatchedWritableBGPSessionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableBGPSessionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableBGPSessionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableBGPSessionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableBGPSessionRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableBGPSessionRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableBGPSessionRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableBGPSessionRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


