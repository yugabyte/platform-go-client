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

// DrConfigRestartForm dr config restart form
type DrConfigRestartForm struct {
	BootstrapParams *RestartBootstrapParams `json:"bootstrapParams,omitempty"`
	// Primary Universe DB IDs
	Dbs []string `json:"dbs"`
}

// NewDrConfigRestartForm instantiates a new DrConfigRestartForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrConfigRestartForm(dbs []string) *DrConfigRestartForm {
	this := DrConfigRestartForm{}
	this.Dbs = dbs
	return &this
}

// NewDrConfigRestartFormWithDefaults instantiates a new DrConfigRestartForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrConfigRestartFormWithDefaults() *DrConfigRestartForm {
	this := DrConfigRestartForm{}
	return &this
}

// GetBootstrapParams returns the BootstrapParams field value if set, zero value otherwise.
func (o *DrConfigRestartForm) GetBootstrapParams() RestartBootstrapParams {
	if o == nil || o.BootstrapParams == nil {
		var ret RestartBootstrapParams
		return ret
	}
	return *o.BootstrapParams
}

// GetBootstrapParamsOk returns a tuple with the BootstrapParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigRestartForm) GetBootstrapParamsOk() (*RestartBootstrapParams, bool) {
	if o == nil || o.BootstrapParams == nil {
		return nil, false
	}
	return o.BootstrapParams, true
}

// HasBootstrapParams returns a boolean if a field has been set.
func (o *DrConfigRestartForm) HasBootstrapParams() bool {
	if o != nil && o.BootstrapParams != nil {
		return true
	}

	return false
}

// SetBootstrapParams gets a reference to the given RestartBootstrapParams and assigns it to the BootstrapParams field.
func (o *DrConfigRestartForm) SetBootstrapParams(v RestartBootstrapParams) {
	o.BootstrapParams = &v
}

// GetDbs returns the Dbs field value
func (o *DrConfigRestartForm) GetDbs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dbs
}

// GetDbsOk returns a tuple with the Dbs field value
// and a boolean to check if the value has been set.
func (o *DrConfigRestartForm) GetDbsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dbs, true
}

// SetDbs sets field value
func (o *DrConfigRestartForm) SetDbs(v []string) {
	o.Dbs = v
}

func (o DrConfigRestartForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootstrapParams != nil {
		toSerialize["bootstrapParams"] = o.BootstrapParams
	}
	if true {
		toSerialize["dbs"] = o.Dbs
	}
	return json.Marshal(toSerialize)
}

type NullableDrConfigRestartForm struct {
	value *DrConfigRestartForm
	isSet bool
}

func (v NullableDrConfigRestartForm) Get() *DrConfigRestartForm {
	return v.value
}

func (v *NullableDrConfigRestartForm) Set(val *DrConfigRestartForm) {
	v.value = val
	v.isSet = true
}

func (v NullableDrConfigRestartForm) IsSet() bool {
	return v.isSet
}

func (v *NullableDrConfigRestartForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrConfigRestartForm(val *DrConfigRestartForm) *NullableDrConfigRestartForm {
	return &NullableDrConfigRestartForm{value: val, isSet: true}
}

func (v NullableDrConfigRestartForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrConfigRestartForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

