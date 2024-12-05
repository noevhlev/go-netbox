# BGPPeerGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**ImportPolicies** | Pointer to **[]int32** |  | [optional] 
**ExportPolicies** | Pointer to **[]int32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewBGPPeerGroupRequest

`func NewBGPPeerGroupRequest(name string, description string, ) *BGPPeerGroupRequest`

NewBGPPeerGroupRequest instantiates a new BGPPeerGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPPeerGroupRequestWithDefaults

`func NewBGPPeerGroupRequestWithDefaults() *BGPPeerGroupRequest`

NewBGPPeerGroupRequestWithDefaults instantiates a new BGPPeerGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BGPPeerGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGPPeerGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGPPeerGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BGPPeerGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BGPPeerGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BGPPeerGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImportPolicies

`func (o *BGPPeerGroupRequest) GetImportPolicies() []*int32`

GetImportPolicies returns the ImportPolicies field if non-nil, zero value otherwise.

### GetImportPoliciesOk

`func (o *BGPPeerGroupRequest) GetImportPoliciesOk() (*[]*int32, bool)`

GetImportPoliciesOk returns a tuple with the ImportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicies

`func (o *BGPPeerGroupRequest) SetImportPolicies(v []*int32)`

SetImportPolicies sets ImportPolicies field to given value.

### HasImportPolicies

`func (o *BGPPeerGroupRequest) HasImportPolicies() bool`

HasImportPolicies returns a boolean if a field has been set.

### GetExportPolicies

`func (o *BGPPeerGroupRequest) GetExportPolicies() []*int32`

GetExportPolicies returns the ExportPolicies field if non-nil, zero value otherwise.

### GetExportPoliciesOk

`func (o *BGPPeerGroupRequest) GetExportPoliciesOk() (*[]*int32, bool)`

GetExportPoliciesOk returns a tuple with the ExportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicies

`func (o *BGPPeerGroupRequest) SetExportPolicies(v []*int32)`

SetExportPolicies sets ExportPolicies field to given value.

### HasExportPolicies

`func (o *BGPPeerGroupRequest) HasExportPolicies() bool`

HasExportPolicies returns a boolean if a field has been set.

### GetComments

`func (o *BGPPeerGroupRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BGPPeerGroupRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BGPPeerGroupRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BGPPeerGroupRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


