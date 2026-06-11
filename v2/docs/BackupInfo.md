# BackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Backup UUID. | [optional] [readonly] 
**BaseBackupUuid** | Pointer to **string** | Base backup UUID for incremental chains. | [optional] [readonly] 
**State** | Pointer to [**BackupState**](BackupState.md) |  | [optional] 
**StorageConfigUuid** | Pointer to **string** | Storage config UUID used for the backup. | [optional] [readonly] 
**KmsConfigUuid** | Pointer to **string** | KMS config UUID used to encrypt the backup. | [optional] [readonly] 
**TaskUuid** | Pointer to **string** | Task UUID for the backup operation. | [optional] [readonly] 
**CreateTime** | Pointer to **time.Time** | Time when the backup was created. | [optional] [readonly] 
**UpdateTime** | Pointer to **time.Time** | Time when the backup was last updated. | [optional] [readonly] 
**CompletionTime** | Pointer to **time.Time** | Time when the backup completed. | [optional] [readonly] 
**TotalBackupSizeInBytes** | Pointer to **int64** | Total backup size, in bytes. | [optional] [readonly] 
**Sse** | Pointer to **bool** | Whether server-side encryption is enabled for the backup. | [optional] [readonly] 
**TableByTableBackup** | Pointer to **bool** | Whether the backup was taken table by table. | [optional] [readonly] 
**HasIncrementalBackups** | Pointer to **bool** | Whether incremental backups exist for this backup chain. | [optional] [readonly] 
**LastIncrementalBackupTime** | Pointer to **time.Time** | Time of the most recent incremental backup in the chain. | [optional] [readonly] 
**LastBackupState** | Pointer to [**BackupState**](BackupState.md) |  | [optional] 
**IsStorageConfigPresent** | Pointer to **bool** | Whether the storage config still exists. | [optional] [readonly] 
**IsUniversePresent** | Pointer to **bool** | Whether the source universe still exists. | [optional] [readonly] 
**FullChainSizeInBytes** | Pointer to **int64** | Total size of the full backup chain, in bytes. | [optional] [readonly] 

## Methods

### NewBackupInfo

`func NewBackupInfo() *BackupInfo`

NewBackupInfo instantiates a new BackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupInfoWithDefaults

`func NewBackupInfoWithDefaults() *BackupInfo`

NewBackupInfoWithDefaults instantiates a new BackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *BackupInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BackupInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BackupInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BackupInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetBaseBackupUuid

`func (o *BackupInfo) GetBaseBackupUuid() string`

GetBaseBackupUuid returns the BaseBackupUuid field if non-nil, zero value otherwise.

### GetBaseBackupUuidOk

`func (o *BackupInfo) GetBaseBackupUuidOk() (*string, bool)`

GetBaseBackupUuidOk returns a tuple with the BaseBackupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBackupUuid

`func (o *BackupInfo) SetBaseBackupUuid(v string)`

SetBaseBackupUuid sets BaseBackupUuid field to given value.

### HasBaseBackupUuid

`func (o *BackupInfo) HasBaseBackupUuid() bool`

HasBaseBackupUuid returns a boolean if a field has been set.

### GetState

`func (o *BackupInfo) GetState() BackupState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupInfo) GetStateOk() (*BackupState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupInfo) SetState(v BackupState)`

SetState sets State field to given value.

### HasState

`func (o *BackupInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageConfigUuid

`func (o *BackupInfo) GetStorageConfigUuid() string`

GetStorageConfigUuid returns the StorageConfigUuid field if non-nil, zero value otherwise.

### GetStorageConfigUuidOk

`func (o *BackupInfo) GetStorageConfigUuidOk() (*string, bool)`

GetStorageConfigUuidOk returns a tuple with the StorageConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUuid

`func (o *BackupInfo) SetStorageConfigUuid(v string)`

SetStorageConfigUuid sets StorageConfigUuid field to given value.

### HasStorageConfigUuid

`func (o *BackupInfo) HasStorageConfigUuid() bool`

HasStorageConfigUuid returns a boolean if a field has been set.

### GetKmsConfigUuid

`func (o *BackupInfo) GetKmsConfigUuid() string`

GetKmsConfigUuid returns the KmsConfigUuid field if non-nil, zero value otherwise.

### GetKmsConfigUuidOk

`func (o *BackupInfo) GetKmsConfigUuidOk() (*string, bool)`

GetKmsConfigUuidOk returns a tuple with the KmsConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUuid

`func (o *BackupInfo) SetKmsConfigUuid(v string)`

SetKmsConfigUuid sets KmsConfigUuid field to given value.

### HasKmsConfigUuid

`func (o *BackupInfo) HasKmsConfigUuid() bool`

HasKmsConfigUuid returns a boolean if a field has been set.

### GetTaskUuid

`func (o *BackupInfo) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *BackupInfo) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *BackupInfo) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.

### HasTaskUuid

`func (o *BackupInfo) HasTaskUuid() bool`

HasTaskUuid returns a boolean if a field has been set.

### GetCreateTime

`func (o *BackupInfo) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *BackupInfo) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *BackupInfo) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *BackupInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *BackupInfo) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *BackupInfo) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *BackupInfo) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *BackupInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *BackupInfo) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *BackupInfo) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *BackupInfo) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *BackupInfo) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetTotalBackupSizeInBytes

`func (o *BackupInfo) GetTotalBackupSizeInBytes() int64`

GetTotalBackupSizeInBytes returns the TotalBackupSizeInBytes field if non-nil, zero value otherwise.

### GetTotalBackupSizeInBytesOk

`func (o *BackupInfo) GetTotalBackupSizeInBytesOk() (*int64, bool)`

GetTotalBackupSizeInBytesOk returns a tuple with the TotalBackupSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackupSizeInBytes

`func (o *BackupInfo) SetTotalBackupSizeInBytes(v int64)`

SetTotalBackupSizeInBytes sets TotalBackupSizeInBytes field to given value.

### HasTotalBackupSizeInBytes

`func (o *BackupInfo) HasTotalBackupSizeInBytes() bool`

HasTotalBackupSizeInBytes returns a boolean if a field has been set.

### GetSse

`func (o *BackupInfo) GetSse() bool`

GetSse returns the Sse field if non-nil, zero value otherwise.

### GetSseOk

`func (o *BackupInfo) GetSseOk() (*bool, bool)`

GetSseOk returns a tuple with the Sse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSse

`func (o *BackupInfo) SetSse(v bool)`

SetSse sets Sse field to given value.

### HasSse

`func (o *BackupInfo) HasSse() bool`

HasSse returns a boolean if a field has been set.

### GetTableByTableBackup

`func (o *BackupInfo) GetTableByTableBackup() bool`

GetTableByTableBackup returns the TableByTableBackup field if non-nil, zero value otherwise.

### GetTableByTableBackupOk

`func (o *BackupInfo) GetTableByTableBackupOk() (*bool, bool)`

GetTableByTableBackupOk returns a tuple with the TableByTableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableByTableBackup

`func (o *BackupInfo) SetTableByTableBackup(v bool)`

SetTableByTableBackup sets TableByTableBackup field to given value.

### HasTableByTableBackup

`func (o *BackupInfo) HasTableByTableBackup() bool`

HasTableByTableBackup returns a boolean if a field has been set.

### GetHasIncrementalBackups

`func (o *BackupInfo) GetHasIncrementalBackups() bool`

GetHasIncrementalBackups returns the HasIncrementalBackups field if non-nil, zero value otherwise.

### GetHasIncrementalBackupsOk

`func (o *BackupInfo) GetHasIncrementalBackupsOk() (*bool, bool)`

GetHasIncrementalBackupsOk returns a tuple with the HasIncrementalBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIncrementalBackups

`func (o *BackupInfo) SetHasIncrementalBackups(v bool)`

SetHasIncrementalBackups sets HasIncrementalBackups field to given value.

### HasHasIncrementalBackups

`func (o *BackupInfo) HasHasIncrementalBackups() bool`

HasHasIncrementalBackups returns a boolean if a field has been set.

### GetLastIncrementalBackupTime

`func (o *BackupInfo) GetLastIncrementalBackupTime() time.Time`

GetLastIncrementalBackupTime returns the LastIncrementalBackupTime field if non-nil, zero value otherwise.

### GetLastIncrementalBackupTimeOk

`func (o *BackupInfo) GetLastIncrementalBackupTimeOk() (*time.Time, bool)`

GetLastIncrementalBackupTimeOk returns a tuple with the LastIncrementalBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIncrementalBackupTime

`func (o *BackupInfo) SetLastIncrementalBackupTime(v time.Time)`

SetLastIncrementalBackupTime sets LastIncrementalBackupTime field to given value.

### HasLastIncrementalBackupTime

`func (o *BackupInfo) HasLastIncrementalBackupTime() bool`

HasLastIncrementalBackupTime returns a boolean if a field has been set.

### GetLastBackupState

`func (o *BackupInfo) GetLastBackupState() BackupState`

GetLastBackupState returns the LastBackupState field if non-nil, zero value otherwise.

### GetLastBackupStateOk

`func (o *BackupInfo) GetLastBackupStateOk() (*BackupState, bool)`

GetLastBackupStateOk returns a tuple with the LastBackupState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupState

`func (o *BackupInfo) SetLastBackupState(v BackupState)`

SetLastBackupState sets LastBackupState field to given value.

### HasLastBackupState

`func (o *BackupInfo) HasLastBackupState() bool`

HasLastBackupState returns a boolean if a field has been set.

### GetIsStorageConfigPresent

`func (o *BackupInfo) GetIsStorageConfigPresent() bool`

GetIsStorageConfigPresent returns the IsStorageConfigPresent field if non-nil, zero value otherwise.

### GetIsStorageConfigPresentOk

`func (o *BackupInfo) GetIsStorageConfigPresentOk() (*bool, bool)`

GetIsStorageConfigPresentOk returns a tuple with the IsStorageConfigPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStorageConfigPresent

`func (o *BackupInfo) SetIsStorageConfigPresent(v bool)`

SetIsStorageConfigPresent sets IsStorageConfigPresent field to given value.

### HasIsStorageConfigPresent

`func (o *BackupInfo) HasIsStorageConfigPresent() bool`

HasIsStorageConfigPresent returns a boolean if a field has been set.

### GetIsUniversePresent

`func (o *BackupInfo) GetIsUniversePresent() bool`

GetIsUniversePresent returns the IsUniversePresent field if non-nil, zero value otherwise.

### GetIsUniversePresentOk

`func (o *BackupInfo) GetIsUniversePresentOk() (*bool, bool)`

GetIsUniversePresentOk returns a tuple with the IsUniversePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUniversePresent

`func (o *BackupInfo) SetIsUniversePresent(v bool)`

SetIsUniversePresent sets IsUniversePresent field to given value.

### HasIsUniversePresent

`func (o *BackupInfo) HasIsUniversePresent() bool`

HasIsUniversePresent returns a boolean if a field has been set.

### GetFullChainSizeInBytes

`func (o *BackupInfo) GetFullChainSizeInBytes() int64`

GetFullChainSizeInBytes returns the FullChainSizeInBytes field if non-nil, zero value otherwise.

### GetFullChainSizeInBytesOk

`func (o *BackupInfo) GetFullChainSizeInBytesOk() (*int64, bool)`

GetFullChainSizeInBytesOk returns a tuple with the FullChainSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullChainSizeInBytes

`func (o *BackupInfo) SetFullChainSizeInBytes(v int64)`

SetFullChainSizeInBytes sets FullChainSizeInBytes field to given value.

### HasFullChainSizeInBytes

`func (o *BackupInfo) HasFullChainSizeInBytes() bool`

HasFullChainSizeInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


