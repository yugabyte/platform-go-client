# EncryptionAtRestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionAtRestEnabled** | Pointer to **bool** | Whether a universe is currently encrypted at rest | [optional] 
**KmsConfigUUID** | Pointer to **string** | KMS configuration UUID | [optional] 
**OpType** | Pointer to **string** | Operation type: enable, disable, or rotate the universe key/encryption at rest | [optional] 
**Type** | Pointer to **string** | Whether to generate a data key or just retrieve the CMK ARN | [optional] 

## Methods

### NewEncryptionAtRestConfig

`func NewEncryptionAtRestConfig() *EncryptionAtRestConfig`

NewEncryptionAtRestConfig instantiates a new EncryptionAtRestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionAtRestConfigWithDefaults

`func NewEncryptionAtRestConfigWithDefaults() *EncryptionAtRestConfig`

NewEncryptionAtRestConfigWithDefaults instantiates a new EncryptionAtRestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionAtRestEnabled

`func (o *EncryptionAtRestConfig) GetEncryptionAtRestEnabled() bool`

GetEncryptionAtRestEnabled returns the EncryptionAtRestEnabled field if non-nil, zero value otherwise.

### GetEncryptionAtRestEnabledOk

`func (o *EncryptionAtRestConfig) GetEncryptionAtRestEnabledOk() (*bool, bool)`

GetEncryptionAtRestEnabledOk returns a tuple with the EncryptionAtRestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestEnabled

`func (o *EncryptionAtRestConfig) SetEncryptionAtRestEnabled(v bool)`

SetEncryptionAtRestEnabled sets EncryptionAtRestEnabled field to given value.

### HasEncryptionAtRestEnabled

`func (o *EncryptionAtRestConfig) HasEncryptionAtRestEnabled() bool`

HasEncryptionAtRestEnabled returns a boolean if a field has been set.

### GetKmsConfigUUID

`func (o *EncryptionAtRestConfig) GetKmsConfigUUID() string`

GetKmsConfigUUID returns the KmsConfigUUID field if non-nil, zero value otherwise.

### GetKmsConfigUUIDOk

`func (o *EncryptionAtRestConfig) GetKmsConfigUUIDOk() (*string, bool)`

GetKmsConfigUUIDOk returns a tuple with the KmsConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUUID

`func (o *EncryptionAtRestConfig) SetKmsConfigUUID(v string)`

SetKmsConfigUUID sets KmsConfigUUID field to given value.

### HasKmsConfigUUID

`func (o *EncryptionAtRestConfig) HasKmsConfigUUID() bool`

HasKmsConfigUUID returns a boolean if a field has been set.

### GetOpType

`func (o *EncryptionAtRestConfig) GetOpType() string`

GetOpType returns the OpType field if non-nil, zero value otherwise.

### GetOpTypeOk

`func (o *EncryptionAtRestConfig) GetOpTypeOk() (*string, bool)`

GetOpTypeOk returns a tuple with the OpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpType

`func (o *EncryptionAtRestConfig) SetOpType(v string)`

SetOpType sets OpType field to given value.

### HasOpType

`func (o *EncryptionAtRestConfig) HasOpType() bool`

HasOpType returns a boolean if a field has been set.

### GetType

`func (o *EncryptionAtRestConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EncryptionAtRestConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EncryptionAtRestConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EncryptionAtRestConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


