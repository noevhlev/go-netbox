# PurchaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supplier** | [**NestedSupplierRequest**](NestedSupplierRequest.md) |  | 
**Name** | **string** |  | 
**Status** | [**NestedPurchaseStatus**](NestedPurchaseStatus.md) |  | 
**Date** | Pointer to **NullableString** | Date when this purchase was made | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPurchaseRequest

`func NewPurchaseRequest(supplier NestedSupplierRequest, name string, status NestedPurchaseStatus, ) *PurchaseRequest`

NewPurchaseRequest instantiates a new PurchaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseRequestWithDefaults

`func NewPurchaseRequestWithDefaults() *PurchaseRequest`

NewPurchaseRequestWithDefaults instantiates a new PurchaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupplier

`func (o *PurchaseRequest) GetSupplier() NestedSupplierRequest`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *PurchaseRequest) GetSupplierOk() (*NestedSupplierRequest, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *PurchaseRequest) SetSupplier(v NestedSupplierRequest)`

SetSupplier sets Supplier field to given value.


### GetName

`func (o *PurchaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PurchaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PurchaseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PurchaseRequest) GetStatus() NestedPurchaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PurchaseRequest) GetStatusOk() (*NestedPurchaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PurchaseRequest) SetStatus(v NestedPurchaseStatus)`

SetStatus sets Status field to given value.


### GetDate

`func (o *PurchaseRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PurchaseRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PurchaseRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PurchaseRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *PurchaseRequest) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *PurchaseRequest) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *PurchaseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PurchaseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PurchaseRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PurchaseRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PurchaseRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PurchaseRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PurchaseRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PurchaseRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PurchaseRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PurchaseRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PurchaseRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PurchaseRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PurchaseRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PurchaseRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


