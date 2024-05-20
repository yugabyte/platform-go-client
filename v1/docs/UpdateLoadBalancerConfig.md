# UpdateLoadBalancerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstTry** | **bool** |  | 
**Name** | **string** |  | 
**RetryLimit** | **int32** |  | 
**SleepMultiplier** | **int32** |  | 
**TaskInfo** | **string** |  | 

## Methods

### NewUpdateLoadBalancerConfig

`func NewUpdateLoadBalancerConfig(firstTry bool, name string, retryLimit int32, sleepMultiplier int32, taskInfo string, ) *UpdateLoadBalancerConfig`

NewUpdateLoadBalancerConfig instantiates a new UpdateLoadBalancerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerConfigWithDefaults

`func NewUpdateLoadBalancerConfigWithDefaults() *UpdateLoadBalancerConfig`

NewUpdateLoadBalancerConfigWithDefaults instantiates a new UpdateLoadBalancerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstTry

`func (o *UpdateLoadBalancerConfig) GetFirstTry() bool`

GetFirstTry returns the FirstTry field if non-nil, zero value otherwise.

### GetFirstTryOk

`func (o *UpdateLoadBalancerConfig) GetFirstTryOk() (*bool, bool)`

GetFirstTryOk returns a tuple with the FirstTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTry

`func (o *UpdateLoadBalancerConfig) SetFirstTry(v bool)`

SetFirstTry sets FirstTry field to given value.


### GetName

`func (o *UpdateLoadBalancerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoadBalancerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoadBalancerConfig) SetName(v string)`

SetName sets Name field to given value.


### GetRetryLimit

`func (o *UpdateLoadBalancerConfig) GetRetryLimit() int32`

GetRetryLimit returns the RetryLimit field if non-nil, zero value otherwise.

### GetRetryLimitOk

`func (o *UpdateLoadBalancerConfig) GetRetryLimitOk() (*int32, bool)`

GetRetryLimitOk returns a tuple with the RetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryLimit

`func (o *UpdateLoadBalancerConfig) SetRetryLimit(v int32)`

SetRetryLimit sets RetryLimit field to given value.


### GetSleepMultiplier

`func (o *UpdateLoadBalancerConfig) GetSleepMultiplier() int32`

GetSleepMultiplier returns the SleepMultiplier field if non-nil, zero value otherwise.

### GetSleepMultiplierOk

`func (o *UpdateLoadBalancerConfig) GetSleepMultiplierOk() (*int32, bool)`

GetSleepMultiplierOk returns a tuple with the SleepMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepMultiplier

`func (o *UpdateLoadBalancerConfig) SetSleepMultiplier(v int32)`

SetSleepMultiplier sets SleepMultiplier field to given value.


### GetTaskInfo

`func (o *UpdateLoadBalancerConfig) GetTaskInfo() string`

GetTaskInfo returns the TaskInfo field if non-nil, zero value otherwise.

### GetTaskInfoOk

`func (o *UpdateLoadBalancerConfig) GetTaskInfoOk() (*string, bool)`

GetTaskInfoOk returns a tuple with the TaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfo

`func (o *UpdateLoadBalancerConfig) SetTaskInfo(v string)`

SetTaskInfo sets TaskInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


