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

// ParametersForAdvancedRestorePreflightChecks struct for ParametersForAdvancedRestorePreflightChecks
type ParametersForAdvancedRestorePreflightChecks struct {
	// List of backup locations to restore from
	BackupLocations []string `json:"backupLocations"`
	// Point in restore timestamp in millis
	RestoreToPointInTimeMillis *int64 `json:"restoreToPointInTimeMillis,omitempty"`
	// Storage config UUID
	StorageConfigUUID string `json:"storageConfigUUID"`
	// Target universe UUID
	UniverseUUID string `json:"universeUUID"`
}

// NewParametersForAdvancedRestorePreflightChecks instantiates a new ParametersForAdvancedRestorePreflightChecks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParametersForAdvancedRestorePreflightChecks(backupLocations []string, storageConfigUUID string, universeUUID string) *ParametersForAdvancedRestorePreflightChecks {
	this := ParametersForAdvancedRestorePreflightChecks{}
	this.BackupLocations = backupLocations
	this.StorageConfigUUID = storageConfigUUID
	this.UniverseUUID = universeUUID
	return &this
}

// NewParametersForAdvancedRestorePreflightChecksWithDefaults instantiates a new ParametersForAdvancedRestorePreflightChecks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersForAdvancedRestorePreflightChecksWithDefaults() *ParametersForAdvancedRestorePreflightChecks {
	this := ParametersForAdvancedRestorePreflightChecks{}
	return &this
}

// GetBackupLocations returns the BackupLocations field value
func (o *ParametersForAdvancedRestorePreflightChecks) GetBackupLocations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BackupLocations
}

// GetBackupLocationsOk returns a tuple with the BackupLocations field value
// and a boolean to check if the value has been set.
func (o *ParametersForAdvancedRestorePreflightChecks) GetBackupLocationsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupLocations, true
}

// SetBackupLocations sets field value
func (o *ParametersForAdvancedRestorePreflightChecks) SetBackupLocations(v []string) {
	o.BackupLocations = v
}

// GetRestoreToPointInTimeMillis returns the RestoreToPointInTimeMillis field value if set, zero value otherwise.
func (o *ParametersForAdvancedRestorePreflightChecks) GetRestoreToPointInTimeMillis() int64 {
	if o == nil || o.RestoreToPointInTimeMillis == nil {
		var ret int64
		return ret
	}
	return *o.RestoreToPointInTimeMillis
}

// GetRestoreToPointInTimeMillisOk returns a tuple with the RestoreToPointInTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParametersForAdvancedRestorePreflightChecks) GetRestoreToPointInTimeMillisOk() (*int64, bool) {
	if o == nil || o.RestoreToPointInTimeMillis == nil {
		return nil, false
	}
	return o.RestoreToPointInTimeMillis, true
}

// HasRestoreToPointInTimeMillis returns a boolean if a field has been set.
func (o *ParametersForAdvancedRestorePreflightChecks) HasRestoreToPointInTimeMillis() bool {
	if o != nil && o.RestoreToPointInTimeMillis != nil {
		return true
	}

	return false
}

// SetRestoreToPointInTimeMillis gets a reference to the given int64 and assigns it to the RestoreToPointInTimeMillis field.
func (o *ParametersForAdvancedRestorePreflightChecks) SetRestoreToPointInTimeMillis(v int64) {
	o.RestoreToPointInTimeMillis = &v
}

// GetStorageConfigUUID returns the StorageConfigUUID field value
func (o *ParametersForAdvancedRestorePreflightChecks) GetStorageConfigUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageConfigUUID
}

// GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field value
// and a boolean to check if the value has been set.
func (o *ParametersForAdvancedRestorePreflightChecks) GetStorageConfigUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageConfigUUID, true
}

// SetStorageConfigUUID sets field value
func (o *ParametersForAdvancedRestorePreflightChecks) SetStorageConfigUUID(v string) {
	o.StorageConfigUUID = v
}

// GetUniverseUUID returns the UniverseUUID field value
func (o *ParametersForAdvancedRestorePreflightChecks) GetUniverseUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniverseUUID
}

// GetUniverseUUIDOk returns a tuple with the UniverseUUID field value
// and a boolean to check if the value has been set.
func (o *ParametersForAdvancedRestorePreflightChecks) GetUniverseUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseUUID, true
}

// SetUniverseUUID sets field value
func (o *ParametersForAdvancedRestorePreflightChecks) SetUniverseUUID(v string) {
	o.UniverseUUID = v
}

func (o ParametersForAdvancedRestorePreflightChecks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupLocations"] = o.BackupLocations
	}
	if o.RestoreToPointInTimeMillis != nil {
		toSerialize["restoreToPointInTimeMillis"] = o.RestoreToPointInTimeMillis
	}
	if true {
		toSerialize["storageConfigUUID"] = o.StorageConfigUUID
	}
	if true {
		toSerialize["universeUUID"] = o.UniverseUUID
	}
	return json.Marshal(toSerialize)
}

type NullableParametersForAdvancedRestorePreflightChecks struct {
	value *ParametersForAdvancedRestorePreflightChecks
	isSet bool
}

func (v NullableParametersForAdvancedRestorePreflightChecks) Get() *ParametersForAdvancedRestorePreflightChecks {
	return v.value
}

func (v *NullableParametersForAdvancedRestorePreflightChecks) Set(val *ParametersForAdvancedRestorePreflightChecks) {
	v.value = val
	v.isSet = true
}

func (v NullableParametersForAdvancedRestorePreflightChecks) IsSet() bool {
	return v.isSet
}

func (v *NullableParametersForAdvancedRestorePreflightChecks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParametersForAdvancedRestorePreflightChecks(val *ParametersForAdvancedRestorePreflightChecks) *NullableParametersForAdvancedRestorePreflightChecks {
	return &NullableParametersForAdvancedRestorePreflightChecks{value: val, isSet: true}
}

func (v NullableParametersForAdvancedRestorePreflightChecks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParametersForAdvancedRestorePreflightChecks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

