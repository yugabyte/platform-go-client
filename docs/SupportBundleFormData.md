# SupportBundleFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | **[]string** | List of components to be included in the support bundle | 
**EndDate** | **time.Time** | End date to filter logs till | 
**MaxCoreFileSize** | Pointer to **int64** | Max size in bytes of the recent collected cores (if any) | [optional] 
**MaxNumRecentCores** | Pointer to **int32** | Max number of the most recent cores to collect (if any) | [optional] 
**PromDumpEndDate** | Pointer to **time.Time** | End date to filter prometheus metrics till | [optional] 
**PromDumpStartDate** | Pointer to **time.Time** | Start date to filter prometheus metrics from | [optional] 
**PromQueries** | Pointer to **map[string]string** | Map of query names to custom PromQL queries to collect in promdump | [optional] 
**PrometheusMetricsTypes** | Pointer to **[]string** | List of exports to be included in the prometheus dump | [optional] 
**StartDate** | **time.Time** | Start date to filter logs from | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


