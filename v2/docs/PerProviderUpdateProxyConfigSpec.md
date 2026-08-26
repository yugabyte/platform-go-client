# PerProviderUpdateProxyConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Cloud provider UUID. | 
**ProxyConfig** | Pointer to [**NodeProxyConfig**](NodeProxyConfig.md) |  | [optional] 
**RegionNetworking** | Pointer to [**map[string]NodeProxyConfig**](NodeProxyConfig.md) | Proxy settings overridden per region identified by region code. | [optional] 
**AzNetworking** | Pointer to [**map[string]NodeProxyConfig**](NodeProxyConfig.md) | Proxy settings overridden per Availability Zone identified by AZ code. | [optional] 

## Methods

### NewPerProviderUpdateProxyConfigSpec

`func NewPerProviderUpdateProxyConfigSpec(provider string, ) *PerProviderUpdateProxyConfigSpec`

NewPerProviderUpdateProxyConfigSpec instantiates a new PerProviderUpdateProxyConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerProviderUpdateProxyConfigSpecWithDefaults

`func NewPerProviderUpdateProxyConfigSpecWithDefaults() *PerProviderUpdateProxyConfigSpec`

NewPerProviderUpdateProxyConfigSpecWithDefaults instantiates a new PerProviderUpdateProxyConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *PerProviderUpdateProxyConfigSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PerProviderUpdateProxyConfigSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PerProviderUpdateProxyConfigSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProxyConfig

`func (o *PerProviderUpdateProxyConfigSpec) GetProxyConfig() NodeProxyConfig`

GetProxyConfig returns the ProxyConfig field if non-nil, zero value otherwise.

### GetProxyConfigOk

`func (o *PerProviderUpdateProxyConfigSpec) GetProxyConfigOk() (*NodeProxyConfig, bool)`

GetProxyConfigOk returns a tuple with the ProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfig

`func (o *PerProviderUpdateProxyConfigSpec) SetProxyConfig(v NodeProxyConfig)`

SetProxyConfig sets ProxyConfig field to given value.

### HasProxyConfig

`func (o *PerProviderUpdateProxyConfigSpec) HasProxyConfig() bool`

HasProxyConfig returns a boolean if a field has been set.

### GetRegionNetworking

`func (o *PerProviderUpdateProxyConfigSpec) GetRegionNetworking() map[string]NodeProxyConfig`

GetRegionNetworking returns the RegionNetworking field if non-nil, zero value otherwise.

### GetRegionNetworkingOk

`func (o *PerProviderUpdateProxyConfigSpec) GetRegionNetworkingOk() (*map[string]NodeProxyConfig, bool)`

GetRegionNetworkingOk returns a tuple with the RegionNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNetworking

`func (o *PerProviderUpdateProxyConfigSpec) SetRegionNetworking(v map[string]NodeProxyConfig)`

SetRegionNetworking sets RegionNetworking field to given value.

### HasRegionNetworking

`func (o *PerProviderUpdateProxyConfigSpec) HasRegionNetworking() bool`

HasRegionNetworking returns a boolean if a field has been set.

### GetAzNetworking

`func (o *PerProviderUpdateProxyConfigSpec) GetAzNetworking() map[string]NodeProxyConfig`

GetAzNetworking returns the AzNetworking field if non-nil, zero value otherwise.

### GetAzNetworkingOk

`func (o *PerProviderUpdateProxyConfigSpec) GetAzNetworkingOk() (*map[string]NodeProxyConfig, bool)`

GetAzNetworkingOk returns a tuple with the AzNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNetworking

`func (o *PerProviderUpdateProxyConfigSpec) SetAzNetworking(v map[string]NodeProxyConfig)`

SetAzNetworking sets AzNetworking field to given value.

### HasAzNetworking

`func (o *PerProviderUpdateProxyConfigSpec) HasAzNetworking() bool`

HasAzNetworking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


