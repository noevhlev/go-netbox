# InventoryItemTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Slug** | **string** |  | 
**Manufacturer** | [**NestedManufacturerRequest**](NestedManufacturerRequest.md) |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**InventoryitemGroup** | Pointer to [**NullableNestedInventoryItemGroupRequest**](NestedInventoryItemGroupRequest.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInventoryItemTypeRequest

`func NewInventoryItemTypeRequest(model string, slug string, manufacturer NestedManufacturerRequest, ) *InventoryItemTypeRequest`

NewInventoryItemTypeRequest instantiates a new InventoryItemTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemTypeRequestWithDefaults

`func NewInventoryItemTypeRequestWithDefaults() *InventoryItemTypeRequest`

NewInventoryItemTypeRequestWithDefaults instantiates a new InventoryItemTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *InventoryItemTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InventoryItemTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InventoryItemTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *InventoryItemTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *InventoryItemTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *InventoryItemTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetManufacturer

`func (o *InventoryItemTypeRequest) GetManufacturer() NestedManufacturerRequest`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InventoryItemTypeRequest) GetManufacturerOk() (*NestedManufacturerRequest, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InventoryItemTypeRequest) SetManufacturer(v NestedManufacturerRequest)`

SetManufacturer sets Manufacturer field to given value.


### GetPartNumber

`func (o *InventoryItemTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *InventoryItemTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *InventoryItemTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *InventoryItemTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetInventoryitemGroup

`func (o *InventoryItemTypeRequest) GetInventoryitemGroup() NestedInventoryItemGroupRequest`

GetInventoryitemGroup returns the InventoryitemGroup field if non-nil, zero value otherwise.

### GetInventoryitemGroupOk

`func (o *InventoryItemTypeRequest) GetInventoryitemGroupOk() (*NestedInventoryItemGroupRequest, bool)`

GetInventoryitemGroupOk returns a tuple with the InventoryitemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemGroup

`func (o *InventoryItemTypeRequest) SetInventoryitemGroup(v NestedInventoryItemGroupRequest)`

SetInventoryitemGroup sets InventoryitemGroup field to given value.

### HasInventoryitemGroup

`func (o *InventoryItemTypeRequest) HasInventoryitemGroup() bool`

HasInventoryitemGroup returns a boolean if a field has been set.

### SetInventoryitemGroupNil

`func (o *InventoryItemTypeRequest) SetInventoryitemGroupNil(b bool)`

 SetInventoryitemGroupNil sets the value for InventoryitemGroup to be an explicit nil

### UnsetInventoryitemGroup
`func (o *InventoryItemTypeRequest) UnsetInventoryitemGroup()`

UnsetInventoryitemGroup ensures that no value is present for InventoryitemGroup, not even an explicit nil
### GetComments

`func (o *InventoryItemTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *InventoryItemTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *InventoryItemTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *InventoryItemTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *InventoryItemTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventoryItemTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventoryItemTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InventoryItemTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *InventoryItemTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InventoryItemTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InventoryItemTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InventoryItemTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


