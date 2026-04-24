# MetricsTelemetrySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScrapeIntervalSeconds** | Pointer to **int32** | Scrape interval in seconds. Applied on all scrape jobs commonly. | [optional] [default to 30]
**ScrapeTimeoutSeconds** | Pointer to **int32** | Scrape timeout in seconds. Applied on all scrape jobs commonly. | [optional] [default to 20]
**CollectionLevel** | Pointer to **string** | The level of metrics collection. Allowed values are:   - ALL: Collect all available metrics.   - NORMAL: Collect standard set of metrics.   - TABLE_OFF: Disable table level metrics collection.   - MINIMAL: Collect minimal set of metrics.   - OFF: Disable metrics collection.  | [optional] [default to "NORMAL"]
**ScrapeConfigTargets** | Pointer to [**[]ScrapeConfigTargetType**](ScrapeConfigTargetType.md) | Set of target types to include in scrape configuration. If not specified, all supported target types will be included. Allowed values are:   - MASTER_EXPORT: Master server metrics   - TSERVER_EXPORT: Tserver metrics   - YSQL_EXPORT: YSQL server metrics   - CQL_EXPORT: YCQL server metrics   - NODE_EXPORT: Node exporter metrics   - NODE_AGENT_EXPORT: Node agent metrics   - OTEL_EXPORT: OpenTelemetry collector internal metrics  | [optional] [default to [MASTER_EXPORT, TSERVER_EXPORT, YSQL_EXPORT, CQL_EXPORT, NODE_EXPORT, NODE_AGENT_EXPORT, OTEL_EXPORT]]
**Exporters** | Pointer to [**[]TelemetryExporterEntry**](TelemetryExporterEntry.md) | List of exporters. Empty &#x3D; no export. | [optional] 

## Methods

### NewMetricsTelemetrySpec

`func NewMetricsTelemetrySpec() *MetricsTelemetrySpec`

NewMetricsTelemetrySpec instantiates a new MetricsTelemetrySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsTelemetrySpecWithDefaults

`func NewMetricsTelemetrySpecWithDefaults() *MetricsTelemetrySpec`

NewMetricsTelemetrySpecWithDefaults instantiates a new MetricsTelemetrySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScrapeIntervalSeconds

`func (o *MetricsTelemetrySpec) GetScrapeIntervalSeconds() int32`

GetScrapeIntervalSeconds returns the ScrapeIntervalSeconds field if non-nil, zero value otherwise.

### GetScrapeIntervalSecondsOk

`func (o *MetricsTelemetrySpec) GetScrapeIntervalSecondsOk() (*int32, bool)`

GetScrapeIntervalSecondsOk returns a tuple with the ScrapeIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeIntervalSeconds

`func (o *MetricsTelemetrySpec) SetScrapeIntervalSeconds(v int32)`

SetScrapeIntervalSeconds sets ScrapeIntervalSeconds field to given value.

### HasScrapeIntervalSeconds

`func (o *MetricsTelemetrySpec) HasScrapeIntervalSeconds() bool`

HasScrapeIntervalSeconds returns a boolean if a field has been set.

### GetScrapeTimeoutSeconds

`func (o *MetricsTelemetrySpec) GetScrapeTimeoutSeconds() int32`

GetScrapeTimeoutSeconds returns the ScrapeTimeoutSeconds field if non-nil, zero value otherwise.

### GetScrapeTimeoutSecondsOk

`func (o *MetricsTelemetrySpec) GetScrapeTimeoutSecondsOk() (*int32, bool)`

GetScrapeTimeoutSecondsOk returns a tuple with the ScrapeTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeTimeoutSeconds

`func (o *MetricsTelemetrySpec) SetScrapeTimeoutSeconds(v int32)`

SetScrapeTimeoutSeconds sets ScrapeTimeoutSeconds field to given value.

### HasScrapeTimeoutSeconds

`func (o *MetricsTelemetrySpec) HasScrapeTimeoutSeconds() bool`

HasScrapeTimeoutSeconds returns a boolean if a field has been set.

### GetCollectionLevel

`func (o *MetricsTelemetrySpec) GetCollectionLevel() string`

GetCollectionLevel returns the CollectionLevel field if non-nil, zero value otherwise.

### GetCollectionLevelOk

`func (o *MetricsTelemetrySpec) GetCollectionLevelOk() (*string, bool)`

GetCollectionLevelOk returns a tuple with the CollectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLevel

`func (o *MetricsTelemetrySpec) SetCollectionLevel(v string)`

SetCollectionLevel sets CollectionLevel field to given value.

### HasCollectionLevel

`func (o *MetricsTelemetrySpec) HasCollectionLevel() bool`

HasCollectionLevel returns a boolean if a field has been set.

### GetScrapeConfigTargets

`func (o *MetricsTelemetrySpec) GetScrapeConfigTargets() []ScrapeConfigTargetType`

GetScrapeConfigTargets returns the ScrapeConfigTargets field if non-nil, zero value otherwise.

### GetScrapeConfigTargetsOk

`func (o *MetricsTelemetrySpec) GetScrapeConfigTargetsOk() (*[]ScrapeConfigTargetType, bool)`

GetScrapeConfigTargetsOk returns a tuple with the ScrapeConfigTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeConfigTargets

`func (o *MetricsTelemetrySpec) SetScrapeConfigTargets(v []ScrapeConfigTargetType)`

SetScrapeConfigTargets sets ScrapeConfigTargets field to given value.

### HasScrapeConfigTargets

`func (o *MetricsTelemetrySpec) HasScrapeConfigTargets() bool`

HasScrapeConfigTargets returns a boolean if a field has been set.

### GetExporters

`func (o *MetricsTelemetrySpec) GetExporters() []TelemetryExporterEntry`

GetExporters returns the Exporters field if non-nil, zero value otherwise.

### GetExportersOk

`func (o *MetricsTelemetrySpec) GetExportersOk() (*[]TelemetryExporterEntry, bool)`

GetExportersOk returns a tuple with the Exporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporters

`func (o *MetricsTelemetrySpec) SetExporters(v []TelemetryExporterEntry)`

SetExporters sets Exporters field to given value.

### HasExporters

`func (o *MetricsTelemetrySpec) HasExporters() bool`

HasExporters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


