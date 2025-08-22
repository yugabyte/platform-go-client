# UniverseQueryLogsExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**RollingUpgrade** | Pointer to **bool** | Perform a rolling upgrade where only one node is upgraded at a time. This is the default behavior. False will perform a non-rolling upgrade where all nodes are upgraded at the same  | [optional] [default to true]
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 
**InstallOtelCollector** | **bool** | Flag to install OpenTelemetry Collector | 
**QueryLogConfig** | [**QueryLogConfig**](QueryLogConfig.md) |  | 

## Methods

### NewUniverseQueryLogsExport

`func NewUniverseQueryLogsExport(installOtelCollector bool, queryLogConfig QueryLogConfig, ) *UniverseQueryLogsExport`

NewUniverseQueryLogsExport instantiates a new UniverseQueryLogsExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseQueryLogsExportWithDefaults

`func NewUniverseQueryLogsExportWithDefaults() *UniverseQueryLogsExport`

NewUniverseQueryLogsExportWithDefaults instantiates a new UniverseQueryLogsExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseQueryLogsExport) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseQueryLogsExport) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseQueryLogsExport) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseQueryLogsExport) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseQueryLogsExport) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseQueryLogsExport) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseQueryLogsExport) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseQueryLogsExport) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetRollingUpgrade

`func (o *UniverseQueryLogsExport) GetRollingUpgrade() bool`

GetRollingUpgrade returns the RollingUpgrade field if non-nil, zero value otherwise.

### GetRollingUpgradeOk

`func (o *UniverseQueryLogsExport) GetRollingUpgradeOk() (*bool, bool)`

GetRollingUpgradeOk returns a tuple with the RollingUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpgrade

`func (o *UniverseQueryLogsExport) SetRollingUpgrade(v bool)`

SetRollingUpgrade sets RollingUpgrade field to given value.

### HasRollingUpgrade

`func (o *UniverseQueryLogsExport) HasRollingUpgrade() bool`

HasRollingUpgrade returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseQueryLogsExport) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseQueryLogsExport) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseQueryLogsExport) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseQueryLogsExport) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetInstallOtelCollector

`func (o *UniverseQueryLogsExport) GetInstallOtelCollector() bool`

GetInstallOtelCollector returns the InstallOtelCollector field if non-nil, zero value otherwise.

### GetInstallOtelCollectorOk

`func (o *UniverseQueryLogsExport) GetInstallOtelCollectorOk() (*bool, bool)`

GetInstallOtelCollectorOk returns a tuple with the InstallOtelCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallOtelCollector

`func (o *UniverseQueryLogsExport) SetInstallOtelCollector(v bool)`

SetInstallOtelCollector sets InstallOtelCollector field to given value.


### GetQueryLogConfig

`func (o *UniverseQueryLogsExport) GetQueryLogConfig() QueryLogConfig`

GetQueryLogConfig returns the QueryLogConfig field if non-nil, zero value otherwise.

### GetQueryLogConfigOk

`func (o *UniverseQueryLogsExport) GetQueryLogConfigOk() (*QueryLogConfig, bool)`

GetQueryLogConfigOk returns a tuple with the QueryLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryLogConfig

`func (o *UniverseQueryLogsExport) SetQueryLogConfig(v QueryLogConfig)`

SetQueryLogConfig sets QueryLogConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


