/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// NodeActionFormData struct for NodeActionFormData
type NodeActionFormData struct {
	// Should ignore errors and proceed with the node action
	Force *bool `json:"force,omitempty"`
	NodeAction string `json:"nodeAction"`
}

// NewNodeActionFormData instantiates a new NodeActionFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeActionFormData(nodeAction string) *NodeActionFormData {
	this := NodeActionFormData{}
	this.NodeAction = nodeAction
	return &this
}

// NewNodeActionFormDataWithDefaults instantiates a new NodeActionFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeActionFormDataWithDefaults() *NodeActionFormData {
	this := NodeActionFormData{}
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *NodeActionFormData) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeActionFormData) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *NodeActionFormData) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *NodeActionFormData) SetForce(v bool) {
	o.Force = &v
}

// GetNodeAction returns the NodeAction field value
func (o *NodeActionFormData) GetNodeAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeAction
}

// GetNodeActionOk returns a tuple with the NodeAction field value
// and a boolean to check if the value has been set.
func (o *NodeActionFormData) GetNodeActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeAction, true
}

// SetNodeAction sets field value
func (o *NodeActionFormData) SetNodeAction(v string) {
	o.NodeAction = v
}

func (o NodeActionFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	if true {
		toSerialize["nodeAction"] = o.NodeAction
	}
	return json.Marshal(toSerialize)
}

type NullableNodeActionFormData struct {
	value *NodeActionFormData
	isSet bool
}

func (v NullableNodeActionFormData) Get() *NodeActionFormData {
	return v.value
}

func (v *NullableNodeActionFormData) Set(val *NodeActionFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeActionFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeActionFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeActionFormData(val *NodeActionFormData) *NullableNodeActionFormData {
	return &NullableNodeActionFormData{value: val, isSet: true}
}

func (v NullableNodeActionFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeActionFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


