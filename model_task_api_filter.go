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

// TaskApiFilter struct for TaskApiFilter
type TaskApiFilter struct {
	// The end date to filter paged query.
	DateRangeEnd *time.Time `json:"dateRangeEnd,omitempty"`
	// The start date to filter paged query.
	DateRangeStart *time.Time `json:"dateRangeStart,omitempty"`
	Status []string `json:"status"`
	TargetList []string `json:"targetList"`
	TargetUUIDList []string `json:"targetUUIDList"`
	TypeList []string `json:"typeList"`
	TypeNameList []string `json:"typeNameList"`
}

// NewTaskApiFilter instantiates a new TaskApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskApiFilter(status []string, targetList []string, targetUUIDList []string, typeList []string, typeNameList []string) *TaskApiFilter {
	this := TaskApiFilter{}
	this.Status = status
	this.TargetList = targetList
	this.TargetUUIDList = targetUUIDList
	this.TypeList = typeList
	this.TypeNameList = typeNameList
	return &this
}

// NewTaskApiFilterWithDefaults instantiates a new TaskApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskApiFilterWithDefaults() *TaskApiFilter {
	this := TaskApiFilter{}
	return &this
}

// GetDateRangeEnd returns the DateRangeEnd field value if set, zero value otherwise.
func (o *TaskApiFilter) GetDateRangeEnd() time.Time {
	if o == nil || o.DateRangeEnd == nil {
		var ret time.Time
		return ret
	}
	return *o.DateRangeEnd
}

// GetDateRangeEndOk returns a tuple with the DateRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskApiFilter) GetDateRangeEndOk() (*time.Time, bool) {
	if o == nil || o.DateRangeEnd == nil {
		return nil, false
	}
	return o.DateRangeEnd, true
}

// HasDateRangeEnd returns a boolean if a field has been set.
func (o *TaskApiFilter) HasDateRangeEnd() bool {
	if o != nil && o.DateRangeEnd != nil {
		return true
	}

	return false
}

// SetDateRangeEnd gets a reference to the given time.Time and assigns it to the DateRangeEnd field.
func (o *TaskApiFilter) SetDateRangeEnd(v time.Time) {
	o.DateRangeEnd = &v
}

// GetDateRangeStart returns the DateRangeStart field value if set, zero value otherwise.
func (o *TaskApiFilter) GetDateRangeStart() time.Time {
	if o == nil || o.DateRangeStart == nil {
		var ret time.Time
		return ret
	}
	return *o.DateRangeStart
}

// GetDateRangeStartOk returns a tuple with the DateRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskApiFilter) GetDateRangeStartOk() (*time.Time, bool) {
	if o == nil || o.DateRangeStart == nil {
		return nil, false
	}
	return o.DateRangeStart, true
}

// HasDateRangeStart returns a boolean if a field has been set.
func (o *TaskApiFilter) HasDateRangeStart() bool {
	if o != nil && o.DateRangeStart != nil {
		return true
	}

	return false
}

// SetDateRangeStart gets a reference to the given time.Time and assigns it to the DateRangeStart field.
func (o *TaskApiFilter) SetDateRangeStart(v time.Time) {
	o.DateRangeStart = &v
}

// GetStatus returns the Status field value
func (o *TaskApiFilter) GetStatus() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskApiFilter) GetStatusOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TaskApiFilter) SetStatus(v []string) {
	o.Status = v
}

// GetTargetList returns the TargetList field value
func (o *TaskApiFilter) GetTargetList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetList
}

// GetTargetListOk returns a tuple with the TargetList field value
// and a boolean to check if the value has been set.
func (o *TaskApiFilter) GetTargetListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetList, true
}

// SetTargetList sets field value
func (o *TaskApiFilter) SetTargetList(v []string) {
	o.TargetList = v
}

// GetTargetUUIDList returns the TargetUUIDList field value
func (o *TaskApiFilter) GetTargetUUIDList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetUUIDList
}

// GetTargetUUIDListOk returns a tuple with the TargetUUIDList field value
// and a boolean to check if the value has been set.
func (o *TaskApiFilter) GetTargetUUIDListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetUUIDList, true
}

// SetTargetUUIDList sets field value
func (o *TaskApiFilter) SetTargetUUIDList(v []string) {
	o.TargetUUIDList = v
}

// GetTypeList returns the TypeList field value
func (o *TaskApiFilter) GetTypeList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TypeList
}

// GetTypeListOk returns a tuple with the TypeList field value
// and a boolean to check if the value has been set.
func (o *TaskApiFilter) GetTypeListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TypeList, true
}

// SetTypeList sets field value
func (o *TaskApiFilter) SetTypeList(v []string) {
	o.TypeList = v
}

// GetTypeNameList returns the TypeNameList field value
func (o *TaskApiFilter) GetTypeNameList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TypeNameList
}

// GetTypeNameListOk returns a tuple with the TypeNameList field value
// and a boolean to check if the value has been set.
func (o *TaskApiFilter) GetTypeNameListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TypeNameList, true
}

// SetTypeNameList sets field value
func (o *TaskApiFilter) SetTypeNameList(v []string) {
	o.TypeNameList = v
}

func (o TaskApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateRangeEnd != nil {
		toSerialize["dateRangeEnd"] = o.DateRangeEnd
	}
	if o.DateRangeStart != nil {
		toSerialize["dateRangeStart"] = o.DateRangeStart
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["targetList"] = o.TargetList
	}
	if true {
		toSerialize["targetUUIDList"] = o.TargetUUIDList
	}
	if true {
		toSerialize["typeList"] = o.TypeList
	}
	if true {
		toSerialize["typeNameList"] = o.TypeNameList
	}
	return json.Marshal(toSerialize)
}

type NullableTaskApiFilter struct {
	value *TaskApiFilter
	isSet bool
}

func (v NullableTaskApiFilter) Get() *TaskApiFilter {
	return v.value
}

func (v *NullableTaskApiFilter) Set(val *TaskApiFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskApiFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskApiFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskApiFilter(val *TaskApiFilter) *NullableTaskApiFilter {
	return &NullableTaskApiFilter{value: val, isSet: true}
}

func (v NullableTaskApiFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskApiFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


