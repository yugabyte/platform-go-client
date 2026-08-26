# ResizeProviderAzNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TserverSpecification** | Pointer to [**ResizeProviderNodeSpec**](ResizeProviderNodeSpec.md) |  | [optional] 
**MasterSpecification** | Pointer to [**ResizeProviderNodeSpec**](ResizeProviderNodeSpec.md) |  | [optional] 
**AzCode** | **string** | Availability zone code within the region. | 

## Methods

### NewResizeProviderAzNodesSpec

`func NewResizeProviderAzNodesSpec(azCode string, ) *ResizeProviderAzNodesSpec`

NewResizeProviderAzNodesSpec instantiates a new ResizeProviderAzNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeProviderAzNodesSpecWithDefaults

`func NewResizeProviderAzNodesSpecWithDefaults() *ResizeProviderAzNodesSpec`

NewResizeProviderAzNodesSpecWithDefaults instantiates a new ResizeProviderAzNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserverSpecification

`func (o *ResizeProviderAzNodesSpec) GetTserverSpecification() ResizeProviderNodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *ResizeProviderAzNodesSpec) GetTserverSpecificationOk() (*ResizeProviderNodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *ResizeProviderAzNodesSpec) SetTserverSpecification(v ResizeProviderNodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *ResizeProviderAzNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.

### GetMasterSpecification

`func (o *ResizeProviderAzNodesSpec) GetMasterSpecification() ResizeProviderNodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *ResizeProviderAzNodesSpec) GetMasterSpecificationOk() (*ResizeProviderNodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *ResizeProviderAzNodesSpec) SetMasterSpecification(v ResizeProviderNodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *ResizeProviderAzNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetAzCode

`func (o *ResizeProviderAzNodesSpec) GetAzCode() string`

GetAzCode returns the AzCode field if non-nil, zero value otherwise.

### GetAzCodeOk

`func (o *ResizeProviderAzNodesSpec) GetAzCodeOk() (*string, bool)`

GetAzCodeOk returns a tuple with the AzCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCode

`func (o *ResizeProviderAzNodesSpec) SetAzCode(v string)`

SetAzCode sets AzCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


