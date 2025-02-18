# RestoreBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** | Action type | [optional] 
**AlterLoadBalancer** | Pointer to **bool** | Alter load balancer state | [optional] 
**BackupStorageInfoList** | Pointer to [**[]BackupStorageInfo**](BackupStorageInfo.md) | Backup&#39;s storage info to restore | [optional] 
**Category** | Pointer to **string** | Category of the backup | [optional] 
**CmkArn** | Pointer to **string** | Amazon Resource Name (ARN) of the CMK | [optional] 
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**CreatingUser** | [**Users**](Users.md) |  | 
**CustomerUUID** | Pointer to **string** | Customer UUID | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**DisableChecksum** | Pointer to **bool** | Disable checksum | [optional] 
**DisableMultipart** | Pointer to **bool** | Disable multipart upload | [optional] 
**EnableVerboseLogs** | Pointer to **bool** | Is verbose logging enabled | [optional] 
**EnableYbc** | Pointer to **bool** |  | [optional] 
**EncryptionAtRestConfig** | Pointer to [**EncryptionAtRestConfig**](EncryptionAtRestConfig.md) |  | [optional] 
**ErrorString** | Pointer to **string** | Error message | [optional] 
**ExpectedUniverseVersion** | Pointer to **int32** | Expected universe version | [optional] 
**ExtraDependencies** | Pointer to [**ExtraDependencies**](ExtraDependencies.md) |  | [optional] 
**InstallYbc** | Pointer to **bool** |  | [optional] 
**KmsConfigUUID** | Pointer to **string** | KMS configuration UUID | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**NodeIp** | Pointer to **string** | Node IP | [optional] 
**Parallelism** | Pointer to **int32** | Number of concurrent commands to run on nodes over SSH | [optional] 
**PlatformUrl** | **string** |  | 
**PlatformVersion** | **string** |  | 
**PrefixUUID** | Pointer to **string** | Prefix UUID | [optional] 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID of a retry | [optional] 
**RestoreTimeStamp** | Pointer to **string** | Restore TimeStamp | [optional] 
**RestoreToPointInTimeMillis** | Pointer to **int64** | Restore timestamp in millis | [optional] 
**SleepAfterMasterRestartMillis** | **int32** |  | 
**SleepAfterTServerRestartMillis** | **int32** |  | 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] [readonly] 
**StorageConfigUUID** | Pointer to **string** | Storage config uuid | [optional] 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] [readonly] 
**UniverseUUID** | **string** | Universe UUID | 
**UseTablespaces** | Pointer to **bool** | Is tablespaces information included | [optional] 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 
**YbcInstalled** | Pointer to **bool** |  | [optional] 
**YbcSoftwareVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewRestoreBackupParams

`func NewRestoreBackupParams(creatingUser Users, platformUrl string, platformVersion string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, universeUUID string, ) *RestoreBackupParams`

NewRestoreBackupParams instantiates a new RestoreBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreBackupParamsWithDefaults

`func NewRestoreBackupParamsWithDefaults() *RestoreBackupParams`

NewRestoreBackupParamsWithDefaults instantiates a new RestoreBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *RestoreBackupParams) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *RestoreBackupParams) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *RestoreBackupParams) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *RestoreBackupParams) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetAlterLoadBalancer

`func (o *RestoreBackupParams) GetAlterLoadBalancer() bool`

GetAlterLoadBalancer returns the AlterLoadBalancer field if non-nil, zero value otherwise.

### GetAlterLoadBalancerOk

`func (o *RestoreBackupParams) GetAlterLoadBalancerOk() (*bool, bool)`

GetAlterLoadBalancerOk returns a tuple with the AlterLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterLoadBalancer

`func (o *RestoreBackupParams) SetAlterLoadBalancer(v bool)`

SetAlterLoadBalancer sets AlterLoadBalancer field to given value.

### HasAlterLoadBalancer

`func (o *RestoreBackupParams) HasAlterLoadBalancer() bool`

HasAlterLoadBalancer returns a boolean if a field has been set.

### GetBackupStorageInfoList

`func (o *RestoreBackupParams) GetBackupStorageInfoList() []BackupStorageInfo`

GetBackupStorageInfoList returns the BackupStorageInfoList field if non-nil, zero value otherwise.

### GetBackupStorageInfoListOk

`func (o *RestoreBackupParams) GetBackupStorageInfoListOk() (*[]BackupStorageInfo, bool)`

GetBackupStorageInfoListOk returns a tuple with the BackupStorageInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorageInfoList

`func (o *RestoreBackupParams) SetBackupStorageInfoList(v []BackupStorageInfo)`

SetBackupStorageInfoList sets BackupStorageInfoList field to given value.

### HasBackupStorageInfoList

`func (o *RestoreBackupParams) HasBackupStorageInfoList() bool`

HasBackupStorageInfoList returns a boolean if a field has been set.

### GetCategory

`func (o *RestoreBackupParams) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RestoreBackupParams) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RestoreBackupParams) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *RestoreBackupParams) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCmkArn

`func (o *RestoreBackupParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *RestoreBackupParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *RestoreBackupParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *RestoreBackupParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *RestoreBackupParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *RestoreBackupParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *RestoreBackupParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *RestoreBackupParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *RestoreBackupParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *RestoreBackupParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *RestoreBackupParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.

### HasCurrentYbcTaskId

`func (o *RestoreBackupParams) HasCurrentYbcTaskId() bool`

HasCurrentYbcTaskId returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *RestoreBackupParams) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *RestoreBackupParams) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *RestoreBackupParams) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *RestoreBackupParams) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *RestoreBackupParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *RestoreBackupParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *RestoreBackupParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *RestoreBackupParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetDisableChecksum

`func (o *RestoreBackupParams) GetDisableChecksum() bool`

GetDisableChecksum returns the DisableChecksum field if non-nil, zero value otherwise.

### GetDisableChecksumOk

`func (o *RestoreBackupParams) GetDisableChecksumOk() (*bool, bool)`

GetDisableChecksumOk returns a tuple with the DisableChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableChecksum

`func (o *RestoreBackupParams) SetDisableChecksum(v bool)`

SetDisableChecksum sets DisableChecksum field to given value.

### HasDisableChecksum

`func (o *RestoreBackupParams) HasDisableChecksum() bool`

HasDisableChecksum returns a boolean if a field has been set.

### GetDisableMultipart

`func (o *RestoreBackupParams) GetDisableMultipart() bool`

GetDisableMultipart returns the DisableMultipart field if non-nil, zero value otherwise.

### GetDisableMultipartOk

`func (o *RestoreBackupParams) GetDisableMultipartOk() (*bool, bool)`

GetDisableMultipartOk returns a tuple with the DisableMultipart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMultipart

`func (o *RestoreBackupParams) SetDisableMultipart(v bool)`

SetDisableMultipart sets DisableMultipart field to given value.

### HasDisableMultipart

`func (o *RestoreBackupParams) HasDisableMultipart() bool`

HasDisableMultipart returns a boolean if a field has been set.

### GetEnableVerboseLogs

`func (o *RestoreBackupParams) GetEnableVerboseLogs() bool`

GetEnableVerboseLogs returns the EnableVerboseLogs field if non-nil, zero value otherwise.

### GetEnableVerboseLogsOk

`func (o *RestoreBackupParams) GetEnableVerboseLogsOk() (*bool, bool)`

GetEnableVerboseLogsOk returns a tuple with the EnableVerboseLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVerboseLogs

`func (o *RestoreBackupParams) SetEnableVerboseLogs(v bool)`

SetEnableVerboseLogs sets EnableVerboseLogs field to given value.

### HasEnableVerboseLogs

`func (o *RestoreBackupParams) HasEnableVerboseLogs() bool`

HasEnableVerboseLogs returns a boolean if a field has been set.

### GetEnableYbc

`func (o *RestoreBackupParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *RestoreBackupParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *RestoreBackupParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *RestoreBackupParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *RestoreBackupParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *RestoreBackupParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *RestoreBackupParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *RestoreBackupParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *RestoreBackupParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *RestoreBackupParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *RestoreBackupParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *RestoreBackupParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *RestoreBackupParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *RestoreBackupParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *RestoreBackupParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *RestoreBackupParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *RestoreBackupParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *RestoreBackupParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *RestoreBackupParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *RestoreBackupParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetInstallYbc

`func (o *RestoreBackupParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *RestoreBackupParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *RestoreBackupParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *RestoreBackupParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetKmsConfigUUID

`func (o *RestoreBackupParams) GetKmsConfigUUID() string`

GetKmsConfigUUID returns the KmsConfigUUID field if non-nil, zero value otherwise.

### GetKmsConfigUUIDOk

`func (o *RestoreBackupParams) GetKmsConfigUUIDOk() (*string, bool)`

GetKmsConfigUUIDOk returns a tuple with the KmsConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUUID

`func (o *RestoreBackupParams) SetKmsConfigUUID(v string)`

SetKmsConfigUUID sets KmsConfigUUID field to given value.

### HasKmsConfigUUID

`func (o *RestoreBackupParams) HasKmsConfigUUID() bool`

HasKmsConfigUUID returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *RestoreBackupParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *RestoreBackupParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *RestoreBackupParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *RestoreBackupParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *RestoreBackupParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *RestoreBackupParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *RestoreBackupParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *RestoreBackupParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNodeIp

`func (o *RestoreBackupParams) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *RestoreBackupParams) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *RestoreBackupParams) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *RestoreBackupParams) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### GetParallelism

`func (o *RestoreBackupParams) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *RestoreBackupParams) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *RestoreBackupParams) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *RestoreBackupParams) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *RestoreBackupParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *RestoreBackupParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *RestoreBackupParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPlatformVersion

`func (o *RestoreBackupParams) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *RestoreBackupParams) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *RestoreBackupParams) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.


### GetPrefixUUID

`func (o *RestoreBackupParams) GetPrefixUUID() string`

GetPrefixUUID returns the PrefixUUID field if non-nil, zero value otherwise.

### GetPrefixUUIDOk

`func (o *RestoreBackupParams) GetPrefixUUIDOk() (*string, bool)`

GetPrefixUUIDOk returns a tuple with the PrefixUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixUUID

`func (o *RestoreBackupParams) SetPrefixUUID(v string)`

SetPrefixUUID sets PrefixUUID field to given value.

### HasPrefixUUID

`func (o *RestoreBackupParams) HasPrefixUUID() bool`

HasPrefixUUID returns a boolean if a field has been set.

### GetPreviousTaskUUID

`func (o *RestoreBackupParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *RestoreBackupParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *RestoreBackupParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *RestoreBackupParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRestoreTimeStamp

`func (o *RestoreBackupParams) GetRestoreTimeStamp() string`

GetRestoreTimeStamp returns the RestoreTimeStamp field if non-nil, zero value otherwise.

### GetRestoreTimeStampOk

`func (o *RestoreBackupParams) GetRestoreTimeStampOk() (*string, bool)`

GetRestoreTimeStampOk returns a tuple with the RestoreTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeStamp

`func (o *RestoreBackupParams) SetRestoreTimeStamp(v string)`

SetRestoreTimeStamp sets RestoreTimeStamp field to given value.

### HasRestoreTimeStamp

`func (o *RestoreBackupParams) HasRestoreTimeStamp() bool`

HasRestoreTimeStamp returns a boolean if a field has been set.

### GetRestoreToPointInTimeMillis

`func (o *RestoreBackupParams) GetRestoreToPointInTimeMillis() int64`

GetRestoreToPointInTimeMillis returns the RestoreToPointInTimeMillis field if non-nil, zero value otherwise.

### GetRestoreToPointInTimeMillisOk

`func (o *RestoreBackupParams) GetRestoreToPointInTimeMillisOk() (*int64, bool)`

GetRestoreToPointInTimeMillisOk returns a tuple with the RestoreToPointInTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToPointInTimeMillis

`func (o *RestoreBackupParams) SetRestoreToPointInTimeMillis(v int64)`

SetRestoreToPointInTimeMillis sets RestoreToPointInTimeMillis field to given value.

### HasRestoreToPointInTimeMillis

`func (o *RestoreBackupParams) HasRestoreToPointInTimeMillis() bool`

HasRestoreToPointInTimeMillis returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *RestoreBackupParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *RestoreBackupParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *RestoreBackupParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *RestoreBackupParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *RestoreBackupParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *RestoreBackupParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSourceXClusterConfigs

`func (o *RestoreBackupParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *RestoreBackupParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *RestoreBackupParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *RestoreBackupParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *RestoreBackupParams) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *RestoreBackupParams) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *RestoreBackupParams) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.

### HasStorageConfigUUID

`func (o *RestoreBackupParams) HasStorageConfigUUID() bool`

HasStorageConfigUUID returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *RestoreBackupParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *RestoreBackupParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *RestoreBackupParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *RestoreBackupParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *RestoreBackupParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *RestoreBackupParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *RestoreBackupParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.


### GetUseTablespaces

`func (o *RestoreBackupParams) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *RestoreBackupParams) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *RestoreBackupParams) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.

### HasUseTablespaces

`func (o *RestoreBackupParams) HasUseTablespaces() bool`

HasUseTablespaces returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *RestoreBackupParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *RestoreBackupParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *RestoreBackupParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *RestoreBackupParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *RestoreBackupParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *RestoreBackupParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *RestoreBackupParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *RestoreBackupParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *RestoreBackupParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *RestoreBackupParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *RestoreBackupParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *RestoreBackupParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


