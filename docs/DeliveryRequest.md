# DeliveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purchase** | [**NestedPurchaseRequest**](NestedPurchaseRequest.md) |  | 
**Name** | **string** |  | 
**Date** | Pointer to **NullableString** | Date when this delivery was made | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ReceivingContact** | Pointer to [**NullableNestedContactRequest**](NestedContactRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeliveryRequest

`func NewDeliveryRequest(purchase NestedPurchaseRequest, name string, ) *DeliveryRequest`

NewDeliveryRequest instantiates a new DeliveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryRequestWithDefaults

`func NewDeliveryRequestWithDefaults() *DeliveryRequest`

NewDeliveryRequestWithDefaults instantiates a new DeliveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchase

`func (o *DeliveryRequest) GetPurchase() NestedPurchaseRequest`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *DeliveryRequest) GetPurchaseOk() (*NestedPurchaseRequest, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *DeliveryRequest) SetPurchase(v NestedPurchaseRequest)`

SetPurchase sets Purchase field to given value.


### GetName

`func (o *DeliveryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDate

`func (o *DeliveryRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DeliveryRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DeliveryRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *DeliveryRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *DeliveryRequest) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *DeliveryRequest) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *DeliveryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeliveryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeliveryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeliveryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *DeliveryRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DeliveryRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DeliveryRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DeliveryRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetReceivingContact

`func (o *DeliveryRequest) GetReceivingContact() NestedContactRequest`

GetReceivingContact returns the ReceivingContact field if non-nil, zero value otherwise.

### GetReceivingContactOk

`func (o *DeliveryRequest) GetReceivingContactOk() (*NestedContactRequest, bool)`

GetReceivingContactOk returns a tuple with the ReceivingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingContact

`func (o *DeliveryRequest) SetReceivingContact(v NestedContactRequest)`

SetReceivingContact sets ReceivingContact field to given value.

### HasReceivingContact

`func (o *DeliveryRequest) HasReceivingContact() bool`

HasReceivingContact returns a boolean if a field has been set.

### SetReceivingContactNil

`func (o *DeliveryRequest) SetReceivingContactNil(b bool)`

 SetReceivingContactNil sets the value for ReceivingContact to be an explicit nil

### UnsetReceivingContact
`func (o *DeliveryRequest) UnsetReceivingContact()`

UnsetReceivingContact ensures that no value is present for ReceivingContact, not even an explicit nil
### GetTags

`func (o *DeliveryRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeliveryRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeliveryRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeliveryRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DeliveryRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DeliveryRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DeliveryRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DeliveryRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


