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

// DrConfigSetTablesForm dr config set tables form
type DrConfigSetTablesForm struct {
	// Whether or not YBA should also include all index tables from any provided main tables.
	AutoIncludeIndexTables *bool `json:"autoIncludeIndexTables,omitempty"`
	BootstrapParams *RestartBootstrapParams `json:"bootstrapParams,omitempty"`
	// Source universe table IDs
	Tables *[]string `json:"tables,omitempty"`
}

// NewDrConfigSetTablesForm instantiates a new DrConfigSetTablesForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrConfigSetTablesForm() *DrConfigSetTablesForm {
	this := DrConfigSetTablesForm{}
	return &this
}

// NewDrConfigSetTablesFormWithDefaults instantiates a new DrConfigSetTablesForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrConfigSetTablesFormWithDefaults() *DrConfigSetTablesForm {
	this := DrConfigSetTablesForm{}
	return &this
}

// GetAutoIncludeIndexTables returns the AutoIncludeIndexTables field value if set, zero value otherwise.
func (o *DrConfigSetTablesForm) GetAutoIncludeIndexTables() bool {
	if o == nil || o.AutoIncludeIndexTables == nil {
		var ret bool
		return ret
	}
	return *o.AutoIncludeIndexTables
}

// GetAutoIncludeIndexTablesOk returns a tuple with the AutoIncludeIndexTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigSetTablesForm) GetAutoIncludeIndexTablesOk() (*bool, bool) {
	if o == nil || o.AutoIncludeIndexTables == nil {
		return nil, false
	}
	return o.AutoIncludeIndexTables, true
}

// HasAutoIncludeIndexTables returns a boolean if a field has been set.
func (o *DrConfigSetTablesForm) HasAutoIncludeIndexTables() bool {
	if o != nil && o.AutoIncludeIndexTables != nil {
		return true
	}

	return false
}

// SetAutoIncludeIndexTables gets a reference to the given bool and assigns it to the AutoIncludeIndexTables field.
func (o *DrConfigSetTablesForm) SetAutoIncludeIndexTables(v bool) {
	o.AutoIncludeIndexTables = &v
}

// GetBootstrapParams returns the BootstrapParams field value if set, zero value otherwise.
func (o *DrConfigSetTablesForm) GetBootstrapParams() RestartBootstrapParams {
	if o == nil || o.BootstrapParams == nil {
		var ret RestartBootstrapParams
		return ret
	}
	return *o.BootstrapParams
}

// GetBootstrapParamsOk returns a tuple with the BootstrapParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigSetTablesForm) GetBootstrapParamsOk() (*RestartBootstrapParams, bool) {
	if o == nil || o.BootstrapParams == nil {
		return nil, false
	}
	return o.BootstrapParams, true
}

// HasBootstrapParams returns a boolean if a field has been set.
func (o *DrConfigSetTablesForm) HasBootstrapParams() bool {
	if o != nil && o.BootstrapParams != nil {
		return true
	}

	return false
}

// SetBootstrapParams gets a reference to the given RestartBootstrapParams and assigns it to the BootstrapParams field.
func (o *DrConfigSetTablesForm) SetBootstrapParams(v RestartBootstrapParams) {
	o.BootstrapParams = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *DrConfigSetTablesForm) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return *o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigSetTablesForm) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *DrConfigSetTablesForm) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *DrConfigSetTablesForm) SetTables(v []string) {
	o.Tables = &v
}

func (o DrConfigSetTablesForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoIncludeIndexTables != nil {
		toSerialize["autoIncludeIndexTables"] = o.AutoIncludeIndexTables
	}
	if o.BootstrapParams != nil {
		toSerialize["bootstrapParams"] = o.BootstrapParams
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	return json.Marshal(toSerialize)
}

type NullableDrConfigSetTablesForm struct {
	value *DrConfigSetTablesForm
	isSet bool
}

func (v NullableDrConfigSetTablesForm) Get() *DrConfigSetTablesForm {
	return v.value
}

func (v *NullableDrConfigSetTablesForm) Set(val *DrConfigSetTablesForm) {
	v.value = val
	v.isSet = true
}

func (v NullableDrConfigSetTablesForm) IsSet() bool {
	return v.isSet
}

func (v *NullableDrConfigSetTablesForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrConfigSetTablesForm(val *DrConfigSetTablesForm) *NullableDrConfigSetTablesForm {
	return &NullableDrConfigSetTablesForm{value: val, isSet: true}
}

func (v NullableDrConfigSetTablesForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrConfigSetTablesForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


