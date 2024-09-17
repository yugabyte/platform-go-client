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

// DrConfigEditForm drConfig edit form
type DrConfigEditForm struct {
	BootstrapParams *RestartBootstrapParams `json:"bootstrapParams,omitempty"`
	PitrParams *PitrParams `json:"pitrParams,omitempty"`
}

// NewDrConfigEditForm instantiates a new DrConfigEditForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrConfigEditForm() *DrConfigEditForm {
	this := DrConfigEditForm{}
	return &this
}

// NewDrConfigEditFormWithDefaults instantiates a new DrConfigEditForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrConfigEditFormWithDefaults() *DrConfigEditForm {
	this := DrConfigEditForm{}
	return &this
}

// GetBootstrapParams returns the BootstrapParams field value if set, zero value otherwise.
func (o *DrConfigEditForm) GetBootstrapParams() RestartBootstrapParams {
	if o == nil || o.BootstrapParams == nil {
		var ret RestartBootstrapParams
		return ret
	}
	return *o.BootstrapParams
}

// GetBootstrapParamsOk returns a tuple with the BootstrapParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigEditForm) GetBootstrapParamsOk() (*RestartBootstrapParams, bool) {
	if o == nil || o.BootstrapParams == nil {
		return nil, false
	}
	return o.BootstrapParams, true
}

// HasBootstrapParams returns a boolean if a field has been set.
func (o *DrConfigEditForm) HasBootstrapParams() bool {
	if o != nil && o.BootstrapParams != nil {
		return true
	}

	return false
}

// SetBootstrapParams gets a reference to the given RestartBootstrapParams and assigns it to the BootstrapParams field.
func (o *DrConfigEditForm) SetBootstrapParams(v RestartBootstrapParams) {
	o.BootstrapParams = &v
}

// GetPitrParams returns the PitrParams field value if set, zero value otherwise.
func (o *DrConfigEditForm) GetPitrParams() PitrParams {
	if o == nil || o.PitrParams == nil {
		var ret PitrParams
		return ret
	}
	return *o.PitrParams
}

// GetPitrParamsOk returns a tuple with the PitrParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigEditForm) GetPitrParamsOk() (*PitrParams, bool) {
	if o == nil || o.PitrParams == nil {
		return nil, false
	}
	return o.PitrParams, true
}

// HasPitrParams returns a boolean if a field has been set.
func (o *DrConfigEditForm) HasPitrParams() bool {
	if o != nil && o.PitrParams != nil {
		return true
	}

	return false
}

// SetPitrParams gets a reference to the given PitrParams and assigns it to the PitrParams field.
func (o *DrConfigEditForm) SetPitrParams(v PitrParams) {
	o.PitrParams = &v
}

func (o DrConfigEditForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootstrapParams != nil {
		toSerialize["bootstrapParams"] = o.BootstrapParams
	}
	if o.PitrParams != nil {
		toSerialize["pitrParams"] = o.PitrParams
	}
	return json.Marshal(toSerialize)
}

type NullableDrConfigEditForm struct {
	value *DrConfigEditForm
	isSet bool
}

func (v NullableDrConfigEditForm) Get() *DrConfigEditForm {
	return v.value
}

func (v *NullableDrConfigEditForm) Set(val *DrConfigEditForm) {
	v.value = val
	v.isSet = true
}

func (v NullableDrConfigEditForm) IsSet() bool {
	return v.isSet
}

func (v *NullableDrConfigEditForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrConfigEditForm(val *DrConfigEditForm) *NullableDrConfigEditForm {
	return &NullableDrConfigEditForm{value: val, isSet: true}
}

func (v NullableDrConfigEditForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrConfigEditForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


