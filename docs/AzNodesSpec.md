# AzNodesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzCode** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**MasterSpecification** | Pointer to [**NodeSpec**](NodeSpec.md) |  | [optional] 
**TserverSpecification** | Pointer to [**NodeSpec**](NodeSpec.md) |  | [optional] 

## Methods

### NewAzNodesSpec

`func NewAzNodesSpec(code string, ) *AzNodesSpec`

NewAzNodesSpec instantiates a new AzNodesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzNodesSpecWithDefaults

`func NewAzNodesSpecWithDefaults() *AzNodesSpec`

NewAzNodesSpecWithDefaults instantiates a new AzNodesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzCode

`func (o *AzNodesSpec) GetAzCode() string`

GetAzCode returns the AzCode field if non-nil, zero value otherwise.

### GetAzCodeOk

`func (o *AzNodesSpec) GetAzCodeOk() (*string, bool)`

GetAzCodeOk returns a tuple with the AzCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCode

`func (o *AzNodesSpec) SetAzCode(v string)`

SetAzCode sets AzCode field to given value.

### HasAzCode

`func (o *AzNodesSpec) HasAzCode() bool`

HasAzCode returns a boolean if a field has been set.

### GetCode

`func (o *AzNodesSpec) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AzNodesSpec) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AzNodesSpec) SetCode(v string)`

SetCode sets Code field to given value.


### GetMasterSpecification

`func (o *AzNodesSpec) GetMasterSpecification() NodeSpec`

GetMasterSpecification returns the MasterSpecification field if non-nil, zero value otherwise.

### GetMasterSpecificationOk

`func (o *AzNodesSpec) GetMasterSpecificationOk() (*NodeSpec, bool)`

GetMasterSpecificationOk returns a tuple with the MasterSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterSpecification

`func (o *AzNodesSpec) SetMasterSpecification(v NodeSpec)`

SetMasterSpecification sets MasterSpecification field to given value.

### HasMasterSpecification

`func (o *AzNodesSpec) HasMasterSpecification() bool`

HasMasterSpecification returns a boolean if a field has been set.

### GetTserverSpecification

`func (o *AzNodesSpec) GetTserverSpecification() NodeSpec`

GetTserverSpecification returns the TserverSpecification field if non-nil, zero value otherwise.

### GetTserverSpecificationOk

`func (o *AzNodesSpec) GetTserverSpecificationOk() (*NodeSpec, bool)`

GetTserverSpecificationOk returns a tuple with the TserverSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverSpecification

`func (o *AzNodesSpec) SetTserverSpecification(v NodeSpec)`

SetTserverSpecification sets TserverSpecification field to given value.

### HasTserverSpecification

`func (o *AzNodesSpec) HasTserverSpecification() bool`

HasTserverSpecification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


