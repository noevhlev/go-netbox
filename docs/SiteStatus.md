# SiteStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PatchedWritableSiteRequestStatus**](PatchedWritableSiteRequestStatus.md) |  | [optional] 
**Label** | Pointer to [**SiteStatusLabel**](SiteStatusLabel.md) |  | [optional] 

## Methods

### NewSiteStatus

`func NewSiteStatus() *SiteStatus`

NewSiteStatus instantiates a new SiteStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteStatusWithDefaults

`func NewSiteStatusWithDefaults() *SiteStatus`

NewSiteStatusWithDefaults instantiates a new SiteStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SiteStatus) GetValue() PatchedWritableSiteRequestStatus`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SiteStatus) GetValueOk() (*PatchedWritableSiteRequestStatus, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SiteStatus) SetValue(v PatchedWritableSiteRequestStatus)`

SetValue sets Value field to given value.

### HasValue

`func (o *SiteStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *SiteStatus) GetLabel() SiteStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SiteStatus) GetLabelOk() (*SiteStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SiteStatus) SetLabel(v SiteStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SiteStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


