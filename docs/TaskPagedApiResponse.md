# TaskPagedApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]CustomerTaskData**](CustomerTaskData.md) |  | 
**HasNext** | **bool** |  | 
**HasPrev** | **bool** |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewTaskPagedApiResponse

`func NewTaskPagedApiResponse(entities []CustomerTaskData, hasNext bool, hasPrev bool, totalCount int32, ) *TaskPagedApiResponse`

NewTaskPagedApiResponse instantiates a new TaskPagedApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPagedApiResponseWithDefaults

`func NewTaskPagedApiResponseWithDefaults() *TaskPagedApiResponse`

NewTaskPagedApiResponseWithDefaults instantiates a new TaskPagedApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *TaskPagedApiResponse) GetEntities() []CustomerTaskData`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TaskPagedApiResponse) GetEntitiesOk() (*[]CustomerTaskData, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *TaskPagedApiResponse) SetEntities(v []CustomerTaskData)`

SetEntities sets Entities field to given value.


### GetHasNext

`func (o *TaskPagedApiResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *TaskPagedApiResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *TaskPagedApiResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrev

`func (o *TaskPagedApiResponse) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *TaskPagedApiResponse) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *TaskPagedApiResponse) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.


### GetTotalCount

`func (o *TaskPagedApiResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TaskPagedApiResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TaskPagedApiResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


