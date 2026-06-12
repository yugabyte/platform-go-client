# KmsConfigurationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | KMS configuration unique identifier. | [optional] [readonly] 
**InUse** | Pointer to **bool** | Whether this KMS config is attached to any universe. | [optional] [readonly] 
**UniverseDetails** | Pointer to [**[]UniverseDetailSubset**](UniverseDetailSubset.md) | Universes using this KMS configuration (subset fields only). | [optional] [readonly] 

## Methods

### NewKmsConfigurationInfo

`func NewKmsConfigurationInfo() *KmsConfigurationInfo`

NewKmsConfigurationInfo instantiates a new KmsConfigurationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigurationInfoWithDefaults

`func NewKmsConfigurationInfoWithDefaults() *KmsConfigurationInfo`

NewKmsConfigurationInfoWithDefaults instantiates a new KmsConfigurationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *KmsConfigurationInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *KmsConfigurationInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *KmsConfigurationInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *KmsConfigurationInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetInUse

`func (o *KmsConfigurationInfo) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *KmsConfigurationInfo) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *KmsConfigurationInfo) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *KmsConfigurationInfo) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetUniverseDetails

`func (o *KmsConfigurationInfo) GetUniverseDetails() []UniverseDetailSubset`

GetUniverseDetails returns the UniverseDetails field if non-nil, zero value otherwise.

### GetUniverseDetailsOk

`func (o *KmsConfigurationInfo) GetUniverseDetailsOk() (*[]UniverseDetailSubset, bool)`

GetUniverseDetailsOk returns a tuple with the UniverseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDetails

`func (o *KmsConfigurationInfo) SetUniverseDetails(v []UniverseDetailSubset)`

SetUniverseDetails sets UniverseDetails field to given value.

### HasUniverseDetails

`func (o *KmsConfigurationInfo) HasUniverseDetails() bool`

HasUniverseDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


