# XClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **time.Time** | Create time of the xCluster config | [optional] 
**Imported** | Pointer to **bool** | Whether this xCluster replication config was imported | [optional] 
**ModifyTime** | Pointer to **time.Time** | Last modify time of the xCluster config | [optional] 
**Name** | Pointer to **string** | XCluster config name | [optional] 
**Paused** | Pointer to **bool** | Whether this xCluster replication config is paused | [optional] 
**ReplicationGroupName** | Pointer to **string** | Replication group name in DB | [optional] 
**SourceActive** | Pointer to **bool** | Whether the source is active in txn xCluster | [optional] 
**SourceUniverseUUID** | Pointer to **string** | Source Universe UUID | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**TableDetails** | Pointer to [**[]XClusterTableConfig**](XClusterTableConfig.md) | Tables participating in this xCluster config | [optional] 
**TableType** | Pointer to **string** | tableType | [optional] 
**Tables** | Pointer to **[]string** |  | [optional] [readonly] 
**TargetActive** | Pointer to **bool** | Whether the source is active in txn xCluster | [optional] 
**TargetUniverseUUID** | Pointer to **string** | Target Universe UUID | [optional] 
**TxnTableDetails** | Pointer to [**XClusterTableConfig**](XClusterTableConfig.md) |  | [optional] 
**TxnTableReplicationGroupName** | Pointer to **string** | Replication group name that replicates the transaction status table | [optional] 
**Type** | Pointer to **string** | Whether the config is txn xCluster | [optional] 
**Uuid** | Pointer to **string** | XCluster config UUID | [optional] 

## Methods

### NewXClusterConfig

`func NewXClusterConfig() *XClusterConfig`

NewXClusterConfig instantiates a new XClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterConfigWithDefaults

`func NewXClusterConfigWithDefaults() *XClusterConfig`

NewXClusterConfigWithDefaults instantiates a new XClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *XClusterConfig) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *XClusterConfig) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *XClusterConfig) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *XClusterConfig) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetImported

`func (o *XClusterConfig) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *XClusterConfig) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *XClusterConfig) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *XClusterConfig) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetModifyTime

`func (o *XClusterConfig) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *XClusterConfig) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *XClusterConfig) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *XClusterConfig) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *XClusterConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XClusterConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XClusterConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *XClusterConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaused

`func (o *XClusterConfig) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *XClusterConfig) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *XClusterConfig) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *XClusterConfig) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetReplicationGroupName

`func (o *XClusterConfig) GetReplicationGroupName() string`

GetReplicationGroupName returns the ReplicationGroupName field if non-nil, zero value otherwise.

### GetReplicationGroupNameOk

`func (o *XClusterConfig) GetReplicationGroupNameOk() (*string, bool)`

GetReplicationGroupNameOk returns a tuple with the ReplicationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroupName

`func (o *XClusterConfig) SetReplicationGroupName(v string)`

SetReplicationGroupName sets ReplicationGroupName field to given value.

### HasReplicationGroupName

`func (o *XClusterConfig) HasReplicationGroupName() bool`

HasReplicationGroupName returns a boolean if a field has been set.

### GetSourceActive

`func (o *XClusterConfig) GetSourceActive() bool`

GetSourceActive returns the SourceActive field if non-nil, zero value otherwise.

### GetSourceActiveOk

`func (o *XClusterConfig) GetSourceActiveOk() (*bool, bool)`

GetSourceActiveOk returns a tuple with the SourceActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceActive

`func (o *XClusterConfig) SetSourceActive(v bool)`

SetSourceActive sets SourceActive field to given value.

### HasSourceActive

`func (o *XClusterConfig) HasSourceActive() bool`

HasSourceActive returns a boolean if a field has been set.

### GetSourceUniverseUUID

`func (o *XClusterConfig) GetSourceUniverseUUID() string`

GetSourceUniverseUUID returns the SourceUniverseUUID field if non-nil, zero value otherwise.

### GetSourceUniverseUUIDOk

`func (o *XClusterConfig) GetSourceUniverseUUIDOk() (*string, bool)`

GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseUUID

`func (o *XClusterConfig) SetSourceUniverseUUID(v string)`

SetSourceUniverseUUID sets SourceUniverseUUID field to given value.

### HasSourceUniverseUUID

`func (o *XClusterConfig) HasSourceUniverseUUID() bool`

HasSourceUniverseUUID returns a boolean if a field has been set.

### GetStatus

`func (o *XClusterConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XClusterConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XClusterConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XClusterConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTableDetails

`func (o *XClusterConfig) GetTableDetails() []XClusterTableConfig`

GetTableDetails returns the TableDetails field if non-nil, zero value otherwise.

### GetTableDetailsOk

`func (o *XClusterConfig) GetTableDetailsOk() (*[]XClusterTableConfig, bool)`

GetTableDetailsOk returns a tuple with the TableDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableDetails

`func (o *XClusterConfig) SetTableDetails(v []XClusterTableConfig)`

SetTableDetails sets TableDetails field to given value.

### HasTableDetails

`func (o *XClusterConfig) HasTableDetails() bool`

HasTableDetails returns a boolean if a field has been set.

### GetTableType

`func (o *XClusterConfig) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *XClusterConfig) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *XClusterConfig) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *XClusterConfig) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetTables

`func (o *XClusterConfig) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *XClusterConfig) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *XClusterConfig) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *XClusterConfig) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetTargetActive

`func (o *XClusterConfig) GetTargetActive() bool`

GetTargetActive returns the TargetActive field if non-nil, zero value otherwise.

### GetTargetActiveOk

`func (o *XClusterConfig) GetTargetActiveOk() (*bool, bool)`

GetTargetActiveOk returns a tuple with the TargetActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetActive

`func (o *XClusterConfig) SetTargetActive(v bool)`

SetTargetActive sets TargetActive field to given value.

### HasTargetActive

`func (o *XClusterConfig) HasTargetActive() bool`

HasTargetActive returns a boolean if a field has been set.

### GetTargetUniverseUUID

`func (o *XClusterConfig) GetTargetUniverseUUID() string`

GetTargetUniverseUUID returns the TargetUniverseUUID field if non-nil, zero value otherwise.

### GetTargetUniverseUUIDOk

`func (o *XClusterConfig) GetTargetUniverseUUIDOk() (*string, bool)`

GetTargetUniverseUUIDOk returns a tuple with the TargetUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseUUID

`func (o *XClusterConfig) SetTargetUniverseUUID(v string)`

SetTargetUniverseUUID sets TargetUniverseUUID field to given value.

### HasTargetUniverseUUID

`func (o *XClusterConfig) HasTargetUniverseUUID() bool`

HasTargetUniverseUUID returns a boolean if a field has been set.

### GetTxnTableDetails

`func (o *XClusterConfig) GetTxnTableDetails() XClusterTableConfig`

GetTxnTableDetails returns the TxnTableDetails field if non-nil, zero value otherwise.

### GetTxnTableDetailsOk

`func (o *XClusterConfig) GetTxnTableDetailsOk() (*XClusterTableConfig, bool)`

GetTxnTableDetailsOk returns a tuple with the TxnTableDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnTableDetails

`func (o *XClusterConfig) SetTxnTableDetails(v XClusterTableConfig)`

SetTxnTableDetails sets TxnTableDetails field to given value.

### HasTxnTableDetails

`func (o *XClusterConfig) HasTxnTableDetails() bool`

HasTxnTableDetails returns a boolean if a field has been set.

### GetTxnTableReplicationGroupName

`func (o *XClusterConfig) GetTxnTableReplicationGroupName() string`

GetTxnTableReplicationGroupName returns the TxnTableReplicationGroupName field if non-nil, zero value otherwise.

### GetTxnTableReplicationGroupNameOk

`func (o *XClusterConfig) GetTxnTableReplicationGroupNameOk() (*string, bool)`

GetTxnTableReplicationGroupNameOk returns a tuple with the TxnTableReplicationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnTableReplicationGroupName

`func (o *XClusterConfig) SetTxnTableReplicationGroupName(v string)`

SetTxnTableReplicationGroupName sets TxnTableReplicationGroupName field to given value.

### HasTxnTableReplicationGroupName

`func (o *XClusterConfig) HasTxnTableReplicationGroupName() bool`

HasTxnTableReplicationGroupName returns a boolean if a field has been set.

### GetType

`func (o *XClusterConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *XClusterConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *XClusterConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *XClusterConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *XClusterConfig) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *XClusterConfig) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *XClusterConfig) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *XClusterConfig) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


