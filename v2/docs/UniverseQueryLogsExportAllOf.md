# UniverseQueryLogsExportAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallOtelCollector** | **bool** | Flag to install OpenTelemetry Collector | 
**QueryLogConfig** | [**QueryLogConfig**](QueryLogConfig.md) |  | 

## Methods

### NewUniverseQueryLogsExportAllOf

`func NewUniverseQueryLogsExportAllOf(installOtelCollector bool, queryLogConfig QueryLogConfig, ) *UniverseQueryLogsExportAllOf`

NewUniverseQueryLogsExportAllOf instantiates a new UniverseQueryLogsExportAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseQueryLogsExportAllOfWithDefaults

`func NewUniverseQueryLogsExportAllOfWithDefaults() *UniverseQueryLogsExportAllOf`

NewUniverseQueryLogsExportAllOfWithDefaults instantiates a new UniverseQueryLogsExportAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallOtelCollector

`func (o *UniverseQueryLogsExportAllOf) GetInstallOtelCollector() bool`

GetInstallOtelCollector returns the InstallOtelCollector field if non-nil, zero value otherwise.

### GetInstallOtelCollectorOk

`func (o *UniverseQueryLogsExportAllOf) GetInstallOtelCollectorOk() (*bool, bool)`

GetInstallOtelCollectorOk returns a tuple with the InstallOtelCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallOtelCollector

`func (o *UniverseQueryLogsExportAllOf) SetInstallOtelCollector(v bool)`

SetInstallOtelCollector sets InstallOtelCollector field to given value.


### GetQueryLogConfig

`func (o *UniverseQueryLogsExportAllOf) GetQueryLogConfig() QueryLogConfig`

GetQueryLogConfig returns the QueryLogConfig field if non-nil, zero value otherwise.

### GetQueryLogConfigOk

`func (o *UniverseQueryLogsExportAllOf) GetQueryLogConfigOk() (*QueryLogConfig, bool)`

GetQueryLogConfigOk returns a tuple with the QueryLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryLogConfig

`func (o *UniverseQueryLogsExportAllOf) SetQueryLogConfig(v QueryLogConfig)`

SetQueryLogConfig sets QueryLogConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


