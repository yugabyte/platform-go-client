# UniversePerformanceAdvisorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**EndTime** | Pointer to **time.Time** | Time perf advisor run was finished | [optional] [readonly] 
**Manual** | Pointer to **bool** | Scheduled or manual run | [optional] [readonly] 
**ScheduleTime** | Pointer to **time.Time** | Time perf advisor run was scheduled | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | Time perf advisor run was started | [optional] [readonly] 
**State** | **string** | State | [readonly] 
**UniverseUUID** | **string** | Universe UUID | [readonly] 
**Uuid** | **string** | Perf advisor run UUID | [readonly] 

## Methods

### NewUniversePerformanceAdvisorStatus

`func NewUniversePerformanceAdvisorStatus(customerUUID string, state string, universeUUID string, uuid string, ) *UniversePerformanceAdvisorStatus`

NewUniversePerformanceAdvisorStatus instantiates a new UniversePerformanceAdvisorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniversePerformanceAdvisorStatusWithDefaults

`func NewUniversePerformanceAdvisorStatusWithDefaults() *UniversePerformanceAdvisorStatus`

NewUniversePerformanceAdvisorStatusWithDefaults instantiates a new UniversePerformanceAdvisorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerUUID

`func (o *UniversePerformanceAdvisorStatus) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *UniversePerformanceAdvisorStatus) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *UniversePerformanceAdvisorStatus) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetEndTime

`func (o *UniversePerformanceAdvisorStatus) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UniversePerformanceAdvisorStatus) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UniversePerformanceAdvisorStatus) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UniversePerformanceAdvisorStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetManual

`func (o *UniversePerformanceAdvisorStatus) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *UniversePerformanceAdvisorStatus) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *UniversePerformanceAdvisorStatus) SetManual(v bool)`

SetManual sets Manual field to given value.

### HasManual

`func (o *UniversePerformanceAdvisorStatus) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetScheduleTime

`func (o *UniversePerformanceAdvisorStatus) GetScheduleTime() time.Time`

GetScheduleTime returns the ScheduleTime field if non-nil, zero value otherwise.

### GetScheduleTimeOk

`func (o *UniversePerformanceAdvisorStatus) GetScheduleTimeOk() (*time.Time, bool)`

GetScheduleTimeOk returns a tuple with the ScheduleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTime

`func (o *UniversePerformanceAdvisorStatus) SetScheduleTime(v time.Time)`

SetScheduleTime sets ScheduleTime field to given value.

### HasScheduleTime

`func (o *UniversePerformanceAdvisorStatus) HasScheduleTime() bool`

HasScheduleTime returns a boolean if a field has been set.

### GetStartTime

`func (o *UniversePerformanceAdvisorStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UniversePerformanceAdvisorStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UniversePerformanceAdvisorStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UniversePerformanceAdvisorStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *UniversePerformanceAdvisorStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UniversePerformanceAdvisorStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UniversePerformanceAdvisorStatus) SetState(v string)`

SetState sets State field to given value.


### GetUniverseUUID

`func (o *UniversePerformanceAdvisorStatus) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *UniversePerformanceAdvisorStatus) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *UniversePerformanceAdvisorStatus) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.


### GetUuid

`func (o *UniversePerformanceAdvisorStatus) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UniversePerformanceAdvisorStatus) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UniversePerformanceAdvisorStatus) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


