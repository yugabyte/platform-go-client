# AZOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CgroupSize** | Pointer to **int32** |  | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**PerProcess** | Pointer to [**map[string]PerProcessDetails**](PerProcessDetails.md) |  | [optional] 
**ProxyConfig** | Pointer to [**ProxyConfig**](ProxyConfig.md) |  | [optional] 

## Methods

### NewAZOverrides

`func NewAZOverrides() *AZOverrides`

NewAZOverrides instantiates a new AZOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAZOverridesWithDefaults

`func NewAZOverridesWithDefaults() *AZOverrides`

NewAZOverridesWithDefaults instantiates a new AZOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCgroupSize

`func (o *AZOverrides) GetCgroupSize() int32`

GetCgroupSize returns the CgroupSize field if non-nil, zero value otherwise.

### GetCgroupSizeOk

`func (o *AZOverrides) GetCgroupSizeOk() (*int32, bool)`

GetCgroupSizeOk returns a tuple with the CgroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupSize

`func (o *AZOverrides) SetCgroupSize(v int32)`

SetCgroupSize sets CgroupSize field to given value.

### HasCgroupSize

`func (o *AZOverrides) HasCgroupSize() bool`

HasCgroupSize returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *AZOverrides) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *AZOverrides) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *AZOverrides) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *AZOverrides) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetInstanceType

`func (o *AZOverrides) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AZOverrides) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AZOverrides) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AZOverrides) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetPerProcess

`func (o *AZOverrides) GetPerProcess() map[string]PerProcessDetails`

GetPerProcess returns the PerProcess field if non-nil, zero value otherwise.

### GetPerProcessOk

`func (o *AZOverrides) GetPerProcessOk() (*map[string]PerProcessDetails, bool)`

GetPerProcessOk returns a tuple with the PerProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerProcess

`func (o *AZOverrides) SetPerProcess(v map[string]PerProcessDetails)`

SetPerProcess sets PerProcess field to given value.

### HasPerProcess

`func (o *AZOverrides) HasPerProcess() bool`

HasPerProcess returns a boolean if a field has been set.

### GetProxyConfig

`func (o *AZOverrides) GetProxyConfig() ProxyConfig`

GetProxyConfig returns the ProxyConfig field if non-nil, zero value otherwise.

### GetProxyConfigOk

`func (o *AZOverrides) GetProxyConfigOk() (*ProxyConfig, bool)`

GetProxyConfigOk returns a tuple with the ProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfig

`func (o *AZOverrides) SetProxyConfig(v ProxyConfig)`

SetProxyConfig sets ProxyConfig field to given value.

### HasProxyConfig

`func (o *AZOverrides) HasProxyConfig() bool`

HasProxyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


