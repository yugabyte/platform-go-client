# ProviderRegionNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TserverSpecification** | Pointer to [**ProviderNodeSpec**](ProviderNodeSpec.md) |  | [optional] 
**MasterSpecification** | Pointer to [**ProviderNodeSpec**](ProviderNodeSpec.md) |  | [optional] 
**RegionCode** | **string** | Region code within the cloud provider. | 
**AzNodesSpecs** | Pointer to [**[]ProviderAzNodesSpec**](ProviderAzNodesSpec.md) | Per-availability-zone node specification overrides. | [optional] 

## Methods

### NewProviderRegionNodesSpec

`func NewProviderRegionNodesSpec(regionCode string, ) *ProviderRegionNodesSpec`

NewProviderRegionNodesSpec instantiates a new ProviderRegionNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRegionNodesSpecWithDefaults

`func NewProviderRegionNodesSpecWithDefaults() *ProviderRegionNodesSpec`

NewProviderRegionNodesSpecWithDefaults instantiates a new ProviderRegionNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserverSpecification

`func (o *ProviderRegionNodesSpec) GetTserverSpecification() ProviderNodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *ProviderRegionNodesSpec) GetTserverSpecificationOk() (*ProviderNodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *ProviderRegionNodesSpec) SetTserverSpecification(v ProviderNodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *ProviderRegionNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.

### GetMasterSpecification

`func (o *ProviderRegionNodesSpec) GetMasterSpecification() ProviderNodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *ProviderRegionNodesSpec) GetMasterSpecificationOk() (*ProviderNodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *ProviderRegionNodesSpec) SetMasterSpecification(v ProviderNodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *ProviderRegionNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetRegionCode

`func (o *ProviderRegionNodesSpec) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ProviderRegionNodesSpec) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ProviderRegionNodesSpec) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetAzNodesSpecs

`func (o *ProviderRegionNodesSpec) GetAzNodesSpecs() []ProviderAzNodesSpec`

GetAzNodesSpecs returns the AzNodesSpecs field if non-nil, zero value otherwise.

### GetAzNodesSpecsOk

`func (o *ProviderRegionNodesSpec) GetAzNodesSpecsOk() (*[]ProviderAzNodesSpec, bool)`

GetAzNodesSpecsOk returns a tuple with the AzNodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNodesSpecs

`func (o *ProviderRegionNodesSpec) SetAzNodesSpecs(v []ProviderAzNodesSpec)`

SetAzNodesSpecs sets AzNodesSpecs field to given value.

### HasAzNodesSpecs

`func (o *ProviderRegionNodesSpec) HasAzNodesSpecs() bool`

HasAzNodesSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


