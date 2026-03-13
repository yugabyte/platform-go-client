# RunScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScriptOptions** | [**ScriptOptions**](ScriptOptions.md) |  | 
**Nodes** | Pointer to [**NodeSelection**](NodeSelection.md) |  | [optional] 

## Methods

### NewRunScriptRequest

`func NewRunScriptRequest(scriptOptions ScriptOptions, ) *RunScriptRequest`

NewRunScriptRequest instantiates a new RunScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunScriptRequestWithDefaults

`func NewRunScriptRequestWithDefaults() *RunScriptRequest`

NewRunScriptRequestWithDefaults instantiates a new RunScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScriptOptions

`func (o *RunScriptRequest) GetScriptOptions() ScriptOptions`

GetScriptOptions returns the ScriptOptions field if non-nil, zero value otherwise.

### GetScriptOptionsOk

`func (o *RunScriptRequest) GetScriptOptionsOk() (*ScriptOptions, bool)`

GetScriptOptionsOk returns a tuple with the ScriptOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptOptions

`func (o *RunScriptRequest) SetScriptOptions(v ScriptOptions)`

SetScriptOptions sets ScriptOptions field to given value.


### GetNodes

`func (o *RunScriptRequest) GetNodes() NodeSelection`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *RunScriptRequest) GetNodesOk() (*NodeSelection, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *RunScriptRequest) SetNodes(v NodeSelection)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *RunScriptRequest) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


