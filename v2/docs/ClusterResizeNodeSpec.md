# ClusterResizeNodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type for tserver/master nodes of cluster that determines the cpu and memory resources. | [optional] 
**StorageSpec** | Pointer to [**ClusterResizeStorageSpec**](ClusterResizeStorageSpec.md) |  | [optional] 
**CgroupSize** | Pointer to **int32** | Amount of memory in MB to limit the postgres process using the ysql cgroup. The value should be greater than 0. When set to 0 it results in no cgroup limits. For a read replica cluster, setting this value to null or -1 would inherit this value from the primary cluster. Applicable only for nodes running as Linux VMs on AWS/GCP/Azure Cloud Provider. Only used internally by YBM. | [optional] 
**Tserver** | Pointer to [**PerProcessResizeNodeSpec**](PerProcessResizeNodeSpec.md) |  | [optional] 
**Master** | Pointer to [**PerProcessResizeNodeSpec**](PerProcessResizeNodeSpec.md) |  | [optional] 
**K8sMasterResourceSpec** | Pointer to [**K8SNodeResourceSpec**](K8SNodeResourceSpec.md) |  | [optional] 
**K8sTserverResourceSpec** | Pointer to [**K8SNodeResourceSpec**](K8SNodeResourceSpec.md) |  | [optional] 
**AzNodeSpec** | Pointer to [**map[string]AvailabilityZoneResizeNodeSpec**](AvailabilityZoneResizeNodeSpec.md) | Granular node settings overridden per Availability Zone identified by AZ uuid. | [optional] 

## Methods

### NewClusterResizeNodeSpec

`func NewClusterResizeNodeSpec() *ClusterResizeNodeSpec`

NewClusterResizeNodeSpec instantiates a new ClusterResizeNodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResizeNodeSpecWithDefaults

`func NewClusterResizeNodeSpecWithDefaults() *ClusterResizeNodeSpec`

NewClusterResizeNodeSpecWithDefaults instantiates a new ClusterResizeNodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *ClusterResizeNodeSpec) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterResizeNodeSpec) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterResizeNodeSpec) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ClusterResizeNodeSpec) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetStorageSpec

`func (o *ClusterResizeNodeSpec) GetStorageSpec() ClusterResizeStorageSpec`

GetStorageSpec returns the StorageSpec field if non-nil, zero value otherwise.

### GetStorageSpecOk

`func (o *ClusterResizeNodeSpec) GetStorageSpecOk() (*ClusterResizeStorageSpec, bool)`

GetStorageSpecOk returns a tuple with the StorageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpec

`func (o *ClusterResizeNodeSpec) SetStorageSpec(v ClusterResizeStorageSpec)`

SetStorageSpec sets StorageSpec field to given value.

### HasStorageSpec

`func (o *ClusterResizeNodeSpec) HasStorageSpec() bool`

HasStorageSpec returns a boolean if a field has been set.

### GetCgroupSize

`func (o *ClusterResizeNodeSpec) GetCgroupSize() int32`

GetCgroupSize returns the CgroupSize field if non-nil, zero value otherwise.

### GetCgroupSizeOk

`func (o *ClusterResizeNodeSpec) GetCgroupSizeOk() (*int32, bool)`

GetCgroupSizeOk returns a tuple with the CgroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupSize

`func (o *ClusterResizeNodeSpec) SetCgroupSize(v int32)`

SetCgroupSize sets CgroupSize field to given value.

### HasCgroupSize

`func (o *ClusterResizeNodeSpec) HasCgroupSize() bool`

HasCgroupSize returns a boolean if a field has been set.

### GetTserver

`func (o *ClusterResizeNodeSpec) GetTserver() PerProcessResizeNodeSpec`

GetTserver returns the Tserver field if non-nil, zero value otherwise.

### GetTserverOk

`func (o *ClusterResizeNodeSpec) GetTserverOk() (*PerProcessResizeNodeSpec, bool)`

GetTserverOk returns a tuple with the Tserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserver

`func (o *ClusterResizeNodeSpec) SetTserver(v PerProcessResizeNodeSpec)`

SetTserver sets Tserver field to given value.

### HasTserver

`func (o *ClusterResizeNodeSpec) HasTserver() bool`

HasTserver returns a boolean if a field has been set.

### GetMaster

`func (o *ClusterResizeNodeSpec) GetMaster() PerProcessResizeNodeSpec`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ClusterResizeNodeSpec) GetMasterOk() (*PerProcessResizeNodeSpec, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ClusterResizeNodeSpec) SetMaster(v PerProcessResizeNodeSpec)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ClusterResizeNodeSpec) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetK8sMasterResourceSpec

`func (o *ClusterResizeNodeSpec) GetK8sMasterResourceSpec() K8SNodeResourceSpec`

GetK8sMasterResourceSpec returns the K8sMasterResourceSpec field if non-nil, zero value otherwise.

### GetK8sMasterResourceSpecOk

`func (o *ClusterResizeNodeSpec) GetK8sMasterResourceSpecOk() (*K8SNodeResourceSpec, bool)`

GetK8sMasterResourceSpecOk returns a tuple with the K8sMasterResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sMasterResourceSpec

`func (o *ClusterResizeNodeSpec) SetK8sMasterResourceSpec(v K8SNodeResourceSpec)`

SetK8sMasterResourceSpec sets K8sMasterResourceSpec field to given value.

### HasK8sMasterResourceSpec

`func (o *ClusterResizeNodeSpec) HasK8sMasterResourceSpec() bool`

HasK8sMasterResourceSpec returns a boolean if a field has been set.

### GetK8sTserverResourceSpec

`func (o *ClusterResizeNodeSpec) GetK8sTserverResourceSpec() K8SNodeResourceSpec`

GetK8sTserverResourceSpec returns the K8sTserverResourceSpec field if non-nil, zero value otherwise.

### GetK8sTserverResourceSpecOk

`func (o *ClusterResizeNodeSpec) GetK8sTserverResourceSpecOk() (*K8SNodeResourceSpec, bool)`

GetK8sTserverResourceSpecOk returns a tuple with the K8sTserverResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sTserverResourceSpec

`func (o *ClusterResizeNodeSpec) SetK8sTserverResourceSpec(v K8SNodeResourceSpec)`

SetK8sTserverResourceSpec sets K8sTserverResourceSpec field to given value.

### HasK8sTserverResourceSpec

`func (o *ClusterResizeNodeSpec) HasK8sTserverResourceSpec() bool`

HasK8sTserverResourceSpec returns a boolean if a field has been set.

### GetAzNodeSpec

`func (o *ClusterResizeNodeSpec) GetAzNodeSpec() map[string]AvailabilityZoneResizeNodeSpec`

GetAzNodeSpec returns the AzNodeSpec field if non-nil, zero value otherwise.

### GetAzNodeSpecOk

`func (o *ClusterResizeNodeSpec) GetAzNodeSpecOk() (*map[string]AvailabilityZoneResizeNodeSpec, bool)`

GetAzNodeSpecOk returns a tuple with the AzNodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNodeSpec

`func (o *ClusterResizeNodeSpec) SetAzNodeSpec(v map[string]AvailabilityZoneResizeNodeSpec)`

SetAzNodeSpec sets AzNodeSpec field to given value.

### HasAzNodeSpec

`func (o *ClusterResizeNodeSpec) HasAzNodeSpec() bool`

HasAzNodeSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


