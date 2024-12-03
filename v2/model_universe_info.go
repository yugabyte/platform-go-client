/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"time"
)

// UniverseInfo These are read-only system generated properties of a Universe. Returned as part of a Universe resource.
type UniverseInfo struct {
	// UUID of the Universe
	UniverseUuid *string `json:"universe_uuid,omitempty"`
	// Universe version
	Version *int32 `json:"version,omitempty"`
	// Universe creation date
	CreationDate *time.Time `json:"creation_date,omitempty"`
	CreatingUser *User `json:"creating_user,omitempty"`
	// CPU Arch of DB nodes.
	Arch *string `json:"arch,omitempty"`
	// DNS name
	DnsName *string `json:"dns_name,omitempty"`
	// YBC Software version installed in DB nodes of this Universe
	YbcSoftwareVersion *string `json:"ybc_software_version,omitempty"`
	// YBAnywhere host name that is managing this Universe
	YbaUrl *string `json:"yba_url,omitempty"`
	// A globally unique name generated as a combination of the customer id and the universe name. This is used as the prefix of node names in the universe. Can be configured at the time of universe creation.
	NodePrefix *string `json:"node_prefix,omitempty"`
	EncryptionAtRestInfo *EncryptionAtRestInfo `json:"encryption_at_rest_info,omitempty"`
	// Whether a create/edit/destroy intent on the universe is currently running.
	UpdateInProgress *bool `json:"update_in_progress,omitempty"`
	// Type of task which set updateInProgress flag.
	UpdatingTask *string `json:"updating_task,omitempty"`
	// UUID of task which set updateInProgress flag.
	UpdatingTaskUuid *string `json:"updating_task_uuid,omitempty"`
	// Whether the latest operation on this universe has successfully completed. Is updated for each operation on the universe.
	UpdateSucceeded *bool `json:"update_succeeded,omitempty"`
	// UUID of the last task that will be retried on this Universe
	PreviousTaskUuid *string `json:"previous_task_uuid,omitempty"`
	// Whether the universe is in the paused state
	UniversePaused *bool `json:"universe_paused,omitempty"`
	// UUID of last failed task that applied modification to cluster state
	PlacementModificationTaskUuid *string `json:"placement_modification_task_uuid,omitempty"`
	// The state of the last YugabyteDB software upgrade operation on this universe
	SoftwareUpgradeState *string `json:"software_upgrade_state,omitempty"`
	// Whether a rollback of the last YugabyteDB upgrade operation is allowed
	IsSoftwareRollbackAllowed *bool `json:"is_software_rollback_allowed,omitempty"`
	PreviousYbSoftwareDetails *YbSoftwareDetails `json:"previous_yb_software_details,omitempty"`
	AllowedTasksOnFailure *AllowedTasksOnFailure `json:"allowed_tasks_on_failure,omitempty"`
	// Set to true if nodes of this Universe can be resized without a full move
	NodesResizeAvailable *bool `json:"nodes_resize_available,omitempty"`
	// Whether this Universe is created and controlled by the Kubernetes Operator
	IsKubernetesOperatorControlled *bool `json:"is_kubernetes_operator_controlled,omitempty"`
	// Whether OpenTelemetry Collector is enabled for universe
	OtelCollectorEnabled *bool `json:"otel_collector_enabled,omitempty"`
	XClusterInfo *XClusterInfo `json:"x_cluster_info,omitempty"`
	RollMaxBatchSize *RollMaxBatchSize `json:"roll_max_batch_size,omitempty"`
	Clusters *[]ClusterInfo `json:"clusters,omitempty"`
	// Node details
	NodeDetailsSet *[]NodeDetails `json:"node_details_set,omitempty"`
	// UUIDs of DR configs where this universe is the source (primary)
	DrConfigUuidsAsSource *[]string `json:"dr_config_uuids_as_source,omitempty"`
	// UUIDs of DR configs where this universe is the target (secondary)
	DrConfigUuidsAsTarget *[]string `json:"dr_config_uuids_as_target,omitempty"`
	Resources *UniverseResourceDetails `json:"resources,omitempty"`
	// Sample command
	SampleAppCommandTxt *string `json:"sample_app_command_txt,omitempty"`
}

// NewUniverseInfo instantiates a new UniverseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseInfo() *UniverseInfo {
	this := UniverseInfo{}
	return &this
}

// NewUniverseInfoWithDefaults instantiates a new UniverseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseInfoWithDefaults() *UniverseInfo {
	this := UniverseInfo{}
	return &this
}

// GetUniverseUuid returns the UniverseUuid field value if set, zero value otherwise.
func (o *UniverseInfo) GetUniverseUuid() string {
	if o == nil || o.UniverseUuid == nil {
		var ret string
		return ret
	}
	return *o.UniverseUuid
}

// GetUniverseUuidOk returns a tuple with the UniverseUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetUniverseUuidOk() (*string, bool) {
	if o == nil || o.UniverseUuid == nil {
		return nil, false
	}
	return o.UniverseUuid, true
}

// HasUniverseUuid returns a boolean if a field has been set.
func (o *UniverseInfo) HasUniverseUuid() bool {
	if o != nil && o.UniverseUuid != nil {
		return true
	}

	return false
}

// SetUniverseUuid gets a reference to the given string and assigns it to the UniverseUuid field.
func (o *UniverseInfo) SetUniverseUuid(v string) {
	o.UniverseUuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UniverseInfo) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UniverseInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *UniverseInfo) SetVersion(v int32) {
	o.Version = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *UniverseInfo) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *UniverseInfo) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *UniverseInfo) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetCreatingUser returns the CreatingUser field value if set, zero value otherwise.
func (o *UniverseInfo) GetCreatingUser() User {
	if o == nil || o.CreatingUser == nil {
		var ret User
		return ret
	}
	return *o.CreatingUser
}

// GetCreatingUserOk returns a tuple with the CreatingUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetCreatingUserOk() (*User, bool) {
	if o == nil || o.CreatingUser == nil {
		return nil, false
	}
	return o.CreatingUser, true
}

// HasCreatingUser returns a boolean if a field has been set.
func (o *UniverseInfo) HasCreatingUser() bool {
	if o != nil && o.CreatingUser != nil {
		return true
	}

	return false
}

// SetCreatingUser gets a reference to the given User and assigns it to the CreatingUser field.
func (o *UniverseInfo) SetCreatingUser(v User) {
	o.CreatingUser = &v
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *UniverseInfo) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *UniverseInfo) HasArch() bool {
	if o != nil && o.Arch != nil {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *UniverseInfo) SetArch(v string) {
	o.Arch = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *UniverseInfo) GetDnsName() string {
	if o == nil || o.DnsName == nil {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetDnsNameOk() (*string, bool) {
	if o == nil || o.DnsName == nil {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *UniverseInfo) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *UniverseInfo) SetDnsName(v string) {
	o.DnsName = &v
}

// GetYbcSoftwareVersion returns the YbcSoftwareVersion field value if set, zero value otherwise.
func (o *UniverseInfo) GetYbcSoftwareVersion() string {
	if o == nil || o.YbcSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbcSoftwareVersion
}

// GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetYbcSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbcSoftwareVersion == nil {
		return nil, false
	}
	return o.YbcSoftwareVersion, true
}

// HasYbcSoftwareVersion returns a boolean if a field has been set.
func (o *UniverseInfo) HasYbcSoftwareVersion() bool {
	if o != nil && o.YbcSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbcSoftwareVersion gets a reference to the given string and assigns it to the YbcSoftwareVersion field.
func (o *UniverseInfo) SetYbcSoftwareVersion(v string) {
	o.YbcSoftwareVersion = &v
}

// GetYbaUrl returns the YbaUrl field value if set, zero value otherwise.
func (o *UniverseInfo) GetYbaUrl() string {
	if o == nil || o.YbaUrl == nil {
		var ret string
		return ret
	}
	return *o.YbaUrl
}

// GetYbaUrlOk returns a tuple with the YbaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetYbaUrlOk() (*string, bool) {
	if o == nil || o.YbaUrl == nil {
		return nil, false
	}
	return o.YbaUrl, true
}

// HasYbaUrl returns a boolean if a field has been set.
func (o *UniverseInfo) HasYbaUrl() bool {
	if o != nil && o.YbaUrl != nil {
		return true
	}

	return false
}

// SetYbaUrl gets a reference to the given string and assigns it to the YbaUrl field.
func (o *UniverseInfo) SetYbaUrl(v string) {
	o.YbaUrl = &v
}

// GetNodePrefix returns the NodePrefix field value if set, zero value otherwise.
func (o *UniverseInfo) GetNodePrefix() string {
	if o == nil || o.NodePrefix == nil {
		var ret string
		return ret
	}
	return *o.NodePrefix
}

// GetNodePrefixOk returns a tuple with the NodePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetNodePrefixOk() (*string, bool) {
	if o == nil || o.NodePrefix == nil {
		return nil, false
	}
	return o.NodePrefix, true
}

// HasNodePrefix returns a boolean if a field has been set.
func (o *UniverseInfo) HasNodePrefix() bool {
	if o != nil && o.NodePrefix != nil {
		return true
	}

	return false
}

// SetNodePrefix gets a reference to the given string and assigns it to the NodePrefix field.
func (o *UniverseInfo) SetNodePrefix(v string) {
	o.NodePrefix = &v
}

// GetEncryptionAtRestInfo returns the EncryptionAtRestInfo field value if set, zero value otherwise.
func (o *UniverseInfo) GetEncryptionAtRestInfo() EncryptionAtRestInfo {
	if o == nil || o.EncryptionAtRestInfo == nil {
		var ret EncryptionAtRestInfo
		return ret
	}
	return *o.EncryptionAtRestInfo
}

// GetEncryptionAtRestInfoOk returns a tuple with the EncryptionAtRestInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetEncryptionAtRestInfoOk() (*EncryptionAtRestInfo, bool) {
	if o == nil || o.EncryptionAtRestInfo == nil {
		return nil, false
	}
	return o.EncryptionAtRestInfo, true
}

// HasEncryptionAtRestInfo returns a boolean if a field has been set.
func (o *UniverseInfo) HasEncryptionAtRestInfo() bool {
	if o != nil && o.EncryptionAtRestInfo != nil {
		return true
	}

	return false
}

// SetEncryptionAtRestInfo gets a reference to the given EncryptionAtRestInfo and assigns it to the EncryptionAtRestInfo field.
func (o *UniverseInfo) SetEncryptionAtRestInfo(v EncryptionAtRestInfo) {
	o.EncryptionAtRestInfo = &v
}

// GetUpdateInProgress returns the UpdateInProgress field value if set, zero value otherwise.
func (o *UniverseInfo) GetUpdateInProgress() bool {
	if o == nil || o.UpdateInProgress == nil {
		var ret bool
		return ret
	}
	return *o.UpdateInProgress
}

// GetUpdateInProgressOk returns a tuple with the UpdateInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetUpdateInProgressOk() (*bool, bool) {
	if o == nil || o.UpdateInProgress == nil {
		return nil, false
	}
	return o.UpdateInProgress, true
}

// HasUpdateInProgress returns a boolean if a field has been set.
func (o *UniverseInfo) HasUpdateInProgress() bool {
	if o != nil && o.UpdateInProgress != nil {
		return true
	}

	return false
}

// SetUpdateInProgress gets a reference to the given bool and assigns it to the UpdateInProgress field.
func (o *UniverseInfo) SetUpdateInProgress(v bool) {
	o.UpdateInProgress = &v
}

// GetUpdatingTask returns the UpdatingTask field value if set, zero value otherwise.
func (o *UniverseInfo) GetUpdatingTask() string {
	if o == nil || o.UpdatingTask == nil {
		var ret string
		return ret
	}
	return *o.UpdatingTask
}

// GetUpdatingTaskOk returns a tuple with the UpdatingTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetUpdatingTaskOk() (*string, bool) {
	if o == nil || o.UpdatingTask == nil {
		return nil, false
	}
	return o.UpdatingTask, true
}

// HasUpdatingTask returns a boolean if a field has been set.
func (o *UniverseInfo) HasUpdatingTask() bool {
	if o != nil && o.UpdatingTask != nil {
		return true
	}

	return false
}

// SetUpdatingTask gets a reference to the given string and assigns it to the UpdatingTask field.
func (o *UniverseInfo) SetUpdatingTask(v string) {
	o.UpdatingTask = &v
}

// GetUpdatingTaskUuid returns the UpdatingTaskUuid field value if set, zero value otherwise.
func (o *UniverseInfo) GetUpdatingTaskUuid() string {
	if o == nil || o.UpdatingTaskUuid == nil {
		var ret string
		return ret
	}
	return *o.UpdatingTaskUuid
}

// GetUpdatingTaskUuidOk returns a tuple with the UpdatingTaskUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetUpdatingTaskUuidOk() (*string, bool) {
	if o == nil || o.UpdatingTaskUuid == nil {
		return nil, false
	}
	return o.UpdatingTaskUuid, true
}

// HasUpdatingTaskUuid returns a boolean if a field has been set.
func (o *UniverseInfo) HasUpdatingTaskUuid() bool {
	if o != nil && o.UpdatingTaskUuid != nil {
		return true
	}

	return false
}

// SetUpdatingTaskUuid gets a reference to the given string and assigns it to the UpdatingTaskUuid field.
func (o *UniverseInfo) SetUpdatingTaskUuid(v string) {
	o.UpdatingTaskUuid = &v
}

// GetUpdateSucceeded returns the UpdateSucceeded field value if set, zero value otherwise.
func (o *UniverseInfo) GetUpdateSucceeded() bool {
	if o == nil || o.UpdateSucceeded == nil {
		var ret bool
		return ret
	}
	return *o.UpdateSucceeded
}

// GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetUpdateSucceededOk() (*bool, bool) {
	if o == nil || o.UpdateSucceeded == nil {
		return nil, false
	}
	return o.UpdateSucceeded, true
}

// HasUpdateSucceeded returns a boolean if a field has been set.
func (o *UniverseInfo) HasUpdateSucceeded() bool {
	if o != nil && o.UpdateSucceeded != nil {
		return true
	}

	return false
}

// SetUpdateSucceeded gets a reference to the given bool and assigns it to the UpdateSucceeded field.
func (o *UniverseInfo) SetUpdateSucceeded(v bool) {
	o.UpdateSucceeded = &v
}

// GetPreviousTaskUuid returns the PreviousTaskUuid field value if set, zero value otherwise.
func (o *UniverseInfo) GetPreviousTaskUuid() string {
	if o == nil || o.PreviousTaskUuid == nil {
		var ret string
		return ret
	}
	return *o.PreviousTaskUuid
}

// GetPreviousTaskUuidOk returns a tuple with the PreviousTaskUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetPreviousTaskUuidOk() (*string, bool) {
	if o == nil || o.PreviousTaskUuid == nil {
		return nil, false
	}
	return o.PreviousTaskUuid, true
}

// HasPreviousTaskUuid returns a boolean if a field has been set.
func (o *UniverseInfo) HasPreviousTaskUuid() bool {
	if o != nil && o.PreviousTaskUuid != nil {
		return true
	}

	return false
}

// SetPreviousTaskUuid gets a reference to the given string and assigns it to the PreviousTaskUuid field.
func (o *UniverseInfo) SetPreviousTaskUuid(v string) {
	o.PreviousTaskUuid = &v
}

// GetUniversePaused returns the UniversePaused field value if set, zero value otherwise.
func (o *UniverseInfo) GetUniversePaused() bool {
	if o == nil || o.UniversePaused == nil {
		var ret bool
		return ret
	}
	return *o.UniversePaused
}

// GetUniversePausedOk returns a tuple with the UniversePaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetUniversePausedOk() (*bool, bool) {
	if o == nil || o.UniversePaused == nil {
		return nil, false
	}
	return o.UniversePaused, true
}

// HasUniversePaused returns a boolean if a field has been set.
func (o *UniverseInfo) HasUniversePaused() bool {
	if o != nil && o.UniversePaused != nil {
		return true
	}

	return false
}

// SetUniversePaused gets a reference to the given bool and assigns it to the UniversePaused field.
func (o *UniverseInfo) SetUniversePaused(v bool) {
	o.UniversePaused = &v
}

// GetPlacementModificationTaskUuid returns the PlacementModificationTaskUuid field value if set, zero value otherwise.
func (o *UniverseInfo) GetPlacementModificationTaskUuid() string {
	if o == nil || o.PlacementModificationTaskUuid == nil {
		var ret string
		return ret
	}
	return *o.PlacementModificationTaskUuid
}

// GetPlacementModificationTaskUuidOk returns a tuple with the PlacementModificationTaskUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetPlacementModificationTaskUuidOk() (*string, bool) {
	if o == nil || o.PlacementModificationTaskUuid == nil {
		return nil, false
	}
	return o.PlacementModificationTaskUuid, true
}

// HasPlacementModificationTaskUuid returns a boolean if a field has been set.
func (o *UniverseInfo) HasPlacementModificationTaskUuid() bool {
	if o != nil && o.PlacementModificationTaskUuid != nil {
		return true
	}

	return false
}

// SetPlacementModificationTaskUuid gets a reference to the given string and assigns it to the PlacementModificationTaskUuid field.
func (o *UniverseInfo) SetPlacementModificationTaskUuid(v string) {
	o.PlacementModificationTaskUuid = &v
}

// GetSoftwareUpgradeState returns the SoftwareUpgradeState field value if set, zero value otherwise.
func (o *UniverseInfo) GetSoftwareUpgradeState() string {
	if o == nil || o.SoftwareUpgradeState == nil {
		var ret string
		return ret
	}
	return *o.SoftwareUpgradeState
}

// GetSoftwareUpgradeStateOk returns a tuple with the SoftwareUpgradeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetSoftwareUpgradeStateOk() (*string, bool) {
	if o == nil || o.SoftwareUpgradeState == nil {
		return nil, false
	}
	return o.SoftwareUpgradeState, true
}

// HasSoftwareUpgradeState returns a boolean if a field has been set.
func (o *UniverseInfo) HasSoftwareUpgradeState() bool {
	if o != nil && o.SoftwareUpgradeState != nil {
		return true
	}

	return false
}

// SetSoftwareUpgradeState gets a reference to the given string and assigns it to the SoftwareUpgradeState field.
func (o *UniverseInfo) SetSoftwareUpgradeState(v string) {
	o.SoftwareUpgradeState = &v
}

// GetIsSoftwareRollbackAllowed returns the IsSoftwareRollbackAllowed field value if set, zero value otherwise.
func (o *UniverseInfo) GetIsSoftwareRollbackAllowed() bool {
	if o == nil || o.IsSoftwareRollbackAllowed == nil {
		var ret bool
		return ret
	}
	return *o.IsSoftwareRollbackAllowed
}

// GetIsSoftwareRollbackAllowedOk returns a tuple with the IsSoftwareRollbackAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetIsSoftwareRollbackAllowedOk() (*bool, bool) {
	if o == nil || o.IsSoftwareRollbackAllowed == nil {
		return nil, false
	}
	return o.IsSoftwareRollbackAllowed, true
}

// HasIsSoftwareRollbackAllowed returns a boolean if a field has been set.
func (o *UniverseInfo) HasIsSoftwareRollbackAllowed() bool {
	if o != nil && o.IsSoftwareRollbackAllowed != nil {
		return true
	}

	return false
}

// SetIsSoftwareRollbackAllowed gets a reference to the given bool and assigns it to the IsSoftwareRollbackAllowed field.
func (o *UniverseInfo) SetIsSoftwareRollbackAllowed(v bool) {
	o.IsSoftwareRollbackAllowed = &v
}

// GetPreviousYbSoftwareDetails returns the PreviousYbSoftwareDetails field value if set, zero value otherwise.
func (o *UniverseInfo) GetPreviousYbSoftwareDetails() YbSoftwareDetails {
	if o == nil || o.PreviousYbSoftwareDetails == nil {
		var ret YbSoftwareDetails
		return ret
	}
	return *o.PreviousYbSoftwareDetails
}

// GetPreviousYbSoftwareDetailsOk returns a tuple with the PreviousYbSoftwareDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetPreviousYbSoftwareDetailsOk() (*YbSoftwareDetails, bool) {
	if o == nil || o.PreviousYbSoftwareDetails == nil {
		return nil, false
	}
	return o.PreviousYbSoftwareDetails, true
}

// HasPreviousYbSoftwareDetails returns a boolean if a field has been set.
func (o *UniverseInfo) HasPreviousYbSoftwareDetails() bool {
	if o != nil && o.PreviousYbSoftwareDetails != nil {
		return true
	}

	return false
}

// SetPreviousYbSoftwareDetails gets a reference to the given YbSoftwareDetails and assigns it to the PreviousYbSoftwareDetails field.
func (o *UniverseInfo) SetPreviousYbSoftwareDetails(v YbSoftwareDetails) {
	o.PreviousYbSoftwareDetails = &v
}

// GetAllowedTasksOnFailure returns the AllowedTasksOnFailure field value if set, zero value otherwise.
func (o *UniverseInfo) GetAllowedTasksOnFailure() AllowedTasksOnFailure {
	if o == nil || o.AllowedTasksOnFailure == nil {
		var ret AllowedTasksOnFailure
		return ret
	}
	return *o.AllowedTasksOnFailure
}

// GetAllowedTasksOnFailureOk returns a tuple with the AllowedTasksOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetAllowedTasksOnFailureOk() (*AllowedTasksOnFailure, bool) {
	if o == nil || o.AllowedTasksOnFailure == nil {
		return nil, false
	}
	return o.AllowedTasksOnFailure, true
}

// HasAllowedTasksOnFailure returns a boolean if a field has been set.
func (o *UniverseInfo) HasAllowedTasksOnFailure() bool {
	if o != nil && o.AllowedTasksOnFailure != nil {
		return true
	}

	return false
}

// SetAllowedTasksOnFailure gets a reference to the given AllowedTasksOnFailure and assigns it to the AllowedTasksOnFailure field.
func (o *UniverseInfo) SetAllowedTasksOnFailure(v AllowedTasksOnFailure) {
	o.AllowedTasksOnFailure = &v
}

// GetNodesResizeAvailable returns the NodesResizeAvailable field value if set, zero value otherwise.
func (o *UniverseInfo) GetNodesResizeAvailable() bool {
	if o == nil || o.NodesResizeAvailable == nil {
		var ret bool
		return ret
	}
	return *o.NodesResizeAvailable
}

// GetNodesResizeAvailableOk returns a tuple with the NodesResizeAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetNodesResizeAvailableOk() (*bool, bool) {
	if o == nil || o.NodesResizeAvailable == nil {
		return nil, false
	}
	return o.NodesResizeAvailable, true
}

// HasNodesResizeAvailable returns a boolean if a field has been set.
func (o *UniverseInfo) HasNodesResizeAvailable() bool {
	if o != nil && o.NodesResizeAvailable != nil {
		return true
	}

	return false
}

// SetNodesResizeAvailable gets a reference to the given bool and assigns it to the NodesResizeAvailable field.
func (o *UniverseInfo) SetNodesResizeAvailable(v bool) {
	o.NodesResizeAvailable = &v
}

// GetIsKubernetesOperatorControlled returns the IsKubernetesOperatorControlled field value if set, zero value otherwise.
func (o *UniverseInfo) GetIsKubernetesOperatorControlled() bool {
	if o == nil || o.IsKubernetesOperatorControlled == nil {
		var ret bool
		return ret
	}
	return *o.IsKubernetesOperatorControlled
}

// GetIsKubernetesOperatorControlledOk returns a tuple with the IsKubernetesOperatorControlled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetIsKubernetesOperatorControlledOk() (*bool, bool) {
	if o == nil || o.IsKubernetesOperatorControlled == nil {
		return nil, false
	}
	return o.IsKubernetesOperatorControlled, true
}

// HasIsKubernetesOperatorControlled returns a boolean if a field has been set.
func (o *UniverseInfo) HasIsKubernetesOperatorControlled() bool {
	if o != nil && o.IsKubernetesOperatorControlled != nil {
		return true
	}

	return false
}

// SetIsKubernetesOperatorControlled gets a reference to the given bool and assigns it to the IsKubernetesOperatorControlled field.
func (o *UniverseInfo) SetIsKubernetesOperatorControlled(v bool) {
	o.IsKubernetesOperatorControlled = &v
}

// GetOtelCollectorEnabled returns the OtelCollectorEnabled field value if set, zero value otherwise.
func (o *UniverseInfo) GetOtelCollectorEnabled() bool {
	if o == nil || o.OtelCollectorEnabled == nil {
		var ret bool
		return ret
	}
	return *o.OtelCollectorEnabled
}

// GetOtelCollectorEnabledOk returns a tuple with the OtelCollectorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetOtelCollectorEnabledOk() (*bool, bool) {
	if o == nil || o.OtelCollectorEnabled == nil {
		return nil, false
	}
	return o.OtelCollectorEnabled, true
}

// HasOtelCollectorEnabled returns a boolean if a field has been set.
func (o *UniverseInfo) HasOtelCollectorEnabled() bool {
	if o != nil && o.OtelCollectorEnabled != nil {
		return true
	}

	return false
}

// SetOtelCollectorEnabled gets a reference to the given bool and assigns it to the OtelCollectorEnabled field.
func (o *UniverseInfo) SetOtelCollectorEnabled(v bool) {
	o.OtelCollectorEnabled = &v
}

// GetXClusterInfo returns the XClusterInfo field value if set, zero value otherwise.
func (o *UniverseInfo) GetXClusterInfo() XClusterInfo {
	if o == nil || o.XClusterInfo == nil {
		var ret XClusterInfo
		return ret
	}
	return *o.XClusterInfo
}

// GetXClusterInfoOk returns a tuple with the XClusterInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetXClusterInfoOk() (*XClusterInfo, bool) {
	if o == nil || o.XClusterInfo == nil {
		return nil, false
	}
	return o.XClusterInfo, true
}

// HasXClusterInfo returns a boolean if a field has been set.
func (o *UniverseInfo) HasXClusterInfo() bool {
	if o != nil && o.XClusterInfo != nil {
		return true
	}

	return false
}

// SetXClusterInfo gets a reference to the given XClusterInfo and assigns it to the XClusterInfo field.
func (o *UniverseInfo) SetXClusterInfo(v XClusterInfo) {
	o.XClusterInfo = &v
}

// GetRollMaxBatchSize returns the RollMaxBatchSize field value if set, zero value otherwise.
func (o *UniverseInfo) GetRollMaxBatchSize() RollMaxBatchSize {
	if o == nil || o.RollMaxBatchSize == nil {
		var ret RollMaxBatchSize
		return ret
	}
	return *o.RollMaxBatchSize
}

// GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool) {
	if o == nil || o.RollMaxBatchSize == nil {
		return nil, false
	}
	return o.RollMaxBatchSize, true
}

// HasRollMaxBatchSize returns a boolean if a field has been set.
func (o *UniverseInfo) HasRollMaxBatchSize() bool {
	if o != nil && o.RollMaxBatchSize != nil {
		return true
	}

	return false
}

// SetRollMaxBatchSize gets a reference to the given RollMaxBatchSize and assigns it to the RollMaxBatchSize field.
func (o *UniverseInfo) SetRollMaxBatchSize(v RollMaxBatchSize) {
	o.RollMaxBatchSize = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *UniverseInfo) GetClusters() []ClusterInfo {
	if o == nil || o.Clusters == nil {
		var ret []ClusterInfo
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetClustersOk() (*[]ClusterInfo, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *UniverseInfo) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []ClusterInfo and assigns it to the Clusters field.
func (o *UniverseInfo) SetClusters(v []ClusterInfo) {
	o.Clusters = &v
}

// GetNodeDetailsSet returns the NodeDetailsSet field value if set, zero value otherwise.
func (o *UniverseInfo) GetNodeDetailsSet() []NodeDetails {
	if o == nil || o.NodeDetailsSet == nil {
		var ret []NodeDetails
		return ret
	}
	return *o.NodeDetailsSet
}

// GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetNodeDetailsSetOk() (*[]NodeDetails, bool) {
	if o == nil || o.NodeDetailsSet == nil {
		return nil, false
	}
	return o.NodeDetailsSet, true
}

// HasNodeDetailsSet returns a boolean if a field has been set.
func (o *UniverseInfo) HasNodeDetailsSet() bool {
	if o != nil && o.NodeDetailsSet != nil {
		return true
	}

	return false
}

// SetNodeDetailsSet gets a reference to the given []NodeDetails and assigns it to the NodeDetailsSet field.
func (o *UniverseInfo) SetNodeDetailsSet(v []NodeDetails) {
	o.NodeDetailsSet = &v
}

// GetDrConfigUuidsAsSource returns the DrConfigUuidsAsSource field value if set, zero value otherwise.
func (o *UniverseInfo) GetDrConfigUuidsAsSource() []string {
	if o == nil || o.DrConfigUuidsAsSource == nil {
		var ret []string
		return ret
	}
	return *o.DrConfigUuidsAsSource
}

// GetDrConfigUuidsAsSourceOk returns a tuple with the DrConfigUuidsAsSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetDrConfigUuidsAsSourceOk() (*[]string, bool) {
	if o == nil || o.DrConfigUuidsAsSource == nil {
		return nil, false
	}
	return o.DrConfigUuidsAsSource, true
}

// HasDrConfigUuidsAsSource returns a boolean if a field has been set.
func (o *UniverseInfo) HasDrConfigUuidsAsSource() bool {
	if o != nil && o.DrConfigUuidsAsSource != nil {
		return true
	}

	return false
}

// SetDrConfigUuidsAsSource gets a reference to the given []string and assigns it to the DrConfigUuidsAsSource field.
func (o *UniverseInfo) SetDrConfigUuidsAsSource(v []string) {
	o.DrConfigUuidsAsSource = &v
}

// GetDrConfigUuidsAsTarget returns the DrConfigUuidsAsTarget field value if set, zero value otherwise.
func (o *UniverseInfo) GetDrConfigUuidsAsTarget() []string {
	if o == nil || o.DrConfigUuidsAsTarget == nil {
		var ret []string
		return ret
	}
	return *o.DrConfigUuidsAsTarget
}

// GetDrConfigUuidsAsTargetOk returns a tuple with the DrConfigUuidsAsTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetDrConfigUuidsAsTargetOk() (*[]string, bool) {
	if o == nil || o.DrConfigUuidsAsTarget == nil {
		return nil, false
	}
	return o.DrConfigUuidsAsTarget, true
}

// HasDrConfigUuidsAsTarget returns a boolean if a field has been set.
func (o *UniverseInfo) HasDrConfigUuidsAsTarget() bool {
	if o != nil && o.DrConfigUuidsAsTarget != nil {
		return true
	}

	return false
}

// SetDrConfigUuidsAsTarget gets a reference to the given []string and assigns it to the DrConfigUuidsAsTarget field.
func (o *UniverseInfo) SetDrConfigUuidsAsTarget(v []string) {
	o.DrConfigUuidsAsTarget = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *UniverseInfo) GetResources() UniverseResourceDetails {
	if o == nil || o.Resources == nil {
		var ret UniverseResourceDetails
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetResourcesOk() (*UniverseResourceDetails, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *UniverseInfo) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given UniverseResourceDetails and assigns it to the Resources field.
func (o *UniverseInfo) SetResources(v UniverseResourceDetails) {
	o.Resources = &v
}

// GetSampleAppCommandTxt returns the SampleAppCommandTxt field value if set, zero value otherwise.
func (o *UniverseInfo) GetSampleAppCommandTxt() string {
	if o == nil || o.SampleAppCommandTxt == nil {
		var ret string
		return ret
	}
	return *o.SampleAppCommandTxt
}

// GetSampleAppCommandTxtOk returns a tuple with the SampleAppCommandTxt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseInfo) GetSampleAppCommandTxtOk() (*string, bool) {
	if o == nil || o.SampleAppCommandTxt == nil {
		return nil, false
	}
	return o.SampleAppCommandTxt, true
}

// HasSampleAppCommandTxt returns a boolean if a field has been set.
func (o *UniverseInfo) HasSampleAppCommandTxt() bool {
	if o != nil && o.SampleAppCommandTxt != nil {
		return true
	}

	return false
}

// SetSampleAppCommandTxt gets a reference to the given string and assigns it to the SampleAppCommandTxt field.
func (o *UniverseInfo) SetSampleAppCommandTxt(v string) {
	o.SampleAppCommandTxt = &v
}

func (o UniverseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UniverseUuid != nil {
		toSerialize["universe_uuid"] = o.UniverseUuid
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.CreatingUser != nil {
		toSerialize["creating_user"] = o.CreatingUser
	}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	if o.DnsName != nil {
		toSerialize["dns_name"] = o.DnsName
	}
	if o.YbcSoftwareVersion != nil {
		toSerialize["ybc_software_version"] = o.YbcSoftwareVersion
	}
	if o.YbaUrl != nil {
		toSerialize["yba_url"] = o.YbaUrl
	}
	if o.NodePrefix != nil {
		toSerialize["node_prefix"] = o.NodePrefix
	}
	if o.EncryptionAtRestInfo != nil {
		toSerialize["encryption_at_rest_info"] = o.EncryptionAtRestInfo
	}
	if o.UpdateInProgress != nil {
		toSerialize["update_in_progress"] = o.UpdateInProgress
	}
	if o.UpdatingTask != nil {
		toSerialize["updating_task"] = o.UpdatingTask
	}
	if o.UpdatingTaskUuid != nil {
		toSerialize["updating_task_uuid"] = o.UpdatingTaskUuid
	}
	if o.UpdateSucceeded != nil {
		toSerialize["update_succeeded"] = o.UpdateSucceeded
	}
	if o.PreviousTaskUuid != nil {
		toSerialize["previous_task_uuid"] = o.PreviousTaskUuid
	}
	if o.UniversePaused != nil {
		toSerialize["universe_paused"] = o.UniversePaused
	}
	if o.PlacementModificationTaskUuid != nil {
		toSerialize["placement_modification_task_uuid"] = o.PlacementModificationTaskUuid
	}
	if o.SoftwareUpgradeState != nil {
		toSerialize["software_upgrade_state"] = o.SoftwareUpgradeState
	}
	if o.IsSoftwareRollbackAllowed != nil {
		toSerialize["is_software_rollback_allowed"] = o.IsSoftwareRollbackAllowed
	}
	if o.PreviousYbSoftwareDetails != nil {
		toSerialize["previous_yb_software_details"] = o.PreviousYbSoftwareDetails
	}
	if o.AllowedTasksOnFailure != nil {
		toSerialize["allowed_tasks_on_failure"] = o.AllowedTasksOnFailure
	}
	if o.NodesResizeAvailable != nil {
		toSerialize["nodes_resize_available"] = o.NodesResizeAvailable
	}
	if o.IsKubernetesOperatorControlled != nil {
		toSerialize["is_kubernetes_operator_controlled"] = o.IsKubernetesOperatorControlled
	}
	if o.OtelCollectorEnabled != nil {
		toSerialize["otel_collector_enabled"] = o.OtelCollectorEnabled
	}
	if o.XClusterInfo != nil {
		toSerialize["x_cluster_info"] = o.XClusterInfo
	}
	if o.RollMaxBatchSize != nil {
		toSerialize["roll_max_batch_size"] = o.RollMaxBatchSize
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.NodeDetailsSet != nil {
		toSerialize["node_details_set"] = o.NodeDetailsSet
	}
	if o.DrConfigUuidsAsSource != nil {
		toSerialize["dr_config_uuids_as_source"] = o.DrConfigUuidsAsSource
	}
	if o.DrConfigUuidsAsTarget != nil {
		toSerialize["dr_config_uuids_as_target"] = o.DrConfigUuidsAsTarget
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.SampleAppCommandTxt != nil {
		toSerialize["sample_app_command_txt"] = o.SampleAppCommandTxt
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseInfo struct {
	value *UniverseInfo
	isSet bool
}

func (v NullableUniverseInfo) Get() *UniverseInfo {
	return v.value
}

func (v *NullableUniverseInfo) Set(val *UniverseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseInfo(val *UniverseInfo) *NullableUniverseInfo {
	return &NullableUniverseInfo{value: val, isSet: true}
}

func (v NullableUniverseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

