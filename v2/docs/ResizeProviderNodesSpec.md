# ResizeProviderNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TserverSpecification** | Pointer to [**ResizeProviderNodeSpec**](ResizeProviderNodeSpec.md) |  | [optional] 
**MasterSpecification** | Pointer to [**ResizeProviderNodeSpec**](ResizeProviderNodeSpec.md) |  | [optional] 

## Methods

### NewResizeProviderNodesSpec

`func NewResizeProviderNodesSpec() *ResizeProviderNodesSpec`

NewResizeProviderNodesSpec instantiates a new ResizeProviderNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeProviderNodesSpecWithDefaults

`func NewResizeProviderNodesSpecWithDefaults() *ResizeProviderNodesSpec`

NewResizeProviderNodesSpecWithDefaults instantiates a new ResizeProviderNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserverSpecification

`func (o *ResizeProviderNodesSpec) GetTserverSpecification() ResizeProviderNodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *ResizeProviderNodesSpec) GetTserverSpecificationOk() (*ResizeProviderNodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *ResizeProviderNodesSpec) SetTserverSpecification(v ResizeProviderNodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *ResizeProviderNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.

### GetMasterSpecification

`func (o *ResizeProviderNodesSpec) GetMasterSpecification() ResizeProviderNodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *ResizeProviderNodesSpec) GetMasterSpecificationOk() (*ResizeProviderNodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *ResizeProviderNodesSpec) SetMasterSpecification(v ResizeProviderNodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *ResizeProviderNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


