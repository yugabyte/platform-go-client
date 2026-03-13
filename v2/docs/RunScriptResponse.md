# RunScriptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**ExecutionSummary**](ExecutionSummary.md) |  | 
**Results** | [**map[string]NodeScriptResult**](NodeScriptResult.md) | Map of node name to execution result | 

## Methods

### NewRunScriptResponse

`func NewRunScriptResponse(summary ExecutionSummary, results map[string]NodeScriptResult, ) *RunScriptResponse`

NewRunScriptResponse instantiates a new RunScriptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunScriptResponseWithDefaults

`func NewRunScriptResponseWithDefaults() *RunScriptResponse`

NewRunScriptResponseWithDefaults instantiates a new RunScriptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *RunScriptResponse) GetSummary() ExecutionSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RunScriptResponse) GetSummaryOk() (*ExecutionSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RunScriptResponse) SetSummary(v ExecutionSummary)`

SetSummary sets Summary field to given value.


### GetResults

`func (o *RunScriptResponse) GetResults() map[string]NodeScriptResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RunScriptResponse) GetResultsOk() (*map[string]NodeScriptResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RunScriptResponse) SetResults(v map[string]NodeScriptResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


