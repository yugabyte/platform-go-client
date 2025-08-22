# UniverseMetricsExporterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalTags** | **map[string]string** | Additional tags | 
**ExporterUuid** | **string** | Exporter uuid | [readonly] 
**MetricsPrefix** | Pointer to **string** | Custom prefix to add to all exported metric names. Can be customised to each exporter differently. | [optional] 
**SendBatchMaxSize** | Pointer to **int32** | Maximum batch size for sending metrics. Can be customised to each exporter differently. | [optional] 
**SendBatchSize** | Pointer to **int32** | Batch size for sending metrics. Can be customised to each exporter differently. | [optional] 
**SendBatchTimeoutSeconds** | Pointer to **int32** | Batch timeout in seconds for sending metrics. Can be customised to each exporter differently. | [optional] 

## Methods

### NewUniverseMetricsExporterConfig

`func NewUniverseMetricsExporterConfig(additionalTags map[string]string, exporterUuid string, ) *UniverseMetricsExporterConfig`

NewUniverseMetricsExporterConfig instantiates a new UniverseMetricsExporterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseMetricsExporterConfigWithDefaults

`func NewUniverseMetricsExporterConfigWithDefaults() *UniverseMetricsExporterConfig`

NewUniverseMetricsExporterConfigWithDefaults instantiates a new UniverseMetricsExporterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalTags

`func (o *UniverseMetricsExporterConfig) GetAdditionalTags() map[string]string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *UniverseMetricsExporterConfig) GetAdditionalTagsOk() (*map[string]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *UniverseMetricsExporterConfig) SetAdditionalTags(v map[string]string)`

SetAdditionalTags sets AdditionalTags field to given value.


### GetExporterUuid

`func (o *UniverseMetricsExporterConfig) GetExporterUuid() string`

GetExporterUuid returns the ExporterUuid field if non-nil, zero value otherwise.

### GetExporterUuidOk

`func (o *UniverseMetricsExporterConfig) GetExporterUuidOk() (*string, bool)`

GetExporterUuidOk returns a tuple with the ExporterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporterUuid

`func (o *UniverseMetricsExporterConfig) SetExporterUuid(v string)`

SetExporterUuid sets ExporterUuid field to given value.


### GetMetricsPrefix

`func (o *UniverseMetricsExporterConfig) GetMetricsPrefix() string`

GetMetricsPrefix returns the MetricsPrefix field if non-nil, zero value otherwise.

### GetMetricsPrefixOk

`func (o *UniverseMetricsExporterConfig) GetMetricsPrefixOk() (*string, bool)`

GetMetricsPrefixOk returns a tuple with the MetricsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPrefix

`func (o *UniverseMetricsExporterConfig) SetMetricsPrefix(v string)`

SetMetricsPrefix sets MetricsPrefix field to given value.

### HasMetricsPrefix

`func (o *UniverseMetricsExporterConfig) HasMetricsPrefix() bool`

HasMetricsPrefix returns a boolean if a field has been set.

### GetSendBatchMaxSize

`func (o *UniverseMetricsExporterConfig) GetSendBatchMaxSize() int32`

GetSendBatchMaxSize returns the SendBatchMaxSize field if non-nil, zero value otherwise.

### GetSendBatchMaxSizeOk

`func (o *UniverseMetricsExporterConfig) GetSendBatchMaxSizeOk() (*int32, bool)`

GetSendBatchMaxSizeOk returns a tuple with the SendBatchMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchMaxSize

`func (o *UniverseMetricsExporterConfig) SetSendBatchMaxSize(v int32)`

SetSendBatchMaxSize sets SendBatchMaxSize field to given value.

### HasSendBatchMaxSize

`func (o *UniverseMetricsExporterConfig) HasSendBatchMaxSize() bool`

HasSendBatchMaxSize returns a boolean if a field has been set.

### GetSendBatchSize

`func (o *UniverseMetricsExporterConfig) GetSendBatchSize() int32`

GetSendBatchSize returns the SendBatchSize field if non-nil, zero value otherwise.

### GetSendBatchSizeOk

`func (o *UniverseMetricsExporterConfig) GetSendBatchSizeOk() (*int32, bool)`

GetSendBatchSizeOk returns a tuple with the SendBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchSize

`func (o *UniverseMetricsExporterConfig) SetSendBatchSize(v int32)`

SetSendBatchSize sets SendBatchSize field to given value.

### HasSendBatchSize

`func (o *UniverseMetricsExporterConfig) HasSendBatchSize() bool`

HasSendBatchSize returns a boolean if a field has been set.

### GetSendBatchTimeoutSeconds

`func (o *UniverseMetricsExporterConfig) GetSendBatchTimeoutSeconds() int32`

GetSendBatchTimeoutSeconds returns the SendBatchTimeoutSeconds field if non-nil, zero value otherwise.

### GetSendBatchTimeoutSecondsOk

`func (o *UniverseMetricsExporterConfig) GetSendBatchTimeoutSecondsOk() (*int32, bool)`

GetSendBatchTimeoutSecondsOk returns a tuple with the SendBatchTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchTimeoutSeconds

`func (o *UniverseMetricsExporterConfig) SetSendBatchTimeoutSeconds(v int32)`

SetSendBatchTimeoutSeconds sets SendBatchTimeoutSeconds field to given value.

### HasSendBatchTimeoutSeconds

`func (o *UniverseMetricsExporterConfig) HasSendBatchTimeoutSeconds() bool`

HasSendBatchTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


