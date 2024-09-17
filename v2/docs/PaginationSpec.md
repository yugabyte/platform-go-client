# PaginationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | Start offset of the records. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of records to be fetched. | [optional] 
**Direction** | Pointer to **string** | Sort order of the records. | [optional] 

## Methods

### NewPaginationSpec

`func NewPaginationSpec() *PaginationSpec`

NewPaginationSpec instantiates a new PaginationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationSpecWithDefaults

`func NewPaginationSpecWithDefaults() *PaginationSpec`

NewPaginationSpecWithDefaults instantiates a new PaginationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *PaginationSpec) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaginationSpec) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaginationSpec) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PaginationSpec) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PaginationSpec) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationSpec) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationSpec) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationSpec) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetDirection

`func (o *PaginationSpec) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PaginationSpec) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PaginationSpec) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PaginationSpec) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


