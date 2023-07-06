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

// NodeData struct for NodeData
type NodeData struct {
	Details []string `json:"details"`
	HasError bool `json:"has_error"`
	HasWarning bool `json:"has_warning"`
	Message string `json:"message"`
	Metrics []Metric `json:"metrics"`
	MetricsOnly bool `json:"metrics_only"`
	Node string `json:"node"`
	NodeName string `json:"node_name"`
	Process string `json:"process"`
	TimestampIso *time.Time `json:"timestamp_iso,omitempty"`
}

// NewNodeData instantiates a new NodeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeData(details []string, hasError bool, hasWarning bool, message string, metrics []Metric, metricsOnly bool, node string, nodeName string, process string, ) *NodeData {
	this := NodeData{}
	this.Details = details
	this.HasError = hasError
	this.HasWarning = hasWarning
	this.Message = message
	this.Metrics = metrics
	this.MetricsOnly = metricsOnly
	this.Node = node
	this.NodeName = nodeName
	this.Process = process
	return &this
}

// NewNodeDataWithDefaults instantiates a new NodeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeDataWithDefaults() *NodeData {
	this := NodeData{}
	return &this
}

// GetDetails returns the Details field value
func (o *NodeData) GetDetails() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetDetailsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *NodeData) SetDetails(v []string) {
	o.Details = v
}

// GetHasError returns the HasError field value
func (o *NodeData) GetHasError() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HasError
}

// GetHasErrorOk returns a tuple with the HasError field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetHasErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasError, true
}

// SetHasError sets field value
func (o *NodeData) SetHasError(v bool) {
	o.HasError = v
}

// GetHasWarning returns the HasWarning field value
func (o *NodeData) GetHasWarning() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HasWarning
}

// GetHasWarningOk returns a tuple with the HasWarning field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetHasWarningOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasWarning, true
}

// SetHasWarning sets field value
func (o *NodeData) SetHasWarning(v bool) {
	o.HasWarning = v
}

// GetMessage returns the Message field value
func (o *NodeData) GetMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NodeData) SetMessage(v string) {
	o.Message = v
}

// GetMetrics returns the Metrics field value
func (o *NodeData) GetMetrics() []Metric {
	if o == nil  {
		var ret []Metric
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetMetricsOk() (*[]Metric, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *NodeData) SetMetrics(v []Metric) {
	o.Metrics = v
}

// GetMetricsOnly returns the MetricsOnly field value
func (o *NodeData) GetMetricsOnly() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.MetricsOnly
}

// GetMetricsOnlyOk returns a tuple with the MetricsOnly field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetMetricsOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricsOnly, true
}

// SetMetricsOnly sets field value
func (o *NodeData) SetMetricsOnly(v bool) {
	o.MetricsOnly = v
}

// GetNode returns the Node field value
func (o *NodeData) GetNode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetNodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *NodeData) SetNode(v string) {
	o.Node = v
}

// GetNodeName returns the NodeName field value
func (o *NodeData) GetNodeName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetNodeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *NodeData) SetNodeName(v string) {
	o.NodeName = v
}

// GetProcess returns the Process field value
func (o *NodeData) GetProcess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Process
}

// GetProcessOk returns a tuple with the Process field value
// and a boolean to check if the value has been set.
func (o *NodeData) GetProcessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Process, true
}

// SetProcess sets field value
func (o *NodeData) SetProcess(v string) {
	o.Process = v
}

// GetTimestampIso returns the TimestampIso field value if set, zero value otherwise.
func (o *NodeData) GetTimestampIso() time.Time {
	if o == nil || o.TimestampIso == nil {
		var ret time.Time
		return ret
	}
	return *o.TimestampIso
}

// GetTimestampIsoOk returns a tuple with the TimestampIso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeData) GetTimestampIsoOk() (*time.Time, bool) {
	if o == nil || o.TimestampIso == nil {
		return nil, false
	}
	return o.TimestampIso, true
}

// HasTimestampIso returns a boolean if a field has been set.
func (o *NodeData) HasTimestampIso() bool {
	if o != nil && o.TimestampIso != nil {
		return true
	}

	return false
}

// SetTimestampIso gets a reference to the given time.Time and assigns it to the TimestampIso field.
func (o *NodeData) SetTimestampIso(v time.Time) {
	o.TimestampIso = &v
}

func (o NodeData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["details"] = o.Details
	}
	if true {
		toSerialize["has_error"] = o.HasError
	}
	if true {
		toSerialize["has_warning"] = o.HasWarning
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["metrics"] = o.Metrics
	}
	if true {
		toSerialize["metrics_only"] = o.MetricsOnly
	}
	if true {
		toSerialize["node"] = o.Node
	}
	if true {
		toSerialize["node_name"] = o.NodeName
	}
	if true {
		toSerialize["process"] = o.Process
	}
	if o.TimestampIso != nil {
		toSerialize["timestamp_iso"] = o.TimestampIso
	}
	return json.Marshal(toSerialize)
}

type NullableNodeData struct {
	value *NodeData
	isSet bool
}

func (v NullableNodeData) Get() *NodeData {
	return v.value
}

func (v *NullableNodeData) Set(val *NodeData) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeData) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeData(val *NodeData) *NullableNodeData {
	return &NullableNodeData{value: val, isSet: true}
}

func (v NullableNodeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


