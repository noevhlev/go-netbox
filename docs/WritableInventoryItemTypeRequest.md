# WritableInventoryItemTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Slug** | **string** |  | 
**Manufacturer** | **int32** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**InventoryitemGroup** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableInventoryItemTypeRequest

`func NewWritableInventoryItemTypeRequest(model string, slug string, manufacturer int32, ) *WritableInventoryItemTypeRequest`

NewWritableInventoryItemTypeRequest instantiates a new WritableInventoryItemTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableInventoryItemTypeRequestWithDefaults

`func NewWritableInventoryItemTypeRequestWithDefaults() *WritableInventoryItemTypeRequest`

NewWritableInventoryItemTypeRequestWithDefaults instantiates a new WritableInventoryItemTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *WritableInventoryItemTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *WritableInventoryItemTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *WritableInventoryItemTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *WritableInventoryItemTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableInventoryItemTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableInventoryItemTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetManufacturer

`func (o *WritableInventoryItemTypeRequest) GetManufacturer() int32`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *WritableInventoryItemTypeRequest) GetManufacturerOk() (*int32, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *WritableInventoryItemTypeRequest) SetManufacturer(v int32)`

SetManufacturer sets Manufacturer field to given value.


### GetPartNumber

`func (o *WritableInventoryItemTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *WritableInventoryItemTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *WritableInventoryItemTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *WritableInventoryItemTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetInventoryitemGroup

`func (o *WritableInventoryItemTypeRequest) GetInventoryitemGroup() int32`

GetInventoryitemGroup returns the InventoryitemGroup field if non-nil, zero value otherwise.

### GetInventoryitemGroupOk

`func (o *WritableInventoryItemTypeRequest) GetInventoryitemGroupOk() (*int32, bool)`

GetInventoryitemGroupOk returns a tuple with the InventoryitemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemGroup

`func (o *WritableInventoryItemTypeRequest) SetInventoryitemGroup(v int32)`

SetInventoryitemGroup sets InventoryitemGroup field to given value.

### HasInventoryitemGroup

`func (o *WritableInventoryItemTypeRequest) HasInventoryitemGroup() bool`

HasInventoryitemGroup returns a boolean if a field has been set.

### SetInventoryitemGroupNil

`func (o *WritableInventoryItemTypeRequest) SetInventoryitemGroupNil(b bool)`

 SetInventoryitemGroupNil sets the value for InventoryitemGroup to be an explicit nil

### UnsetInventoryitemGroup
`func (o *WritableInventoryItemTypeRequest) UnsetInventoryitemGroup()`

UnsetInventoryitemGroup ensures that no value is present for InventoryitemGroup, not even an explicit nil
### GetComments

`func (o *WritableInventoryItemTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableInventoryItemTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableInventoryItemTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableInventoryItemTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableInventoryItemTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableInventoryItemTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableInventoryItemTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableInventoryItemTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableInventoryItemTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableInventoryItemTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableInventoryItemTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableInventoryItemTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


