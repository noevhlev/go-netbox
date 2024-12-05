# BGPSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**BGPSessionStatusValue**](BGPSessionStatusValue.md) |  | [optional] 
**Site** | Pointer to [**NullableBriefSiteRequest**](BriefSiteRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Device** | Pointer to [**NullableBriefDeviceRequest**](BriefDeviceRequest.md) |  | [optional] 
**LocalAddress** | [**BriefIPAddressRequest**](BriefIPAddressRequest.md) |  | 
**RemoteAddress** | [**BriefIPAddressRequest**](BriefIPAddressRequest.md) |  | 
**LocalAs** | [**BriefASNRequest**](BriefASNRequest.md) |  | 
**RemoteAs** | [**BriefASNRequest**](BriefASNRequest.md) |  | 
**PeerGroup** | Pointer to [**NullableBriefBGPPeerGroupRequest**](BriefBGPPeerGroupRequest.md) |  | [optional] 
**ImportPolicies** | Pointer to **[]int32** |  | [optional] 
**ExportPolicies** | Pointer to **[]int32** |  | [optional] 
**PrefixListIn** | Pointer to [**NullableBriefPrefixListRequest**](BriefPrefixListRequest.md) |  | [optional] 
**PrefixListOut** | Pointer to [**NullableBriefPrefixListRequest**](BriefPrefixListRequest.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewBGPSessionRequest

`func NewBGPSessionRequest(localAddress BriefIPAddressRequest, remoteAddress BriefIPAddressRequest, localAs BriefASNRequest, remoteAs BriefASNRequest, ) *BGPSessionRequest`

NewBGPSessionRequest instantiates a new BGPSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPSessionRequestWithDefaults

`func NewBGPSessionRequestWithDefaults() *BGPSessionRequest`

NewBGPSessionRequestWithDefaults instantiates a new BGPSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *BGPSessionRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BGPSessionRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BGPSessionRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BGPSessionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BGPSessionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BGPSessionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BGPSessionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BGPSessionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetStatus

`func (o *BGPSessionRequest) GetStatus() BGPSessionStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BGPSessionRequest) GetStatusOk() (*BGPSessionStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BGPSessionRequest) SetStatus(v BGPSessionStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BGPSessionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSite

`func (o *BGPSessionRequest) GetSite() BriefSiteRequest`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *BGPSessionRequest) GetSiteOk() (*BriefSiteRequest, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *BGPSessionRequest) SetSite(v BriefSiteRequest)`

SetSite sets Site field to given value.

### HasSite

`func (o *BGPSessionRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *BGPSessionRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *BGPSessionRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetTenant

`func (o *BGPSessionRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BGPSessionRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BGPSessionRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BGPSessionRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BGPSessionRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BGPSessionRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDevice

`func (o *BGPSessionRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BGPSessionRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BGPSessionRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BGPSessionRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *BGPSessionRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *BGPSessionRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetLocalAddress

`func (o *BGPSessionRequest) GetLocalAddress() BriefIPAddressRequest`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *BGPSessionRequest) GetLocalAddressOk() (*BriefIPAddressRequest, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *BGPSessionRequest) SetLocalAddress(v BriefIPAddressRequest)`

SetLocalAddress sets LocalAddress field to given value.


### GetRemoteAddress

`func (o *BGPSessionRequest) GetRemoteAddress() BriefIPAddressRequest`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *BGPSessionRequest) GetRemoteAddressOk() (*BriefIPAddressRequest, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *BGPSessionRequest) SetRemoteAddress(v BriefIPAddressRequest)`

SetRemoteAddress sets RemoteAddress field to given value.


### GetLocalAs

`func (o *BGPSessionRequest) GetLocalAs() BriefASNRequest`

GetLocalAs returns the LocalAs field if non-nil, zero value otherwise.

### GetLocalAsOk

`func (o *BGPSessionRequest) GetLocalAsOk() (*BriefASNRequest, bool)`

GetLocalAsOk returns a tuple with the LocalAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAs

`func (o *BGPSessionRequest) SetLocalAs(v BriefASNRequest)`

SetLocalAs sets LocalAs field to given value.


### GetRemoteAs

`func (o *BGPSessionRequest) GetRemoteAs() BriefASNRequest`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *BGPSessionRequest) GetRemoteAsOk() (*BriefASNRequest, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *BGPSessionRequest) SetRemoteAs(v BriefASNRequest)`

SetRemoteAs sets RemoteAs field to given value.


### GetPeerGroup

`func (o *BGPSessionRequest) GetPeerGroup() BriefBGPPeerGroupRequest`

GetPeerGroup returns the PeerGroup field if non-nil, zero value otherwise.

### GetPeerGroupOk

`func (o *BGPSessionRequest) GetPeerGroupOk() (*BriefBGPPeerGroupRequest, bool)`

GetPeerGroupOk returns a tuple with the PeerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGroup

`func (o *BGPSessionRequest) SetPeerGroup(v BriefBGPPeerGroupRequest)`

SetPeerGroup sets PeerGroup field to given value.

### HasPeerGroup

`func (o *BGPSessionRequest) HasPeerGroup() bool`

HasPeerGroup returns a boolean if a field has been set.

### SetPeerGroupNil

`func (o *BGPSessionRequest) SetPeerGroupNil(b bool)`

 SetPeerGroupNil sets the value for PeerGroup to be an explicit nil

### UnsetPeerGroup
`func (o *BGPSessionRequest) UnsetPeerGroup()`

UnsetPeerGroup ensures that no value is present for PeerGroup, not even an explicit nil
### GetImportPolicies

`func (o *BGPSessionRequest) GetImportPolicies() []*int32`

GetImportPolicies returns the ImportPolicies field if non-nil, zero value otherwise.

### GetImportPoliciesOk

`func (o *BGPSessionRequest) GetImportPoliciesOk() (*[]*int32, bool)`

GetImportPoliciesOk returns a tuple with the ImportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicies

`func (o *BGPSessionRequest) SetImportPolicies(v []*int32)`

SetImportPolicies sets ImportPolicies field to given value.

### HasImportPolicies

`func (o *BGPSessionRequest) HasImportPolicies() bool`

HasImportPolicies returns a boolean if a field has been set.

### GetExportPolicies

`func (o *BGPSessionRequest) GetExportPolicies() []*int32`

GetExportPolicies returns the ExportPolicies field if non-nil, zero value otherwise.

### GetExportPoliciesOk

`func (o *BGPSessionRequest) GetExportPoliciesOk() (*[]*int32, bool)`

GetExportPoliciesOk returns a tuple with the ExportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicies

`func (o *BGPSessionRequest) SetExportPolicies(v []*int32)`

SetExportPolicies sets ExportPolicies field to given value.

### HasExportPolicies

`func (o *BGPSessionRequest) HasExportPolicies() bool`

HasExportPolicies returns a boolean if a field has been set.

### GetPrefixListIn

`func (o *BGPSessionRequest) GetPrefixListIn() BriefPrefixListRequest`

GetPrefixListIn returns the PrefixListIn field if non-nil, zero value otherwise.

### GetPrefixListInOk

`func (o *BGPSessionRequest) GetPrefixListInOk() (*BriefPrefixListRequest, bool)`

GetPrefixListInOk returns a tuple with the PrefixListIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListIn

`func (o *BGPSessionRequest) SetPrefixListIn(v BriefPrefixListRequest)`

SetPrefixListIn sets PrefixListIn field to given value.

### HasPrefixListIn

`func (o *BGPSessionRequest) HasPrefixListIn() bool`

HasPrefixListIn returns a boolean if a field has been set.

### SetPrefixListInNil

`func (o *BGPSessionRequest) SetPrefixListInNil(b bool)`

 SetPrefixListInNil sets the value for PrefixListIn to be an explicit nil

### UnsetPrefixListIn
`func (o *BGPSessionRequest) UnsetPrefixListIn()`

UnsetPrefixListIn ensures that no value is present for PrefixListIn, not even an explicit nil
### GetPrefixListOut

`func (o *BGPSessionRequest) GetPrefixListOut() BriefPrefixListRequest`

GetPrefixListOut returns the PrefixListOut field if non-nil, zero value otherwise.

### GetPrefixListOutOk

`func (o *BGPSessionRequest) GetPrefixListOutOk() (*BriefPrefixListRequest, bool)`

GetPrefixListOutOk returns a tuple with the PrefixListOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListOut

`func (o *BGPSessionRequest) SetPrefixListOut(v BriefPrefixListRequest)`

SetPrefixListOut sets PrefixListOut field to given value.

### HasPrefixListOut

`func (o *BGPSessionRequest) HasPrefixListOut() bool`

HasPrefixListOut returns a boolean if a field has been set.

### SetPrefixListOutNil

`func (o *BGPSessionRequest) SetPrefixListOutNil(b bool)`

 SetPrefixListOutNil sets the value for PrefixListOut to be an explicit nil

### UnsetPrefixListOut
`func (o *BGPSessionRequest) UnsetPrefixListOut()`

UnsetPrefixListOut ensures that no value is present for PrefixListOut, not even an explicit nil
### GetName

`func (o *BGPSessionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGPSessionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGPSessionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BGPSessionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BGPSessionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BGPSessionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BGPSessionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BGPSessionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BGPSessionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BGPSessionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *BGPSessionRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BGPSessionRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BGPSessionRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BGPSessionRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


