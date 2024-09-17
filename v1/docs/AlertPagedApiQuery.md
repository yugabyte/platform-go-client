# AlertPagedApiQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** |  | 
**Filter** | [**AlertApiFilter**](AlertApiFilter.md) |  | 
**Limit** | **int32** |  | 
**NeedTotalCount** | **bool** |  | 
**Offset** | **int32** |  | 
**SortBy** | **string** |  | 

## Methods

### NewAlertPagedApiQuery

`func NewAlertPagedApiQuery(direction string, filter AlertApiFilter, limit int32, needTotalCount bool, offset int32, sortBy string, ) *AlertPagedApiQuery`

NewAlertPagedApiQuery instantiates a new AlertPagedApiQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertPagedApiQueryWithDefaults

`func NewAlertPagedApiQueryWithDefaults() *AlertPagedApiQuery`

NewAlertPagedApiQueryWithDefaults instantiates a new AlertPagedApiQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *AlertPagedApiQuery) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AlertPagedApiQuery) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AlertPagedApiQuery) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetFilter

`func (o *AlertPagedApiQuery) GetFilter() AlertApiFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AlertPagedApiQuery) GetFilterOk() (*AlertApiFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AlertPagedApiQuery) SetFilter(v AlertApiFilter)`

SetFilter sets Filter field to given value.


### GetLimit

`func (o *AlertPagedApiQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AlertPagedApiQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AlertPagedApiQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNeedTotalCount

`func (o *AlertPagedApiQuery) GetNeedTotalCount() bool`

GetNeedTotalCount returns the NeedTotalCount field if non-nil, zero value otherwise.

### GetNeedTotalCountOk

`func (o *AlertPagedApiQuery) GetNeedTotalCountOk() (*bool, bool)`

GetNeedTotalCountOk returns a tuple with the NeedTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedTotalCount

`func (o *AlertPagedApiQuery) SetNeedTotalCount(v bool)`

SetNeedTotalCount sets NeedTotalCount field to given value.


### GetOffset

`func (o *AlertPagedApiQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AlertPagedApiQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AlertPagedApiQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetSortBy

`func (o *AlertPagedApiQuery) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *AlertPagedApiQuery) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *AlertPagedApiQuery) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


