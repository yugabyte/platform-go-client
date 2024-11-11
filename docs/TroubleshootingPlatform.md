# TroubleshootingPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | **string** | YBA API Token | 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**MetricsScrapePeriodSecs** | **int64** | Metrics Scrape Period Seconds | 
**MetricsUrl** | **string** | Metrics URL | 
**TpApiToken** | Pointer to **string** | TP API Token | [optional] 
**TpUrl** | **string** | Troubleshooting Platform URL | 
**Uuid** | **string** | Troubleshooting Platform UUID | [readonly] 
**YbaUrl** | **string** | YBA URL | 

## Methods

### NewTroubleshootingPlatform

`func NewTroubleshootingPlatform(apiToken string, customerUUID string, metricsScrapePeriodSecs int64, metricsUrl string, tpUrl string, uuid string, ybaUrl string, ) *TroubleshootingPlatform`

NewTroubleshootingPlatform instantiates a new TroubleshootingPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTroubleshootingPlatformWithDefaults

`func NewTroubleshootingPlatformWithDefaults() *TroubleshootingPlatform`

NewTroubleshootingPlatformWithDefaults instantiates a new TroubleshootingPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *TroubleshootingPlatform) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *TroubleshootingPlatform) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *TroubleshootingPlatform) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetCustomerUUID

`func (o *TroubleshootingPlatform) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *TroubleshootingPlatform) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *TroubleshootingPlatform) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetMetricsScrapePeriodSecs

`func (o *TroubleshootingPlatform) GetMetricsScrapePeriodSecs() int64`

GetMetricsScrapePeriodSecs returns the MetricsScrapePeriodSecs field if non-nil, zero value otherwise.

### GetMetricsScrapePeriodSecsOk

`func (o *TroubleshootingPlatform) GetMetricsScrapePeriodSecsOk() (*int64, bool)`

GetMetricsScrapePeriodSecsOk returns a tuple with the MetricsScrapePeriodSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsScrapePeriodSecs

`func (o *TroubleshootingPlatform) SetMetricsScrapePeriodSecs(v int64)`

SetMetricsScrapePeriodSecs sets MetricsScrapePeriodSecs field to given value.


### GetMetricsUrl

`func (o *TroubleshootingPlatform) GetMetricsUrl() string`

GetMetricsUrl returns the MetricsUrl field if non-nil, zero value otherwise.

### GetMetricsUrlOk

`func (o *TroubleshootingPlatform) GetMetricsUrlOk() (*string, bool)`

GetMetricsUrlOk returns a tuple with the MetricsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUrl

`func (o *TroubleshootingPlatform) SetMetricsUrl(v string)`

SetMetricsUrl sets MetricsUrl field to given value.


### GetTpApiToken

`func (o *TroubleshootingPlatform) GetTpApiToken() string`

GetTpApiToken returns the TpApiToken field if non-nil, zero value otherwise.

### GetTpApiTokenOk

`func (o *TroubleshootingPlatform) GetTpApiTokenOk() (*string, bool)`

GetTpApiTokenOk returns a tuple with the TpApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpApiToken

`func (o *TroubleshootingPlatform) SetTpApiToken(v string)`

SetTpApiToken sets TpApiToken field to given value.

### HasTpApiToken

`func (o *TroubleshootingPlatform) HasTpApiToken() bool`

HasTpApiToken returns a boolean if a field has been set.

### GetTpUrl

`func (o *TroubleshootingPlatform) GetTpUrl() string`

GetTpUrl returns the TpUrl field if non-nil, zero value otherwise.

### GetTpUrlOk

`func (o *TroubleshootingPlatform) GetTpUrlOk() (*string, bool)`

GetTpUrlOk returns a tuple with the TpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpUrl

`func (o *TroubleshootingPlatform) SetTpUrl(v string)`

SetTpUrl sets TpUrl field to given value.


### GetUuid

`func (o *TroubleshootingPlatform) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TroubleshootingPlatform) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TroubleshootingPlatform) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetYbaUrl

`func (o *TroubleshootingPlatform) GetYbaUrl() string`

GetYbaUrl returns the YbaUrl field if non-nil, zero value otherwise.

### GetYbaUrlOk

`func (o *TroubleshootingPlatform) GetYbaUrlOk() (*string, bool)`

GetYbaUrlOk returns a tuple with the YbaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbaUrl

`func (o *TroubleshootingPlatform) SetYbaUrl(v string)`

SetYbaUrl sets YbaUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


