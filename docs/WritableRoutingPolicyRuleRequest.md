# WritableRoutingPolicyRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**MatchIpAddress** | Pointer to **[]int32** |  | [optional] 
**RoutingPolicy** | **int32** |  | 
**MatchCommunity** | Pointer to **[]int32** |  | [optional] 
**MatchCustom** | Pointer to **interface{}** |  | [optional] 
**SetActions** | Pointer to **interface{}** |  | [optional] 
**MatchIpv6Address** | Pointer to **[]int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewWritableRoutingPolicyRuleRequest

`func NewWritableRoutingPolicyRuleRequest(index int32, action CommunityListRuleAction, routingPolicy int32, ) *WritableRoutingPolicyRuleRequest`

NewWritableRoutingPolicyRuleRequest instantiates a new WritableRoutingPolicyRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableRoutingPolicyRuleRequestWithDefaults

`func NewWritableRoutingPolicyRuleRequestWithDefaults() *WritableRoutingPolicyRuleRequest`

NewWritableRoutingPolicyRuleRequestWithDefaults instantiates a new WritableRoutingPolicyRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *WritableRoutingPolicyRuleRequest) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WritableRoutingPolicyRuleRequest) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WritableRoutingPolicyRuleRequest) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetAction

`func (o *WritableRoutingPolicyRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WritableRoutingPolicyRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WritableRoutingPolicyRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetMatchIpAddress

`func (o *WritableRoutingPolicyRuleRequest) GetMatchIpAddress() []*int32`

GetMatchIpAddress returns the MatchIpAddress field if non-nil, zero value otherwise.

### GetMatchIpAddressOk

`func (o *WritableRoutingPolicyRuleRequest) GetMatchIpAddressOk() (*[]*int32, bool)`

GetMatchIpAddressOk returns a tuple with the MatchIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpAddress

`func (o *WritableRoutingPolicyRuleRequest) SetMatchIpAddress(v []*int32)`

SetMatchIpAddress sets MatchIpAddress field to given value.

### HasMatchIpAddress

`func (o *WritableRoutingPolicyRuleRequest) HasMatchIpAddress() bool`

HasMatchIpAddress returns a boolean if a field has been set.

### GetRoutingPolicy

`func (o *WritableRoutingPolicyRuleRequest) GetRoutingPolicy() int32`

GetRoutingPolicy returns the RoutingPolicy field if non-nil, zero value otherwise.

### GetRoutingPolicyOk

`func (o *WritableRoutingPolicyRuleRequest) GetRoutingPolicyOk() (*int32, bool)`

GetRoutingPolicyOk returns a tuple with the RoutingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicy

`func (o *WritableRoutingPolicyRuleRequest) SetRoutingPolicy(v int32)`

SetRoutingPolicy sets RoutingPolicy field to given value.


### GetMatchCommunity

`func (o *WritableRoutingPolicyRuleRequest) GetMatchCommunity() []*int32`

GetMatchCommunity returns the MatchCommunity field if non-nil, zero value otherwise.

### GetMatchCommunityOk

`func (o *WritableRoutingPolicyRuleRequest) GetMatchCommunityOk() (*[]*int32, bool)`

GetMatchCommunityOk returns a tuple with the MatchCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCommunity

`func (o *WritableRoutingPolicyRuleRequest) SetMatchCommunity(v []*int32)`

SetMatchCommunity sets MatchCommunity field to given value.

### HasMatchCommunity

`func (o *WritableRoutingPolicyRuleRequest) HasMatchCommunity() bool`

HasMatchCommunity returns a boolean if a field has been set.

### GetMatchCustom

`func (o *WritableRoutingPolicyRuleRequest) GetMatchCustom() interface{}`

GetMatchCustom returns the MatchCustom field if non-nil, zero value otherwise.

### GetMatchCustomOk

`func (o *WritableRoutingPolicyRuleRequest) GetMatchCustomOk() (*interface{}, bool)`

GetMatchCustomOk returns a tuple with the MatchCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCustom

`func (o *WritableRoutingPolicyRuleRequest) SetMatchCustom(v interface{})`

SetMatchCustom sets MatchCustom field to given value.

### HasMatchCustom

`func (o *WritableRoutingPolicyRuleRequest) HasMatchCustom() bool`

HasMatchCustom returns a boolean if a field has been set.

### SetMatchCustomNil

`func (o *WritableRoutingPolicyRuleRequest) SetMatchCustomNil(b bool)`

 SetMatchCustomNil sets the value for MatchCustom to be an explicit nil

### UnsetMatchCustom
`func (o *WritableRoutingPolicyRuleRequest) UnsetMatchCustom()`

UnsetMatchCustom ensures that no value is present for MatchCustom, not even an explicit nil
### GetSetActions

`func (o *WritableRoutingPolicyRuleRequest) GetSetActions() interface{}`

GetSetActions returns the SetActions field if non-nil, zero value otherwise.

### GetSetActionsOk

`func (o *WritableRoutingPolicyRuleRequest) GetSetActionsOk() (*interface{}, bool)`

GetSetActionsOk returns a tuple with the SetActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetActions

`func (o *WritableRoutingPolicyRuleRequest) SetSetActions(v interface{})`

SetSetActions sets SetActions field to given value.

### HasSetActions

`func (o *WritableRoutingPolicyRuleRequest) HasSetActions() bool`

HasSetActions returns a boolean if a field has been set.

### SetSetActionsNil

`func (o *WritableRoutingPolicyRuleRequest) SetSetActionsNil(b bool)`

 SetSetActionsNil sets the value for SetActions to be an explicit nil

### UnsetSetActions
`func (o *WritableRoutingPolicyRuleRequest) UnsetSetActions()`

UnsetSetActions ensures that no value is present for SetActions, not even an explicit nil
### GetMatchIpv6Address

`func (o *WritableRoutingPolicyRuleRequest) GetMatchIpv6Address() []int32`

GetMatchIpv6Address returns the MatchIpv6Address field if non-nil, zero value otherwise.

### GetMatchIpv6AddressOk

`func (o *WritableRoutingPolicyRuleRequest) GetMatchIpv6AddressOk() (*[]int32, bool)`

GetMatchIpv6AddressOk returns a tuple with the MatchIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpv6Address

`func (o *WritableRoutingPolicyRuleRequest) SetMatchIpv6Address(v []int32)`

SetMatchIpv6Address sets MatchIpv6Address field to given value.

### HasMatchIpv6Address

`func (o *WritableRoutingPolicyRuleRequest) HasMatchIpv6Address() bool`

HasMatchIpv6Address returns a boolean if a field has been set.

### GetDescription

`func (o *WritableRoutingPolicyRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableRoutingPolicyRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableRoutingPolicyRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableRoutingPolicyRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *WritableRoutingPolicyRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableRoutingPolicyRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableRoutingPolicyRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableRoutingPolicyRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableRoutingPolicyRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableRoutingPolicyRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableRoutingPolicyRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableRoutingPolicyRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComments

`func (o *WritableRoutingPolicyRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableRoutingPolicyRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableRoutingPolicyRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableRoutingPolicyRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


