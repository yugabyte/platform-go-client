# PerProcessNodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type for tserver/master nodes of cluster that determines the cpu and memory resources. | [optional] 
**StorageSpec** | Pointer to [**ClusterStorageSpec**](ClusterStorageSpec.md) |  | [optional] 

## Methods

### NewPerProcessNodeSpec

`func NewPerProcessNodeSpec() *PerProcessNodeSpec`

NewPerProcessNodeSpec instantiates a new PerProcessNodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerProcessNodeSpecWithDefaults

`func NewPerProcessNodeSpecWithDefaults() *PerProcessNodeSpec`

NewPerProcessNodeSpecWithDefaults instantiates a new PerProcessNodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *PerProcessNodeSpec) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *PerProcessNodeSpec) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *PerProcessNodeSpec) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *PerProcessNodeSpec) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetStorageSpec

`func (o *PerProcessNodeSpec) GetStorageSpec() ClusterStorageSpec`

GetStorageSpec returns the StorageSpec field if non-nil, zero value otherwise.

### GetStorageSpecOk

`func (o *PerProcessNodeSpec) GetStorageSpecOk() (*ClusterStorageSpec, bool)`

GetStorageSpecOk returns a tuple with the StorageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpec

`func (o *PerProcessNodeSpec) SetStorageSpec(v ClusterStorageSpec)`

SetStorageSpec sets StorageSpec field to given value.

### HasStorageSpec

`func (o *PerProcessNodeSpec) HasStorageSpec() bool`

HasStorageSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


