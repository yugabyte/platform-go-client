/*
 * Yugabyte Platform APIs
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

// BackupResp struct for BackupResp
type BackupResp struct {
	BackupType string `json:"backupType"`
	BackupUUID string `json:"backupUUID"`
	CreateTime time.Time `json:"createTime"`
	CustomerUUID string `json:"customerUUID"`
	ExpiryTime time.Time `json:"expiryTime"`
	IsStorageConfigPresent bool `json:"isStorageConfigPresent"`
	IsUniversePresent bool `json:"isUniversePresent"`
	OnDemand bool `json:"onDemand"`
	ResponseList []KeyspaceTablesList `json:"responseList"`
	ScheduleUUID string `json:"scheduleUUID"`
	Sse bool `json:"sse"`
	State string `json:"state"`
	StorageConfigUUID string `json:"storageConfigUUID"`
	TotalBackupSizeInBytes int64 `json:"totalBackupSizeInBytes"`
	UniverseName string `json:"universeName"`
	UniverseUUID string `json:"universeUUID"`
	UpdateTime time.Time `json:"updateTime"`
}

// NewBackupResp instantiates a new BackupResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupResp(backupType string, backupUUID string, createTime time.Time, customerUUID string, expiryTime time.Time, isStorageConfigPresent bool, isUniversePresent bool, onDemand bool, responseList []KeyspaceTablesList, scheduleUUID string, sse bool, state string, storageConfigUUID string, totalBackupSizeInBytes int64, universeName string, universeUUID string, updateTime time.Time, ) *BackupResp {
	this := BackupResp{}
	this.BackupType = backupType
	this.BackupUUID = backupUUID
	this.CreateTime = createTime
	this.CustomerUUID = customerUUID
	this.ExpiryTime = expiryTime
	this.IsStorageConfigPresent = isStorageConfigPresent
	this.IsUniversePresent = isUniversePresent
	this.OnDemand = onDemand
	this.ResponseList = responseList
	this.ScheduleUUID = scheduleUUID
	this.Sse = sse
	this.State = state
	this.StorageConfigUUID = storageConfigUUID
	this.TotalBackupSizeInBytes = totalBackupSizeInBytes
	this.UniverseName = universeName
	this.UniverseUUID = universeUUID
	this.UpdateTime = updateTime
	return &this
}

// NewBackupRespWithDefaults instantiates a new BackupResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRespWithDefaults() *BackupResp {
	this := BackupResp{}
	return &this
}

// GetBackupType returns the BackupType field value
func (o *BackupResp) GetBackupType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetBackupTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupType, true
}

// SetBackupType sets field value
func (o *BackupResp) SetBackupType(v string) {
	o.BackupType = v
}

// GetBackupUUID returns the BackupUUID field value
func (o *BackupResp) GetBackupUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BackupUUID
}

// GetBackupUUIDOk returns a tuple with the BackupUUID field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetBackupUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupUUID, true
}

// SetBackupUUID sets field value
func (o *BackupResp) SetBackupUUID(v string) {
	o.BackupUUID = v
}

// GetCreateTime returns the CreateTime field value
func (o *BackupResp) GetCreateTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *BackupResp) SetCreateTime(v time.Time) {
	o.CreateTime = v
}

// GetCustomerUUID returns the CustomerUUID field value
func (o *BackupResp) GetCustomerUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetCustomerUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerUUID, true
}

// SetCustomerUUID sets field value
func (o *BackupResp) SetCustomerUUID(v string) {
	o.CustomerUUID = v
}

// GetExpiryTime returns the ExpiryTime field value
func (o *BackupResp) GetExpiryTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetExpiryTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiryTime, true
}

// SetExpiryTime sets field value
func (o *BackupResp) SetExpiryTime(v time.Time) {
	o.ExpiryTime = v
}

// GetIsStorageConfigPresent returns the IsStorageConfigPresent field value
func (o *BackupResp) GetIsStorageConfigPresent() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsStorageConfigPresent
}

// GetIsStorageConfigPresentOk returns a tuple with the IsStorageConfigPresent field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetIsStorageConfigPresentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsStorageConfigPresent, true
}

// SetIsStorageConfigPresent sets field value
func (o *BackupResp) SetIsStorageConfigPresent(v bool) {
	o.IsStorageConfigPresent = v
}

// GetIsUniversePresent returns the IsUniversePresent field value
func (o *BackupResp) GetIsUniversePresent() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsUniversePresent
}

// GetIsUniversePresentOk returns a tuple with the IsUniversePresent field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetIsUniversePresentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsUniversePresent, true
}

// SetIsUniversePresent sets field value
func (o *BackupResp) SetIsUniversePresent(v bool) {
	o.IsUniversePresent = v
}

// GetOnDemand returns the OnDemand field value
func (o *BackupResp) GetOnDemand() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.OnDemand
}

// GetOnDemandOk returns a tuple with the OnDemand field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetOnDemandOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OnDemand, true
}

// SetOnDemand sets field value
func (o *BackupResp) SetOnDemand(v bool) {
	o.OnDemand = v
}

// GetResponseList returns the ResponseList field value
func (o *BackupResp) GetResponseList() []KeyspaceTablesList {
	if o == nil  {
		var ret []KeyspaceTablesList
		return ret
	}

	return o.ResponseList
}

// GetResponseListOk returns a tuple with the ResponseList field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetResponseListOk() (*[]KeyspaceTablesList, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseList, true
}

// SetResponseList sets field value
func (o *BackupResp) SetResponseList(v []KeyspaceTablesList) {
	o.ResponseList = v
}

// GetScheduleUUID returns the ScheduleUUID field value
func (o *BackupResp) GetScheduleUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ScheduleUUID
}

// GetScheduleUUIDOk returns a tuple with the ScheduleUUID field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetScheduleUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScheduleUUID, true
}

// SetScheduleUUID sets field value
func (o *BackupResp) SetScheduleUUID(v string) {
	o.ScheduleUUID = v
}

// GetSse returns the Sse field value
func (o *BackupResp) GetSse() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Sse
}

// GetSseOk returns a tuple with the Sse field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetSseOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sse, true
}

// SetSse sets field value
func (o *BackupResp) SetSse(v bool) {
	o.Sse = v
}

// GetState returns the State field value
func (o *BackupResp) GetState() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *BackupResp) SetState(v string) {
	o.State = v
}

// GetStorageConfigUUID returns the StorageConfigUUID field value
func (o *BackupResp) GetStorageConfigUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.StorageConfigUUID
}

// GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetStorageConfigUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageConfigUUID, true
}

// SetStorageConfigUUID sets field value
func (o *BackupResp) SetStorageConfigUUID(v string) {
	o.StorageConfigUUID = v
}

// GetTotalBackupSizeInBytes returns the TotalBackupSizeInBytes field value
func (o *BackupResp) GetTotalBackupSizeInBytes() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.TotalBackupSizeInBytes
}

// GetTotalBackupSizeInBytesOk returns a tuple with the TotalBackupSizeInBytes field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetTotalBackupSizeInBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalBackupSizeInBytes, true
}

// SetTotalBackupSizeInBytes sets field value
func (o *BackupResp) SetTotalBackupSizeInBytes(v int64) {
	o.TotalBackupSizeInBytes = v
}

// GetUniverseName returns the UniverseName field value
func (o *BackupResp) GetUniverseName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniverseName
}

// GetUniverseNameOk returns a tuple with the UniverseName field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetUniverseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseName, true
}

// SetUniverseName sets field value
func (o *BackupResp) SetUniverseName(v string) {
	o.UniverseName = v
}

// GetUniverseUUID returns the UniverseUUID field value
func (o *BackupResp) GetUniverseUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniverseUUID
}

// GetUniverseUUIDOk returns a tuple with the UniverseUUID field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetUniverseUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseUUID, true
}

// SetUniverseUUID sets field value
func (o *BackupResp) SetUniverseUUID(v string) {
	o.UniverseUUID = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *BackupResp) GetUpdateTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *BackupResp) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *BackupResp) SetUpdateTime(v time.Time) {
	o.UpdateTime = v
}

func (o BackupResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupType"] = o.BackupType
	}
	if true {
		toSerialize["backupUUID"] = o.BackupUUID
	}
	if true {
		toSerialize["createTime"] = o.CreateTime
	}
	if true {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if true {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if true {
		toSerialize["isStorageConfigPresent"] = o.IsStorageConfigPresent
	}
	if true {
		toSerialize["isUniversePresent"] = o.IsUniversePresent
	}
	if true {
		toSerialize["onDemand"] = o.OnDemand
	}
	if true {
		toSerialize["responseList"] = o.ResponseList
	}
	if true {
		toSerialize["scheduleUUID"] = o.ScheduleUUID
	}
	if true {
		toSerialize["sse"] = o.Sse
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["storageConfigUUID"] = o.StorageConfigUUID
	}
	if true {
		toSerialize["totalBackupSizeInBytes"] = o.TotalBackupSizeInBytes
	}
	if true {
		toSerialize["universeName"] = o.UniverseName
	}
	if true {
		toSerialize["universeUUID"] = o.UniverseUUID
	}
	if true {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableBackupResp struct {
	value *BackupResp
	isSet bool
}

func (v NullableBackupResp) Get() *BackupResp {
	return v.value
}

func (v *NullableBackupResp) Set(val *BackupResp) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupResp) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupResp(val *BackupResp) *NullableBackupResp {
	return &NullableBackupResp{value: val, isSet: true}
}

func (v NullableBackupResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


