# DrConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Disaster recovery config name. | 
**PrimaryUniverseUuid** | **string** | UUID of the primary (source) universe. | 
**DrReplicaUniverseUuid** | **string** | UUID of the DR replica (target) universe. | 
**TableType** | [**XClusterTableType**](XClusterTableType.md) |  | 
**Type** | **string** | DR config type. | 
**BootstrapParams** | Pointer to [**DrConfigBootstrapParams**](DrConfigBootstrapParams.md) |  | [optional] 
**PitrRetentionPeriodSec** | Pointer to **int64** | PITR retention period in seconds for the DR config. | [optional] 
**PitrSnapshotIntervalSec** | Pointer to **int64** | PITR snapshot interval in seconds for the DR config. | [optional] 
**Tables** | Pointer to **[]string** | Table IDs included in replication for transaction-scoped (Txn) DR configs. | [optional] 
**Dbs** | Pointer to **[]string** | Database (namespace) IDs included in replication for DB-scoped DR configs. | [optional] 
**Webhooks** | Pointer to [**[]DrConfigWebhook**](DrConfigWebhook.md) | Webhooks configured for this DR config. | [optional] 

## Methods

### NewDrConfigSpec

`func NewDrConfigSpec(name string, primaryUniverseUuid string, drReplicaUniverseUuid string, tableType XClusterTableType, type_ string, ) *DrConfigSpec`

NewDrConfigSpec instantiates a new DrConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigSpecWithDefaults

`func NewDrConfigSpecWithDefaults() *DrConfigSpec`

NewDrConfigSpecWithDefaults instantiates a new DrConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DrConfigSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DrConfigSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DrConfigSpec) SetName(v string)`

SetName sets Name field to given value.


### GetPrimaryUniverseUuid

`func (o *DrConfigSpec) GetPrimaryUniverseUuid() string`

GetPrimaryUniverseUuid returns the PrimaryUniverseUuid field if non-nil, zero value otherwise.

### GetPrimaryUniverseUuidOk

`func (o *DrConfigSpec) GetPrimaryUniverseUuidOk() (*string, bool)`

GetPrimaryUniverseUuidOk returns a tuple with the PrimaryUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUniverseUuid

`func (o *DrConfigSpec) SetPrimaryUniverseUuid(v string)`

SetPrimaryUniverseUuid sets PrimaryUniverseUuid field to given value.


### GetDrReplicaUniverseUuid

`func (o *DrConfigSpec) GetDrReplicaUniverseUuid() string`

GetDrReplicaUniverseUuid returns the DrReplicaUniverseUuid field if non-nil, zero value otherwise.

### GetDrReplicaUniverseUuidOk

`func (o *DrConfigSpec) GetDrReplicaUniverseUuidOk() (*string, bool)`

GetDrReplicaUniverseUuidOk returns a tuple with the DrReplicaUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrReplicaUniverseUuid

`func (o *DrConfigSpec) SetDrReplicaUniverseUuid(v string)`

SetDrReplicaUniverseUuid sets DrReplicaUniverseUuid field to given value.


### GetTableType

`func (o *DrConfigSpec) GetTableType() XClusterTableType`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *DrConfigSpec) GetTableTypeOk() (*XClusterTableType, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *DrConfigSpec) SetTableType(v XClusterTableType)`

SetTableType sets TableType field to given value.


### GetType

`func (o *DrConfigSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DrConfigSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DrConfigSpec) SetType(v string)`

SetType sets Type field to given value.


### GetBootstrapParams

`func (o *DrConfigSpec) GetBootstrapParams() DrConfigBootstrapParams`

GetBootstrapParams returns the BootstrapParams field if non-nil, zero value otherwise.

### GetBootstrapParamsOk

`func (o *DrConfigSpec) GetBootstrapParamsOk() (*DrConfigBootstrapParams, bool)`

GetBootstrapParamsOk returns a tuple with the BootstrapParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapParams

`func (o *DrConfigSpec) SetBootstrapParams(v DrConfigBootstrapParams)`

SetBootstrapParams sets BootstrapParams field to given value.

### HasBootstrapParams

`func (o *DrConfigSpec) HasBootstrapParams() bool`

HasBootstrapParams returns a boolean if a field has been set.

### GetPitrRetentionPeriodSec

`func (o *DrConfigSpec) GetPitrRetentionPeriodSec() int64`

GetPitrRetentionPeriodSec returns the PitrRetentionPeriodSec field if non-nil, zero value otherwise.

### GetPitrRetentionPeriodSecOk

`func (o *DrConfigSpec) GetPitrRetentionPeriodSecOk() (*int64, bool)`

GetPitrRetentionPeriodSecOk returns a tuple with the PitrRetentionPeriodSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrRetentionPeriodSec

`func (o *DrConfigSpec) SetPitrRetentionPeriodSec(v int64)`

SetPitrRetentionPeriodSec sets PitrRetentionPeriodSec field to given value.

### HasPitrRetentionPeriodSec

`func (o *DrConfigSpec) HasPitrRetentionPeriodSec() bool`

HasPitrRetentionPeriodSec returns a boolean if a field has been set.

### GetPitrSnapshotIntervalSec

`func (o *DrConfigSpec) GetPitrSnapshotIntervalSec() int64`

GetPitrSnapshotIntervalSec returns the PitrSnapshotIntervalSec field if non-nil, zero value otherwise.

### GetPitrSnapshotIntervalSecOk

`func (o *DrConfigSpec) GetPitrSnapshotIntervalSecOk() (*int64, bool)`

GetPitrSnapshotIntervalSecOk returns a tuple with the PitrSnapshotIntervalSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrSnapshotIntervalSec

`func (o *DrConfigSpec) SetPitrSnapshotIntervalSec(v int64)`

SetPitrSnapshotIntervalSec sets PitrSnapshotIntervalSec field to given value.

### HasPitrSnapshotIntervalSec

`func (o *DrConfigSpec) HasPitrSnapshotIntervalSec() bool`

HasPitrSnapshotIntervalSec returns a boolean if a field has been set.

### GetTables

`func (o *DrConfigSpec) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *DrConfigSpec) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *DrConfigSpec) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *DrConfigSpec) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetDbs

`func (o *DrConfigSpec) GetDbs() []string`

GetDbs returns the Dbs field if non-nil, zero value otherwise.

### GetDbsOk

`func (o *DrConfigSpec) GetDbsOk() (*[]string, bool)`

GetDbsOk returns a tuple with the Dbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbs

`func (o *DrConfigSpec) SetDbs(v []string)`

SetDbs sets Dbs field to given value.

### HasDbs

`func (o *DrConfigSpec) HasDbs() bool`

HasDbs returns a boolean if a field has been set.

### GetWebhooks

`func (o *DrConfigSpec) GetWebhooks() []DrConfigWebhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *DrConfigSpec) GetWebhooksOk() (*[]DrConfigWebhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *DrConfigSpec) SetWebhooks(v []DrConfigWebhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *DrConfigSpec) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


