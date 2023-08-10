# RestorePreflightParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupLocations** | **[]string** | List of backup locations | 
**BackupUUID** | Pointer to **string** | The backup of which the restore is being attempted | [optional] 
**StorageConfigUUID** | **string** | Storage config UUID | 
**UniverseUUID** | **string** | Target universe UUID | 

## Methods

### NewRestorePreflightParams

`func NewRestorePreflightParams(backupLocations []string, storageConfigUUID string, universeUUID string, ) *RestorePreflightParams`

NewRestorePreflightParams instantiates a new RestorePreflightParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorePreflightParamsWithDefaults

`func NewRestorePreflightParamsWithDefaults() *RestorePreflightParams`

NewRestorePreflightParamsWithDefaults instantiates a new RestorePreflightParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupLocations

`func (o *RestorePreflightParams) GetBackupLocations() []string`

GetBackupLocations returns the BackupLocations field if non-nil, zero value otherwise.

### GetBackupLocationsOk

`func (o *RestorePreflightParams) GetBackupLocationsOk() (*[]string, bool)`

GetBackupLocationsOk returns a tuple with the BackupLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocations

`func (o *RestorePreflightParams) SetBackupLocations(v []string)`

SetBackupLocations sets BackupLocations field to given value.


### GetBackupUUID

`func (o *RestorePreflightParams) GetBackupUUID() string`

GetBackupUUID returns the BackupUUID field if non-nil, zero value otherwise.

### GetBackupUUIDOk

`func (o *RestorePreflightParams) GetBackupUUIDOk() (*string, bool)`

GetBackupUUIDOk returns a tuple with the BackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUID

`func (o *RestorePreflightParams) SetBackupUUID(v string)`

SetBackupUUID sets BackupUUID field to given value.

### HasBackupUUID

`func (o *RestorePreflightParams) HasBackupUUID() bool`

HasBackupUUID returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *RestorePreflightParams) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *RestorePreflightParams) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *RestorePreflightParams) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetUniverseUUID

`func (o *RestorePreflightParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *RestorePreflightParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *RestorePreflightParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


