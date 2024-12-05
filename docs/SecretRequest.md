# SecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **int32** |  | 
**Role** | [**BriefSecretRoleRequest**](BriefSecretRoleRequest.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Plaintext** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSecretRequest

`func NewSecretRequest(assignedObjectType string, assignedObjectId int32, role BriefSecretRoleRequest, plaintext string, ) *SecretRequest`

NewSecretRequest instantiates a new SecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretRequestWithDefaults

`func NewSecretRequestWithDefaults() *SecretRequest`

NewSecretRequestWithDefaults instantiates a new SecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedObjectType

`func (o *SecretRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *SecretRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *SecretRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *SecretRequest) GetAssignedObjectId() int32`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *SecretRequest) GetAssignedObjectIdOk() (*int32, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *SecretRequest) SetAssignedObjectId(v int32)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetRole

`func (o *SecretRequest) GetRole() BriefSecretRoleRequest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SecretRequest) GetRoleOk() (*BriefSecretRoleRequest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SecretRequest) SetRole(v BriefSecretRoleRequest)`

SetRole sets Role field to given value.


### GetName

`func (o *SecretRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaintext

`func (o *SecretRequest) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *SecretRequest) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *SecretRequest) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### GetDescription

`func (o *SecretRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecretRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecretRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecretRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *SecretRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *SecretRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *SecretRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *SecretRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *SecretRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecretRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecretRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecretRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *SecretRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SecretRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SecretRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SecretRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


