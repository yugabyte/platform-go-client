# BackupScheduleEditParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** | Cron expression for scheduling | [optional] 
**FrequencyTimeUnit** | Pointer to **string** | Time Unit for frequency | [optional] 
**IncrementalBackupFrequency** | Pointer to **int64** | Frequency of incremental backup schedule | [optional] 
**IncrementalBackupFrequencyTimeUnit** | Pointer to **string** | TimeUnit for incremental Backup Schedule frequency | [optional] 
**SchedulingFrequency** | Pointer to **int64** | Frequency of the schedule | [optional] 
**TimeBeforeDelete** | Pointer to **int64** | Time before deleting the backup from storage, in milliseconds | [optional] 
**UsePrivileges** | Pointer to **bool** | Backup privileges for roles | [optional] 
**UseRoles** | Pointer to **bool** | Backup global ysql roles | [optional] 
**UseTablespaces** | Pointer to **bool** | Is tablespaces information included | [optional] 

## Methods

### NewBackupScheduleEditParams

`func NewBackupScheduleEditParams() *BackupScheduleEditParams`

NewBackupScheduleEditParams instantiates a new BackupScheduleEditParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleEditParamsWithDefaults

`func NewBackupScheduleEditParamsWithDefaults() *BackupScheduleEditParams`

NewBackupScheduleEditParamsWithDefaults instantiates a new BackupScheduleEditParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *BackupScheduleEditParams) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *BackupScheduleEditParams) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *BackupScheduleEditParams) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *BackupScheduleEditParams) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetFrequencyTimeUnit

`func (o *BackupScheduleEditParams) GetFrequencyTimeUnit() string`

GetFrequencyTimeUnit returns the FrequencyTimeUnit field if non-nil, zero value otherwise.

### GetFrequencyTimeUnitOk

`func (o *BackupScheduleEditParams) GetFrequencyTimeUnitOk() (*string, bool)`

GetFrequencyTimeUnitOk returns a tuple with the FrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyTimeUnit

`func (o *BackupScheduleEditParams) SetFrequencyTimeUnit(v string)`

SetFrequencyTimeUnit sets FrequencyTimeUnit field to given value.

### HasFrequencyTimeUnit

`func (o *BackupScheduleEditParams) HasFrequencyTimeUnit() bool`

HasFrequencyTimeUnit returns a boolean if a field has been set.

### GetIncrementalBackupFrequency

`func (o *BackupScheduleEditParams) GetIncrementalBackupFrequency() int64`

GetIncrementalBackupFrequency returns the IncrementalBackupFrequency field if non-nil, zero value otherwise.

### GetIncrementalBackupFrequencyOk

`func (o *BackupScheduleEditParams) GetIncrementalBackupFrequencyOk() (*int64, bool)`

GetIncrementalBackupFrequencyOk returns a tuple with the IncrementalBackupFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupFrequency

`func (o *BackupScheduleEditParams) SetIncrementalBackupFrequency(v int64)`

SetIncrementalBackupFrequency sets IncrementalBackupFrequency field to given value.

### HasIncrementalBackupFrequency

`func (o *BackupScheduleEditParams) HasIncrementalBackupFrequency() bool`

HasIncrementalBackupFrequency returns a boolean if a field has been set.

### GetIncrementalBackupFrequencyTimeUnit

`func (o *BackupScheduleEditParams) GetIncrementalBackupFrequencyTimeUnit() string`

GetIncrementalBackupFrequencyTimeUnit returns the IncrementalBackupFrequencyTimeUnit field if non-nil, zero value otherwise.

### GetIncrementalBackupFrequencyTimeUnitOk

`func (o *BackupScheduleEditParams) GetIncrementalBackupFrequencyTimeUnitOk() (*string, bool)`

GetIncrementalBackupFrequencyTimeUnitOk returns a tuple with the IncrementalBackupFrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupFrequencyTimeUnit

`func (o *BackupScheduleEditParams) SetIncrementalBackupFrequencyTimeUnit(v string)`

SetIncrementalBackupFrequencyTimeUnit sets IncrementalBackupFrequencyTimeUnit field to given value.

### HasIncrementalBackupFrequencyTimeUnit

`func (o *BackupScheduleEditParams) HasIncrementalBackupFrequencyTimeUnit() bool`

HasIncrementalBackupFrequencyTimeUnit returns a boolean if a field has been set.

### GetSchedulingFrequency

`func (o *BackupScheduleEditParams) GetSchedulingFrequency() int64`

GetSchedulingFrequency returns the SchedulingFrequency field if non-nil, zero value otherwise.

### GetSchedulingFrequencyOk

`func (o *BackupScheduleEditParams) GetSchedulingFrequencyOk() (*int64, bool)`

GetSchedulingFrequencyOk returns a tuple with the SchedulingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingFrequency

`func (o *BackupScheduleEditParams) SetSchedulingFrequency(v int64)`

SetSchedulingFrequency sets SchedulingFrequency field to given value.

### HasSchedulingFrequency

`func (o *BackupScheduleEditParams) HasSchedulingFrequency() bool`

HasSchedulingFrequency returns a boolean if a field has been set.

### GetTimeBeforeDelete

`func (o *BackupScheduleEditParams) GetTimeBeforeDelete() int64`

GetTimeBeforeDelete returns the TimeBeforeDelete field if non-nil, zero value otherwise.

### GetTimeBeforeDeleteOk

`func (o *BackupScheduleEditParams) GetTimeBeforeDeleteOk() (*int64, bool)`

GetTimeBeforeDeleteOk returns a tuple with the TimeBeforeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBeforeDelete

`func (o *BackupScheduleEditParams) SetTimeBeforeDelete(v int64)`

SetTimeBeforeDelete sets TimeBeforeDelete field to given value.

### HasTimeBeforeDelete

`func (o *BackupScheduleEditParams) HasTimeBeforeDelete() bool`

HasTimeBeforeDelete returns a boolean if a field has been set.

### GetUsePrivileges

`func (o *BackupScheduleEditParams) GetUsePrivileges() bool`

GetUsePrivileges returns the UsePrivileges field if non-nil, zero value otherwise.

### GetUsePrivilegesOk

`func (o *BackupScheduleEditParams) GetUsePrivilegesOk() (*bool, bool)`

GetUsePrivilegesOk returns a tuple with the UsePrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePrivileges

`func (o *BackupScheduleEditParams) SetUsePrivileges(v bool)`

SetUsePrivileges sets UsePrivileges field to given value.

### HasUsePrivileges

`func (o *BackupScheduleEditParams) HasUsePrivileges() bool`

HasUsePrivileges returns a boolean if a field has been set.

### GetUseRoles

`func (o *BackupScheduleEditParams) GetUseRoles() bool`

GetUseRoles returns the UseRoles field if non-nil, zero value otherwise.

### GetUseRolesOk

`func (o *BackupScheduleEditParams) GetUseRolesOk() (*bool, bool)`

GetUseRolesOk returns a tuple with the UseRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRoles

`func (o *BackupScheduleEditParams) SetUseRoles(v bool)`

SetUseRoles sets UseRoles field to given value.

### HasUseRoles

`func (o *BackupScheduleEditParams) HasUseRoles() bool`

HasUseRoles returns a boolean if a field has been set.

### GetUseTablespaces

`func (o *BackupScheduleEditParams) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *BackupScheduleEditParams) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *BackupScheduleEditParams) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.

### HasUseTablespaces

`func (o *BackupScheduleEditParams) HasUseTablespaces() bool`

HasUseTablespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


