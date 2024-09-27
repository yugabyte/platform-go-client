# UniverseEditEncryptionInTransit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**RollingUpgrade** | Pointer to **bool** | Perform a rolling upgrade where only one node is upgraded at a time. This is the default behavior. False will perform a non-rolling upgrade where all nodes are upgraded at the same  | [optional] [default to true]
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 
**NodeToNode** | Pointer to **bool** | Control encryption in transit between YugabyteDB nodes | [optional] 
**ClientToNode** | Pointer to **bool** | Control encryption in transit between YugabyteDB nodes and clients | [optional] 
**RootCa** | Pointer to **string** | Root CA cert for node to node encryption. Required if node_to_node is true | [optional] 
**ClientRootCa** | Pointer to **string** | Root CA used for node to client encryption. Required if client_to_node is true | [optional] 

## Methods

### NewUniverseEditEncryptionInTransit

`func NewUniverseEditEncryptionInTransit() *UniverseEditEncryptionInTransit`

NewUniverseEditEncryptionInTransit instantiates a new UniverseEditEncryptionInTransit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseEditEncryptionInTransitWithDefaults

`func NewUniverseEditEncryptionInTransitWithDefaults() *UniverseEditEncryptionInTransit`

NewUniverseEditEncryptionInTransitWithDefaults instantiates a new UniverseEditEncryptionInTransit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseEditEncryptionInTransit) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseEditEncryptionInTransit) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseEditEncryptionInTransit) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseEditEncryptionInTransit) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseEditEncryptionInTransit) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseEditEncryptionInTransit) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseEditEncryptionInTransit) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseEditEncryptionInTransit) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetRollingUpgrade

`func (o *UniverseEditEncryptionInTransit) GetRollingUpgrade() bool`

GetRollingUpgrade returns the RollingUpgrade field if non-nil, zero value otherwise.

### GetRollingUpgradeOk

`func (o *UniverseEditEncryptionInTransit) GetRollingUpgradeOk() (*bool, bool)`

GetRollingUpgradeOk returns a tuple with the RollingUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpgrade

`func (o *UniverseEditEncryptionInTransit) SetRollingUpgrade(v bool)`

SetRollingUpgrade sets RollingUpgrade field to given value.

### HasRollingUpgrade

`func (o *UniverseEditEncryptionInTransit) HasRollingUpgrade() bool`

HasRollingUpgrade returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseEditEncryptionInTransit) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseEditEncryptionInTransit) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseEditEncryptionInTransit) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseEditEncryptionInTransit) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetNodeToNode

`func (o *UniverseEditEncryptionInTransit) GetNodeToNode() bool`

GetNodeToNode returns the NodeToNode field if non-nil, zero value otherwise.

### GetNodeToNodeOk

`func (o *UniverseEditEncryptionInTransit) GetNodeToNodeOk() (*bool, bool)`

GetNodeToNodeOk returns a tuple with the NodeToNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeToNode

`func (o *UniverseEditEncryptionInTransit) SetNodeToNode(v bool)`

SetNodeToNode sets NodeToNode field to given value.

### HasNodeToNode

`func (o *UniverseEditEncryptionInTransit) HasNodeToNode() bool`

HasNodeToNode returns a boolean if a field has been set.

### GetClientToNode

`func (o *UniverseEditEncryptionInTransit) GetClientToNode() bool`

GetClientToNode returns the ClientToNode field if non-nil, zero value otherwise.

### GetClientToNodeOk

`func (o *UniverseEditEncryptionInTransit) GetClientToNodeOk() (*bool, bool)`

GetClientToNodeOk returns a tuple with the ClientToNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToNode

`func (o *UniverseEditEncryptionInTransit) SetClientToNode(v bool)`

SetClientToNode sets ClientToNode field to given value.

### HasClientToNode

`func (o *UniverseEditEncryptionInTransit) HasClientToNode() bool`

HasClientToNode returns a boolean if a field has been set.

### GetRootCa

`func (o *UniverseEditEncryptionInTransit) GetRootCa() string`

GetRootCa returns the RootCa field if non-nil, zero value otherwise.

### GetRootCaOk

`func (o *UniverseEditEncryptionInTransit) GetRootCaOk() (*string, bool)`

GetRootCaOk returns a tuple with the RootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCa

`func (o *UniverseEditEncryptionInTransit) SetRootCa(v string)`

SetRootCa sets RootCa field to given value.

### HasRootCa

`func (o *UniverseEditEncryptionInTransit) HasRootCa() bool`

HasRootCa returns a boolean if a field has been set.

### GetClientRootCa

`func (o *UniverseEditEncryptionInTransit) GetClientRootCa() string`

GetClientRootCa returns the ClientRootCa field if non-nil, zero value otherwise.

### GetClientRootCaOk

`func (o *UniverseEditEncryptionInTransit) GetClientRootCaOk() (*string, bool)`

GetClientRootCaOk returns a tuple with the ClientRootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCa

`func (o *UniverseEditEncryptionInTransit) SetClientRootCa(v string)`

SetClientRootCa sets ClientRootCa field to given value.

### HasClientRootCa

`func (o *UniverseEditEncryptionInTransit) HasClientRootCa() bool`

HasClientRootCa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


