# PerProviderResizeNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Cloud provider UUID. | 
**NodesSpec** | [**ResizeProviderRootNodesSpec**](ResizeProviderRootNodesSpec.md) |  | 

## Methods

### NewPerProviderResizeNodesSpec

`func NewPerProviderResizeNodesSpec(provider string, nodesSpec ResizeProviderRootNodesSpec, ) *PerProviderResizeNodesSpec`

NewPerProviderResizeNodesSpec instantiates a new PerProviderResizeNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerProviderResizeNodesSpecWithDefaults

`func NewPerProviderResizeNodesSpecWithDefaults() *PerProviderResizeNodesSpec`

NewPerProviderResizeNodesSpecWithDefaults instantiates a new PerProviderResizeNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *PerProviderResizeNodesSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PerProviderResizeNodesSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PerProviderResizeNodesSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetNodesSpec

`func (o *PerProviderResizeNodesSpec) GetNodesSpec() ResizeProviderRootNodesSpec`

GetNodesSpec returns the NodesSpec field if non-nil, zero value otherwise.

### GetNodesSpecOk

`func (o *PerProviderResizeNodesSpec) GetNodesSpecOk() (*ResizeProviderRootNodesSpec, bool)`

GetNodesSpecOk returns a tuple with the NodesSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesSpec

`func (o *PerProviderResizeNodesSpec) SetNodesSpec(v ResizeProviderRootNodesSpec)`

SetNodesSpec sets NodesSpec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


