# BGPSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Display** | **string** |  | [readonly] 
**Status** | Pointer to [**BGPSessionStatus**](BGPSessionStatus.md) |  | [optional] 
**Site** | Pointer to [**NullableBriefSite**](BriefSite.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenant**](BriefTenant.md) |  | [optional] 
**Device** | Pointer to [**NullableBriefDevice**](BriefDevice.md) |  | [optional] 
**LocalAddress** | [**BriefIPAddress**](BriefIPAddress.md) |  | 
**RemoteAddress** | [**BriefIPAddress**](BriefIPAddress.md) |  | 
**LocalAs** | [**BriefASN**](BriefASN.md) |  | 
**RemoteAs** | [**BriefASN**](BriefASN.md) |  | 
**PeerGroup** | Pointer to [**NullableBriefBGPPeerGroup**](BriefBGPPeerGroup.md) |  | [optional] 
**ImportPolicies** | Pointer to [**[]RoutingPolicy**](RoutingPolicy.md) |  | [optional] 
**ExportPolicies** | Pointer to [**[]RoutingPolicy**](RoutingPolicy.md) |  | [optional] 
**PrefixListIn** | Pointer to [**NullableBriefPrefixList**](BriefPrefixList.md) |  | [optional] 
**PrefixListOut** | Pointer to [**NullableBriefPrefixList**](BriefPrefixList.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewBGPSession

`func NewBGPSession(id int32, url string, display string, localAddress BriefIPAddress, remoteAddress BriefIPAddress, localAs BriefASN, remoteAs BriefASN, created NullableTime, lastUpdated NullableTime, ) *BGPSession`

NewBGPSession instantiates a new BGPSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPSessionWithDefaults

`func NewBGPSessionWithDefaults() *BGPSession`

NewBGPSessionWithDefaults instantiates a new BGPSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BGPSession) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BGPSession) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BGPSession) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BGPSession) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BGPSession) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BGPSession) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTags

`func (o *BGPSession) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BGPSession) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BGPSession) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BGPSession) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BGPSession) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BGPSession) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BGPSession) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BGPSession) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDisplay

`func (o *BGPSession) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BGPSession) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BGPSession) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetStatus

`func (o *BGPSession) GetStatus() BGPSessionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BGPSession) GetStatusOk() (*BGPSessionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BGPSession) SetStatus(v BGPSessionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BGPSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSite

`func (o *BGPSession) GetSite() BriefSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *BGPSession) GetSiteOk() (*BriefSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *BGPSession) SetSite(v BriefSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *BGPSession) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *BGPSession) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *BGPSession) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetTenant

`func (o *BGPSession) GetTenant() BriefTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BGPSession) GetTenantOk() (*BriefTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BGPSession) SetTenant(v BriefTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BGPSession) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BGPSession) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BGPSession) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDevice

`func (o *BGPSession) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BGPSession) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BGPSession) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BGPSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *BGPSession) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *BGPSession) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetLocalAddress

`func (o *BGPSession) GetLocalAddress() BriefIPAddress`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *BGPSession) GetLocalAddressOk() (*BriefIPAddress, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *BGPSession) SetLocalAddress(v BriefIPAddress)`

SetLocalAddress sets LocalAddress field to given value.


### GetRemoteAddress

`func (o *BGPSession) GetRemoteAddress() BriefIPAddress`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *BGPSession) GetRemoteAddressOk() (*BriefIPAddress, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *BGPSession) SetRemoteAddress(v BriefIPAddress)`

SetRemoteAddress sets RemoteAddress field to given value.


### GetLocalAs

`func (o *BGPSession) GetLocalAs() BriefASN`

GetLocalAs returns the LocalAs field if non-nil, zero value otherwise.

### GetLocalAsOk

`func (o *BGPSession) GetLocalAsOk() (*BriefASN, bool)`

GetLocalAsOk returns a tuple with the LocalAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAs

`func (o *BGPSession) SetLocalAs(v BriefASN)`

SetLocalAs sets LocalAs field to given value.


### GetRemoteAs

`func (o *BGPSession) GetRemoteAs() BriefASN`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *BGPSession) GetRemoteAsOk() (*BriefASN, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *BGPSession) SetRemoteAs(v BriefASN)`

SetRemoteAs sets RemoteAs field to given value.


### GetPeerGroup

`func (o *BGPSession) GetPeerGroup() BriefBGPPeerGroup`

GetPeerGroup returns the PeerGroup field if non-nil, zero value otherwise.

### GetPeerGroupOk

`func (o *BGPSession) GetPeerGroupOk() (*BriefBGPPeerGroup, bool)`

GetPeerGroupOk returns a tuple with the PeerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGroup

`func (o *BGPSession) SetPeerGroup(v BriefBGPPeerGroup)`

SetPeerGroup sets PeerGroup field to given value.

### HasPeerGroup

`func (o *BGPSession) HasPeerGroup() bool`

HasPeerGroup returns a boolean if a field has been set.

### SetPeerGroupNil

`func (o *BGPSession) SetPeerGroupNil(b bool)`

 SetPeerGroupNil sets the value for PeerGroup to be an explicit nil

### UnsetPeerGroup
`func (o *BGPSession) UnsetPeerGroup()`

UnsetPeerGroup ensures that no value is present for PeerGroup, not even an explicit nil
### GetImportPolicies

`func (o *BGPSession) GetImportPolicies() []*RoutingPolicy`

GetImportPolicies returns the ImportPolicies field if non-nil, zero value otherwise.

### GetImportPoliciesOk

`func (o *BGPSession) GetImportPoliciesOk() (*[]*RoutingPolicy, bool)`

GetImportPoliciesOk returns a tuple with the ImportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicies

`func (o *BGPSession) SetImportPolicies(v []*RoutingPolicy)`

SetImportPolicies sets ImportPolicies field to given value.

### HasImportPolicies

`func (o *BGPSession) HasImportPolicies() bool`

HasImportPolicies returns a boolean if a field has been set.

### GetExportPolicies

`func (o *BGPSession) GetExportPolicies() []*RoutingPolicy`

GetExportPolicies returns the ExportPolicies field if non-nil, zero value otherwise.

### GetExportPoliciesOk

`func (o *BGPSession) GetExportPoliciesOk() (*[]*RoutingPolicy, bool)`

GetExportPoliciesOk returns a tuple with the ExportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicies

`func (o *BGPSession) SetExportPolicies(v []*RoutingPolicy)`

SetExportPolicies sets ExportPolicies field to given value.

### HasExportPolicies

`func (o *BGPSession) HasExportPolicies() bool`

HasExportPolicies returns a boolean if a field has been set.

### GetPrefixListIn

`func (o *BGPSession) GetPrefixListIn() BriefPrefixList`

GetPrefixListIn returns the PrefixListIn field if non-nil, zero value otherwise.

### GetPrefixListInOk

`func (o *BGPSession) GetPrefixListInOk() (*BriefPrefixList, bool)`

GetPrefixListInOk returns a tuple with the PrefixListIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListIn

`func (o *BGPSession) SetPrefixListIn(v BriefPrefixList)`

SetPrefixListIn sets PrefixListIn field to given value.

### HasPrefixListIn

`func (o *BGPSession) HasPrefixListIn() bool`

HasPrefixListIn returns a boolean if a field has been set.

### SetPrefixListInNil

`func (o *BGPSession) SetPrefixListInNil(b bool)`

 SetPrefixListInNil sets the value for PrefixListIn to be an explicit nil

### UnsetPrefixListIn
`func (o *BGPSession) UnsetPrefixListIn()`

UnsetPrefixListIn ensures that no value is present for PrefixListIn, not even an explicit nil
### GetPrefixListOut

`func (o *BGPSession) GetPrefixListOut() BriefPrefixList`

GetPrefixListOut returns the PrefixListOut field if non-nil, zero value otherwise.

### GetPrefixListOutOk

`func (o *BGPSession) GetPrefixListOutOk() (*BriefPrefixList, bool)`

GetPrefixListOutOk returns a tuple with the PrefixListOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListOut

`func (o *BGPSession) SetPrefixListOut(v BriefPrefixList)`

SetPrefixListOut sets PrefixListOut field to given value.

### HasPrefixListOut

`func (o *BGPSession) HasPrefixListOut() bool`

HasPrefixListOut returns a boolean if a field has been set.

### SetPrefixListOutNil

`func (o *BGPSession) SetPrefixListOutNil(b bool)`

 SetPrefixListOutNil sets the value for PrefixListOut to be an explicit nil

### UnsetPrefixListOut
`func (o *BGPSession) UnsetPrefixListOut()`

UnsetPrefixListOut ensures that no value is present for PrefixListOut, not even an explicit nil
### GetCreated

`func (o *BGPSession) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BGPSession) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BGPSession) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *BGPSession) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *BGPSession) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *BGPSession) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BGPSession) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BGPSession) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *BGPSession) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *BGPSession) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetName

`func (o *BGPSession) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGPSession) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGPSession) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BGPSession) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BGPSession) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BGPSession) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BGPSession) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BGPSession) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BGPSession) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BGPSession) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *BGPSession) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BGPSession) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BGPSession) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BGPSession) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


