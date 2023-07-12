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
	"time"
)

// Backup A single backup. Includes the backup's status, expiration time, and configuration.
type Backup struct {
	BackupInfo *BackupTableParams `json:"backupInfo,omitempty"`
	// Backup UUID
	BackupUUID *string `json:"backupUUID,omitempty"`
	// Base backup UUID
	BaseBackupUUID *string `json:"baseBackupUUID,omitempty"`
	// Category of the backup
	Category *string `json:"category,omitempty"`
	// Backup completion time
	CompletionTime *time.Time `json:"completionTime,omitempty"`
	// Backup creation time
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Customer UUID that owns this backup
	CustomerUUID *string `json:"customerUUID,omitempty"`
	// Expiry time (unix timestamp) of the backup
	Expiry *time.Time `json:"expiry,omitempty"`
	// Time unit for backup expiry time
	ExpiryTimeUnit *string `json:"expiryTimeUnit,omitempty"`
	IncrementalBackup bool `json:"incrementalBackup"`
	ParentBackup bool `json:"parentBackup"`
	// Schedule Policy Name, if this backup is part of a schedule
	ScheduleName *string `json:"scheduleName,omitempty"`
	// Schedule UUID, if this backup is part of a schedule
	ScheduleUUID *string `json:"scheduleUUID,omitempty"`
	// State of the backup
	State *string `json:"state,omitempty"`
	// Storage Config UUID that created this backup
	StorageConfigUUID *string `json:"storageConfigUUID,omitempty"`
	// Backup UUID
	TaskUUID *string `json:"taskUUID,omitempty"`
	// Universe name that created this backup
	UniverseName *string `json:"universeName,omitempty"`
	// Backup update time
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Version of the backup in a category
	Version *string `json:"version,omitempty"`
}

// NewBackup instantiates a new Backup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackup(incrementalBackup bool, parentBackup bool, ) *Backup {
	this := Backup{}
	this.IncrementalBackup = incrementalBackup
	this.ParentBackup = parentBackup
	return &this
}

// NewBackupWithDefaults instantiates a new Backup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupWithDefaults() *Backup {
	this := Backup{}
	return &this
}

// GetBackupInfo returns the BackupInfo field value if set, zero value otherwise.
func (o *Backup) GetBackupInfo() BackupTableParams {
	if o == nil || o.BackupInfo == nil {
		var ret BackupTableParams
		return ret
	}
	return *o.BackupInfo
}

// GetBackupInfoOk returns a tuple with the BackupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetBackupInfoOk() (*BackupTableParams, bool) {
	if o == nil || o.BackupInfo == nil {
		return nil, false
	}
	return o.BackupInfo, true
}

// HasBackupInfo returns a boolean if a field has been set.
func (o *Backup) HasBackupInfo() bool {
	if o != nil && o.BackupInfo != nil {
		return true
	}

	return false
}

// SetBackupInfo gets a reference to the given BackupTableParams and assigns it to the BackupInfo field.
func (o *Backup) SetBackupInfo(v BackupTableParams) {
	o.BackupInfo = &v
}

// GetBackupUUID returns the BackupUUID field value if set, zero value otherwise.
func (o *Backup) GetBackupUUID() string {
	if o == nil || o.BackupUUID == nil {
		var ret string
		return ret
	}
	return *o.BackupUUID
}

// GetBackupUUIDOk returns a tuple with the BackupUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetBackupUUIDOk() (*string, bool) {
	if o == nil || o.BackupUUID == nil {
		return nil, false
	}
	return o.BackupUUID, true
}

// HasBackupUUID returns a boolean if a field has been set.
func (o *Backup) HasBackupUUID() bool {
	if o != nil && o.BackupUUID != nil {
		return true
	}

	return false
}

// SetBackupUUID gets a reference to the given string and assigns it to the BackupUUID field.
func (o *Backup) SetBackupUUID(v string) {
	o.BackupUUID = &v
}

// GetBaseBackupUUID returns the BaseBackupUUID field value if set, zero value otherwise.
func (o *Backup) GetBaseBackupUUID() string {
	if o == nil || o.BaseBackupUUID == nil {
		var ret string
		return ret
	}
	return *o.BaseBackupUUID
}

// GetBaseBackupUUIDOk returns a tuple with the BaseBackupUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetBaseBackupUUIDOk() (*string, bool) {
	if o == nil || o.BaseBackupUUID == nil {
		return nil, false
	}
	return o.BaseBackupUUID, true
}

// HasBaseBackupUUID returns a boolean if a field has been set.
func (o *Backup) HasBaseBackupUUID() bool {
	if o != nil && o.BaseBackupUUID != nil {
		return true
	}

	return false
}

// SetBaseBackupUUID gets a reference to the given string and assigns it to the BaseBackupUUID field.
func (o *Backup) SetBaseBackupUUID(v string) {
	o.BaseBackupUUID = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Backup) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Backup) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Backup) SetCategory(v string) {
	o.Category = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *Backup) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *Backup) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *Backup) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Backup) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Backup) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Backup) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetCustomerUUID returns the CustomerUUID field value if set, zero value otherwise.
func (o *Backup) GetCustomerUUID() string {
	if o == nil || o.CustomerUUID == nil {
		var ret string
		return ret
	}
	return *o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetCustomerUUIDOk() (*string, bool) {
	if o == nil || o.CustomerUUID == nil {
		return nil, false
	}
	return o.CustomerUUID, true
}

// HasCustomerUUID returns a boolean if a field has been set.
func (o *Backup) HasCustomerUUID() bool {
	if o != nil && o.CustomerUUID != nil {
		return true
	}

	return false
}

// SetCustomerUUID gets a reference to the given string and assigns it to the CustomerUUID field.
func (o *Backup) SetCustomerUUID(v string) {
	o.CustomerUUID = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *Backup) GetExpiry() time.Time {
	if o == nil || o.Expiry == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetExpiryOk() (*time.Time, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *Backup) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *Backup) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetExpiryTimeUnit returns the ExpiryTimeUnit field value if set, zero value otherwise.
func (o *Backup) GetExpiryTimeUnit() string {
	if o == nil || o.ExpiryTimeUnit == nil {
		var ret string
		return ret
	}
	return *o.ExpiryTimeUnit
}

// GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetExpiryTimeUnitOk() (*string, bool) {
	if o == nil || o.ExpiryTimeUnit == nil {
		return nil, false
	}
	return o.ExpiryTimeUnit, true
}

// HasExpiryTimeUnit returns a boolean if a field has been set.
func (o *Backup) HasExpiryTimeUnit() bool {
	if o != nil && o.ExpiryTimeUnit != nil {
		return true
	}

	return false
}

// SetExpiryTimeUnit gets a reference to the given string and assigns it to the ExpiryTimeUnit field.
func (o *Backup) SetExpiryTimeUnit(v string) {
	o.ExpiryTimeUnit = &v
}

// GetIncrementalBackup returns the IncrementalBackup field value
func (o *Backup) GetIncrementalBackup() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IncrementalBackup
}

// GetIncrementalBackupOk returns a tuple with the IncrementalBackup field value
// and a boolean to check if the value has been set.
func (o *Backup) GetIncrementalBackupOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncrementalBackup, true
}

// SetIncrementalBackup sets field value
func (o *Backup) SetIncrementalBackup(v bool) {
	o.IncrementalBackup = v
}

// GetParentBackup returns the ParentBackup field value
func (o *Backup) GetParentBackup() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.ParentBackup
}

// GetParentBackupOk returns a tuple with the ParentBackup field value
// and a boolean to check if the value has been set.
func (o *Backup) GetParentBackupOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParentBackup, true
}

// SetParentBackup sets field value
func (o *Backup) SetParentBackup(v bool) {
	o.ParentBackup = v
}

// GetScheduleName returns the ScheduleName field value if set, zero value otherwise.
func (o *Backup) GetScheduleName() string {
	if o == nil || o.ScheduleName == nil {
		var ret string
		return ret
	}
	return *o.ScheduleName
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetScheduleNameOk() (*string, bool) {
	if o == nil || o.ScheduleName == nil {
		return nil, false
	}
	return o.ScheduleName, true
}

// HasScheduleName returns a boolean if a field has been set.
func (o *Backup) HasScheduleName() bool {
	if o != nil && o.ScheduleName != nil {
		return true
	}

	return false
}

// SetScheduleName gets a reference to the given string and assigns it to the ScheduleName field.
func (o *Backup) SetScheduleName(v string) {
	o.ScheduleName = &v
}

// GetScheduleUUID returns the ScheduleUUID field value if set, zero value otherwise.
func (o *Backup) GetScheduleUUID() string {
	if o == nil || o.ScheduleUUID == nil {
		var ret string
		return ret
	}
	return *o.ScheduleUUID
}

// GetScheduleUUIDOk returns a tuple with the ScheduleUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetScheduleUUIDOk() (*string, bool) {
	if o == nil || o.ScheduleUUID == nil {
		return nil, false
	}
	return o.ScheduleUUID, true
}

// HasScheduleUUID returns a boolean if a field has been set.
func (o *Backup) HasScheduleUUID() bool {
	if o != nil && o.ScheduleUUID != nil {
		return true
	}

	return false
}

// SetScheduleUUID gets a reference to the given string and assigns it to the ScheduleUUID field.
func (o *Backup) SetScheduleUUID(v string) {
	o.ScheduleUUID = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Backup) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Backup) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Backup) SetState(v string) {
	o.State = &v
}

// GetStorageConfigUUID returns the StorageConfigUUID field value if set, zero value otherwise.
func (o *Backup) GetStorageConfigUUID() string {
	if o == nil || o.StorageConfigUUID == nil {
		var ret string
		return ret
	}
	return *o.StorageConfigUUID
}

// GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetStorageConfigUUIDOk() (*string, bool) {
	if o == nil || o.StorageConfigUUID == nil {
		return nil, false
	}
	return o.StorageConfigUUID, true
}

// HasStorageConfigUUID returns a boolean if a field has been set.
func (o *Backup) HasStorageConfigUUID() bool {
	if o != nil && o.StorageConfigUUID != nil {
		return true
	}

	return false
}

// SetStorageConfigUUID gets a reference to the given string and assigns it to the StorageConfigUUID field.
func (o *Backup) SetStorageConfigUUID(v string) {
	o.StorageConfigUUID = &v
}

// GetTaskUUID returns the TaskUUID field value if set, zero value otherwise.
func (o *Backup) GetTaskUUID() string {
	if o == nil || o.TaskUUID == nil {
		var ret string
		return ret
	}
	return *o.TaskUUID
}

// GetTaskUUIDOk returns a tuple with the TaskUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetTaskUUIDOk() (*string, bool) {
	if o == nil || o.TaskUUID == nil {
		return nil, false
	}
	return o.TaskUUID, true
}

// HasTaskUUID returns a boolean if a field has been set.
func (o *Backup) HasTaskUUID() bool {
	if o != nil && o.TaskUUID != nil {
		return true
	}

	return false
}

// SetTaskUUID gets a reference to the given string and assigns it to the TaskUUID field.
func (o *Backup) SetTaskUUID(v string) {
	o.TaskUUID = &v
}

// GetUniverseName returns the UniverseName field value if set, zero value otherwise.
func (o *Backup) GetUniverseName() string {
	if o == nil || o.UniverseName == nil {
		var ret string
		return ret
	}
	return *o.UniverseName
}

// GetUniverseNameOk returns a tuple with the UniverseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetUniverseNameOk() (*string, bool) {
	if o == nil || o.UniverseName == nil {
		return nil, false
	}
	return o.UniverseName, true
}

// HasUniverseName returns a boolean if a field has been set.
func (o *Backup) HasUniverseName() bool {
	if o != nil && o.UniverseName != nil {
		return true
	}

	return false
}

// SetUniverseName gets a reference to the given string and assigns it to the UniverseName field.
func (o *Backup) SetUniverseName(v string) {
	o.UniverseName = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Backup) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Backup) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *Backup) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Backup) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Backup) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Backup) SetVersion(v string) {
	o.Version = &v
}

func (o Backup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupInfo != nil {
		toSerialize["backupInfo"] = o.BackupInfo
	}
	if o.BackupUUID != nil {
		toSerialize["backupUUID"] = o.BackupUUID
	}
	if o.BaseBackupUUID != nil {
		toSerialize["baseBackupUUID"] = o.BaseBackupUUID
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CompletionTime != nil {
		toSerialize["completionTime"] = o.CompletionTime
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.CustomerUUID != nil {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if o.ExpiryTimeUnit != nil {
		toSerialize["expiryTimeUnit"] = o.ExpiryTimeUnit
	}
	if true {
		toSerialize["incrementalBackup"] = o.IncrementalBackup
	}
	if true {
		toSerialize["parentBackup"] = o.ParentBackup
	}
	if o.ScheduleName != nil {
		toSerialize["scheduleName"] = o.ScheduleName
	}
	if o.ScheduleUUID != nil {
		toSerialize["scheduleUUID"] = o.ScheduleUUID
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StorageConfigUUID != nil {
		toSerialize["storageConfigUUID"] = o.StorageConfigUUID
	}
	if o.TaskUUID != nil {
		toSerialize["taskUUID"] = o.TaskUUID
	}
	if o.UniverseName != nil {
		toSerialize["universeName"] = o.UniverseName
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBackup struct {
	value *Backup
	isSet bool
}

func (v NullableBackup) Get() *Backup {
	return v.value
}

func (v *NullableBackup) Set(val *Backup) {
	v.value = val
	v.isSet = true
}

func (v NullableBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackup(val *Backup) *NullableBackup {
	return &NullableBackup{value: val, isSet: true}
}

func (v NullableBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


