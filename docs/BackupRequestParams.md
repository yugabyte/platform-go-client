# BackupRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlterLoadBalancer** | Pointer to **bool** | Alter load balancer state | [optional] 
**BackupType** | Pointer to **string** | Backup type | [optional] 
**CmkArn** | Pointer to **string** | Amazon Resource Name (ARN) of the CMK | [optional] 
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**CronExpression** | Pointer to **string** | Cron expression for a recurring backup | [optional] 
**CustomerUUID** | Pointer to **string** | Customer UUID | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**EnableVerboseLogs** | Pointer to **bool** | Is verbose logging enabled | [optional] 
**EncryptionAtRestConfig** | Pointer to [**EncryptionAtRestConfig**](EncryptionAtRestConfig.md) |  | [optional] 
**ErrorString** | Pointer to **string** | Error message | [optional] 
**ExpectedUniverseVersion** | Pointer to **int32** | Expected universe version | [optional] 
**ExtraDependencies** | Pointer to [**ExtraDependencies**](ExtraDependencies.md) |  | [optional] 
**FirstTry** | Pointer to **bool** | Whether this task has been tried before | [optional] 
**IgnoreErrors** | Pointer to **bool** | Should table backup errors be ignored | [optional] 
**KeyspaceTableList** | Pointer to [**[]KeyspaceTable**](KeyspaceTable.md) | Backup info | [optional] 
**KmsConfigUUID** | Pointer to **string** | KMS configuration UUID | [optional] 
**MinNumBackupsToRetain** | Pointer to **int32** | Minimum number of backups to retain for a particular backup schedule | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 
**NodeExporterUser** | Pointer to **string** | Node exporter user | [optional] 
**Parallelism** | Pointer to **int32** | Number of concurrent commands to run on nodes over SSH | [optional] 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID only if this task is a retry | [optional] 
**ScheduleUUID** | Pointer to **string** | Schedule UUID | [optional] 
**SchedulingFrequency** | Pointer to **int64** | Frequency to run the backup, in milliseconds | [optional] 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] [readonly] 
**Sse** | Pointer to **bool** | Is SSE | [optional] 
**StorageConfigUUID** | **string** | Storage configuration UUID | 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] [readonly] 
**TimeBeforeDelete** | Pointer to **int64** | Time before deleting the backup from storage, in milliseconds | [optional] 
**UniverseUUID** | **string** | Universe UUID | 
**YbPrevSoftwareVersion** | Pointer to **string** | Previous software version | [optional] 

## Methods

### NewBackupRequestParams

`func NewBackupRequestParams(storageConfigUUID string, universeUUID string, ) *BackupRequestParams`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


