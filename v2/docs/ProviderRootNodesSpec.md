# ProviderRootNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TserverSpecification** | Pointer to [**ProviderNodeSpec**](ProviderNodeSpec.md) |  | [optional] 
**MasterSpecification** | Pointer to [**ProviderNodeSpec**](ProviderNodeSpec.md) |  | [optional] 
**RegionNodesSpecs** | Pointer to [**[]ProviderRegionNodesSpec**](ProviderRegionNodesSpec.md) | Per-region node specification overrides. | [optional] 

## Methods

### NewProviderRootNodesSpec

`func NewProviderRootNodesSpec() *ProviderRootNodesSpec`

NewProviderRootNodesSpec instantiates a new ProviderRootNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRootNodesSpecWithDefaults

`func NewProviderRootNodesSpecWithDefaults() *ProviderRootNodesSpec`

NewProviderRootNodesSpecWithDefaults instantiates a new ProviderRootNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserverSpecification

`func (o *ProviderRootNodesSpec) GetTserverSpecification() ProviderNodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *ProviderRootNodesSpec) GetTserverSpecificationOk() (*ProviderNodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *ProviderRootNodesSpec) SetTserverSpecification(v ProviderNodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *ProviderRootNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.

### GetMasterSpecification

`func (o *ProviderRootNodesSpec) GetMasterSpecification() ProviderNodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *ProviderRootNodesSpec) GetMasterSpecificationOk() (*ProviderNodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *ProviderRootNodesSpec) SetMasterSpecification(v ProviderNodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *ProviderRootNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetRegionNodesSpecs

`func (o *ProviderRootNodesSpec) GetRegionNodesSpecs() []ProviderRegionNodesSpec`

GetRegionNodesSpecs returns the RegionNodesSpecs field if non-nil, zero value otherwise.

### GetRegionNodesSpecsOk

`func (o *ProviderRootNodesSpec) GetRegionNodesSpecsOk() (*[]ProviderRegionNodesSpec, bool)`

GetRegionNodesSpecsOk returns a tuple with the RegionNodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNodesSpecs

`func (o *ProviderRootNodesSpec) SetRegionNodesSpecs(v []ProviderRegionNodesSpec)`

SetRegionNodesSpecs sets RegionNodesSpecs field to given value.

### HasRegionNodesSpecs

`func (o *ProviderRootNodesSpec) HasRegionNodesSpecs() bool`

HasRegionNodesSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


