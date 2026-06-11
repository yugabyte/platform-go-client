# BackupApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRangeStart** | Pointer to **time.Time** | Include backups created on or after this time. | [optional] 
**DateRangeEnd** | Pointer to **time.Time** | Include backups created on or before this time. | [optional] 
**States** | Pointer to [**[]BackupState**](BackupState.md) | Backup states to include. | [optional] 
**KeyspaceList** | Pointer to **[]string** | Keyspace names to include. | [optional] 
**UniverseNameList** | Pointer to **[]string** | Universe names to include. | [optional] 
**StorageConfigUuidList** | Pointer to **[]string** | Storage config UUIDs to include. | [optional] 
**ScheduleUuidList** | Pointer to **[]string** | Backup schedule UUIDs to include. | [optional] 
**UniverseUuidList** | Pointer to **[]string** | Universe UUIDs to include. | [optional] 
**BackupUuidList** | Pointer to **[]string** | Backup UUIDs to include. | [optional] 
**OnlyShowDeletedUniverses** | Pointer to **bool** | When true, include backups whose universe has been deleted. | [optional] 
**OnlyShowDeletedConfigs** | Pointer to **bool** | When true, include backups whose storage config has been deleted. | [optional] 
**ShowHidden** | Pointer to **bool** | When true, include hidden backups. | [optional] 

## Methods

### NewBackupApiFilter

`func NewBackupApiFilter() *BackupApiFilter`

NewBackupApiFilter instantiates a new BackupApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupApiFilterWithDefaults

`func NewBackupApiFilterWithDefaults() *BackupApiFilter`

NewBackupApiFilterWithDefaults instantiates a new BackupApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRangeStart

`func (o *BackupApiFilter) GetDateRangeStart() time.Time`

GetDateRangeStart returns the DateRangeStart field if non-nil, zero value otherwise.

### GetDateRangeStartOk

`func (o *BackupApiFilter) GetDateRangeStartOk() (*time.Time, bool)`

GetDateRangeStartOk returns a tuple with the DateRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRangeStart

`func (o *BackupApiFilter) SetDateRangeStart(v time.Time)`

SetDateRangeStart sets DateRangeStart field to given value.

### HasDateRangeStart

`func (o *BackupApiFilter) HasDateRangeStart() bool`

HasDateRangeStart returns a boolean if a field has been set.

### GetDateRangeEnd

`func (o *BackupApiFilter) GetDateRangeEnd() time.Time`

GetDateRangeEnd returns the DateRangeEnd field if non-nil, zero value otherwise.

### GetDateRangeEndOk

`func (o *BackupApiFilter) GetDateRangeEndOk() (*time.Time, bool)`

GetDateRangeEndOk returns a tuple with the DateRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRangeEnd

`func (o *BackupApiFilter) SetDateRangeEnd(v time.Time)`

SetDateRangeEnd sets DateRangeEnd field to given value.

### HasDateRangeEnd

`func (o *BackupApiFilter) HasDateRangeEnd() bool`

HasDateRangeEnd returns a boolean if a field has been set.

### GetStates

`func (o *BackupApiFilter) GetStates() []BackupState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *BackupApiFilter) GetStatesOk() (*[]BackupState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *BackupApiFilter) SetStates(v []BackupState)`

SetStates sets States field to given value.

### HasStates

`func (o *BackupApiFilter) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetKeyspaceList

`func (o *BackupApiFilter) GetKeyspaceList() []string`

GetKeyspaceList returns the KeyspaceList field if non-nil, zero value otherwise.

### GetKeyspaceListOk

`func (o *BackupApiFilter) GetKeyspaceListOk() (*[]string, bool)`

GetKeyspaceListOk returns a tuple with the KeyspaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceList

`func (o *BackupApiFilter) SetKeyspaceList(v []string)`

SetKeyspaceList sets KeyspaceList field to given value.

### HasKeyspaceList

`func (o *BackupApiFilter) HasKeyspaceList() bool`

HasKeyspaceList returns a boolean if a field has been set.

### GetUniverseNameList

`func (o *BackupApiFilter) GetUniverseNameList() []string`

GetUniverseNameList returns the UniverseNameList field if non-nil, zero value otherwise.

### GetUniverseNameListOk

`func (o *BackupApiFilter) GetUniverseNameListOk() (*[]string, bool)`

GetUniverseNameListOk returns a tuple with the UniverseNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseNameList

`func (o *BackupApiFilter) SetUniverseNameList(v []string)`

SetUniverseNameList sets UniverseNameList field to given value.

### HasUniverseNameList

`func (o *BackupApiFilter) HasUniverseNameList() bool`

HasUniverseNameList returns a boolean if a field has been set.

### GetStorageConfigUuidList

`func (o *BackupApiFilter) GetStorageConfigUuidList() []string`

GetStorageConfigUuidList returns the StorageConfigUuidList field if non-nil, zero value otherwise.

### GetStorageConfigUuidListOk

`func (o *BackupApiFilter) GetStorageConfigUuidListOk() (*[]string, bool)`

GetStorageConfigUuidListOk returns a tuple with the StorageConfigUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUuidList

`func (o *BackupApiFilter) SetStorageConfigUuidList(v []string)`

SetStorageConfigUuidList sets StorageConfigUuidList field to given value.

### HasStorageConfigUuidList

`func (o *BackupApiFilter) HasStorageConfigUuidList() bool`

HasStorageConfigUuidList returns a boolean if a field has been set.

### GetScheduleUuidList

`func (o *BackupApiFilter) GetScheduleUuidList() []string`

GetScheduleUuidList returns the ScheduleUuidList field if non-nil, zero value otherwise.

### GetScheduleUuidListOk

`func (o *BackupApiFilter) GetScheduleUuidListOk() (*[]string, bool)`

GetScheduleUuidListOk returns a tuple with the ScheduleUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUuidList

`func (o *BackupApiFilter) SetScheduleUuidList(v []string)`

SetScheduleUuidList sets ScheduleUuidList field to given value.

### HasScheduleUuidList

`func (o *BackupApiFilter) HasScheduleUuidList() bool`

HasScheduleUuidList returns a boolean if a field has been set.

### GetUniverseUuidList

`func (o *BackupApiFilter) GetUniverseUuidList() []string`

GetUniverseUuidList returns the UniverseUuidList field if non-nil, zero value otherwise.

### GetUniverseUuidListOk

`func (o *BackupApiFilter) GetUniverseUuidListOk() (*[]string, bool)`

GetUniverseUuidListOk returns a tuple with the UniverseUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUuidList

`func (o *BackupApiFilter) SetUniverseUuidList(v []string)`

SetUniverseUuidList sets UniverseUuidList field to given value.

### HasUniverseUuidList

`func (o *BackupApiFilter) HasUniverseUuidList() bool`

HasUniverseUuidList returns a boolean if a field has been set.

### GetBackupUuidList

`func (o *BackupApiFilter) GetBackupUuidList() []string`

GetBackupUuidList returns the BackupUuidList field if non-nil, zero value otherwise.

### GetBackupUuidListOk

`func (o *BackupApiFilter) GetBackupUuidListOk() (*[]string, bool)`

GetBackupUuidListOk returns a tuple with the BackupUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUuidList

`func (o *BackupApiFilter) SetBackupUuidList(v []string)`

SetBackupUuidList sets BackupUuidList field to given value.

### HasBackupUuidList

`func (o *BackupApiFilter) HasBackupUuidList() bool`

HasBackupUuidList returns a boolean if a field has been set.

### GetOnlyShowDeletedUniverses

`func (o *BackupApiFilter) GetOnlyShowDeletedUniverses() bool`

GetOnlyShowDeletedUniverses returns the OnlyShowDeletedUniverses field if non-nil, zero value otherwise.

### GetOnlyShowDeletedUniversesOk

`func (o *BackupApiFilter) GetOnlyShowDeletedUniversesOk() (*bool, bool)`

GetOnlyShowDeletedUniversesOk returns a tuple with the OnlyShowDeletedUniverses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyShowDeletedUniverses

`func (o *BackupApiFilter) SetOnlyShowDeletedUniverses(v bool)`

SetOnlyShowDeletedUniverses sets OnlyShowDeletedUniverses field to given value.

### HasOnlyShowDeletedUniverses

`func (o *BackupApiFilter) HasOnlyShowDeletedUniverses() bool`

HasOnlyShowDeletedUniverses returns a boolean if a field has been set.

### GetOnlyShowDeletedConfigs

`func (o *BackupApiFilter) GetOnlyShowDeletedConfigs() bool`

GetOnlyShowDeletedConfigs returns the OnlyShowDeletedConfigs field if non-nil, zero value otherwise.

### GetOnlyShowDeletedConfigsOk

`func (o *BackupApiFilter) GetOnlyShowDeletedConfigsOk() (*bool, bool)`

GetOnlyShowDeletedConfigsOk returns a tuple with the OnlyShowDeletedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyShowDeletedConfigs

`func (o *BackupApiFilter) SetOnlyShowDeletedConfigs(v bool)`

SetOnlyShowDeletedConfigs sets OnlyShowDeletedConfigs field to given value.

### HasOnlyShowDeletedConfigs

`func (o *BackupApiFilter) HasOnlyShowDeletedConfigs() bool`

HasOnlyShowDeletedConfigs returns a boolean if a field has been set.

### GetShowHidden

`func (o *BackupApiFilter) GetShowHidden() bool`

GetShowHidden returns the ShowHidden field if non-nil, zero value otherwise.

### GetShowHiddenOk

`func (o *BackupApiFilter) GetShowHiddenOk() (*bool, bool)`

GetShowHiddenOk returns a tuple with the ShowHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHidden

`func (o *BackupApiFilter) SetShowHidden(v bool)`

SetShowHidden sets ShowHidden field to given value.

### HasShowHidden

`func (o *BackupApiFilter) HasShowHidden() bool`

HasShowHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


