# BriefPrefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Family** | [**AggregateFamily**](AggregateFamily.md) |  | 
**Prefix** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Depth** | **int32** |  | [readonly] 

## Methods

### NewBriefPrefix

`func NewBriefPrefix(id int32, url string, display string, family AggregateFamily, prefix string, depth int32, ) *BriefPrefix`

NewBriefPrefix instantiates a new BriefPrefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefPrefixWithDefaults

`func NewBriefPrefixWithDefaults() *BriefPrefix`

NewBriefPrefixWithDefaults instantiates a new BriefPrefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefPrefix) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefPrefix) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefPrefix) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefPrefix) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefPrefix) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefPrefix) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefPrefix) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefPrefix) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefPrefix) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetFamily

`func (o *BriefPrefix) GetFamily() AggregateFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *BriefPrefix) GetFamilyOk() (*AggregateFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *BriefPrefix) SetFamily(v AggregateFamily)`

SetFamily sets Family field to given value.


### GetPrefix

`func (o *BriefPrefix) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BriefPrefix) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BriefPrefix) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetDescription

`func (o *BriefPrefix) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefPrefix) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefPrefix) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefPrefix) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDepth

`func (o *BriefPrefix) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *BriefPrefix) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *BriefPrefix) SetDepth(v int32)`

SetDepth sets Depth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


