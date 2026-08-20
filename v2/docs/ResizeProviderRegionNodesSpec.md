# ResizeProviderRegionNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TserverSpecification** | Pointer to [**ResizeProviderNodeSpec**](ResizeProviderNodeSpec.md) |  | [optional] 
**MasterSpecification** | Pointer to [**ResizeProviderNodeSpec**](ResizeProviderNodeSpec.md) |  | [optional] 
**RegionCode** | **string** | Region code within the cloud provider. | 
**AzNodesSpecs** | Pointer to [**[]ResizeProviderAzNodesSpec**](ResizeProviderAzNodesSpec.md) | Per-availability-zone resize node specification overrides. | [optional] 

## Methods

### NewResizeProviderRegionNodesSpec

`func NewResizeProviderRegionNodesSpec(regionCode string, ) *ResizeProviderRegionNodesSpec`

NewResizeProviderRegionNodesSpec instantiates a new ResizeProviderRegionNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeProviderRegionNodesSpecWithDefaults

`func NewResizeProviderRegionNodesSpecWithDefaults() *ResizeProviderRegionNodesSpec`

NewResizeProviderRegionNodesSpecWithDefaults instantiates a new ResizeProviderRegionNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserverSpecification

`func (o *ResizeProviderRegionNodesSpec) GetTserverSpecification() ResizeProviderNodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *ResizeProviderRegionNodesSpec) GetTserverSpecificationOk() (*ResizeProviderNodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *ResizeProviderRegionNodesSpec) SetTserverSpecification(v ResizeProviderNodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *ResizeProviderRegionNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.

### GetMasterSpecification

`func (o *ResizeProviderRegionNodesSpec) GetMasterSpecification() ResizeProviderNodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *ResizeProviderRegionNodesSpec) GetMasterSpecificationOk() (*ResizeProviderNodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *ResizeProviderRegionNodesSpec) SetMasterSpecification(v ResizeProviderNodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *ResizeProviderRegionNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetRegionCode

`func (o *ResizeProviderRegionNodesSpec) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ResizeProviderRegionNodesSpec) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ResizeProviderRegionNodesSpec) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetAzNodesSpecs

`func (o *ResizeProviderRegionNodesSpec) GetAzNodesSpecs() []ResizeProviderAzNodesSpec`

GetAzNodesSpecs returns the AzNodesSpecs field if non-nil, zero value otherwise.

### GetAzNodesSpecsOk

`func (o *ResizeProviderRegionNodesSpec) GetAzNodesSpecsOk() (*[]ResizeProviderAzNodesSpec, bool)`

GetAzNodesSpecsOk returns a tuple with the AzNodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNodesSpecs

`func (o *ResizeProviderRegionNodesSpec) SetAzNodesSpecs(v []ResizeProviderAzNodesSpec)`

SetAzNodesSpecs sets AzNodesSpecs field to given value.

### HasAzNodesSpecs

`func (o *ResizeProviderRegionNodesSpec) HasAzNodesSpecs() bool`

HasAzNodesSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


