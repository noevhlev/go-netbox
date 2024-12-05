# PatchedPrefixListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Family** | Pointer to [**PatchedPrefixListRequestFamily**](PatchedPrefixListRequestFamily.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedPrefixListRequest

`func NewPatchedPrefixListRequest() *PatchedPrefixListRequest`

NewPatchedPrefixListRequest instantiates a new PatchedPrefixListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPrefixListRequestWithDefaults

`func NewPatchedPrefixListRequestWithDefaults() *PatchedPrefixListRequest`

NewPatchedPrefixListRequestWithDefaults instantiates a new PatchedPrefixListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedPrefixListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPrefixListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPrefixListRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPrefixListRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedPrefixListRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedPrefixListRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedPrefixListRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedPrefixListRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFamily

`func (o *PatchedPrefixListRequest) GetFamily() PatchedPrefixListRequestFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *PatchedPrefixListRequest) GetFamilyOk() (*PatchedPrefixListRequestFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *PatchedPrefixListRequest) SetFamily(v PatchedPrefixListRequestFamily)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *PatchedPrefixListRequest) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetTags

`func (o *PatchedPrefixListRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedPrefixListRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedPrefixListRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedPrefixListRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedPrefixListRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedPrefixListRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedPrefixListRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedPrefixListRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComments

`func (o *PatchedPrefixListRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedPrefixListRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedPrefixListRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedPrefixListRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


