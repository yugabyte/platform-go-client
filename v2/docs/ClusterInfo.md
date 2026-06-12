# ClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | cluster uuid | [optional] 
**GeoPartitioned** | Pointer to **bool** | Whether universe is treated as geo partitioned (i.e. tablespaces are created). | [optional] 
**SpotPrice** | Pointer to **float64** | TBD | [optional] 

## Methods

### NewClusterInfo

`func NewClusterInfo() *ClusterInfo`

NewClusterInfo instantiates a new ClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfoWithDefaults

`func NewClusterInfoWithDefaults() *ClusterInfo`

NewClusterInfoWithDefaults instantiates a new ClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ClusterInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ClusterInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ClusterInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ClusterInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetGeoPartitioned

`func (o *ClusterInfo) GetGeoPartitioned() bool`

GetGeoPartitioned returns the GeoPartitioned field if non-nil, zero value otherwise.

### GetGeoPartitionedOk

`func (o *ClusterInfo) GetGeoPartitionedOk() (*bool, bool)`

GetGeoPartitionedOk returns a tuple with the GeoPartitioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoPartitioned

`func (o *ClusterInfo) SetGeoPartitioned(v bool)`

SetGeoPartitioned sets GeoPartitioned field to given value.

### HasGeoPartitioned

`func (o *ClusterInfo) HasGeoPartitioned() bool`

HasGeoPartitioned returns a boolean if a field has been set.

### GetSpotPrice

`func (o *ClusterInfo) GetSpotPrice() float64`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *ClusterInfo) GetSpotPriceOk() (*float64, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *ClusterInfo) SetSpotPrice(v float64)`

SetSpotPrice sets SpotPrice field to given value.

### HasSpotPrice

`func (o *ClusterInfo) HasSpotPrice() bool`

HasSpotPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


