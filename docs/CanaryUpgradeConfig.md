# CanaryUpgradeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PauseAfterMasters** | Pointer to **bool** | WARNING: This is a preview API that could change. Whether to pause after all masters are upgraded | [optional] 
**PrimaryClusterAZSteps** | Pointer to [**[]AZUpgradeStep**](AZUpgradeStep.md) | WARNING: This is a preview API that could change. Primary cluster: ordered list of AZ upgrade steps. List order &#x3D; upgrade order. Each step can optionally pause after tserver upgrade. Null &#x3D; use default AZ order with no pause points. | [optional] 
**ReadReplicaClusterAZSteps** | Pointer to [**[]AZUpgradeStep**](AZUpgradeStep.md) | WARNING: This is a preview API that could change. Read replica cluster: ordered list of AZ upgrade steps. Only applicable if universe has read replica clusters. Null &#x3D; use default AZ order with no pause points. | [optional] 

## Methods

### NewCanaryUpgradeConfig

`func NewCanaryUpgradeConfig() *CanaryUpgradeConfig`

NewCanaryUpgradeConfig instantiates a new CanaryUpgradeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanaryUpgradeConfigWithDefaults

`func NewCanaryUpgradeConfigWithDefaults() *CanaryUpgradeConfig`

NewCanaryUpgradeConfigWithDefaults instantiates a new CanaryUpgradeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPauseAfterMasters

`func (o *CanaryUpgradeConfig) GetPauseAfterMasters() bool`

GetPauseAfterMasters returns the PauseAfterMasters field if non-nil, zero value otherwise.

### GetPauseAfterMastersOk

`func (o *CanaryUpgradeConfig) GetPauseAfterMastersOk() (*bool, bool)`

GetPauseAfterMastersOk returns a tuple with the PauseAfterMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseAfterMasters

`func (o *CanaryUpgradeConfig) SetPauseAfterMasters(v bool)`

SetPauseAfterMasters sets PauseAfterMasters field to given value.

### HasPauseAfterMasters

`func (o *CanaryUpgradeConfig) HasPauseAfterMasters() bool`

HasPauseAfterMasters returns a boolean if a field has been set.

### GetPrimaryClusterAZSteps

`func (o *CanaryUpgradeConfig) GetPrimaryClusterAZSteps() []AZUpgradeStep`

GetPrimaryClusterAZSteps returns the PrimaryClusterAZSteps field if non-nil, zero value otherwise.

### GetPrimaryClusterAZStepsOk

`func (o *CanaryUpgradeConfig) GetPrimaryClusterAZStepsOk() (*[]AZUpgradeStep, bool)`

GetPrimaryClusterAZStepsOk returns a tuple with the PrimaryClusterAZSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryClusterAZSteps

`func (o *CanaryUpgradeConfig) SetPrimaryClusterAZSteps(v []AZUpgradeStep)`

SetPrimaryClusterAZSteps sets PrimaryClusterAZSteps field to given value.

### HasPrimaryClusterAZSteps

`func (o *CanaryUpgradeConfig) HasPrimaryClusterAZSteps() bool`

HasPrimaryClusterAZSteps returns a boolean if a field has been set.

### GetReadReplicaClusterAZSteps

`func (o *CanaryUpgradeConfig) GetReadReplicaClusterAZSteps() []AZUpgradeStep`

GetReadReplicaClusterAZSteps returns the ReadReplicaClusterAZSteps field if non-nil, zero value otherwise.

### GetReadReplicaClusterAZStepsOk

`func (o *CanaryUpgradeConfig) GetReadReplicaClusterAZStepsOk() (*[]AZUpgradeStep, bool)`

GetReadReplicaClusterAZStepsOk returns a tuple with the ReadReplicaClusterAZSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadReplicaClusterAZSteps

`func (o *CanaryUpgradeConfig) SetReadReplicaClusterAZSteps(v []AZUpgradeStep)`

SetReadReplicaClusterAZSteps sets ReadReplicaClusterAZSteps field to given value.

### HasReadReplicaClusterAZSteps

`func (o *CanaryUpgradeConfig) HasReadReplicaClusterAZSteps() bool`

HasReadReplicaClusterAZSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


