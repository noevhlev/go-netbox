# InventoryItemType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**Manufacturer** | [**NestedManufacturer**](NestedManufacturer.md) |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**InventoryitemGroup** | Pointer to [**NullableNestedInventoryItemGroup**](NestedInventoryItemGroup.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**AssetCount** | **int32** |  | [readonly] 

## Methods

### NewInventoryItemType

`func NewInventoryItemType(id int32, url string, display string, model string, slug string, manufacturer NestedManufacturer, created NullableTime, lastUpdated NullableTime, assetCount int32, ) *InventoryItemType`

NewInventoryItemType instantiates a new InventoryItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemTypeWithDefaults

`func NewInventoryItemTypeWithDefaults() *InventoryItemType`

NewInventoryItemTypeWithDefaults instantiates a new InventoryItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryItemType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryItemType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryItemType) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *InventoryItemType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InventoryItemType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InventoryItemType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *InventoryItemType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InventoryItemType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InventoryItemType) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetModel

`func (o *InventoryItemType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InventoryItemType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InventoryItemType) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *InventoryItemType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *InventoryItemType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *InventoryItemType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetManufacturer

`func (o *InventoryItemType) GetManufacturer() NestedManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InventoryItemType) GetManufacturerOk() (*NestedManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InventoryItemType) SetManufacturer(v NestedManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetPartNumber

`func (o *InventoryItemType) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *InventoryItemType) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *InventoryItemType) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *InventoryItemType) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetInventoryitemGroup

`func (o *InventoryItemType) GetInventoryitemGroup() NestedInventoryItemGroup`

GetInventoryitemGroup returns the InventoryitemGroup field if non-nil, zero value otherwise.

### GetInventoryitemGroupOk

`func (o *InventoryItemType) GetInventoryitemGroupOk() (*NestedInventoryItemGroup, bool)`

GetInventoryitemGroupOk returns a tuple with the InventoryitemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemGroup

`func (o *InventoryItemType) SetInventoryitemGroup(v NestedInventoryItemGroup)`

SetInventoryitemGroup sets InventoryitemGroup field to given value.

### HasInventoryitemGroup

`func (o *InventoryItemType) HasInventoryitemGroup() bool`

HasInventoryitemGroup returns a boolean if a field has been set.

### SetInventoryitemGroupNil

`func (o *InventoryItemType) SetInventoryitemGroupNil(b bool)`

 SetInventoryitemGroupNil sets the value for InventoryitemGroup to be an explicit nil

### UnsetInventoryitemGroup
`func (o *InventoryItemType) UnsetInventoryitemGroup()`

UnsetInventoryitemGroup ensures that no value is present for InventoryitemGroup, not even an explicit nil
### GetComments

`func (o *InventoryItemType) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *InventoryItemType) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *InventoryItemType) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *InventoryItemType) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *InventoryItemType) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventoryItemType) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventoryItemType) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InventoryItemType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *InventoryItemType) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InventoryItemType) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InventoryItemType) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InventoryItemType) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *InventoryItemType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InventoryItemType) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InventoryItemType) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *InventoryItemType) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *InventoryItemType) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *InventoryItemType) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InventoryItemType) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InventoryItemType) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *InventoryItemType) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *InventoryItemType) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetAssetCount

`func (o *InventoryItemType) GetAssetCount() int32`

GetAssetCount returns the AssetCount field if non-nil, zero value otherwise.

### GetAssetCountOk

`func (o *InventoryItemType) GetAssetCountOk() (*int32, bool)`

GetAssetCountOk returns a tuple with the AssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCount

`func (o *InventoryItemType) SetAssetCount(v int32)`

SetAssetCount sets AssetCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


