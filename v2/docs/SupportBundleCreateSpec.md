# SupportBundleCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | [**[]SupportBundleComponentType**](SupportBundleComponentType.md) | List of components to be included in the support bundle. On Kubernetes universes the server drops what a yb-master / yb-tserver pod cannot produce (see the component spec fields below), and a component type is dropped once every one of its sub-components has been dropped. The persisted &#x60;SupportBundleSpec&#x60; and the bundle&#39;s manifest name only what was kept. Rejected with a 400 when nothing collectable remains. K8sInfo is dropped on non-Kubernetes universes.  | 
**StartDate** | Pointer to **time.Time** | Start date to filter logs from. Defaults to 24 hours before end_date. | [optional] 
**EndDate** | Pointer to **time.Time** | End date to filter logs till. Defaults to now. | [optional] 
**NodeNames** | Pointer to **[]string** | Names of the universe nodes to collect node-level components from. When omitted or empty, every node in the universe is considered. Global-level components (PrometheusMetrics, YbaMetadata, TabletReport, PerfAdvisor, K8sInfo, ApplicationLogs) are collected once for the universe and are not scoped by this list. Rejected with a 400 when none of the given names exist in the universe; names that do not exist are ignored when at least one matches. Not applicable to YBA-only support bundles.  | [optional] 
**FilesComponentSpecs** | Pointer to [**[]FilesComponentSpec**](FilesComponentSpec.md) | Specs driving the generic node-level FilesComponent. Required when components contains FilesComponent. On Kubernetes universes the SystemLogs and NodeAgent specs are dropped, since a pod has neither the host&#39;s /var/log nor a node agent, and the YbcLogs spec is collected from tserver pods only.  | [optional] 
**BashComponentSpecs** | Pointer to [**[]BashComponentSpec**](BashComponentSpec.md) | Specs driving the node-level BashComponent. Required when components contains BashComponent. On Kubernetes universes the collect_cdc_state and collect_threadz_memz_rpcz_stats specs are collected from tserver pods only, and collect_heap_snapshot skips its tserver endpoints on a master pod.  | [optional] 
**YsqlComponentSpecs** | Pointer to [**[]YSQLComponentSpec**](YSQLComponentSpec.md) | Specs driving the node-level YSQLComponent. Required when components contains YSQLComponent. Collected from tserver pods only on Kubernetes universes.  | [optional] 
**YcqlComponentSpecs** | Pointer to [**[]YCQLComponentSpec**](YCQLComponentSpec.md) | Specs driving the node-level YCQLComponent. Required when components contains YCQLComponent. Collected from tserver pods only on Kubernetes universes.  | [optional] 
**YbAdminComponentSpecs** | Pointer to [**[]YbAdminComponentSpec**](YbAdminComponentSpec.md) | Specs driving the node-level YbAdminComponent. Required when components contains YbAdminComponent. | [optional] 
**YbaComponentSpecs** | Pointer to [**[]YbaComponentSpec**](YbaComponentSpec.md) | Specs driving the global-level YBAComponent. Required when components contains YBAComponent. | [optional] 
**FilterPgAuditLogs** | Pointer to **bool** | Specifies if Postgres audit logs should be filtered out when collecting universe logs. | [optional] 
**MaxCoreFileSize** | Pointer to **int64** | Max size in bytes of the recent collected cores (if any). | [optional] 
**MaxNumRecentCores** | Pointer to **int32** | Max number of the most recent cores to collect (if any). | [optional] 
**PaDumpStartDate** | Pointer to **time.Time** | Start date to filter Perf Advisor data. | [optional] 
**PaDumpEndDate** | Pointer to **time.Time** | End date to filter Perf Advisor data. | [optional] 
**PaMetricsFormat** | Pointer to [**PrometheusMetricsFormat**](PrometheusMetricsFormat.md) |  | [optional] 
**PromDumpDownSample** | Pointer to **bool** | Whether to downsample Prometheus remote-read export data points. | [optional] 
**PromDumpStartDate** | Pointer to **time.Time** | Start date to filter prometheus metrics from. | [optional] 
**PromDumpEndDate** | Pointer to **time.Time** | End date to filter prometheus metrics till. | [optional] 
**PromExportType** | Pointer to [**PromExportType**](PromExportType.md) |  | [optional] 
**PromMetricsFormat** | Pointer to [**PrometheusMetricsFormat**](PrometheusMetricsFormat.md) |  | [optional] 
**PromQueries** | Pointer to **map[string]string** | Map of query names to custom PromQL queries to collect in promdump. | [optional] 
**PrometheusMetricsTypes** | Pointer to [**[]PrometheusMetricsType**](PrometheusMetricsType.md) | List of exports to be included in the prometheus dump. NODE_EXPORT is dropped on Kubernetes universes, where node_exporter is not scraped.  | [optional] 
**BatchDurationPromDumpMins** | Pointer to **int32** | Batch duration for the prometheus dump (in minutes). | [optional] 
**StepPromDumpSecs** | Pointer to **int32** | Metrics downsample step (in seconds) for the prometheus dump. | [optional] 

## Methods

### NewSupportBundleCreateSpec

`func NewSupportBundleCreateSpec(components []SupportBundleComponentType, ) *SupportBundleCreateSpec`

NewSupportBundleCreateSpec instantiates a new SupportBundleCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportBundleCreateSpecWithDefaults

`func NewSupportBundleCreateSpecWithDefaults() *SupportBundleCreateSpec`

NewSupportBundleCreateSpecWithDefaults instantiates a new SupportBundleCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *SupportBundleCreateSpec) GetComponents() []SupportBundleComponentType`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SupportBundleCreateSpec) GetComponentsOk() (*[]SupportBundleComponentType, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SupportBundleCreateSpec) SetComponents(v []SupportBundleComponentType)`

SetComponents sets Components field to given value.


### GetStartDate

`func (o *SupportBundleCreateSpec) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SupportBundleCreateSpec) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SupportBundleCreateSpec) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SupportBundleCreateSpec) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SupportBundleCreateSpec) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SupportBundleCreateSpec) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SupportBundleCreateSpec) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SupportBundleCreateSpec) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNodeNames

`func (o *SupportBundleCreateSpec) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *SupportBundleCreateSpec) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *SupportBundleCreateSpec) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.

### HasNodeNames

`func (o *SupportBundleCreateSpec) HasNodeNames() bool`

HasNodeNames returns a boolean if a field has been set.

### GetFilesComponentSpecs

`func (o *SupportBundleCreateSpec) GetFilesComponentSpecs() []FilesComponentSpec`

GetFilesComponentSpecs returns the FilesComponentSpecs field if non-nil, zero value otherwise.

### GetFilesComponentSpecsOk

`func (o *SupportBundleCreateSpec) GetFilesComponentSpecsOk() (*[]FilesComponentSpec, bool)`

GetFilesComponentSpecsOk returns a tuple with the FilesComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesComponentSpecs

`func (o *SupportBundleCreateSpec) SetFilesComponentSpecs(v []FilesComponentSpec)`

SetFilesComponentSpecs sets FilesComponentSpecs field to given value.

### HasFilesComponentSpecs

`func (o *SupportBundleCreateSpec) HasFilesComponentSpecs() bool`

HasFilesComponentSpecs returns a boolean if a field has been set.

### GetBashComponentSpecs

`func (o *SupportBundleCreateSpec) GetBashComponentSpecs() []BashComponentSpec`

GetBashComponentSpecs returns the BashComponentSpecs field if non-nil, zero value otherwise.

### GetBashComponentSpecsOk

`func (o *SupportBundleCreateSpec) GetBashComponentSpecsOk() (*[]BashComponentSpec, bool)`

GetBashComponentSpecsOk returns a tuple with the BashComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBashComponentSpecs

`func (o *SupportBundleCreateSpec) SetBashComponentSpecs(v []BashComponentSpec)`

SetBashComponentSpecs sets BashComponentSpecs field to given value.

### HasBashComponentSpecs

`func (o *SupportBundleCreateSpec) HasBashComponentSpecs() bool`

HasBashComponentSpecs returns a boolean if a field has been set.

### GetYsqlComponentSpecs

`func (o *SupportBundleCreateSpec) GetYsqlComponentSpecs() []YSQLComponentSpec`

GetYsqlComponentSpecs returns the YsqlComponentSpecs field if non-nil, zero value otherwise.

### GetYsqlComponentSpecsOk

`func (o *SupportBundleCreateSpec) GetYsqlComponentSpecsOk() (*[]YSQLComponentSpec, bool)`

GetYsqlComponentSpecsOk returns a tuple with the YsqlComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlComponentSpecs

`func (o *SupportBundleCreateSpec) SetYsqlComponentSpecs(v []YSQLComponentSpec)`

SetYsqlComponentSpecs sets YsqlComponentSpecs field to given value.

### HasYsqlComponentSpecs

`func (o *SupportBundleCreateSpec) HasYsqlComponentSpecs() bool`

HasYsqlComponentSpecs returns a boolean if a field has been set.

### GetYcqlComponentSpecs

`func (o *SupportBundleCreateSpec) GetYcqlComponentSpecs() []YCQLComponentSpec`

GetYcqlComponentSpecs returns the YcqlComponentSpecs field if non-nil, zero value otherwise.

### GetYcqlComponentSpecsOk

`func (o *SupportBundleCreateSpec) GetYcqlComponentSpecsOk() (*[]YCQLComponentSpec, bool)`

GetYcqlComponentSpecsOk returns a tuple with the YcqlComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlComponentSpecs

`func (o *SupportBundleCreateSpec) SetYcqlComponentSpecs(v []YCQLComponentSpec)`

SetYcqlComponentSpecs sets YcqlComponentSpecs field to given value.

### HasYcqlComponentSpecs

`func (o *SupportBundleCreateSpec) HasYcqlComponentSpecs() bool`

HasYcqlComponentSpecs returns a boolean if a field has been set.

### GetYbAdminComponentSpecs

`func (o *SupportBundleCreateSpec) GetYbAdminComponentSpecs() []YbAdminComponentSpec`

GetYbAdminComponentSpecs returns the YbAdminComponentSpecs field if non-nil, zero value otherwise.

### GetYbAdminComponentSpecsOk

`func (o *SupportBundleCreateSpec) GetYbAdminComponentSpecsOk() (*[]YbAdminComponentSpec, bool)`

GetYbAdminComponentSpecsOk returns a tuple with the YbAdminComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbAdminComponentSpecs

`func (o *SupportBundleCreateSpec) SetYbAdminComponentSpecs(v []YbAdminComponentSpec)`

SetYbAdminComponentSpecs sets YbAdminComponentSpecs field to given value.

### HasYbAdminComponentSpecs

`func (o *SupportBundleCreateSpec) HasYbAdminComponentSpecs() bool`

HasYbAdminComponentSpecs returns a boolean if a field has been set.

### GetYbaComponentSpecs

`func (o *SupportBundleCreateSpec) GetYbaComponentSpecs() []YbaComponentSpec`

GetYbaComponentSpecs returns the YbaComponentSpecs field if non-nil, zero value otherwise.

### GetYbaComponentSpecsOk

`func (o *SupportBundleCreateSpec) GetYbaComponentSpecsOk() (*[]YbaComponentSpec, bool)`

GetYbaComponentSpecsOk returns a tuple with the YbaComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbaComponentSpecs

`func (o *SupportBundleCreateSpec) SetYbaComponentSpecs(v []YbaComponentSpec)`

SetYbaComponentSpecs sets YbaComponentSpecs field to given value.

### HasYbaComponentSpecs

`func (o *SupportBundleCreateSpec) HasYbaComponentSpecs() bool`

HasYbaComponentSpecs returns a boolean if a field has been set.

### GetFilterPgAuditLogs

`func (o *SupportBundleCreateSpec) GetFilterPgAuditLogs() bool`

GetFilterPgAuditLogs returns the FilterPgAuditLogs field if non-nil, zero value otherwise.

### GetFilterPgAuditLogsOk

`func (o *SupportBundleCreateSpec) GetFilterPgAuditLogsOk() (*bool, bool)`

GetFilterPgAuditLogsOk returns a tuple with the FilterPgAuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPgAuditLogs

`func (o *SupportBundleCreateSpec) SetFilterPgAuditLogs(v bool)`

SetFilterPgAuditLogs sets FilterPgAuditLogs field to given value.

### HasFilterPgAuditLogs

`func (o *SupportBundleCreateSpec) HasFilterPgAuditLogs() bool`

HasFilterPgAuditLogs returns a boolean if a field has been set.

### GetMaxCoreFileSize

`func (o *SupportBundleCreateSpec) GetMaxCoreFileSize() int64`

GetMaxCoreFileSize returns the MaxCoreFileSize field if non-nil, zero value otherwise.

### GetMaxCoreFileSizeOk

`func (o *SupportBundleCreateSpec) GetMaxCoreFileSizeOk() (*int64, bool)`

GetMaxCoreFileSizeOk returns a tuple with the MaxCoreFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoreFileSize

`func (o *SupportBundleCreateSpec) SetMaxCoreFileSize(v int64)`

SetMaxCoreFileSize sets MaxCoreFileSize field to given value.

### HasMaxCoreFileSize

`func (o *SupportBundleCreateSpec) HasMaxCoreFileSize() bool`

HasMaxCoreFileSize returns a boolean if a field has been set.

### GetMaxNumRecentCores

`func (o *SupportBundleCreateSpec) GetMaxNumRecentCores() int32`

GetMaxNumRecentCores returns the MaxNumRecentCores field if non-nil, zero value otherwise.

### GetMaxNumRecentCoresOk

`func (o *SupportBundleCreateSpec) GetMaxNumRecentCoresOk() (*int32, bool)`

GetMaxNumRecentCoresOk returns a tuple with the MaxNumRecentCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumRecentCores

`func (o *SupportBundleCreateSpec) SetMaxNumRecentCores(v int32)`

SetMaxNumRecentCores sets MaxNumRecentCores field to given value.

### HasMaxNumRecentCores

`func (o *SupportBundleCreateSpec) HasMaxNumRecentCores() bool`

HasMaxNumRecentCores returns a boolean if a field has been set.

### GetPaDumpStartDate

`func (o *SupportBundleCreateSpec) GetPaDumpStartDate() time.Time`

GetPaDumpStartDate returns the PaDumpStartDate field if non-nil, zero value otherwise.

### GetPaDumpStartDateOk

`func (o *SupportBundleCreateSpec) GetPaDumpStartDateOk() (*time.Time, bool)`

GetPaDumpStartDateOk returns a tuple with the PaDumpStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDumpStartDate

`func (o *SupportBundleCreateSpec) SetPaDumpStartDate(v time.Time)`

SetPaDumpStartDate sets PaDumpStartDate field to given value.

### HasPaDumpStartDate

`func (o *SupportBundleCreateSpec) HasPaDumpStartDate() bool`

HasPaDumpStartDate returns a boolean if a field has been set.

### GetPaDumpEndDate

`func (o *SupportBundleCreateSpec) GetPaDumpEndDate() time.Time`

GetPaDumpEndDate returns the PaDumpEndDate field if non-nil, zero value otherwise.

### GetPaDumpEndDateOk

`func (o *SupportBundleCreateSpec) GetPaDumpEndDateOk() (*time.Time, bool)`

GetPaDumpEndDateOk returns a tuple with the PaDumpEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDumpEndDate

`func (o *SupportBundleCreateSpec) SetPaDumpEndDate(v time.Time)`

SetPaDumpEndDate sets PaDumpEndDate field to given value.

### HasPaDumpEndDate

`func (o *SupportBundleCreateSpec) HasPaDumpEndDate() bool`

HasPaDumpEndDate returns a boolean if a field has been set.

### GetPaMetricsFormat

`func (o *SupportBundleCreateSpec) GetPaMetricsFormat() PrometheusMetricsFormat`

GetPaMetricsFormat returns the PaMetricsFormat field if non-nil, zero value otherwise.

### GetPaMetricsFormatOk

`func (o *SupportBundleCreateSpec) GetPaMetricsFormatOk() (*PrometheusMetricsFormat, bool)`

GetPaMetricsFormatOk returns a tuple with the PaMetricsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaMetricsFormat

`func (o *SupportBundleCreateSpec) SetPaMetricsFormat(v PrometheusMetricsFormat)`

SetPaMetricsFormat sets PaMetricsFormat field to given value.

### HasPaMetricsFormat

`func (o *SupportBundleCreateSpec) HasPaMetricsFormat() bool`

HasPaMetricsFormat returns a boolean if a field has been set.

### GetPromDumpDownSample

`func (o *SupportBundleCreateSpec) GetPromDumpDownSample() bool`

GetPromDumpDownSample returns the PromDumpDownSample field if non-nil, zero value otherwise.

### GetPromDumpDownSampleOk

`func (o *SupportBundleCreateSpec) GetPromDumpDownSampleOk() (*bool, bool)`

GetPromDumpDownSampleOk returns a tuple with the PromDumpDownSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpDownSample

`func (o *SupportBundleCreateSpec) SetPromDumpDownSample(v bool)`

SetPromDumpDownSample sets PromDumpDownSample field to given value.

### HasPromDumpDownSample

`func (o *SupportBundleCreateSpec) HasPromDumpDownSample() bool`

HasPromDumpDownSample returns a boolean if a field has been set.

### GetPromDumpStartDate

`func (o *SupportBundleCreateSpec) GetPromDumpStartDate() time.Time`

GetPromDumpStartDate returns the PromDumpStartDate field if non-nil, zero value otherwise.

### GetPromDumpStartDateOk

`func (o *SupportBundleCreateSpec) GetPromDumpStartDateOk() (*time.Time, bool)`

GetPromDumpStartDateOk returns a tuple with the PromDumpStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpStartDate

`func (o *SupportBundleCreateSpec) SetPromDumpStartDate(v time.Time)`

SetPromDumpStartDate sets PromDumpStartDate field to given value.

### HasPromDumpStartDate

`func (o *SupportBundleCreateSpec) HasPromDumpStartDate() bool`

HasPromDumpStartDate returns a boolean if a field has been set.

### GetPromDumpEndDate

`func (o *SupportBundleCreateSpec) GetPromDumpEndDate() time.Time`

GetPromDumpEndDate returns the PromDumpEndDate field if non-nil, zero value otherwise.

### GetPromDumpEndDateOk

`func (o *SupportBundleCreateSpec) GetPromDumpEndDateOk() (*time.Time, bool)`

GetPromDumpEndDateOk returns a tuple with the PromDumpEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpEndDate

`func (o *SupportBundleCreateSpec) SetPromDumpEndDate(v time.Time)`

SetPromDumpEndDate sets PromDumpEndDate field to given value.

### HasPromDumpEndDate

`func (o *SupportBundleCreateSpec) HasPromDumpEndDate() bool`

HasPromDumpEndDate returns a boolean if a field has been set.

### GetPromExportType

`func (o *SupportBundleCreateSpec) GetPromExportType() PromExportType`

GetPromExportType returns the PromExportType field if non-nil, zero value otherwise.

### GetPromExportTypeOk

`func (o *SupportBundleCreateSpec) GetPromExportTypeOk() (*PromExportType, bool)`

GetPromExportTypeOk returns a tuple with the PromExportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromExportType

`func (o *SupportBundleCreateSpec) SetPromExportType(v PromExportType)`

SetPromExportType sets PromExportType field to given value.

### HasPromExportType

`func (o *SupportBundleCreateSpec) HasPromExportType() bool`

HasPromExportType returns a boolean if a field has been set.

### GetPromMetricsFormat

`func (o *SupportBundleCreateSpec) GetPromMetricsFormat() PrometheusMetricsFormat`

GetPromMetricsFormat returns the PromMetricsFormat field if non-nil, zero value otherwise.

### GetPromMetricsFormatOk

`func (o *SupportBundleCreateSpec) GetPromMetricsFormatOk() (*PrometheusMetricsFormat, bool)`

GetPromMetricsFormatOk returns a tuple with the PromMetricsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromMetricsFormat

`func (o *SupportBundleCreateSpec) SetPromMetricsFormat(v PrometheusMetricsFormat)`

SetPromMetricsFormat sets PromMetricsFormat field to given value.

### HasPromMetricsFormat

`func (o *SupportBundleCreateSpec) HasPromMetricsFormat() bool`

HasPromMetricsFormat returns a boolean if a field has been set.

### GetPromQueries

`func (o *SupportBundleCreateSpec) GetPromQueries() map[string]string`

GetPromQueries returns the PromQueries field if non-nil, zero value otherwise.

### GetPromQueriesOk

`func (o *SupportBundleCreateSpec) GetPromQueriesOk() (*map[string]string, bool)`

GetPromQueriesOk returns a tuple with the PromQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromQueries

`func (o *SupportBundleCreateSpec) SetPromQueries(v map[string]string)`

SetPromQueries sets PromQueries field to given value.

### HasPromQueries

`func (o *SupportBundleCreateSpec) HasPromQueries() bool`

HasPromQueries returns a boolean if a field has been set.

### GetPrometheusMetricsTypes

`func (o *SupportBundleCreateSpec) GetPrometheusMetricsTypes() []PrometheusMetricsType`

GetPrometheusMetricsTypes returns the PrometheusMetricsTypes field if non-nil, zero value otherwise.

### GetPrometheusMetricsTypesOk

`func (o *SupportBundleCreateSpec) GetPrometheusMetricsTypesOk() (*[]PrometheusMetricsType, bool)`

GetPrometheusMetricsTypesOk returns a tuple with the PrometheusMetricsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusMetricsTypes

`func (o *SupportBundleCreateSpec) SetPrometheusMetricsTypes(v []PrometheusMetricsType)`

SetPrometheusMetricsTypes sets PrometheusMetricsTypes field to given value.

### HasPrometheusMetricsTypes

`func (o *SupportBundleCreateSpec) HasPrometheusMetricsTypes() bool`

HasPrometheusMetricsTypes returns a boolean if a field has been set.

### GetBatchDurationPromDumpMins

`func (o *SupportBundleCreateSpec) GetBatchDurationPromDumpMins() int32`

GetBatchDurationPromDumpMins returns the BatchDurationPromDumpMins field if non-nil, zero value otherwise.

### GetBatchDurationPromDumpMinsOk

`func (o *SupportBundleCreateSpec) GetBatchDurationPromDumpMinsOk() (*int32, bool)`

GetBatchDurationPromDumpMinsOk returns a tuple with the BatchDurationPromDumpMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchDurationPromDumpMins

`func (o *SupportBundleCreateSpec) SetBatchDurationPromDumpMins(v int32)`

SetBatchDurationPromDumpMins sets BatchDurationPromDumpMins field to given value.

### HasBatchDurationPromDumpMins

`func (o *SupportBundleCreateSpec) HasBatchDurationPromDumpMins() bool`

HasBatchDurationPromDumpMins returns a boolean if a field has been set.

### GetStepPromDumpSecs

`func (o *SupportBundleCreateSpec) GetStepPromDumpSecs() int32`

GetStepPromDumpSecs returns the StepPromDumpSecs field if non-nil, zero value otherwise.

### GetStepPromDumpSecsOk

`func (o *SupportBundleCreateSpec) GetStepPromDumpSecsOk() (*int32, bool)`

GetStepPromDumpSecsOk returns a tuple with the StepPromDumpSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepPromDumpSecs

`func (o *SupportBundleCreateSpec) SetStepPromDumpSecs(v int32)`

SetStepPromDumpSecs sets StepPromDumpSecs field to given value.

### HasStepPromDumpSecs

`func (o *SupportBundleCreateSpec) HasStepPromDumpSecs() bool`

HasStepPromDumpSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


