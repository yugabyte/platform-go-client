# MultiTableBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** | Action type | [optional] 
**BackupList** | Pointer to [**[]BackupTableParams**](BackupTableParams.md) | Backups | [optional] 
**BackupType** | Pointer to **string** | Backup type | [optional] 
**BackupUuid** | Pointer to **string** | Backup UUID | [optional] 
**CmkArn** | Pointer to **string** | Amazon Resource Name (ARN) of the CMK | [optional] 
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**Controller** | Pointer to **string** | Controller type | [optional] 
**CronExpression** | Pointer to **string** | Cron expression for a recurring backup | [optional] 
**CustomerUUID** | **string** |  | 
**CustomerUuid** | Pointer to **string** | Customer UUID | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**EnableVerboseLogs** | Pointer to **bool** | Is verbose logging enabled | [optional] 
**EncryptionAtRestConfig** | Pointer to [**EncryptionAtRestConfig**](EncryptionAtRestConfig.md) |  | [optional] 
**ErrorString** | Pointer to **string** | Error message | [optional] 
**ExpectedUniverseVersion** | Pointer to **int32** | Expected universe version | [optional] 
**ExtraDependencies** | Pointer to [**ExtraDependencies**](ExtraDependencies.md) |  | [optional] 
**FirstTry** | Pointer to **bool** | Whether this task has been tried before | [optional] 
**IgnoreErrors** | Pointer to **bool** | Should table backup errors be ignored | [optional] 
**Keyspace** | Pointer to **string** | Key space | [optional] 
**KmsConfigUUID** | Pointer to **string** | KMS configuration UUID | [optional] 
**MinNumBackupsToRetain** | Pointer to **int32** | Minimum number of backups to retain for a particular backup schedule | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**Parallelism** | Pointer to **int32** | Number of concurrent commands to run on nodes over SSH | [optional] 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID only if this task is a retry | [optional] 
**RestoreTimeStamp** | Pointer to **string** | Restore TimeStamp | [optional] 
**ScheduleUUID** | Pointer to **string** | Schedule UUID | [optional] 
**SchedulingFrequency** | Pointer to **int64** | Frequency to run the backup, in milliseconds | [optional] 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] [readonly] 
**Sse** | Pointer to **bool** | Is SSE | [optional] 
**StorageConfigUUID** | **string** | Storage configuration UUID | 
**StorageLocation** | Pointer to **string** | Storage location | [optional] 
**TableName** | Pointer to **string** | Table name | [optional] 
**TableNameList** | Pointer to **[]string** | Tables | [optional] 
**TableUUID** | Pointer to **string** | Table UUID | [optional] 
**TableUUIDList** | **[]string** |  | 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] [readonly] 
**TimeBeforeDelete** | Pointer to **int64** | Time before deleting the backup from storage, in milliseconds | [optional] 
**TransactionalBackup** | Pointer to **bool** | Is backup transactional across tables | [optional] 
**UniverseUUID** | Pointer to **string** | Associated universe UUID | [optional] 
**UseTablespaces** | Pointer to **bool** | Is tablespaces information included | [optional] 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 

## Methods

### NewMultiTableBackupParams

`func NewMultiTableBackupParams(customerUUID string, storageConfigUUID string, tableUUIDList []string, ) *MultiTableBackupParams`

NewMultiTableBackupParams instantiates a new MultiTableBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiTableBackupParamsWithDefaults

`func NewMultiTableBackupParamsWithDefaults() *MultiTableBackupParams`

NewMultiTableBackupParamsWithDefaults instantiates a new MultiTableBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *MultiTableBackupParams) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *MultiTableBackupParams) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *MultiTableBackupParams) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *MultiTableBackupParams) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetBackupList

`func (o *MultiTableBackupParams) GetBackupList() []BackupTableParams`

GetBackupList returns the BackupList field if non-nil, zero value otherwise.

### GetBackupListOk

`func (o *MultiTableBackupParams) GetBackupListOk() (*[]BackupTableParams, bool)`

GetBackupListOk returns a tuple with the BackupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupList

`func (o *MultiTableBackupParams) SetBackupList(v []BackupTableParams)`

SetBackupList sets BackupList field to given value.

### HasBackupList

`func (o *MultiTableBackupParams) HasBackupList() bool`

HasBackupList returns a boolean if a field has been set.

### GetBackupType

`func (o *MultiTableBackupParams) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *MultiTableBackupParams) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *MultiTableBackupParams) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *MultiTableBackupParams) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetBackupUuid

`func (o *MultiTableBackupParams) GetBackupUuid() string`

GetBackupUuid returns the BackupUuid field if non-nil, zero value otherwise.

### GetBackupUuidOk

`func (o *MultiTableBackupParams) GetBackupUuidOk() (*string, bool)`

GetBackupUuidOk returns a tuple with the BackupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUuid

`func (o *MultiTableBackupParams) SetBackupUuid(v string)`

SetBackupUuid sets BackupUuid field to given value.

### HasBackupUuid

`func (o *MultiTableBackupParams) HasBackupUuid() bool`

HasBackupUuid returns a boolean if a field has been set.

### GetCmkArn

`func (o *MultiTableBackupParams) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *MultiTableBackupParams) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *MultiTableBackupParams) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *MultiTableBackupParams) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### GetCommunicationPorts

`func (o *MultiTableBackupParams) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *MultiTableBackupParams) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *MultiTableBackupParams) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *MultiTableBackupParams) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetController

`func (o *MultiTableBackupParams) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *MultiTableBackupParams) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *MultiTableBackupParams) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *MultiTableBackupParams) HasController() bool`

HasController returns a boolean if a field has been set.

### GetCronExpression

`func (o *MultiTableBackupParams) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *MultiTableBackupParams) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *MultiTableBackupParams) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *MultiTableBackupParams) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *MultiTableBackupParams) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *MultiTableBackupParams) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *MultiTableBackupParams) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetCustomerUuid

`func (o *MultiTableBackupParams) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *MultiTableBackupParams) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *MultiTableBackupParams) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *MultiTableBackupParams) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *MultiTableBackupParams) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *MultiTableBackupParams) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *MultiTableBackupParams) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *MultiTableBackupParams) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableVerboseLogs

`func (o *MultiTableBackupParams) GetEnableVerboseLogs() bool`

GetEnableVerboseLogs returns the EnableVerboseLogs field if non-nil, zero value otherwise.

### GetEnableVerboseLogsOk

`func (o *MultiTableBackupParams) GetEnableVerboseLogsOk() (*bool, bool)`

GetEnableVerboseLogsOk returns a tuple with the EnableVerboseLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVerboseLogs

`func (o *MultiTableBackupParams) SetEnableVerboseLogs(v bool)`

SetEnableVerboseLogs sets EnableVerboseLogs field to given value.

### HasEnableVerboseLogs

`func (o *MultiTableBackupParams) HasEnableVerboseLogs() bool`

HasEnableVerboseLogs returns a boolean if a field has been set.

### GetEncryptionAtRestConfig

`func (o *MultiTableBackupParams) GetEncryptionAtRestConfig() EncryptionAtRestConfig`

GetEncryptionAtRestConfig returns the EncryptionAtRestConfig field if non-nil, zero value otherwise.

### GetEncryptionAtRestConfigOk

`func (o *MultiTableBackupParams) GetEncryptionAtRestConfigOk() (*EncryptionAtRestConfig, bool)`

GetEncryptionAtRestConfigOk returns a tuple with the EncryptionAtRestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestConfig

`func (o *MultiTableBackupParams) SetEncryptionAtRestConfig(v EncryptionAtRestConfig)`

SetEncryptionAtRestConfig sets EncryptionAtRestConfig field to given value.

### HasEncryptionAtRestConfig

`func (o *MultiTableBackupParams) HasEncryptionAtRestConfig() bool`

HasEncryptionAtRestConfig returns a boolean if a field has been set.

### GetErrorString

`func (o *MultiTableBackupParams) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *MultiTableBackupParams) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *MultiTableBackupParams) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *MultiTableBackupParams) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *MultiTableBackupParams) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *MultiTableBackupParams) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *MultiTableBackupParams) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.

### HasExpectedUniverseVersion

`func (o *MultiTableBackupParams) HasExpectedUniverseVersion() bool`

HasExpectedUniverseVersion returns a boolean if a field has been set.

### GetExtraDependencies

`func (o *MultiTableBackupParams) GetExtraDependencies() ExtraDependencies`

GetExtraDependencies returns the ExtraDependencies field if non-nil, zero value otherwise.

### GetExtraDependenciesOk

`func (o *MultiTableBackupParams) GetExtraDependenciesOk() (*ExtraDependencies, bool)`

GetExtraDependenciesOk returns a tuple with the ExtraDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDependencies

`func (o *MultiTableBackupParams) SetExtraDependencies(v ExtraDependencies)`

SetExtraDependencies sets ExtraDependencies field to given value.

### HasExtraDependencies

`func (o *MultiTableBackupParams) HasExtraDependencies() bool`

HasExtraDependencies returns a boolean if a field has been set.

### GetFirstTry

`func (o *MultiTableBackupParams) GetFirstTry() bool`

GetFirstTry returns the FirstTry field if non-nil, zero value otherwise.

### GetFirstTryOk

`func (o *MultiTableBackupParams) GetFirstTryOk() (*bool, bool)`

GetFirstTryOk returns a tuple with the FirstTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTry

`func (o *MultiTableBackupParams) SetFirstTry(v bool)`

SetFirstTry sets FirstTry field to given value.

### HasFirstTry

`func (o *MultiTableBackupParams) HasFirstTry() bool`

HasFirstTry returns a boolean if a field has been set.

### GetIgnoreErrors

`func (o *MultiTableBackupParams) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *MultiTableBackupParams) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *MultiTableBackupParams) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.

### HasIgnoreErrors

`func (o *MultiTableBackupParams) HasIgnoreErrors() bool`

HasIgnoreErrors returns a boolean if a field has been set.

### GetKeyspace

`func (o *MultiTableBackupParams) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *MultiTableBackupParams) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *MultiTableBackupParams) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *MultiTableBackupParams) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetKmsConfigUUID

`func (o *MultiTableBackupParams) GetKmsConfigUUID() string`

GetKmsConfigUUID returns the KmsConfigUUID field if non-nil, zero value otherwise.

### GetKmsConfigUUIDOk

`func (o *MultiTableBackupParams) GetKmsConfigUUIDOk() (*string, bool)`

GetKmsConfigUUIDOk returns a tuple with the KmsConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUUID

`func (o *MultiTableBackupParams) SetKmsConfigUUID(v string)`

SetKmsConfigUUID sets KmsConfigUUID field to given value.

### HasKmsConfigUUID

`func (o *MultiTableBackupParams) HasKmsConfigUUID() bool`

HasKmsConfigUUID returns a boolean if a field has been set.

### GetMinNumBackupsToRetain

`func (o *MultiTableBackupParams) GetMinNumBackupsToRetain() int32`

GetMinNumBackupsToRetain returns the MinNumBackupsToRetain field if non-nil, zero value otherwise.

### GetMinNumBackupsToRetainOk

`func (o *MultiTableBackupParams) GetMinNumBackupsToRetainOk() (*int32, bool)`

GetMinNumBackupsToRetainOk returns a tuple with the MinNumBackupsToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumBackupsToRetain

`func (o *MultiTableBackupParams) SetMinNumBackupsToRetain(v int32)`

SetMinNumBackupsToRetain sets MinNumBackupsToRetain field to given value.

### HasMinNumBackupsToRetain

`func (o *MultiTableBackupParams) HasMinNumBackupsToRetain() bool`

HasMinNumBackupsToRetain returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *MultiTableBackupParams) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *MultiTableBackupParams) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *MultiTableBackupParams) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *MultiTableBackupParams) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *MultiTableBackupParams) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *MultiTableBackupParams) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *MultiTableBackupParams) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *MultiTableBackupParams) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetParallelism

`func (o *MultiTableBackupParams) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *MultiTableBackupParams) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *MultiTableBackupParams) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *MultiTableBackupParams) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetPreviousTaskUUID

`func (o *MultiTableBackupParams) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *MultiTableBackupParams) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *MultiTableBackupParams) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *MultiTableBackupParams) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetRestoreTimeStamp

`func (o *MultiTableBackupParams) GetRestoreTimeStamp() string`

GetRestoreTimeStamp returns the RestoreTimeStamp field if non-nil, zero value otherwise.

### GetRestoreTimeStampOk

`func (o *MultiTableBackupParams) GetRestoreTimeStampOk() (*string, bool)`

GetRestoreTimeStampOk returns a tuple with the RestoreTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeStamp

`func (o *MultiTableBackupParams) SetRestoreTimeStamp(v string)`

SetRestoreTimeStamp sets RestoreTimeStamp field to given value.

### HasRestoreTimeStamp

`func (o *MultiTableBackupParams) HasRestoreTimeStamp() bool`

HasRestoreTimeStamp returns a boolean if a field has been set.

### GetScheduleUUID

`func (o *MultiTableBackupParams) GetScheduleUUID() string`

GetScheduleUUID returns the ScheduleUUID field if non-nil, zero value otherwise.

### GetScheduleUUIDOk

`func (o *MultiTableBackupParams) GetScheduleUUIDOk() (*string, bool)`

GetScheduleUUIDOk returns a tuple with the ScheduleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUID

`func (o *MultiTableBackupParams) SetScheduleUUID(v string)`

SetScheduleUUID sets ScheduleUUID field to given value.

### HasScheduleUUID

`func (o *MultiTableBackupParams) HasScheduleUUID() bool`

HasScheduleUUID returns a boolean if a field has been set.

### GetSchedulingFrequency

`func (o *MultiTableBackupParams) GetSchedulingFrequency() int64`

GetSchedulingFrequency returns the SchedulingFrequency field if non-nil, zero value otherwise.

### GetSchedulingFrequencyOk

`func (o *MultiTableBackupParams) GetSchedulingFrequencyOk() (*int64, bool)`

GetSchedulingFrequencyOk returns a tuple with the SchedulingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingFrequency

`func (o *MultiTableBackupParams) SetSchedulingFrequency(v int64)`

SetSchedulingFrequency sets SchedulingFrequency field to given value.

### HasSchedulingFrequency

`func (o *MultiTableBackupParams) HasSchedulingFrequency() bool`

HasSchedulingFrequency returns a boolean if a field has been set.

### GetSourceXClusterConfigs

`func (o *MultiTableBackupParams) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *MultiTableBackupParams) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *MultiTableBackupParams) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *MultiTableBackupParams) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetSse

`func (o *MultiTableBackupParams) GetSse() bool`

GetSse returns the Sse field if non-nil, zero value otherwise.

### GetSseOk

`func (o *MultiTableBackupParams) GetSseOk() (*bool, bool)`

GetSseOk returns a tuple with the Sse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSse

`func (o *MultiTableBackupParams) SetSse(v bool)`

SetSse sets Sse field to given value.

### HasSse

`func (o *MultiTableBackupParams) HasSse() bool`

HasSse returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *MultiTableBackupParams) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *MultiTableBackupParams) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *MultiTableBackupParams) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetStorageLocation

`func (o *MultiTableBackupParams) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *MultiTableBackupParams) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *MultiTableBackupParams) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *MultiTableBackupParams) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### GetTableName

`func (o *MultiTableBackupParams) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *MultiTableBackupParams) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *MultiTableBackupParams) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *MultiTableBackupParams) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetTableNameList

`func (o *MultiTableBackupParams) GetTableNameList() []string`

GetTableNameList returns the TableNameList field if non-nil, zero value otherwise.

### GetTableNameListOk

`func (o *MultiTableBackupParams) GetTableNameListOk() (*[]string, bool)`

GetTableNameListOk returns a tuple with the TableNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNameList

`func (o *MultiTableBackupParams) SetTableNameList(v []string)`

SetTableNameList sets TableNameList field to given value.

### HasTableNameList

`func (o *MultiTableBackupParams) HasTableNameList() bool`

HasTableNameList returns a boolean if a field has been set.

### GetTableUUID

`func (o *MultiTableBackupParams) GetTableUUID() string`

GetTableUUID returns the TableUUID field if non-nil, zero value otherwise.

### GetTableUUIDOk

`func (o *MultiTableBackupParams) GetTableUUIDOk() (*string, bool)`

GetTableUUIDOk returns a tuple with the TableUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUID

`func (o *MultiTableBackupParams) SetTableUUID(v string)`

SetTableUUID sets TableUUID field to given value.

### HasTableUUID

`func (o *MultiTableBackupParams) HasTableUUID() bool`

HasTableUUID returns a boolean if a field has been set.

### GetTableUUIDList

`func (o *MultiTableBackupParams) GetTableUUIDList() []string`

GetTableUUIDList returns the TableUUIDList field if non-nil, zero value otherwise.

### GetTableUUIDListOk

`func (o *MultiTableBackupParams) GetTableUUIDListOk() (*[]string, bool)`

GetTableUUIDListOk returns a tuple with the TableUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUIDList

`func (o *MultiTableBackupParams) SetTableUUIDList(v []string)`

SetTableUUIDList sets TableUUIDList field to given value.


### GetTargetXClusterConfigs

`func (o *MultiTableBackupParams) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *MultiTableBackupParams) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *MultiTableBackupParams) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *MultiTableBackupParams) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.

### GetTimeBeforeDelete

`func (o *MultiTableBackupParams) GetTimeBeforeDelete() int64`

GetTimeBeforeDelete returns the TimeBeforeDelete field if non-nil, zero value otherwise.

### GetTimeBeforeDeleteOk

`func (o *MultiTableBackupParams) GetTimeBeforeDeleteOk() (*int64, bool)`

GetTimeBeforeDeleteOk returns a tuple with the TimeBeforeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBeforeDelete

`func (o *MultiTableBackupParams) SetTimeBeforeDelete(v int64)`

SetTimeBeforeDelete sets TimeBeforeDelete field to given value.

### HasTimeBeforeDelete

`func (o *MultiTableBackupParams) HasTimeBeforeDelete() bool`

HasTimeBeforeDelete returns a boolean if a field has been set.

### GetTransactionalBackup

`func (o *MultiTableBackupParams) GetTransactionalBackup() bool`

GetTransactionalBackup returns the TransactionalBackup field if non-nil, zero value otherwise.

### GetTransactionalBackupOk

`func (o *MultiTableBackupParams) GetTransactionalBackupOk() (*bool, bool)`

GetTransactionalBackupOk returns a tuple with the TransactionalBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionalBackup

`func (o *MultiTableBackupParams) SetTransactionalBackup(v bool)`

SetTransactionalBackup sets TransactionalBackup field to given value.

### HasTransactionalBackup

`func (o *MultiTableBackupParams) HasTransactionalBackup() bool`

HasTransactionalBackup returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *MultiTableBackupParams) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *MultiTableBackupParams) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *MultiTableBackupParams) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *MultiTableBackupParams) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetUseTablespaces

`func (o *MultiTableBackupParams) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *MultiTableBackupParams) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *MultiTableBackupParams) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.

### HasUseTablespaces

`func (o *MultiTableBackupParams) HasUseTablespaces() bool`

HasUseTablespaces returns a boolean if a field has been set.

### GetYbPrevSoftwareVersion

`func (o *MultiTableBackupParams) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *MultiTableBackupParams) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *MultiTableBackupParams) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *MultiTableBackupParams) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


