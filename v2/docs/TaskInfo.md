# TaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Task UUID. | [optional] [readonly] 
**Title** | Pointer to **string** | Human-readable task title. | [optional] [readonly] 
**Target** | Pointer to **string** | Target resource type name. | [optional] [readonly] 
**TargetUuid** | Pointer to **string** | UUID of the target resource. | [optional] [readonly] 
**Type** | Pointer to **string** | Task action type. | [optional] [readonly] 
**TypeName** | Pointer to **string** | Friendly task type name. | [optional] [readonly] 
**Details** | Pointer to [**TaskDetails**](TaskDetails.md) |  | [optional] 
**CorrelationId** | Pointer to **string** | Correlation id for the task, if any. | [optional] [readonly] 
**CreateTime** | Pointer to **time.Time** | Task creation time. | [optional] [readonly] 
**CompletionTime** | Pointer to **time.Time** | Task completion time, if finished. | [optional] [readonly] 
**Status** | Pointer to **string** | Current task status. | [optional] [readonly] 
**PercentComplete** | Pointer to **int32** | Percentage of task completion. | [optional] [readonly] 
**Abortable** | Pointer to **bool** | Whether the task can be aborted. | [optional] [readonly] 
**Retryable** | Pointer to **bool** | Whether the task can be retried. | [optional] [readonly] 
**CanRollback** | Pointer to **bool** | Whether the task can be rolled back. | [optional] [readonly] 
**OriginalTaskUuid** | Pointer to **string** | UUID of the first task in the retry/rollback chain (clean universe state), if any. Carried forward on retries and rollbacks. Distinct from previousTaskUUID inherit linkage.  | [optional] [readonly] 
**UserEmail** | Pointer to **string** | Email of the user who started the task. | [optional] [readonly] 

## Methods

### NewTaskInfo

`func NewTaskInfo() *TaskInfo`

NewTaskInfo instantiates a new TaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInfoWithDefaults

`func NewTaskInfoWithDefaults() *TaskInfo`

NewTaskInfoWithDefaults instantiates a new TaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *TaskInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TaskInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TaskInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TaskInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetTitle

`func (o *TaskInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTarget

`func (o *TaskInfo) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TaskInfo) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TaskInfo) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TaskInfo) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetUuid

`func (o *TaskInfo) GetTargetUuid() string`

GetTargetUuid returns the TargetUuid field if non-nil, zero value otherwise.

### GetTargetUuidOk

`func (o *TaskInfo) GetTargetUuidOk() (*string, bool)`

GetTargetUuidOk returns a tuple with the TargetUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUuid

`func (o *TaskInfo) SetTargetUuid(v string)`

SetTargetUuid sets TargetUuid field to given value.

### HasTargetUuid

`func (o *TaskInfo) HasTargetUuid() bool`

HasTargetUuid returns a boolean if a field has been set.

### GetType

`func (o *TaskInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeName

`func (o *TaskInfo) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *TaskInfo) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *TaskInfo) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *TaskInfo) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetDetails

`func (o *TaskInfo) GetDetails() TaskDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TaskInfo) GetDetailsOk() (*TaskDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TaskInfo) SetDetails(v TaskDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *TaskInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCorrelationId

`func (o *TaskInfo) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *TaskInfo) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *TaskInfo) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *TaskInfo) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetCreateTime

`func (o *TaskInfo) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *TaskInfo) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *TaskInfo) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *TaskInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *TaskInfo) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *TaskInfo) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *TaskInfo) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *TaskInfo) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetStatus

`func (o *TaskInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPercentComplete

`func (o *TaskInfo) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *TaskInfo) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *TaskInfo) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *TaskInfo) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetAbortable

`func (o *TaskInfo) GetAbortable() bool`

GetAbortable returns the Abortable field if non-nil, zero value otherwise.

### GetAbortableOk

`func (o *TaskInfo) GetAbortableOk() (*bool, bool)`

GetAbortableOk returns a tuple with the Abortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortable

`func (o *TaskInfo) SetAbortable(v bool)`

SetAbortable sets Abortable field to given value.

### HasAbortable

`func (o *TaskInfo) HasAbortable() bool`

HasAbortable returns a boolean if a field has been set.

### GetRetryable

`func (o *TaskInfo) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *TaskInfo) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *TaskInfo) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *TaskInfo) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetCanRollback

`func (o *TaskInfo) GetCanRollback() bool`

GetCanRollback returns the CanRollback field if non-nil, zero value otherwise.

### GetCanRollbackOk

`func (o *TaskInfo) GetCanRollbackOk() (*bool, bool)`

GetCanRollbackOk returns a tuple with the CanRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRollback

`func (o *TaskInfo) SetCanRollback(v bool)`

SetCanRollback sets CanRollback field to given value.

### HasCanRollback

`func (o *TaskInfo) HasCanRollback() bool`

HasCanRollback returns a boolean if a field has been set.

### GetOriginalTaskUuid

`func (o *TaskInfo) GetOriginalTaskUuid() string`

GetOriginalTaskUuid returns the OriginalTaskUuid field if non-nil, zero value otherwise.

### GetOriginalTaskUuidOk

`func (o *TaskInfo) GetOriginalTaskUuidOk() (*string, bool)`

GetOriginalTaskUuidOk returns a tuple with the OriginalTaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTaskUuid

`func (o *TaskInfo) SetOriginalTaskUuid(v string)`

SetOriginalTaskUuid sets OriginalTaskUuid field to given value.

### HasOriginalTaskUuid

`func (o *TaskInfo) HasOriginalTaskUuid() bool`

HasOriginalTaskUuid returns a boolean if a field has been set.

### GetUserEmail

`func (o *TaskInfo) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *TaskInfo) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *TaskInfo) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *TaskInfo) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


