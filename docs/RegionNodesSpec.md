# RegionNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzNodesSpecs** | Pointer to [**[]AzNodesSpec**](AzNodesSpec.md) |  | [optional] 
**Code** | **string** |  | 
**MasterSpecification** | Pointer to [**NodeSpec**](NodeSpec.md) |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**TserverSpecification** | Pointer to [**NodeSpec**](NodeSpec.md) |  | [optional] 

## Methods

### NewRegionNodesSpec

`func NewRegionNodesSpec(code string, ) *RegionNodesSpec`

NewRegionNodesSpec instantiates a new RegionNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionNodesSpecWithDefaults

`func NewRegionNodesSpecWithDefaults() *RegionNodesSpec`

NewRegionNodesSpecWithDefaults instantiates a new RegionNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzNodesSpecs

`func (o *RegionNodesSpec) GetAzNodesSpecs() []AzNodesSpec`

GetAzNodesSpecs returns the AzNodesSpecs field if non-nil, zero value otherwise.

### GetAzNodesSpecsOk

`func (o *RegionNodesSpec) GetAzNodesSpecsOk() (*[]AzNodesSpec, bool)`

GetAzNodesSpecsOk returns a tuple with the AzNodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNodesSpecs

`func (o *RegionNodesSpec) SetAzNodesSpecs(v []AzNodesSpec)`

SetAzNodesSpecs sets AzNodesSpecs field to given value.

### HasAzNodesSpecs

`func (o *RegionNodesSpec) HasAzNodesSpecs() bool`

HasAzNodesSpecs returns a boolean if a field has been set.

### GetCode

`func (o *RegionNodesSpec) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RegionNodesSpec) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RegionNodesSpec) SetCode(v string)`

SetCode sets Code field to given value.


### GetMasterSpecification

`func (o *RegionNodesSpec) GetMasterSpecification() NodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *RegionNodesSpec) GetMasterSpecificationOk() (*NodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *RegionNodesSpec) SetMasterSpecification(v NodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *RegionNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetRegionCode

`func (o *RegionNodesSpec) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *RegionNodesSpec) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *RegionNodesSpec) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *RegionNodesSpec) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetTserverSpecification

`func (o *RegionNodesSpec) GetTserverSpecification() NodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *RegionNodesSpec) GetTserverSpecificationOk() (*NodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *RegionNodesSpec) SetTserverSpecification(v NodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *RegionNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


