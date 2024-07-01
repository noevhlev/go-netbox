# WritableAssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Can be used to quickly identify a particular asset | [optional] 
**AssetTag** | Pointer to **NullableString** | Identifier assigned by owner | [optional] 
**Serial** | Pointer to **NullableString** | Identifier assigned by manufacturer | [optional] 
**Status** | [**AssetStatus**](AssetStatus.md) |  | 
**DeviceType** | Pointer to **NullableInt32** |  | [optional] 
**Device** | Pointer to **NullableInt32** |  | [optional] 
**ModuleType** | Pointer to **NullableInt32** |  | [optional] 
**Module** | Pointer to **NullableInt32** |  | [optional] 
**InventoryitemType** | Pointer to **NullableInt32** |  | [optional] 
**Inventoryitem** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to **NullableInt32** | Tenant using this asset | [optional] 
**Contact** | Pointer to **NullableInt32** | Contact using this asset | [optional] 
**StorageLocation** | Pointer to **NullableInt32** | Where is this asset stored when not in use | [optional] 
**Owner** | Pointer to **NullableInt32** | Who owns this asset | [optional] 
**Delivery** | Pointer to **NullableInt32** | Delivery this asset was part of | [optional] 
**Purchase** | Pointer to **NullableInt32** | Purchase through which this asset was purchased | [optional] 
**WarrantyStart** | Pointer to **NullableString** | First date warranty for this asset is valid | [optional] 
**WarrantyEnd** | Pointer to **NullableString** | Last date warranty for this asset is valid | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableAssetRequest

`func NewWritableAssetRequest(status AssetStatus, ) *WritableAssetRequest`

NewWritableAssetRequest instantiates a new WritableAssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableAssetRequestWithDefaults

`func NewWritableAssetRequestWithDefaults() *WritableAssetRequest`

NewWritableAssetRequestWithDefaults instantiates a new WritableAssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableAssetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableAssetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableAssetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WritableAssetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssetTag

`func (o *WritableAssetRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *WritableAssetRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *WritableAssetRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *WritableAssetRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *WritableAssetRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *WritableAssetRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSerial

`func (o *WritableAssetRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *WritableAssetRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *WritableAssetRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *WritableAssetRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *WritableAssetRequest) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *WritableAssetRequest) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetStatus

`func (o *WritableAssetRequest) GetStatus() AssetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableAssetRequest) GetStatusOk() (*AssetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableAssetRequest) SetStatus(v AssetStatus)`

SetStatus sets Status field to given value.


### GetDeviceType

`func (o *WritableAssetRequest) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WritableAssetRequest) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WritableAssetRequest) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WritableAssetRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *WritableAssetRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *WritableAssetRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetDevice

`func (o *WritableAssetRequest) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableAssetRequest) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableAssetRequest) SetDevice(v int32)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *WritableAssetRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *WritableAssetRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *WritableAssetRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetModuleType

`func (o *WritableAssetRequest) GetModuleType() int32`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *WritableAssetRequest) GetModuleTypeOk() (*int32, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *WritableAssetRequest) SetModuleType(v int32)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *WritableAssetRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *WritableAssetRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *WritableAssetRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetModule

`func (o *WritableAssetRequest) GetModule() int32`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WritableAssetRequest) GetModuleOk() (*int32, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WritableAssetRequest) SetModule(v int32)`

SetModule sets Module field to given value.

### HasModule

`func (o *WritableAssetRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *WritableAssetRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *WritableAssetRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetInventoryitemType

`func (o *WritableAssetRequest) GetInventoryitemType() int32`

GetInventoryitemType returns the InventoryitemType field if non-nil, zero value otherwise.

### GetInventoryitemTypeOk

`func (o *WritableAssetRequest) GetInventoryitemTypeOk() (*int32, bool)`

GetInventoryitemTypeOk returns a tuple with the InventoryitemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemType

`func (o *WritableAssetRequest) SetInventoryitemType(v int32)`

SetInventoryitemType sets InventoryitemType field to given value.

### HasInventoryitemType

`func (o *WritableAssetRequest) HasInventoryitemType() bool`

HasInventoryitemType returns a boolean if a field has been set.

### SetInventoryitemTypeNil

`func (o *WritableAssetRequest) SetInventoryitemTypeNil(b bool)`

 SetInventoryitemTypeNil sets the value for InventoryitemType to be an explicit nil

### UnsetInventoryitemType
`func (o *WritableAssetRequest) UnsetInventoryitemType()`

UnsetInventoryitemType ensures that no value is present for InventoryitemType, not even an explicit nil
### GetInventoryitem

`func (o *WritableAssetRequest) GetInventoryitem() int32`

GetInventoryitem returns the Inventoryitem field if non-nil, zero value otherwise.

### GetInventoryitemOk

`func (o *WritableAssetRequest) GetInventoryitemOk() (*int32, bool)`

GetInventoryitemOk returns a tuple with the Inventoryitem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitem

`func (o *WritableAssetRequest) SetInventoryitem(v int32)`

SetInventoryitem sets Inventoryitem field to given value.

### HasInventoryitem

`func (o *WritableAssetRequest) HasInventoryitem() bool`

HasInventoryitem returns a boolean if a field has been set.

### SetInventoryitemNil

`func (o *WritableAssetRequest) SetInventoryitemNil(b bool)`

 SetInventoryitemNil sets the value for Inventoryitem to be an explicit nil

### UnsetInventoryitem
`func (o *WritableAssetRequest) UnsetInventoryitem()`

UnsetInventoryitem ensures that no value is present for Inventoryitem, not even an explicit nil
### GetTenant

`func (o *WritableAssetRequest) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableAssetRequest) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableAssetRequest) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableAssetRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableAssetRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableAssetRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetContact

`func (o *WritableAssetRequest) GetContact() int32`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *WritableAssetRequest) GetContactOk() (*int32, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *WritableAssetRequest) SetContact(v int32)`

SetContact sets Contact field to given value.

### HasContact

`func (o *WritableAssetRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *WritableAssetRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *WritableAssetRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetStorageLocation

`func (o *WritableAssetRequest) GetStorageLocation() int32`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *WritableAssetRequest) GetStorageLocationOk() (*int32, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *WritableAssetRequest) SetStorageLocation(v int32)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *WritableAssetRequest) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### SetStorageLocationNil

`func (o *WritableAssetRequest) SetStorageLocationNil(b bool)`

 SetStorageLocationNil sets the value for StorageLocation to be an explicit nil

### UnsetStorageLocation
`func (o *WritableAssetRequest) UnsetStorageLocation()`

UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil
### GetOwner

`func (o *WritableAssetRequest) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WritableAssetRequest) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WritableAssetRequest) SetOwner(v int32)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WritableAssetRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WritableAssetRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WritableAssetRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetDelivery

`func (o *WritableAssetRequest) GetDelivery() int32`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *WritableAssetRequest) GetDeliveryOk() (*int32, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *WritableAssetRequest) SetDelivery(v int32)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *WritableAssetRequest) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### SetDeliveryNil

`func (o *WritableAssetRequest) SetDeliveryNil(b bool)`

 SetDeliveryNil sets the value for Delivery to be an explicit nil

### UnsetDelivery
`func (o *WritableAssetRequest) UnsetDelivery()`

UnsetDelivery ensures that no value is present for Delivery, not even an explicit nil
### GetPurchase

`func (o *WritableAssetRequest) GetPurchase() int32`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *WritableAssetRequest) GetPurchaseOk() (*int32, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *WritableAssetRequest) SetPurchase(v int32)`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *WritableAssetRequest) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### SetPurchaseNil

`func (o *WritableAssetRequest) SetPurchaseNil(b bool)`

 SetPurchaseNil sets the value for Purchase to be an explicit nil

### UnsetPurchase
`func (o *WritableAssetRequest) UnsetPurchase()`

UnsetPurchase ensures that no value is present for Purchase, not even an explicit nil
### GetWarrantyStart

`func (o *WritableAssetRequest) GetWarrantyStart() string`

GetWarrantyStart returns the WarrantyStart field if non-nil, zero value otherwise.

### GetWarrantyStartOk

`func (o *WritableAssetRequest) GetWarrantyStartOk() (*string, bool)`

GetWarrantyStartOk returns a tuple with the WarrantyStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyStart

`func (o *WritableAssetRequest) SetWarrantyStart(v string)`

SetWarrantyStart sets WarrantyStart field to given value.

### HasWarrantyStart

`func (o *WritableAssetRequest) HasWarrantyStart() bool`

HasWarrantyStart returns a boolean if a field has been set.

### SetWarrantyStartNil

`func (o *WritableAssetRequest) SetWarrantyStartNil(b bool)`

 SetWarrantyStartNil sets the value for WarrantyStart to be an explicit nil

### UnsetWarrantyStart
`func (o *WritableAssetRequest) UnsetWarrantyStart()`

UnsetWarrantyStart ensures that no value is present for WarrantyStart, not even an explicit nil
### GetWarrantyEnd

`func (o *WritableAssetRequest) GetWarrantyEnd() string`

GetWarrantyEnd returns the WarrantyEnd field if non-nil, zero value otherwise.

### GetWarrantyEndOk

`func (o *WritableAssetRequest) GetWarrantyEndOk() (*string, bool)`

GetWarrantyEndOk returns a tuple with the WarrantyEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEnd

`func (o *WritableAssetRequest) SetWarrantyEnd(v string)`

SetWarrantyEnd sets WarrantyEnd field to given value.

### HasWarrantyEnd

`func (o *WritableAssetRequest) HasWarrantyEnd() bool`

HasWarrantyEnd returns a boolean if a field has been set.

### SetWarrantyEndNil

`func (o *WritableAssetRequest) SetWarrantyEndNil(b bool)`

 SetWarrantyEndNil sets the value for WarrantyEnd to be an explicit nil

### UnsetWarrantyEnd
`func (o *WritableAssetRequest) UnsetWarrantyEnd()`

UnsetWarrantyEnd ensures that no value is present for WarrantyEnd, not even an explicit nil
### GetComments

`func (o *WritableAssetRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableAssetRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableAssetRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableAssetRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableAssetRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableAssetRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableAssetRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableAssetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableAssetRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableAssetRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableAssetRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableAssetRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


