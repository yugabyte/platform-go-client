# RestoreKeyspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSizeFromStorageLocation** | **int64** |  | 
**CompleteTime** | Pointer to **time.Time** | RestoreKeyspace task completion time | [optional] 
**CreateTime** | Pointer to **time.Time** | RestoreKeyspace task creation time | [optional] 
**RestoreUUID** | Pointer to **string** | Universe-level Restore UUID | [optional] [readonly] 
**SourceKeyspace** | Pointer to **string** | Source keyspace name | [optional] [readonly] 
**State** | Pointer to **string** | State of the keyspace restore | [optional] [readonly] 
**StorageLocation** | Pointer to **string** | Storage location name | [optional] [readonly] 
**TargetKeyspace** | Pointer to **string** | Target keyspace name | [optional] [readonly] 
**TaskUUID** | Pointer to **string** | Restore Keyspace task UUID | [optional] [readonly] 
**Uuid** | Pointer to **string** | Restore keyspace UUID | [optional] [readonly] 

## Methods

### NewRestoreKeyspace

`func NewRestoreKeyspace(backupSizeFromStorageLocation int64, ) *RestoreKeyspace`

NewRestoreKeyspace instantiates a new RestoreKeyspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreKeyspaceWithDefaults

`func NewRestoreKeyspaceWithDefaults() *RestoreKeyspace`

NewRestoreKeyspaceWithDefaults instantiates a new RestoreKeyspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupSizeFromStorageLocation

`func (o *RestoreKeyspace) GetBackupSizeFromStorageLocation() int64`

GetBackupSizeFromStorageLocation returns the BackupSizeFromStorageLocation field if non-nil, zero value otherwise.

### GetBackupSizeFromStorageLocationOk

`func (o *RestoreKeyspace) GetBackupSizeFromStorageLocationOk() (*int64, bool)`

GetBackupSizeFromStorageLocationOk returns a tuple with the BackupSizeFromStorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSizeFromStorageLocation

`func (o *RestoreKeyspace) SetBackupSizeFromStorageLocation(v int64)`

SetBackupSizeFromStorageLocation sets BackupSizeFromStorageLocation field to given value.


### GetCompleteTime

`func (o *RestoreKeyspace) GetCompleteTime() time.Time`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *RestoreKeyspace) GetCompleteTimeOk() (*time.Time, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *RestoreKeyspace) SetCompleteTime(v time.Time)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *RestoreKeyspace) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.

### GetCreateTime

`func (o *RestoreKeyspace) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *RestoreKeyspace) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *RestoreKeyspace) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *RestoreKeyspace) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetRestoreUUID

`func (o *RestoreKeyspace) GetRestoreUUID() string`

GetRestoreUUID returns the RestoreUUID field if non-nil, zero value otherwise.

### GetRestoreUUIDOk

`func (o *RestoreKeyspace) GetRestoreUUIDOk() (*string, bool)`

GetRestoreUUIDOk returns a tuple with the RestoreUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUUID

`func (o *RestoreKeyspace) SetRestoreUUID(v string)`

SetRestoreUUID sets RestoreUUID field to given value.

### HasRestoreUUID

`func (o *RestoreKeyspace) HasRestoreUUID() bool`

HasRestoreUUID returns a boolean if a field has been set.

### GetSourceKeyspace

`func (o *RestoreKeyspace) GetSourceKeyspace() string`

GetSourceKeyspace returns the SourceKeyspace field if non-nil, zero value otherwise.

### GetSourceKeyspaceOk

`func (o *RestoreKeyspace) GetSourceKeyspaceOk() (*string, bool)`

GetSourceKeyspaceOk returns a tuple with the SourceKeyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKeyspace

`func (o *RestoreKeyspace) SetSourceKeyspace(v string)`

SetSourceKeyspace sets SourceKeyspace field to given value.

### HasSourceKeyspace

`func (o *RestoreKeyspace) HasSourceKeyspace() bool`

HasSourceKeyspace returns a boolean if a field has been set.

### GetState

`func (o *RestoreKeyspace) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RestoreKeyspace) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RestoreKeyspace) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RestoreKeyspace) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageLocation

`func (o *RestoreKeyspace) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *RestoreKeyspace) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *RestoreKeyspace) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *RestoreKeyspace) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### GetTargetKeyspace

`func (o *RestoreKeyspace) GetTargetKeyspace() string`

GetTargetKeyspace returns the TargetKeyspace field if non-nil, zero value otherwise.

### GetTargetKeyspaceOk

`func (o *RestoreKeyspace) GetTargetKeyspaceOk() (*string, bool)`

GetTargetKeyspaceOk returns a tuple with the TargetKeyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKeyspace

`func (o *RestoreKeyspace) SetTargetKeyspace(v string)`

SetTargetKeyspace sets TargetKeyspace field to given value.

### HasTargetKeyspace

`func (o *RestoreKeyspace) HasTargetKeyspace() bool`

HasTargetKeyspace returns a boolean if a field has been set.

### GetTaskUUID

`func (o *RestoreKeyspace) GetTaskUUID() string`

GetTaskUUID returns the TaskUUID field if non-nil, zero value otherwise.

### GetTaskUUIDOk

`func (o *RestoreKeyspace) GetTaskUUIDOk() (*string, bool)`

GetTaskUUIDOk returns a tuple with the TaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUUID

`func (o *RestoreKeyspace) SetTaskUUID(v string)`

SetTaskUUID sets TaskUUID field to given value.

### HasTaskUUID

`func (o *RestoreKeyspace) HasTaskUUID() bool`

HasTaskUUID returns a boolean if a field has been set.

### GetUuid

`func (o *RestoreKeyspace) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RestoreKeyspace) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RestoreKeyspace) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RestoreKeyspace) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


