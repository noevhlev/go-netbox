# SessionKeyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** |  | 
**PreserveKey** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSessionKeyCreateRequest

`func NewSessionKeyCreateRequest(privateKey string, ) *SessionKeyCreateRequest`

NewSessionKeyCreateRequest instantiates a new SessionKeyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionKeyCreateRequestWithDefaults

`func NewSessionKeyCreateRequestWithDefaults() *SessionKeyCreateRequest`

NewSessionKeyCreateRequestWithDefaults instantiates a new SessionKeyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *SessionKeyCreateRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SessionKeyCreateRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SessionKeyCreateRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPreserveKey

`func (o *SessionKeyCreateRequest) GetPreserveKey() bool`

GetPreserveKey returns the PreserveKey field if non-nil, zero value otherwise.

### GetPreserveKeyOk

`func (o *SessionKeyCreateRequest) GetPreserveKeyOk() (*bool, bool)`

GetPreserveKeyOk returns a tuple with the PreserveKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveKey

`func (o *SessionKeyCreateRequest) SetPreserveKey(v bool)`

SetPreserveKey sets PreserveKey field to given value.

### HasPreserveKey

`func (o *SessionKeyCreateRequest) HasPreserveKey() bool`

HasPreserveKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


