# TaskExecutorShutdownResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | True if task executor shutdown was initiated by this call.  | [optional] 

## Methods

### NewTaskExecutorShutdownResp

`func NewTaskExecutorShutdownResp() *TaskExecutorShutdownResp`

NewTaskExecutorShutdownResp instantiates a new TaskExecutorShutdownResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskExecutorShutdownRespWithDefaults

`func NewTaskExecutorShutdownRespWithDefaults() *TaskExecutorShutdownResp`

NewTaskExecutorShutdownRespWithDefaults instantiates a new TaskExecutorShutdownResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TaskExecutorShutdownResp) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TaskExecutorShutdownResp) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TaskExecutorShutdownResp) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TaskExecutorShutdownResp) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


