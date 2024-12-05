# SessionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Userkey** | [**BriefUserKey**](BriefUserKey.md) |  | 
**SessionKey** | **string** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 

## Methods

### NewSessionKey

`func NewSessionKey(pk int32, id int32, url string, display string, userkey BriefUserKey, sessionKey string, created time.Time, ) *SessionKey`

NewSessionKey instantiates a new SessionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionKeyWithDefaults

`func NewSessionKeyWithDefaults() *SessionKey`

NewSessionKeyWithDefaults instantiates a new SessionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *SessionKey) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *SessionKey) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *SessionKey) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetId

`func (o *SessionKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionKey) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *SessionKey) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SessionKey) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SessionKey) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *SessionKey) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SessionKey) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SessionKey) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUserkey

`func (o *SessionKey) GetUserkey() BriefUserKey`

GetUserkey returns the Userkey field if non-nil, zero value otherwise.

### GetUserkeyOk

`func (o *SessionKey) GetUserkeyOk() (*BriefUserKey, bool)`

GetUserkeyOk returns a tuple with the Userkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserkey

`func (o *SessionKey) SetUserkey(v BriefUserKey)`

SetUserkey sets Userkey field to given value.


### GetSessionKey

`func (o *SessionKey) GetSessionKey() string`

GetSessionKey returns the SessionKey field if non-nil, zero value otherwise.

### GetSessionKeyOk

`func (o *SessionKey) GetSessionKeyOk() (*string, bool)`

GetSessionKeyOk returns a tuple with the SessionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionKey

`func (o *SessionKey) SetSessionKey(v string)`

SetSessionKey sets SessionKey field to given value.


### GetCreated

`func (o *SessionKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SessionKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SessionKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

