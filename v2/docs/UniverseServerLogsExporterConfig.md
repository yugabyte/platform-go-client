# UniverseServerLogsExporterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalTags** | Pointer to **map[string]string** | Additional tags | [optional] 
**ExporterUuid** | **string** | Exporter uuid | 
**SendBatchMaxSize** | Pointer to **int32** | Send batch max size | [optional] [default to 1000]
**SendBatchSize** | Pointer to **int32** | Send batch size | [optional] [default to 100]
**SendBatchTimeoutSeconds** | Pointer to **int32** | Send batch timeout in seconds | [optional] [default to 10]
**MemoryLimitMib** | Pointer to **int32** | Memory limit in MiB for the OpenTelemetry Collector process in the config file. | [optional] [default to 2048]
**MemoryLimitCheckIntervalSeconds** | Pointer to **int32** | Check interval in seconds for the MemoryLimiterProcessor. | [optional] [default to 10]

## Methods

### NewUniverseServerLogsExporterConfig

`func NewUniverseServerLogsExporterConfig(exporterUuid string, ) *UniverseServerLogsExporterConfig`

NewUniverseServerLogsExporterConfig instantiates a new UniverseServerLogsExporterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseServerLogsExporterConfigWithDefaults

`func NewUniverseServerLogsExporterConfigWithDefaults() *UniverseServerLogsExporterConfig`

NewUniverseServerLogsExporterConfigWithDefaults instantiates a new UniverseServerLogsExporterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalTags

`func (o *UniverseServerLogsExporterConfig) GetAdditionalTags() map[string]string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *UniverseServerLogsExporterConfig) GetAdditionalTagsOk() (*map[string]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *UniverseServerLogsExporterConfig) SetAdditionalTags(v map[string]string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *UniverseServerLogsExporterConfig) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetExporterUuid

`func (o *UniverseServerLogsExporterConfig) GetExporterUuid() string`

GetExporterUuid returns the ExporterUuid field if non-nil, zero value otherwise.

### GetExporterUuidOk

`func (o *UniverseServerLogsExporterConfig) GetExporterUuidOk() (*string, bool)`

GetExporterUuidOk returns a tuple with the ExporterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporterUuid

`func (o *UniverseServerLogsExporterConfig) SetExporterUuid(v string)`

SetExporterUuid sets ExporterUuid field to given value.


### GetSendBatchMaxSize

`func (o *UniverseServerLogsExporterConfig) GetSendBatchMaxSize() int32`

GetSendBatchMaxSize returns the SendBatchMaxSize field if non-nil, zero value otherwise.

### GetSendBatchMaxSizeOk

`func (o *UniverseServerLogsExporterConfig) GetSendBatchMaxSizeOk() (*int32, bool)`

GetSendBatchMaxSizeOk returns a tuple with the SendBatchMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchMaxSize

`func (o *UniverseServerLogsExporterConfig) SetSendBatchMaxSize(v int32)`

SetSendBatchMaxSize sets SendBatchMaxSize field to given value.

### HasSendBatchMaxSize

`func (o *UniverseServerLogsExporterConfig) HasSendBatchMaxSize() bool`

HasSendBatchMaxSize returns a boolean if a field has been set.

### GetSendBatchSize

`func (o *UniverseServerLogsExporterConfig) GetSendBatchSize() int32`

GetSendBatchSize returns the SendBatchSize field if non-nil, zero value otherwise.

### GetSendBatchSizeOk

`func (o *UniverseServerLogsExporterConfig) GetSendBatchSizeOk() (*int32, bool)`

GetSendBatchSizeOk returns a tuple with the SendBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchSize

`func (o *UniverseServerLogsExporterConfig) SetSendBatchSize(v int32)`

SetSendBatchSize sets SendBatchSize field to given value.

### HasSendBatchSize

`func (o *UniverseServerLogsExporterConfig) HasSendBatchSize() bool`

HasSendBatchSize returns a boolean if a field has been set.

### GetSendBatchTimeoutSeconds

`func (o *UniverseServerLogsExporterConfig) GetSendBatchTimeoutSeconds() int32`

GetSendBatchTimeoutSeconds returns the SendBatchTimeoutSeconds field if non-nil, zero value otherwise.

### GetSendBatchTimeoutSecondsOk

`func (o *UniverseServerLogsExporterConfig) GetSendBatchTimeoutSecondsOk() (*int32, bool)`

GetSendBatchTimeoutSecondsOk returns a tuple with the SendBatchTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchTimeoutSeconds

`func (o *UniverseServerLogsExporterConfig) SetSendBatchTimeoutSeconds(v int32)`

SetSendBatchTimeoutSeconds sets SendBatchTimeoutSeconds field to given value.

### HasSendBatchTimeoutSeconds

`func (o *UniverseServerLogsExporterConfig) HasSendBatchTimeoutSeconds() bool`

HasSendBatchTimeoutSeconds returns a boolean if a field has been set.

### GetMemoryLimitMib

`func (o *UniverseServerLogsExporterConfig) GetMemoryLimitMib() int32`

GetMemoryLimitMib returns the MemoryLimitMib field if non-nil, zero value otherwise.

### GetMemoryLimitMibOk

`func (o *UniverseServerLogsExporterConfig) GetMemoryLimitMibOk() (*int32, bool)`

GetMemoryLimitMibOk returns a tuple with the MemoryLimitMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitMib

`func (o *UniverseServerLogsExporterConfig) SetMemoryLimitMib(v int32)`

SetMemoryLimitMib sets MemoryLimitMib field to given value.

### HasMemoryLimitMib

`func (o *UniverseServerLogsExporterConfig) HasMemoryLimitMib() bool`

HasMemoryLimitMib returns a boolean if a field has been set.

### GetMemoryLimitCheckIntervalSeconds

`func (o *UniverseServerLogsExporterConfig) GetMemoryLimitCheckIntervalSeconds() int32`

GetMemoryLimitCheckIntervalSeconds returns the MemoryLimitCheckIntervalSeconds field if non-nil, zero value otherwise.

### GetMemoryLimitCheckIntervalSecondsOk

`func (o *UniverseServerLogsExporterConfig) GetMemoryLimitCheckIntervalSecondsOk() (*int32, bool)`

GetMemoryLimitCheckIntervalSecondsOk returns a tuple with the MemoryLimitCheckIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitCheckIntervalSeconds

`func (o *UniverseServerLogsExporterConfig) SetMemoryLimitCheckIntervalSeconds(v int32)`

SetMemoryLimitCheckIntervalSeconds sets MemoryLimitCheckIntervalSeconds field to given value.

### HasMemoryLimitCheckIntervalSeconds

`func (o *UniverseServerLogsExporterConfig) HasMemoryLimitCheckIntervalSeconds() bool`

HasMemoryLimitCheckIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


