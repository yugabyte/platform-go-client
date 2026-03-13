# ScheduleResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BacklogStatus** | **bool** |  | 
**BackupInfo** | [**BackupInfo**](BackupInfo.md) |  | 
**CronExpression** | **string** |  | 
**CustomerUUID** | **string** |  | 
**FailureCount** | **int32** |  | 
**Frequency** | **int64** |  | 
**FrequencyTimeUnit** | **string** |  | 
**IncrementBacklogStatus** | **bool** |  | 
**IncrementalBackupFrequency** | **int64** |  | 
**IncrementalBackupFrequencyTimeUnit** | **string** |  | 
**NextExpectedTask** | Pointer to **time.Time** | Next expected task time | [optional] 
**PrevCompletedTask** | Pointer to **time.Time** | Previous completed task time | [optional] 
**RunningState** | **bool** |  | 
**ScheduleName** | **string** |  | 
**ScheduleUUID** | **string** |  | 
**Status** | **string** |  | 
**TableByTableBackup** | **bool** |  | 
**TaskType** | **string** |  | 
**UseLocalTimezone** | **bool** |  | 

## Methods

### NewScheduleResp

`func NewScheduleResp(backlogStatus bool, backupInfo BackupInfo, cronExpression string, customerUUID string, failureCount int32, frequency int64, frequencyTimeUnit string, incrementBacklogStatus bool, incrementalBackupFrequency int64, incrementalBackupFrequencyTimeUnit string, runningState bool, scheduleName string, scheduleUUID string, status string, tableByTableBackup bool, taskType string, useLocalTimezone bool, ) *ScheduleResp`

NewScheduleResp instantiates a new ScheduleResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRespWithDefaults

`func NewScheduleRespWithDefaults() *ScheduleResp`

NewScheduleRespWithDefaults instantiates a new ScheduleResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacklogStatus

`func (o *ScheduleResp) GetBacklogStatus() bool`

GetBacklogStatus returns the BacklogStatus field if non-nil, zero value otherwise.

### GetBacklogStatusOk

`func (o *ScheduleResp) GetBacklogStatusOk() (*bool, bool)`

GetBacklogStatusOk returns a tuple with the BacklogStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklogStatus

`func (o *ScheduleResp) SetBacklogStatus(v bool)`

SetBacklogStatus sets BacklogStatus field to given value.


### GetBackupInfo

`func (o *ScheduleResp) GetBackupInfo() BackupInfo`

GetBackupInfo returns the BackupInfo field if non-nil, zero value otherwise.

### GetBackupInfoOk

`func (o *ScheduleResp) GetBackupInfoOk() (*BackupInfo, bool)`

GetBackupInfoOk returns a tuple with the BackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInfo

`func (o *ScheduleResp) SetBackupInfo(v BackupInfo)`

SetBackupInfo sets BackupInfo field to given value.


### GetCronExpression

`func (o *ScheduleResp) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ScheduleResp) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ScheduleResp) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetCustomerUUID

`func (o *ScheduleResp) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *ScheduleResp) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *ScheduleResp) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetFailureCount

`func (o *ScheduleResp) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *ScheduleResp) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *ScheduleResp) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.


### GetFrequency

`func (o *ScheduleResp) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ScheduleResp) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ScheduleResp) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.


### GetFrequencyTimeUnit

`func (o *ScheduleResp) GetFrequencyTimeUnit() string`

GetFrequencyTimeUnit returns the FrequencyTimeUnit field if non-nil, zero value otherwise.

### GetFrequencyTimeUnitOk

`func (o *ScheduleResp) GetFrequencyTimeUnitOk() (*string, bool)`

GetFrequencyTimeUnitOk returns a tuple with the FrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyTimeUnit

`func (o *ScheduleResp) SetFrequencyTimeUnit(v string)`

SetFrequencyTimeUnit sets FrequencyTimeUnit field to given value.


### GetIncrementBacklogStatus

`func (o *ScheduleResp) GetIncrementBacklogStatus() bool`

GetIncrementBacklogStatus returns the IncrementBacklogStatus field if non-nil, zero value otherwise.

### GetIncrementBacklogStatusOk

`func (o *ScheduleResp) GetIncrementBacklogStatusOk() (*bool, bool)`

GetIncrementBacklogStatusOk returns a tuple with the IncrementBacklogStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementBacklogStatus

`func (o *ScheduleResp) SetIncrementBacklogStatus(v bool)`

SetIncrementBacklogStatus sets IncrementBacklogStatus field to given value.


### GetIncrementalBackupFrequency

`func (o *ScheduleResp) GetIncrementalBackupFrequency() int64`

GetIncrementalBackupFrequency returns the IncrementalBackupFrequency field if non-nil, zero value otherwise.

### GetIncrementalBackupFrequencyOk

`func (o *ScheduleResp) GetIncrementalBackupFrequencyOk() (*int64, bool)`

GetIncrementalBackupFrequencyOk returns a tuple with the IncrementalBackupFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupFrequency

`func (o *ScheduleResp) SetIncrementalBackupFrequency(v int64)`

SetIncrementalBackupFrequency sets IncrementalBackupFrequency field to given value.


### GetIncrementalBackupFrequencyTimeUnit

`func (o *ScheduleResp) GetIncrementalBackupFrequencyTimeUnit() string`

GetIncrementalBackupFrequencyTimeUnit returns the IncrementalBackupFrequencyTimeUnit field if non-nil, zero value otherwise.

### GetIncrementalBackupFrequencyTimeUnitOk

`func (o *ScheduleResp) GetIncrementalBackupFrequencyTimeUnitOk() (*string, bool)`

GetIncrementalBackupFrequencyTimeUnitOk returns a tuple with the IncrementalBackupFrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupFrequencyTimeUnit

`func (o *ScheduleResp) SetIncrementalBackupFrequencyTimeUnit(v string)`

SetIncrementalBackupFrequencyTimeUnit sets IncrementalBackupFrequencyTimeUnit field to given value.


### GetNextExpectedTask

`func (o *ScheduleResp) GetNextExpectedTask() time.Time`

GetNextExpectedTask returns the NextExpectedTask field if non-nil, zero value otherwise.

### GetNextExpectedTaskOk

`func (o *ScheduleResp) GetNextExpectedTaskOk() (*time.Time, bool)`

GetNextExpectedTaskOk returns a tuple with the NextExpectedTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExpectedTask

`func (o *ScheduleResp) SetNextExpectedTask(v time.Time)`

SetNextExpectedTask sets NextExpectedTask field to given value.

### HasNextExpectedTask

`func (o *ScheduleResp) HasNextExpectedTask() bool`

HasNextExpectedTask returns a boolean if a field has been set.

### GetPrevCompletedTask

`func (o *ScheduleResp) GetPrevCompletedTask() time.Time`

GetPrevCompletedTask returns the PrevCompletedTask field if non-nil, zero value otherwise.

### GetPrevCompletedTaskOk

`func (o *ScheduleResp) GetPrevCompletedTaskOk() (*time.Time, bool)`

GetPrevCompletedTaskOk returns a tuple with the PrevCompletedTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCompletedTask

`func (o *ScheduleResp) SetPrevCompletedTask(v time.Time)`

SetPrevCompletedTask sets PrevCompletedTask field to given value.

### HasPrevCompletedTask

`func (o *ScheduleResp) HasPrevCompletedTask() bool`

HasPrevCompletedTask returns a boolean if a field has been set.

### GetRunningState

`func (o *ScheduleResp) GetRunningState() bool`

GetRunningState returns the RunningState field if non-nil, zero value otherwise.

### GetRunningStateOk

`func (o *ScheduleResp) GetRunningStateOk() (*bool, bool)`

GetRunningStateOk returns a tuple with the RunningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningState

`func (o *ScheduleResp) SetRunningState(v bool)`

SetRunningState sets RunningState field to given value.


### GetScheduleName

`func (o *ScheduleResp) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *ScheduleResp) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *ScheduleResp) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.


### GetScheduleUUID

`func (o *ScheduleResp) GetScheduleUUID() string`

GetScheduleUUID returns the ScheduleUUID field if non-nil, zero value otherwise.

### GetScheduleUUIDOk

`func (o *ScheduleResp) GetScheduleUUIDOk() (*string, bool)`

GetScheduleUUIDOk returns a tuple with the ScheduleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUID

`func (o *ScheduleResp) SetScheduleUUID(v string)`

SetScheduleUUID sets ScheduleUUID field to given value.


### GetStatus

`func (o *ScheduleResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleResp) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTableByTableBackup

`func (o *ScheduleResp) GetTableByTableBackup() bool`

GetTableByTableBackup returns the TableByTableBackup field if non-nil, zero value otherwise.

### GetTableByTableBackupOk

`func (o *ScheduleResp) GetTableByTableBackupOk() (*bool, bool)`

GetTableByTableBackupOk returns a tuple with the TableByTableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableByTableBackup

`func (o *ScheduleResp) SetTableByTableBackup(v bool)`

SetTableByTableBackup sets TableByTableBackup field to given value.


### GetTaskType

`func (o *ScheduleResp) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ScheduleResp) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ScheduleResp) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.


### GetUseLocalTimezone

`func (o *ScheduleResp) GetUseLocalTimezone() bool`

GetUseLocalTimezone returns the UseLocalTimezone field if non-nil, zero value otherwise.

### GetUseLocalTimezoneOk

`func (o *ScheduleResp) GetUseLocalTimezoneOk() (*bool, bool)`

GetUseLocalTimezoneOk returns a tuple with the UseLocalTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalTimezone

`func (o *ScheduleResp) SetUseLocalTimezone(v bool)`

SetUseLocalTimezone sets UseLocalTimezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


