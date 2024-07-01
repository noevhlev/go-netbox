# PatchedWritablePurchaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supplier** | Pointer to **int32** | Legal entity this purchase was made at | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**NestedPurchaseStatus**](NestedPurchaseStatus.md) |  | [optional] 
**Date** | Pointer to **NullableString** | Date when this purchase was made | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritablePurchaseRequest

`func NewPatchedWritablePurchaseRequest() *PatchedWritablePurchaseRequest`

NewPatchedWritablePurchaseRequest instantiates a new PatchedWritablePurchaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePurchaseRequestWithDefaults

`func NewPatchedWritablePurchaseRequestWithDefaults() *PatchedWritablePurchaseRequest`

NewPatchedWritablePurchaseRequestWithDefaults instantiates a new PatchedWritablePurchaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupplier

`func (o *PatchedWritablePurchaseRequest) GetSupplier() int32`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *PatchedWritablePurchaseRequest) GetSupplierOk() (*int32, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *PatchedWritablePurchaseRequest) SetSupplier(v int32)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *PatchedWritablePurchaseRequest) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritablePurchaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritablePurchaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritablePurchaseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritablePurchaseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritablePurchaseRequest) GetStatus() NestedPurchaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritablePurchaseRequest) GetStatusOk() (*NestedPurchaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritablePurchaseRequest) SetStatus(v NestedPurchaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritablePurchaseRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDate

`func (o *PatchedWritablePurchaseRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PatchedWritablePurchaseRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PatchedWritablePurchaseRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PatchedWritablePurchaseRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *PatchedWritablePurchaseRequest) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *PatchedWritablePurchaseRequest) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *PatchedWritablePurchaseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritablePurchaseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritablePurchaseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritablePurchaseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritablePurchaseRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritablePurchaseRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritablePurchaseRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritablePurchaseRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritablePurchaseRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritablePurchaseRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritablePurchaseRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritablePurchaseRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritablePurchaseRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePurchaseRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePurchaseRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePurchaseRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


