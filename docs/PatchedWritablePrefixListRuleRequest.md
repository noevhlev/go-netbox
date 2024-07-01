# PatchedWritablePrefixListRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**PrefixList** | Pointer to **int32** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Action** | Pointer to [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | [optional] 
**PrefixCustom** | Pointer to **NullableString** |  | [optional] 
**Ge** | Pointer to **NullableInt32** |  | [optional] 
**Le** | Pointer to **NullableInt32** |  | [optional] 
**Prefix** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedWritablePrefixListRuleRequest

`func NewPatchedWritablePrefixListRuleRequest() *PatchedWritablePrefixListRuleRequest`

NewPatchedWritablePrefixListRuleRequest instantiates a new PatchedWritablePrefixListRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePrefixListRuleRequestWithDefaults

`func NewPatchedWritablePrefixListRuleRequestWithDefaults() *PatchedWritablePrefixListRuleRequest`

NewPatchedWritablePrefixListRuleRequestWithDefaults instantiates a new PatchedWritablePrefixListRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PatchedWritablePrefixListRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritablePrefixListRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritablePrefixListRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritablePrefixListRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritablePrefixListRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePrefixListRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePrefixListRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePrefixListRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPrefixList

`func (o *PatchedWritablePrefixListRuleRequest) GetPrefixList() int32`

GetPrefixList returns the PrefixList field if non-nil, zero value otherwise.

### GetPrefixListOk

`func (o *PatchedWritablePrefixListRuleRequest) GetPrefixListOk() (*int32, bool)`

GetPrefixListOk returns a tuple with the PrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixList

`func (o *PatchedWritablePrefixListRuleRequest) SetPrefixList(v int32)`

SetPrefixList sets PrefixList field to given value.

### HasPrefixList

`func (o *PatchedWritablePrefixListRuleRequest) HasPrefixList() bool`

HasPrefixList returns a boolean if a field has been set.

### GetIndex

`func (o *PatchedWritablePrefixListRuleRequest) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PatchedWritablePrefixListRuleRequest) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PatchedWritablePrefixListRuleRequest) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PatchedWritablePrefixListRuleRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetAction

`func (o *PatchedWritablePrefixListRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedWritablePrefixListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedWritablePrefixListRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedWritablePrefixListRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPrefixCustom

`func (o *PatchedWritablePrefixListRuleRequest) GetPrefixCustom() string`

GetPrefixCustom returns the PrefixCustom field if non-nil, zero value otherwise.

### GetPrefixCustomOk

`func (o *PatchedWritablePrefixListRuleRequest) GetPrefixCustomOk() (*string, bool)`

GetPrefixCustomOk returns a tuple with the PrefixCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCustom

`func (o *PatchedWritablePrefixListRuleRequest) SetPrefixCustom(v string)`

SetPrefixCustom sets PrefixCustom field to given value.

### HasPrefixCustom

`func (o *PatchedWritablePrefixListRuleRequest) HasPrefixCustom() bool`

HasPrefixCustom returns a boolean if a field has been set.

### SetPrefixCustomNil

`func (o *PatchedWritablePrefixListRuleRequest) SetPrefixCustomNil(b bool)`

 SetPrefixCustomNil sets the value for PrefixCustom to be an explicit nil

### UnsetPrefixCustom
`func (o *PatchedWritablePrefixListRuleRequest) UnsetPrefixCustom()`

UnsetPrefixCustom ensures that no value is present for PrefixCustom, not even an explicit nil
### GetGe

`func (o *PatchedWritablePrefixListRuleRequest) GetGe() int32`

GetGe returns the Ge field if non-nil, zero value otherwise.

### GetGeOk

`func (o *PatchedWritablePrefixListRuleRequest) GetGeOk() (*int32, bool)`

GetGeOk returns a tuple with the Ge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGe

`func (o *PatchedWritablePrefixListRuleRequest) SetGe(v int32)`

SetGe sets Ge field to given value.

### HasGe

`func (o *PatchedWritablePrefixListRuleRequest) HasGe() bool`

HasGe returns a boolean if a field has been set.

### SetGeNil

`func (o *PatchedWritablePrefixListRuleRequest) SetGeNil(b bool)`

 SetGeNil sets the value for Ge to be an explicit nil

### UnsetGe
`func (o *PatchedWritablePrefixListRuleRequest) UnsetGe()`

UnsetGe ensures that no value is present for Ge, not even an explicit nil
### GetLe

`func (o *PatchedWritablePrefixListRuleRequest) GetLe() int32`

GetLe returns the Le field if non-nil, zero value otherwise.

### GetLeOk

`func (o *PatchedWritablePrefixListRuleRequest) GetLeOk() (*int32, bool)`

GetLeOk returns a tuple with the Le field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLe

`func (o *PatchedWritablePrefixListRuleRequest) SetLe(v int32)`

SetLe sets Le field to given value.

### HasLe

`func (o *PatchedWritablePrefixListRuleRequest) HasLe() bool`

HasLe returns a boolean if a field has been set.

### SetLeNil

`func (o *PatchedWritablePrefixListRuleRequest) SetLeNil(b bool)`

 SetLeNil sets the value for Le to be an explicit nil

### UnsetLe
`func (o *PatchedWritablePrefixListRuleRequest) UnsetLe()`

UnsetLe ensures that no value is present for Le, not even an explicit nil
### GetPrefix

`func (o *PatchedWritablePrefixListRuleRequest) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedWritablePrefixListRuleRequest) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedWritablePrefixListRuleRequest) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedWritablePrefixListRuleRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *PatchedWritablePrefixListRuleRequest) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *PatchedWritablePrefixListRuleRequest) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetComments

`func (o *PatchedWritablePrefixListRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritablePrefixListRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritablePrefixListRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritablePrefixListRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


