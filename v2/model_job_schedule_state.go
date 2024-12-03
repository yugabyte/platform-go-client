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
	"fmt"
)

// JobScheduleState Job schedule state of a schedule that is a part of JobScheduleInfo.
type JobScheduleState string

// List of JobScheduleState
const (
	ACTIVE JobScheduleState = "ACTIVE"
	INACTIVE JobScheduleState = "INACTIVE"
)

var allowedJobScheduleStateEnumValues = []JobScheduleState{
	"ACTIVE",
	"INACTIVE",
}

func (v *JobScheduleState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobScheduleState(value)
	for _, existing := range allowedJobScheduleStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobScheduleState", value)
}

// NewJobScheduleStateFromValue returns a pointer to a valid JobScheduleState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobScheduleStateFromValue(v string) (*JobScheduleState, error) {
	ev := JobScheduleState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobScheduleState: valid values are %v", v, allowedJobScheduleStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobScheduleState) IsValid() bool {
	for _, existing := range allowedJobScheduleStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobScheduleState value
func (v JobScheduleState) Ptr() *JobScheduleState {
	return &v
}

type NullableJobScheduleState struct {
	value *JobScheduleState
	isSet bool
}

func (v NullableJobScheduleState) Get() *JobScheduleState {
	return v.value
}

func (v *NullableJobScheduleState) Set(val *JobScheduleState) {
	v.value = val
	v.isSet = true
}

func (v NullableJobScheduleState) IsSet() bool {
	return v.isSet
}

func (v *NullableJobScheduleState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobScheduleState(val *JobScheduleState) *NullableJobScheduleState {
	return &NullableJobScheduleState{value: val, isSet: true}
}

func (v NullableJobScheduleState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobScheduleState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
