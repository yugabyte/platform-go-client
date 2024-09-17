# JobScheduleUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | **bool** | Disable or enable the job schedule. | 
**IntervalSecs** | **int64** | Delay interval between two jobs. | 
**SnoozeUntil** | Pointer to **time.Time** | Time until the job should be snoozed. | [optional] 
**Type** | [**JobScheduleType**](JobScheduleType.md) |  | 

## Methods

### NewJobScheduleUpdateSpec

`func NewJobScheduleUpdateSpec(disabled bool, intervalSecs int64, type_ JobScheduleType, ) *JobScheduleUpdateSpec`

NewJobScheduleUpdateSpec instantiates a new JobScheduleUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobScheduleUpdateSpecWithDefaults

`func NewJobScheduleUpdateSpecWithDefaults() *JobScheduleUpdateSpec`

NewJobScheduleUpdateSpecWithDefaults instantiates a new JobScheduleUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *JobScheduleUpdateSpec) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *JobScheduleUpdateSpec) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *JobScheduleUpdateSpec) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetIntervalSecs

`func (o *JobScheduleUpdateSpec) GetIntervalSecs() int64`

GetIntervalSecs returns the IntervalSecs field if non-nil, zero value otherwise.

### GetIntervalSecsOk

`func (o *JobScheduleUpdateSpec) GetIntervalSecsOk() (*int64, bool)`

GetIntervalSecsOk returns a tuple with the IntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSecs

`func (o *JobScheduleUpdateSpec) SetIntervalSecs(v int64)`

SetIntervalSecs sets IntervalSecs field to given value.


### GetSnoozeUntil

`func (o *JobScheduleUpdateSpec) GetSnoozeUntil() time.Time`

GetSnoozeUntil returns the SnoozeUntil field if non-nil, zero value otherwise.

### GetSnoozeUntilOk

`func (o *JobScheduleUpdateSpec) GetSnoozeUntilOk() (*time.Time, bool)`

GetSnoozeUntilOk returns a tuple with the SnoozeUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoozeUntil

`func (o *JobScheduleUpdateSpec) SetSnoozeUntil(v time.Time)`

SetSnoozeUntil sets SnoozeUntil field to given value.

### HasSnoozeUntil

`func (o *JobScheduleUpdateSpec) HasSnoozeUntil() bool`

HasSnoozeUntil returns a boolean if a field has been set.

### GetType

`func (o *JobScheduleUpdateSpec) GetType() JobScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobScheduleUpdateSpec) GetTypeOk() (*JobScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobScheduleUpdateSpec) SetType(v JobScheduleType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


