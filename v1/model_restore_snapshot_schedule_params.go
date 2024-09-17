/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// RestoreSnapshotScheduleParams Restore snapshot schedule parameters
type RestoreSnapshotScheduleParams struct {
	// Amazon Resource Name (ARN) of the CMK
	CmkArn *string `json:"cmkArn,omitempty"`
	CommunicationPorts *CommunicationPorts `json:"communicationPorts,omitempty"`
	CreatingUser Users `json:"creatingUser"`
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`
	EnableYbc *bool `json:"enableYbc,omitempty"`
	EncryptionAtRestConfig *EncryptionAtRestConfig `json:"encryptionAtRestConfig,omitempty"`
	// Error message
	ErrorString *string `json:"errorString,omitempty"`
	// Expected universe version
	ExpectedUniverseVersion *int32 `json:"expectedUniverseVersion,omitempty"`
	ExtraDependencies *ExtraDependencies `json:"extraDependencies,omitempty"`
	InstallYbc *bool `json:"installYbc,omitempty"`
	// Node details
	NodeDetailsSet *[]NodeDetails `json:"nodeDetailsSet,omitempty"`
	// Node exporter user
	NodeExporterUser *string `json:"nodeExporterUser,omitempty"`
	// PITR Config UUID
	PitrConfigUUID *string `json:"pitrConfigUUID,omitempty"`
	PlatformUrl string `json:"platformUrl"`
	PlatformVersion string `json:"platformVersion"`
	// Previous task UUID of a retry
	PreviousTaskUUID *string `json:"previousTaskUUID,omitempty"`
	// Restore Time In millis
	RestoreTimeInMillis *int64 `json:"restoreTimeInMillis,omitempty"`
	SleepAfterMasterRestartMillis int32 `json:"sleepAfterMasterRestartMillis"`
	SleepAfterTServerRestartMillis int32 `json:"sleepAfterTServerRestartMillis"`
	// The source universe's xcluster replication relationships
	SourceXClusterConfigs *[]string `json:"sourceXClusterConfigs,omitempty"`
	// The target universe's xcluster replication relationships
	TargetXClusterConfigs *[]string `json:"targetXClusterConfigs,omitempty"`
	// Universe UUID
	UniverseUUID *string `json:"universeUUID,omitempty"`
	// Previous software version
	YbPrevSoftwareVersion *string `json:"ybPrevSoftwareVersion,omitempty"`
	YbcInstalled *bool `json:"ybcInstalled,omitempty"`
	YbcSoftwareVersion *string `json:"ybcSoftwareVersion,omitempty"`
}

// NewRestoreSnapshotScheduleParams instantiates a new RestoreSnapshotScheduleParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreSnapshotScheduleParams(creatingUser Users, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32) *RestoreSnapshotScheduleParams {
	this := RestoreSnapshotScheduleParams{}
	this.CreatingUser = creatingUser
	this.PlatformUrl = platformUrl
	this.PlatformVersion = platformVersion
	this.SleepAfterMasterRestartMillis = sleepAfterMasterRestartMillis
	this.SleepAfterTServerRestartMillis = sleepAfterTServerRestartMillis
	return &this
}

// NewRestoreSnapshotScheduleParamsWithDefaults instantiates a new RestoreSnapshotScheduleParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreSnapshotScheduleParamsWithDefaults() *RestoreSnapshotScheduleParams {
	this := RestoreSnapshotScheduleParams{}
	return &this
}

// GetCmkArn returns the CmkArn field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetCmkArn() string {
	if o == nil || o.CmkArn == nil {
		var ret string
		return ret
	}
	return *o.CmkArn
}

// GetCmkArnOk returns a tuple with the CmkArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetCmkArnOk() (*string, bool) {
	if o == nil || o.CmkArn == nil {
		return nil, false
	}
	return o.CmkArn, true
}

// HasCmkArn returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasCmkArn() bool {
	if o != nil && o.CmkArn != nil {
		return true
	}

	return false
}

// SetCmkArn gets a reference to the given string and assigns it to the CmkArn field.
func (o *RestoreSnapshotScheduleParams) SetCmkArn(v string) {
	o.CmkArn = &v
}

// GetCommunicationPorts returns the CommunicationPorts field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetCommunicationPorts() CommunicationPorts {
	if o == nil || o.CommunicationPorts == nil {
		var ret CommunicationPorts
		return ret
	}
	return *o.CommunicationPorts
}

// GetCommunicationPortsOk returns a tuple with the CommunicationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetCommunicationPortsOk() (*CommunicationPorts, bool) {
	if o == nil || o.CommunicationPorts == nil {
		return nil, false
	}
	return o.CommunicationPorts, true
}

// HasCommunicationPorts returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasCommunicationPorts() bool {
	if o != nil && o.CommunicationPorts != nil {
		return true
	}

	return false
}

// SetCommunicationPorts gets a reference to the given CommunicationPorts and assigns it to the CommunicationPorts field.
func (o *RestoreSnapshotScheduleParams) SetCommunicationPorts(v CommunicationPorts) {
	o.CommunicationPorts = &v
}

// GetCreatingUser returns the CreatingUser field value
func (o *RestoreSnapshotScheduleParams) GetCreatingUser() Users {
	if o == nil {
		var ret Users
		return ret
	}

	return o.CreatingUser
}

// GetCreatingUserOk returns a tuple with the CreatingUser field value
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetCreatingUserOk() (*Users, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatingUser, true
}

// SetCreatingUser sets field value
func (o *RestoreSnapshotScheduleParams) SetCreatingUser(v Users) {
	o.CreatingUser = v
}

// GetDeviceInfo returns the DeviceInfo field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetDeviceInfo() DeviceInfo {
	if o == nil || o.DeviceInfo == nil {
		var ret DeviceInfo
		return ret
	}
	return *o.DeviceInfo
}

// GetDeviceInfoOk returns a tuple with the DeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetDeviceInfoOk() (*DeviceInfo, bool) {
	if o == nil || o.DeviceInfo == nil {
		return nil, false
	}
	return o.DeviceInfo, true
}

// HasDeviceInfo returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasDeviceInfo() bool {
	if o != nil && o.DeviceInfo != nil {
		return true
	}

	return false
}

// SetDeviceInfo gets a reference to the given DeviceInfo and assigns it to the DeviceInfo field.
func (o *RestoreSnapshotScheduleParams) SetDeviceInfo(v DeviceInfo) {
	o.DeviceInfo = &v
}

// GetEnableYbc returns the EnableYbc field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetEnableYbc() bool {
	if o == nil || o.EnableYbc == nil {
		var ret bool
		return ret
	}
	return *o.EnableYbc
}

// GetEnableYbcOk returns a tuple with the EnableYbc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetEnableYbcOk() (*bool, bool) {
	if o == nil || o.EnableYbc == nil {
		return nil, false
	}
	return o.EnableYbc, true
}

// HasEnableYbc returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasEnableYbc() bool {
	if o != nil && o.EnableYbc != nil {
		return true
	}

	return false
}

// SetEnableYbc gets a reference to the given bool and assigns it to the EnableYbc field.
func (o *RestoreSnapshotScheduleParams) SetEnableYbc(v bool) {
	o.EnableYbc = &v
}

// GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig {
	if o == nil || o.EncryptionAtRestConfig == nil {
		var ret EncryptionAtRestConfig
		return ret
	}
	return *o.EncryptionAtRestConfig
}

// GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool) {
	if o == nil || o.EncryptionAtRestConfig == nil {
		return nil, false
	}
	return o.EncryptionAtRestConfig, true
}

// HasEncryptionAtRestConfig returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasEncryptionAtRestConfig() bool {
	if o != nil && o.EncryptionAtRestConfig != nil {
		return true
	}

	return false
}

// SetEncryptionAtRestConfig gets a reference to the given EncryptionAtRestConfig and assigns it to the EncryptionAtRestConfig field.
func (o *RestoreSnapshotScheduleParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig) {
	o.EncryptionAtRestConfig = &v
}

// GetErrorString returns the ErrorString field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetErrorString() string {
	if o == nil || o.ErrorString == nil {
		var ret string
		return ret
	}
	return *o.ErrorString
}

// GetErrorStringOk returns a tuple with the ErrorString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetErrorStringOk() (*string, bool) {
	if o == nil || o.ErrorString == nil {
		return nil, false
	}
	return o.ErrorString, true
}

// HasErrorString returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasErrorString() bool {
	if o != nil && o.ErrorString != nil {
		return true
	}

	return false
}

// SetErrorString gets a reference to the given string and assigns it to the ErrorString field.
func (o *RestoreSnapshotScheduleParams) SetErrorString(v string) {
	o.ErrorString = &v
}

// GetExpectedUniverseVersion returns the ExpectedUniverseVersion field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetExpectedUniverseVersion() int32 {
	if o == nil || o.ExpectedUniverseVersion == nil {
		var ret int32
		return ret
	}
	return *o.ExpectedUniverseVersion
}

// GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetExpectedUniverseVersionOk() (*int32, bool) {
	if o == nil || o.ExpectedUniverseVersion == nil {
		return nil, false
	}
	return o.ExpectedUniverseVersion, true
}

// HasExpectedUniverseVersion returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasExpectedUniverseVersion() bool {
	if o != nil && o.ExpectedUniverseVersion != nil {
		return true
	}

	return false
}

// SetExpectedUniverseVersion gets a reference to the given int32 and assigns it to the ExpectedUniverseVersion field.
func (o *RestoreSnapshotScheduleParams) SetExpectedUniverseVersion(v int32) {
	o.ExpectedUniverseVersion = &v
}

// GetExtraDependencies returns the ExtraDependencies field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetExtraDependencies() ExtraDependencies {
	if o == nil || o.ExtraDependencies == nil {
		var ret ExtraDependencies
		return ret
	}
	return *o.ExtraDependencies
}

// GetExtraDependenciesOk returns a tuple with the ExtraDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetExtraDependenciesOk() (*ExtraDependencies, bool) {
	if o == nil || o.ExtraDependencies == nil {
		return nil, false
	}
	return o.ExtraDependencies, true
}

// HasExtraDependencies returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasExtraDependencies() bool {
	if o != nil && o.ExtraDependencies != nil {
		return true
	}

	return false
}

// SetExtraDependencies gets a reference to the given ExtraDependencies and assigns it to the ExtraDependencies field.
func (o *RestoreSnapshotScheduleParams) SetExtraDependencies(v ExtraDependencies) {
	o.ExtraDependencies = &v
}

// GetInstallYbc returns the InstallYbc field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetInstallYbc() bool {
	if o == nil || o.InstallYbc == nil {
		var ret bool
		return ret
	}
	return *o.InstallYbc
}

// GetInstallYbcOk returns a tuple with the InstallYbc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetInstallYbcOk() (*bool, bool) {
	if o == nil || o.InstallYbc == nil {
		return nil, false
	}
	return o.InstallYbc, true
}

// HasInstallYbc returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasInstallYbc() bool {
	if o != nil && o.InstallYbc != nil {
		return true
	}

	return false
}

// SetInstallYbc gets a reference to the given bool and assigns it to the InstallYbc field.
func (o *RestoreSnapshotScheduleParams) SetInstallYbc(v bool) {
	o.InstallYbc = &v
}

// GetNodeDetailsSet returns the NodeDetailsSet field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetNodeDetailsSet() []NodeDetails {
	if o == nil || o.NodeDetailsSet == nil {
		var ret []NodeDetails
		return ret
	}
	return *o.NodeDetailsSet
}

// GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool) {
	if o == nil || o.NodeDetailsSet == nil {
		return nil, false
	}
	return o.NodeDetailsSet, true
}

// HasNodeDetailsSet returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasNodeDetailsSet() bool {
	if o != nil && o.NodeDetailsSet != nil {
		return true
	}

	return false
}

// SetNodeDetailsSet gets a reference to the given []NodeDetails and assigns it to the NodeDetailsSet field.
func (o *RestoreSnapshotScheduleParams) SetNodeDetailsSet(v []NodeDetails) {
	o.NodeDetailsSet = &v
}

// GetNodeExporterUser returns the NodeExporterUser field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetNodeExporterUser() string {
	if o == nil || o.NodeExporterUser == nil {
		var ret string
		return ret
	}
	return *o.NodeExporterUser
}

// GetNodeExporterUserOk returns a tuple with the NodeExporterUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetNodeExporterUserOk() (*string, bool) {
	if o == nil || o.NodeExporterUser == nil {
		return nil, false
	}
	return o.NodeExporterUser, true
}

// HasNodeExporterUser returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasNodeExporterUser() bool {
	if o != nil && o.NodeExporterUser != nil {
		return true
	}

	return false
}

// SetNodeExporterUser gets a reference to the given string and assigns it to the NodeExporterUser field.
func (o *RestoreSnapshotScheduleParams) SetNodeExporterUser(v string) {
	o.NodeExporterUser = &v
}

// GetPitrConfigUUID returns the PitrConfigUUID field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetPitrConfigUUID() string {
	if o == nil || o.PitrConfigUUID == nil {
		var ret string
		return ret
	}
	return *o.PitrConfigUUID
}

// GetPitrConfigUUIDOk returns a tuple with the PitrConfigUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetPitrConfigUUIDOk() (*string, bool) {
	if o == nil || o.PitrConfigUUID == nil {
		return nil, false
	}
	return o.PitrConfigUUID, true
}

// HasPitrConfigUUID returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasPitrConfigUUID() bool {
	if o != nil && o.PitrConfigUUID != nil {
		return true
	}

	return false
}

// SetPitrConfigUUID gets a reference to the given string and assigns it to the PitrConfigUUID field.
func (o *RestoreSnapshotScheduleParams) SetPitrConfigUUID(v string) {
	o.PitrConfigUUID = &v
}

// GetPlatformUrl returns the PlatformUrl field value
func (o *RestoreSnapshotScheduleParams) GetPlatformUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformUrl
}

// GetPlatformUrlOk returns a tuple with the PlatformUrl field value
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetPlatformUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlatformUrl, true
}

// SetPlatformUrl sets field value
func (o *RestoreSnapshotScheduleParams) SetPlatformUrl(v string) {
	o.PlatformUrl = v
}

// GetPlatformVersion returns the PlatformVersion field value
func (o *RestoreSnapshotScheduleParams) GetPlatformVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformVersion
}

// GetPlatformVersionOk returns a tuple with the PlatformVersion field value
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetPlatformVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlatformVersion, true
}

// SetPlatformVersion sets field value
func (o *RestoreSnapshotScheduleParams) SetPlatformVersion(v string) {
	o.PlatformVersion = v
}

// GetPreviousTaskUUID returns the PreviousTaskUUID field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetPreviousTaskUUID() string {
	if o == nil || o.PreviousTaskUUID == nil {
		var ret string
		return ret
	}
	return *o.PreviousTaskUUID
}

// GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetPreviousTaskUUIDOk() (*string, bool) {
	if o == nil || o.PreviousTaskUUID == nil {
		return nil, false
	}
	return o.PreviousTaskUUID, true
}

// HasPreviousTaskUUID returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasPreviousTaskUUID() bool {
	if o != nil && o.PreviousTaskUUID != nil {
		return true
	}

	return false
}

// SetPreviousTaskUUID gets a reference to the given string and assigns it to the PreviousTaskUUID field.
func (o *RestoreSnapshotScheduleParams) SetPreviousTaskUUID(v string) {
	o.PreviousTaskUUID = &v
}

// GetRestoreTimeInMillis returns the RestoreTimeInMillis field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetRestoreTimeInMillis() int64 {
	if o == nil || o.RestoreTimeInMillis == nil {
		var ret int64
		return ret
	}
	return *o.RestoreTimeInMillis
}

// GetRestoreTimeInMillisOk returns a tuple with the RestoreTimeInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetRestoreTimeInMillisOk() (*int64, bool) {
	if o == nil || o.RestoreTimeInMillis == nil {
		return nil, false
	}
	return o.RestoreTimeInMillis, true
}

// HasRestoreTimeInMillis returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasRestoreTimeInMillis() bool {
	if o != nil && o.RestoreTimeInMillis != nil {
		return true
	}

	return false
}

// SetRestoreTimeInMillis gets a reference to the given int64 and assigns it to the RestoreTimeInMillis field.
func (o *RestoreSnapshotScheduleParams) SetRestoreTimeInMillis(v int64) {
	o.RestoreTimeInMillis = &v
}

// GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field value
func (o *RestoreSnapshotScheduleParams) GetSleepAfterMasterRestartMillis() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SleepAfterMasterRestartMillis
}

// GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field value
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SleepAfterMasterRestartMillis, true
}

// SetSleepAfterMasterRestartMillis sets field value
func (o *RestoreSnapshotScheduleParams) SetSleepAfterMasterRestartMillis(v int32) {
	o.SleepAfterMasterRestartMillis = v
}

// GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field value
func (o *RestoreSnapshotScheduleParams) GetSleepAfterTServerRestartMillis() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SleepAfterTServerRestartMillis
}

// GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field value
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SleepAfterTServerRestartMillis, true
}

// SetSleepAfterTServerRestartMillis sets field value
func (o *RestoreSnapshotScheduleParams) SetSleepAfterTServerRestartMillis(v int32) {
	o.SleepAfterTServerRestartMillis = v
}

// GetSourceXClusterConfigs returns the SourceXClusterConfigs field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetSourceXClusterConfigs() []string {
	if o == nil || o.SourceXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.SourceXClusterConfigs
}

// GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetSourceXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.SourceXClusterConfigs == nil {
		return nil, false
	}
	return o.SourceXClusterConfigs, true
}

// HasSourceXClusterConfigs returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasSourceXClusterConfigs() bool {
	if o != nil && o.SourceXClusterConfigs != nil {
		return true
	}

	return false
}

// SetSourceXClusterConfigs gets a reference to the given []string and assigns it to the SourceXClusterConfigs field.
func (o *RestoreSnapshotScheduleParams) SetSourceXClusterConfigs(v []string) {
	o.SourceXClusterConfigs = &v
}

// GetTargetXClusterConfigs returns the TargetXClusterConfigs field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetTargetXClusterConfigs() []string {
	if o == nil || o.TargetXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.TargetXClusterConfigs
}

// GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetTargetXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.TargetXClusterConfigs == nil {
		return nil, false
	}
	return o.TargetXClusterConfigs, true
}

// HasTargetXClusterConfigs returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasTargetXClusterConfigs() bool {
	if o != nil && o.TargetXClusterConfigs != nil {
		return true
	}

	return false
}

// SetTargetXClusterConfigs gets a reference to the given []string and assigns it to the TargetXClusterConfigs field.
func (o *RestoreSnapshotScheduleParams) SetTargetXClusterConfigs(v []string) {
	o.TargetXClusterConfigs = &v
}

// GetUniverseUUID returns the UniverseUUID field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetUniverseUUID() string {
	if o == nil || o.UniverseUUID == nil {
		var ret string
		return ret
	}
	return *o.UniverseUUID
}

// GetUniverseUUIDOk returns a tuple with the UniverseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetUniverseUUIDOk() (*string, bool) {
	if o == nil || o.UniverseUUID == nil {
		return nil, false
	}
	return o.UniverseUUID, true
}

// HasUniverseUUID returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasUniverseUUID() bool {
	if o != nil && o.UniverseUUID != nil {
		return true
	}

	return false
}

// SetUniverseUUID gets a reference to the given string and assigns it to the UniverseUUID field.
func (o *RestoreSnapshotScheduleParams) SetUniverseUUID(v string) {
	o.UniverseUUID = &v
}

// GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetYbPrevSoftwareVersion() string {
	if o == nil || o.YbPrevSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbPrevSoftwareVersion
}

// GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetYbPrevSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbPrevSoftwareVersion == nil {
		return nil, false
	}
	return o.YbPrevSoftwareVersion, true
}

// HasYbPrevSoftwareVersion returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasYbPrevSoftwareVersion() bool {
	if o != nil && o.YbPrevSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbPrevSoftwareVersion gets a reference to the given string and assigns it to the YbPrevSoftwareVersion field.
func (o *RestoreSnapshotScheduleParams) SetYbPrevSoftwareVersion(v string) {
	o.YbPrevSoftwareVersion = &v
}

// GetYbcInstalled returns the YbcInstalled field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetYbcInstalled() bool {
	if o == nil || o.YbcInstalled == nil {
		var ret bool
		return ret
	}
	return *o.YbcInstalled
}

// GetYbcInstalledOk returns a tuple with the YbcInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetYbcInstalledOk() (*bool, bool) {
	if o == nil || o.YbcInstalled == nil {
		return nil, false
	}
	return o.YbcInstalled, true
}

// HasYbcInstalled returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasYbcInstalled() bool {
	if o != nil && o.YbcInstalled != nil {
		return true
	}

	return false
}

// SetYbcInstalled gets a reference to the given bool and assigns it to the YbcInstalled field.
func (o *RestoreSnapshotScheduleParams) SetYbcInstalled(v bool) {
	o.YbcInstalled = &v
}

// GetYbcSoftwareVersion returns the YbcSoftwareVersion field value if set, zero value otherwise.
func (o *RestoreSnapshotScheduleParams) GetYbcSoftwareVersion() string {
	if o == nil || o.YbcSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbcSoftwareVersion
}

// GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreSnapshotScheduleParams) GetYbcSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbcSoftwareVersion == nil {
		return nil, false
	}
	return o.YbcSoftwareVersion, true
}

// HasYbcSoftwareVersion returns a boolean if a field has been set.
func (o *RestoreSnapshotScheduleParams) HasYbcSoftwareVersion() bool {
	if o != nil && o.YbcSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbcSoftwareVersion gets a reference to the given string and assigns it to the YbcSoftwareVersion field.
func (o *RestoreSnapshotScheduleParams) SetYbcSoftwareVersion(v string) {
	o.YbcSoftwareVersion = &v
}

func (o RestoreSnapshotScheduleParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CmkArn != nil {
		toSerialize["cmkArn"] = o.CmkArn
	}
	if o.CommunicationPorts != nil {
		toSerialize["communicationPorts"] = o.CommunicationPorts
	}
	if true {
		toSerialize["creatingUser"] = o.CreatingUser
	}
	if o.DeviceInfo != nil {
		toSerialize["deviceInfo"] = o.DeviceInfo
	}
	if o.EnableYbc != nil {
		toSerialize["enableYbc"] = o.EnableYbc
	}
	if o.EncryptionAtRestConfig != nil {
		toSerialize["encryptionAtRestConfig"] = o.EncryptionAtRestConfig
	}
	if o.ErrorString != nil {
		toSerialize["errorString"] = o.ErrorString
	}
	if o.ExpectedUniverseVersion != nil {
		toSerialize["expectedUniverseVersion"] = o.ExpectedUniverseVersion
	}
	if o.ExtraDependencies != nil {
		toSerialize["extraDependencies"] = o.ExtraDependencies
	}
	if o.InstallYbc != nil {
		toSerialize["installYbc"] = o.InstallYbc
	}
	if o.NodeDetailsSet != nil {
		toSerialize["nodeDetailsSet"] = o.NodeDetailsSet
	}
	if o.NodeExporterUser != nil {
		toSerialize["nodeExporterUser"] = o.NodeExporterUser
	}
	if o.PitrConfigUUID != nil {
		toSerialize["pitrConfigUUID"] = o.PitrConfigUUID
	}
	if true {
		toSerialize["platformUrl"] = o.PlatformUrl
	}
	if true {
		toSerialize["platformVersion"] = o.PlatformVersion
	}
	if o.PreviousTaskUUID != nil {
		toSerialize["previousTaskUUID"] = o.PreviousTaskUUID
	}
	if o.RestoreTimeInMillis != nil {
		toSerialize["restoreTimeInMillis"] = o.RestoreTimeInMillis
	}
	if true {
		toSerialize["sleepAfterMasterRestartMillis"] = o.SleepAfterMasterRestartMillis
	}
	if true {
		toSerialize["sleepAfterTServerRestartMillis"] = o.SleepAfterTServerRestartMillis
	}
	if o.SourceXClusterConfigs != nil {
		toSerialize["sourceXClusterConfigs"] = o.SourceXClusterConfigs
	}
	if o.TargetXClusterConfigs != nil {
		toSerialize["targetXClusterConfigs"] = o.TargetXClusterConfigs
	}
	if o.UniverseUUID != nil {
		toSerialize["universeUUID"] = o.UniverseUUID
	}
	if o.YbPrevSoftwareVersion != nil {
		toSerialize["ybPrevSoftwareVersion"] = o.YbPrevSoftwareVersion
	}
	if o.YbcInstalled != nil {
		toSerialize["ybcInstalled"] = o.YbcInstalled
	}
	if o.YbcSoftwareVersion != nil {
		toSerialize["ybcSoftwareVersion"] = o.YbcSoftwareVersion
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreSnapshotScheduleParams struct {
	value *RestoreSnapshotScheduleParams
	isSet bool
}

func (v NullableRestoreSnapshotScheduleParams) Get() *RestoreSnapshotScheduleParams {
	return v.value
}

func (v *NullableRestoreSnapshotScheduleParams) Set(val *RestoreSnapshotScheduleParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreSnapshotScheduleParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreSnapshotScheduleParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreSnapshotScheduleParams(val *RestoreSnapshotScheduleParams) *NullableRestoreSnapshotScheduleParams {
	return &NullableRestoreSnapshotScheduleParams{value: val, isSet: true}
}

func (v NullableRestoreSnapshotScheduleParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreSnapshotScheduleParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


