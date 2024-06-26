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

// PerProcessFlags struct for PerProcessFlags
type PerProcessFlags struct {
	Value map[string]map[string]string `json:"value"`
}

// NewPerProcessFlags instantiates a new PerProcessFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerProcessFlags(value map[string]map[string]string) *PerProcessFlags {
	this := PerProcessFlags{}
	this.Value = value
	return &this
}

// NewPerProcessFlagsWithDefaults instantiates a new PerProcessFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerProcessFlagsWithDefaults() *PerProcessFlags {
	this := PerProcessFlags{}
	return &this
}

// GetValue returns the Value field value
func (o *PerProcessFlags) GetValue() map[string]map[string]string {
	if o == nil {
		var ret map[string]map[string]string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PerProcessFlags) GetValueOk() (*map[string]map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PerProcessFlags) SetValue(v map[string]map[string]string) {
	o.Value = v
}

func (o PerProcessFlags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePerProcessFlags struct {
	value *PerProcessFlags
	isSet bool
}

func (v NullablePerProcessFlags) Get() *PerProcessFlags {
	return v.value
}

func (v *NullablePerProcessFlags) Set(val *PerProcessFlags) {
	v.value = val
	v.isSet = true
}

func (v NullablePerProcessFlags) IsSet() bool {
	return v.isSet
}

func (v *NullablePerProcessFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerProcessFlags(val *PerProcessFlags) *NullablePerProcessFlags {
	return &NullablePerProcessFlags{value: val, isSet: true}
}

func (v NullablePerProcessFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerProcessFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


