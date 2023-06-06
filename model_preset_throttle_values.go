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

// PresetThrottleValues struct for PresetThrottleValues
type PresetThrottleValues struct {
	DefaultValue int32 `json:"defaultValue"`
	MaxValue int32 `json:"maxValue"`
	MinValue int32 `json:"minValue"`
}

// NewPresetThrottleValues instantiates a new PresetThrottleValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresetThrottleValues(defaultValue int32, maxValue int32, minValue int32, ) *PresetThrottleValues {
	this := PresetThrottleValues{}
	this.DefaultValue = defaultValue
	this.MaxValue = maxValue
	this.MinValue = minValue
	return &this
}

// NewPresetThrottleValuesWithDefaults instantiates a new PresetThrottleValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresetThrottleValuesWithDefaults() *PresetThrottleValues {
	this := PresetThrottleValues{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value
func (o *PresetThrottleValues) GetDefaultValue() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *PresetThrottleValues) GetDefaultValueOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value
func (o *PresetThrottleValues) SetDefaultValue(v int32) {
	o.DefaultValue = v
}

// GetMaxValue returns the MaxValue field value
func (o *PresetThrottleValues) GetMaxValue() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value
// and a boolean to check if the value has been set.
func (o *PresetThrottleValues) GetMaxValueOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxValue, true
}

// SetMaxValue sets field value
func (o *PresetThrottleValues) SetMaxValue(v int32) {
	o.MaxValue = v
}

// GetMinValue returns the MinValue field value
func (o *PresetThrottleValues) GetMinValue() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value
// and a boolean to check if the value has been set.
func (o *PresetThrottleValues) GetMinValueOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinValue, true
}

// SetMinValue sets field value
func (o *PresetThrottleValues) SetMinValue(v int32) {
	o.MinValue = v
}

func (o PresetThrottleValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if true {
		toSerialize["maxValue"] = o.MaxValue
	}
	if true {
		toSerialize["minValue"] = o.MinValue
	}
	return json.Marshal(toSerialize)
}

type NullablePresetThrottleValues struct {
	value *PresetThrottleValues
	isSet bool
}

func (v NullablePresetThrottleValues) Get() *PresetThrottleValues {
	return v.value
}

func (v *NullablePresetThrottleValues) Set(val *PresetThrottleValues) {
	v.value = val
	v.isSet = true
}

func (v NullablePresetThrottleValues) IsSet() bool {
	return v.isSet
}

func (v *NullablePresetThrottleValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresetThrottleValues(val *PresetThrottleValues) *NullablePresetThrottleValues {
	return &NullablePresetThrottleValues{value: val, isSet: true}
}

func (v NullablePresetThrottleValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresetThrottleValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

