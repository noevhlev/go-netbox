# Supplier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**AssetCount** | **int32** |  | [readonly] 
**PurchaseCount** | **int32** |  | [readonly] 
**DeliveryCount** | **int32** |  | [readonly] 

## Methods

### NewSupplier

`func NewSupplier(id int32, url string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, assetCount int32, purchaseCount int32, deliveryCount int32, ) *Supplier`

NewSupplier instantiates a new Supplier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierWithDefaults

`func NewSupplierWithDefaults() *Supplier`

NewSupplierWithDefaults instantiates a new Supplier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Supplier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Supplier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Supplier) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Supplier) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Supplier) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Supplier) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *Supplier) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Supplier) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Supplier) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *Supplier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Supplier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Supplier) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Supplier) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Supplier) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Supplier) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *Supplier) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Supplier) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Supplier) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Supplier) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *Supplier) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Supplier) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Supplier) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Supplier) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Supplier) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Supplier) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Supplier) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Supplier) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Supplier) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Supplier) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Supplier) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Supplier) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Supplier) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Supplier) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Supplier) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Supplier) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Supplier) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Supplier) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Supplier) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Supplier) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Supplier) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Supplier) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetAssetCount

`func (o *Supplier) GetAssetCount() int32`

GetAssetCount returns the AssetCount field if non-nil, zero value otherwise.

### GetAssetCountOk

`func (o *Supplier) GetAssetCountOk() (*int32, bool)`

GetAssetCountOk returns a tuple with the AssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCount

`func (o *Supplier) SetAssetCount(v int32)`

SetAssetCount sets AssetCount field to given value.


### GetPurchaseCount

`func (o *Supplier) GetPurchaseCount() int32`

GetPurchaseCount returns the PurchaseCount field if non-nil, zero value otherwise.

### GetPurchaseCountOk

`func (o *Supplier) GetPurchaseCountOk() (*int32, bool)`

GetPurchaseCountOk returns a tuple with the PurchaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCount

`func (o *Supplier) SetPurchaseCount(v int32)`

SetPurchaseCount sets PurchaseCount field to given value.


### GetDeliveryCount

`func (o *Supplier) GetDeliveryCount() int32`

GetDeliveryCount returns the DeliveryCount field if non-nil, zero value otherwise.

### GetDeliveryCountOk

`func (o *Supplier) GetDeliveryCountOk() (*int32, bool)`

GetDeliveryCountOk returns a tuple with the DeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCount

`func (o *Supplier) SetDeliveryCount(v int32)`

SetDeliveryCount sets DeliveryCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


