# GFlagsUpgradeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalServicesStateData** | Pointer to [**AdditionalServicesStateData**](AdditionalServicesStateData.md) |  | [optional] 
**AllowInsecure** | Pointer to **bool** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
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
**InstallYbc** | Pointer to **bool** |  | [optional] 
**IsKubernetesOperatorControlled** | Pointer to **bool** |  | [optional] 
**IsSoftwareRollbackAllowed** | Pointer to **bool** | Available since YBA version 2.20.2.0 | [optional] [readonly] 
**ItestS3PackagePath** | Pointer to **string** |  | [optional] 
**KubernetesUpgradeSupported** | **bool** |  | 
**MasterGFlags** | **map[string]string** |  | 
**MastersInDefaultRegion** | Pointer to **bool** |  | [optional] 
**NextClusterIndex** | Pointer to **int32** |  | [optional] 
**NodeAgentMissing** | Pointer to **bool** | YbaApi Internal. True if a node agent for missing in any of the nodes | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**NodePrefix** | Pointer to **string** |  | [optional] 
**NodesResizeAvailable** | Pointer to **bool** |  | [optional] 
**OtelCollectorEnabled** | Pointer to **bool** | YbaApi Internal. OpenTelemetry Collector enabled for universe | [optional] 
**PlacementModificationTaskUuid** | Pointer to **string** |  | [optional] 
**PlatformUrl** | **string** |  | 
**PlatformVersion** | Pointer to **string** |  | [optional] [readonly] 
**PrevYBSoftwareConfig** | Pointer to [**PrevYBSoftwareConfig**](PrevYBSoftwareConfig.md) |  | [optional] 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID of a retry | [optional] 
**RemotePackagePath** | Pointer to **string** |  | [optional] 
**ResetAZConfig** | Pointer to **bool** |  | [optional] 
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 
**RootAndClientRootCASame** | Pointer to **bool** |  | [optional] 
**RootCA** | Pointer to **string** |  | [optional] 
**RunOnlyPrechecks** | Pointer to **bool** | YbaApi Internal. Run only prechecks during task run | [optional] 
**SetTxnTableWaitCountFlag** | Pointer to **bool** |  | [optional] 
**SkipNodeChecks** | Pointer to **bool** | YbaApi Internal. Whether to skip node prechecks while performing rolling upgrade | [optional] 
**SleepAfterMasterRestartMillis** | **int32** |  | 
**SleepAfterTServerRestartMillis** | **int32** |  | 
**SoftwareUpgradeState** | Pointer to **string** |  | [optional] 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] [readonly] 
**SshUserOverride** | Pointer to **string** |  | [optional] 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] [readonly] 
**TserverGFlags** | **map[string]string** |  | 
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

### NewGFlagsUpgradeParams

`func NewGFlagsUpgradeParams(clusters []Cluster, creatingUser Users, kubernetesUpgradeSupported bool, masterGFlags map[string]string, platformUrl string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, tserverGFlags map[string]string, upgradeOption string, ) *GFlagsUpgradeParams`

NewGFlagsUpgradeParams instantiates a new GFlagsUpgradeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFlagsUpgradeParamsWithDefaults

`func NewGFlagsUpgradeParamsWithDefaults() *GFlagsUpgradeParams`

NewGFlagsUpgradeParamsWithDefaults instantiates a new GFlagsUpgradeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalServicesStateData

`func (o *GFlagsUpgradeParams) GetAdditionalServicesStateData() AdditionalServicesStateData`

GetAdditionalServicesStateData returns the AdditionalServicesStateData field if non-nil, zero value otherwise.

### GetAdditionalServicesStateDataOk

`func (o *GFlagsUpgradeParams) GetAdditionalServicesStateDataOk() (*AdditionalServicesStateData, bool)`

GetAdditionalServicesStateDataOk returns a tuple with the AdditionalServicesStateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalServicesStateData

`func (o *GFlagsUpgradeParams) SetAdditionalServicesStateData(v AdditionalServicesStateData)`

SetAdditionalServicesStateData sets AdditionalServicesStateData field to given value.

### HasAdditionalServicesStateData

`func (o *GFlagsUpgradeParams) HasAdditionalServicesStateData() bool`

HasAdditionalServicesStateData returns a boolean if a field has been set.

### GetAllowInsecure

`func (o *GFlagsUpgradeParams) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *GFlagsUpgradeParams) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *GFlagsUpgradeParams) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *GFlagsUpgradeParams) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetArch

`func (o *GFlagsUpgradeParams) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *GFlagsUpgradeParams) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *GFlagsUpgradeParams) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *GFlagsUpgradeParams) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCapability

`func (o *GFlagsUpgradeParams) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *GFlagsUpgradeParams) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *GFlagsUpgradeParams) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *GFlagsUpgradeParams) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetClientRootCA

`func (o *GFlagsUpgradeParams) GetClientRootCA() string`

GetClientRootCA returns the ClientRootCA field if non-nil, zero value otherwise.

### GetClientRootCAOk

`func (o *GFlagsUpgradeParams) GetClientRootCAOk() (*string, bool)`

GetClientRootCAOk returns a tuple with the ClientRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCA

`func (o *GFlagsUpgradeParams) SetClientRootCA(v string)`

SetClientRootCA sets ClientRootCA field to given value.

### HasClientRootCA

`func (o *GFlagsUpgradeParams) HasClientRootCA() bool`

HasClientRootCA returns a boolean if a field has been set.

### GetClusters

`func (o *GFlagsUpgradeParams) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *GFlagsUpgradeParams) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *GFlagsUpgradeParams) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.


### GetCmkArn

`func (o *GFlagsUpgradeParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *GFlagsUpgradeParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *GFlagsUpgradeParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *GFlagsUpgradeParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *GFlagsUpgradeParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *GFlagsUpgradeParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *GFlagsUpgradeParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *GFlagsUpgradeParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *GFlagsUpgradeParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *GFlagsUpgradeParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *GFlagsUpgradeParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCurrentClusterType

`func (o *GFlagsUpgradeParams) GetCurrentClusterType() string`

GetCurrentClusterType returns the CurrentClusterType field if non-nil, zero value otherwise.

### GetCurrentClusterTypeOk

`func (o *GFlagsUpgradeParams) GetCurrentClusterTypeOk() (*string, bool)`

GetCurrentClusterTypeOk returns a tuple with the CurrentClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClusterType

`func (o *GFlagsUpgradeParams) SetCurrentClusterType(v string)`

SetCurrentClusterType sets CurrentClusterType field to given value.

### HasCurrentClusterType

`func (o *GFlagsUpgradeParams) HasCurrentClusterType() bool`

HasCurrentClusterType returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *GFlagsUpgradeParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *GFlagsUpgradeParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *GFlagsUpgradeParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *GFlagsUpgradeParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableYbc

`func (o *GFlagsUpgradeParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *GFlagsUpgradeParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *GFlagsUpgradeParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *GFlagsUpgradeParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *GFlagsUpgradeParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *GFlagsUpgradeParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *GFlagsUpgradeParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *GFlagsUpgradeParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *GFlagsUpgradeParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *GFlagsUpgradeParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *GFlagsUpgradeParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *GFlagsUpgradeParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *GFlagsUpgradeParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *GFlagsUpgradeParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *GFlagsUpgradeParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *GFlagsUpgradeParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *GFlagsUpgradeParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *GFlagsUpgradeParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *GFlagsUpgradeParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *GFlagsUpgradeParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetImportedState

`func (o *GFlagsUpgradeParams) GetImportedState() string`

GetImportedState returns the ImportedState field if non-nil, zero value otherwise.

### GetImportedStateOk

`func (o *GFlagsUpgradeParams) GetImportedStateOk() (*string, bool)`

GetImportedStateOk returns a tuple with the ImportedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedState

`func (o *GFlagsUpgradeParams) SetImportedState(v string)`

SetImportedState sets ImportedState field to given value.

### HasImportedState

`func (o *GFlagsUpgradeParams) HasImportedState() bool`

HasImportedState returns a boolean if a field has been set.

### GetInstallNodeAgent

`func (o *GFlagsUpgradeParams) GetInstallNodeAgent() bool`

GetInstallNodeAgent returns the InstallNodeAgent field if non-nil, zero value otherwise.

### GetInstallNodeAgentOk

`func (o *GFlagsUpgradeParams) GetInstallNodeAgentOk() (*bool, bool)`

GetInstallNodeAgentOk returns a tuple with the InstallNodeAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallNodeAgent

`func (o *GFlagsUpgradeParams) SetInstallNodeAgent(v bool)`

SetInstallNodeAgent sets InstallNodeAgent field to given value.

### HasInstallNodeAgent

`func (o *GFlagsUpgradeParams) HasInstallNodeAgent() bool`

HasInstallNodeAgent returns a boolean if a field has been set.

### GetInstallYbc

`func (o *GFlagsUpgradeParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *GFlagsUpgradeParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *GFlagsUpgradeParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *GFlagsUpgradeParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetIsKubernetesOperatorControlled

`func (o *GFlagsUpgradeParams) GetIsKubernetesOperatorControlled() bool`

GetIsKubernetesOperatorControlled returns the IsKubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetIsKubernetesOperatorControlledOk

`func (o *GFlagsUpgradeParams) GetIsKubernetesOperatorControlledOk() (*bool, bool)`

GetIsKubernetesOperatorControlledOk returns a tuple with the IsKubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetesOperatorControlled

`func (o *GFlagsUpgradeParams) SetIsKubernetesOperatorControlled(v bool)`

SetIsKubernetesOperatorControlled sets IsKubernetesOperatorControlled field to given value.

### HasIsKubernetesOperatorControlled

`func (o *GFlagsUpgradeParams) HasIsKubernetesOperatorControlled() bool`

HasIsKubernetesOperatorControlled returns a boolean if a field has been set.

### GetIsSoftwareRollbackAllowed

`func (o *GFlagsUpgradeParams) GetIsSoftwareRollbackAllowed() bool`

GetIsSoftwareRollbackAllowed returns the IsSoftwareRollbackAllowed field if non-nil, zero value otherwise.

### GetIsSoftwareRollbackAllowedOk

`func (o *GFlagsUpgradeParams) GetIsSoftwareRollbackAllowedOk() (*bool, bool)`

GetIsSoftwareRollbackAllowedOk returns a tuple with the IsSoftwareRollbackAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftwareRollbackAllowed

`func (o *GFlagsUpgradeParams) SetIsSoftwareRollbackAllowed(v bool)`

SetIsSoftwareRollbackAllowed sets IsSoftwareRollbackAllowed field to given value.

### HasIsSoftwareRollbackAllowed

`func (o *GFlagsUpgradeParams) HasIsSoftwareRollbackAllowed() bool`

HasIsSoftwareRollbackAllowed returns a boolean if a field has been set.

### GetItestS3PackagePath

`func (o *GFlagsUpgradeParams) GetItestS3PackagePath() string`

GetItestS3PackagePath returns the ItestS3PackagePath field if non-nil, zero value otherwise.

### GetItestS3PackagePathOk

`func (o *GFlagsUpgradeParams) GetItestS3PackagePathOk() (*string, bool)`

GetItestS3PackagePathOk returns a tuple with the ItestS3PackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItestS3PackagePath

`func (o *GFlagsUpgradeParams) SetItestS3PackagePath(v string)`

SetItestS3PackagePath sets ItestS3PackagePath field to given value.

### HasItestS3PackagePath

`func (o *GFlagsUpgradeParams) HasItestS3PackagePath() bool`

HasItestS3PackagePath returns a boolean if a field has been set.

### GetKubernetesUpgradeSupported

`func (o *GFlagsUpgradeParams) GetKubernetesUpgradeSupported() bool`

GetKubernetesUpgradeSupported returns the KubernetesUpgradeSupported field if non-nil, zero value otherwise.

### GetKubernetesUpgradeSupportedOk

`func (o *GFlagsUpgradeParams) GetKubernetesUpgradeSupportedOk() (*bool, bool)`

GetKubernetesUpgradeSupportedOk returns a tuple with the KubernetesUpgradeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesUpgradeSupported

`func (o *GFlagsUpgradeParams) SetKubernetesUpgradeSupported(v bool)`

SetKubernetesUpgradeSupported sets KubernetesUpgradeSupported field to given value.


### GetMasterGFlags

`func (o *GFlagsUpgradeParams) GetMasterGFlags() map[string]string`

GetMasterGFlags returns the MasterGFlags field if non-nil, zero value otherwise.

### GetMasterGFlagsOk

`func (o *GFlagsUpgradeParams) GetMasterGFlagsOk() (*map[string]string, bool)`

GetMasterGFlagsOk returns a tuple with the MasterGFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterGFlags

`func (o *GFlagsUpgradeParams) SetMasterGFlags(v map[string]string)`

SetMasterGFlags sets MasterGFlags field to given value.


### GetMastersInDefaultRegion

`func (o *GFlagsUpgradeParams) GetMastersInDefaultRegion() bool`

GetMastersInDefaultRegion returns the MastersInDefaultRegion field if non-nil, zero value otherwise.

### GetMastersInDefaultRegionOk

`func (o *GFlagsUpgradeParams) GetMastersInDefaultRegionOk() (*bool, bool)`

GetMastersInDefaultRegionOk returns a tuple with the MastersInDefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastersInDefaultRegion

`func (o *GFlagsUpgradeParams) SetMastersInDefaultRegion(v bool)`

SetMastersInDefaultRegion sets MastersInDefaultRegion field to given value.

### HasMastersInDefaultRegion

`func (o *GFlagsUpgradeParams) HasMastersInDefaultRegion() bool`

HasMastersInDefaultRegion returns a boolean if a field has been set.

### GetNextClusterIndex

`func (o *GFlagsUpgradeParams) GetNextClusterIndex() int32`

GetNextClusterIndex returns the NextClusterIndex field if non-nil, zero value otherwise.

### GetNextClusterIndexOk

`func (o *GFlagsUpgradeParams) GetNextClusterIndexOk() (*int32, bool)`

GetNextClusterIndexOk returns a tuple with the NextClusterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextClusterIndex

`func (o *GFlagsUpgradeParams) SetNextClusterIndex(v int32)`

SetNextClusterIndex sets NextClusterIndex field to given value.

### HasNextClusterIndex

`func (o *GFlagsUpgradeParams) HasNextClusterIndex() bool`

HasNextClusterIndex returns a boolean if a field has been set.

### GetNodeAgentMissing

`func (o *GFlagsUpgradeParams) GetNodeAgentMissing() bool`

GetNodeAgentMissing returns the NodeAgentMissing field if non-nil, zero value otherwise.

### GetNodeAgentMissingOk

`func (o *GFlagsUpgradeParams) GetNodeAgentMissingOk() (*bool, bool)`

GetNodeAgentMissingOk returns a tuple with the NodeAgentMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAgentMissing

`func (o *GFlagsUpgradeParams) SetNodeAgentMissing(v bool)`

SetNodeAgentMissing sets NodeAgentMissing field to given value.

### HasNodeAgentMissing

`func (o *GFlagsUpgradeParams) HasNodeAgentMissing() bool`

HasNodeAgentMissing returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *GFlagsUpgradeParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *GFlagsUpgradeParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *GFlagsUpgradeParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *GFlagsUpgradeParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *GFlagsUpgradeParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *GFlagsUpgradeParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *GFlagsUpgradeParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *GFlagsUpgradeParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNodePrefix

`func (o *GFlagsUpgradeParams) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *GFlagsUpgradeParams) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *GFlagsUpgradeParams) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *GFlagsUpgradeParams) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetNodesResizeAvailable

`func (o *GFlagsUpgradeParams) GetNodesResizeAvailable() bool`

GetNodesResizeAvailable returns the NodesResizeAvailable field if non-nil, zero value otherwise.

### GetNodesResizeAvailableOk

`func (o *GFlagsUpgradeParams) GetNodesResizeAvailableOk() (*bool, bool)`

GetNodesResizeAvailableOk returns a tuple with the NodesResizeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesResizeAvailable

`func (o *GFlagsUpgradeParams) SetNodesResizeAvailable(v bool)`

SetNodesResizeAvailable sets NodesResizeAvailable field to given value.

### HasNodesResizeAvailable

`func (o *GFlagsUpgradeParams) HasNodesResizeAvailable() bool`

HasNodesResizeAvailable returns a boolean if a field has been set.

### GetOtelCollectorEnabled

`func (o *GFlagsUpgradeParams) GetOtelCollectorEnabled() bool`

GetOtelCollectorEnabled returns the OtelCollectorEnabled field if non-nil, zero value otherwise.

### GetOtelCollectorEnabledOk

`func (o *GFlagsUpgradeParams) GetOtelCollectorEnabledOk() (*bool, bool)`

GetOtelCollectorEnabledOk returns a tuple with the OtelCollectorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelCollectorEnabled

`func (o *GFlagsUpgradeParams) SetOtelCollectorEnabled(v bool)`

SetOtelCollectorEnabled sets OtelCollectorEnabled field to given value.

### HasOtelCollectorEnabled

`func (o *GFlagsUpgradeParams) HasOtelCollectorEnabled() bool`

HasOtelCollectorEnabled returns a boolean if a field has been set.

### GetPlacementModificationTaskUuid

`func (o *GFlagsUpgradeParams) GetPlacementModificationTaskUuid() string`

GetPlacementModificationTaskUuid returns the PlacementModificationTaskUuid field if non-nil, zero value otherwise.

### GetPlacementModificationTaskUuidOk

`func (o *GFlagsUpgradeParams) GetPlacementModificationTaskUuidOk() (*string, bool)`

GetPlacementModificationTaskUuidOk returns a tuple with the PlacementModificationTaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementModificationTaskUuid

`func (o *GFlagsUpgradeParams) SetPlacementModificationTaskUuid(v string)`

SetPlacementModificationTaskUuid sets PlacementModificationTaskUuid field to given value.

### HasPlacementModificationTaskUuid

`func (o *GFlagsUpgradeParams) HasPlacementModificationTaskUuid() bool`

HasPlacementModificationTaskUuid returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *GFlagsUpgradeParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *GFlagsUpgradeParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *GFlagsUpgradeParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPlatformVersion

`func (o *GFlagsUpgradeParams) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *GFlagsUpgradeParams) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *GFlagsUpgradeParams) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *GFlagsUpgradeParams) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetPrevYBSoftwareConfig

`func (o *GFlagsUpgradeParams) GetPrevYBSoftwareConfig() PrevYBSoftwareConfig`

GetPrevYBSoftwareConfig returns the PrevYBSoftwareConfig field if non-nil, zero value otherwise.

### GetPrevYBSoftwareConfigOk

`func (o *GFlagsUpgradeParams) GetPrevYBSoftwareConfigOk() (*PrevYBSoftwareConfig, bool)`

GetPrevYBSoftwareConfigOk returns a tuple with the PrevYBSoftwareConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevYBSoftwareConfig

`func (o *GFlagsUpgradeParams) SetPrevYBSoftwareConfig(v PrevYBSoftwareConfig)`

SetPrevYBSoftwareConfig sets PrevYBSoftwareConfig field to given value.

### HasPrevYBSoftwareConfig

`func (o *GFlagsUpgradeParams) HasPrevYBSoftwareConfig() bool`

HasPrevYBSoftwareConfig returns a boolean if a field has been set.

### GetPreviousTaskUUID

`func (o *GFlagsUpgradeParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *GFlagsUpgradeParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *GFlagsUpgradeParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *GFlagsUpgradeParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRemotePackagePath

`func (o *GFlagsUpgradeParams) GetRemotePackagePath() string`

GetRemotePackagePath returns the RemotePackagePath field if non-nil, zero value otherwise.

### GetRemotePackagePathOk

`func (o *GFlagsUpgradeParams) GetRemotePackagePathOk() (*string, bool)`

GetRemotePackagePathOk returns a tuple with the RemotePackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePackagePath

`func (o *GFlagsUpgradeParams) SetRemotePackagePath(v string)`

SetRemotePackagePath sets RemotePackagePath field to given value.

### HasRemotePackagePath

`func (o *GFlagsUpgradeParams) HasRemotePackagePath() bool`

HasRemotePackagePath returns a boolean if a field has been set.

### GetResetAZConfig

`func (o *GFlagsUpgradeParams) GetResetAZConfig() bool`

GetResetAZConfig returns the ResetAZConfig field if non-nil, zero value otherwise.

### GetResetAZConfigOk

`func (o *GFlagsUpgradeParams) GetResetAZConfigOk() (*bool, bool)`

GetResetAZConfigOk returns a tuple with the ResetAZConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAZConfig

`func (o *GFlagsUpgradeParams) SetResetAZConfig(v bool)`

SetResetAZConfig sets ResetAZConfig field to given value.

### HasResetAZConfig

`func (o *GFlagsUpgradeParams) HasResetAZConfig() bool`

HasResetAZConfig returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *GFlagsUpgradeParams) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *GFlagsUpgradeParams) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *GFlagsUpgradeParams) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *GFlagsUpgradeParams) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetRootAndClientRootCASame

`func (o *GFlagsUpgradeParams) GetRootAndClientRootCASame() bool`

GetRootAndClientRootCASame returns the RootAndClientRootCASame field if non-nil, zero value otherwise.

### GetRootAndClientRootCASameOk

`func (o *GFlagsUpgradeParams) GetRootAndClientRootCASameOk() (*bool, bool)`

GetRootAndClientRootCASameOk returns a tuple with the RootAndClientRootCASame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootAndClientRootCASame

`func (o *GFlagsUpgradeParams) SetRootAndClientRootCASame(v bool)`

SetRootAndClientRootCASame sets RootAndClientRootCASame field to given value.

### HasRootAndClientRootCASame

`func (o *GFlagsUpgradeParams) HasRootAndClientRootCASame() bool`

HasRootAndClientRootCASame returns a boolean if a field has been set.

### GetRootCA

`func (o *GFlagsUpgradeParams) GetRootCA() string`

GetRootCA returns the RootCA field if non-nil, zero value otherwise.

### GetRootCAOk

`func (o *GFlagsUpgradeParams) GetRootCAOk() (*string, bool)`

GetRootCAOk returns a tuple with the RootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCA

`func (o *GFlagsUpgradeParams) SetRootCA(v string)`

SetRootCA sets RootCA field to given value.

### HasRootCA

`func (o *GFlagsUpgradeParams) HasRootCA() bool`

HasRootCA returns a boolean if a field has been set.

### GetRunOnlyPrechecks

`func (o *GFlagsUpgradeParams) GetRunOnlyPrechecks() bool`

GetRunOnlyPrechecks returns the RunOnlyPrechecks field if non-nil, zero value otherwise.

### GetRunOnlyPrechecksOk

`func (o *GFlagsUpgradeParams) GetRunOnlyPrechecksOk() (*bool, bool)`

GetRunOnlyPrechecksOk returns a tuple with the RunOnlyPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnlyPrechecks

`func (o *GFlagsUpgradeParams) SetRunOnlyPrechecks(v bool)`

SetRunOnlyPrechecks sets RunOnlyPrechecks field to given value.

### HasRunOnlyPrechecks

`func (o *GFlagsUpgradeParams) HasRunOnlyPrechecks() bool`

HasRunOnlyPrechecks returns a boolean if a field has been set.

### GetSetTxnTableWaitCountFlag

`func (o *GFlagsUpgradeParams) GetSetTxnTableWaitCountFlag() bool`

GetSetTxnTableWaitCountFlag returns the SetTxnTableWaitCountFlag field if non-nil, zero value otherwise.

### GetSetTxnTableWaitCountFlagOk

`func (o *GFlagsUpgradeParams) GetSetTxnTableWaitCountFlagOk() (*bool, bool)`

GetSetTxnTableWaitCountFlagOk returns a tuple with the SetTxnTableWaitCountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTxnTableWaitCountFlag

`func (o *GFlagsUpgradeParams) SetSetTxnTableWaitCountFlag(v bool)`

SetSetTxnTableWaitCountFlag sets SetTxnTableWaitCountFlag field to given value.

### HasSetTxnTableWaitCountFlag

`func (o *GFlagsUpgradeParams) HasSetTxnTableWaitCountFlag() bool`

HasSetTxnTableWaitCountFlag returns a boolean if a field has been set.

### GetSkipNodeChecks

`func (o *GFlagsUpgradeParams) GetSkipNodeChecks() bool`

GetSkipNodeChecks returns the SkipNodeChecks field if non-nil, zero value otherwise.

### GetSkipNodeChecksOk

`func (o *GFlagsUpgradeParams) GetSkipNodeChecksOk() (*bool, bool)`

GetSkipNodeChecksOk returns a tuple with the SkipNodeChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNodeChecks

`func (o *GFlagsUpgradeParams) SetSkipNodeChecks(v bool)`

SetSkipNodeChecks sets SkipNodeChecks field to given value.

### HasSkipNodeChecks

`func (o *GFlagsUpgradeParams) HasSkipNodeChecks() bool`

HasSkipNodeChecks returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *GFlagsUpgradeParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *GFlagsUpgradeParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *GFlagsUpgradeParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *GFlagsUpgradeParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *GFlagsUpgradeParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *GFlagsUpgradeParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSoftwareUpgradeState

`func (o *GFlagsUpgradeParams) GetSoftwareUpgradeState() string`

GetSoftwareUpgradeState returns the SoftwareUpgradeState field if non-nil, zero value otherwise.

### GetSoftwareUpgradeStateOk

`func (o *GFlagsUpgradeParams) GetSoftwareUpgradeStateOk() (*string, bool)`

GetSoftwareUpgradeStateOk returns a tuple with the SoftwareUpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpgradeState

`func (o *GFlagsUpgradeParams) SetSoftwareUpgradeState(v string)`

SetSoftwareUpgradeState sets SoftwareUpgradeState field to given value.

### HasSoftwareUpgradeState

`func (o *GFlagsUpgradeParams) HasSoftwareUpgradeState() bool`

HasSoftwareUpgradeState returns a boolean if a field has been set.

### GetSourceXClusterConfigs

`func (o *GFlagsUpgradeParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *GFlagsUpgradeParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *GFlagsUpgradeParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *GFlagsUpgradeParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSshUserOverride

`func (o *GFlagsUpgradeParams) GetSshUserOverride() string`

GetSshUserOverride returns the SshUserOverride field if non-nil, zero value otherwise.

### GetSshUserOverrideOk

`func (o *GFlagsUpgradeParams) GetSshUserOverrideOk() (*string, bool)`

GetSshUserOverrideOk returns a tuple with the SshUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverride

`func (o *GFlagsUpgradeParams) SetSshUserOverride(v string)`

SetSshUserOverride sets SshUserOverride field to given value.

### HasSshUserOverride

`func (o *GFlagsUpgradeParams) HasSshUserOverride() bool`

HasSshUserOverride returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *GFlagsUpgradeParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *GFlagsUpgradeParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *GFlagsUpgradeParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *GFlagsUpgradeParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetTserverGFlags

`func (o *GFlagsUpgradeParams) GetTserverGFlags() map[string]string`

GetTserverGFlags returns the TserverGFlags field if non-nil, zero value otherwise.

### GetTserverGFlagsOk

`func (o *GFlagsUpgradeParams) GetTserverGFlagsOk() (*map[string]string, bool)`

GetTserverGFlagsOk returns a tuple with the TserverGFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverGFlags

`func (o *GFlagsUpgradeParams) SetTserverGFlags(v map[string]string)`

SetTserverGFlags sets TserverGFlags field to given value.


### GetUniversePaused

`func (o *GFlagsUpgradeParams) GetUniversePaused() bool`

GetUniversePaused returns the UniversePaused field if non-nil, zero value otherwise.

### GetUniversePausedOk

`func (o *GFlagsUpgradeParams) GetUniversePausedOk() (*bool, bool)`

GetUniversePausedOk returns a tuple with the UniversePaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversePaused

`func (o *GFlagsUpgradeParams) SetUniversePaused(v bool)`

SetUniversePaused sets UniversePaused field to given value.

### HasUniversePaused

`func (o *GFlagsUpgradeParams) HasUniversePaused() bool`

HasUniversePaused returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *GFlagsUpgradeParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *GFlagsUpgradeParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *GFlagsUpgradeParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *GFlagsUpgradeParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUpdateInProgress

`func (o *GFlagsUpgradeParams) GetUpdateInProgress() bool`

GetUpdateInProgress returns the UpdateInProgress field if non-nil, zero value otherwise.

### GetUpdateInProgressOk

`func (o *GFlagsUpgradeParams) GetUpdateInProgressOk() (*bool, bool)`

GetUpdateInProgressOk returns a tuple with the UpdateInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInProgress

`func (o *GFlagsUpgradeParams) SetUpdateInProgress(v bool)`

SetUpdateInProgress sets UpdateInProgress field to given value.

### HasUpdateInProgress

`func (o *GFlagsUpgradeParams) HasUpdateInProgress() bool`

HasUpdateInProgress returns a boolean if a field has been set.

### GetUpdateOptions

`func (o *GFlagsUpgradeParams) GetUpdateOptions() []string`

GetUpdateOptions returns the UpdateOptions field if non-nil, zero value otherwise.

### GetUpdateOptionsOk

`func (o *GFlagsUpgradeParams) GetUpdateOptionsOk() (*[]string, bool)`

GetUpdateOptionsOk returns a tuple with the UpdateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOptions

`func (o *GFlagsUpgradeParams) SetUpdateOptions(v []string)`

SetUpdateOptions sets UpdateOptions field to given value.

### HasUpdateOptions

`func (o *GFlagsUpgradeParams) HasUpdateOptions() bool`

HasUpdateOptions returns a boolean if a field has been set.

### GetUpdateSucceeded

`func (o *GFlagsUpgradeParams) GetUpdateSucceeded() bool`

GetUpdateSucceeded returns the UpdateSucceeded field if non-nil, zero value otherwise.

### GetUpdateSucceededOk

`func (o *GFlagsUpgradeParams) GetUpdateSucceededOk() (*bool, bool)`

GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceeded

`func (o *GFlagsUpgradeParams) SetUpdateSucceeded(v bool)`

SetUpdateSucceeded sets UpdateSucceeded field to given value.

### HasUpdateSucceeded

`func (o *GFlagsUpgradeParams) HasUpdateSucceeded() bool`

HasUpdateSucceeded returns a boolean if a field has been set.

### GetUpdatingTask

`func (o *GFlagsUpgradeParams) GetUpdatingTask() string`

GetUpdatingTask returns the UpdatingTask field if non-nil, zero value otherwise.

### GetUpdatingTaskOk

`func (o *GFlagsUpgradeParams) GetUpdatingTaskOk() (*string, bool)`

GetUpdatingTaskOk returns a tuple with the UpdatingTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTask

`func (o *GFlagsUpgradeParams) SetUpdatingTask(v string)`

SetUpdatingTask sets UpdatingTask field to given value.

### HasUpdatingTask

`func (o *GFlagsUpgradeParams) HasUpdatingTask() bool`

HasUpdatingTask returns a boolean if a field has been set.

### GetUpdatingTaskUUID

`func (o *GFlagsUpgradeParams) GetUpdatingTaskUUID() string`

GetUpdatingTaskUUID returns the UpdatingTaskUUID field if non-nil, zero value otherwise.

### GetUpdatingTaskUUIDOk

`func (o *GFlagsUpgradeParams) GetUpdatingTaskUUIDOk() (*string, bool)`

GetUpdatingTaskUUIDOk returns a tuple with the UpdatingTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTaskUUID

`func (o *GFlagsUpgradeParams) SetUpdatingTaskUUID(v string)`

SetUpdatingTaskUUID sets UpdatingTaskUUID field to given value.

### HasUpdatingTaskUUID

`func (o *GFlagsUpgradeParams) HasUpdatingTaskUUID() bool`

HasUpdatingTaskUUID returns a boolean if a field has been set.

### GetUpgradeOption

`func (o *GFlagsUpgradeParams) GetUpgradeOption() string`

GetUpgradeOption returns the UpgradeOption field if non-nil, zero value otherwise.

### GetUpgradeOptionOk

`func (o *GFlagsUpgradeParams) GetUpgradeOptionOk() (*string, bool)`

GetUpgradeOptionOk returns a tuple with the UpgradeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOption

`func (o *GFlagsUpgradeParams) SetUpgradeOption(v string)`

SetUpgradeOption sets UpgradeOption field to given value.


### GetUseNewHelmNamingStyle

`func (o *GFlagsUpgradeParams) GetUseNewHelmNamingStyle() bool`

GetUseNewHelmNamingStyle returns the UseNewHelmNamingStyle field if non-nil, zero value otherwise.

### GetUseNewHelmNamingStyleOk

`func (o *GFlagsUpgradeParams) GetUseNewHelmNamingStyleOk() (*bool, bool)`

GetUseNewHelmNamingStyleOk returns a tuple with the UseNewHelmNamingStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNewHelmNamingStyle

`func (o *GFlagsUpgradeParams) SetUseNewHelmNamingStyle(v bool)`

SetUseNewHelmNamingStyle sets UseNewHelmNamingStyle field to given value.

### HasUseNewHelmNamingStyle

`func (o *GFlagsUpgradeParams) HasUseNewHelmNamingStyle() bool`

HasUseNewHelmNamingStyle returns a boolean if a field has been set.

### GetUserAZSelected

`func (o *GFlagsUpgradeParams) GetUserAZSelected() bool`

GetUserAZSelected returns the UserAZSelected field if non-nil, zero value otherwise.

### GetUserAZSelectedOk

`func (o *GFlagsUpgradeParams) GetUserAZSelectedOk() (*bool, bool)`

GetUserAZSelectedOk returns a tuple with the UserAZSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAZSelected

`func (o *GFlagsUpgradeParams) SetUserAZSelected(v bool)`

SetUserAZSelected sets UserAZSelected field to given value.

### HasUserAZSelected

`func (o *GFlagsUpgradeParams) HasUserAZSelected() bool`

HasUserAZSelected returns a boolean if a field has been set.

### GetXclusterInfo

`func (o *GFlagsUpgradeParams) GetXclusterInfo() XClusterInfo`

GetXclusterInfo returns the XclusterInfo field if non-nil, zero value otherwise.

### GetXclusterInfoOk

`func (o *GFlagsUpgradeParams) GetXclusterInfoOk() (*XClusterInfo, bool)`

GetXclusterInfoOk returns a tuple with the XclusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterInfo

`func (o *GFlagsUpgradeParams) SetXclusterInfo(v XClusterInfo)`

SetXclusterInfo sets XclusterInfo field to given value.

### HasXclusterInfo

`func (o *GFlagsUpgradeParams) HasXclusterInfo() bool`

HasXclusterInfo returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *GFlagsUpgradeParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *GFlagsUpgradeParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *GFlagsUpgradeParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *GFlagsUpgradeParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *GFlagsUpgradeParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *GFlagsUpgradeParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *GFlagsUpgradeParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *GFlagsUpgradeParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *GFlagsUpgradeParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *GFlagsUpgradeParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *GFlagsUpgradeParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *GFlagsUpgradeParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


