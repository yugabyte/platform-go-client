# MetricQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** | End time | [optional] 
**IsRecharts** | Pointer to **bool** | Is Recharts | [optional] 
**Metrics** | **[]string** | Metrics | 
**NodeName** | Pointer to **string** | Node name | [optional] 
**NodePrefix** | Pointer to **string** | Node prefix | [optional] 
**Start** | **int64** | Start time | 

## Methods

### NewMetricQueryParams

`func NewMetricQueryParams(metrics []string, start int64, ) *MetricQueryParams`

NewMetricQueryParams instantiates a new MetricQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricQueryParamsWithDefaults

`func NewMetricQueryParamsWithDefaults() *MetricQueryParams`

NewMetricQueryParamsWithDefaults instantiates a new MetricQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *MetricQueryParams) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MetricQueryParams) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MetricQueryParams) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MetricQueryParams) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetIsRecharts

`func (o *MetricQueryParams) GetIsRecharts() bool`

GetIsRecharts returns the IsRecharts field if non-nil, zero value otherwise.

### GetIsRechartsOk

`func (o *MetricQueryParams) GetIsRechartsOk() (*bool, bool)`

GetIsRechartsOk returns a tuple with the IsRecharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecharts

`func (o *MetricQueryParams) SetIsRecharts(v bool)`

SetIsRecharts sets IsRecharts field to given value.

### HasIsRecharts

`func (o *MetricQueryParams) HasIsRecharts() bool`

HasIsRecharts returns a boolean if a field has been set.

### GetMetrics

`func (o *MetricQueryParams) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricQueryParams) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricQueryParams) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.


### GetNodeName

`func (o *MetricQueryParams) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *MetricQueryParams) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *MetricQueryParams) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *MetricQueryParams) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodePrefix

`func (o *MetricQueryParams) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *MetricQueryParams) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *MetricQueryParams) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *MetricQueryParams) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetStart

`func (o *MetricQueryParams) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MetricQueryParams) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MetricQueryParams) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


