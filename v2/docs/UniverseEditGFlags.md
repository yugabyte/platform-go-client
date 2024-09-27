# UniverseEditGFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**UpgradeOption** | Pointer to **string** | Option for an upgrade to be rolling (one node at a time) or non-rolling (all nodes at once, with downtime)  | [optional] [default to "Rolling"]
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 
**KubernetesResourceDetails** | Pointer to [**KubernetesResourceDetails**](KubernetesResourceDetails.md) |  | [optional] 
**UniverseGflags** | Pointer to [**map[string]ClusterGFlags**](ClusterGFlags.md) | GFlags for each cluster uuid of this universe | [optional] 

## Methods

### NewUniverseEditGFlags

`func NewUniverseEditGFlags() *UniverseEditGFlags`

NewUniverseEditGFlags instantiates a new UniverseEditGFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseEditGFlagsWithDefaults

`func NewUniverseEditGFlagsWithDefaults() *UniverseEditGFlags`

NewUniverseEditGFlagsWithDefaults instantiates a new UniverseEditGFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseEditGFlags) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseEditGFlags) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseEditGFlags) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseEditGFlags) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseEditGFlags) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseEditGFlags) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseEditGFlags) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseEditGFlags) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetUpgradeOption

`func (o *UniverseEditGFlags) GetUpgradeOption() string`

GetUpgradeOption returns the UpgradeOption field if non-nil, zero value otherwise.

### GetUpgradeOptionOk

`func (o *UniverseEditGFlags) GetUpgradeOptionOk() (*string, bool)`

GetUpgradeOptionOk returns a tuple with the UpgradeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOption

`func (o *UniverseEditGFlags) SetUpgradeOption(v string)`

SetUpgradeOption sets UpgradeOption field to given value.

### HasUpgradeOption

`func (o *UniverseEditGFlags) HasUpgradeOption() bool`

HasUpgradeOption returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseEditGFlags) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseEditGFlags) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseEditGFlags) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseEditGFlags) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetKubernetesResourceDetails

`func (o *UniverseEditGFlags) GetKubernetesResourceDetails() KubernetesResourceDetails`

GetKubernetesResourceDetails returns the KubernetesResourceDetails field if non-nil, zero value otherwise.

### GetKubernetesResourceDetailsOk

`func (o *UniverseEditGFlags) GetKubernetesResourceDetailsOk() (*KubernetesResourceDetails, bool)`

GetKubernetesResourceDetailsOk returns a tuple with the KubernetesResourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesResourceDetails

`func (o *UniverseEditGFlags) SetKubernetesResourceDetails(v KubernetesResourceDetails)`

SetKubernetesResourceDetails sets KubernetesResourceDetails field to given value.

### HasKubernetesResourceDetails

`func (o *UniverseEditGFlags) HasKubernetesResourceDetails() bool`

HasKubernetesResourceDetails returns a boolean if a field has been set.

### GetUniverseGflags

`func (o *UniverseEditGFlags) GetUniverseGflags() map[string]ClusterGFlags`

GetUniverseGflags returns the UniverseGflags field if non-nil, zero value otherwise.

### GetUniverseGflagsOk

`func (o *UniverseEditGFlags) GetUniverseGflagsOk() (*map[string]ClusterGFlags, bool)`

GetUniverseGflagsOk returns a tuple with the UniverseGflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseGflags

`func (o *UniverseEditGFlags) SetUniverseGflags(v map[string]ClusterGFlags)`

SetUniverseGflags sets UniverseGflags field to given value.

### HasUniverseGflags

`func (o *UniverseEditGFlags) HasUniverseGflags() bool`

HasUniverseGflags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


