# ClusterPartitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | System generated partition uuid used to lookup corresponding GeoPartitionSpec. This is not a user input. | [optional] [readonly] 
**Name** | **string** | The name of geo partition | 
**DefaultPartition** | **bool** | Whether the partition is default (all masters are put only into this partition) | 
**ReplicationFactor** | **int32** | The replication factor for the partition. | 
**Placement** | [**ClusterPlacementSpec**](ClusterPlacementSpec.md) |  | 

## Methods

### NewClusterPartitionSpec

`func NewClusterPartitionSpec(name string, defaultPartition bool, replicationFactor int32, placement ClusterPlacementSpec, ) *ClusterPartitionSpec`

NewClusterPartitionSpec instantiates a new ClusterPartitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPartitionSpecWithDefaults

`func NewClusterPartitionSpecWithDefaults() *ClusterPartitionSpec`

NewClusterPartitionSpecWithDefaults instantiates a new ClusterPartitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ClusterPartitionSpec) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ClusterPartitionSpec) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ClusterPartitionSpec) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ClusterPartitionSpec) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ClusterPartitionSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterPartitionSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterPartitionSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultPartition

`func (o *ClusterPartitionSpec) GetDefaultPartition() bool`

GetDefaultPartition returns the DefaultPartition field if non-nil, zero value otherwise.

### GetDefaultPartitionOk

`func (o *ClusterPartitionSpec) GetDefaultPartitionOk() (*bool, bool)`

GetDefaultPartitionOk returns a tuple with the DefaultPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPartition

`func (o *ClusterPartitionSpec) SetDefaultPartition(v bool)`

SetDefaultPartition sets DefaultPartition field to given value.


### GetReplicationFactor

`func (o *ClusterPartitionSpec) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *ClusterPartitionSpec) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *ClusterPartitionSpec) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.


### GetPlacement

`func (o *ClusterPartitionSpec) GetPlacement() ClusterPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *ClusterPartitionSpec) GetPlacementOk() (*ClusterPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *ClusterPartitionSpec) SetPlacement(v ClusterPlacementSpec)`

SetPlacement sets Placement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


