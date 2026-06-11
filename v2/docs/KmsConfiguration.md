# KmsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**KmsConfigurationSpec**](KmsConfigurationSpec.md) |  | [optional] 
**Info** | Pointer to [**KmsConfigurationInfo**](KmsConfigurationInfo.md) |  | [optional] 

## Methods

### NewKmsConfiguration

`func NewKmsConfiguration() *KmsConfiguration`

NewKmsConfiguration instantiates a new KmsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigurationWithDefaults

`func NewKmsConfigurationWithDefaults() *KmsConfiguration`

NewKmsConfigurationWithDefaults instantiates a new KmsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *KmsConfiguration) GetSpec() KmsConfigurationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KmsConfiguration) GetSpecOk() (*KmsConfigurationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KmsConfiguration) SetSpec(v KmsConfigurationSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *KmsConfiguration) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetInfo

`func (o *KmsConfiguration) GetInfo() KmsConfigurationInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *KmsConfiguration) GetInfoOk() (*KmsConfigurationInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *KmsConfiguration) SetInfo(v KmsConfigurationInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *KmsConfiguration) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


