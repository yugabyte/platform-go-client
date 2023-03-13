# SystemdUpgradeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInsecure** | Pointer to **bool** |  | [optional] 
**Capability** | Pointer to **string** |  | [optional] 
**ClientRootCA** | Pointer to **string** |  | [optional] 
**Clusters** | [**[]Cluster**](Cluster.md) |  | 
**CmkArn** | Pointer to **string** | Amazon Resource Name (ARN) of the CMK | [optional] 
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**CreatingUser** | [**Users**](Users.md) |  | 
**CurrentClusterType** | Pointer to **string** |  | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**EnableYbc** | Pointer to **bool** |  | [optional] 
**EncryptionAtRestConfig** | Pointer to [**EncryptionAtRestConfig**](EncryptionAtRestConfig.md) |  | [optional] 
**ErrorString** | Pointer to **string** | Error message | [optional] 
**ExpectedUniverseVersion** | Pointer to **int32** | Expected universe version | [optional] 
**ExtraDependencies** | Pointer to [**ExtraDependencies**](ExtraDependencies.md) |  | [optional] 
**ImportedState** | Pointer to **string** |  | [optional] 
**InstallYbc** | Pointer to **bool** |  | [optional] 
**ItestS3PackagePath** | Pointer to **string** |  | [optional] 
**KubernetesUpgradeSupported** | **bool** |  | 
**MastersInDefaultRegion** | Pointer to **bool** |  | [optional] 
**NextClusterIndex** | Pointer to **int32** |  | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**NodePrefix** | Pointer to **string** |  | [optional] 
**NodesResizeAvailable** | Pointer to **bool** |  | [optional] 
**PlatformUrl** | **string** |  | 
**PlatformVersion** | **string** |  | 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID of a retry | [optional] 
**RemotePackagePath** | Pointer to **string** |  | [optional] 
**ResetAZConfig** | Pointer to **bool** |  | [optional] 
**RootAndClientRootCASame** | Pointer to **bool** |  | [optional] 
**RootCA** | Pointer to **string** |  | [optional] 
**SetTxnTableWaitCountFlag** | Pointer to **bool** |  | [optional] 
**SleepAfterMasterRestartMillis** | **int32** |  | 
**SleepAfterTServerRestartMillis** | **int32** |  | 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] [readonly] 
**SshUserOverride** | Pointer to **string** |  | [optional] 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] [readonly] 
**UniversePaused** | Pointer to **bool** |  | [optional] 
**UniverseUUID** | Pointer to **string** | Associated universe UUID | [optional] 
**UpdateInProgress** | Pointer to **bool** |  | [optional] 
**UpdateOptions** | Pointer to **[]string** |  | [optional] 
**UpdateSucceeded** | Pointer to **bool** |  | [optional] 
**UpdatingTask** | Pointer to **string** |  | [optional] 
**UpdatingTaskUUID** | Pointer to **string** |  | [optional] 
**UpgradeOption** | **string** |  | 
**UseNewHelmNamingStyle** | Pointer to **bool** |  | [optional] 
**UserAZSelected** | Pointer to **bool** |  | [optional] 
**XclusterInfo** | Pointer to [**XClusterInfo**](XClusterInfo.md) |  | [optional] 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 
**YbcInstalled** | Pointer to **bool** |  | [optional] 
**YbcSoftwareVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemdUpgradeParams

`func NewSystemdUpgradeParams(clusters []Cluster, creatingUser Users, kubernetesUpgradeSupported bool, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, upgradeOption string, ) *SystemdUpgradeParams`

NewSystemdUpgradeParams instantiates a new SystemdUpgradeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemdUpgradeParamsWithDefaults

`func NewSystemdUpgradeParamsWithDefaults() *SystemdUpgradeParams`

NewSystemdUpgradeParamsWithDefaults instantiates a new SystemdUpgradeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *SystemdUpgradeParams) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *SystemdUpgradeParams) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *SystemdUpgradeParams) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *SystemdUpgradeParams) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetCapability

`func (o *SystemdUpgradeParams) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *SystemdUpgradeParams) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *SystemdUpgradeParams) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *SystemdUpgradeParams) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetClientRootCA

`func (o *SystemdUpgradeParams) GetClientRootCA() string`

GetClientRootCA returns the ClientRootCA field if non-nil, zero value otherwise.

### GetClientRootCAOk

`func (o *SystemdUpgradeParams) GetClientRootCAOk() (*string, bool)`

GetClientRootCAOk returns a tuple with the ClientRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCA

`func (o *SystemdUpgradeParams) SetClientRootCA(v string)`

SetClientRootCA sets ClientRootCA field to given value.

### HasClientRootCA

`func (o *SystemdUpgradeParams) HasClientRootCA() bool`

HasClientRootCA returns a boolean if a field has been set.

### GetClusters

`func (o *SystemdUpgradeParams) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SystemdUpgradeParams) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SystemdUpgradeParams) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.


### GetCmkArn

`func (o *SystemdUpgradeParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *SystemdUpgradeParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *SystemdUpgradeParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *SystemdUpgradeParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *SystemdUpgradeParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *SystemdUpgradeParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *SystemdUpgradeParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *SystemdUpgradeParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *SystemdUpgradeParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *SystemdUpgradeParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *SystemdUpgradeParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCurrentClusterType

`func (o *SystemdUpgradeParams) GetCurrentClusterType() string`

GetCurrentClusterType returns the CurrentClusterType field if non-nil, zero value otherwise.

### GetCurrentClusterTypeOk

`func (o *SystemdUpgradeParams) GetCurrentClusterTypeOk() (*string, bool)`

GetCurrentClusterTypeOk returns a tuple with the CurrentClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClusterType

`func (o *SystemdUpgradeParams) SetCurrentClusterType(v string)`

SetCurrentClusterType sets CurrentClusterType field to given value.

### HasCurrentClusterType

`func (o *SystemdUpgradeParams) HasCurrentClusterType() bool`

HasCurrentClusterType returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *SystemdUpgradeParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *SystemdUpgradeParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *SystemdUpgradeParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *SystemdUpgradeParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableYbc

`func (o *SystemdUpgradeParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *SystemdUpgradeParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *SystemdUpgradeParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *SystemdUpgradeParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *SystemdUpgradeParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *SystemdUpgradeParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *SystemdUpgradeParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *SystemdUpgradeParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *SystemdUpgradeParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *SystemdUpgradeParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *SystemdUpgradeParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *SystemdUpgradeParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *SystemdUpgradeParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *SystemdUpgradeParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *SystemdUpgradeParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *SystemdUpgradeParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *SystemdUpgradeParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *SystemdUpgradeParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *SystemdUpgradeParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *SystemdUpgradeParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetImportedState

`func (o *SystemdUpgradeParams) GetImportedState() string`

GetImportedState returns the ImportedState field if non-nil, zero value otherwise.

### GetImportedStateOk

`func (o *SystemdUpgradeParams) GetImportedStateOk() (*string, bool)`

GetImportedStateOk returns a tuple with the ImportedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedState

`func (o *SystemdUpgradeParams) SetImportedState(v string)`

SetImportedState sets ImportedState field to given value.

### HasImportedState

`func (o *SystemdUpgradeParams) HasImportedState() bool`

HasImportedState returns a boolean if a field has been set.

### GetInstallYbc

`func (o *SystemdUpgradeParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *SystemdUpgradeParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *SystemdUpgradeParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *SystemdUpgradeParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetItestS3PackagePath

`func (o *SystemdUpgradeParams) GetItestS3PackagePath() string`

GetItestS3PackagePath returns the ItestS3PackagePath field if non-nil, zero value otherwise.

### GetItestS3PackagePathOk

`func (o *SystemdUpgradeParams) GetItestS3PackagePathOk() (*string, bool)`

GetItestS3PackagePathOk returns a tuple with the ItestS3PackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItestS3PackagePath

`func (o *SystemdUpgradeParams) SetItestS3PackagePath(v string)`

SetItestS3PackagePath sets ItestS3PackagePath field to given value.

### HasItestS3PackagePath

`func (o *SystemdUpgradeParams) HasItestS3PackagePath() bool`

HasItestS3PackagePath returns a boolean if a field has been set.

### GetKubernetesUpgradeSupported

`func (o *SystemdUpgradeParams) GetKubernetesUpgradeSupported() bool`

GetKubernetesUpgradeSupported returns the KubernetesUpgradeSupported field if non-nil, zero value otherwise.

### GetKubernetesUpgradeSupportedOk

`func (o *SystemdUpgradeParams) GetKubernetesUpgradeSupportedOk() (*bool, bool)`

GetKubernetesUpgradeSupportedOk returns a tuple with the KubernetesUpgradeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesUpgradeSupported

`func (o *SystemdUpgradeParams) SetKubernetesUpgradeSupported(v bool)`

SetKubernetesUpgradeSupported sets KubernetesUpgradeSupported field to given value.


### GetMastersInDefaultRegion

`func (o *SystemdUpgradeParams) GetMastersInDefaultRegion() bool`

GetMastersInDefaultRegion returns the MastersInDefaultRegion field if non-nil, zero value otherwise.

### GetMastersInDefaultRegionOk

`func (o *SystemdUpgradeParams) GetMastersInDefaultRegionOk() (*bool, bool)`

GetMastersInDefaultRegionOk returns a tuple with the MastersInDefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastersInDefaultRegion

`func (o *SystemdUpgradeParams) SetMastersInDefaultRegion(v bool)`

SetMastersInDefaultRegion sets MastersInDefaultRegion field to given value.

### HasMastersInDefaultRegion

`func (o *SystemdUpgradeParams) HasMastersInDefaultRegion() bool`

HasMastersInDefaultRegion returns a boolean if a field has been set.

### GetNextClusterIndex

`func (o *SystemdUpgradeParams) GetNextClusterIndex() int32`

GetNextClusterIndex returns the NextClusterIndex field if non-nil, zero value otherwise.

### GetNextClusterIndexOk

`func (o *SystemdUpgradeParams) GetNextClusterIndexOk() (*int32, bool)`

GetNextClusterIndexOk returns a tuple with the NextClusterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextClusterIndex

`func (o *SystemdUpgradeParams) SetNextClusterIndex(v int32)`

SetNextClusterIndex sets NextClusterIndex field to given value.

### HasNextClusterIndex

`func (o *SystemdUpgradeParams) HasNextClusterIndex() bool`

HasNextClusterIndex returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *SystemdUpgradeParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *SystemdUpgradeParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *SystemdUpgradeParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *SystemdUpgradeParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *SystemdUpgradeParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *SystemdUpgradeParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *SystemdUpgradeParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *SystemdUpgradeParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNodePrefix

`func (o *SystemdUpgradeParams) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *SystemdUpgradeParams) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *SystemdUpgradeParams) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *SystemdUpgradeParams) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetNodesResizeAvailable

`func (o *SystemdUpgradeParams) GetNodesResizeAvailable() bool`

GetNodesResizeAvailable returns the NodesResizeAvailable field if non-nil, zero value otherwise.

### GetNodesResizeAvailableOk

`func (o *SystemdUpgradeParams) GetNodesResizeAvailableOk() (*bool, bool)`

GetNodesResizeAvailableOk returns a tuple with the NodesResizeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesResizeAvailable

`func (o *SystemdUpgradeParams) SetNodesResizeAvailable(v bool)`

SetNodesResizeAvailable sets NodesResizeAvailable field to given value.

### HasNodesResizeAvailable

`func (o *SystemdUpgradeParams) HasNodesResizeAvailable() bool`

HasNodesResizeAvailable returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *SystemdUpgradeParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *SystemdUpgradeParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *SystemdUpgradeParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPlatformVersion

`func (o *SystemdUpgradeParams) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *SystemdUpgradeParams) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *SystemdUpgradeParams) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.


### GetPreviousTaskUUID

`func (o *SystemdUpgradeParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *SystemdUpgradeParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *SystemdUpgradeParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *SystemdUpgradeParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRemotePackagePath

`func (o *SystemdUpgradeParams) GetRemotePackagePath() string`

GetRemotePackagePath returns the RemotePackagePath field if non-nil, zero value otherwise.

### GetRemotePackagePathOk

`func (o *SystemdUpgradeParams) GetRemotePackagePathOk() (*string, bool)`

GetRemotePackagePathOk returns a tuple with the RemotePackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePackagePath

`func (o *SystemdUpgradeParams) SetRemotePackagePath(v string)`

SetRemotePackagePath sets RemotePackagePath field to given value.

### HasRemotePackagePath

`func (o *SystemdUpgradeParams) HasRemotePackagePath() bool`

HasRemotePackagePath returns a boolean if a field has been set.

### GetResetAZConfig

`func (o *SystemdUpgradeParams) GetResetAZConfig() bool`

GetResetAZConfig returns the ResetAZConfig field if non-nil, zero value otherwise.

### GetResetAZConfigOk

`func (o *SystemdUpgradeParams) GetResetAZConfigOk() (*bool, bool)`

GetResetAZConfigOk returns a tuple with the ResetAZConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAZConfig

`func (o *SystemdUpgradeParams) SetResetAZConfig(v bool)`

SetResetAZConfig sets ResetAZConfig field to given value.

### HasResetAZConfig

`func (o *SystemdUpgradeParams) HasResetAZConfig() bool`

HasResetAZConfig returns a boolean if a field has been set.

### GetRootAndClientRootCASame

`func (o *SystemdUpgradeParams) GetRootAndClientRootCASame() bool`

GetRootAndClientRootCASame returns the RootAndClientRootCASame field if non-nil, zero value otherwise.

### GetRootAndClientRootCASameOk

`func (o *SystemdUpgradeParams) GetRootAndClientRootCASameOk() (*bool, bool)`

GetRootAndClientRootCASameOk returns a tuple with the RootAndClientRootCASame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootAndClientRootCASame

`func (o *SystemdUpgradeParams) SetRootAndClientRootCASame(v bool)`

SetRootAndClientRootCASame sets RootAndClientRootCASame field to given value.

### HasRootAndClientRootCASame

`func (o *SystemdUpgradeParams) HasRootAndClientRootCASame() bool`

HasRootAndClientRootCASame returns a boolean if a field has been set.

### GetRootCA

`func (o *SystemdUpgradeParams) GetRootCA() string`

GetRootCA returns the RootCA field if non-nil, zero value otherwise.

### GetRootCAOk

`func (o *SystemdUpgradeParams) GetRootCAOk() (*string, bool)`

GetRootCAOk returns a tuple with the RootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCA

`func (o *SystemdUpgradeParams) SetRootCA(v string)`

SetRootCA sets RootCA field to given value.

### HasRootCA

`func (o *SystemdUpgradeParams) HasRootCA() bool`

HasRootCA returns a boolean if a field has been set.

### GetSetTxnTableWaitCountFlag

`func (o *SystemdUpgradeParams) GetSetTxnTableWaitCountFlag() bool`

GetSetTxnTableWaitCountFlag returns the SetTxnTableWaitCountFlag field if non-nil, zero value otherwise.

### GetSetTxnTableWaitCountFlagOk

`func (o *SystemdUpgradeParams) GetSetTxnTableWaitCountFlagOk() (*bool, bool)`

GetSetTxnTableWaitCountFlagOk returns a tuple with the SetTxnTableWaitCountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTxnTableWaitCountFlag

`func (o *SystemdUpgradeParams) SetSetTxnTableWaitCountFlag(v bool)`

SetSetTxnTableWaitCountFlag sets SetTxnTableWaitCountFlag field to given value.

### HasSetTxnTableWaitCountFlag

`func (o *SystemdUpgradeParams) HasSetTxnTableWaitCountFlag() bool`

HasSetTxnTableWaitCountFlag returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *SystemdUpgradeParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *SystemdUpgradeParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *SystemdUpgradeParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *SystemdUpgradeParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *SystemdUpgradeParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *SystemdUpgradeParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSourceXClusterConfigs

`func (o *SystemdUpgradeParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *SystemdUpgradeParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *SystemdUpgradeParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *SystemdUpgradeParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSshUserOverride

`func (o *SystemdUpgradeParams) GetSshUserOverride() string`

GetSshUserOverride returns the SshUserOverride field if non-nil, zero value otherwise.

### GetSshUserOverrideOk

`func (o *SystemdUpgradeParams) GetSshUserOverrideOk() (*string, bool)`

GetSshUserOverrideOk returns a tuple with the SshUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverride

`func (o *SystemdUpgradeParams) SetSshUserOverride(v string)`

SetSshUserOverride sets SshUserOverride field to given value.

### HasSshUserOverride

`func (o *SystemdUpgradeParams) HasSshUserOverride() bool`

HasSshUserOverride returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *SystemdUpgradeParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *SystemdUpgradeParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *SystemdUpgradeParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *SystemdUpgradeParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetUniversePaused

`func (o *SystemdUpgradeParams) GetUniversePaused() bool`

GetUniversePaused returns the UniversePaused field if non-nil, zero value otherwise.

### GetUniversePausedOk

`func (o *SystemdUpgradeParams) GetUniversePausedOk() (*bool, bool)`

GetUniversePausedOk returns a tuple with the UniversePaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversePaused

`func (o *SystemdUpgradeParams) SetUniversePaused(v bool)`

SetUniversePaused sets UniversePaused field to given value.

### HasUniversePaused

`func (o *SystemdUpgradeParams) HasUniversePaused() bool`

HasUniversePaused returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *SystemdUpgradeParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *SystemdUpgradeParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *SystemdUpgradeParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *SystemdUpgradeParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUpdateInProgress

`func (o *SystemdUpgradeParams) GetUpdateInProgress() bool`

GetUpdateInProgress returns the UpdateInProgress field if non-nil, zero value otherwise.

### GetUpdateInProgressOk

`func (o *SystemdUpgradeParams) GetUpdateInProgressOk() (*bool, bool)`

GetUpdateInProgressOk returns a tuple with the UpdateInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInProgress

`func (o *SystemdUpgradeParams) SetUpdateInProgress(v bool)`

SetUpdateInProgress sets UpdateInProgress field to given value.

### HasUpdateInProgress

`func (o *SystemdUpgradeParams) HasUpdateInProgress() bool`

HasUpdateInProgress returns a boolean if a field has been set.

### GetUpdateOptions

`func (o *SystemdUpgradeParams) GetUpdateOptions() []string`

GetUpdateOptions returns the UpdateOptions field if non-nil, zero value otherwise.

### GetUpdateOptionsOk

`func (o *SystemdUpgradeParams) GetUpdateOptionsOk() (*[]string, bool)`

GetUpdateOptionsOk returns a tuple with the UpdateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOptions

`func (o *SystemdUpgradeParams) SetUpdateOptions(v []string)`

SetUpdateOptions sets UpdateOptions field to given value.

### HasUpdateOptions

`func (o *SystemdUpgradeParams) HasUpdateOptions() bool`

HasUpdateOptions returns a boolean if a field has been set.

### GetUpdateSucceeded

`func (o *SystemdUpgradeParams) GetUpdateSucceeded() bool`

GetUpdateSucceeded returns the UpdateSucceeded field if non-nil, zero value otherwise.

### GetUpdateSucceededOk

`func (o *SystemdUpgradeParams) GetUpdateSucceededOk() (*bool, bool)`

GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceeded

`func (o *SystemdUpgradeParams) SetUpdateSucceeded(v bool)`

SetUpdateSucceeded sets UpdateSucceeded field to given value.

### HasUpdateSucceeded

`func (o *SystemdUpgradeParams) HasUpdateSucceeded() bool`

HasUpdateSucceeded returns a boolean if a field has been set.

### GetUpdatingTask

`func (o *SystemdUpgradeParams) GetUpdatingTask() string`

GetUpdatingTask returns the UpdatingTask field if non-nil, zero value otherwise.

### GetUpdatingTaskOk

`func (o *SystemdUpgradeParams) GetUpdatingTaskOk() (*string, bool)`

GetUpdatingTaskOk returns a tuple with the UpdatingTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTask

`func (o *SystemdUpgradeParams) SetUpdatingTask(v string)`

SetUpdatingTask sets UpdatingTask field to given value.

### HasUpdatingTask

`func (o *SystemdUpgradeParams) HasUpdatingTask() bool`

HasUpdatingTask returns a boolean if a field has been set.

### GetUpdatingTaskUUID

`func (o *SystemdUpgradeParams) GetUpdatingTaskUUID() string`

GetUpdatingTaskUUID returns the UpdatingTaskUUID field if non-nil, zero value otherwise.

### GetUpdatingTaskUUIDOk

`func (o *SystemdUpgradeParams) GetUpdatingTaskUUIDOk() (*string, bool)`

GetUpdatingTaskUUIDOk returns a tuple with the UpdatingTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTaskUUID

`func (o *SystemdUpgradeParams) SetUpdatingTaskUUID(v string)`

SetUpdatingTaskUUID sets UpdatingTaskUUID field to given value.

### HasUpdatingTaskUUID

`func (o *SystemdUpgradeParams) HasUpdatingTaskUUID() bool`

HasUpdatingTaskUUID returns a boolean if a field has been set.

### GetUpgradeOption

`func (o *SystemdUpgradeParams) GetUpgradeOption() string`

GetUpgradeOption returns the UpgradeOption field if non-nil, zero value otherwise.

### GetUpgradeOptionOk

`func (o *SystemdUpgradeParams) GetUpgradeOptionOk() (*string, bool)`

GetUpgradeOptionOk returns a tuple with the UpgradeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOption

`func (o *SystemdUpgradeParams) SetUpgradeOption(v string)`

SetUpgradeOption sets UpgradeOption field to given value.


### GetUseNewHelmNamingStyle

`func (o *SystemdUpgradeParams) GetUseNewHelmNamingStyle() bool`

GetUseNewHelmNamingStyle returns the UseNewHelmNamingStyle field if non-nil, zero value otherwise.

### GetUseNewHelmNamingStyleOk

`func (o *SystemdUpgradeParams) GetUseNewHelmNamingStyleOk() (*bool, bool)`

GetUseNewHelmNamingStyleOk returns a tuple with the UseNewHelmNamingStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNewHelmNamingStyle

`func (o *SystemdUpgradeParams) SetUseNewHelmNamingStyle(v bool)`

SetUseNewHelmNamingStyle sets UseNewHelmNamingStyle field to given value.

### HasUseNewHelmNamingStyle

`func (o *SystemdUpgradeParams) HasUseNewHelmNamingStyle() bool`

HasUseNewHelmNamingStyle returns a boolean if a field has been set.

### GetUserAZSelected

`func (o *SystemdUpgradeParams) GetUserAZSelected() bool`

GetUserAZSelected returns the UserAZSelected field if non-nil, zero value otherwise.

### GetUserAZSelectedOk

`func (o *SystemdUpgradeParams) GetUserAZSelectedOk() (*bool, bool)`

GetUserAZSelectedOk returns a tuple with the UserAZSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAZSelected

`func (o *SystemdUpgradeParams) SetUserAZSelected(v bool)`

SetUserAZSelected sets UserAZSelected field to given value.

### HasUserAZSelected

`func (o *SystemdUpgradeParams) HasUserAZSelected() bool`

HasUserAZSelected returns a boolean if a field has been set.

### GetXclusterInfo

`func (o *SystemdUpgradeParams) GetXclusterInfo() XClusterInfo`

GetXclusterInfo returns the XclusterInfo field if non-nil, zero value otherwise.

### GetXclusterInfoOk

`func (o *SystemdUpgradeParams) GetXclusterInfoOk() (*XClusterInfo, bool)`

GetXclusterInfoOk returns a tuple with the XclusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterInfo

`func (o *SystemdUpgradeParams) SetXclusterInfo(v XClusterInfo)`

SetXclusterInfo sets XclusterInfo field to given value.

### HasXclusterInfo

`func (o *SystemdUpgradeParams) HasXclusterInfo() bool`

HasXclusterInfo returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *SystemdUpgradeParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *SystemdUpgradeParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *SystemdUpgradeParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *SystemdUpgradeParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *SystemdUpgradeParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *SystemdUpgradeParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *SystemdUpgradeParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *SystemdUpgradeParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *SystemdUpgradeParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *SystemdUpgradeParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *SystemdUpgradeParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *SystemdUpgradeParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


