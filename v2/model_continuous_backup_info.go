/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"time"
)

// ContinuousBackupInfo ContinuousBackupInfo  These are read-only system generated properties of a continuous backup configuration. Returned as part of the Continuous Backup resource. 
type ContinuousBackupInfo struct {
	// UUID of the Continuous Backup Config
	Uuid *string `json:"uuid,omitempty"`
	// bucket or directory where backups are stored
	StorageLocation *string `json:"storage_location,omitempty"`
	// time of last backup stored
	LastBackup *time.Time `json:"last_backup,omitempty"`
}

// NewContinuousBackupInfo instantiates a new ContinuousBackupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousBackupInfo() *ContinuousBackupInfo {
	this := ContinuousBackupInfo{}
	return &this
}

// NewContinuousBackupInfoWithDefaults instantiates a new ContinuousBackupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousBackupInfoWithDefaults() *ContinuousBackupInfo {
	this := ContinuousBackupInfo{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ContinuousBackupInfo) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousBackupInfo) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ContinuousBackupInfo) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ContinuousBackupInfo) SetUuid(v string) {
	o.Uuid = &v
}

// GetStorageLocation returns the StorageLocation field value if set, zero value otherwise.
func (o *ContinuousBackupInfo) GetStorageLocation() string {
	if o == nil || o.StorageLocation == nil {
		var ret string
		return ret
	}
	return *o.StorageLocation
}

// GetStorageLocationOk returns a tuple with the StorageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousBackupInfo) GetStorageLocationOk() (*string, bool) {
	if o == nil || o.StorageLocation == nil {
		return nil, false
	}
	return o.StorageLocation, true
}

// HasStorageLocation returns a boolean if a field has been set.
func (o *ContinuousBackupInfo) HasStorageLocation() bool {
	if o != nil && o.StorageLocation != nil {
		return true
	}

	return false
}

// SetStorageLocation gets a reference to the given string and assigns it to the StorageLocation field.
func (o *ContinuousBackupInfo) SetStorageLocation(v string) {
	o.StorageLocation = &v
}

// GetLastBackup returns the LastBackup field value if set, zero value otherwise.
func (o *ContinuousBackupInfo) GetLastBackup() time.Time {
	if o == nil || o.LastBackup == nil {
		var ret time.Time
		return ret
	}
	return *o.LastBackup
}

// GetLastBackupOk returns a tuple with the LastBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousBackupInfo) GetLastBackupOk() (*time.Time, bool) {
	if o == nil || o.LastBackup == nil {
		return nil, false
	}
	return o.LastBackup, true
}

// HasLastBackup returns a boolean if a field has been set.
func (o *ContinuousBackupInfo) HasLastBackup() bool {
	if o != nil && o.LastBackup != nil {
		return true
	}

	return false
}

// SetLastBackup gets a reference to the given time.Time and assigns it to the LastBackup field.
func (o *ContinuousBackupInfo) SetLastBackup(v time.Time) {
	o.LastBackup = &v
}

func (o ContinuousBackupInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.StorageLocation != nil {
		toSerialize["storage_location"] = o.StorageLocation
	}
	if o.LastBackup != nil {
		toSerialize["last_backup"] = o.LastBackup
	}
	return json.Marshal(toSerialize)
}

type NullableContinuousBackupInfo struct {
	value *ContinuousBackupInfo
	isSet bool
}

func (v NullableContinuousBackupInfo) Get() *ContinuousBackupInfo {
	return v.value
}

func (v *NullableContinuousBackupInfo) Set(val *ContinuousBackupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousBackupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousBackupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousBackupInfo(val *ContinuousBackupInfo) *NullableContinuousBackupInfo {
	return &NullableContinuousBackupInfo{value: val, isSet: true}
}

func (v NullableContinuousBackupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousBackupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


