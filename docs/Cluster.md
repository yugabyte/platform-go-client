# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterType** | **string** |  | 
**Index** | Pointer to **int32** |  | [optional] 
**Partitions** | Pointer to [**[]PartitionInfo**](PartitionInfo.md) | WARNING: This is a preview API that could change. Geo partitions for cluster | [optional] 
**PlacementInfo** | Pointer to [**PlacementInfo**](PlacementInfo.md) |  | [optional] 
**Regions** | Pointer to [**[]Region**](Region.md) |  | [optional] [readonly] 
**UserIntent** | [**UserIntent**](UserIntent.md) |  | 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewCluster

`func NewCluster(clusterType string, userIntent UserIntent, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterType

`func (o *Cluster) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *Cluster) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *Cluster) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.


### GetIndex

`func (o *Cluster) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Cluster) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Cluster) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Cluster) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetPartitions

`func (o *Cluster) GetPartitions() []PartitionInfo`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *Cluster) GetPartitionsOk() (*[]PartitionInfo, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *Cluster) SetPartitions(v []PartitionInfo)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *Cluster) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetPlacementInfo

`func (o *Cluster) GetPlacementInfo() PlacementInfo`

GetPlacementInfo returns the PlacementInfo field if non-nil, zero value otherwise.

### GetPlacementInfoOk

`func (o *Cluster) GetPlacementInfoOk() (*PlacementInfo, bool)`

GetPlacementInfoOk returns a tuple with the PlacementInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementInfo

`func (o *Cluster) SetPlacementInfo(v PlacementInfo)`

SetPlacementInfo sets PlacementInfo field to given value.

### HasPlacementInfo

`func (o *Cluster) HasPlacementInfo() bool`

HasPlacementInfo returns a boolean if a field has been set.

### GetRegions

`func (o *Cluster) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Cluster) GetRegionsOk() (*[]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Cluster) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Cluster) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetUserIntent

`func (o *Cluster) GetUserIntent() UserIntent`

GetUserIntent returns the UserIntent field if non-nil, zero value otherwise.

### GetUserIntentOk

`func (o *Cluster) GetUserIntentOk() (*UserIntent, bool)`

GetUserIntentOk returns a tuple with the UserIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIntent

`func (o *Cluster) SetUserIntent(v UserIntent)`

SetUserIntent sets UserIntent field to given value.


### GetUuid

`func (o *Cluster) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Cluster) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Cluster) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Cluster) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


