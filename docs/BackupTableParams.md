# BackupTableParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** | Action type | [optional] 
**AllTables** | Pointer to **bool** | Backup all tables in Keyspace | [optional] 
**AlterLoadBalancer** | Pointer to **bool** | Alter load balancer state | [optional] 
**BackupList** | Pointer to [**[]BackupTableParams**](BackupTableParams.md) | Backups | [optional] 
**BackupSizeInBytes** | Pointer to **int64** | Backup size in bytes | [optional] 
**BackupType** | Pointer to **string** | Backup type | [optional] 
**BackupUuid** | Pointer to **string** | Backup UUID | [optional] 
**BaseBackupUUID** | Pointer to **string** | Base backup UUID | [optional] 
**CmkArn** | Pointer to **string** | Amazon Resource Name (ARN) of the CMK | [optional] 
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**CreatingUser** | [**Users**](Users.md) |  | 
**CronExpression** | Pointer to **string** | Cron expression for a recurring backup | [optional] 
**CustomerUuid** | Pointer to **string** | Customer UUID | [optional] 
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
**FullChainSizeInBytes** | Pointer to **int64** | Incremental backups chain size | [optional] 
**IgnoreErrors** | Pointer to **bool** | Should table backup errors be ignored | [optional] 
**InstallYbc** | Pointer to **bool** |  | [optional] 
**IsFullBackup** | Pointer to **bool** | Full Table type backup | [optional] 
**Keyspace** | Pointer to **string** | Key space | [optional] 
**KmsConfigUUID** | Pointer to **string** | KMS configuration UUID | [optional] 
**MinNumBackupsToRetain** | Pointer to **int32** | Minimum number of backups to retain for a particular backup schedule | [optional] 
**NewOwner** | Pointer to **string** | User name of the new tables owner | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**OldOwner** | Pointer to **string** | User name of the current tables owner | [optional] 
**Parallelism** | Pointer to **int32** | Number of concurrent commands to run on nodes over SSH | [optional] 
**PlatformUrl** | **string** |  | 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID of a retry | [optional] 
**RegionLocations** | Pointer to [**[]RegionLocations**](RegionLocations.md) | Per region locations | [optional] 
**RestoreTimeStamp** | Pointer to **string** | Restore TimeStamp | [optional] 
**ScheduleUUID** | Pointer to **string** | Schedule UUID | [optional] 
**SchedulingFrequency** | Pointer to **int64** | Frequency to run the backup, in milliseconds | [optional] 
**SleepAfterMasterRestartMillis** | **int32** |  | 
**SleepAfterTServerRestartMillis** | **int32** |  | 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] [readonly] 
**Sse** | Pointer to **bool** | Is SSE | [optional] 
**StorageConfigType** | Pointer to **string** | Type of backup storage config | [optional] 
**StorageConfigUUID** | **string** | Storage configuration UUID | 
**StorageLocation** | Pointer to **string** | Storage location | [optional] 
**TableName** | Pointer to **string** | Table name | [optional] 
**TableNameList** | Pointer to **[]string** | Tables | [optional] 
**TableUUID** | Pointer to **string** | Table UUID | [optional] 
**TableUUIDList** | Pointer to **[]string** | Table UUIDs | [optional] 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] [readonly] 
**TimeBeforeDelete** | Pointer to **int64** | Time before deleting the backup from storage, in milliseconds | [optional] 
**TimeTakenPartial** | **int64** |  | 
**TransactionalBackup** | Pointer to **bool** | Is backup transactional across tables | [optional] 
**UniverseUUID** | Pointer to **string** | Associated universe UUID | [optional] 
**UseTablespaces** | Pointer to **bool** | Is tablespaces information included | [optional] 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 
**YbcInstalled** | Pointer to **bool** |  | [optional] 
**YbcSoftwareVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewBackupTableParams

`func NewBackupTableParams(creatingUser Users, platformUrl string, sleepAfterMasterRestartMillis int32, sleepAfterTServerRestartMillis int32, storageConfigUUID string, timeTakenPartial int64, ) *BackupTableParams`

NewBackupTableParams instantiates a new BackupTableParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupTableParamsWithDefaults

`func NewBackupTableParamsWithDefaults() *BackupTableParams`

NewBackupTableParamsWithDefaults instantiates a new BackupTableParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *BackupTableParams) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *BackupTableParams) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *BackupTableParams) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *BackupTableParams) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetAllTables

`func (o *BackupTableParams) GetAllTables() bool`

GetAllTables returns the AllTables field if non-nil, zero value otherwise.

### GetAllTablesOk

`func (o *BackupTableParams) GetAllTablesOk() (*bool, bool)`

GetAllTablesOk returns a tuple with the AllTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTables

`func (o *BackupTableParams) SetAllTables(v bool)`

SetAllTables sets AllTables field to given value.

### HasAllTables

`func (o *BackupTableParams) HasAllTables() bool`

HasAllTables returns a boolean if a field has been set.

### GetAlterLoadBalancer

`func (o *BackupTableParams) GetAlterLoadBalancer() bool`

GetAlterLoadBalancer returns the AlterLoadBalancer field if non-nil, zero value otherwise.

### GetAlterLoadBalancerOk

`func (o *BackupTableParams) GetAlterLoadBalancerOk() (*bool, bool)`

GetAlterLoadBalancerOk returns a tuple with the AlterLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterLoadBalancer

`func (o *BackupTableParams) SetAlterLoadBalancer(v bool)`

SetAlterLoadBalancer sets AlterLoadBalancer field to given value.

### HasAlterLoadBalancer

`func (o *BackupTableParams) HasAlterLoadBalancer() bool`

HasAlterLoadBalancer returns a boolean if a field has been set.

### GetBackupList

`func (o *BackupTableParams) GetBackupList() []BackupTableParams`

GetBackupList returns the BackupList field if non-nil, zero value otherwise.

### GetBackupListOk

`func (o *BackupTableParams) GetBackupListOk() (*[]BackupTableParams, bool)`

GetBackupListOk returns a tuple with the BackupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupList

`func (o *BackupTableParams) SetBackupList(v []BackupTableParams)`

SetBackupList sets BackupList field to given value.

### HasBackupList

`func (o *BackupTableParams) HasBackupList() bool`

HasBackupList returns a boolean if a field has been set.

### GetBackupSizeInBytes

`func (o *BackupTableParams) GetBackupSizeInBytes() int64`

GetBackupSizeInBytes returns the BackupSizeInBytes field if non-nil, zero value otherwise.

### GetBackupSizeInBytesOk

`func (o *BackupTableParams) GetBackupSizeInBytesOk() (*int64, bool)`

GetBackupSizeInBytesOk returns a tuple with the BackupSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSizeInBytes

`func (o *BackupTableParams) SetBackupSizeInBytes(v int64)`

SetBackupSizeInBytes sets BackupSizeInBytes field to given value.

### HasBackupSizeInBytes

`func (o *BackupTableParams) HasBackupSizeInBytes() bool`

HasBackupSizeInBytes returns a boolean if a field has been set.

### GetBackupType

`func (o *BackupTableParams) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupTableParams) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupTableParams) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *BackupTableParams) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetBackupUuid

`func (o *BackupTableParams) GetBackupUuid() string`

GetBackupUuid returns the BackupUuid field if non-nil, zero value otherwise.

### GetBackupUuidOk

`func (o *BackupTableParams) GetBackupUuidOk() (*string, bool)`

GetBackupUuidOk returns a tuple with the BackupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUuid

`func (o *BackupTableParams) SetBackupUuid(v string)`

SetBackupUuid sets BackupUuid field to given value.

### HasBackupUuid

`func (o *BackupTableParams) HasBackupUuid() bool`

HasBackupUuid returns a boolean if a field has been set.

### GetBaseBackupUUID

`func (o *BackupTableParams) GetBaseBackupUUID() string`

GetBaseBackupUUID returns the BaseBackupUUID field if non-nil, zero value otherwise.

### GetBaseBackupUUIDOk

`func (o *BackupTableParams) GetBaseBackupUUIDOk() (*string, bool)`

GetBaseBackupUUIDOk returns a tuple with the BaseBackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBackupUUID

`func (o *BackupTableParams) SetBaseBackupUUID(v string)`

SetBaseBackupUUID sets BaseBackupUUID field to given value.

### HasBaseBackupUUID

`func (o *BackupTableParams) HasBaseBackupUUID() bool`

HasBaseBackupUUID returns a boolean if a field has been set.

### GetCmkArn

`func (o *BackupTableParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *BackupTableParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *BackupTableParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *BackupTableParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *BackupTableParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *BackupTableParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *BackupTableParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *BackupTableParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *BackupTableParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *BackupTableParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *BackupTableParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCronExpression

`func (o *BackupTableParams) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *BackupTableParams) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *BackupTableParams) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *BackupTableParams) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *BackupTableParams) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *BackupTableParams) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *BackupTableParams) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *BackupTableParams) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *BackupTableParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *BackupTableParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *BackupTableParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *BackupTableParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetDisableChecksum

`func (o *BackupTableParams) GetDisableChecksum() bool`

GetDisableChecksum returns the DisableChecksum field if non-nil, zero value otherwise.

### GetDisableChecksumOk

`func (o *BackupTableParams) GetDisableChecksumOk() (*bool, bool)`

GetDisableChecksumOk returns a tuple with the DisableChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableChecksum

`func (o *BackupTableParams) SetDisableChecksum(v bool)`

SetDisableChecksum sets DisableChecksum field to given value.

### HasDisableChecksum

`func (o *BackupTableParams) HasDisableChecksum() bool`

HasDisableChecksum returns a boolean if a field has been set.

### GetDisableMultipart

`func (o *BackupTableParams) GetDisableMultipart() bool`

GetDisableMultipart returns the DisableMultipart field if non-nil, zero value otherwise.

### GetDisableMultipartOk

`func (o *BackupTableParams) GetDisableMultipartOk() (*bool, bool)`

GetDisableMultipartOk returns a tuple with the DisableMultipart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMultipart

`func (o *BackupTableParams) SetDisableMultipart(v bool)`

SetDisableMultipart sets DisableMultipart field to given value.

### HasDisableMultipart

`func (o *BackupTableParams) HasDisableMultipart() bool`

HasDisableMultipart returns a boolean if a field has been set.

### GetDisableParallelism

`func (o *BackupTableParams) GetDisableParallelism() bool`

GetDisableParallelism returns the DisableParallelism field if non-nil, zero value otherwise.

### GetDisableParallelismOk

`func (o *BackupTableParams) GetDisableParallelismOk() (*bool, bool)`

GetDisableParallelismOk returns a tuple with the DisableParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableParallelism

`func (o *BackupTableParams) SetDisableParallelism(v bool)`

SetDisableParallelism sets DisableParallelism field to given value.

### HasDisableParallelism

`func (o *BackupTableParams) HasDisableParallelism() bool`

HasDisableParallelism returns a boolean if a field has been set.

### GetEnableVerboseLogs

`func (o *BackupTableParams) GetEnableVerboseLogs() bool`

GetEnableVerboseLogs returns the EnableVerboseLogs field if non-nil, zero value otherwise.

### GetEnableVerboseLogsOk

`func (o *BackupTableParams) GetEnableVerboseLogsOk() (*bool, bool)`

GetEnableVerboseLogsOk returns a tuple with the EnableVerboseLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVerboseLogs

`func (o *BackupTableParams) SetEnableVerboseLogs(v bool)`

SetEnableVerboseLogs sets EnableVerboseLogs field to given value.

### HasEnableVerboseLogs

`func (o *BackupTableParams) HasEnableVerboseLogs() bool`

HasEnableVerboseLogs returns a boolean if a field has been set.

### GetEnableYbc

`func (o *BackupTableParams) GetEnableYbc() bool`

GetEnableYbc returns the EnableYbc field if non-nil, zero value otherwise.

### GetEnableYbcOk

`func (o *BackupTableParams) GetEnableYbcOk() (*bool, bool)`

GetEnableYbcOk returns a tuple with the EnableYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYbc

`func (o *BackupTableParams) SetEnableYbc(v bool)`

SetEnableYbc sets EnableYbc field to given value.

### HasEnableYbc

`func (o *BackupTableParams) HasEnableYbc() bool`

HasEnableYbc returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *BackupTableParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *BackupTableParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *BackupTableParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *BackupTableParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *BackupTableParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *BackupTableParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *BackupTableParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *BackupTableParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *BackupTableParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *BackupTableParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *BackupTableParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *BackupTableParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExpiryTimeUnit

`func (o *BackupTableParams) GetExpiryTimeUnit() string`

GetExpiryTimeUnit returns the ExpiryTimeUnit field if non-nil, zero value otherwise.

### GetExpiryTimeUnitOk

`func (o *BackupTableParams) GetExpiryTimeUnitOk() (*string, bool)`

GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUnit

`func (o *BackupTableParams) SetExpiryTimeUnit(v string)`

SetExpiryTimeUnit sets ExpiryTimeUnit field to given value.

### HasExpiryTimeUnit

`func (o *BackupTableParams) HasExpiryTimeUnit() bool`

HasExpiryTimeUnit returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *BackupTableParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *BackupTableParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *BackupTableParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *BackupTableParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetFirstTry

`func (o *BackupTableParams) GetFirstTry() bool`

GetFirstTry returns the FirstTry field if non-nil, zero value otherwise.

### GetFirstTryOk

`func (o *BackupTableParams) GetFirstTryOk() (*bool, bool)`

GetFirstTryOk returns a tuple with the FirstTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTry

`func (o *BackupTableParams) SetFirstTry(v bool)`

SetFirstTry sets FirstTry field to given value.

### HasFirstTry

`func (o *BackupTableParams) HasFirstTry() bool`

HasFirstTry returns a boolean if a field has been set.

### GetFullChainSizeInBytes

`func (o *BackupTableParams) GetFullChainSizeInBytes() int64`

GetFullChainSizeInBytes returns the FullChainSizeInBytes field if non-nil, zero value otherwise.

### GetFullChainSizeInBytesOk

`func (o *BackupTableParams) GetFullChainSizeInBytesOk() (*int64, bool)`

GetFullChainSizeInBytesOk returns a tuple with the FullChainSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullChainSizeInBytes

`func (o *BackupTableParams) SetFullChainSizeInBytes(v int64)`

SetFullChainSizeInBytes sets FullChainSizeInBytes field to given value.

### HasFullChainSizeInBytes

`func (o *BackupTableParams) HasFullChainSizeInBytes() bool`

HasFullChainSizeInBytes returns a boolean if a field has been set.

### GetIgnoreErrors

`func (o *BackupTableParams) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *BackupTableParams) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *BackupTableParams) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.

### HasIgnoreErrors

`func (o *BackupTableParams) HasIgnoreErrors() bool`

HasIgnoreErrors returns a boolean if a field has been set.

### GetInstallYbc

`func (o *BackupTableParams) GetInstallYbc() bool`

GetInstallYbc returns the InstallYbc field if non-nil, zero value otherwise.

### GetInstallYbcOk

`func (o *BackupTableParams) GetInstallYbcOk() (*bool, bool)`

GetInstallYbcOk returns a tuple with the InstallYbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallYbc

`func (o *BackupTableParams) SetInstallYbc(v bool)`

SetInstallYbc sets InstallYbc field to given value.

### HasInstallYbc

`func (o *BackupTableParams) HasInstallYbc() bool`

HasInstallYbc returns a boolean if a field has been set.

### GetIsFullBackup

`func (o *BackupTableParams) GetIsFullBackup() bool`

GetIsFullBackup returns the IsFullBackup field if non-nil, zero value otherwise.

### GetIsFullBackupOk

`func (o *BackupTableParams) GetIsFullBackupOk() (*bool, bool)`

GetIsFullBackupOk returns a tuple with the IsFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullBackup

`func (o *BackupTableParams) SetIsFullBackup(v bool)`

SetIsFullBackup sets IsFullBackup field to given value.

### HasIsFullBackup

`func (o *BackupTableParams) HasIsFullBackup() bool`

HasIsFullBackup returns a boolean if a field has been set.

### GetKeyspace

`func (o *BackupTableParams) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *BackupTableParams) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *BackupTableParams) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *BackupTableParams) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetKmsConfigUUID

`func (o *BackupTableParams) GetKmsConfigUUID() string`

GetKmsConfigUUID returns the KmsConfigUUID field if non-nil, zero value otherwise.

### GetKmsConfigUUIDOk

`func (o *BackupTableParams) GetKmsConfigUUIDOk() (*string, bool)`

GetKmsConfigUUIDOk returns a tuple with the KmsConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUUID

`func (o *BackupTableParams) SetKmsConfigUUID(v string)`

SetKmsConfigUUID sets KmsConfigUUID field to given value.

### HasKmsConfigUUID

`func (o *BackupTableParams) HasKmsConfigUUID() bool`

HasKmsConfigUUID returns a boolean if a field has been set.

### GetMinNumBackupsToRetain

`func (o *BackupTableParams) GetMinNumBackupsToRetain() int32`

GetMinNumBackupsToRetain returns the MinNumBackupsToRetain field if non-nil, zero value otherwise.

### GetMinNumBackupsToRetainOk

`func (o *BackupTableParams) GetMinNumBackupsToRetainOk() (*int32, bool)`

GetMinNumBackupsToRetainOk returns a tuple with the MinNumBackupsToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumBackupsToRetain

`func (o *BackupTableParams) SetMinNumBackupsToRetain(v int32)`

SetMinNumBackupsToRetain sets MinNumBackupsToRetain field to given value.

### HasMinNumBackupsToRetain

`func (o *BackupTableParams) HasMinNumBackupsToRetain() bool`

HasMinNumBackupsToRetain returns a boolean if a field has been set.

### GetNewOwner

`func (o *BackupTableParams) GetNewOwner() string`

GetNewOwner returns the NewOwner field if non-nil, zero value otherwise.

### GetNewOwnerOk

`func (o *BackupTableParams) GetNewOwnerOk() (*string, bool)`

GetNewOwnerOk returns a tuple with the NewOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOwner

`func (o *BackupTableParams) SetNewOwner(v string)`

SetNewOwner sets NewOwner field to given value.

### HasNewOwner

`func (o *BackupTableParams) HasNewOwner() bool`

HasNewOwner returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *BackupTableParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *BackupTableParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *BackupTableParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *BackupTableParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *BackupTableParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *BackupTableParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *BackupTableParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *BackupTableParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetOldOwner

`func (o *BackupTableParams) GetOldOwner() string`

GetOldOwner returns the OldOwner field if non-nil, zero value otherwise.

### GetOldOwnerOk

`func (o *BackupTableParams) GetOldOwnerOk() (*string, bool)`

GetOldOwnerOk returns a tuple with the OldOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldOwner

`func (o *BackupTableParams) SetOldOwner(v string)`

SetOldOwner sets OldOwner field to given value.

### HasOldOwner

`func (o *BackupTableParams) HasOldOwner() bool`

HasOldOwner returns a boolean if a field has been set.

### GetParallelism

`func (o *BackupTableParams) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *BackupTableParams) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *BackupTableParams) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *BackupTableParams) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *BackupTableParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *BackupTableParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *BackupTableParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPreviousTaskUUID

`func (o *BackupTableParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *BackupTableParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *BackupTableParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *BackupTableParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRegionLocations

`func (o *BackupTableParams) GetRegionLocations() []RegionLocations`

GetRegionLocations returns the RegionLocations field if non-nil, zero value otherwise.

### GetRegionLocationsOk

`func (o *BackupTableParams) GetRegionLocationsOk() (*[]RegionLocations, bool)`

GetRegionLocationsOk returns a tuple with the RegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLocations

`func (o *BackupTableParams) SetRegionLocations(v []RegionLocations)`

SetRegionLocations sets RegionLocations field to given value.

### HasRegionLocations

`func (o *BackupTableParams) HasRegionLocations() bool`

HasRegionLocations returns a boolean if a field has been set.

### GetRestoreTimeStamp

`func (o *BackupTableParams) GetRestoreTimeStamp() string`

GetRestoreTimeStamp returns the RestoreTimeStamp field if non-nil, zero value otherwise.

### GetRestoreTimeStampOk

`func (o *BackupTableParams) GetRestoreTimeStampOk() (*string, bool)`

GetRestoreTimeStampOk returns a tuple with the RestoreTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeStamp

`func (o *BackupTableParams) SetRestoreTimeStamp(v string)`

SetRestoreTimeStamp sets RestoreTimeStamp field to given value.

### HasRestoreTimeStamp

`func (o *BackupTableParams) HasRestoreTimeStamp() bool`

HasRestoreTimeStamp returns a boolean if a field has been set.

### GetScheduleUUID

`func (o *BackupTableParams) GetScheduleUUID() string`

GetScheduleUUID returns the ScheduleUUID field if non-nil, zero value otherwise.

### GetScheduleUUIDOk

`func (o *BackupTableParams) GetScheduleUUIDOk() (*string, bool)`

GetScheduleUUIDOk returns a tuple with the ScheduleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUID

`func (o *BackupTableParams) SetScheduleUUID(v string)`

SetScheduleUUID sets ScheduleUUID field to given value.

### HasScheduleUUID

`func (o *BackupTableParams) HasScheduleUUID() bool`

HasScheduleUUID returns a boolean if a field has been set.

### GetSchedulingFrequency

`func (o *BackupTableParams) GetSchedulingFrequency() int64`

GetSchedulingFrequency returns the SchedulingFrequency field if non-nil, zero value otherwise.

### GetSchedulingFrequencyOk

`func (o *BackupTableParams) GetSchedulingFrequencyOk() (*int64, bool)`

GetSchedulingFrequencyOk returns a tuple with the SchedulingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingFrequency

`func (o *BackupTableParams) SetSchedulingFrequency(v int64)`

SetSchedulingFrequency sets SchedulingFrequency field to given value.

### HasSchedulingFrequency

`func (o *BackupTableParams) HasSchedulingFrequency() bool`

HasSchedulingFrequency returns a boolean if a field has been set.

### GetSleepAfterMasterRestartMillis

`func (o *BackupTableParams) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *BackupTableParams) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *BackupTableParams) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.


### GetSleepAfterTServerRestartMillis

`func (o *BackupTableParams) GetSleepAfterTServerRestartMillis() int32`

GetSleepAfterTServerRestartMillis returns the SleepAfterTServerRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTServerRestartMillisOk

`func (o *BackupTableParams) GetSleepAfterTServerRestartMillisOk() (*int32, bool)`

GetSleepAfterTServerRestartMillisOk returns a tuple with the SleepAfterTServerRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTServerRestartMillis

`func (o *BackupTableParams) SetSleepAfterTServerRestartMillis(v int32)`

SetSleepAfterTServerRestartMillis sets SleepAfterTServerRestartMillis field to given value.


### GetSourceXClusterConfigs

`func (o *BackupTableParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *BackupTableParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *BackupTableParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *BackupTableParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSse

`func (o *BackupTableParams) GetSse() bool`

GetSse returns the Sse field if non-nil, zero value otherwise.

### GetSseOk

`func (o *BackupTableParams) GetSseOk() (*bool, bool)`

GetSseOk returns a tuple with the Sse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSse

`func (o *BackupTableParams) SetSse(v bool)`

SetSse sets Sse field to given value.

### HasSse

`func (o *BackupTableParams) HasSse() bool`

HasSse returns a boolean if a field has been set.

### GetStorageConfigType

`func (o *BackupTableParams) GetStorageConfigType() string`

GetStorageConfigType returns the StorageConfigType field if non-nil, zero value otherwise.

### GetStorageConfigTypeOk

`func (o *BackupTableParams) GetStorageConfigTypeOk() (*string, bool)`

GetStorageConfigTypeOk returns a tuple with the StorageConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigType

`func (o *BackupTableParams) SetStorageConfigType(v string)`

SetStorageConfigType sets StorageConfigType field to given value.

### HasStorageConfigType

`func (o *BackupTableParams) HasStorageConfigType() bool`

HasStorageConfigType returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *BackupTableParams) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *BackupTableParams) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *BackupTableParams) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetStorageLocation

`func (o *BackupTableParams) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *BackupTableParams) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *BackupTableParams) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *BackupTableParams) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### GetTableName

`func (o *BackupTableParams) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *BackupTableParams) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *BackupTableParams) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *BackupTableParams) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetTableNameList

`func (o *BackupTableParams) GetTableNameList() []string`

GetTableNameList returns the TableNameList field if non-nil, zero value otherwise.

### GetTableNameListOk

`func (o *BackupTableParams) GetTableNameListOk() (*[]string, bool)`

GetTableNameListOk returns a tuple with the TableNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNameList

`func (o *BackupTableParams) SetTableNameList(v []string)`

SetTableNameList sets TableNameList field to given value.

### HasTableNameList

`func (o *BackupTableParams) HasTableNameList() bool`

HasTableNameList returns a boolean if a field has been set.

### GetTableUUID

`func (o *BackupTableParams) GetTableUUID() string`

GetTableUUID returns the TableUUID field if non-nil, zero value otherwise.

### GetTableUUIDOk

`func (o *BackupTableParams) GetTableUUIDOk() (*string, bool)`

GetTableUUIDOk returns a tuple with the TableUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUID

`func (o *BackupTableParams) SetTableUUID(v string)`

SetTableUUID sets TableUUID field to given value.

### HasTableUUID

`func (o *BackupTableParams) HasTableUUID() bool`

HasTableUUID returns a boolean if a field has been set.

### GetTableUUIDList

`func (o *BackupTableParams) GetTableUUIDList() []string`

GetTableUUIDList returns the TableUUIDList field if non-nil, zero value otherwise.

### GetTableUUIDListOk

`func (o *BackupTableParams) GetTableUUIDListOk() (*[]string, bool)`

GetTableUUIDListOk returns a tuple with the TableUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUIDList

`func (o *BackupTableParams) SetTableUUIDList(v []string)`

SetTableUUIDList sets TableUUIDList field to given value.

### HasTableUUIDList

`func (o *BackupTableParams) HasTableUUIDList() bool`

HasTableUUIDList returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *BackupTableParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *BackupTableParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *BackupTableParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *BackupTableParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetTimeBeforeDelete

`func (o *BackupTableParams) GetTimeBeforeDelete() int64`

GetTimeBeforeDelete returns the TimeBeforeDelete field if non-nil, zero value otherwise.

### GetTimeBeforeDeleteOk

`func (o *BackupTableParams) GetTimeBeforeDeleteOk() (*int64, bool)`

GetTimeBeforeDeleteOk returns a tuple with the TimeBeforeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBeforeDelete

`func (o *BackupTableParams) SetTimeBeforeDelete(v int64)`

SetTimeBeforeDelete sets TimeBeforeDelete field to given value.

### HasTimeBeforeDelete

`func (o *BackupTableParams) HasTimeBeforeDelete() bool`

HasTimeBeforeDelete returns a boolean if a field has been set.

### GetTimeTakenPartial

`func (o *BackupTableParams) GetTimeTakenPartial() int64`

GetTimeTakenPartial returns the TimeTakenPartial field if non-nil, zero value otherwise.

### GetTimeTakenPartialOk

`func (o *BackupTableParams) GetTimeTakenPartialOk() (*int64, bool)`

GetTimeTakenPartialOk returns a tuple with the TimeTakenPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenPartial

`func (o *BackupTableParams) SetTimeTakenPartial(v int64)`

SetTimeTakenPartial sets TimeTakenPartial field to given value.


### GetTransactionalBackup

`func (o *BackupTableParams) GetTransactionalBackup() bool`

GetTransactionalBackup returns the TransactionalBackup field if non-nil, zero value otherwise.

### GetTransactionalBackupOk

`func (o *BackupTableParams) GetTransactionalBackupOk() (*bool, bool)`

GetTransactionalBackupOk returns a tuple with the TransactionalBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionalBackup

`func (o *BackupTableParams) SetTransactionalBackup(v bool)`

SetTransactionalBackup sets TransactionalBackup field to given value.

### HasTransactionalBackup

`func (o *BackupTableParams) HasTransactionalBackup() bool`

HasTransactionalBackup returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *BackupTableParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *BackupTableParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *BackupTableParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *BackupTableParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUseTablespaces

`func (o *BackupTableParams) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *BackupTableParams) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *BackupTableParams) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.

### HasUseTablespaces

`func (o *BackupTableParams) HasUseTablespaces() bool`

HasUseTablespaces returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *BackupTableParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *BackupTableParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *BackupTableParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *BackupTableParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbcInstalled

`func (o *BackupTableParams) GetYbcInstalled() bool`

GetYbcInstalled returns the YbcInstalled field if non-nil, zero value otherwise.

### GetYbcInstalledOk

`func (o *BackupTableParams) GetYbcInstalledOk() (*bool, bool)`

GetYbcInstalledOk returns a tuple with the YbcInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcInstalled

`func (o *BackupTableParams) SetYbcInstalled(v bool)`

SetYbcInstalled sets YbcInstalled field to given value.

### HasYbcInstalled

`func (o *BackupTableParams) HasYbcInstalled() bool`

HasYbcInstalled returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *BackupTableParams) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *BackupTableParams) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *BackupTableParams) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *BackupTableParams) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


