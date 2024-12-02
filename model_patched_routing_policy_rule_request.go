/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedRoutingPolicyRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedRoutingPolicyRuleRequest{}

// PatchedRoutingPolicyRuleRequest Adds support for custom fields and tags.
type PatchedRoutingPolicyRuleRequest struct {
	Index                *int32                     `json:"index,omitempty"`
	Action               *CommunityListRuleAction   `json:"action,omitempty"`
	MatchIpAddress       []*int32                   `json:"match_ip_address,omitempty"`
	RoutingPolicy        *BriefRoutingPolicyRequest `json:"routing_policy,omitempty"`
	MatchCommunity       []*int32                   `json:"match_community,omitempty"`
	MatchCommunityList   []*int32                   `json:"match_community_list,omitempty"`
	MatchCustom          interface{}                `json:"match_custom,omitempty"`
	SetActions           interface{}                `json:"set_actions,omitempty"`
	MatchIpv6Address     []int32                    `json:"match_ipv6_address,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	Tags                 []NestedTagRequest         `json:"tags,omitempty"`
	CustomFields         map[string]interface{}     `json:"custom_fields,omitempty"`
	Comments             *string                    `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedRoutingPolicyRuleRequest PatchedRoutingPolicyRuleRequest

// NewPatchedRoutingPolicyRuleRequest instantiates a new PatchedRoutingPolicyRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRoutingPolicyRuleRequest() *PatchedRoutingPolicyRuleRequest {
	this := PatchedRoutingPolicyRuleRequest{}
	return &this
}

// NewPatchedRoutingPolicyRuleRequestWithDefaults instantiates a new PatchedRoutingPolicyRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRoutingPolicyRuleRequestWithDefaults() *PatchedRoutingPolicyRuleRequest {
	this := PatchedRoutingPolicyRuleRequest{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *PatchedRoutingPolicyRuleRequest) SetIndex(v int32) {
	o.Index = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetAction() CommunityListRuleAction {
	if o == nil || IsNil(o.Action) {
		var ret CommunityListRuleAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetActionOk() (*CommunityListRuleAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given CommunityListRuleAction and assigns it to the Action field.
func (o *PatchedRoutingPolicyRuleRequest) SetAction(v CommunityListRuleAction) {
	o.Action = &v
}

// GetMatchIpAddress returns the MatchIpAddress field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetMatchIpAddress() []*int32 {
	if o == nil || IsNil(o.MatchIpAddress) {
		var ret []*int32
		return ret
	}
	return o.MatchIpAddress
}

// GetMatchIpAddressOk returns a tuple with the MatchIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetMatchIpAddressOk() ([]*int32, bool) {
	if o == nil || IsNil(o.MatchIpAddress) {
		return nil, false
	}
	return o.MatchIpAddress, true
}

// HasMatchIpAddress returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasMatchIpAddress() bool {
	if o != nil && !IsNil(o.MatchIpAddress) {
		return true
	}

	return false
}

// SetMatchIpAddress gets a reference to the given []*int32 and assigns it to the MatchIpAddress field.
func (o *PatchedRoutingPolicyRuleRequest) SetMatchIpAddress(v []*int32) {
	o.MatchIpAddress = v
}

// GetRoutingPolicy returns the RoutingPolicy field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetRoutingPolicy() BriefRoutingPolicyRequest {
	if o == nil || IsNil(o.RoutingPolicy) {
		var ret BriefRoutingPolicyRequest
		return ret
	}
	return *o.RoutingPolicy
}

// GetRoutingPolicyOk returns a tuple with the RoutingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetRoutingPolicyOk() (*BriefRoutingPolicyRequest, bool) {
	if o == nil || IsNil(o.RoutingPolicy) {
		return nil, false
	}
	return o.RoutingPolicy, true
}

// HasRoutingPolicy returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasRoutingPolicy() bool {
	if o != nil && !IsNil(o.RoutingPolicy) {
		return true
	}

	return false
}

// SetRoutingPolicy gets a reference to the given BriefRoutingPolicyRequest and assigns it to the RoutingPolicy field.
func (o *PatchedRoutingPolicyRuleRequest) SetRoutingPolicy(v BriefRoutingPolicyRequest) {
	o.RoutingPolicy = &v
}

// GetMatchCommunity returns the MatchCommunity field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetMatchCommunity() []*int32 {
	if o == nil || IsNil(o.MatchCommunity) {
		var ret []*int32
		return ret
	}
	return o.MatchCommunity
}

// GetMatchCommunityOk returns a tuple with the MatchCommunity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetMatchCommunityOk() ([]*int32, bool) {
	if o == nil || IsNil(o.MatchCommunity) {
		return nil, false
	}
	return o.MatchCommunity, true
}

// HasMatchCommunity returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasMatchCommunity() bool {
	if o != nil && !IsNil(o.MatchCommunity) {
		return true
	}

	return false
}

// SetMatchCommunity gets a reference to the given []*int32 and assigns it to the MatchCommunity field.
func (o *PatchedRoutingPolicyRuleRequest) SetMatchCommunity(v []*int32) {
	o.MatchCommunity = v
}

// GetMatchCommunityList returns the MatchCommunityList field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetMatchCommunityList() []*int32 {
	if o == nil || IsNil(o.MatchCommunityList) {
		var ret []*int32
		return ret
	}
	return o.MatchCommunityList
}

// GetMatchCommunityListOk returns a tuple with the MatchCommunityList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetMatchCommunityListOk() ([]*int32, bool) {
	if o == nil || IsNil(o.MatchCommunityList) {
		return nil, false
	}
	return o.MatchCommunityList, true
}

// HasMatchCommunityList returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasMatchCommunityList() bool {
	if o != nil && !IsNil(o.MatchCommunityList) {
		return true
	}

	return false
}

// SetMatchCommunityList gets a reference to the given []*int32 and assigns it to the MatchCommunityList field.
func (o *PatchedRoutingPolicyRuleRequest) SetMatchCommunityList(v []*int32) {
	o.MatchCommunityList = v
}

// GetMatchCustom returns the MatchCustom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRoutingPolicyRuleRequest) GetMatchCustom() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MatchCustom
}

// GetMatchCustomOk returns a tuple with the MatchCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRoutingPolicyRuleRequest) GetMatchCustomOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MatchCustom) {
		return nil, false
	}
	return &o.MatchCustom, true
}

// HasMatchCustom returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasMatchCustom() bool {
	if o != nil && !IsNil(o.MatchCustom) {
		return true
	}

	return false
}

// SetMatchCustom gets a reference to the given interface{} and assigns it to the MatchCustom field.
func (o *PatchedRoutingPolicyRuleRequest) SetMatchCustom(v interface{}) {
	o.MatchCustom = v
}

// GetSetActions returns the SetActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRoutingPolicyRuleRequest) GetSetActions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SetActions
}

// GetSetActionsOk returns a tuple with the SetActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRoutingPolicyRuleRequest) GetSetActionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SetActions) {
		return nil, false
	}
	return &o.SetActions, true
}

// HasSetActions returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasSetActions() bool {
	if o != nil && !IsNil(o.SetActions) {
		return true
	}

	return false
}

// SetSetActions gets a reference to the given interface{} and assigns it to the SetActions field.
func (o *PatchedRoutingPolicyRuleRequest) SetSetActions(v interface{}) {
	o.SetActions = v
}

// GetMatchIpv6Address returns the MatchIpv6Address field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetMatchIpv6Address() []int32 {
	if o == nil || IsNil(o.MatchIpv6Address) {
		var ret []int32
		return ret
	}
	return o.MatchIpv6Address
}

// GetMatchIpv6AddressOk returns a tuple with the MatchIpv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetMatchIpv6AddressOk() ([]int32, bool) {
	if o == nil || IsNil(o.MatchIpv6Address) {
		return nil, false
	}
	return o.MatchIpv6Address, true
}

// HasMatchIpv6Address returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasMatchIpv6Address() bool {
	if o != nil && !IsNil(o.MatchIpv6Address) {
		return true
	}

	return false
}

// SetMatchIpv6Address gets a reference to the given []int32 and assigns it to the MatchIpv6Address field.
func (o *PatchedRoutingPolicyRuleRequest) SetMatchIpv6Address(v []int32) {
	o.MatchIpv6Address = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedRoutingPolicyRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedRoutingPolicyRuleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedRoutingPolicyRuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedRoutingPolicyRuleRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoutingPolicyRuleRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedRoutingPolicyRuleRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedRoutingPolicyRuleRequest) SetComments(v string) {
	o.Comments = &v
}

func (o PatchedRoutingPolicyRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedRoutingPolicyRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.MatchIpAddress) {
		toSerialize["match_ip_address"] = o.MatchIpAddress
	}
	if !IsNil(o.RoutingPolicy) {
		toSerialize["routing_policy"] = o.RoutingPolicy
	}
	if !IsNil(o.MatchCommunity) {
		toSerialize["match_community"] = o.MatchCommunity
	}
	if !IsNil(o.MatchCommunityList) {
		toSerialize["match_community_list"] = o.MatchCommunityList
	}
	if o.MatchCustom != nil {
		toSerialize["match_custom"] = o.MatchCustom
	}
	if o.SetActions != nil {
		toSerialize["set_actions"] = o.SetActions
	}
	if !IsNil(o.MatchIpv6Address) {
		toSerialize["match_ipv6_address"] = o.MatchIpv6Address
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedRoutingPolicyRuleRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedRoutingPolicyRuleRequest := _PatchedRoutingPolicyRuleRequest{}

	err = json.Unmarshal(data, &varPatchedRoutingPolicyRuleRequest)

	if err != nil {
		return err
	}

	*o = PatchedRoutingPolicyRuleRequest(varPatchedRoutingPolicyRuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "index")
		delete(additionalProperties, "action")
		delete(additionalProperties, "match_ip_address")
		delete(additionalProperties, "routing_policy")
		delete(additionalProperties, "match_community")
		delete(additionalProperties, "match_community_list")
		delete(additionalProperties, "match_custom")
		delete(additionalProperties, "set_actions")
		delete(additionalProperties, "match_ipv6_address")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedRoutingPolicyRuleRequest struct {
	value *PatchedRoutingPolicyRuleRequest
	isSet bool
}

func (v NullablePatchedRoutingPolicyRuleRequest) Get() *PatchedRoutingPolicyRuleRequest {
	return v.value
}

func (v *NullablePatchedRoutingPolicyRuleRequest) Set(val *PatchedRoutingPolicyRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRoutingPolicyRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRoutingPolicyRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRoutingPolicyRuleRequest(val *PatchedRoutingPolicyRuleRequest) *NullablePatchedRoutingPolicyRuleRequest {
	return &NullablePatchedRoutingPolicyRuleRequest{value: val, isSet: true}
}

func (v NullablePatchedRoutingPolicyRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRoutingPolicyRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
