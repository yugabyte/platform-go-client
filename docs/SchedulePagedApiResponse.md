# SchedulePagedApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]ScheduleResp**](ScheduleResp.md) |  | 
**HasNext** | **bool** |  | 
**HasPrev** | **bool** |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewSchedulePagedApiResponse

`func NewSchedulePagedApiResponse(entities []ScheduleResp, hasNext bool, hasPrev bool, totalCount int32, ) *SchedulePagedApiResponse`

NewSchedulePagedApiResponse instantiates a new SchedulePagedApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePagedApiResponseWithDefaults

`func NewSchedulePagedApiResponseWithDefaults() *SchedulePagedApiResponse`

NewSchedulePagedApiResponseWithDefaults instantiates a new SchedulePagedApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *SchedulePagedApiResponse) GetEntities() []ScheduleResp`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SchedulePagedApiResponse) GetEntitiesOk() (*[]ScheduleResp, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SchedulePagedApiResponse) SetEntities(v []ScheduleResp)`

SetEntities sets Entities field to given value.


### GetHasNext

`func (o *SchedulePagedApiResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *SchedulePagedApiResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *SchedulePagedApiResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrev

`func (o *SchedulePagedApiResponse) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *SchedulePagedApiResponse) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *SchedulePagedApiResponse) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.


### GetTotalCount

`func (o *SchedulePagedApiResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SchedulePagedApiResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SchedulePagedApiResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


