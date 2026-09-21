# TaskExecutorShutdownStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsShutdownInitiated** | Pointer to **bool** | True if the task executor has been signaled to shut down.  | [optional] [readonly] 
**IsShutdownComplete** | Pointer to **bool** | True if application shutdown hooks have finished after the task executor drained.  | [optional] [readonly] 
**NumRunningTasks** | Pointer to **int32** | Number of parent tasks still registered with the executor.  | [optional] [readonly] 

## Methods

### NewTaskExecutorShutdownStatus

`func NewTaskExecutorShutdownStatus() *TaskExecutorShutdownStatus`

NewTaskExecutorShutdownStatus instantiates a new TaskExecutorShutdownStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskExecutorShutdownStatusWithDefaults

`func NewTaskExecutorShutdownStatusWithDefaults() *TaskExecutorShutdownStatus`

NewTaskExecutorShutdownStatusWithDefaults instantiates a new TaskExecutorShutdownStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsShutdownInitiated

`func (o *TaskExecutorShutdownStatus) GetIsShutdownInitiated() bool`

GetIsShutdownInitiated returns the IsShutdownInitiated field if non-nil, zero value otherwise.

### GetIsShutdownInitiatedOk

`func (o *TaskExecutorShutdownStatus) GetIsShutdownInitiatedOk() (*bool, bool)`

GetIsShutdownInitiatedOk returns a tuple with the IsShutdownInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShutdownInitiated

`func (o *TaskExecutorShutdownStatus) SetIsShutdownInitiated(v bool)`

SetIsShutdownInitiated sets IsShutdownInitiated field to given value.

### HasIsShutdownInitiated

`func (o *TaskExecutorShutdownStatus) HasIsShutdownInitiated() bool`

HasIsShutdownInitiated returns a boolean if a field has been set.

### GetIsShutdownComplete

`func (o *TaskExecutorShutdownStatus) GetIsShutdownComplete() bool`

GetIsShutdownComplete returns the IsShutdownComplete field if non-nil, zero value otherwise.

### GetIsShutdownCompleteOk

`func (o *TaskExecutorShutdownStatus) GetIsShutdownCompleteOk() (*bool, bool)`

GetIsShutdownCompleteOk returns a tuple with the IsShutdownComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShutdownComplete

`func (o *TaskExecutorShutdownStatus) SetIsShutdownComplete(v bool)`

SetIsShutdownComplete sets IsShutdownComplete field to given value.

### HasIsShutdownComplete

`func (o *TaskExecutorShutdownStatus) HasIsShutdownComplete() bool`

HasIsShutdownComplete returns a boolean if a field has been set.

### GetNumRunningTasks

`func (o *TaskExecutorShutdownStatus) GetNumRunningTasks() int32`

GetNumRunningTasks returns the NumRunningTasks field if non-nil, zero value otherwise.

### GetNumRunningTasksOk

`func (o *TaskExecutorShutdownStatus) GetNumRunningTasksOk() (*int32, bool)`

GetNumRunningTasksOk returns a tuple with the NumRunningTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRunningTasks

`func (o *TaskExecutorShutdownStatus) SetNumRunningTasks(v int32)`

SetNumRunningTasks sets NumRunningTasks field to given value.

### HasNumRunningTasks

`func (o *TaskExecutorShutdownStatus) HasNumRunningTasks() bool`

HasNumRunningTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


