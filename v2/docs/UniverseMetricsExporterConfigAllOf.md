# UniverseMetricsExporterConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendBatchMaxSize** | Pointer to **int32** | Maximum batch size for sending metrics. Can be customised to each exporter differently. | [optional] [default to 1000]
**SendBatchSize** | Pointer to **int32** | Batch size for sending metrics. Can be customised to each exporter differently. | [optional] [default to 100]
**SendBatchTimeoutSeconds** | Pointer to **int32** | Batch timeout in seconds for sending metrics. Can be customised to each exporter differently. | [optional] [default to 10]
**MetricsPrefix** | Pointer to **string** | Custom prefix to be added to all the metrics. Can be customised to each exporter differently. | [optional] [default to "ybdb."]

## Methods

### NewUniverseMetricsExporterConfigAllOf

`func NewUniverseMetricsExporterConfigAllOf() *UniverseMetricsExporterConfigAllOf`

NewUniverseMetricsExporterConfigAllOf instantiates a new UniverseMetricsExporterConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseMetricsExporterConfigAllOfWithDefaults

`func NewUniverseMetricsExporterConfigAllOfWithDefaults() *UniverseMetricsExporterConfigAllOf`

NewUniverseMetricsExporterConfigAllOfWithDefaults instantiates a new UniverseMetricsExporterConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendBatchMaxSize

`func (o *UniverseMetricsExporterConfigAllOf) GetSendBatchMaxSize() int32`

GetSendBatchMaxSize returns the SendBatchMaxSize field if non-nil, zero value otherwise.

### GetSendBatchMaxSizeOk

`func (o *UniverseMetricsExporterConfigAllOf) GetSendBatchMaxSizeOk() (*int32, bool)`

GetSendBatchMaxSizeOk returns a tuple with the SendBatchMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchMaxSize

`func (o *UniverseMetricsExporterConfigAllOf) SetSendBatchMaxSize(v int32)`

SetSendBatchMaxSize sets SendBatchMaxSize field to given value.

### HasSendBatchMaxSize

`func (o *UniverseMetricsExporterConfigAllOf) HasSendBatchMaxSize() bool`

HasSendBatchMaxSize returns a boolean if a field has been set.

### GetSendBatchSize

`func (o *UniverseMetricsExporterConfigAllOf) GetSendBatchSize() int32`

GetSendBatchSize returns the SendBatchSize field if non-nil, zero value otherwise.

### GetSendBatchSizeOk

`func (o *UniverseMetricsExporterConfigAllOf) GetSendBatchSizeOk() (*int32, bool)`

GetSendBatchSizeOk returns a tuple with the SendBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchSize

`func (o *UniverseMetricsExporterConfigAllOf) SetSendBatchSize(v int32)`

SetSendBatchSize sets SendBatchSize field to given value.

### HasSendBatchSize

`func (o *UniverseMetricsExporterConfigAllOf) HasSendBatchSize() bool`

HasSendBatchSize returns a boolean if a field has been set.

### GetSendBatchTimeoutSeconds

`func (o *UniverseMetricsExporterConfigAllOf) GetSendBatchTimeoutSeconds() int32`

GetSendBatchTimeoutSeconds returns the SendBatchTimeoutSeconds field if non-nil, zero value otherwise.

### GetSendBatchTimeoutSecondsOk

`func (o *UniverseMetricsExporterConfigAllOf) GetSendBatchTimeoutSecondsOk() (*int32, bool)`

GetSendBatchTimeoutSecondsOk returns a tuple with the SendBatchTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchTimeoutSeconds

`func (o *UniverseMetricsExporterConfigAllOf) SetSendBatchTimeoutSeconds(v int32)`

SetSendBatchTimeoutSeconds sets SendBatchTimeoutSeconds field to given value.

### HasSendBatchTimeoutSeconds

`func (o *UniverseMetricsExporterConfigAllOf) HasSendBatchTimeoutSeconds() bool`

HasSendBatchTimeoutSeconds returns a boolean if a field has been set.

### GetMetricsPrefix

`func (o *UniverseMetricsExporterConfigAllOf) GetMetricsPrefix() string`

GetMetricsPrefix returns the MetricsPrefix field if non-nil, zero value otherwise.

### GetMetricsPrefixOk

`func (o *UniverseMetricsExporterConfigAllOf) GetMetricsPrefixOk() (*string, bool)`

GetMetricsPrefixOk returns a tuple with the MetricsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPrefix

`func (o *UniverseMetricsExporterConfigAllOf) SetMetricsPrefix(v string)`

SetMetricsPrefix sets MetricsPrefix field to given value.

### HasMetricsPrefix

`func (o *UniverseMetricsExporterConfigAllOf) HasMetricsPrefix() bool`

HasMetricsPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


