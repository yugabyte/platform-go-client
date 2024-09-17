# RestoreSnapshotScheduleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CmkArn** | Pointer to **string** | Amazon Resource Name (ARN) of the CMK | [optional] 
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**CreatingUser** | [**Users**](Users.md) |  | 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**EnableYbc** | Pointer to **bool** |  | [optional] 
**EncryptionAtRestConfig** | Pointer to [**EncryptionAtRestConfig**](EncryptionAtRestConfig.md) |  | [optional] 
**ErrorString** | Pointer to **string** | Error message | [optional] 
**ExpectedUniverseVersion** | Pointer to **int32** | Expected universe version | [optional] 
**ExtraDependencies** | Pointer to [**ExtraDependencies**](ExtraDependencies.md) |  | [optional] 
**InstallYbc** | Pointer to **bool** |  | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**PitrConfigUUID** | Pointer to **string** | PITR Config UUID | [optional] 
**PlatformUrl** | **string** |  | 
**PlatformVersion** | **string** |  | 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID of a retry | [optional] 
**RestoreTimeInMillis** | Pointer to **int64** | Restore Time In millis | [optional] 
**SleepAfterMasterRestartMillis** | **int32** |  | 
**SleepAfterTServerRestartMillis** | **int32** |  | 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] [readonly] 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] [readonly] 
**UniverseUUID** | Pointer to **string** | Universe UUID | [optional] 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 
**YbcInstalled** | Pointer to **bool** |  | [optional] 
**YbcSoftwareVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewRestoreSnapshotScheduleParams

`func NewRestoreSnapshotScheduleParams(creatingUser Users, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, ) *RestoreSnapshotScheduleParams`

NewRestoreSnapshotScheduleParams instantiates a new RestoreSnapshotScheduleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSnapshotScheduleParamsWithDefaults

`func NewRestoreSnapshotScheduleParamsWithDefaults() *RestoreSnapshotScheduleParams`

NewRestoreSnapshotScheduleParamsWithDefaults instantiates a new RestoreSnapshotScheduleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmkArn

`func (o *RestoreSnapshotScheduleParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *RestoreSnapshotScheduleParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *RestoreSnapshotScheduleParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *RestoreSnapshotScheduleParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *RestoreSnapshotScheduleParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *RestoreSnapshotScheduleParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *RestoreSnapshotScheduleParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *RestoreSnapshotScheduleParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *RestoreSnapshotScheduleParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *RestoreSnapshotScheduleParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *RestoreSnapshotScheduleParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetDeviceInfo

`func (o *RestoreSnapshotScheduleParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *RestoreSnapshotScheduleParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *RestoreSnapshotScheduleParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *RestoreSnapshotScheduleParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableYbc

`func (o *RestoreSnapshotScheduleParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *RestoreSnapshotScheduleParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *RestoreSnapshotScheduleParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *RestoreSnapshotScheduleParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *RestoreSnapshotScheduleParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *RestoreSnapshotScheduleParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *RestoreSnapshotScheduleParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *RestoreSnapshotScheduleParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *RestoreSnapshotScheduleParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *RestoreSnapshotScheduleParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *RestoreSnapshotScheduleParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *RestoreSnapshotScheduleParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *RestoreSnapshotScheduleParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *RestoreSnapshotScheduleParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *RestoreSnapshotScheduleParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *RestoreSnapshotScheduleParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *RestoreSnapshotScheduleParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *RestoreSnapshotScheduleParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *RestoreSnapshotScheduleParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *RestoreSnapshotScheduleParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetInstallYbc

`func (o *RestoreSnapshotScheduleParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *RestoreSnapshotScheduleParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *RestoreSnapshotScheduleParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *RestoreSnapshotScheduleParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *RestoreSnapshotScheduleParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *RestoreSnapshotScheduleParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *RestoreSnapshotScheduleParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *RestoreSnapshotScheduleParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *RestoreSnapshotScheduleParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *RestoreSnapshotScheduleParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *RestoreSnapshotScheduleParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *RestoreSnapshotScheduleParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetPitrConfigUUID

`func (o *RestoreSnapshotScheduleParams) GetPitrConfigUUID() string`

GetPitrConfigUUID returns the PitrConfigUUID field if non-nil, zero value otherwise.

### GetPitrConfigUUIDOk

`func (o *RestoreSnapshotScheduleParams) GetPitrConfigUUIDOk() (*string, bool)`

GetPitrConfigUUIDOk returns a tuple with the PitrConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrConfigUUID

`func (o *RestoreSnapshotScheduleParams) SetPitrConfigUUID(v string)`

SetPitrConfigUUID sets PitrConfigUUID field to given value.

### HasPitrConfigUUID

`func (o *RestoreSnapshotScheduleParams) HasPitrConfigUUID() bool`

HasPitrConfigUUID returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *RestoreSnapshotScheduleParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *RestoreSnapshotScheduleParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *RestoreSnapshotScheduleParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPlatformVersion

`func (o *RestoreSnapshotScheduleParams) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *RestoreSnapshotScheduleParams) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *RestoreSnapshotScheduleParams) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.


### GetPreviousTaskUUID

`func (o *RestoreSnapshotScheduleParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *RestoreSnapshotScheduleParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *RestoreSnapshotScheduleParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *RestoreSnapshotScheduleParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRestoreTimeInMillis

`func (o *RestoreSnapshotScheduleParams) GetRestoreTimeInMillis() int64`

GetRestoreTimeInMillis returns the RestoreTimeInMillis field if non-nil, zero value otherwise.

### GetRestoreTimeInMillisOk

`func (o *RestoreSnapshotScheduleParams) GetRestoreTimeInMillisOk() (*int64, bool)`

GetRestoreTimeInMillisOk returns a tuple with the RestoreTimeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeInMillis

`func (o *RestoreSnapshotScheduleParams) SetRestoreTimeInMillis(v int64)`

SetRestoreTimeInMillis sets RestoreTimeInMillis field to given value.

### HasRestoreTimeInMillis

`func (o *RestoreSnapshotScheduleParams) HasRestoreTimeInMillis() bool`

HasRestoreTimeInMillis returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *RestoreSnapshotScheduleParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *RestoreSnapshotScheduleParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *RestoreSnapshotScheduleParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *RestoreSnapshotScheduleParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *RestoreSnapshotScheduleParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *RestoreSnapshotScheduleParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSourceXClusterConfigs

`func (o *RestoreSnapshotScheduleParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *RestoreSnapshotScheduleParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *RestoreSnapshotScheduleParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *RestoreSnapshotScheduleParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *RestoreSnapshotScheduleParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *RestoreSnapshotScheduleParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *RestoreSnapshotScheduleParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *RestoreSnapshotScheduleParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *RestoreSnapshotScheduleParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *RestoreSnapshotScheduleParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *RestoreSnapshotScheduleParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *RestoreSnapshotScheduleParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *RestoreSnapshotScheduleParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *RestoreSnapshotScheduleParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *RestoreSnapshotScheduleParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *RestoreSnapshotScheduleParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *RestoreSnapshotScheduleParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *RestoreSnapshotScheduleParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *RestoreSnapshotScheduleParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *RestoreSnapshotScheduleParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *RestoreSnapshotScheduleParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *RestoreSnapshotScheduleParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *RestoreSnapshotScheduleParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *RestoreSnapshotScheduleParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


