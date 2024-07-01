# PatchedWritableDeliveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purchase** | Pointer to **int32** | Purchase that this delivery is part of | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **NullableString** | Date when this delivery was made | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ReceivingContact** | Pointer to **NullableInt32** | Contact that accepted this delivery | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableDeliveryRequest

`func NewPatchedWritableDeliveryRequest() *PatchedWritableDeliveryRequest`

NewPatchedWritableDeliveryRequest instantiates a new PatchedWritableDeliveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableDeliveryRequestWithDefaults

`func NewPatchedWritableDeliveryRequestWithDefaults() *PatchedWritableDeliveryRequest`

NewPatchedWritableDeliveryRequestWithDefaults instantiates a new PatchedWritableDeliveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchase

`func (o *PatchedWritableDeliveryRequest) GetPurchase() int32`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *PatchedWritableDeliveryRequest) GetPurchaseOk() (*int32, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *PatchedWritableDeliveryRequest) SetPurchase(v int32)`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *PatchedWritableDeliveryRequest) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableDeliveryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableDeliveryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableDeliveryRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableDeliveryRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDate

`func (o *PatchedWritableDeliveryRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PatchedWritableDeliveryRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PatchedWritableDeliveryRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PatchedWritableDeliveryRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *PatchedWritableDeliveryRequest) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *PatchedWritableDeliveryRequest) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *PatchedWritableDeliveryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableDeliveryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableDeliveryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableDeliveryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableDeliveryRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableDeliveryRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableDeliveryRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableDeliveryRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetReceivingContact

`func (o *PatchedWritableDeliveryRequest) GetReceivingContact() int32`

GetReceivingContact returns the ReceivingContact field if non-nil, zero value otherwise.

### GetReceivingContactOk

`func (o *PatchedWritableDeliveryRequest) GetReceivingContactOk() (*int32, bool)`

GetReceivingContactOk returns a tuple with the ReceivingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingContact

`func (o *PatchedWritableDeliveryRequest) SetReceivingContact(v int32)`

SetReceivingContact sets ReceivingContact field to given value.

### HasReceivingContact

`func (o *PatchedWritableDeliveryRequest) HasReceivingContact() bool`

HasReceivingContact returns a boolean if a field has been set.

### SetReceivingContactNil

`func (o *PatchedWritableDeliveryRequest) SetReceivingContactNil(b bool)`

 SetReceivingContactNil sets the value for ReceivingContact to be an explicit nil

### UnsetReceivingContact
`func (o *PatchedWritableDeliveryRequest) UnsetReceivingContact()`

UnsetReceivingContact ensures that no value is present for ReceivingContact, not even an explicit nil
### GetTags

`func (o *PatchedWritableDeliveryRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableDeliveryRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableDeliveryRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableDeliveryRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableDeliveryRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableDeliveryRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableDeliveryRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableDeliveryRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


