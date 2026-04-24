# ExportTelemetryUpgradeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingUpgrade** | Pointer to **bool** | If true, perform rolling restart; otherwise non-rolling. | [optional] [default to true]
**DelayBetweenMasterServers** | Pointer to **int32** | Delay in seconds between restarting master servers during rolling upgrade. | [optional] [default to 0]
**DelayBetweenTserverServers** | Pointer to **int32** | Delay in seconds between restarting tserver servers during rolling upgrade. | [optional] [default to 0]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Sleep duration in milliseconds after each tserver restart. | [optional] [default to 180000]
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Sleep duration in milliseconds after each master restart. | [optional] [default to 180000]

## Methods

### NewExportTelemetryUpgradeOptions

`func NewExportTelemetryUpgradeOptions() *ExportTelemetryUpgradeOptions`

NewExportTelemetryUpgradeOptions instantiates a new ExportTelemetryUpgradeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTelemetryUpgradeOptionsWithDefaults

`func NewExportTelemetryUpgradeOptionsWithDefaults() *ExportTelemetryUpgradeOptions`

NewExportTelemetryUpgradeOptionsWithDefaults instantiates a new ExportTelemetryUpgradeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingUpgrade

`func (o *ExportTelemetryUpgradeOptions) GetRollingUpgrade() bool`

GetRollingUpgrade returns the RollingUpgrade field if non-nil, zero value otherwise.

### GetRollingUpgradeOk

`func (o *ExportTelemetryUpgradeOptions) GetRollingUpgradeOk() (*bool, bool)`

GetRollingUpgradeOk returns a tuple with the RollingUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpgrade

`func (o *ExportTelemetryUpgradeOptions) SetRollingUpgrade(v bool)`

SetRollingUpgrade sets RollingUpgrade field to given value.

### HasRollingUpgrade

`func (o *ExportTelemetryUpgradeOptions) HasRollingUpgrade() bool`

HasRollingUpgrade returns a boolean if a field has been set.

### GetDelayBetweenMasterServers

`func (o *ExportTelemetryUpgradeOptions) GetDelayBetweenMasterServers() int32`

GetDelayBetweenMasterServers returns the DelayBetweenMasterServers field if non-nil, zero value otherwise.

### GetDelayBetweenMasterServersOk

`func (o *ExportTelemetryUpgradeOptions) GetDelayBetweenMasterServersOk() (*int32, bool)`

GetDelayBetweenMasterServersOk returns a tuple with the DelayBetweenMasterServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBetweenMasterServers

`func (o *ExportTelemetryUpgradeOptions) SetDelayBetweenMasterServers(v int32)`

SetDelayBetweenMasterServers sets DelayBetweenMasterServers field to given value.

### HasDelayBetweenMasterServers

`func (o *ExportTelemetryUpgradeOptions) HasDelayBetweenMasterServers() bool`

HasDelayBetweenMasterServers returns a boolean if a field has been set.

### GetDelayBetweenTserverServers

`func (o *ExportTelemetryUpgradeOptions) GetDelayBetweenTserverServers() int32`

GetDelayBetweenTserverServers returns the DelayBetweenTserverServers field if non-nil, zero value otherwise.

### GetDelayBetweenTserverServersOk

`func (o *ExportTelemetryUpgradeOptions) GetDelayBetweenTserverServersOk() (*int32, bool)`

GetDelayBetweenTserverServersOk returns a tuple with the DelayBetweenTserverServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBetweenTserverServers

`func (o *ExportTelemetryUpgradeOptions) SetDelayBetweenTserverServers(v int32)`

SetDelayBetweenTserverServers sets DelayBetweenTserverServers field to given value.

### HasDelayBetweenTserverServers

`func (o *ExportTelemetryUpgradeOptions) HasDelayBetweenTserverServers() bool`

HasDelayBetweenTserverServers returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *ExportTelemetryUpgradeOptions) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *ExportTelemetryUpgradeOptions) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *ExportTelemetryUpgradeOptions) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *ExportTelemetryUpgradeOptions) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *ExportTelemetryUpgradeOptions) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *ExportTelemetryUpgradeOptions) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *ExportTelemetryUpgradeOptions) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *ExportTelemetryUpgradeOptions) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


