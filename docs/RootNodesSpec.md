# RootNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterSpecification** | Pointer to [**NodeSpec**](NodeSpec.md) |  | [optional] 
**RegionNodesSpecs** | Pointer to [**[]RegionNodesSpec**](RegionNodesSpec.md) |  | [optional] 
**TserverSpecification** | Pointer to [**NodeSpec**](NodeSpec.md) |  | [optional] 

## Methods

### NewRootNodesSpec

`func NewRootNodesSpec() *RootNodesSpec`

NewRootNodesSpec instantiates a new RootNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootNodesSpecWithDefaults

`func NewRootNodesSpecWithDefaults() *RootNodesSpec`

NewRootNodesSpecWithDefaults instantiates a new RootNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterSpecification

`func (o *RootNodesSpec) GetMasterSpecification() NodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *RootNodesSpec) GetMasterSpecificationOk() (*NodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *RootNodesSpec) SetMasterSpecification(v NodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *RootNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetRegionNodesSpecs

`func (o *RootNodesSpec) GetRegionNodesSpecs() []RegionNodesSpec`

GetRegionNodesSpecs returns the RegionNodesSpecs field if non-nil, zero value otherwise.

### GetRegionNodesSpecsOk

`func (o *RootNodesSpec) GetRegionNodesSpecsOk() (*[]RegionNodesSpec, bool)`

GetRegionNodesSpecsOk returns a tuple with the RegionNodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNodesSpecs

`func (o *RootNodesSpec) SetRegionNodesSpecs(v []RegionNodesSpec)`

SetRegionNodesSpecs sets RegionNodesSpecs field to given value.

### HasRegionNodesSpecs

`func (o *RootNodesSpec) HasRegionNodesSpecs() bool`

HasRegionNodesSpecs returns a boolean if a field has been set.

### GetTserverSpecification

`func (o *RootNodesSpec) GetTserverSpecification() NodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *RootNodesSpec) GetTserverSpecificationOk() (*NodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *RootNodesSpec) SetTserverSpecification(v NodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *RootNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


