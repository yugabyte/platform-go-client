# PerProcessResizeNodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type for tserver/master nodes of cluster that determines the cpu and memory resources. | [optional] 
**StorageSpec** | Pointer to [**ClusterResizeStorageSpec**](ClusterResizeStorageSpec.md) |  | [optional] 

## Methods

### NewPerProcessResizeNodeSpec

`func NewPerProcessResizeNodeSpec() *PerProcessResizeNodeSpec`

NewPerProcessResizeNodeSpec instantiates a new PerProcessResizeNodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerProcessResizeNodeSpecWithDefaults

`func NewPerProcessResizeNodeSpecWithDefaults() *PerProcessResizeNodeSpec`

NewPerProcessResizeNodeSpecWithDefaults instantiates a new PerProcessResizeNodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *PerProcessResizeNodeSpec) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *PerProcessResizeNodeSpec) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *PerProcessResizeNodeSpec) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *PerProcessResizeNodeSpec) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetStorageSpec

`func (o *PerProcessResizeNodeSpec) GetStorageSpec() ClusterResizeStorageSpec`

GetStorageSpec returns the StorageSpec field if non-nil, zero value otherwise.

### GetStorageSpecOk

`func (o *PerProcessResizeNodeSpec) GetStorageSpecOk() (*ClusterResizeStorageSpec, bool)`

GetStorageSpecOk returns a tuple with the StorageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpec

`func (o *PerProcessResizeNodeSpec) SetStorageSpec(v ClusterResizeStorageSpec)`

SetStorageSpec sets StorageSpec field to given value.

### HasStorageSpec

`func (o *PerProcessResizeNodeSpec) HasStorageSpec() bool`

HasStorageSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


