# NestedPurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Supplier** | [**NestedSupplier**](NestedSupplier.md) |  | [readonly] 
**Name** | **string** |  | 
**Status** | [**NestedPurchaseStatus**](NestedPurchaseStatus.md) |  | 
**Date** | Pointer to **NullableString** | Date when this purchase was made | [optional] 

## Methods

### NewNestedPurchase

`func NewNestedPurchase(id int32, url string, display string, supplier NestedSupplier, name string, status NestedPurchaseStatus, ) *NestedPurchase`

NewNestedPurchase instantiates a new NestedPurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedPurchaseWithDefaults

`func NewNestedPurchaseWithDefaults() *NestedPurchase`

NewNestedPurchaseWithDefaults instantiates a new NestedPurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedPurchase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedPurchase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedPurchase) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedPurchase) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedPurchase) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedPurchase) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NestedPurchase) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedPurchase) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedPurchase) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetSupplier

`func (o *NestedPurchase) GetSupplier() NestedSupplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *NestedPurchase) GetSupplierOk() (*NestedSupplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *NestedPurchase) SetSupplier(v NestedSupplier)`

SetSupplier sets Supplier field to given value.


### GetName

`func (o *NestedPurchase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedPurchase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedPurchase) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *NestedPurchase) GetStatus() NestedPurchaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NestedPurchase) GetStatusOk() (*NestedPurchaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NestedPurchase) SetStatus(v NestedPurchaseStatus)`

SetStatus sets Status field to given value.


### GetDate

`func (o *NestedPurchase) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NestedPurchase) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NestedPurchase) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *NestedPurchase) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *NestedPurchase) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *NestedPurchase) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


