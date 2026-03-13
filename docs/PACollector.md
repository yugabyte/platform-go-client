# PACollector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | **string** | YBA API Token | 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**MetricsPassword** | Pointer to **string** | Metrics API Password | [optional] 
**MetricsScrapePeriodSecs** | **int64** | Metrics Scrape Period Seconds | 
**MetricsUrl** | **string** | Metrics URL | 
**MetricsUsername** | Pointer to **string** | Metrics API Username | [optional] 
**PaApiToken** | Pointer to **string** | PA Collector API Token | [optional] 
**PaUrl** | **string** | PA Collector URL | 
**Uuid** | **string** | PA Collector UUID | [readonly] 
**YbaUrl** | **string** | YBA URL | 

## Methods

### NewPACollector

`func NewPACollector(apiToken string, customerUUID string, metricsScrapePeriodSecs int64, metricsUrl string, paUrl string, uuid string, ybaUrl string, ) *PACollector`

NewPACollector instantiates a new PACollector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPACollectorWithDefaults

`func NewPACollectorWithDefaults() *PACollector`

NewPACollectorWithDefaults instantiates a new PACollector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *PACollector) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *PACollector) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *PACollector) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetCustomerUUID

`func (o *PACollector) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *PACollector) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *PACollector) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetMetricsPassword

`func (o *PACollector) GetMetricsPassword() string`

GetMetricsPassword returns the MetricsPassword field if non-nil, zero value otherwise.

### GetMetricsPasswordOk

`func (o *PACollector) GetMetricsPasswordOk() (*string, bool)`

GetMetricsPasswordOk returns a tuple with the MetricsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPassword

`func (o *PACollector) SetMetricsPassword(v string)`

SetMetricsPassword sets MetricsPassword field to given value.

### HasMetricsPassword

`func (o *PACollector) HasMetricsPassword() bool`

HasMetricsPassword returns a boolean if a field has been set.

### GetMetricsScrapePeriodSecs

`func (o *PACollector) GetMetricsScrapePeriodSecs() int64`

GetMetricsScrapePeriodSecs returns the MetricsScrapePeriodSecs field if non-nil, zero value otherwise.

### GetMetricsScrapePeriodSecsOk

`func (o *PACollector) GetMetricsScrapePeriodSecsOk() (*int64, bool)`

GetMetricsScrapePeriodSecsOk returns a tuple with the MetricsScrapePeriodSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsScrapePeriodSecs

`func (o *PACollector) SetMetricsScrapePeriodSecs(v int64)`

SetMetricsScrapePeriodSecs sets MetricsScrapePeriodSecs field to given value.


### GetMetricsUrl

`func (o *PACollector) GetMetricsUrl() string`

GetMetricsUrl returns the MetricsUrl field if non-nil, zero value otherwise.

### GetMetricsUrlOk

`func (o *PACollector) GetMetricsUrlOk() (*string, bool)`

GetMetricsUrlOk returns a tuple with the MetricsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUrl

`func (o *PACollector) SetMetricsUrl(v string)`

SetMetricsUrl sets MetricsUrl field to given value.


### GetMetricsUsername

`func (o *PACollector) GetMetricsUsername() string`

GetMetricsUsername returns the MetricsUsername field if non-nil, zero value otherwise.

### GetMetricsUsernameOk

`func (o *PACollector) GetMetricsUsernameOk() (*string, bool)`

GetMetricsUsernameOk returns a tuple with the MetricsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUsername

`func (o *PACollector) SetMetricsUsername(v string)`

SetMetricsUsername sets MetricsUsername field to given value.

### HasMetricsUsername

`func (o *PACollector) HasMetricsUsername() bool`

HasMetricsUsername returns a boolean if a field has been set.

### GetPaApiToken

`func (o *PACollector) GetPaApiToken() string`

GetPaApiToken returns the PaApiToken field if non-nil, zero value otherwise.

### GetPaApiTokenOk

`func (o *PACollector) GetPaApiTokenOk() (*string, bool)`

GetPaApiTokenOk returns a tuple with the PaApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaApiToken

`func (o *PACollector) SetPaApiToken(v string)`

SetPaApiToken sets PaApiToken field to given value.

### HasPaApiToken

`func (o *PACollector) HasPaApiToken() bool`

HasPaApiToken returns a boolean if a field has been set.

### GetPaUrl

`func (o *PACollector) GetPaUrl() string`

GetPaUrl returns the PaUrl field if non-nil, zero value otherwise.

### GetPaUrlOk

`func (o *PACollector) GetPaUrlOk() (*string, bool)`

GetPaUrlOk returns a tuple with the PaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaUrl

`func (o *PACollector) SetPaUrl(v string)`

SetPaUrl sets PaUrl field to given value.


### GetUuid

`func (o *PACollector) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PACollector) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PACollector) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetYbaUrl

`func (o *PACollector) GetYbaUrl() string`

GetYbaUrl returns the YbaUrl field if non-nil, zero value otherwise.

### GetYbaUrlOk

`func (o *PACollector) GetYbaUrlOk() (*string, bool)`

GetYbaUrlOk returns a tuple with the YbaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbaUrl

`func (o *PACollector) SetYbaUrl(v string)`

SetYbaUrl sets YbaUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


