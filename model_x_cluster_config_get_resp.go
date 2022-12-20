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

// XClusterConfigGetResp xcluster get response
type XClusterConfigGetResp struct {
	// Create time of the xCluster config
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Lag metric data
	Lag map[string]interface{} `json:"lag"`
	// Last modify time of the xCluster config
	ModifyTime *time.Time `json:"modifyTime,omitempty"`
	// XCluster config name
	Name *string `json:"name,omitempty"`
	// Whether this xCluster replication config is paused
	Paused *bool `json:"paused,omitempty"`
	// Replication group name in DB
	ReplicationGroupName *string `json:"replicationGroupName,omitempty"`
	// Source Universe UUID
	SourceUniverseUUID *string `json:"sourceUniverseUUID,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
	TableDetails []XClusterTableConfig `json:"tableDetails"`
	Tables *[]string `json:"tables,omitempty"`
	// Target Universe UUID
	TargetUniverseUUID *string `json:"targetUniverseUUID,omitempty"`
	// XCluster config UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewXClusterConfigGetResp instantiates a new XClusterConfigGetResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXClusterConfigGetResp(lag map[string]interface{}, tableDetails []XClusterTableConfig, ) *XClusterConfigGetResp {
	this := XClusterConfigGetResp{}
	this.Lag = lag
	this.TableDetails = tableDetails
	return &this
}

// NewXClusterConfigGetRespWithDefaults instantiates a new XClusterConfigGetResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXClusterConfigGetRespWithDefaults() *XClusterConfigGetResp {
	this := XClusterConfigGetResp{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *XClusterConfigGetResp) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetLag returns the Lag field value
func (o *XClusterConfigGetResp) GetLag() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Lag
}

// GetLagOk returns a tuple with the Lag field value
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetLagOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Lag, true
}

// SetLag sets field value
func (o *XClusterConfigGetResp) SetLag(v map[string]interface{}) {
	o.Lag = v
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetModifyTime() time.Time {
	if o == nil || o.ModifyTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetModifyTimeOk() (*time.Time, bool) {
	if o == nil || o.ModifyTime == nil {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasModifyTime() bool {
	if o != nil && o.ModifyTime != nil {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given time.Time and assigns it to the ModifyTime field.
func (o *XClusterConfigGetResp) SetModifyTime(v time.Time) {
	o.ModifyTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XClusterConfigGetResp) SetName(v string) {
	o.Name = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetPaused() bool {
	if o == nil || o.Paused == nil {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetPausedOk() (*bool, bool) {
	if o == nil || o.Paused == nil {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasPaused() bool {
	if o != nil && o.Paused != nil {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *XClusterConfigGetResp) SetPaused(v bool) {
	o.Paused = &v
}

// GetReplicationGroupName returns the ReplicationGroupName field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetReplicationGroupName() string {
	if o == nil || o.ReplicationGroupName == nil {
		var ret string
		return ret
	}
	return *o.ReplicationGroupName
}

// GetReplicationGroupNameOk returns a tuple with the ReplicationGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetReplicationGroupNameOk() (*string, bool) {
	if o == nil || o.ReplicationGroupName == nil {
		return nil, false
	}
	return o.ReplicationGroupName, true
}

// HasReplicationGroupName returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasReplicationGroupName() bool {
	if o != nil && o.ReplicationGroupName != nil {
		return true
	}

	return false
}

// SetReplicationGroupName gets a reference to the given string and assigns it to the ReplicationGroupName field.
func (o *XClusterConfigGetResp) SetReplicationGroupName(v string) {
	o.ReplicationGroupName = &v
}

// GetSourceUniverseUUID returns the SourceUniverseUUID field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetSourceUniverseUUID() string {
	if o == nil || o.SourceUniverseUUID == nil {
		var ret string
		return ret
	}
	return *o.SourceUniverseUUID
}

// GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetSourceUniverseUUIDOk() (*string, bool) {
	if o == nil || o.SourceUniverseUUID == nil {
		return nil, false
	}
	return o.SourceUniverseUUID, true
}

// HasSourceUniverseUUID returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasSourceUniverseUUID() bool {
	if o != nil && o.SourceUniverseUUID != nil {
		return true
	}

	return false
}

// SetSourceUniverseUUID gets a reference to the given string and assigns it to the SourceUniverseUUID field.
func (o *XClusterConfigGetResp) SetSourceUniverseUUID(v string) {
	o.SourceUniverseUUID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *XClusterConfigGetResp) SetStatus(v string) {
	o.Status = &v
}

// GetTableDetails returns the TableDetails field value
func (o *XClusterConfigGetResp) GetTableDetails() []XClusterTableConfig {
	if o == nil  {
		var ret []XClusterTableConfig
		return ret
	}

	return o.TableDetails
}

// GetTableDetailsOk returns a tuple with the TableDetails field value
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetTableDetailsOk() (*[]XClusterTableConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TableDetails, true
}

// SetTableDetails sets field value
func (o *XClusterConfigGetResp) SetTableDetails(v []XClusterTableConfig) {
	o.TableDetails = v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return *o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *XClusterConfigGetResp) SetTables(v []string) {
	o.Tables = &v
}

// GetTargetUniverseUUID returns the TargetUniverseUUID field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetTargetUniverseUUID() string {
	if o == nil || o.TargetUniverseUUID == nil {
		var ret string
		return ret
	}
	return *o.TargetUniverseUUID
}

// GetTargetUniverseUUIDOk returns a tuple with the TargetUniverseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetTargetUniverseUUIDOk() (*string, bool) {
	if o == nil || o.TargetUniverseUUID == nil {
		return nil, false
	}
	return o.TargetUniverseUUID, true
}

// HasTargetUniverseUUID returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasTargetUniverseUUID() bool {
	if o != nil && o.TargetUniverseUUID != nil {
		return true
	}

	return false
}

// SetTargetUniverseUUID gets a reference to the given string and assigns it to the TargetUniverseUUID field.
func (o *XClusterConfigGetResp) SetTargetUniverseUUID(v string) {
	o.TargetUniverseUUID = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *XClusterConfigGetResp) SetUuid(v string) {
	o.Uuid = &v
}

func (o XClusterConfigGetResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if true {
		toSerialize["lag"] = o.Lag
	}
	if o.ModifyTime != nil {
		toSerialize["modifyTime"] = o.ModifyTime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Paused != nil {
		toSerialize["paused"] = o.Paused
	}
	if o.ReplicationGroupName != nil {
		toSerialize["replicationGroupName"] = o.ReplicationGroupName
	}
	if o.SourceUniverseUUID != nil {
		toSerialize["sourceUniverseUUID"] = o.SourceUniverseUUID
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["tableDetails"] = o.TableDetails
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	if o.TargetUniverseUUID != nil {
		toSerialize["targetUniverseUUID"] = o.TargetUniverseUUID
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableXClusterConfigGetResp struct {
	value *XClusterConfigGetResp
	isSet bool
}

func (v NullableXClusterConfigGetResp) Get() *XClusterConfigGetResp {
	return v.value
}

func (v *NullableXClusterConfigGetResp) Set(val *XClusterConfigGetResp) {
	v.value = val
	v.isSet = true
}

func (v NullableXClusterConfigGetResp) IsSet() bool {
	return v.isSet
}

func (v *NullableXClusterConfigGetResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXClusterConfigGetResp(val *XClusterConfigGetResp) *NullableXClusterConfigGetResp {
	return &NullableXClusterConfigGetResp{value: val, isSet: true}
}

func (v NullableXClusterConfigGetResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXClusterConfigGetResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


