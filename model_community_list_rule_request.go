/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the CommunityListRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityListRuleRequest{}

// CommunityListRuleRequest Adds support for custom fields and tags.
type CommunityListRuleRequest struct {
	Tags                 []NestedTagRequest            `json:"tags,omitempty"`
	CustomFields         map[string]interface{}        `json:"custom_fields,omitempty"`
	Description          *string                       `json:"description,omitempty"`
	CommunityList        BriefCommunityListRequest     `json:"community_list"`
	Action               CommunityListRuleAction       `json:"action"`
	Community            NullableBriefCommunityRequest `json:"community,omitempty"`
	Comments             *string                       `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommunityListRuleRequest CommunityListRuleRequest

// NewCommunityListRuleRequest instantiates a new CommunityListRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityListRuleRequest(communityList BriefCommunityListRequest, action CommunityListRuleAction) *CommunityListRuleRequest {
	this := CommunityListRuleRequest{}
	this.CommunityList = communityList
	this.Action = action
	return &this
}

// NewCommunityListRuleRequestWithDefaults instantiates a new CommunityListRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityListRuleRequestWithDefaults() *CommunityListRuleRequest {
	this := CommunityListRuleRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CommunityListRuleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityListRuleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CommunityListRuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *CommunityListRuleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CommunityListRuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityListRuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CommunityListRuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CommunityListRuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommunityListRuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityListRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommunityListRuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommunityListRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCommunityList returns the CommunityList field value
func (o *CommunityListRuleRequest) GetCommunityList() BriefCommunityListRequest {
	if o == nil {
		var ret BriefCommunityListRequest
		return ret
	}

	return o.CommunityList
}

// GetCommunityListOk returns a tuple with the CommunityList field value
// and a boolean to check if the value has been set.
func (o *CommunityListRuleRequest) GetCommunityListOk() (*BriefCommunityListRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunityList, true
}

// SetCommunityList sets field value
func (o *CommunityListRuleRequest) SetCommunityList(v BriefCommunityListRequest) {
	o.CommunityList = v
}

// GetAction returns the Action field value
func (o *CommunityListRuleRequest) GetAction() CommunityListRuleAction {
	if o == nil {
		var ret CommunityListRuleAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CommunityListRuleRequest) GetActionOk() (*CommunityListRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *CommunityListRuleRequest) SetAction(v CommunityListRuleAction) {
	o.Action = v
}

// GetCommunity returns the Community field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommunityListRuleRequest) GetCommunity() BriefCommunityRequest {
	if o == nil || IsNil(o.Community.Get()) {
		var ret BriefCommunityRequest
		return ret
	}
	return *o.Community.Get()
}

// GetCommunityOk returns a tuple with the Community field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityListRuleRequest) GetCommunityOk() (*BriefCommunityRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Community.Get(), o.Community.IsSet()
}

// HasCommunity returns a boolean if a field has been set.
func (o *CommunityListRuleRequest) HasCommunity() bool {
	if o != nil && o.Community.IsSet() {
		return true
	}

	return false
}

// SetCommunity gets a reference to the given NullableBriefCommunityRequest and assigns it to the Community field.
func (o *CommunityListRuleRequest) SetCommunity(v BriefCommunityRequest) {
	o.Community.Set(&v)
}

// SetCommunityNil sets the value for Community to be an explicit nil
func (o *CommunityListRuleRequest) SetCommunityNil() {
	o.Community.Set(nil)
}

// UnsetCommunity ensures that no value is present for Community, not even an explicit nil
func (o *CommunityListRuleRequest) UnsetCommunity() {
	o.Community.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *CommunityListRuleRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityListRuleRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *CommunityListRuleRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *CommunityListRuleRequest) SetComments(v string) {
	o.Comments = &v
}

func (o CommunityListRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityListRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["community_list"] = o.CommunityList
	toSerialize["action"] = o.Action
	if o.Community.IsSet() {
		toSerialize["community"] = o.Community.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommunityListRuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"community_list",
		"action",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommunityListRuleRequest := _CommunityListRuleRequest{}

	err = json.Unmarshal(data, &varCommunityListRuleRequest)

	if err != nil {
		return err
	}

	*o = CommunityListRuleRequest(varCommunityListRuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "description")
		delete(additionalProperties, "community_list")
		delete(additionalProperties, "action")
		delete(additionalProperties, "community")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommunityListRuleRequest struct {
	value *CommunityListRuleRequest
	isSet bool
}

func (v NullableCommunityListRuleRequest) Get() *CommunityListRuleRequest {
	return v.value
}

func (v *NullableCommunityListRuleRequest) Set(val *CommunityListRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityListRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityListRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityListRuleRequest(val *CommunityListRuleRequest) *NullableCommunityListRuleRequest {
	return &NullableCommunityListRuleRequest{value: val, isSet: true}
}

func (v NullableCommunityListRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityListRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
