# AuditLogConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInsecure** | Pointer to **bool** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**AuditLogConfig** | [**AuditLogConfig**](AuditLogConfig.md) |  | 
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
**InstallNodeAgent** | Pointer to **bool** | YbaApi Internal. Install node agent in background if it is true | [optional] 
**InstallOtelCollector** | **bool** |  | 
**InstallYbc** | Pointer to **bool** |  | [optional] 
**IsKubernetesOperatorControlled** | Pointer to **bool** |  | [optional] 
**IsSoftwareRollbackAllowed** | Pointer to **bool** | Available since YBA version 2.20.2.0 | [optional] [readonly] 
**ItestS3PackagePath** | Pointer to **string** |  | [optional] 
**KubernetesUpgradeSupported** | **bool** |  | 
**MastersInDefaultRegion** | Pointer to **bool** |  | [optional] 
**NextClusterIndex** | Pointer to **int32** |  | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**NodePrefix** | Pointer to **string** |  | [optional] 
**NodesResizeAvailable** | Pointer to **bool** |  | [optional] 
**OtelCollectorEnabled** | Pointer to **bool** | YbaApi Internal. OpenTelemetry Collector enabled for universe | [optional] 
**PlacementModificationTaskUuid** | Pointer to **string** |  | [optional] 
**PlatformUrl** | **string** |  | 
**PlatformVersion** | **string** |  | 
**PrevYBSoftwareConfig** | Pointer to [**PrevYBSoftwareConfig**](PrevYBSoftwareConfig.md) |  | [optional] 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID of a retry | [optional] 
**RemotePackagePath** | Pointer to **string** |  | [optional] 
**ResetAZConfig** | Pointer to **bool** |  | [optional] 
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 
**RootAndClientRootCASame** | Pointer to **bool** |  | [optional] 
**RootCA** | Pointer to **string** |  | [optional] 
**SetTxnTableWaitCountFlag** | Pointer to **bool** |  | [optional] 
**SkipNodeChecks** | Pointer to **bool** | YbaApi Internal. Whether to skip node prechecks while performing rolling upgrade | [optional] 
**SleepAfterMasterRestartMillis** | **int32** |  | 
**SleepAfterTServerRestartMillis** | **int32** |  | 
**SoftwareUpgradeState** | Pointer to **string** |  | [optional] 
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

### NewAuditLogConfigParams

`func NewAuditLogConfigParams(auditLogConfig AuditLogConfig, clusters []Cluster, creatingUser Users, installOtelCollector bool, kubernetesUpgradeSupported bool, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, upgradeOption string, ) *AuditLogConfigParams`

NewAuditLogConfigParams instantiates a new AuditLogConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogConfigParamsWithDefaults

`func NewAuditLogConfigParamsWithDefaults() *AuditLogConfigParams`

NewAuditLogConfigParamsWithDefaults instantiates a new AuditLogConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *AuditLogConfigParams) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *AuditLogConfigParams) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *AuditLogConfigParams) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *AuditLogConfigParams) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetArch

`func (o *AuditLogConfigParams) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *AuditLogConfigParams) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *AuditLogConfigParams) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *AuditLogConfigParams) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetAuditLogConfig

`func (o *AuditLogConfigParams) GetAuditLogConfig() AuditLogConfig`

GetAuditLogConfig returns the AuditLogConfig field if non-nil, zero value otherwise.

### GetAuditLogConfigOk

`func (o *AuditLogConfigParams) GetAuditLogConfigOk() (*AuditLogConfig, bool)`

GetAuditLogConfigOk returns a tuple with the AuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogConfig

`func (o *AuditLogConfigParams) SetAuditLogConfig(v AuditLogConfig)`

SetAuditLogConfig sets AuditLogConfig field to given value.


### GetCapability

`func (o *AuditLogConfigParams) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *AuditLogConfigParams) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *AuditLogConfigParams) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *AuditLogConfigParams) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetClientRootCA

`func (o *AuditLogConfigParams) GetClientRootCA() string`

GetClientRootCA returns the ClientRootCA field if non-nil, zero value otherwise.

### GetClientRootCAOk

`func (o *AuditLogConfigParams) GetClientRootCAOk() (*string, bool)`

GetClientRootCAOk returns a tuple with the ClientRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCA

`func (o *AuditLogConfigParams) SetClientRootCA(v string)`

SetClientRootCA sets ClientRootCA field to given value.

### HasClientRootCA

`func (o *AuditLogConfigParams) HasClientRootCA() bool`

HasClientRootCA returns a boolean if a field has been set.

### GetClusters

`func (o *AuditLogConfigParams) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *AuditLogConfigParams) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *AuditLogConfigParams) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.


### GetCmkArn

`func (o *AuditLogConfigParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *AuditLogConfigParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *AuditLogConfigParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *AuditLogConfigParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *AuditLogConfigParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *AuditLogConfigParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *AuditLogConfigParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *AuditLogConfigParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *AuditLogConfigParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *AuditLogConfigParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *AuditLogConfigParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCurrentClusterType

`func (o *AuditLogConfigParams) GetCurrentClusterType() string`

GetCurrentClusterType returns the CurrentClusterType field if non-nil, zero value otherwise.

### GetCurrentClusterTypeOk

`func (o *AuditLogConfigParams) GetCurrentClusterTypeOk() (*string, bool)`

GetCurrentClusterTypeOk returns a tuple with the CurrentClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClusterType

`func (o *AuditLogConfigParams) SetCurrentClusterType(v string)`

SetCurrentClusterType sets CurrentClusterType field to given value.

### HasCurrentClusterType

`func (o *AuditLogConfigParams) HasCurrentClusterType() bool`

HasCurrentClusterType returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *AuditLogConfigParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *AuditLogConfigParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *AuditLogConfigParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *AuditLogConfigParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableYbc

`func (o *AuditLogConfigParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *AuditLogConfigParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *AuditLogConfigParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *AuditLogConfigParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *AuditLogConfigParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *AuditLogConfigParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *AuditLogConfigParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *AuditLogConfigParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *AuditLogConfigParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *AuditLogConfigParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *AuditLogConfigParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *AuditLogConfigParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *AuditLogConfigParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *AuditLogConfigParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *AuditLogConfigParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *AuditLogConfigParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *AuditLogConfigParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *AuditLogConfigParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *AuditLogConfigParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *AuditLogConfigParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetImportedState

`func (o *AuditLogConfigParams) GetImportedState() string`

GetImportedState returns the ImportedState field if non-nil, zero value otherwise.

### GetImportedStateOk

`func (o *AuditLogConfigParams) GetImportedStateOk() (*string, bool)`

GetImportedStateOk returns a tuple with the ImportedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedState

`func (o *AuditLogConfigParams) SetImportedState(v string)`

SetImportedState sets ImportedState field to given value.

### HasImportedState

`func (o *AuditLogConfigParams) HasImportedState() bool`

HasImportedState returns a boolean if a field has been set.

### GetInstallNodeAgent

`func (o *AuditLogConfigParams) GetInstallNodeAgent() bool`

GetInstallNodeAgent returns the InstallNodeAgent field if non-nil, zero value otherwise.

### GetInstallNodeAgentOk

`func (o *AuditLogConfigParams) GetInstallNodeAgentOk() (*bool, bool)`

GetInstallNodeAgentOk returns a tuple with the InstallNodeAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallNodeAgent

`func (o *AuditLogConfigParams) SetInstallNodeAgent(v bool)`

SetInstallNodeAgent sets InstallNodeAgent field to given value.

### HasInstallNodeAgent

`func (o *AuditLogConfigParams) HasInstallNodeAgent() bool`

HasInstallNodeAgent returns a boolean if a field has been set.

### GetInstallOtelCollector

`func (o *AuditLogConfigParams) GetInstallOtelCollector() bool`

GetInstallOtelCollector returns the InstallOtelCollector field if non-nil, zero value otherwise.

### GetInstallOtelCollectorOk

`func (o *AuditLogConfigParams) GetInstallOtelCollectorOk() (*bool, bool)`

GetInstallOtelCollectorOk returns a tuple with the InstallOtelCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallOtelCollector

`func (o *AuditLogConfigParams) SetInstallOtelCollector(v bool)`

SetInstallOtelCollector sets InstallOtelCollector field to given value.


### GetInstallYbc

`func (o *AuditLogConfigParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *AuditLogConfigParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *AuditLogConfigParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *AuditLogConfigParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetIsKubernetesOperatorControlled

`func (o *AuditLogConfigParams) GetIsKubernetesOperatorControlled() bool`

GetIsKubernetesOperatorControlled returns the IsKubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetIsKubernetesOperatorControlledOk

`func (o *AuditLogConfigParams) GetIsKubernetesOperatorControlledOk() (*bool, bool)`

GetIsKubernetesOperatorControlledOk returns a tuple with the IsKubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetesOperatorControlled

`func (o *AuditLogConfigParams) SetIsKubernetesOperatorControlled(v bool)`

SetIsKubernetesOperatorControlled sets IsKubernetesOperatorControlled field to given value.

### HasIsKubernetesOperatorControlled

`func (o *AuditLogConfigParams) HasIsKubernetesOperatorControlled() bool`

HasIsKubernetesOperatorControlled returns a boolean if a field has been set.

### GetIsSoftwareRollbackAllowed

`func (o *AuditLogConfigParams) GetIsSoftwareRollbackAllowed() bool`

GetIsSoftwareRollbackAllowed returns the IsSoftwareRollbackAllowed field if non-nil, zero value otherwise.

### GetIsSoftwareRollbackAllowedOk

`func (o *AuditLogConfigParams) GetIsSoftwareRollbackAllowedOk() (*bool, bool)`

GetIsSoftwareRollbackAllowedOk returns a tuple with the IsSoftwareRollbackAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftwareRollbackAllowed

`func (o *AuditLogConfigParams) SetIsSoftwareRollbackAllowed(v bool)`

SetIsSoftwareRollbackAllowed sets IsSoftwareRollbackAllowed field to given value.

### HasIsSoftwareRollbackAllowed

`func (o *AuditLogConfigParams) HasIsSoftwareRollbackAllowed() bool`

HasIsSoftwareRollbackAllowed returns a boolean if a field has been set.

### GetItestS3PackagePath

`func (o *AuditLogConfigParams) GetItestS3PackagePath() string`

GetItestS3PackagePath returns the ItestS3PackagePath field if non-nil, zero value otherwise.

### GetItestS3PackagePathOk

`func (o *AuditLogConfigParams) GetItestS3PackagePathOk() (*string, bool)`

GetItestS3PackagePathOk returns a tuple with the ItestS3PackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItestS3PackagePath

`func (o *AuditLogConfigParams) SetItestS3PackagePath(v string)`

SetItestS3PackagePath sets ItestS3PackagePath field to given value.

### HasItestS3PackagePath

`func (o *AuditLogConfigParams) HasItestS3PackagePath() bool`

HasItestS3PackagePath returns a boolean if a field has been set.

### GetKubernetesUpgradeSupported

`func (o *AuditLogConfigParams) GetKubernetesUpgradeSupported() bool`

GetKubernetesUpgradeSupported returns the KubernetesUpgradeSupported field if non-nil, zero value otherwise.

### GetKubernetesUpgradeSupportedOk

`func (o *AuditLogConfigParams) GetKubernetesUpgradeSupportedOk() (*bool, bool)`

GetKubernetesUpgradeSupportedOk returns a tuple with the KubernetesUpgradeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesUpgradeSupported

`func (o *AuditLogConfigParams) SetKubernetesUpgradeSupported(v bool)`

SetKubernetesUpgradeSupported sets KubernetesUpgradeSupported field to given value.


### GetMastersInDefaultRegion

`func (o *AuditLogConfigParams) GetMastersInDefaultRegion() bool`

GetMastersInDefaultRegion returns the MastersInDefaultRegion field if non-nil, zero value otherwise.

### GetMastersInDefaultRegionOk

`func (o *AuditLogConfigParams) GetMastersInDefaultRegionOk() (*bool, bool)`

GetMastersInDefaultRegionOk returns a tuple with the MastersInDefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastersInDefaultRegion

`func (o *AuditLogConfigParams) SetMastersInDefaultRegion(v bool)`

SetMastersInDefaultRegion sets MastersInDefaultRegion field to given value.

### HasMastersInDefaultRegion

`func (o *AuditLogConfigParams) HasMastersInDefaultRegion() bool`

HasMastersInDefaultRegion returns a boolean if a field has been set.

### GetNextClusterIndex

`func (o *AuditLogConfigParams) GetNextClusterIndex() int32`

GetNextClusterIndex returns the NextClusterIndex field if non-nil, zero value otherwise.

### GetNextClusterIndexOk

`func (o *AuditLogConfigParams) GetNextClusterIndexOk() (*int32, bool)`

GetNextClusterIndexOk returns a tuple with the NextClusterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextClusterIndex

`func (o *AuditLogConfigParams) SetNextClusterIndex(v int32)`

SetNextClusterIndex sets NextClusterIndex field to given value.

### HasNextClusterIndex

`func (o *AuditLogConfigParams) HasNextClusterIndex() bool`

HasNextClusterIndex returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *AuditLogConfigParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *AuditLogConfigParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *AuditLogConfigParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *AuditLogConfigParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *AuditLogConfigParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *AuditLogConfigParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *AuditLogConfigParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *AuditLogConfigParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNodePrefix

`func (o *AuditLogConfigParams) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *AuditLogConfigParams) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *AuditLogConfigParams) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *AuditLogConfigParams) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetNodesResizeAvailable

`func (o *AuditLogConfigParams) GetNodesResizeAvailable() bool`

GetNodesResizeAvailable returns the NodesResizeAvailable field if non-nil, zero value otherwise.

### GetNodesResizeAvailableOk

`func (o *AuditLogConfigParams) GetNodesResizeAvailableOk() (*bool, bool)`

GetNodesResizeAvailableOk returns a tuple with the NodesResizeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesResizeAvailable

`func (o *AuditLogConfigParams) SetNodesResizeAvailable(v bool)`

SetNodesResizeAvailable sets NodesResizeAvailable field to given value.

### HasNodesResizeAvailable

`func (o *AuditLogConfigParams) HasNodesResizeAvailable() bool`

HasNodesResizeAvailable returns a boolean if a field has been set.

### GetOtelCollectorEnabled

`func (o *AuditLogConfigParams) GetOtelCollectorEnabled() bool`

GetOtelCollectorEnabled returns the OtelCollectorEnabled field if non-nil, zero value otherwise.

### GetOtelCollectorEnabledOk

`func (o *AuditLogConfigParams) GetOtelCollectorEnabledOk() (*bool, bool)`

GetOtelCollectorEnabledOk returns a tuple with the OtelCollectorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelCollectorEnabled

`func (o *AuditLogConfigParams) SetOtelCollectorEnabled(v bool)`

SetOtelCollectorEnabled sets OtelCollectorEnabled field to given value.

### HasOtelCollectorEnabled

`func (o *AuditLogConfigParams) HasOtelCollectorEnabled() bool`

HasOtelCollectorEnabled returns a boolean if a field has been set.

### GetPlacementModificationTaskUuid

`func (o *AuditLogConfigParams) GetPlacementModificationTaskUuid() string`

GetPlacementModificationTaskUuid returns the PlacementModificationTaskUuid field if non-nil, zero value otherwise.

### GetPlacementModificationTaskUuidOk

`func (o *AuditLogConfigParams) GetPlacementModificationTaskUuidOk() (*string, bool)`

GetPlacementModificationTaskUuidOk returns a tuple with the PlacementModificationTaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementModificationTaskUuid

`func (o *AuditLogConfigParams) SetPlacementModificationTaskUuid(v string)`

SetPlacementModificationTaskUuid sets PlacementModificationTaskUuid field to given value.

### HasPlacementModificationTaskUuid

`func (o *AuditLogConfigParams) HasPlacementModificationTaskUuid() bool`

HasPlacementModificationTaskUuid returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *AuditLogConfigParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *AuditLogConfigParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *AuditLogConfigParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPlatformVersion

`func (o *AuditLogConfigParams) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *AuditLogConfigParams) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *AuditLogConfigParams) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.


### GetPrevYBSoftwareConfig

`func (o *AuditLogConfigParams) GetPrevYBSoftwareConfig() PrevYBSoftwareConfig`

GetPrevYBSoftwareConfig returns the PrevYBSoftwareConfig field if non-nil, zero value otherwise.

### GetPrevYBSoftwareConfigOk

`func (o *AuditLogConfigParams) GetPrevYBSoftwareConfigOk() (*PrevYBSoftwareConfig, bool)`

GetPrevYBSoftwareConfigOk returns a tuple with the PrevYBSoftwareConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevYBSoftwareConfig

`func (o *AuditLogConfigParams) SetPrevYBSoftwareConfig(v PrevYBSoftwareConfig)`

SetPrevYBSoftwareConfig sets PrevYBSoftwareConfig field to given value.

### HasPrevYBSoftwareConfig

`func (o *AuditLogConfigParams) HasPrevYBSoftwareConfig() bool`

HasPrevYBSoftwareConfig returns a boolean if a field has been set.

### GetPreviousTaskUUID

`func (o *AuditLogConfigParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *AuditLogConfigParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *AuditLogConfigParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *AuditLogConfigParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRemotePackagePath

`func (o *AuditLogConfigParams) GetRemotePackagePath() string`

GetRemotePackagePath returns the RemotePackagePath field if non-nil, zero value otherwise.

### GetRemotePackagePathOk

`func (o *AuditLogConfigParams) GetRemotePackagePathOk() (*string, bool)`

GetRemotePackagePathOk returns a tuple with the RemotePackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePackagePath

`func (o *AuditLogConfigParams) SetRemotePackagePath(v string)`

SetRemotePackagePath sets RemotePackagePath field to given value.

### HasRemotePackagePath

`func (o *AuditLogConfigParams) HasRemotePackagePath() bool`

HasRemotePackagePath returns a boolean if a field has been set.

### GetResetAZConfig

`func (o *AuditLogConfigParams) GetResetAZConfig() bool`

GetResetAZConfig returns the ResetAZConfig field if non-nil, zero value otherwise.

### GetResetAZConfigOk

`func (o *AuditLogConfigParams) GetResetAZConfigOk() (*bool, bool)`

GetResetAZConfigOk returns a tuple with the ResetAZConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAZConfig

`func (o *AuditLogConfigParams) SetResetAZConfig(v bool)`

SetResetAZConfig sets ResetAZConfig field to given value.

### HasResetAZConfig

`func (o *AuditLogConfigParams) HasResetAZConfig() bool`

HasResetAZConfig returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *AuditLogConfigParams) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *AuditLogConfigParams) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *AuditLogConfigParams) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *AuditLogConfigParams) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetRootAndClientRootCASame

`func (o *AuditLogConfigParams) GetRootAndClientRootCASame() bool`

GetRootAndClientRootCASame returns the RootAndClientRootCASame field if non-nil, zero value otherwise.

### GetRootAndClientRootCASameOk

`func (o *AuditLogConfigParams) GetRootAndClientRootCASameOk() (*bool, bool)`

GetRootAndClientRootCASameOk returns a tuple with the RootAndClientRootCASame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootAndClientRootCASame

`func (o *AuditLogConfigParams) SetRootAndClientRootCASame(v bool)`

SetRootAndClientRootCASame sets RootAndClientRootCASame field to given value.

### HasRootAndClientRootCASame

`func (o *AuditLogConfigParams) HasRootAndClientRootCASame() bool`

HasRootAndClientRootCASame returns a boolean if a field has been set.

### GetRootCA

`func (o *AuditLogConfigParams) GetRootCA() string`

GetRootCA returns the RootCA field if non-nil, zero value otherwise.

### GetRootCAOk

`func (o *AuditLogConfigParams) GetRootCAOk() (*string, bool)`

GetRootCAOk returns a tuple with the RootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCA

`func (o *AuditLogConfigParams) SetRootCA(v string)`

SetRootCA sets RootCA field to given value.

### HasRootCA

`func (o *AuditLogConfigParams) HasRootCA() bool`

HasRootCA returns a boolean if a field has been set.

### GetSetTxnTableWaitCountFlag

`func (o *AuditLogConfigParams) GetSetTxnTableWaitCountFlag() bool`

GetSetTxnTableWaitCountFlag returns the SetTxnTableWaitCountFlag field if non-nil, zero value otherwise.

### GetSetTxnTableWaitCountFlagOk

`func (o *AuditLogConfigParams) GetSetTxnTableWaitCountFlagOk() (*bool, bool)`

GetSetTxnTableWaitCountFlagOk returns a tuple with the SetTxnTableWaitCountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTxnTableWaitCountFlag

`func (o *AuditLogConfigParams) SetSetTxnTableWaitCountFlag(v bool)`

SetSetTxnTableWaitCountFlag sets SetTxnTableWaitCountFlag field to given value.

### HasSetTxnTableWaitCountFlag

`func (o *AuditLogConfigParams) HasSetTxnTableWaitCountFlag() bool`

HasSetTxnTableWaitCountFlag returns a boolean if a field has been set.

### GetSkipNodeChecks

`func (o *AuditLogConfigParams) GetSkipNodeChecks() bool`

GetSkipNodeChecks returns the SkipNodeChecks field if non-nil, zero value otherwise.

### GetSkipNodeChecksOk

`func (o *AuditLogConfigParams) GetSkipNodeChecksOk() (*bool, bool)`

GetSkipNodeChecksOk returns a tuple with the SkipNodeChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNodeChecks

`func (o *AuditLogConfigParams) SetSkipNodeChecks(v bool)`

SetSkipNodeChecks sets SkipNodeChecks field to given value.

### HasSkipNodeChecks

`func (o *AuditLogConfigParams) HasSkipNodeChecks() bool`

HasSkipNodeChecks returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *AuditLogConfigParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *AuditLogConfigParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *AuditLogConfigParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *AuditLogConfigParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *AuditLogConfigParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *AuditLogConfigParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSoftwareUpgradeState

`func (o *AuditLogConfigParams) GetSoftwareUpgradeState() string`

GetSoftwareUpgradeState returns the SoftwareUpgradeState field if non-nil, zero value otherwise.

### GetSoftwareUpgradeStateOk

`func (o *AuditLogConfigParams) GetSoftwareUpgradeStateOk() (*string, bool)`

GetSoftwareUpgradeStateOk returns a tuple with the SoftwareUpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpgradeState

`func (o *AuditLogConfigParams) SetSoftwareUpgradeState(v string)`

SetSoftwareUpgradeState sets SoftwareUpgradeState field to given value.

### HasSoftwareUpgradeState

`func (o *AuditLogConfigParams) HasSoftwareUpgradeState() bool`

HasSoftwareUpgradeState returns a boolean if a field has been set.

### GetSourceXClusterConfigs

`func (o *AuditLogConfigParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *AuditLogConfigParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *AuditLogConfigParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *AuditLogConfigParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSshUserOverride

`func (o *AuditLogConfigParams) GetSshUserOverride() string`

GetSshUserOverride returns the SshUserOverride field if non-nil, zero value otherwise.

### GetSshUserOverrideOk

`func (o *AuditLogConfigParams) GetSshUserOverrideOk() (*string, bool)`

GetSshUserOverrideOk returns a tuple with the SshUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverride

`func (o *AuditLogConfigParams) SetSshUserOverride(v string)`

SetSshUserOverride sets SshUserOverride field to given value.

### HasSshUserOverride

`func (o *AuditLogConfigParams) HasSshUserOverride() bool`

HasSshUserOverride returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *AuditLogConfigParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *AuditLogConfigParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *AuditLogConfigParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *AuditLogConfigParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetUniversePaused

`func (o *AuditLogConfigParams) GetUniversePaused() bool`

GetUniversePaused returns the UniversePaused field if non-nil, zero value otherwise.

### GetUniversePausedOk

`func (o *AuditLogConfigParams) GetUniversePausedOk() (*bool, bool)`

GetUniversePausedOk returns a tuple with the UniversePaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversePaused

`func (o *AuditLogConfigParams) SetUniversePaused(v bool)`

SetUniversePaused sets UniversePaused field to given value.

### HasUniversePaused

`func (o *AuditLogConfigParams) HasUniversePaused() bool`

HasUniversePaused returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *AuditLogConfigParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *AuditLogConfigParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *AuditLogConfigParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *AuditLogConfigParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUpdateInProgress

`func (o *AuditLogConfigParams) GetUpdateInProgress() bool`

GetUpdateInProgress returns the UpdateInProgress field if non-nil, zero value otherwise.

### GetUpdateInProgressOk

`func (o *AuditLogConfigParams) GetUpdateInProgressOk() (*bool, bool)`

GetUpdateInProgressOk returns a tuple with the UpdateInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInProgress

`func (o *AuditLogConfigParams) SetUpdateInProgress(v bool)`

SetUpdateInProgress sets UpdateInProgress field to given value.

### HasUpdateInProgress

`func (o *AuditLogConfigParams) HasUpdateInProgress() bool`

HasUpdateInProgress returns a boolean if a field has been set.

### GetUpdateOptions

`func (o *AuditLogConfigParams) GetUpdateOptions() []string`

GetUpdateOptions returns the UpdateOptions field if non-nil, zero value otherwise.

### GetUpdateOptionsOk

`func (o *AuditLogConfigParams) GetUpdateOptionsOk() (*[]string, bool)`

GetUpdateOptionsOk returns a tuple with the UpdateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOptions

`func (o *AuditLogConfigParams) SetUpdateOptions(v []string)`

SetUpdateOptions sets UpdateOptions field to given value.

### HasUpdateOptions

`func (o *AuditLogConfigParams) HasUpdateOptions() bool`

HasUpdateOptions returns a boolean if a field has been set.

### GetUpdateSucceeded

`func (o *AuditLogConfigParams) GetUpdateSucceeded() bool`

GetUpdateSucceeded returns the UpdateSucceeded field if non-nil, zero value otherwise.

### GetUpdateSucceededOk

`func (o *AuditLogConfigParams) GetUpdateSucceededOk() (*bool, bool)`

GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceeded

`func (o *AuditLogConfigParams) SetUpdateSucceeded(v bool)`

SetUpdateSucceeded sets UpdateSucceeded field to given value.

### HasUpdateSucceeded

`func (o *AuditLogConfigParams) HasUpdateSucceeded() bool`

HasUpdateSucceeded returns a boolean if a field has been set.

### GetUpdatingTask

`func (o *AuditLogConfigParams) GetUpdatingTask() string`

GetUpdatingTask returns the UpdatingTask field if non-nil, zero value otherwise.

### GetUpdatingTaskOk

`func (o *AuditLogConfigParams) GetUpdatingTaskOk() (*string, bool)`

GetUpdatingTaskOk returns a tuple with the UpdatingTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTask

`func (o *AuditLogConfigParams) SetUpdatingTask(v string)`

SetUpdatingTask sets UpdatingTask field to given value.

### HasUpdatingTask

`func (o *AuditLogConfigParams) HasUpdatingTask() bool`

HasUpdatingTask returns a boolean if a field has been set.

### GetUpdatingTaskUUID

`func (o *AuditLogConfigParams) GetUpdatingTaskUUID() string`

GetUpdatingTaskUUID returns the UpdatingTaskUUID field if non-nil, zero value otherwise.

### GetUpdatingTaskUUIDOk

`func (o *AuditLogConfigParams) GetUpdatingTaskUUIDOk() (*string, bool)`

GetUpdatingTaskUUIDOk returns a tuple with the UpdatingTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTaskUUID

`func (o *AuditLogConfigParams) SetUpdatingTaskUUID(v string)`

SetUpdatingTaskUUID sets UpdatingTaskUUID field to given value.

### HasUpdatingTaskUUID

`func (o *AuditLogConfigParams) HasUpdatingTaskUUID() bool`

HasUpdatingTaskUUID returns a boolean if a field has been set.

### GetUpgradeOption

`func (o *AuditLogConfigParams) GetUpgradeOption() string`

GetUpgradeOption returns the UpgradeOption field if non-nil, zero value otherwise.

### GetUpgradeOptionOk

`func (o *AuditLogConfigParams) GetUpgradeOptionOk() (*string, bool)`

GetUpgradeOptionOk returns a tuple with the UpgradeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOption

`func (o *AuditLogConfigParams) SetUpgradeOption(v string)`

SetUpgradeOption sets UpgradeOption field to given value.


### GetUseNewHelmNamingStyle

`func (o *AuditLogConfigParams) GetUseNewHelmNamingStyle() bool`

GetUseNewHelmNamingStyle returns the UseNewHelmNamingStyle field if non-nil, zero value otherwise.

### GetUseNewHelmNamingStyleOk

`func (o *AuditLogConfigParams) GetUseNewHelmNamingStyleOk() (*bool, bool)`

GetUseNewHelmNamingStyleOk returns a tuple with the UseNewHelmNamingStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNewHelmNamingStyle

`func (o *AuditLogConfigParams) SetUseNewHelmNamingStyle(v bool)`

SetUseNewHelmNamingStyle sets UseNewHelmNamingStyle field to given value.

### HasUseNewHelmNamingStyle

`func (o *AuditLogConfigParams) HasUseNewHelmNamingStyle() bool`

HasUseNewHelmNamingStyle returns a boolean if a field has been set.

### GetUserAZSelected

`func (o *AuditLogConfigParams) GetUserAZSelected() bool`

GetUserAZSelected returns the UserAZSelected field if non-nil, zero value otherwise.

### GetUserAZSelectedOk

`func (o *AuditLogConfigParams) GetUserAZSelectedOk() (*bool, bool)`

GetUserAZSelectedOk returns a tuple with the UserAZSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAZSelected

`func (o *AuditLogConfigParams) SetUserAZSelected(v bool)`

SetUserAZSelected sets UserAZSelected field to given value.

### HasUserAZSelected

`func (o *AuditLogConfigParams) HasUserAZSelected() bool`

HasUserAZSelected returns a boolean if a field has been set.

### GetXclusterInfo

`func (o *AuditLogConfigParams) GetXclusterInfo() XClusterInfo`

GetXclusterInfo returns the XclusterInfo field if non-nil, zero value otherwise.

### GetXclusterInfoOk

`func (o *AuditLogConfigParams) GetXclusterInfoOk() (*XClusterInfo, bool)`

GetXclusterInfoOk returns a tuple with the XclusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterInfo

`func (o *AuditLogConfigParams) SetXclusterInfo(v XClusterInfo)`

SetXclusterInfo sets XclusterInfo field to given value.

### HasXclusterInfo

`func (o *AuditLogConfigParams) HasXclusterInfo() bool`

HasXclusterInfo returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *AuditLogConfigParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *AuditLogConfigParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *AuditLogConfigParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *AuditLogConfigParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *AuditLogConfigParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *AuditLogConfigParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *AuditLogConfigParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *AuditLogConfigParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *AuditLogConfigParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *AuditLogConfigParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *AuditLogConfigParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *AuditLogConfigParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


