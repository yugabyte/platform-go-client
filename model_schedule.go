/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ywclient

import (
	"encoding/json"
	"time"
)

// Schedule Backup schedule
type Schedule struct {
	// Backlog status of schedule arose due to conflicts
	BacklogStatus *bool `json:"backlogStatus,omitempty"`
	// Cron expression for the schedule
	CronExpression *string `json:"cronExpression,omitempty"`
	// Number of failed backup attempts
	FailureCount *int32 `json:"failureCount,omitempty"`
	// Frequency of the schedule, in milli seconds
	Frequency *int64 `json:"frequency,omitempty"`
	// Time unit of frequency
	FrequencyTimeUnit *string `json:"frequencyTimeUnit,omitempty"`
	// Backlog status of schedule of incremental backups arose due to conflicts
	IncrementBacklogStatus *bool `json:"incrementBacklogStatus,omitempty"`
	// Time on which schedule is expected to run for incremental backups
	NextIncrementScheduleTaskTime *time.Time `json:"nextIncrementScheduleTaskTime,omitempty"`
	// Time on which schedule is expected to run
	NextScheduleTaskTime *time.Time `json:"nextScheduleTaskTime,omitempty"`
	// Owner UUID for the schedule
	OwnerUUID *string `json:"ownerUUID,omitempty"`
	// Running state of the schedule
	RunningState *bool `json:"runningState,omitempty"`
	// Name of the schedule
	ScheduleName *string `json:"scheduleName,omitempty"`
	// Schedule UUID
	ScheduleUUID *string `json:"scheduleUUID,omitempty"`
	// Status of the task. Possible values are _Active_, _Paused_, or _Stopped_.
	Status *string `json:"status,omitempty"`
	// Type of task to be scheduled.
	TaskType *string `json:"taskType,omitempty"`
	// User who created the schedule policy
	UserEmail *string `json:"userEmail,omitempty"`
}

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule() *Schedule {
	this := Schedule{}
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetBacklogStatus returns the BacklogStatus field value if set, zero value otherwise.
func (o *Schedule) GetBacklogStatus() bool {
	if o == nil || o.BacklogStatus == nil {
		var ret bool
		return ret
	}
	return *o.BacklogStatus
}

// GetBacklogStatusOk returns a tuple with the BacklogStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetBacklogStatusOk() (*bool, bool) {
	if o == nil || o.BacklogStatus == nil {
		return nil, false
	}
	return o.BacklogStatus, true
}

// HasBacklogStatus returns a boolean if a field has been set.
func (o *Schedule) HasBacklogStatus() bool {
	if o != nil && o.BacklogStatus != nil {
		return true
	}

	return false
}

// SetBacklogStatus gets a reference to the given bool and assigns it to the BacklogStatus field.
func (o *Schedule) SetBacklogStatus(v bool) {
	o.BacklogStatus = &v
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *Schedule) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *Schedule) HasCronExpression() bool {
	if o != nil && o.CronExpression != nil {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *Schedule) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetFailureCount returns the FailureCount field value if set, zero value otherwise.
func (o *Schedule) GetFailureCount() int32 {
	if o == nil || o.FailureCount == nil {
		var ret int32
		return ret
	}
	return *o.FailureCount
}

// GetFailureCountOk returns a tuple with the FailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetFailureCountOk() (*int32, bool) {
	if o == nil || o.FailureCount == nil {
		return nil, false
	}
	return o.FailureCount, true
}

// HasFailureCount returns a boolean if a field has been set.
func (o *Schedule) HasFailureCount() bool {
	if o != nil && o.FailureCount != nil {
		return true
	}

	return false
}

// SetFailureCount gets a reference to the given int32 and assigns it to the FailureCount field.
func (o *Schedule) SetFailureCount(v int32) {
	o.FailureCount = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *Schedule) GetFrequency() int64 {
	if o == nil || o.Frequency == nil {
		var ret int64
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetFrequencyOk() (*int64, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *Schedule) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given int64 and assigns it to the Frequency field.
func (o *Schedule) SetFrequency(v int64) {
	o.Frequency = &v
}

// GetFrequencyTimeUnit returns the FrequencyTimeUnit field value if set, zero value otherwise.
func (o *Schedule) GetFrequencyTimeUnit() string {
	if o == nil || o.FrequencyTimeUnit == nil {
		var ret string
		return ret
	}
	return *o.FrequencyTimeUnit
}

// GetFrequencyTimeUnitOk returns a tuple with the FrequencyTimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetFrequencyTimeUnitOk() (*string, bool) {
	if o == nil || o.FrequencyTimeUnit == nil {
		return nil, false
	}
	return o.FrequencyTimeUnit, true
}

// HasFrequencyTimeUnit returns a boolean if a field has been set.
func (o *Schedule) HasFrequencyTimeUnit() bool {
	if o != nil && o.FrequencyTimeUnit != nil {
		return true
	}

	return false
}

// SetFrequencyTimeUnit gets a reference to the given string and assigns it to the FrequencyTimeUnit field.
func (o *Schedule) SetFrequencyTimeUnit(v string) {
	o.FrequencyTimeUnit = &v
}

// GetIncrementBacklogStatus returns the IncrementBacklogStatus field value if set, zero value otherwise.
func (o *Schedule) GetIncrementBacklogStatus() bool {
	if o == nil || o.IncrementBacklogStatus == nil {
		var ret bool
		return ret
	}
	return *o.IncrementBacklogStatus
}

// GetIncrementBacklogStatusOk returns a tuple with the IncrementBacklogStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetIncrementBacklogStatusOk() (*bool, bool) {
	if o == nil || o.IncrementBacklogStatus == nil {
		return nil, false
	}
	return o.IncrementBacklogStatus, true
}

// HasIncrementBacklogStatus returns a boolean if a field has been set.
func (o *Schedule) HasIncrementBacklogStatus() bool {
	if o != nil && o.IncrementBacklogStatus != nil {
		return true
	}

	return false
}

// SetIncrementBacklogStatus gets a reference to the given bool and assigns it to the IncrementBacklogStatus field.
func (o *Schedule) SetIncrementBacklogStatus(v bool) {
	o.IncrementBacklogStatus = &v
}

// GetNextIncrementScheduleTaskTime returns the NextIncrementScheduleTaskTime field value if set, zero value otherwise.
func (o *Schedule) GetNextIncrementScheduleTaskTime() time.Time {
	if o == nil || o.NextIncrementScheduleTaskTime == nil {
		var ret time.Time
		return ret
	}
	return *o.NextIncrementScheduleTaskTime
}

// GetNextIncrementScheduleTaskTimeOk returns a tuple with the NextIncrementScheduleTaskTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetNextIncrementScheduleTaskTimeOk() (*time.Time, bool) {
	if o == nil || o.NextIncrementScheduleTaskTime == nil {
		return nil, false
	}
	return o.NextIncrementScheduleTaskTime, true
}

// HasNextIncrementScheduleTaskTime returns a boolean if a field has been set.
func (o *Schedule) HasNextIncrementScheduleTaskTime() bool {
	if o != nil && o.NextIncrementScheduleTaskTime != nil {
		return true
	}

	return false
}

// SetNextIncrementScheduleTaskTime gets a reference to the given time.Time and assigns it to the NextIncrementScheduleTaskTime field.
func (o *Schedule) SetNextIncrementScheduleTaskTime(v time.Time) {
	o.NextIncrementScheduleTaskTime = &v
}

// GetNextScheduleTaskTime returns the NextScheduleTaskTime field value if set, zero value otherwise.
func (o *Schedule) GetNextScheduleTaskTime() time.Time {
	if o == nil || o.NextScheduleTaskTime == nil {
		var ret time.Time
		return ret
	}
	return *o.NextScheduleTaskTime
}

// GetNextScheduleTaskTimeOk returns a tuple with the NextScheduleTaskTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetNextScheduleTaskTimeOk() (*time.Time, bool) {
	if o == nil || o.NextScheduleTaskTime == nil {
		return nil, false
	}
	return o.NextScheduleTaskTime, true
}

// HasNextScheduleTaskTime returns a boolean if a field has been set.
func (o *Schedule) HasNextScheduleTaskTime() bool {
	if o != nil && o.NextScheduleTaskTime != nil {
		return true
	}

	return false
}

// SetNextScheduleTaskTime gets a reference to the given time.Time and assigns it to the NextScheduleTaskTime field.
func (o *Schedule) SetNextScheduleTaskTime(v time.Time) {
	o.NextScheduleTaskTime = &v
}

// GetOwnerUUID returns the OwnerUUID field value if set, zero value otherwise.
func (o *Schedule) GetOwnerUUID() string {
	if o == nil || o.OwnerUUID == nil {
		var ret string
		return ret
	}
	return *o.OwnerUUID
}

// GetOwnerUUIDOk returns a tuple with the OwnerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetOwnerUUIDOk() (*string, bool) {
	if o == nil || o.OwnerUUID == nil {
		return nil, false
	}
	return o.OwnerUUID, true
}

// HasOwnerUUID returns a boolean if a field has been set.
func (o *Schedule) HasOwnerUUID() bool {
	if o != nil && o.OwnerUUID != nil {
		return true
	}

	return false
}

// SetOwnerUUID gets a reference to the given string and assigns it to the OwnerUUID field.
func (o *Schedule) SetOwnerUUID(v string) {
	o.OwnerUUID = &v
}

// GetRunningState returns the RunningState field value if set, zero value otherwise.
func (o *Schedule) GetRunningState() bool {
	if o == nil || o.RunningState == nil {
		var ret bool
		return ret
	}
	return *o.RunningState
}

// GetRunningStateOk returns a tuple with the RunningState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetRunningStateOk() (*bool, bool) {
	if o == nil || o.RunningState == nil {
		return nil, false
	}
	return o.RunningState, true
}

// HasRunningState returns a boolean if a field has been set.
func (o *Schedule) HasRunningState() bool {
	if o != nil && o.RunningState != nil {
		return true
	}

	return false
}

// SetRunningState gets a reference to the given bool and assigns it to the RunningState field.
func (o *Schedule) SetRunningState(v bool) {
	o.RunningState = &v
}

// GetScheduleName returns the ScheduleName field value if set, zero value otherwise.
func (o *Schedule) GetScheduleName() string {
	if o == nil || o.ScheduleName == nil {
		var ret string
		return ret
	}
	return *o.ScheduleName
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetScheduleNameOk() (*string, bool) {
	if o == nil || o.ScheduleName == nil {
		return nil, false
	}
	return o.ScheduleName, true
}

// HasScheduleName returns a boolean if a field has been set.
func (o *Schedule) HasScheduleName() bool {
	if o != nil && o.ScheduleName != nil {
		return true
	}

	return false
}

// SetScheduleName gets a reference to the given string and assigns it to the ScheduleName field.
func (o *Schedule) SetScheduleName(v string) {
	o.ScheduleName = &v
}

// GetScheduleUUID returns the ScheduleUUID field value if set, zero value otherwise.
func (o *Schedule) GetScheduleUUID() string {
	if o == nil || o.ScheduleUUID == nil {
		var ret string
		return ret
	}
	return *o.ScheduleUUID
}

// GetScheduleUUIDOk returns a tuple with the ScheduleUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetScheduleUUIDOk() (*string, bool) {
	if o == nil || o.ScheduleUUID == nil {
		return nil, false
	}
	return o.ScheduleUUID, true
}

// HasScheduleUUID returns a boolean if a field has been set.
func (o *Schedule) HasScheduleUUID() bool {
	if o != nil && o.ScheduleUUID != nil {
		return true
	}

	return false
}

// SetScheduleUUID gets a reference to the given string and assigns it to the ScheduleUUID field.
func (o *Schedule) SetScheduleUUID(v string) {
	o.ScheduleUUID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Schedule) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Schedule) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Schedule) SetStatus(v string) {
	o.Status = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *Schedule) GetTaskType() string {
	if o == nil || o.TaskType == nil {
		var ret string
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTaskTypeOk() (*string, bool) {
	if o == nil || o.TaskType == nil {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *Schedule) HasTaskType() bool {
	if o != nil && o.TaskType != nil {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given string and assigns it to the TaskType field.
func (o *Schedule) SetTaskType(v string) {
	o.TaskType = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *Schedule) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetUserEmailOk() (*string, bool) {
	if o == nil || o.UserEmail == nil {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *Schedule) HasUserEmail() bool {
	if o != nil && o.UserEmail != nil {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *Schedule) SetUserEmail(v string) {
	o.UserEmail = &v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BacklogStatus != nil {
		toSerialize["backlogStatus"] = o.BacklogStatus
	}
	if o.CronExpression != nil {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if o.FailureCount != nil {
		toSerialize["failureCount"] = o.FailureCount
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.FrequencyTimeUnit != nil {
		toSerialize["frequencyTimeUnit"] = o.FrequencyTimeUnit
	}
	if o.IncrementBacklogStatus != nil {
		toSerialize["incrementBacklogStatus"] = o.IncrementBacklogStatus
	}
	if o.NextIncrementScheduleTaskTime != nil {
		toSerialize["nextIncrementScheduleTaskTime"] = o.NextIncrementScheduleTaskTime
	}
	if o.NextScheduleTaskTime != nil {
		toSerialize["nextScheduleTaskTime"] = o.NextScheduleTaskTime
	}
	if o.OwnerUUID != nil {
		toSerialize["ownerUUID"] = o.OwnerUUID
	}
	if o.RunningState != nil {
		toSerialize["runningState"] = o.RunningState
	}
	if o.ScheduleName != nil {
		toSerialize["scheduleName"] = o.ScheduleName
	}
	if o.ScheduleUUID != nil {
		toSerialize["scheduleUUID"] = o.ScheduleUUID
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TaskType != nil {
		toSerialize["taskType"] = o.TaskType
	}
	if o.UserEmail != nil {
		toSerialize["userEmail"] = o.UserEmail
	}
	return json.Marshal(toSerialize)
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


