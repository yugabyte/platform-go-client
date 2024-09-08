# TroubleshootingPlatformDetailsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | **string** | API Token | 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**InUseStatus** | Pointer to **string** | In Use Status | [optional] 
**MetricsScrapePeriodSecs** | **int64** | Metrics Scrape Period Seconds | 
**MetricsUrl** | **string** | Metrics URL | 
**TpUrl** | **string** | Troubleshooting Platform URL | 
**Uuid** | **string** | Troubleshooting Platform UUID | [readonly] 
**YbaUrl** | **string** | YBA URL | 

## Methods

### NewTroubleshootingPlatformDetailsModel

`func NewTroubleshootingPlatformDetailsModel(apiToken string, customerUUID string, metricsScrapePeriodSecs int64, metricsUrl string, tpUrl string, uuid string, ybaUrl string, ) *TroubleshootingPlatformDetailsModel`

NewTroubleshootingPlatformDetailsModel instantiates a new TroubleshootingPlatformDetailsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTroubleshootingPlatformDetailsModelWithDefaults

`func NewTroubleshootingPlatformDetailsModelWithDefaults() *TroubleshootingPlatformDetailsModel`

NewTroubleshootingPlatformDetailsModelWithDefaults instantiates a new TroubleshootingPlatformDetailsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *TroubleshootingPlatformDetailsModel) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *TroubleshootingPlatformDetailsModel) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *TroubleshootingPlatformDetailsModel) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetCustomerUUID

`func (o *TroubleshootingPlatformDetailsModel) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *TroubleshootingPlatformDetailsModel) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *TroubleshootingPlatformDetailsModel) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetInUseStatus

`func (o *TroubleshootingPlatformDetailsModel) GetInUseStatus() string`

GetInUseStatus returns the InUseStatus field if non-nil, zero value otherwise.

### GetInUseStatusOk

`func (o *TroubleshootingPlatformDetailsModel) GetInUseStatusOk() (*string, bool)`

GetInUseStatusOk returns a tuple with the InUseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUseStatus

`func (o *TroubleshootingPlatformDetailsModel) SetInUseStatus(v string)`

SetInUseStatus sets InUseStatus field to given value.

### HasInUseStatus

`func (o *TroubleshootingPlatformDetailsModel) HasInUseStatus() bool`

HasInUseStatus returns a boolean if a field has been set.

### GetMetricsScrapePeriodSecs

`func (o *TroubleshootingPlatformDetailsModel) GetMetricsScrapePeriodSecs() int64`

GetMetricsScrapePeriodSecs returns the MetricsScrapePeriodSecs field if non-nil, zero value otherwise.

### GetMetricsScrapePeriodSecsOk

`func (o *TroubleshootingPlatformDetailsModel) GetMetricsScrapePeriodSecsOk() (*int64, bool)`

GetMetricsScrapePeriodSecsOk returns a tuple with the MetricsScrapePeriodSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsScrapePeriodSecs

`func (o *TroubleshootingPlatformDetailsModel) SetMetricsScrapePeriodSecs(v int64)`

SetMetricsScrapePeriodSecs sets MetricsScrapePeriodSecs field to given value.


### GetMetricsUrl

`func (o *TroubleshootingPlatformDetailsModel) GetMetricsUrl() string`

GetMetricsUrl returns the MetricsUrl field if non-nil, zero value otherwise.

### GetMetricsUrlOk

`func (o *TroubleshootingPlatformDetailsModel) GetMetricsUrlOk() (*string, bool)`

GetMetricsUrlOk returns a tuple with the MetricsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUrl

`func (o *TroubleshootingPlatformDetailsModel) SetMetricsUrl(v string)`

SetMetricsUrl sets MetricsUrl field to given value.


### GetTpUrl

`func (o *TroubleshootingPlatformDetailsModel) GetTpUrl() string`

GetTpUrl returns the TpUrl field if non-nil, zero value otherwise.

### GetTpUrlOk

`func (o *TroubleshootingPlatformDetailsModel) GetTpUrlOk() (*string, bool)`

GetTpUrlOk returns a tuple with the TpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpUrl

`func (o *TroubleshootingPlatformDetailsModel) SetTpUrl(v string)`

SetTpUrl sets TpUrl field to given value.


### GetUuid

`func (o *TroubleshootingPlatformDetailsModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TroubleshootingPlatformDetailsModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TroubleshootingPlatformDetailsModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetYbaUrl

`func (o *TroubleshootingPlatformDetailsModel) GetYbaUrl() string`

GetYbaUrl returns the YbaUrl field if non-nil, zero value otherwise.

### GetYbaUrlOk

`func (o *TroubleshootingPlatformDetailsModel) GetYbaUrlOk() (*string, bool)`

GetYbaUrlOk returns a tuple with the YbaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbaUrl

`func (o *TroubleshootingPlatformDetailsModel) SetYbaUrl(v string)`

SetYbaUrl sets YbaUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


