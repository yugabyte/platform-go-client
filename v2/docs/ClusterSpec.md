# ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | System generated cluster uuid used to lookup corresponding ClusterInfo. This is not a user input. | [optional] [readonly] 
**ClusterType** | **string** | Cluster type can be one of PRIMARY, ASYNC (for ReadOnly), ADDON | 
**NumNodes** | **int32** | The number of nodes (tservers) to provision in this cluster | 
**ReplicationFactor** | Pointer to **int32** | The number of copies of data to maintain in this cluster. Defaults to 3. | [optional] [default to 3]
**NodeSpec** | [**ClusterNodeSpec**](ClusterNodeSpec.md) |  | 
**NetworkingSpec** | Pointer to [**ClusterNetworkingSpec**](ClusterNetworkingSpec.md) |  | [optional] 
**ProviderSpec** | [**ClusterProviderSpec**](ClusterProviderSpec.md) |  | 
**PlacementSpec** | Pointer to [**ClusterPlacementSpec**](ClusterPlacementSpec.md) |  | [optional] 
**UseSpotInstance** | Pointer to **bool** | Whether to use spot instances for nodes in aws/gcp. Used in dev/test environments. | [optional] 
**InstanceTags** | Pointer to **map[string]string** | A map of strings representing a set of Tags and Values to apply on nodes in the aws/gcp/azu cloud. See https://docs.yugabyte.com/preview/yugabyte-platform/manage-deployments/instance-tags/. | [optional] 
**AuditLogConfig** | Pointer to [**AuditLogConfig**](AuditLogConfig.md) |  | [optional] 
**Gflags** | Pointer to [**ClusterGFlags**](ClusterGFlags.md) |  | [optional] 

## Methods

### NewClusterSpec

`func NewClusterSpec(clusterType string, numNodes int32, nodeSpec ClusterNodeSpec, providerSpec ClusterProviderSpec, ) *ClusterSpec`

NewClusterSpec instantiates a new ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSpecWithDefaults

`func NewClusterSpecWithDefaults() *ClusterSpec`

NewClusterSpecWithDefaults instantiates a new ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ClusterSpec) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ClusterSpec) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ClusterSpec) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ClusterSpec) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetClusterType

`func (o *ClusterSpec) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ClusterSpec) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ClusterSpec) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.


### GetNumNodes

`func (o *ClusterSpec) GetNumNodes() int32`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *ClusterSpec) GetNumNodesOk() (*int32, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *ClusterSpec) SetNumNodes(v int32)`

SetNumNodes sets NumNodes field to given value.


### GetReplicationFactor

`func (o *ClusterSpec) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *ClusterSpec) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *ClusterSpec) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *ClusterSpec) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetNodeSpec

`func (o *ClusterSpec) GetNodeSpec() ClusterNodeSpec`

GetNodeSpec returns the NodeSpec field if non-nil, zero value otherwise.

### GetNodeSpecOk

`func (o *ClusterSpec) GetNodeSpecOk() (*ClusterNodeSpec, bool)`

GetNodeSpecOk returns a tuple with the NodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSpec

`func (o *ClusterSpec) SetNodeSpec(v ClusterNodeSpec)`

SetNodeSpec sets NodeSpec field to given value.


### GetNetworkingSpec

`func (o *ClusterSpec) GetNetworkingSpec() ClusterNetworkingSpec`

GetNetworkingSpec returns the NetworkingSpec field if non-nil, zero value otherwise.

### GetNetworkingSpecOk

`func (o *ClusterSpec) GetNetworkingSpecOk() (*ClusterNetworkingSpec, bool)`

GetNetworkingSpecOk returns a tuple with the NetworkingSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingSpec

`func (o *ClusterSpec) SetNetworkingSpec(v ClusterNetworkingSpec)`

SetNetworkingSpec sets NetworkingSpec field to given value.

### HasNetworkingSpec

`func (o *ClusterSpec) HasNetworkingSpec() bool`

HasNetworkingSpec returns a boolean if a field has been set.

### GetProviderSpec

`func (o *ClusterSpec) GetProviderSpec() ClusterProviderSpec`

GetProviderSpec returns the ProviderSpec field if non-nil, zero value otherwise.

### GetProviderSpecOk

`func (o *ClusterSpec) GetProviderSpecOk() (*ClusterProviderSpec, bool)`

GetProviderSpecOk returns a tuple with the ProviderSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSpec

`func (o *ClusterSpec) SetProviderSpec(v ClusterProviderSpec)`

SetProviderSpec sets ProviderSpec field to given value.


### GetPlacementSpec

`func (o *ClusterSpec) GetPlacementSpec() ClusterPlacementSpec`

GetPlacementSpec returns the PlacementSpec field if non-nil, zero value otherwise.

### GetPlacementSpecOk

`func (o *ClusterSpec) GetPlacementSpecOk() (*ClusterPlacementSpec, bool)`

GetPlacementSpecOk returns a tuple with the PlacementSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementSpec

`func (o *ClusterSpec) SetPlacementSpec(v ClusterPlacementSpec)`

SetPlacementSpec sets PlacementSpec field to given value.

### HasPlacementSpec

`func (o *ClusterSpec) HasPlacementSpec() bool`

HasPlacementSpec returns a boolean if a field has been set.

### GetUseSpotInstance

`func (o *ClusterSpec) GetUseSpotInstance() bool`

GetUseSpotInstance returns the UseSpotInstance field if non-nil, zero value otherwise.

### GetUseSpotInstanceOk

`func (o *ClusterSpec) GetUseSpotInstanceOk() (*bool, bool)`

GetUseSpotInstanceOk returns a tuple with the UseSpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpotInstance

`func (o *ClusterSpec) SetUseSpotInstance(v bool)`

SetUseSpotInstance sets UseSpotInstance field to given value.

### HasUseSpotInstance

`func (o *ClusterSpec) HasUseSpotInstance() bool`

HasUseSpotInstance returns a boolean if a field has been set.

### GetInstanceTags

`func (o *ClusterSpec) GetInstanceTags() map[string]string`

GetInstanceTags returns the InstanceTags field if non-nil, zero value otherwise.

### GetInstanceTagsOk

`func (o *ClusterSpec) GetInstanceTagsOk() (*map[string]string, bool)`

GetInstanceTagsOk returns a tuple with the InstanceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTags

`func (o *ClusterSpec) SetInstanceTags(v map[string]string)`

SetInstanceTags sets InstanceTags field to given value.

### HasInstanceTags

`func (o *ClusterSpec) HasInstanceTags() bool`

HasInstanceTags returns a boolean if a field has been set.

### GetAuditLogConfig

`func (o *ClusterSpec) GetAuditLogConfig() AuditLogConfig`

GetAuditLogConfig returns the AuditLogConfig field if non-nil, zero value otherwise.

### GetAuditLogConfigOk

`func (o *ClusterSpec) GetAuditLogConfigOk() (*AuditLogConfig, bool)`

GetAuditLogConfigOk returns a tuple with the AuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogConfig

`func (o *ClusterSpec) SetAuditLogConfig(v AuditLogConfig)`

SetAuditLogConfig sets AuditLogConfig field to given value.

### HasAuditLogConfig

`func (o *ClusterSpec) HasAuditLogConfig() bool`

HasAuditLogConfig returns a boolean if a field has been set.

### GetGflags

`func (o *ClusterSpec) GetGflags() ClusterGFlags`

GetGflags returns the Gflags field if non-nil, zero value otherwise.

### GetGflagsOk

`func (o *ClusterSpec) GetGflagsOk() (*ClusterGFlags, bool)`

GetGflagsOk returns a tuple with the Gflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGflags

`func (o *ClusterSpec) SetGflags(v ClusterGFlags)`

SetGflags sets Gflags field to given value.

### HasGflags

`func (o *ClusterSpec) HasGflags() bool`

HasGflags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


