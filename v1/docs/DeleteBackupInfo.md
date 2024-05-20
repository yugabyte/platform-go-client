# DeleteBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupUUID** | **string** | backup UUID | 
**StorageConfigUUID** | Pointer to **string** | storage config UUID | [optional] 

## Methods

### NewDeleteBackupInfo

`func NewDeleteBackupInfo(backupUUID string, ) *DeleteBackupInfo`

NewDeleteBackupInfo instantiates a new DeleteBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBackupInfoWithDefaults

`func NewDeleteBackupInfoWithDefaults() *DeleteBackupInfo`

NewDeleteBackupInfoWithDefaults instantiates a new DeleteBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupUUID

`func (o *DeleteBackupInfo) GetBackupUUID() string`

GetBackupUUID returns the BackupUUID field if non-nil, zero value otherwise.

### GetBackupUUIDOk

`func (o *DeleteBackupInfo) GetBackupUUIDOk() (*string, bool)`

GetBackupUUIDOk returns a tuple with the BackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUID

`func (o *DeleteBackupInfo) SetBackupUUID(v string)`

SetBackupUUID sets BackupUUID field to given value.


### GetStorageConfigUUID

`func (o *DeleteBackupInfo) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *DeleteBackupInfo) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *DeleteBackupInfo) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.

### HasStorageConfigUUID

`func (o *DeleteBackupInfo) HasStorageConfigUUID() bool`

HasStorageConfigUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


