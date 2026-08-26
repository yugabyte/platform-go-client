# DrConfigTableDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableId** | Pointer to **string** | Table ID. | [optional] [readonly] 
**StreamId** | Pointer to **string** | Stream ID when replication is set up, or bootstrap ID when bootstrapped. | [optional] [readonly] 
**Status** | Pointer to [**DrConfigReplicationDetailStatus**](DrConfigReplicationDetailStatus.md) |  | [optional] 
**IndexTable** | Pointer to **bool** | Whether this is an index table whose main table is in replication. | [optional] [readonly] 
**NeedBootstrap** | Pointer to **bool** | Whether this table needs bootstrap for replication setup. | [optional] [readonly] 
**ReplicationSetupDone** | Pointer to **bool** | Whether replication is set up for this table. | [optional] [readonly] 
**BootstrapCreateTime** | Pointer to **time.Time** | Time when bootstrap for this table was created. | [optional] [readonly] 
**RestoreTime** | Pointer to **time.Time** | Time of the last restore attempt to the target universe. | [optional] [readonly] 
**ReplicationStatusErrors** | Pointer to **[]string** | Human-readable replication error messages, if any. | [optional] [readonly] 
**SourceTableInfo** | Pointer to [**TableInfo**](TableInfo.md) |  | [optional] 
**TargetTableInfo** | Pointer to [**TableInfo**](TableInfo.md) |  | [optional] 

## Methods

### NewDrConfigTableDetail

`func NewDrConfigTableDetail() *DrConfigTableDetail`

NewDrConfigTableDetail instantiates a new DrConfigTableDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigTableDetailWithDefaults

`func NewDrConfigTableDetailWithDefaults() *DrConfigTableDetail`

NewDrConfigTableDetailWithDefaults instantiates a new DrConfigTableDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableId

`func (o *DrConfigTableDetail) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *DrConfigTableDetail) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *DrConfigTableDetail) SetTableId(v string)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *DrConfigTableDetail) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetStreamId

`func (o *DrConfigTableDetail) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *DrConfigTableDetail) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *DrConfigTableDetail) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *DrConfigTableDetail) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetStatus

`func (o *DrConfigTableDetail) GetStatus() DrConfigReplicationDetailStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DrConfigTableDetail) GetStatusOk() (*DrConfigReplicationDetailStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DrConfigTableDetail) SetStatus(v DrConfigReplicationDetailStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DrConfigTableDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIndexTable

`func (o *DrConfigTableDetail) GetIndexTable() bool`

GetIndexTable returns the IndexTable field if non-nil, zero value otherwise.

### GetIndexTableOk

`func (o *DrConfigTableDetail) GetIndexTableOk() (*bool, bool)`

GetIndexTableOk returns a tuple with the IndexTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexTable

`func (o *DrConfigTableDetail) SetIndexTable(v bool)`

SetIndexTable sets IndexTable field to given value.

### HasIndexTable

`func (o *DrConfigTableDetail) HasIndexTable() bool`

HasIndexTable returns a boolean if a field has been set.

### GetNeedBootstrap

`func (o *DrConfigTableDetail) GetNeedBootstrap() bool`

GetNeedBootstrap returns the NeedBootstrap field if non-nil, zero value otherwise.

### GetNeedBootstrapOk

`func (o *DrConfigTableDetail) GetNeedBootstrapOk() (*bool, bool)`

GetNeedBootstrapOk returns a tuple with the NeedBootstrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedBootstrap

`func (o *DrConfigTableDetail) SetNeedBootstrap(v bool)`

SetNeedBootstrap sets NeedBootstrap field to given value.

### HasNeedBootstrap

`func (o *DrConfigTableDetail) HasNeedBootstrap() bool`

HasNeedBootstrap returns a boolean if a field has been set.

### GetReplicationSetupDone

`func (o *DrConfigTableDetail) GetReplicationSetupDone() bool`

GetReplicationSetupDone returns the ReplicationSetupDone field if non-nil, zero value otherwise.

### GetReplicationSetupDoneOk

`func (o *DrConfigTableDetail) GetReplicationSetupDoneOk() (*bool, bool)`

GetReplicationSetupDoneOk returns a tuple with the ReplicationSetupDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSetupDone

`func (o *DrConfigTableDetail) SetReplicationSetupDone(v bool)`

SetReplicationSetupDone sets ReplicationSetupDone field to given value.

### HasReplicationSetupDone

`func (o *DrConfigTableDetail) HasReplicationSetupDone() bool`

HasReplicationSetupDone returns a boolean if a field has been set.

### GetBootstrapCreateTime

`func (o *DrConfigTableDetail) GetBootstrapCreateTime() time.Time`

GetBootstrapCreateTime returns the BootstrapCreateTime field if non-nil, zero value otherwise.

### GetBootstrapCreateTimeOk

`func (o *DrConfigTableDetail) GetBootstrapCreateTimeOk() (*time.Time, bool)`

GetBootstrapCreateTimeOk returns a tuple with the BootstrapCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapCreateTime

`func (o *DrConfigTableDetail) SetBootstrapCreateTime(v time.Time)`

SetBootstrapCreateTime sets BootstrapCreateTime field to given value.

### HasBootstrapCreateTime

`func (o *DrConfigTableDetail) HasBootstrapCreateTime() bool`

HasBootstrapCreateTime returns a boolean if a field has been set.

### GetRestoreTime

`func (o *DrConfigTableDetail) GetRestoreTime() time.Time`

GetRestoreTime returns the RestoreTime field if non-nil, zero value otherwise.

### GetRestoreTimeOk

`func (o *DrConfigTableDetail) GetRestoreTimeOk() (*time.Time, bool)`

GetRestoreTimeOk returns a tuple with the RestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTime

`func (o *DrConfigTableDetail) SetRestoreTime(v time.Time)`

SetRestoreTime sets RestoreTime field to given value.

### HasRestoreTime

`func (o *DrConfigTableDetail) HasRestoreTime() bool`

HasRestoreTime returns a boolean if a field has been set.

### GetReplicationStatusErrors

`func (o *DrConfigTableDetail) GetReplicationStatusErrors() []string`

GetReplicationStatusErrors returns the ReplicationStatusErrors field if non-nil, zero value otherwise.

### GetReplicationStatusErrorsOk

`func (o *DrConfigTableDetail) GetReplicationStatusErrorsOk() (*[]string, bool)`

GetReplicationStatusErrorsOk returns a tuple with the ReplicationStatusErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatusErrors

`func (o *DrConfigTableDetail) SetReplicationStatusErrors(v []string)`

SetReplicationStatusErrors sets ReplicationStatusErrors field to given value.

### HasReplicationStatusErrors

`func (o *DrConfigTableDetail) HasReplicationStatusErrors() bool`

HasReplicationStatusErrors returns a boolean if a field has been set.

### GetSourceTableInfo

`func (o *DrConfigTableDetail) GetSourceTableInfo() TableInfo`

GetSourceTableInfo returns the SourceTableInfo field if non-nil, zero value otherwise.

### GetSourceTableInfoOk

`func (o *DrConfigTableDetail) GetSourceTableInfoOk() (*TableInfo, bool)`

GetSourceTableInfoOk returns a tuple with the SourceTableInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableInfo

`func (o *DrConfigTableDetail) SetSourceTableInfo(v TableInfo)`

SetSourceTableInfo sets SourceTableInfo field to given value.

### HasSourceTableInfo

`func (o *DrConfigTableDetail) HasSourceTableInfo() bool`

HasSourceTableInfo returns a boolean if a field has been set.

### GetTargetTableInfo

`func (o *DrConfigTableDetail) GetTargetTableInfo() TableInfo`

GetTargetTableInfo returns the TargetTableInfo field if non-nil, zero value otherwise.

### GetTargetTableInfoOk

`func (o *DrConfigTableDetail) GetTargetTableInfoOk() (*TableInfo, bool)`

GetTargetTableInfoOk returns a tuple with the TargetTableInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTableInfo

`func (o *DrConfigTableDetail) SetTargetTableInfo(v TableInfo)`

SetTargetTableInfo sets TargetTableInfo field to given value.

### HasTargetTableInfo

`func (o *DrConfigTableDetail) HasTargetTableInfo() bool`

HasTargetTableInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


