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

// MetricQueryParams Metrics request data
type MetricQueryParams struct {
	// End time
	End *int64 `json:"end,omitempty"`
	// Is Recharts
	IsRecharts *bool `json:"isRecharts,omitempty"`
	// Metrics
	Metrics []string `json:"metrics"`
	// Node name
	NodeName *string `json:"nodeName,omitempty"`
	// Node prefix
	NodePrefix *string `json:"nodePrefix,omitempty"`
	// Start time
	Start int64 `json:"start"`
}

// NewMetricQueryParams instantiates a new MetricQueryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricQueryParams(metrics []string, start int64, ) *MetricQueryParams {
	this := MetricQueryParams{}
	this.Metrics = metrics
	this.Start = start
	return &this
}

// NewMetricQueryParamsWithDefaults instantiates a new MetricQueryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricQueryParamsWithDefaults() *MetricQueryParams {
	this := MetricQueryParams{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *MetricQueryParams) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricQueryParams) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *MetricQueryParams) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *MetricQueryParams) SetEnd(v int64) {
	o.End = &v
}

// GetIsRecharts returns the IsRecharts field value if set, zero value otherwise.
func (o *MetricQueryParams) GetIsRecharts() bool {
	if o == nil || o.IsRecharts == nil {
		var ret bool
		return ret
	}
	return *o.IsRecharts
}

// GetIsRechartsOk returns a tuple with the IsRecharts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricQueryParams) GetIsRechartsOk() (*bool, bool) {
	if o == nil || o.IsRecharts == nil {
		return nil, false
	}
	return o.IsRecharts, true
}

// HasIsRecharts returns a boolean if a field has been set.
func (o *MetricQueryParams) HasIsRecharts() bool {
	if o != nil && o.IsRecharts != nil {
		return true
	}

	return false
}

// SetIsRecharts gets a reference to the given bool and assigns it to the IsRecharts field.
func (o *MetricQueryParams) SetIsRecharts(v bool) {
	o.IsRecharts = &v
}

// GetMetrics returns the Metrics field value
func (o *MetricQueryParams) GetMetrics() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *MetricQueryParams) GetMetricsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *MetricQueryParams) SetMetrics(v []string) {
	o.Metrics = v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *MetricQueryParams) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricQueryParams) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *MetricQueryParams) HasNodeName() bool {
	if o != nil && o.NodeName != nil {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *MetricQueryParams) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodePrefix returns the NodePrefix field value if set, zero value otherwise.
func (o *MetricQueryParams) GetNodePrefix() string {
	if o == nil || o.NodePrefix == nil {
		var ret string
		return ret
	}
	return *o.NodePrefix
}

// GetNodePrefixOk returns a tuple with the NodePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricQueryParams) GetNodePrefixOk() (*string, bool) {
	if o == nil || o.NodePrefix == nil {
		return nil, false
	}
	return o.NodePrefix, true
}

// HasNodePrefix returns a boolean if a field has been set.
func (o *MetricQueryParams) HasNodePrefix() bool {
	if o != nil && o.NodePrefix != nil {
		return true
	}

	return false
}

// SetNodePrefix gets a reference to the given string and assigns it to the NodePrefix field.
func (o *MetricQueryParams) SetNodePrefix(v string) {
	o.NodePrefix = &v
}

// GetStart returns the Start field value
func (o *MetricQueryParams) GetStart() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *MetricQueryParams) GetStartOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *MetricQueryParams) SetStart(v int64) {
	o.Start = v
}

func (o MetricQueryParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.IsRecharts != nil {
		toSerialize["isRecharts"] = o.IsRecharts
	}
	if true {
		toSerialize["metrics"] = o.Metrics
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.NodePrefix != nil {
		toSerialize["nodePrefix"] = o.NodePrefix
	}
	if true {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableMetricQueryParams struct {
	value *MetricQueryParams
	isSet bool
}

func (v NullableMetricQueryParams) Get() *MetricQueryParams {
	return v.value
}

func (v *NullableMetricQueryParams) Set(val *MetricQueryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricQueryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricQueryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricQueryParams(val *MetricQueryParams) *NullableMetricQueryParams {
	return &NullableMetricQueryParams{value: val, isSet: true}
}

func (v NullableMetricQueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricQueryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


