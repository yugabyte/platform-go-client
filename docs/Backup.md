# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCategory** | **string** |  | 
**BackupInfo** | Pointer to [**BackupTableParams**](BackupTableParams.md) |  | [optional] 
**BackupUUID** | Pointer to **string** | Backup UUID | [optional] [readonly] 
**BaseBackupUUID** | Pointer to **string** | Base backup UUID | [optional] [readonly] 
**Category** | Pointer to **string** | Category of the backup | [optional] 
**CompletionTime** | Pointer to **time.Time** | Backup completion time | [optional] [readonly] 
**CreateTime** | **time.Time** |  | 
**CustomerUUID** | Pointer to **string** | Customer UUID that owns this backup | [optional] 
**Expiry** | Pointer to **time.Time** | Expiry time (unix timestamp) of the backup | [optional] 
**ExpiryTimeUnit** | Pointer to **string** | Time unit for backup expiry time | [optional] 
**IncrementalBackup** | **bool** |  | 
**ParentBackup** | **bool** |  | 
**ScheduleUUID** | Pointer to **string** | Schedule UUID, if this backup is part of a schedule | [optional] 
**State** | Pointer to **string** | State of the backup | [optional] [readonly] 
**StorageConfigUUID** | Pointer to **string** | Storage Config UUID that created this backup | [optional] 
**TaskUUID** | Pointer to **string** | Backup UUID | [optional] [readonly] 
**UniverseName** | Pointer to **string** | Universe name that created this backup | [optional] 
**UniverseUUID** | Pointer to **string** | Universe UUID that created this backup | [optional] 
**UpdateTime** | **time.Time** |  | 
**Version** | Pointer to **string** | Version of the backup in a category | [optional] 

## Methods

### NewBackup

`func NewBackup(backupCategory string, createTime time.Time, incrementalBackup bool, parentBackup bool, updateTime time.Time, ) *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCategory

`func (o *Backup) GetBackupCategory() string`

GetBackupCategory returns the BackupCategory field if non-nil, zero value otherwise.

### GetBackupCategoryOk

`func (o *Backup) GetBackupCategoryOk() (*string, bool)`

GetBackupCategoryOk returns a tuple with the BackupCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCategory

`func (o *Backup) SetBackupCategory(v string)`

SetBackupCategory sets BackupCategory field to given value.


### GetBackupInfo

`func (o *Backup) GetBackupInfo() BackupTableParams`

GetBackupInfo returns the BackupInfo field if non-nil, zero value otherwise.

### GetBackupInfoOk

`func (o *Backup) GetBackupInfoOk() (*BackupTableParams, bool)`

GetBackupInfoOk returns a tuple with the BackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInfo

`func (o *Backup) SetBackupInfo(v BackupTableParams)`

SetBackupInfo sets BackupInfo field to given value.

### HasBackupInfo

`func (o *Backup) HasBackupInfo() bool`

HasBackupInfo returns a boolean if a field has been set.

### GetBackupUUID

`func (o *Backup) GetBackupUUID() string`

GetBackupUUID returns the BackupUUID field if non-nil, zero value otherwise.

### GetBackupUUIDOk

`func (o *Backup) GetBackupUUIDOk() (*string, bool)`

GetBackupUUIDOk returns a tuple with the BackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUID

`func (o *Backup) SetBackupUUID(v string)`

SetBackupUUID sets BackupUUID field to given value.

### HasBackupUUID

`func (o *Backup) HasBackupUUID() bool`

HasBackupUUID returns a boolean if a field has been set.

### GetBaseBackupUUID

`func (o *Backup) GetBaseBackupUUID() string`

GetBaseBackupUUID returns the BaseBackupUUID field if non-nil, zero value otherwise.

### GetBaseBackupUUIDOk

`func (o *Backup) GetBaseBackupUUIDOk() (*string, bool)`

GetBaseBackupUUIDOk returns a tuple with the BaseBackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBackupUUID

`func (o *Backup) SetBaseBackupUUID(v string)`

SetBaseBackupUUID sets BaseBackupUUID field to given value.

### HasBaseBackupUUID

`func (o *Backup) HasBaseBackupUUID() bool`

HasBaseBackupUUID returns a boolean if a field has been set.

### GetCategory

`func (o *Backup) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Backup) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Backup) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Backup) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompletionTime

`func (o *Backup) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *Backup) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *Backup) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *Backup) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetCreateTime

`func (o *Backup) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Backup) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Backup) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetCustomerUUID

`func (o *Backup) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *Backup) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *Backup) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *Backup) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetExpiry

`func (o *Backup) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *Backup) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *Backup) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *Backup) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetExpiryTimeUnit

`func (o *Backup) GetExpiryTimeUnit() string`

GetExpiryTimeUnit returns the ExpiryTimeUnit field if non-nil, zero value otherwise.

### GetExpiryTimeUnitOk

`func (o *Backup) GetExpiryTimeUnitOk() (*string, bool)`

GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUnit

`func (o *Backup) SetExpiryTimeUnit(v string)`

SetExpiryTimeUnit sets ExpiryTimeUnit field to given value.

### HasExpiryTimeUnit

`func (o *Backup) HasExpiryTimeUnit() bool`

HasExpiryTimeUnit returns a boolean if a field has been set.

### GetIncrementalBackup

`func (o *Backup) GetIncrementalBackup() bool`

GetIncrementalBackup returns the IncrementalBackup field if non-nil, zero value otherwise.

### GetIncrementalBackupOk

`func (o *Backup) GetIncrementalBackupOk() (*bool, bool)`

GetIncrementalBackupOk returns a tuple with the IncrementalBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackup

`func (o *Backup) SetIncrementalBackup(v bool)`

SetIncrementalBackup sets IncrementalBackup field to given value.


### GetParentBackup

`func (o *Backup) GetParentBackup() bool`

GetParentBackup returns the ParentBackup field if non-nil, zero value otherwise.

### GetParentBackupOk

`func (o *Backup) GetParentBackupOk() (*bool, bool)`

GetParentBackupOk returns a tuple with the ParentBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackup

`func (o *Backup) SetParentBackup(v bool)`

SetParentBackup sets ParentBackup field to given value.


### GetScheduleUUID

`func (o *Backup) GetScheduleUUID() string`

GetScheduleUUID returns the ScheduleUUID field if non-nil, zero value otherwise.

### GetScheduleUUIDOk

`func (o *Backup) GetScheduleUUIDOk() (*string, bool)`

GetScheduleUUIDOk returns a tuple with the ScheduleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUID

`func (o *Backup) SetScheduleUUID(v string)`

SetScheduleUUID sets ScheduleUUID field to given value.

### HasScheduleUUID

`func (o *Backup) HasScheduleUUID() bool`

HasScheduleUUID returns a boolean if a field has been set.

### GetState

`func (o *Backup) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Backup) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Backup) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Backup) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *Backup) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *Backup) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *Backup) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.

### HasStorageConfigUUID

`func (o *Backup) HasStorageConfigUUID() bool`

HasStorageConfigUUID returns a boolean if a field has been set.

### GetTaskUUID

`func (o *Backup) GetTaskUUID() string`

GetTaskUUID returns the TaskUUID field if non-nil, zero value otherwise.

### GetTaskUUIDOk

`func (o *Backup) GetTaskUUIDOk() (*string, bool)`

GetTaskUUIDOk returns a tuple with the TaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUUID

`func (o *Backup) SetTaskUUID(v string)`

SetTaskUUID sets TaskUUID field to given value.

### HasTaskUUID

`func (o *Backup) HasTaskUUID() bool`

HasTaskUUID returns a boolean if a field has been set.

### GetUniverseName

`func (o *Backup) GetUniverseName() string`

GetUniverseName returns the UniverseName field if non-nil, zero value otherwise.

### GetUniverseNameOk

`func (o *Backup) GetUniverseNameOk() (*string, bool)`

GetUniverseNameOk returns a tuple with the UniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseName

`func (o *Backup) SetUniverseName(v string)`

SetUniverseName sets UniverseName field to given value.

### HasUniverseName

`func (o *Backup) HasUniverseName() bool`

HasUniverseName returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *Backup) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *Backup) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *Backup) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *Backup) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Backup) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Backup) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Backup) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.


### GetVersion

`func (o *Backup) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Backup) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Backup) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Backup) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


