/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the SavedFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedFilter{}

// SavedFilter Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type SavedFilter struct {
	Id                   int32         `json:"id"`
	Url                  string        `json:"url"`
	Display              string        `json:"display"`
	ContentTypes         []string      `json:"content_types"`
	Name                 string        `json:"name"`
	Slug                 string        `json:"slug"`
	Description          *string       `json:"description,omitempty"`
	User                 NullableInt32 `json:"user,omitempty"`
	Weight               *int32        `json:"weight,omitempty"`
	Enabled              *bool         `json:"enabled,omitempty"`
	Shared               *bool         `json:"shared,omitempty"`
	Parameters           interface{}   `json:"parameters"`
	Created              NullableTime  `json:"created"`
	LastUpdated          NullableTime  `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _SavedFilter SavedFilter

// NewSavedFilter instantiates a new SavedFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedFilter(id int32, url string, display string, contentTypes []string, name string, slug string, parameters interface{}, created NullableTime, lastUpdated NullableTime) *SavedFilter {
	this := SavedFilter{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.ContentTypes = contentTypes
	this.Name = name
	this.Slug = slug
	this.Parameters = parameters
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewSavedFilterWithDefaults instantiates a new SavedFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedFilterWithDefaults() *SavedFilter {
	this := SavedFilter{}
	return &this
}

// GetId returns the Id field value
func (o *SavedFilter) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SavedFilter) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *SavedFilter) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SavedFilter) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *SavedFilter) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *SavedFilter) SetDisplay(v string) {
	o.Display = v
}

// GetContentTypes returns the ContentTypes field value
func (o *SavedFilter) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *SavedFilter) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *SavedFilter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SavedFilter) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *SavedFilter) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *SavedFilter) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SavedFilter) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SavedFilter) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SavedFilter) SetDescription(v string) {
	o.Description = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedFilter) GetUser() int32 {
	if o == nil || IsNil(o.User.Get()) {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedFilter) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *SavedFilter) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *SavedFilter) SetUser(v int32) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *SavedFilter) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *SavedFilter) UnsetUser() {
	o.User.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *SavedFilter) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *SavedFilter) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *SavedFilter) SetWeight(v int32) {
	o.Weight = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SavedFilter) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SavedFilter) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SavedFilter) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *SavedFilter) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedFilter) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *SavedFilter) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *SavedFilter) SetShared(v bool) {
	o.Shared = &v
}

// GetParameters returns the Parameters field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SavedFilter) GetParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedFilter) GetParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *SavedFilter) SetParameters(v interface{}) {
	o.Parameters = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *SavedFilter) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedFilter) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *SavedFilter) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *SavedFilter) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedFilter) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *SavedFilter) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o SavedFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SavedFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"content_types",
		"name",
		"slug",
		"parameters",
		"created",
		"last_updated",
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

	varSavedFilter := _SavedFilter{}

	err = json.Unmarshal(data, &varSavedFilter)

	if err != nil {
		return err
	}

	*o = SavedFilter(varSavedFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "user")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "shared")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavedFilter struct {
	value *SavedFilter
	isSet bool
}

func (v NullableSavedFilter) Get() *SavedFilter {
	return v.value
}

func (v *NullableSavedFilter) Set(val *SavedFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedFilter(val *SavedFilter) *NullableSavedFilter {
	return &NullableSavedFilter{value: val, isSet: true}
}

func (v NullableSavedFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
