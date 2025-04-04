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
	Dbs *[]string `json:"dbs,omitempty"`
	// YbaApi Internal. Whether this xCluster replication config was imported
	Imported *bool `json:"imported,omitempty"`
	// WARNING: This is a preview API that could change. The keyspace name that the xCluster task is working on; used for disaster recovery
	KeyspacePending *string `json:"keyspacePending,omitempty"`
	// Lag metric data
	Lag map[string]interface{} `json:"lag"`
	// Last modify time of the xCluster config
	ModifyTime *time.Time `json:"modifyTime,omitempty"`
	// XCluster config name
	Name *string `json:"name,omitempty"`
	// Namespaces participating in this xCluster config
	NamespaceDetails *[]XClusterNamespaceConfig `json:"namespaceDetails,omitempty"`
	Namespaces []XClusterNamespaceConfig `json:"namespaces"`
	// Whether this xCluster replication config is paused
	Paused *bool `json:"paused,omitempty"`
	// WARNING: This is a preview API that could change. The list of PITR configs used for the txn xCluster config
	PitrConfigs *[]PitrConfig `json:"pitrConfigs,omitempty"`
	// Replication group name in the target universe cluster config
	ReplicationGroupName *string `json:"replicationGroupName,omitempty"`
	// WARNING: This is a preview API that could change. Whether this xCluster config is used as a secondary config for a DR config
	Secondary *bool `json:"secondary,omitempty"`
	// Whether the source is active in txn xCluster
	SourceActive *bool `json:"sourceActive,omitempty"`
	// WARNING: This is a preview API that could change. The replication status of the source universe; used for disaster recovery
	SourceUniverseState *string `json:"sourceUniverseState,omitempty"`
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
	// WARNING: This is a preview API that could change. The replication status of the target universe; used for disaster recovery
	TargetUniverseState *string `json:"targetUniverseState,omitempty"`
	// Target Universe UUID
	TargetUniverseUUID *string `json:"targetUniverseUUID,omitempty"`
	// Whether the config is basic, txn, or db scoped xCluster
	Type *string `json:"type,omitempty"`
	// WARNING: This is a preview API that could change. Whether the xCluster config is used as part of a DR config
	UsedForDr *bool `json:"usedForDr,omitempty"`
	// XCluster config UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewXClusterConfigGetResp instantiates a new XClusterConfigGetResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXClusterConfigGetResp(lag map[string]interface{}, namespaces []XClusterNamespaceConfig) *XClusterConfigGetResp {
	this := XClusterConfigGetResp{}
	this.Lag = lag
	this.Namespaces = namespaces
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

// GetDbs returns the Dbs field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetDbs() []string {
	if o == nil || o.Dbs == nil {
		var ret []string
		return ret
	}
	return *o.Dbs
}

// GetDbsOk returns a tuple with the Dbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetDbsOk() (*[]string, bool) {
	if o == nil || o.Dbs == nil {
		return nil, false
	}
	return o.Dbs, true
}

// HasDbs returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasDbs() bool {
	if o != nil && o.Dbs != nil {
		return true
	}

	return false
}

// SetDbs gets a reference to the given []string and assigns it to the Dbs field.
func (o *XClusterConfigGetResp) SetDbs(v []string) {
	o.Dbs = &v
}

// GetImported returns the Imported field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetImported() bool {
	if o == nil || o.Imported == nil {
		var ret bool
		return ret
	}
	return *o.Imported
}

// GetImportedOk returns a tuple with the Imported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetImportedOk() (*bool, bool) {
	if o == nil || o.Imported == nil {
		return nil, false
	}
	return o.Imported, true
}

// HasImported returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasImported() bool {
	if o != nil && o.Imported != nil {
		return true
	}

	return false
}

// SetImported gets a reference to the given bool and assigns it to the Imported field.
func (o *XClusterConfigGetResp) SetImported(v bool) {
	o.Imported = &v
}

// GetKeyspacePending returns the KeyspacePending field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetKeyspacePending() string {
	if o == nil || o.KeyspacePending == nil {
		var ret string
		return ret
	}
	return *o.KeyspacePending
}

// GetKeyspacePendingOk returns a tuple with the KeyspacePending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetKeyspacePendingOk() (*string, bool) {
	if o == nil || o.KeyspacePending == nil {
		return nil, false
	}
	return o.KeyspacePending, true
}

// HasKeyspacePending returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasKeyspacePending() bool {
	if o != nil && o.KeyspacePending != nil {
		return true
	}

	return false
}

// SetKeyspacePending gets a reference to the given string and assigns it to the KeyspacePending field.
func (o *XClusterConfigGetResp) SetKeyspacePending(v string) {
	o.KeyspacePending = &v
}

// GetLag returns the Lag field value
func (o *XClusterConfigGetResp) GetLag() map[string]interface{} {
	if o == nil {
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

// GetNamespaceDetails returns the NamespaceDetails field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetNamespaceDetails() []XClusterNamespaceConfig {
	if o == nil || o.NamespaceDetails == nil {
		var ret []XClusterNamespaceConfig
		return ret
	}
	return *o.NamespaceDetails
}

// GetNamespaceDetailsOk returns a tuple with the NamespaceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetNamespaceDetailsOk() (*[]XClusterNamespaceConfig, bool) {
	if o == nil || o.NamespaceDetails == nil {
		return nil, false
	}
	return o.NamespaceDetails, true
}

// HasNamespaceDetails returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasNamespaceDetails() bool {
	if o != nil && o.NamespaceDetails != nil {
		return true
	}

	return false
}

// SetNamespaceDetails gets a reference to the given []XClusterNamespaceConfig and assigns it to the NamespaceDetails field.
func (o *XClusterConfigGetResp) SetNamespaceDetails(v []XClusterNamespaceConfig) {
	o.NamespaceDetails = &v
}

// GetNamespaces returns the Namespaces field value
func (o *XClusterConfigGetResp) GetNamespaces() []XClusterNamespaceConfig {
	if o == nil {
		var ret []XClusterNamespaceConfig
		return ret
	}

	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetNamespacesOk() (*[]XClusterNamespaceConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Namespaces, true
}

// SetNamespaces sets field value
func (o *XClusterConfigGetResp) SetNamespaces(v []XClusterNamespaceConfig) {
	o.Namespaces = v
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

// GetPitrConfigs returns the PitrConfigs field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetPitrConfigs() []PitrConfig {
	if o == nil || o.PitrConfigs == nil {
		var ret []PitrConfig
		return ret
	}
	return *o.PitrConfigs
}

// GetPitrConfigsOk returns a tuple with the PitrConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetPitrConfigsOk() (*[]PitrConfig, bool) {
	if o == nil || o.PitrConfigs == nil {
		return nil, false
	}
	return o.PitrConfigs, true
}

// HasPitrConfigs returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasPitrConfigs() bool {
	if o != nil && o.PitrConfigs != nil {
		return true
	}

	return false
}

// SetPitrConfigs gets a reference to the given []PitrConfig and assigns it to the PitrConfigs field.
func (o *XClusterConfigGetResp) SetPitrConfigs(v []PitrConfig) {
	o.PitrConfigs = &v
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

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetSecondary() bool {
	if o == nil || o.Secondary == nil {
		var ret bool
		return ret
	}
	return *o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetSecondaryOk() (*bool, bool) {
	if o == nil || o.Secondary == nil {
		return nil, false
	}
	return o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasSecondary() bool {
	if o != nil && o.Secondary != nil {
		return true
	}

	return false
}

// SetSecondary gets a reference to the given bool and assigns it to the Secondary field.
func (o *XClusterConfigGetResp) SetSecondary(v bool) {
	o.Secondary = &v
}

// GetSourceActive returns the SourceActive field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetSourceActive() bool {
	if o == nil || o.SourceActive == nil {
		var ret bool
		return ret
	}
	return *o.SourceActive
}

// GetSourceActiveOk returns a tuple with the SourceActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetSourceActiveOk() (*bool, bool) {
	if o == nil || o.SourceActive == nil {
		return nil, false
	}
	return o.SourceActive, true
}

// HasSourceActive returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasSourceActive() bool {
	if o != nil && o.SourceActive != nil {
		return true
	}

	return false
}

// SetSourceActive gets a reference to the given bool and assigns it to the SourceActive field.
func (o *XClusterConfigGetResp) SetSourceActive(v bool) {
	o.SourceActive = &v
}

// GetSourceUniverseState returns the SourceUniverseState field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetSourceUniverseState() string {
	if o == nil || o.SourceUniverseState == nil {
		var ret string
		return ret
	}
	return *o.SourceUniverseState
}

// GetSourceUniverseStateOk returns a tuple with the SourceUniverseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetSourceUniverseStateOk() (*string, bool) {
	if o == nil || o.SourceUniverseState == nil {
		return nil, false
	}
	return o.SourceUniverseState, true
}

// HasSourceUniverseState returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasSourceUniverseState() bool {
	if o != nil && o.SourceUniverseState != nil {
		return true
	}

	return false
}

// SetSourceUniverseState gets a reference to the given string and assigns it to the SourceUniverseState field.
func (o *XClusterConfigGetResp) SetSourceUniverseState(v string) {
	o.SourceUniverseState = &v
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

// GetTableDetails returns the TableDetails field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetTableDetails() []XClusterTableConfig {
	if o == nil || o.TableDetails == nil {
		var ret []XClusterTableConfig
		return ret
	}
	return *o.TableDetails
}

// GetTableDetailsOk returns a tuple with the TableDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetTableDetailsOk() (*[]XClusterTableConfig, bool) {
	if o == nil || o.TableDetails == nil {
		return nil, false
	}
	return o.TableDetails, true
}

// HasTableDetails returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasTableDetails() bool {
	if o != nil && o.TableDetails != nil {
		return true
	}

	return false
}

// SetTableDetails gets a reference to the given []XClusterTableConfig and assigns it to the TableDetails field.
func (o *XClusterConfigGetResp) SetTableDetails(v []XClusterTableConfig) {
	o.TableDetails = &v
}

// GetTableType returns the TableType field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetTableType() string {
	if o == nil || o.TableType == nil {
		var ret string
		return ret
	}
	return *o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetTableTypeOk() (*string, bool) {
	if o == nil || o.TableType == nil {
		return nil, false
	}
	return o.TableType, true
}

// HasTableType returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasTableType() bool {
	if o != nil && o.TableType != nil {
		return true
	}

	return false
}

// SetTableType gets a reference to the given string and assigns it to the TableType field.
func (o *XClusterConfigGetResp) SetTableType(v string) {
	o.TableType = &v
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

// GetTargetActive returns the TargetActive field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetTargetActive() bool {
	if o == nil || o.TargetActive == nil {
		var ret bool
		return ret
	}
	return *o.TargetActive
}

// GetTargetActiveOk returns a tuple with the TargetActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetTargetActiveOk() (*bool, bool) {
	if o == nil || o.TargetActive == nil {
		return nil, false
	}
	return o.TargetActive, true
}

// HasTargetActive returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasTargetActive() bool {
	if o != nil && o.TargetActive != nil {
		return true
	}

	return false
}

// SetTargetActive gets a reference to the given bool and assigns it to the TargetActive field.
func (o *XClusterConfigGetResp) SetTargetActive(v bool) {
	o.TargetActive = &v
}

// GetTargetUniverseState returns the TargetUniverseState field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetTargetUniverseState() string {
	if o == nil || o.TargetUniverseState == nil {
		var ret string
		return ret
	}
	return *o.TargetUniverseState
}

// GetTargetUniverseStateOk returns a tuple with the TargetUniverseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetTargetUniverseStateOk() (*string, bool) {
	if o == nil || o.TargetUniverseState == nil {
		return nil, false
	}
	return o.TargetUniverseState, true
}

// HasTargetUniverseState returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasTargetUniverseState() bool {
	if o != nil && o.TargetUniverseState != nil {
		return true
	}

	return false
}

// SetTargetUniverseState gets a reference to the given string and assigns it to the TargetUniverseState field.
func (o *XClusterConfigGetResp) SetTargetUniverseState(v string) {
	o.TargetUniverseState = &v
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

// GetType returns the Type field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *XClusterConfigGetResp) SetType(v string) {
	o.Type = &v
}

// GetUsedForDr returns the UsedForDr field value if set, zero value otherwise.
func (o *XClusterConfigGetResp) GetUsedForDr() bool {
	if o == nil || o.UsedForDr == nil {
		var ret bool
		return ret
	}
	return *o.UsedForDr
}

// GetUsedForDrOk returns a tuple with the UsedForDr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigGetResp) GetUsedForDrOk() (*bool, bool) {
	if o == nil || o.UsedForDr == nil {
		return nil, false
	}
	return o.UsedForDr, true
}

// HasUsedForDr returns a boolean if a field has been set.
func (o *XClusterConfigGetResp) HasUsedForDr() bool {
	if o != nil && o.UsedForDr != nil {
		return true
	}

	return false
}

// SetUsedForDr gets a reference to the given bool and assigns it to the UsedForDr field.
func (o *XClusterConfigGetResp) SetUsedForDr(v bool) {
	o.UsedForDr = &v
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
	if o.Dbs != nil {
		toSerialize["dbs"] = o.Dbs
	}
	if o.Imported != nil {
		toSerialize["imported"] = o.Imported
	}
	if o.KeyspacePending != nil {
		toSerialize["keyspacePending"] = o.KeyspacePending
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
	if o.NamespaceDetails != nil {
		toSerialize["namespaceDetails"] = o.NamespaceDetails
	}
	if true {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Paused != nil {
		toSerialize["paused"] = o.Paused
	}
	if o.PitrConfigs != nil {
		toSerialize["pitrConfigs"] = o.PitrConfigs
	}
	if o.ReplicationGroupName != nil {
		toSerialize["replicationGroupName"] = o.ReplicationGroupName
	}
	if o.Secondary != nil {
		toSerialize["secondary"] = o.Secondary
	}
	if o.SourceActive != nil {
		toSerialize["sourceActive"] = o.SourceActive
	}
	if o.SourceUniverseState != nil {
		toSerialize["sourceUniverseState"] = o.SourceUniverseState
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
	if o.TargetUniverseState != nil {
		toSerialize["targetUniverseState"] = o.TargetUniverseState
	}
	if o.TargetUniverseUUID != nil {
		toSerialize["targetUniverseUUID"] = o.TargetUniverseUUID
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UsedForDr != nil {
		toSerialize["usedForDr"] = o.UsedForDr
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


