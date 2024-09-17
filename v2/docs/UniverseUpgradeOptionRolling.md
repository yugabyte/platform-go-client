# UniverseUpgradeOptionRolling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingUpgrade** | Pointer to **bool** | Perform a rolling upgrade where only one node is upgraded at a time. This is the default behavior. False will perform a non-rolling upgrade where all nodes are upgraded at the same  | [optional] [default to true]
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 

## Methods

### NewUniverseUpgradeOptionRolling

`func NewUniverseUpgradeOptionRolling() *UniverseUpgradeOptionRolling`

NewUniverseUpgradeOptionRolling instantiates a new UniverseUpgradeOptionRolling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseUpgradeOptionRollingWithDefaults

`func NewUniverseUpgradeOptionRollingWithDefaults() *UniverseUpgradeOptionRolling`

NewUniverseUpgradeOptionRollingWithDefaults instantiates a new UniverseUpgradeOptionRolling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingUpgrade

`func (o *UniverseUpgradeOptionRolling) GetRollingUpgrade() bool`

GetRollingUpgrade returns the RollingUpgrade field if non-nil, zero value otherwise.

### GetRollingUpgradeOk

`func (o *UniverseUpgradeOptionRolling) GetRollingUpgradeOk() (*bool, bool)`

GetRollingUpgradeOk returns a tuple with the RollingUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpgrade

`func (o *UniverseUpgradeOptionRolling) SetRollingUpgrade(v bool)`

SetRollingUpgrade sets RollingUpgrade field to given value.

### HasRollingUpgrade

`func (o *UniverseUpgradeOptionRolling) HasRollingUpgrade() bool`

HasRollingUpgrade returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseUpgradeOptionRolling) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseUpgradeOptionRolling) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseUpgradeOptionRolling) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseUpgradeOptionRolling) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


