# Purchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Supplier** | [**NestedSupplier**](NestedSupplier.md) |  | 
**Name** | **string** |  | 
**Status** | [**NestedPurchaseStatus**](NestedPurchaseStatus.md) |  | 
**Date** | Pointer to **NullableString** | Date when this purchase was made | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**AssetCount** | **int32** |  | [readonly] 
**DeliveryCount** | **int32** |  | [readonly] 

## Methods

### NewPurchase

`func NewPurchase(id int32, url string, display string, supplier NestedSupplier, name string, status NestedPurchaseStatus, created NullableTime, lastUpdated NullableTime, assetCount int32, deliveryCount int32, ) *Purchase`

NewPurchase instantiates a new Purchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseWithDefaults

`func NewPurchaseWithDefaults() *Purchase`

NewPurchaseWithDefaults instantiates a new Purchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Purchase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Purchase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Purchase) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Purchase) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Purchase) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Purchase) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *Purchase) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Purchase) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Purchase) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetSupplier

`func (o *Purchase) GetSupplier() NestedSupplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *Purchase) GetSupplierOk() (*NestedSupplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *Purchase) SetSupplier(v NestedSupplier)`

SetSupplier sets Supplier field to given value.


### GetName

`func (o *Purchase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Purchase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Purchase) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Purchase) GetStatus() NestedPurchaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Purchase) GetStatusOk() (*NestedPurchaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Purchase) SetStatus(v NestedPurchaseStatus)`

SetStatus sets Status field to given value.


### GetDate

`func (o *Purchase) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Purchase) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Purchase) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Purchase) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *Purchase) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *Purchase) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *Purchase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Purchase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Purchase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Purchase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *Purchase) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Purchase) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Purchase) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Purchase) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Purchase) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Purchase) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Purchase) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Purchase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Purchase) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Purchase) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Purchase) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Purchase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Purchase) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Purchase) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Purchase) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Purchase) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Purchase) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Purchase) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Purchase) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Purchase) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Purchase) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Purchase) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetAssetCount

`func (o *Purchase) GetAssetCount() int32`

GetAssetCount returns the AssetCount field if non-nil, zero value otherwise.

### GetAssetCountOk

`func (o *Purchase) GetAssetCountOk() (*int32, bool)`

GetAssetCountOk returns a tuple with the AssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCount

`func (o *Purchase) SetAssetCount(v int32)`

SetAssetCount sets AssetCount field to given value.


### GetDeliveryCount

`func (o *Purchase) GetDeliveryCount() int32`

GetDeliveryCount returns the DeliveryCount field if non-nil, zero value otherwise.

### GetDeliveryCountOk

`func (o *Purchase) GetDeliveryCountOk() (*int32, bool)`

GetDeliveryCountOk returns a tuple with the DeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCount

`func (o *Purchase) SetDeliveryCount(v int32)`

SetDeliveryCount sets DeliveryCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


