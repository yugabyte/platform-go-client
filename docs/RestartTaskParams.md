# RestartTaskParams

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

### NewRestartTaskParams

`func NewRestartTaskParams(clusters []Cluster, creatingUser Users, kubernetesUpgradeSupported bool, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, upgradeOption string, ) *RestartTaskParams`

NewRestartTaskParams instantiates a new RestartTaskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartTaskParamsWithDefaults

`func NewRestartTaskParamsWithDefaults() *RestartTaskParams`

NewRestartTaskParamsWithDefaults instantiates a new RestartTaskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *RestartTaskParams) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *RestartTaskParams) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *RestartTaskParams) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *RestartTaskParams) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetCapability

`func (o *RestartTaskParams) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *RestartTaskParams) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *RestartTaskParams) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *RestartTaskParams) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetClientRootCA

`func (o *RestartTaskParams) GetClientRootCA() string`

GetClientRootCA returns the ClientRootCA field if non-nil, zero value otherwise.

### GetClientRootCAOk

`func (o *RestartTaskParams) GetClientRootCAOk() (*string, bool)`

GetClientRootCAOk returns a tuple with the ClientRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCA

`func (o *RestartTaskParams) SetClientRootCA(v string)`

SetClientRootCA sets ClientRootCA field to given value.

### HasClientRootCA

`func (o *RestartTaskParams) HasClientRootCA() bool`

HasClientRootCA returns a boolean if a field has been set.

### GetClusters

`func (o *RestartTaskParams) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *RestartTaskParams) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *RestartTaskParams) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.


### GetCmkArn

`func (o *RestartTaskParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *RestartTaskParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *RestartTaskParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *RestartTaskParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *RestartTaskParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *RestartTaskParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *RestartTaskParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *RestartTaskParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *RestartTaskParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *RestartTaskParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *RestartTaskParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCurrentClusterType

`func (o *RestartTaskParams) GetCurrentClusterType() string`

GetCurrentClusterType returns the CurrentClusterType field if non-nil, zero value otherwise.

### GetCurrentClusterTypeOk

`func (o *RestartTaskParams) GetCurrentClusterTypeOk() (*string, bool)`

GetCurrentClusterTypeOk returns a tuple with the CurrentClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClusterType

`func (o *RestartTaskParams) SetCurrentClusterType(v string)`

SetCurrentClusterType sets CurrentClusterType field to given value.

### HasCurrentClusterType

`func (o *RestartTaskParams) HasCurrentClusterType() bool`

HasCurrentClusterType returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *RestartTaskParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *RestartTaskParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *RestartTaskParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *RestartTaskParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableYbc

`func (o *RestartTaskParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *RestartTaskParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *RestartTaskParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *RestartTaskParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *RestartTaskParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *RestartTaskParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *RestartTaskParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *RestartTaskParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *RestartTaskParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *RestartTaskParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *RestartTaskParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *RestartTaskParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *RestartTaskParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *RestartTaskParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *RestartTaskParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *RestartTaskParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *RestartTaskParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *RestartTaskParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *RestartTaskParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *RestartTaskParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetImportedState

`func (o *RestartTaskParams) GetImportedState() string`

GetImportedState returns the ImportedState field if non-nil, zero value otherwise.

### GetImportedStateOk

`func (o *RestartTaskParams) GetImportedStateOk() (*string, bool)`

GetImportedStateOk returns a tuple with the ImportedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedState

`func (o *RestartTaskParams) SetImportedState(v string)`

SetImportedState sets ImportedState field to given value.

### HasImportedState

`func (o *RestartTaskParams) HasImportedState() bool`

HasImportedState returns a boolean if a field has been set.

### GetInstallYbc

`func (o *RestartTaskParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *RestartTaskParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *RestartTaskParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *RestartTaskParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetItestS3PackagePath

`func (o *RestartTaskParams) GetItestS3PackagePath() string`

GetItestS3PackagePath returns the ItestS3PackagePath field if non-nil, zero value otherwise.

### GetItestS3PackagePathOk

`func (o *RestartTaskParams) GetItestS3PackagePathOk() (*string, bool)`

GetItestS3PackagePathOk returns a tuple with the ItestS3PackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItestS3PackagePath

`func (o *RestartTaskParams) SetItestS3PackagePath(v string)`

SetItestS3PackagePath sets ItestS3PackagePath field to given value.

### HasItestS3PackagePath

`func (o *RestartTaskParams) HasItestS3PackagePath() bool`

HasItestS3PackagePath returns a boolean if a field has been set.

### GetKubernetesUpgradeSupported

`func (o *RestartTaskParams) GetKubernetesUpgradeSupported() bool`

GetKubernetesUpgradeSupported returns the KubernetesUpgradeSupported field if non-nil, zero value otherwise.

### GetKubernetesUpgradeSupportedOk

`func (o *RestartTaskParams) GetKubernetesUpgradeSupportedOk() (*bool, bool)`

GetKubernetesUpgradeSupportedOk returns a tuple with the KubernetesUpgradeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesUpgradeSupported

`func (o *RestartTaskParams) SetKubernetesUpgradeSupported(v bool)`

SetKubernetesUpgradeSupported sets KubernetesUpgradeSupported field to given value.


### GetMastersInDefaultRegion

`func (o *RestartTaskParams) GetMastersInDefaultRegion() bool`

GetMastersInDefaultRegion returns the MastersInDefaultRegion field if non-nil, zero value otherwise.

### GetMastersInDefaultRegionOk

`func (o *RestartTaskParams) GetMastersInDefaultRegionOk() (*bool, bool)`

GetMastersInDefaultRegionOk returns a tuple with the MastersInDefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastersInDefaultRegion

`func (o *RestartTaskParams) SetMastersInDefaultRegion(v bool)`

SetMastersInDefaultRegion sets MastersInDefaultRegion field to given value.

### HasMastersInDefaultRegion

`func (o *RestartTaskParams) HasMastersInDefaultRegion() bool`

HasMastersInDefaultRegion returns a boolean if a field has been set.

### GetNextClusterIndex

`func (o *RestartTaskParams) GetNextClusterIndex() int32`

GetNextClusterIndex returns the NextClusterIndex field if non-nil, zero value otherwise.

### GetNextClusterIndexOk

`func (o *RestartTaskParams) GetNextClusterIndexOk() (*int32, bool)`

GetNextClusterIndexOk returns a tuple with the NextClusterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextClusterIndex

`func (o *RestartTaskParams) SetNextClusterIndex(v int32)`

SetNextClusterIndex sets NextClusterIndex field to given value.

### HasNextClusterIndex

`func (o *RestartTaskParams) HasNextClusterIndex() bool`

HasNextClusterIndex returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *RestartTaskParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *RestartTaskParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *RestartTaskParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *RestartTaskParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *RestartTaskParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *RestartTaskParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *RestartTaskParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *RestartTaskParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNodePrefix

`func (o *RestartTaskParams) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *RestartTaskParams) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *RestartTaskParams) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *RestartTaskParams) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetNodesResizeAvailable

`func (o *RestartTaskParams) GetNodesResizeAvailable() bool`

GetNodesResizeAvailable returns the NodesResizeAvailable field if non-nil, zero value otherwise.

### GetNodesResizeAvailableOk

`func (o *RestartTaskParams) GetNodesResizeAvailableOk() (*bool, bool)`

GetNodesResizeAvailableOk returns a tuple with the NodesResizeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesResizeAvailable

`func (o *RestartTaskParams) SetNodesResizeAvailable(v bool)`

SetNodesResizeAvailable sets NodesResizeAvailable field to given value.

### HasNodesResizeAvailable

`func (o *RestartTaskParams) HasNodesResizeAvailable() bool`

HasNodesResizeAvailable returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *RestartTaskParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *RestartTaskParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *RestartTaskParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPlatformVersion

`func (o *RestartTaskParams) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *RestartTaskParams) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *RestartTaskParams) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.


### GetPreviousTaskUUID

`func (o *RestartTaskParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *RestartTaskParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *RestartTaskParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *RestartTaskParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRemotePackagePath

`func (o *RestartTaskParams) GetRemotePackagePath() string`

GetRemotePackagePath returns the RemotePackagePath field if non-nil, zero value otherwise.

### GetRemotePackagePathOk

`func (o *RestartTaskParams) GetRemotePackagePathOk() (*string, bool)`

GetRemotePackagePathOk returns a tuple with the RemotePackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePackagePath

`func (o *RestartTaskParams) SetRemotePackagePath(v string)`

SetRemotePackagePath sets RemotePackagePath field to given value.

### HasRemotePackagePath

`func (o *RestartTaskParams) HasRemotePackagePath() bool`

HasRemotePackagePath returns a boolean if a field has been set.

### GetResetAZConfig

`func (o *RestartTaskParams) GetResetAZConfig() bool`

GetResetAZConfig returns the ResetAZConfig field if non-nil, zero value otherwise.

### GetResetAZConfigOk

`func (o *RestartTaskParams) GetResetAZConfigOk() (*bool, bool)`

GetResetAZConfigOk returns a tuple with the ResetAZConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAZConfig

`func (o *RestartTaskParams) SetResetAZConfig(v bool)`

SetResetAZConfig sets ResetAZConfig field to given value.

### HasResetAZConfig

`func (o *RestartTaskParams) HasResetAZConfig() bool`

HasResetAZConfig returns a boolean if a field has been set.

### GetRootAndClientRootCASame

`func (o *RestartTaskParams) GetRootAndClientRootCASame() bool`

GetRootAndClientRootCASame returns the RootAndClientRootCASame field if non-nil, zero value otherwise.

### GetRootAndClientRootCASameOk

`func (o *RestartTaskParams) GetRootAndClientRootCASameOk() (*bool, bool)`

GetRootAndClientRootCASameOk returns a tuple with the RootAndClientRootCASame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootAndClientRootCASame

`func (o *RestartTaskParams) SetRootAndClientRootCASame(v bool)`

SetRootAndClientRootCASame sets RootAndClientRootCASame field to given value.

### HasRootAndClientRootCASame

`func (o *RestartTaskParams) HasRootAndClientRootCASame() bool`

HasRootAndClientRootCASame returns a boolean if a field has been set.

### GetRootCA

`func (o *RestartTaskParams) GetRootCA() string`

GetRootCA returns the RootCA field if non-nil, zero value otherwise.

### GetRootCAOk

`func (o *RestartTaskParams) GetRootCAOk() (*string, bool)`

GetRootCAOk returns a tuple with the RootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCA

`func (o *RestartTaskParams) SetRootCA(v string)`

SetRootCA sets RootCA field to given value.

### HasRootCA

`func (o *RestartTaskParams) HasRootCA() bool`

HasRootCA returns a boolean if a field has been set.

### GetSetTxnTableWaitCountFlag

`func (o *RestartTaskParams) GetSetTxnTableWaitCountFlag() bool`

GetSetTxnTableWaitCountFlag returns the SetTxnTableWaitCountFlag field if non-nil, zero value otherwise.

### GetSetTxnTableWaitCountFlagOk

`func (o *RestartTaskParams) GetSetTxnTableWaitCountFlagOk() (*bool, bool)`

GetSetTxnTableWaitCountFlagOk returns a tuple with the SetTxnTableWaitCountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTxnTableWaitCountFlag

`func (o *RestartTaskParams) SetSetTxnTableWaitCountFlag(v bool)`

SetSetTxnTableWaitCountFlag sets SetTxnTableWaitCountFlag field to given value.

### HasSetTxnTableWaitCountFlag

`func (o *RestartTaskParams) HasSetTxnTableWaitCountFlag() bool`

HasSetTxnTableWaitCountFlag returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *RestartTaskParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *RestartTaskParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *RestartTaskParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *RestartTaskParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *RestartTaskParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *RestartTaskParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSourceXClusterConfigs

`func (o *RestartTaskParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *RestartTaskParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *RestartTaskParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *RestartTaskParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSshUserOverride

`func (o *RestartTaskParams) GetSshUserOverride() string`

GetSshUserOverride returns the SshUserOverride field if non-nil, zero value otherwise.

### GetSshUserOverrideOk

`func (o *RestartTaskParams) GetSshUserOverrideOk() (*string, bool)`

GetSshUserOverrideOk returns a tuple with the SshUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverride

`func (o *RestartTaskParams) SetSshUserOverride(v string)`

SetSshUserOverride sets SshUserOverride field to given value.

### HasSshUserOverride

`func (o *RestartTaskParams) HasSshUserOverride() bool`

HasSshUserOverride returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *RestartTaskParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *RestartTaskParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *RestartTaskParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *RestartTaskParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetUniversePaused

`func (o *RestartTaskParams) GetUniversePaused() bool`

GetUniversePaused returns the UniversePaused field if non-nil, zero value otherwise.

### GetUniversePausedOk

`func (o *RestartTaskParams) GetUniversePausedOk() (*bool, bool)`

GetUniversePausedOk returns a tuple with the UniversePaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversePaused

`func (o *RestartTaskParams) SetUniversePaused(v bool)`

SetUniversePaused sets UniversePaused field to given value.

### HasUniversePaused

`func (o *RestartTaskParams) HasUniversePaused() bool`

HasUniversePaused returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *RestartTaskParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *RestartTaskParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *RestartTaskParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *RestartTaskParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUpdateInProgress

`func (o *RestartTaskParams) GetUpdateInProgress() bool`

GetUpdateInProgress returns the UpdateInProgress field if non-nil, zero value otherwise.

### GetUpdateInProgressOk

`func (o *RestartTaskParams) GetUpdateInProgressOk() (*bool, bool)`

GetUpdateInProgressOk returns a tuple with the UpdateInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInProgress

`func (o *RestartTaskParams) SetUpdateInProgress(v bool)`

SetUpdateInProgress sets UpdateInProgress field to given value.

### HasUpdateInProgress

`func (o *RestartTaskParams) HasUpdateInProgress() bool`

HasUpdateInProgress returns a boolean if a field has been set.

### GetUpdateOptions

`func (o *RestartTaskParams) GetUpdateOptions() []string`

GetUpdateOptions returns the UpdateOptions field if non-nil, zero value otherwise.

### GetUpdateOptionsOk

`func (o *RestartTaskParams) GetUpdateOptionsOk() (*[]string, bool)`

GetUpdateOptionsOk returns a tuple with the UpdateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOptions

`func (o *RestartTaskParams) SetUpdateOptions(v []string)`

SetUpdateOptions sets UpdateOptions field to given value.

### HasUpdateOptions

`func (o *RestartTaskParams) HasUpdateOptions() bool`

HasUpdateOptions returns a boolean if a field has been set.

### GetUpdateSucceeded

`func (o *RestartTaskParams) GetUpdateSucceeded() bool`

GetUpdateSucceeded returns the UpdateSucceeded field if non-nil, zero value otherwise.

### GetUpdateSucceededOk

`func (o *RestartTaskParams) GetUpdateSucceededOk() (*bool, bool)`

GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceeded

`func (o *RestartTaskParams) SetUpdateSucceeded(v bool)`

SetUpdateSucceeded sets UpdateSucceeded field to given value.

### HasUpdateSucceeded

`func (o *RestartTaskParams) HasUpdateSucceeded() bool`

HasUpdateSucceeded returns a boolean if a field has been set.

### GetUpdatingTask

`func (o *RestartTaskParams) GetUpdatingTask() string`

GetUpdatingTask returns the UpdatingTask field if non-nil, zero value otherwise.

### GetUpdatingTaskOk

`func (o *RestartTaskParams) GetUpdatingTaskOk() (*string, bool)`

GetUpdatingTaskOk returns a tuple with the UpdatingTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTask

`func (o *RestartTaskParams) SetUpdatingTask(v string)`

SetUpdatingTask sets UpdatingTask field to given value.

### HasUpdatingTask

`func (o *RestartTaskParams) HasUpdatingTask() bool`

HasUpdatingTask returns a boolean if a field has been set.

### GetUpdatingTaskUUID

`func (o *RestartTaskParams) GetUpdatingTaskUUID() string`

GetUpdatingTaskUUID returns the UpdatingTaskUUID field if non-nil, zero value otherwise.

### GetUpdatingTaskUUIDOk

`func (o *RestartTaskParams) GetUpdatingTaskUUIDOk() (*string, bool)`

GetUpdatingTaskUUIDOk returns a tuple with the UpdatingTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTaskUUID

`func (o *RestartTaskParams) SetUpdatingTaskUUID(v string)`

SetUpdatingTaskUUID sets UpdatingTaskUUID field to given value.

### HasUpdatingTaskUUID

`func (o *RestartTaskParams) HasUpdatingTaskUUID() bool`

HasUpdatingTaskUUID returns a boolean if a field has been set.

### GetUpgradeOption

`func (o *RestartTaskParams) GetUpgradeOption() string`

GetUpgradeOption returns the UpgradeOption field if non-nil, zero value otherwise.

### GetUpgradeOptionOk

`func (o *RestartTaskParams) GetUpgradeOptionOk() (*string, bool)`

GetUpgradeOptionOk returns a tuple with the UpgradeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOption

`func (o *RestartTaskParams) SetUpgradeOption(v string)`

SetUpgradeOption sets UpgradeOption field to given value.


### GetUseNewHelmNamingStyle

`func (o *RestartTaskParams) GetUseNewHelmNamingStyle() bool`

GetUseNewHelmNamingStyle returns the UseNewHelmNamingStyle field if non-nil, zero value otherwise.

### GetUseNewHelmNamingStyleOk

`func (o *RestartTaskParams) GetUseNewHelmNamingStyleOk() (*bool, bool)`

GetUseNewHelmNamingStyleOk returns a tuple with the UseNewHelmNamingStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNewHelmNamingStyle

`func (o *RestartTaskParams) SetUseNewHelmNamingStyle(v bool)`

SetUseNewHelmNamingStyle sets UseNewHelmNamingStyle field to given value.

### HasUseNewHelmNamingStyle

`func (o *RestartTaskParams) HasUseNewHelmNamingStyle() bool`

HasUseNewHelmNamingStyle returns a boolean if a field has been set.

### GetUserAZSelected

`func (o *RestartTaskParams) GetUserAZSelected() bool`

GetUserAZSelected returns the UserAZSelected field if non-nil, zero value otherwise.

### GetUserAZSelectedOk

`func (o *RestartTaskParams) GetUserAZSelectedOk() (*bool, bool)`

GetUserAZSelectedOk returns a tuple with the UserAZSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAZSelected

`func (o *RestartTaskParams) SetUserAZSelected(v bool)`

SetUserAZSelected sets UserAZSelected field to given value.

### HasUserAZSelected

`func (o *RestartTaskParams) HasUserAZSelected() bool`

HasUserAZSelected returns a boolean if a field has been set.

### GetXclusterInfo

`func (o *RestartTaskParams) GetXclusterInfo() XClusterInfo`

GetXclusterInfo returns the XclusterInfo field if non-nil, zero value otherwise.

### GetXclusterInfoOk

`func (o *RestartTaskParams) GetXclusterInfoOk() (*XClusterInfo, bool)`

GetXclusterInfoOk returns a tuple with the XclusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterInfo

`func (o *RestartTaskParams) SetXclusterInfo(v XClusterInfo)`

SetXclusterInfo sets XclusterInfo field to given value.

### HasXclusterInfo

`func (o *RestartTaskParams) HasXclusterInfo() bool`

HasXclusterInfo returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *RestartTaskParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *RestartTaskParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *RestartTaskParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *RestartTaskParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *RestartTaskParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *RestartTaskParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *RestartTaskParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *RestartTaskParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *RestartTaskParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *RestartTaskParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *RestartTaskParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *RestartTaskParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


