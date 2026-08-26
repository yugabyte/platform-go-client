# ProviderNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TserverSpecification** | Pointer to [**ProviderNodeSpec**](ProviderNodeSpec.md) |  | [optional] 
**MasterSpecification** | Pointer to [**ProviderNodeSpec**](ProviderNodeSpec.md) |  | [optional] 

## Methods

### NewProviderNodesSpec

`func NewProviderNodesSpec() *ProviderNodesSpec`

NewProviderNodesSpec instantiates a new ProviderNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderNodesSpecWithDefaults

`func NewProviderNodesSpecWithDefaults() *ProviderNodesSpec`

NewProviderNodesSpecWithDefaults instantiates a new ProviderNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserverSpecification

`func (o *ProviderNodesSpec) GetTserverSpecification() ProviderNodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *ProviderNodesSpec) GetTserverSpecificationOk() (*ProviderNodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *ProviderNodesSpec) SetTserverSpecification(v ProviderNodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *ProviderNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.

### GetMasterSpecification

`func (o *ProviderNodesSpec) GetMasterSpecification() ProviderNodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *ProviderNodesSpec) GetMasterSpecificationOk() (*ProviderNodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *ProviderNodesSpec) SetMasterSpecification(v ProviderNodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *ProviderNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


