# JobInstanceApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartWindowSecs** | Pointer to **int64** | Filter by start time falling within the time window from now. | [optional] 
**State** | Pointer to [**JobInstanceState**](JobInstanceState.md) |  | [optional] 

## Methods

### NewJobInstanceApiFilter

`func NewJobInstanceApiFilter() *JobInstanceApiFilter`

NewJobInstanceApiFilter instantiates a new JobInstanceApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInstanceApiFilterWithDefaults

`func NewJobInstanceApiFilterWithDefaults() *JobInstanceApiFilter`

NewJobInstanceApiFilterWithDefaults instantiates a new JobInstanceApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartWindowSecs

`func (o *JobInstanceApiFilter) GetStartWindowSecs() int64`

GetStartWindowSecs returns the StartWindowSecs field if non-nil, zero value otherwise.

### GetStartWindowSecsOk

`func (o *JobInstanceApiFilter) GetStartWindowSecsOk() (*int64, bool)`

GetStartWindowSecsOk returns a tuple with the StartWindowSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWindowSecs

`func (o *JobInstanceApiFilter) SetStartWindowSecs(v int64)`

SetStartWindowSecs sets StartWindowSecs field to given value.

### HasStartWindowSecs

`func (o *JobInstanceApiFilter) HasStartWindowSecs() bool`

HasStartWindowSecs returns a boolean if a field has been set.

### GetState

`func (o *JobInstanceApiFilter) GetState() JobInstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JobInstanceApiFilter) GetStateOk() (*JobInstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JobInstanceApiFilter) SetState(v JobInstanceState)`

SetState sets State field to given value.

### HasState

`func (o *JobInstanceApiFilter) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


