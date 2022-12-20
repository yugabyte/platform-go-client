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

// XClusterConfigNeedBootstrapFormData struct for XClusterConfigNeedBootstrapFormData
type XClusterConfigNeedBootstrapFormData struct {
	// Source universe table IDs to check whether they need bootstrap
	Tables []string `json:"tables"`
}

// NewXClusterConfigNeedBootstrapFormData instantiates a new XClusterConfigNeedBootstrapFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXClusterConfigNeedBootstrapFormData(tables []string, ) *XClusterConfigNeedBootstrapFormData {
	this := XClusterConfigNeedBootstrapFormData{}
	this.Tables = tables
	return &this
}

// NewXClusterConfigNeedBootstrapFormDataWithDefaults instantiates a new XClusterConfigNeedBootstrapFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXClusterConfigNeedBootstrapFormDataWithDefaults() *XClusterConfigNeedBootstrapFormData {
	this := XClusterConfigNeedBootstrapFormData{}
	return &this
}

// GetTables returns the Tables field value
func (o *XClusterConfigNeedBootstrapFormData) GetTables() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value
// and a boolean to check if the value has been set.
func (o *XClusterConfigNeedBootstrapFormData) GetTablesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tables, true
}

// SetTables sets field value
func (o *XClusterConfigNeedBootstrapFormData) SetTables(v []string) {
	o.Tables = v
}

func (o XClusterConfigNeedBootstrapFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tables"] = o.Tables
	}
	return json.Marshal(toSerialize)
}

type NullableXClusterConfigNeedBootstrapFormData struct {
	value *XClusterConfigNeedBootstrapFormData
	isSet bool
}

func (v NullableXClusterConfigNeedBootstrapFormData) Get() *XClusterConfigNeedBootstrapFormData {
	return v.value
}

func (v *NullableXClusterConfigNeedBootstrapFormData) Set(val *XClusterConfigNeedBootstrapFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableXClusterConfigNeedBootstrapFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableXClusterConfigNeedBootstrapFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXClusterConfigNeedBootstrapFormData(val *XClusterConfigNeedBootstrapFormData) *NullableXClusterConfigNeedBootstrapFormData {
	return &NullableXClusterConfigNeedBootstrapFormData{value: val, isSet: true}
}

func (v NullableXClusterConfigNeedBootstrapFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXClusterConfigNeedBootstrapFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


