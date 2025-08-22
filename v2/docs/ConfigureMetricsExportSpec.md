# ConfigureMetricsExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**InstallOtelCollector** | Pointer to **bool** | Similar to DBAL request body. Optional - default false. Requires true if \&quot;universe_metrics_exporter_config\&quot; is not empty and \&quot;universeDetails.otelCollectorEnabled\&quot; is false.  | [optional] [default to false]
**MetricsExportConfig** | Pointer to [**MetricsExportConfig**](MetricsExportConfig.md) |  | [optional] 

## Methods

### NewConfigureMetricsExportSpec

`func NewConfigureMetricsExportSpec() *ConfigureMetricsExportSpec`

NewConfigureMetricsExportSpec instantiates a new ConfigureMetricsExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureMetricsExportSpecWithDefaults

`func NewConfigureMetricsExportSpecWithDefaults() *ConfigureMetricsExportSpec`

NewConfigureMetricsExportSpecWithDefaults instantiates a new ConfigureMetricsExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *ConfigureMetricsExportSpec) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *ConfigureMetricsExportSpec) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *ConfigureMetricsExportSpec) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *ConfigureMetricsExportSpec) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *ConfigureMetricsExportSpec) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *ConfigureMetricsExportSpec) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *ConfigureMetricsExportSpec) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *ConfigureMetricsExportSpec) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetInstallOtelCollector

`func (o *ConfigureMetricsExportSpec) GetInstallOtelCollector() bool`

GetInstallOtelCollector returns the InstallOtelCollector field if non-nil, zero value otherwise.

### GetInstallOtelCollectorOk

`func (o *ConfigureMetricsExportSpec) GetInstallOtelCollectorOk() (*bool, bool)`

GetInstallOtelCollectorOk returns a tuple with the InstallOtelCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallOtelCollector

`func (o *ConfigureMetricsExportSpec) SetInstallOtelCollector(v bool)`

SetInstallOtelCollector sets InstallOtelCollector field to given value.

### HasInstallOtelCollector

`func (o *ConfigureMetricsExportSpec) HasInstallOtelCollector() bool`

HasInstallOtelCollector returns a boolean if a field has been set.

### GetMetricsExportConfig

`func (o *ConfigureMetricsExportSpec) GetMetricsExportConfig() MetricsExportConfig`

GetMetricsExportConfig returns the MetricsExportConfig field if non-nil, zero value otherwise.

### GetMetricsExportConfigOk

`func (o *ConfigureMetricsExportSpec) GetMetricsExportConfigOk() (*MetricsExportConfig, bool)`

GetMetricsExportConfigOk returns a tuple with the MetricsExportConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsExportConfig

`func (o *ConfigureMetricsExportSpec) SetMetricsExportConfig(v MetricsExportConfig)`

SetMetricsExportConfig sets MetricsExportConfig field to given value.

### HasMetricsExportConfig

`func (o *ConfigureMetricsExportSpec) HasMetricsExportConfig() bool`

HasMetricsExportConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


