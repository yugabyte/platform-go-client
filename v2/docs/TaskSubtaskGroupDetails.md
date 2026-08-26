# TaskSubtaskGroupDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Short title for this subtask group. | [optional] [readonly] 
**Description** | Pointer to **string** | Longer description of what this subtask group represents. | [optional] [readonly] 
**State** | Pointer to **string** | Aggregated state for this subtask group. | [optional] [readonly] 
**ExtraDetails** | Pointer to **[]map[string]interface{}** | Additional JSON blobs for this subtask group (progress, errors, etc.). | [optional] [readonly] 

## Methods

### NewTaskSubtaskGroupDetails

`func NewTaskSubtaskGroupDetails() *TaskSubtaskGroupDetails`

NewTaskSubtaskGroupDetails instantiates a new TaskSubtaskGroupDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSubtaskGroupDetailsWithDefaults

`func NewTaskSubtaskGroupDetailsWithDefaults() *TaskSubtaskGroupDetails`

NewTaskSubtaskGroupDetailsWithDefaults instantiates a new TaskSubtaskGroupDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TaskSubtaskGroupDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskSubtaskGroupDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskSubtaskGroupDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskSubtaskGroupDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TaskSubtaskGroupDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskSubtaskGroupDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskSubtaskGroupDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskSubtaskGroupDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *TaskSubtaskGroupDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskSubtaskGroupDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskSubtaskGroupDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TaskSubtaskGroupDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetExtraDetails

`func (o *TaskSubtaskGroupDetails) GetExtraDetails() []map[string]interface{}`

GetExtraDetails returns the ExtraDetails field if non-nil, zero value otherwise.

### GetExtraDetailsOk

`func (o *TaskSubtaskGroupDetails) GetExtraDetailsOk() (*[]map[string]interface{}, bool)`

GetExtraDetailsOk returns a tuple with the ExtraDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDetails

`func (o *TaskSubtaskGroupDetails) SetExtraDetails(v []map[string]interface{})`

SetExtraDetails sets ExtraDetails field to given value.

### HasExtraDetails

`func (o *TaskSubtaskGroupDetails) HasExtraDetails() bool`

HasExtraDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


