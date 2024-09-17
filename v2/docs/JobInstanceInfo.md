# JobInstanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Creation time | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | End time | [optional] [readonly] 
**JobScheduleUuid** | Pointer to **string** | Job schedule UUID | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | Start time | [optional] [readonly] 
**State** | Pointer to [**JobInstanceState**](JobInstanceState.md) |  | [optional] 
**Uuid** | Pointer to **string** | Job instance UUID | [optional] [readonly] 

## Methods

### NewJobInstanceInfo

`func NewJobInstanceInfo() *JobInstanceInfo`

NewJobInstanceInfo instantiates a new JobInstanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInstanceInfoWithDefaults

`func NewJobInstanceInfoWithDefaults() *JobInstanceInfo`

NewJobInstanceInfoWithDefaults instantiates a new JobInstanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *JobInstanceInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobInstanceInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobInstanceInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *JobInstanceInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndTime

`func (o *JobInstanceInfo) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *JobInstanceInfo) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *JobInstanceInfo) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *JobInstanceInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetJobScheduleUuid

`func (o *JobInstanceInfo) GetJobScheduleUuid() string`

GetJobScheduleUuid returns the JobScheduleUuid field if non-nil, zero value otherwise.

### GetJobScheduleUuidOk

`func (o *JobInstanceInfo) GetJobScheduleUuidOk() (*string, bool)`

GetJobScheduleUuidOk returns a tuple with the JobScheduleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobScheduleUuid

`func (o *JobInstanceInfo) SetJobScheduleUuid(v string)`

SetJobScheduleUuid sets JobScheduleUuid field to given value.

### HasJobScheduleUuid

`func (o *JobInstanceInfo) HasJobScheduleUuid() bool`

HasJobScheduleUuid returns a boolean if a field has been set.

### GetStartTime

`func (o *JobInstanceInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *JobInstanceInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *JobInstanceInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *JobInstanceInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *JobInstanceInfo) GetState() JobInstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JobInstanceInfo) GetStateOk() (*JobInstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JobInstanceInfo) SetState(v JobInstanceState)`

SetState sets State field to given value.

### HasState

`func (o *JobInstanceInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *JobInstanceInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *JobInstanceInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *JobInstanceInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *JobInstanceInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


