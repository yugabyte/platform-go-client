# UniverseRestart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**RollingRestart** | Pointer to **bool** | Perform a rolling restart of the universe. Otherwise, all nodes will be restarted at the same time. | [optional] [default to true]
**RestartType** | Pointer to **string** | The method to reboot the node. This is not required for kubernetes universes, as the pods  will get restarted no matter what. \&quot;HARD\&quot; reboots are not supported today.  OS: Restarts the node via the operating system. SERVICE: Restart the YugabyteDB Process only (master, tserver, etc).  | [optional] [default to "SERVICE"]

## Methods

### NewUniverseRestart

`func NewUniverseRestart() *UniverseRestart`

NewUniverseRestart instantiates a new UniverseRestart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseRestartWithDefaults

`func NewUniverseRestartWithDefaults() *UniverseRestart`

NewUniverseRestartWithDefaults instantiates a new UniverseRestart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseRestart) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseRestart) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseRestart) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseRestart) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseRestart) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseRestart) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseRestart) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseRestart) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetRollingRestart

`func (o *UniverseRestart) GetRollingRestart() bool`

GetRollingRestart returns the RollingRestart field if non-nil, zero value otherwise.

### GetRollingRestartOk

`func (o *UniverseRestart) GetRollingRestartOk() (*bool, bool)`

GetRollingRestartOk returns a tuple with the RollingRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingRestart

`func (o *UniverseRestart) SetRollingRestart(v bool)`

SetRollingRestart sets RollingRestart field to given value.

### HasRollingRestart

`func (o *UniverseRestart) HasRollingRestart() bool`

HasRollingRestart returns a boolean if a field has been set.

### GetRestartType

`func (o *UniverseRestart) GetRestartType() string`

GetRestartType returns the RestartType field if non-nil, zero value otherwise.

### GetRestartTypeOk

`func (o *UniverseRestart) GetRestartTypeOk() (*string, bool)`

GetRestartTypeOk returns a tuple with the RestartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartType

`func (o *UniverseRestart) SetRestartType(v string)`

SetRestartType sets RestartType field to given value.

### HasRestartType

`func (o *UniverseRestart) HasRestartType() bool`

HasRestartType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


