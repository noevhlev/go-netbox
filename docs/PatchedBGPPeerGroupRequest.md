# PatchedBGPPeerGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ImportPolicies** | Pointer to **[]int32** |  | [optional] 
**ExportPolicies** | Pointer to **[]int32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedBGPPeerGroupRequest

`func NewPatchedBGPPeerGroupRequest() *PatchedBGPPeerGroupRequest`

NewPatchedBGPPeerGroupRequest instantiates a new PatchedBGPPeerGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBGPPeerGroupRequestWithDefaults

`func NewPatchedBGPPeerGroupRequestWithDefaults() *PatchedBGPPeerGroupRequest`

NewPatchedBGPPeerGroupRequestWithDefaults instantiates a new PatchedBGPPeerGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedBGPPeerGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBGPPeerGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBGPPeerGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBGPPeerGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBGPPeerGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBGPPeerGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBGPPeerGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBGPPeerGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImportPolicies

`func (o *PatchedBGPPeerGroupRequest) GetImportPolicies() []*int32`

GetImportPolicies returns the ImportPolicies field if non-nil, zero value otherwise.

### GetImportPoliciesOk

`func (o *PatchedBGPPeerGroupRequest) GetImportPoliciesOk() (*[]*int32, bool)`

GetImportPoliciesOk returns a tuple with the ImportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicies

`func (o *PatchedBGPPeerGroupRequest) SetImportPolicies(v []*int32)`

SetImportPolicies sets ImportPolicies field to given value.

### HasImportPolicies

`func (o *PatchedBGPPeerGroupRequest) HasImportPolicies() bool`

HasImportPolicies returns a boolean if a field has been set.

### GetExportPolicies

`func (o *PatchedBGPPeerGroupRequest) GetExportPolicies() []*int32`

GetExportPolicies returns the ExportPolicies field if non-nil, zero value otherwise.

### GetExportPoliciesOk

`func (o *PatchedBGPPeerGroupRequest) GetExportPoliciesOk() (*[]*int32, bool)`

GetExportPoliciesOk returns a tuple with the ExportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicies

`func (o *PatchedBGPPeerGroupRequest) SetExportPolicies(v []*int32)`

SetExportPolicies sets ExportPolicies field to given value.

### HasExportPolicies

`func (o *PatchedBGPPeerGroupRequest) HasExportPolicies() bool`

HasExportPolicies returns a boolean if a field has been set.

### GetComments

`func (o *PatchedBGPPeerGroupRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBGPPeerGroupRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBGPPeerGroupRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBGPPeerGroupRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


