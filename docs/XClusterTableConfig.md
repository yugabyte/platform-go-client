# XClusterTableConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupUuid** | Pointer to **string** | The backup config used to do bootstrapping for this table | [optional] 
**BootstrapCreateTime** | Pointer to **time.Time** | Time of the bootstrap of the table | [optional] 
**RestoreTime** | Pointer to **time.Time** | Time of the last try to restore data to the target universe | [optional] 
**RestoreUuid** | Pointer to **string** | The restore config used to do bootstrapping for this table | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**StreamId** | Pointer to **string** | Stream ID if replication is setup; bootstrap ID if the table is bootstrapped | [optional] 
**TableId** | Pointer to **string** | Table ID | [optional] 

## Methods

### NewXClusterTableConfig

`func NewXClusterTableConfig() *XClusterTableConfig`

NewXClusterTableConfig instantiates a new XClusterTableConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterTableConfigWithDefaults

`func NewXClusterTableConfigWithDefaults() *XClusterTableConfig`

NewXClusterTableConfigWithDefaults instantiates a new XClusterTableConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupUuid

`func (o *XClusterTableConfig) GetBackupUuid() string`

GetBackupUuid returns the BackupUuid field if non-nil, zero value otherwise.

### GetBackupUuidOk

`func (o *XClusterTableConfig) GetBackupUuidOk() (*string, bool)`

GetBackupUuidOk returns a tuple with the BackupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUuid

`func (o *XClusterTableConfig) SetBackupUuid(v string)`

SetBackupUuid sets BackupUuid field to given value.

### HasBackupUuid

`func (o *XClusterTableConfig) HasBackupUuid() bool`

HasBackupUuid returns a boolean if a field has been set.

### GetBootstrapCreateTime

`func (o *XClusterTableConfig) GetBootstrapCreateTime() time.Time`

GetBootstrapCreateTime returns the BootstrapCreateTime field if non-nil, zero value otherwise.

### GetBootstrapCreateTimeOk

`func (o *XClusterTableConfig) GetBootstrapCreateTimeOk() (*time.Time, bool)`

GetBootstrapCreateTimeOk returns a tuple with the BootstrapCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapCreateTime

`func (o *XClusterTableConfig) SetBootstrapCreateTime(v time.Time)`

SetBootstrapCreateTime sets BootstrapCreateTime field to given value.

### HasBootstrapCreateTime

`func (o *XClusterTableConfig) HasBootstrapCreateTime() bool`

HasBootstrapCreateTime returns a boolean if a field has been set.

### GetRestoreTime

`func (o *XClusterTableConfig) GetRestoreTime() time.Time`

GetRestoreTime returns the RestoreTime field if non-nil, zero value otherwise.

### GetRestoreTimeOk

`func (o *XClusterTableConfig) GetRestoreTimeOk() (*time.Time, bool)`

GetRestoreTimeOk returns a tuple with the RestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTime

`func (o *XClusterTableConfig) SetRestoreTime(v time.Time)`

SetRestoreTime sets RestoreTime field to given value.

### HasRestoreTime

`func (o *XClusterTableConfig) HasRestoreTime() bool`

HasRestoreTime returns a boolean if a field has been set.

### GetRestoreUuid

`func (o *XClusterTableConfig) GetRestoreUuid() string`

GetRestoreUuid returns the RestoreUuid field if non-nil, zero value otherwise.

### GetRestoreUuidOk

`func (o *XClusterTableConfig) GetRestoreUuidOk() (*string, bool)`

GetRestoreUuidOk returns a tuple with the RestoreUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUuid

`func (o *XClusterTableConfig) SetRestoreUuid(v string)`

SetRestoreUuid sets RestoreUuid field to given value.

### HasRestoreUuid

`func (o *XClusterTableConfig) HasRestoreUuid() bool`

HasRestoreUuid returns a boolean if a field has been set.

### GetStatus

`func (o *XClusterTableConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XClusterTableConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XClusterTableConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XClusterTableConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStreamId

`func (o *XClusterTableConfig) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *XClusterTableConfig) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *XClusterTableConfig) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *XClusterTableConfig) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTableId

`func (o *XClusterTableConfig) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *XClusterTableConfig) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *XClusterTableConfig) SetTableId(v string)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *XClusterTableConfig) HasTableId() bool`

HasTableId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


