# TaskExecutorShutdownSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortTimeSeconds** | **int64** | Grace period in seconds before in-flight tasks are aborted.  | 

## Methods

### NewTaskExecutorShutdownSpec

`func NewTaskExecutorShutdownSpec(abortTimeSeconds int64, ) *TaskExecutorShutdownSpec`

NewTaskExecutorShutdownSpec instantiates a new TaskExecutorShutdownSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskExecutorShutdownSpecWithDefaults

`func NewTaskExecutorShutdownSpecWithDefaults() *TaskExecutorShutdownSpec`

NewTaskExecutorShutdownSpecWithDefaults instantiates a new TaskExecutorShutdownSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortTimeSeconds

`func (o *TaskExecutorShutdownSpec) GetAbortTimeSeconds() int64`

GetAbortTimeSeconds returns the AbortTimeSeconds field if non-nil, zero value otherwise.

### GetAbortTimeSecondsOk

`func (o *TaskExecutorShutdownSpec) GetAbortTimeSecondsOk() (*int64, bool)`

GetAbortTimeSecondsOk returns a tuple with the AbortTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortTimeSeconds

`func (o *TaskExecutorShutdownSpec) SetAbortTimeSeconds(v int64)`

SetAbortTimeSeconds sets AbortTimeSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


