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

// checks if the ASNRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ASNRange{}

// ASNRange Adds support for custom fields and tags.
type ASNRange struct {
	Id                   int32                  `json:"id"`
	Url                  string                 `json:"url"`
	Display              string                 `json:"display"`
	Name                 string                 `json:"name"`
	Slug                 string                 `json:"slug"`
	Rir                  BriefRIR               `json:"rir"`
	Start                int64                  `json:"start"`
	End                  int64                  `json:"end"`
	Tenant               NullableBriefTenant    `json:"tenant,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AsnCount             int32                  `json:"asn_count"`
	AdditionalProperties map[string]interface{}
}

type _ASNRange ASNRange

// NewASNRange instantiates a new ASNRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewASNRange(id int32, url string, display string, name string, slug string, rir BriefRIR, start int64, end int64, created NullableTime, lastUpdated NullableTime, asnCount int32) *ASNRange {
	this := ASNRange{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Rir = rir
	this.Start = start
	this.End = end
	this.Created = created
	this.LastUpdated = lastUpdated
	this.AsnCount = asnCount
	return &this
}

// NewASNRangeWithDefaults instantiates a new ASNRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewASNRangeWithDefaults() *ASNRange {
	this := ASNRange{}
	return &this
}

// GetId returns the Id field value
func (o *ASNRange) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ASNRange) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ASNRange) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ASNRange) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ASNRange) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ASNRange) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *ASNRange) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ASNRange) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *ASNRange) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *ASNRange) SetSlug(v string) {
	o.Slug = v
}

// GetRir returns the Rir field value
func (o *ASNRange) GetRir() BriefRIR {
	if o == nil {
		var ret BriefRIR
		return ret
	}

	return o.Rir
}

// GetRirOk returns a tuple with the Rir field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetRirOk() (*BriefRIR, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rir, true
}

// SetRir sets field value
func (o *ASNRange) SetRir(v BriefRIR) {
	o.Rir = v
}

// GetStart returns the Start field value
func (o *ASNRange) GetStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ASNRange) SetStart(v int64) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *ASNRange) GetEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ASNRange) SetEnd(v int64) {
	o.End = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ASNRange) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASNRange) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *ASNRange) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *ASNRange) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *ASNRange) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *ASNRange) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ASNRange) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASNRange) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ASNRange) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ASNRange) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ASNRange) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASNRange) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ASNRange) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *ASNRange) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ASNRange) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASNRange) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ASNRange) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ASNRange) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ASNRange) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASNRange) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ASNRange) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ASNRange) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASNRange) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ASNRange) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetAsnCount returns the AsnCount field value
func (o *ASNRange) GetAsnCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AsnCount
}

// GetAsnCountOk returns a tuple with the AsnCount field value
// and a boolean to check if the value has been set.
func (o *ASNRange) GetAsnCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AsnCount, true
}

// SetAsnCount sets field value
func (o *ASNRange) SetAsnCount(v int32) {
	o.AsnCount = v
}

func (o ASNRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ASNRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["rir"] = o.Rir
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["asn_count"] = o.AsnCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ASNRange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
		"rir",
		"start",
		"end",
		"created",
		"last_updated",
		"asn_count",
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

	varASNRange := _ASNRange{}

	err = json.Unmarshal(data, &varASNRange)

	if err != nil {
		return err
	}

	*o = ASNRange(varASNRange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "rir")
		delete(additionalProperties, "start")
		delete(additionalProperties, "end")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "asn_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableASNRange struct {
	value *ASNRange
	isSet bool
}

func (v NullableASNRange) Get() *ASNRange {
	return v.value
}

func (v *NullableASNRange) Set(val *ASNRange) {
	v.value = val
	v.isSet = true
}

func (v NullableASNRange) IsSet() bool {
	return v.isSet
}

func (v *NullableASNRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableASNRange(val *ASNRange) *NullableASNRange {
	return &NullableASNRange{value: val, isSet: true}
}

func (v NullableASNRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableASNRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
