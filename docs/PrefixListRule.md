# PrefixListRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Display** | **string** |  | [readonly] 
**PrefixList** | [**NestedPrefixList**](NestedPrefixList.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Index** | **int32** |  | 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**PrefixCustom** | Pointer to **NullableString** |  | [optional] 
**Ge** | Pointer to **NullableInt32** |  | [optional] 
**Le** | Pointer to **NullableInt32** |  | [optional] 
**Prefix** | Pointer to [**NullableNestedPrefix**](NestedPrefix.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPrefixListRule

`func NewPrefixListRule(id int32, display string, prefixList NestedPrefixList, created NullableTime, lastUpdated NullableTime, index int32, action CommunityListRuleAction, ) *PrefixListRule`

NewPrefixListRule instantiates a new PrefixListRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixListRuleWithDefaults

`func NewPrefixListRuleWithDefaults() *PrefixListRule`

NewPrefixListRuleWithDefaults instantiates a new PrefixListRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrefixListRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrefixListRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrefixListRule) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *PrefixListRule) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrefixListRule) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrefixListRule) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PrefixListRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PrefixListRule) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PrefixListRule) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PrefixListRule) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PrefixListRule) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDisplay

`func (o *PrefixListRule) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PrefixListRule) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PrefixListRule) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetPrefixList

`func (o *PrefixListRule) GetPrefixList() NestedPrefixList`

GetPrefixList returns the PrefixList field if non-nil, zero value otherwise.

### GetPrefixListOk

`func (o *PrefixListRule) GetPrefixListOk() (*NestedPrefixList, bool)`

GetPrefixListOk returns a tuple with the PrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixList

`func (o *PrefixListRule) SetPrefixList(v NestedPrefixList)`

SetPrefixList sets PrefixList field to given value.


### GetCreated

`func (o *PrefixListRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PrefixListRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PrefixListRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *PrefixListRule) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PrefixListRule) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *PrefixListRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PrefixListRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PrefixListRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *PrefixListRule) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *PrefixListRule) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetIndex

`func (o *PrefixListRule) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PrefixListRule) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PrefixListRule) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetAction

`func (o *PrefixListRule) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PrefixListRule) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PrefixListRule) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetPrefixCustom

`func (o *PrefixListRule) GetPrefixCustom() string`

GetPrefixCustom returns the PrefixCustom field if non-nil, zero value otherwise.

### GetPrefixCustomOk

`func (o *PrefixListRule) GetPrefixCustomOk() (*string, bool)`

GetPrefixCustomOk returns a tuple with the PrefixCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCustom

`func (o *PrefixListRule) SetPrefixCustom(v string)`

SetPrefixCustom sets PrefixCustom field to given value.

### HasPrefixCustom

`func (o *PrefixListRule) HasPrefixCustom() bool`

HasPrefixCustom returns a boolean if a field has been set.

### SetPrefixCustomNil

`func (o *PrefixListRule) SetPrefixCustomNil(b bool)`

 SetPrefixCustomNil sets the value for PrefixCustom to be an explicit nil

### UnsetPrefixCustom
`func (o *PrefixListRule) UnsetPrefixCustom()`

UnsetPrefixCustom ensures that no value is present for PrefixCustom, not even an explicit nil
### GetGe

`func (o *PrefixListRule) GetGe() int32`

GetGe returns the Ge field if non-nil, zero value otherwise.

### GetGeOk

`func (o *PrefixListRule) GetGeOk() (*int32, bool)`

GetGeOk returns a tuple with the Ge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGe

`func (o *PrefixListRule) SetGe(v int32)`

SetGe sets Ge field to given value.

### HasGe

`func (o *PrefixListRule) HasGe() bool`

HasGe returns a boolean if a field has been set.

### SetGeNil

`func (o *PrefixListRule) SetGeNil(b bool)`

 SetGeNil sets the value for Ge to be an explicit nil

### UnsetGe
`func (o *PrefixListRule) UnsetGe()`

UnsetGe ensures that no value is present for Ge, not even an explicit nil
### GetLe

`func (o *PrefixListRule) GetLe() int32`

GetLe returns the Le field if non-nil, zero value otherwise.

### GetLeOk

`func (o *PrefixListRule) GetLeOk() (*int32, bool)`

GetLeOk returns a tuple with the Le field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLe

`func (o *PrefixListRule) SetLe(v int32)`

SetLe sets Le field to given value.

### HasLe

`func (o *PrefixListRule) HasLe() bool`

HasLe returns a boolean if a field has been set.

### SetLeNil

`func (o *PrefixListRule) SetLeNil(b bool)`

 SetLeNil sets the value for Le to be an explicit nil

### UnsetLe
`func (o *PrefixListRule) UnsetLe()`

UnsetLe ensures that no value is present for Le, not even an explicit nil
### GetPrefix

`func (o *PrefixListRule) GetPrefix() NestedPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PrefixListRule) GetPrefixOk() (*NestedPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PrefixListRule) SetPrefix(v NestedPrefix)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PrefixListRule) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *PrefixListRule) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *PrefixListRule) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetComments

`func (o *PrefixListRule) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PrefixListRule) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PrefixListRule) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PrefixListRule) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


