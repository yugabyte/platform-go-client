# TaskApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRangeStart** | Pointer to **time.Time** | Include tasks created on or after this time (filters &#x60;create_time&#x60;). May be set alone (open-ended upper bound) or together with &#x60;date_range_end&#x60;.  | [optional] 
**DateRangeEnd** | Pointer to **time.Time** | Include tasks created on or before this time (filters &#x60;create_time&#x60;). May be set alone (open-ended lower bound) or together with &#x60;date_range_start&#x60;.  | [optional] 
**CompletionDateRangeStart** | Pointer to **time.Time** | Include tasks that completed on or after this time. May be set alone (open-ended upper bound) or together with &#x60;completion_date_range_end&#x60;. In-progress tasks don&#39;t have a completion time yet and are always excluded when either completion bound is set (SQL comparisons do not match nulls).  | [optional] 
**CompletionDateRangeEnd** | Pointer to **time.Time** | Include tasks that completed on or before this time. May be set alone (open-ended lower bound) or together with &#x60;completion_date_range_start&#x60;. In-progress tasks don&#39;t have a completion time yet and are always excluded when either completion bound is set (SQL comparisons do not match nulls).  | [optional] 
**TargetList** | Pointer to **[]string** | Filter by target resource types. | [optional] 
**TargetUuidList** | Pointer to **[]string** | Filter by target resource UUIDs. | [optional] 
**TypeList** | Pointer to **[]string** | Filter by task action types. | [optional] 
**TypeNameList** | Pointer to **[]string** | Filter by friendly task type names. | [optional] 
**Status** | Pointer to **[]string** | Filter by task statuses. | [optional] 

## Methods

### NewTaskApiFilter

`func NewTaskApiFilter() *TaskApiFilter`

NewTaskApiFilter instantiates a new TaskApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskApiFilterWithDefaults

`func NewTaskApiFilterWithDefaults() *TaskApiFilter`

NewTaskApiFilterWithDefaults instantiates a new TaskApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRangeStart

`func (o *TaskApiFilter) GetDateRangeStart() time.Time`

GetDateRangeStart returns the DateRangeStart field if non-nil, zero value otherwise.

### GetDateRangeStartOk

`func (o *TaskApiFilter) GetDateRangeStartOk() (*time.Time, bool)`

GetDateRangeStartOk returns a tuple with the DateRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRangeStart

`func (o *TaskApiFilter) SetDateRangeStart(v time.Time)`

SetDateRangeStart sets DateRangeStart field to given value.

### HasDateRangeStart

`func (o *TaskApiFilter) HasDateRangeStart() bool`

HasDateRangeStart returns a boolean if a field has been set.

### GetDateRangeEnd

`func (o *TaskApiFilter) GetDateRangeEnd() time.Time`

GetDateRangeEnd returns the DateRangeEnd field if non-nil, zero value otherwise.

### GetDateRangeEndOk

`func (o *TaskApiFilter) GetDateRangeEndOk() (*time.Time, bool)`

GetDateRangeEndOk returns a tuple with the DateRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRangeEnd

`func (o *TaskApiFilter) SetDateRangeEnd(v time.Time)`

SetDateRangeEnd sets DateRangeEnd field to given value.

### HasDateRangeEnd

`func (o *TaskApiFilter) HasDateRangeEnd() bool`

HasDateRangeEnd returns a boolean if a field has been set.

### GetCompletionDateRangeStart

`func (o *TaskApiFilter) GetCompletionDateRangeStart() time.Time`

GetCompletionDateRangeStart returns the CompletionDateRangeStart field if non-nil, zero value otherwise.

### GetCompletionDateRangeStartOk

`func (o *TaskApiFilter) GetCompletionDateRangeStartOk() (*time.Time, bool)`

GetCompletionDateRangeStartOk returns a tuple with the CompletionDateRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateRangeStart

`func (o *TaskApiFilter) SetCompletionDateRangeStart(v time.Time)`

SetCompletionDateRangeStart sets CompletionDateRangeStart field to given value.

### HasCompletionDateRangeStart

`func (o *TaskApiFilter) HasCompletionDateRangeStart() bool`

HasCompletionDateRangeStart returns a boolean if a field has been set.

### GetCompletionDateRangeEnd

`func (o *TaskApiFilter) GetCompletionDateRangeEnd() time.Time`

GetCompletionDateRangeEnd returns the CompletionDateRangeEnd field if non-nil, zero value otherwise.

### GetCompletionDateRangeEndOk

`func (o *TaskApiFilter) GetCompletionDateRangeEndOk() (*time.Time, bool)`

GetCompletionDateRangeEndOk returns a tuple with the CompletionDateRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateRangeEnd

`func (o *TaskApiFilter) SetCompletionDateRangeEnd(v time.Time)`

SetCompletionDateRangeEnd sets CompletionDateRangeEnd field to given value.

### HasCompletionDateRangeEnd

`func (o *TaskApiFilter) HasCompletionDateRangeEnd() bool`

HasCompletionDateRangeEnd returns a boolean if a field has been set.

### GetTargetList

`func (o *TaskApiFilter) GetTargetList() []string`

GetTargetList returns the TargetList field if non-nil, zero value otherwise.

### GetTargetListOk

`func (o *TaskApiFilter) GetTargetListOk() (*[]string, bool)`

GetTargetListOk returns a tuple with the TargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetList

`func (o *TaskApiFilter) SetTargetList(v []string)`

SetTargetList sets TargetList field to given value.

### HasTargetList

`func (o *TaskApiFilter) HasTargetList() bool`

HasTargetList returns a boolean if a field has been set.

### GetTargetUuidList

`func (o *TaskApiFilter) GetTargetUuidList() []string`

GetTargetUuidList returns the TargetUuidList field if non-nil, zero value otherwise.

### GetTargetUuidListOk

`func (o *TaskApiFilter) GetTargetUuidListOk() (*[]string, bool)`

GetTargetUuidListOk returns a tuple with the TargetUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUuidList

`func (o *TaskApiFilter) SetTargetUuidList(v []string)`

SetTargetUuidList sets TargetUuidList field to given value.

### HasTargetUuidList

`func (o *TaskApiFilter) HasTargetUuidList() bool`

HasTargetUuidList returns a boolean if a field has been set.

### GetTypeList

`func (o *TaskApiFilter) GetTypeList() []string`

GetTypeList returns the TypeList field if non-nil, zero value otherwise.

### GetTypeListOk

`func (o *TaskApiFilter) GetTypeListOk() (*[]string, bool)`

GetTypeListOk returns a tuple with the TypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeList

`func (o *TaskApiFilter) SetTypeList(v []string)`

SetTypeList sets TypeList field to given value.

### HasTypeList

`func (o *TaskApiFilter) HasTypeList() bool`

HasTypeList returns a boolean if a field has been set.

### GetTypeNameList

`func (o *TaskApiFilter) GetTypeNameList() []string`

GetTypeNameList returns the TypeNameList field if non-nil, zero value otherwise.

### GetTypeNameListOk

`func (o *TaskApiFilter) GetTypeNameListOk() (*[]string, bool)`

GetTypeNameListOk returns a tuple with the TypeNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeNameList

`func (o *TaskApiFilter) SetTypeNameList(v []string)`

SetTypeNameList sets TypeNameList field to given value.

### HasTypeNameList

`func (o *TaskApiFilter) HasTypeNameList() bool`

HasTypeNameList returns a boolean if a field has been set.

### GetStatus

`func (o *TaskApiFilter) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskApiFilter) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskApiFilter) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskApiFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


