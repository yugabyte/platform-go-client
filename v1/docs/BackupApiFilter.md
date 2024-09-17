# BackupApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupUUIDList** | **[]string** |  | 
**DateRangeEnd** | Pointer to **time.Time** | The end date for backup filter. | [optional] 
**DateRangeStart** | Pointer to **time.Time** | The start date for backup filter. | [optional] 
**KeyspaceList** | **[]string** |  | 
**OnlyShowDeletedConfigs** | **bool** |  | 
**OnlyShowDeletedUniverses** | **bool** |  | 
**ScheduleUUIDList** | **[]string** |  | 
**ShowHidden** | **bool** |  | 
**States** | **[]string** |  | 
**StorageConfigUUIDList** | **[]string** |  | 
**UniverseNameList** | **[]string** |  | 
**UniverseUUIDList** | **[]string** |  | 

## Methods

### NewBackupApiFilter

`func NewBackupApiFilter(backupUUIDList []string, keyspaceList []string, onlyShowDeletedConfigs bool, onlyShowDeletedUniverses bool, scheduleUUIDList []string, showHidden bool, states []string, storageConfigUUIDList []string, universeNameList []string, universeUUIDList []string, ) *BackupApiFilter`

NewBackupApiFilter instantiates a new BackupApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupApiFilterWithDefaults

`func NewBackupApiFilterWithDefaults() *BackupApiFilter`

NewBackupApiFilterWithDefaults instantiates a new BackupApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupUUIDList

`func (o *BackupApiFilter) GetBackupUUIDList() []string`

GetBackupUUIDList returns the BackupUUIDList field if non-nil, zero value otherwise.

### GetBackupUUIDListOk

`func (o *BackupApiFilter) GetBackupUUIDListOk() (*[]string, bool)`

GetBackupUUIDListOk returns a tuple with the BackupUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUIDList

`func (o *BackupApiFilter) SetBackupUUIDList(v []string)`

SetBackupUUIDList sets BackupUUIDList field to given value.


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


### GetScheduleUUIDList

`func (o *BackupApiFilter) GetScheduleUUIDList() []string`

GetScheduleUUIDList returns the ScheduleUUIDList field if non-nil, zero value otherwise.

### GetScheduleUUIDListOk

`func (o *BackupApiFilter) GetScheduleUUIDListOk() (*[]string, bool)`

GetScheduleUUIDListOk returns a tuple with the ScheduleUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUUIDList

`func (o *BackupApiFilter) SetScheduleUUIDList(v []string)`

SetScheduleUUIDList sets ScheduleUUIDList field to given value.


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


### GetStates

`func (o *BackupApiFilter) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *BackupApiFilter) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *BackupApiFilter) SetStates(v []string)`

SetStates sets States field to given value.


### GetStorageConfigUUIDList

`func (o *BackupApiFilter) GetStorageConfigUUIDList() []string`

GetStorageConfigUUIDList returns the StorageConfigUUIDList field if non-nil, zero value otherwise.

### GetStorageConfigUUIDListOk

`func (o *BackupApiFilter) GetStorageConfigUUIDListOk() (*[]string, bool)`

GetStorageConfigUUIDListOk returns a tuple with the StorageConfigUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUIDList

`func (o *BackupApiFilter) SetStorageConfigUUIDList(v []string)`

SetStorageConfigUUIDList sets StorageConfigUUIDList field to given value.


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


### GetUniverseUUIDList

`func (o *BackupApiFilter) GetUniverseUUIDList() []string`

GetUniverseUUIDList returns the UniverseUUIDList field if non-nil, zero value otherwise.

### GetUniverseUUIDListOk

`func (o *BackupApiFilter) GetUniverseUUIDListOk() (*[]string, bool)`

GetUniverseUUIDListOk returns a tuple with the UniverseUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUIDList

`func (o *BackupApiFilter) SetUniverseUUIDList(v []string)`

SetUniverseUUIDList sets UniverseUUIDList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


