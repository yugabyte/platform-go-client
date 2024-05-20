# SchedulePagedApiQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** |  | 
**Filter** | [**ScheduleApiFilter**](ScheduleApiFilter.md) |  | 
**Limit** | **int32** |  | 
**NeedTotalCount** | **bool** |  | 
**Offset** | **int32** |  | 
**SortBy** | **string** |  | 

## Methods

### NewSchedulePagedApiQuery

`func NewSchedulePagedApiQuery(direction string, filter ScheduleApiFilter, limit int32, needTotalCount bool, offset int32, sortBy string, ) *SchedulePagedApiQuery`

NewSchedulePagedApiQuery instantiates a new SchedulePagedApiQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePagedApiQueryWithDefaults

`func NewSchedulePagedApiQueryWithDefaults() *SchedulePagedApiQuery`

NewSchedulePagedApiQueryWithDefaults instantiates a new SchedulePagedApiQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SchedulePagedApiQuery) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SchedulePagedApiQuery) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SchedulePagedApiQuery) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetFilter

`func (o *SchedulePagedApiQuery) GetFilter() ScheduleApiFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SchedulePagedApiQuery) GetFilterOk() (*ScheduleApiFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SchedulePagedApiQuery) SetFilter(v ScheduleApiFilter)`

SetFilter sets Filter field to given value.


### GetLimit

`func (o *SchedulePagedApiQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SchedulePagedApiQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SchedulePagedApiQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNeedTotalCount

`func (o *SchedulePagedApiQuery) GetNeedTotalCount() bool`

GetNeedTotalCount returns the NeedTotalCount field if non-nil, zero value otherwise.

### GetNeedTotalCountOk

`func (o *SchedulePagedApiQuery) GetNeedTotalCountOk() (*bool, bool)`

GetNeedTotalCountOk returns a tuple with the NeedTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedTotalCount

`func (o *SchedulePagedApiQuery) SetNeedTotalCount(v bool)`

SetNeedTotalCount sets NeedTotalCount field to given value.


### GetOffset

`func (o *SchedulePagedApiQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SchedulePagedApiQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SchedulePagedApiQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetSortBy

`func (o *SchedulePagedApiQuery) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *SchedulePagedApiQuery) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *SchedulePagedApiQuery) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


