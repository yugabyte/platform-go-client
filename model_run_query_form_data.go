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

// RunQueryFormData struct for RunQueryFormData
type RunQueryFormData struct {
	DbName string `json:"db_name"`
	NodeName string `json:"node_name"`
	Query string `json:"query"`
	TableType string `json:"table_type"`
}

// NewRunQueryFormData instantiates a new RunQueryFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunQueryFormData(dbName string, nodeName string, query string, tableType string) *RunQueryFormData {
	this := RunQueryFormData{}
	this.DbName = dbName
	this.NodeName = nodeName
	this.Query = query
	this.TableType = tableType
	return &this
}

// NewRunQueryFormDataWithDefaults instantiates a new RunQueryFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunQueryFormDataWithDefaults() *RunQueryFormData {
	this := RunQueryFormData{}
	return &this
}

// GetDbName returns the DbName field value
func (o *RunQueryFormData) GetDbName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *RunQueryFormData) GetDbNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value
func (o *RunQueryFormData) SetDbName(v string) {
	o.DbName = v
}

// GetNodeName returns the NodeName field value
func (o *RunQueryFormData) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *RunQueryFormData) GetNodeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *RunQueryFormData) SetNodeName(v string) {
	o.NodeName = v
}

// GetQuery returns the Query field value
func (o *RunQueryFormData) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *RunQueryFormData) GetQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *RunQueryFormData) SetQuery(v string) {
	o.Query = v
}

// GetTableType returns the TableType field value
func (o *RunQueryFormData) GetTableType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value
// and a boolean to check if the value has been set.
func (o *RunQueryFormData) GetTableTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TableType, true
}

// SetTableType sets field value
func (o *RunQueryFormData) SetTableType(v string) {
	o.TableType = v
}

func (o RunQueryFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["db_name"] = o.DbName
	}
	if true {
		toSerialize["node_name"] = o.NodeName
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["table_type"] = o.TableType
	}
	return json.Marshal(toSerialize)
}

type NullableRunQueryFormData struct {
	value *RunQueryFormData
	isSet bool
}

func (v NullableRunQueryFormData) Get() *RunQueryFormData {
	return v.value
}

func (v *NullableRunQueryFormData) Set(val *RunQueryFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableRunQueryFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableRunQueryFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunQueryFormData(val *RunQueryFormData) *NullableRunQueryFormData {
	return &NullableRunQueryFormData{value: val, isSet: true}
}

func (v NullableRunQueryFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunQueryFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


