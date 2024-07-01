# WritablePrefixListRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**PrefixList** | **int32** |  | 
**Index** | **int32** |  | 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**PrefixCustom** | Pointer to **NullableString** |  | [optional] 
**Ge** | Pointer to **NullableInt32** |  | [optional] 
**Le** | Pointer to **NullableInt32** |  | [optional] 
**Prefix** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewWritablePrefixListRuleRequest

`func NewWritablePrefixListRuleRequest(prefixList int32, index int32, action CommunityListRuleAction, ) *WritablePrefixListRuleRequest`

NewWritablePrefixListRuleRequest instantiates a new WritablePrefixListRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePrefixListRuleRequestWithDefaults

`func NewWritablePrefixListRuleRequestWithDefaults() *WritablePrefixListRuleRequest`

NewWritablePrefixListRuleRequestWithDefaults instantiates a new WritablePrefixListRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *WritablePrefixListRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePrefixListRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePrefixListRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePrefixListRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePrefixListRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePrefixListRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePrefixListRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePrefixListRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPrefixList

`func (o *WritablePrefixListRuleRequest) GetPrefixList() int32`

GetPrefixList returns the PrefixList field if non-nil, zero value otherwise.

### GetPrefixListOk

`func (o *WritablePrefixListRuleRequest) GetPrefixListOk() (*int32, bool)`

GetPrefixListOk returns a tuple with the PrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixList

`func (o *WritablePrefixListRuleRequest) SetPrefixList(v int32)`

SetPrefixList sets PrefixList field to given value.


### GetIndex

`func (o *WritablePrefixListRuleRequest) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WritablePrefixListRuleRequest) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WritablePrefixListRuleRequest) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetAction

`func (o *WritablePrefixListRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WritablePrefixListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WritablePrefixListRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetPrefixCustom

`func (o *WritablePrefixListRuleRequest) GetPrefixCustom() string`

GetPrefixCustom returns the PrefixCustom field if non-nil, zero value otherwise.

### GetPrefixCustomOk

`func (o *WritablePrefixListRuleRequest) GetPrefixCustomOk() (*string, bool)`

GetPrefixCustomOk returns a tuple with the PrefixCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCustom

`func (o *WritablePrefixListRuleRequest) SetPrefixCustom(v string)`

SetPrefixCustom sets PrefixCustom field to given value.

### HasPrefixCustom

`func (o *WritablePrefixListRuleRequest) HasPrefixCustom() bool`

HasPrefixCustom returns a boolean if a field has been set.

### SetPrefixCustomNil

`func (o *WritablePrefixListRuleRequest) SetPrefixCustomNil(b bool)`

 SetPrefixCustomNil sets the value for PrefixCustom to be an explicit nil

### UnsetPrefixCustom
`func (o *WritablePrefixListRuleRequest) UnsetPrefixCustom()`

UnsetPrefixCustom ensures that no value is present for PrefixCustom, not even an explicit nil
### GetGe

`func (o *WritablePrefixListRuleRequest) GetGe() int32`

GetGe returns the Ge field if non-nil, zero value otherwise.

### GetGeOk

`func (o *WritablePrefixListRuleRequest) GetGeOk() (*int32, bool)`

GetGeOk returns a tuple with the Ge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGe

`func (o *WritablePrefixListRuleRequest) SetGe(v int32)`

SetGe sets Ge field to given value.

### HasGe

`func (o *WritablePrefixListRuleRequest) HasGe() bool`

HasGe returns a boolean if a field has been set.

### SetGeNil

`func (o *WritablePrefixListRuleRequest) SetGeNil(b bool)`

 SetGeNil sets the value for Ge to be an explicit nil

### UnsetGe
`func (o *WritablePrefixListRuleRequest) UnsetGe()`

UnsetGe ensures that no value is present for Ge, not even an explicit nil
### GetLe

`func (o *WritablePrefixListRuleRequest) GetLe() int32`

GetLe returns the Le field if non-nil, zero value otherwise.

### GetLeOk

`func (o *WritablePrefixListRuleRequest) GetLeOk() (*int32, bool)`

GetLeOk returns a tuple with the Le field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLe

`func (o *WritablePrefixListRuleRequest) SetLe(v int32)`

SetLe sets Le field to given value.

### HasLe

`func (o *WritablePrefixListRuleRequest) HasLe() bool`

HasLe returns a boolean if a field has been set.

### SetLeNil

`func (o *WritablePrefixListRuleRequest) SetLeNil(b bool)`

 SetLeNil sets the value for Le to be an explicit nil

### UnsetLe
`func (o *WritablePrefixListRuleRequest) UnsetLe()`

UnsetLe ensures that no value is present for Le, not even an explicit nil
### GetPrefix

`func (o *WritablePrefixListRuleRequest) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WritablePrefixListRuleRequest) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WritablePrefixListRuleRequest) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *WritablePrefixListRuleRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *WritablePrefixListRuleRequest) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *WritablePrefixListRuleRequest) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetComments

`func (o *WritablePrefixListRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritablePrefixListRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritablePrefixListRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritablePrefixListRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


