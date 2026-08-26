# UniverseResizeNodesCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Cluster UUID | 
**NodeSpec** | Pointer to [**ClusterResizeNodeSpec**](ClusterResizeNodeSpec.md) |  | [optional] 
**ProviderNodesSpecs** | Pointer to [**[]PerProviderResizeNodesSpec**](PerProviderResizeNodesSpec.md) | Proposed resize node settings per provider for multicloud clusters. | [optional] 
**Gflags** | Pointer to [**ClusterGFlags**](ClusterGFlags.md) |  | [optional] 

## Methods

### NewUniverseResizeNodesCluster

`func NewUniverseResizeNodesCluster(uuid string, ) *UniverseResizeNodesCluster`

NewUniverseResizeNodesCluster instantiates a new UniverseResizeNodesCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseResizeNodesClusterWithDefaults

`func NewUniverseResizeNodesClusterWithDefaults() *UniverseResizeNodesCluster`

NewUniverseResizeNodesClusterWithDefaults instantiates a new UniverseResizeNodesCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *UniverseResizeNodesCluster) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UniverseResizeNodesCluster) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UniverseResizeNodesCluster) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetNodeSpec

`func (o *UniverseResizeNodesCluster) GetNodeSpec() ClusterResizeNodeSpec`

GetNodeSpec returns the NodeSpec field if non-nil, zero value otherwise.

### GetNodeSpecOk

`func (o *UniverseResizeNodesCluster) GetNodeSpecOk() (*ClusterResizeNodeSpec, bool)`

GetNodeSpecOk returns a tuple with the NodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSpec

`func (o *UniverseResizeNodesCluster) SetNodeSpec(v ClusterResizeNodeSpec)`

SetNodeSpec sets NodeSpec field to given value.

### HasNodeSpec

`func (o *UniverseResizeNodesCluster) HasNodeSpec() bool`

HasNodeSpec returns a boolean if a field has been set.

### GetProviderNodesSpecs

`func (o *UniverseResizeNodesCluster) GetProviderNodesSpecs() []PerProviderResizeNodesSpec`

GetProviderNodesSpecs returns the ProviderNodesSpecs field if non-nil, zero value otherwise.

### GetProviderNodesSpecsOk

`func (o *UniverseResizeNodesCluster) GetProviderNodesSpecsOk() (*[]PerProviderResizeNodesSpec, bool)`

GetProviderNodesSpecsOk returns a tuple with the ProviderNodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNodesSpecs

`func (o *UniverseResizeNodesCluster) SetProviderNodesSpecs(v []PerProviderResizeNodesSpec)`

SetProviderNodesSpecs sets ProviderNodesSpecs field to given value.

### HasProviderNodesSpecs

`func (o *UniverseResizeNodesCluster) HasProviderNodesSpecs() bool`

HasProviderNodesSpecs returns a boolean if a field has been set.

### GetGflags

`func (o *UniverseResizeNodesCluster) GetGflags() ClusterGFlags`

GetGflags returns the Gflags field if non-nil, zero value otherwise.

### GetGflagsOk

`func (o *UniverseResizeNodesCluster) GetGflagsOk() (*ClusterGFlags, bool)`

GetGflagsOk returns a tuple with the Gflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGflags

`func (o *UniverseResizeNodesCluster) SetGflags(v ClusterGFlags)`

SetGflags sets Gflags field to given value.

### HasGflags

`func (o *UniverseResizeNodesCluster) HasGflags() bool`

HasGflags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


