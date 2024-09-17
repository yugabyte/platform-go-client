# PerfAdvisorManualRunStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveRun** | Pointer to [**UniversePerformanceAdvisorStatus**](UniversePerformanceAdvisorStatus.md) |  | [optional] 
**Message** | Pointer to **string** | API response message. | [optional] [readonly] 
**Success** | Pointer to **bool** | API operation status. A value of true indicates the operation was successful. | [optional] [readonly] 

## Methods

### NewPerfAdvisorManualRunStatus

`func NewPerfAdvisorManualRunStatus() *PerfAdvisorManualRunStatus`

NewPerfAdvisorManualRunStatus instantiates a new PerfAdvisorManualRunStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfAdvisorManualRunStatusWithDefaults

`func NewPerfAdvisorManualRunStatusWithDefaults() *PerfAdvisorManualRunStatus`

NewPerfAdvisorManualRunStatusWithDefaults instantiates a new PerfAdvisorManualRunStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveRun

`func (o *PerfAdvisorManualRunStatus) GetActiveRun() UniversePerformanceAdvisorStatus`

GetActiveRun returns the ActiveRun field if non-nil, zero value otherwise.

### GetActiveRunOk

`func (o *PerfAdvisorManualRunStatus) GetActiveRunOk() (*UniversePerformanceAdvisorStatus, bool)`

GetActiveRunOk returns a tuple with the ActiveRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRun

`func (o *PerfAdvisorManualRunStatus) SetActiveRun(v UniversePerformanceAdvisorStatus)`

SetActiveRun sets ActiveRun field to given value.

### HasActiveRun

`func (o *PerfAdvisorManualRunStatus) HasActiveRun() bool`

HasActiveRun returns a boolean if a field has been set.

### GetMessage

`func (o *PerfAdvisorManualRunStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PerfAdvisorManualRunStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PerfAdvisorManualRunStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PerfAdvisorManualRunStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *PerfAdvisorManualRunStatus) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PerfAdvisorManualRunStatus) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PerfAdvisorManualRunStatus) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PerfAdvisorManualRunStatus) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


