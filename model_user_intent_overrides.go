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
)

// UserIntentOverrides WARNING: This is a preview API that could change. User Intent overrides
type UserIntentOverrides struct {
	AzOverrides *map[string]AZOverrides `json:"azOverrides,omitempty"`
	PerProcess *map[string]PerProcessDetails `json:"perProcess,omitempty"`
}

// NewUserIntentOverrides instantiates a new UserIntentOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIntentOverrides() *UserIntentOverrides {
	this := UserIntentOverrides{}
	return &this
}

// NewUserIntentOverridesWithDefaults instantiates a new UserIntentOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIntentOverridesWithDefaults() *UserIntentOverrides {
	this := UserIntentOverrides{}
	return &this
}

// GetAzOverrides returns the AzOverrides field value if set, zero value otherwise.
func (o *UserIntentOverrides) GetAzOverrides() map[string]AZOverrides {
	if o == nil || o.AzOverrides == nil {
		var ret map[string]AZOverrides
		return ret
	}
	return *o.AzOverrides
}

// GetAzOverridesOk returns a tuple with the AzOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIntentOverrides) GetAzOverridesOk() (*map[string]AZOverrides, bool) {
	if o == nil || o.AzOverrides == nil {
		return nil, false
	}
	return o.AzOverrides, true
}

// HasAzOverrides returns a boolean if a field has been set.
func (o *UserIntentOverrides) HasAzOverrides() bool {
	if o != nil && o.AzOverrides != nil {
		return true
	}

	return false
}

// SetAzOverrides gets a reference to the given map[string]AZOverrides and assigns it to the AzOverrides field.
func (o *UserIntentOverrides) SetAzOverrides(v map[string]AZOverrides) {
	o.AzOverrides = &v
}

// GetPerProcess returns the PerProcess field value if set, zero value otherwise.
func (o *UserIntentOverrides) GetPerProcess() map[string]PerProcessDetails {
	if o == nil || o.PerProcess == nil {
		var ret map[string]PerProcessDetails
		return ret
	}
	return *o.PerProcess
}

// GetPerProcessOk returns a tuple with the PerProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIntentOverrides) GetPerProcessOk() (*map[string]PerProcessDetails, bool) {
	if o == nil || o.PerProcess == nil {
		return nil, false
	}
	return o.PerProcess, true
}

// HasPerProcess returns a boolean if a field has been set.
func (o *UserIntentOverrides) HasPerProcess() bool {
	if o != nil && o.PerProcess != nil {
		return true
	}

	return false
}

// SetPerProcess gets a reference to the given map[string]PerProcessDetails and assigns it to the PerProcess field.
func (o *UserIntentOverrides) SetPerProcess(v map[string]PerProcessDetails) {
	o.PerProcess = &v
}

func (o UserIntentOverrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzOverrides != nil {
		toSerialize["azOverrides"] = o.AzOverrides
	}
	if o.PerProcess != nil {
		toSerialize["perProcess"] = o.PerProcess
	}
	return json.Marshal(toSerialize)
}

type NullableUserIntentOverrides struct {
	value *UserIntentOverrides
	isSet bool
}

func (v NullableUserIntentOverrides) Get() *UserIntentOverrides {
	return v.value
}

func (v *NullableUserIntentOverrides) Set(val *UserIntentOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIntentOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIntentOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIntentOverrides(val *UserIntentOverrides) *NullableUserIntentOverrides {
	return &NullableUserIntentOverrides{value: val, isSet: true}
}

func (v NullableUserIntentOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIntentOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

