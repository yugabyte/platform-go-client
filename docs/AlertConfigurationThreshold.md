# AlertConfigurationThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | **string** | Threshold condition (greater than, or less than) | 
**Threshold** | **float64** | Threshold value | 

## Methods

### NewAlertConfigurationThreshold

`func NewAlertConfigurationThreshold(condition string, threshold float64, ) *AlertConfigurationThreshold`

NewAlertConfigurationThreshold instantiates a new AlertConfigurationThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertConfigurationThresholdWithDefaults

`func NewAlertConfigurationThresholdWithDefaults() *AlertConfigurationThreshold`

NewAlertConfigurationThresholdWithDefaults instantiates a new AlertConfigurationThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *AlertConfigurationThreshold) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AlertConfigurationThreshold) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AlertConfigurationThreshold) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetThreshold

`func (o *AlertConfigurationThreshold) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AlertConfigurationThreshold) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AlertConfigurationThreshold) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


