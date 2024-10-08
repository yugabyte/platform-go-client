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
)

// JobInstanceApiFilter API filter for job instance that is a part of JobInstancePagedQuerySpec.
type JobInstanceApiFilter struct {
	// Filter by start time falling within the time window from now.
	StartWindowSecs *int64 `json:"start_window_secs,omitempty"`
	State *JobInstanceState `json:"state,omitempty"`
}

// NewJobInstanceApiFilter instantiates a new JobInstanceApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobInstanceApiFilter() *JobInstanceApiFilter {
	this := JobInstanceApiFilter{}
	return &this
}

// NewJobInstanceApiFilterWithDefaults instantiates a new JobInstanceApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobInstanceApiFilterWithDefaults() *JobInstanceApiFilter {
	this := JobInstanceApiFilter{}
	return &this
}

// GetStartWindowSecs returns the StartWindowSecs field value if set, zero value otherwise.
func (o *JobInstanceApiFilter) GetStartWindowSecs() int64 {
	if o == nil || o.StartWindowSecs == nil {
		var ret int64
		return ret
	}
	return *o.StartWindowSecs
}

// GetStartWindowSecsOk returns a tuple with the StartWindowSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstanceApiFilter) GetStartWindowSecsOk() (*int64, bool) {
	if o == nil || o.StartWindowSecs == nil {
		return nil, false
	}
	return o.StartWindowSecs, true
}

// HasStartWindowSecs returns a boolean if a field has been set.
func (o *JobInstanceApiFilter) HasStartWindowSecs() bool {
	if o != nil && o.StartWindowSecs != nil {
		return true
	}

	return false
}

// SetStartWindowSecs gets a reference to the given int64 and assigns it to the StartWindowSecs field.
func (o *JobInstanceApiFilter) SetStartWindowSecs(v int64) {
	o.StartWindowSecs = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *JobInstanceApiFilter) GetState() JobInstanceState {
	if o == nil || o.State == nil {
		var ret JobInstanceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstanceApiFilter) GetStateOk() (*JobInstanceState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *JobInstanceApiFilter) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given JobInstanceState and assigns it to the State field.
func (o *JobInstanceApiFilter) SetState(v JobInstanceState) {
	o.State = &v
}

func (o JobInstanceApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartWindowSecs != nil {
		toSerialize["start_window_secs"] = o.StartWindowSecs
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableJobInstanceApiFilter struct {
	value *JobInstanceApiFilter
	isSet bool
}

func (v NullableJobInstanceApiFilter) Get() *JobInstanceApiFilter {
	return v.value
}

func (v *NullableJobInstanceApiFilter) Set(val *JobInstanceApiFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableJobInstanceApiFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableJobInstanceApiFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobInstanceApiFilter(val *JobInstanceApiFilter) *NullableJobInstanceApiFilter {
	return &NullableJobInstanceApiFilter{value: val, isSet: true}
}

func (v NullableJobInstanceApiFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobInstanceApiFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


