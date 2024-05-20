# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Help** | **string** |  | 
**Name** | **string** |  | 
**Unit** | **string** |  | 
**Values** | [**[]MetricValue**](MetricValue.md) |  | 

## Methods

### NewMetric

`func NewMetric(help string, name string, unit string, values []MetricValue, ) *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelp

`func (o *Metric) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *Metric) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *Metric) SetHelp(v string)`

SetHelp sets Help field to given value.


### GetName

`func (o *Metric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metric) SetName(v string)`

SetName sets Name field to given value.


### GetUnit

`func (o *Metric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Metric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Metric) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetValues

`func (o *Metric) GetValues() []MetricValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Metric) GetValuesOk() (*[]MetricValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Metric) SetValues(v []MetricValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


