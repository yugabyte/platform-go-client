/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"time"
)

// JobInstanceInfo Job instance info representing the execution instance of a JobSchedule.
type JobInstanceInfo struct {
	// Creation time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// End time
	EndTime *time.Time `json:"end_time,omitempty"`
	// Job schedule UUID
	JobScheduleUuid *string `json:"job_schedule_uuid,omitempty"`
	// Start time
	StartTime *time.Time `json:"start_time,omitempty"`
	State *JobInstanceState `json:"state,omitempty"`
	// Job instance UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewJobInstanceInfo instantiates a new JobInstanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobInstanceInfo() *JobInstanceInfo {
	this := JobInstanceInfo{}
	return &this
}

// NewJobInstanceInfoWithDefaults instantiates a new JobInstanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobInstanceInfoWithDefaults() *JobInstanceInfo {
	this := JobInstanceInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *JobInstanceInfo) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstanceInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *JobInstanceInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *JobInstanceInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *JobInstanceInfo) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstanceInfo) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *JobInstanceInfo) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *JobInstanceInfo) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetJobScheduleUuid returns the JobScheduleUuid field value if set, zero value otherwise.
func (o *JobInstanceInfo) GetJobScheduleUuid() string {
	if o == nil || o.JobScheduleUuid == nil {
		var ret string
		return ret
	}
	return *o.JobScheduleUuid
}

// GetJobScheduleUuidOk returns a tuple with the JobScheduleUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstanceInfo) GetJobScheduleUuidOk() (*string, bool) {
	if o == nil || o.JobScheduleUuid == nil {
		return nil, false
	}
	return o.JobScheduleUuid, true
}

// HasJobScheduleUuid returns a boolean if a field has been set.
func (o *JobInstanceInfo) HasJobScheduleUuid() bool {
	if o != nil && o.JobScheduleUuid != nil {
		return true
	}

	return false
}

// SetJobScheduleUuid gets a reference to the given string and assigns it to the JobScheduleUuid field.
func (o *JobInstanceInfo) SetJobScheduleUuid(v string) {
	o.JobScheduleUuid = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *JobInstanceInfo) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstanceInfo) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *JobInstanceInfo) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *JobInstanceInfo) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *JobInstanceInfo) GetState() JobInstanceState {
	if o == nil || o.State == nil {
		var ret JobInstanceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstanceInfo) GetStateOk() (*JobInstanceState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *JobInstanceInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given JobInstanceState and assigns it to the State field.
func (o *JobInstanceInfo) SetState(v JobInstanceState) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *JobInstanceInfo) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstanceInfo) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *JobInstanceInfo) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *JobInstanceInfo) SetUuid(v string) {
	o.Uuid = &v
}

func (o JobInstanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.JobScheduleUuid != nil {
		toSerialize["job_schedule_uuid"] = o.JobScheduleUuid
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableJobInstanceInfo struct {
	value *JobInstanceInfo
	isSet bool
}

func (v NullableJobInstanceInfo) Get() *JobInstanceInfo {
	return v.value
}

func (v *NullableJobInstanceInfo) Set(val *JobInstanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJobInstanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJobInstanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobInstanceInfo(val *JobInstanceInfo) *NullableJobInstanceInfo {
	return &NullableJobInstanceInfo{value: val, isSet: true}
}

func (v NullableJobInstanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobInstanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

