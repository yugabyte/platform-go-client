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

// MetricValue struct for MetricValue
type MetricValue struct {
	Labels []MetricLabel `json:"labels"`
	Value float64 `json:"value"`
}

// NewMetricValue instantiates a new MetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricValue(labels []MetricLabel, value float64) *MetricValue {
	this := MetricValue{}
	this.Labels = labels
	this.Value = value
	return &this
}

// NewMetricValueWithDefaults instantiates a new MetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricValueWithDefaults() *MetricValue {
	this := MetricValue{}
	return &this
}

// GetLabels returns the Labels field value
func (o *MetricValue) GetLabels() []MetricLabel {
	if o == nil {
		var ret []MetricLabel
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *MetricValue) GetLabelsOk() (*[]MetricLabel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *MetricValue) SetLabels(v []MetricLabel) {
	o.Labels = v
}

// GetValue returns the Value field value
func (o *MetricValue) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MetricValue) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MetricValue) SetValue(v float64) {
	o.Value = v
}

func (o MetricValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMetricValue struct {
	value *MetricValue
	isSet bool
}

func (v NullableMetricValue) Get() *MetricValue {
	return v.value
}

func (v *NullableMetricValue) Set(val *MetricValue) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricValue) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricValue(val *MetricValue) *NullableMetricValue {
	return &NullableMetricValue{value: val, isSet: true}
}

func (v NullableMetricValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


