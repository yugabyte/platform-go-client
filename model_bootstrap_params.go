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

// BootstrapParams Bootstrap parameters
type BootstrapParams struct {
	BackupRequestParams BootstarpBackupParams `json:"backupRequestParams"`
	// Source Universe table IDs that need bootstrapping; must be a subset of tables in the main body
	Tables []string `json:"tables"`
}

// NewBootstrapParams instantiates a new BootstrapParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootstrapParams(backupRequestParams BootstarpBackupParams, tables []string) *BootstrapParams {
	this := BootstrapParams{}
	this.BackupRequestParams = backupRequestParams
	this.Tables = tables
	return &this
}

// NewBootstrapParamsWithDefaults instantiates a new BootstrapParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootstrapParamsWithDefaults() *BootstrapParams {
	this := BootstrapParams{}
	return &this
}

// GetBackupRequestParams returns the BackupRequestParams field value
func (o *BootstrapParams) GetBackupRequestParams() BootstarpBackupParams {
	if o == nil {
		var ret BootstarpBackupParams
		return ret
	}

	return o.BackupRequestParams
}

// GetBackupRequestParamsOk returns a tuple with the BackupRequestParams field value
// and a boolean to check if the value has been set.
func (o *BootstrapParams) GetBackupRequestParamsOk() (*BootstarpBackupParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupRequestParams, true
}

// SetBackupRequestParams sets field value
func (o *BootstrapParams) SetBackupRequestParams(v BootstarpBackupParams) {
	o.BackupRequestParams = v
}

// GetTables returns the Tables field value
func (o *BootstrapParams) GetTables() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value
// and a boolean to check if the value has been set.
func (o *BootstrapParams) GetTablesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tables, true
}

// SetTables sets field value
func (o *BootstrapParams) SetTables(v []string) {
	o.Tables = v
}

func (o BootstrapParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupRequestParams"] = o.BackupRequestParams
	}
	if true {
		toSerialize["tables"] = o.Tables
	}
	return json.Marshal(toSerialize)
}

type NullableBootstrapParams struct {
	value *BootstrapParams
	isSet bool
}

func (v NullableBootstrapParams) Get() *BootstrapParams {
	return v.value
}

func (v *NullableBootstrapParams) Set(val *BootstrapParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBootstrapParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBootstrapParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootstrapParams(val *BootstrapParams) *NullableBootstrapParams {
	return &NullableBootstrapParams{value: val, isSet: true}
}

func (v NullableBootstrapParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootstrapParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


