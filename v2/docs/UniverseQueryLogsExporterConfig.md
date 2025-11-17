# UniverseQueryLogsExporterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalTags** | Pointer to **map[string]string** | Additional tags | [optional] 
**ExporterUuid** | **string** | Exporter uuid | 
**SendBatchMaxSize** | Pointer to **int32** | Send batch max size | [optional] [default to 1000]
**SendBatchSize** | Pointer to **int32** | Send batch max size | [optional] [default to 100]
**SendBatchTimeoutSeconds** | Pointer to **int32** | Send batch timeout in seconds | [optional] [default to 10]
**MemoryLimitMib** | Pointer to **int32** | Memory limit in MiB for the OpenTelemetry Collector process in the config file. | [optional] [default to 2048]
**MemoryLimitCheckIntervalSeconds** | Pointer to **int32** | Check interval in seconds for the MemoryLimiterProcessor. | [optional] [default to 10]

## Methods

### NewUniverseQueryLogsExporterConfig

`func NewUniverseQueryLogsExporterConfig(exporterUuid string, ) *UniverseQueryLogsExporterConfig`

NewUniverseQueryLogsExporterConfig instantiates a new UniverseQueryLogsExporterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseQueryLogsExporterConfigWithDefaults

`func NewUniverseQueryLogsExporterConfigWithDefaults() *UniverseQueryLogsExporterConfig`

NewUniverseQueryLogsExporterConfigWithDefaults instantiates a new UniverseQueryLogsExporterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalTags

`func (o *UniverseQueryLogsExporterConfig) GetAdditionalTags() map[string]string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *UniverseQueryLogsExporterConfig) GetAdditionalTagsOk() (*map[string]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *UniverseQueryLogsExporterConfig) SetAdditionalTags(v map[string]string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *UniverseQueryLogsExporterConfig) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetExporterUuid

`func (o *UniverseQueryLogsExporterConfig) GetExporterUuid() string`

GetExporterUuid returns the ExporterUuid field if non-nil, zero value otherwise.

### GetExporterUuidOk

`func (o *UniverseQueryLogsExporterConfig) GetExporterUuidOk() (*string, bool)`

GetExporterUuidOk returns a tuple with the ExporterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporterUuid

`func (o *UniverseQueryLogsExporterConfig) SetExporterUuid(v string)`

SetExporterUuid sets ExporterUuid field to given value.


### GetSendBatchMaxSize

`func (o *UniverseQueryLogsExporterConfig) GetSendBatchMaxSize() int32`

GetSendBatchMaxSize returns the SendBatchMaxSize field if non-nil, zero value otherwise.

### GetSendBatchMaxSizeOk

`func (o *UniverseQueryLogsExporterConfig) GetSendBatchMaxSizeOk() (*int32, bool)`

GetSendBatchMaxSizeOk returns a tuple with the SendBatchMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchMaxSize

`func (o *UniverseQueryLogsExporterConfig) SetSendBatchMaxSize(v int32)`

SetSendBatchMaxSize sets SendBatchMaxSize field to given value.

### HasSendBatchMaxSize

`func (o *UniverseQueryLogsExporterConfig) HasSendBatchMaxSize() bool`

HasSendBatchMaxSize returns a boolean if a field has been set.

### GetSendBatchSize

`func (o *UniverseQueryLogsExporterConfig) GetSendBatchSize() int32`

GetSendBatchSize returns the SendBatchSize field if non-nil, zero value otherwise.

### GetSendBatchSizeOk

`func (o *UniverseQueryLogsExporterConfig) GetSendBatchSizeOk() (*int32, bool)`

GetSendBatchSizeOk returns a tuple with the SendBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchSize

`func (o *UniverseQueryLogsExporterConfig) SetSendBatchSize(v int32)`

SetSendBatchSize sets SendBatchSize field to given value.

### HasSendBatchSize

`func (o *UniverseQueryLogsExporterConfig) HasSendBatchSize() bool`

HasSendBatchSize returns a boolean if a field has been set.

### GetSendBatchTimeoutSeconds

`func (o *UniverseQueryLogsExporterConfig) GetSendBatchTimeoutSeconds() int32`

GetSendBatchTimeoutSeconds returns the SendBatchTimeoutSeconds field if non-nil, zero value otherwise.

### GetSendBatchTimeoutSecondsOk

`func (o *UniverseQueryLogsExporterConfig) GetSendBatchTimeoutSecondsOk() (*int32, bool)`

GetSendBatchTimeoutSecondsOk returns a tuple with the SendBatchTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchTimeoutSeconds

`func (o *UniverseQueryLogsExporterConfig) SetSendBatchTimeoutSeconds(v int32)`

SetSendBatchTimeoutSeconds sets SendBatchTimeoutSeconds field to given value.

### HasSendBatchTimeoutSeconds

`func (o *UniverseQueryLogsExporterConfig) HasSendBatchTimeoutSeconds() bool`

HasSendBatchTimeoutSeconds returns a boolean if a field has been set.

### GetMemoryLimitMib

`func (o *UniverseQueryLogsExporterConfig) GetMemoryLimitMib() int32`

GetMemoryLimitMib returns the MemoryLimitMib field if non-nil, zero value otherwise.

### GetMemoryLimitMibOk

`func (o *UniverseQueryLogsExporterConfig) GetMemoryLimitMibOk() (*int32, bool)`

GetMemoryLimitMibOk returns a tuple with the MemoryLimitMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitMib

`func (o *UniverseQueryLogsExporterConfig) SetMemoryLimitMib(v int32)`

SetMemoryLimitMib sets MemoryLimitMib field to given value.

### HasMemoryLimitMib

`func (o *UniverseQueryLogsExporterConfig) HasMemoryLimitMib() bool`

HasMemoryLimitMib returns a boolean if a field has been set.

### GetMemoryLimitCheckIntervalSeconds

`func (o *UniverseQueryLogsExporterConfig) GetMemoryLimitCheckIntervalSeconds() int32`

GetMemoryLimitCheckIntervalSeconds returns the MemoryLimitCheckIntervalSeconds field if non-nil, zero value otherwise.

### GetMemoryLimitCheckIntervalSecondsOk

`func (o *UniverseQueryLogsExporterConfig) GetMemoryLimitCheckIntervalSecondsOk() (*int32, bool)`

GetMemoryLimitCheckIntervalSecondsOk returns a tuple with the MemoryLimitCheckIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitCheckIntervalSeconds

`func (o *UniverseQueryLogsExporterConfig) SetMemoryLimitCheckIntervalSeconds(v int32)`

SetMemoryLimitCheckIntervalSeconds sets MemoryLimitCheckIntervalSeconds field to given value.

### HasMemoryLimitCheckIntervalSeconds

`func (o *UniverseQueryLogsExporterConfig) HasMemoryLimitCheckIntervalSeconds() bool`

HasMemoryLimitCheckIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


