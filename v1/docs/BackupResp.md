# BackupResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupType** | **string** |  | 
**Category** | **string** |  | 
**CommonBackupInfo** | [**CommonBackupInfo**](CommonBackupInfo.md) |  | 
**CustomerUUID** | **string** |  | 
**ExpiryTime** | Pointer to **time.Time** | The expiry time for backup. | [optional] 
**ExpiryTimeUnit** | **string** |  | 
**FullChainSizeInBytes** | **int64** |  | 
**HasIncrementalBackups** | **bool** |  | 
**IsFullBackup** | **bool** |  | 
**IsStorageConfigPresent** | **bool** |  | 
**IsUniversePresent** | **bool** |  | 
**LastBackupState** | **string** |  | 
**LastIncrementalBackupTime** | Pointer to **time.Time** | Time for last incremenatal backup. | [optional] 
**OnDemand** | **bool** |  | 
**ScheduleName** | **string** |  | 
**ScheduleUUID** | **string** |  | 
**StorageConfigType** | **string** |  | 
**UniverseName** | **string** |  | 
**UniverseUUID** | **string** |  | 
**UseTablespaces** | **bool** |  | 

## Methods

### NewBackupResp

`func NewBackupResp(backupType string, category string, commonBackupInfo CommonBackupInfo, customerUUID string, expiryTimeUnit string, fullChainSizeInBytes int64, hasIncrementalBackups bool, isFullBackup bool, isStorageConfigPresent bool, isUniversePresent bool, lastBackupState string, onDemand bool, scheduleName string, scheduleUUID string, storageConfigType string, universeName string, universeUUID string, useTablespaces bool, ) *BackupResp`

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


### GetCategory

`func (o *BackupResp) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BackupResp) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BackupResp) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCommonBackupInfo

`func (o *BackupResp) GetCommonBackupInfo() CommonBackupInfo`

GetCommonBackupInfo returns the CommonBackupInfo field if non-nil, zero value otherwise.

### GetCommonBackupInfoOk

`func (o *BackupResp) GetCommonBackupInfoOk() (*CommonBackupInfo, bool)`

GetCommonBackupInfoOk returns a tuple with the CommonBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonBackupInfo

`func (o *BackupResp) SetCommonBackupInfo(v CommonBackupInfo)`

SetCommonBackupInfo sets CommonBackupInfo field to given value.


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

### HasExpiryTime

`func (o *BackupResp) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetExpiryTimeUnit

`func (o *BackupResp) GetExpiryTimeUnit() string`

GetExpiryTimeUnit returns the ExpiryTimeUnit field if non-nil, zero value otherwise.

### GetExpiryTimeUnitOk

`func (o *BackupResp) GetExpiryTimeUnitOk() (*string, bool)`

GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUnit

`func (o *BackupResp) SetExpiryTimeUnit(v string)`

SetExpiryTimeUnit sets ExpiryTimeUnit field to given value.


### GetFullChainSizeInBytes

`func (o *BackupResp) GetFullChainSizeInBytes() int64`

GetFullChainSizeInBytes returns the FullChainSizeInBytes field if non-nil, zero value otherwise.

### GetFullChainSizeInBytesOk

`func (o *BackupResp) GetFullChainSizeInBytesOk() (*int64, bool)`

GetFullChainSizeInBytesOk returns a tuple with the FullChainSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullChainSizeInBytes

`func (o *BackupResp) SetFullChainSizeInBytes(v int64)`

SetFullChainSizeInBytes sets FullChainSizeInBytes field to given value.


### GetHasIncrementalBackups

`func (o *BackupResp) GetHasIncrementalBackups() bool`

GetHasIncrementalBackups returns the HasIncrementalBackups field if non-nil, zero value otherwise.

### GetHasIncrementalBackupsOk

`func (o *BackupResp) GetHasIncrementalBackupsOk() (*bool, bool)`

GetHasIncrementalBackupsOk returns a tuple with the HasIncrementalBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIncrementalBackups

`func (o *BackupResp) SetHasIncrementalBackups(v bool)`

SetHasIncrementalBackups sets HasIncrementalBackups field to given value.


### GetIsFullBackup

`func (o *BackupResp) GetIsFullBackup() bool`

GetIsFullBackup returns the IsFullBackup field if non-nil, zero value otherwise.

### GetIsFullBackupOk

`func (o *BackupResp) GetIsFullBackupOk() (*bool, bool)`

GetIsFullBackupOk returns a tuple with the IsFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullBackup

`func (o *BackupResp) SetIsFullBackup(v bool)`

SetIsFullBackup sets IsFullBackup field to given value.


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


### GetLastBackupState

`func (o *BackupResp) GetLastBackupState() string`

GetLastBackupState returns the LastBackupState field if non-nil, zero value otherwise.

### GetLastBackupStateOk

`func (o *BackupResp) GetLastBackupStateOk() (*string, bool)`

GetLastBackupStateOk returns a tuple with the LastBackupState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupState

`func (o *BackupResp) SetLastBackupState(v string)`

SetLastBackupState sets LastBackupState field to given value.


### GetLastIncrementalBackupTime

`func (o *BackupResp) GetLastIncrementalBackupTime() time.Time`

GetLastIncrementalBackupTime returns the LastIncrementalBackupTime field if non-nil, zero value otherwise.

### GetLastIncrementalBackupTimeOk

`func (o *BackupResp) GetLastIncrementalBackupTimeOk() (*time.Time, bool)`

GetLastIncrementalBackupTimeOk returns a tuple with the LastIncrementalBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIncrementalBackupTime

`func (o *BackupResp) SetLastIncrementalBackupTime(v time.Time)`

SetLastIncrementalBackupTime sets LastIncrementalBackupTime field to given value.

### HasLastIncrementalBackupTime

`func (o *BackupResp) HasLastIncrementalBackupTime() bool`

HasLastIncrementalBackupTime returns a boolean if a field has been set.

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


### GetScheduleName

`func (o *BackupResp) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupResp) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupResp) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.


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


### GetStorageConfigType

`func (o *BackupResp) GetStorageConfigType() string`

GetStorageConfigType returns the StorageConfigType field if non-nil, zero value otherwise.

### GetStorageConfigTypeOk

`func (o *BackupResp) GetStorageConfigTypeOk() (*string, bool)`

GetStorageConfigTypeOk returns a tuple with the StorageConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigType

`func (o *BackupResp) SetStorageConfigType(v string)`

SetStorageConfigType sets StorageConfigType field to given value.


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


### GetUseTablespaces

`func (o *BackupResp) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *BackupResp) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *BackupResp) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


