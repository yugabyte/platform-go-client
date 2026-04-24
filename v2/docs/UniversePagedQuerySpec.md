# UniversePagedQuerySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | Start offset of the records. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of records to be fetched. | [optional] 
**Direction** | Pointer to **string** | Sort order of the records. | [optional] 
**Filter** | Pointer to [**UniverseListApiFilter**](UniverseListApiFilter.md) |  | [optional] 

## Methods

### NewUniversePagedQuerySpec

`func NewUniversePagedQuerySpec() *UniversePagedQuerySpec`

NewUniversePagedQuerySpec instantiates a new UniversePagedQuerySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniversePagedQuerySpecWithDefaults

`func NewUniversePagedQuerySpecWithDefaults() *UniversePagedQuerySpec`

NewUniversePagedQuerySpecWithDefaults instantiates a new UniversePagedQuerySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *UniversePagedQuerySpec) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UniversePagedQuerySpec) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UniversePagedQuerySpec) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *UniversePagedQuerySpec) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *UniversePagedQuerySpec) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UniversePagedQuerySpec) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UniversePagedQuerySpec) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UniversePagedQuerySpec) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetDirection

`func (o *UniversePagedQuerySpec) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *UniversePagedQuerySpec) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *UniversePagedQuerySpec) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *UniversePagedQuerySpec) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFilter

`func (o *UniversePagedQuerySpec) GetFilter() UniverseListApiFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UniversePagedQuerySpec) GetFilterOk() (*UniverseListApiFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UniversePagedQuerySpec) SetFilter(v UniverseListApiFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UniversePagedQuerySpec) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


