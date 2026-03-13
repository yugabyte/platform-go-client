# SupportBundleFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchDurationPromDumpMins** | Pointer to **int32** | Batch duration for the prometheus dump (in minutes). Overrides global default. Use with stepPromDumpSecs to get longer historical trends while keeping the same number of data points | [optional] 
**Components** | **[]string** | List of components to be included in the support bundle | 
**EndDate** | **time.Time** | End date to filter logs till | 
**FilterPgAuditLogs** | Pointer to **bool** | Specifies if Postgres audit logs should be filtered out when collecting universe logs. | [optional] 
**MaxCoreFileSize** | Pointer to **int64** | Max size in bytes of the recent collected cores (if any) | [optional] 
**MaxNumRecentCores** | Pointer to **int32** | Max number of the most recent cores to collect (if any) | [optional] 
**PaDumpEndDate** | Pointer to **time.Time** | End date to filter Perf Advisor data | [optional] 
**PaDumpStartDate** | Pointer to **time.Time** | Start date to filter Perf Advisor data | [optional] 
**PaMetricsFormat** | Pointer to **string** | Specifies metrics format. | [optional] 
**PromDumpDownSample** | Pointer to **bool** | When promExportType is REMOTE_READ, whether to downsample raw data points by yb.support_bundle.step_prom_dump_secs or stepPromDumpSecs. Ignored for PROMQL (always downsampled). Default true. | [optional] 
**PromDumpEndDate** | Pointer to **time.Time** | End date to filter prometheus metrics till | [optional] 
**PromDumpStartDate** | Pointer to **time.Time** | Start date to filter prometheus metrics from | [optional] 
**PromExportType** | Pointer to **string** | How to export Prometheus metrics: PROMQL (query_range) or REMOTE_READ. Default PROMQL for backward compatibility. | [optional] 
**PromMetricsFormat** | Pointer to **string** | When promExportType is REMOTE_READ, format for remote read export: PROMQL_JSON or PROM_CHUNK. Default PROMQL_JSON for backward compatibility. | [optional] 
**PromQueries** | Pointer to **map[string]string** | Map of query names to custom PromQL queries to collect in promdump | [optional] 
**PrometheusMetricsTypes** | Pointer to **[]string** | List of exports to be included in the prometheus dump | [optional] 
**StartDate** | **time.Time** | Start date to filter logs from | 
**StepPromDumpSecs** | Pointer to **int32** | Metrics downsample step (in seconds). Overrides global default. Use with batchDurationPromDumpMins to get longer historical trends while keeping the same number of data points | [optional] 

## Methods

### NewSupportBundleFormData

`func NewSupportBundleFormData(components []string, endDate time.Time, startDate time.Time, ) *SupportBundleFormData`

NewSupportBundleFormData instantiates a new SupportBundleFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportBundleFormDataWithDefaults

`func NewSupportBundleFormDataWithDefaults() *SupportBundleFormData`

NewSupportBundleFormDataWithDefaults instantiates a new SupportBundleFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchDurationPromDumpMins

`func (o *SupportBundleFormData) GetBatchDurationPromDumpMins() int32`

GetBatchDurationPromDumpMins returns the BatchDurationPromDumpMins field if non-nil, zero value otherwise.

### GetBatchDurationPromDumpMinsOk

`func (o *SupportBundleFormData) GetBatchDurationPromDumpMinsOk() (*int32, bool)`

GetBatchDurationPromDumpMinsOk returns a tuple with the BatchDurationPromDumpMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchDurationPromDumpMins

`func (o *SupportBundleFormData) SetBatchDurationPromDumpMins(v int32)`

SetBatchDurationPromDumpMins sets BatchDurationPromDumpMins field to given value.

### HasBatchDurationPromDumpMins

`func (o *SupportBundleFormData) HasBatchDurationPromDumpMins() bool`

HasBatchDurationPromDumpMins returns a boolean if a field has been set.

### GetComponents

`func (o *SupportBundleFormData) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SupportBundleFormData) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SupportBundleFormData) SetComponents(v []string)`

SetComponents sets Components field to given value.


### GetEndDate

`func (o *SupportBundleFormData) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SupportBundleFormData) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SupportBundleFormData) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetFilterPgAuditLogs

`func (o *SupportBundleFormData) GetFilterPgAuditLogs() bool`

GetFilterPgAuditLogs returns the FilterPgAuditLogs field if non-nil, zero value otherwise.

### GetFilterPgAuditLogsOk

`func (o *SupportBundleFormData) GetFilterPgAuditLogsOk() (*bool, bool)`

GetFilterPgAuditLogsOk returns a tuple with the FilterPgAuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPgAuditLogs

`func (o *SupportBundleFormData) SetFilterPgAuditLogs(v bool)`

SetFilterPgAuditLogs sets FilterPgAuditLogs field to given value.

### HasFilterPgAuditLogs

`func (o *SupportBundleFormData) HasFilterPgAuditLogs() bool`

HasFilterPgAuditLogs returns a boolean if a field has been set.

### GetMaxCoreFileSize

`func (o *SupportBundleFormData) GetMaxCoreFileSize() int64`

GetMaxCoreFileSize returns the MaxCoreFileSize field if non-nil, zero value otherwise.

### GetMaxCoreFileSizeOk

`func (o *SupportBundleFormData) GetMaxCoreFileSizeOk() (*int64, bool)`

GetMaxCoreFileSizeOk returns a tuple with the MaxCoreFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoreFileSize

`func (o *SupportBundleFormData) SetMaxCoreFileSize(v int64)`

SetMaxCoreFileSize sets MaxCoreFileSize field to given value.

### HasMaxCoreFileSize

`func (o *SupportBundleFormData) HasMaxCoreFileSize() bool`

HasMaxCoreFileSize returns a boolean if a field has been set.

### GetMaxNumRecentCores

`func (o *SupportBundleFormData) GetMaxNumRecentCores() int32`

GetMaxNumRecentCores returns the MaxNumRecentCores field if non-nil, zero value otherwise.

### GetMaxNumRecentCoresOk

`func (o *SupportBundleFormData) GetMaxNumRecentCoresOk() (*int32, bool)`

GetMaxNumRecentCoresOk returns a tuple with the MaxNumRecentCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumRecentCores

`func (o *SupportBundleFormData) SetMaxNumRecentCores(v int32)`

SetMaxNumRecentCores sets MaxNumRecentCores field to given value.

### HasMaxNumRecentCores

`func (o *SupportBundleFormData) HasMaxNumRecentCores() bool`

HasMaxNumRecentCores returns a boolean if a field has been set.

### GetPaDumpEndDate

`func (o *SupportBundleFormData) GetPaDumpEndDate() time.Time`

GetPaDumpEndDate returns the PaDumpEndDate field if non-nil, zero value otherwise.

### GetPaDumpEndDateOk

`func (o *SupportBundleFormData) GetPaDumpEndDateOk() (*time.Time, bool)`

GetPaDumpEndDateOk returns a tuple with the PaDumpEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDumpEndDate

`func (o *SupportBundleFormData) SetPaDumpEndDate(v time.Time)`

SetPaDumpEndDate sets PaDumpEndDate field to given value.

### HasPaDumpEndDate

`func (o *SupportBundleFormData) HasPaDumpEndDate() bool`

HasPaDumpEndDate returns a boolean if a field has been set.

### GetPaDumpStartDate

`func (o *SupportBundleFormData) GetPaDumpStartDate() time.Time`

GetPaDumpStartDate returns the PaDumpStartDate field if non-nil, zero value otherwise.

### GetPaDumpStartDateOk

`func (o *SupportBundleFormData) GetPaDumpStartDateOk() (*time.Time, bool)`

GetPaDumpStartDateOk returns a tuple with the PaDumpStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDumpStartDate

`func (o *SupportBundleFormData) SetPaDumpStartDate(v time.Time)`

SetPaDumpStartDate sets PaDumpStartDate field to given value.

### HasPaDumpStartDate

`func (o *SupportBundleFormData) HasPaDumpStartDate() bool`

HasPaDumpStartDate returns a boolean if a field has been set.

### GetPaMetricsFormat

`func (o *SupportBundleFormData) GetPaMetricsFormat() string`

GetPaMetricsFormat returns the PaMetricsFormat field if non-nil, zero value otherwise.

### GetPaMetricsFormatOk

`func (o *SupportBundleFormData) GetPaMetricsFormatOk() (*string, bool)`

GetPaMetricsFormatOk returns a tuple with the PaMetricsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaMetricsFormat

`func (o *SupportBundleFormData) SetPaMetricsFormat(v string)`

SetPaMetricsFormat sets PaMetricsFormat field to given value.

### HasPaMetricsFormat

`func (o *SupportBundleFormData) HasPaMetricsFormat() bool`

HasPaMetricsFormat returns a boolean if a field has been set.

### GetPromDumpDownSample

`func (o *SupportBundleFormData) GetPromDumpDownSample() bool`

GetPromDumpDownSample returns the PromDumpDownSample field if non-nil, zero value otherwise.

### GetPromDumpDownSampleOk

`func (o *SupportBundleFormData) GetPromDumpDownSampleOk() (*bool, bool)`

GetPromDumpDownSampleOk returns a tuple with the PromDumpDownSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpDownSample

`func (o *SupportBundleFormData) SetPromDumpDownSample(v bool)`

SetPromDumpDownSample sets PromDumpDownSample field to given value.

### HasPromDumpDownSample

`func (o *SupportBundleFormData) HasPromDumpDownSample() bool`

HasPromDumpDownSample returns a boolean if a field has been set.

### GetPromDumpEndDate

`func (o *SupportBundleFormData) GetPromDumpEndDate() time.Time`

GetPromDumpEndDate returns the PromDumpEndDate field if non-nil, zero value otherwise.

### GetPromDumpEndDateOk

`func (o *SupportBundleFormData) GetPromDumpEndDateOk() (*time.Time, bool)`

GetPromDumpEndDateOk returns a tuple with the PromDumpEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpEndDate

`func (o *SupportBundleFormData) SetPromDumpEndDate(v time.Time)`

SetPromDumpEndDate sets PromDumpEndDate field to given value.

### HasPromDumpEndDate

`func (o *SupportBundleFormData) HasPromDumpEndDate() bool`

HasPromDumpEndDate returns a boolean if a field has been set.

### GetPromDumpStartDate

`func (o *SupportBundleFormData) GetPromDumpStartDate() time.Time`

GetPromDumpStartDate returns the PromDumpStartDate field if non-nil, zero value otherwise.

### GetPromDumpStartDateOk

`func (o *SupportBundleFormData) GetPromDumpStartDateOk() (*time.Time, bool)`

GetPromDumpStartDateOk returns a tuple with the PromDumpStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpStartDate

`func (o *SupportBundleFormData) SetPromDumpStartDate(v time.Time)`

SetPromDumpStartDate sets PromDumpStartDate field to given value.

### HasPromDumpStartDate

`func (o *SupportBundleFormData) HasPromDumpStartDate() bool`

HasPromDumpStartDate returns a boolean if a field has been set.

### GetPromExportType

`func (o *SupportBundleFormData) GetPromExportType() string`

GetPromExportType returns the PromExportType field if non-nil, zero value otherwise.

### GetPromExportTypeOk

`func (o *SupportBundleFormData) GetPromExportTypeOk() (*string, bool)`

GetPromExportTypeOk returns a tuple with the PromExportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromExportType

`func (o *SupportBundleFormData) SetPromExportType(v string)`

SetPromExportType sets PromExportType field to given value.

### HasPromExportType

`func (o *SupportBundleFormData) HasPromExportType() bool`

HasPromExportType returns a boolean if a field has been set.

### GetPromMetricsFormat

`func (o *SupportBundleFormData) GetPromMetricsFormat() string`

GetPromMetricsFormat returns the PromMetricsFormat field if non-nil, zero value otherwise.

### GetPromMetricsFormatOk

`func (o *SupportBundleFormData) GetPromMetricsFormatOk() (*string, bool)`

GetPromMetricsFormatOk returns a tuple with the PromMetricsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromMetricsFormat

`func (o *SupportBundleFormData) SetPromMetricsFormat(v string)`

SetPromMetricsFormat sets PromMetricsFormat field to given value.

### HasPromMetricsFormat

`func (o *SupportBundleFormData) HasPromMetricsFormat() bool`

HasPromMetricsFormat returns a boolean if a field has been set.

### GetPromQueries

`func (o *SupportBundleFormData) GetPromQueries() map[string]string`

GetPromQueries returns the PromQueries field if non-nil, zero value otherwise.

### GetPromQueriesOk

`func (o *SupportBundleFormData) GetPromQueriesOk() (*map[string]string, bool)`

GetPromQueriesOk returns a tuple with the PromQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromQueries

`func (o *SupportBundleFormData) SetPromQueries(v map[string]string)`

SetPromQueries sets PromQueries field to given value.

### HasPromQueries

`func (o *SupportBundleFormData) HasPromQueries() bool`

HasPromQueries returns a boolean if a field has been set.

### GetPrometheusMetricsTypes

`func (o *SupportBundleFormData) GetPrometheusMetricsTypes() []string`

GetPrometheusMetricsTypes returns the PrometheusMetricsTypes field if non-nil, zero value otherwise.

### GetPrometheusMetricsTypesOk

`func (o *SupportBundleFormData) GetPrometheusMetricsTypesOk() (*[]string, bool)`

GetPrometheusMetricsTypesOk returns a tuple with the PrometheusMetricsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusMetricsTypes

`func (o *SupportBundleFormData) SetPrometheusMetricsTypes(v []string)`

SetPrometheusMetricsTypes sets PrometheusMetricsTypes field to given value.

### HasPrometheusMetricsTypes

`func (o *SupportBundleFormData) HasPrometheusMetricsTypes() bool`

HasPrometheusMetricsTypes returns a boolean if a field has been set.

### GetStartDate

`func (o *SupportBundleFormData) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SupportBundleFormData) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SupportBundleFormData) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetStepPromDumpSecs

`func (o *SupportBundleFormData) GetStepPromDumpSecs() int32`

GetStepPromDumpSecs returns the StepPromDumpSecs field if non-nil, zero value otherwise.

### GetStepPromDumpSecsOk

`func (o *SupportBundleFormData) GetStepPromDumpSecsOk() (*int32, bool)`

GetStepPromDumpSecsOk returns a tuple with the StepPromDumpSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepPromDumpSecs

`func (o *SupportBundleFormData) SetStepPromDumpSecs(v int32)`

SetStepPromDumpSecs sets StepPromDumpSecs field to given value.

### HasStepPromDumpSecs

`func (o *SupportBundleFormData) HasStepPromDumpSecs() bool`

HasStepPromDumpSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


