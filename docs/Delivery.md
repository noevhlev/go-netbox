# Delivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Purchase** | [**NestedPurchase**](NestedPurchase.md) |  | 
**Name** | **string** |  | 
**Date** | Pointer to **NullableString** | Date when this delivery was made | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ReceivingContact** | Pointer to [**NullableNestedContact**](NestedContact.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**AssetCount** | **int32** |  | [readonly] 

## Methods

### NewDelivery

`func NewDelivery(id int32, url string, display string, purchase NestedPurchase, name string, created NullableTime, lastUpdated NullableTime, assetCount int32, ) *Delivery`

NewDelivery instantiates a new Delivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryWithDefaults

`func NewDeliveryWithDefaults() *Delivery`

NewDeliveryWithDefaults instantiates a new Delivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Delivery) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Delivery) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Delivery) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Delivery) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Delivery) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Delivery) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *Delivery) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Delivery) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Delivery) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetPurchase

`func (o *Delivery) GetPurchase() NestedPurchase`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *Delivery) GetPurchaseOk() (*NestedPurchase, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *Delivery) SetPurchase(v NestedPurchase)`

SetPurchase sets Purchase field to given value.


### GetName

`func (o *Delivery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Delivery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Delivery) SetName(v string)`

SetName sets Name field to given value.


### GetDate

`func (o *Delivery) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Delivery) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Delivery) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Delivery) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *Delivery) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *Delivery) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *Delivery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Delivery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Delivery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Delivery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *Delivery) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Delivery) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Delivery) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Delivery) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetReceivingContact

`func (o *Delivery) GetReceivingContact() NestedContact`

GetReceivingContact returns the ReceivingContact field if non-nil, zero value otherwise.

### GetReceivingContactOk

`func (o *Delivery) GetReceivingContactOk() (*NestedContact, bool)`

GetReceivingContactOk returns a tuple with the ReceivingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingContact

`func (o *Delivery) SetReceivingContact(v NestedContact)`

SetReceivingContact sets ReceivingContact field to given value.

### HasReceivingContact

`func (o *Delivery) HasReceivingContact() bool`

HasReceivingContact returns a boolean if a field has been set.

### SetReceivingContactNil

`func (o *Delivery) SetReceivingContactNil(b bool)`

 SetReceivingContactNil sets the value for ReceivingContact to be an explicit nil

### UnsetReceivingContact
`func (o *Delivery) UnsetReceivingContact()`

UnsetReceivingContact ensures that no value is present for ReceivingContact, not even an explicit nil
### GetTags

`func (o *Delivery) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Delivery) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Delivery) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Delivery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Delivery) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Delivery) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Delivery) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Delivery) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Delivery) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Delivery) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Delivery) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Delivery) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Delivery) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Delivery) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Delivery) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Delivery) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Delivery) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Delivery) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetAssetCount

`func (o *Delivery) GetAssetCount() int32`

GetAssetCount returns the AssetCount field if non-nil, zero value otherwise.

### GetAssetCountOk

`func (o *Delivery) GetAssetCountOk() (*int32, bool)`

GetAssetCountOk returns a tuple with the AssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCount

`func (o *Delivery) SetAssetCount(v int32)`

SetAssetCount sets AssetCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


