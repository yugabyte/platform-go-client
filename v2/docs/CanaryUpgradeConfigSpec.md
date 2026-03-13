# CanaryUpgradeConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PauseAfterMasters** | Pointer to **bool** | Whether to pause after all masters are upgraded | [optional] [default to false]
**PrimaryClusterAzSteps** | Pointer to [**[]SoftwareUpgradeAZStep**](SoftwareUpgradeAZStep.md) | Primary cluster: ordered list of AZ upgrade steps. List order &#x3D; upgrade order. Each step can optionally pause after tserver upgrade. Null &#x3D; use default AZ order.  | [optional] 
**ReadReplicaClusterAzSteps** | Pointer to [**[]SoftwareUpgradeAZStep**](SoftwareUpgradeAZStep.md) | Read replica cluster: ordered list of AZ upgrade steps. Only applicable if universe has read replica clusters. Null &#x3D; use default AZ order.  | [optional] 

## Methods

### NewCanaryUpgradeConfigSpec

`func NewCanaryUpgradeConfigSpec() *CanaryUpgradeConfigSpec`

NewCanaryUpgradeConfigSpec instantiates a new CanaryUpgradeConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanaryUpgradeConfigSpecWithDefaults

`func NewCanaryUpgradeConfigSpecWithDefaults() *CanaryUpgradeConfigSpec`

NewCanaryUpgradeConfigSpecWithDefaults instantiates a new CanaryUpgradeConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPauseAfterMasters

`func (o *CanaryUpgradeConfigSpec) GetPauseAfterMasters() bool`

GetPauseAfterMasters returns the PauseAfterMasters field if non-nil, zero value otherwise.

### GetPauseAfterMastersOk

`func (o *CanaryUpgradeConfigSpec) GetPauseAfterMastersOk() (*bool, bool)`

GetPauseAfterMastersOk returns a tuple with the PauseAfterMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseAfterMasters

`func (o *CanaryUpgradeConfigSpec) SetPauseAfterMasters(v bool)`

SetPauseAfterMasters sets PauseAfterMasters field to given value.

### HasPauseAfterMasters

`func (o *CanaryUpgradeConfigSpec) HasPauseAfterMasters() bool`

HasPauseAfterMasters returns a boolean if a field has been set.

### GetPrimaryClusterAzSteps

`func (o *CanaryUpgradeConfigSpec) GetPrimaryClusterAzSteps() []SoftwareUpgradeAZStep`

GetPrimaryClusterAzSteps returns the PrimaryClusterAzSteps field if non-nil, zero value otherwise.

### GetPrimaryClusterAzStepsOk

`func (o *CanaryUpgradeConfigSpec) GetPrimaryClusterAzStepsOk() (*[]SoftwareUpgradeAZStep, bool)`

GetPrimaryClusterAzStepsOk returns a tuple with the PrimaryClusterAzSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryClusterAzSteps

`func (o *CanaryUpgradeConfigSpec) SetPrimaryClusterAzSteps(v []SoftwareUpgradeAZStep)`

SetPrimaryClusterAzSteps sets PrimaryClusterAzSteps field to given value.

### HasPrimaryClusterAzSteps

`func (o *CanaryUpgradeConfigSpec) HasPrimaryClusterAzSteps() bool`

HasPrimaryClusterAzSteps returns a boolean if a field has been set.

### GetReadReplicaClusterAzSteps

`func (o *CanaryUpgradeConfigSpec) GetReadReplicaClusterAzSteps() []SoftwareUpgradeAZStep`

GetReadReplicaClusterAzSteps returns the ReadReplicaClusterAzSteps field if non-nil, zero value otherwise.

### GetReadReplicaClusterAzStepsOk

`func (o *CanaryUpgradeConfigSpec) GetReadReplicaClusterAzStepsOk() (*[]SoftwareUpgradeAZStep, bool)`

GetReadReplicaClusterAzStepsOk returns a tuple with the ReadReplicaClusterAzSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadReplicaClusterAzSteps

`func (o *CanaryUpgradeConfigSpec) SetReadReplicaClusterAzSteps(v []SoftwareUpgradeAZStep)`

SetReadReplicaClusterAzSteps sets ReadReplicaClusterAzSteps field to given value.

### HasReadReplicaClusterAzSteps

`func (o *CanaryUpgradeConfigSpec) HasReadReplicaClusterAzSteps() bool`

HasReadReplicaClusterAzSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


