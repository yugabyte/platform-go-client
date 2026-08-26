# ProviderNodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type for tserver/master nodes that determines the cpu and memory resources. | [optional] 
**StorageSpec** | Pointer to [**ClusterStorageSpec**](ClusterStorageSpec.md) |  | [optional] 
**CgroupSize** | Pointer to **int32** | Amount of memory in MB to limit the postgres process using the ysql cgroup. The value should be greater than 0. When set to 0 it results in no cgroup limits. Applicable only for nodes running as Linux VMs on AWS/GCP/Azure Cloud Provider. | [optional] 
**BackupProxyConfig** | Pointer to [**NodeProxyConfig**](NodeProxyConfig.md) |  | [optional] 
**K8sNodeResourceSpec** | Pointer to [**K8SNodeResourceSpec**](K8SNodeResourceSpec.md) |  | [optional] 

## Methods

### NewProviderNodeSpec

`func NewProviderNodeSpec() *ProviderNodeSpec`

NewProviderNodeSpec instantiates a new ProviderNodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderNodeSpecWithDefaults

`func NewProviderNodeSpecWithDefaults() *ProviderNodeSpec`

NewProviderNodeSpecWithDefaults instantiates a new ProviderNodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *ProviderNodeSpec) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ProviderNodeSpec) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ProviderNodeSpec) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ProviderNodeSpec) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetStorageSpec

`func (o *ProviderNodeSpec) GetStorageSpec() ClusterStorageSpec`

GetStorageSpec returns the StorageSpec field if non-nil, zero value otherwise.

### GetStorageSpecOk

`func (o *ProviderNodeSpec) GetStorageSpecOk() (*ClusterStorageSpec, bool)`

GetStorageSpecOk returns a tuple with the StorageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpec

`func (o *ProviderNodeSpec) SetStorageSpec(v ClusterStorageSpec)`

SetStorageSpec sets StorageSpec field to given value.

### HasStorageSpec

`func (o *ProviderNodeSpec) HasStorageSpec() bool`

HasStorageSpec returns a boolean if a field has been set.

### GetCgroupSize

`func (o *ProviderNodeSpec) GetCgroupSize() int32`

GetCgroupSize returns the CgroupSize field if non-nil, zero value otherwise.

### GetCgroupSizeOk

`func (o *ProviderNodeSpec) GetCgroupSizeOk() (*int32, bool)`

GetCgroupSizeOk returns a tuple with the CgroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupSize

`func (o *ProviderNodeSpec) SetCgroupSize(v int32)`

SetCgroupSize sets CgroupSize field to given value.

### HasCgroupSize

`func (o *ProviderNodeSpec) HasCgroupSize() bool`

HasCgroupSize returns a boolean if a field has been set.

### GetBackupProxyConfig

`func (o *ProviderNodeSpec) GetBackupProxyConfig() NodeProxyConfig`

GetBackupProxyConfig returns the BackupProxyConfig field if non-nil, zero value otherwise.

### GetBackupProxyConfigOk

`func (o *ProviderNodeSpec) GetBackupProxyConfigOk() (*NodeProxyConfig, bool)`

GetBackupProxyConfigOk returns a tuple with the BackupProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProxyConfig

`func (o *ProviderNodeSpec) SetBackupProxyConfig(v NodeProxyConfig)`

SetBackupProxyConfig sets BackupProxyConfig field to given value.

### HasBackupProxyConfig

`func (o *ProviderNodeSpec) HasBackupProxyConfig() bool`

HasBackupProxyConfig returns a boolean if a field has been set.

### GetK8sNodeResourceSpec

`func (o *ProviderNodeSpec) GetK8sNodeResourceSpec() K8SNodeResourceSpec`

GetK8sNodeResourceSpec returns the K8sNodeResourceSpec field if non-nil, zero value otherwise.

### GetK8sNodeResourceSpecOk

`func (o *ProviderNodeSpec) GetK8sNodeResourceSpecOk() (*K8SNodeResourceSpec, bool)`

GetK8sNodeResourceSpecOk returns a tuple with the K8sNodeResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNodeResourceSpec

`func (o *ProviderNodeSpec) SetK8sNodeResourceSpec(v K8SNodeResourceSpec)`

SetK8sNodeResourceSpec sets K8sNodeResourceSpec field to given value.

### HasK8sNodeResourceSpec

`func (o *ProviderNodeSpec) HasK8sNodeResourceSpec() bool`

HasK8sNodeResourceSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


