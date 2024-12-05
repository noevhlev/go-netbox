# WritableBGPSessionRequest

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

### NewWritableBGPSessionRequest

`func NewWritableBGPSessionRequest(localAddress BriefIPAddressRequest, remoteAddress BriefIPAddressRequest, localAs BriefASNRequest, remoteAs BriefASNRequest, ) *WritableBGPSessionRequest`

NewWritableBGPSessionRequest instantiates a new WritableBGPSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableBGPSessionRequestWithDefaults

`func NewWritableBGPSessionRequestWithDefaults() *WritableBGPSessionRequest`

NewWritableBGPSessionRequestWithDefaults instantiates a new WritableBGPSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *WritableBGPSessionRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableBGPSessionRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableBGPSessionRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableBGPSessionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableBGPSessionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableBGPSessionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableBGPSessionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableBGPSessionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetStatus

`func (o *WritableBGPSessionRequest) GetStatus() BGPSessionStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableBGPSessionRequest) GetStatusOk() (*BGPSessionStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableBGPSessionRequest) SetStatus(v BGPSessionStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableBGPSessionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSite

`func (o *WritableBGPSessionRequest) GetSite() BriefSiteRequest`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritableBGPSessionRequest) GetSiteOk() (*BriefSiteRequest, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritableBGPSessionRequest) SetSite(v BriefSiteRequest)`

SetSite sets Site field to given value.

### HasSite

`func (o *WritableBGPSessionRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *WritableBGPSessionRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *WritableBGPSessionRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetTenant

`func (o *WritableBGPSessionRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableBGPSessionRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableBGPSessionRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableBGPSessionRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableBGPSessionRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableBGPSessionRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDevice

`func (o *WritableBGPSessionRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableBGPSessionRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableBGPSessionRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *WritableBGPSessionRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *WritableBGPSessionRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *WritableBGPSessionRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetLocalAddress

`func (o *WritableBGPSessionRequest) GetLocalAddress() BriefIPAddressRequest`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *WritableBGPSessionRequest) GetLocalAddressOk() (*BriefIPAddressRequest, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *WritableBGPSessionRequest) SetLocalAddress(v BriefIPAddressRequest)`

SetLocalAddress sets LocalAddress field to given value.


### GetRemoteAddress

`func (o *WritableBGPSessionRequest) GetRemoteAddress() BriefIPAddressRequest`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *WritableBGPSessionRequest) GetRemoteAddressOk() (*BriefIPAddressRequest, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *WritableBGPSessionRequest) SetRemoteAddress(v BriefIPAddressRequest)`

SetRemoteAddress sets RemoteAddress field to given value.


### GetLocalAs

`func (o *WritableBGPSessionRequest) GetLocalAs() BriefASNRequest`

GetLocalAs returns the LocalAs field if non-nil, zero value otherwise.

### GetLocalAsOk

`func (o *WritableBGPSessionRequest) GetLocalAsOk() (*BriefASNRequest, bool)`

GetLocalAsOk returns a tuple with the LocalAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAs

`func (o *WritableBGPSessionRequest) SetLocalAs(v BriefASNRequest)`

SetLocalAs sets LocalAs field to given value.


### GetRemoteAs

`func (o *WritableBGPSessionRequest) GetRemoteAs() BriefASNRequest`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *WritableBGPSessionRequest) GetRemoteAsOk() (*BriefASNRequest, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *WritableBGPSessionRequest) SetRemoteAs(v BriefASNRequest)`

SetRemoteAs sets RemoteAs field to given value.


### GetPeerGroup

`func (o *WritableBGPSessionRequest) GetPeerGroup() BriefBGPPeerGroupRequest`

GetPeerGroup returns the PeerGroup field if non-nil, zero value otherwise.

### GetPeerGroupOk

`func (o *WritableBGPSessionRequest) GetPeerGroupOk() (*BriefBGPPeerGroupRequest, bool)`

GetPeerGroupOk returns a tuple with the PeerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGroup

`func (o *WritableBGPSessionRequest) SetPeerGroup(v BriefBGPPeerGroupRequest)`

SetPeerGroup sets PeerGroup field to given value.

### HasPeerGroup

`func (o *WritableBGPSessionRequest) HasPeerGroup() bool`

HasPeerGroup returns a boolean if a field has been set.

### SetPeerGroupNil

`func (o *WritableBGPSessionRequest) SetPeerGroupNil(b bool)`

 SetPeerGroupNil sets the value for PeerGroup to be an explicit nil

### UnsetPeerGroup
`func (o *WritableBGPSessionRequest) UnsetPeerGroup()`

UnsetPeerGroup ensures that no value is present for PeerGroup, not even an explicit nil
### GetImportPolicies

`func (o *WritableBGPSessionRequest) GetImportPolicies() []*int32`

GetImportPolicies returns the ImportPolicies field if non-nil, zero value otherwise.

### GetImportPoliciesOk

`func (o *WritableBGPSessionRequest) GetImportPoliciesOk() (*[]*int32, bool)`

GetImportPoliciesOk returns a tuple with the ImportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicies

`func (o *WritableBGPSessionRequest) SetImportPolicies(v []*int32)`

SetImportPolicies sets ImportPolicies field to given value.

### HasImportPolicies

`func (o *WritableBGPSessionRequest) HasImportPolicies() bool`

HasImportPolicies returns a boolean if a field has been set.

### GetExportPolicies

`func (o *WritableBGPSessionRequest) GetExportPolicies() []*int32`

GetExportPolicies returns the ExportPolicies field if non-nil, zero value otherwise.

### GetExportPoliciesOk

`func (o *WritableBGPSessionRequest) GetExportPoliciesOk() (*[]*int32, bool)`

GetExportPoliciesOk returns a tuple with the ExportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicies

`func (o *WritableBGPSessionRequest) SetExportPolicies(v []*int32)`

SetExportPolicies sets ExportPolicies field to given value.

### HasExportPolicies

`func (o *WritableBGPSessionRequest) HasExportPolicies() bool`

HasExportPolicies returns a boolean if a field has been set.

### GetPrefixListIn

`func (o *WritableBGPSessionRequest) GetPrefixListIn() BriefPrefixListRequest`

GetPrefixListIn returns the PrefixListIn field if non-nil, zero value otherwise.

### GetPrefixListInOk

`func (o *WritableBGPSessionRequest) GetPrefixListInOk() (*BriefPrefixListRequest, bool)`

GetPrefixListInOk returns a tuple with the PrefixListIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListIn

`func (o *WritableBGPSessionRequest) SetPrefixListIn(v BriefPrefixListRequest)`

SetPrefixListIn sets PrefixListIn field to given value.

### HasPrefixListIn

`func (o *WritableBGPSessionRequest) HasPrefixListIn() bool`

HasPrefixListIn returns a boolean if a field has been set.

### SetPrefixListInNil

`func (o *WritableBGPSessionRequest) SetPrefixListInNil(b bool)`

 SetPrefixListInNil sets the value for PrefixListIn to be an explicit nil

### UnsetPrefixListIn
`func (o *WritableBGPSessionRequest) UnsetPrefixListIn()`

UnsetPrefixListIn ensures that no value is present for PrefixListIn, not even an explicit nil
### GetPrefixListOut

`func (o *WritableBGPSessionRequest) GetPrefixListOut() BriefPrefixListRequest`

GetPrefixListOut returns the PrefixListOut field if non-nil, zero value otherwise.

### GetPrefixListOutOk

`func (o *WritableBGPSessionRequest) GetPrefixListOutOk() (*BriefPrefixListRequest, bool)`

GetPrefixListOutOk returns a tuple with the PrefixListOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListOut

`func (o *WritableBGPSessionRequest) SetPrefixListOut(v BriefPrefixListRequest)`

SetPrefixListOut sets PrefixListOut field to given value.

### HasPrefixListOut

`func (o *WritableBGPSessionRequest) HasPrefixListOut() bool`

HasPrefixListOut returns a boolean if a field has been set.

### SetPrefixListOutNil

`func (o *WritableBGPSessionRequest) SetPrefixListOutNil(b bool)`

 SetPrefixListOutNil sets the value for PrefixListOut to be an explicit nil

### UnsetPrefixListOut
`func (o *WritableBGPSessionRequest) UnsetPrefixListOut()`

UnsetPrefixListOut ensures that no value is present for PrefixListOut, not even an explicit nil
### GetName

`func (o *WritableBGPSessionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableBGPSessionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableBGPSessionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WritableBGPSessionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WritableBGPSessionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WritableBGPSessionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *WritableBGPSessionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableBGPSessionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableBGPSessionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableBGPSessionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableBGPSessionRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableBGPSessionRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableBGPSessionRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableBGPSessionRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


