# NodeScriptResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeName** | **string** | Name of the node | 
**NodeAddress** | Pointer to **string** | IP address or hostname of the node | [optional] 
**ExitCode** | **int32** | Exit code of the script. 0 indicates success. | 
**Stdout** | Pointer to **string** | Standard output from the script execution | [optional] 
**Stderr** | Pointer to **string** | Standard error from the script execution | [optional] 
**ExecutionTimeMs** | Pointer to **int64** | Time taken to execute the script in milliseconds | [optional] 
**Success** | **bool** | Whether the script execution was successful (exit_code &#x3D;&#x3D; 0) | 
**ErrorMessage** | Pointer to **string** | Error message if the execution failed (e.g., timeout, connection error) | [optional] 

## Methods

### NewNodeScriptResult

`func NewNodeScriptResult(nodeName string, exitCode int32, success bool, ) *NodeScriptResult`

NewNodeScriptResult instantiates a new NodeScriptResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeScriptResultWithDefaults

`func NewNodeScriptResultWithDefaults() *NodeScriptResult`

NewNodeScriptResultWithDefaults instantiates a new NodeScriptResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeName

`func (o *NodeScriptResult) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *NodeScriptResult) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *NodeScriptResult) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetNodeAddress

`func (o *NodeScriptResult) GetNodeAddress() string`

GetNodeAddress returns the NodeAddress field if non-nil, zero value otherwise.

### GetNodeAddressOk

`func (o *NodeScriptResult) GetNodeAddressOk() (*string, bool)`

GetNodeAddressOk returns a tuple with the NodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddress

`func (o *NodeScriptResult) SetNodeAddress(v string)`

SetNodeAddress sets NodeAddress field to given value.

### HasNodeAddress

`func (o *NodeScriptResult) HasNodeAddress() bool`

HasNodeAddress returns a boolean if a field has been set.

### GetExitCode

`func (o *NodeScriptResult) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *NodeScriptResult) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *NodeScriptResult) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetStdout

`func (o *NodeScriptResult) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *NodeScriptResult) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *NodeScriptResult) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *NodeScriptResult) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetStderr

`func (o *NodeScriptResult) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *NodeScriptResult) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *NodeScriptResult) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *NodeScriptResult) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetExecutionTimeMs

`func (o *NodeScriptResult) GetExecutionTimeMs() int64`

GetExecutionTimeMs returns the ExecutionTimeMs field if non-nil, zero value otherwise.

### GetExecutionTimeMsOk

`func (o *NodeScriptResult) GetExecutionTimeMsOk() (*int64, bool)`

GetExecutionTimeMsOk returns a tuple with the ExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeMs

`func (o *NodeScriptResult) SetExecutionTimeMs(v int64)`

SetExecutionTimeMs sets ExecutionTimeMs field to given value.

### HasExecutionTimeMs

`func (o *NodeScriptResult) HasExecutionTimeMs() bool`

HasExecutionTimeMs returns a boolean if a field has been set.

### GetSuccess

`func (o *NodeScriptResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *NodeScriptResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *NodeScriptResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrorMessage

`func (o *NodeScriptResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NodeScriptResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NodeScriptResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NodeScriptResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


