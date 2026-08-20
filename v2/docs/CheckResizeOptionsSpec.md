# CheckResizeOptionsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | **string** | Cluster UUID to evaluate. | 
**NodeSpec** | Pointer to [**ClusterNodeSpec**](ClusterNodeSpec.md) |  | [optional] 
**ProviderNodesSpecs** | Pointer to [**[]PerProviderResizeNodesSpec**](PerProviderResizeNodesSpec.md) | Proposed node settings per provider for multicloud clusters. | [optional] 

## Methods

### NewCheckResizeOptionsSpec

`func NewCheckResizeOptionsSpec(clusterUuid string, ) *CheckResizeOptionsSpec`

NewCheckResizeOptionsSpec instantiates a new CheckResizeOptionsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckResizeOptionsSpecWithDefaults

`func NewCheckResizeOptionsSpecWithDefaults() *CheckResizeOptionsSpec`

NewCheckResizeOptionsSpecWithDefaults instantiates a new CheckResizeOptionsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *CheckResizeOptionsSpec) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *CheckResizeOptionsSpec) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *CheckResizeOptionsSpec) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.


### GetNodeSpec

`func (o *CheckResizeOptionsSpec) GetNodeSpec() ClusterNodeSpec`

GetNodeSpec returns the NodeSpec field if non-nil, zero value otherwise.

### GetNodeSpecOk

`func (o *CheckResizeOptionsSpec) GetNodeSpecOk() (*ClusterNodeSpec, bool)`

GetNodeSpecOk returns a tuple with the NodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSpec

`func (o *CheckResizeOptionsSpec) SetNodeSpec(v ClusterNodeSpec)`

SetNodeSpec sets NodeSpec field to given value.

### HasNodeSpec

`func (o *CheckResizeOptionsSpec) HasNodeSpec() bool`

HasNodeSpec returns a boolean if a field has been set.

### GetProviderNodesSpecs

`func (o *CheckResizeOptionsSpec) GetProviderNodesSpecs() []PerProviderResizeNodesSpec`

GetProviderNodesSpecs returns the ProviderNodesSpecs field if non-nil, zero value otherwise.

### GetProviderNodesSpecsOk

`func (o *CheckResizeOptionsSpec) GetProviderNodesSpecsOk() (*[]PerProviderResizeNodesSpec, bool)`

GetProviderNodesSpecsOk returns a tuple with the ProviderNodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNodesSpecs

`func (o *CheckResizeOptionsSpec) SetProviderNodesSpecs(v []PerProviderResizeNodesSpec)`

SetProviderNodesSpecs sets ProviderNodesSpecs field to given value.

### HasProviderNodesSpecs

`func (o *CheckResizeOptionsSpec) HasProviderNodesSpecs() bool`

HasProviderNodesSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


