# RoutingPolicyRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**MatchIpAddress** | Pointer to **[]int32** |  | [optional] 
**RoutingPolicy** | [**BriefRoutingPolicyRequest**](BriefRoutingPolicyRequest.md) |  | 
**MatchCommunity** | Pointer to **[]int32** |  | [optional] 
**MatchCommunityList** | Pointer to **[]int32** |  | [optional] 
**MatchCustom** | Pointer to **interface{}** |  | [optional] 
**SetActions** | Pointer to **interface{}** |  | [optional] 
**MatchIpv6Address** | Pointer to **[]int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewRoutingPolicyRuleRequest

`func NewRoutingPolicyRuleRequest(index int32, action CommunityListRuleAction, routingPolicy BriefRoutingPolicyRequest, ) *RoutingPolicyRuleRequest`

NewRoutingPolicyRuleRequest instantiates a new RoutingPolicyRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyRuleRequestWithDefaults

`func NewRoutingPolicyRuleRequestWithDefaults() *RoutingPolicyRuleRequest`

NewRoutingPolicyRuleRequestWithDefaults instantiates a new RoutingPolicyRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *RoutingPolicyRuleRequest) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RoutingPolicyRuleRequest) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RoutingPolicyRuleRequest) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetAction

`func (o *RoutingPolicyRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RoutingPolicyRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RoutingPolicyRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetMatchIpAddress

`func (o *RoutingPolicyRuleRequest) GetMatchIpAddress() []*int32`

GetMatchIpAddress returns the MatchIpAddress field if non-nil, zero value otherwise.

### GetMatchIpAddressOk

`func (o *RoutingPolicyRuleRequest) GetMatchIpAddressOk() (*[]*int32, bool)`

GetMatchIpAddressOk returns a tuple with the MatchIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpAddress

`func (o *RoutingPolicyRuleRequest) SetMatchIpAddress(v []*int32)`

SetMatchIpAddress sets MatchIpAddress field to given value.

### HasMatchIpAddress

`func (o *RoutingPolicyRuleRequest) HasMatchIpAddress() bool`

HasMatchIpAddress returns a boolean if a field has been set.

### GetRoutingPolicy

`func (o *RoutingPolicyRuleRequest) GetRoutingPolicy() BriefRoutingPolicyRequest`

GetRoutingPolicy returns the RoutingPolicy field if non-nil, zero value otherwise.

### GetRoutingPolicyOk

`func (o *RoutingPolicyRuleRequest) GetRoutingPolicyOk() (*BriefRoutingPolicyRequest, bool)`

GetRoutingPolicyOk returns a tuple with the RoutingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicy

`func (o *RoutingPolicyRuleRequest) SetRoutingPolicy(v BriefRoutingPolicyRequest)`

SetRoutingPolicy sets RoutingPolicy field to given value.


### GetMatchCommunity

`func (o *RoutingPolicyRuleRequest) GetMatchCommunity() []*int32`

GetMatchCommunity returns the MatchCommunity field if non-nil, zero value otherwise.

### GetMatchCommunityOk

`func (o *RoutingPolicyRuleRequest) GetMatchCommunityOk() (*[]*int32, bool)`

GetMatchCommunityOk returns a tuple with the MatchCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCommunity

`func (o *RoutingPolicyRuleRequest) SetMatchCommunity(v []*int32)`

SetMatchCommunity sets MatchCommunity field to given value.

### HasMatchCommunity

`func (o *RoutingPolicyRuleRequest) HasMatchCommunity() bool`

HasMatchCommunity returns a boolean if a field has been set.

### GetMatchCommunityList

`func (o *RoutingPolicyRuleRequest) GetMatchCommunityList() []*int32`

GetMatchCommunityList returns the MatchCommunityList field if non-nil, zero value otherwise.

### GetMatchCommunityListOk

`func (o *RoutingPolicyRuleRequest) GetMatchCommunityListOk() (*[]*int32, bool)`

GetMatchCommunityListOk returns a tuple with the MatchCommunityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCommunityList

`func (o *RoutingPolicyRuleRequest) SetMatchCommunityList(v []*int32)`

SetMatchCommunityList sets MatchCommunityList field to given value.

### HasMatchCommunityList

`func (o *RoutingPolicyRuleRequest) HasMatchCommunityList() bool`

HasMatchCommunityList returns a boolean if a field has been set.

### GetMatchCustom

`func (o *RoutingPolicyRuleRequest) GetMatchCustom() interface{}`

GetMatchCustom returns the MatchCustom field if non-nil, zero value otherwise.

### GetMatchCustomOk

`func (o *RoutingPolicyRuleRequest) GetMatchCustomOk() (*interface{}, bool)`

GetMatchCustomOk returns a tuple with the MatchCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCustom

`func (o *RoutingPolicyRuleRequest) SetMatchCustom(v interface{})`

SetMatchCustom sets MatchCustom field to given value.

### HasMatchCustom

`func (o *RoutingPolicyRuleRequest) HasMatchCustom() bool`

HasMatchCustom returns a boolean if a field has been set.

### SetMatchCustomNil

`func (o *RoutingPolicyRuleRequest) SetMatchCustomNil(b bool)`

 SetMatchCustomNil sets the value for MatchCustom to be an explicit nil

### UnsetMatchCustom
`func (o *RoutingPolicyRuleRequest) UnsetMatchCustom()`

UnsetMatchCustom ensures that no value is present for MatchCustom, not even an explicit nil
### GetSetActions

`func (o *RoutingPolicyRuleRequest) GetSetActions() interface{}`

GetSetActions returns the SetActions field if non-nil, zero value otherwise.

### GetSetActionsOk

`func (o *RoutingPolicyRuleRequest) GetSetActionsOk() (*interface{}, bool)`

GetSetActionsOk returns a tuple with the SetActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetActions

`func (o *RoutingPolicyRuleRequest) SetSetActions(v interface{})`

SetSetActions sets SetActions field to given value.

### HasSetActions

`func (o *RoutingPolicyRuleRequest) HasSetActions() bool`

HasSetActions returns a boolean if a field has been set.

### SetSetActionsNil

`func (o *RoutingPolicyRuleRequest) SetSetActionsNil(b bool)`

 SetSetActionsNil sets the value for SetActions to be an explicit nil

### UnsetSetActions
`func (o *RoutingPolicyRuleRequest) UnsetSetActions()`

UnsetSetActions ensures that no value is present for SetActions, not even an explicit nil
### GetMatchIpv6Address

`func (o *RoutingPolicyRuleRequest) GetMatchIpv6Address() []int32`

GetMatchIpv6Address returns the MatchIpv6Address field if non-nil, zero value otherwise.

### GetMatchIpv6AddressOk

`func (o *RoutingPolicyRuleRequest) GetMatchIpv6AddressOk() (*[]int32, bool)`

GetMatchIpv6AddressOk returns a tuple with the MatchIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpv6Address

`func (o *RoutingPolicyRuleRequest) SetMatchIpv6Address(v []int32)`

SetMatchIpv6Address sets MatchIpv6Address field to given value.

### HasMatchIpv6Address

`func (o *RoutingPolicyRuleRequest) HasMatchIpv6Address() bool`

HasMatchIpv6Address returns a boolean if a field has been set.

### GetDescription

`func (o *RoutingPolicyRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingPolicyRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingPolicyRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingPolicyRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *RoutingPolicyRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingPolicyRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingPolicyRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingPolicyRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *RoutingPolicyRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RoutingPolicyRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RoutingPolicyRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RoutingPolicyRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComments

`func (o *RoutingPolicyRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RoutingPolicyRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RoutingPolicyRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RoutingPolicyRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


