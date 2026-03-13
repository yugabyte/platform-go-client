# BackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupType** | **string** |  | 
**ExpiryTimeUnit** | **string** |  | 
**FullBackup** | **bool** |  | 
**KeyspaceList** | [**[]KeyspaceTablesList**](KeyspaceTablesList.md) |  | 
**Parallelism** | **int64** |  | 
**PointInTimeRestoreEnabled** | **bool** |  | 
**StorageConfigUUID** | **string** |  | 
**TimeBeforeDelete** | **int64** |  | 
**UniverseUUID** | **string** |  | 
**UseRoles** | **bool** |  | 
**UseTablespaces** | **bool** |  | 

## Methods

### NewBackupInfo

`func NewBackupInfo(backupType string, expiryTimeUnit string, fullBackup bool, keyspaceList []KeyspaceTablesList, parallelism int64, pointInTimeRestoreEnabled bool, storageConfigUUID string, timeBeforeDelete int64, universeUUID string, useRoles bool, useTablespaces bool, ) *BackupInfo`

NewBackupInfo instantiates a new BackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupInfoWithDefaults

`func NewBackupInfoWithDefaults() *BackupInfo`

NewBackupInfoWithDefaults instantiates a new BackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupType

`func (o *BackupInfo) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupInfo) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupInfo) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.


### GetExpiryTimeUnit

`func (o *BackupInfo) GetExpiryTimeUnit() string`

GetExpiryTimeUnit returns the ExpiryTimeUnit field if non-nil, zero value otherwise.

### GetExpiryTimeUnitOk

`func (o *BackupInfo) GetExpiryTimeUnitOk() (*string, bool)`

GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUnit

`func (o *BackupInfo) SetExpiryTimeUnit(v string)`

SetExpiryTimeUnit sets ExpiryTimeUnit field to given value.


### GetFullBackup

`func (o *BackupInfo) GetFullBackup() bool`

GetFullBackup returns the FullBackup field if non-nil, zero value otherwise.

### GetFullBackupOk

`func (o *BackupInfo) GetFullBackupOk() (*bool, bool)`

GetFullBackupOk returns a tuple with the FullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackup

`func (o *BackupInfo) SetFullBackup(v bool)`

SetFullBackup sets FullBackup field to given value.


### GetKeyspaceList

`func (o *BackupInfo) GetKeyspaceList() []KeyspaceTablesList`

GetKeyspaceList returns the KeyspaceList field if non-nil, zero value otherwise.

### GetKeyspaceListOk

`func (o *BackupInfo) GetKeyspaceListOk() (*[]KeyspaceTablesList, bool)`

GetKeyspaceListOk returns a tuple with the KeyspaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceList

`func (o *BackupInfo) SetKeyspaceList(v []KeyspaceTablesList)`

SetKeyspaceList sets KeyspaceList field to given value.


### GetParallelism

`func (o *BackupInfo) GetParallelism() int64`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *BackupInfo) GetParallelismOk() (*int64, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *BackupInfo) SetParallelism(v int64)`

SetParallelism sets Parallelism field to given value.


### GetPointInTimeRestoreEnabled

`func (o *BackupInfo) GetPointInTimeRestoreEnabled() bool`

GetPointInTimeRestoreEnabled returns the PointInTimeRestoreEnabled field if non-nil, zero value otherwise.

### GetPointInTimeRestoreEnabledOk

`func (o *BackupInfo) GetPointInTimeRestoreEnabledOk() (*bool, bool)`

GetPointInTimeRestoreEnabledOk returns a tuple with the PointInTimeRestoreEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeRestoreEnabled

`func (o *BackupInfo) SetPointInTimeRestoreEnabled(v bool)`

SetPointInTimeRestoreEnabled sets PointInTimeRestoreEnabled field to given value.


### GetStorageConfigUUID

`func (o *BackupInfo) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *BackupInfo) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *BackupInfo) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetTimeBeforeDelete

`func (o *BackupInfo) GetTimeBeforeDelete() int64`

GetTimeBeforeDelete returns the TimeBeforeDelete field if non-nil, zero value otherwise.

### GetTimeBeforeDeleteOk

`func (o *BackupInfo) GetTimeBeforeDeleteOk() (*int64, bool)`

GetTimeBeforeDeleteOk returns a tuple with the TimeBeforeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBeforeDelete

`func (o *BackupInfo) SetTimeBeforeDelete(v int64)`

SetTimeBeforeDelete sets TimeBeforeDelete field to given value.


### GetUniverseUUID

`func (o *BackupInfo) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *BackupInfo) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *BackupInfo) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.


### GetUseRoles

`func (o *BackupInfo) GetUseRoles() bool`

GetUseRoles returns the UseRoles field if non-nil, zero value otherwise.

### GetUseRolesOk

`func (o *BackupInfo) GetUseRolesOk() (*bool, bool)`

GetUseRolesOk returns a tuple with the UseRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRoles

`func (o *BackupInfo) SetUseRoles(v bool)`

SetUseRoles sets UseRoles field to given value.


### GetUseTablespaces

`func (o *BackupInfo) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *BackupInfo) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *BackupInfo) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


