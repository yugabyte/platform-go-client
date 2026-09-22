# UniverseVMImageUpgradeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. If unset, runtime config is used. | [optional] 
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. If unset, runtime config is used. | [optional] 
**ImageBundles** | Pointer to [**[]VMImageClusterSpec**](VMImageClusterSpec.md) |  | [optional] 
**ForceUpgrade** | Pointer to **bool** | Whether to skip checking current OS for each node. | [optional] 
**YbSoftwareVersion** | Pointer to **string** | The YugabyteDB software version to install during os upgrade. | [optional] 

## Methods

### NewUniverseVMImageUpgradeSpec

`func NewUniverseVMImageUpgradeSpec() *UniverseVMImageUpgradeSpec`

NewUniverseVMImageUpgradeSpec instantiates a new UniverseVMImageUpgradeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseVMImageUpgradeSpecWithDefaults

`func NewUniverseVMImageUpgradeSpecWithDefaults() *UniverseVMImageUpgradeSpec`

NewUniverseVMImageUpgradeSpecWithDefaults instantiates a new UniverseVMImageUpgradeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseVMImageUpgradeSpec) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseVMImageUpgradeSpec) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseVMImageUpgradeSpec) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseVMImageUpgradeSpec) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseVMImageUpgradeSpec) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseVMImageUpgradeSpec) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseVMImageUpgradeSpec) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseVMImageUpgradeSpec) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetImageBundles

`func (o *UniverseVMImageUpgradeSpec) GetImageBundles() []VMImageClusterSpec`

GetImageBundles returns the ImageBundles field if non-nil, zero value otherwise.

### GetImageBundlesOk

`func (o *UniverseVMImageUpgradeSpec) GetImageBundlesOk() (*[]VMImageClusterSpec, bool)`

GetImageBundlesOk returns a tuple with the ImageBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundles

`func (o *UniverseVMImageUpgradeSpec) SetImageBundles(v []VMImageClusterSpec)`

SetImageBundles sets ImageBundles field to given value.

### HasImageBundles

`func (o *UniverseVMImageUpgradeSpec) HasImageBundles() bool`

HasImageBundles returns a boolean if a field has been set.

### GetForceUpgrade

`func (o *UniverseVMImageUpgradeSpec) GetForceUpgrade() bool`

GetForceUpgrade returns the ForceUpgrade field if non-nil, zero value otherwise.

### GetForceUpgradeOk

`func (o *UniverseVMImageUpgradeSpec) GetForceUpgradeOk() (*bool, bool)`

GetForceUpgradeOk returns a tuple with the ForceUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUpgrade

`func (o *UniverseVMImageUpgradeSpec) SetForceUpgrade(v bool)`

SetForceUpgrade sets ForceUpgrade field to given value.

### HasForceUpgrade

`func (o *UniverseVMImageUpgradeSpec) HasForceUpgrade() bool`

HasForceUpgrade returns a boolean if a field has been set.

### GetYbSoftwareVersion

`func (o *UniverseVMImageUpgradeSpec) GetYbSoftwareVersion() string`

GetYbSoftwareVersion returns the YbSoftwareVersion field if non-nil, zero value otherwise.

### GetYbSoftwareVersionOk

`func (o *UniverseVMImageUpgradeSpec) GetYbSoftwareVersionOk() (*string, bool)`

GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbSoftwareVersion

`func (o *UniverseVMImageUpgradeSpec) SetYbSoftwareVersion(v string)`

SetYbSoftwareVersion sets YbSoftwareVersion field to given value.

### HasYbSoftwareVersion

`func (o *UniverseVMImageUpgradeSpec) HasYbSoftwareVersion() bool`

HasYbSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


