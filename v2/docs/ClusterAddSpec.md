# ClusterAddSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterType** | **string** | Cluster type can be one of READ_REPLICA, ADDON | 
**NumNodes** | **int32** | Set the number of nodes (tservers) to provision in this cluster | 
**NodeSpec** | [**ClusterNodeSpec**](ClusterNodeSpec.md) |  | 
**ProviderSpec** | [**ClusterProviderEditSpec**](ClusterProviderEditSpec.md) |  | 
**PlacementSpec** | Pointer to [**ClusterPlacementSpec**](ClusterPlacementSpec.md) |  | [optional] 
**PartitionsSpec** | Pointer to [**[]ClusterPartitionSpec**](ClusterPartitionSpec.md) |  | [optional] 
**InstanceTags** | Pointer to **map[string]string** | A map of strings representing a set of Tags and Values to apply on nodes in the aws/gcp/azu cloud. See https://docs.yugabyte.com/preview/yugabyte-platform/manage-deployments/instance-tags/. | [optional] 
**Gflags** | Pointer to [**ClusterGFlags**](ClusterGFlags.md) |  | [optional] 

## Methods

### NewClusterAddSpec

`func NewClusterAddSpec(clusterType string, numNodes int32, nodeSpec ClusterNodeSpec, providerSpec ClusterProviderEditSpec, ) *ClusterAddSpec`

NewClusterAddSpec instantiates a new ClusterAddSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAddSpecWithDefaults

`func NewClusterAddSpecWithDefaults() *ClusterAddSpec`

NewClusterAddSpecWithDefaults instantiates a new ClusterAddSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterType

`func (o *ClusterAddSpec) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ClusterAddSpec) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ClusterAddSpec) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.


### GetNumNodes

`func (o *ClusterAddSpec) GetNumNodes() int32`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *ClusterAddSpec) GetNumNodesOk() (*int32, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *ClusterAddSpec) SetNumNodes(v int32)`

SetNumNodes sets NumNodes field to given value.


### GetNodeSpec

`func (o *ClusterAddSpec) GetNodeSpec() ClusterNodeSpec`

GetNodeSpec returns the NodeSpec field if non-nil, zero value otherwise.

### GetNodeSpecOk

`func (o *ClusterAddSpec) GetNodeSpecOk() (*ClusterNodeSpec, bool)`

GetNodeSpecOk returns a tuple with the NodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSpec

`func (o *ClusterAddSpec) SetNodeSpec(v ClusterNodeSpec)`

SetNodeSpec sets NodeSpec field to given value.


### GetProviderSpec

`func (o *ClusterAddSpec) GetProviderSpec() ClusterProviderEditSpec`

GetProviderSpec returns the ProviderSpec field if non-nil, zero value otherwise.

### GetProviderSpecOk

`func (o *ClusterAddSpec) GetProviderSpecOk() (*ClusterProviderEditSpec, bool)`

GetProviderSpecOk returns a tuple with the ProviderSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSpec

`func (o *ClusterAddSpec) SetProviderSpec(v ClusterProviderEditSpec)`

SetProviderSpec sets ProviderSpec field to given value.


### GetPlacementSpec

`func (o *ClusterAddSpec) GetPlacementSpec() ClusterPlacementSpec`

GetPlacementSpec returns the PlacementSpec field if non-nil, zero value otherwise.

### GetPlacementSpecOk

`func (o *ClusterAddSpec) GetPlacementSpecOk() (*ClusterPlacementSpec, bool)`

GetPlacementSpecOk returns a tuple with the PlacementSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementSpec

`func (o *ClusterAddSpec) SetPlacementSpec(v ClusterPlacementSpec)`

SetPlacementSpec sets PlacementSpec field to given value.

### HasPlacementSpec

`func (o *ClusterAddSpec) HasPlacementSpec() bool`

HasPlacementSpec returns a boolean if a field has been set.

### GetPartitionsSpec

`func (o *ClusterAddSpec) GetPartitionsSpec() []ClusterPartitionSpec`

GetPartitionsSpec returns the PartitionsSpec field if non-nil, zero value otherwise.

### GetPartitionsSpecOk

`func (o *ClusterAddSpec) GetPartitionsSpecOk() (*[]ClusterPartitionSpec, bool)`

GetPartitionsSpecOk returns a tuple with the PartitionsSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionsSpec

`func (o *ClusterAddSpec) SetPartitionsSpec(v []ClusterPartitionSpec)`

SetPartitionsSpec sets PartitionsSpec field to given value.

### HasPartitionsSpec

`func (o *ClusterAddSpec) HasPartitionsSpec() bool`

HasPartitionsSpec returns a boolean if a field has been set.

### GetInstanceTags

`func (o *ClusterAddSpec) GetInstanceTags() map[string]string`

GetInstanceTags returns the InstanceTags field if non-nil, zero value otherwise.

### GetInstanceTagsOk

`func (o *ClusterAddSpec) GetInstanceTagsOk() (*map[string]string, bool)`

GetInstanceTagsOk returns a tuple with the InstanceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTags

`func (o *ClusterAddSpec) SetInstanceTags(v map[string]string)`

SetInstanceTags sets InstanceTags field to given value.

### HasInstanceTags

`func (o *ClusterAddSpec) HasInstanceTags() bool`

HasInstanceTags returns a boolean if a field has been set.

### GetGflags

`func (o *ClusterAddSpec) GetGflags() ClusterGFlags`

GetGflags returns the Gflags field if non-nil, zero value otherwise.

### GetGflagsOk

`func (o *ClusterAddSpec) GetGflagsOk() (*ClusterGFlags, bool)`

GetGflagsOk returns a tuple with the Gflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGflags

`func (o *ClusterAddSpec) SetGflags(v ClusterGFlags)`

SetGflags sets Gflags field to given value.

### HasGflags

`func (o *ClusterAddSpec) HasGflags() bool`

HasGflags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


