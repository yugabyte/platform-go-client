# ClusterEditSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The system generated cluster uuid to edit. This can be fetched from ClusterInfo. | 
**NumNodes** | Pointer to **int32** | Set the number of nodes (tservers) to provision in this cluster | [optional] 
**NodeSpec** | Pointer to [**ClusterNodeSpec**](ClusterNodeSpec.md) |  | [optional] 
**ProviderSpec** | Pointer to [**ClusterProviderEditSpec**](ClusterProviderEditSpec.md) |  | [optional] 
**PlacementSpec** | Pointer to [**ClusterPlacementSpec**](ClusterPlacementSpec.md) |  | [optional] 
**InstanceTags** | Pointer to **map[string]string** | A map of strings representing a set of Tags and Values to apply on nodes in the aws/gcp/azu cloud. See https://docs.yugabyte.com/preview/yugabyte-platform/manage-deployments/instance-tags/. | [optional] 

## Methods

### NewClusterEditSpec

`func NewClusterEditSpec(uuid string, ) *ClusterEditSpec`

NewClusterEditSpec instantiates a new ClusterEditSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEditSpecWithDefaults

`func NewClusterEditSpecWithDefaults() *ClusterEditSpec`

NewClusterEditSpecWithDefaults instantiates a new ClusterEditSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ClusterEditSpec) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ClusterEditSpec) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ClusterEditSpec) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetNumNodes

`func (o *ClusterEditSpec) GetNumNodes() int32`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *ClusterEditSpec) GetNumNodesOk() (*int32, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *ClusterEditSpec) SetNumNodes(v int32)`

SetNumNodes sets NumNodes field to given value.

### HasNumNodes

`func (o *ClusterEditSpec) HasNumNodes() bool`

HasNumNodes returns a boolean if a field has been set.

### GetNodeSpec

`func (o *ClusterEditSpec) GetNodeSpec() ClusterNodeSpec`

GetNodeSpec returns the NodeSpec field if non-nil, zero value otherwise.

### GetNodeSpecOk

`func (o *ClusterEditSpec) GetNodeSpecOk() (*ClusterNodeSpec, bool)`

GetNodeSpecOk returns a tuple with the NodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSpec

`func (o *ClusterEditSpec) SetNodeSpec(v ClusterNodeSpec)`

SetNodeSpec sets NodeSpec field to given value.

### HasNodeSpec

`func (o *ClusterEditSpec) HasNodeSpec() bool`

HasNodeSpec returns a boolean if a field has been set.

### GetProviderSpec

`func (o *ClusterEditSpec) GetProviderSpec() ClusterProviderEditSpec`

GetProviderSpec returns the ProviderSpec field if non-nil, zero value otherwise.

### GetProviderSpecOk

`func (o *ClusterEditSpec) GetProviderSpecOk() (*ClusterProviderEditSpec, bool)`

GetProviderSpecOk returns a tuple with the ProviderSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSpec

`func (o *ClusterEditSpec) SetProviderSpec(v ClusterProviderEditSpec)`

SetProviderSpec sets ProviderSpec field to given value.

### HasProviderSpec

`func (o *ClusterEditSpec) HasProviderSpec() bool`

HasProviderSpec returns a boolean if a field has been set.

### GetPlacementSpec

`func (o *ClusterEditSpec) GetPlacementSpec() ClusterPlacementSpec`

GetPlacementSpec returns the PlacementSpec field if non-nil, zero value otherwise.

### GetPlacementSpecOk

`func (o *ClusterEditSpec) GetPlacementSpecOk() (*ClusterPlacementSpec, bool)`

GetPlacementSpecOk returns a tuple with the PlacementSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementSpec

`func (o *ClusterEditSpec) SetPlacementSpec(v ClusterPlacementSpec)`

SetPlacementSpec sets PlacementSpec field to given value.

### HasPlacementSpec

`func (o *ClusterEditSpec) HasPlacementSpec() bool`

HasPlacementSpec returns a boolean if a field has been set.

### GetInstanceTags

`func (o *ClusterEditSpec) GetInstanceTags() map[string]string`

GetInstanceTags returns the InstanceTags field if non-nil, zero value otherwise.

### GetInstanceTagsOk

`func (o *ClusterEditSpec) GetInstanceTagsOk() (*map[string]string, bool)`

GetInstanceTagsOk returns a tuple with the InstanceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTags

`func (o *ClusterEditSpec) SetInstanceTags(v map[string]string)`

SetInstanceTags sets InstanceTags field to given value.

### HasInstanceTags

`func (o *ClusterEditSpec) HasInstanceTags() bool`

HasInstanceTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


