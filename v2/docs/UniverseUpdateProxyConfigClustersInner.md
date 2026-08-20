# UniverseUpdateProxyConfigClustersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The system generated cluster uuid to update. This can be fetched from ClusterInfo. | 
**NetworkingSpec** | Pointer to [**UpdateProxyConfigSpec**](UpdateProxyConfigSpec.md) |  | [optional] 
**ProviderProxySpecs** | Pointer to [**[]PerProviderUpdateProxyConfigSpec**](PerProviderUpdateProxyConfigSpec.md) | Proxy settings per provider for multicloud clusters. | [optional] 

## Methods

### NewUniverseUpdateProxyConfigClustersInner

`func NewUniverseUpdateProxyConfigClustersInner(uuid string, ) *UniverseUpdateProxyConfigClustersInner`

NewUniverseUpdateProxyConfigClustersInner instantiates a new UniverseUpdateProxyConfigClustersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseUpdateProxyConfigClustersInnerWithDefaults

`func NewUniverseUpdateProxyConfigClustersInnerWithDefaults() *UniverseUpdateProxyConfigClustersInner`

NewUniverseUpdateProxyConfigClustersInnerWithDefaults instantiates a new UniverseUpdateProxyConfigClustersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *UniverseUpdateProxyConfigClustersInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UniverseUpdateProxyConfigClustersInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UniverseUpdateProxyConfigClustersInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetNetworkingSpec

`func (o *UniverseUpdateProxyConfigClustersInner) GetNetworkingSpec() UpdateProxyConfigSpec`

GetNetworkingSpec returns the NetworkingSpec field if non-nil, zero value otherwise.

### GetNetworkingSpecOk

`func (o *UniverseUpdateProxyConfigClustersInner) GetNetworkingSpecOk() (*UpdateProxyConfigSpec, bool)`

GetNetworkingSpecOk returns a tuple with the NetworkingSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingSpec

`func (o *UniverseUpdateProxyConfigClustersInner) SetNetworkingSpec(v UpdateProxyConfigSpec)`

SetNetworkingSpec sets NetworkingSpec field to given value.

### HasNetworkingSpec

`func (o *UniverseUpdateProxyConfigClustersInner) HasNetworkingSpec() bool`

HasNetworkingSpec returns a boolean if a field has been set.

### GetProviderProxySpecs

`func (o *UniverseUpdateProxyConfigClustersInner) GetProviderProxySpecs() []PerProviderUpdateProxyConfigSpec`

GetProviderProxySpecs returns the ProviderProxySpecs field if non-nil, zero value otherwise.

### GetProviderProxySpecsOk

`func (o *UniverseUpdateProxyConfigClustersInner) GetProviderProxySpecsOk() (*[]PerProviderUpdateProxyConfigSpec, bool)`

GetProviderProxySpecsOk returns a tuple with the ProviderProxySpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderProxySpecs

`func (o *UniverseUpdateProxyConfigClustersInner) SetProviderProxySpecs(v []PerProviderUpdateProxyConfigSpec)`

SetProviderProxySpecs sets ProviderProxySpecs field to given value.

### HasProviderProxySpecs

`func (o *UniverseUpdateProxyConfigClustersInner) HasProviderProxySpecs() bool`

HasProviderProxySpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


