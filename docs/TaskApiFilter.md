# TaskApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRangeEnd** | Pointer to **time.Time** | The end date to filter paged query. | [optional] 
**DateRangeStart** | Pointer to **time.Time** | The start date to filter paged query. | [optional] 
**Status** | **[]string** |  | 
**TargetList** | **[]string** |  | 
**TargetUUIDList** | **[]string** |  | 
**TypeList** | **[]string** |  | 
**TypeNameList** | **[]string** |  | 

## Methods

### NewTaskApiFilter

`func NewTaskApiFilter(status []string, targetList []string, targetUUIDList []string, typeList []string, typeNameList []string, ) *TaskApiFilter`

NewTaskApiFilter instantiates a new TaskApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskApiFilterWithDefaults

`func NewTaskApiFilterWithDefaults() *TaskApiFilter`

NewTaskApiFilterWithDefaults instantiates a new TaskApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetTargetUUIDList

`func (o *TaskApiFilter) GetTargetUUIDList() []string`

GetTargetUUIDList returns the TargetUUIDList field if non-nil, zero value otherwise.

### GetTargetUUIDListOk

`func (o *TaskApiFilter) GetTargetUUIDListOk() (*[]string, bool)`

GetTargetUUIDListOk returns a tuple with the TargetUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUUIDList

`func (o *TaskApiFilter) SetTargetUUIDList(v []string)`

SetTargetUUIDList sets TargetUUIDList field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


