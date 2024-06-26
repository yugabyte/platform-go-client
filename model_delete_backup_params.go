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

// DeleteBackupParams struct for DeleteBackupParams
type DeleteBackupParams struct {
	// Backups to be deleted
	DeleteBackupInfos []DeleteBackupInfo `json:"deleteBackupInfos"`
	// Delete Backups forcefully
	DeleteForcefully *bool `json:"deleteForcefully,omitempty"`
}

// NewDeleteBackupParams instantiates a new DeleteBackupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBackupParams(deleteBackupInfos []DeleteBackupInfo) *DeleteBackupParams {
	this := DeleteBackupParams{}
	this.DeleteBackupInfos = deleteBackupInfos
	return &this
}

// NewDeleteBackupParamsWithDefaults instantiates a new DeleteBackupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBackupParamsWithDefaults() *DeleteBackupParams {
	this := DeleteBackupParams{}
	return &this
}

// GetDeleteBackupInfos returns the DeleteBackupInfos field value
func (o *DeleteBackupParams) GetDeleteBackupInfos() []DeleteBackupInfo {
	if o == nil {
		var ret []DeleteBackupInfo
		return ret
	}

	return o.DeleteBackupInfos
}

// GetDeleteBackupInfosOk returns a tuple with the DeleteBackupInfos field value
// and a boolean to check if the value has been set.
func (o *DeleteBackupParams) GetDeleteBackupInfosOk() (*[]DeleteBackupInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeleteBackupInfos, true
}

// SetDeleteBackupInfos sets field value
func (o *DeleteBackupParams) SetDeleteBackupInfos(v []DeleteBackupInfo) {
	o.DeleteBackupInfos = v
}

// GetDeleteForcefully returns the DeleteForcefully field value if set, zero value otherwise.
func (o *DeleteBackupParams) GetDeleteForcefully() bool {
	if o == nil || o.DeleteForcefully == nil {
		var ret bool
		return ret
	}
	return *o.DeleteForcefully
}

// GetDeleteForcefullyOk returns a tuple with the DeleteForcefully field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBackupParams) GetDeleteForcefullyOk() (*bool, bool) {
	if o == nil || o.DeleteForcefully == nil {
		return nil, false
	}
	return o.DeleteForcefully, true
}

// HasDeleteForcefully returns a boolean if a field has been set.
func (o *DeleteBackupParams) HasDeleteForcefully() bool {
	if o != nil && o.DeleteForcefully != nil {
		return true
	}

	return false
}

// SetDeleteForcefully gets a reference to the given bool and assigns it to the DeleteForcefully field.
func (o *DeleteBackupParams) SetDeleteForcefully(v bool) {
	o.DeleteForcefully = &v
}

func (o DeleteBackupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deleteBackupInfos"] = o.DeleteBackupInfos
	}
	if o.DeleteForcefully != nil {
		toSerialize["deleteForcefully"] = o.DeleteForcefully
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteBackupParams struct {
	value *DeleteBackupParams
	isSet bool
}

func (v NullableDeleteBackupParams) Get() *DeleteBackupParams {
	return v.value
}

func (v *NullableDeleteBackupParams) Set(val *DeleteBackupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBackupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBackupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBackupParams(val *DeleteBackupParams) *NullableDeleteBackupParams {
	return &NullableDeleteBackupParams{value: val, isSet: true}
}

func (v NullableDeleteBackupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBackupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


