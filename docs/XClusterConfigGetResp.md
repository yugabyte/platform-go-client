# XClusterConfigGetResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticDdlMode** | Pointer to **bool** |  | [optional] [readonly] 
**CreateTime** | Pointer to **time.Time** | Create time of the xCluster config | [optional] 
**Dbs** | Pointer to **[]string** |  | [optional] [readonly] 
**Imported** | Pointer to **bool** | YbaApi Internal. Whether this xCluster replication config was imported | [optional] 
**KeyspacePending** | Pointer to **string** | WARNING: This is a preview API that could change. The keyspace name that the xCluster task is working on; used for disaster recovery | [optional] 
**Lag** | **map[string]interface{}** | Lag metric data | 
**ModifyTime** | Pointer to **time.Time** | Last modify time of the xCluster config | [optional] 
**Name** | Pointer to **string** | XCluster config name | [optional] 
**NamespaceDetails** | Pointer to [**[]XClusterNamespaceConfig**](XClusterNamespaceConfig.md) | Namespaces participating in this xCluster config | [optional] [readonly] 
**Namespaces** | [**[]XClusterNamespaceConfig**](XClusterNamespaceConfig.md) |  | 
**Paused** | Pointer to **bool** | Whether this xCluster replication config is paused | [optional] 
**PitrConfigs** | Pointer to [**[]PitrConfig**](PitrConfig.md) | WARNING: This is a preview API that could change. The list of PITR configs used for the txn xCluster config | [optional] 
**ReplicationGroupName** | Pointer to **string** | Replication group name in the target universe cluster config | [optional] [readonly] 
**Secondary** | Pointer to **bool** | WARNING: This is a preview API that could change. Whether this xCluster config is used as a secondary config for a DR config | [optional] 
**SourceActive** | Pointer to **bool** | Whether the source is active in txn xCluster | [optional] 
**SourceUniverseState** | Pointer to **string** | WARNING: This is a preview API that could change. The replication status of the source universe; used for disaster recovery | [optional] 
**SourceUniverseUUID** | Pointer to **string** | Source Universe UUID | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**TableDetails** | Pointer to [**[]XClusterTableConfig**](XClusterTableConfig.md) | Tables participating in this xCluster config | [optional] 
**TableType** | Pointer to **string** | tableType | [optional] 
**Tables** | Pointer to **[]string** |  | [optional] [readonly] 
**TargetActive** | Pointer to **bool** | Whether the target is active in txn xCluster | [optional] 
**TargetUniverseState** | Pointer to **string** | WARNING: This is a preview API that could change. The replication status of the target universe; used for disaster recovery | [optional] 
**TargetUniverseUUID** | Pointer to **string** | Target Universe UUID | [optional] 
**Type** | Pointer to **string** | Whether the config is basic, txn, or db scoped xCluster | [optional] 
**UsedForDr** | Pointer to **bool** | WARNING: This is a preview API that could change. Whether the xCluster config is used as part of a DR config | [optional] [readonly] 
**Uuid** | Pointer to **string** | XCluster config UUID | [optional] 

## Methods

### NewXClusterConfigGetResp

`func NewXClusterConfigGetResp(lag map[string]interface{}, namespaces []XClusterNamespaceConfig, ) *XClusterConfigGetResp`

NewXClusterConfigGetResp instantiates a new XClusterConfigGetResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterConfigGetRespWithDefaults

`func NewXClusterConfigGetRespWithDefaults() *XClusterConfigGetResp`

NewXClusterConfigGetRespWithDefaults instantiates a new XClusterConfigGetResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticDdlMode

`func (o *XClusterConfigGetResp) GetAutomaticDdlMode() bool`

GetAutomaticDdlMode returns the AutomaticDdlMode field if non-nil, zero value otherwise.

### GetAutomaticDdlModeOk

`func (o *XClusterConfigGetResp) GetAutomaticDdlModeOk() (*bool, bool)`

GetAutomaticDdlModeOk returns a tuple with the AutomaticDdlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticDdlMode

`func (o *XClusterConfigGetResp) SetAutomaticDdlMode(v bool)`

SetAutomaticDdlMode sets AutomaticDdlMode field to given value.

### HasAutomaticDdlMode

`func (o *XClusterConfigGetResp) HasAutomaticDdlMode() bool`

HasAutomaticDdlMode returns a boolean if a field has been set.

### GetCreateTime

`func (o *XClusterConfigGetResp) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *XClusterConfigGetResp) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *XClusterConfigGetResp) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *XClusterConfigGetResp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDbs

`func (o *XClusterConfigGetResp) GetDbs() []string`

GetDbs returns the Dbs field if non-nil, zero value otherwise.

### GetDbsOk

`func (o *XClusterConfigGetResp) GetDbsOk() (*[]string, bool)`

GetDbsOk returns a tuple with the Dbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbs

`func (o *XClusterConfigGetResp) SetDbs(v []string)`

SetDbs sets Dbs field to given value.

### HasDbs

`func (o *XClusterConfigGetResp) HasDbs() bool`

HasDbs returns a boolean if a field has been set.

### GetImported

`func (o *XClusterConfigGetResp) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *XClusterConfigGetResp) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *XClusterConfigGetResp) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *XClusterConfigGetResp) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetKeyspacePending

`func (o *XClusterConfigGetResp) GetKeyspacePending() string`

GetKeyspacePending returns the KeyspacePending field if non-nil, zero value otherwise.

### GetKeyspacePendingOk

`func (o *XClusterConfigGetResp) GetKeyspacePendingOk() (*string, bool)`

GetKeyspacePendingOk returns a tuple with the KeyspacePending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspacePending

`func (o *XClusterConfigGetResp) SetKeyspacePending(v string)`

SetKeyspacePending sets KeyspacePending field to given value.

### HasKeyspacePending

`func (o *XClusterConfigGetResp) HasKeyspacePending() bool`

HasKeyspacePending returns a boolean if a field has been set.

### GetLag

`func (o *XClusterConfigGetResp) GetLag() map[string]interface{}`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *XClusterConfigGetResp) GetLagOk() (*map[string]interface{}, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *XClusterConfigGetResp) SetLag(v map[string]interface{})`

SetLag sets Lag field to given value.


### GetModifyTime

`func (o *XClusterConfigGetResp) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *XClusterConfigGetResp) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *XClusterConfigGetResp) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *XClusterConfigGetResp) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *XClusterConfigGetResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XClusterConfigGetResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XClusterConfigGetResp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *XClusterConfigGetResp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceDetails

`func (o *XClusterConfigGetResp) GetNamespaceDetails() []XClusterNamespaceConfig`

GetNamespaceDetails returns the NamespaceDetails field if non-nil, zero value otherwise.

### GetNamespaceDetailsOk

`func (o *XClusterConfigGetResp) GetNamespaceDetailsOk() (*[]XClusterNamespaceConfig, bool)`

GetNamespaceDetailsOk returns a tuple with the NamespaceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceDetails

`func (o *XClusterConfigGetResp) SetNamespaceDetails(v []XClusterNamespaceConfig)`

SetNamespaceDetails sets NamespaceDetails field to given value.

### HasNamespaceDetails

`func (o *XClusterConfigGetResp) HasNamespaceDetails() bool`

HasNamespaceDetails returns a boolean if a field has been set.

### GetNamespaces

`func (o *XClusterConfigGetResp) GetNamespaces() []XClusterNamespaceConfig`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *XClusterConfigGetResp) GetNamespacesOk() (*[]XClusterNamespaceConfig, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *XClusterConfigGetResp) SetNamespaces(v []XClusterNamespaceConfig)`

SetNamespaces sets Namespaces field to given value.


### GetPaused

`func (o *XClusterConfigGetResp) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *XClusterConfigGetResp) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *XClusterConfigGetResp) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *XClusterConfigGetResp) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPitrConfigs

`func (o *XClusterConfigGetResp) GetPitrConfigs() []PitrConfig`

GetPitrConfigs returns the PitrConfigs field if non-nil, zero value otherwise.

### GetPitrConfigsOk

`func (o *XClusterConfigGetResp) GetPitrConfigsOk() (*[]PitrConfig, bool)`

GetPitrConfigsOk returns a tuple with the PitrConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrConfigs

`func (o *XClusterConfigGetResp) SetPitrConfigs(v []PitrConfig)`

SetPitrConfigs sets PitrConfigs field to given value.

### HasPitrConfigs

`func (o *XClusterConfigGetResp) HasPitrConfigs() bool`

HasPitrConfigs returns a boolean if a field has been set.

### GetReplicationGroupName

`func (o *XClusterConfigGetResp) GetReplicationGroupName() string`

GetReplicationGroupName returns the ReplicationGroupName field if non-nil, zero value otherwise.

### GetReplicationGroupNameOk

`func (o *XClusterConfigGetResp) GetReplicationGroupNameOk() (*string, bool)`

GetReplicationGroupNameOk returns a tuple with the ReplicationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroupName

`func (o *XClusterConfigGetResp) SetReplicationGroupName(v string)`

SetReplicationGroupName sets ReplicationGroupName field to given value.

### HasReplicationGroupName

`func (o *XClusterConfigGetResp) HasReplicationGroupName() bool`

HasReplicationGroupName returns a boolean if a field has been set.

### GetSecondary

`func (o *XClusterConfigGetResp) GetSecondary() bool`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *XClusterConfigGetResp) GetSecondaryOk() (*bool, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *XClusterConfigGetResp) SetSecondary(v bool)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *XClusterConfigGetResp) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetSourceActive

`func (o *XClusterConfigGetResp) GetSourceActive() bool`

GetSourceActive returns the SourceActive field if non-nil, zero value otherwise.

### GetSourceActiveOk

`func (o *XClusterConfigGetResp) GetSourceActiveOk() (*bool, bool)`

GetSourceActiveOk returns a tuple with the SourceActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceActive

`func (o *XClusterConfigGetResp) SetSourceActive(v bool)`

SetSourceActive sets SourceActive field to given value.

### HasSourceActive

`func (o *XClusterConfigGetResp) HasSourceActive() bool`

HasSourceActive returns a boolean if a field has been set.

### GetSourceUniverseState

`func (o *XClusterConfigGetResp) GetSourceUniverseState() string`

GetSourceUniverseState returns the SourceUniverseState field if non-nil, zero value otherwise.

### GetSourceUniverseStateOk

`func (o *XClusterConfigGetResp) GetSourceUniverseStateOk() (*string, bool)`

GetSourceUniverseStateOk returns a tuple with the SourceUniverseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseState

`func (o *XClusterConfigGetResp) SetSourceUniverseState(v string)`

SetSourceUniverseState sets SourceUniverseState field to given value.

### HasSourceUniverseState

`func (o *XClusterConfigGetResp) HasSourceUniverseState() bool`

HasSourceUniverseState returns a boolean if a field has been set.

### GetSourceUniverseUUID

`func (o *XClusterConfigGetResp) GetSourceUniverseUUID() string`

GetSourceUniverseUUID returns the SourceUniverseUUID field if non-nil, zero value otherwise.

### GetSourceUniverseUUIDOk

`func (o *XClusterConfigGetResp) GetSourceUniverseUUIDOk() (*string, bool)`

GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseUUID

`func (o *XClusterConfigGetResp) SetSourceUniverseUUID(v string)`

SetSourceUniverseUUID sets SourceUniverseUUID field to given value.

### HasSourceUniverseUUID

`func (o *XClusterConfigGetResp) HasSourceUniverseUUID() bool`

HasSourceUniverseUUID returns a boolean if a field has been set.

### GetStatus

`func (o *XClusterConfigGetResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XClusterConfigGetResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XClusterConfigGetResp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XClusterConfigGetResp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTableDetails

`func (o *XClusterConfigGetResp) GetTableDetails() []XClusterTableConfig`

GetTableDetails returns the TableDetails field if non-nil, zero value otherwise.

### GetTableDetailsOk

`func (o *XClusterConfigGetResp) GetTableDetailsOk() (*[]XClusterTableConfig, bool)`

GetTableDetailsOk returns a tuple with the TableDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableDetails

`func (o *XClusterConfigGetResp) SetTableDetails(v []XClusterTableConfig)`

SetTableDetails sets TableDetails field to given value.

### HasTableDetails

`func (o *XClusterConfigGetResp) HasTableDetails() bool`

HasTableDetails returns a boolean if a field has been set.

### GetTableType

`func (o *XClusterConfigGetResp) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *XClusterConfigGetResp) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *XClusterConfigGetResp) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *XClusterConfigGetResp) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetTables

`func (o *XClusterConfigGetResp) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *XClusterConfigGetResp) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *XClusterConfigGetResp) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *XClusterConfigGetResp) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetTargetActive

`func (o *XClusterConfigGetResp) GetTargetActive() bool`

GetTargetActive returns the TargetActive field if non-nil, zero value otherwise.

### GetTargetActiveOk

`func (o *XClusterConfigGetResp) GetTargetActiveOk() (*bool, bool)`

GetTargetActiveOk returns a tuple with the TargetActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetActive

`func (o *XClusterConfigGetResp) SetTargetActive(v bool)`

SetTargetActive sets TargetActive field to given value.

### HasTargetActive

`func (o *XClusterConfigGetResp) HasTargetActive() bool`

HasTargetActive returns a boolean if a field has been set.

### GetTargetUniverseState

`func (o *XClusterConfigGetResp) GetTargetUniverseState() string`

GetTargetUniverseState returns the TargetUniverseState field if non-nil, zero value otherwise.

### GetTargetUniverseStateOk

`func (o *XClusterConfigGetResp) GetTargetUniverseStateOk() (*string, bool)`

GetTargetUniverseStateOk returns a tuple with the TargetUniverseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseState

`func (o *XClusterConfigGetResp) SetTargetUniverseState(v string)`

SetTargetUniverseState sets TargetUniverseState field to given value.

### HasTargetUniverseState

`func (o *XClusterConfigGetResp) HasTargetUniverseState() bool`

HasTargetUniverseState returns a boolean if a field has been set.

### GetTargetUniverseUUID

`func (o *XClusterConfigGetResp) GetTargetUniverseUUID() string`

GetTargetUniverseUUID returns the TargetUniverseUUID field if non-nil, zero value otherwise.

### GetTargetUniverseUUIDOk

`func (o *XClusterConfigGetResp) GetTargetUniverseUUIDOk() (*string, bool)`

GetTargetUniverseUUIDOk returns a tuple with the TargetUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseUUID

`func (o *XClusterConfigGetResp) SetTargetUniverseUUID(v string)`

SetTargetUniverseUUID sets TargetUniverseUUID field to given value.

### HasTargetUniverseUUID

`func (o *XClusterConfigGetResp) HasTargetUniverseUUID() bool`

HasTargetUniverseUUID returns a boolean if a field has been set.

### GetType

`func (o *XClusterConfigGetResp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *XClusterConfigGetResp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *XClusterConfigGetResp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *XClusterConfigGetResp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsedForDr

`func (o *XClusterConfigGetResp) GetUsedForDr() bool`

GetUsedForDr returns the UsedForDr field if non-nil, zero value otherwise.

### GetUsedForDrOk

`func (o *XClusterConfigGetResp) GetUsedForDrOk() (*bool, bool)`

GetUsedForDrOk returns a tuple with the UsedForDr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedForDr

`func (o *XClusterConfigGetResp) SetUsedForDr(v bool)`

SetUsedForDr sets UsedForDr field to given value.

### HasUsedForDr

`func (o *XClusterConfigGetResp) HasUsedForDr() bool`

HasUsedForDr returns a boolean if a field has been set.

### GetUuid

`func (o *XClusterConfigGetResp) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *XClusterConfigGetResp) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *XClusterConfigGetResp) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *XClusterConfigGetResp) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


