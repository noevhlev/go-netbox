# CommunityListRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CommunityList** | [**BriefCommunityListRequest**](BriefCommunityListRequest.md) |  | 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**Community** | Pointer to [**NullableBriefCommunityRequest**](BriefCommunityRequest.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewCommunityListRuleRequest

`func NewCommunityListRuleRequest(communityList BriefCommunityListRequest, action CommunityListRuleAction, ) *CommunityListRuleRequest`

NewCommunityListRuleRequest instantiates a new CommunityListRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunityListRuleRequestWithDefaults

`func NewCommunityListRuleRequestWithDefaults() *CommunityListRuleRequest`

NewCommunityListRuleRequestWithDefaults instantiates a new CommunityListRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *CommunityListRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CommunityListRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CommunityListRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CommunityListRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CommunityListRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CommunityListRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CommunityListRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CommunityListRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDescription

`func (o *CommunityListRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommunityListRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommunityListRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommunityListRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCommunityList

`func (o *CommunityListRuleRequest) GetCommunityList() BriefCommunityListRequest`

GetCommunityList returns the CommunityList field if non-nil, zero value otherwise.

### GetCommunityListOk

`func (o *CommunityListRuleRequest) GetCommunityListOk() (*BriefCommunityListRequest, bool)`

GetCommunityListOk returns a tuple with the CommunityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityList

`func (o *CommunityListRuleRequest) SetCommunityList(v BriefCommunityListRequest)`

SetCommunityList sets CommunityList field to given value.


### GetAction

`func (o *CommunityListRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CommunityListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CommunityListRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetCommunity

`func (o *CommunityListRuleRequest) GetCommunity() BriefCommunityRequest`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *CommunityListRuleRequest) GetCommunityOk() (*BriefCommunityRequest, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *CommunityListRuleRequest) SetCommunity(v BriefCommunityRequest)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *CommunityListRuleRequest) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### SetCommunityNil

`func (o *CommunityListRuleRequest) SetCommunityNil(b bool)`

 SetCommunityNil sets the value for Community to be an explicit nil

### UnsetCommunity
`func (o *CommunityListRuleRequest) UnsetCommunity()`

UnsetCommunity ensures that no value is present for Community, not even an explicit nil
### GetComments

`func (o *CommunityListRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CommunityListRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CommunityListRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CommunityListRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


