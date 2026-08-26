# ProviderAzNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TserverSpecification** | Pointer to [**ProviderNodeSpec**](ProviderNodeSpec.md) |  | [optional] 
**MasterSpecification** | Pointer to [**ProviderNodeSpec**](ProviderNodeSpec.md) |  | [optional] 
**AzCode** | **string** | Availability zone code within the region. | 

## Methods

### NewProviderAzNodesSpec

`func NewProviderAzNodesSpec(azCode string, ) *ProviderAzNodesSpec`

NewProviderAzNodesSpec instantiates a new ProviderAzNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderAzNodesSpecWithDefaults

`func NewProviderAzNodesSpecWithDefaults() *ProviderAzNodesSpec`

NewProviderAzNodesSpecWithDefaults instantiates a new ProviderAzNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserverSpecification

`func (o *ProviderAzNodesSpec) GetTserverSpecification() ProviderNodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *ProviderAzNodesSpec) GetTserverSpecificationOk() (*ProviderNodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *ProviderAzNodesSpec) SetTserverSpecification(v ProviderNodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *ProviderAzNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.

### GetMasterSpecification

`func (o *ProviderAzNodesSpec) GetMasterSpecification() ProviderNodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *ProviderAzNodesSpec) GetMasterSpecificationOk() (*ProviderNodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *ProviderAzNodesSpec) SetMasterSpecification(v ProviderNodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *ProviderAzNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetAzCode

`func (o *ProviderAzNodesSpec) GetAzCode() string`

GetAzCode returns the AzCode field if non-nil, zero value otherwise.

### GetAzCodeOk

`func (o *ProviderAzNodesSpec) GetAzCodeOk() (*string, bool)`

GetAzCodeOk returns a tuple with the AzCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCode

`func (o *ProviderAzNodesSpec) SetAzCode(v string)`

SetAzCode sets AzCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


