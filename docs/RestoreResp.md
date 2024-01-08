# RestoreResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCreatedOnDate** | Pointer to **time.Time** | Backup details. | [optional] 
**BackupType** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **time.Time** | Restore creation time. | [optional] 
**CustomerUUID** | **string** |  | 
**IsSourceUniversePresent** | **bool** |  | 
**RestoreKeyspaceList** | [**[]RestoreKeyspace**](RestoreKeyspace.md) |  | 
**RestoreSizeInBytes** | **int64** |  | 
**RestoreUUID** | **string** |  | 
**SourceUniverseName** | **string** |  | 
**SourceUniverseUUID** | **string** |  | 
**State** | **string** |  | 
**TargetUniverseName** | **string** |  | 
**UniverseUUID** | **string** |  | 
**UpdateTime** | Pointer to **time.Time** | Restore update time. | [optional] 

## Methods

### NewRestoreResp

`func NewRestoreResp(customerUUID string, isSourceUniversePresent bool, restoreKeyspaceList []RestoreKeyspace, restoreSizeInBytes int64, restoreUUID string, sourceUniverseName string, sourceUniverseUUID string, state string, targetUniverseName string, universeUUID string, ) *RestoreResp`

NewRestoreResp instantiates a new RestoreResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreRespWithDefaults

`func NewRestoreRespWithDefaults() *RestoreResp`

NewRestoreRespWithDefaults instantiates a new RestoreResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCreatedOnDate

`func (o *RestoreResp) GetBackupCreatedOnDate() time.Time`

GetBackupCreatedOnDate returns the BackupCreatedOnDate field if non-nil, zero value otherwise.

### GetBackupCreatedOnDateOk

`func (o *RestoreResp) GetBackupCreatedOnDateOk() (*time.Time, bool)`

GetBackupCreatedOnDateOk returns a tuple with the BackupCreatedOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCreatedOnDate

`func (o *RestoreResp) SetBackupCreatedOnDate(v time.Time)`

SetBackupCreatedOnDate sets BackupCreatedOnDate field to given value.

### HasBackupCreatedOnDate

`func (o *RestoreResp) HasBackupCreatedOnDate() bool`

HasBackupCreatedOnDate returns a boolean if a field has been set.

### GetBackupType

`func (o *RestoreResp) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *RestoreResp) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *RestoreResp) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *RestoreResp) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetCreateTime

`func (o *RestoreResp) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *RestoreResp) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *RestoreResp) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *RestoreResp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *RestoreResp) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *RestoreResp) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *RestoreResp) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetIsSourceUniversePresent

`func (o *RestoreResp) GetIsSourceUniversePresent() bool`

GetIsSourceUniversePresent returns the IsSourceUniversePresent field if non-nil, zero value otherwise.

### GetIsSourceUniversePresentOk

`func (o *RestoreResp) GetIsSourceUniversePresentOk() (*bool, bool)`

GetIsSourceUniversePresentOk returns a tuple with the IsSourceUniversePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceUniversePresent

`func (o *RestoreResp) SetIsSourceUniversePresent(v bool)`

SetIsSourceUniversePresent sets IsSourceUniversePresent field to given value.


### GetRestoreKeyspaceList

`func (o *RestoreResp) GetRestoreKeyspaceList() []RestoreKeyspace`

GetRestoreKeyspaceList returns the RestoreKeyspaceList field if non-nil, zero value otherwise.

### GetRestoreKeyspaceListOk

`func (o *RestoreResp) GetRestoreKeyspaceListOk() (*[]RestoreKeyspace, bool)`

GetRestoreKeyspaceListOk returns a tuple with the RestoreKeyspaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreKeyspaceList

`func (o *RestoreResp) SetRestoreKeyspaceList(v []RestoreKeyspace)`

SetRestoreKeyspaceList sets RestoreKeyspaceList field to given value.


### GetRestoreSizeInBytes

`func (o *RestoreResp) GetRestoreSizeInBytes() int64`

GetRestoreSizeInBytes returns the RestoreSizeInBytes field if non-nil, zero value otherwise.

### GetRestoreSizeInBytesOk

`func (o *RestoreResp) GetRestoreSizeInBytesOk() (*int64, bool)`

GetRestoreSizeInBytesOk returns a tuple with the RestoreSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSizeInBytes

`func (o *RestoreResp) SetRestoreSizeInBytes(v int64)`

SetRestoreSizeInBytes sets RestoreSizeInBytes field to given value.


### GetRestoreUUID

`func (o *RestoreResp) GetRestoreUUID() string`

GetRestoreUUID returns the RestoreUUID field if non-nil, zero value otherwise.

### GetRestoreUUIDOk

`func (o *RestoreResp) GetRestoreUUIDOk() (*string, bool)`

GetRestoreUUIDOk returns a tuple with the RestoreUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUUID

`func (o *RestoreResp) SetRestoreUUID(v string)`

SetRestoreUUID sets RestoreUUID field to given value.


### GetSourceUniverseName

`func (o *RestoreResp) GetSourceUniverseName() string`

GetSourceUniverseName returns the SourceUniverseName field if non-nil, zero value otherwise.

### GetSourceUniverseNameOk

`func (o *RestoreResp) GetSourceUniverseNameOk() (*string, bool)`

GetSourceUniverseNameOk returns a tuple with the SourceUniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseName

`func (o *RestoreResp) SetSourceUniverseName(v string)`

SetSourceUniverseName sets SourceUniverseName field to given value.


### GetSourceUniverseUUID

`func (o *RestoreResp) GetSourceUniverseUUID() string`

GetSourceUniverseUUID returns the SourceUniverseUUID field if non-nil, zero value otherwise.

### GetSourceUniverseUUIDOk

`func (o *RestoreResp) GetSourceUniverseUUIDOk() (*string, bool)`

GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseUUID

`func (o *RestoreResp) SetSourceUniverseUUID(v string)`

SetSourceUniverseUUID sets SourceUniverseUUID field to given value.


### GetState

`func (o *RestoreResp) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RestoreResp) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RestoreResp) SetState(v string)`

SetState sets State field to given value.


### GetTargetUniverseName

`func (o *RestoreResp) GetTargetUniverseName() string`

GetTargetUniverseName returns the TargetUniverseName field if non-nil, zero value otherwise.

### GetTargetUniverseNameOk

`func (o *RestoreResp) GetTargetUniverseNameOk() (*string, bool)`

GetTargetUniverseNameOk returns a tuple with the TargetUniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseName

`func (o *RestoreResp) SetTargetUniverseName(v string)`

SetTargetUniverseName sets TargetUniverseName field to given value.


### GetUniverseUUID

`func (o *RestoreResp) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *RestoreResp) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *RestoreResp) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.


### GetUpdateTime

`func (o *RestoreResp) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *RestoreResp) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *RestoreResp) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *RestoreResp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


