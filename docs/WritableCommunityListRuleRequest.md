# WritableCommunityListRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CommunityList** | **int32** |  | 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**Community** | **int32** |  | 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewWritableCommunityListRuleRequest

`func NewWritableCommunityListRuleRequest(communityList int32, action CommunityListRuleAction, community int32, ) *WritableCommunityListRuleRequest`

NewWritableCommunityListRuleRequest instantiates a new WritableCommunityListRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCommunityListRuleRequestWithDefaults

`func NewWritableCommunityListRuleRequestWithDefaults() *WritableCommunityListRuleRequest`

NewWritableCommunityListRuleRequestWithDefaults instantiates a new WritableCommunityListRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *WritableCommunityListRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCommunityListRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCommunityListRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCommunityListRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCommunityListRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCommunityListRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCommunityListRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCommunityListRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCommunityList

`func (o *WritableCommunityListRuleRequest) GetCommunityList() int32`

GetCommunityList returns the CommunityList field if non-nil, zero value otherwise.

### GetCommunityListOk

`func (o *WritableCommunityListRuleRequest) GetCommunityListOk() (*int32, bool)`

GetCommunityListOk returns a tuple with the CommunityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityList

`func (o *WritableCommunityListRuleRequest) SetCommunityList(v int32)`

SetCommunityList sets CommunityList field to given value.


### GetAction

`func (o *WritableCommunityListRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WritableCommunityListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WritableCommunityListRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetCommunity

`func (o *WritableCommunityListRuleRequest) GetCommunity() int32`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *WritableCommunityListRuleRequest) GetCommunityOk() (*int32, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *WritableCommunityListRuleRequest) SetCommunity(v int32)`

SetCommunity sets Community field to given value.


### GetComments

`func (o *WritableCommunityListRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableCommunityListRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableCommunityListRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableCommunityListRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


