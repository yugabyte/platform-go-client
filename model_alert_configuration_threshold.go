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

// AlertConfigurationThreshold Alert configuration threshold. Conditions can be either greater than a specified value, or less than a specified value.
type AlertConfigurationThreshold struct {
	// Threshold condition (greater than, or less than)
	Condition string `json:"condition"`
	// Threshold value
	Threshold float64 `json:"threshold"`
}

// NewAlertConfigurationThreshold instantiates a new AlertConfigurationThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertConfigurationThreshold(condition string, threshold float64) *AlertConfigurationThreshold {
	this := AlertConfigurationThreshold{}
	this.Condition = condition
	this.Threshold = threshold
	return &this
}

// NewAlertConfigurationThresholdWithDefaults instantiates a new AlertConfigurationThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertConfigurationThresholdWithDefaults() *AlertConfigurationThreshold {
	this := AlertConfigurationThreshold{}
	return &this
}

// GetCondition returns the Condition field value
func (o *AlertConfigurationThreshold) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationThreshold) GetConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *AlertConfigurationThreshold) SetCondition(v string) {
	o.Condition = v
}

// GetThreshold returns the Threshold field value
func (o *AlertConfigurationThreshold) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationThreshold) GetThresholdOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *AlertConfigurationThreshold) SetThreshold(v float64) {
	o.Threshold = v
}

func (o AlertConfigurationThreshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["condition"] = o.Condition
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	return json.Marshal(toSerialize)
}

type NullableAlertConfigurationThreshold struct {
	value *AlertConfigurationThreshold
	isSet bool
}

func (v NullableAlertConfigurationThreshold) Get() *AlertConfigurationThreshold {
	return v.value
}

func (v *NullableAlertConfigurationThreshold) Set(val *AlertConfigurationThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertConfigurationThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertConfigurationThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertConfigurationThreshold(val *AlertConfigurationThreshold) *NullableAlertConfigurationThreshold {
	return &NullableAlertConfigurationThreshold{value: val, isSet: true}
}

func (v NullableAlertConfigurationThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertConfigurationThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


