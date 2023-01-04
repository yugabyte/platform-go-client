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

// TriggerHealthCheckResult The response type for triggering a health check. Contains the timestamp of when the health check was triggered.
type TriggerHealthCheckResult struct {
	// The ISO-8601 timestamp when the health check was triggered.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewTriggerHealthCheckResult instantiates a new TriggerHealthCheckResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerHealthCheckResult() *TriggerHealthCheckResult {
	this := TriggerHealthCheckResult{}
	return &this
}

// NewTriggerHealthCheckResultWithDefaults instantiates a new TriggerHealthCheckResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerHealthCheckResultWithDefaults() *TriggerHealthCheckResult {
	this := TriggerHealthCheckResult{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TriggerHealthCheckResult) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHealthCheckResult) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TriggerHealthCheckResult) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TriggerHealthCheckResult) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o TriggerHealthCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableTriggerHealthCheckResult struct {
	value *TriggerHealthCheckResult
	isSet bool
}

func (v NullableTriggerHealthCheckResult) Get() *TriggerHealthCheckResult {
	return v.value
}

func (v *NullableTriggerHealthCheckResult) Set(val *TriggerHealthCheckResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerHealthCheckResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerHealthCheckResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerHealthCheckResult(val *TriggerHealthCheckResult) *NullableTriggerHealthCheckResult {
	return &NullableTriggerHealthCheckResult{value: val, isSet: true}
}

func (v NullableTriggerHealthCheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerHealthCheckResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

