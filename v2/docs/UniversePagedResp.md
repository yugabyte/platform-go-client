# UniversePagedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | There are more next records if true. | [optional] 
**HasPrev** | Pointer to **bool** | There are more previous records if true. | [optional] 
**TotalCount** | Pointer to **int32** | Total number of records. | [optional] 
**Entities** | Pointer to [**[]Universe**](Universe.md) |  | [optional] 

## Methods

### NewUniversePagedResp

`func NewUniversePagedResp() *UniversePagedResp`

NewUniversePagedResp instantiates a new UniversePagedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniversePagedRespWithDefaults

`func NewUniversePagedRespWithDefaults() *UniversePagedResp`

NewUniversePagedRespWithDefaults instantiates a new UniversePagedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *UniversePagedResp) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *UniversePagedResp) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *UniversePagedResp) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *UniversePagedResp) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *UniversePagedResp) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *UniversePagedResp) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *UniversePagedResp) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *UniversePagedResp) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetTotalCount

`func (o *UniversePagedResp) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UniversePagedResp) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UniversePagedResp) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *UniversePagedResp) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetEntities

`func (o *UniversePagedResp) GetEntities() []Universe`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *UniversePagedResp) GetEntitiesOk() (*[]Universe, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *UniversePagedResp) SetEntities(v []Universe)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *UniversePagedResp) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


