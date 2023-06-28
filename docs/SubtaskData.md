# SubtaskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | Creation time (unix timestamp) of the task | [optional] 
**ErrorString** | Pointer to **string** | Failed SubTask Error message | [optional] 
**SubTaskGroupType** | Pointer to **string** | Failed SubTask Group Type | [optional] 
**SubTaskState** | Pointer to **string** | Failed SubTask State | [optional] 
**SubTaskType** | Pointer to **string** | Failed SubTask Type | [optional] 
**SubTaskUUID** | Pointer to **string** | Failed SubTask UUID | [optional] 

## Methods

### NewSubtaskData

`func NewSubtaskData() *SubtaskData`

NewSubtaskData instantiates a new SubtaskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubtaskDataWithDefaults

`func NewSubtaskDataWithDefaults() *SubtaskData`

NewSubtaskDataWithDefaults instantiates a new SubtaskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *SubtaskData) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *SubtaskData) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *SubtaskData) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *SubtaskData) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetErrorString

`func (o *SubtaskData) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *SubtaskData) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *SubtaskData) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *SubtaskData) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetSubTaskGroupType

`func (o *SubtaskData) GetSubTaskGroupType() string`

GetSubTaskGroupType returns the SubTaskGroupType field if non-nil, zero value otherwise.

### GetSubTaskGroupTypeOk

`func (o *SubtaskData) GetSubTaskGroupTypeOk() (*string, bool)`

GetSubTaskGroupTypeOk returns a tuple with the SubTaskGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTaskGroupType

`func (o *SubtaskData) SetSubTaskGroupType(v string)`

SetSubTaskGroupType sets SubTaskGroupType field to given value.

### HasSubTaskGroupType

`func (o *SubtaskData) HasSubTaskGroupType() bool`

HasSubTaskGroupType returns a boolean if a field has been set.

### GetSubTaskState

`func (o *SubtaskData) GetSubTaskState() string`

GetSubTaskState returns the SubTaskState field if non-nil, zero value otherwise.

### GetSubTaskStateOk

`func (o *SubtaskData) GetSubTaskStateOk() (*string, bool)`

GetSubTaskStateOk returns a tuple with the SubTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTaskState

`func (o *SubtaskData) SetSubTaskState(v string)`

SetSubTaskState sets SubTaskState field to given value.

### HasSubTaskState

`func (o *SubtaskData) HasSubTaskState() bool`

HasSubTaskState returns a boolean if a field has been set.

### GetSubTaskType

`func (o *SubtaskData) GetSubTaskType() string`

GetSubTaskType returns the SubTaskType field if non-nil, zero value otherwise.

### GetSubTaskTypeOk

`func (o *SubtaskData) GetSubTaskTypeOk() (*string, bool)`

GetSubTaskTypeOk returns a tuple with the SubTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTaskType

`func (o *SubtaskData) SetSubTaskType(v string)`

SetSubTaskType sets SubTaskType field to given value.

### HasSubTaskType

`func (o *SubtaskData) HasSubTaskType() bool`

HasSubTaskType returns a boolean if a field has been set.

### GetSubTaskUUID

`func (o *SubtaskData) GetSubTaskUUID() string`

GetSubTaskUUID returns the SubTaskUUID field if non-nil, zero value otherwise.

### GetSubTaskUUIDOk

`func (o *SubtaskData) GetSubTaskUUIDOk() (*string, bool)`

GetSubTaskUUIDOk returns a tuple with the SubTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTaskUUID

`func (o *SubtaskData) SetSubTaskUUID(v string)`

SetSubTaskUUID sets SubTaskUUID field to given value.

### HasSubTaskUUID

`func (o *SubtaskData) HasSubTaskUUID() bool`

HasSubTaskUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


