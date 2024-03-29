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

// RestorePreflightResponse struct for RestorePreflightResponse
type RestorePreflightResponse struct {
	// Backup Category
	BackupCategory *string `json:"backupCategory,omitempty"`
	// Whether backup was KMS encrypted
	HasKMSHistory *bool `json:"hasKMSHistory,omitempty"`
	// Map of backup location and backup-info object
	PerLocationBackupInfoMap *map[string]PerLocationBackupInfo `json:"perLocationBackupInfoMap,omitempty"`
}

// NewRestorePreflightResponse instantiates a new RestorePreflightResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestorePreflightResponse() *RestorePreflightResponse {
	this := RestorePreflightResponse{}
	return &this
}

// NewRestorePreflightResponseWithDefaults instantiates a new RestorePreflightResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestorePreflightResponseWithDefaults() *RestorePreflightResponse {
	this := RestorePreflightResponse{}
	return &this
}

// GetBackupCategory returns the BackupCategory field value if set, zero value otherwise.
func (o *RestorePreflightResponse) GetBackupCategory() string {
	if o == nil || o.BackupCategory == nil {
		var ret string
		return ret
	}
	return *o.BackupCategory
}

// GetBackupCategoryOk returns a tuple with the BackupCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightResponse) GetBackupCategoryOk() (*string, bool) {
	if o == nil || o.BackupCategory == nil {
		return nil, false
	}
	return o.BackupCategory, true
}

// HasBackupCategory returns a boolean if a field has been set.
func (o *RestorePreflightResponse) HasBackupCategory() bool {
	if o != nil && o.BackupCategory != nil {
		return true
	}

	return false
}

// SetBackupCategory gets a reference to the given string and assigns it to the BackupCategory field.
func (o *RestorePreflightResponse) SetBackupCategory(v string) {
	o.BackupCategory = &v
}

// GetHasKMSHistory returns the HasKMSHistory field value if set, zero value otherwise.
func (o *RestorePreflightResponse) GetHasKMSHistory() bool {
	if o == nil || o.HasKMSHistory == nil {
		var ret bool
		return ret
	}
	return *o.HasKMSHistory
}

// GetHasKMSHistoryOk returns a tuple with the HasKMSHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightResponse) GetHasKMSHistoryOk() (*bool, bool) {
	if o == nil || o.HasKMSHistory == nil {
		return nil, false
	}
	return o.HasKMSHistory, true
}

// HasHasKMSHistory returns a boolean if a field has been set.
func (o *RestorePreflightResponse) HasHasKMSHistory() bool {
	if o != nil && o.HasKMSHistory != nil {
		return true
	}

	return false
}

// SetHasKMSHistory gets a reference to the given bool and assigns it to the HasKMSHistory field.
func (o *RestorePreflightResponse) SetHasKMSHistory(v bool) {
	o.HasKMSHistory = &v
}

// GetPerLocationBackupInfoMap returns the PerLocationBackupInfoMap field value if set, zero value otherwise.
func (o *RestorePreflightResponse) GetPerLocationBackupInfoMap() map[string]PerLocationBackupInfo {
	if o == nil || o.PerLocationBackupInfoMap == nil {
		var ret map[string]PerLocationBackupInfo
		return ret
	}
	return *o.PerLocationBackupInfoMap
}

// GetPerLocationBackupInfoMapOk returns a tuple with the PerLocationBackupInfoMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightResponse) GetPerLocationBackupInfoMapOk() (*map[string]PerLocationBackupInfo, bool) {
	if o == nil || o.PerLocationBackupInfoMap == nil {
		return nil, false
	}
	return o.PerLocationBackupInfoMap, true
}

// HasPerLocationBackupInfoMap returns a boolean if a field has been set.
func (o *RestorePreflightResponse) HasPerLocationBackupInfoMap() bool {
	if o != nil && o.PerLocationBackupInfoMap != nil {
		return true
	}

	return false
}

// SetPerLocationBackupInfoMap gets a reference to the given map[string]PerLocationBackupInfo and assigns it to the PerLocationBackupInfoMap field.
func (o *RestorePreflightResponse) SetPerLocationBackupInfoMap(v map[string]PerLocationBackupInfo) {
	o.PerLocationBackupInfoMap = &v
}

func (o RestorePreflightResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupCategory != nil {
		toSerialize["backupCategory"] = o.BackupCategory
	}
	if o.HasKMSHistory != nil {
		toSerialize["hasKMSHistory"] = o.HasKMSHistory
	}
	if o.PerLocationBackupInfoMap != nil {
		toSerialize["perLocationBackupInfoMap"] = o.PerLocationBackupInfoMap
	}
	return json.Marshal(toSerialize)
}

type NullableRestorePreflightResponse struct {
	value *RestorePreflightResponse
	isSet bool
}

func (v NullableRestorePreflightResponse) Get() *RestorePreflightResponse {
	return v.value
}

func (v *NullableRestorePreflightResponse) Set(val *RestorePreflightResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRestorePreflightResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRestorePreflightResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestorePreflightResponse(val *RestorePreflightResponse) *NullableRestorePreflightResponse {
	return &NullableRestorePreflightResponse{value: val, isSet: true}
}

func (v NullableRestorePreflightResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestorePreflightResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


