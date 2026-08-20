# TaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskDetails** | Pointer to [**[]TaskSubtaskGroupDetails**](TaskSubtaskGroupDetails.md) | User-facing subtask groups for this task. | [optional] [readonly] 
**VersionNumbers** | Pointer to [**TaskVersionNumbers**](TaskVersionNumbers.md) |  | [optional] 
**SoftwareUpgradeProgress** | Pointer to [**SoftwareUpgradeProgress**](SoftwareUpgradeProgress.md) |  | [optional] 
**AuditLogConfig** | Pointer to [**AuditLogConfig**](AuditLogConfig.md) |  | [optional] 
**QueryLogConfig** | Pointer to [**QueryLogConfig**](QueryLogConfig.md) |  | [optional] 
**MetricsExportConfig** | Pointer to [**MetricsExportConfig**](MetricsExportConfig.md) |  | [optional] 
**ModifiedExportTypes** | Pointer to [**[]ExportType**](ExportType.md) | Telemetry config types (audit logs, query logs, metrics) being modified by this telemetry export configure task. Computed by diffing the requested config against the currently stored config at submit time.  | [optional] [readonly] 

## Methods

### NewTaskDetails

`func NewTaskDetails() *TaskDetails`

NewTaskDetails instantiates a new TaskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDetailsWithDefaults

`func NewTaskDetailsWithDefaults() *TaskDetails`

NewTaskDetailsWithDefaults instantiates a new TaskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskDetails

`func (o *TaskDetails) GetTaskDetails() []TaskSubtaskGroupDetails`

GetTaskDetails returns the TaskDetails field if non-nil, zero value otherwise.

### GetTaskDetailsOk

`func (o *TaskDetails) GetTaskDetailsOk() (*[]TaskSubtaskGroupDetails, bool)`

GetTaskDetailsOk returns a tuple with the TaskDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDetails

`func (o *TaskDetails) SetTaskDetails(v []TaskSubtaskGroupDetails)`

SetTaskDetails sets TaskDetails field to given value.

### HasTaskDetails

`func (o *TaskDetails) HasTaskDetails() bool`

HasTaskDetails returns a boolean if a field has been set.

### GetVersionNumbers

`func (o *TaskDetails) GetVersionNumbers() TaskVersionNumbers`

GetVersionNumbers returns the VersionNumbers field if non-nil, zero value otherwise.

### GetVersionNumbersOk

`func (o *TaskDetails) GetVersionNumbersOk() (*TaskVersionNumbers, bool)`

GetVersionNumbersOk returns a tuple with the VersionNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumbers

`func (o *TaskDetails) SetVersionNumbers(v TaskVersionNumbers)`

SetVersionNumbers sets VersionNumbers field to given value.

### HasVersionNumbers

`func (o *TaskDetails) HasVersionNumbers() bool`

HasVersionNumbers returns a boolean if a field has been set.

### GetSoftwareUpgradeProgress

`func (o *TaskDetails) GetSoftwareUpgradeProgress() SoftwareUpgradeProgress`

GetSoftwareUpgradeProgress returns the SoftwareUpgradeProgress field if non-nil, zero value otherwise.

### GetSoftwareUpgradeProgressOk

`func (o *TaskDetails) GetSoftwareUpgradeProgressOk() (*SoftwareUpgradeProgress, bool)`

GetSoftwareUpgradeProgressOk returns a tuple with the SoftwareUpgradeProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpgradeProgress

`func (o *TaskDetails) SetSoftwareUpgradeProgress(v SoftwareUpgradeProgress)`

SetSoftwareUpgradeProgress sets SoftwareUpgradeProgress field to given value.

### HasSoftwareUpgradeProgress

`func (o *TaskDetails) HasSoftwareUpgradeProgress() bool`

HasSoftwareUpgradeProgress returns a boolean if a field has been set.

### GetAuditLogConfig

`func (o *TaskDetails) GetAuditLogConfig() AuditLogConfig`

GetAuditLogConfig returns the AuditLogConfig field if non-nil, zero value otherwise.

### GetAuditLogConfigOk

`func (o *TaskDetails) GetAuditLogConfigOk() (*AuditLogConfig, bool)`

GetAuditLogConfigOk returns a tuple with the AuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogConfig

`func (o *TaskDetails) SetAuditLogConfig(v AuditLogConfig)`

SetAuditLogConfig sets AuditLogConfig field to given value.

### HasAuditLogConfig

`func (o *TaskDetails) HasAuditLogConfig() bool`

HasAuditLogConfig returns a boolean if a field has been set.

### GetQueryLogConfig

`func (o *TaskDetails) GetQueryLogConfig() QueryLogConfig`

GetQueryLogConfig returns the QueryLogConfig field if non-nil, zero value otherwise.

### GetQueryLogConfigOk

`func (o *TaskDetails) GetQueryLogConfigOk() (*QueryLogConfig, bool)`

GetQueryLogConfigOk returns a tuple with the QueryLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryLogConfig

`func (o *TaskDetails) SetQueryLogConfig(v QueryLogConfig)`

SetQueryLogConfig sets QueryLogConfig field to given value.

### HasQueryLogConfig

`func (o *TaskDetails) HasQueryLogConfig() bool`

HasQueryLogConfig returns a boolean if a field has been set.

### GetMetricsExportConfig

`func (o *TaskDetails) GetMetricsExportConfig() MetricsExportConfig`

GetMetricsExportConfig returns the MetricsExportConfig field if non-nil, zero value otherwise.

### GetMetricsExportConfigOk

`func (o *TaskDetails) GetMetricsExportConfigOk() (*MetricsExportConfig, bool)`

GetMetricsExportConfigOk returns a tuple with the MetricsExportConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsExportConfig

`func (o *TaskDetails) SetMetricsExportConfig(v MetricsExportConfig)`

SetMetricsExportConfig sets MetricsExportConfig field to given value.

### HasMetricsExportConfig

`func (o *TaskDetails) HasMetricsExportConfig() bool`

HasMetricsExportConfig returns a boolean if a field has been set.

### GetModifiedExportTypes

`func (o *TaskDetails) GetModifiedExportTypes() []ExportType`

GetModifiedExportTypes returns the ModifiedExportTypes field if non-nil, zero value otherwise.

### GetModifiedExportTypesOk

`func (o *TaskDetails) GetModifiedExportTypesOk() (*[]ExportType, bool)`

GetModifiedExportTypesOk returns a tuple with the ModifiedExportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedExportTypes

`func (o *TaskDetails) SetModifiedExportTypes(v []ExportType)`

SetModifiedExportTypes sets ModifiedExportTypes field to given value.

### HasModifiedExportTypes

`func (o *TaskDetails) HasModifiedExportTypes() bool`

HasModifiedExportTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


