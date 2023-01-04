# MetricSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | Metric name | 
**NodeAggregation** | Pointer to **string** | Way of metrics aggregation across nodes | [optional] 
**ReturnAggregatedValue** | Pointer to **bool** | Defines if we return &#39;mean&#39; line with node lines in case we split nodes | [optional] 
**SplitCount** | Pointer to **int32** | Defines how many node lines we return in case we split by nodes/tables/etc. | [optional] 
**SplitMode** | Pointer to **string** | Controls if we split nodes, tables, etc. into own lines OR aggregate  and how we select lines in case of split query | [optional] 
**SplitType** | Pointer to **string** | Defines set of labels, which we use for a split query | [optional] 
**TimeAggregation** | Pointer to **string** | Way of metrics aggregation over time | [optional] 

## Methods

### NewMetricSettings

`func NewMetricSettings(metric string, ) *MetricSettings`

NewMetricSettings instantiates a new MetricSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSettingsWithDefaults

`func NewMetricSettingsWithDefaults() *MetricSettings`

NewMetricSettingsWithDefaults instantiates a new MetricSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *MetricSettings) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricSettings) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricSettings) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetNodeAggregation

`func (o *MetricSettings) GetNodeAggregation() string`

GetNodeAggregation returns the NodeAggregation field if non-nil, zero value otherwise.

### GetNodeAggregationOk

`func (o *MetricSettings) GetNodeAggregationOk() (*string, bool)`

GetNodeAggregationOk returns a tuple with the NodeAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAggregation

`func (o *MetricSettings) SetNodeAggregation(v string)`

SetNodeAggregation sets NodeAggregation field to given value.

### HasNodeAggregation

`func (o *MetricSettings) HasNodeAggregation() bool`

HasNodeAggregation returns a boolean if a field has been set.

### GetReturnAggregatedValue

`func (o *MetricSettings) GetReturnAggregatedValue() bool`

GetReturnAggregatedValue returns the ReturnAggregatedValue field if non-nil, zero value otherwise.

### GetReturnAggregatedValueOk

`func (o *MetricSettings) GetReturnAggregatedValueOk() (*bool, bool)`

GetReturnAggregatedValueOk returns a tuple with the ReturnAggregatedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAggregatedValue

`func (o *MetricSettings) SetReturnAggregatedValue(v bool)`

SetReturnAggregatedValue sets ReturnAggregatedValue field to given value.

### HasReturnAggregatedValue

`func (o *MetricSettings) HasReturnAggregatedValue() bool`

HasReturnAggregatedValue returns a boolean if a field has been set.

### GetSplitCount

`func (o *MetricSettings) GetSplitCount() int32`

GetSplitCount returns the SplitCount field if non-nil, zero value otherwise.

### GetSplitCountOk

`func (o *MetricSettings) GetSplitCountOk() (*int32, bool)`

GetSplitCountOk returns a tuple with the SplitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitCount

`func (o *MetricSettings) SetSplitCount(v int32)`

SetSplitCount sets SplitCount field to given value.

### HasSplitCount

`func (o *MetricSettings) HasSplitCount() bool`

HasSplitCount returns a boolean if a field has been set.

### GetSplitMode

`func (o *MetricSettings) GetSplitMode() string`

GetSplitMode returns the SplitMode field if non-nil, zero value otherwise.

### GetSplitModeOk

`func (o *MetricSettings) GetSplitModeOk() (*string, bool)`

GetSplitModeOk returns a tuple with the SplitMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitMode

`func (o *MetricSettings) SetSplitMode(v string)`

SetSplitMode sets SplitMode field to given value.

### HasSplitMode

`func (o *MetricSettings) HasSplitMode() bool`

HasSplitMode returns a boolean if a field has been set.

### GetSplitType

`func (o *MetricSettings) GetSplitType() string`

GetSplitType returns the SplitType field if non-nil, zero value otherwise.

### GetSplitTypeOk

`func (o *MetricSettings) GetSplitTypeOk() (*string, bool)`

GetSplitTypeOk returns a tuple with the SplitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitType

`func (o *MetricSettings) SetSplitType(v string)`

SetSplitType sets SplitType field to given value.

### HasSplitType

`func (o *MetricSettings) HasSplitType() bool`

HasSplitType returns a boolean if a field has been set.

### GetTimeAggregation

`func (o *MetricSettings) GetTimeAggregation() string`

GetTimeAggregation returns the TimeAggregation field if non-nil, zero value otherwise.

### GetTimeAggregationOk

`func (o *MetricSettings) GetTimeAggregationOk() (*string, bool)`

GetTimeAggregationOk returns a tuple with the TimeAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAggregation

`func (o *MetricSettings) SetTimeAggregation(v string)`

SetTimeAggregation sets TimeAggregation field to given value.

### HasTimeAggregation

`func (o *MetricSettings) HasTimeAggregation() bool`

HasTimeAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


