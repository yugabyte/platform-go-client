# ExecutionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNodes** | **int32** | Total number of nodes targeted for script execution | 
**SuccessfulNodes** | **int32** | Number of nodes where the script executed successfully (exit_code &#x3D;&#x3D; 0) | 
**FailedNodes** | **int32** | Number of nodes where the script execution failed | 
**TotalExecutionTimeMs** | Pointer to **int64** | Total time taken for the entire operation in milliseconds | [optional] 
**AllSucceeded** | **bool** | Whether the script executed successfully on all targeted nodes | 

## Methods

### NewExecutionSummary

`func NewExecutionSummary(totalNodes int32, successfulNodes int32, failedNodes int32, allSucceeded bool, ) *ExecutionSummary`

NewExecutionSummary instantiates a new ExecutionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionSummaryWithDefaults

`func NewExecutionSummaryWithDefaults() *ExecutionSummary`

NewExecutionSummaryWithDefaults instantiates a new ExecutionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNodes

`func (o *ExecutionSummary) GetTotalNodes() int32`

GetTotalNodes returns the TotalNodes field if non-nil, zero value otherwise.

### GetTotalNodesOk

`func (o *ExecutionSummary) GetTotalNodesOk() (*int32, bool)`

GetTotalNodesOk returns a tuple with the TotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNodes

`func (o *ExecutionSummary) SetTotalNodes(v int32)`

SetTotalNodes sets TotalNodes field to given value.


### GetSuccessfulNodes

`func (o *ExecutionSummary) GetSuccessfulNodes() int32`

GetSuccessfulNodes returns the SuccessfulNodes field if non-nil, zero value otherwise.

### GetSuccessfulNodesOk

`func (o *ExecutionSummary) GetSuccessfulNodesOk() (*int32, bool)`

GetSuccessfulNodesOk returns a tuple with the SuccessfulNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulNodes

`func (o *ExecutionSummary) SetSuccessfulNodes(v int32)`

SetSuccessfulNodes sets SuccessfulNodes field to given value.


### GetFailedNodes

`func (o *ExecutionSummary) GetFailedNodes() int32`

GetFailedNodes returns the FailedNodes field if non-nil, zero value otherwise.

### GetFailedNodesOk

`func (o *ExecutionSummary) GetFailedNodesOk() (*int32, bool)`

GetFailedNodesOk returns a tuple with the FailedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNodes

`func (o *ExecutionSummary) SetFailedNodes(v int32)`

SetFailedNodes sets FailedNodes field to given value.


### GetTotalExecutionTimeMs

`func (o *ExecutionSummary) GetTotalExecutionTimeMs() int64`

GetTotalExecutionTimeMs returns the TotalExecutionTimeMs field if non-nil, zero value otherwise.

### GetTotalExecutionTimeMsOk

`func (o *ExecutionSummary) GetTotalExecutionTimeMsOk() (*int64, bool)`

GetTotalExecutionTimeMsOk returns a tuple with the TotalExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExecutionTimeMs

`func (o *ExecutionSummary) SetTotalExecutionTimeMs(v int64)`

SetTotalExecutionTimeMs sets TotalExecutionTimeMs field to given value.

### HasTotalExecutionTimeMs

`func (o *ExecutionSummary) HasTotalExecutionTimeMs() bool`

HasTotalExecutionTimeMs returns a boolean if a field has been set.

### GetAllSucceeded

`func (o *ExecutionSummary) GetAllSucceeded() bool`

GetAllSucceeded returns the AllSucceeded field if non-nil, zero value otherwise.

### GetAllSucceededOk

`func (o *ExecutionSummary) GetAllSucceededOk() (*bool, bool)`

GetAllSucceededOk returns a tuple with the AllSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSucceeded

`func (o *ExecutionSummary) SetAllSucceeded(v bool)`

SetAllSucceeded sets AllSucceeded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


