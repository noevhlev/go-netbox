# NestedBGPPeerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewNestedBGPPeerGroup

`func NewNestedBGPPeerGroup(id int32, display string, url string, name string, ) *NestedBGPPeerGroup`

NewNestedBGPPeerGroup instantiates a new NestedBGPPeerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedBGPPeerGroupWithDefaults

`func NewNestedBGPPeerGroupWithDefaults() *NestedBGPPeerGroup`

NewNestedBGPPeerGroupWithDefaults instantiates a new NestedBGPPeerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedBGPPeerGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedBGPPeerGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedBGPPeerGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetDisplay

`func (o *NestedBGPPeerGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedBGPPeerGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedBGPPeerGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *NestedBGPPeerGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedBGPPeerGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedBGPPeerGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *NestedBGPPeerGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedBGPPeerGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedBGPPeerGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NestedBGPPeerGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NestedBGPPeerGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NestedBGPPeerGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NestedBGPPeerGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


