# NodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupProxyConfig** | Pointer to [**ProxyConfig**](ProxyConfig.md) |  | [optional] 
**CgroupSize** | Pointer to **int32** |  | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**K8SNodeResourceSpec** | Pointer to [**K8SNodeResourceSpec**](K8SNodeResourceSpec.md) |  | [optional] 

## Methods

### NewNodeSpec

`func NewNodeSpec() *NodeSpec`

NewNodeSpec instantiates a new NodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSpecWithDefaults

`func NewNodeSpecWithDefaults() *NodeSpec`

NewNodeSpecWithDefaults instantiates a new NodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupProxyConfig

`func (o *NodeSpec) GetBackupProxyConfig() ProxyConfig`

GetBackupProxyConfig returns the BackupProxyConfig field if non-nil, zero value otherwise.

### GetBackupProxyConfigOk

`func (o *NodeSpec) GetBackupProxyConfigOk() (*ProxyConfig, bool)`

GetBackupProxyConfigOk returns a tuple with the BackupProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProxyConfig

`func (o *NodeSpec) SetBackupProxyConfig(v ProxyConfig)`

SetBackupProxyConfig sets BackupProxyConfig field to given value.

### HasBackupProxyConfig

`func (o *NodeSpec) HasBackupProxyConfig() bool`

HasBackupProxyConfig returns a boolean if a field has been set.

### GetCgroupSize

`func (o *NodeSpec) GetCgroupSize() int32`

GetCgroupSize returns the CgroupSize field if non-nil, zero value otherwise.

### GetCgroupSizeOk

`func (o *NodeSpec) GetCgroupSizeOk() (*int32, bool)`

GetCgroupSizeOk returns a tuple with the CgroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupSize

`func (o *NodeSpec) SetCgroupSize(v int32)`

SetCgroupSize sets CgroupSize field to given value.

### HasCgroupSize

`func (o *NodeSpec) HasCgroupSize() bool`

HasCgroupSize returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *NodeSpec) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *NodeSpec) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *NodeSpec) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *NodeSpec) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetInstanceType

`func (o *NodeSpec) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *NodeSpec) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *NodeSpec) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *NodeSpec) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetK8SNodeResourceSpec

`func (o *NodeSpec) GetK8SNodeResourceSpec() K8SNodeResourceSpec`

GetK8SNodeResourceSpec returns the K8SNodeResourceSpec field if non-nil, zero value otherwise.

### GetK8SNodeResourceSpecOk

`func (o *NodeSpec) GetK8SNodeResourceSpecOk() (*K8SNodeResourceSpec, bool)`

GetK8SNodeResourceSpecOk returns a tuple with the K8SNodeResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8SNodeResourceSpec

`func (o *NodeSpec) SetK8SNodeResourceSpec(v K8SNodeResourceSpec)`

SetK8SNodeResourceSpec sets K8SNodeResourceSpec field to given value.

### HasK8SNodeResourceSpec

`func (o *NodeSpec) HasK8SNodeResourceSpec() bool`

HasK8SNodeResourceSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


