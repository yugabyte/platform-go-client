# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BacklogStatus** | Pointer to **bool** | Backlog status of schedule arose due to conflicts | [optional] [readonly] 
**CronExpression** | Pointer to **string** | Cron expression for the schedule | [optional] 
**FailureCount** | Pointer to **int32** | Number of failed backup attempts | [optional] [readonly] 
**Frequency** | Pointer to **int64** | Frequency of the schedule, in milli seconds | [optional] 
**FrequencyTimeUnit** | Pointer to **string** | Time unit of frequency | [optional] 
**IncrementBacklogStatus** | Pointer to **bool** | Backlog status of schedule of incremental backups arose due to conflicts | [optional] [readonly] 
**NextIncrementScheduleTaskTime** | Pointer to **time.Time** | Time on which schedule is expected to run for incremental backups | [optional] [readonly] 
**NextScheduleTaskTime** | Pointer to **time.Time** | Time on which schedule is expected to run | [optional] [readonly] 
**OwnerUUID** | Pointer to **string** | Owner UUID for the schedule | [optional] [readonly] 
**RunningState** | Pointer to **bool** | Running state of the schedule | [optional] 
**ScheduleName** | Pointer to **string** | Name of the schedule | [optional] [readonly] 
**ScheduleUUID** | Pointer to **string** | Schedule UUID | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the task. Possible values are _Active_, _Paused_, or _Stopped_. | [optional] [readonly] 
**TaskType** | Pointer to **string** | Type of task to be scheduled. | [optional] 
**UseLocalTimezone** | Pointer to **bool** | Whether to use local timezone with cron expression for the schedule | [optional] 
**UserEmail** | Pointer to **string** | User who created the schedule policy | [optional] [readonly] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacklogStatus

`func (o *Schedule) GetBacklogStatus() bool`

GetBacklogStatus returns the BacklogStatus field if non-nil, zero value otherwise.

### GetBacklogStatusOk

`func (o *Schedule) GetBacklogStatusOk() (*bool, bool)`

GetBacklogStatusOk returns a tuple with the BacklogStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklogStatus

`func (o *Schedule) SetBacklogStatus(v bool)`

SetBacklogStatus sets BacklogStatus field to given value.

### HasBacklogStatus

`func (o *Schedule) HasBacklogStatus() bool`

HasBacklogStatus returns a boolean if a field has been set.

### GetCronExpression

`func (o *Schedule) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *Schedule) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *Schedule) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *Schedule) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetFailureCount

`func (o *Schedule) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *Schedule) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *Schedule) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *Schedule) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetFrequency

`func (o *Schedule) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Schedule) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Schedule) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *Schedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetFrequencyTimeUnit

`func (o *Schedule) GetFrequencyTimeUnit() string`

GetFrequencyTimeUnit returns the FrequencyTimeUnit field if non-nil, zero value otherwise.

### GetFrequencyTimeUnitOk

`func (o *Schedule) GetFrequencyTimeUnitOk() (*string, bool)`

GetFrequencyTimeUnitOk returns a tuple with the FrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyTimeUnit

`func (o *Schedule) SetFrequencyTimeUnit(v string)`

SetFrequencyTimeUnit sets FrequencyTimeUnit field to given value.

### HasFrequencyTimeUnit

`func (o *Schedule) HasFrequencyTimeUnit() bool`

HasFrequencyTimeUnit returns a boolean if a field has been set.

### GetIncrementBacklogStatus

`func (o *Schedule) GetIncrementBacklogStatus() bool`

GetIncrementBacklogStatus returns the IncrementBacklogStatus field if non-nil, zero value otherwise.

### GetIncrementBacklogStatusOk

`func (o *Schedule) GetIncrementBacklogStatusOk() (*bool, bool)`

GetIncrementBacklogStatusOk returns a tuple with the IncrementBacklogStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementBacklogStatus

`func (o *Schedule) SetIncrementBacklogStatus(v bool)`

SetIncrementBacklogStatus sets IncrementBacklogStatus field to given value.

### HasIncrementBacklogStatus

`func (o *Schedule) HasIncrementBacklogStatus() bool`

HasIncrementBacklogStatus returns a boolean if a field has been set.

### GetNextIncrementScheduleTaskTime

`func (o *Schedule) GetNextIncrementScheduleTaskTime() time.Time`

GetNextIncrementScheduleTaskTime returns the NextIncrementScheduleTaskTime field if non-nil, zero value otherwise.

### GetNextIncrementScheduleTaskTimeOk

`func (o *Schedule) GetNextIncrementScheduleTaskTimeOk() (*time.Time, bool)`

GetNextIncrementScheduleTaskTimeOk returns a tuple with the NextIncrementScheduleTaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextIncrementScheduleTaskTime

`func (o *Schedule) SetNextIncrementScheduleTaskTime(v time.Time)`

SetNextIncrementScheduleTaskTime sets NextIncrementScheduleTaskTime field to given value.

### HasNextIncrementScheduleTaskTime

`func (o *Schedule) HasNextIncrementScheduleTaskTime() bool`

HasNextIncrementScheduleTaskTime returns a boolean if a field has been set.

### GetNextScheduleTaskTime

`func (o *Schedule) GetNextScheduleTaskTime() time.Time`

GetNextScheduleTaskTime returns the NextScheduleTaskTime field if non-nil, zero value otherwise.

### GetNextScheduleTaskTimeOk

`func (o *Schedule) GetNextScheduleTaskTimeOk() (*time.Time, bool)`

GetNextScheduleTaskTimeOk returns a tuple with the NextScheduleTaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduleTaskTime

`func (o *Schedule) SetNextScheduleTaskTime(v time.Time)`

SetNextScheduleTaskTime sets NextScheduleTaskTime field to given value.

### HasNextScheduleTaskTime

`func (o *Schedule) HasNextScheduleTaskTime() bool`

HasNextScheduleTaskTime returns a boolean if a field has been set.

### GetOwnerUUID

`func (o *Schedule) GetOwnerUUID() string`

GetOwnerUUID returns the OwnerUUID field if non-nil, zero value otherwise.

### GetOwnerUUIDOk

`func (o *Schedule) GetOwnerUUIDOk() (*string, bool)`

GetOwnerUUIDOk returns a tuple with the OwnerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUUID

`func (o *Schedule) SetOwnerUUID(v string)`

SetOwnerUUID sets OwnerUUID field to given value.

### HasOwnerUUID

`func (o *Schedule) HasOwnerUUID() bool`

HasOwnerUUID returns a boolean if a field has been set.

### GetRunningState

`func (o *Schedule) GetRunningState() bool`

GetRunningState returns the RunningState field if non-nil, zero value otherwise.

### GetRunningStateOk

`func (o *Schedule) GetRunningStateOk() (*bool, bool)`

GetRunningStateOk returns a tuple with the RunningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningState

`func (o *Schedule) SetRunningState(v bool)`

SetRunningState sets RunningState field to given value.

### HasRunningState

`func (o *Schedule) HasRunningState() bool`

HasRunningState returns a boolean if a field has been set.

### GetScheduleName

`func (o *Schedule) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *Schedule) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *Schedule) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *Schedule) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetScheduleUUID

`func (o *Schedule) GetScheduleUUID() string`

GetScheduleUUID returns the ScheduleUUID field if non-nil, zero value otherwise.

### GetScheduleUUIDOk

`func (o *Schedule) GetScheduleUUIDOk() (*string, bool)`

GetScheduleUUIDOk returns a tuple with the ScheduleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUID

`func (o *Schedule) SetScheduleUUID(v string)`

SetScheduleUUID sets ScheduleUUID field to given value.

### HasScheduleUUID

`func (o *Schedule) HasScheduleUUID() bool`

HasScheduleUUID returns a boolean if a field has been set.

### GetStatus

`func (o *Schedule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Schedule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Schedule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Schedule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskType

`func (o *Schedule) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *Schedule) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *Schedule) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *Schedule) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetUseLocalTimezone

`func (o *Schedule) GetUseLocalTimezone() bool`

GetUseLocalTimezone returns the UseLocalTimezone field if non-nil, zero value otherwise.

### GetUseLocalTimezoneOk

`func (o *Schedule) GetUseLocalTimezoneOk() (*bool, bool)`

GetUseLocalTimezoneOk returns a tuple with the UseLocalTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalTimezone

`func (o *Schedule) SetUseLocalTimezone(v bool)`

SetUseLocalTimezone sets UseLocalTimezone field to given value.

### HasUseLocalTimezone

`func (o *Schedule) HasUseLocalTimezone() bool`

HasUseLocalTimezone returns a boolean if a field has been set.

### GetUserEmail

`func (o *Schedule) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *Schedule) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *Schedule) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *Schedule) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


