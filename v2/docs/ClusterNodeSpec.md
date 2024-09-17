# ClusterNodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type for tserver/master nodes of cluster that determines the cpu and memory resources. | [optional] 
**StorageSpec** | Pointer to [**ClusterStorageSpec**](ClusterStorageSpec.md) |  | [optional] 
**CgroupSize** | Pointer to **int32** | Amount of memory in MB to limit the postgres process using the ysql cgroup. The value should be greater than 0. When set to 0 it results in no cgroup limits. For a read replica cluster, setting this value to null or -1 would inherit this value from the primary cluster. Applicable only for nodes running as Linux VMs on AWS/GCP/Azure Cloud Provider. Only used internally by YBM. | [optional] 
**Tserver** | Pointer to [**PerProcessNodeSpec**](PerProcessNodeSpec.md) |  | [optional] 
**Master** | Pointer to [**PerProcessNodeSpec**](PerProcessNodeSpec.md) |  | [optional] 
**DedicatedNodes** | Pointer to **bool** | Whether to run tserver and master processes in dedicated nodes in this cluster. Defaults to false where master and tserver processes share the same node. | [optional] [default to false]
**K8sMasterResourceSpec** | Pointer to [**K8SNodeResourceSpec**](K8SNodeResourceSpec.md) |  | [optional] 
**K8sTserverResourceSpec** | Pointer to [**K8SNodeResourceSpec**](K8SNodeResourceSpec.md) |  | [optional] 
**AzNodeSpec** | Pointer to [**map[string]AvailabilityZoneNodeSpec**](AvailabilityZoneNodeSpec.md) | Granular node settings overridden per Availability Zone identified by AZ uuid. | [optional] 

## Methods

### NewClusterNodeSpec

`func NewClusterNodeSpec() *ClusterNodeSpec`

NewClusterNodeSpec instantiates a new ClusterNodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeSpecWithDefaults

`func NewClusterNodeSpecWithDefaults() *ClusterNodeSpec`

NewClusterNodeSpecWithDefaults instantiates a new ClusterNodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *ClusterNodeSpec) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterNodeSpec) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterNodeSpec) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ClusterNodeSpec) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetStorageSpec

`func (o *ClusterNodeSpec) GetStorageSpec() ClusterStorageSpec`

GetStorageSpec returns the StorageSpec field if non-nil, zero value otherwise.

### GetStorageSpecOk

`func (o *ClusterNodeSpec) GetStorageSpecOk() (*ClusterStorageSpec, bool)`

GetStorageSpecOk returns a tuple with the StorageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpec

`func (o *ClusterNodeSpec) SetStorageSpec(v ClusterStorageSpec)`

SetStorageSpec sets StorageSpec field to given value.

### HasStorageSpec

`func (o *ClusterNodeSpec) HasStorageSpec() bool`

HasStorageSpec returns a boolean if a field has been set.

### GetCgroupSize

`func (o *ClusterNodeSpec) GetCgroupSize() int32`

GetCgroupSize returns the CgroupSize field if non-nil, zero value otherwise.

### GetCgroupSizeOk

`func (o *ClusterNodeSpec) GetCgroupSizeOk() (*int32, bool)`

GetCgroupSizeOk returns a tuple with the CgroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupSize

`func (o *ClusterNodeSpec) SetCgroupSize(v int32)`

SetCgroupSize sets CgroupSize field to given value.

### HasCgroupSize

`func (o *ClusterNodeSpec) HasCgroupSize() bool`

HasCgroupSize returns a boolean if a field has been set.

### GetTserver

`func (o *ClusterNodeSpec) GetTserver() PerProcessNodeSpec`

GetTserver returns the Tserver field if non-nil, zero value otherwise.

### GetTserverOk

`func (o *ClusterNodeSpec) GetTserverOk() (*PerProcessNodeSpec, bool)`

GetTserverOk returns a tuple with the Tserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserver

`func (o *ClusterNodeSpec) SetTserver(v PerProcessNodeSpec)`

SetTserver sets Tserver field to given value.

### HasTserver

`func (o *ClusterNodeSpec) HasTserver() bool`

HasTserver returns a boolean if a field has been set.

### GetMaster

`func (o *ClusterNodeSpec) GetMaster() PerProcessNodeSpec`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ClusterNodeSpec) GetMasterOk() (*PerProcessNodeSpec, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ClusterNodeSpec) SetMaster(v PerProcessNodeSpec)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ClusterNodeSpec) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetDedicatedNodes

`func (o *ClusterNodeSpec) GetDedicatedNodes() bool`

GetDedicatedNodes returns the DedicatedNodes field if non-nil, zero value otherwise.

### GetDedicatedNodesOk

`func (o *ClusterNodeSpec) GetDedicatedNodesOk() (*bool, bool)`

GetDedicatedNodesOk returns a tuple with the DedicatedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedNodes

`func (o *ClusterNodeSpec) SetDedicatedNodes(v bool)`

SetDedicatedNodes sets DedicatedNodes field to given value.

### HasDedicatedNodes

`func (o *ClusterNodeSpec) HasDedicatedNodes() bool`

HasDedicatedNodes returns a boolean if a field has been set.

### GetK8sMasterResourceSpec

`func (o *ClusterNodeSpec) GetK8sMasterResourceSpec() K8SNodeResourceSpec`

GetK8sMasterResourceSpec returns the K8sMasterResourceSpec field if non-nil, zero value otherwise.

### GetK8sMasterResourceSpecOk

`func (o *ClusterNodeSpec) GetK8sMasterResourceSpecOk() (*K8SNodeResourceSpec, bool)`

GetK8sMasterResourceSpecOk returns a tuple with the K8sMasterResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sMasterResourceSpec

`func (o *ClusterNodeSpec) SetK8sMasterResourceSpec(v K8SNodeResourceSpec)`

SetK8sMasterResourceSpec sets K8sMasterResourceSpec field to given value.

### HasK8sMasterResourceSpec

`func (o *ClusterNodeSpec) HasK8sMasterResourceSpec() bool`

HasK8sMasterResourceSpec returns a boolean if a field has been set.

### GetK8sTserverResourceSpec

`func (o *ClusterNodeSpec) GetK8sTserverResourceSpec() K8SNodeResourceSpec`

GetK8sTserverResourceSpec returns the K8sTserverResourceSpec field if non-nil, zero value otherwise.

### GetK8sTserverResourceSpecOk

`func (o *ClusterNodeSpec) GetK8sTserverResourceSpecOk() (*K8SNodeResourceSpec, bool)`

GetK8sTserverResourceSpecOk returns a tuple with the K8sTserverResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sTserverResourceSpec

`func (o *ClusterNodeSpec) SetK8sTserverResourceSpec(v K8SNodeResourceSpec)`

SetK8sTserverResourceSpec sets K8sTserverResourceSpec field to given value.

### HasK8sTserverResourceSpec

`func (o *ClusterNodeSpec) HasK8sTserverResourceSpec() bool`

HasK8sTserverResourceSpec returns a boolean if a field has been set.

### GetAzNodeSpec

`func (o *ClusterNodeSpec) GetAzNodeSpec() map[string]AvailabilityZoneNodeSpec`

GetAzNodeSpec returns the AzNodeSpec field if non-nil, zero value otherwise.

### GetAzNodeSpecOk

`func (o *ClusterNodeSpec) GetAzNodeSpecOk() (*map[string]AvailabilityZoneNodeSpec, bool)`

GetAzNodeSpecOk returns a tuple with the AzNodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNodeSpec

`func (o *ClusterNodeSpec) SetAzNodeSpec(v map[string]AvailabilityZoneNodeSpec)`

SetAzNodeSpec sets AzNodeSpec field to given value.

### HasAzNodeSpec

`func (o *ClusterNodeSpec) HasAzNodeSpec() bool`

HasAzNodeSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


