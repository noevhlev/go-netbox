# AssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Can be used to quickly identify a particular asset | [optional] 
**AssetTag** | Pointer to **NullableString** | Identifier assigned by owner | [optional] 
**Serial** | Pointer to **NullableString** | Identifier assigned by manufacturer | [optional] 
**Status** | [**AssetStatus**](AssetStatus.md) |  | 
**DeviceType** | Pointer to [**NullableNestedDeviceTypeRequest**](NestedDeviceTypeRequest.md) |  | [optional] 
**Device** | Pointer to [**NullableNestedDeviceRequest**](NestedDeviceRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableNestedModuleTypeRequest**](NestedModuleTypeRequest.md) |  | [optional] 
**InventoryitemType** | Pointer to [**NullableNestedInventoryItemTypeRequest**](NestedInventoryItemTypeRequest.md) |  | [optional] 
**Inventoryitem** | Pointer to [**NullableNestedInventoryItemRequest**](NestedInventoryItemRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableNestedTenantRequest**](NestedTenantRequest.md) |  | [optional] 
**Contact** | Pointer to [**NullableNestedContactRequest**](NestedContactRequest.md) |  | [optional] 
**StorageLocation** | Pointer to [**NullableNestedLocationRequest**](NestedLocationRequest.md) |  | [optional] 
**Owner** | Pointer to [**NullableNestedTenantRequest**](NestedTenantRequest.md) |  | [optional] 
**Delivery** | Pointer to [**NullableNestedDeliveryRequest**](NestedDeliveryRequest.md) |  | [optional] 
**Purchase** | Pointer to [**NullableNestedPurchaseRequest**](NestedPurchaseRequest.md) |  | [optional] 
**WarrantyStart** | Pointer to **NullableString** | First date warranty for this asset is valid | [optional] 
**WarrantyEnd** | Pointer to **NullableString** | Last date warranty for this asset is valid | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAssetRequest

`func NewAssetRequest(status AssetStatus, ) *AssetRequest`

NewAssetRequest instantiates a new AssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetRequestWithDefaults

`func NewAssetRequestWithDefaults() *AssetRequest`

NewAssetRequestWithDefaults instantiates a new AssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssetTag

`func (o *AssetRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *AssetRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *AssetRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *AssetRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *AssetRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *AssetRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSerial

`func (o *AssetRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AssetRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AssetRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AssetRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *AssetRequest) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *AssetRequest) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetStatus

`func (o *AssetRequest) GetStatus() AssetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetRequest) GetStatusOk() (*AssetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetRequest) SetStatus(v AssetStatus)`

SetStatus sets Status field to given value.


### GetDeviceType

`func (o *AssetRequest) GetDeviceType() NestedDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AssetRequest) GetDeviceTypeOk() (*NestedDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AssetRequest) SetDeviceType(v NestedDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *AssetRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *AssetRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *AssetRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetDevice

`func (o *AssetRequest) GetDevice() NestedDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AssetRequest) GetDeviceOk() (*NestedDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AssetRequest) SetDevice(v NestedDeviceRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AssetRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *AssetRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *AssetRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetModuleType

`func (o *AssetRequest) GetModuleType() NestedModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *AssetRequest) GetModuleTypeOk() (*NestedModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *AssetRequest) SetModuleType(v NestedModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *AssetRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *AssetRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *AssetRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetInventoryitemType

`func (o *AssetRequest) GetInventoryitemType() NestedInventoryItemTypeRequest`

GetInventoryitemType returns the InventoryitemType field if non-nil, zero value otherwise.

### GetInventoryitemTypeOk

`func (o *AssetRequest) GetInventoryitemTypeOk() (*NestedInventoryItemTypeRequest, bool)`

GetInventoryitemTypeOk returns a tuple with the InventoryitemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemType

`func (o *AssetRequest) SetInventoryitemType(v NestedInventoryItemTypeRequest)`

SetInventoryitemType sets InventoryitemType field to given value.

### HasInventoryitemType

`func (o *AssetRequest) HasInventoryitemType() bool`

HasInventoryitemType returns a boolean if a field has been set.

### SetInventoryitemTypeNil

`func (o *AssetRequest) SetInventoryitemTypeNil(b bool)`

 SetInventoryitemTypeNil sets the value for InventoryitemType to be an explicit nil

### UnsetInventoryitemType
`func (o *AssetRequest) UnsetInventoryitemType()`

UnsetInventoryitemType ensures that no value is present for InventoryitemType, not even an explicit nil
### GetInventoryitem

`func (o *AssetRequest) GetInventoryitem() NestedInventoryItemRequest`

GetInventoryitem returns the Inventoryitem field if non-nil, zero value otherwise.

### GetInventoryitemOk

`func (o *AssetRequest) GetInventoryitemOk() (*NestedInventoryItemRequest, bool)`

GetInventoryitemOk returns a tuple with the Inventoryitem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitem

`func (o *AssetRequest) SetInventoryitem(v NestedInventoryItemRequest)`

SetInventoryitem sets Inventoryitem field to given value.

### HasInventoryitem

`func (o *AssetRequest) HasInventoryitem() bool`

HasInventoryitem returns a boolean if a field has been set.

### SetInventoryitemNil

`func (o *AssetRequest) SetInventoryitemNil(b bool)`

 SetInventoryitemNil sets the value for Inventoryitem to be an explicit nil

### UnsetInventoryitem
`func (o *AssetRequest) UnsetInventoryitem()`

UnsetInventoryitem ensures that no value is present for Inventoryitem, not even an explicit nil
### GetTenant

`func (o *AssetRequest) GetTenant() NestedTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AssetRequest) GetTenantOk() (*NestedTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AssetRequest) SetTenant(v NestedTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AssetRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *AssetRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *AssetRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetContact

`func (o *AssetRequest) GetContact() NestedContactRequest`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AssetRequest) GetContactOk() (*NestedContactRequest, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AssetRequest) SetContact(v NestedContactRequest)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AssetRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *AssetRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *AssetRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetStorageLocation

`func (o *AssetRequest) GetStorageLocation() NestedLocationRequest`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *AssetRequest) GetStorageLocationOk() (*NestedLocationRequest, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *AssetRequest) SetStorageLocation(v NestedLocationRequest)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *AssetRequest) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### SetStorageLocationNil

`func (o *AssetRequest) SetStorageLocationNil(b bool)`

 SetStorageLocationNil sets the value for StorageLocation to be an explicit nil

### UnsetStorageLocation
`func (o *AssetRequest) UnsetStorageLocation()`

UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil
### GetOwner

`func (o *AssetRequest) GetOwner() NestedTenantRequest`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AssetRequest) GetOwnerOk() (*NestedTenantRequest, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AssetRequest) SetOwner(v NestedTenantRequest)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AssetRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *AssetRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *AssetRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetDelivery

`func (o *AssetRequest) GetDelivery() NestedDeliveryRequest`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *AssetRequest) GetDeliveryOk() (*NestedDeliveryRequest, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *AssetRequest) SetDelivery(v NestedDeliveryRequest)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *AssetRequest) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### SetDeliveryNil

`func (o *AssetRequest) SetDeliveryNil(b bool)`

 SetDeliveryNil sets the value for Delivery to be an explicit nil

### UnsetDelivery
`func (o *AssetRequest) UnsetDelivery()`

UnsetDelivery ensures that no value is present for Delivery, not even an explicit nil
### GetPurchase

`func (o *AssetRequest) GetPurchase() NestedPurchaseRequest`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *AssetRequest) GetPurchaseOk() (*NestedPurchaseRequest, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *AssetRequest) SetPurchase(v NestedPurchaseRequest)`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *AssetRequest) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### SetPurchaseNil

`func (o *AssetRequest) SetPurchaseNil(b bool)`

 SetPurchaseNil sets the value for Purchase to be an explicit nil

### UnsetPurchase
`func (o *AssetRequest) UnsetPurchase()`

UnsetPurchase ensures that no value is present for Purchase, not even an explicit nil
### GetWarrantyStart

`func (o *AssetRequest) GetWarrantyStart() string`

GetWarrantyStart returns the WarrantyStart field if non-nil, zero value otherwise.

### GetWarrantyStartOk

`func (o *AssetRequest) GetWarrantyStartOk() (*string, bool)`

GetWarrantyStartOk returns a tuple with the WarrantyStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyStart

`func (o *AssetRequest) SetWarrantyStart(v string)`

SetWarrantyStart sets WarrantyStart field to given value.

### HasWarrantyStart

`func (o *AssetRequest) HasWarrantyStart() bool`

HasWarrantyStart returns a boolean if a field has been set.

### SetWarrantyStartNil

`func (o *AssetRequest) SetWarrantyStartNil(b bool)`

 SetWarrantyStartNil sets the value for WarrantyStart to be an explicit nil

### UnsetWarrantyStart
`func (o *AssetRequest) UnsetWarrantyStart()`

UnsetWarrantyStart ensures that no value is present for WarrantyStart, not even an explicit nil
### GetWarrantyEnd

`func (o *AssetRequest) GetWarrantyEnd() string`

GetWarrantyEnd returns the WarrantyEnd field if non-nil, zero value otherwise.

### GetWarrantyEndOk

`func (o *AssetRequest) GetWarrantyEndOk() (*string, bool)`

GetWarrantyEndOk returns a tuple with the WarrantyEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEnd

`func (o *AssetRequest) SetWarrantyEnd(v string)`

SetWarrantyEnd sets WarrantyEnd field to given value.

### HasWarrantyEnd

`func (o *AssetRequest) HasWarrantyEnd() bool`

HasWarrantyEnd returns a boolean if a field has been set.

### SetWarrantyEndNil

`func (o *AssetRequest) SetWarrantyEndNil(b bool)`

 SetWarrantyEndNil sets the value for WarrantyEnd to be an explicit nil

### UnsetWarrantyEnd
`func (o *AssetRequest) UnsetWarrantyEnd()`

UnsetWarrantyEnd ensures that no value is present for WarrantyEnd, not even an explicit nil
### GetComments

`func (o *AssetRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AssetRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AssetRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AssetRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *AssetRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *AssetRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AssetRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AssetRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AssetRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


