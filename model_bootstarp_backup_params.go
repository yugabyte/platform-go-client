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

// BootstarpBackupParams Backup parameters for bootstrapping
type BootstarpBackupParams struct {
	// Number of concurrent commands used by yb_backup (not ybc) to run on nodes over SSH
	Parallelism *int32 `json:"parallelism,omitempty"`
	// Storage configuration UUID
	StorageConfigUUID string `json:"storageConfigUUID"`
}

// NewBootstarpBackupParams instantiates a new BootstarpBackupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootstarpBackupParams(storageConfigUUID string) *BootstarpBackupParams {
	this := BootstarpBackupParams{}
	this.StorageConfigUUID = storageConfigUUID
	return &this
}

// NewBootstarpBackupParamsWithDefaults instantiates a new BootstarpBackupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootstarpBackupParamsWithDefaults() *BootstarpBackupParams {
	this := BootstarpBackupParams{}
	return &this
}

// GetParallelism returns the Parallelism field value if set, zero value otherwise.
func (o *BootstarpBackupParams) GetParallelism() int32 {
	if o == nil || o.Parallelism == nil {
		var ret int32
		return ret
	}
	return *o.Parallelism
}

// GetParallelismOk returns a tuple with the Parallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstarpBackupParams) GetParallelismOk() (*int32, bool) {
	if o == nil || o.Parallelism == nil {
		return nil, false
	}
	return o.Parallelism, true
}

// HasParallelism returns a boolean if a field has been set.
func (o *BootstarpBackupParams) HasParallelism() bool {
	if o != nil && o.Parallelism != nil {
		return true
	}

	return false
}

// SetParallelism gets a reference to the given int32 and assigns it to the Parallelism field.
func (o *BootstarpBackupParams) SetParallelism(v int32) {
	o.Parallelism = &v
}

// GetStorageConfigUUID returns the StorageConfigUUID field value
func (o *BootstarpBackupParams) GetStorageConfigUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageConfigUUID
}

// GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field value
// and a boolean to check if the value has been set.
func (o *BootstarpBackupParams) GetStorageConfigUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageConfigUUID, true
}

// SetStorageConfigUUID sets field value
func (o *BootstarpBackupParams) SetStorageConfigUUID(v string) {
	o.StorageConfigUUID = v
}

func (o BootstarpBackupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parallelism != nil {
		toSerialize["parallelism"] = o.Parallelism
	}
	if true {
		toSerialize["storageConfigUUID"] = o.StorageConfigUUID
	}
	return json.Marshal(toSerialize)
}

type NullableBootstarpBackupParams struct {
	value *BootstarpBackupParams
	isSet bool
}

func (v NullableBootstarpBackupParams) Get() *BootstarpBackupParams {
	return v.value
}

func (v *NullableBootstarpBackupParams) Set(val *BootstarpBackupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBootstarpBackupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBootstarpBackupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootstarpBackupParams(val *BootstarpBackupParams) *NullableBootstarpBackupParams {
	return &NullableBootstarpBackupParams{value: val, isSet: true}
}

func (v NullableBootstarpBackupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootstarpBackupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


