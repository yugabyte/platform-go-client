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

// NodeConfig A node configuration.
type NodeConfig struct {
	Type string `json:"type"`
	// true
	Value string `json:"value"`
}

// NewNodeConfig instantiates a new NodeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeConfig(type_ string, value string) *NodeConfig {
	this := NodeConfig{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewNodeConfigWithDefaults instantiates a new NodeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeConfigWithDefaults() *NodeConfig {
	this := NodeConfig{}
	return &this
}

// GetType returns the Type field value
func (o *NodeConfig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NodeConfig) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NodeConfig) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *NodeConfig) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NodeConfig) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NodeConfig) SetValue(v string) {
	o.Value = v
}

func (o NodeConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNodeConfig struct {
	value *NodeConfig
	isSet bool
}

func (v NullableNodeConfig) Get() *NodeConfig {
	return v.value
}

func (v *NullableNodeConfig) Set(val *NodeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeConfig(val *NodeConfig) *NullableNodeConfig {
	return &NullableNodeConfig{value: val, isSet: true}
}

func (v NullableNodeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


