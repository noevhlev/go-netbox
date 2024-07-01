# PatchedWritableInventoryItemTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **int32** |  | [optional] 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**InventoryitemGroup** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableInventoryItemTypeRequest

`func NewPatchedWritableInventoryItemTypeRequest() *PatchedWritableInventoryItemTypeRequest`

NewPatchedWritableInventoryItemTypeRequest instantiates a new PatchedWritableInventoryItemTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableInventoryItemTypeRequestWithDefaults

`func NewPatchedWritableInventoryItemTypeRequestWithDefaults() *PatchedWritableInventoryItemTypeRequest`

NewPatchedWritableInventoryItemTypeRequestWithDefaults instantiates a new PatchedWritableInventoryItemTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *PatchedWritableInventoryItemTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PatchedWritableInventoryItemTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PatchedWritableInventoryItemTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PatchedWritableInventoryItemTypeRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedWritableInventoryItemTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedWritableInventoryItemTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedWritableInventoryItemTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedWritableInventoryItemTypeRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetManufacturer

`func (o *PatchedWritableInventoryItemTypeRequest) GetManufacturer() int32`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *PatchedWritableInventoryItemTypeRequest) GetManufacturerOk() (*int32, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *PatchedWritableInventoryItemTypeRequest) SetManufacturer(v int32)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *PatchedWritableInventoryItemTypeRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetPartNumber

`func (o *PatchedWritableInventoryItemTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *PatchedWritableInventoryItemTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *PatchedWritableInventoryItemTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *PatchedWritableInventoryItemTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetInventoryitemGroup

`func (o *PatchedWritableInventoryItemTypeRequest) GetInventoryitemGroup() int32`

GetInventoryitemGroup returns the InventoryitemGroup field if non-nil, zero value otherwise.

### GetInventoryitemGroupOk

`func (o *PatchedWritableInventoryItemTypeRequest) GetInventoryitemGroupOk() (*int32, bool)`

GetInventoryitemGroupOk returns a tuple with the InventoryitemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemGroup

`func (o *PatchedWritableInventoryItemTypeRequest) SetInventoryitemGroup(v int32)`

SetInventoryitemGroup sets InventoryitemGroup field to given value.

### HasInventoryitemGroup

`func (o *PatchedWritableInventoryItemTypeRequest) HasInventoryitemGroup() bool`

HasInventoryitemGroup returns a boolean if a field has been set.

### SetInventoryitemGroupNil

`func (o *PatchedWritableInventoryItemTypeRequest) SetInventoryitemGroupNil(b bool)`

 SetInventoryitemGroupNil sets the value for InventoryitemGroup to be an explicit nil

### UnsetInventoryitemGroup
`func (o *PatchedWritableInventoryItemTypeRequest) UnsetInventoryitemGroup()`

UnsetInventoryitemGroup ensures that no value is present for InventoryitemGroup, not even an explicit nil
### GetComments

`func (o *PatchedWritableInventoryItemTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableInventoryItemTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableInventoryItemTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableInventoryItemTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableInventoryItemTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableInventoryItemTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableInventoryItemTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableInventoryItemTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableInventoryItemTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableInventoryItemTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableInventoryItemTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableInventoryItemTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


