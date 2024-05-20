# CommonBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupUUID** | **string** |  | 
**BaseBackupUUID** | **string** |  | 
**CompletionTime** | Pointer to **time.Time** | Backup completion time. | [optional] 
**CreateTime** | Pointer to **time.Time** | Backup create time. | [optional] 
**KmsConfigUUID** | **string** |  | 
**ResponseList** | [**[]KeyspaceTablesList**](KeyspaceTablesList.md) |  | 
**Sse** | **bool** |  | 
**State** | **string** |  | 
**StorageConfigUUID** | **string** |  | 
**TableByTableBackup** | **bool** |  | 
**TaskUUID** | **string** |  | 
**TotalBackupSizeInBytes** | **int64** |  | 
**UpdateTime** | Pointer to **time.Time** | Backup update time. | [optional] 

## Methods

### NewCommonBackupInfo

`func NewCommonBackupInfo(backupUUID string, baseBackupUUID string, kmsConfigUUID string, responseList []KeyspaceTablesList, sse bool, state string, storageConfigUUID string, tableByTableBackup bool, taskUUID string, totalBackupSizeInBytes int64, ) *CommonBackupInfo`

NewCommonBackupInfo instantiates a new CommonBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonBackupInfoWithDefaults

`func NewCommonBackupInfoWithDefaults() *CommonBackupInfo`

NewCommonBackupInfoWithDefaults instantiates a new CommonBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupUUID

`func (o *CommonBackupInfo) GetBackupUUID() string`

GetBackupUUID returns the BackupUUID field if non-nil, zero value otherwise.

### GetBackupUUIDOk

`func (o *CommonBackupInfo) GetBackupUUIDOk() (*string, bool)`

GetBackupUUIDOk returns a tuple with the BackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUID

`func (o *CommonBackupInfo) SetBackupUUID(v string)`

SetBackupUUID sets BackupUUID field to given value.


### GetBaseBackupUUID

`func (o *CommonBackupInfo) GetBaseBackupUUID() string`

GetBaseBackupUUID returns the BaseBackupUUID field if non-nil, zero value otherwise.

### GetBaseBackupUUIDOk

`func (o *CommonBackupInfo) GetBaseBackupUUIDOk() (*string, bool)`

GetBaseBackupUUIDOk returns a tuple with the BaseBackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBackupUUID

`func (o *CommonBackupInfo) SetBaseBackupUUID(v string)`

SetBaseBackupUUID sets BaseBackupUUID field to given value.


### GetCompletionTime

`func (o *CommonBackupInfo) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *CommonBackupInfo) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *CommonBackupInfo) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *CommonBackupInfo) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetCreateTime

`func (o *CommonBackupInfo) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CommonBackupInfo) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CommonBackupInfo) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CommonBackupInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetKmsConfigUUID

`func (o *CommonBackupInfo) GetKmsConfigUUID() string`

GetKmsConfigUUID returns the KmsConfigUUID field if non-nil, zero value otherwise.

### GetKmsConfigUUIDOk

`func (o *CommonBackupInfo) GetKmsConfigUUIDOk() (*string, bool)`

GetKmsConfigUUIDOk returns a tuple with the KmsConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUUID

`func (o *CommonBackupInfo) SetKmsConfigUUID(v string)`

SetKmsConfigUUID sets KmsConfigUUID field to given value.


### GetResponseList

`func (o *CommonBackupInfo) GetResponseList() []KeyspaceTablesList`

GetResponseList returns the ResponseList field if non-nil, zero value otherwise.

### GetResponseListOk

`func (o *CommonBackupInfo) GetResponseListOk() (*[]KeyspaceTablesList, bool)`

GetResponseListOk returns a tuple with the ResponseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseList

`func (o *CommonBackupInfo) SetResponseList(v []KeyspaceTablesList)`

SetResponseList sets ResponseList field to given value.


### GetSse

`func (o *CommonBackupInfo) GetSse() bool`

GetSse returns the Sse field if non-nil, zero value otherwise.

### GetSseOk

`func (o *CommonBackupInfo) GetSseOk() (*bool, bool)`

GetSseOk returns a tuple with the Sse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSse

`func (o *CommonBackupInfo) SetSse(v bool)`

SetSse sets Sse field to given value.


### GetState

`func (o *CommonBackupInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CommonBackupInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CommonBackupInfo) SetState(v string)`

SetState sets State field to given value.


### GetStorageConfigUUID

`func (o *CommonBackupInfo) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *CommonBackupInfo) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *CommonBackupInfo) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetTableByTableBackup

`func (o *CommonBackupInfo) GetTableByTableBackup() bool`

GetTableByTableBackup returns the TableByTableBackup field if non-nil, zero value otherwise.

### GetTableByTableBackupOk

`func (o *CommonBackupInfo) GetTableByTableBackupOk() (*bool, bool)`

GetTableByTableBackupOk returns a tuple with the TableByTableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableByTableBackup

`func (o *CommonBackupInfo) SetTableByTableBackup(v bool)`

SetTableByTableBackup sets TableByTableBackup field to given value.


### GetTaskUUID

`func (o *CommonBackupInfo) GetTaskUUID() string`

GetTaskUUID returns the TaskUUID field if non-nil, zero value otherwise.

### GetTaskUUIDOk

`func (o *CommonBackupInfo) GetTaskUUIDOk() (*string, bool)`

GetTaskUUIDOk returns a tuple with the TaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUUID

`func (o *CommonBackupInfo) SetTaskUUID(v string)`

SetTaskUUID sets TaskUUID field to given value.


### GetTotalBackupSizeInBytes

`func (o *CommonBackupInfo) GetTotalBackupSizeInBytes() int64`

GetTotalBackupSizeInBytes returns the TotalBackupSizeInBytes field if non-nil, zero value otherwise.

### GetTotalBackupSizeInBytesOk

`func (o *CommonBackupInfo) GetTotalBackupSizeInBytesOk() (*int64, bool)`

GetTotalBackupSizeInBytesOk returns a tuple with the TotalBackupSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackupSizeInBytes

`func (o *CommonBackupInfo) SetTotalBackupSizeInBytes(v int64)`

SetTotalBackupSizeInBytes sets TotalBackupSizeInBytes field to given value.


### GetUpdateTime

`func (o *CommonBackupInfo) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CommonBackupInfo) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CommonBackupInfo) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CommonBackupInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


