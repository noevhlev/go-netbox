# BGPSessionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**BGPSessionStatusValue**](BGPSessionStatusValue.md) |  | [optional] 
**Label** | Pointer to [**BGPSessionStatusLabel**](BGPSessionStatusLabel.md) |  | [optional] 

## Methods

### NewBGPSessionStatus

`func NewBGPSessionStatus() *BGPSessionStatus`

NewBGPSessionStatus instantiates a new BGPSessionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPSessionStatusWithDefaults

`func NewBGPSessionStatusWithDefaults() *BGPSessionStatus`

NewBGPSessionStatusWithDefaults instantiates a new BGPSessionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *BGPSessionStatus) GetValue() BGPSessionStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BGPSessionStatus) GetValueOk() (*BGPSessionStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BGPSessionStatus) SetValue(v BGPSessionStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *BGPSessionStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *BGPSessionStatus) GetLabel() BGPSessionStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BGPSessionStatus) GetLabelOk() (*BGPSessionStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BGPSessionStatus) SetLabel(v BGPSessionStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BGPSessionStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


