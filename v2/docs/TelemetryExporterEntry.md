# TelemetryExporterEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExporterUuid** | **string** | UUID of the telemetry provider (exporter). | 
**AdditionalTags** | Pointer to **map[string]string** | Key-value tags to attach to exported data. | [optional] 
**SendBatchMaxSize** | Pointer to **int32** | Maximum number of metrics to buffer before sending. | [optional] [default to 1000]
**SendBatchSize** | Pointer to **int32** | Target batch size for sending metrics. | [optional] [default to 100]
**SendBatchTimeoutSeconds** | Pointer to **int32** | Timeout in seconds for sending a batch. | [optional] [default to 10]
**MemoryLimitMib** | Pointer to **int32** | Memory limit in MiB for the exporter. | [optional] [default to 2048]
**MemoryLimitCheckIntervalSeconds** | Pointer to **int32** | Interval in seconds between memory limit checks. | [optional] [default to 10]
**MetricsPrefix** | Pointer to **string** | Prefix added to metric names (e.g. ybdb.). | [optional] [default to ""]

## Methods

### NewTelemetryExporterEntry

`func NewTelemetryExporterEntry(exporterUuid string, ) *TelemetryExporterEntry`

NewTelemetryExporterEntry instantiates a new TelemetryExporterEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryExporterEntryWithDefaults

`func NewTelemetryExporterEntryWithDefaults() *TelemetryExporterEntry`

NewTelemetryExporterEntryWithDefaults instantiates a new TelemetryExporterEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExporterUuid

`func (o *TelemetryExporterEntry) GetExporterUuid() string`

GetExporterUuid returns the ExporterUuid field if non-nil, zero value otherwise.

### GetExporterUuidOk

`func (o *TelemetryExporterEntry) GetExporterUuidOk() (*string, bool)`

GetExporterUuidOk returns a tuple with the ExporterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporterUuid

`func (o *TelemetryExporterEntry) SetExporterUuid(v string)`

SetExporterUuid sets ExporterUuid field to given value.


### GetAdditionalTags

`func (o *TelemetryExporterEntry) GetAdditionalTags() map[string]string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *TelemetryExporterEntry) GetAdditionalTagsOk() (*map[string]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *TelemetryExporterEntry) SetAdditionalTags(v map[string]string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *TelemetryExporterEntry) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetSendBatchMaxSize

`func (o *TelemetryExporterEntry) GetSendBatchMaxSize() int32`

GetSendBatchMaxSize returns the SendBatchMaxSize field if non-nil, zero value otherwise.

### GetSendBatchMaxSizeOk

`func (o *TelemetryExporterEntry) GetSendBatchMaxSizeOk() (*int32, bool)`

GetSendBatchMaxSizeOk returns a tuple with the SendBatchMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchMaxSize

`func (o *TelemetryExporterEntry) SetSendBatchMaxSize(v int32)`

SetSendBatchMaxSize sets SendBatchMaxSize field to given value.

### HasSendBatchMaxSize

`func (o *TelemetryExporterEntry) HasSendBatchMaxSize() bool`

HasSendBatchMaxSize returns a boolean if a field has been set.

### GetSendBatchSize

`func (o *TelemetryExporterEntry) GetSendBatchSize() int32`

GetSendBatchSize returns the SendBatchSize field if non-nil, zero value otherwise.

### GetSendBatchSizeOk

`func (o *TelemetryExporterEntry) GetSendBatchSizeOk() (*int32, bool)`

GetSendBatchSizeOk returns a tuple with the SendBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchSize

`func (o *TelemetryExporterEntry) SetSendBatchSize(v int32)`

SetSendBatchSize sets SendBatchSize field to given value.

### HasSendBatchSize

`func (o *TelemetryExporterEntry) HasSendBatchSize() bool`

HasSendBatchSize returns a boolean if a field has been set.

### GetSendBatchTimeoutSeconds

`func (o *TelemetryExporterEntry) GetSendBatchTimeoutSeconds() int32`

GetSendBatchTimeoutSeconds returns the SendBatchTimeoutSeconds field if non-nil, zero value otherwise.

### GetSendBatchTimeoutSecondsOk

`func (o *TelemetryExporterEntry) GetSendBatchTimeoutSecondsOk() (*int32, bool)`

GetSendBatchTimeoutSecondsOk returns a tuple with the SendBatchTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchTimeoutSeconds

`func (o *TelemetryExporterEntry) SetSendBatchTimeoutSeconds(v int32)`

SetSendBatchTimeoutSeconds sets SendBatchTimeoutSeconds field to given value.

### HasSendBatchTimeoutSeconds

`func (o *TelemetryExporterEntry) HasSendBatchTimeoutSeconds() bool`

HasSendBatchTimeoutSeconds returns a boolean if a field has been set.

### GetMemoryLimitMib

`func (o *TelemetryExporterEntry) GetMemoryLimitMib() int32`

GetMemoryLimitMib returns the MemoryLimitMib field if non-nil, zero value otherwise.

### GetMemoryLimitMibOk

`func (o *TelemetryExporterEntry) GetMemoryLimitMibOk() (*int32, bool)`

GetMemoryLimitMibOk returns a tuple with the MemoryLimitMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitMib

`func (o *TelemetryExporterEntry) SetMemoryLimitMib(v int32)`

SetMemoryLimitMib sets MemoryLimitMib field to given value.

### HasMemoryLimitMib

`func (o *TelemetryExporterEntry) HasMemoryLimitMib() bool`

HasMemoryLimitMib returns a boolean if a field has been set.

### GetMemoryLimitCheckIntervalSeconds

`func (o *TelemetryExporterEntry) GetMemoryLimitCheckIntervalSeconds() int32`

GetMemoryLimitCheckIntervalSeconds returns the MemoryLimitCheckIntervalSeconds field if non-nil, zero value otherwise.

### GetMemoryLimitCheckIntervalSecondsOk

`func (o *TelemetryExporterEntry) GetMemoryLimitCheckIntervalSecondsOk() (*int32, bool)`

GetMemoryLimitCheckIntervalSecondsOk returns a tuple with the MemoryLimitCheckIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitCheckIntervalSeconds

`func (o *TelemetryExporterEntry) SetMemoryLimitCheckIntervalSeconds(v int32)`

SetMemoryLimitCheckIntervalSeconds sets MemoryLimitCheckIntervalSeconds field to given value.

### HasMemoryLimitCheckIntervalSeconds

`func (o *TelemetryExporterEntry) HasMemoryLimitCheckIntervalSeconds() bool`

HasMemoryLimitCheckIntervalSeconds returns a boolean if a field has been set.

### GetMetricsPrefix

`func (o *TelemetryExporterEntry) GetMetricsPrefix() string`

GetMetricsPrefix returns the MetricsPrefix field if non-nil, zero value otherwise.

### GetMetricsPrefixOk

`func (o *TelemetryExporterEntry) GetMetricsPrefixOk() (*string, bool)`

GetMetricsPrefixOk returns a tuple with the MetricsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPrefix

`func (o *TelemetryExporterEntry) SetMetricsPrefix(v string)`

SetMetricsPrefix sets MetricsPrefix field to given value.

### HasMetricsPrefix

`func (o *TelemetryExporterEntry) HasMetricsPrefix() bool`

HasMetricsPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


