# AvailabilityZoneResizeNodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type for tserver/master nodes of cluster that determines the cpu and memory resources. | [optional] 
**StorageSpec** | Pointer to [**ClusterResizeStorageSpec**](ClusterResizeStorageSpec.md) |  | [optional] 
**CgroupSize** | Pointer to **int32** | Amount of memory in MB to limit the postgres process using the ysql cgroup. The value should be greater than 0. When set to 0 it results in no cgroup limits. For a read replica cluster, setting this value to null or -1 would inherit this value from the primary cluster. Applicable only for nodes running as Linux VMs on AWS/GCP/Azure Cloud Provider. Only used internally by YBM. | [optional] 
**Tserver** | Pointer to [**PerProcessResizeNodeSpec**](PerProcessResizeNodeSpec.md) |  | [optional] 
**Master** | Pointer to [**PerProcessResizeNodeSpec**](PerProcessResizeNodeSpec.md) |  | [optional] 

## Methods

### NewAvailabilityZoneResizeNodeSpec

`func NewAvailabilityZoneResizeNodeSpec() *AvailabilityZoneResizeNodeSpec`

NewAvailabilityZoneResizeNodeSpec instantiates a new AvailabilityZoneResizeNodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneResizeNodeSpecWithDefaults

`func NewAvailabilityZoneResizeNodeSpecWithDefaults() *AvailabilityZoneResizeNodeSpec`

NewAvailabilityZoneResizeNodeSpecWithDefaults instantiates a new AvailabilityZoneResizeNodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *AvailabilityZoneResizeNodeSpec) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AvailabilityZoneResizeNodeSpec) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AvailabilityZoneResizeNodeSpec) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AvailabilityZoneResizeNodeSpec) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetStorageSpec

`func (o *AvailabilityZoneResizeNodeSpec) GetStorageSpec() ClusterResizeStorageSpec`

GetStorageSpec returns the StorageSpec field if non-nil, zero value otherwise.

### GetStorageSpecOk

`func (o *AvailabilityZoneResizeNodeSpec) GetStorageSpecOk() (*ClusterResizeStorageSpec, bool)`

GetStorageSpecOk returns a tuple with the StorageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpec

`func (o *AvailabilityZoneResizeNodeSpec) SetStorageSpec(v ClusterResizeStorageSpec)`

SetStorageSpec sets StorageSpec field to given value.

### HasStorageSpec

`func (o *AvailabilityZoneResizeNodeSpec) HasStorageSpec() bool`

HasStorageSpec returns a boolean if a field has been set.

### GetCgroupSize

`func (o *AvailabilityZoneResizeNodeSpec) GetCgroupSize() int32`

GetCgroupSize returns the CgroupSize field if non-nil, zero value otherwise.

### GetCgroupSizeOk

`func (o *AvailabilityZoneResizeNodeSpec) GetCgroupSizeOk() (*int32, bool)`

GetCgroupSizeOk returns a tuple with the CgroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupSize

`func (o *AvailabilityZoneResizeNodeSpec) SetCgroupSize(v int32)`

SetCgroupSize sets CgroupSize field to given value.

### HasCgroupSize

`func (o *AvailabilityZoneResizeNodeSpec) HasCgroupSize() bool`

HasCgroupSize returns a boolean if a field has been set.

### GetTserver

`func (o *AvailabilityZoneResizeNodeSpec) GetTserver() PerProcessResizeNodeSpec`

GetTserver returns the Tserver field if non-nil, zero value otherwise.

### GetTserverOk

`func (o *AvailabilityZoneResizeNodeSpec) GetTserverOk() (*PerProcessResizeNodeSpec, bool)`

GetTserverOk returns a tuple with the Tserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserver

`func (o *AvailabilityZoneResizeNodeSpec) SetTserver(v PerProcessResizeNodeSpec)`

SetTserver sets Tserver field to given value.

### HasTserver

`func (o *AvailabilityZoneResizeNodeSpec) HasTserver() bool`

HasTserver returns a boolean if a field has been set.

### GetMaster

`func (o *AvailabilityZoneResizeNodeSpec) GetMaster() PerProcessResizeNodeSpec`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *AvailabilityZoneResizeNodeSpec) GetMasterOk() (*PerProcessResizeNodeSpec, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *AvailabilityZoneResizeNodeSpec) SetMaster(v PerProcessResizeNodeSpec)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *AvailabilityZoneResizeNodeSpec) HasMaster() bool`

HasMaster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


