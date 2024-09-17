# ClusterNodeSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedNodes** | Pointer to **bool** | Whether to run tserver and master processes in dedicated nodes in this cluster. Defaults to false where master and tserver processes share the same node. | [optional] [default to false]
**K8sMasterResourceSpec** | Pointer to [**K8SNodeResourceSpec**](K8SNodeResourceSpec.md) |  | [optional] 
**K8sTserverResourceSpec** | Pointer to [**K8SNodeResourceSpec**](K8SNodeResourceSpec.md) |  | [optional] 
**AzNodeSpec** | Pointer to [**map[string]AvailabilityZoneNodeSpec**](AvailabilityZoneNodeSpec.md) | Granular node settings overridden per Availability Zone identified by AZ uuid. | [optional] 

## Methods

### NewClusterNodeSpecAllOf

`func NewClusterNodeSpecAllOf() *ClusterNodeSpecAllOf`

NewClusterNodeSpecAllOf instantiates a new ClusterNodeSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeSpecAllOfWithDefaults

`func NewClusterNodeSpecAllOfWithDefaults() *ClusterNodeSpecAllOf`

NewClusterNodeSpecAllOfWithDefaults instantiates a new ClusterNodeSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedNodes

`func (o *ClusterNodeSpecAllOf) GetDedicatedNodes() bool`

GetDedicatedNodes returns the DedicatedNodes field if non-nil, zero value otherwise.

### GetDedicatedNodesOk

`func (o *ClusterNodeSpecAllOf) GetDedicatedNodesOk() (*bool, bool)`

GetDedicatedNodesOk returns a tuple with the DedicatedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedNodes

`func (o *ClusterNodeSpecAllOf) SetDedicatedNodes(v bool)`

SetDedicatedNodes sets DedicatedNodes field to given value.

### HasDedicatedNodes

`func (o *ClusterNodeSpecAllOf) HasDedicatedNodes() bool`

HasDedicatedNodes returns a boolean if a field has been set.

### GetK8sMasterResourceSpec

`func (o *ClusterNodeSpecAllOf) GetK8sMasterResourceSpec() K8SNodeResourceSpec`

GetK8sMasterResourceSpec returns the K8sMasterResourceSpec field if non-nil, zero value otherwise.

### GetK8sMasterResourceSpecOk

`func (o *ClusterNodeSpecAllOf) GetK8sMasterResourceSpecOk() (*K8SNodeResourceSpec, bool)`

GetK8sMasterResourceSpecOk returns a tuple with the K8sMasterResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sMasterResourceSpec

`func (o *ClusterNodeSpecAllOf) SetK8sMasterResourceSpec(v K8SNodeResourceSpec)`

SetK8sMasterResourceSpec sets K8sMasterResourceSpec field to given value.

### HasK8sMasterResourceSpec

`func (o *ClusterNodeSpecAllOf) HasK8sMasterResourceSpec() bool`

HasK8sMasterResourceSpec returns a boolean if a field has been set.

### GetK8sTserverResourceSpec

`func (o *ClusterNodeSpecAllOf) GetK8sTserverResourceSpec() K8SNodeResourceSpec`

GetK8sTserverResourceSpec returns the K8sTserverResourceSpec field if non-nil, zero value otherwise.

### GetK8sTserverResourceSpecOk

`func (o *ClusterNodeSpecAllOf) GetK8sTserverResourceSpecOk() (*K8SNodeResourceSpec, bool)`

GetK8sTserverResourceSpecOk returns a tuple with the K8sTserverResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sTserverResourceSpec

`func (o *ClusterNodeSpecAllOf) SetK8sTserverResourceSpec(v K8SNodeResourceSpec)`

SetK8sTserverResourceSpec sets K8sTserverResourceSpec field to given value.

### HasK8sTserverResourceSpec

`func (o *ClusterNodeSpecAllOf) HasK8sTserverResourceSpec() bool`

HasK8sTserverResourceSpec returns a boolean if a field has been set.

### GetAzNodeSpec

`func (o *ClusterNodeSpecAllOf) GetAzNodeSpec() map[string]AvailabilityZoneNodeSpec`

GetAzNodeSpec returns the AzNodeSpec field if non-nil, zero value otherwise.

### GetAzNodeSpecOk

`func (o *ClusterNodeSpecAllOf) GetAzNodeSpecOk() (*map[string]AvailabilityZoneNodeSpec, bool)`

GetAzNodeSpecOk returns a tuple with the AzNodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNodeSpec

`func (o *ClusterNodeSpecAllOf) SetAzNodeSpec(v map[string]AvailabilityZoneNodeSpec)`

SetAzNodeSpec sets AzNodeSpec field to given value.

### HasAzNodeSpec

`func (o *ClusterNodeSpecAllOf) HasAzNodeSpec() bool`

HasAzNodeSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


