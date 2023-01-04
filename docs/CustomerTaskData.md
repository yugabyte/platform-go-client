# CustomerTaskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abortable** | Pointer to **bool** | Customer task abortable | [optional] 
**CompletionTime** | Pointer to **time.Time** | Customer task completion time | [optional] 
**CorrelationId** | Pointer to **string** | Correlation id | [optional] 
**CreateTime** | Pointer to **time.Time** | Customer task creation time | [optional] 
**Id** | Pointer to **string** | Customer task UUID | [optional] 
**PercentComplete** | Pointer to **int32** | Customer task percentage completed | [optional] 
**Retryable** | Pointer to **bool** | Customer task retryable | [optional] 
**Status** | Pointer to **string** | Customer task status | [optional] 
**Target** | Pointer to **string** | Customer task target | [optional] 
**TargetUUID** | Pointer to **string** | Customer task target UUID | [optional] 
**Title** | Pointer to **string** | Customer task title | [optional] 
**Type** | Pointer to **string** | Customer task type | [optional] 
**TypeName** | Pointer to **string** | Customer task type name | [optional] 

## Methods

### NewCustomerTaskData

`func NewCustomerTaskData() *CustomerTaskData`

NewCustomerTaskData instantiates a new CustomerTaskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerTaskDataWithDefaults

`func NewCustomerTaskDataWithDefaults() *CustomerTaskData`

NewCustomerTaskDataWithDefaults instantiates a new CustomerTaskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortable

`func (o *CustomerTaskData) GetAbortable() bool`

GetAbortable returns the Abortable field if non-nil, zero value otherwise.

### GetAbortableOk

`func (o *CustomerTaskData) GetAbortableOk() (*bool, bool)`

GetAbortableOk returns a tuple with the Abortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortable

`func (o *CustomerTaskData) SetAbortable(v bool)`

SetAbortable sets Abortable field to given value.

### HasAbortable

`func (o *CustomerTaskData) HasAbortable() bool`

HasAbortable returns a boolean if a field has been set.

### GetCompletionTime

`func (o *CustomerTaskData) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *CustomerTaskData) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *CustomerTaskData) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *CustomerTaskData) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetCorrelationId

`func (o *CustomerTaskData) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *CustomerTaskData) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *CustomerTaskData) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *CustomerTaskData) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetCreateTime

`func (o *CustomerTaskData) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CustomerTaskData) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CustomerTaskData) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CustomerTaskData) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetId

`func (o *CustomerTaskData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerTaskData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerTaskData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerTaskData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPercentComplete

`func (o *CustomerTaskData) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *CustomerTaskData) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *CustomerTaskData) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *CustomerTaskData) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetRetryable

`func (o *CustomerTaskData) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *CustomerTaskData) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *CustomerTaskData) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *CustomerTaskData) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerTaskData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerTaskData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerTaskData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerTaskData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *CustomerTaskData) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CustomerTaskData) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CustomerTaskData) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CustomerTaskData) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetUUID

`func (o *CustomerTaskData) GetTargetUUID() string`

GetTargetUUID returns the TargetUUID field if non-nil, zero value otherwise.

### GetTargetUUIDOk

`func (o *CustomerTaskData) GetTargetUUIDOk() (*string, bool)`

GetTargetUUIDOk returns a tuple with the TargetUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUUID

`func (o *CustomerTaskData) SetTargetUUID(v string)`

SetTargetUUID sets TargetUUID field to given value.

### HasTargetUUID

`func (o *CustomerTaskData) HasTargetUUID() bool`

HasTargetUUID returns a boolean if a field has been set.

### GetTitle

`func (o *CustomerTaskData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomerTaskData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomerTaskData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomerTaskData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CustomerTaskData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerTaskData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerTaskData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerTaskData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeName

`func (o *CustomerTaskData) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *CustomerTaskData) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *CustomerTaskData) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *CustomerTaskData) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


