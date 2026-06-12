# CustomerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**CustomerConfigSpec**](CustomerConfigSpec.md) |  | [optional] 
**Info** | Pointer to [**CustomerConfigInfo**](CustomerConfigInfo.md) |  | [optional] 

## Methods

### NewCustomerConfig

`func NewCustomerConfig() *CustomerConfig`

NewCustomerConfig instantiates a new CustomerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerConfigWithDefaults

`func NewCustomerConfigWithDefaults() *CustomerConfig`

NewCustomerConfigWithDefaults instantiates a new CustomerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *CustomerConfig) GetSpec() CustomerConfigSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CustomerConfig) GetSpecOk() (*CustomerConfigSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CustomerConfig) SetSpec(v CustomerConfigSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CustomerConfig) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetInfo

`func (o *CustomerConfig) GetInfo() CustomerConfigInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CustomerConfig) GetInfoOk() (*CustomerConfigInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CustomerConfig) SetInfo(v CustomerConfigInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CustomerConfig) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


