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

// EditBackupParams Edit backup parameters
type EditBackupParams struct {
	// Time before deleting the backup from storage, in milliseconds
	TimeBeforeDeleteFromPresentInMillis int64 `json:"timeBeforeDeleteFromPresentInMillis"`
}

// NewEditBackupParams instantiates a new EditBackupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditBackupParams(timeBeforeDeleteFromPresentInMillis int64, ) *EditBackupParams {
	this := EditBackupParams{}
	this.TimeBeforeDeleteFromPresentInMillis = timeBeforeDeleteFromPresentInMillis
	return &this
}

// NewEditBackupParamsWithDefaults instantiates a new EditBackupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditBackupParamsWithDefaults() *EditBackupParams {
	this := EditBackupParams{}
	return &this
}

// GetTimeBeforeDeleteFromPresentInMillis returns the TimeBeforeDeleteFromPresentInMillis field value
func (o *EditBackupParams) GetTimeBeforeDeleteFromPresentInMillis() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.TimeBeforeDeleteFromPresentInMillis
}

// GetTimeBeforeDeleteFromPresentInMillisOk returns a tuple with the TimeBeforeDeleteFromPresentInMillis field value
// and a boolean to check if the value has been set.
func (o *EditBackupParams) GetTimeBeforeDeleteFromPresentInMillisOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeBeforeDeleteFromPresentInMillis, true
}

// SetTimeBeforeDeleteFromPresentInMillis sets field value
func (o *EditBackupParams) SetTimeBeforeDeleteFromPresentInMillis(v int64) {
	o.TimeBeforeDeleteFromPresentInMillis = v
}

func (o EditBackupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timeBeforeDeleteFromPresentInMillis"] = o.TimeBeforeDeleteFromPresentInMillis
	}
	return json.Marshal(toSerialize)
}

type NullableEditBackupParams struct {
	value *EditBackupParams
	isSet bool
}

func (v NullableEditBackupParams) Get() *EditBackupParams {
	return v.value
}

func (v *NullableEditBackupParams) Set(val *EditBackupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEditBackupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEditBackupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditBackupParams(val *EditBackupParams) *NullableEditBackupParams {
	return &NullableEditBackupParams{value: val, isSet: true}
}

func (v NullableEditBackupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditBackupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


