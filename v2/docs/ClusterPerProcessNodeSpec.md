# ClusterPerProcessNodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type for tserver/master nodes of cluster that determines the cpu and memory resources. | [optional] 
**StorageSpec** | Pointer to [**ClusterStorageSpec**](ClusterStorageSpec.md) |  | [optional] 

## Methods

### NewClusterPerProcessNodeSpec

`func NewClusterPerProcessNodeSpec() *ClusterPerProcessNodeSpec`

NewClusterPerProcessNodeSpec instantiates a new ClusterPerProcessNodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPerProcessNodeSpecWithDefaults

`func NewClusterPerProcessNodeSpecWithDefaults() *ClusterPerProcessNodeSpec`

NewClusterPerProcessNodeSpecWithDefaults instantiates a new ClusterPerProcessNodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *ClusterPerProcessNodeSpec) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterPerProcessNodeSpec) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterPerProcessNodeSpec) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ClusterPerProcessNodeSpec) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetStorageSpec

`func (o *ClusterPerProcessNodeSpec) GetStorageSpec() ClusterStorageSpec`

GetStorageSpec returns the StorageSpec field if non-nil, zero value otherwise.

### GetStorageSpecOk

`func (o *ClusterPerProcessNodeSpec) GetStorageSpecOk() (*ClusterStorageSpec, bool)`

GetStorageSpecOk returns a tuple with the StorageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpec

`func (o *ClusterPerProcessNodeSpec) SetStorageSpec(v ClusterStorageSpec)`

SetStorageSpec sets StorageSpec field to given value.

### HasStorageSpec

`func (o *ClusterPerProcessNodeSpec) HasStorageSpec() bool`

HasStorageSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


