# JobSchedulePagedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | There are more next records if true. | [optional] 
**HasPrev** | Pointer to **bool** | There are more previous records if true. | [optional] 
**TotalCount** | Pointer to **int32** | Total number of records. | [optional] 
**Entities** | Pointer to [**[]JobSchedule**](JobSchedule.md) |  | [optional] 

## Methods

### NewJobSchedulePagedResp

`func NewJobSchedulePagedResp() *JobSchedulePagedResp`

NewJobSchedulePagedResp instantiates a new JobSchedulePagedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSchedulePagedRespWithDefaults

`func NewJobSchedulePagedRespWithDefaults() *JobSchedulePagedResp`

NewJobSchedulePagedRespWithDefaults instantiates a new JobSchedulePagedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *JobSchedulePagedResp) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *JobSchedulePagedResp) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *JobSchedulePagedResp) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *JobSchedulePagedResp) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *JobSchedulePagedResp) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *JobSchedulePagedResp) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *JobSchedulePagedResp) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *JobSchedulePagedResp) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetTotalCount

`func (o *JobSchedulePagedResp) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *JobSchedulePagedResp) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *JobSchedulePagedResp) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *JobSchedulePagedResp) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetEntities

`func (o *JobSchedulePagedResp) GetEntities() []JobSchedule`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *JobSchedulePagedResp) GetEntitiesOk() (*[]JobSchedule, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *JobSchedulePagedResp) SetEntities(v []JobSchedule)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *JobSchedulePagedResp) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


