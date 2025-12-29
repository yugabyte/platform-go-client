# PlacementCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The cloud provider id. | [optional] 
**Code** | Pointer to **string** | The cloud provider code. | [optional] 
**RegionList** | Pointer to [**[]PlacementRegion**](PlacementRegion.md) | The list of region in this cloud we want to place data in. | [optional] 
**DefaultRegion** | Pointer to **string** | UUID of default region. For universes with more AZs than RF, the default placement for user tables will be RF AZs in the default region. This is commonly encountered in geo-partitioning use cases. Deprecated, replaced with default partition. | [optional] 
**MastersInDefaultRegion** | Pointer to **bool** | Whether to place all masters in the default region. Defaults to true. | [optional] [default to true]

## Methods

### NewPlacementCloud

`func NewPlacementCloud() *PlacementCloud`

NewPlacementCloud instantiates a new PlacementCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementCloudWithDefaults

`func NewPlacementCloudWithDefaults() *PlacementCloud`

NewPlacementCloudWithDefaults instantiates a new PlacementCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PlacementCloud) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PlacementCloud) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PlacementCloud) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PlacementCloud) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCode

`func (o *PlacementCloud) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PlacementCloud) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PlacementCloud) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PlacementCloud) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRegionList

`func (o *PlacementCloud) GetRegionList() []PlacementRegion`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *PlacementCloud) GetRegionListOk() (*[]PlacementRegion, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *PlacementCloud) SetRegionList(v []PlacementRegion)`

SetRegionList sets RegionList field to given value.

### HasRegionList

`func (o *PlacementCloud) HasRegionList() bool`

HasRegionList returns a boolean if a field has been set.

### GetDefaultRegion

`func (o *PlacementCloud) GetDefaultRegion() string`

GetDefaultRegion returns the DefaultRegion field if non-nil, zero value otherwise.

### GetDefaultRegionOk

`func (o *PlacementCloud) GetDefaultRegionOk() (*string, bool)`

GetDefaultRegionOk returns a tuple with the DefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRegion

`func (o *PlacementCloud) SetDefaultRegion(v string)`

SetDefaultRegion sets DefaultRegion field to given value.

### HasDefaultRegion

`func (o *PlacementCloud) HasDefaultRegion() bool`

HasDefaultRegion returns a boolean if a field has been set.

### GetMastersInDefaultRegion

`func (o *PlacementCloud) GetMastersInDefaultRegion() bool`

GetMastersInDefaultRegion returns the MastersInDefaultRegion field if non-nil, zero value otherwise.

### GetMastersInDefaultRegionOk

`func (o *PlacementCloud) GetMastersInDefaultRegionOk() (*bool, bool)`

GetMastersInDefaultRegionOk returns a tuple with the MastersInDefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastersInDefaultRegion

`func (o *PlacementCloud) SetMastersInDefaultRegion(v bool)`

SetMastersInDefaultRegion sets MastersInDefaultRegion field to given value.

### HasMastersInDefaultRegion

`func (o *PlacementCloud) HasMastersInDefaultRegion() bool`

HasMastersInDefaultRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


