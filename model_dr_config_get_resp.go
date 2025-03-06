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

// DrConfigGetResp disaster recovery get response
type DrConfigGetResp struct {
	BootstrapParams *RestartBootstrapParams `json:"bootstrapParams,omitempty"`
	// Create time of the DR config
	CreateTime *time.Time `json:"createTime,omitempty"`
	// WARNING: This is a preview API that could change. List of db details in replication
	DbDetails *[]XClusterNamespaceConfig `json:"dbDetails,omitempty"`
	// WARNING: This is a preview API that could change. List of db ids in replication
	Dbs *[]string `json:"dbs,omitempty"`
	// Whether the dr replica universe is active
	DrReplicaUniverseActive *bool `json:"drReplicaUniverseActive,omitempty"`
	// WARNING: This is a preview API that could change. The replication status of the dr replica universe.
	DrReplicaUniverseState *string `json:"drReplicaUniverseState,omitempty"`
	// DR Replica Universe UUID
	DrReplicaUniverseUuid *string `json:"drReplicaUniverseUuid,omitempty"`
	// WARNING: This is a preview API that could change. The keyspace name that the current task is working on.
	KeyspacePending *string `json:"keyspacePending,omitempty"`
	// Last modify time of the DR config
	ModifyTime *time.Time `json:"modifyTime,omitempty"`
	// Disaster recovery config name
	Name *string `json:"name,omitempty"`
	// Whether the underlying xCluster config is paused
	Paused *bool `json:"paused,omitempty"`
	// WARNING: This is a preview API that could change. The list of PITR configs used for the underlying txn xCluster config
	PitrConfigs *[]PitrConfig `json:"pitrConfigs,omitempty"`
	// WARNING: This is a preview API that could change. PITR Retention Period in seconds
	PitrRetentionPeriodSec *int64 `json:"pitrRetentionPeriodSec,omitempty"`
	// WARNING: This is a preview API that could change. PITR Retention Period in seconds
	PitrSnapshotIntervalSec *int64 `json:"pitrSnapshotIntervalSec,omitempty"`
	// Whether the primary universe is active
	PrimaryUniverseActive *bool `json:"primaryUniverseActive,omitempty"`
	// WARNING: This is a preview API that could change. The replication status of the primary universe.
	PrimaryUniverseState *string `json:"primaryUniverseState,omitempty"`
	// Primary Universe UUID
	PrimaryUniverseUuid *string `json:"primaryUniverseUuid,omitempty"`
	// Replication group name in the dr replica universe cluster config
	ReplicationGroupName *string `json:"replicationGroupName,omitempty"`
	// The state of the DR config
	State *string `json:"state,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
	// Details for each table in replication
	TableDetails *[]XClusterTableConfig `json:"tableDetails,omitempty"`
	// The type of tables that are being replicated
	TableType *string `json:"tableType,omitempty"`
	// List of table ids in replication
	Tables *[]string `json:"tables,omitempty"`
	// Whether the config is basic, txn, or db scoped xCluster
	Type *string `json:"type,omitempty"`
	// DR config UUID
	Uuid *string `json:"uuid,omitempty"`
	Webhooks []GetWebhookResponse `json:"webhooks"`
	// UUID of the underlying xCluster config that is managing the replication
	XclusterConfigUuid *string `json:"xclusterConfigUuid,omitempty"`
	// YbaApi Internal. The list of xCluster configs' uuids that belong to this dr config
	XclusterConfigsUuid *[]string `json:"xclusterConfigsUuid,omitempty"`
}

// NewDrConfigGetResp instantiates a new DrConfigGetResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrConfigGetResp(webhooks []GetWebhookResponse) *DrConfigGetResp {
	this := DrConfigGetResp{}
	this.Webhooks = webhooks
	return &this
}

// NewDrConfigGetRespWithDefaults instantiates a new DrConfigGetResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrConfigGetRespWithDefaults() *DrConfigGetResp {
	this := DrConfigGetResp{}
	return &this
}

// GetBootstrapParams returns the BootstrapParams field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetBootstrapParams() RestartBootstrapParams {
	if o == nil || o.BootstrapParams == nil {
		var ret RestartBootstrapParams
		return ret
	}
	return *o.BootstrapParams
}

// GetBootstrapParamsOk returns a tuple with the BootstrapParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetBootstrapParamsOk() (*RestartBootstrapParams, bool) {
	if o == nil || o.BootstrapParams == nil {
		return nil, false
	}
	return o.BootstrapParams, true
}

// HasBootstrapParams returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasBootstrapParams() bool {
	if o != nil && o.BootstrapParams != nil {
		return true
	}

	return false
}

// SetBootstrapParams gets a reference to the given RestartBootstrapParams and assigns it to the BootstrapParams field.
func (o *DrConfigGetResp) SetBootstrapParams(v RestartBootstrapParams) {
	o.BootstrapParams = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *DrConfigGetResp) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetDbDetails returns the DbDetails field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetDbDetails() []XClusterNamespaceConfig {
	if o == nil || o.DbDetails == nil {
		var ret []XClusterNamespaceConfig
		return ret
	}
	return *o.DbDetails
}

// GetDbDetailsOk returns a tuple with the DbDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetDbDetailsOk() (*[]XClusterNamespaceConfig, bool) {
	if o == nil || o.DbDetails == nil {
		return nil, false
	}
	return o.DbDetails, true
}

// HasDbDetails returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasDbDetails() bool {
	if o != nil && o.DbDetails != nil {
		return true
	}

	return false
}

// SetDbDetails gets a reference to the given []XClusterNamespaceConfig and assigns it to the DbDetails field.
func (o *DrConfigGetResp) SetDbDetails(v []XClusterNamespaceConfig) {
	o.DbDetails = &v
}

// GetDbs returns the Dbs field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetDbs() []string {
	if o == nil || o.Dbs == nil {
		var ret []string
		return ret
	}
	return *o.Dbs
}

// GetDbsOk returns a tuple with the Dbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetDbsOk() (*[]string, bool) {
	if o == nil || o.Dbs == nil {
		return nil, false
	}
	return o.Dbs, true
}

// HasDbs returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasDbs() bool {
	if o != nil && o.Dbs != nil {
		return true
	}

	return false
}

// SetDbs gets a reference to the given []string and assigns it to the Dbs field.
func (o *DrConfigGetResp) SetDbs(v []string) {
	o.Dbs = &v
}

// GetDrReplicaUniverseActive returns the DrReplicaUniverseActive field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetDrReplicaUniverseActive() bool {
	if o == nil || o.DrReplicaUniverseActive == nil {
		var ret bool
		return ret
	}
	return *o.DrReplicaUniverseActive
}

// GetDrReplicaUniverseActiveOk returns a tuple with the DrReplicaUniverseActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetDrReplicaUniverseActiveOk() (*bool, bool) {
	if o == nil || o.DrReplicaUniverseActive == nil {
		return nil, false
	}
	return o.DrReplicaUniverseActive, true
}

// HasDrReplicaUniverseActive returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasDrReplicaUniverseActive() bool {
	if o != nil && o.DrReplicaUniverseActive != nil {
		return true
	}

	return false
}

// SetDrReplicaUniverseActive gets a reference to the given bool and assigns it to the DrReplicaUniverseActive field.
func (o *DrConfigGetResp) SetDrReplicaUniverseActive(v bool) {
	o.DrReplicaUniverseActive = &v
}

// GetDrReplicaUniverseState returns the DrReplicaUniverseState field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetDrReplicaUniverseState() string {
	if o == nil || o.DrReplicaUniverseState == nil {
		var ret string
		return ret
	}
	return *o.DrReplicaUniverseState
}

// GetDrReplicaUniverseStateOk returns a tuple with the DrReplicaUniverseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetDrReplicaUniverseStateOk() (*string, bool) {
	if o == nil || o.DrReplicaUniverseState == nil {
		return nil, false
	}
	return o.DrReplicaUniverseState, true
}

// HasDrReplicaUniverseState returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasDrReplicaUniverseState() bool {
	if o != nil && o.DrReplicaUniverseState != nil {
		return true
	}

	return false
}

// SetDrReplicaUniverseState gets a reference to the given string and assigns it to the DrReplicaUniverseState field.
func (o *DrConfigGetResp) SetDrReplicaUniverseState(v string) {
	o.DrReplicaUniverseState = &v
}

// GetDrReplicaUniverseUuid returns the DrReplicaUniverseUuid field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetDrReplicaUniverseUuid() string {
	if o == nil || o.DrReplicaUniverseUuid == nil {
		var ret string
		return ret
	}
	return *o.DrReplicaUniverseUuid
}

// GetDrReplicaUniverseUuidOk returns a tuple with the DrReplicaUniverseUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetDrReplicaUniverseUuidOk() (*string, bool) {
	if o == nil || o.DrReplicaUniverseUuid == nil {
		return nil, false
	}
	return o.DrReplicaUniverseUuid, true
}

// HasDrReplicaUniverseUuid returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasDrReplicaUniverseUuid() bool {
	if o != nil && o.DrReplicaUniverseUuid != nil {
		return true
	}

	return false
}

// SetDrReplicaUniverseUuid gets a reference to the given string and assigns it to the DrReplicaUniverseUuid field.
func (o *DrConfigGetResp) SetDrReplicaUniverseUuid(v string) {
	o.DrReplicaUniverseUuid = &v
}

// GetKeyspacePending returns the KeyspacePending field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetKeyspacePending() string {
	if o == nil || o.KeyspacePending == nil {
		var ret string
		return ret
	}
	return *o.KeyspacePending
}

// GetKeyspacePendingOk returns a tuple with the KeyspacePending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetKeyspacePendingOk() (*string, bool) {
	if o == nil || o.KeyspacePending == nil {
		return nil, false
	}
	return o.KeyspacePending, true
}

// HasKeyspacePending returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasKeyspacePending() bool {
	if o != nil && o.KeyspacePending != nil {
		return true
	}

	return false
}

// SetKeyspacePending gets a reference to the given string and assigns it to the KeyspacePending field.
func (o *DrConfigGetResp) SetKeyspacePending(v string) {
	o.KeyspacePending = &v
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetModifyTime() time.Time {
	if o == nil || o.ModifyTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetModifyTimeOk() (*time.Time, bool) {
	if o == nil || o.ModifyTime == nil {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasModifyTime() bool {
	if o != nil && o.ModifyTime != nil {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given time.Time and assigns it to the ModifyTime field.
func (o *DrConfigGetResp) SetModifyTime(v time.Time) {
	o.ModifyTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DrConfigGetResp) SetName(v string) {
	o.Name = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetPaused() bool {
	if o == nil || o.Paused == nil {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetPausedOk() (*bool, bool) {
	if o == nil || o.Paused == nil {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasPaused() bool {
	if o != nil && o.Paused != nil {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *DrConfigGetResp) SetPaused(v bool) {
	o.Paused = &v
}

// GetPitrConfigs returns the PitrConfigs field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetPitrConfigs() []PitrConfig {
	if o == nil || o.PitrConfigs == nil {
		var ret []PitrConfig
		return ret
	}
	return *o.PitrConfigs
}

// GetPitrConfigsOk returns a tuple with the PitrConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetPitrConfigsOk() (*[]PitrConfig, bool) {
	if o == nil || o.PitrConfigs == nil {
		return nil, false
	}
	return o.PitrConfigs, true
}

// HasPitrConfigs returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasPitrConfigs() bool {
	if o != nil && o.PitrConfigs != nil {
		return true
	}

	return false
}

// SetPitrConfigs gets a reference to the given []PitrConfig and assigns it to the PitrConfigs field.
func (o *DrConfigGetResp) SetPitrConfigs(v []PitrConfig) {
	o.PitrConfigs = &v
}

// GetPitrRetentionPeriodSec returns the PitrRetentionPeriodSec field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetPitrRetentionPeriodSec() int64 {
	if o == nil || o.PitrRetentionPeriodSec == nil {
		var ret int64
		return ret
	}
	return *o.PitrRetentionPeriodSec
}

// GetPitrRetentionPeriodSecOk returns a tuple with the PitrRetentionPeriodSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetPitrRetentionPeriodSecOk() (*int64, bool) {
	if o == nil || o.PitrRetentionPeriodSec == nil {
		return nil, false
	}
	return o.PitrRetentionPeriodSec, true
}

// HasPitrRetentionPeriodSec returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasPitrRetentionPeriodSec() bool {
	if o != nil && o.PitrRetentionPeriodSec != nil {
		return true
	}

	return false
}

// SetPitrRetentionPeriodSec gets a reference to the given int64 and assigns it to the PitrRetentionPeriodSec field.
func (o *DrConfigGetResp) SetPitrRetentionPeriodSec(v int64) {
	o.PitrRetentionPeriodSec = &v
}

// GetPitrSnapshotIntervalSec returns the PitrSnapshotIntervalSec field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetPitrSnapshotIntervalSec() int64 {
	if o == nil || o.PitrSnapshotIntervalSec == nil {
		var ret int64
		return ret
	}
	return *o.PitrSnapshotIntervalSec
}

// GetPitrSnapshotIntervalSecOk returns a tuple with the PitrSnapshotIntervalSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetPitrSnapshotIntervalSecOk() (*int64, bool) {
	if o == nil || o.PitrSnapshotIntervalSec == nil {
		return nil, false
	}
	return o.PitrSnapshotIntervalSec, true
}

// HasPitrSnapshotIntervalSec returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasPitrSnapshotIntervalSec() bool {
	if o != nil && o.PitrSnapshotIntervalSec != nil {
		return true
	}

	return false
}

// SetPitrSnapshotIntervalSec gets a reference to the given int64 and assigns it to the PitrSnapshotIntervalSec field.
func (o *DrConfigGetResp) SetPitrSnapshotIntervalSec(v int64) {
	o.PitrSnapshotIntervalSec = &v
}

// GetPrimaryUniverseActive returns the PrimaryUniverseActive field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetPrimaryUniverseActive() bool {
	if o == nil || o.PrimaryUniverseActive == nil {
		var ret bool
		return ret
	}
	return *o.PrimaryUniverseActive
}

// GetPrimaryUniverseActiveOk returns a tuple with the PrimaryUniverseActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetPrimaryUniverseActiveOk() (*bool, bool) {
	if o == nil || o.PrimaryUniverseActive == nil {
		return nil, false
	}
	return o.PrimaryUniverseActive, true
}

// HasPrimaryUniverseActive returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasPrimaryUniverseActive() bool {
	if o != nil && o.PrimaryUniverseActive != nil {
		return true
	}

	return false
}

// SetPrimaryUniverseActive gets a reference to the given bool and assigns it to the PrimaryUniverseActive field.
func (o *DrConfigGetResp) SetPrimaryUniverseActive(v bool) {
	o.PrimaryUniverseActive = &v
}

// GetPrimaryUniverseState returns the PrimaryUniverseState field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetPrimaryUniverseState() string {
	if o == nil || o.PrimaryUniverseState == nil {
		var ret string
		return ret
	}
	return *o.PrimaryUniverseState
}

// GetPrimaryUniverseStateOk returns a tuple with the PrimaryUniverseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetPrimaryUniverseStateOk() (*string, bool) {
	if o == nil || o.PrimaryUniverseState == nil {
		return nil, false
	}
	return o.PrimaryUniverseState, true
}

// HasPrimaryUniverseState returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasPrimaryUniverseState() bool {
	if o != nil && o.PrimaryUniverseState != nil {
		return true
	}

	return false
}

// SetPrimaryUniverseState gets a reference to the given string and assigns it to the PrimaryUniverseState field.
func (o *DrConfigGetResp) SetPrimaryUniverseState(v string) {
	o.PrimaryUniverseState = &v
}

// GetPrimaryUniverseUuid returns the PrimaryUniverseUuid field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetPrimaryUniverseUuid() string {
	if o == nil || o.PrimaryUniverseUuid == nil {
		var ret string
		return ret
	}
	return *o.PrimaryUniverseUuid
}

// GetPrimaryUniverseUuidOk returns a tuple with the PrimaryUniverseUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetPrimaryUniverseUuidOk() (*string, bool) {
	if o == nil || o.PrimaryUniverseUuid == nil {
		return nil, false
	}
	return o.PrimaryUniverseUuid, true
}

// HasPrimaryUniverseUuid returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasPrimaryUniverseUuid() bool {
	if o != nil && o.PrimaryUniverseUuid != nil {
		return true
	}

	return false
}

// SetPrimaryUniverseUuid gets a reference to the given string and assigns it to the PrimaryUniverseUuid field.
func (o *DrConfigGetResp) SetPrimaryUniverseUuid(v string) {
	o.PrimaryUniverseUuid = &v
}

// GetReplicationGroupName returns the ReplicationGroupName field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetReplicationGroupName() string {
	if o == nil || o.ReplicationGroupName == nil {
		var ret string
		return ret
	}
	return *o.ReplicationGroupName
}

// GetReplicationGroupNameOk returns a tuple with the ReplicationGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetReplicationGroupNameOk() (*string, bool) {
	if o == nil || o.ReplicationGroupName == nil {
		return nil, false
	}
	return o.ReplicationGroupName, true
}

// HasReplicationGroupName returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasReplicationGroupName() bool {
	if o != nil && o.ReplicationGroupName != nil {
		return true
	}

	return false
}

// SetReplicationGroupName gets a reference to the given string and assigns it to the ReplicationGroupName field.
func (o *DrConfigGetResp) SetReplicationGroupName(v string) {
	o.ReplicationGroupName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DrConfigGetResp) SetState(v string) {
	o.State = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DrConfigGetResp) SetStatus(v string) {
	o.Status = &v
}

// GetTableDetails returns the TableDetails field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetTableDetails() []XClusterTableConfig {
	if o == nil || o.TableDetails == nil {
		var ret []XClusterTableConfig
		return ret
	}
	return *o.TableDetails
}

// GetTableDetailsOk returns a tuple with the TableDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetTableDetailsOk() (*[]XClusterTableConfig, bool) {
	if o == nil || o.TableDetails == nil {
		return nil, false
	}
	return o.TableDetails, true
}

// HasTableDetails returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasTableDetails() bool {
	if o != nil && o.TableDetails != nil {
		return true
	}

	return false
}

// SetTableDetails gets a reference to the given []XClusterTableConfig and assigns it to the TableDetails field.
func (o *DrConfigGetResp) SetTableDetails(v []XClusterTableConfig) {
	o.TableDetails = &v
}

// GetTableType returns the TableType field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetTableType() string {
	if o == nil || o.TableType == nil {
		var ret string
		return ret
	}
	return *o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetTableTypeOk() (*string, bool) {
	if o == nil || o.TableType == nil {
		return nil, false
	}
	return o.TableType, true
}

// HasTableType returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasTableType() bool {
	if o != nil && o.TableType != nil {
		return true
	}

	return false
}

// SetTableType gets a reference to the given string and assigns it to the TableType field.
func (o *DrConfigGetResp) SetTableType(v string) {
	o.TableType = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return *o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *DrConfigGetResp) SetTables(v []string) {
	o.Tables = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DrConfigGetResp) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DrConfigGetResp) SetUuid(v string) {
	o.Uuid = &v
}

// GetWebhooks returns the Webhooks field value
func (o *DrConfigGetResp) GetWebhooks() []GetWebhookResponse {
	if o == nil {
		var ret []GetWebhookResponse
		return ret
	}

	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetWebhooksOk() (*[]GetWebhookResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Webhooks, true
}

// SetWebhooks sets field value
func (o *DrConfigGetResp) SetWebhooks(v []GetWebhookResponse) {
	o.Webhooks = v
}

// GetXclusterConfigUuid returns the XclusterConfigUuid field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetXclusterConfigUuid() string {
	if o == nil || o.XclusterConfigUuid == nil {
		var ret string
		return ret
	}
	return *o.XclusterConfigUuid
}

// GetXclusterConfigUuidOk returns a tuple with the XclusterConfigUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetXclusterConfigUuidOk() (*string, bool) {
	if o == nil || o.XclusterConfigUuid == nil {
		return nil, false
	}
	return o.XclusterConfigUuid, true
}

// HasXclusterConfigUuid returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasXclusterConfigUuid() bool {
	if o != nil && o.XclusterConfigUuid != nil {
		return true
	}

	return false
}

// SetXclusterConfigUuid gets a reference to the given string and assigns it to the XclusterConfigUuid field.
func (o *DrConfigGetResp) SetXclusterConfigUuid(v string) {
	o.XclusterConfigUuid = &v
}

// GetXclusterConfigsUuid returns the XclusterConfigsUuid field value if set, zero value otherwise.
func (o *DrConfigGetResp) GetXclusterConfigsUuid() []string {
	if o == nil || o.XclusterConfigsUuid == nil {
		var ret []string
		return ret
	}
	return *o.XclusterConfigsUuid
}

// GetXclusterConfigsUuidOk returns a tuple with the XclusterConfigsUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigGetResp) GetXclusterConfigsUuidOk() (*[]string, bool) {
	if o == nil || o.XclusterConfigsUuid == nil {
		return nil, false
	}
	return o.XclusterConfigsUuid, true
}

// HasXclusterConfigsUuid returns a boolean if a field has been set.
func (o *DrConfigGetResp) HasXclusterConfigsUuid() bool {
	if o != nil && o.XclusterConfigsUuid != nil {
		return true
	}

	return false
}

// SetXclusterConfigsUuid gets a reference to the given []string and assigns it to the XclusterConfigsUuid field.
func (o *DrConfigGetResp) SetXclusterConfigsUuid(v []string) {
	o.XclusterConfigsUuid = &v
}

func (o DrConfigGetResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootstrapParams != nil {
		toSerialize["bootstrapParams"] = o.BootstrapParams
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.DbDetails != nil {
		toSerialize["dbDetails"] = o.DbDetails
	}
	if o.Dbs != nil {
		toSerialize["dbs"] = o.Dbs
	}
	if o.DrReplicaUniverseActive != nil {
		toSerialize["drReplicaUniverseActive"] = o.DrReplicaUniverseActive
	}
	if o.DrReplicaUniverseState != nil {
		toSerialize["drReplicaUniverseState"] = o.DrReplicaUniverseState
	}
	if o.DrReplicaUniverseUuid != nil {
		toSerialize["drReplicaUniverseUuid"] = o.DrReplicaUniverseUuid
	}
	if o.KeyspacePending != nil {
		toSerialize["keyspacePending"] = o.KeyspacePending
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
	if o.PitrConfigs != nil {
		toSerialize["pitrConfigs"] = o.PitrConfigs
	}
	if o.PitrRetentionPeriodSec != nil {
		toSerialize["pitrRetentionPeriodSec"] = o.PitrRetentionPeriodSec
	}
	if o.PitrSnapshotIntervalSec != nil {
		toSerialize["pitrSnapshotIntervalSec"] = o.PitrSnapshotIntervalSec
	}
	if o.PrimaryUniverseActive != nil {
		toSerialize["primaryUniverseActive"] = o.PrimaryUniverseActive
	}
	if o.PrimaryUniverseState != nil {
		toSerialize["primaryUniverseState"] = o.PrimaryUniverseState
	}
	if o.PrimaryUniverseUuid != nil {
		toSerialize["primaryUniverseUuid"] = o.PrimaryUniverseUuid
	}
	if o.ReplicationGroupName != nil {
		toSerialize["replicationGroupName"] = o.ReplicationGroupName
	}
	if o.State != nil {
		toSerialize["state"] = o.State
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["webhooks"] = o.Webhooks
	}
	if o.XclusterConfigUuid != nil {
		toSerialize["xclusterConfigUuid"] = o.XclusterConfigUuid
	}
	if o.XclusterConfigsUuid != nil {
		toSerialize["xclusterConfigsUuid"] = o.XclusterConfigsUuid
	}
	return json.Marshal(toSerialize)
}

type NullableDrConfigGetResp struct {
	value *DrConfigGetResp
	isSet bool
}

func (v NullableDrConfigGetResp) Get() *DrConfigGetResp {
	return v.value
}

func (v *NullableDrConfigGetResp) Set(val *DrConfigGetResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDrConfigGetResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDrConfigGetResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrConfigGetResp(val *DrConfigGetResp) *NullableDrConfigGetResp {
	return &NullableDrConfigGetResp{value: val, isSet: true}
}

func (v NullableDrConfigGetResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrConfigGetResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


