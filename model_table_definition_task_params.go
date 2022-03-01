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
)

// TableDefinitionTaskParams struct for TableDefinitionTaskParams
type TableDefinitionTaskParams struct {
	// Amazon Resource Name (ARN) of the CMK
	CmkArn *string `json:"cmkArn,omitempty"`
	CommunicationPorts *CommunicationPorts `json:"communicationPorts,omitempty"`
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`
	EncryptionAtRestConfig *EncryptionAtRestConfig `json:"encryptionAtRestConfig,omitempty"`
	// Error message
	ErrorString *string `json:"errorString,omitempty"`
	// Expected universe version
	ExpectedUniverseVersion *int32 `json:"expectedUniverseVersion,omitempty"`
	ExtraDependencies *ExtraDependencies `json:"extraDependencies,omitempty"`
	// Whether this task has been tried before
	FirstTry *bool `json:"firstTry,omitempty"`
	// Node details
	NodeDetailsSet *[]NodeDetails `json:"nodeDetailsSet,omitempty"`
	// Node exporter user
	NodeExporterUser *string `json:"nodeExporterUser,omitempty"`
	// Previous task UUID only if this task is a retry
	PreviousTaskUUID *string `json:"previousTaskUUID,omitempty"`
	// The source universe's xcluster replication relationships
	SourceXClusterConfigs *[]string `json:"sourceXClusterConfigs,omitempty"`
	TableDetails TableDetails `json:"tableDetails"`
	TableType string `json:"tableType"`
	TableUUID string `json:"tableUUID"`
	// The target universe's xcluster replication relationships
	TargetXClusterConfigs *[]string `json:"targetXClusterConfigs,omitempty"`
	// Associated universe UUID
	UniverseUUID *string `json:"universeUUID,omitempty"`
	// Previous software version
	YbPrevSoftwareVersion *string `json:"ybPrevSoftwareVersion,omitempty"`
}

// NewTableDefinitionTaskParams instantiates a new TableDefinitionTaskParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableDefinitionTaskParams(tableDetails TableDetails, tableType string, tableUUID string, ) *TableDefinitionTaskParams {
	this := TableDefinitionTaskParams{}
	this.TableDetails = tableDetails
	this.TableType = tableType
	this.TableUUID = tableUUID
	return &this
}

// NewTableDefinitionTaskParamsWithDefaults instantiates a new TableDefinitionTaskParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableDefinitionTaskParamsWithDefaults() *TableDefinitionTaskParams {
	this := TableDefinitionTaskParams{}
	return &this
}

// GetCmkArn returns the CmkArn field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetCmkArn() string {
	if o == nil || o.CmkArn == nil {
		var ret string
		return ret
	}
	return *o.CmkArn
}

// GetCmkArnOk returns a tuple with the CmkArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetCmkArnOk() (*string, bool) {
	if o == nil || o.CmkArn == nil {
		return nil, false
	}
	return o.CmkArn, true
}

// HasCmkArn returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasCmkArn() bool {
	if o != nil && o.CmkArn != nil {
		return true
	}

	return false
}

// SetCmkArn gets a reference to the given string and assigns it to the CmkArn field.
func (o *TableDefinitionTaskParams) SetCmkArn(v string) {
	o.CmkArn = &v
}

// GetCommunicationPorts returns the CommunicationPorts field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetCommunicationPorts() CommunicationPorts {
	if o == nil || o.CommunicationPorts == nil {
		var ret CommunicationPorts
		return ret
	}
	return *o.CommunicationPorts
}

// GetCommunicationPortsOk returns a tuple with the CommunicationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetCommunicationPortsOk() (*CommunicationPorts, bool) {
	if o == nil || o.CommunicationPorts == nil {
		return nil, false
	}
	return o.CommunicationPorts, true
}

// HasCommunicationPorts returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasCommunicationPorts() bool {
	if o != nil && o.CommunicationPorts != nil {
		return true
	}

	return false
}

// SetCommunicationPorts gets a reference to the given CommunicationPorts and assigns it to the CommunicationPorts field.
func (o *TableDefinitionTaskParams) SetCommunicationPorts(v CommunicationPorts) {
	o.CommunicationPorts = &v
}

// GetDeviceInfo returns the DeviceInfo field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetDeviceInfo() DeviceInfo {
	if o == nil || o.DeviceInfo == nil {
		var ret DeviceInfo
		return ret
	}
	return *o.DeviceInfo
}

// GetDeviceInfoOk returns a tuple with the DeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetDeviceInfoOk() (*DeviceInfo, bool) {
	if o == nil || o.DeviceInfo == nil {
		return nil, false
	}
	return o.DeviceInfo, true
}

// HasDeviceInfo returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasDeviceInfo() bool {
	if o != nil && o.DeviceInfo != nil {
		return true
	}

	return false
}

// SetDeviceInfo gets a reference to the given DeviceInfo and assigns it to the DeviceInfo field.
func (o *TableDefinitionTaskParams) SetDeviceInfo(v DeviceInfo) {
	o.DeviceInfo = &v
}

// GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig {
	if o == nil || o.EncryptionAtRestConfig == nil {
		var ret EncryptionAtRestConfig
		return ret
	}
	return *o.EncryptionAtRestConfig
}

// GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool) {
	if o == nil || o.EncryptionAtRestConfig == nil {
		return nil, false
	}
	return o.EncryptionAtRestConfig, true
}

// HasEncryptionAtRestConfig returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasEncryptionAtRestConfig() bool {
	if o != nil && o.EncryptionAtRestConfig != nil {
		return true
	}

	return false
}

// SetEncryptionAtRestConfig gets a reference to the given EncryptionAtRestConfig and assigns it to the EncryptionAtRestConfig field.
func (o *TableDefinitionTaskParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig) {
	o.EncryptionAtRestConfig = &v
}

// GetErrorString returns the ErrorString field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetErrorString() string {
	if o == nil || o.ErrorString == nil {
		var ret string
		return ret
	}
	return *o.ErrorString
}

// GetErrorStringOk returns a tuple with the ErrorString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetErrorStringOk() (*string, bool) {
	if o == nil || o.ErrorString == nil {
		return nil, false
	}
	return o.ErrorString, true
}

// HasErrorString returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasErrorString() bool {
	if o != nil && o.ErrorString != nil {
		return true
	}

	return false
}

// SetErrorString gets a reference to the given string and assigns it to the ErrorString field.
func (o *TableDefinitionTaskParams) SetErrorString(v string) {
	o.ErrorString = &v
}

// GetExpectedUniverseVersion returns the ExpectedUniverseVersion field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetExpectedUniverseVersion() int32 {
	if o == nil || o.ExpectedUniverseVersion == nil {
		var ret int32
		return ret
	}
	return *o.ExpectedUniverseVersion
}

// GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetExpectedUniverseVersionOk() (*int32, bool) {
	if o == nil || o.ExpectedUniverseVersion == nil {
		return nil, false
	}
	return o.ExpectedUniverseVersion, true
}

// HasExpectedUniverseVersion returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasExpectedUniverseVersion() bool {
	if o != nil && o.ExpectedUniverseVersion != nil {
		return true
	}

	return false
}

// SetExpectedUniverseVersion gets a reference to the given int32 and assigns it to the ExpectedUniverseVersion field.
func (o *TableDefinitionTaskParams) SetExpectedUniverseVersion(v int32) {
	o.ExpectedUniverseVersion = &v
}

// GetExtraDependencies returns the ExtraDependencies field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetExtraDependencies() ExtraDependencies {
	if o == nil || o.ExtraDependencies == nil {
		var ret ExtraDependencies
		return ret
	}
	return *o.ExtraDependencies
}

// GetExtraDependenciesOk returns a tuple with the ExtraDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetExtraDependenciesOk() (*ExtraDependencies, bool) {
	if o == nil || o.ExtraDependencies == nil {
		return nil, false
	}
	return o.ExtraDependencies, true
}

// HasExtraDependencies returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasExtraDependencies() bool {
	if o != nil && o.ExtraDependencies != nil {
		return true
	}

	return false
}

// SetExtraDependencies gets a reference to the given ExtraDependencies and assigns it to the ExtraDependencies field.
func (o *TableDefinitionTaskParams) SetExtraDependencies(v ExtraDependencies) {
	o.ExtraDependencies = &v
}

// GetFirstTry returns the FirstTry field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetFirstTry() bool {
	if o == nil || o.FirstTry == nil {
		var ret bool
		return ret
	}
	return *o.FirstTry
}

// GetFirstTryOk returns a tuple with the FirstTry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetFirstTryOk() (*bool, bool) {
	if o == nil || o.FirstTry == nil {
		return nil, false
	}
	return o.FirstTry, true
}

// HasFirstTry returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasFirstTry() bool {
	if o != nil && o.FirstTry != nil {
		return true
	}

	return false
}

// SetFirstTry gets a reference to the given bool and assigns it to the FirstTry field.
func (o *TableDefinitionTaskParams) SetFirstTry(v bool) {
	o.FirstTry = &v
}

// GetNodeDetailsSet returns the NodeDetailsSet field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetNodeDetailsSet() []NodeDetails {
	if o == nil || o.NodeDetailsSet == nil {
		var ret []NodeDetails
		return ret
	}
	return *o.NodeDetailsSet
}

// GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool) {
	if o == nil || o.NodeDetailsSet == nil {
		return nil, false
	}
	return o.NodeDetailsSet, true
}

// HasNodeDetailsSet returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasNodeDetailsSet() bool {
	if o != nil && o.NodeDetailsSet != nil {
		return true
	}

	return false
}

// SetNodeDetailsSet gets a reference to the given []NodeDetails and assigns it to the NodeDetailsSet field.
func (o *TableDefinitionTaskParams) SetNodeDetailsSet(v []NodeDetails) {
	o.NodeDetailsSet = &v
}

// GetNodeExporterUser returns the NodeExporterUser field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetNodeExporterUser() string {
	if o == nil || o.NodeExporterUser == nil {
		var ret string
		return ret
	}
	return *o.NodeExporterUser
}

// GetNodeExporterUserOk returns a tuple with the NodeExporterUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetNodeExporterUserOk() (*string, bool) {
	if o == nil || o.NodeExporterUser == nil {
		return nil, false
	}
	return o.NodeExporterUser, true
}

// HasNodeExporterUser returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasNodeExporterUser() bool {
	if o != nil && o.NodeExporterUser != nil {
		return true
	}

	return false
}

// SetNodeExporterUser gets a reference to the given string and assigns it to the NodeExporterUser field.
func (o *TableDefinitionTaskParams) SetNodeExporterUser(v string) {
	o.NodeExporterUser = &v
}

// GetPreviousTaskUUID returns the PreviousTaskUUID field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetPreviousTaskUUID() string {
	if o == nil || o.PreviousTaskUUID == nil {
		var ret string
		return ret
	}
	return *o.PreviousTaskUUID
}

// GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetPreviousTaskUUIDOk() (*string, bool) {
	if o == nil || o.PreviousTaskUUID == nil {
		return nil, false
	}
	return o.PreviousTaskUUID, true
}

// HasPreviousTaskUUID returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasPreviousTaskUUID() bool {
	if o != nil && o.PreviousTaskUUID != nil {
		return true
	}

	return false
}

// SetPreviousTaskUUID gets a reference to the given string and assigns it to the PreviousTaskUUID field.
func (o *TableDefinitionTaskParams) SetPreviousTaskUUID(v string) {
	o.PreviousTaskUUID = &v
}

// GetSourceXClusterConfigs returns the SourceXClusterConfigs field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetSourceXClusterConfigs() []string {
	if o == nil || o.SourceXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.SourceXClusterConfigs
}

// GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetSourceXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.SourceXClusterConfigs == nil {
		return nil, false
	}
	return o.SourceXClusterConfigs, true
}

// HasSourceXClusterConfigs returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasSourceXClusterConfigs() bool {
	if o != nil && o.SourceXClusterConfigs != nil {
		return true
	}

	return false
}

// SetSourceXClusterConfigs gets a reference to the given []string and assigns it to the SourceXClusterConfigs field.
func (o *TableDefinitionTaskParams) SetSourceXClusterConfigs(v []string) {
	o.SourceXClusterConfigs = &v
}

// GetTableDetails returns the TableDetails field value
func (o *TableDefinitionTaskParams) GetTableDetails() TableDetails {
	if o == nil  {
		var ret TableDetails
		return ret
	}

	return o.TableDetails
}

// GetTableDetailsOk returns a tuple with the TableDetails field value
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetTableDetailsOk() (*TableDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TableDetails, true
}

// SetTableDetails sets field value
func (o *TableDefinitionTaskParams) SetTableDetails(v TableDetails) {
	o.TableDetails = v
}

// GetTableType returns the TableType field value
func (o *TableDefinitionTaskParams) GetTableType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetTableTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TableType, true
}

// SetTableType sets field value
func (o *TableDefinitionTaskParams) SetTableType(v string) {
	o.TableType = v
}

// GetTableUUID returns the TableUUID field value
func (o *TableDefinitionTaskParams) GetTableUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TableUUID
}

// GetTableUUIDOk returns a tuple with the TableUUID field value
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetTableUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TableUUID, true
}

// SetTableUUID sets field value
func (o *TableDefinitionTaskParams) SetTableUUID(v string) {
	o.TableUUID = v
}

// GetTargetXClusterConfigs returns the TargetXClusterConfigs field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetTargetXClusterConfigs() []string {
	if o == nil || o.TargetXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.TargetXClusterConfigs
}

// GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetTargetXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.TargetXClusterConfigs == nil {
		return nil, false
	}
	return o.TargetXClusterConfigs, true
}

// HasTargetXClusterConfigs returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasTargetXClusterConfigs() bool {
	if o != nil && o.TargetXClusterConfigs != nil {
		return true
	}

	return false
}

// SetTargetXClusterConfigs gets a reference to the given []string and assigns it to the TargetXClusterConfigs field.
func (o *TableDefinitionTaskParams) SetTargetXClusterConfigs(v []string) {
	o.TargetXClusterConfigs = &v
}

// GetUniverseUUID returns the UniverseUUID field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetUniverseUUID() string {
	if o == nil || o.UniverseUUID == nil {
		var ret string
		return ret
	}
	return *o.UniverseUUID
}

// GetUniverseUUIDOk returns a tuple with the UniverseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetUniverseUUIDOk() (*string, bool) {
	if o == nil || o.UniverseUUID == nil {
		return nil, false
	}
	return o.UniverseUUID, true
}

// HasUniverseUUID returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasUniverseUUID() bool {
	if o != nil && o.UniverseUUID != nil {
		return true
	}

	return false
}

// SetUniverseUUID gets a reference to the given string and assigns it to the UniverseUUID field.
func (o *TableDefinitionTaskParams) SetUniverseUUID(v string) {
	o.UniverseUUID = &v
}

// GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field value if set, zero value otherwise.
func (o *TableDefinitionTaskParams) GetYbPrevSoftwareVersion() string {
	if o == nil || o.YbPrevSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbPrevSoftwareVersion
}

// GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDefinitionTaskParams) GetYbPrevSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbPrevSoftwareVersion == nil {
		return nil, false
	}
	return o.YbPrevSoftwareVersion, true
}

// HasYbPrevSoftwareVersion returns a boolean if a field has been set.
func (o *TableDefinitionTaskParams) HasYbPrevSoftwareVersion() bool {
	if o != nil && o.YbPrevSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbPrevSoftwareVersion gets a reference to the given string and assigns it to the YbPrevSoftwareVersion field.
func (o *TableDefinitionTaskParams) SetYbPrevSoftwareVersion(v string) {
	o.YbPrevSoftwareVersion = &v
}

func (o TableDefinitionTaskParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CmkArn != nil {
		toSerialize["cmkArn"] = o.CmkArn
	}
	if o.CommunicationPorts != nil {
		toSerialize["communicationPorts"] = o.CommunicationPorts
	}
	if o.DeviceInfo != nil {
		toSerialize["deviceInfo"] = o.DeviceInfo
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
	if o.FirstTry != nil {
		toSerialize["firstTry"] = o.FirstTry
	}
	if o.NodeDetailsSet != nil {
		toSerialize["nodeDetailsSet"] = o.NodeDetailsSet
	}
	if o.NodeExporterUser != nil {
		toSerialize["nodeExporterUser"] = o.NodeExporterUser
	}
	if o.PreviousTaskUUID != nil {
		toSerialize["previousTaskUUID"] = o.PreviousTaskUUID
	}
	if o.SourceXClusterConfigs != nil {
		toSerialize["sourceXClusterConfigs"] = o.SourceXClusterConfigs
	}
	if true {
		toSerialize["tableDetails"] = o.TableDetails
	}
	if true {
		toSerialize["tableType"] = o.TableType
	}
	if true {
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
	return json.Marshal(toSerialize)
}

type NullableTableDefinitionTaskParams struct {
	value *TableDefinitionTaskParams
	isSet bool
}

func (v NullableTableDefinitionTaskParams) Get() *TableDefinitionTaskParams {
	return v.value
}

func (v *NullableTableDefinitionTaskParams) Set(val *TableDefinitionTaskParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTableDefinitionTaskParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTableDefinitionTaskParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableDefinitionTaskParams(val *TableDefinitionTaskParams) *NullableTableDefinitionTaskParams {
	return &NullableTableDefinitionTaskParams{value: val, isSet: true}
}

func (v NullableTableDefinitionTaskParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableDefinitionTaskParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


