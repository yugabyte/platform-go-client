# DrConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | DR config UUID. | [optional] [readonly] 
**State** | Pointer to **string** | High-level lifecycle state of the DR config. | [optional] [readonly] 
**PrimaryUniverseState** | Pointer to [**DrConfigUniverseReplicationState**](DrConfigUniverseReplicationState.md) |  | [optional] 
**DrReplicaUniverseState** | Pointer to [**DrConfigUniverseReplicationState**](DrConfigUniverseReplicationState.md) |  | [optional] 
**KeyspacePending** | Pointer to **string** | Keyspace name that the current DR task is working on, if any. | [optional] [readonly] 
**CreateTime** | Pointer to **time.Time** | Time when the DR config was created. | [optional] [readonly] 
**ModifyTime** | Pointer to **time.Time** | Time when the DR config was last modified. | [optional] [readonly] 
**ReplicationGroupName** | Pointer to **string** | Replication group name in the DR replica universe cluster config. | [optional] [readonly] 
**PrimaryUniverseActive** | Pointer to **bool** | Whether the primary universe is active in replication. | [optional] [readonly] 
**DrReplicaUniverseActive** | Pointer to **bool** | Whether the DR replica universe is active in replication. | [optional] [readonly] 

## Methods

### NewDrConfigInfo

`func NewDrConfigInfo() *DrConfigInfo`

NewDrConfigInfo instantiates a new DrConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigInfoWithDefaults

`func NewDrConfigInfoWithDefaults() *DrConfigInfo`

NewDrConfigInfoWithDefaults instantiates a new DrConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DrConfigInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DrConfigInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DrConfigInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DrConfigInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *DrConfigInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DrConfigInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DrConfigInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DrConfigInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPrimaryUniverseState

`func (o *DrConfigInfo) GetPrimaryUniverseState() DrConfigUniverseReplicationState`

GetPrimaryUniverseState returns the PrimaryUniverseState field if non-nil, zero value otherwise.

### GetPrimaryUniverseStateOk

`func (o *DrConfigInfo) GetPrimaryUniverseStateOk() (*DrConfigUniverseReplicationState, bool)`

GetPrimaryUniverseStateOk returns a tuple with the PrimaryUniverseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUniverseState

`func (o *DrConfigInfo) SetPrimaryUniverseState(v DrConfigUniverseReplicationState)`

SetPrimaryUniverseState sets PrimaryUniverseState field to given value.

### HasPrimaryUniverseState

`func (o *DrConfigInfo) HasPrimaryUniverseState() bool`

HasPrimaryUniverseState returns a boolean if a field has been set.

### GetDrReplicaUniverseState

`func (o *DrConfigInfo) GetDrReplicaUniverseState() DrConfigUniverseReplicationState`

GetDrReplicaUniverseState returns the DrReplicaUniverseState field if non-nil, zero value otherwise.

### GetDrReplicaUniverseStateOk

`func (o *DrConfigInfo) GetDrReplicaUniverseStateOk() (*DrConfigUniverseReplicationState, bool)`

GetDrReplicaUniverseStateOk returns a tuple with the DrReplicaUniverseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrReplicaUniverseState

`func (o *DrConfigInfo) SetDrReplicaUniverseState(v DrConfigUniverseReplicationState)`

SetDrReplicaUniverseState sets DrReplicaUniverseState field to given value.

### HasDrReplicaUniverseState

`func (o *DrConfigInfo) HasDrReplicaUniverseState() bool`

HasDrReplicaUniverseState returns a boolean if a field has been set.

### GetKeyspacePending

`func (o *DrConfigInfo) GetKeyspacePending() string`

GetKeyspacePending returns the KeyspacePending field if non-nil, zero value otherwise.

### GetKeyspacePendingOk

`func (o *DrConfigInfo) GetKeyspacePendingOk() (*string, bool)`

GetKeyspacePendingOk returns a tuple with the KeyspacePending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspacePending

`func (o *DrConfigInfo) SetKeyspacePending(v string)`

SetKeyspacePending sets KeyspacePending field to given value.

### HasKeyspacePending

`func (o *DrConfigInfo) HasKeyspacePending() bool`

HasKeyspacePending returns a boolean if a field has been set.

### GetCreateTime

`func (o *DrConfigInfo) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DrConfigInfo) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DrConfigInfo) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DrConfigInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetModifyTime

`func (o *DrConfigInfo) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *DrConfigInfo) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *DrConfigInfo) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *DrConfigInfo) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetReplicationGroupName

`func (o *DrConfigInfo) GetReplicationGroupName() string`

GetReplicationGroupName returns the ReplicationGroupName field if non-nil, zero value otherwise.

### GetReplicationGroupNameOk

`func (o *DrConfigInfo) GetReplicationGroupNameOk() (*string, bool)`

GetReplicationGroupNameOk returns a tuple with the ReplicationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroupName

`func (o *DrConfigInfo) SetReplicationGroupName(v string)`

SetReplicationGroupName sets ReplicationGroupName field to given value.

### HasReplicationGroupName

`func (o *DrConfigInfo) HasReplicationGroupName() bool`

HasReplicationGroupName returns a boolean if a field has been set.

### GetPrimaryUniverseActive

`func (o *DrConfigInfo) GetPrimaryUniverseActive() bool`

GetPrimaryUniverseActive returns the PrimaryUniverseActive field if non-nil, zero value otherwise.

### GetPrimaryUniverseActiveOk

`func (o *DrConfigInfo) GetPrimaryUniverseActiveOk() (*bool, bool)`

GetPrimaryUniverseActiveOk returns a tuple with the PrimaryUniverseActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUniverseActive

`func (o *DrConfigInfo) SetPrimaryUniverseActive(v bool)`

SetPrimaryUniverseActive sets PrimaryUniverseActive field to given value.

### HasPrimaryUniverseActive

`func (o *DrConfigInfo) HasPrimaryUniverseActive() bool`

HasPrimaryUniverseActive returns a boolean if a field has been set.

### GetDrReplicaUniverseActive

`func (o *DrConfigInfo) GetDrReplicaUniverseActive() bool`

GetDrReplicaUniverseActive returns the DrReplicaUniverseActive field if non-nil, zero value otherwise.

### GetDrReplicaUniverseActiveOk

`func (o *DrConfigInfo) GetDrReplicaUniverseActiveOk() (*bool, bool)`

GetDrReplicaUniverseActiveOk returns a tuple with the DrReplicaUniverseActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrReplicaUniverseActive

`func (o *DrConfigInfo) SetDrReplicaUniverseActive(v bool)`

SetDrReplicaUniverseActive sets DrReplicaUniverseActive field to given value.

### HasDrReplicaUniverseActive

`func (o *DrConfigInfo) HasDrReplicaUniverseActive() bool`

HasDrReplicaUniverseActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


