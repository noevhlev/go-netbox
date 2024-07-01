# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | Pointer to **string** | Can be used to quickly identify a particular asset | [optional] 
**AssetTag** | Pointer to **NullableString** | Identifier assigned by owner | [optional] 
**Serial** | Pointer to **NullableString** | Identifier assigned by manufacturer | [optional] 
**Status** | [**AssetStatus**](AssetStatus.md) |  | 
**Kind** | **string** |  | [readonly] 
**DeviceType** | Pointer to [**NullableNestedDeviceType**](NestedDeviceType.md) |  | [optional] 
**Device** | Pointer to [**NullableNestedDevice**](NestedDevice.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableNestedModuleType**](NestedModuleType.md) |  | [optional] 
**Module** | Pointer to [**NullableNestedModule**](NestedModule.md) |  | [optional] 
**InventoryitemType** | Pointer to [**NullableNestedInventoryItemType**](NestedInventoryItemType.md) |  | [optional] 
**Inventoryitem** | Pointer to [**NullableNestedInventoryItem**](NestedInventoryItem.md) |  | [optional] 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**Contact** | Pointer to [**NullableNestedContact**](NestedContact.md) |  | [optional] 
**StorageLocation** | Pointer to [**NullableNestedLocation**](NestedLocation.md) |  | [optional] 
**Owner** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**Delivery** | Pointer to [**NullableNestedDelivery**](NestedDelivery.md) |  | [optional] 
**Purchase** | Pointer to [**NullableNestedPurchase**](NestedPurchase.md) |  | [optional] 
**WarrantyStart** | Pointer to **NullableString** | First date warranty for this asset is valid | [optional] 
**WarrantyEnd** | Pointer to **NullableString** | Last date warranty for this asset is valid | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewAsset

`func NewAsset(id int32, url string, display string, status AssetStatus, kind string, created NullableTime, lastUpdated NullableTime, ) *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Asset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Asset) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Asset) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Asset) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *Asset) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Asset) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Asset) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Asset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssetTag

`func (o *Asset) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *Asset) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *Asset) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *Asset) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *Asset) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *Asset) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSerial

`func (o *Asset) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Asset) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Asset) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Asset) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *Asset) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *Asset) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetStatus

`func (o *Asset) GetStatus() AssetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Asset) GetStatusOk() (*AssetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Asset) SetStatus(v AssetStatus)`

SetStatus sets Status field to given value.


### GetKind

`func (o *Asset) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Asset) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Asset) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDeviceType

`func (o *Asset) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Asset) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Asset) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *Asset) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *Asset) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *Asset) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetDevice

`func (o *Asset) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Asset) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Asset) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Asset) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *Asset) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *Asset) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetModuleType

`func (o *Asset) GetModuleType() NestedModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *Asset) GetModuleTypeOk() (*NestedModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *Asset) SetModuleType(v NestedModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *Asset) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *Asset) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *Asset) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetModule

`func (o *Asset) GetModule() NestedModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *Asset) GetModuleOk() (*NestedModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *Asset) SetModule(v NestedModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *Asset) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *Asset) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *Asset) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetInventoryitemType

`func (o *Asset) GetInventoryitemType() NestedInventoryItemType`

GetInventoryitemType returns the InventoryitemType field if non-nil, zero value otherwise.

### GetInventoryitemTypeOk

`func (o *Asset) GetInventoryitemTypeOk() (*NestedInventoryItemType, bool)`

GetInventoryitemTypeOk returns a tuple with the InventoryitemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemType

`func (o *Asset) SetInventoryitemType(v NestedInventoryItemType)`

SetInventoryitemType sets InventoryitemType field to given value.

### HasInventoryitemType

`func (o *Asset) HasInventoryitemType() bool`

HasInventoryitemType returns a boolean if a field has been set.

### SetInventoryitemTypeNil

`func (o *Asset) SetInventoryitemTypeNil(b bool)`

 SetInventoryitemTypeNil sets the value for InventoryitemType to be an explicit nil

### UnsetInventoryitemType
`func (o *Asset) UnsetInventoryitemType()`

UnsetInventoryitemType ensures that no value is present for InventoryitemType, not even an explicit nil
### GetInventoryitem

`func (o *Asset) GetInventoryitem() NestedInventoryItem`

GetInventoryitem returns the Inventoryitem field if non-nil, zero value otherwise.

### GetInventoryitemOk

`func (o *Asset) GetInventoryitemOk() (*NestedInventoryItem, bool)`

GetInventoryitemOk returns a tuple with the Inventoryitem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitem

`func (o *Asset) SetInventoryitem(v NestedInventoryItem)`

SetInventoryitem sets Inventoryitem field to given value.

### HasInventoryitem

`func (o *Asset) HasInventoryitem() bool`

HasInventoryitem returns a boolean if a field has been set.

### SetInventoryitemNil

`func (o *Asset) SetInventoryitemNil(b bool)`

 SetInventoryitemNil sets the value for Inventoryitem to be an explicit nil

### UnsetInventoryitem
`func (o *Asset) UnsetInventoryitem()`

UnsetInventoryitem ensures that no value is present for Inventoryitem, not even an explicit nil
### GetTenant

`func (o *Asset) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Asset) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Asset) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Asset) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Asset) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Asset) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetContact

`func (o *Asset) GetContact() NestedContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Asset) GetContactOk() (*NestedContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Asset) SetContact(v NestedContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Asset) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Asset) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Asset) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetStorageLocation

`func (o *Asset) GetStorageLocation() NestedLocation`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *Asset) GetStorageLocationOk() (*NestedLocation, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *Asset) SetStorageLocation(v NestedLocation)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *Asset) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### SetStorageLocationNil

`func (o *Asset) SetStorageLocationNil(b bool)`

 SetStorageLocationNil sets the value for StorageLocation to be an explicit nil

### UnsetStorageLocation
`func (o *Asset) UnsetStorageLocation()`

UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil
### GetOwner

`func (o *Asset) GetOwner() NestedTenant`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Asset) GetOwnerOk() (*NestedTenant, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Asset) SetOwner(v NestedTenant)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Asset) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Asset) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Asset) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetDelivery

`func (o *Asset) GetDelivery() NestedDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *Asset) GetDeliveryOk() (*NestedDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *Asset) SetDelivery(v NestedDelivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *Asset) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### SetDeliveryNil

`func (o *Asset) SetDeliveryNil(b bool)`

 SetDeliveryNil sets the value for Delivery to be an explicit nil

### UnsetDelivery
`func (o *Asset) UnsetDelivery()`

UnsetDelivery ensures that no value is present for Delivery, not even an explicit nil
### GetPurchase

`func (o *Asset) GetPurchase() NestedPurchase`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *Asset) GetPurchaseOk() (*NestedPurchase, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *Asset) SetPurchase(v NestedPurchase)`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *Asset) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### SetPurchaseNil

`func (o *Asset) SetPurchaseNil(b bool)`

 SetPurchaseNil sets the value for Purchase to be an explicit nil

### UnsetPurchase
`func (o *Asset) UnsetPurchase()`

UnsetPurchase ensures that no value is present for Purchase, not even an explicit nil
### GetWarrantyStart

`func (o *Asset) GetWarrantyStart() string`

GetWarrantyStart returns the WarrantyStart field if non-nil, zero value otherwise.

### GetWarrantyStartOk

`func (o *Asset) GetWarrantyStartOk() (*string, bool)`

GetWarrantyStartOk returns a tuple with the WarrantyStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyStart

`func (o *Asset) SetWarrantyStart(v string)`

SetWarrantyStart sets WarrantyStart field to given value.

### HasWarrantyStart

`func (o *Asset) HasWarrantyStart() bool`

HasWarrantyStart returns a boolean if a field has been set.

### SetWarrantyStartNil

`func (o *Asset) SetWarrantyStartNil(b bool)`

 SetWarrantyStartNil sets the value for WarrantyStart to be an explicit nil

### UnsetWarrantyStart
`func (o *Asset) UnsetWarrantyStart()`

UnsetWarrantyStart ensures that no value is present for WarrantyStart, not even an explicit nil
### GetWarrantyEnd

`func (o *Asset) GetWarrantyEnd() string`

GetWarrantyEnd returns the WarrantyEnd field if non-nil, zero value otherwise.

### GetWarrantyEndOk

`func (o *Asset) GetWarrantyEndOk() (*string, bool)`

GetWarrantyEndOk returns a tuple with the WarrantyEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEnd

`func (o *Asset) SetWarrantyEnd(v string)`

SetWarrantyEnd sets WarrantyEnd field to given value.

### HasWarrantyEnd

`func (o *Asset) HasWarrantyEnd() bool`

HasWarrantyEnd returns a boolean if a field has been set.

### SetWarrantyEndNil

`func (o *Asset) SetWarrantyEndNil(b bool)`

 SetWarrantyEndNil sets the value for WarrantyEnd to be an explicit nil

### UnsetWarrantyEnd
`func (o *Asset) UnsetWarrantyEnd()`

UnsetWarrantyEnd ensures that no value is present for WarrantyEnd, not even an explicit nil
### GetComments

`func (o *Asset) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Asset) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Asset) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Asset) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Asset) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Asset) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Asset) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Asset) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Asset) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Asset) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Asset) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Asset) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Asset) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Asset) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Asset) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Asset) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Asset) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Asset) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Asset) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Asset) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Asset) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Asset) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


