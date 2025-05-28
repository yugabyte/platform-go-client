# BackupScheduleToggleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunImmediateBackupOnResume** | Pointer to **bool** | Run a full or incremental backup if required when resuming a stopped schedule. When false (default), the full backup will instead run at its normally scheduled time. | [optional] 
**Status** | **string** | State of the schedule | 

## Methods

### NewBackupScheduleToggleParams

`func NewBackupScheduleToggleParams(status string, ) *BackupScheduleToggleParams`

NewBackupScheduleToggleParams instantiates a new BackupScheduleToggleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleToggleParamsWithDefaults

`func NewBackupScheduleToggleParamsWithDefaults() *BackupScheduleToggleParams`

NewBackupScheduleToggleParamsWithDefaults instantiates a new BackupScheduleToggleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunImmediateBackupOnResume

`func (o *BackupScheduleToggleParams) GetRunImmediateBackupOnResume() bool`

GetRunImmediateBackupOnResume returns the RunImmediateBackupOnResume field if non-nil, zero value otherwise.

### GetRunImmediateBackupOnResumeOk

`func (o *BackupScheduleToggleParams) GetRunImmediateBackupOnResumeOk() (*bool, bool)`

GetRunImmediateBackupOnResumeOk returns a tuple with the RunImmediateBackupOnResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunImmediateBackupOnResume

`func (o *BackupScheduleToggleParams) SetRunImmediateBackupOnResume(v bool)`

SetRunImmediateBackupOnResume sets RunImmediateBackupOnResume field to given value.

### HasRunImmediateBackupOnResume

`func (o *BackupScheduleToggleParams) HasRunImmediateBackupOnResume() bool`

HasRunImmediateBackupOnResume returns a boolean if a field has been set.

### GetStatus

`func (o *BackupScheduleToggleParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupScheduleToggleParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupScheduleToggleParams) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


