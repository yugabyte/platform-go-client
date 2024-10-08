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

// BootstrapBackupParams Backup parameters for bootstrapping
type BootstrapBackupParams struct {
	// Number of concurrent commands used by yb_backup (not ybc) to run on nodes over SSH
	Parallelism *int32 `json:"parallelism,omitempty"`
	// Storage configuration UUID
	StorageConfigUUID string `json:"storageConfigUUID"`
}

// NewBootstrapBackupParams instantiates a new BootstrapBackupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootstrapBackupParams(storageConfigUUID string) *BootstrapBackupParams {
	this := BootstrapBackupParams{}
	this.StorageConfigUUID = storageConfigUUID
	return &this
}

// NewBootstrapBackupParamsWithDefaults instantiates a new BootstrapBackupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootstrapBackupParamsWithDefaults() *BootstrapBackupParams {
	this := BootstrapBackupParams{}
	return &this
}

// GetParallelism returns the Parallelism field value if set, zero value otherwise.
func (o *BootstrapBackupParams) GetParallelism() int32 {
	if o == nil || o.Parallelism == nil {
		var ret int32
		return ret
	}
	return *o.Parallelism
}

// GetParallelismOk returns a tuple with the Parallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrapBackupParams) GetParallelismOk() (*int32, bool) {
	if o == nil || o.Parallelism == nil {
		return nil, false
	}
	return o.Parallelism, true
}

// HasParallelism returns a boolean if a field has been set.
func (o *BootstrapBackupParams) HasParallelism() bool {
	if o != nil && o.Parallelism != nil {
		return true
	}

	return false
}

// SetParallelism gets a reference to the given int32 and assigns it to the Parallelism field.
func (o *BootstrapBackupParams) SetParallelism(v int32) {
	o.Parallelism = &v
}

// GetStorageConfigUUID returns the StorageConfigUUID field value
func (o *BootstrapBackupParams) GetStorageConfigUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageConfigUUID
}

// GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field value
// and a boolean to check if the value has been set.
func (o *BootstrapBackupParams) GetStorageConfigUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageConfigUUID, true
}

// SetStorageConfigUUID sets field value
func (o *BootstrapBackupParams) SetStorageConfigUUID(v string) {
	o.StorageConfigUUID = v
}

func (o BootstrapBackupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parallelism != nil {
		toSerialize["parallelism"] = o.Parallelism
	}
	if true {
		toSerialize["storageConfigUUID"] = o.StorageConfigUUID
	}
	return json.Marshal(toSerialize)
}

type NullableBootstrapBackupParams struct {
	value *BootstrapBackupParams
	isSet bool
}

func (v NullableBootstrapBackupParams) Get() *BootstrapBackupParams {
	return v.value
}

func (v *NullableBootstrapBackupParams) Set(val *BootstrapBackupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBootstrapBackupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBootstrapBackupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootstrapBackupParams(val *BootstrapBackupParams) *NullableBootstrapBackupParams {
	return &NullableBootstrapBackupParams{value: val, isSet: true}
}

func (v NullableBootstrapBackupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootstrapBackupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


