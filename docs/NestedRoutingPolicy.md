# NestedRoutingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Display** | **string** |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewNestedRoutingPolicy

`func NewNestedRoutingPolicy(id int32, url string, name string, display string, ) *NestedRoutingPolicy`

NewNestedRoutingPolicy instantiates a new NestedRoutingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedRoutingPolicyWithDefaults

`func NewNestedRoutingPolicyWithDefaults() *NestedRoutingPolicy`

NewNestedRoutingPolicyWithDefaults instantiates a new NestedRoutingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedRoutingPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedRoutingPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedRoutingPolicy) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedRoutingPolicy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedRoutingPolicy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedRoutingPolicy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *NestedRoutingPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedRoutingPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedRoutingPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDisplay

`func (o *NestedRoutingPolicy) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedRoutingPolicy) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedRoutingPolicy) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDescription

`func (o *NestedRoutingPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NestedRoutingPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NestedRoutingPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NestedRoutingPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


