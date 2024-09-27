# UniverseThirdPartySoftwareUpgradeStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**ForceAll** | Pointer to **bool** | force all thirdparty softwares to be upgraded. | [optional] [default to false]

## Methods

### NewUniverseThirdPartySoftwareUpgradeStart

`func NewUniverseThirdPartySoftwareUpgradeStart() *UniverseThirdPartySoftwareUpgradeStart`

NewUniverseThirdPartySoftwareUpgradeStart instantiates a new UniverseThirdPartySoftwareUpgradeStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseThirdPartySoftwareUpgradeStartWithDefaults

`func NewUniverseThirdPartySoftwareUpgradeStartWithDefaults() *UniverseThirdPartySoftwareUpgradeStart`

NewUniverseThirdPartySoftwareUpgradeStartWithDefaults instantiates a new UniverseThirdPartySoftwareUpgradeStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseThirdPartySoftwareUpgradeStart) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseThirdPartySoftwareUpgradeStart) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseThirdPartySoftwareUpgradeStart) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseThirdPartySoftwareUpgradeStart) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseThirdPartySoftwareUpgradeStart) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseThirdPartySoftwareUpgradeStart) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseThirdPartySoftwareUpgradeStart) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseThirdPartySoftwareUpgradeStart) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetForceAll

`func (o *UniverseThirdPartySoftwareUpgradeStart) GetForceAll() bool`

GetForceAll returns the ForceAll field if non-nil, zero value otherwise.

### GetForceAllOk

`func (o *UniverseThirdPartySoftwareUpgradeStart) GetForceAllOk() (*bool, bool)`

GetForceAllOk returns a tuple with the ForceAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAll

`func (o *UniverseThirdPartySoftwareUpgradeStart) SetForceAll(v bool)`

SetForceAll sets ForceAll field to given value.

### HasForceAll

`func (o *UniverseThirdPartySoftwareUpgradeStart) HasForceAll() bool`

HasForceAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


