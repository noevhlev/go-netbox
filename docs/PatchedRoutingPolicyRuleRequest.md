# PatchedRoutingPolicyRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Action** | Pointer to [**CommunityListRuleAction**](CommunityListRuleAction.md) |  | [optional] 
**MatchIpAddress** | Pointer to **[]int32** |  | [optional] 
**RoutingPolicy** | Pointer to [**BriefRoutingPolicyRequest**](BriefRoutingPolicyRequest.md) |  | [optional] 
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

### NewPatchedRoutingPolicyRuleRequest

`func NewPatchedRoutingPolicyRuleRequest() *PatchedRoutingPolicyRuleRequest`

NewPatchedRoutingPolicyRuleRequest instantiates a new PatchedRoutingPolicyRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRoutingPolicyRuleRequestWithDefaults

`func NewPatchedRoutingPolicyRuleRequestWithDefaults() *PatchedRoutingPolicyRuleRequest`

NewPatchedRoutingPolicyRuleRequestWithDefaults instantiates a new PatchedRoutingPolicyRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *PatchedRoutingPolicyRuleRequest) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PatchedRoutingPolicyRuleRequest) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PatchedRoutingPolicyRuleRequest) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PatchedRoutingPolicyRuleRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetAction

`func (o *PatchedRoutingPolicyRuleRequest) GetAction() CommunityListRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedRoutingPolicyRuleRequest) GetActionOk() (*CommunityListRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedRoutingPolicyRuleRequest) SetAction(v CommunityListRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedRoutingPolicyRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMatchIpAddress

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchIpAddress() []*int32`

GetMatchIpAddress returns the MatchIpAddress field if non-nil, zero value otherwise.

### GetMatchIpAddressOk

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchIpAddressOk() (*[]*int32, bool)`

GetMatchIpAddressOk returns a tuple with the MatchIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpAddress

`func (o *PatchedRoutingPolicyRuleRequest) SetMatchIpAddress(v []*int32)`

SetMatchIpAddress sets MatchIpAddress field to given value.

### HasMatchIpAddress

`func (o *PatchedRoutingPolicyRuleRequest) HasMatchIpAddress() bool`

HasMatchIpAddress returns a boolean if a field has been set.

### GetRoutingPolicy

`func (o *PatchedRoutingPolicyRuleRequest) GetRoutingPolicy() BriefRoutingPolicyRequest`

GetRoutingPolicy returns the RoutingPolicy field if non-nil, zero value otherwise.

### GetRoutingPolicyOk

`func (o *PatchedRoutingPolicyRuleRequest) GetRoutingPolicyOk() (*BriefRoutingPolicyRequest, bool)`

GetRoutingPolicyOk returns a tuple with the RoutingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicy

`func (o *PatchedRoutingPolicyRuleRequest) SetRoutingPolicy(v BriefRoutingPolicyRequest)`

SetRoutingPolicy sets RoutingPolicy field to given value.

### HasRoutingPolicy

`func (o *PatchedRoutingPolicyRuleRequest) HasRoutingPolicy() bool`

HasRoutingPolicy returns a boolean if a field has been set.

### GetMatchCommunity

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchCommunity() []*int32`

GetMatchCommunity returns the MatchCommunity field if non-nil, zero value otherwise.

### GetMatchCommunityOk

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchCommunityOk() (*[]*int32, bool)`

GetMatchCommunityOk returns a tuple with the MatchCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCommunity

`func (o *PatchedRoutingPolicyRuleRequest) SetMatchCommunity(v []*int32)`

SetMatchCommunity sets MatchCommunity field to given value.

### HasMatchCommunity

`func (o *PatchedRoutingPolicyRuleRequest) HasMatchCommunity() bool`

HasMatchCommunity returns a boolean if a field has been set.

### GetMatchCommunityList

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchCommunityList() []*int32`

GetMatchCommunityList returns the MatchCommunityList field if non-nil, zero value otherwise.

### GetMatchCommunityListOk

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchCommunityListOk() (*[]*int32, bool)`

GetMatchCommunityListOk returns a tuple with the MatchCommunityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCommunityList

`func (o *PatchedRoutingPolicyRuleRequest) SetMatchCommunityList(v []*int32)`

SetMatchCommunityList sets MatchCommunityList field to given value.

### HasMatchCommunityList

`func (o *PatchedRoutingPolicyRuleRequest) HasMatchCommunityList() bool`

HasMatchCommunityList returns a boolean if a field has been set.

### GetMatchCustom

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchCustom() interface{}`

GetMatchCustom returns the MatchCustom field if non-nil, zero value otherwise.

### GetMatchCustomOk

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchCustomOk() (*interface{}, bool)`

GetMatchCustomOk returns a tuple with the MatchCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCustom

`func (o *PatchedRoutingPolicyRuleRequest) SetMatchCustom(v interface{})`

SetMatchCustom sets MatchCustom field to given value.

### HasMatchCustom

`func (o *PatchedRoutingPolicyRuleRequest) HasMatchCustom() bool`

HasMatchCustom returns a boolean if a field has been set.

### SetMatchCustomNil

`func (o *PatchedRoutingPolicyRuleRequest) SetMatchCustomNil(b bool)`

 SetMatchCustomNil sets the value for MatchCustom to be an explicit nil

### UnsetMatchCustom
`func (o *PatchedRoutingPolicyRuleRequest) UnsetMatchCustom()`

UnsetMatchCustom ensures that no value is present for MatchCustom, not even an explicit nil
### GetSetActions

`func (o *PatchedRoutingPolicyRuleRequest) GetSetActions() interface{}`

GetSetActions returns the SetActions field if non-nil, zero value otherwise.

### GetSetActionsOk

`func (o *PatchedRoutingPolicyRuleRequest) GetSetActionsOk() (*interface{}, bool)`

GetSetActionsOk returns a tuple with the SetActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetActions

`func (o *PatchedRoutingPolicyRuleRequest) SetSetActions(v interface{})`

SetSetActions sets SetActions field to given value.

### HasSetActions

`func (o *PatchedRoutingPolicyRuleRequest) HasSetActions() bool`

HasSetActions returns a boolean if a field has been set.

### SetSetActionsNil

`func (o *PatchedRoutingPolicyRuleRequest) SetSetActionsNil(b bool)`

 SetSetActionsNil sets the value for SetActions to be an explicit nil

### UnsetSetActions
`func (o *PatchedRoutingPolicyRuleRequest) UnsetSetActions()`

UnsetSetActions ensures that no value is present for SetActions, not even an explicit nil
### GetMatchIpv6Address

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchIpv6Address() []int32`

GetMatchIpv6Address returns the MatchIpv6Address field if non-nil, zero value otherwise.

### GetMatchIpv6AddressOk

`func (o *PatchedRoutingPolicyRuleRequest) GetMatchIpv6AddressOk() (*[]int32, bool)`

GetMatchIpv6AddressOk returns a tuple with the MatchIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpv6Address

`func (o *PatchedRoutingPolicyRuleRequest) SetMatchIpv6Address(v []int32)`

SetMatchIpv6Address sets MatchIpv6Address field to given value.

### HasMatchIpv6Address

`func (o *PatchedRoutingPolicyRuleRequest) HasMatchIpv6Address() bool`

HasMatchIpv6Address returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedRoutingPolicyRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedRoutingPolicyRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedRoutingPolicyRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedRoutingPolicyRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedRoutingPolicyRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedRoutingPolicyRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedRoutingPolicyRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedRoutingPolicyRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedRoutingPolicyRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedRoutingPolicyRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedRoutingPolicyRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedRoutingPolicyRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComments

`func (o *PatchedRoutingPolicyRuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedRoutingPolicyRuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedRoutingPolicyRuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedRoutingPolicyRuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


