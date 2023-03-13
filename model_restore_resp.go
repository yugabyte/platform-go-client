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

// RestoreResp struct for RestoreResp
type RestoreResp struct {
	// Restore creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	CustomerUUID string `json:"customerUUID"`
	IsSourceUniversePresent bool `json:"isSourceUniversePresent"`
	RestoreKeyspaceList []RestoreKeyspace `json:"restoreKeyspaceList"`
	RestoreSizeInBytes int64 `json:"restoreSizeInBytes"`
	RestoreUUID string `json:"restoreUUID"`
	SourceUniverseName string `json:"sourceUniverseName"`
	SourceUniverseUUID string `json:"sourceUniverseUUID"`
	State string `json:"state"`
	TargetUniverseName string `json:"targetUniverseName"`
	UniverseUUID string `json:"universeUUID"`
	// Restore update time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
}

// NewRestoreResp instantiates a new RestoreResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreResp(customerUUID string, isSourceUniversePresent bool, restoreKeyspaceList []RestoreKeyspace, restoreSizeInBytes int64, restoreUUID string, sourceUniverseName string, sourceUniverseUUID string, state string, targetUniverseName string, universeUUID string, ) *RestoreResp {
	this := RestoreResp{}
	this.CustomerUUID = customerUUID
	this.IsSourceUniversePresent = isSourceUniversePresent
	this.RestoreKeyspaceList = restoreKeyspaceList
	this.RestoreSizeInBytes = restoreSizeInBytes
	this.RestoreUUID = restoreUUID
	this.SourceUniverseName = sourceUniverseName
	this.SourceUniverseUUID = sourceUniverseUUID
	this.State = state
	this.TargetUniverseName = targetUniverseName
	this.UniverseUUID = universeUUID
	return &this
}

// NewRestoreRespWithDefaults instantiates a new RestoreResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreRespWithDefaults() *RestoreResp {
	this := RestoreResp{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *RestoreResp) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *RestoreResp) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *RestoreResp) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetCustomerUUID returns the CustomerUUID field value
func (o *RestoreResp) GetCustomerUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetCustomerUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerUUID, true
}

// SetCustomerUUID sets field value
func (o *RestoreResp) SetCustomerUUID(v string) {
	o.CustomerUUID = v
}

// GetIsSourceUniversePresent returns the IsSourceUniversePresent field value
func (o *RestoreResp) GetIsSourceUniversePresent() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsSourceUniversePresent
}

// GetIsSourceUniversePresentOk returns a tuple with the IsSourceUniversePresent field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetIsSourceUniversePresentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSourceUniversePresent, true
}

// SetIsSourceUniversePresent sets field value
func (o *RestoreResp) SetIsSourceUniversePresent(v bool) {
	o.IsSourceUniversePresent = v
}

// GetRestoreKeyspaceList returns the RestoreKeyspaceList field value
func (o *RestoreResp) GetRestoreKeyspaceList() []RestoreKeyspace {
	if o == nil  {
		var ret []RestoreKeyspace
		return ret
	}

	return o.RestoreKeyspaceList
}

// GetRestoreKeyspaceListOk returns a tuple with the RestoreKeyspaceList field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetRestoreKeyspaceListOk() (*[]RestoreKeyspace, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RestoreKeyspaceList, true
}

// SetRestoreKeyspaceList sets field value
func (o *RestoreResp) SetRestoreKeyspaceList(v []RestoreKeyspace) {
	o.RestoreKeyspaceList = v
}

// GetRestoreSizeInBytes returns the RestoreSizeInBytes field value
func (o *RestoreResp) GetRestoreSizeInBytes() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.RestoreSizeInBytes
}

// GetRestoreSizeInBytesOk returns a tuple with the RestoreSizeInBytes field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetRestoreSizeInBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RestoreSizeInBytes, true
}

// SetRestoreSizeInBytes sets field value
func (o *RestoreResp) SetRestoreSizeInBytes(v int64) {
	o.RestoreSizeInBytes = v
}

// GetRestoreUUID returns the RestoreUUID field value
func (o *RestoreResp) GetRestoreUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RestoreUUID
}

// GetRestoreUUIDOk returns a tuple with the RestoreUUID field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetRestoreUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RestoreUUID, true
}

// SetRestoreUUID sets field value
func (o *RestoreResp) SetRestoreUUID(v string) {
	o.RestoreUUID = v
}

// GetSourceUniverseName returns the SourceUniverseName field value
func (o *RestoreResp) GetSourceUniverseName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SourceUniverseName
}

// GetSourceUniverseNameOk returns a tuple with the SourceUniverseName field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetSourceUniverseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceUniverseName, true
}

// SetSourceUniverseName sets field value
func (o *RestoreResp) SetSourceUniverseName(v string) {
	o.SourceUniverseName = v
}

// GetSourceUniverseUUID returns the SourceUniverseUUID field value
func (o *RestoreResp) GetSourceUniverseUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SourceUniverseUUID
}

// GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetSourceUniverseUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceUniverseUUID, true
}

// SetSourceUniverseUUID sets field value
func (o *RestoreResp) SetSourceUniverseUUID(v string) {
	o.SourceUniverseUUID = v
}

// GetState returns the State field value
func (o *RestoreResp) GetState() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RestoreResp) SetState(v string) {
	o.State = v
}

// GetTargetUniverseName returns the TargetUniverseName field value
func (o *RestoreResp) GetTargetUniverseName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetUniverseName
}

// GetTargetUniverseNameOk returns a tuple with the TargetUniverseName field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetTargetUniverseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetUniverseName, true
}

// SetTargetUniverseName sets field value
func (o *RestoreResp) SetTargetUniverseName(v string) {
	o.TargetUniverseName = v
}

// GetUniverseUUID returns the UniverseUUID field value
func (o *RestoreResp) GetUniverseUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniverseUUID
}

// GetUniverseUUIDOk returns a tuple with the UniverseUUID field value
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetUniverseUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseUUID, true
}

// SetUniverseUUID sets field value
func (o *RestoreResp) SetUniverseUUID(v string) {
	o.UniverseUUID = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *RestoreResp) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreResp) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *RestoreResp) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *RestoreResp) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

func (o RestoreResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if true {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if true {
		toSerialize["isSourceUniversePresent"] = o.IsSourceUniversePresent
	}
	if true {
		toSerialize["restoreKeyspaceList"] = o.RestoreKeyspaceList
	}
	if true {
		toSerialize["restoreSizeInBytes"] = o.RestoreSizeInBytes
	}
	if true {
		toSerialize["restoreUUID"] = o.RestoreUUID
	}
	if true {
		toSerialize["sourceUniverseName"] = o.SourceUniverseName
	}
	if true {
		toSerialize["sourceUniverseUUID"] = o.SourceUniverseUUID
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["targetUniverseName"] = o.TargetUniverseName
	}
	if true {
		toSerialize["universeUUID"] = o.UniverseUUID
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreResp struct {
	value *RestoreResp
	isSet bool
}

func (v NullableRestoreResp) Get() *RestoreResp {
	return v.value
}

func (v *NullableRestoreResp) Set(val *RestoreResp) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreResp) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreResp(val *RestoreResp) *NullableRestoreResp {
	return &NullableRestoreResp{value: val, isSet: true}
}

func (v NullableRestoreResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


