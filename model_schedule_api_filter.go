/*
 * Yugabyte Platform APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ywclient

import (
	"encoding/json"
)

// ScheduleApiFilter struct for ScheduleApiFilter
type ScheduleApiFilter struct {
	Status []string `json:"status"`
	TaskTypes []string `json:"taskTypes"`
}

// NewScheduleApiFilter instantiates a new ScheduleApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleApiFilter(status []string, taskTypes []string, ) *ScheduleApiFilter {
	this := ScheduleApiFilter{}
	this.Status = status
	this.TaskTypes = taskTypes
	return &this
}

// NewScheduleApiFilterWithDefaults instantiates a new ScheduleApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleApiFilterWithDefaults() *ScheduleApiFilter {
	this := ScheduleApiFilter{}
	return &this
}

// GetStatus returns the Status field value
func (o *ScheduleApiFilter) GetStatus() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ScheduleApiFilter) GetStatusOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ScheduleApiFilter) SetStatus(v []string) {
	o.Status = v
}

// GetTaskTypes returns the TaskTypes field value
func (o *ScheduleApiFilter) GetTaskTypes() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.TaskTypes
}

// GetTaskTypesOk returns a tuple with the TaskTypes field value
// and a boolean to check if the value has been set.
func (o *ScheduleApiFilter) GetTaskTypesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TaskTypes, true
}

// SetTaskTypes sets field value
func (o *ScheduleApiFilter) SetTaskTypes(v []string) {
	o.TaskTypes = v
}

func (o ScheduleApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["taskTypes"] = o.TaskTypes
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleApiFilter struct {
	value *ScheduleApiFilter
	isSet bool
}

func (v NullableScheduleApiFilter) Get() *ScheduleApiFilter {
	return v.value
}

func (v *NullableScheduleApiFilter) Set(val *ScheduleApiFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleApiFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleApiFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleApiFilter(val *ScheduleApiFilter) *NullableScheduleApiFilter {
	return &NullableScheduleApiFilter{value: val, isSet: true}
}

func (v NullableScheduleApiFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleApiFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


