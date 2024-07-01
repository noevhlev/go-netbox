# PatchedWritableCommunityListRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CommunityList** | Pointer to **int32** |  | [optional] 
**Action** | Pointer to [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | [optional] 
**Community** | Pointer to **int32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedWritableCommunityListRuleRequest

`func NewPatchedWritableCommunityListRuleRequest() *PatchedWritableCommunityListRuleRequest`

NewPatchedWritableCommunityListRuleRequest instantiates a new PatchedWritableCommunityListRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCommunityListRuleRequestWithDefaults

`func NewPatchedWritableCommunityListRuleRequestWithDefaults() *PatchedWritableCommunityListRuleRequest`

NewPatchedWritableCommunityListRuleRequestWithDefaults instantiates a new PatchedWritableCommunityListRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PatchedWritableCommunityListRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableCommunityListRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableCommunityListRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableCommunityListRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableCommunityListRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableCommunityListRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableCommunityListRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableCommunityListRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCommunityList

`func (o *PatchedWritableCommunityListRuleRequest) GetCommunityList() int32`

GetCommunityList returns the CommunityList field if non-nil, zero value otherwise.

### GetCommunityListOk

`func (o *PatchedWritableCommunityListRuleRequest) GetCommunityListOk() (*int32, bool)`

GetCommunityListOk returns a tuple with the CommunityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityList

`func (o *PatchedWritableCommunityListRuleRequest) SetCommunityList(v int32)`

SetCommunityList sets CommunityList field to given value.

### HasCommunityList

`func (o *PatchedWritableCommunityListRuleRequest) HasCommunityList() bool`

HasCommunityList returns a boolean if a field has been set.

### GetAction

`func (o *PatchedWritableCommunityListRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedWritableCommunityListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedWritableCommunityListRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedWritableCommunityListRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCommunity

`func (o *PatchedWritableCommunityListRuleRequest) GetCommunity() int32`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *PatchedWritableCommunityListRuleRequest) GetCommunityOk() (*int32, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *PatchedWritableCommunityListRuleRequest) SetCommunity(v int32)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *PatchedWritableCommunityListRuleRequest) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableCommunityListRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableCommunityListRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableCommunityListRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableCommunityListRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


