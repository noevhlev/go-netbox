# NestedPurchaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | [**NestedPurchaseStatus**](NestedPurchaseStatus.md) |  | 
**Date** | Pointer to **NullableString** | Date when this purchase was made | [optional] 

## Methods

### NewNestedPurchaseRequest

`func NewNestedPurchaseRequest(name string, status NestedPurchaseStatus, ) *NestedPurchaseRequest`

NewNestedPurchaseRequest instantiates a new NestedPurchaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedPurchaseRequestWithDefaults

`func NewNestedPurchaseRequestWithDefaults() *NestedPurchaseRequest`

NewNestedPurchaseRequestWithDefaults instantiates a new NestedPurchaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NestedPurchaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedPurchaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedPurchaseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *NestedPurchaseRequest) GetStatus() NestedPurchaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NestedPurchaseRequest) GetStatusOk() (*NestedPurchaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NestedPurchaseRequest) SetStatus(v NestedPurchaseStatus)`

SetStatus sets Status field to given value.


### GetDate

`func (o *NestedPurchaseRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NestedPurchaseRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NestedPurchaseRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *NestedPurchaseRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *NestedPurchaseRequest) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *NestedPurchaseRequest) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


