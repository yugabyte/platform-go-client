# MetricsExportConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScrapeIntervalSeconds** | Pointer to **int32** | Scrape interval in seconds. Applied on all scrape jobs commonly. | [optional] [default to 30]
**ScrapeTimeoutSeconds** | Pointer to **int32** | Scrape timeout in seconds. Applied on all scrape jobs commonly. | [optional] [default to 20]
**CollectionLevel** | Pointer to **string** | The level of metrics collection. Allowed values are:   - ALL: Collect all available metrics.   - NORMAL: Collect standard set of metrics.   - TABLE_OFF: Disable table level metrics collection.   - MINIMAL: Collect minimal set of metrics.   - OFF: Disable metrics collection.  | [optional] [default to "NORMAL"]
**ScrapeConfigTargets** | Pointer to [**[]ScrapeConfigTargetType**](ScrapeConfigTargetType.md) | Set of target types to include in scrape configuration. If not specified, all supported target types will be included. Allowed values are:   - MASTER_EXPORT: Master server metrics   - TSERVER_EXPORT: Tserver metrics     - YSQL_EXPORT: YSQL server metrics   - CQL_EXPORT: YCQL server metrics   - NODE_EXPORT: Node exporter metrics   - NODE_AGENT_EXPORT: Node agent metrics   - OTEL_EXPORT: OpenTelemetry collector internal metrics  | [optional] [default to {"MASTER_EXPORT", "TSERVER_EXPORT", "YSQL_EXPORT", "CQL_EXPORT", "NODE_EXPORT", "NODE_AGENT_EXPORT", "OTEL_EXPORT"}]
**UniverseMetricsExporterConfig** | [**[]UniverseMetricsExporterConfig**](UniverseMetricsExporterConfig.md) | List of universe metrics exporter configurations. If empty, no metrics will be sent anywhere. | 

## Methods

### NewMetricsExportConfig

`func NewMetricsExportConfig(universeMetricsExporterConfig []UniverseMetricsExporterConfig, ) *MetricsExportConfig`

NewMetricsExportConfig instantiates a new MetricsExportConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsExportConfigWithDefaults

`func NewMetricsExportConfigWithDefaults() *MetricsExportConfig`

NewMetricsExportConfigWithDefaults instantiates a new MetricsExportConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScrapeIntervalSeconds

`func (o *MetricsExportConfig) GetScrapeIntervalSeconds() int32`

GetScrapeIntervalSeconds returns the ScrapeIntervalSeconds field if non-nil, zero value otherwise.

### GetScrapeIntervalSecondsOk

`func (o *MetricsExportConfig) GetScrapeIntervalSecondsOk() (*int32, bool)`

GetScrapeIntervalSecondsOk returns a tuple with the ScrapeIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeIntervalSeconds

`func (o *MetricsExportConfig) SetScrapeIntervalSeconds(v int32)`

SetScrapeIntervalSeconds sets ScrapeIntervalSeconds field to given value.

### HasScrapeIntervalSeconds

`func (o *MetricsExportConfig) HasScrapeIntervalSeconds() bool`

HasScrapeIntervalSeconds returns a boolean if a field has been set.

### GetScrapeTimeoutSeconds

`func (o *MetricsExportConfig) GetScrapeTimeoutSeconds() int32`

GetScrapeTimeoutSeconds returns the ScrapeTimeoutSeconds field if non-nil, zero value otherwise.

### GetScrapeTimeoutSecondsOk

`func (o *MetricsExportConfig) GetScrapeTimeoutSecondsOk() (*int32, bool)`

GetScrapeTimeoutSecondsOk returns a tuple with the ScrapeTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeTimeoutSeconds

`func (o *MetricsExportConfig) SetScrapeTimeoutSeconds(v int32)`

SetScrapeTimeoutSeconds sets ScrapeTimeoutSeconds field to given value.

### HasScrapeTimeoutSeconds

`func (o *MetricsExportConfig) HasScrapeTimeoutSeconds() bool`

HasScrapeTimeoutSeconds returns a boolean if a field has been set.

### GetCollectionLevel

`func (o *MetricsExportConfig) GetCollectionLevel() string`

GetCollectionLevel returns the CollectionLevel field if non-nil, zero value otherwise.

### GetCollectionLevelOk

`func (o *MetricsExportConfig) GetCollectionLevelOk() (*string, bool)`

GetCollectionLevelOk returns a tuple with the CollectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLevel

`func (o *MetricsExportConfig) SetCollectionLevel(v string)`

SetCollectionLevel sets CollectionLevel field to given value.

### HasCollectionLevel

`func (o *MetricsExportConfig) HasCollectionLevel() bool`

HasCollectionLevel returns a boolean if a field has been set.

### GetScrapeConfigTargets

`func (o *MetricsExportConfig) GetScrapeConfigTargets() []ScrapeConfigTargetType`

GetScrapeConfigTargets returns the ScrapeConfigTargets field if non-nil, zero value otherwise.

### GetScrapeConfigTargetsOk

`func (o *MetricsExportConfig) GetScrapeConfigTargetsOk() (*[]ScrapeConfigTargetType, bool)`

GetScrapeConfigTargetsOk returns a tuple with the ScrapeConfigTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeConfigTargets

`func (o *MetricsExportConfig) SetScrapeConfigTargets(v []ScrapeConfigTargetType)`

SetScrapeConfigTargets sets ScrapeConfigTargets field to given value.

### HasScrapeConfigTargets

`func (o *MetricsExportConfig) HasScrapeConfigTargets() bool`

HasScrapeConfigTargets returns a boolean if a field has been set.

### GetUniverseMetricsExporterConfig

`func (o *MetricsExportConfig) GetUniverseMetricsExporterConfig() []UniverseMetricsExporterConfig`

GetUniverseMetricsExporterConfig returns the UniverseMetricsExporterConfig field if non-nil, zero value otherwise.

### GetUniverseMetricsExporterConfigOk

`func (o *MetricsExportConfig) GetUniverseMetricsExporterConfigOk() (*[]UniverseMetricsExporterConfig, bool)`

GetUniverseMetricsExporterConfigOk returns a tuple with the UniverseMetricsExporterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseMetricsExporterConfig

`func (o *MetricsExportConfig) SetUniverseMetricsExporterConfig(v []UniverseMetricsExporterConfig)`

SetUniverseMetricsExporterConfig sets UniverseMetricsExporterConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


