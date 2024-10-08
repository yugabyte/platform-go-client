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

// BackupApiFilter struct for BackupApiFilter
type BackupApiFilter struct {
	BackupUUIDList []string `json:"backupUUIDList"`
	// The end date for backup filter.
	DateRangeEnd *time.Time `json:"dateRangeEnd,omitempty"`
	// The start date for backup filter.
	DateRangeStart *time.Time `json:"dateRangeStart,omitempty"`
	KeyspaceList []string `json:"keyspaceList"`
	OnlyShowDeletedConfigs bool `json:"onlyShowDeletedConfigs"`
	OnlyShowDeletedUniverses bool `json:"onlyShowDeletedUniverses"`
	ScheduleUUIDList []string `json:"scheduleUUIDList"`
	ShowHidden bool `json:"showHidden"`
	States []string `json:"states"`
	StorageConfigUUIDList []string `json:"storageConfigUUIDList"`
	UniverseNameList []string `json:"universeNameList"`
	UniverseUUIDList []string `json:"universeUUIDList"`
}

// NewBackupApiFilter instantiates a new BackupApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupApiFilter(backupUUIDList []string, keyspaceList []string, onlyShowDeletedConfigs bool, onlyShowDeletedUniverses bool, scheduleUUIDList []string, showHidden bool, states []string, storageConfigUUIDList []string, universeNameList []string, universeUUIDList []string) *BackupApiFilter {
	this := BackupApiFilter{}
	this.BackupUUIDList = backupUUIDList
	this.KeyspaceList = keyspaceList
	this.OnlyShowDeletedConfigs = onlyShowDeletedConfigs
	this.OnlyShowDeletedUniverses = onlyShowDeletedUniverses
	this.ScheduleUUIDList = scheduleUUIDList
	this.ShowHidden = showHidden
	this.States = states
	this.StorageConfigUUIDList = storageConfigUUIDList
	this.UniverseNameList = universeNameList
	this.UniverseUUIDList = universeUUIDList
	return &this
}

// NewBackupApiFilterWithDefaults instantiates a new BackupApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupApiFilterWithDefaults() *BackupApiFilter {
	this := BackupApiFilter{}
	return &this
}

// GetBackupUUIDList returns the BackupUUIDList field value
func (o *BackupApiFilter) GetBackupUUIDList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BackupUUIDList
}

// GetBackupUUIDListOk returns a tuple with the BackupUUIDList field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetBackupUUIDListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupUUIDList, true
}

// SetBackupUUIDList sets field value
func (o *BackupApiFilter) SetBackupUUIDList(v []string) {
	o.BackupUUIDList = v
}

// GetDateRangeEnd returns the DateRangeEnd field value if set, zero value otherwise.
func (o *BackupApiFilter) GetDateRangeEnd() time.Time {
	if o == nil || o.DateRangeEnd == nil {
		var ret time.Time
		return ret
	}
	return *o.DateRangeEnd
}

// GetDateRangeEndOk returns a tuple with the DateRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetDateRangeEndOk() (*time.Time, bool) {
	if o == nil || o.DateRangeEnd == nil {
		return nil, false
	}
	return o.DateRangeEnd, true
}

// HasDateRangeEnd returns a boolean if a field has been set.
func (o *BackupApiFilter) HasDateRangeEnd() bool {
	if o != nil && o.DateRangeEnd != nil {
		return true
	}

	return false
}

// SetDateRangeEnd gets a reference to the given time.Time and assigns it to the DateRangeEnd field.
func (o *BackupApiFilter) SetDateRangeEnd(v time.Time) {
	o.DateRangeEnd = &v
}

// GetDateRangeStart returns the DateRangeStart field value if set, zero value otherwise.
func (o *BackupApiFilter) GetDateRangeStart() time.Time {
	if o == nil || o.DateRangeStart == nil {
		var ret time.Time
		return ret
	}
	return *o.DateRangeStart
}

// GetDateRangeStartOk returns a tuple with the DateRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetDateRangeStartOk() (*time.Time, bool) {
	if o == nil || o.DateRangeStart == nil {
		return nil, false
	}
	return o.DateRangeStart, true
}

// HasDateRangeStart returns a boolean if a field has been set.
func (o *BackupApiFilter) HasDateRangeStart() bool {
	if o != nil && o.DateRangeStart != nil {
		return true
	}

	return false
}

// SetDateRangeStart gets a reference to the given time.Time and assigns it to the DateRangeStart field.
func (o *BackupApiFilter) SetDateRangeStart(v time.Time) {
	o.DateRangeStart = &v
}

// GetKeyspaceList returns the KeyspaceList field value
func (o *BackupApiFilter) GetKeyspaceList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.KeyspaceList
}

// GetKeyspaceListOk returns a tuple with the KeyspaceList field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetKeyspaceListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyspaceList, true
}

// SetKeyspaceList sets field value
func (o *BackupApiFilter) SetKeyspaceList(v []string) {
	o.KeyspaceList = v
}

// GetOnlyShowDeletedConfigs returns the OnlyShowDeletedConfigs field value
func (o *BackupApiFilter) GetOnlyShowDeletedConfigs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OnlyShowDeletedConfigs
}

// GetOnlyShowDeletedConfigsOk returns a tuple with the OnlyShowDeletedConfigs field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetOnlyShowDeletedConfigsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OnlyShowDeletedConfigs, true
}

// SetOnlyShowDeletedConfigs sets field value
func (o *BackupApiFilter) SetOnlyShowDeletedConfigs(v bool) {
	o.OnlyShowDeletedConfigs = v
}

// GetOnlyShowDeletedUniverses returns the OnlyShowDeletedUniverses field value
func (o *BackupApiFilter) GetOnlyShowDeletedUniverses() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OnlyShowDeletedUniverses
}

// GetOnlyShowDeletedUniversesOk returns a tuple with the OnlyShowDeletedUniverses field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetOnlyShowDeletedUniversesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OnlyShowDeletedUniverses, true
}

// SetOnlyShowDeletedUniverses sets field value
func (o *BackupApiFilter) SetOnlyShowDeletedUniverses(v bool) {
	o.OnlyShowDeletedUniverses = v
}

// GetScheduleUUIDList returns the ScheduleUUIDList field value
func (o *BackupApiFilter) GetScheduleUUIDList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScheduleUUIDList
}

// GetScheduleUUIDListOk returns a tuple with the ScheduleUUIDList field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetScheduleUUIDListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScheduleUUIDList, true
}

// SetScheduleUUIDList sets field value
func (o *BackupApiFilter) SetScheduleUUIDList(v []string) {
	o.ScheduleUUIDList = v
}

// GetShowHidden returns the ShowHidden field value
func (o *BackupApiFilter) GetShowHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowHidden
}

// GetShowHiddenOk returns a tuple with the ShowHidden field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetShowHiddenOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ShowHidden, true
}

// SetShowHidden sets field value
func (o *BackupApiFilter) SetShowHidden(v bool) {
	o.ShowHidden = v
}

// GetStates returns the States field value
func (o *BackupApiFilter) GetStates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.States
}

// GetStatesOk returns a tuple with the States field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetStatesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.States, true
}

// SetStates sets field value
func (o *BackupApiFilter) SetStates(v []string) {
	o.States = v
}

// GetStorageConfigUUIDList returns the StorageConfigUUIDList field value
func (o *BackupApiFilter) GetStorageConfigUUIDList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StorageConfigUUIDList
}

// GetStorageConfigUUIDListOk returns a tuple with the StorageConfigUUIDList field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetStorageConfigUUIDListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageConfigUUIDList, true
}

// SetStorageConfigUUIDList sets field value
func (o *BackupApiFilter) SetStorageConfigUUIDList(v []string) {
	o.StorageConfigUUIDList = v
}

// GetUniverseNameList returns the UniverseNameList field value
func (o *BackupApiFilter) GetUniverseNameList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UniverseNameList
}

// GetUniverseNameListOk returns a tuple with the UniverseNameList field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetUniverseNameListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseNameList, true
}

// SetUniverseNameList sets field value
func (o *BackupApiFilter) SetUniverseNameList(v []string) {
	o.UniverseNameList = v
}

// GetUniverseUUIDList returns the UniverseUUIDList field value
func (o *BackupApiFilter) GetUniverseUUIDList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UniverseUUIDList
}

// GetUniverseUUIDListOk returns a tuple with the UniverseUUIDList field value
// and a boolean to check if the value has been set.
func (o *BackupApiFilter) GetUniverseUUIDListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseUUIDList, true
}

// SetUniverseUUIDList sets field value
func (o *BackupApiFilter) SetUniverseUUIDList(v []string) {
	o.UniverseUUIDList = v
}

func (o BackupApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupUUIDList"] = o.BackupUUIDList
	}
	if o.DateRangeEnd != nil {
		toSerialize["dateRangeEnd"] = o.DateRangeEnd
	}
	if o.DateRangeStart != nil {
		toSerialize["dateRangeStart"] = o.DateRangeStart
	}
	if true {
		toSerialize["keyspaceList"] = o.KeyspaceList
	}
	if true {
		toSerialize["onlyShowDeletedConfigs"] = o.OnlyShowDeletedConfigs
	}
	if true {
		toSerialize["onlyShowDeletedUniverses"] = o.OnlyShowDeletedUniverses
	}
	if true {
		toSerialize["scheduleUUIDList"] = o.ScheduleUUIDList
	}
	if true {
		toSerialize["showHidden"] = o.ShowHidden
	}
	if true {
		toSerialize["states"] = o.States
	}
	if true {
		toSerialize["storageConfigUUIDList"] = o.StorageConfigUUIDList
	}
	if true {
		toSerialize["universeNameList"] = o.UniverseNameList
	}
	if true {
		toSerialize["universeUUIDList"] = o.UniverseUUIDList
	}
	return json.Marshal(toSerialize)
}

type NullableBackupApiFilter struct {
	value *BackupApiFilter
	isSet bool
}

func (v NullableBackupApiFilter) Get() *BackupApiFilter {
	return v.value
}

func (v *NullableBackupApiFilter) Set(val *BackupApiFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupApiFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupApiFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupApiFilter(val *BackupApiFilter) *NullableBackupApiFilter {
	return &NullableBackupApiFilter{value: val, isSet: true}
}

func (v NullableBackupApiFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupApiFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


