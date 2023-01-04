# SupportBundleFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | **[]string** | List of components to be included in the support bundle | 
**EndDate** | **time.Time** | End date to filter logs till | 
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


