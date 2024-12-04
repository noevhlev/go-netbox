# PrefixListRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**PrefixList** | [**BriefPrefixListRequest**](BriefPrefixListRequest.md) |  | 
**Index** | **int32** |  | 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**PrefixCustom** | Pointer to **NullableString** |  | [optional] 
**Ge** | Pointer to **NullableInt32** |  | [optional] 
**Le** | Pointer to **NullableInt32** |  | [optional] 
**Prefix** | Pointer to [**NullableBriefPrefixRequest**](BriefPrefixRequest.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPrefixListRuleRequest

`func NewPrefixListRuleRequest(prefixList BriefPrefixListRequest, index int32, action CommunityListRuleAction, ) *PrefixListRuleRequest`

NewPrefixListRuleRequest instantiates a new PrefixListRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixListRuleRequestWithDefaults

`func NewPrefixListRuleRequestWithDefaults() *PrefixListRuleRequest`

NewPrefixListRuleRequestWithDefaults instantiates a new PrefixListRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PrefixListRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrefixListRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrefixListRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrefixListRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PrefixListRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrefixListRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrefixListRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PrefixListRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PrefixListRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PrefixListRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PrefixListRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PrefixListRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPrefixList

`func (o *PrefixListRuleRequest) GetPrefixList() BriefPrefixListRequest`

GetPrefixList returns the PrefixList field if non-nil, zero value otherwise.

### GetPrefixListOk

`func (o *PrefixListRuleRequest) GetPrefixListOk() (*BriefPrefixListRequest, bool)`

GetPrefixListOk returns a tuple with the PrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixList

`func (o *PrefixListRuleRequest) SetPrefixList(v BriefPrefixListRequest)`

SetPrefixList sets PrefixList field to given value.


### GetIndex

`func (o *PrefixListRuleRequest) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PrefixListRuleRequest) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PrefixListRuleRequest) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetAction

`func (o *PrefixListRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PrefixListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PrefixListRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetPrefixCustom

`func (o *PrefixListRuleRequest) GetPrefixCustom() string`

GetPrefixCustom returns the PrefixCustom field if non-nil, zero value otherwise.

### GetPrefixCustomOk

`func (o *PrefixListRuleRequest) GetPrefixCustomOk() (*string, bool)`

GetPrefixCustomOk returns a tuple with the PrefixCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCustom

`func (o *PrefixListRuleRequest) SetPrefixCustom(v string)`

SetPrefixCustom sets PrefixCustom field to given value.

### HasPrefixCustom

`func (o *PrefixListRuleRequest) HasPrefixCustom() bool`

HasPrefixCustom returns a boolean if a field has been set.

### SetPrefixCustomNil

`func (o *PrefixListRuleRequest) SetPrefixCustomNil(b bool)`

 SetPrefixCustomNil sets the value for PrefixCustom to be an explicit nil

### UnsetPrefixCustom
`func (o *PrefixListRuleRequest) UnsetPrefixCustom()`

UnsetPrefixCustom ensures that no value is present for PrefixCustom, not even an explicit nil
### GetGe

`func (o *PrefixListRuleRequest) GetGe() int32`

GetGe returns the Ge field if non-nil, zero value otherwise.

### GetGeOk

`func (o *PrefixListRuleRequest) GetGeOk() (*int32, bool)`

GetGeOk returns a tuple with the Ge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGe

`func (o *PrefixListRuleRequest) SetGe(v int32)`

SetGe sets Ge field to given value.

### HasGe

`func (o *PrefixListRuleRequest) HasGe() bool`

HasGe returns a boolean if a field has been set.

### SetGeNil

`func (o *PrefixListRuleRequest) SetGeNil(b bool)`

 SetGeNil sets the value for Ge to be an explicit nil

### UnsetGe
`func (o *PrefixListRuleRequest) UnsetGe()`

UnsetGe ensures that no value is present for Ge, not even an explicit nil
### GetLe

`func (o *PrefixListRuleRequest) GetLe() int32`

GetLe returns the Le field if non-nil, zero value otherwise.

### GetLeOk

`func (o *PrefixListRuleRequest) GetLeOk() (*int32, bool)`

GetLeOk returns a tuple with the Le field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLe

`func (o *PrefixListRuleRequest) SetLe(v int32)`

SetLe sets Le field to given value.

### HasLe

`func (o *PrefixListRuleRequest) HasLe() bool`

HasLe returns a boolean if a field has been set.

### SetLeNil

`func (o *PrefixListRuleRequest) SetLeNil(b bool)`

 SetLeNil sets the value for Le to be an explicit nil

### UnsetLe
`func (o *PrefixListRuleRequest) UnsetLe()`

UnsetLe ensures that no value is present for Le, not even an explicit nil
### GetPrefix

`func (o *PrefixListRuleRequest) GetPrefix() BriefPrefixRequest`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PrefixListRuleRequest) GetPrefixOk() (*BriefPrefixRequest, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PrefixListRuleRequest) SetPrefix(v BriefPrefixRequest)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PrefixListRuleRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *PrefixListRuleRequest) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *PrefixListRuleRequest) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetComments

`func (o *PrefixListRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PrefixListRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PrefixListRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PrefixListRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


