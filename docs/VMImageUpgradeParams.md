# VMImageUpgradeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**ForceVMImageUpgrade** | **bool** |  | 
**ImageBundleUUID** | Pointer to **string** | ImageBundle to be used for upgrade. &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.21.1.0.&lt;/b&gt; Use imageBundles instead. | [optional] 
**ImageBundles** | Pointer to [**[]ImageBundleUpgradeInfo**](ImageBundleUpgradeInfo.md) | Available since YBA version 2.21.1.0. ImageBundles for provider to be used for upgrade | [optional] 
**ImportedState** | Pointer to **string** |  | [optional] 
**InstallNodeAgent** | Pointer to **bool** | YbaApi Internal. Install node agent in background if it is true | [optional] 
**InstallYbc** | Pointer to **bool** |  | [optional] 
**IsKubernetesOperatorControlled** | Pointer to **bool** |  | [optional] 
**IsSoftwareRollbackAllowed** | Pointer to **bool** | Available since YBA version 2.20.2.0 | [optional] [readonly] 
**ItestS3PackagePath** | Pointer to **string** |  | [optional] 
**KubernetesUpgradeSupported** | **bool** |  | 
**MachineImages** | Pointer to **map[string]string** | Map of region UUID to AMI name. &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.18.0.0.&lt;/b&gt; Use imageBundle instead. | [optional] 
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
**SshUserOverrideMap** | Pointer to **map[string]string** | Map of region UUID to SSH User override. &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.18.0.0.&lt;/b&gt; Use imageBundle instead. | [optional] 
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
**YbSoftwareVersion** | **string** |  | 
**YbcInstalled** | Pointer to **bool** |  | [optional] 
**YbcSoftwareVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewVMImageUpgradeParams

`func NewVMImageUpgradeParams(clusters []Cluster, creatingUser Users, forceVMImageUpgrade bool, kubernetesUpgradeSupported bool, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, upgradeOption string, ybSoftwareVersion string, ) *VMImageUpgradeParams`

NewVMImageUpgradeParams instantiates a new VMImageUpgradeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMImageUpgradeParamsWithDefaults

`func NewVMImageUpgradeParamsWithDefaults() *VMImageUpgradeParams`

NewVMImageUpgradeParamsWithDefaults instantiates a new VMImageUpgradeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *VMImageUpgradeParams) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *VMImageUpgradeParams) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *VMImageUpgradeParams) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *VMImageUpgradeParams) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetArch

`func (o *VMImageUpgradeParams) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *VMImageUpgradeParams) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *VMImageUpgradeParams) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *VMImageUpgradeParams) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCapability

`func (o *VMImageUpgradeParams) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *VMImageUpgradeParams) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *VMImageUpgradeParams) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *VMImageUpgradeParams) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetClientRootCA

`func (o *VMImageUpgradeParams) GetClientRootCA() string`

GetClientRootCA returns the ClientRootCA field if non-nil, zero value otherwise.

### GetClientRootCAOk

`func (o *VMImageUpgradeParams) GetClientRootCAOk() (*string, bool)`

GetClientRootCAOk returns a tuple with the ClientRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCA

`func (o *VMImageUpgradeParams) SetClientRootCA(v string)`

SetClientRootCA sets ClientRootCA field to given value.

### HasClientRootCA

`func (o *VMImageUpgradeParams) HasClientRootCA() bool`

HasClientRootCA returns a boolean if a field has been set.

### GetClusters

`func (o *VMImageUpgradeParams) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *VMImageUpgradeParams) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *VMImageUpgradeParams) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.


### GetCmkArn

`func (o *VMImageUpgradeParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *VMImageUpgradeParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *VMImageUpgradeParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *VMImageUpgradeParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *VMImageUpgradeParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *VMImageUpgradeParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *VMImageUpgradeParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *VMImageUpgradeParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *VMImageUpgradeParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *VMImageUpgradeParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *VMImageUpgradeParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCurrentClusterType

`func (o *VMImageUpgradeParams) GetCurrentClusterType() string`

GetCurrentClusterType returns the CurrentClusterType field if non-nil, zero value otherwise.

### GetCurrentClusterTypeOk

`func (o *VMImageUpgradeParams) GetCurrentClusterTypeOk() (*string, bool)`

GetCurrentClusterTypeOk returns a tuple with the CurrentClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClusterType

`func (o *VMImageUpgradeParams) SetCurrentClusterType(v string)`

SetCurrentClusterType sets CurrentClusterType field to given value.

### HasCurrentClusterType

`func (o *VMImageUpgradeParams) HasCurrentClusterType() bool`

HasCurrentClusterType returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *VMImageUpgradeParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *VMImageUpgradeParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *VMImageUpgradeParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *VMImageUpgradeParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableYbc

`func (o *VMImageUpgradeParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *VMImageUpgradeParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *VMImageUpgradeParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *VMImageUpgradeParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *VMImageUpgradeParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *VMImageUpgradeParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *VMImageUpgradeParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *VMImageUpgradeParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *VMImageUpgradeParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *VMImageUpgradeParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *VMImageUpgradeParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *VMImageUpgradeParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *VMImageUpgradeParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *VMImageUpgradeParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *VMImageUpgradeParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *VMImageUpgradeParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *VMImageUpgradeParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *VMImageUpgradeParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *VMImageUpgradeParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *VMImageUpgradeParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetForceVMImageUpgrade

`func (o *VMImageUpgradeParams) GetForceVMImageUpgrade() bool`

GetForceVMImageUpgrade returns the ForceVMImageUpgrade field if non-nil, zero value otherwise.

### GetForceVMImageUpgradeOk

`func (o *VMImageUpgradeParams) GetForceVMImageUpgradeOk() (*bool, bool)`

GetForceVMImageUpgradeOk returns a tuple with the ForceVMImageUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceVMImageUpgrade

`func (o *VMImageUpgradeParams) SetForceVMImageUpgrade(v bool)`

SetForceVMImageUpgrade sets ForceVMImageUpgrade field to given value.


### GetImageBundleUUID

`func (o *VMImageUpgradeParams) GetImageBundleUUID() string`

GetImageBundleUUID returns the ImageBundleUUID field if non-nil, zero value otherwise.

### GetImageBundleUUIDOk

`func (o *VMImageUpgradeParams) GetImageBundleUUIDOk() (*string, bool)`

GetImageBundleUUIDOk returns a tuple with the ImageBundleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundleUUID

`func (o *VMImageUpgradeParams) SetImageBundleUUID(v string)`

SetImageBundleUUID sets ImageBundleUUID field to given value.

### HasImageBundleUUID

`func (o *VMImageUpgradeParams) HasImageBundleUUID() bool`

HasImageBundleUUID returns a boolean if a field has been set.

### GetImageBundles

`func (o *VMImageUpgradeParams) GetImageBundles() []ImageBundleUpgradeInfo`

GetImageBundles returns the ImageBundles field if non-nil, zero value otherwise.

### GetImageBundlesOk

`func (o *VMImageUpgradeParams) GetImageBundlesOk() (*[]ImageBundleUpgradeInfo, bool)`

GetImageBundlesOk returns a tuple with the ImageBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundles

`func (o *VMImageUpgradeParams) SetImageBundles(v []ImageBundleUpgradeInfo)`

SetImageBundles sets ImageBundles field to given value.

### HasImageBundles

`func (o *VMImageUpgradeParams) HasImageBundles() bool`

HasImageBundles returns a boolean if a field has been set.

### GetImportedState

`func (o *VMImageUpgradeParams) GetImportedState() string`

GetImportedState returns the ImportedState field if non-nil, zero value otherwise.

### GetImportedStateOk

`func (o *VMImageUpgradeParams) GetImportedStateOk() (*string, bool)`

GetImportedStateOk returns a tuple with the ImportedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedState

`func (o *VMImageUpgradeParams) SetImportedState(v string)`

SetImportedState sets ImportedState field to given value.

### HasImportedState

`func (o *VMImageUpgradeParams) HasImportedState() bool`

HasImportedState returns a boolean if a field has been set.

### GetInstallNodeAgent

`func (o *VMImageUpgradeParams) GetInstallNodeAgent() bool`

GetInstallNodeAgent returns the InstallNodeAgent field if non-nil, zero value otherwise.

### GetInstallNodeAgentOk

`func (o *VMImageUpgradeParams) GetInstallNodeAgentOk() (*bool, bool)`

GetInstallNodeAgentOk returns a tuple with the InstallNodeAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallNodeAgent

`func (o *VMImageUpgradeParams) SetInstallNodeAgent(v bool)`

SetInstallNodeAgent sets InstallNodeAgent field to given value.

### HasInstallNodeAgent

`func (o *VMImageUpgradeParams) HasInstallNodeAgent() bool`

HasInstallNodeAgent returns a boolean if a field has been set.

### GetInstallYbc

`func (o *VMImageUpgradeParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *VMImageUpgradeParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *VMImageUpgradeParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *VMImageUpgradeParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetIsKubernetesOperatorControlled

`func (o *VMImageUpgradeParams) GetIsKubernetesOperatorControlled() bool`

GetIsKubernetesOperatorControlled returns the IsKubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetIsKubernetesOperatorControlledOk

`func (o *VMImageUpgradeParams) GetIsKubernetesOperatorControlledOk() (*bool, bool)`

GetIsKubernetesOperatorControlledOk returns a tuple with the IsKubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetesOperatorControlled

`func (o *VMImageUpgradeParams) SetIsKubernetesOperatorControlled(v bool)`

SetIsKubernetesOperatorControlled sets IsKubernetesOperatorControlled field to given value.

### HasIsKubernetesOperatorControlled

`func (o *VMImageUpgradeParams) HasIsKubernetesOperatorControlled() bool`

HasIsKubernetesOperatorControlled returns a boolean if a field has been set.

### GetIsSoftwareRollbackAllowed

`func (o *VMImageUpgradeParams) GetIsSoftwareRollbackAllowed() bool`

GetIsSoftwareRollbackAllowed returns the IsSoftwareRollbackAllowed field if non-nil, zero value otherwise.

### GetIsSoftwareRollbackAllowedOk

`func (o *VMImageUpgradeParams) GetIsSoftwareRollbackAllowedOk() (*bool, bool)`

GetIsSoftwareRollbackAllowedOk returns a tuple with the IsSoftwareRollbackAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftwareRollbackAllowed

`func (o *VMImageUpgradeParams) SetIsSoftwareRollbackAllowed(v bool)`

SetIsSoftwareRollbackAllowed sets IsSoftwareRollbackAllowed field to given value.

### HasIsSoftwareRollbackAllowed

`func (o *VMImageUpgradeParams) HasIsSoftwareRollbackAllowed() bool`

HasIsSoftwareRollbackAllowed returns a boolean if a field has been set.

### GetItestS3PackagePath

`func (o *VMImageUpgradeParams) GetItestS3PackagePath() string`

GetItestS3PackagePath returns the ItestS3PackagePath field if non-nil, zero value otherwise.

### GetItestS3PackagePathOk

`func (o *VMImageUpgradeParams) GetItestS3PackagePathOk() (*string, bool)`

GetItestS3PackagePathOk returns a tuple with the ItestS3PackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItestS3PackagePath

`func (o *VMImageUpgradeParams) SetItestS3PackagePath(v string)`

SetItestS3PackagePath sets ItestS3PackagePath field to given value.

### HasItestS3PackagePath

`func (o *VMImageUpgradeParams) HasItestS3PackagePath() bool`

HasItestS3PackagePath returns a boolean if a field has been set.

### GetKubernetesUpgradeSupported

`func (o *VMImageUpgradeParams) GetKubernetesUpgradeSupported() bool`

GetKubernetesUpgradeSupported returns the KubernetesUpgradeSupported field if non-nil, zero value otherwise.

### GetKubernetesUpgradeSupportedOk

`func (o *VMImageUpgradeParams) GetKubernetesUpgradeSupportedOk() (*bool, bool)`

GetKubernetesUpgradeSupportedOk returns a tuple with the KubernetesUpgradeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesUpgradeSupported

`func (o *VMImageUpgradeParams) SetKubernetesUpgradeSupported(v bool)`

SetKubernetesUpgradeSupported sets KubernetesUpgradeSupported field to given value.


### GetMachineImages

`func (o *VMImageUpgradeParams) GetMachineImages() map[string]string`

GetMachineImages returns the MachineImages field if non-nil, zero value otherwise.

### GetMachineImagesOk

`func (o *VMImageUpgradeParams) GetMachineImagesOk() (*map[string]string, bool)`

GetMachineImagesOk returns a tuple with the MachineImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineImages

`func (o *VMImageUpgradeParams) SetMachineImages(v map[string]string)`

SetMachineImages sets MachineImages field to given value.

### HasMachineImages

`func (o *VMImageUpgradeParams) HasMachineImages() bool`

HasMachineImages returns a boolean if a field has been set.

### GetMastersInDefaultRegion

`func (o *VMImageUpgradeParams) GetMastersInDefaultRegion() bool`

GetMastersInDefaultRegion returns the MastersInDefaultRegion field if non-nil, zero value otherwise.

### GetMastersInDefaultRegionOk

`func (o *VMImageUpgradeParams) GetMastersInDefaultRegionOk() (*bool, bool)`

GetMastersInDefaultRegionOk returns a tuple with the MastersInDefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastersInDefaultRegion

`func (o *VMImageUpgradeParams) SetMastersInDefaultRegion(v bool)`

SetMastersInDefaultRegion sets MastersInDefaultRegion field to given value.

### HasMastersInDefaultRegion

`func (o *VMImageUpgradeParams) HasMastersInDefaultRegion() bool`

HasMastersInDefaultRegion returns a boolean if a field has been set.

### GetNextClusterIndex

`func (o *VMImageUpgradeParams) GetNextClusterIndex() int32`

GetNextClusterIndex returns the NextClusterIndex field if non-nil, zero value otherwise.

### GetNextClusterIndexOk

`func (o *VMImageUpgradeParams) GetNextClusterIndexOk() (*int32, bool)`

GetNextClusterIndexOk returns a tuple with the NextClusterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextClusterIndex

`func (o *VMImageUpgradeParams) SetNextClusterIndex(v int32)`

SetNextClusterIndex sets NextClusterIndex field to given value.

### HasNextClusterIndex

`func (o *VMImageUpgradeParams) HasNextClusterIndex() bool`

HasNextClusterIndex returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *VMImageUpgradeParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *VMImageUpgradeParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *VMImageUpgradeParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *VMImageUpgradeParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *VMImageUpgradeParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *VMImageUpgradeParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *VMImageUpgradeParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *VMImageUpgradeParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNodePrefix

`func (o *VMImageUpgradeParams) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *VMImageUpgradeParams) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *VMImageUpgradeParams) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *VMImageUpgradeParams) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetNodesResizeAvailable

`func (o *VMImageUpgradeParams) GetNodesResizeAvailable() bool`

GetNodesResizeAvailable returns the NodesResizeAvailable field if non-nil, zero value otherwise.

### GetNodesResizeAvailableOk

`func (o *VMImageUpgradeParams) GetNodesResizeAvailableOk() (*bool, bool)`

GetNodesResizeAvailableOk returns a tuple with the NodesResizeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesResizeAvailable

`func (o *VMImageUpgradeParams) SetNodesResizeAvailable(v bool)`

SetNodesResizeAvailable sets NodesResizeAvailable field to given value.

### HasNodesResizeAvailable

`func (o *VMImageUpgradeParams) HasNodesResizeAvailable() bool`

HasNodesResizeAvailable returns a boolean if a field has been set.

### GetOtelCollectorEnabled

`func (o *VMImageUpgradeParams) GetOtelCollectorEnabled() bool`

GetOtelCollectorEnabled returns the OtelCollectorEnabled field if non-nil, zero value otherwise.

### GetOtelCollectorEnabledOk

`func (o *VMImageUpgradeParams) GetOtelCollectorEnabledOk() (*bool, bool)`

GetOtelCollectorEnabledOk returns a tuple with the OtelCollectorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelCollectorEnabled

`func (o *VMImageUpgradeParams) SetOtelCollectorEnabled(v bool)`

SetOtelCollectorEnabled sets OtelCollectorEnabled field to given value.

### HasOtelCollectorEnabled

`func (o *VMImageUpgradeParams) HasOtelCollectorEnabled() bool`

HasOtelCollectorEnabled returns a boolean if a field has been set.

### GetPlacementModificationTaskUuid

`func (o *VMImageUpgradeParams) GetPlacementModificationTaskUuid() string`

GetPlacementModificationTaskUuid returns the PlacementModificationTaskUuid field if non-nil, zero value otherwise.

### GetPlacementModificationTaskUuidOk

`func (o *VMImageUpgradeParams) GetPlacementModificationTaskUuidOk() (*string, bool)`

GetPlacementModificationTaskUuidOk returns a tuple with the PlacementModificationTaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementModificationTaskUuid

`func (o *VMImageUpgradeParams) SetPlacementModificationTaskUuid(v string)`

SetPlacementModificationTaskUuid sets PlacementModificationTaskUuid field to given value.

### HasPlacementModificationTaskUuid

`func (o *VMImageUpgradeParams) HasPlacementModificationTaskUuid() bool`

HasPlacementModificationTaskUuid returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *VMImageUpgradeParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *VMImageUpgradeParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *VMImageUpgradeParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPlatformVersion

`func (o *VMImageUpgradeParams) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *VMImageUpgradeParams) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *VMImageUpgradeParams) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.


### GetPrevYBSoftwareConfig

`func (o *VMImageUpgradeParams) GetPrevYBSoftwareConfig() PrevYBSoftwareConfig`

GetPrevYBSoftwareConfig returns the PrevYBSoftwareConfig field if non-nil, zero value otherwise.

### GetPrevYBSoftwareConfigOk

`func (o *VMImageUpgradeParams) GetPrevYBSoftwareConfigOk() (*PrevYBSoftwareConfig, bool)`

GetPrevYBSoftwareConfigOk returns a tuple with the PrevYBSoftwareConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevYBSoftwareConfig

`func (o *VMImageUpgradeParams) SetPrevYBSoftwareConfig(v PrevYBSoftwareConfig)`

SetPrevYBSoftwareConfig sets PrevYBSoftwareConfig field to given value.

### HasPrevYBSoftwareConfig

`func (o *VMImageUpgradeParams) HasPrevYBSoftwareConfig() bool`

HasPrevYBSoftwareConfig returns a boolean if a field has been set.

### GetPreviousTaskUUID

`func (o *VMImageUpgradeParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *VMImageUpgradeParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *VMImageUpgradeParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *VMImageUpgradeParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRemotePackagePath

`func (o *VMImageUpgradeParams) GetRemotePackagePath() string`

GetRemotePackagePath returns the RemotePackagePath field if non-nil, zero value otherwise.

### GetRemotePackagePathOk

`func (o *VMImageUpgradeParams) GetRemotePackagePathOk() (*string, bool)`

GetRemotePackagePathOk returns a tuple with the RemotePackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePackagePath

`func (o *VMImageUpgradeParams) SetRemotePackagePath(v string)`

SetRemotePackagePath sets RemotePackagePath field to given value.

### HasRemotePackagePath

`func (o *VMImageUpgradeParams) HasRemotePackagePath() bool`

HasRemotePackagePath returns a boolean if a field has been set.

### GetResetAZConfig

`func (o *VMImageUpgradeParams) GetResetAZConfig() bool`

GetResetAZConfig returns the ResetAZConfig field if non-nil, zero value otherwise.

### GetResetAZConfigOk

`func (o *VMImageUpgradeParams) GetResetAZConfigOk() (*bool, bool)`

GetResetAZConfigOk returns a tuple with the ResetAZConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAZConfig

`func (o *VMImageUpgradeParams) SetResetAZConfig(v bool)`

SetResetAZConfig sets ResetAZConfig field to given value.

### HasResetAZConfig

`func (o *VMImageUpgradeParams) HasResetAZConfig() bool`

HasResetAZConfig returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *VMImageUpgradeParams) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *VMImageUpgradeParams) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *VMImageUpgradeParams) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *VMImageUpgradeParams) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetRootAndClientRootCASame

`func (o *VMImageUpgradeParams) GetRootAndClientRootCASame() bool`

GetRootAndClientRootCASame returns the RootAndClientRootCASame field if non-nil, zero value otherwise.

### GetRootAndClientRootCASameOk

`func (o *VMImageUpgradeParams) GetRootAndClientRootCASameOk() (*bool, bool)`

GetRootAndClientRootCASameOk returns a tuple with the RootAndClientRootCASame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootAndClientRootCASame

`func (o *VMImageUpgradeParams) SetRootAndClientRootCASame(v bool)`

SetRootAndClientRootCASame sets RootAndClientRootCASame field to given value.

### HasRootAndClientRootCASame

`func (o *VMImageUpgradeParams) HasRootAndClientRootCASame() bool`

HasRootAndClientRootCASame returns a boolean if a field has been set.

### GetRootCA

`func (o *VMImageUpgradeParams) GetRootCA() string`

GetRootCA returns the RootCA field if non-nil, zero value otherwise.

### GetRootCAOk

`func (o *VMImageUpgradeParams) GetRootCAOk() (*string, bool)`

GetRootCAOk returns a tuple with the RootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCA

`func (o *VMImageUpgradeParams) SetRootCA(v string)`

SetRootCA sets RootCA field to given value.

### HasRootCA

`func (o *VMImageUpgradeParams) HasRootCA() bool`

HasRootCA returns a boolean if a field has been set.

### GetSetTxnTableWaitCountFlag

`func (o *VMImageUpgradeParams) GetSetTxnTableWaitCountFlag() bool`

GetSetTxnTableWaitCountFlag returns the SetTxnTableWaitCountFlag field if non-nil, zero value otherwise.

### GetSetTxnTableWaitCountFlagOk

`func (o *VMImageUpgradeParams) GetSetTxnTableWaitCountFlagOk() (*bool, bool)`

GetSetTxnTableWaitCountFlagOk returns a tuple with the SetTxnTableWaitCountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTxnTableWaitCountFlag

`func (o *VMImageUpgradeParams) SetSetTxnTableWaitCountFlag(v bool)`

SetSetTxnTableWaitCountFlag sets SetTxnTableWaitCountFlag field to given value.

### HasSetTxnTableWaitCountFlag

`func (o *VMImageUpgradeParams) HasSetTxnTableWaitCountFlag() bool`

HasSetTxnTableWaitCountFlag returns a boolean if a field has been set.

### GetSkipNodeChecks

`func (o *VMImageUpgradeParams) GetSkipNodeChecks() bool`

GetSkipNodeChecks returns the SkipNodeChecks field if non-nil, zero value otherwise.

### GetSkipNodeChecksOk

`func (o *VMImageUpgradeParams) GetSkipNodeChecksOk() (*bool, bool)`

GetSkipNodeChecksOk returns a tuple with the SkipNodeChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNodeChecks

`func (o *VMImageUpgradeParams) SetSkipNodeChecks(v bool)`

SetSkipNodeChecks sets SkipNodeChecks field to given value.

### HasSkipNodeChecks

`func (o *VMImageUpgradeParams) HasSkipNodeChecks() bool`

HasSkipNodeChecks returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *VMImageUpgradeParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *VMImageUpgradeParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *VMImageUpgradeParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *VMImageUpgradeParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *VMImageUpgradeParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *VMImageUpgradeParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSoftwareUpgradeState

`func (o *VMImageUpgradeParams) GetSoftwareUpgradeState() string`

GetSoftwareUpgradeState returns the SoftwareUpgradeState field if non-nil, zero value otherwise.

### GetSoftwareUpgradeStateOk

`func (o *VMImageUpgradeParams) GetSoftwareUpgradeStateOk() (*string, bool)`

GetSoftwareUpgradeStateOk returns a tuple with the SoftwareUpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpgradeState

`func (o *VMImageUpgradeParams) SetSoftwareUpgradeState(v string)`

SetSoftwareUpgradeState sets SoftwareUpgradeState field to given value.

### HasSoftwareUpgradeState

`func (o *VMImageUpgradeParams) HasSoftwareUpgradeState() bool`

HasSoftwareUpgradeState returns a boolean if a field has been set.

### GetSourceXClusterConfigs

`func (o *VMImageUpgradeParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *VMImageUpgradeParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *VMImageUpgradeParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *VMImageUpgradeParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSshUserOverride

`func (o *VMImageUpgradeParams) GetSshUserOverride() string`

GetSshUserOverride returns the SshUserOverride field if non-nil, zero value otherwise.

### GetSshUserOverrideOk

`func (o *VMImageUpgradeParams) GetSshUserOverrideOk() (*string, bool)`

GetSshUserOverrideOk returns a tuple with the SshUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverride

`func (o *VMImageUpgradeParams) SetSshUserOverride(v string)`

SetSshUserOverride sets SshUserOverride field to given value.

### HasSshUserOverride

`func (o *VMImageUpgradeParams) HasSshUserOverride() bool`

HasSshUserOverride returns a boolean if a field has been set.

### GetSshUserOverrideMap

`func (o *VMImageUpgradeParams) GetSshUserOverrideMap() map[string]string`

GetSshUserOverrideMap returns the SshUserOverrideMap field if non-nil, zero value otherwise.

### GetSshUserOverrideMapOk

`func (o *VMImageUpgradeParams) GetSshUserOverrideMapOk() (*map[string]string, bool)`

GetSshUserOverrideMapOk returns a tuple with the SshUserOverrideMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverrideMap

`func (o *VMImageUpgradeParams) SetSshUserOverrideMap(v map[string]string)`

SetSshUserOverrideMap sets SshUserOverrideMap field to given value.

### HasSshUserOverrideMap

`func (o *VMImageUpgradeParams) HasSshUserOverrideMap() bool`

HasSshUserOverrideMap returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *VMImageUpgradeParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *VMImageUpgradeParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *VMImageUpgradeParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *VMImageUpgradeParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetUniversePaused

`func (o *VMImageUpgradeParams) GetUniversePaused() bool`

GetUniversePaused returns the UniversePaused field if non-nil, zero value otherwise.

### GetUniversePausedOk

`func (o *VMImageUpgradeParams) GetUniversePausedOk() (*bool, bool)`

GetUniversePausedOk returns a tuple with the UniversePaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversePaused

`func (o *VMImageUpgradeParams) SetUniversePaused(v bool)`

SetUniversePaused sets UniversePaused field to given value.

### HasUniversePaused

`func (o *VMImageUpgradeParams) HasUniversePaused() bool`

HasUniversePaused returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *VMImageUpgradeParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *VMImageUpgradeParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *VMImageUpgradeParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *VMImageUpgradeParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUpdateInProgress

`func (o *VMImageUpgradeParams) GetUpdateInProgress() bool`

GetUpdateInProgress returns the UpdateInProgress field if non-nil, zero value otherwise.

### GetUpdateInProgressOk

`func (o *VMImageUpgradeParams) GetUpdateInProgressOk() (*bool, bool)`

GetUpdateInProgressOk returns a tuple with the UpdateInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInProgress

`func (o *VMImageUpgradeParams) SetUpdateInProgress(v bool)`

SetUpdateInProgress sets UpdateInProgress field to given value.

### HasUpdateInProgress

`func (o *VMImageUpgradeParams) HasUpdateInProgress() bool`

HasUpdateInProgress returns a boolean if a field has been set.

### GetUpdateOptions

`func (o *VMImageUpgradeParams) GetUpdateOptions() []string`

GetUpdateOptions returns the UpdateOptions field if non-nil, zero value otherwise.

### GetUpdateOptionsOk

`func (o *VMImageUpgradeParams) GetUpdateOptionsOk() (*[]string, bool)`

GetUpdateOptionsOk returns a tuple with the UpdateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOptions

`func (o *VMImageUpgradeParams) SetUpdateOptions(v []string)`

SetUpdateOptions sets UpdateOptions field to given value.

### HasUpdateOptions

`func (o *VMImageUpgradeParams) HasUpdateOptions() bool`

HasUpdateOptions returns a boolean if a field has been set.

### GetUpdateSucceeded

`func (o *VMImageUpgradeParams) GetUpdateSucceeded() bool`

GetUpdateSucceeded returns the UpdateSucceeded field if non-nil, zero value otherwise.

### GetUpdateSucceededOk

`func (o *VMImageUpgradeParams) GetUpdateSucceededOk() (*bool, bool)`

GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceeded

`func (o *VMImageUpgradeParams) SetUpdateSucceeded(v bool)`

SetUpdateSucceeded sets UpdateSucceeded field to given value.

### HasUpdateSucceeded

`func (o *VMImageUpgradeParams) HasUpdateSucceeded() bool`

HasUpdateSucceeded returns a boolean if a field has been set.

### GetUpdatingTask

`func (o *VMImageUpgradeParams) GetUpdatingTask() string`

GetUpdatingTask returns the UpdatingTask field if non-nil, zero value otherwise.

### GetUpdatingTaskOk

`func (o *VMImageUpgradeParams) GetUpdatingTaskOk() (*string, bool)`

GetUpdatingTaskOk returns a tuple with the UpdatingTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTask

`func (o *VMImageUpgradeParams) SetUpdatingTask(v string)`

SetUpdatingTask sets UpdatingTask field to given value.

### HasUpdatingTask

`func (o *VMImageUpgradeParams) HasUpdatingTask() bool`

HasUpdatingTask returns a boolean if a field has been set.

### GetUpdatingTaskUUID

`func (o *VMImageUpgradeParams) GetUpdatingTaskUUID() string`

GetUpdatingTaskUUID returns the UpdatingTaskUUID field if non-nil, zero value otherwise.

### GetUpdatingTaskUUIDOk

`func (o *VMImageUpgradeParams) GetUpdatingTaskUUIDOk() (*string, bool)`

GetUpdatingTaskUUIDOk returns a tuple with the UpdatingTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTaskUUID

`func (o *VMImageUpgradeParams) SetUpdatingTaskUUID(v string)`

SetUpdatingTaskUUID sets UpdatingTaskUUID field to given value.

### HasUpdatingTaskUUID

`func (o *VMImageUpgradeParams) HasUpdatingTaskUUID() bool`

HasUpdatingTaskUUID returns a boolean if a field has been set.

### GetUpgradeOption

`func (o *VMImageUpgradeParams) GetUpgradeOption() string`

GetUpgradeOption returns the UpgradeOption field if non-nil, zero value otherwise.

### GetUpgradeOptionOk

`func (o *VMImageUpgradeParams) GetUpgradeOptionOk() (*string, bool)`

GetUpgradeOptionOk returns a tuple with the UpgradeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOption

`func (o *VMImageUpgradeParams) SetUpgradeOption(v string)`

SetUpgradeOption sets UpgradeOption field to given value.


### GetUseNewHelmNamingStyle

`func (o *VMImageUpgradeParams) GetUseNewHelmNamingStyle() bool`

GetUseNewHelmNamingStyle returns the UseNewHelmNamingStyle field if non-nil, zero value otherwise.

### GetUseNewHelmNamingStyleOk

`func (o *VMImageUpgradeParams) GetUseNewHelmNamingStyleOk() (*bool, bool)`

GetUseNewHelmNamingStyleOk returns a tuple with the UseNewHelmNamingStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNewHelmNamingStyle

`func (o *VMImageUpgradeParams) SetUseNewHelmNamingStyle(v bool)`

SetUseNewHelmNamingStyle sets UseNewHelmNamingStyle field to given value.

### HasUseNewHelmNamingStyle

`func (o *VMImageUpgradeParams) HasUseNewHelmNamingStyle() bool`

HasUseNewHelmNamingStyle returns a boolean if a field has been set.

### GetUserAZSelected

`func (o *VMImageUpgradeParams) GetUserAZSelected() bool`

GetUserAZSelected returns the UserAZSelected field if non-nil, zero value otherwise.

### GetUserAZSelectedOk

`func (o *VMImageUpgradeParams) GetUserAZSelectedOk() (*bool, bool)`

GetUserAZSelectedOk returns a tuple with the UserAZSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAZSelected

`func (o *VMImageUpgradeParams) SetUserAZSelected(v bool)`

SetUserAZSelected sets UserAZSelected field to given value.

### HasUserAZSelected

`func (o *VMImageUpgradeParams) HasUserAZSelected() bool`

HasUserAZSelected returns a boolean if a field has been set.

### GetXclusterInfo

`func (o *VMImageUpgradeParams) GetXclusterInfo() XClusterInfo`

GetXclusterInfo returns the XclusterInfo field if non-nil, zero value otherwise.

### GetXclusterInfoOk

`func (o *VMImageUpgradeParams) GetXclusterInfoOk() (*XClusterInfo, bool)`

GetXclusterInfoOk returns a tuple with the XclusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterInfo

`func (o *VMImageUpgradeParams) SetXclusterInfo(v XClusterInfo)`

SetXclusterInfo sets XclusterInfo field to given value.

### HasXclusterInfo

`func (o *VMImageUpgradeParams) HasXclusterInfo() bool`

HasXclusterInfo returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *VMImageUpgradeParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *VMImageUpgradeParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *VMImageUpgradeParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *VMImageUpgradeParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbSoftwareVersion

`func (o *VMImageUpgradeParams) GetYbSoftwareVersion() string`

GetYbSoftwareVersion returns the YbSoftwareVersion field if non-nil, zero value otherwise.

### GetYbSoftwareVersionOk

`func (o *VMImageUpgradeParams) GetYbSoftwareVersionOk() (*string, bool)`

GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbSoftwareVersion

`func (o *VMImageUpgradeParams) SetYbSoftwareVersion(v string)`

SetYbSoftwareVersion sets YbSoftwareVersion field to given value.


### GetYbcInstalled

`func (o *VMImageUpgradeParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *VMImageUpgradeParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *VMImageUpgradeParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *VMImageUpgradeParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *VMImageUpgradeParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *VMImageUpgradeParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *VMImageUpgradeParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *VMImageUpgradeParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


