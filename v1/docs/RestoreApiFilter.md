# RestoreApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRangeEnd** | Pointer to **time.Time** | The end date to filter paged query. | [optional] 
**DateRangeStart** | Pointer to **time.Time** | The start date to filter paged query. | [optional] 
**OnlyShowDeletedSourceUniverses** | **bool** |  | 
**RestoreUUIDList** | **[]string** |  | 
**SourceUniverseNameList** | **[]string** |  | 
**States** | **[]string** |  | 
**StorageConfigUUIDList** | **[]string** |  | 
**UniverseNameList** | **[]string** |  | 
**UniverseUUIDList** | **[]string** |  | 

## Methods

### NewRestoreApiFilter

`func NewRestoreApiFilter(onlyShowDeletedSourceUniverses bool, restoreUUIDList []string, sourceUniverseNameList []string, states []string, storageConfigUUIDList []string, universeNameList []string, universeUUIDList []string, ) *RestoreApiFilter`

NewRestoreApiFilter instantiates a new RestoreApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreApiFilterWithDefaults

`func NewRestoreApiFilterWithDefaults() *RestoreApiFilter`

NewRestoreApiFilterWithDefaults instantiates a new RestoreApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetRestoreUUIDList

`func (o *RestoreApiFilter) GetRestoreUUIDList() []string`

GetRestoreUUIDList returns the RestoreUUIDList field if non-nil, zero value otherwise.

### GetRestoreUUIDListOk

`func (o *RestoreApiFilter) GetRestoreUUIDListOk() (*[]string, bool)`

GetRestoreUUIDListOk returns a tuple with the RestoreUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUUIDList

`func (o *RestoreApiFilter) SetRestoreUUIDList(v []string)`

SetRestoreUUIDList sets RestoreUUIDList field to given value.


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


### GetStates

`func (o *RestoreApiFilter) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *RestoreApiFilter) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *RestoreApiFilter) SetStates(v []string)`

SetStates sets States field to given value.


### GetStorageConfigUUIDList

`func (o *RestoreApiFilter) GetStorageConfigUUIDList() []string`

GetStorageConfigUUIDList returns the StorageConfigUUIDList field if non-nil, zero value otherwise.

### GetStorageConfigUUIDListOk

`func (o *RestoreApiFilter) GetStorageConfigUUIDListOk() (*[]string, bool)`

GetStorageConfigUUIDListOk returns a tuple with the StorageConfigUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUIDList

`func (o *RestoreApiFilter) SetStorageConfigUUIDList(v []string)`

SetStorageConfigUUIDList sets StorageConfigUUIDList field to given value.


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


### GetUniverseUUIDList

`func (o *RestoreApiFilter) GetUniverseUUIDList() []string`

GetUniverseUUIDList returns the UniverseUUIDList field if non-nil, zero value otherwise.

### GetUniverseUUIDListOk

`func (o *RestoreApiFilter) GetUniverseUUIDListOk() (*[]string, bool)`

GetUniverseUUIDListOk returns a tuple with the UniverseUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUIDList

`func (o *RestoreApiFilter) SetUniverseUUIDList(v []string)`

SetUniverseUUIDList sets UniverseUUIDList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


