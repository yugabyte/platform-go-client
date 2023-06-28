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

// XClusterConfig xcluster config object
type XClusterConfig struct {
	// Create time of the xCluster config
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Whether this xCluster replication config was imported
	Imported *bool `json:"imported,omitempty"`
	// Last modify time of the xCluster config
	ModifyTime *time.Time `json:"modifyTime,omitempty"`
	// XCluster config name
	Name *string `json:"name,omitempty"`
	// Whether this xCluster replication config is paused
	Paused *bool `json:"paused,omitempty"`
	// Replication group name in DB
	ReplicationGroupName *string `json:"replicationGroupName,omitempty"`
	// Whether the source is active in txn xCluster
	SourceActive *bool `json:"sourceActive,omitempty"`
	// Source Universe UUID
	SourceUniverseUUID *string `json:"sourceUniverseUUID,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
	// Tables participating in this xCluster config
	TableDetails *[]XClusterTableConfig `json:"tableDetails,omitempty"`
	// tableType
	TableType *string `json:"tableType,omitempty"`
	Tables *[]string `json:"tables,omitempty"`
	// Whether the target is active in txn xCluster
	TargetActive *bool `json:"targetActive,omitempty"`
	// Target Universe UUID
	TargetUniverseUUID *string `json:"targetUniverseUUID,omitempty"`
	// Whether the config is txn xCluster
	Type *string `json:"type,omitempty"`
	// XCluster config UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewXClusterConfig instantiates a new XClusterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXClusterConfig() *XClusterConfig {
	this := XClusterConfig{}
	return &this
}

// NewXClusterConfigWithDefaults instantiates a new XClusterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXClusterConfigWithDefaults() *XClusterConfig {
	this := XClusterConfig{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *XClusterConfig) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *XClusterConfig) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *XClusterConfig) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetImported returns the Imported field value if set, zero value otherwise.
func (o *XClusterConfig) GetImported() bool {
	if o == nil || o.Imported == nil {
		var ret bool
		return ret
	}
	return *o.Imported
}

// GetImportedOk returns a tuple with the Imported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetImportedOk() (*bool, bool) {
	if o == nil || o.Imported == nil {
		return nil, false
	}
	return o.Imported, true
}

// HasImported returns a boolean if a field has been set.
func (o *XClusterConfig) HasImported() bool {
	if o != nil && o.Imported != nil {
		return true
	}

	return false
}

// SetImported gets a reference to the given bool and assigns it to the Imported field.
func (o *XClusterConfig) SetImported(v bool) {
	o.Imported = &v
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *XClusterConfig) GetModifyTime() time.Time {
	if o == nil || o.ModifyTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetModifyTimeOk() (*time.Time, bool) {
	if o == nil || o.ModifyTime == nil {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *XClusterConfig) HasModifyTime() bool {
	if o != nil && o.ModifyTime != nil {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given time.Time and assigns it to the ModifyTime field.
func (o *XClusterConfig) SetModifyTime(v time.Time) {
	o.ModifyTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XClusterConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XClusterConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XClusterConfig) SetName(v string) {
	o.Name = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *XClusterConfig) GetPaused() bool {
	if o == nil || o.Paused == nil {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetPausedOk() (*bool, bool) {
	if o == nil || o.Paused == nil {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *XClusterConfig) HasPaused() bool {
	if o != nil && o.Paused != nil {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *XClusterConfig) SetPaused(v bool) {
	o.Paused = &v
}

// GetReplicationGroupName returns the ReplicationGroupName field value if set, zero value otherwise.
func (o *XClusterConfig) GetReplicationGroupName() string {
	if o == nil || o.ReplicationGroupName == nil {
		var ret string
		return ret
	}
	return *o.ReplicationGroupName
}

// GetReplicationGroupNameOk returns a tuple with the ReplicationGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetReplicationGroupNameOk() (*string, bool) {
	if o == nil || o.ReplicationGroupName == nil {
		return nil, false
	}
	return o.ReplicationGroupName, true
}

// HasReplicationGroupName returns a boolean if a field has been set.
func (o *XClusterConfig) HasReplicationGroupName() bool {
	if o != nil && o.ReplicationGroupName != nil {
		return true
	}

	return false
}

// SetReplicationGroupName gets a reference to the given string and assigns it to the ReplicationGroupName field.
func (o *XClusterConfig) SetReplicationGroupName(v string) {
	o.ReplicationGroupName = &v
}

// GetSourceActive returns the SourceActive field value if set, zero value otherwise.
func (o *XClusterConfig) GetSourceActive() bool {
	if o == nil || o.SourceActive == nil {
		var ret bool
		return ret
	}
	return *o.SourceActive
}

// GetSourceActiveOk returns a tuple with the SourceActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetSourceActiveOk() (*bool, bool) {
	if o == nil || o.SourceActive == nil {
		return nil, false
	}
	return o.SourceActive, true
}

// HasSourceActive returns a boolean if a field has been set.
func (o *XClusterConfig) HasSourceActive() bool {
	if o != nil && o.SourceActive != nil {
		return true
	}

	return false
}

// SetSourceActive gets a reference to the given bool and assigns it to the SourceActive field.
func (o *XClusterConfig) SetSourceActive(v bool) {
	o.SourceActive = &v
}

// GetSourceUniverseUUID returns the SourceUniverseUUID field value if set, zero value otherwise.
func (o *XClusterConfig) GetSourceUniverseUUID() string {
	if o == nil || o.SourceUniverseUUID == nil {
		var ret string
		return ret
	}
	return *o.SourceUniverseUUID
}

// GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetSourceUniverseUUIDOk() (*string, bool) {
	if o == nil || o.SourceUniverseUUID == nil {
		return nil, false
	}
	return o.SourceUniverseUUID, true
}

// HasSourceUniverseUUID returns a boolean if a field has been set.
func (o *XClusterConfig) HasSourceUniverseUUID() bool {
	if o != nil && o.SourceUniverseUUID != nil {
		return true
	}

	return false
}

// SetSourceUniverseUUID gets a reference to the given string and assigns it to the SourceUniverseUUID field.
func (o *XClusterConfig) SetSourceUniverseUUID(v string) {
	o.SourceUniverseUUID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *XClusterConfig) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *XClusterConfig) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *XClusterConfig) SetStatus(v string) {
	o.Status = &v
}

// GetTableDetails returns the TableDetails field value if set, zero value otherwise.
func (o *XClusterConfig) GetTableDetails() []XClusterTableConfig {
	if o == nil || o.TableDetails == nil {
		var ret []XClusterTableConfig
		return ret
	}
	return *o.TableDetails
}

// GetTableDetailsOk returns a tuple with the TableDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetTableDetailsOk() (*[]XClusterTableConfig, bool) {
	if o == nil || o.TableDetails == nil {
		return nil, false
	}
	return o.TableDetails, true
}

// HasTableDetails returns a boolean if a field has been set.
func (o *XClusterConfig) HasTableDetails() bool {
	if o != nil && o.TableDetails != nil {
		return true
	}

	return false
}

// SetTableDetails gets a reference to the given []XClusterTableConfig and assigns it to the TableDetails field.
func (o *XClusterConfig) SetTableDetails(v []XClusterTableConfig) {
	o.TableDetails = &v
}

// GetTableType returns the TableType field value if set, zero value otherwise.
func (o *XClusterConfig) GetTableType() string {
	if o == nil || o.TableType == nil {
		var ret string
		return ret
	}
	return *o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetTableTypeOk() (*string, bool) {
	if o == nil || o.TableType == nil {
		return nil, false
	}
	return o.TableType, true
}

// HasTableType returns a boolean if a field has been set.
func (o *XClusterConfig) HasTableType() bool {
	if o != nil && o.TableType != nil {
		return true
	}

	return false
}

// SetTableType gets a reference to the given string and assigns it to the TableType field.
func (o *XClusterConfig) SetTableType(v string) {
	o.TableType = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *XClusterConfig) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return *o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *XClusterConfig) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *XClusterConfig) SetTables(v []string) {
	o.Tables = &v
}

// GetTargetActive returns the TargetActive field value if set, zero value otherwise.
func (o *XClusterConfig) GetTargetActive() bool {
	if o == nil || o.TargetActive == nil {
		var ret bool
		return ret
	}
	return *o.TargetActive
}

// GetTargetActiveOk returns a tuple with the TargetActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetTargetActiveOk() (*bool, bool) {
	if o == nil || o.TargetActive == nil {
		return nil, false
	}
	return o.TargetActive, true
}

// HasTargetActive returns a boolean if a field has been set.
func (o *XClusterConfig) HasTargetActive() bool {
	if o != nil && o.TargetActive != nil {
		return true
	}

	return false
}

// SetTargetActive gets a reference to the given bool and assigns it to the TargetActive field.
func (o *XClusterConfig) SetTargetActive(v bool) {
	o.TargetActive = &v
}

// GetTargetUniverseUUID returns the TargetUniverseUUID field value if set, zero value otherwise.
func (o *XClusterConfig) GetTargetUniverseUUID() string {
	if o == nil || o.TargetUniverseUUID == nil {
		var ret string
		return ret
	}
	return *o.TargetUniverseUUID
}

// GetTargetUniverseUUIDOk returns a tuple with the TargetUniverseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetTargetUniverseUUIDOk() (*string, bool) {
	if o == nil || o.TargetUniverseUUID == nil {
		return nil, false
	}
	return o.TargetUniverseUUID, true
}

// HasTargetUniverseUUID returns a boolean if a field has been set.
func (o *XClusterConfig) HasTargetUniverseUUID() bool {
	if o != nil && o.TargetUniverseUUID != nil {
		return true
	}

	return false
}

// SetTargetUniverseUUID gets a reference to the given string and assigns it to the TargetUniverseUUID field.
func (o *XClusterConfig) SetTargetUniverseUUID(v string) {
	o.TargetUniverseUUID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *XClusterConfig) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *XClusterConfig) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *XClusterConfig) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *XClusterConfig) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfig) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *XClusterConfig) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *XClusterConfig) SetUuid(v string) {
	o.Uuid = &v
}

func (o XClusterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.Imported != nil {
		toSerialize["imported"] = o.Imported
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
	if o.SourceActive != nil {
		toSerialize["sourceActive"] = o.SourceActive
	}
	if o.SourceUniverseUUID != nil {
		toSerialize["sourceUniverseUUID"] = o.SourceUniverseUUID
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TableDetails != nil {
		toSerialize["tableDetails"] = o.TableDetails
	}
	if o.TableType != nil {
		toSerialize["tableType"] = o.TableType
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	if o.TargetActive != nil {
		toSerialize["targetActive"] = o.TargetActive
	}
	if o.TargetUniverseUUID != nil {
		toSerialize["targetUniverseUUID"] = o.TargetUniverseUUID
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableXClusterConfig struct {
	value *XClusterConfig
	isSet bool
}

func (v NullableXClusterConfig) Get() *XClusterConfig {
	return v.value
}

func (v *NullableXClusterConfig) Set(val *XClusterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableXClusterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableXClusterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXClusterConfig(val *XClusterConfig) *NullableXClusterConfig {
	return &NullableXClusterConfig{value: val, isSet: true}
}

func (v NullableXClusterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXClusterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


