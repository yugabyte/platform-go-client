# BackupRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlterLoadBalancer** | Pointer to **bool** | Alter load balancer state | [optional] 
**BackupType** | Pointer to **string** | Backup type | [optional] 
**BackupUUID** | **string** |  | 
**BaseBackupUUID** | Pointer to **string** | UUID of the parent backup | [optional] 
**CmkArn** | Pointer to **string** | Amazon Resource Name (ARN) of the CMK | [optional] 
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**CreatingUser** | [**Users**](Users.md) |  | 
**CronExpression** | Pointer to **string** | Cron expression for a recurring backup | [optional] 
**CurrentIdx** | **int32** |  | 
**CurrentYbcTaskId** | **string** |  | 
**CustomerUUID** | Pointer to **string** | Customer UUID | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**DisableChecksum** | Pointer to **bool** | Disable checksum | [optional] 
**DisableMultipart** | Pointer to **bool** | Disable multipart upload | [optional] 
**DisableParallelism** | Pointer to **bool** | Don&#39;t add -m flag during gsutil upload dir command | [optional] 
**EnableVerboseLogs** | Pointer to **bool** | Is verbose logging enabled | [optional] 
**EnableYbc** | Pointer to **bool** |  | [optional] 
**EncryptionAtRestConfig** | Pointer to [**EncryptionAtRestConfig**](EncryptionAtRestConfig.md) |  | [optional] 
**ErrorString** | Pointer to **string** | Error message | [optional] 
**ExpectedUniverseVersion** | Pointer to **int32** | Expected universe version | [optional] 
**ExpiryTimeUnit** | Pointer to **string** | Time unit for backup expiry time | [optional] 
**ExtraDependencies** | Pointer to [**ExtraDependencies**](ExtraDependencies.md) |  | [optional] 
**FirstTry** | Pointer to **bool** | Whether this task has been tried before | [optional] 
**FrequencyTimeUnit** | Pointer to **string** | Time unit for user input schedule frequency | [optional] 
**IgnoreErrors** | Pointer to **bool** | Should table backup errors be ignored | [optional] 
**IncrementalBackupFrequency** | Pointer to **int64** | Frequency of incremental backups | [optional] 
**IncrementalBackupFrequencyTimeUnit** | Pointer to **string** | Time unit for user input incremental backup schedule frequency | [optional] 
**InstallYbc** | Pointer to **bool** |  | [optional] 
**KeyspaceTableList** | Pointer to [**[]KeyspaceTable**](KeyspaceTable.md) | Backup info | [optional] 
**KmsConfigUUID** | Pointer to **string** | KMS configuration UUID | [optional] 
**MinNumBackupsToRetain** | Pointer to **int32** | Minimum number of backups to retain for a particular backup schedule | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**Parallelism** | Pointer to **int32** | Number of concurrent commands to run on nodes over SSH | [optional] 
**PlatformUrl** | **string** |  | 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID of a retry | [optional] 
**ScheduleName** | Pointer to **string** | Schedule Name | [optional] 
**ScheduleUUID** | Pointer to **string** | Schedule UUID | [optional] 
**SchedulingFrequency** | Pointer to **int64** | Frequency to run the backup, in milliseconds | [optional] 
**SleepAfterMasterRestartMillis** | **int32** |  | 
**SleepAfterTServerRestartMillis** | **int32** |  | 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] [readonly] 
**Sse** | Pointer to **bool** | Is SSE | [optional] 
**StorageConfigUUID** | **string** | Storage configuration UUID | 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] [readonly] 
**TimeBeforeDelete** | Pointer to **int64** | Time before deleting the backup from storage, in milliseconds | [optional] 
**UniverseUUID** | **string** | Universe UUID | 
**UseTablespaces** | Pointer to **bool** | Is tablespaces information included | [optional] 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 
**YbcInstalled** | Pointer to **bool** |  | [optional] 
**YbcSoftwareVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewBackupRequestParams

`func NewBackupRequestParams(backupUUID string, creatingUser Users, currentIdx int32, currentYbcTaskId string, platformUrl string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, storageConfigUUID string, universeUUID string, ) *BackupRequestParams`

NewBackupRequestParams instantiates a new BackupRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRequestParamsWithDefaults

`func NewBackupRequestParamsWithDefaults() *BackupRequestParams`

NewBackupRequestParamsWithDefaults instantiates a new BackupRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlterLoadBalancer

`func (o *BackupRequestParams) GetAlterLoadBalancer() bool`

GetAlterLoadBalancer returns the AlterLoadBalancer field if non-nil, zero value otherwise.

### GetAlterLoadBalancerOk

`func (o *BackupRequestParams) GetAlterLoadBalancerOk() (*bool, bool)`

GetAlterLoadBalancerOk returns a tuple with the AlterLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterLoadBalancer

`func (o *BackupRequestParams) SetAlterLoadBalancer(v bool)`

SetAlterLoadBalancer sets AlterLoadBalancer field to given value.

### HasAlterLoadBalancer

`func (o *BackupRequestParams) HasAlterLoadBalancer() bool`

HasAlterLoadBalancer returns a boolean if a field has been set.

### GetBackupType

`func (o *BackupRequestParams) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupRequestParams) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupRequestParams) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *BackupRequestParams) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetBackupUUID

`func (o *BackupRequestParams) GetBackupUUID() string`

GetBackupUUID returns the BackupUUID field if non-nil, zero value otherwise.

### GetBackupUUIDOk

`func (o *BackupRequestParams) GetBackupUUIDOk() (*string, bool)`

GetBackupUUIDOk returns a tuple with the BackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUID

`func (o *BackupRequestParams) SetBackupUUID(v string)`

SetBackupUUID sets BackupUUID field to given value.


### GetBaseBackupUUID

`func (o *BackupRequestParams) GetBaseBackupUUID() string`

GetBaseBackupUUID returns the BaseBackupUUID field if non-nil, zero value otherwise.

### GetBaseBackupUUIDOk

`func (o *BackupRequestParams) GetBaseBackupUUIDOk() (*string, bool)`

GetBaseBackupUUIDOk returns a tuple with the BaseBackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBackupUUID

`func (o *BackupRequestParams) SetBaseBackupUUID(v string)`

SetBaseBackupUUID sets BaseBackupUUID field to given value.

### HasBaseBackupUUID

`func (o *BackupRequestParams) HasBaseBackupUUID() bool`

HasBaseBackupUUID returns a boolean if a field has been set.

### GetCmkArn

`func (o *BackupRequestParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *BackupRequestParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *BackupRequestParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *BackupRequestParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *BackupRequestParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *BackupRequestParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *BackupRequestParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *BackupRequestParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *BackupRequestParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *BackupRequestParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *BackupRequestParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCronExpression

`func (o *BackupRequestParams) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *BackupRequestParams) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *BackupRequestParams) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *BackupRequestParams) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetCurrentIdx

`func (o *BackupRequestParams) GetCurrentIdx() int32`

GetCurrentIdx returns the CurrentIdx field if non-nil, zero value otherwise.

### GetCurrentIdxOk

`func (o *BackupRequestParams) GetCurrentIdxOk() (*int32, bool)`

GetCurrentIdxOk returns a tuple with the CurrentIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIdx

`func (o *BackupRequestParams) SetCurrentIdx(v int32)`

SetCurrentIdx sets CurrentIdx field to given value.


### GetCurrentYbcTaskId

`func (o *BackupRequestParams) GetCurrentYbcTaskId() string`

GetCurrentYbcTaskId returns the CurrentYbcTaskId field if non-nil, zero value otherwise.

### GetCurrentYbcTaskIdOk

`func (o *BackupRequestParams) GetCurrentYbcTaskIdOk() (*string, bool)`

GetCurrentYbcTaskIdOk returns a tuple with the CurrentYbcTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentYbcTaskId

`func (o *BackupRequestParams) SetCurrentYbcTaskId(v string)`

SetCurrentYbcTaskId sets CurrentYbcTaskId field to given value.


### GetCustomerUUID

`func (o *BackupRequestParams) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *BackupRequestParams) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *BackupRequestParams) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *BackupRequestParams) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *BackupRequestParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *BackupRequestParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *BackupRequestParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *BackupRequestParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetDisableChecksum

`func (o *BackupRequestParams) GetDisableChecksum() bool`

GetDisableChecksum returns the DisableChecksum field if non-nil, zero value otherwise.

### GetDisableChecksumOk

`func (o *BackupRequestParams) GetDisableChecksumOk() (*bool, bool)`

GetDisableChecksumOk returns a tuple with the DisableChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableChecksum

`func (o *BackupRequestParams) SetDisableChecksum(v bool)`

SetDisableChecksum sets DisableChecksum field to given value.

### HasDisableChecksum

`func (o *BackupRequestParams) HasDisableChecksum() bool`

HasDisableChecksum returns a boolean if a field has been set.

### GetDisableMultipart

`func (o *BackupRequestParams) GetDisableMultipart() bool`

GetDisableMultipart returns the DisableMultipart field if non-nil, zero value otherwise.

### GetDisableMultipartOk

`func (o *BackupRequestParams) GetDisableMultipartOk() (*bool, bool)`

GetDisableMultipartOk returns a tuple with the DisableMultipart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMultipart

`func (o *BackupRequestParams) SetDisableMultipart(v bool)`

SetDisableMultipart sets DisableMultipart field to given value.

### HasDisableMultipart

`func (o *BackupRequestParams) HasDisableMultipart() bool`

HasDisableMultipart returns a boolean if a field has been set.

### GetDisableParallelism

`func (o *BackupRequestParams) GetDisableParallelism() bool`

GetDisableParallelism returns the DisableParallelism field if non-nil, zero value otherwise.

### GetDisableParallelismOk

`func (o *BackupRequestParams) GetDisableParallelismOk() (*bool, bool)`

GetDisableParallelismOk returns a tuple with the DisableParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableParallelism

`func (o *BackupRequestParams) SetDisableParallelism(v bool)`

SetDisableParallelism sets DisableParallelism field to given value.

### HasDisableParallelism

`func (o *BackupRequestParams) HasDisableParallelism() bool`

HasDisableParallelism returns a boolean if a field has been set.

### GetEnableVerboseLogs

`func (o *BackupRequestParams) GetEnableVerboseLogs() bool`

GetEnableVerboseLogs returns the EnableVerboseLogs field if non-nil, zero value otherwise.

### GetEnableVerboseLogsOk

`func (o *BackupRequestParams) GetEnableVerboseLogsOk() (*bool, bool)`

GetEnableVerboseLogsOk returns a tuple with the EnableVerboseLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVerboseLogs

`func (o *BackupRequestParams) SetEnableVerboseLogs(v bool)`

SetEnableVerboseLogs sets EnableVerboseLogs field to given value.

### HasEnableVerboseLogs

`func (o *BackupRequestParams) HasEnableVerboseLogs() bool`

HasEnableVerboseLogs returns a boolean if a field has been set.

### GetEnableYbc

`func (o *BackupRequestParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *BackupRequestParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *BackupRequestParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *BackupRequestParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *BackupRequestParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *BackupRequestParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *BackupRequestParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *BackupRequestParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *BackupRequestParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *BackupRequestParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *BackupRequestParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *BackupRequestParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *BackupRequestParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *BackupRequestParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *BackupRequestParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *BackupRequestParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExpiryTimeUnit

`func (o *BackupRequestParams) GetExpiryTimeUnit() string`

GetExpiryTimeUnit returns the ExpiryTimeUnit field if non-nil, zero value otherwise.

### GetExpiryTimeUnitOk

`func (o *BackupRequestParams) GetExpiryTimeUnitOk() (*string, bool)`

GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUnit

`func (o *BackupRequestParams) SetExpiryTimeUnit(v string)`

SetExpiryTimeUnit sets ExpiryTimeUnit field to given value.

### HasExpiryTimeUnit

`func (o *BackupRequestParams) HasExpiryTimeUnit() bool`

HasExpiryTimeUnit returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *BackupRequestParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *BackupRequestParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *BackupRequestParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *BackupRequestParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetFirstTry

`func (o *BackupRequestParams) GetFirstTry() bool`

GetFirstTry returns the FirstTry field if non-nil, zero value otherwise.

### GetFirstTryOk

`func (o *BackupRequestParams) GetFirstTryOk() (*bool, bool)`

GetFirstTryOk returns a tuple with the FirstTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTry

`func (o *BackupRequestParams) SetFirstTry(v bool)`

SetFirstTry sets FirstTry field to given value.

### HasFirstTry

`func (o *BackupRequestParams) HasFirstTry() bool`

HasFirstTry returns a boolean if a field has been set.

### GetFrequencyTimeUnit

`func (o *BackupRequestParams) GetFrequencyTimeUnit() string`

GetFrequencyTimeUnit returns the FrequencyTimeUnit field if non-nil, zero value otherwise.

### GetFrequencyTimeUnitOk

`func (o *BackupRequestParams) GetFrequencyTimeUnitOk() (*string, bool)`

GetFrequencyTimeUnitOk returns a tuple with the FrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyTimeUnit

`func (o *BackupRequestParams) SetFrequencyTimeUnit(v string)`

SetFrequencyTimeUnit sets FrequencyTimeUnit field to given value.

### HasFrequencyTimeUnit

`func (o *BackupRequestParams) HasFrequencyTimeUnit() bool`

HasFrequencyTimeUnit returns a boolean if a field has been set.

### GetIgnoreErrors

`func (o *BackupRequestParams) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *BackupRequestParams) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *BackupRequestParams) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.

### HasIgnoreErrors

`func (o *BackupRequestParams) HasIgnoreErrors() bool`

HasIgnoreErrors returns a boolean if a field has been set.

### GetIncrementalBackupFrequency

`func (o *BackupRequestParams) GetIncrementalBackupFrequency() int64`

GetIncrementalBackupFrequency returns the IncrementalBackupFrequency field if non-nil, zero value otherwise.

### GetIncrementalBackupFrequencyOk

`func (o *BackupRequestParams) GetIncrementalBackupFrequencyOk() (*int64, bool)`

GetIncrementalBackupFrequencyOk returns a tuple with the IncrementalBackupFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupFrequency

`func (o *BackupRequestParams) SetIncrementalBackupFrequency(v int64)`

SetIncrementalBackupFrequency sets IncrementalBackupFrequency field to given value.

### HasIncrementalBackupFrequency

`func (o *BackupRequestParams) HasIncrementalBackupFrequency() bool`

HasIncrementalBackupFrequency returns a boolean if a field has been set.

### GetIncrementalBackupFrequencyTimeUnit

`func (o *BackupRequestParams) GetIncrementalBackupFrequencyTimeUnit() string`

GetIncrementalBackupFrequencyTimeUnit returns the IncrementalBackupFrequencyTimeUnit field if non-nil, zero value otherwise.

### GetIncrementalBackupFrequencyTimeUnitOk

`func (o *BackupRequestParams) GetIncrementalBackupFrequencyTimeUnitOk() (*string, bool)`

GetIncrementalBackupFrequencyTimeUnitOk returns a tuple with the IncrementalBackupFrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupFrequencyTimeUnit

`func (o *BackupRequestParams) SetIncrementalBackupFrequencyTimeUnit(v string)`

SetIncrementalBackupFrequencyTimeUnit sets IncrementalBackupFrequencyTimeUnit field to given value.

### HasIncrementalBackupFrequencyTimeUnit

`func (o *BackupRequestParams) HasIncrementalBackupFrequencyTimeUnit() bool`

HasIncrementalBackupFrequencyTimeUnit returns a boolean if a field has been set.

### GetInstallYbc

`func (o *BackupRequestParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *BackupRequestParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *BackupRequestParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *BackupRequestParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetKeyspaceTableList

`func (o *BackupRequestParams) GetKeyspaceTableList() []KeyspaceTable`

GetKeyspaceTableList returns the KeyspaceTableList field if non-nil, zero value otherwise.

### GetKeyspaceTableListOk

`func (o *BackupRequestParams) GetKeyspaceTableListOk() (*[]KeyspaceTable, bool)`

GetKeyspaceTableListOk returns a tuple with the KeyspaceTableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceTableList

`func (o *BackupRequestParams) SetKeyspaceTableList(v []KeyspaceTable)`

SetKeyspaceTableList sets KeyspaceTableList field to given value.

### HasKeyspaceTableList

`func (o *BackupRequestParams) HasKeyspaceTableList() bool`

HasKeyspaceTableList returns a boolean if a field has been set.

### GetKmsConfigUUID

`func (o *BackupRequestParams) GetKmsConfigUUID() string`

GetKmsConfigUUID returns the KmsConfigUUID field if non-nil, zero value otherwise.

### GetKmsConfigUUIDOk

`func (o *BackupRequestParams) GetKmsConfigUUIDOk() (*string, bool)`

GetKmsConfigUUIDOk returns a tuple with the KmsConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUUID

`func (o *BackupRequestParams) SetKmsConfigUUID(v string)`

SetKmsConfigUUID sets KmsConfigUUID field to given value.

### HasKmsConfigUUID

`func (o *BackupRequestParams) HasKmsConfigUUID() bool`

HasKmsConfigUUID returns a boolean if a field has been set.

### GetMinNumBackupsToRetain

`func (o *BackupRequestParams) GetMinNumBackupsToRetain() int32`

GetMinNumBackupsToRetain returns the MinNumBackupsToRetain field if non-nil, zero value otherwise.

### GetMinNumBackupsToRetainOk

`func (o *BackupRequestParams) GetMinNumBackupsToRetainOk() (*int32, bool)`

GetMinNumBackupsToRetainOk returns a tuple with the MinNumBackupsToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumBackupsToRetain

`func (o *BackupRequestParams) SetMinNumBackupsToRetain(v int32)`

SetMinNumBackupsToRetain sets MinNumBackupsToRetain field to given value.

### HasMinNumBackupsToRetain

`func (o *BackupRequestParams) HasMinNumBackupsToRetain() bool`

HasMinNumBackupsToRetain returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *BackupRequestParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *BackupRequestParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *BackupRequestParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *BackupRequestParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *BackupRequestParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *BackupRequestParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *BackupRequestParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *BackupRequestParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetParallelism

`func (o *BackupRequestParams) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *BackupRequestParams) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *BackupRequestParams) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *BackupRequestParams) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *BackupRequestParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *BackupRequestParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *BackupRequestParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPreviousTaskUUID

`func (o *BackupRequestParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *BackupRequestParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *BackupRequestParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *BackupRequestParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetScheduleName

`func (o *BackupRequestParams) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupRequestParams) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupRequestParams) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *BackupRequestParams) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetScheduleUUID

`func (o *BackupRequestParams) GetScheduleUUID() string`

GetScheduleUUID returns the ScheduleUUID field if non-nil, zero value otherwise.

### GetScheduleUUIDOk

`func (o *BackupRequestParams) GetScheduleUUIDOk() (*string, bool)`

GetScheduleUUIDOk returns a tuple with the ScheduleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUID

`func (o *BackupRequestParams) SetScheduleUUID(v string)`

SetScheduleUUID sets ScheduleUUID field to given value.

### HasScheduleUUID

`func (o *BackupRequestParams) HasScheduleUUID() bool`

HasScheduleUUID returns a boolean if a field has been set.

### GetSchedulingFrequency

`func (o *BackupRequestParams) GetSchedulingFrequency() int64`

GetSchedulingFrequency returns the SchedulingFrequency field if non-nil, zero value otherwise.

### GetSchedulingFrequencyOk

`func (o *BackupRequestParams) GetSchedulingFrequencyOk() (*int64, bool)`

GetSchedulingFrequencyOk returns a tuple with the SchedulingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingFrequency

`func (o *BackupRequestParams) SetSchedulingFrequency(v int64)`

SetSchedulingFrequency sets SchedulingFrequency field to given value.

### HasSchedulingFrequency

`func (o *BackupRequestParams) HasSchedulingFrequency() bool`

HasSchedulingFrequency returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *BackupRequestParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *BackupRequestParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *BackupRequestParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *BackupRequestParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *BackupRequestParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *BackupRequestParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSourceXClusterConfigs

`func (o *BackupRequestParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *BackupRequestParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *BackupRequestParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *BackupRequestParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSse

`func (o *BackupRequestParams) GetSse() bool`

GetSse returns the Sse field if non-nil, zero value otherwise.

### GetSseOk

`func (o *BackupRequestParams) GetSseOk() (*bool, bool)`

GetSseOk returns a tuple with the Sse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSse

`func (o *BackupRequestParams) SetSse(v bool)`

SetSse sets Sse field to given value.

### HasSse

`func (o *BackupRequestParams) HasSse() bool`

HasSse returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *BackupRequestParams) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *BackupRequestParams) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *BackupRequestParams) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetTargetXClusterConfigs

`func (o *BackupRequestParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *BackupRequestParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *BackupRequestParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *BackupRequestParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetTimeBeforeDelete

`func (o *BackupRequestParams) GetTimeBeforeDelete() int64`

GetTimeBeforeDelete returns the TimeBeforeDelete field if non-nil, zero value otherwise.

### GetTimeBeforeDeleteOk

`func (o *BackupRequestParams) GetTimeBeforeDeleteOk() (*int64, bool)`

GetTimeBeforeDeleteOk returns a tuple with the TimeBeforeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBeforeDelete

`func (o *BackupRequestParams) SetTimeBeforeDelete(v int64)`

SetTimeBeforeDelete sets TimeBeforeDelete field to given value.

### HasTimeBeforeDelete

`func (o *BackupRequestParams) HasTimeBeforeDelete() bool`

HasTimeBeforeDelete returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *BackupRequestParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *BackupRequestParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *BackupRequestParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.


### GetUseTablespaces

`func (o *BackupRequestParams) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *BackupRequestParams) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *BackupRequestParams) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.

### HasUseTablespaces

`func (o *BackupRequestParams) HasUseTablespaces() bool`

HasUseTablespaces returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *BackupRequestParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *BackupRequestParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *BackupRequestParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *BackupRequestParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *BackupRequestParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *BackupRequestParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *BackupRequestParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *BackupRequestParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *BackupRequestParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *BackupRequestParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *BackupRequestParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *BackupRequestParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


