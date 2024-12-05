# PrefixListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Family** | [**PatchedPrefixListRequestFamily**](PatchedPrefixListRequestFamily.md) |  | 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPrefixListRequest

`func NewPrefixListRequest(name string, description string, family PatchedPrefixListRequestFamily, ) *PrefixListRequest`

NewPrefixListRequest instantiates a new PrefixListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixListRequestWithDefaults

`func NewPrefixListRequestWithDefaults() *PrefixListRequest`

NewPrefixListRequestWithDefaults instantiates a new PrefixListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrefixListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrefixListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrefixListRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrefixListRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrefixListRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrefixListRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFamily

`func (o *PrefixListRequest) GetFamily() PatchedPrefixListRequestFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *PrefixListRequest) GetFamilyOk() (*PatchedPrefixListRequestFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *PrefixListRequest) SetFamily(v PatchedPrefixListRequestFamily)`

SetFamily sets Family field to given value.


### GetTags

`func (o *PrefixListRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrefixListRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrefixListRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PrefixListRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PrefixListRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PrefixListRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PrefixListRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PrefixListRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComments

`func (o *PrefixListRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PrefixListRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PrefixListRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PrefixListRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


