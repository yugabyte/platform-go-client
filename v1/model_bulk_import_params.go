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

// BulkImportParams Bulk import parameters
type BulkImportParams struct {
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
	// Instance count
	InstanceCount *int32 `json:"instanceCount,omitempty"`
	// Key space
	Keyspace *string `json:"keyspace,omitempty"`
	// Node details
	NodeDetailsSet *[]NodeDetails `json:"nodeDetailsSet,omitempty"`
	// Node exporter user
	NodeExporterUser *string `json:"nodeExporterUser,omitempty"`
	PlatformUrl string `json:"platformUrl"`
	PlatformVersion string `json:"platformVersion"`
	// Previous task UUID of a retry
	PreviousTaskUUID *string `json:"previousTaskUUID,omitempty"`
	// S3 bucket URL
	S3Bucket string `json:"s3Bucket"`
	SleepAfterMasterRestartMillis int32 `json:"sleepAfterMasterRestartMillis"`
	SleepAfterTServerRestartMillis int32 `json:"sleepAfterTServerRestartMillis"`
	// The source universe's xcluster replication relationships
	SourceXClusterConfigs *[]string `json:"sourceXClusterConfigs,omitempty"`
	// Is SSE
	Sse *bool `json:"sse,omitempty"`
	// Table name
	TableName *string `json:"tableName,omitempty"`
	// Table UUID
	TableUUID *string `json:"tableUUID,omitempty"`
	// The target universe's xcluster replication relationships
	TargetXClusterConfigs *[]string `json:"targetXClusterConfigs,omitempty"`
	// Associated universe UUID
	UniverseUUID *string `json:"universeUUID,omitempty"`
	// Previous software version
	YbPrevSoftwareVersion *string `json:"ybPrevSoftwareVersion,omitempty"`
	YbcInstalled *bool `json:"ybcInstalled,omitempty"`
	YbcSoftwareVersion *string `json:"ybcSoftwareVersion,omitempty"`
}

// NewBulkImportParams instantiates a new BulkImportParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkImportParams(creatingUser Users, platformUrl string, platformVersion string, s3Bucket string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32) *BulkImportParams {
	this := BulkImportParams{}
	this.CreatingUser = creatingUser
	this.PlatformUrl = platformUrl
	this.PlatformVersion = platformVersion
	this.S3Bucket = s3Bucket
	this.SleepAfterMasterRestartMillis = sleepAfterMasterRestartMillis
	this.SleepAfterTServerRestartMillis = sleepAfterTServerRestartMillis
	return &this
}

// NewBulkImportParamsWithDefaults instantiates a new BulkImportParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkImportParamsWithDefaults() *BulkImportParams {
	this := BulkImportParams{}
	return &this
}

// GetCmkArn returns the CmkArn field value if set, zero value otherwise.
func (o *BulkImportParams) GetCmkArn() string {
	if o == nil || o.CmkArn == nil {
		var ret string
		return ret
	}
	return *o.CmkArn
}

// GetCmkArnOk returns a tuple with the CmkArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetCmkArnOk() (*string, bool) {
	if o == nil || o.CmkArn == nil {
		return nil, false
	}
	return o.CmkArn, true
}

// HasCmkArn returns a boolean if a field has been set.
func (o *BulkImportParams) HasCmkArn() bool {
	if o != nil && o.CmkArn != nil {
		return true
	}

	return false
}

// SetCmkArn gets a reference to the given string and assigns it to the CmkArn field.
func (o *BulkImportParams) SetCmkArn(v string) {
	o.CmkArn = &v
}

// GetCommunicationPorts returns the CommunicationPorts field value if set, zero value otherwise.
func (o *BulkImportParams) GetCommunicationPorts() CommunicationPorts {
	if o == nil || o.CommunicationPorts == nil {
		var ret CommunicationPorts
		return ret
	}
	return *o.CommunicationPorts
}

// GetCommunicationPortsOk returns a tuple with the CommunicationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetCommunicationPortsOk() (*CommunicationPorts, bool) {
	if o == nil || o.CommunicationPorts == nil {
		return nil, false
	}
	return o.CommunicationPorts, true
}

// HasCommunicationPorts returns a boolean if a field has been set.
func (o *BulkImportParams) HasCommunicationPorts() bool {
	if o != nil && o.CommunicationPorts != nil {
		return true
	}

	return false
}

// SetCommunicationPorts gets a reference to the given CommunicationPorts and assigns it to the CommunicationPorts field.
func (o *BulkImportParams) SetCommunicationPorts(v CommunicationPorts) {
	o.CommunicationPorts = &v
}

// GetCreatingUser returns the CreatingUser field value
func (o *BulkImportParams) GetCreatingUser() Users {
	if o == nil {
		var ret Users
		return ret
	}

	return o.CreatingUser
}

// GetCreatingUserOk returns a tuple with the CreatingUser field value
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetCreatingUserOk() (*Users, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatingUser, true
}

// SetCreatingUser sets field value
func (o *BulkImportParams) SetCreatingUser(v Users) {
	o.CreatingUser = v
}

// GetDeviceInfo returns the DeviceInfo field value if set, zero value otherwise.
func (o *BulkImportParams) GetDeviceInfo() DeviceInfo {
	if o == nil || o.DeviceInfo == nil {
		var ret DeviceInfo
		return ret
	}
	return *o.DeviceInfo
}

// GetDeviceInfoOk returns a tuple with the DeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetDeviceInfoOk() (*DeviceInfo, bool) {
	if o == nil || o.DeviceInfo == nil {
		return nil, false
	}
	return o.DeviceInfo, true
}

// HasDeviceInfo returns a boolean if a field has been set.
func (o *BulkImportParams) HasDeviceInfo() bool {
	if o != nil && o.DeviceInfo != nil {
		return true
	}

	return false
}

// SetDeviceInfo gets a reference to the given DeviceInfo and assigns it to the DeviceInfo field.
func (o *BulkImportParams) SetDeviceInfo(v DeviceInfo) {
	o.DeviceInfo = &v
}

// GetEnableYbc returns the EnableYbc field value if set, zero value otherwise.
func (o *BulkImportParams) GetEnableYbc() bool {
	if o == nil || o.EnableYbc == nil {
		var ret bool
		return ret
	}
	return *o.EnableYbc
}

// GetEnableYbcOk returns a tuple with the EnableYbc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetEnableYbcOk() (*bool, bool) {
	if o == nil || o.EnableYbc == nil {
		return nil, false
	}
	return o.EnableYbc, true
}

// HasEnableYbc returns a boolean if a field has been set.
func (o *BulkImportParams) HasEnableYbc() bool {
	if o != nil && o.EnableYbc != nil {
		return true
	}

	return false
}

// SetEnableYbc gets a reference to the given bool and assigns it to the EnableYbc field.
func (o *BulkImportParams) SetEnableYbc(v bool) {
	o.EnableYbc = &v
}

// GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field value if set, zero value otherwise.
func (o *BulkImportParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig {
	if o == nil || o.EncryptionAtRestConfig == nil {
		var ret EncryptionAtRestConfig
		return ret
	}
	return *o.EncryptionAtRestConfig
}

// GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool) {
	if o == nil || o.EncryptionAtRestConfig == nil {
		return nil, false
	}
	return o.EncryptionAtRestConfig, true
}

// HasEncryptionAtRestConfig returns a boolean if a field has been set.
func (o *BulkImportParams) HasEncryptionAtRestConfig() bool {
	if o != nil && o.EncryptionAtRestConfig != nil {
		return true
	}

	return false
}

// SetEncryptionAtRestConfig gets a reference to the given EncryptionAtRestConfig and assigns it to the EncryptionAtRestConfig field.
func (o *BulkImportParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig) {
	o.EncryptionAtRestConfig = &v
}

// GetErrorString returns the ErrorString field value if set, zero value otherwise.
func (o *BulkImportParams) GetErrorString() string {
	if o == nil || o.ErrorString == nil {
		var ret string
		return ret
	}
	return *o.ErrorString
}

// GetErrorStringOk returns a tuple with the ErrorString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetErrorStringOk() (*string, bool) {
	if o == nil || o.ErrorString == nil {
		return nil, false
	}
	return o.ErrorString, true
}

// HasErrorString returns a boolean if a field has been set.
func (o *BulkImportParams) HasErrorString() bool {
	if o != nil && o.ErrorString != nil {
		return true
	}

	return false
}

// SetErrorString gets a reference to the given string and assigns it to the ErrorString field.
func (o *BulkImportParams) SetErrorString(v string) {
	o.ErrorString = &v
}

// GetExpectedUniverseVersion returns the ExpectedUniverseVersion field value if set, zero value otherwise.
func (o *BulkImportParams) GetExpectedUniverseVersion() int32 {
	if o == nil || o.ExpectedUniverseVersion == nil {
		var ret int32
		return ret
	}
	return *o.ExpectedUniverseVersion
}

// GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetExpectedUniverseVersionOk() (*int32, bool) {
	if o == nil || o.ExpectedUniverseVersion == nil {
		return nil, false
	}
	return o.ExpectedUniverseVersion, true
}

// HasExpectedUniverseVersion returns a boolean if a field has been set.
func (o *BulkImportParams) HasExpectedUniverseVersion() bool {
	if o != nil && o.ExpectedUniverseVersion != nil {
		return true
	}

	return false
}

// SetExpectedUniverseVersion gets a reference to the given int32 and assigns it to the ExpectedUniverseVersion field.
func (o *BulkImportParams) SetExpectedUniverseVersion(v int32) {
	o.ExpectedUniverseVersion = &v
}

// GetExtraDependencies returns the ExtraDependencies field value if set, zero value otherwise.
func (o *BulkImportParams) GetExtraDependencies() ExtraDependencies {
	if o == nil || o.ExtraDependencies == nil {
		var ret ExtraDependencies
		return ret
	}
	return *o.ExtraDependencies
}

// GetExtraDependenciesOk returns a tuple with the ExtraDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetExtraDependenciesOk() (*ExtraDependencies, bool) {
	if o == nil || o.ExtraDependencies == nil {
		return nil, false
	}
	return o.ExtraDependencies, true
}

// HasExtraDependencies returns a boolean if a field has been set.
func (o *BulkImportParams) HasExtraDependencies() bool {
	if o != nil && o.ExtraDependencies != nil {
		return true
	}

	return false
}

// SetExtraDependencies gets a reference to the given ExtraDependencies and assigns it to the ExtraDependencies field.
func (o *BulkImportParams) SetExtraDependencies(v ExtraDependencies) {
	o.ExtraDependencies = &v
}

// GetInstallYbc returns the InstallYbc field value if set, zero value otherwise.
func (o *BulkImportParams) GetInstallYbc() bool {
	if o == nil || o.InstallYbc == nil {
		var ret bool
		return ret
	}
	return *o.InstallYbc
}

// GetInstallYbcOk returns a tuple with the InstallYbc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetInstallYbcOk() (*bool, bool) {
	if o == nil || o.InstallYbc == nil {
		return nil, false
	}
	return o.InstallYbc, true
}

// HasInstallYbc returns a boolean if a field has been set.
func (o *BulkImportParams) HasInstallYbc() bool {
	if o != nil && o.InstallYbc != nil {
		return true
	}

	return false
}

// SetInstallYbc gets a reference to the given bool and assigns it to the InstallYbc field.
func (o *BulkImportParams) SetInstallYbc(v bool) {
	o.InstallYbc = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *BulkImportParams) GetInstanceCount() int32 {
	if o == nil || o.InstanceCount == nil {
		var ret int32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetInstanceCountOk() (*int32, bool) {
	if o == nil || o.InstanceCount == nil {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *BulkImportParams) HasInstanceCount() bool {
	if o != nil && o.InstanceCount != nil {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *BulkImportParams) SetInstanceCount(v int32) {
	o.InstanceCount = &v
}

// GetKeyspace returns the Keyspace field value if set, zero value otherwise.
func (o *BulkImportParams) GetKeyspace() string {
	if o == nil || o.Keyspace == nil {
		var ret string
		return ret
	}
	return *o.Keyspace
}

// GetKeyspaceOk returns a tuple with the Keyspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetKeyspaceOk() (*string, bool) {
	if o == nil || o.Keyspace == nil {
		return nil, false
	}
	return o.Keyspace, true
}

// HasKeyspace returns a boolean if a field has been set.
func (o *BulkImportParams) HasKeyspace() bool {
	if o != nil && o.Keyspace != nil {
		return true
	}

	return false
}

// SetKeyspace gets a reference to the given string and assigns it to the Keyspace field.
func (o *BulkImportParams) SetKeyspace(v string) {
	o.Keyspace = &v
}

// GetNodeDetailsSet returns the NodeDetailsSet field value if set, zero value otherwise.
func (o *BulkImportParams) GetNodeDetailsSet() []NodeDetails {
	if o == nil || o.NodeDetailsSet == nil {
		var ret []NodeDetails
		return ret
	}
	return *o.NodeDetailsSet
}

// GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool) {
	if o == nil || o.NodeDetailsSet == nil {
		return nil, false
	}
	return o.NodeDetailsSet, true
}

// HasNodeDetailsSet returns a boolean if a field has been set.
func (o *BulkImportParams) HasNodeDetailsSet() bool {
	if o != nil && o.NodeDetailsSet != nil {
		return true
	}

	return false
}

// SetNodeDetailsSet gets a reference to the given []NodeDetails and assigns it to the NodeDetailsSet field.
func (o *BulkImportParams) SetNodeDetailsSet(v []NodeDetails) {
	o.NodeDetailsSet = &v
}

// GetNodeExporterUser returns the NodeExporterUser field value if set, zero value otherwise.
func (o *BulkImportParams) GetNodeExporterUser() string {
	if o == nil || o.NodeExporterUser == nil {
		var ret string
		return ret
	}
	return *o.NodeExporterUser
}

// GetNodeExporterUserOk returns a tuple with the NodeExporterUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetNodeExporterUserOk() (*string, bool) {
	if o == nil || o.NodeExporterUser == nil {
		return nil, false
	}
	return o.NodeExporterUser, true
}

// HasNodeExporterUser returns a boolean if a field has been set.
func (o *BulkImportParams) HasNodeExporterUser() bool {
	if o != nil && o.NodeExporterUser != nil {
		return true
	}

	return false
}

// SetNodeExporterUser gets a reference to the given string and assigns it to the NodeExporterUser field.
func (o *BulkImportParams) SetNodeExporterUser(v string) {
	o.NodeExporterUser = &v
}

// GetPlatformUrl returns the PlatformUrl field value
func (o *BulkImportParams) GetPlatformUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformUrl
}

// GetPlatformUrlOk returns a tuple with the PlatformUrl field value
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetPlatformUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlatformUrl, true
}

// SetPlatformUrl sets field value
func (o *BulkImportParams) SetPlatformUrl(v string) {
	o.PlatformUrl = v
}

// GetPlatformVersion returns the PlatformVersion field value
func (o *BulkImportParams) GetPlatformVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformVersion
}

// GetPlatformVersionOk returns a tuple with the PlatformVersion field value
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetPlatformVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlatformVersion, true
}

// SetPlatformVersion sets field value
func (o *BulkImportParams) SetPlatformVersion(v string) {
	o.PlatformVersion = v
}

// GetPreviousTaskUUID returns the PreviousTaskUUID field value if set, zero value otherwise.
func (o *BulkImportParams) GetPreviousTaskUUID() string {
	if o == nil || o.PreviousTaskUUID == nil {
		var ret string
		return ret
	}
	return *o.PreviousTaskUUID
}

// GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetPreviousTaskUUIDOk() (*string, bool) {
	if o == nil || o.PreviousTaskUUID == nil {
		return nil, false
	}
	return o.PreviousTaskUUID, true
}

// HasPreviousTaskUUID returns a boolean if a field has been set.
func (o *BulkImportParams) HasPreviousTaskUUID() bool {
	if o != nil && o.PreviousTaskUUID != nil {
		return true
	}

	return false
}

// SetPreviousTaskUUID gets a reference to the given string and assigns it to the PreviousTaskUUID field.
func (o *BulkImportParams) SetPreviousTaskUUID(v string) {
	o.PreviousTaskUUID = &v
}

// GetS3Bucket returns the S3Bucket field value
func (o *BulkImportParams) GetS3Bucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.S3Bucket
}

// GetS3BucketOk returns a tuple with the S3Bucket field value
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetS3BucketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.S3Bucket, true
}

// SetS3Bucket sets field value
func (o *BulkImportParams) SetS3Bucket(v string) {
	o.S3Bucket = v
}

// GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field value
func (o *BulkImportParams) GetSleepAfterMasterRestartMillis() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SleepAfterMasterRestartMillis
}

// GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field value
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SleepAfterMasterRestartMillis, true
}

// SetSleepAfterMasterRestartMillis sets field value
func (o *BulkImportParams) SetSleepAfterMasterRestartMillis(v int32) {
	o.SleepAfterMasterRestartMillis = v
}

// GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field value
func (o *BulkImportParams) GetSleepAfterTServerRestartMillis() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SleepAfterTServerRestartMillis
}

// GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field value
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SleepAfterTServerRestartMillis, true
}

// SetSleepAfterTServerRestartMillis sets field value
func (o *BulkImportParams) SetSleepAfterTServerRestartMillis(v int32) {
	o.SleepAfterTServerRestartMillis = v
}

// GetSourceXClusterConfigs returns the SourceXClusterConfigs field value if set, zero value otherwise.
func (o *BulkImportParams) GetSourceXClusterConfigs() []string {
	if o == nil || o.SourceXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.SourceXClusterConfigs
}

// GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetSourceXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.SourceXClusterConfigs == nil {
		return nil, false
	}
	return o.SourceXClusterConfigs, true
}

// HasSourceXClusterConfigs returns a boolean if a field has been set.
func (o *BulkImportParams) HasSourceXClusterConfigs() bool {
	if o != nil && o.SourceXClusterConfigs != nil {
		return true
	}

	return false
}

// SetSourceXClusterConfigs gets a reference to the given []string and assigns it to the SourceXClusterConfigs field.
func (o *BulkImportParams) SetSourceXClusterConfigs(v []string) {
	o.SourceXClusterConfigs = &v
}

// GetSse returns the Sse field value if set, zero value otherwise.
func (o *BulkImportParams) GetSse() bool {
	if o == nil || o.Sse == nil {
		var ret bool
		return ret
	}
	return *o.Sse
}

// GetSseOk returns a tuple with the Sse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetSseOk() (*bool, bool) {
	if o == nil || o.Sse == nil {
		return nil, false
	}
	return o.Sse, true
}

// HasSse returns a boolean if a field has been set.
func (o *BulkImportParams) HasSse() bool {
	if o != nil && o.Sse != nil {
		return true
	}

	return false
}

// SetSse gets a reference to the given bool and assigns it to the Sse field.
func (o *BulkImportParams) SetSse(v bool) {
	o.Sse = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *BulkImportParams) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *BulkImportParams) HasTableName() bool {
	if o != nil && o.TableName != nil {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *BulkImportParams) SetTableName(v string) {
	o.TableName = &v
}

// GetTableUUID returns the TableUUID field value if set, zero value otherwise.
func (o *BulkImportParams) GetTableUUID() string {
	if o == nil || o.TableUUID == nil {
		var ret string
		return ret
	}
	return *o.TableUUID
}

// GetTableUUIDOk returns a tuple with the TableUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetTableUUIDOk() (*string, bool) {
	if o == nil || o.TableUUID == nil {
		return nil, false
	}
	return o.TableUUID, true
}

// HasTableUUID returns a boolean if a field has been set.
func (o *BulkImportParams) HasTableUUID() bool {
	if o != nil && o.TableUUID != nil {
		return true
	}

	return false
}

// SetTableUUID gets a reference to the given string and assigns it to the TableUUID field.
func (o *BulkImportParams) SetTableUUID(v string) {
	o.TableUUID = &v
}

// GetTargetXClusterConfigs returns the TargetXClusterConfigs field value if set, zero value otherwise.
func (o *BulkImportParams) GetTargetXClusterConfigs() []string {
	if o == nil || o.TargetXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.TargetXClusterConfigs
}

// GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetTargetXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.TargetXClusterConfigs == nil {
		return nil, false
	}
	return o.TargetXClusterConfigs, true
}

// HasTargetXClusterConfigs returns a boolean if a field has been set.
func (o *BulkImportParams) HasTargetXClusterConfigs() bool {
	if o != nil && o.TargetXClusterConfigs != nil {
		return true
	}

	return false
}

// SetTargetXClusterConfigs gets a reference to the given []string and assigns it to the TargetXClusterConfigs field.
func (o *BulkImportParams) SetTargetXClusterConfigs(v []string) {
	o.TargetXClusterConfigs = &v
}

// GetUniverseUUID returns the UniverseUUID field value if set, zero value otherwise.
func (o *BulkImportParams) GetUniverseUUID() string {
	if o == nil || o.UniverseUUID == nil {
		var ret string
		return ret
	}
	return *o.UniverseUUID
}

// GetUniverseUUIDOk returns a tuple with the UniverseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetUniverseUUIDOk() (*string, bool) {
	if o == nil || o.UniverseUUID == nil {
		return nil, false
	}
	return o.UniverseUUID, true
}

// HasUniverseUUID returns a boolean if a field has been set.
func (o *BulkImportParams) HasUniverseUUID() bool {
	if o != nil && o.UniverseUUID != nil {
		return true
	}

	return false
}

// SetUniverseUUID gets a reference to the given string and assigns it to the UniverseUUID field.
func (o *BulkImportParams) SetUniverseUUID(v string) {
	o.UniverseUUID = &v
}

// GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field value if set, zero value otherwise.
func (o *BulkImportParams) GetYbPrevSoftwareVersion() string {
	if o == nil || o.YbPrevSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbPrevSoftwareVersion
}

// GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetYbPrevSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbPrevSoftwareVersion == nil {
		return nil, false
	}
	return o.YbPrevSoftwareVersion, true
}

// HasYbPrevSoftwareVersion returns a boolean if a field has been set.
func (o *BulkImportParams) HasYbPrevSoftwareVersion() bool {
	if o != nil && o.YbPrevSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbPrevSoftwareVersion gets a reference to the given string and assigns it to the YbPrevSoftwareVersion field.
func (o *BulkImportParams) SetYbPrevSoftwareVersion(v string) {
	o.YbPrevSoftwareVersion = &v
}

// GetYbcInstalled returns the YbcInstalled field value if set, zero value otherwise.
func (o *BulkImportParams) GetYbcInstalled() bool {
	if o == nil || o.YbcInstalled == nil {
		var ret bool
		return ret
	}
	return *o.YbcInstalled
}

// GetYbcInstalledOk returns a tuple with the YbcInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetYbcInstalledOk() (*bool, bool) {
	if o == nil || o.YbcInstalled == nil {
		return nil, false
	}
	return o.YbcInstalled, true
}

// HasYbcInstalled returns a boolean if a field has been set.
func (o *BulkImportParams) HasYbcInstalled() bool {
	if o != nil && o.YbcInstalled != nil {
		return true
	}

	return false
}

// SetYbcInstalled gets a reference to the given bool and assigns it to the YbcInstalled field.
func (o *BulkImportParams) SetYbcInstalled(v bool) {
	o.YbcInstalled = &v
}

// GetYbcSoftwareVersion returns the YbcSoftwareVersion field value if set, zero value otherwise.
func (o *BulkImportParams) GetYbcSoftwareVersion() string {
	if o == nil || o.YbcSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbcSoftwareVersion
}

// GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkImportParams) GetYbcSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbcSoftwareVersion == nil {
		return nil, false
	}
	return o.YbcSoftwareVersion, true
}

// HasYbcSoftwareVersion returns a boolean if a field has been set.
func (o *BulkImportParams) HasYbcSoftwareVersion() bool {
	if o != nil && o.YbcSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbcSoftwareVersion gets a reference to the given string and assigns it to the YbcSoftwareVersion field.
func (o *BulkImportParams) SetYbcSoftwareVersion(v string) {
	o.YbcSoftwareVersion = &v
}

func (o BulkImportParams) MarshalJSON() ([]byte, error) {
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
	if o.InstanceCount != nil {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if o.Keyspace != nil {
		toSerialize["keyspace"] = o.Keyspace
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
	if true {
		toSerialize["s3Bucket"] = o.S3Bucket
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
	if o.Sse != nil {
		toSerialize["sse"] = o.Sse
	}
	if o.TableName != nil {
		toSerialize["tableName"] = o.TableName
	}
	if o.TableUUID != nil {
		toSerialize["tableUUID"] = o.TableUUID
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

type NullableBulkImportParams struct {
	value *BulkImportParams
	isSet bool
}

func (v NullableBulkImportParams) Get() *BulkImportParams {
	return v.value
}

func (v *NullableBulkImportParams) Set(val *BulkImportParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkImportParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkImportParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkImportParams(val *BulkImportParams) *NullableBulkImportParams {
	return &NullableBulkImportParams{value: val, isSet: true}
}

func (v NullableBulkImportParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkImportParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


