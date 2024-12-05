# BGPPeerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**ImportPolicies** | Pointer to [**[]RoutingPolicy**](RoutingPolicy.md) |  | [optional] 
**ExportPolicies** | Pointer to [**[]RoutingPolicy**](RoutingPolicy.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewBGPPeerGroup

`func NewBGPPeerGroup(id int32, url string, display string, name string, description string, ) *BGPPeerGroup`

NewBGPPeerGroup instantiates a new BGPPeerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPPeerGroupWithDefaults

`func NewBGPPeerGroupWithDefaults() *BGPPeerGroup`

NewBGPPeerGroupWithDefaults instantiates a new BGPPeerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BGPPeerGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BGPPeerGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BGPPeerGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BGPPeerGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BGPPeerGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BGPPeerGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BGPPeerGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BGPPeerGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BGPPeerGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BGPPeerGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGPPeerGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGPPeerGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BGPPeerGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BGPPeerGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BGPPeerGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImportPolicies

`func (o *BGPPeerGroup) GetImportPolicies() []*RoutingPolicy`

GetImportPolicies returns the ImportPolicies field if non-nil, zero value otherwise.

### GetImportPoliciesOk

`func (o *BGPPeerGroup) GetImportPoliciesOk() (*[]*RoutingPolicy, bool)`

GetImportPoliciesOk returns a tuple with the ImportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicies

`func (o *BGPPeerGroup) SetImportPolicies(v []*RoutingPolicy)`

SetImportPolicies sets ImportPolicies field to given value.

### HasImportPolicies

`func (o *BGPPeerGroup) HasImportPolicies() bool`

HasImportPolicies returns a boolean if a field has been set.

### GetExportPolicies

`func (o *BGPPeerGroup) GetExportPolicies() []*RoutingPolicy`

GetExportPolicies returns the ExportPolicies field if non-nil, zero value otherwise.

### GetExportPoliciesOk

`func (o *BGPPeerGroup) GetExportPoliciesOk() (*[]*RoutingPolicy, bool)`

GetExportPoliciesOk returns a tuple with the ExportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicies

`func (o *BGPPeerGroup) SetExportPolicies(v []*RoutingPolicy)`

SetExportPolicies sets ExportPolicies field to given value.

### HasExportPolicies

`func (o *BGPPeerGroup) HasExportPolicies() bool`

HasExportPolicies returns a boolean if a field has been set.

### GetComments

`func (o *BGPPeerGroup) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BGPPeerGroup) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BGPPeerGroup) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BGPPeerGroup) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


