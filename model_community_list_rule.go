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
	"time"
)

// checks if the CommunityListRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityListRule{}

// CommunityListRule Adds support for custom fields and tags.
type CommunityListRule struct {
	Id                   int32                   `json:"id"`
	Tags                 []NestedTag             `json:"tags,omitempty"`
	CustomFields         map[string]interface{}  `json:"custom_fields,omitempty"`
	Display              string                  `json:"display"`
	Description          *string                 `json:"description,omitempty"`
	CommunityList        BriefCommunityList      `json:"community_list"`
	Created              NullableTime            `json:"created"`
	LastUpdated          NullableTime            `json:"last_updated"`
	Action               CommunityListRuleAction `json:"action"`
	Community            NullableBriefCommunity  `json:"community,omitempty"`
	Comments             *string                 `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommunityListRule CommunityListRule

// NewCommunityListRule instantiates a new CommunityListRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityListRule(id int32, display string, communityList BriefCommunityList, created NullableTime, lastUpdated NullableTime, action CommunityListRuleAction) *CommunityListRule {
	this := CommunityListRule{}
	this.Id = id
	this.Display = display
	this.CommunityList = communityList
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Action = action
	return &this
}

// NewCommunityListRuleWithDefaults instantiates a new CommunityListRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityListRuleWithDefaults() *CommunityListRule {
	this := CommunityListRule{}
	return &this
}

// GetId returns the Id field value
func (o *CommunityListRule) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommunityListRule) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommunityListRule) SetId(v int32) {
	o.Id = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CommunityListRule) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityListRule) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CommunityListRule) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *CommunityListRule) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CommunityListRule) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityListRule) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CommunityListRule) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CommunityListRule) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *CommunityListRule) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CommunityListRule) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CommunityListRule) SetDisplay(v string) {
	o.Display = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommunityListRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityListRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommunityListRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommunityListRule) SetDescription(v string) {
	o.Description = &v
}

// GetCommunityList returns the CommunityList field value
func (o *CommunityListRule) GetCommunityList() BriefCommunityList {
	if o == nil {
		var ret BriefCommunityList
		return ret
	}

	return o.CommunityList
}

// GetCommunityListOk returns a tuple with the CommunityList field value
// and a boolean to check if the value has been set.
func (o *CommunityListRule) GetCommunityListOk() (*BriefCommunityList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunityList, true
}

// SetCommunityList sets field value
func (o *CommunityListRule) SetCommunityList(v BriefCommunityList) {
	o.CommunityList = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CommunityListRule) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityListRule) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *CommunityListRule) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CommunityListRule) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityListRule) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *CommunityListRule) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetAction returns the Action field value
func (o *CommunityListRule) GetAction() CommunityListRuleAction {
	if o == nil {
		var ret CommunityListRuleAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CommunityListRule) GetActionOk() (*CommunityListRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *CommunityListRule) SetAction(v CommunityListRuleAction) {
	o.Action = v
}

// GetCommunity returns the Community field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommunityListRule) GetCommunity() BriefCommunity {
	if o == nil || IsNil(o.Community.Get()) {
		var ret BriefCommunity
		return ret
	}
	return *o.Community.Get()
}

// GetCommunityOk returns a tuple with the Community field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityListRule) GetCommunityOk() (*BriefCommunity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Community.Get(), o.Community.IsSet()
}

// HasCommunity returns a boolean if a field has been set.
func (o *CommunityListRule) HasCommunity() bool {
	if o != nil && o.Community.IsSet() {
		return true
	}

	return false
}

// SetCommunity gets a reference to the given NullableBriefCommunity and assigns it to the Community field.
func (o *CommunityListRule) SetCommunity(v BriefCommunity) {
	o.Community.Set(&v)
}

// SetCommunityNil sets the value for Community to be an explicit nil
func (o *CommunityListRule) SetCommunityNil() {
	o.Community.Set(nil)
}

// UnsetCommunity ensures that no value is present for Community, not even an explicit nil
func (o *CommunityListRule) UnsetCommunity() {
	o.Community.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *CommunityListRule) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityListRule) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *CommunityListRule) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *CommunityListRule) SetComments(v string) {
	o.Comments = &v
}

func (o CommunityListRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityListRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["display"] = o.Display
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["community_list"] = o.CommunityList
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
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

func (o *CommunityListRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"display",
		"community_list",
		"created",
		"last_updated",
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

	varCommunityListRule := _CommunityListRule{}

	err = json.Unmarshal(data, &varCommunityListRule)

	if err != nil {
		return err
	}

	*o = CommunityListRule(varCommunityListRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "display")
		delete(additionalProperties, "description")
		delete(additionalProperties, "community_list")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "action")
		delete(additionalProperties, "community")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommunityListRule struct {
	value *CommunityListRule
	isSet bool
}

func (v NullableCommunityListRule) Get() *CommunityListRule {
	return v.value
}

func (v *NullableCommunityListRule) Set(val *CommunityListRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityListRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityListRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityListRule(val *CommunityListRule) *NullableCommunityListRule {
	return &NullableCommunityListRule{value: val, isSet: true}
}

func (v NullableCommunityListRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityListRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
