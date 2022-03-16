# BackupResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupType** | **string** |  | 
**BackupUUID** | **string** |  | 
**CompletionTime** | **time.Time** |  | 
**CreateTime** | **time.Time** |  | 
**CustomerUUID** | **string** |  | 
**ExpiryTime** | **time.Time** |  | 
**IsStorageConfigPresent** | **bool** |  | 
**IsUniversePresent** | **bool** |  | 
**OnDemand** | **bool** |  | 
**ResponseList** | [**[]KeyspaceTablesList**](KeyspaceTablesList.md) |  | 
**ScheduleUUID** | **string** |  | 
**Sse** | **bool** |  | 
**State** | **string** |  | 
**StorageConfigUUID** | **string** |  | 
**TaskUUID** | **string** |  | 
**TotalBackupSizeInBytes** | **int64** |  | 
**UniverseName** | **string** |  | 
**UniverseUUID** | **string** |  | 
**UpdateTime** | **time.Time** |  | 

## Methods

### NewBackupResp

`func NewBackupResp(backupType string, backupUUID string, completionTime time.Time, createTime time.Time, customerUUID string, expiryTime time.Time, isStorageConfigPresent bool, isUniversePresent bool, onDemand bool, responseList []KeyspaceTablesList, scheduleUUID string, sse bool, state string, storageConfigUUID string, taskUUID string, totalBackupSizeInBytes int64, universeName string, universeUUID string, updateTime time.Time, ) *BackupResp`

NewBackupResp instantiates a new BackupResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRespWithDefaults

`func NewBackupRespWithDefaults() *BackupResp`

NewBackupRespWithDefaults instantiates a new BackupResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupType

`func (o *BackupResp) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupResp) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupResp) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.


### GetBackupUUID

`func (o *BackupResp) GetBackupUUID() string`

GetBackupUUID returns the BackupUUID field if non-nil, zero value otherwise.

### GetBackupUUIDOk

`func (o *BackupResp) GetBackupUUIDOk() (*string, bool)`

GetBackupUUIDOk returns a tuple with the BackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUID

`func (o *BackupResp) SetBackupUUID(v string)`

SetBackupUUID sets BackupUUID field to given value.


### GetCompletionTime

`func (o *BackupResp) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *BackupResp) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *BackupResp) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.


### GetCreateTime

`func (o *BackupResp) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *BackupResp) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *BackupResp) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetCustomerUUID

`func (o *BackupResp) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *BackupResp) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *BackupResp) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetExpiryTime

`func (o *BackupResp) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *BackupResp) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *BackupResp) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.


### GetIsStorageConfigPresent

`func (o *BackupResp) GetIsStorageConfigPresent() bool`

GetIsStorageConfigPresent returns the IsStorageConfigPresent field if non-nil, zero value otherwise.

### GetIsStorageConfigPresentOk

`func (o *BackupResp) GetIsStorageConfigPresentOk() (*bool, bool)`

GetIsStorageConfigPresentOk returns a tuple with the IsStorageConfigPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStorageConfigPresent

`func (o *BackupResp) SetIsStorageConfigPresent(v bool)`

SetIsStorageConfigPresent sets IsStorageConfigPresent field to given value.


### GetIsUniversePresent

`func (o *BackupResp) GetIsUniversePresent() bool`

GetIsUniversePresent returns the IsUniversePresent field if non-nil, zero value otherwise.

### GetIsUniversePresentOk

`func (o *BackupResp) GetIsUniversePresentOk() (*bool, bool)`

GetIsUniversePresentOk returns a tuple with the IsUniversePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUniversePresent

`func (o *BackupResp) SetIsUniversePresent(v bool)`

SetIsUniversePresent sets IsUniversePresent field to given value.


### GetOnDemand

`func (o *BackupResp) GetOnDemand() bool`

GetOnDemand returns the OnDemand field if non-nil, zero value otherwise.

### GetOnDemandOk

`func (o *BackupResp) GetOnDemandOk() (*bool, bool)`

GetOnDemandOk returns a tuple with the OnDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemand

`func (o *BackupResp) SetOnDemand(v bool)`

SetOnDemand sets OnDemand field to given value.


### GetResponseList

`func (o *BackupResp) GetResponseList() []KeyspaceTablesList`

GetResponseList returns the ResponseList field if non-nil, zero value otherwise.

### GetResponseListOk

`func (o *BackupResp) GetResponseListOk() (*[]KeyspaceTablesList, bool)`

GetResponseListOk returns a tuple with the ResponseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseList

`func (o *BackupResp) SetResponseList(v []KeyspaceTablesList)`

SetResponseList sets ResponseList field to given value.


### GetScheduleUUID

`func (o *BackupResp) GetScheduleUUID() string`

GetScheduleUUID returns the ScheduleUUID field if non-nil, zero value otherwise.

### GetScheduleUUIDOk

`func (o *BackupResp) GetScheduleUUIDOk() (*string, bool)`

GetScheduleUUIDOk returns a tuple with the ScheduleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUID

`func (o *BackupResp) SetScheduleUUID(v string)`

SetScheduleUUID sets ScheduleUUID field to given value.


### GetSse

`func (o *BackupResp) GetSse() bool`

GetSse returns the Sse field if non-nil, zero value otherwise.

### GetSseOk

`func (o *BackupResp) GetSseOk() (*bool, bool)`

GetSseOk returns a tuple with the Sse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSse

`func (o *BackupResp) SetSse(v bool)`

SetSse sets Sse field to given value.


### GetState

`func (o *BackupResp) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupResp) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupResp) SetState(v string)`

SetState sets State field to given value.


### GetStorageConfigUUID

`func (o *BackupResp) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *BackupResp) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *BackupResp) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetTaskUUID

`func (o *BackupResp) GetTaskUUID() string`

GetTaskUUID returns the TaskUUID field if non-nil, zero value otherwise.

### GetTaskUUIDOk

`func (o *BackupResp) GetTaskUUIDOk() (*string, bool)`

GetTaskUUIDOk returns a tuple with the TaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUUID

`func (o *BackupResp) SetTaskUUID(v string)`

SetTaskUUID sets TaskUUID field to given value.


### GetTotalBackupSizeInBytes

`func (o *BackupResp) GetTotalBackupSizeInBytes() int64`

GetTotalBackupSizeInBytes returns the TotalBackupSizeInBytes field if non-nil, zero value otherwise.

### GetTotalBackupSizeInBytesOk

`func (o *BackupResp) GetTotalBackupSizeInBytesOk() (*int64, bool)`

GetTotalBackupSizeInBytesOk returns a tuple with the TotalBackupSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackupSizeInBytes

`func (o *BackupResp) SetTotalBackupSizeInBytes(v int64)`

SetTotalBackupSizeInBytes sets TotalBackupSizeInBytes field to given value.


### GetUniverseName

`func (o *BackupResp) GetUniverseName() string`

GetUniverseName returns the UniverseName field if non-nil, zero value otherwise.

### GetUniverseNameOk

`func (o *BackupResp) GetUniverseNameOk() (*string, bool)`

GetUniverseNameOk returns a tuple with the UniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseName

`func (o *BackupResp) SetUniverseName(v string)`

SetUniverseName sets UniverseName field to given value.


### GetUniverseUUID

`func (o *BackupResp) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *BackupResp) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *BackupResp) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.


### GetUpdateTime

`func (o *BackupResp) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *BackupResp) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *BackupResp) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


