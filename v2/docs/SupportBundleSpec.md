# SupportBundleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]SupportBundleComponentType**](SupportBundleComponentType.md) | Components collected into this support bundle. Names what the server kept after dropping whatever the universe cannot produce, so it can be a subset of what the request asked for.  | [optional] 
**StartDate** | Pointer to **time.Time** | Start date logs were filtered from. | [optional] 
**EndDate** | Pointer to **time.Time** | End date logs were filtered till. | [optional] 
**NodeNames** | Pointer to **[]string** | Universe nodes that node-level components were collected from. Empty or absent means every node in the universe was considered. Global-level components are collected once for the universe and are not scoped by this list.  | [optional] 
**FilesComponentSpecs** | Pointer to [**[]FilesComponentSpec**](FilesComponentSpec.md) | Specs that drove the generic node-level FilesComponent. | [optional] 
**BashComponentSpecs** | Pointer to [**[]BashComponentSpec**](BashComponentSpec.md) | Specs that drove the node-level BashComponent. | [optional] 
**YsqlComponentSpecs** | Pointer to [**[]YSQLComponentSpec**](YSQLComponentSpec.md) | Specs that drove the node-level YSQLComponent. | [optional] 
**YcqlComponentSpecs** | Pointer to [**[]YCQLComponentSpec**](YCQLComponentSpec.md) | Specs that drove the node-level YCQLComponent. | [optional] 
**YbAdminComponentSpecs** | Pointer to [**[]YbAdminComponentSpec**](YbAdminComponentSpec.md) | Specs that drove the node-level YbAdminComponent. | [optional] 
**YbaComponentSpecs** | Pointer to [**[]YbaComponentSpec**](YbaComponentSpec.md) | Specs that drove the global-level YBAComponent. | [optional] 
**MaxCoreFileSize** | Pointer to **int64** | Max size in bytes of the collected cores (if any). | [optional] 
**MaxNumRecentCores** | Pointer to **int32** | Max number of most recent cores collected (if any). | [optional] 
**PromDumpStartDate** | Pointer to **time.Time** | Start date Prometheus metrics were filtered from. | [optional] 
**PromDumpEndDate** | Pointer to **time.Time** | End date Prometheus metrics were filtered till. | [optional] 
**PromMetricsFormat** | Pointer to [**PrometheusMetricsFormat**](PrometheusMetricsFormat.md) |  | [optional] 
**PromMetricsStepSec** | Pointer to **int32** | Downsample step in seconds applied to the Prometheus metrics dump. | [optional] 
**PrometheusMetricsTypes** | Pointer to [**[]PrometheusMetricsType**](PrometheusMetricsType.md) | Prometheus exports included in the metrics dump. | [optional] 
**PaDumpStartDate** | Pointer to **time.Time** | Start date Perf Advisor data was filtered from. | [optional] 
**PaDumpEndDate** | Pointer to **time.Time** | End date Perf Advisor data was filtered till. | [optional] 
**PaMetricsFormat** | Pointer to [**PrometheusMetricsFormat**](PrometheusMetricsFormat.md) |  | [optional] 

## Methods

### NewSupportBundleSpec

`func NewSupportBundleSpec() *SupportBundleSpec`

NewSupportBundleSpec instantiates a new SupportBundleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportBundleSpecWithDefaults

`func NewSupportBundleSpecWithDefaults() *SupportBundleSpec`

NewSupportBundleSpecWithDefaults instantiates a new SupportBundleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *SupportBundleSpec) GetComponents() []SupportBundleComponentType`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SupportBundleSpec) GetComponentsOk() (*[]SupportBundleComponentType, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SupportBundleSpec) SetComponents(v []SupportBundleComponentType)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SupportBundleSpec) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetStartDate

`func (o *SupportBundleSpec) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SupportBundleSpec) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SupportBundleSpec) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SupportBundleSpec) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SupportBundleSpec) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SupportBundleSpec) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SupportBundleSpec) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SupportBundleSpec) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNodeNames

`func (o *SupportBundleSpec) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *SupportBundleSpec) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *SupportBundleSpec) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.

### HasNodeNames

`func (o *SupportBundleSpec) HasNodeNames() bool`

HasNodeNames returns a boolean if a field has been set.

### GetFilesComponentSpecs

`func (o *SupportBundleSpec) GetFilesComponentSpecs() []FilesComponentSpec`

GetFilesComponentSpecs returns the FilesComponentSpecs field if non-nil, zero value otherwise.

### GetFilesComponentSpecsOk

`func (o *SupportBundleSpec) GetFilesComponentSpecsOk() (*[]FilesComponentSpec, bool)`

GetFilesComponentSpecsOk returns a tuple with the FilesComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesComponentSpecs

`func (o *SupportBundleSpec) SetFilesComponentSpecs(v []FilesComponentSpec)`

SetFilesComponentSpecs sets FilesComponentSpecs field to given value.

### HasFilesComponentSpecs

`func (o *SupportBundleSpec) HasFilesComponentSpecs() bool`

HasFilesComponentSpecs returns a boolean if a field has been set.

### GetBashComponentSpecs

`func (o *SupportBundleSpec) GetBashComponentSpecs() []BashComponentSpec`

GetBashComponentSpecs returns the BashComponentSpecs field if non-nil, zero value otherwise.

### GetBashComponentSpecsOk

`func (o *SupportBundleSpec) GetBashComponentSpecsOk() (*[]BashComponentSpec, bool)`

GetBashComponentSpecsOk returns a tuple with the BashComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBashComponentSpecs

`func (o *SupportBundleSpec) SetBashComponentSpecs(v []BashComponentSpec)`

SetBashComponentSpecs sets BashComponentSpecs field to given value.

### HasBashComponentSpecs

`func (o *SupportBundleSpec) HasBashComponentSpecs() bool`

HasBashComponentSpecs returns a boolean if a field has been set.

### GetYsqlComponentSpecs

`func (o *SupportBundleSpec) GetYsqlComponentSpecs() []YSQLComponentSpec`

GetYsqlComponentSpecs returns the YsqlComponentSpecs field if non-nil, zero value otherwise.

### GetYsqlComponentSpecsOk

`func (o *SupportBundleSpec) GetYsqlComponentSpecsOk() (*[]YSQLComponentSpec, bool)`

GetYsqlComponentSpecsOk returns a tuple with the YsqlComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlComponentSpecs

`func (o *SupportBundleSpec) SetYsqlComponentSpecs(v []YSQLComponentSpec)`

SetYsqlComponentSpecs sets YsqlComponentSpecs field to given value.

### HasYsqlComponentSpecs

`func (o *SupportBundleSpec) HasYsqlComponentSpecs() bool`

HasYsqlComponentSpecs returns a boolean if a field has been set.

### GetYcqlComponentSpecs

`func (o *SupportBundleSpec) GetYcqlComponentSpecs() []YCQLComponentSpec`

GetYcqlComponentSpecs returns the YcqlComponentSpecs field if non-nil, zero value otherwise.

### GetYcqlComponentSpecsOk

`func (o *SupportBundleSpec) GetYcqlComponentSpecsOk() (*[]YCQLComponentSpec, bool)`

GetYcqlComponentSpecsOk returns a tuple with the YcqlComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlComponentSpecs

`func (o *SupportBundleSpec) SetYcqlComponentSpecs(v []YCQLComponentSpec)`

SetYcqlComponentSpecs sets YcqlComponentSpecs field to given value.

### HasYcqlComponentSpecs

`func (o *SupportBundleSpec) HasYcqlComponentSpecs() bool`

HasYcqlComponentSpecs returns a boolean if a field has been set.

### GetYbAdminComponentSpecs

`func (o *SupportBundleSpec) GetYbAdminComponentSpecs() []YbAdminComponentSpec`

GetYbAdminComponentSpecs returns the YbAdminComponentSpecs field if non-nil, zero value otherwise.

### GetYbAdminComponentSpecsOk

`func (o *SupportBundleSpec) GetYbAdminComponentSpecsOk() (*[]YbAdminComponentSpec, bool)`

GetYbAdminComponentSpecsOk returns a tuple with the YbAdminComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbAdminComponentSpecs

`func (o *SupportBundleSpec) SetYbAdminComponentSpecs(v []YbAdminComponentSpec)`

SetYbAdminComponentSpecs sets YbAdminComponentSpecs field to given value.

### HasYbAdminComponentSpecs

`func (o *SupportBundleSpec) HasYbAdminComponentSpecs() bool`

HasYbAdminComponentSpecs returns a boolean if a field has been set.

### GetYbaComponentSpecs

`func (o *SupportBundleSpec) GetYbaComponentSpecs() []YbaComponentSpec`

GetYbaComponentSpecs returns the YbaComponentSpecs field if non-nil, zero value otherwise.

### GetYbaComponentSpecsOk

`func (o *SupportBundleSpec) GetYbaComponentSpecsOk() (*[]YbaComponentSpec, bool)`

GetYbaComponentSpecsOk returns a tuple with the YbaComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbaComponentSpecs

`func (o *SupportBundleSpec) SetYbaComponentSpecs(v []YbaComponentSpec)`

SetYbaComponentSpecs sets YbaComponentSpecs field to given value.

### HasYbaComponentSpecs

`func (o *SupportBundleSpec) HasYbaComponentSpecs() bool`

HasYbaComponentSpecs returns a boolean if a field has been set.

### GetMaxCoreFileSize

`func (o *SupportBundleSpec) GetMaxCoreFileSize() int64`

GetMaxCoreFileSize returns the MaxCoreFileSize field if non-nil, zero value otherwise.

### GetMaxCoreFileSizeOk

`func (o *SupportBundleSpec) GetMaxCoreFileSizeOk() (*int64, bool)`

GetMaxCoreFileSizeOk returns a tuple with the MaxCoreFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoreFileSize

`func (o *SupportBundleSpec) SetMaxCoreFileSize(v int64)`

SetMaxCoreFileSize sets MaxCoreFileSize field to given value.

### HasMaxCoreFileSize

`func (o *SupportBundleSpec) HasMaxCoreFileSize() bool`

HasMaxCoreFileSize returns a boolean if a field has been set.

### GetMaxNumRecentCores

`func (o *SupportBundleSpec) GetMaxNumRecentCores() int32`

GetMaxNumRecentCores returns the MaxNumRecentCores field if non-nil, zero value otherwise.

### GetMaxNumRecentCoresOk

`func (o *SupportBundleSpec) GetMaxNumRecentCoresOk() (*int32, bool)`

GetMaxNumRecentCoresOk returns a tuple with the MaxNumRecentCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumRecentCores

`func (o *SupportBundleSpec) SetMaxNumRecentCores(v int32)`

SetMaxNumRecentCores sets MaxNumRecentCores field to given value.

### HasMaxNumRecentCores

`func (o *SupportBundleSpec) HasMaxNumRecentCores() bool`

HasMaxNumRecentCores returns a boolean if a field has been set.

### GetPromDumpStartDate

`func (o *SupportBundleSpec) GetPromDumpStartDate() time.Time`

GetPromDumpStartDate returns the PromDumpStartDate field if non-nil, zero value otherwise.

### GetPromDumpStartDateOk

`func (o *SupportBundleSpec) GetPromDumpStartDateOk() (*time.Time, bool)`

GetPromDumpStartDateOk returns a tuple with the PromDumpStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpStartDate

`func (o *SupportBundleSpec) SetPromDumpStartDate(v time.Time)`

SetPromDumpStartDate sets PromDumpStartDate field to given value.

### HasPromDumpStartDate

`func (o *SupportBundleSpec) HasPromDumpStartDate() bool`

HasPromDumpStartDate returns a boolean if a field has been set.

### GetPromDumpEndDate

`func (o *SupportBundleSpec) GetPromDumpEndDate() time.Time`

GetPromDumpEndDate returns the PromDumpEndDate field if non-nil, zero value otherwise.

### GetPromDumpEndDateOk

`func (o *SupportBundleSpec) GetPromDumpEndDateOk() (*time.Time, bool)`

GetPromDumpEndDateOk returns a tuple with the PromDumpEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpEndDate

`func (o *SupportBundleSpec) SetPromDumpEndDate(v time.Time)`

SetPromDumpEndDate sets PromDumpEndDate field to given value.

### HasPromDumpEndDate

`func (o *SupportBundleSpec) HasPromDumpEndDate() bool`

HasPromDumpEndDate returns a boolean if a field has been set.

### GetPromMetricsFormat

`func (o *SupportBundleSpec) GetPromMetricsFormat() PrometheusMetricsFormat`

GetPromMetricsFormat returns the PromMetricsFormat field if non-nil, zero value otherwise.

### GetPromMetricsFormatOk

`func (o *SupportBundleSpec) GetPromMetricsFormatOk() (*PrometheusMetricsFormat, bool)`

GetPromMetricsFormatOk returns a tuple with the PromMetricsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromMetricsFormat

`func (o *SupportBundleSpec) SetPromMetricsFormat(v PrometheusMetricsFormat)`

SetPromMetricsFormat sets PromMetricsFormat field to given value.

### HasPromMetricsFormat

`func (o *SupportBundleSpec) HasPromMetricsFormat() bool`

HasPromMetricsFormat returns a boolean if a field has been set.

### GetPromMetricsStepSec

`func (o *SupportBundleSpec) GetPromMetricsStepSec() int32`

GetPromMetricsStepSec returns the PromMetricsStepSec field if non-nil, zero value otherwise.

### GetPromMetricsStepSecOk

`func (o *SupportBundleSpec) GetPromMetricsStepSecOk() (*int32, bool)`

GetPromMetricsStepSecOk returns a tuple with the PromMetricsStepSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromMetricsStepSec

`func (o *SupportBundleSpec) SetPromMetricsStepSec(v int32)`

SetPromMetricsStepSec sets PromMetricsStepSec field to given value.

### HasPromMetricsStepSec

`func (o *SupportBundleSpec) HasPromMetricsStepSec() bool`

HasPromMetricsStepSec returns a boolean if a field has been set.

### GetPrometheusMetricsTypes

`func (o *SupportBundleSpec) GetPrometheusMetricsTypes() []PrometheusMetricsType`

GetPrometheusMetricsTypes returns the PrometheusMetricsTypes field if non-nil, zero value otherwise.

### GetPrometheusMetricsTypesOk

`func (o *SupportBundleSpec) GetPrometheusMetricsTypesOk() (*[]PrometheusMetricsType, bool)`

GetPrometheusMetricsTypesOk returns a tuple with the PrometheusMetricsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusMetricsTypes

`func (o *SupportBundleSpec) SetPrometheusMetricsTypes(v []PrometheusMetricsType)`

SetPrometheusMetricsTypes sets PrometheusMetricsTypes field to given value.

### HasPrometheusMetricsTypes

`func (o *SupportBundleSpec) HasPrometheusMetricsTypes() bool`

HasPrometheusMetricsTypes returns a boolean if a field has been set.

### GetPaDumpStartDate

`func (o *SupportBundleSpec) GetPaDumpStartDate() time.Time`

GetPaDumpStartDate returns the PaDumpStartDate field if non-nil, zero value otherwise.

### GetPaDumpStartDateOk

`func (o *SupportBundleSpec) GetPaDumpStartDateOk() (*time.Time, bool)`

GetPaDumpStartDateOk returns a tuple with the PaDumpStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDumpStartDate

`func (o *SupportBundleSpec) SetPaDumpStartDate(v time.Time)`

SetPaDumpStartDate sets PaDumpStartDate field to given value.

### HasPaDumpStartDate

`func (o *SupportBundleSpec) HasPaDumpStartDate() bool`

HasPaDumpStartDate returns a boolean if a field has been set.

### GetPaDumpEndDate

`func (o *SupportBundleSpec) GetPaDumpEndDate() time.Time`

GetPaDumpEndDate returns the PaDumpEndDate field if non-nil, zero value otherwise.

### GetPaDumpEndDateOk

`func (o *SupportBundleSpec) GetPaDumpEndDateOk() (*time.Time, bool)`

GetPaDumpEndDateOk returns a tuple with the PaDumpEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDumpEndDate

`func (o *SupportBundleSpec) SetPaDumpEndDate(v time.Time)`

SetPaDumpEndDate sets PaDumpEndDate field to given value.

### HasPaDumpEndDate

`func (o *SupportBundleSpec) HasPaDumpEndDate() bool`

HasPaDumpEndDate returns a boolean if a field has been set.

### GetPaMetricsFormat

`func (o *SupportBundleSpec) GetPaMetricsFormat() PrometheusMetricsFormat`

GetPaMetricsFormat returns the PaMetricsFormat field if non-nil, zero value otherwise.

### GetPaMetricsFormatOk

`func (o *SupportBundleSpec) GetPaMetricsFormatOk() (*PrometheusMetricsFormat, bool)`

GetPaMetricsFormatOk returns a tuple with the PaMetricsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaMetricsFormat

`func (o *SupportBundleSpec) SetPaMetricsFormat(v PrometheusMetricsFormat)`

SetPaMetricsFormat sets PaMetricsFormat field to given value.

### HasPaMetricsFormat

`func (o *SupportBundleSpec) HasPaMetricsFormat() bool`

HasPaMetricsFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


