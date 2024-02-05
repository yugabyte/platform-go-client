# SupportBundleFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | **[]string** | List of components to be included in the support bundle | 
**EndDate** | **time.Time** | End date to filter logs till | 
**MaxCoreFileSize** | Pointer to **int64** | Max size in bytes of the recent collected cores (if any) | [optional] 
**MaxNumRecentCores** | Pointer to **int32** | Max number of the most recent cores to collect (if any) | [optional] 
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


