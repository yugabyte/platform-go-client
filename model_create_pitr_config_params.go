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
)

// CreatePitrConfigParams Create PITR config parameters
type CreatePitrConfigParams struct {
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
	// Time interval between snapshots
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty"`
	// PITR config name
	Name *string `json:"name,omitempty"`
	// Node details
	NodeDetailsSet *[]NodeDetails `json:"nodeDetailsSet,omitempty"`
	// Node exporter user
	NodeExporterUser *string `json:"nodeExporterUser,omitempty"`
	PlatformUrl string `json:"platformUrl"`
	PlatformVersion string `json:"platformVersion"`
	// Previous task UUID of a retry
	PreviousTaskUUID *string `json:"previousTaskUUID,omitempty"`
	// Retention period of a snapshot
	RetentionPeriodInSeconds *int64 `json:"retentionPeriodInSeconds,omitempty"`
	SleepAfterMasterRestartMillis int32 `json:"sleepAfterMasterRestartMillis"`
	SleepAfterTServerRestartMillis int32 `json:"sleepAfterTServerRestartMillis"`
	// The source universe's xcluster replication relationships
	SourceXClusterConfigs *[]string `json:"sourceXClusterConfigs,omitempty"`
	// The target universe's xcluster replication relationships
	TargetXClusterConfigs *[]string `json:"targetXClusterConfigs,omitempty"`
	// Previous software version
	YbPrevSoftwareVersion *string `json:"ybPrevSoftwareVersion,omitempty"`
	YbcInstalled *bool `json:"ybcInstalled,omitempty"`
	YbcSoftwareVersion *string `json:"ybcSoftwareVersion,omitempty"`
}

// NewCreatePitrConfigParams instantiates a new CreatePitrConfigParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePitrConfigParams(creatingUser Users, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32) *CreatePitrConfigParams {
	this := CreatePitrConfigParams{}
	this.CreatingUser = creatingUser
	this.PlatformUrl = platformUrl
	this.PlatformVersion = platformVersion
	this.SleepAfterMasterRestartMillis = sleepAfterMasterRestartMillis
	this.SleepAfterTServerRestartMillis = sleepAfterTServerRestartMillis
	return &this
}

// NewCreatePitrConfigParamsWithDefaults instantiates a new CreatePitrConfigParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePitrConfigParamsWithDefaults() *CreatePitrConfigParams {
	this := CreatePitrConfigParams{}
	return &this
}

// GetCmkArn returns the CmkArn field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetCmkArn() string {
	if o == nil || o.CmkArn == nil {
		var ret string
		return ret
	}
	return *o.CmkArn
}

// GetCmkArnOk returns a tuple with the CmkArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetCmkArnOk() (*string, bool) {
	if o == nil || o.CmkArn == nil {
		return nil, false
	}
	return o.CmkArn, true
}

// HasCmkArn returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasCmkArn() bool {
	if o != nil && o.CmkArn != nil {
		return true
	}

	return false
}

// SetCmkArn gets a reference to the given string and assigns it to the CmkArn field.
func (o *CreatePitrConfigParams) SetCmkArn(v string) {
	o.CmkArn = &v
}

// GetCommunicationPorts returns the CommunicationPorts field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetCommunicationPorts() CommunicationPorts {
	if o == nil || o.CommunicationPorts == nil {
		var ret CommunicationPorts
		return ret
	}
	return *o.CommunicationPorts
}

// GetCommunicationPortsOk returns a tuple with the CommunicationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetCommunicationPortsOk() (*CommunicationPorts, bool) {
	if o == nil || o.CommunicationPorts == nil {
		return nil, false
	}
	return o.CommunicationPorts, true
}

// HasCommunicationPorts returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasCommunicationPorts() bool {
	if o != nil && o.CommunicationPorts != nil {
		return true
	}

	return false
}

// SetCommunicationPorts gets a reference to the given CommunicationPorts and assigns it to the CommunicationPorts field.
func (o *CreatePitrConfigParams) SetCommunicationPorts(v CommunicationPorts) {
	o.CommunicationPorts = &v
}

// GetCreatingUser returns the CreatingUser field value
func (o *CreatePitrConfigParams) GetCreatingUser() Users {
	if o == nil {
		var ret Users
		return ret
	}

	return o.CreatingUser
}

// GetCreatingUserOk returns a tuple with the CreatingUser field value
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetCreatingUserOk() (*Users, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatingUser, true
}

// SetCreatingUser sets field value
func (o *CreatePitrConfigParams) SetCreatingUser(v Users) {
	o.CreatingUser = v
}

// GetDeviceInfo returns the DeviceInfo field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetDeviceInfo() DeviceInfo {
	if o == nil || o.DeviceInfo == nil {
		var ret DeviceInfo
		return ret
	}
	return *o.DeviceInfo
}

// GetDeviceInfoOk returns a tuple with the DeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetDeviceInfoOk() (*DeviceInfo, bool) {
	if o == nil || o.DeviceInfo == nil {
		return nil, false
	}
	return o.DeviceInfo, true
}

// HasDeviceInfo returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasDeviceInfo() bool {
	if o != nil && o.DeviceInfo != nil {
		return true
	}

	return false
}

// SetDeviceInfo gets a reference to the given DeviceInfo and assigns it to the DeviceInfo field.
func (o *CreatePitrConfigParams) SetDeviceInfo(v DeviceInfo) {
	o.DeviceInfo = &v
}

// GetEnableYbc returns the EnableYbc field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetEnableYbc() bool {
	if o == nil || o.EnableYbc == nil {
		var ret bool
		return ret
	}
	return *o.EnableYbc
}

// GetEnableYbcOk returns a tuple with the EnableYbc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetEnableYbcOk() (*bool, bool) {
	if o == nil || o.EnableYbc == nil {
		return nil, false
	}
	return o.EnableYbc, true
}

// HasEnableYbc returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasEnableYbc() bool {
	if o != nil && o.EnableYbc != nil {
		return true
	}

	return false
}

// SetEnableYbc gets a reference to the given bool and assigns it to the EnableYbc field.
func (o *CreatePitrConfigParams) SetEnableYbc(v bool) {
	o.EnableYbc = &v
}

// GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig {
	if o == nil || o.EncryptionAtRestConfig == nil {
		var ret EncryptionAtRestConfig
		return ret
	}
	return *o.EncryptionAtRestConfig
}

// GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool) {
	if o == nil || o.EncryptionAtRestConfig == nil {
		return nil, false
	}
	return o.EncryptionAtRestConfig, true
}

// HasEncryptionAtRestConfig returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasEncryptionAtRestConfig() bool {
	if o != nil && o.EncryptionAtRestConfig != nil {
		return true
	}

	return false
}

// SetEncryptionAtRestConfig gets a reference to the given EncryptionAtRestConfig and assigns it to the EncryptionAtRestConfig field.
func (o *CreatePitrConfigParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig) {
	o.EncryptionAtRestConfig = &v
}

// GetErrorString returns the ErrorString field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetErrorString() string {
	if o == nil || o.ErrorString == nil {
		var ret string
		return ret
	}
	return *o.ErrorString
}

// GetErrorStringOk returns a tuple with the ErrorString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetErrorStringOk() (*string, bool) {
	if o == nil || o.ErrorString == nil {
		return nil, false
	}
	return o.ErrorString, true
}

// HasErrorString returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasErrorString() bool {
	if o != nil && o.ErrorString != nil {
		return true
	}

	return false
}

// SetErrorString gets a reference to the given string and assigns it to the ErrorString field.
func (o *CreatePitrConfigParams) SetErrorString(v string) {
	o.ErrorString = &v
}

// GetExpectedUniverseVersion returns the ExpectedUniverseVersion field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetExpectedUniverseVersion() int32 {
	if o == nil || o.ExpectedUniverseVersion == nil {
		var ret int32
		return ret
	}
	return *o.ExpectedUniverseVersion
}

// GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetExpectedUniverseVersionOk() (*int32, bool) {
	if o == nil || o.ExpectedUniverseVersion == nil {
		return nil, false
	}
	return o.ExpectedUniverseVersion, true
}

// HasExpectedUniverseVersion returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasExpectedUniverseVersion() bool {
	if o != nil && o.ExpectedUniverseVersion != nil {
		return true
	}

	return false
}

// SetExpectedUniverseVersion gets a reference to the given int32 and assigns it to the ExpectedUniverseVersion field.
func (o *CreatePitrConfigParams) SetExpectedUniverseVersion(v int32) {
	o.ExpectedUniverseVersion = &v
}

// GetExtraDependencies returns the ExtraDependencies field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetExtraDependencies() ExtraDependencies {
	if o == nil || o.ExtraDependencies == nil {
		var ret ExtraDependencies
		return ret
	}
	return *o.ExtraDependencies
}

// GetExtraDependenciesOk returns a tuple with the ExtraDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetExtraDependenciesOk() (*ExtraDependencies, bool) {
	if o == nil || o.ExtraDependencies == nil {
		return nil, false
	}
	return o.ExtraDependencies, true
}

// HasExtraDependencies returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasExtraDependencies() bool {
	if o != nil && o.ExtraDependencies != nil {
		return true
	}

	return false
}

// SetExtraDependencies gets a reference to the given ExtraDependencies and assigns it to the ExtraDependencies field.
func (o *CreatePitrConfigParams) SetExtraDependencies(v ExtraDependencies) {
	o.ExtraDependencies = &v
}

// GetInstallYbc returns the InstallYbc field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetInstallYbc() bool {
	if o == nil || o.InstallYbc == nil {
		var ret bool
		return ret
	}
	return *o.InstallYbc
}

// GetInstallYbcOk returns a tuple with the InstallYbc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetInstallYbcOk() (*bool, bool) {
	if o == nil || o.InstallYbc == nil {
		return nil, false
	}
	return o.InstallYbc, true
}

// HasInstallYbc returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasInstallYbc() bool {
	if o != nil && o.InstallYbc != nil {
		return true
	}

	return false
}

// SetInstallYbc gets a reference to the given bool and assigns it to the InstallYbc field.
func (o *CreatePitrConfigParams) SetInstallYbc(v bool) {
	o.InstallYbc = &v
}

// GetIntervalInSeconds returns the IntervalInSeconds field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetIntervalInSeconds() int64 {
	if o == nil || o.IntervalInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.IntervalInSeconds
}

// GetIntervalInSecondsOk returns a tuple with the IntervalInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetIntervalInSecondsOk() (*int64, bool) {
	if o == nil || o.IntervalInSeconds == nil {
		return nil, false
	}
	return o.IntervalInSeconds, true
}

// HasIntervalInSeconds returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasIntervalInSeconds() bool {
	if o != nil && o.IntervalInSeconds != nil {
		return true
	}

	return false
}

// SetIntervalInSeconds gets a reference to the given int64 and assigns it to the IntervalInSeconds field.
func (o *CreatePitrConfigParams) SetIntervalInSeconds(v int64) {
	o.IntervalInSeconds = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreatePitrConfigParams) SetName(v string) {
	o.Name = &v
}

// GetNodeDetailsSet returns the NodeDetailsSet field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetNodeDetailsSet() []NodeDetails {
	if o == nil || o.NodeDetailsSet == nil {
		var ret []NodeDetails
		return ret
	}
	return *o.NodeDetailsSet
}

// GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool) {
	if o == nil || o.NodeDetailsSet == nil {
		return nil, false
	}
	return o.NodeDetailsSet, true
}

// HasNodeDetailsSet returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasNodeDetailsSet() bool {
	if o != nil && o.NodeDetailsSet != nil {
		return true
	}

	return false
}

// SetNodeDetailsSet gets a reference to the given []NodeDetails and assigns it to the NodeDetailsSet field.
func (o *CreatePitrConfigParams) SetNodeDetailsSet(v []NodeDetails) {
	o.NodeDetailsSet = &v
}

// GetNodeExporterUser returns the NodeExporterUser field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetNodeExporterUser() string {
	if o == nil || o.NodeExporterUser == nil {
		var ret string
		return ret
	}
	return *o.NodeExporterUser
}

// GetNodeExporterUserOk returns a tuple with the NodeExporterUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetNodeExporterUserOk() (*string, bool) {
	if o == nil || o.NodeExporterUser == nil {
		return nil, false
	}
	return o.NodeExporterUser, true
}

// HasNodeExporterUser returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasNodeExporterUser() bool {
	if o != nil && o.NodeExporterUser != nil {
		return true
	}

	return false
}

// SetNodeExporterUser gets a reference to the given string and assigns it to the NodeExporterUser field.
func (o *CreatePitrConfigParams) SetNodeExporterUser(v string) {
	o.NodeExporterUser = &v
}

// GetPlatformUrl returns the PlatformUrl field value
func (o *CreatePitrConfigParams) GetPlatformUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformUrl
}

// GetPlatformUrlOk returns a tuple with the PlatformUrl field value
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetPlatformUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlatformUrl, true
}

// SetPlatformUrl sets field value
func (o *CreatePitrConfigParams) SetPlatformUrl(v string) {
	o.PlatformUrl = v
}

// GetPlatformVersion returns the PlatformVersion field value
func (o *CreatePitrConfigParams) GetPlatformVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformVersion
}

// GetPlatformVersionOk returns a tuple with the PlatformVersion field value
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetPlatformVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlatformVersion, true
}

// SetPlatformVersion sets field value
func (o *CreatePitrConfigParams) SetPlatformVersion(v string) {
	o.PlatformVersion = v
}

// GetPreviousTaskUUID returns the PreviousTaskUUID field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetPreviousTaskUUID() string {
	if o == nil || o.PreviousTaskUUID == nil {
		var ret string
		return ret
	}
	return *o.PreviousTaskUUID
}

// GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetPreviousTaskUUIDOk() (*string, bool) {
	if o == nil || o.PreviousTaskUUID == nil {
		return nil, false
	}
	return o.PreviousTaskUUID, true
}

// HasPreviousTaskUUID returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasPreviousTaskUUID() bool {
	if o != nil && o.PreviousTaskUUID != nil {
		return true
	}

	return false
}

// SetPreviousTaskUUID gets a reference to the given string and assigns it to the PreviousTaskUUID field.
func (o *CreatePitrConfigParams) SetPreviousTaskUUID(v string) {
	o.PreviousTaskUUID = &v
}

// GetRetentionPeriodInSeconds returns the RetentionPeriodInSeconds field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetRetentionPeriodInSeconds() int64 {
	if o == nil || o.RetentionPeriodInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.RetentionPeriodInSeconds
}

// GetRetentionPeriodInSecondsOk returns a tuple with the RetentionPeriodInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetRetentionPeriodInSecondsOk() (*int64, bool) {
	if o == nil || o.RetentionPeriodInSeconds == nil {
		return nil, false
	}
	return o.RetentionPeriodInSeconds, true
}

// HasRetentionPeriodInSeconds returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasRetentionPeriodInSeconds() bool {
	if o != nil && o.RetentionPeriodInSeconds != nil {
		return true
	}

	return false
}

// SetRetentionPeriodInSeconds gets a reference to the given int64 and assigns it to the RetentionPeriodInSeconds field.
func (o *CreatePitrConfigParams) SetRetentionPeriodInSeconds(v int64) {
	o.RetentionPeriodInSeconds = &v
}

// GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field value
func (o *CreatePitrConfigParams) GetSleepAfterMasterRestartMillis() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SleepAfterMasterRestartMillis
}

// GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field value
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SleepAfterMasterRestartMillis, true
}

// SetSleepAfterMasterRestartMillis sets field value
func (o *CreatePitrConfigParams) SetSleepAfterMasterRestartMillis(v int32) {
	o.SleepAfterMasterRestartMillis = v
}

// GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field value
func (o *CreatePitrConfigParams) GetSleepAfterTServerRestartMillis() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SleepAfterTServerRestartMillis
}

// GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field value
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SleepAfterTServerRestartMillis, true
}

// SetSleepAfterTServerRestartMillis sets field value
func (o *CreatePitrConfigParams) SetSleepAfterTServerRestartMillis(v int32) {
	o.SleepAfterTServerRestartMillis = v
}

// GetSourceXClusterConfigs returns the SourceXClusterConfigs field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetSourceXClusterConfigs() []string {
	if o == nil || o.SourceXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.SourceXClusterConfigs
}

// GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetSourceXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.SourceXClusterConfigs == nil {
		return nil, false
	}
	return o.SourceXClusterConfigs, true
}

// HasSourceXClusterConfigs returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasSourceXClusterConfigs() bool {
	if o != nil && o.SourceXClusterConfigs != nil {
		return true
	}

	return false
}

// SetSourceXClusterConfigs gets a reference to the given []string and assigns it to the SourceXClusterConfigs field.
func (o *CreatePitrConfigParams) SetSourceXClusterConfigs(v []string) {
	o.SourceXClusterConfigs = &v
}

// GetTargetXClusterConfigs returns the TargetXClusterConfigs field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetTargetXClusterConfigs() []string {
	if o == nil || o.TargetXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.TargetXClusterConfigs
}

// GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetTargetXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.TargetXClusterConfigs == nil {
		return nil, false
	}
	return o.TargetXClusterConfigs, true
}

// HasTargetXClusterConfigs returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasTargetXClusterConfigs() bool {
	if o != nil && o.TargetXClusterConfigs != nil {
		return true
	}

	return false
}

// SetTargetXClusterConfigs gets a reference to the given []string and assigns it to the TargetXClusterConfigs field.
func (o *CreatePitrConfigParams) SetTargetXClusterConfigs(v []string) {
	o.TargetXClusterConfigs = &v
}

// GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetYbPrevSoftwareVersion() string {
	if o == nil || o.YbPrevSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbPrevSoftwareVersion
}

// GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetYbPrevSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbPrevSoftwareVersion == nil {
		return nil, false
	}
	return o.YbPrevSoftwareVersion, true
}

// HasYbPrevSoftwareVersion returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasYbPrevSoftwareVersion() bool {
	if o != nil && o.YbPrevSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbPrevSoftwareVersion gets a reference to the given string and assigns it to the YbPrevSoftwareVersion field.
func (o *CreatePitrConfigParams) SetYbPrevSoftwareVersion(v string) {
	o.YbPrevSoftwareVersion = &v
}

// GetYbcInstalled returns the YbcInstalled field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetYbcInstalled() bool {
	if o == nil || o.YbcInstalled == nil {
		var ret bool
		return ret
	}
	return *o.YbcInstalled
}

// GetYbcInstalledOk returns a tuple with the YbcInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetYbcInstalledOk() (*bool, bool) {
	if o == nil || o.YbcInstalled == nil {
		return nil, false
	}
	return o.YbcInstalled, true
}

// HasYbcInstalled returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasYbcInstalled() bool {
	if o != nil && o.YbcInstalled != nil {
		return true
	}

	return false
}

// SetYbcInstalled gets a reference to the given bool and assigns it to the YbcInstalled field.
func (o *CreatePitrConfigParams) SetYbcInstalled(v bool) {
	o.YbcInstalled = &v
}

// GetYbcSoftwareVersion returns the YbcSoftwareVersion field value if set, zero value otherwise.
func (o *CreatePitrConfigParams) GetYbcSoftwareVersion() string {
	if o == nil || o.YbcSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbcSoftwareVersion
}

// GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePitrConfigParams) GetYbcSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbcSoftwareVersion == nil {
		return nil, false
	}
	return o.YbcSoftwareVersion, true
}

// HasYbcSoftwareVersion returns a boolean if a field has been set.
func (o *CreatePitrConfigParams) HasYbcSoftwareVersion() bool {
	if o != nil && o.YbcSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbcSoftwareVersion gets a reference to the given string and assigns it to the YbcSoftwareVersion field.
func (o *CreatePitrConfigParams) SetYbcSoftwareVersion(v string) {
	o.YbcSoftwareVersion = &v
}

func (o CreatePitrConfigParams) MarshalJSON() ([]byte, error) {
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
	if o.IntervalInSeconds != nil {
		toSerialize["intervalInSeconds"] = o.IntervalInSeconds
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NodeDetailsSet != nil {
		toSerialize["nodeDetailsSet"] = o.NodeDetailsSet
	}
	if o.NodeExporterUser != nil {
		toSerialize["nodeExporterUser"] = o.NodeExporterUser
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
	if o.RetentionPeriodInSeconds != nil {
		toSerialize["retentionPeriodInSeconds"] = o.RetentionPeriodInSeconds
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

type NullableCreatePitrConfigParams struct {
	value *CreatePitrConfigParams
	isSet bool
}

func (v NullableCreatePitrConfigParams) Get() *CreatePitrConfigParams {
	return v.value
}

func (v *NullableCreatePitrConfigParams) Set(val *CreatePitrConfigParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePitrConfigParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePitrConfigParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePitrConfigParams(val *CreatePitrConfigParams) *NullableCreatePitrConfigParams {
	return &NullableCreatePitrConfigParams{value: val, isSet: true}
}

func (v NullableCreatePitrConfigParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePitrConfigParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


