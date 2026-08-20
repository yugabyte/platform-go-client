# RestoreApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRangeStart** | Pointer to **time.Time** | Include restores created on or after this time. | [optional] 
**DateRangeEnd** | Pointer to **time.Time** | Include restores created on or before this time. | [optional] 
**States** | Pointer to [**[]RestoreState**](RestoreState.md) | Restore states to include. | [optional] 
**UniverseNameList** | Pointer to **[]string** | Target universe names to include. | [optional] 
**SourceUniverseNameList** | Pointer to **[]string** | Source universe names to include. | [optional] 
**StorageConfigUuidList** | Pointer to **[]string** | Storage config UUIDs to include. | [optional] 
**UniverseUuidList** | Pointer to **[]string** | Target universe UUIDs to include. | [optional] 
**RestoreUuidList** | Pointer to **[]string** | Restore UUIDs to include. | [optional] 
**OnlyShowDeletedSourceUniverses** | Pointer to **bool** | When true, include restores whose source universe has been deleted. | [optional] 
**ShowHidden** | Pointer to **bool** | When true, include hidden restores. | [optional] 

## Methods

### NewRestoreApiFilter

`func NewRestoreApiFilter() *RestoreApiFilter`

NewRestoreApiFilter instantiates a new RestoreApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreApiFilterWithDefaults

`func NewRestoreApiFilterWithDefaults() *RestoreApiFilter`

NewRestoreApiFilterWithDefaults instantiates a new RestoreApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRangeStart

`func (o *RestoreApiFilter) GetDateRangeStart() time.Time`

GetDateRangeStart returns the DateRangeStart field if non-nil, zero value otherwise.

### GetDateRangeStartOk

`func (o *RestoreApiFilter) GetDateRangeStartOk() (*time.Time, bool)`

GetDateRangeStartOk returns a tuple with the DateRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRangeStart

`func (o *RestoreApiFilter) SetDateRangeStart(v time.Time)`

SetDateRangeStart sets DateRangeStart field to given value.

### HasDateRangeStart

`func (o *RestoreApiFilter) HasDateRangeStart() bool`

HasDateRangeStart returns a boolean if a field has been set.

### GetDateRangeEnd

`func (o *RestoreApiFilter) GetDateRangeEnd() time.Time`

GetDateRangeEnd returns the DateRangeEnd field if non-nil, zero value otherwise.

### GetDateRangeEndOk

`func (o *RestoreApiFilter) GetDateRangeEndOk() (*time.Time, bool)`

GetDateRangeEndOk returns a tuple with the DateRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRangeEnd

`func (o *RestoreApiFilter) SetDateRangeEnd(v time.Time)`

SetDateRangeEnd sets DateRangeEnd field to given value.

### HasDateRangeEnd

`func (o *RestoreApiFilter) HasDateRangeEnd() bool`

HasDateRangeEnd returns a boolean if a field has been set.

### GetStates

`func (o *RestoreApiFilter) GetStates() []RestoreState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *RestoreApiFilter) GetStatesOk() (*[]RestoreState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *RestoreApiFilter) SetStates(v []RestoreState)`

SetStates sets States field to given value.

### HasStates

`func (o *RestoreApiFilter) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetUniverseNameList

`func (o *RestoreApiFilter) GetUniverseNameList() []string`

GetUniverseNameList returns the UniverseNameList field if non-nil, zero value otherwise.

### GetUniverseNameListOk

`func (o *RestoreApiFilter) GetUniverseNameListOk() (*[]string, bool)`

GetUniverseNameListOk returns a tuple with the UniverseNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseNameList

`func (o *RestoreApiFilter) SetUniverseNameList(v []string)`

SetUniverseNameList sets UniverseNameList field to given value.

### HasUniverseNameList

`func (o *RestoreApiFilter) HasUniverseNameList() bool`

HasUniverseNameList returns a boolean if a field has been set.

### GetSourceUniverseNameList

`func (o *RestoreApiFilter) GetSourceUniverseNameList() []string`

GetSourceUniverseNameList returns the SourceUniverseNameList field if non-nil, zero value otherwise.

### GetSourceUniverseNameListOk

`func (o *RestoreApiFilter) GetSourceUniverseNameListOk() (*[]string, bool)`

GetSourceUniverseNameListOk returns a tuple with the SourceUniverseNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseNameList

`func (o *RestoreApiFilter) SetSourceUniverseNameList(v []string)`

SetSourceUniverseNameList sets SourceUniverseNameList field to given value.

### HasSourceUniverseNameList

`func (o *RestoreApiFilter) HasSourceUniverseNameList() bool`

HasSourceUniverseNameList returns a boolean if a field has been set.

### GetStorageConfigUuidList

`func (o *RestoreApiFilter) GetStorageConfigUuidList() []string`

GetStorageConfigUuidList returns the StorageConfigUuidList field if non-nil, zero value otherwise.

### GetStorageConfigUuidListOk

`func (o *RestoreApiFilter) GetStorageConfigUuidListOk() (*[]string, bool)`

GetStorageConfigUuidListOk returns a tuple with the StorageConfigUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUuidList

`func (o *RestoreApiFilter) SetStorageConfigUuidList(v []string)`

SetStorageConfigUuidList sets StorageConfigUuidList field to given value.

### HasStorageConfigUuidList

`func (o *RestoreApiFilter) HasStorageConfigUuidList() bool`

HasStorageConfigUuidList returns a boolean if a field has been set.

### GetUniverseUuidList

`func (o *RestoreApiFilter) GetUniverseUuidList() []string`

GetUniverseUuidList returns the UniverseUuidList field if non-nil, zero value otherwise.

### GetUniverseUuidListOk

`func (o *RestoreApiFilter) GetUniverseUuidListOk() (*[]string, bool)`

GetUniverseUuidListOk returns a tuple with the UniverseUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUuidList

`func (o *RestoreApiFilter) SetUniverseUuidList(v []string)`

SetUniverseUuidList sets UniverseUuidList field to given value.

### HasUniverseUuidList

`func (o *RestoreApiFilter) HasUniverseUuidList() bool`

HasUniverseUuidList returns a boolean if a field has been set.

### GetRestoreUuidList

`func (o *RestoreApiFilter) GetRestoreUuidList() []string`

GetRestoreUuidList returns the RestoreUuidList field if non-nil, zero value otherwise.

### GetRestoreUuidListOk

`func (o *RestoreApiFilter) GetRestoreUuidListOk() (*[]string, bool)`

GetRestoreUuidListOk returns a tuple with the RestoreUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUuidList

`func (o *RestoreApiFilter) SetRestoreUuidList(v []string)`

SetRestoreUuidList sets RestoreUuidList field to given value.

### HasRestoreUuidList

`func (o *RestoreApiFilter) HasRestoreUuidList() bool`

HasRestoreUuidList returns a boolean if a field has been set.

### GetOnlyShowDeletedSourceUniverses

`func (o *RestoreApiFilter) GetOnlyShowDeletedSourceUniverses() bool`

GetOnlyShowDeletedSourceUniverses returns the OnlyShowDeletedSourceUniverses field if non-nil, zero value otherwise.

### GetOnlyShowDeletedSourceUniversesOk

`func (o *RestoreApiFilter) GetOnlyShowDeletedSourceUniversesOk() (*bool, bool)`

GetOnlyShowDeletedSourceUniversesOk returns a tuple with the OnlyShowDeletedSourceUniverses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyShowDeletedSourceUniverses

`func (o *RestoreApiFilter) SetOnlyShowDeletedSourceUniverses(v bool)`

SetOnlyShowDeletedSourceUniverses sets OnlyShowDeletedSourceUniverses field to given value.

### HasOnlyShowDeletedSourceUniverses

`func (o *RestoreApiFilter) HasOnlyShowDeletedSourceUniverses() bool`

HasOnlyShowDeletedSourceUniverses returns a boolean if a field has been set.

### GetShowHidden

`func (o *RestoreApiFilter) GetShowHidden() bool`

GetShowHidden returns the ShowHidden field if non-nil, zero value otherwise.

### GetShowHiddenOk

`func (o *RestoreApiFilter) GetShowHiddenOk() (*bool, bool)`

GetShowHiddenOk returns a tuple with the ShowHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHidden

`func (o *RestoreApiFilter) SetShowHidden(v bool)`

SetShowHidden sets ShowHidden field to given value.

### HasShowHidden

`func (o *RestoreApiFilter) HasShowHidden() bool`

HasShowHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


