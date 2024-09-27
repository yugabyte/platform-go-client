# UniverseRollbackUpgradeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**RollingUpgrade** | Pointer to **bool** | Perform a rolling upgrade where only one node is upgraded at a time. This is the default behavior. False will perform a non-rolling upgrade where all nodes are upgraded at the same  | [optional] [default to true]
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 

## Methods

### NewUniverseRollbackUpgradeReq

`func NewUniverseRollbackUpgradeReq() *UniverseRollbackUpgradeReq`

NewUniverseRollbackUpgradeReq instantiates a new UniverseRollbackUpgradeReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseRollbackUpgradeReqWithDefaults

`func NewUniverseRollbackUpgradeReqWithDefaults() *UniverseRollbackUpgradeReq`

NewUniverseRollbackUpgradeReqWithDefaults instantiates a new UniverseRollbackUpgradeReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseRollbackUpgradeReq) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseRollbackUpgradeReq) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseRollbackUpgradeReq) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseRollbackUpgradeReq) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseRollbackUpgradeReq) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseRollbackUpgradeReq) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseRollbackUpgradeReq) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseRollbackUpgradeReq) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetRollingUpgrade

`func (o *UniverseRollbackUpgradeReq) GetRollingUpgrade() bool`

GetRollingUpgrade returns the RollingUpgrade field if non-nil, zero value otherwise.

### GetRollingUpgradeOk

`func (o *UniverseRollbackUpgradeReq) GetRollingUpgradeOk() (*bool, bool)`

GetRollingUpgradeOk returns a tuple with the RollingUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpgrade

`func (o *UniverseRollbackUpgradeReq) SetRollingUpgrade(v bool)`

SetRollingUpgrade sets RollingUpgrade field to given value.

### HasRollingUpgrade

`func (o *UniverseRollbackUpgradeReq) HasRollingUpgrade() bool`

HasRollingUpgrade returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseRollbackUpgradeReq) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseRollbackUpgradeReq) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseRollbackUpgradeReq) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseRollbackUpgradeReq) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


