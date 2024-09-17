# JobScheduleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Creation time | [optional] [readonly] 
**CustomerUuid** | Pointer to **string** | Customer UUID | [optional] [readonly] 
**ExecutionCount** | Pointer to **int64** | Total execution count | [optional] [readonly] 
**FailedCount** | Pointer to **int64** | Total failed count | [optional] [readonly] 
**LastEndTime** | Pointer to **time.Time** | Last end time | [optional] [readonly] 
**LastJobInstanceUuid** | Pointer to **string** | Last job instance UUID | [optional] [readonly] 
**LastStartTime** | Pointer to **time.Time** | Last start time | [optional] [readonly] 
**NextStartTime** | Pointer to **time.Time** | Next start time | [optional] [readonly] 
**State** | Pointer to [**JobScheduleState**](JobScheduleState.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Updated time | [optional] [readonly] 
**Uuid** | Pointer to **string** | Job schedule UUID | [optional] [readonly] 

## Methods

### NewJobScheduleInfo

`func NewJobScheduleInfo() *JobScheduleInfo`

NewJobScheduleInfo instantiates a new JobScheduleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobScheduleInfoWithDefaults

`func NewJobScheduleInfoWithDefaults() *JobScheduleInfo`

NewJobScheduleInfoWithDefaults instantiates a new JobScheduleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *JobScheduleInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobScheduleInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobScheduleInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *JobScheduleInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *JobScheduleInfo) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *JobScheduleInfo) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *JobScheduleInfo) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *JobScheduleInfo) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetExecutionCount

`func (o *JobScheduleInfo) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *JobScheduleInfo) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *JobScheduleInfo) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *JobScheduleInfo) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *JobScheduleInfo) GetFailedCount() int64`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *JobScheduleInfo) GetFailedCountOk() (*int64, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *JobScheduleInfo) SetFailedCount(v int64)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *JobScheduleInfo) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetLastEndTime

`func (o *JobScheduleInfo) GetLastEndTime() time.Time`

GetLastEndTime returns the LastEndTime field if non-nil, zero value otherwise.

### GetLastEndTimeOk

`func (o *JobScheduleInfo) GetLastEndTimeOk() (*time.Time, bool)`

GetLastEndTimeOk returns a tuple with the LastEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEndTime

`func (o *JobScheduleInfo) SetLastEndTime(v time.Time)`

SetLastEndTime sets LastEndTime field to given value.

### HasLastEndTime

`func (o *JobScheduleInfo) HasLastEndTime() bool`

HasLastEndTime returns a boolean if a field has been set.

### GetLastJobInstanceUuid

`func (o *JobScheduleInfo) GetLastJobInstanceUuid() string`

GetLastJobInstanceUuid returns the LastJobInstanceUuid field if non-nil, zero value otherwise.

### GetLastJobInstanceUuidOk

`func (o *JobScheduleInfo) GetLastJobInstanceUuidOk() (*string, bool)`

GetLastJobInstanceUuidOk returns a tuple with the LastJobInstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastJobInstanceUuid

`func (o *JobScheduleInfo) SetLastJobInstanceUuid(v string)`

SetLastJobInstanceUuid sets LastJobInstanceUuid field to given value.

### HasLastJobInstanceUuid

`func (o *JobScheduleInfo) HasLastJobInstanceUuid() bool`

HasLastJobInstanceUuid returns a boolean if a field has been set.

### GetLastStartTime

`func (o *JobScheduleInfo) GetLastStartTime() time.Time`

GetLastStartTime returns the LastStartTime field if non-nil, zero value otherwise.

### GetLastStartTimeOk

`func (o *JobScheduleInfo) GetLastStartTimeOk() (*time.Time, bool)`

GetLastStartTimeOk returns a tuple with the LastStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStartTime

`func (o *JobScheduleInfo) SetLastStartTime(v time.Time)`

SetLastStartTime sets LastStartTime field to given value.

### HasLastStartTime

`func (o *JobScheduleInfo) HasLastStartTime() bool`

HasLastStartTime returns a boolean if a field has been set.

### GetNextStartTime

`func (o *JobScheduleInfo) GetNextStartTime() time.Time`

GetNextStartTime returns the NextStartTime field if non-nil, zero value otherwise.

### GetNextStartTimeOk

`func (o *JobScheduleInfo) GetNextStartTimeOk() (*time.Time, bool)`

GetNextStartTimeOk returns a tuple with the NextStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartTime

`func (o *JobScheduleInfo) SetNextStartTime(v time.Time)`

SetNextStartTime sets NextStartTime field to given value.

### HasNextStartTime

`func (o *JobScheduleInfo) HasNextStartTime() bool`

HasNextStartTime returns a boolean if a field has been set.

### GetState

`func (o *JobScheduleInfo) GetState() JobScheduleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JobScheduleInfo) GetStateOk() (*JobScheduleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JobScheduleInfo) SetState(v JobScheduleState)`

SetState sets State field to given value.

### HasState

`func (o *JobScheduleInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *JobScheduleInfo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *JobScheduleInfo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *JobScheduleInfo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *JobScheduleInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *JobScheduleInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *JobScheduleInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *JobScheduleInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *JobScheduleInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


