# NodeAgentPagedApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]NodeAgentResp**](NodeAgentResp.md) |  | 
**HasNext** | **bool** |  | 
**HasPrev** | **bool** |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewNodeAgentPagedApiResponse

`func NewNodeAgentPagedApiResponse(entities []NodeAgentResp, hasNext bool, hasPrev bool, totalCount int32, ) *NodeAgentPagedApiResponse`

NewNodeAgentPagedApiResponse instantiates a new NodeAgentPagedApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAgentPagedApiResponseWithDefaults

`func NewNodeAgentPagedApiResponseWithDefaults() *NodeAgentPagedApiResponse`

NewNodeAgentPagedApiResponseWithDefaults instantiates a new NodeAgentPagedApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *NodeAgentPagedApiResponse) GetEntities() []NodeAgentResp`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *NodeAgentPagedApiResponse) GetEntitiesOk() (*[]NodeAgentResp, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *NodeAgentPagedApiResponse) SetEntities(v []NodeAgentResp)`

SetEntities sets Entities field to given value.


### GetHasNext

`func (o *NodeAgentPagedApiResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *NodeAgentPagedApiResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *NodeAgentPagedApiResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrev

`func (o *NodeAgentPagedApiResponse) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *NodeAgentPagedApiResponse) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *NodeAgentPagedApiResponse) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.


### GetTotalCount

`func (o *NodeAgentPagedApiResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *NodeAgentPagedApiResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *NodeAgentPagedApiResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


