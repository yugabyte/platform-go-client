# ConfigureMetricsExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


