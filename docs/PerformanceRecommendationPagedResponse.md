# PerformanceRecommendationPagedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]PerformanceRecommendation**](PerformanceRecommendation.md) |  | 
**HasNext** | **bool** |  | 
**HasPrev** | **bool** |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewPerformanceRecommendationPagedResponse

`func NewPerformanceRecommendationPagedResponse(entities []PerformanceRecommendation, hasNext bool, hasPrev bool, totalCount int32, ) *PerformanceRecommendationPagedResponse`

NewPerformanceRecommendationPagedResponse instantiates a new PerformanceRecommendationPagedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceRecommendationPagedResponseWithDefaults

`func NewPerformanceRecommendationPagedResponseWithDefaults() *PerformanceRecommendationPagedResponse`

NewPerformanceRecommendationPagedResponseWithDefaults instantiates a new PerformanceRecommendationPagedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *PerformanceRecommendationPagedResponse) GetEntities() []PerformanceRecommendation`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PerformanceRecommendationPagedResponse) GetEntitiesOk() (*[]PerformanceRecommendation, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PerformanceRecommendationPagedResponse) SetEntities(v []PerformanceRecommendation)`

SetEntities sets Entities field to given value.


### GetHasNext

`func (o *PerformanceRecommendationPagedResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PerformanceRecommendationPagedResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PerformanceRecommendationPagedResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrev

`func (o *PerformanceRecommendationPagedResponse) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *PerformanceRecommendationPagedResponse) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *PerformanceRecommendationPagedResponse) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.


### GetTotalCount

`func (o *PerformanceRecommendationPagedResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PerformanceRecommendationPagedResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PerformanceRecommendationPagedResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


