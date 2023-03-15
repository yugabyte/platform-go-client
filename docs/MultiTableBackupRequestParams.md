# MultiTableBackupRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** | Action type | [optional] 
**AlterLoadBalancer** | Pointer to **bool** | Alter load balancer state | [optional] 
**BackupList** | Pointer to [**[]BackupTableParams**](BackupTableParams.md) | Backups | [optional] 
**BackupSizeInBytes** | Pointer to **int64** | Backup size in bytes | [optional] 
**BackupType** | Pointer to **string** | Backup type | [optional] 
**BackupUuid** | Pointer to **string** | Backup UUID | [optional] 
**CmkArn** | Pointer to **string** | Amazon Resource Name (ARN) of the CMK | [optional] 
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**CreatingUser** | [**Users**](Users.md) |  | 
**CronExpression** | Pointer to **string** | Cron expression for a recurring backup | [optional] 
**CustomerUUID** | Pointer to **string** | Customer UUID | [optional] 
**CustomerUuid** | Pointer to **string** | Customer UUID | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**DisableChecksum** | Pointer to **bool** | Disable checksum | [optional] 
**EnableVerboseLogs** | Pointer to **bool** | Is verbose logging enabled | [optional] 
**EncryptionAtRestConfig** | Pointer to [**EncryptionAtRestConfig**](EncryptionAtRestConfig.md) |  | [optional] 
**ErrorString** | Pointer to **string** | Error message | [optional] 
**ExpectedUniverseVersion** | Pointer to **int32** | Expected universe version | [optional] 
**ExpiryTimeUnit** | Pointer to **string** | Time unit for backup expiry time | [optional] 
**ExtraDependencies** | Pointer to [**ExtraDependencies**](ExtraDependencies.md) |  | [optional] 
**FirstTry** | Pointer to **bool** | Whether this task has been tried before | [optional] 
**IgnoreErrors** | Pointer to **bool** | Should table backup errors be ignored | [optional] 
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
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID only if this task is a retry | [optional] 
**RegionLocations** | Pointer to [**[]RegionLocations**](RegionLocations.md) | Per region locations | [optional] 
**RestoreTimeStamp** | Pointer to **string** | Restore TimeStamp | [optional] 
**ScheduleUUID** | Pointer to **string** | Schedule UUID | [optional] 
**SchedulingFrequency** | Pointer to **int64** | Frequency to run the backup, in milliseconds | [optional] 
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
**TransactionalBackup** | Pointer to **bool** | Is backup transactional across tables | [optional] 
**UniverseUUID** | Pointer to **string** | Associated universe UUID | [optional] 
**UseTablespaces** | Pointer to **bool** | Is tablespaces information included | [optional] 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 

## Methods

### NewMultiTableBackupRequestParams

`func NewMultiTableBackupRequestParams(creatingUser Users, platformUrl string, storageConfigUUID string, ) *MultiTableBackupRequestParams`

NewMultiTableBackupRequestParams instantiates a new MultiTableBackupRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiTableBackupRequestParamsWithDefaults

`func NewMultiTableBackupRequestParamsWithDefaults() *MultiTableBackupRequestParams`

NewMultiTableBackupRequestParamsWithDefaults instantiates a new MultiTableBackupRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *MultiTableBackupRequestParams) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *MultiTableBackupRequestParams) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *MultiTableBackupRequestParams) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *MultiTableBackupRequestParams) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetAlterLoadBalancer

`func (o *MultiTableBackupRequestParams) GetAlterLoadBalancer() bool`

GetAlterLoadBalancer returns the AlterLoadBalancer field if non-nil, zero value otherwise.

### GetAlterLoadBalancerOk

`func (o *MultiTableBackupRequestParams) GetAlterLoadBalancerOk() (*bool, bool)`

GetAlterLoadBalancerOk returns a tuple with the AlterLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterLoadBalancer

`func (o *MultiTableBackupRequestParams) SetAlterLoadBalancer(v bool)`

SetAlterLoadBalancer sets AlterLoadBalancer field to given value.

### HasAlterLoadBalancer

`func (o *MultiTableBackupRequestParams) HasAlterLoadBalancer() bool`

HasAlterLoadBalancer returns a boolean if a field has been set.

### GetBackupList

`func (o *MultiTableBackupRequestParams) GetBackupList() []BackupTableParams`

GetBackupList returns the BackupList field if non-nil, zero value otherwise.

### GetBackupListOk

`func (o *MultiTableBackupRequestParams) GetBackupListOk() (*[]BackupTableParams, bool)`

GetBackupListOk returns a tuple with the BackupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupList

`func (o *MultiTableBackupRequestParams) SetBackupList(v []BackupTableParams)`

SetBackupList sets BackupList field to given value.

### HasBackupList

`func (o *MultiTableBackupRequestParams) HasBackupList() bool`

HasBackupList returns a boolean if a field has been set.

### GetBackupSizeInBytes

`func (o *MultiTableBackupRequestParams) GetBackupSizeInBytes() int64`

GetBackupSizeInBytes returns the BackupSizeInBytes field if non-nil, zero value otherwise.

### GetBackupSizeInBytesOk

`func (o *MultiTableBackupRequestParams) GetBackupSizeInBytesOk() (*int64, bool)`

GetBackupSizeInBytesOk returns a tuple with the BackupSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSizeInBytes

`func (o *MultiTableBackupRequestParams) SetBackupSizeInBytes(v int64)`

SetBackupSizeInBytes sets BackupSizeInBytes field to given value.

### HasBackupSizeInBytes

`func (o *MultiTableBackupRequestParams) HasBackupSizeInBytes() bool`

HasBackupSizeInBytes returns a boolean if a field has been set.

### GetBackupType

`func (o *MultiTableBackupRequestParams) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *MultiTableBackupRequestParams) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *MultiTableBackupRequestParams) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *MultiTableBackupRequestParams) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetBackupUuid

`func (o *MultiTableBackupRequestParams) GetBackupUuid() string`

GetBackupUuid returns the BackupUuid field if non-nil, zero value otherwise.

### GetBackupUuidOk

`func (o *MultiTableBackupRequestParams) GetBackupUuidOk() (*string, bool)`

GetBackupUuidOk returns a tuple with the BackupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUuid

`func (o *MultiTableBackupRequestParams) SetBackupUuid(v string)`

SetBackupUuid sets BackupUuid field to given value.

### HasBackupUuid

`func (o *MultiTableBackupRequestParams) HasBackupUuid() bool`

HasBackupUuid returns a boolean if a field has been set.

### GetCmkArn

`func (o *MultiTableBackupRequestParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *MultiTableBackupRequestParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *MultiTableBackupRequestParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *MultiTableBackupRequestParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *MultiTableBackupRequestParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *MultiTableBackupRequestParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *MultiTableBackupRequestParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *MultiTableBackupRequestParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetCreatingUser

`func (o *MultiTableBackupRequestParams) GetCreatingUser() Users`

GetCreatingUser returns the CreatingUser field if non-nil, zero value otherwise.

### GetCreatingUserOk

`func (o *MultiTableBackupRequestParams) GetCreatingUserOk() (*Users, bool)`

GetCreatingUserOk returns a tuple with the CreatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingUser

`func (o *MultiTableBackupRequestParams) SetCreatingUser(v Users)`

SetCreatingUser sets CreatingUser field to given value.


### GetCronExpression

`func (o *MultiTableBackupRequestParams) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *MultiTableBackupRequestParams) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *MultiTableBackupRequestParams) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *MultiTableBackupRequestParams) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *MultiTableBackupRequestParams) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *MultiTableBackupRequestParams) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *MultiTableBackupRequestParams) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *MultiTableBackupRequestParams) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *MultiTableBackupRequestParams) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *MultiTableBackupRequestParams) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *MultiTableBackupRequestParams) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *MultiTableBackupRequestParams) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *MultiTableBackupRequestParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *MultiTableBackupRequestParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *MultiTableBackupRequestParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *MultiTableBackupRequestParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetDisableChecksum

`func (o *MultiTableBackupRequestParams) GetDisableChecksum() bool`

GetDisableChecksum returns the DisableChecksum field if non-nil, zero value otherwise.

### GetDisableChecksumOk

`func (o *MultiTableBackupRequestParams) GetDisableChecksumOk() (*bool, bool)`

GetDisableChecksumOk returns a tuple with the DisableChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableChecksum

`func (o *MultiTableBackupRequestParams) SetDisableChecksum(v bool)`

SetDisableChecksum sets DisableChecksum field to given value.

### HasDisableChecksum

`func (o *MultiTableBackupRequestParams) HasDisableChecksum() bool`

HasDisableChecksum returns a boolean if a field has been set.

### GetEnableVerboseLogs

`func (o *MultiTableBackupRequestParams) GetEnableVerboseLogs() bool`

GetEnableVerboseLogs returns the EnableVerboseLogs field if non-nil, zero value otherwise.

### GetEnableVerboseLogsOk

`func (o *MultiTableBackupRequestParams) GetEnableVerboseLogsOk() (*bool, bool)`

GetEnableVerboseLogsOk returns a tuple with the EnableVerboseLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVerboseLogs

`func (o *MultiTableBackupRequestParams) SetEnableVerboseLogs(v bool)`

SetEnableVerboseLogs sets EnableVerboseLogs field to given value.

### HasEnableVerboseLogs

`func (o *MultiTableBackupRequestParams) HasEnableVerboseLogs() bool`

HasEnableVerboseLogs returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *MultiTableBackupRequestParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *MultiTableBackupRequestParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *MultiTableBackupRequestParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *MultiTableBackupRequestParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *MultiTableBackupRequestParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *MultiTableBackupRequestParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *MultiTableBackupRequestParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *MultiTableBackupRequestParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *MultiTableBackupRequestParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *MultiTableBackupRequestParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *MultiTableBackupRequestParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *MultiTableBackupRequestParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExpiryTimeUnit

`func (o *MultiTableBackupRequestParams) GetExpiryTimeUnit() string`

GetExpiryTimeUnit returns the ExpiryTimeUnit field if non-nil, zero value otherwise.

### GetExpiryTimeUnitOk

`func (o *MultiTableBackupRequestParams) GetExpiryTimeUnitOk() (*string, bool)`

GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUnit

`func (o *MultiTableBackupRequestParams) SetExpiryTimeUnit(v string)`

SetExpiryTimeUnit sets ExpiryTimeUnit field to given value.

### HasExpiryTimeUnit

`func (o *MultiTableBackupRequestParams) HasExpiryTimeUnit() bool`

HasExpiryTimeUnit returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *MultiTableBackupRequestParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *MultiTableBackupRequestParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *MultiTableBackupRequestParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *MultiTableBackupRequestParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetFirstTry

`func (o *MultiTableBackupRequestParams) GetFirstTry() bool`

GetFirstTry returns the FirstTry field if non-nil, zero value otherwise.

### GetFirstTryOk

`func (o *MultiTableBackupRequestParams) GetFirstTryOk() (*bool, bool)`

GetFirstTryOk returns a tuple with the FirstTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTry

`func (o *MultiTableBackupRequestParams) SetFirstTry(v bool)`

SetFirstTry sets FirstTry field to given value.

### HasFirstTry

`func (o *MultiTableBackupRequestParams) HasFirstTry() bool`

HasFirstTry returns a boolean if a field has been set.

### GetIgnoreErrors

`func (o *MultiTableBackupRequestParams) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *MultiTableBackupRequestParams) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *MultiTableBackupRequestParams) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.

### HasIgnoreErrors

`func (o *MultiTableBackupRequestParams) HasIgnoreErrors() bool`

HasIgnoreErrors returns a boolean if a field has been set.

### GetIsFullBackup

`func (o *MultiTableBackupRequestParams) GetIsFullBackup() bool`

GetIsFullBackup returns the IsFullBackup field if non-nil, zero value otherwise.

### GetIsFullBackupOk

`func (o *MultiTableBackupRequestParams) GetIsFullBackupOk() (*bool, bool)`

GetIsFullBackupOk returns a tuple with the IsFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullBackup

`func (o *MultiTableBackupRequestParams) SetIsFullBackup(v bool)`

SetIsFullBackup sets IsFullBackup field to given value.

### HasIsFullBackup

`func (o *MultiTableBackupRequestParams) HasIsFullBackup() bool`

HasIsFullBackup returns a boolean if a field has been set.

### GetKeyspace

`func (o *MultiTableBackupRequestParams) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *MultiTableBackupRequestParams) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *MultiTableBackupRequestParams) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *MultiTableBackupRequestParams) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetKmsConfigUUID

`func (o *MultiTableBackupRequestParams) GetKmsConfigUUID() string`

GetKmsConfigUUID returns the KmsConfigUUID field if non-nil, zero value otherwise.

### GetKmsConfigUUIDOk

`func (o *MultiTableBackupRequestParams) GetKmsConfigUUIDOk() (*string, bool)`

GetKmsConfigUUIDOk returns a tuple with the KmsConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUUID

`func (o *MultiTableBackupRequestParams) SetKmsConfigUUID(v string)`

SetKmsConfigUUID sets KmsConfigUUID field to given value.

### HasKmsConfigUUID

`func (o *MultiTableBackupRequestParams) HasKmsConfigUUID() bool`

HasKmsConfigUUID returns a boolean if a field has been set.

### GetMinNumBackupsToRetain

`func (o *MultiTableBackupRequestParams) GetMinNumBackupsToRetain() int32`

GetMinNumBackupsToRetain returns the MinNumBackupsToRetain field if non-nil, zero value otherwise.

### GetMinNumBackupsToRetainOk

`func (o *MultiTableBackupRequestParams) GetMinNumBackupsToRetainOk() (*int32, bool)`

GetMinNumBackupsToRetainOk returns a tuple with the MinNumBackupsToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumBackupsToRetain

`func (o *MultiTableBackupRequestParams) SetMinNumBackupsToRetain(v int32)`

SetMinNumBackupsToRetain sets MinNumBackupsToRetain field to given value.

### HasMinNumBackupsToRetain

`func (o *MultiTableBackupRequestParams) HasMinNumBackupsToRetain() bool`

HasMinNumBackupsToRetain returns a boolean if a field has been set.

### GetNewOwner

`func (o *MultiTableBackupRequestParams) GetNewOwner() string`

GetNewOwner returns the NewOwner field if non-nil, zero value otherwise.

### GetNewOwnerOk

`func (o *MultiTableBackupRequestParams) GetNewOwnerOk() (*string, bool)`

GetNewOwnerOk returns a tuple with the NewOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOwner

`func (o *MultiTableBackupRequestParams) SetNewOwner(v string)`

SetNewOwner sets NewOwner field to given value.

### HasNewOwner

`func (o *MultiTableBackupRequestParams) HasNewOwner() bool`

HasNewOwner returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *MultiTableBackupRequestParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *MultiTableBackupRequestParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *MultiTableBackupRequestParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *MultiTableBackupRequestParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *MultiTableBackupRequestParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *MultiTableBackupRequestParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *MultiTableBackupRequestParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *MultiTableBackupRequestParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetOldOwner

`func (o *MultiTableBackupRequestParams) GetOldOwner() string`

GetOldOwner returns the OldOwner field if non-nil, zero value otherwise.

### GetOldOwnerOk

`func (o *MultiTableBackupRequestParams) GetOldOwnerOk() (*string, bool)`

GetOldOwnerOk returns a tuple with the OldOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldOwner

`func (o *MultiTableBackupRequestParams) SetOldOwner(v string)`

SetOldOwner sets OldOwner field to given value.

### HasOldOwner

`func (o *MultiTableBackupRequestParams) HasOldOwner() bool`

HasOldOwner returns a boolean if a field has been set.

### GetParallelism

`func (o *MultiTableBackupRequestParams) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *MultiTableBackupRequestParams) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *MultiTableBackupRequestParams) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *MultiTableBackupRequestParams) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *MultiTableBackupRequestParams) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *MultiTableBackupRequestParams) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *MultiTableBackupRequestParams) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.


### GetPreviousTaskUUID

`func (o *MultiTableBackupRequestParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *MultiTableBackupRequestParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *MultiTableBackupRequestParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *MultiTableBackupRequestParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRegionLocations

`func (o *MultiTableBackupRequestParams) GetRegionLocations() []RegionLocations`

GetRegionLocations returns the RegionLocations field if non-nil, zero value otherwise.

### GetRegionLocationsOk

`func (o *MultiTableBackupRequestParams) GetRegionLocationsOk() (*[]RegionLocations, bool)`

GetRegionLocationsOk returns a tuple with the RegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLocations

`func (o *MultiTableBackupRequestParams) SetRegionLocations(v []RegionLocations)`

SetRegionLocations sets RegionLocations field to given value.

### HasRegionLocations

`func (o *MultiTableBackupRequestParams) HasRegionLocations() bool`

HasRegionLocations returns a boolean if a field has been set.

### GetRestoreTimeStamp

`func (o *MultiTableBackupRequestParams) GetRestoreTimeStamp() string`

GetRestoreTimeStamp returns the RestoreTimeStamp field if non-nil, zero value otherwise.

### GetRestoreTimeStampOk

`func (o *MultiTableBackupRequestParams) GetRestoreTimeStampOk() (*string, bool)`

GetRestoreTimeStampOk returns a tuple with the RestoreTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeStamp

`func (o *MultiTableBackupRequestParams) SetRestoreTimeStamp(v string)`

SetRestoreTimeStamp sets RestoreTimeStamp field to given value.

### HasRestoreTimeStamp

`func (o *MultiTableBackupRequestParams) HasRestoreTimeStamp() bool`

HasRestoreTimeStamp returns a boolean if a field has been set.

### GetScheduleUUID

`func (o *MultiTableBackupRequestParams) GetScheduleUUID() string`

GetScheduleUUID returns the ScheduleUUID field if non-nil, zero value otherwise.

### GetScheduleUUIDOk

`func (o *MultiTableBackupRequestParams) GetScheduleUUIDOk() (*string, bool)`

GetScheduleUUIDOk returns a tuple with the ScheduleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUID

`func (o *MultiTableBackupRequestParams) SetScheduleUUID(v string)`

SetScheduleUUID sets ScheduleUUID field to given value.

### HasScheduleUUID

`func (o *MultiTableBackupRequestParams) HasScheduleUUID() bool`

HasScheduleUUID returns a boolean if a field has been set.

### GetSchedulingFrequency

`func (o *MultiTableBackupRequestParams) GetSchedulingFrequency() int64`

GetSchedulingFrequency returns the SchedulingFrequency field if non-nil, zero value otherwise.

### GetSchedulingFrequencyOk

`func (o *MultiTableBackupRequestParams) GetSchedulingFrequencyOk() (*int64, bool)`

GetSchedulingFrequencyOk returns a tuple with the SchedulingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingFrequency

`func (o *MultiTableBackupRequestParams) SetSchedulingFrequency(v int64)`

SetSchedulingFrequency sets SchedulingFrequency field to given value.

### HasSchedulingFrequency

`func (o *MultiTableBackupRequestParams) HasSchedulingFrequency() bool`

HasSchedulingFrequency returns a boolean if a field has been set.

### GetSourceXClusterConfigs

`func (o *MultiTableBackupRequestParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *MultiTableBackupRequestParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *MultiTableBackupRequestParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *MultiTableBackupRequestParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSse

`func (o *MultiTableBackupRequestParams) GetSse() bool`

GetSse returns the Sse field if non-nil, zero value otherwise.

### GetSseOk

`func (o *MultiTableBackupRequestParams) GetSseOk() (*bool, bool)`

GetSseOk returns a tuple with the Sse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSse

`func (o *MultiTableBackupRequestParams) SetSse(v bool)`

SetSse sets Sse field to given value.

### HasSse

`func (o *MultiTableBackupRequestParams) HasSse() bool`

HasSse returns a boolean if a field has been set.

### GetStorageConfigType

`func (o *MultiTableBackupRequestParams) GetStorageConfigType() string`

GetStorageConfigType returns the StorageConfigType field if non-nil, zero value otherwise.

### GetStorageConfigTypeOk

`func (o *MultiTableBackupRequestParams) GetStorageConfigTypeOk() (*string, bool)`

GetStorageConfigTypeOk returns a tuple with the StorageConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigType

`func (o *MultiTableBackupRequestParams) SetStorageConfigType(v string)`

SetStorageConfigType sets StorageConfigType field to given value.

### HasStorageConfigType

`func (o *MultiTableBackupRequestParams) HasStorageConfigType() bool`

HasStorageConfigType returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *MultiTableBackupRequestParams) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *MultiTableBackupRequestParams) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *MultiTableBackupRequestParams) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetStorageLocation

`func (o *MultiTableBackupRequestParams) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *MultiTableBackupRequestParams) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *MultiTableBackupRequestParams) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *MultiTableBackupRequestParams) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### GetTableName

`func (o *MultiTableBackupRequestParams) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *MultiTableBackupRequestParams) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *MultiTableBackupRequestParams) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *MultiTableBackupRequestParams) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetTableNameList

`func (o *MultiTableBackupRequestParams) GetTableNameList() []string`

GetTableNameList returns the TableNameList field if non-nil, zero value otherwise.

### GetTableNameListOk

`func (o *MultiTableBackupRequestParams) GetTableNameListOk() (*[]string, bool)`

GetTableNameListOk returns a tuple with the TableNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNameList

`func (o *MultiTableBackupRequestParams) SetTableNameList(v []string)`

SetTableNameList sets TableNameList field to given value.

### HasTableNameList

`func (o *MultiTableBackupRequestParams) HasTableNameList() bool`

HasTableNameList returns a boolean if a field has been set.

### GetTableUUID

`func (o *MultiTableBackupRequestParams) GetTableUUID() string`

GetTableUUID returns the TableUUID field if non-nil, zero value otherwise.

### GetTableUUIDOk

`func (o *MultiTableBackupRequestParams) GetTableUUIDOk() (*string, bool)`

GetTableUUIDOk returns a tuple with the TableUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUID

`func (o *MultiTableBackupRequestParams) SetTableUUID(v string)`

SetTableUUID sets TableUUID field to given value.

### HasTableUUID

`func (o *MultiTableBackupRequestParams) HasTableUUID() bool`

HasTableUUID returns a boolean if a field has been set.

### GetTableUUIDList

`func (o *MultiTableBackupRequestParams) GetTableUUIDList() []string`

GetTableUUIDList returns the TableUUIDList field if non-nil, zero value otherwise.

### GetTableUUIDListOk

`func (o *MultiTableBackupRequestParams) GetTableUUIDListOk() (*[]string, bool)`

GetTableUUIDListOk returns a tuple with the TableUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUIDList

`func (o *MultiTableBackupRequestParams) SetTableUUIDList(v []string)`

SetTableUUIDList sets TableUUIDList field to given value.

### HasTableUUIDList

`func (o *MultiTableBackupRequestParams) HasTableUUIDList() bool`

HasTableUUIDList returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *MultiTableBackupRequestParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *MultiTableBackupRequestParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *MultiTableBackupRequestParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *MultiTableBackupRequestParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetTimeBeforeDelete

`func (o *MultiTableBackupRequestParams) GetTimeBeforeDelete() int64`

GetTimeBeforeDelete returns the TimeBeforeDelete field if non-nil, zero value otherwise.

### GetTimeBeforeDeleteOk

`func (o *MultiTableBackupRequestParams) GetTimeBeforeDeleteOk() (*int64, bool)`

GetTimeBeforeDeleteOk returns a tuple with the TimeBeforeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBeforeDelete

`func (o *MultiTableBackupRequestParams) SetTimeBeforeDelete(v int64)`

SetTimeBeforeDelete sets TimeBeforeDelete field to given value.

### HasTimeBeforeDelete

`func (o *MultiTableBackupRequestParams) HasTimeBeforeDelete() bool`

HasTimeBeforeDelete returns a boolean if a field has been set.

### GetTransactionalBackup

`func (o *MultiTableBackupRequestParams) GetTransactionalBackup() bool`

GetTransactionalBackup returns the TransactionalBackup field if non-nil, zero value otherwise.

### GetTransactionalBackupOk

`func (o *MultiTableBackupRequestParams) GetTransactionalBackupOk() (*bool, bool)`

GetTransactionalBackupOk returns a tuple with the TransactionalBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionalBackup

`func (o *MultiTableBackupRequestParams) SetTransactionalBackup(v bool)`

SetTransactionalBackup sets TransactionalBackup field to given value.

### HasTransactionalBackup

`func (o *MultiTableBackupRequestParams) HasTransactionalBackup() bool`

HasTransactionalBackup returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *MultiTableBackupRequestParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *MultiTableBackupRequestParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *MultiTableBackupRequestParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *MultiTableBackupRequestParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUseTablespaces

`func (o *MultiTableBackupRequestParams) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *MultiTableBackupRequestParams) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *MultiTableBackupRequestParams) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.

### HasUseTablespaces

`func (o *MultiTableBackupRequestParams) HasUseTablespaces() bool`

HasUseTablespaces returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *MultiTableBackupRequestParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *MultiTableBackupRequestParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *MultiTableBackupRequestParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *MultiTableBackupRequestParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


