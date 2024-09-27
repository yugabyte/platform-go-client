# K8SNodeResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCoreCount** | Pointer to **float64** | Number of CPU cores for tserver/master pods | [optional] [default to 2]
**MemoryGib** | Pointer to **float64** | Memory in GiB for tserver/master pods | [optional] [default to 4]

## Methods

### NewK8SNodeResourceSpec

`func NewK8SNodeResourceSpec() *K8SNodeResourceSpec`

NewK8SNodeResourceSpec instantiates a new K8SNodeResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8SNodeResourceSpecWithDefaults

`func NewK8SNodeResourceSpecWithDefaults() *K8SNodeResourceSpec`

NewK8SNodeResourceSpecWithDefaults instantiates a new K8SNodeResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCoreCount

`func (o *K8SNodeResourceSpec) GetCpuCoreCount() float64`

GetCpuCoreCount returns the CpuCoreCount field if non-nil, zero value otherwise.

### GetCpuCoreCountOk

`func (o *K8SNodeResourceSpec) GetCpuCoreCountOk() (*float64, bool)`

GetCpuCoreCountOk returns a tuple with the CpuCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoreCount

`func (o *K8SNodeResourceSpec) SetCpuCoreCount(v float64)`

SetCpuCoreCount sets CpuCoreCount field to given value.

### HasCpuCoreCount

`func (o *K8SNodeResourceSpec) HasCpuCoreCount() bool`

HasCpuCoreCount returns a boolean if a field has been set.

### GetMemoryGib

`func (o *K8SNodeResourceSpec) GetMemoryGib() float64`

GetMemoryGib returns the MemoryGib field if non-nil, zero value otherwise.

### GetMemoryGibOk

`func (o *K8SNodeResourceSpec) GetMemoryGibOk() (*float64, bool)`

GetMemoryGibOk returns a tuple with the MemoryGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGib

`func (o *K8SNodeResourceSpec) SetMemoryGib(v float64)`

SetMemoryGib sets MemoryGib field to given value.

### HasMemoryGib

`func (o *K8SNodeResourceSpec) HasMemoryGib() bool`

HasMemoryGib returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


