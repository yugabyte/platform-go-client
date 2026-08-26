# BundleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BashComponentSpecs** | Pointer to [**[]BashComponentSpec**](BashComponentSpec.md) | Specs that drive the node-level BashComponent (if requested). | [optional] 
**Components** | **[]string** |  | 
**FilesComponentSpecs** | Pointer to [**[]FilesComponentSpec**](FilesComponentSpec.md) | Specs that drive the generic node-level FilesComponent (if requested). | [optional] 
**MaxCoreFileSize** | Pointer to **int64** | Max size of the collected cores (if any) | [optional] 
**MaxNumRecentCores** | Pointer to **int32** | Max number of most recent cores to collect (if any) | [optional] 
**NodeNames** | Pointer to **[]string** | Names of the universe nodes node-level components were collected from. Empty or null means every node in the universe was considered. | [optional] 
**PaDumpEndDate** | Pointer to **time.Time** | End date to filter Perf Advisor data | [optional] 
**PaDumpStartDate** | Pointer to **time.Time** | Start date to filter Perf Advisor data | [optional] 
**PaMetricsFormat** | Pointer to **string** | Specifies PA Dump metrics format. | [optional] 
**PromDumpEndDate** | Pointer to **time.Time** | End date to filter prometheus metrics till | [optional] 
**PromDumpStartDate** | Pointer to **time.Time** | Start date to filter prometheus metrics from | [optional] 
**PromMetricsFormat** | Pointer to **string** | Specifies Prom Dump metrics format. | [optional] 
**PromMetricsStepSec** | Pointer to **int32** | Specifies Prom Dump metrics step in seconds. | [optional] 
**PrometheusMetricsTypes** | Pointer to **[]string** | List of exports to be included in the prometheus dump | [optional] 
**YbAdminComponentSpecs** | Pointer to [**[]YbAdminComponentSpec**](YbAdminComponentSpec.md) | Specs that drive the node-level YbAdminComponent (if requested). | [optional] 
**YbaComponentSpecs** | Pointer to [**[]YbaComponentSpec**](YbaComponentSpec.md) | Specs that drive the global-level YBAComponent (if requested). | [optional] 
**YcqlComponentSpecs** | Pointer to [**[]YCQLComponentSpec**](YCQLComponentSpec.md) | Specs that drive the node-level YCQLComponent (if requested). | [optional] 
**YsqlComponentSpecs** | Pointer to [**[]YSQLComponentSpec**](YSQLComponentSpec.md) | Specs that drive the node-level YSQLComponent (if requested). | [optional] 

## Methods

### NewBundleDetails

`func NewBundleDetails(components []string, ) *BundleDetails`

NewBundleDetails instantiates a new BundleDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleDetailsWithDefaults

`func NewBundleDetailsWithDefaults() *BundleDetails`

NewBundleDetailsWithDefaults instantiates a new BundleDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBashComponentSpecs

`func (o *BundleDetails) GetBashComponentSpecs() []BashComponentSpec`

GetBashComponentSpecs returns the BashComponentSpecs field if non-nil, zero value otherwise.

### GetBashComponentSpecsOk

`func (o *BundleDetails) GetBashComponentSpecsOk() (*[]BashComponentSpec, bool)`

GetBashComponentSpecsOk returns a tuple with the BashComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBashComponentSpecs

`func (o *BundleDetails) SetBashComponentSpecs(v []BashComponentSpec)`

SetBashComponentSpecs sets BashComponentSpecs field to given value.

### HasBashComponentSpecs

`func (o *BundleDetails) HasBashComponentSpecs() bool`

HasBashComponentSpecs returns a boolean if a field has been set.

### GetComponents

`func (o *BundleDetails) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BundleDetails) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BundleDetails) SetComponents(v []string)`

SetComponents sets Components field to given value.


### GetFilesComponentSpecs

`func (o *BundleDetails) GetFilesComponentSpecs() []FilesComponentSpec`

GetFilesComponentSpecs returns the FilesComponentSpecs field if non-nil, zero value otherwise.

### GetFilesComponentSpecsOk

`func (o *BundleDetails) GetFilesComponentSpecsOk() (*[]FilesComponentSpec, bool)`

GetFilesComponentSpecsOk returns a tuple with the FilesComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesComponentSpecs

`func (o *BundleDetails) SetFilesComponentSpecs(v []FilesComponentSpec)`

SetFilesComponentSpecs sets FilesComponentSpecs field to given value.

### HasFilesComponentSpecs

`func (o *BundleDetails) HasFilesComponentSpecs() bool`

HasFilesComponentSpecs returns a boolean if a field has been set.

### GetMaxCoreFileSize

`func (o *BundleDetails) GetMaxCoreFileSize() int64`

GetMaxCoreFileSize returns the MaxCoreFileSize field if non-nil, zero value otherwise.

### GetMaxCoreFileSizeOk

`func (o *BundleDetails) GetMaxCoreFileSizeOk() (*int64, bool)`

GetMaxCoreFileSizeOk returns a tuple with the MaxCoreFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoreFileSize

`func (o *BundleDetails) SetMaxCoreFileSize(v int64)`

SetMaxCoreFileSize sets MaxCoreFileSize field to given value.

### HasMaxCoreFileSize

`func (o *BundleDetails) HasMaxCoreFileSize() bool`

HasMaxCoreFileSize returns a boolean if a field has been set.

### GetMaxNumRecentCores

`func (o *BundleDetails) GetMaxNumRecentCores() int32`

GetMaxNumRecentCores returns the MaxNumRecentCores field if non-nil, zero value otherwise.

### GetMaxNumRecentCoresOk

`func (o *BundleDetails) GetMaxNumRecentCoresOk() (*int32, bool)`

GetMaxNumRecentCoresOk returns a tuple with the MaxNumRecentCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumRecentCores

`func (o *BundleDetails) SetMaxNumRecentCores(v int32)`

SetMaxNumRecentCores sets MaxNumRecentCores field to given value.

### HasMaxNumRecentCores

`func (o *BundleDetails) HasMaxNumRecentCores() bool`

HasMaxNumRecentCores returns a boolean if a field has been set.

### GetNodeNames

`func (o *BundleDetails) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *BundleDetails) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *BundleDetails) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.

### HasNodeNames

`func (o *BundleDetails) HasNodeNames() bool`

HasNodeNames returns a boolean if a field has been set.

### GetPaDumpEndDate

`func (o *BundleDetails) GetPaDumpEndDate() time.Time`

GetPaDumpEndDate returns the PaDumpEndDate field if non-nil, zero value otherwise.

### GetPaDumpEndDateOk

`func (o *BundleDetails) GetPaDumpEndDateOk() (*time.Time, bool)`

GetPaDumpEndDateOk returns a tuple with the PaDumpEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDumpEndDate

`func (o *BundleDetails) SetPaDumpEndDate(v time.Time)`

SetPaDumpEndDate sets PaDumpEndDate field to given value.

### HasPaDumpEndDate

`func (o *BundleDetails) HasPaDumpEndDate() bool`

HasPaDumpEndDate returns a boolean if a field has been set.

### GetPaDumpStartDate

`func (o *BundleDetails) GetPaDumpStartDate() time.Time`

GetPaDumpStartDate returns the PaDumpStartDate field if non-nil, zero value otherwise.

### GetPaDumpStartDateOk

`func (o *BundleDetails) GetPaDumpStartDateOk() (*time.Time, bool)`

GetPaDumpStartDateOk returns a tuple with the PaDumpStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDumpStartDate

`func (o *BundleDetails) SetPaDumpStartDate(v time.Time)`

SetPaDumpStartDate sets PaDumpStartDate field to given value.

### HasPaDumpStartDate

`func (o *BundleDetails) HasPaDumpStartDate() bool`

HasPaDumpStartDate returns a boolean if a field has been set.

### GetPaMetricsFormat

`func (o *BundleDetails) GetPaMetricsFormat() string`

GetPaMetricsFormat returns the PaMetricsFormat field if non-nil, zero value otherwise.

### GetPaMetricsFormatOk

`func (o *BundleDetails) GetPaMetricsFormatOk() (*string, bool)`

GetPaMetricsFormatOk returns a tuple with the PaMetricsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaMetricsFormat

`func (o *BundleDetails) SetPaMetricsFormat(v string)`

SetPaMetricsFormat sets PaMetricsFormat field to given value.

### HasPaMetricsFormat

`func (o *BundleDetails) HasPaMetricsFormat() bool`

HasPaMetricsFormat returns a boolean if a field has been set.

### GetPromDumpEndDate

`func (o *BundleDetails) GetPromDumpEndDate() time.Time`

GetPromDumpEndDate returns the PromDumpEndDate field if non-nil, zero value otherwise.

### GetPromDumpEndDateOk

`func (o *BundleDetails) GetPromDumpEndDateOk() (*time.Time, bool)`

GetPromDumpEndDateOk returns a tuple with the PromDumpEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpEndDate

`func (o *BundleDetails) SetPromDumpEndDate(v time.Time)`

SetPromDumpEndDate sets PromDumpEndDate field to given value.

### HasPromDumpEndDate

`func (o *BundleDetails) HasPromDumpEndDate() bool`

HasPromDumpEndDate returns a boolean if a field has been set.

### GetPromDumpStartDate

`func (o *BundleDetails) GetPromDumpStartDate() time.Time`

GetPromDumpStartDate returns the PromDumpStartDate field if non-nil, zero value otherwise.

### GetPromDumpStartDateOk

`func (o *BundleDetails) GetPromDumpStartDateOk() (*time.Time, bool)`

GetPromDumpStartDateOk returns a tuple with the PromDumpStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromDumpStartDate

`func (o *BundleDetails) SetPromDumpStartDate(v time.Time)`

SetPromDumpStartDate sets PromDumpStartDate field to given value.

### HasPromDumpStartDate

`func (o *BundleDetails) HasPromDumpStartDate() bool`

HasPromDumpStartDate returns a boolean if a field has been set.

### GetPromMetricsFormat

`func (o *BundleDetails) GetPromMetricsFormat() string`

GetPromMetricsFormat returns the PromMetricsFormat field if non-nil, zero value otherwise.

### GetPromMetricsFormatOk

`func (o *BundleDetails) GetPromMetricsFormatOk() (*string, bool)`

GetPromMetricsFormatOk returns a tuple with the PromMetricsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromMetricsFormat

`func (o *BundleDetails) SetPromMetricsFormat(v string)`

SetPromMetricsFormat sets PromMetricsFormat field to given value.

### HasPromMetricsFormat

`func (o *BundleDetails) HasPromMetricsFormat() bool`

HasPromMetricsFormat returns a boolean if a field has been set.

### GetPromMetricsStepSec

`func (o *BundleDetails) GetPromMetricsStepSec() int32`

GetPromMetricsStepSec returns the PromMetricsStepSec field if non-nil, zero value otherwise.

### GetPromMetricsStepSecOk

`func (o *BundleDetails) GetPromMetricsStepSecOk() (*int32, bool)`

GetPromMetricsStepSecOk returns a tuple with the PromMetricsStepSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromMetricsStepSec

`func (o *BundleDetails) SetPromMetricsStepSec(v int32)`

SetPromMetricsStepSec sets PromMetricsStepSec field to given value.

### HasPromMetricsStepSec

`func (o *BundleDetails) HasPromMetricsStepSec() bool`

HasPromMetricsStepSec returns a boolean if a field has been set.

### GetPrometheusMetricsTypes

`func (o *BundleDetails) GetPrometheusMetricsTypes() []string`

GetPrometheusMetricsTypes returns the PrometheusMetricsTypes field if non-nil, zero value otherwise.

### GetPrometheusMetricsTypesOk

`func (o *BundleDetails) GetPrometheusMetricsTypesOk() (*[]string, bool)`

GetPrometheusMetricsTypesOk returns a tuple with the PrometheusMetricsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusMetricsTypes

`func (o *BundleDetails) SetPrometheusMetricsTypes(v []string)`

SetPrometheusMetricsTypes sets PrometheusMetricsTypes field to given value.

### HasPrometheusMetricsTypes

`func (o *BundleDetails) HasPrometheusMetricsTypes() bool`

HasPrometheusMetricsTypes returns a boolean if a field has been set.

### GetYbAdminComponentSpecs

`func (o *BundleDetails) GetYbAdminComponentSpecs() []YbAdminComponentSpec`

GetYbAdminComponentSpecs returns the YbAdminComponentSpecs field if non-nil, zero value otherwise.

### GetYbAdminComponentSpecsOk

`func (o *BundleDetails) GetYbAdminComponentSpecsOk() (*[]YbAdminComponentSpec, bool)`

GetYbAdminComponentSpecsOk returns a tuple with the YbAdminComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbAdminComponentSpecs

`func (o *BundleDetails) SetYbAdminComponentSpecs(v []YbAdminComponentSpec)`

SetYbAdminComponentSpecs sets YbAdminComponentSpecs field to given value.

### HasYbAdminComponentSpecs

`func (o *BundleDetails) HasYbAdminComponentSpecs() bool`

HasYbAdminComponentSpecs returns a boolean if a field has been set.

### GetYbaComponentSpecs

`func (o *BundleDetails) GetYbaComponentSpecs() []YbaComponentSpec`

GetYbaComponentSpecs returns the YbaComponentSpecs field if non-nil, zero value otherwise.

### GetYbaComponentSpecsOk

`func (o *BundleDetails) GetYbaComponentSpecsOk() (*[]YbaComponentSpec, bool)`

GetYbaComponentSpecsOk returns a tuple with the YbaComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbaComponentSpecs

`func (o *BundleDetails) SetYbaComponentSpecs(v []YbaComponentSpec)`

SetYbaComponentSpecs sets YbaComponentSpecs field to given value.

### HasYbaComponentSpecs

`func (o *BundleDetails) HasYbaComponentSpecs() bool`

HasYbaComponentSpecs returns a boolean if a field has been set.

### GetYcqlComponentSpecs

`func (o *BundleDetails) GetYcqlComponentSpecs() []YCQLComponentSpec`

GetYcqlComponentSpecs returns the YcqlComponentSpecs field if non-nil, zero value otherwise.

### GetYcqlComponentSpecsOk

`func (o *BundleDetails) GetYcqlComponentSpecsOk() (*[]YCQLComponentSpec, bool)`

GetYcqlComponentSpecsOk returns a tuple with the YcqlComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlComponentSpecs

`func (o *BundleDetails) SetYcqlComponentSpecs(v []YCQLComponentSpec)`

SetYcqlComponentSpecs sets YcqlComponentSpecs field to given value.

### HasYcqlComponentSpecs

`func (o *BundleDetails) HasYcqlComponentSpecs() bool`

HasYcqlComponentSpecs returns a boolean if a field has been set.

### GetYsqlComponentSpecs

`func (o *BundleDetails) GetYsqlComponentSpecs() []YSQLComponentSpec`

GetYsqlComponentSpecs returns the YsqlComponentSpecs field if non-nil, zero value otherwise.

### GetYsqlComponentSpecsOk

`func (o *BundleDetails) GetYsqlComponentSpecsOk() (*[]YSQLComponentSpec, bool)`

GetYsqlComponentSpecsOk returns a tuple with the YsqlComponentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlComponentSpecs

`func (o *BundleDetails) SetYsqlComponentSpecs(v []YSQLComponentSpec)`

SetYsqlComponentSpecs sets YsqlComponentSpecs field to given value.

### HasYsqlComponentSpecs

`func (o *BundleDetails) HasYsqlComponentSpecs() bool`

HasYsqlComponentSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


