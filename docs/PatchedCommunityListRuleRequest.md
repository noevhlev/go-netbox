# PatchedCommunityListRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CommunityList** | Pointer to [**BriefCommunityListRequest**](BriefCommunityListRequest.md) |  | [optional] 
**Action** | Pointer to [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | [optional] 
**Community** | Pointer to [**NullableBriefCommunityRequest**](BriefCommunityRequest.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedCommunityListRuleRequest

`func NewPatchedCommunityListRuleRequest() *PatchedCommunityListRuleRequest`

NewPatchedCommunityListRuleRequest instantiates a new PatchedCommunityListRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCommunityListRuleRequestWithDefaults

`func NewPatchedCommunityListRuleRequestWithDefaults() *PatchedCommunityListRuleRequest`

NewPatchedCommunityListRuleRequestWithDefaults instantiates a new PatchedCommunityListRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PatchedCommunityListRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedCommunityListRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedCommunityListRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedCommunityListRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedCommunityListRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedCommunityListRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedCommunityListRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedCommunityListRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedCommunityListRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedCommunityListRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedCommunityListRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedCommunityListRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCommunityList

`func (o *PatchedCommunityListRuleRequest) GetCommunityList() BriefCommunityListRequest`

GetCommunityList returns the CommunityList field if non-nil, zero value otherwise.

### GetCommunityListOk

`func (o *PatchedCommunityListRuleRequest) GetCommunityListOk() (*BriefCommunityListRequest, bool)`

GetCommunityListOk returns a tuple with the CommunityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityList

`func (o *PatchedCommunityListRuleRequest) SetCommunityList(v BriefCommunityListRequest)`

SetCommunityList sets CommunityList field to given value.

### HasCommunityList

`func (o *PatchedCommunityListRuleRequest) HasCommunityList() bool`

HasCommunityList returns a boolean if a field has been set.

### GetAction

`func (o *PatchedCommunityListRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedCommunityListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedCommunityListRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedCommunityListRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCommunity

`func (o *PatchedCommunityListRuleRequest) GetCommunity() BriefCommunityRequest`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *PatchedCommunityListRuleRequest) GetCommunityOk() (*BriefCommunityRequest, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *PatchedCommunityListRuleRequest) SetCommunity(v BriefCommunityRequest)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *PatchedCommunityListRuleRequest) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### SetCommunityNil

`func (o *PatchedCommunityListRuleRequest) SetCommunityNil(b bool)`

 SetCommunityNil sets the value for Community to be an explicit nil

### UnsetCommunity
`func (o *PatchedCommunityListRuleRequest) UnsetCommunity()`

UnsetCommunity ensures that no value is present for Community, not even an explicit nil
### GetComments

`func (o *PatchedCommunityListRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedCommunityListRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedCommunityListRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedCommunityListRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


