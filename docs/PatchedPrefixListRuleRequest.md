# PatchedPrefixListRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**PrefixList** | Pointer to [**BriefPrefixListRequest**](BriefPrefixListRequest.md) |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Action** | Pointer to [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | [optional] 
**PrefixCustom** | Pointer to **NullableString** |  | [optional] 
**Ge** | Pointer to **NullableInt32** |  | [optional] 
**Le** | Pointer to **NullableInt32** |  | [optional] 
**Prefix** | Pointer to [**NullableBriefPrefixRequest**](BriefPrefixRequest.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedPrefixListRuleRequest

`func NewPatchedPrefixListRuleRequest() *PatchedPrefixListRuleRequest`

NewPatchedPrefixListRuleRequest instantiates a new PatchedPrefixListRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPrefixListRuleRequestWithDefaults

`func NewPatchedPrefixListRuleRequestWithDefaults() *PatchedPrefixListRuleRequest`

NewPatchedPrefixListRuleRequestWithDefaults instantiates a new PatchedPrefixListRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PatchedPrefixListRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedPrefixListRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedPrefixListRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedPrefixListRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedPrefixListRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedPrefixListRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedPrefixListRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedPrefixListRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedPrefixListRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedPrefixListRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedPrefixListRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedPrefixListRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPrefixList

`func (o *PatchedPrefixListRuleRequest) GetPrefixList() BriefPrefixListRequest`

GetPrefixList returns the PrefixList field if non-nil, zero value otherwise.

### GetPrefixListOk

`func (o *PatchedPrefixListRuleRequest) GetPrefixListOk() (*BriefPrefixListRequest, bool)`

GetPrefixListOk returns a tuple with the PrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixList

`func (o *PatchedPrefixListRuleRequest) SetPrefixList(v BriefPrefixListRequest)`

SetPrefixList sets PrefixList field to given value.

### HasPrefixList

`func (o *PatchedPrefixListRuleRequest) HasPrefixList() bool`

HasPrefixList returns a boolean if a field has been set.

### GetIndex

`func (o *PatchedPrefixListRuleRequest) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PatchedPrefixListRuleRequest) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PatchedPrefixListRuleRequest) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PatchedPrefixListRuleRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetAction

`func (o *PatchedPrefixListRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedPrefixListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedPrefixListRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedPrefixListRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPrefixCustom

`func (o *PatchedPrefixListRuleRequest) GetPrefixCustom() string`

GetPrefixCustom returns the PrefixCustom field if non-nil, zero value otherwise.

### GetPrefixCustomOk

`func (o *PatchedPrefixListRuleRequest) GetPrefixCustomOk() (*string, bool)`

GetPrefixCustomOk returns a tuple with the PrefixCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCustom

`func (o *PatchedPrefixListRuleRequest) SetPrefixCustom(v string)`

SetPrefixCustom sets PrefixCustom field to given value.

### HasPrefixCustom

`func (o *PatchedPrefixListRuleRequest) HasPrefixCustom() bool`

HasPrefixCustom returns a boolean if a field has been set.

### SetPrefixCustomNil

`func (o *PatchedPrefixListRuleRequest) SetPrefixCustomNil(b bool)`

 SetPrefixCustomNil sets the value for PrefixCustom to be an explicit nil

### UnsetPrefixCustom
`func (o *PatchedPrefixListRuleRequest) UnsetPrefixCustom()`

UnsetPrefixCustom ensures that no value is present for PrefixCustom, not even an explicit nil
### GetGe

`func (o *PatchedPrefixListRuleRequest) GetGe() int32`

GetGe returns the Ge field if non-nil, zero value otherwise.

### GetGeOk

`func (o *PatchedPrefixListRuleRequest) GetGeOk() (*int32, bool)`

GetGeOk returns a tuple with the Ge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGe

`func (o *PatchedPrefixListRuleRequest) SetGe(v int32)`

SetGe sets Ge field to given value.

### HasGe

`func (o *PatchedPrefixListRuleRequest) HasGe() bool`

HasGe returns a boolean if a field has been set.

### SetGeNil

`func (o *PatchedPrefixListRuleRequest) SetGeNil(b bool)`

 SetGeNil sets the value for Ge to be an explicit nil

### UnsetGe
`func (o *PatchedPrefixListRuleRequest) UnsetGe()`

UnsetGe ensures that no value is present for Ge, not even an explicit nil
### GetLe

`func (o *PatchedPrefixListRuleRequest) GetLe() int32`

GetLe returns the Le field if non-nil, zero value otherwise.

### GetLeOk

`func (o *PatchedPrefixListRuleRequest) GetLeOk() (*int32, bool)`

GetLeOk returns a tuple with the Le field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLe

`func (o *PatchedPrefixListRuleRequest) SetLe(v int32)`

SetLe sets Le field to given value.

### HasLe

`func (o *PatchedPrefixListRuleRequest) HasLe() bool`

HasLe returns a boolean if a field has been set.

### SetLeNil

`func (o *PatchedPrefixListRuleRequest) SetLeNil(b bool)`

 SetLeNil sets the value for Le to be an explicit nil

### UnsetLe
`func (o *PatchedPrefixListRuleRequest) UnsetLe()`

UnsetLe ensures that no value is present for Le, not even an explicit nil
### GetPrefix

`func (o *PatchedPrefixListRuleRequest) GetPrefix() BriefPrefixRequest`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedPrefixListRuleRequest) GetPrefixOk() (*BriefPrefixRequest, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedPrefixListRuleRequest) SetPrefix(v BriefPrefixRequest)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedPrefixListRuleRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *PatchedPrefixListRuleRequest) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *PatchedPrefixListRuleRequest) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetComments

`func (o *PatchedPrefixListRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedPrefixListRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedPrefixListRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedPrefixListRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


