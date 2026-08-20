# ResizeProviderRootNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TserverSpecification** | Pointer to [**ResizeProviderNodeSpec**](ResizeProviderNodeSpec.md) |  | [optional] 
**MasterSpecification** | Pointer to [**ResizeProviderNodeSpec**](ResizeProviderNodeSpec.md) |  | [optional] 
**RegionNodesSpecs** | Pointer to [**[]ResizeProviderRegionNodesSpec**](ResizeProviderRegionNodesSpec.md) | Per-region resize node specification overrides. | [optional] 

## Methods

### NewResizeProviderRootNodesSpec

`func NewResizeProviderRootNodesSpec() *ResizeProviderRootNodesSpec`

NewResizeProviderRootNodesSpec instantiates a new ResizeProviderRootNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeProviderRootNodesSpecWithDefaults

`func NewResizeProviderRootNodesSpecWithDefaults() *ResizeProviderRootNodesSpec`

NewResizeProviderRootNodesSpecWithDefaults instantiates a new ResizeProviderRootNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserverSpecification

`func (o *ResizeProviderRootNodesSpec) GetTserverSpecification() ResizeProviderNodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *ResizeProviderRootNodesSpec) GetTserverSpecificationOk() (*ResizeProviderNodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *ResizeProviderRootNodesSpec) SetTserverSpecification(v ResizeProviderNodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *ResizeProviderRootNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.

### GetMasterSpecification

`func (o *ResizeProviderRootNodesSpec) GetMasterSpecification() ResizeProviderNodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *ResizeProviderRootNodesSpec) GetMasterSpecificationOk() (*ResizeProviderNodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *ResizeProviderRootNodesSpec) SetMasterSpecification(v ResizeProviderNodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *ResizeProviderRootNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetRegionNodesSpecs

`func (o *ResizeProviderRootNodesSpec) GetRegionNodesSpecs() []ResizeProviderRegionNodesSpec`

GetRegionNodesSpecs returns the RegionNodesSpecs field if non-nil, zero value otherwise.

### GetRegionNodesSpecsOk

`func (o *ResizeProviderRootNodesSpec) GetRegionNodesSpecsOk() (*[]ResizeProviderRegionNodesSpec, bool)`

GetRegionNodesSpecsOk returns a tuple with the RegionNodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNodesSpecs

`func (o *ResizeProviderRootNodesSpec) SetRegionNodesSpecs(v []ResizeProviderRegionNodesSpec)`

SetRegionNodesSpecs sets RegionNodesSpecs field to given value.

### HasRegionNodesSpecs

`func (o *ResizeProviderRootNodesSpec) HasRegionNodesSpecs() bool`

HasRegionNodesSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


