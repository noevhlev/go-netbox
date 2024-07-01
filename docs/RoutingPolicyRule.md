# RoutingPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Index** | **int32** |  | 
**Display** | **string** |  | [readonly] 
**Action** | [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | 
**MatchIpAddress** | Pointer to **[]int32** |  | [optional] 
**RoutingPolicy** | [**NestedRoutingPolicy**](NestedRoutingPolicy.md) |  | 
**MatchCommunity** | Pointer to **[]int32** |  | [optional] 
**MatchCustom** | Pointer to **interface{}** |  | [optional] 
**SetActions** | Pointer to **interface{}** |  | [optional] 
**MatchIpv6Address** | Pointer to **[]int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewRoutingPolicyRule

`func NewRoutingPolicyRule(id int32, index int32, display string, action CommunityListRuleAction, routingPolicy NestedRoutingPolicy, ) *RoutingPolicyRule`

NewRoutingPolicyRule instantiates a new RoutingPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyRuleWithDefaults

`func NewRoutingPolicyRuleWithDefaults() *RoutingPolicyRule`

NewRoutingPolicyRuleWithDefaults instantiates a new RoutingPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoutingPolicyRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoutingPolicyRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoutingPolicyRule) SetId(v int32)`

SetId sets Id field to given value.


### GetIndex

`func (o *RoutingPolicyRule) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RoutingPolicyRule) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RoutingPolicyRule) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetDisplay

`func (o *RoutingPolicyRule) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RoutingPolicyRule) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RoutingPolicyRule) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetAction

`func (o *RoutingPolicyRule) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RoutingPolicyRule) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RoutingPolicyRule) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.


### GetMatchIpAddress

`func (o *RoutingPolicyRule) GetMatchIpAddress() []*int32`

GetMatchIpAddress returns the MatchIpAddress field if non-nil, zero value otherwise.

### GetMatchIpAddressOk

`func (o *RoutingPolicyRule) GetMatchIpAddressOk() (*[]*int32, bool)`

GetMatchIpAddressOk returns a tuple with the MatchIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpAddress

`func (o *RoutingPolicyRule) SetMatchIpAddress(v []*int32)`

SetMatchIpAddress sets MatchIpAddress field to given value.

### HasMatchIpAddress

`func (o *RoutingPolicyRule) HasMatchIpAddress() bool`

HasMatchIpAddress returns a boolean if a field has been set.

### GetRoutingPolicy

`func (o *RoutingPolicyRule) GetRoutingPolicy() NestedRoutingPolicy`

GetRoutingPolicy returns the RoutingPolicy field if non-nil, zero value otherwise.

### GetRoutingPolicyOk

`func (o *RoutingPolicyRule) GetRoutingPolicyOk() (*NestedRoutingPolicy, bool)`

GetRoutingPolicyOk returns a tuple with the RoutingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicy

`func (o *RoutingPolicyRule) SetRoutingPolicy(v NestedRoutingPolicy)`

SetRoutingPolicy sets RoutingPolicy field to given value.


### GetMatchCommunity

`func (o *RoutingPolicyRule) GetMatchCommunity() []*int32`

GetMatchCommunity returns the MatchCommunity field if non-nil, zero value otherwise.

### GetMatchCommunityOk

`func (o *RoutingPolicyRule) GetMatchCommunityOk() (*[]*int32, bool)`

GetMatchCommunityOk returns a tuple with the MatchCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCommunity

`func (o *RoutingPolicyRule) SetMatchCommunity(v []*int32)`

SetMatchCommunity sets MatchCommunity field to given value.

### HasMatchCommunity

`func (o *RoutingPolicyRule) HasMatchCommunity() bool`

HasMatchCommunity returns a boolean if a field has been set.

### GetMatchCustom

`func (o *RoutingPolicyRule) GetMatchCustom() interface{}`

GetMatchCustom returns the MatchCustom field if non-nil, zero value otherwise.

### GetMatchCustomOk

`func (o *RoutingPolicyRule) GetMatchCustomOk() (*interface{}, bool)`

GetMatchCustomOk returns a tuple with the MatchCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCustom

`func (o *RoutingPolicyRule) SetMatchCustom(v interface{})`

SetMatchCustom sets MatchCustom field to given value.

### HasMatchCustom

`func (o *RoutingPolicyRule) HasMatchCustom() bool`

HasMatchCustom returns a boolean if a field has been set.

### SetMatchCustomNil

`func (o *RoutingPolicyRule) SetMatchCustomNil(b bool)`

 SetMatchCustomNil sets the value for MatchCustom to be an explicit nil

### UnsetMatchCustom
`func (o *RoutingPolicyRule) UnsetMatchCustom()`

UnsetMatchCustom ensures that no value is present for MatchCustom, not even an explicit nil
### GetSetActions

`func (o *RoutingPolicyRule) GetSetActions() interface{}`

GetSetActions returns the SetActions field if non-nil, zero value otherwise.

### GetSetActionsOk

`func (o *RoutingPolicyRule) GetSetActionsOk() (*interface{}, bool)`

GetSetActionsOk returns a tuple with the SetActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetActions

`func (o *RoutingPolicyRule) SetSetActions(v interface{})`

SetSetActions sets SetActions field to given value.

### HasSetActions

`func (o *RoutingPolicyRule) HasSetActions() bool`

HasSetActions returns a boolean if a field has been set.

### SetSetActionsNil

`func (o *RoutingPolicyRule) SetSetActionsNil(b bool)`

 SetSetActionsNil sets the value for SetActions to be an explicit nil

### UnsetSetActions
`func (o *RoutingPolicyRule) UnsetSetActions()`

UnsetSetActions ensures that no value is present for SetActions, not even an explicit nil
### GetMatchIpv6Address

`func (o *RoutingPolicyRule) GetMatchIpv6Address() []int32`

GetMatchIpv6Address returns the MatchIpv6Address field if non-nil, zero value otherwise.

### GetMatchIpv6AddressOk

`func (o *RoutingPolicyRule) GetMatchIpv6AddressOk() (*[]int32, bool)`

GetMatchIpv6AddressOk returns a tuple with the MatchIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpv6Address

`func (o *RoutingPolicyRule) SetMatchIpv6Address(v []int32)`

SetMatchIpv6Address sets MatchIpv6Address field to given value.

### HasMatchIpv6Address

`func (o *RoutingPolicyRule) HasMatchIpv6Address() bool`

HasMatchIpv6Address returns a boolean if a field has been set.

### GetDescription

`func (o *RoutingPolicyRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingPolicyRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingPolicyRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingPolicyRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *RoutingPolicyRule) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingPolicyRule) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingPolicyRule) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingPolicyRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *RoutingPolicyRule) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RoutingPolicyRule) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RoutingPolicyRule) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RoutingPolicyRule) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComments

`func (o *RoutingPolicyRule) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RoutingPolicyRule) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RoutingPolicyRule) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RoutingPolicyRule) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


