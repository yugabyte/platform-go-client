# UpdateProxyConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProxyConfig** | Pointer to [**NodeProxyConfig**](NodeProxyConfig.md) |  | [optional] 
**AzNetworking** | Pointer to [**map[string]NodeProxyConfig**](NodeProxyConfig.md) | Proxy settings overridden per Availability Zone identified by AZ uuid. | [optional] 

## Methods

### NewUpdateProxyConfigSpec

`func NewUpdateProxyConfigSpec() *UpdateProxyConfigSpec`

NewUpdateProxyConfigSpec instantiates a new UpdateProxyConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProxyConfigSpecWithDefaults

`func NewUpdateProxyConfigSpecWithDefaults() *UpdateProxyConfigSpec`

NewUpdateProxyConfigSpecWithDefaults instantiates a new UpdateProxyConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProxyConfig

`func (o *UpdateProxyConfigSpec) GetProxyConfig() NodeProxyConfig`

GetProxyConfig returns the ProxyConfig field if non-nil, zero value otherwise.

### GetProxyConfigOk

`func (o *UpdateProxyConfigSpec) GetProxyConfigOk() (*NodeProxyConfig, bool)`

GetProxyConfigOk returns a tuple with the ProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfig

`func (o *UpdateProxyConfigSpec) SetProxyConfig(v NodeProxyConfig)`

SetProxyConfig sets ProxyConfig field to given value.

### HasProxyConfig

`func (o *UpdateProxyConfigSpec) HasProxyConfig() bool`

HasProxyConfig returns a boolean if a field has been set.

### GetAzNetworking

`func (o *UpdateProxyConfigSpec) GetAzNetworking() map[string]NodeProxyConfig`

GetAzNetworking returns the AzNetworking field if non-nil, zero value otherwise.

### GetAzNetworkingOk

`func (o *UpdateProxyConfigSpec) GetAzNetworkingOk() (*map[string]NodeProxyConfig, bool)`

GetAzNetworkingOk returns a tuple with the AzNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNetworking

`func (o *UpdateProxyConfigSpec) SetAzNetworking(v map[string]NodeProxyConfig)`

SetAzNetworking sets AzNetworking field to given value.

### HasAzNetworking

`func (o *UpdateProxyConfigSpec) HasAzNetworking() bool`

HasAzNetworking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


