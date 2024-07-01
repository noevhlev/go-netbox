# PatchedWritableRoutingPolicyRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Action** | Pointer to [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | [optional] 
**MatchIpAddress** | Pointer to **[]int32** |  | [optional] 
**RoutingPolicy** | Pointer to **int32** |  | [optional] 
**MatchCommunity** | Pointer to **[]int32** |  | [optional] 
**MatchCustom** | Pointer to **interface{}** |  | [optional] 
**SetActions** | Pointer to **interface{}** |  | [optional] 
**MatchIpv6Address** | Pointer to **[]int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedWritableRoutingPolicyRuleRequest

`func NewPatchedWritableRoutingPolicyRuleRequest() *PatchedWritableRoutingPolicyRuleRequest`

NewPatchedWritableRoutingPolicyRuleRequest instantiates a new PatchedWritableRoutingPolicyRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableRoutingPolicyRuleRequestWithDefaults

`func NewPatchedWritableRoutingPolicyRuleRequestWithDefaults() *PatchedWritableRoutingPolicyRuleRequest`

NewPatchedWritableRoutingPolicyRuleRequestWithDefaults instantiates a new PatchedWritableRoutingPolicyRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetAction

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMatchIpAddress

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetMatchIpAddress() []*int32`

GetMatchIpAddress returns the MatchIpAddress field if non-nil, zero value otherwise.

### GetMatchIpAddressOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetMatchIpAddressOk() (*[]*int32, bool)`

GetMatchIpAddressOk returns a tuple with the MatchIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpAddress

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetMatchIpAddress(v []*int32)`

SetMatchIpAddress sets MatchIpAddress field to given value.

### HasMatchIpAddress

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasMatchIpAddress() bool`

HasMatchIpAddress returns a boolean if a field has been set.

### GetRoutingPolicy

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetRoutingPolicy() int32`

GetRoutingPolicy returns the RoutingPolicy field if non-nil, zero value otherwise.

### GetRoutingPolicyOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetRoutingPolicyOk() (*int32, bool)`

GetRoutingPolicyOk returns a tuple with the RoutingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicy

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetRoutingPolicy(v int32)`

SetRoutingPolicy sets RoutingPolicy field to given value.

### HasRoutingPolicy

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasRoutingPolicy() bool`

HasRoutingPolicy returns a boolean if a field has been set.

### GetMatchCommunity

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetMatchCommunity() []*int32`

GetMatchCommunity returns the MatchCommunity field if non-nil, zero value otherwise.

### GetMatchCommunityOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetMatchCommunityOk() (*[]*int32, bool)`

GetMatchCommunityOk returns a tuple with the MatchCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCommunity

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetMatchCommunity(v []*int32)`

SetMatchCommunity sets MatchCommunity field to given value.

### HasMatchCommunity

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasMatchCommunity() bool`

HasMatchCommunity returns a boolean if a field has been set.

### GetMatchCustom

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetMatchCustom() interface{}`

GetMatchCustom returns the MatchCustom field if non-nil, zero value otherwise.

### GetMatchCustomOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetMatchCustomOk() (*interface{}, bool)`

GetMatchCustomOk returns a tuple with the MatchCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCustom

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetMatchCustom(v interface{})`

SetMatchCustom sets MatchCustom field to given value.

### HasMatchCustom

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasMatchCustom() bool`

HasMatchCustom returns a boolean if a field has been set.

### SetMatchCustomNil

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetMatchCustomNil(b bool)`

 SetMatchCustomNil sets the value for MatchCustom to be an explicit nil

### UnsetMatchCustom
`func (o *PatchedWritableRoutingPolicyRuleRequest) UnsetMatchCustom()`

UnsetMatchCustom ensures that no value is present for MatchCustom, not even an explicit nil
### GetSetActions

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetSetActions() interface{}`

GetSetActions returns the SetActions field if non-nil, zero value otherwise.

### GetSetActionsOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetSetActionsOk() (*interface{}, bool)`

GetSetActionsOk returns a tuple with the SetActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetActions

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetSetActions(v interface{})`

SetSetActions sets SetActions field to given value.

### HasSetActions

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasSetActions() bool`

HasSetActions returns a boolean if a field has been set.

### SetSetActionsNil

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetSetActionsNil(b bool)`

 SetSetActionsNil sets the value for SetActions to be an explicit nil

### UnsetSetActions
`func (o *PatchedWritableRoutingPolicyRuleRequest) UnsetSetActions()`

UnsetSetActions ensures that no value is present for SetActions, not even an explicit nil
### GetMatchIpv6Address

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetMatchIpv6Address() []int32`

GetMatchIpv6Address returns the MatchIpv6Address field if non-nil, zero value otherwise.

### GetMatchIpv6AddressOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetMatchIpv6AddressOk() (*[]int32, bool)`

GetMatchIpv6AddressOk returns a tuple with the MatchIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpv6Address

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetMatchIpv6Address(v []int32)`

SetMatchIpv6Address sets MatchIpv6Address field to given value.

### HasMatchIpv6Address

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasMatchIpv6Address() bool`

HasMatchIpv6Address returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableRoutingPolicyRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableRoutingPolicyRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableRoutingPolicyRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


