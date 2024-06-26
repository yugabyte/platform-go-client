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

// CustomCertInfo struct for CustomCertInfo
type CustomCertInfo struct {
	ClientCertPath string `json:"clientCertPath"`
	ClientKeyPath string `json:"clientKeyPath"`
	NodeCertPath string `json:"nodeCertPath"`
	NodeKeyPath string `json:"nodeKeyPath"`
	RootCertPath string `json:"rootCertPath"`
}

// NewCustomCertInfo instantiates a new CustomCertInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomCertInfo(clientCertPath string, clientKeyPath string, nodeCertPath string, nodeKeyPath string, rootCertPath string) *CustomCertInfo {
	this := CustomCertInfo{}
	this.ClientCertPath = clientCertPath
	this.ClientKeyPath = clientKeyPath
	this.NodeCertPath = nodeCertPath
	this.NodeKeyPath = nodeKeyPath
	this.RootCertPath = rootCertPath
	return &this
}

// NewCustomCertInfoWithDefaults instantiates a new CustomCertInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomCertInfoWithDefaults() *CustomCertInfo {
	this := CustomCertInfo{}
	return &this
}

// GetClientCertPath returns the ClientCertPath field value
func (o *CustomCertInfo) GetClientCertPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientCertPath
}

// GetClientCertPathOk returns a tuple with the ClientCertPath field value
// and a boolean to check if the value has been set.
func (o *CustomCertInfo) GetClientCertPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientCertPath, true
}

// SetClientCertPath sets field value
func (o *CustomCertInfo) SetClientCertPath(v string) {
	o.ClientCertPath = v
}

// GetClientKeyPath returns the ClientKeyPath field value
func (o *CustomCertInfo) GetClientKeyPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientKeyPath
}

// GetClientKeyPathOk returns a tuple with the ClientKeyPath field value
// and a boolean to check if the value has been set.
func (o *CustomCertInfo) GetClientKeyPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientKeyPath, true
}

// SetClientKeyPath sets field value
func (o *CustomCertInfo) SetClientKeyPath(v string) {
	o.ClientKeyPath = v
}

// GetNodeCertPath returns the NodeCertPath field value
func (o *CustomCertInfo) GetNodeCertPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeCertPath
}

// GetNodeCertPathOk returns a tuple with the NodeCertPath field value
// and a boolean to check if the value has been set.
func (o *CustomCertInfo) GetNodeCertPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeCertPath, true
}

// SetNodeCertPath sets field value
func (o *CustomCertInfo) SetNodeCertPath(v string) {
	o.NodeCertPath = v
}

// GetNodeKeyPath returns the NodeKeyPath field value
func (o *CustomCertInfo) GetNodeKeyPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeKeyPath
}

// GetNodeKeyPathOk returns a tuple with the NodeKeyPath field value
// and a boolean to check if the value has been set.
func (o *CustomCertInfo) GetNodeKeyPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeKeyPath, true
}

// SetNodeKeyPath sets field value
func (o *CustomCertInfo) SetNodeKeyPath(v string) {
	o.NodeKeyPath = v
}

// GetRootCertPath returns the RootCertPath field value
func (o *CustomCertInfo) GetRootCertPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootCertPath
}

// GetRootCertPathOk returns a tuple with the RootCertPath field value
// and a boolean to check if the value has been set.
func (o *CustomCertInfo) GetRootCertPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RootCertPath, true
}

// SetRootCertPath sets field value
func (o *CustomCertInfo) SetRootCertPath(v string) {
	o.RootCertPath = v
}

func (o CustomCertInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientCertPath"] = o.ClientCertPath
	}
	if true {
		toSerialize["clientKeyPath"] = o.ClientKeyPath
	}
	if true {
		toSerialize["nodeCertPath"] = o.NodeCertPath
	}
	if true {
		toSerialize["nodeKeyPath"] = o.NodeKeyPath
	}
	if true {
		toSerialize["rootCertPath"] = o.RootCertPath
	}
	return json.Marshal(toSerialize)
}

type NullableCustomCertInfo struct {
	value *CustomCertInfo
	isSet bool
}

func (v NullableCustomCertInfo) Get() *CustomCertInfo {
	return v.value
}

func (v *NullableCustomCertInfo) Set(val *CustomCertInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomCertInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomCertInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomCertInfo(val *CustomCertInfo) *NullableCustomCertInfo {
	return &NullableCustomCertInfo{value: val, isSet: true}
}

func (v NullableCustomCertInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomCertInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


