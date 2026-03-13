# PACollectorDetailsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | **string** | YBA API Token | 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**InUseStatus** | Pointer to **string** | In Use Status | [optional] 
**MetricsPassword** | Pointer to **string** | Metrics API Password | [optional] 
**MetricsScrapePeriodSecs** | **int64** | Metrics Scrape Period Seconds | 
**MetricsUrl** | **string** | Metrics URL | 
**MetricsUsername** | Pointer to **string** | Metrics API Username | [optional] 
**PaApiToken** | Pointer to **string** | PA Collector API Token | [optional] 
**PaUrl** | **string** | PA Collector URL | 
**Uuid** | **string** | PA Collector UUID | [readonly] 
**YbaUrl** | **string** | YBA URL | 

## Methods

### NewPACollectorDetailsModel

`func NewPACollectorDetailsModel(apiToken string, customerUUID string, metricsScrapePeriodSecs int64, metricsUrl string, paUrl string, uuid string, ybaUrl string, ) *PACollectorDetailsModel`

NewPACollectorDetailsModel instantiates a new PACollectorDetailsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPACollectorDetailsModelWithDefaults

`func NewPACollectorDetailsModelWithDefaults() *PACollectorDetailsModel`

NewPACollectorDetailsModelWithDefaults instantiates a new PACollectorDetailsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *PACollectorDetailsModel) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *PACollectorDetailsModel) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *PACollectorDetailsModel) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetCustomerUUID

`func (o *PACollectorDetailsModel) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *PACollectorDetailsModel) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *PACollectorDetailsModel) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetInUseStatus

`func (o *PACollectorDetailsModel) GetInUseStatus() string`

GetInUseStatus returns the InUseStatus field if non-nil, zero value otherwise.

### GetInUseStatusOk

`func (o *PACollectorDetailsModel) GetInUseStatusOk() (*string, bool)`

GetInUseStatusOk returns a tuple with the InUseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUseStatus

`func (o *PACollectorDetailsModel) SetInUseStatus(v string)`

SetInUseStatus sets InUseStatus field to given value.

### HasInUseStatus

`func (o *PACollectorDetailsModel) HasInUseStatus() bool`

HasInUseStatus returns a boolean if a field has been set.

### GetMetricsPassword

`func (o *PACollectorDetailsModel) GetMetricsPassword() string`

GetMetricsPassword returns the MetricsPassword field if non-nil, zero value otherwise.

### GetMetricsPasswordOk

`func (o *PACollectorDetailsModel) GetMetricsPasswordOk() (*string, bool)`

GetMetricsPasswordOk returns a tuple with the MetricsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPassword

`func (o *PACollectorDetailsModel) SetMetricsPassword(v string)`

SetMetricsPassword sets MetricsPassword field to given value.

### HasMetricsPassword

`func (o *PACollectorDetailsModel) HasMetricsPassword() bool`

HasMetricsPassword returns a boolean if a field has been set.

### GetMetricsScrapePeriodSecs

`func (o *PACollectorDetailsModel) GetMetricsScrapePeriodSecs() int64`

GetMetricsScrapePeriodSecs returns the MetricsScrapePeriodSecs field if non-nil, zero value otherwise.

### GetMetricsScrapePeriodSecsOk

`func (o *PACollectorDetailsModel) GetMetricsScrapePeriodSecsOk() (*int64, bool)`

GetMetricsScrapePeriodSecsOk returns a tuple with the MetricsScrapePeriodSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsScrapePeriodSecs

`func (o *PACollectorDetailsModel) SetMetricsScrapePeriodSecs(v int64)`

SetMetricsScrapePeriodSecs sets MetricsScrapePeriodSecs field to given value.


### GetMetricsUrl

`func (o *PACollectorDetailsModel) GetMetricsUrl() string`

GetMetricsUrl returns the MetricsUrl field if non-nil, zero value otherwise.

### GetMetricsUrlOk

`func (o *PACollectorDetailsModel) GetMetricsUrlOk() (*string, bool)`

GetMetricsUrlOk returns a tuple with the MetricsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUrl

`func (o *PACollectorDetailsModel) SetMetricsUrl(v string)`

SetMetricsUrl sets MetricsUrl field to given value.


### GetMetricsUsername

`func (o *PACollectorDetailsModel) GetMetricsUsername() string`

GetMetricsUsername returns the MetricsUsername field if non-nil, zero value otherwise.

### GetMetricsUsernameOk

`func (o *PACollectorDetailsModel) GetMetricsUsernameOk() (*string, bool)`

GetMetricsUsernameOk returns a tuple with the MetricsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUsername

`func (o *PACollectorDetailsModel) SetMetricsUsername(v string)`

SetMetricsUsername sets MetricsUsername field to given value.

### HasMetricsUsername

`func (o *PACollectorDetailsModel) HasMetricsUsername() bool`

HasMetricsUsername returns a boolean if a field has been set.

### GetPaApiToken

`func (o *PACollectorDetailsModel) GetPaApiToken() string`

GetPaApiToken returns the PaApiToken field if non-nil, zero value otherwise.

### GetPaApiTokenOk

`func (o *PACollectorDetailsModel) GetPaApiTokenOk() (*string, bool)`

GetPaApiTokenOk returns a tuple with the PaApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaApiToken

`func (o *PACollectorDetailsModel) SetPaApiToken(v string)`

SetPaApiToken sets PaApiToken field to given value.

### HasPaApiToken

`func (o *PACollectorDetailsModel) HasPaApiToken() bool`

HasPaApiToken returns a boolean if a field has been set.

### GetPaUrl

`func (o *PACollectorDetailsModel) GetPaUrl() string`

GetPaUrl returns the PaUrl field if non-nil, zero value otherwise.

### GetPaUrlOk

`func (o *PACollectorDetailsModel) GetPaUrlOk() (*string, bool)`

GetPaUrlOk returns a tuple with the PaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaUrl

`func (o *PACollectorDetailsModel) SetPaUrl(v string)`

SetPaUrl sets PaUrl field to given value.


### GetUuid

`func (o *PACollectorDetailsModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PACollectorDetailsModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PACollectorDetailsModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetYbaUrl

`func (o *PACollectorDetailsModel) GetYbaUrl() string`

GetYbaUrl returns the YbaUrl field if non-nil, zero value otherwise.

### GetYbaUrlOk

`func (o *PACollectorDetailsModel) GetYbaUrlOk() (*string, bool)`

GetYbaUrlOk returns a tuple with the YbaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbaUrl

`func (o *PACollectorDetailsModel) SetYbaUrl(v string)`

SetYbaUrl sets YbaUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


