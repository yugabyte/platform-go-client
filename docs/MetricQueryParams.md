# MetricQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recharts** | **bool** |  | 
**XclusterConfigUuid** | **string** |  | 

## Methods

### NewMetricQueryParams

`func NewMetricQueryParams(recharts bool, xclusterConfigUuid string, ) *MetricQueryParams`

NewMetricQueryParams instantiates a new MetricQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricQueryParamsWithDefaults

`func NewMetricQueryParamsWithDefaults() *MetricQueryParams`

NewMetricQueryParamsWithDefaults instantiates a new MetricQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecharts

`func (o *MetricQueryParams) GetRecharts() bool`

GetRecharts returns the Recharts field if non-nil, zero value otherwise.

### GetRechartsOk

`func (o *MetricQueryParams) GetRechartsOk() (*bool, bool)`

GetRechartsOk returns a tuple with the Recharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecharts

`func (o *MetricQueryParams) SetRecharts(v bool)`

SetRecharts sets Recharts field to given value.


### GetXclusterConfigUuid

`func (o *MetricQueryParams) GetXclusterConfigUuid() string`

GetXclusterConfigUuid returns the XclusterConfigUuid field if non-nil, zero value otherwise.

### GetXclusterConfigUuidOk

`func (o *MetricQueryParams) GetXclusterConfigUuidOk() (*string, bool)`

GetXclusterConfigUuidOk returns a tuple with the XclusterConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterConfigUuid

`func (o *MetricQueryParams) SetXclusterConfigUuid(v string)`

SetXclusterConfigUuid sets XclusterConfigUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


