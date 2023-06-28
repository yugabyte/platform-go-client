# UniverseDefinitionTaskParamsResp

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
**IsKubernetesOperatorControlled** | Pointer to **bool** |  | [optional] 
**ItestS3PackagePath** | Pointer to **string** |  | [optional] 
**MastersInDefaultRegion** | Pointer to **bool** |  | [optional] 
**NextClusterIndex** | Pointer to **int32** |  | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetailsResp**](NodeDetailsResp.md) |  | [optional] [readonly] 
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
**UseNewHelmNamingStyle** | Pointer to **bool** |  | [optional] 
**UserAZSelected** | Pointer to **bool** |  | [optional] 
**XclusterInfo** | Pointer to [**XClusterInfo**](XClusterInfo.md) |  | [optional] 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 
**YbcInstalled** | Pointer to **bool** |  | [optional] 
**YbcSoftwareVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewUniverseDefinitionTaskParamsResp

`func NewUniverseDefinitionTaskParamsResp(clusters []Cluster, creatingUser Users, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, ) *UniverseDefinitionTaskParamsResp`

NewUniverseDefinitionTaskParamsResp instantiates a new UniverseDefinitionTaskParamsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseDefinitionTaskParamsRespWithDefaults

`func NewUniverseDefinitionTaskParamsRespWithDefaults() *UniverseDefinitionTaskParamsResp`

NewUniverseDefinitionTaskParamsRespWithDefaults instantiates a new UniverseDefinitionTaskParamsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *UniverseDefinitionTaskParamsResp) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *UniverseDefinitionTaskParamsResp) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *UniverseDefinitionTaskParamsResp) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *UniverseDefinitionTaskParamsResp) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetCapability

`func (o *UniverseDefinitionTaskParamsResp) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *UniverseDefinitionTaskParamsResp) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *UniverseDefinitionTaskParamsResp) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *UniverseDefinitionTaskParamsResp) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetClientRootCA

`func (o *UniverseDefinitionTaskParamsResp) GetClientRootCA() string`

GetClientRootCA returns the ClientRootCA field if non-nil, zero value otherwise.

### GetClientRootCAOk

`func (o *UniverseDefinitionTaskParamsResp) GetClientRootCAOk() (*string, bool)`

GetClientRootCAOk returns a tuple with the ClientRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCA

`func (o *UniverseDefinitionTaskParamsResp) SetClientRootCA(v string)`

SetClientRootCA sets ClientRootCA field to given value.

### HasClientRootCA

`func (o *UniverseDefinitionTaskParamsResp) HasClientRootCA() bool`

HasClientRootCA returns a boolean if a field has been set.

### GetClusters

`func (o *UniverseDefinitionTaskParamsResp) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *UniverseDefinitionTaskParamsResp) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *UniverseDefinitionTaskParamsResp) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.


### GetCmkArn

`func (o *UniverseDefinitionTaskParamsResp) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *UniverseDefinitionTaskParamsResp) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *UniverseDefinitionTaskParamsResp) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *UniverseDefinitionTaskParamsResp) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *UniverseDefinitionTaskParamsResp) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *UniverseDefinitionTaskParamsResp) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *UniverseDefinitionTaskParamsResp) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *UniverseDefinitionTaskParamsResp) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *UniverseDefinitionTaskParamsResp) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *UniverseDefinitionTaskParamsResp) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *UniverseDefinitionTaskParamsResp) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCurrentClusterType

`func (o *UniverseDefinitionTaskParamsResp) GetCurrentClusterType() string`

GetCurrentClusterType returns the CurrentClusterType field if non-nil, zero value otherwise.

### GetCurrentClusterTypeOk

`func (o *UniverseDefinitionTaskParamsResp) GetCurrentClusterTypeOk() (*string, bool)`

GetCurrentClusterTypeOk returns a tuple with the CurrentClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClusterType

`func (o *UniverseDefinitionTaskParamsResp) SetCurrentClusterType(v string)`

SetCurrentClusterType sets CurrentClusterType field to given value.

### HasCurrentClusterType

`func (o *UniverseDefinitionTaskParamsResp) HasCurrentClusterType() bool`

HasCurrentClusterType returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *UniverseDefinitionTaskParamsResp) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *UniverseDefinitionTaskParamsResp) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *UniverseDefinitionTaskParamsResp) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *UniverseDefinitionTaskParamsResp) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableYbc

`func (o *UniverseDefinitionTaskParamsResp) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *UniverseDefinitionTaskParamsResp) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *UniverseDefinitionTaskParamsResp) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *UniverseDefinitionTaskParamsResp) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *UniverseDefinitionTaskParamsResp) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *UniverseDefinitionTaskParamsResp) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *UniverseDefinitionTaskParamsResp) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *UniverseDefinitionTaskParamsResp) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *UniverseDefinitionTaskParamsResp) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *UniverseDefinitionTaskParamsResp) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *UniverseDefinitionTaskParamsResp) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *UniverseDefinitionTaskParamsResp) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *UniverseDefinitionTaskParamsResp) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *UniverseDefinitionTaskParamsResp) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *UniverseDefinitionTaskParamsResp) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *UniverseDefinitionTaskParamsResp) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *UniverseDefinitionTaskParamsResp) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *UniverseDefinitionTaskParamsResp) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *UniverseDefinitionTaskParamsResp) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *UniverseDefinitionTaskParamsResp) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetImportedState

`func (o *UniverseDefinitionTaskParamsResp) GetImportedState() string`

GetImportedState returns the ImportedState field if non-nil, zero value otherwise.

### GetImportedStateOk

`func (o *UniverseDefinitionTaskParamsResp) GetImportedStateOk() (*string, bool)`

GetImportedStateOk returns a tuple with the ImportedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedState

`func (o *UniverseDefinitionTaskParamsResp) SetImportedState(v string)`

SetImportedState sets ImportedState field to given value.

### HasImportedState

`func (o *UniverseDefinitionTaskParamsResp) HasImportedState() bool`

HasImportedState returns a boolean if a field has been set.

### GetInstallYbc

`func (o *UniverseDefinitionTaskParamsResp) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *UniverseDefinitionTaskParamsResp) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *UniverseDefinitionTaskParamsResp) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *UniverseDefinitionTaskParamsResp) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetIsKubernetesOperatorControlled

`func (o *UniverseDefinitionTaskParamsResp) GetIsKubernetesOperatorControlled() bool`

GetIsKubernetesOperatorControlled returns the IsKubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetIsKubernetesOperatorControlledOk

`func (o *UniverseDefinitionTaskParamsResp) GetIsKubernetesOperatorControlledOk() (*bool, bool)`

GetIsKubernetesOperatorControlledOk returns a tuple with the IsKubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetesOperatorControlled

`func (o *UniverseDefinitionTaskParamsResp) SetIsKubernetesOperatorControlled(v bool)`

SetIsKubernetesOperatorControlled sets IsKubernetesOperatorControlled field to given value.

### HasIsKubernetesOperatorControlled

`func (o *UniverseDefinitionTaskParamsResp) HasIsKubernetesOperatorControlled() bool`

HasIsKubernetesOperatorControlled returns a boolean if a field has been set.

### GetItestS3PackagePath

`func (o *UniverseDefinitionTaskParamsResp) GetItestS3PackagePath() string`

GetItestS3PackagePath returns the ItestS3PackagePath field if non-nil, zero value otherwise.

### GetItestS3PackagePathOk

`func (o *UniverseDefinitionTaskParamsResp) GetItestS3PackagePathOk() (*string, bool)`

GetItestS3PackagePathOk returns a tuple with the ItestS3PackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItestS3PackagePath

`func (o *UniverseDefinitionTaskParamsResp) SetItestS3PackagePath(v string)`

SetItestS3PackagePath sets ItestS3PackagePath field to given value.

### HasItestS3PackagePath

`func (o *UniverseDefinitionTaskParamsResp) HasItestS3PackagePath() bool`

HasItestS3PackagePath returns a boolean if a field has been set.

### GetMastersInDefaultRegion

`func (o *UniverseDefinitionTaskParamsResp) GetMastersInDefaultRegion() bool`

GetMastersInDefaultRegion returns the MastersInDefaultRegion field if non-nil, zero value otherwise.

### GetMastersInDefaultRegionOk

`func (o *UniverseDefinitionTaskParamsResp) GetMastersInDefaultRegionOk() (*bool, bool)`

GetMastersInDefaultRegionOk returns a tuple with the MastersInDefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastersInDefaultRegion

`func (o *UniverseDefinitionTaskParamsResp) SetMastersInDefaultRegion(v bool)`

SetMastersInDefaultRegion sets MastersInDefaultRegion field to given value.

### HasMastersInDefaultRegion

`func (o *UniverseDefinitionTaskParamsResp) HasMastersInDefaultRegion() bool`

HasMastersInDefaultRegion returns a boolean if a field has been set.

### GetNextClusterIndex

`func (o *UniverseDefinitionTaskParamsResp) GetNextClusterIndex() int32`

GetNextClusterIndex returns the NextClusterIndex field if non-nil, zero value otherwise.

### GetNextClusterIndexOk

`func (o *UniverseDefinitionTaskParamsResp) GetNextClusterIndexOk() (*int32, bool)`

GetNextClusterIndexOk returns a tuple with the NextClusterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextClusterIndex

`func (o *UniverseDefinitionTaskParamsResp) SetNextClusterIndex(v int32)`

SetNextClusterIndex sets NextClusterIndex field to given value.

### HasNextClusterIndex

`func (o *UniverseDefinitionTaskParamsResp) HasNextClusterIndex() bool`

HasNextClusterIndex returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *UniverseDefinitionTaskParamsResp) GetNodeDetailsSet() []NodeDetailsResp`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *UniverseDefinitionTaskParamsResp) GetNodeDetailsSetOk() (*[]NodeDetailsResp, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *UniverseDefinitionTaskParamsResp) SetNodeDetailsSet(v []NodeDetailsResp)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *UniverseDefinitionTaskParamsResp) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *UniverseDefinitionTaskParamsResp) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *UniverseDefinitionTaskParamsResp) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *UniverseDefinitionTaskParamsResp) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *UniverseDefinitionTaskParamsResp) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNodePrefix

`func (o *UniverseDefinitionTaskParamsResp) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *UniverseDefinitionTaskParamsResp) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *UniverseDefinitionTaskParamsResp) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *UniverseDefinitionTaskParamsResp) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetNodesResizeAvailable

`func (o *UniverseDefinitionTaskParamsResp) GetNodesResizeAvailable() bool`

GetNodesResizeAvailable returns the NodesResizeAvailable field if non-nil, zero value otherwise.

### GetNodesResizeAvailableOk

`func (o *UniverseDefinitionTaskParamsResp) GetNodesResizeAvailableOk() (*bool, bool)`

GetNodesResizeAvailableOk returns a tuple with the NodesResizeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesResizeAvailable

`func (o *UniverseDefinitionTaskParamsResp) SetNodesResizeAvailable(v bool)`

SetNodesResizeAvailable sets NodesResizeAvailable field to given value.

### HasNodesResizeAvailable

`func (o *UniverseDefinitionTaskParamsResp) HasNodesResizeAvailable() bool`

HasNodesResizeAvailable returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *UniverseDefinitionTaskParamsResp) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *UniverseDefinitionTaskParamsResp) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *UniverseDefinitionTaskParamsResp) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPlatformVersion

`func (o *UniverseDefinitionTaskParamsResp) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *UniverseDefinitionTaskParamsResp) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *UniverseDefinitionTaskParamsResp) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.


### GetPreviousTaskUUID

`func (o *UniverseDefinitionTaskParamsResp) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *UniverseDefinitionTaskParamsResp) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *UniverseDefinitionTaskParamsResp) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *UniverseDefinitionTaskParamsResp) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRemotePackagePath

`func (o *UniverseDefinitionTaskParamsResp) GetRemotePackagePath() string`

GetRemotePackagePath returns the RemotePackagePath field if non-nil, zero value otherwise.

### GetRemotePackagePathOk

`func (o *UniverseDefinitionTaskParamsResp) GetRemotePackagePathOk() (*string, bool)`

GetRemotePackagePathOk returns a tuple with the RemotePackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePackagePath

`func (o *UniverseDefinitionTaskParamsResp) SetRemotePackagePath(v string)`

SetRemotePackagePath sets RemotePackagePath field to given value.

### HasRemotePackagePath

`func (o *UniverseDefinitionTaskParamsResp) HasRemotePackagePath() bool`

HasRemotePackagePath returns a boolean if a field has been set.

### GetResetAZConfig

`func (o *UniverseDefinitionTaskParamsResp) GetResetAZConfig() bool`

GetResetAZConfig returns the ResetAZConfig field if non-nil, zero value otherwise.

### GetResetAZConfigOk

`func (o *UniverseDefinitionTaskParamsResp) GetResetAZConfigOk() (*bool, bool)`

GetResetAZConfigOk returns a tuple with the ResetAZConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAZConfig

`func (o *UniverseDefinitionTaskParamsResp) SetResetAZConfig(v bool)`

SetResetAZConfig sets ResetAZConfig field to given value.

### HasResetAZConfig

`func (o *UniverseDefinitionTaskParamsResp) HasResetAZConfig() bool`

HasResetAZConfig returns a boolean if a field has been set.

### GetRootAndClientRootCASame

`func (o *UniverseDefinitionTaskParamsResp) GetRootAndClientRootCASame() bool`

GetRootAndClientRootCASame returns the RootAndClientRootCASame field if non-nil, zero value otherwise.

### GetRootAndClientRootCASameOk

`func (o *UniverseDefinitionTaskParamsResp) GetRootAndClientRootCASameOk() (*bool, bool)`

GetRootAndClientRootCASameOk returns a tuple with the RootAndClientRootCASame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootAndClientRootCASame

`func (o *UniverseDefinitionTaskParamsResp) SetRootAndClientRootCASame(v bool)`

SetRootAndClientRootCASame sets RootAndClientRootCASame field to given value.

### HasRootAndClientRootCASame

`func (o *UniverseDefinitionTaskParamsResp) HasRootAndClientRootCASame() bool`

HasRootAndClientRootCASame returns a boolean if a field has been set.

### GetRootCA

`func (o *UniverseDefinitionTaskParamsResp) GetRootCA() string`

GetRootCA returns the RootCA field if non-nil, zero value otherwise.

### GetRootCAOk

`func (o *UniverseDefinitionTaskParamsResp) GetRootCAOk() (*string, bool)`

GetRootCAOk returns a tuple with the RootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCA

`func (o *UniverseDefinitionTaskParamsResp) SetRootCA(v string)`

SetRootCA sets RootCA field to given value.

### HasRootCA

`func (o *UniverseDefinitionTaskParamsResp) HasRootCA() bool`

HasRootCA returns a boolean if a field has been set.

### GetSetTxnTableWaitCountFlag

`func (o *UniverseDefinitionTaskParamsResp) GetSetTxnTableWaitCountFlag() bool`

GetSetTxnTableWaitCountFlag returns the SetTxnTableWaitCountFlag field if non-nil, zero value otherwise.

### GetSetTxnTableWaitCountFlagOk

`func (o *UniverseDefinitionTaskParamsResp) GetSetTxnTableWaitCountFlagOk() (*bool, bool)`

GetSetTxnTableWaitCountFlagOk returns a tuple with the SetTxnTableWaitCountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTxnTableWaitCountFlag

`func (o *UniverseDefinitionTaskParamsResp) SetSetTxnTableWaitCountFlag(v bool)`

SetSetTxnTableWaitCountFlag sets SetTxnTableWaitCountFlag field to given value.

### HasSetTxnTableWaitCountFlag

`func (o *UniverseDefinitionTaskParamsResp) HasSetTxnTableWaitCountFlag() bool`

HasSetTxnTableWaitCountFlag returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *UniverseDefinitionTaskParamsResp) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseDefinitionTaskParamsResp) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseDefinitionTaskParamsResp) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *UniverseDefinitionTaskParamsResp) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *UniverseDefinitionTaskParamsResp) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *UniverseDefinitionTaskParamsResp) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSourceXClusterConfigs

`func (o *UniverseDefinitionTaskParamsResp) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *UniverseDefinitionTaskParamsResp) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *UniverseDefinitionTaskParamsResp) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *UniverseDefinitionTaskParamsResp) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSshUserOverride

`func (o *UniverseDefinitionTaskParamsResp) GetSshUserOverride() string`

GetSshUserOverride returns the SshUserOverride field if non-nil, zero value otherwise.

### GetSshUserOverrideOk

`func (o *UniverseDefinitionTaskParamsResp) GetSshUserOverrideOk() (*string, bool)`

GetSshUserOverrideOk returns a tuple with the SshUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverride

`func (o *UniverseDefinitionTaskParamsResp) SetSshUserOverride(v string)`

SetSshUserOverride sets SshUserOverride field to given value.

### HasSshUserOverride

`func (o *UniverseDefinitionTaskParamsResp) HasSshUserOverride() bool`

HasSshUserOverride returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *UniverseDefinitionTaskParamsResp) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *UniverseDefinitionTaskParamsResp) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *UniverseDefinitionTaskParamsResp) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *UniverseDefinitionTaskParamsResp) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetUniversePaused

`func (o *UniverseDefinitionTaskParamsResp) GetUniversePaused() bool`

GetUniversePaused returns the UniversePaused field if non-nil, zero value otherwise.

### GetUniversePausedOk

`func (o *UniverseDefinitionTaskParamsResp) GetUniversePausedOk() (*bool, bool)`

GetUniversePausedOk returns a tuple with the UniversePaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversePaused

`func (o *UniverseDefinitionTaskParamsResp) SetUniversePaused(v bool)`

SetUniversePaused sets UniversePaused field to given value.

### HasUniversePaused

`func (o *UniverseDefinitionTaskParamsResp) HasUniversePaused() bool`

HasUniversePaused returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *UniverseDefinitionTaskParamsResp) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *UniverseDefinitionTaskParamsResp) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *UniverseDefinitionTaskParamsResp) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *UniverseDefinitionTaskParamsResp) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUpdateInProgress

`func (o *UniverseDefinitionTaskParamsResp) GetUpdateInProgress() bool`

GetUpdateInProgress returns the UpdateInProgress field if non-nil, zero value otherwise.

### GetUpdateInProgressOk

`func (o *UniverseDefinitionTaskParamsResp) GetUpdateInProgressOk() (*bool, bool)`

GetUpdateInProgressOk returns a tuple with the UpdateInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInProgress

`func (o *UniverseDefinitionTaskParamsResp) SetUpdateInProgress(v bool)`

SetUpdateInProgress sets UpdateInProgress field to given value.

### HasUpdateInProgress

`func (o *UniverseDefinitionTaskParamsResp) HasUpdateInProgress() bool`

HasUpdateInProgress returns a boolean if a field has been set.

### GetUpdateOptions

`func (o *UniverseDefinitionTaskParamsResp) GetUpdateOptions() []string`

GetUpdateOptions returns the UpdateOptions field if non-nil, zero value otherwise.

### GetUpdateOptionsOk

`func (o *UniverseDefinitionTaskParamsResp) GetUpdateOptionsOk() (*[]string, bool)`

GetUpdateOptionsOk returns a tuple with the UpdateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOptions

`func (o *UniverseDefinitionTaskParamsResp) SetUpdateOptions(v []string)`

SetUpdateOptions sets UpdateOptions field to given value.

### HasUpdateOptions

`func (o *UniverseDefinitionTaskParamsResp) HasUpdateOptions() bool`

HasUpdateOptions returns a boolean if a field has been set.

### GetUpdateSucceeded

`func (o *UniverseDefinitionTaskParamsResp) GetUpdateSucceeded() bool`

GetUpdateSucceeded returns the UpdateSucceeded field if non-nil, zero value otherwise.

### GetUpdateSucceededOk

`func (o *UniverseDefinitionTaskParamsResp) GetUpdateSucceededOk() (*bool, bool)`

GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceeded

`func (o *UniverseDefinitionTaskParamsResp) SetUpdateSucceeded(v bool)`

SetUpdateSucceeded sets UpdateSucceeded field to given value.

### HasUpdateSucceeded

`func (o *UniverseDefinitionTaskParamsResp) HasUpdateSucceeded() bool`

HasUpdateSucceeded returns a boolean if a field has been set.

### GetUpdatingTask

`func (o *UniverseDefinitionTaskParamsResp) GetUpdatingTask() string`

GetUpdatingTask returns the UpdatingTask field if non-nil, zero value otherwise.

### GetUpdatingTaskOk

`func (o *UniverseDefinitionTaskParamsResp) GetUpdatingTaskOk() (*string, bool)`

GetUpdatingTaskOk returns a tuple with the UpdatingTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTask

`func (o *UniverseDefinitionTaskParamsResp) SetUpdatingTask(v string)`

SetUpdatingTask sets UpdatingTask field to given value.

### HasUpdatingTask

`func (o *UniverseDefinitionTaskParamsResp) HasUpdatingTask() bool`

HasUpdatingTask returns a boolean if a field has been set.

### GetUpdatingTaskUUID

`func (o *UniverseDefinitionTaskParamsResp) GetUpdatingTaskUUID() string`

GetUpdatingTaskUUID returns the UpdatingTaskUUID field if non-nil, zero value otherwise.

### GetUpdatingTaskUUIDOk

`func (o *UniverseDefinitionTaskParamsResp) GetUpdatingTaskUUIDOk() (*string, bool)`

GetUpdatingTaskUUIDOk returns a tuple with the UpdatingTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTaskUUID

`func (o *UniverseDefinitionTaskParamsResp) SetUpdatingTaskUUID(v string)`

SetUpdatingTaskUUID sets UpdatingTaskUUID field to given value.

### HasUpdatingTaskUUID

`func (o *UniverseDefinitionTaskParamsResp) HasUpdatingTaskUUID() bool`

HasUpdatingTaskUUID returns a boolean if a field has been set.

### GetUseNewHelmNamingStyle

`func (o *UniverseDefinitionTaskParamsResp) GetUseNewHelmNamingStyle() bool`

GetUseNewHelmNamingStyle returns the UseNewHelmNamingStyle field if non-nil, zero value otherwise.

### GetUseNewHelmNamingStyleOk

`func (o *UniverseDefinitionTaskParamsResp) GetUseNewHelmNamingStyleOk() (*bool, bool)`

GetUseNewHelmNamingStyleOk returns a tuple with the UseNewHelmNamingStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNewHelmNamingStyle

`func (o *UniverseDefinitionTaskParamsResp) SetUseNewHelmNamingStyle(v bool)`

SetUseNewHelmNamingStyle sets UseNewHelmNamingStyle field to given value.

### HasUseNewHelmNamingStyle

`func (o *UniverseDefinitionTaskParamsResp) HasUseNewHelmNamingStyle() bool`

HasUseNewHelmNamingStyle returns a boolean if a field has been set.

### GetUserAZSelected

`func (o *UniverseDefinitionTaskParamsResp) GetUserAZSelected() bool`

GetUserAZSelected returns the UserAZSelected field if non-nil, zero value otherwise.

### GetUserAZSelectedOk

`func (o *UniverseDefinitionTaskParamsResp) GetUserAZSelectedOk() (*bool, bool)`

GetUserAZSelectedOk returns a tuple with the UserAZSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAZSelected

`func (o *UniverseDefinitionTaskParamsResp) SetUserAZSelected(v bool)`

SetUserAZSelected sets UserAZSelected field to given value.

### HasUserAZSelected

`func (o *UniverseDefinitionTaskParamsResp) HasUserAZSelected() bool`

HasUserAZSelected returns a boolean if a field has been set.

### GetXclusterInfo

`func (o *UniverseDefinitionTaskParamsResp) GetXclusterInfo() XClusterInfo`

GetXclusterInfo returns the XclusterInfo field if non-nil, zero value otherwise.

### GetXclusterInfoOk

`func (o *UniverseDefinitionTaskParamsResp) GetXclusterInfoOk() (*XClusterInfo, bool)`

GetXclusterInfoOk returns a tuple with the XclusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterInfo

`func (o *UniverseDefinitionTaskParamsResp) SetXclusterInfo(v XClusterInfo)`

SetXclusterInfo sets XclusterInfo field to given value.

### HasXclusterInfo

`func (o *UniverseDefinitionTaskParamsResp) HasXclusterInfo() bool`

HasXclusterInfo returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *UniverseDefinitionTaskParamsResp) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *UniverseDefinitionTaskParamsResp) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *UniverseDefinitionTaskParamsResp) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *UniverseDefinitionTaskParamsResp) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *UniverseDefinitionTaskParamsResp) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *UniverseDefinitionTaskParamsResp) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *UniverseDefinitionTaskParamsResp) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *UniverseDefinitionTaskParamsResp) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *UniverseDefinitionTaskParamsResp) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *UniverseDefinitionTaskParamsResp) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *UniverseDefinitionTaskParamsResp) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *UniverseDefinitionTaskParamsResp) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


