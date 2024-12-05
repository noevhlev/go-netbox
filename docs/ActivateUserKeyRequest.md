# ActivateUserKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** |  | 
**UserKeys** | **[]interface{}** |  | 

## Methods

### NewActivateUserKeyRequest

`func NewActivateUserKeyRequest(privateKey string, userKeys []interface{}, ) *ActivateUserKeyRequest`

NewActivateUserKeyRequest instantiates a new ActivateUserKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateUserKeyRequestWithDefaults

`func NewActivateUserKeyRequestWithDefaults() *ActivateUserKeyRequest`

NewActivateUserKeyRequestWithDefaults instantiates a new ActivateUserKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *ActivateUserKeyRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ActivateUserKeyRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ActivateUserKeyRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetUserKeys

`func (o *ActivateUserKeyRequest) GetUserKeys() []interface{}`

GetUserKeys returns the UserKeys field if non-nil, zero value otherwise.

### GetUserKeysOk

`func (o *ActivateUserKeyRequest) GetUserKeysOk() (*[]interface{}, bool)`

GetUserKeysOk returns a tuple with the UserKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeys

`func (o *ActivateUserKeyRequest) SetUserKeys(v []interface{})`

SetUserKeys sets UserKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


