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

// DeleteBackupInfo struct for DeleteBackupInfo
type DeleteBackupInfo struct {
	// backup UUID
	BackupUUID string `json:"backupUUID"`
	// storage config UUID
	StorageConfigUUID *string `json:"storageConfigUUID,omitempty"`
}

// NewDeleteBackupInfo instantiates a new DeleteBackupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBackupInfo(backupUUID string, ) *DeleteBackupInfo {
	this := DeleteBackupInfo{}
	this.BackupUUID = backupUUID
	return &this
}

// NewDeleteBackupInfoWithDefaults instantiates a new DeleteBackupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBackupInfoWithDefaults() *DeleteBackupInfo {
	this := DeleteBackupInfo{}
	return &this
}

// GetBackupUUID returns the BackupUUID field value
func (o *DeleteBackupInfo) GetBackupUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BackupUUID
}

// GetBackupUUIDOk returns a tuple with the BackupUUID field value
// and a boolean to check if the value has been set.
func (o *DeleteBackupInfo) GetBackupUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupUUID, true
}

// SetBackupUUID sets field value
func (o *DeleteBackupInfo) SetBackupUUID(v string) {
	o.BackupUUID = v
}

// GetStorageConfigUUID returns the StorageConfigUUID field value if set, zero value otherwise.
func (o *DeleteBackupInfo) GetStorageConfigUUID() string {
	if o == nil || o.StorageConfigUUID == nil {
		var ret string
		return ret
	}
	return *o.StorageConfigUUID
}

// GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBackupInfo) GetStorageConfigUUIDOk() (*string, bool) {
	if o == nil || o.StorageConfigUUID == nil {
		return nil, false
	}
	return o.StorageConfigUUID, true
}

// HasStorageConfigUUID returns a boolean if a field has been set.
func (o *DeleteBackupInfo) HasStorageConfigUUID() bool {
	if o != nil && o.StorageConfigUUID != nil {
		return true
	}

	return false
}

// SetStorageConfigUUID gets a reference to the given string and assigns it to the StorageConfigUUID field.
func (o *DeleteBackupInfo) SetStorageConfigUUID(v string) {
	o.StorageConfigUUID = &v
}

func (o DeleteBackupInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupUUID"] = o.BackupUUID
	}
	if o.StorageConfigUUID != nil {
		toSerialize["storageConfigUUID"] = o.StorageConfigUUID
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteBackupInfo struct {
	value *DeleteBackupInfo
	isSet bool
}

func (v NullableDeleteBackupInfo) Get() *DeleteBackupInfo {
	return v.value
}

func (v *NullableDeleteBackupInfo) Set(val *DeleteBackupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBackupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBackupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBackupInfo(val *DeleteBackupInfo) *NullableDeleteBackupInfo {
	return &NullableDeleteBackupInfo{value: val, isSet: true}
}

func (v NullableDeleteBackupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBackupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

