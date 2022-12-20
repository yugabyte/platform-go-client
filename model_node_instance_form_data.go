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

// NodeInstanceFormData struct for NodeInstanceFormData
type NodeInstanceFormData struct {
	// Node instances
	Nodes []NodeInstanceData `json:"nodes"`
}

// NewNodeInstanceFormData instantiates a new NodeInstanceFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeInstanceFormData(nodes []NodeInstanceData, ) *NodeInstanceFormData {
	this := NodeInstanceFormData{}
	this.Nodes = nodes
	return &this
}

// NewNodeInstanceFormDataWithDefaults instantiates a new NodeInstanceFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeInstanceFormDataWithDefaults() *NodeInstanceFormData {
	this := NodeInstanceFormData{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *NodeInstanceFormData) GetNodes() []NodeInstanceData {
	if o == nil  {
		var ret []NodeInstanceData
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *NodeInstanceFormData) GetNodesOk() (*[]NodeInstanceData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *NodeInstanceFormData) SetNodes(v []NodeInstanceData) {
	o.Nodes = v
}

func (o NodeInstanceFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nodes"] = o.Nodes
	}
	return json.Marshal(toSerialize)
}

type NullableNodeInstanceFormData struct {
	value *NodeInstanceFormData
	isSet bool
}

func (v NullableNodeInstanceFormData) Get() *NodeInstanceFormData {
	return v.value
}

func (v *NullableNodeInstanceFormData) Set(val *NodeInstanceFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeInstanceFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeInstanceFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeInstanceFormData(val *NodeInstanceFormData) *NullableNodeInstanceFormData {
	return &NullableNodeInstanceFormData{value: val, isSet: true}
}

func (v NullableNodeInstanceFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeInstanceFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


