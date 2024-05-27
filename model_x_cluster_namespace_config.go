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

// XClusterNamespaceConfig struct for XClusterNamespaceConfig
type XClusterNamespaceConfig struct {
	SourceNamespaceId string `json:"sourceNamespaceId"`
}

// NewXClusterNamespaceConfig instantiates a new XClusterNamespaceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXClusterNamespaceConfig(sourceNamespaceId string) *XClusterNamespaceConfig {
	this := XClusterNamespaceConfig{}
	this.SourceNamespaceId = sourceNamespaceId
	return &this
}

// NewXClusterNamespaceConfigWithDefaults instantiates a new XClusterNamespaceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXClusterNamespaceConfigWithDefaults() *XClusterNamespaceConfig {
	this := XClusterNamespaceConfig{}
	return &this
}

// GetSourceNamespaceId returns the SourceNamespaceId field value
func (o *XClusterNamespaceConfig) GetSourceNamespaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceNamespaceId
}

// GetSourceNamespaceIdOk returns a tuple with the SourceNamespaceId field value
// and a boolean to check if the value has been set.
func (o *XClusterNamespaceConfig) GetSourceNamespaceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceNamespaceId, true
}

// SetSourceNamespaceId sets field value
func (o *XClusterNamespaceConfig) SetSourceNamespaceId(v string) {
	o.SourceNamespaceId = v
}

func (o XClusterNamespaceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceNamespaceId"] = o.SourceNamespaceId
	}
	return json.Marshal(toSerialize)
}

type NullableXClusterNamespaceConfig struct {
	value *XClusterNamespaceConfig
	isSet bool
}

func (v NullableXClusterNamespaceConfig) Get() *XClusterNamespaceConfig {
	return v.value
}

func (v *NullableXClusterNamespaceConfig) Set(val *XClusterNamespaceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableXClusterNamespaceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableXClusterNamespaceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXClusterNamespaceConfig(val *XClusterNamespaceConfig) *NullableXClusterNamespaceConfig {
	return &NullableXClusterNamespaceConfig{value: val, isSet: true}
}

func (v NullableXClusterNamespaceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXClusterNamespaceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

