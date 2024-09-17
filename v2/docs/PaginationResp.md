# PaginationResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | There are more next records if true. | [optional] 
**HasPrev** | Pointer to **bool** | There are more previous records if true. | [optional] 
**TotalCount** | Pointer to **int32** | Total number of records. | [optional] 

## Methods

### NewPaginationResp

`func NewPaginationResp() *PaginationResp`

NewPaginationResp instantiates a new PaginationResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationRespWithDefaults

`func NewPaginationRespWithDefaults() *PaginationResp`

NewPaginationRespWithDefaults instantiates a new PaginationResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *PaginationResp) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PaginationResp) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PaginationResp) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PaginationResp) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *PaginationResp) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *PaginationResp) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *PaginationResp) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *PaginationResp) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginationResp) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginationResp) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginationResp) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginationResp) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


