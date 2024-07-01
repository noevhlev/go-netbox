# CommunityListRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Display** | **string** |  | [readonly] 
**CommunityList** | [**NestedCommunityList**](NestedCommunityList.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**Community** | Pointer to [**NullableNestedCommunity**](NestedCommunity.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewCommunityListRule

`func NewCommunityListRule(id int32, display string, communityList NestedCommunityList, created NullableTime, lastUpdated NullableTime, action CommunityListRuleAction, ) *CommunityListRule`

NewCommunityListRule instantiates a new CommunityListRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunityListRuleWithDefaults

`func NewCommunityListRuleWithDefaults() *CommunityListRule`

NewCommunityListRuleWithDefaults instantiates a new CommunityListRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommunityListRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunityListRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunityListRule) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *CommunityListRule) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CommunityListRule) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CommunityListRule) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CommunityListRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CommunityListRule) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CommunityListRule) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CommunityListRule) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CommunityListRule) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDisplay

`func (o *CommunityListRule) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CommunityListRule) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CommunityListRule) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetCommunityList

`func (o *CommunityListRule) GetCommunityList() NestedCommunityList`

GetCommunityList returns the CommunityList field if non-nil, zero value otherwise.

### GetCommunityListOk

`func (o *CommunityListRule) GetCommunityListOk() (*NestedCommunityList, bool)`

GetCommunityListOk returns a tuple with the CommunityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityList

`func (o *CommunityListRule) SetCommunityList(v NestedCommunityList)`

SetCommunityList sets CommunityList field to given value.


### GetCreated

`func (o *CommunityListRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CommunityListRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CommunityListRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CommunityListRule) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CommunityListRule) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CommunityListRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CommunityListRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CommunityListRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CommunityListRule) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CommunityListRule) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetAction

`func (o *CommunityListRule) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CommunityListRule) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CommunityListRule) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetCommunity

`func (o *CommunityListRule) GetCommunity() NestedCommunity`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *CommunityListRule) GetCommunityOk() (*NestedCommunity, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *CommunityListRule) SetCommunity(v NestedCommunity)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *CommunityListRule) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### SetCommunityNil

`func (o *CommunityListRule) SetCommunityNil(b bool)`

 SetCommunityNil sets the value for Community to be an explicit nil

### UnsetCommunity
`func (o *CommunityListRule) UnsetCommunity()`

UnsetCommunity ensures that no value is present for Community, not even an explicit nil
### GetComments

`func (o *CommunityListRule) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CommunityListRule) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CommunityListRule) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CommunityListRule) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


