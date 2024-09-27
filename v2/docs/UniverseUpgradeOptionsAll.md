# UniverseUpgradeOptionsAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeOption** | Pointer to **string** | Option for an upgrade to be rolling (one node at a time) or non-rolling (all nodes at once, with downtime)  | [optional] [default to "Rolling"]
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 

## Methods

### NewUniverseUpgradeOptionsAll

`func NewUniverseUpgradeOptionsAll() *UniverseUpgradeOptionsAll`

NewUniverseUpgradeOptionsAll instantiates a new UniverseUpgradeOptionsAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseUpgradeOptionsAllWithDefaults

`func NewUniverseUpgradeOptionsAllWithDefaults() *UniverseUpgradeOptionsAll`

NewUniverseUpgradeOptionsAllWithDefaults instantiates a new UniverseUpgradeOptionsAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeOption

`func (o *UniverseUpgradeOptionsAll) GetUpgradeOption() string`

GetUpgradeOption returns the UpgradeOption field if non-nil, zero value otherwise.

### GetUpgradeOptionOk

`func (o *UniverseUpgradeOptionsAll) GetUpgradeOptionOk() (*string, bool)`

GetUpgradeOptionOk returns a tuple with the UpgradeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOption

`func (o *UniverseUpgradeOptionsAll) SetUpgradeOption(v string)`

SetUpgradeOption sets UpgradeOption field to given value.

### HasUpgradeOption

`func (o *UniverseUpgradeOptionsAll) HasUpgradeOption() bool`

HasUpgradeOption returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseUpgradeOptionsAll) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseUpgradeOptionsAll) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseUpgradeOptionsAll) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseUpgradeOptionsAll) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


