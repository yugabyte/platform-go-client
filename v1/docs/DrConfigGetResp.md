# DrConfigGetResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapParams** | Pointer to [**RestartBootstrapParams**](RestartBootstrapParams.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** | Create time of the DR config | [optional] 
**DbDetails** | Pointer to [**[]XClusterNamespaceConfig**](XClusterNamespaceConfig.md) | WARNING: This is a preview API that could change. List of db details in replication | [optional] 
**Dbs** | Pointer to **[]string** | WARNING: This is a preview API that could change. List of db ids in replication | [optional] 
**DrReplicaUniverseActive** | Pointer to **bool** | Whether the dr replica universe is active | [optional] 
**DrReplicaUniverseState** | Pointer to **string** | WARNING: This is a preview API that could change. The replication status of the dr replica universe. | [optional] 
**DrReplicaUniverseUuid** | Pointer to **string** | DR Replica Universe UUID | [optional] 
**KeyspacePending** | Pointer to **string** | WARNING: This is a preview API that could change. The keyspace name that the current task is working on. | [optional] 
**ModifyTime** | Pointer to **time.Time** | Last modify time of the DR config | [optional] 
**Name** | Pointer to **string** | Disaster recovery config name | [optional] 
**Paused** | Pointer to **bool** | Whether the underlying xCluster config is paused | [optional] 
**PitrConfigs** | Pointer to [**[]PitrConfig**](PitrConfig.md) | WARNING: This is a preview API that could change. The list of PITR configs used for the underlying txn xCluster config | [optional] 
**PitrRetentionPeriodSec** | Pointer to **int64** | WARNING: This is a preview API that could change. PITR Retention Period in seconds | [optional] 
**PitrSnapshotIntervalSec** | Pointer to **int64** | WARNING: This is a preview API that could change. PITR Retention Period in seconds | [optional] 
**PrimaryUniverseActive** | Pointer to **bool** | Whether the primary universe is active | [optional] 
**PrimaryUniverseState** | Pointer to **string** | WARNING: This is a preview API that could change. The replication status of the primary universe. | [optional] 
**PrimaryUniverseUuid** | Pointer to **string** | Primary Universe UUID | [optional] 
**ReplicationGroupName** | Pointer to **string** | Replication group name in the dr replica universe cluster config | [optional] 
**State** | Pointer to **string** | The state of the DR config | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**TableDetails** | Pointer to [**[]XClusterTableConfig**](XClusterTableConfig.md) | Details for each table in replication | [optional] 
**TableType** | Pointer to **string** | The type of tables that are being replicated | [optional] 
**Tables** | Pointer to **[]string** | List of table ids in replication | [optional] 
**Type** | Pointer to **string** | Whether the config is basic, txn, or db scoped xCluster | [optional] 
**Uuid** | Pointer to **string** | DR config UUID | [optional] 
**XclusterConfigUuid** | Pointer to **string** | UUID of the underlying xCluster config that is managing the replication | [optional] 

## Methods

### NewDrConfigGetResp

`func NewDrConfigGetResp() *DrConfigGetResp`

NewDrConfigGetResp instantiates a new DrConfigGetResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigGetRespWithDefaults

`func NewDrConfigGetRespWithDefaults() *DrConfigGetResp`

NewDrConfigGetRespWithDefaults instantiates a new DrConfigGetResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapParams

`func (o *DrConfigGetResp) GetBootstrapParams() RestartBootstrapParams`

GetBootstrapParams returns the BootstrapParams field if non-nil, zero value otherwise.

### GetBootstrapParamsOk

`func (o *DrConfigGetResp) GetBootstrapParamsOk() (*RestartBootstrapParams, bool)`

GetBootstrapParamsOk returns a tuple with the BootstrapParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapParams

`func (o *DrConfigGetResp) SetBootstrapParams(v RestartBootstrapParams)`

SetBootstrapParams sets BootstrapParams field to given value.

### HasBootstrapParams

`func (o *DrConfigGetResp) HasBootstrapParams() bool`

HasBootstrapParams returns a boolean if a field has been set.

### GetCreateTime

`func (o *DrConfigGetResp) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DrConfigGetResp) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DrConfigGetResp) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DrConfigGetResp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDbDetails

`func (o *DrConfigGetResp) GetDbDetails() []XClusterNamespaceConfig`

GetDbDetails returns the DbDetails field if non-nil, zero value otherwise.

### GetDbDetailsOk

`func (o *DrConfigGetResp) GetDbDetailsOk() (*[]XClusterNamespaceConfig, bool)`

GetDbDetailsOk returns a tuple with the DbDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDetails

`func (o *DrConfigGetResp) SetDbDetails(v []XClusterNamespaceConfig)`

SetDbDetails sets DbDetails field to given value.

### HasDbDetails

`func (o *DrConfigGetResp) HasDbDetails() bool`

HasDbDetails returns a boolean if a field has been set.

### GetDbs

`func (o *DrConfigGetResp) GetDbs() []string`

GetDbs returns the Dbs field if non-nil, zero value otherwise.

### GetDbsOk

`func (o *DrConfigGetResp) GetDbsOk() (*[]string, bool)`

GetDbsOk returns a tuple with the Dbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbs

`func (o *DrConfigGetResp) SetDbs(v []string)`

SetDbs sets Dbs field to given value.

### HasDbs

`func (o *DrConfigGetResp) HasDbs() bool`

HasDbs returns a boolean if a field has been set.

### GetDrReplicaUniverseActive

`func (o *DrConfigGetResp) GetDrReplicaUniverseActive() bool`

GetDrReplicaUniverseActive returns the DrReplicaUniverseActive field if non-nil, zero value otherwise.

### GetDrReplicaUniverseActiveOk

`func (o *DrConfigGetResp) GetDrReplicaUniverseActiveOk() (*bool, bool)`

GetDrReplicaUniverseActiveOk returns a tuple with the DrReplicaUniverseActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrReplicaUniverseActive

`func (o *DrConfigGetResp) SetDrReplicaUniverseActive(v bool)`

SetDrReplicaUniverseActive sets DrReplicaUniverseActive field to given value.

### HasDrReplicaUniverseActive

`func (o *DrConfigGetResp) HasDrReplicaUniverseActive() bool`

HasDrReplicaUniverseActive returns a boolean if a field has been set.

### GetDrReplicaUniverseState

`func (o *DrConfigGetResp) GetDrReplicaUniverseState() string`

GetDrReplicaUniverseState returns the DrReplicaUniverseState field if non-nil, zero value otherwise.

### GetDrReplicaUniverseStateOk

`func (o *DrConfigGetResp) GetDrReplicaUniverseStateOk() (*string, bool)`

GetDrReplicaUniverseStateOk returns a tuple with the DrReplicaUniverseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrReplicaUniverseState

`func (o *DrConfigGetResp) SetDrReplicaUniverseState(v string)`

SetDrReplicaUniverseState sets DrReplicaUniverseState field to given value.

### HasDrReplicaUniverseState

`func (o *DrConfigGetResp) HasDrReplicaUniverseState() bool`

HasDrReplicaUniverseState returns a boolean if a field has been set.

### GetDrReplicaUniverseUuid

`func (o *DrConfigGetResp) GetDrReplicaUniverseUuid() string`

GetDrReplicaUniverseUuid returns the DrReplicaUniverseUuid field if non-nil, zero value otherwise.

### GetDrReplicaUniverseUuidOk

`func (o *DrConfigGetResp) GetDrReplicaUniverseUuidOk() (*string, bool)`

GetDrReplicaUniverseUuidOk returns a tuple with the DrReplicaUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrReplicaUniverseUuid

`func (o *DrConfigGetResp) SetDrReplicaUniverseUuid(v string)`

SetDrReplicaUniverseUuid sets DrReplicaUniverseUuid field to given value.

### HasDrReplicaUniverseUuid

`func (o *DrConfigGetResp) HasDrReplicaUniverseUuid() bool`

HasDrReplicaUniverseUuid returns a boolean if a field has been set.

### GetKeyspacePending

`func (o *DrConfigGetResp) GetKeyspacePending() string`

GetKeyspacePending returns the KeyspacePending field if non-nil, zero value otherwise.

### GetKeyspacePendingOk

`func (o *DrConfigGetResp) GetKeyspacePendingOk() (*string, bool)`

GetKeyspacePendingOk returns a tuple with the KeyspacePending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspacePending

`func (o *DrConfigGetResp) SetKeyspacePending(v string)`

SetKeyspacePending sets KeyspacePending field to given value.

### HasKeyspacePending

`func (o *DrConfigGetResp) HasKeyspacePending() bool`

HasKeyspacePending returns a boolean if a field has been set.

### GetModifyTime

`func (o *DrConfigGetResp) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *DrConfigGetResp) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *DrConfigGetResp) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *DrConfigGetResp) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *DrConfigGetResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DrConfigGetResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DrConfigGetResp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DrConfigGetResp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaused

`func (o *DrConfigGetResp) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *DrConfigGetResp) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *DrConfigGetResp) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *DrConfigGetResp) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPitrConfigs

`func (o *DrConfigGetResp) GetPitrConfigs() []PitrConfig`

GetPitrConfigs returns the PitrConfigs field if non-nil, zero value otherwise.

### GetPitrConfigsOk

`func (o *DrConfigGetResp) GetPitrConfigsOk() (*[]PitrConfig, bool)`

GetPitrConfigsOk returns a tuple with the PitrConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrConfigs

`func (o *DrConfigGetResp) SetPitrConfigs(v []PitrConfig)`

SetPitrConfigs sets PitrConfigs field to given value.

### HasPitrConfigs

`func (o *DrConfigGetResp) HasPitrConfigs() bool`

HasPitrConfigs returns a boolean if a field has been set.

### GetPitrRetentionPeriodSec

`func (o *DrConfigGetResp) GetPitrRetentionPeriodSec() int64`

GetPitrRetentionPeriodSec returns the PitrRetentionPeriodSec field if non-nil, zero value otherwise.

### GetPitrRetentionPeriodSecOk

`func (o *DrConfigGetResp) GetPitrRetentionPeriodSecOk() (*int64, bool)`

GetPitrRetentionPeriodSecOk returns a tuple with the PitrRetentionPeriodSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrRetentionPeriodSec

`func (o *DrConfigGetResp) SetPitrRetentionPeriodSec(v int64)`

SetPitrRetentionPeriodSec sets PitrRetentionPeriodSec field to given value.

### HasPitrRetentionPeriodSec

`func (o *DrConfigGetResp) HasPitrRetentionPeriodSec() bool`

HasPitrRetentionPeriodSec returns a boolean if a field has been set.

### GetPitrSnapshotIntervalSec

`func (o *DrConfigGetResp) GetPitrSnapshotIntervalSec() int64`

GetPitrSnapshotIntervalSec returns the PitrSnapshotIntervalSec field if non-nil, zero value otherwise.

### GetPitrSnapshotIntervalSecOk

`func (o *DrConfigGetResp) GetPitrSnapshotIntervalSecOk() (*int64, bool)`

GetPitrSnapshotIntervalSecOk returns a tuple with the PitrSnapshotIntervalSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrSnapshotIntervalSec

`func (o *DrConfigGetResp) SetPitrSnapshotIntervalSec(v int64)`

SetPitrSnapshotIntervalSec sets PitrSnapshotIntervalSec field to given value.

### HasPitrSnapshotIntervalSec

`func (o *DrConfigGetResp) HasPitrSnapshotIntervalSec() bool`

HasPitrSnapshotIntervalSec returns a boolean if a field has been set.

### GetPrimaryUniverseActive

`func (o *DrConfigGetResp) GetPrimaryUniverseActive() bool`

GetPrimaryUniverseActive returns the PrimaryUniverseActive field if non-nil, zero value otherwise.

### GetPrimaryUniverseActiveOk

`func (o *DrConfigGetResp) GetPrimaryUniverseActiveOk() (*bool, bool)`

GetPrimaryUniverseActiveOk returns a tuple with the PrimaryUniverseActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUniverseActive

`func (o *DrConfigGetResp) SetPrimaryUniverseActive(v bool)`

SetPrimaryUniverseActive sets PrimaryUniverseActive field to given value.

### HasPrimaryUniverseActive

`func (o *DrConfigGetResp) HasPrimaryUniverseActive() bool`

HasPrimaryUniverseActive returns a boolean if a field has been set.

### GetPrimaryUniverseState

`func (o *DrConfigGetResp) GetPrimaryUniverseState() string`

GetPrimaryUniverseState returns the PrimaryUniverseState field if non-nil, zero value otherwise.

### GetPrimaryUniverseStateOk

`func (o *DrConfigGetResp) GetPrimaryUniverseStateOk() (*string, bool)`

GetPrimaryUniverseStateOk returns a tuple with the PrimaryUniverseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUniverseState

`func (o *DrConfigGetResp) SetPrimaryUniverseState(v string)`

SetPrimaryUniverseState sets PrimaryUniverseState field to given value.

### HasPrimaryUniverseState

`func (o *DrConfigGetResp) HasPrimaryUniverseState() bool`

HasPrimaryUniverseState returns a boolean if a field has been set.

### GetPrimaryUniverseUuid

`func (o *DrConfigGetResp) GetPrimaryUniverseUuid() string`

GetPrimaryUniverseUuid returns the PrimaryUniverseUuid field if non-nil, zero value otherwise.

### GetPrimaryUniverseUuidOk

`func (o *DrConfigGetResp) GetPrimaryUniverseUuidOk() (*string, bool)`

GetPrimaryUniverseUuidOk returns a tuple with the PrimaryUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUniverseUuid

`func (o *DrConfigGetResp) SetPrimaryUniverseUuid(v string)`

SetPrimaryUniverseUuid sets PrimaryUniverseUuid field to given value.

### HasPrimaryUniverseUuid

`func (o *DrConfigGetResp) HasPrimaryUniverseUuid() bool`

HasPrimaryUniverseUuid returns a boolean if a field has been set.

### GetReplicationGroupName

`func (o *DrConfigGetResp) GetReplicationGroupName() string`

GetReplicationGroupName returns the ReplicationGroupName field if non-nil, zero value otherwise.

### GetReplicationGroupNameOk

`func (o *DrConfigGetResp) GetReplicationGroupNameOk() (*string, bool)`

GetReplicationGroupNameOk returns a tuple with the ReplicationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroupName

`func (o *DrConfigGetResp) SetReplicationGroupName(v string)`

SetReplicationGroupName sets ReplicationGroupName field to given value.

### HasReplicationGroupName

`func (o *DrConfigGetResp) HasReplicationGroupName() bool`

HasReplicationGroupName returns a boolean if a field has been set.

### GetState

`func (o *DrConfigGetResp) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DrConfigGetResp) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DrConfigGetResp) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DrConfigGetResp) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *DrConfigGetResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DrConfigGetResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DrConfigGetResp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DrConfigGetResp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTableDetails

`func (o *DrConfigGetResp) GetTableDetails() []XClusterTableConfig`

GetTableDetails returns the TableDetails field if non-nil, zero value otherwise.

### GetTableDetailsOk

`func (o *DrConfigGetResp) GetTableDetailsOk() (*[]XClusterTableConfig, bool)`

GetTableDetailsOk returns a tuple with the TableDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableDetails

`func (o *DrConfigGetResp) SetTableDetails(v []XClusterTableConfig)`

SetTableDetails sets TableDetails field to given value.

### HasTableDetails

`func (o *DrConfigGetResp) HasTableDetails() bool`

HasTableDetails returns a boolean if a field has been set.

### GetTableType

`func (o *DrConfigGetResp) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *DrConfigGetResp) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *DrConfigGetResp) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *DrConfigGetResp) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetTables

`func (o *DrConfigGetResp) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *DrConfigGetResp) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *DrConfigGetResp) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *DrConfigGetResp) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetType

`func (o *DrConfigGetResp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DrConfigGetResp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DrConfigGetResp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DrConfigGetResp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *DrConfigGetResp) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DrConfigGetResp) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DrConfigGetResp) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DrConfigGetResp) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetXclusterConfigUuid

`func (o *DrConfigGetResp) GetXclusterConfigUuid() string`

GetXclusterConfigUuid returns the XclusterConfigUuid field if non-nil, zero value otherwise.

### GetXclusterConfigUuidOk

`func (o *DrConfigGetResp) GetXclusterConfigUuidOk() (*string, bool)`

GetXclusterConfigUuidOk returns a tuple with the XclusterConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterConfigUuid

`func (o *DrConfigGetResp) SetXclusterConfigUuid(v string)`

SetXclusterConfigUuid sets XclusterConfigUuid field to given value.

### HasXclusterConfigUuid

`func (o *DrConfigGetResp) HasXclusterConfigUuid() bool`

HasXclusterConfigUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


