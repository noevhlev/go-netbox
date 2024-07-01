# NestedDeliveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Date** | Pointer to **NullableString** | Date when this delivery was made | [optional] 

## Methods

### NewNestedDeliveryRequest

`func NewNestedDeliveryRequest(name string, ) *NestedDeliveryRequest`

NewNestedDeliveryRequest instantiates a new NestedDeliveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedDeliveryRequestWithDefaults

`func NewNestedDeliveryRequestWithDefaults() *NestedDeliveryRequest`

NewNestedDeliveryRequestWithDefaults instantiates a new NestedDeliveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NestedDeliveryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedDeliveryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedDeliveryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDate

`func (o *NestedDeliveryRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NestedDeliveryRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NestedDeliveryRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *NestedDeliveryRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *NestedDeliveryRequest) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *NestedDeliveryRequest) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


