# EncryptionAtRestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KmsConfigUuid** | Pointer to **string** | The KMS Configuration associated with the encryption keys being used on this Universe | [optional] 

## Methods

### NewEncryptionAtRestSpec

`func NewEncryptionAtRestSpec() *EncryptionAtRestSpec`

NewEncryptionAtRestSpec instantiates a new EncryptionAtRestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionAtRestSpecWithDefaults

`func NewEncryptionAtRestSpecWithDefaults() *EncryptionAtRestSpec`

NewEncryptionAtRestSpecWithDefaults instantiates a new EncryptionAtRestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKmsConfigUuid

`func (o *EncryptionAtRestSpec) GetKmsConfigUuid() string`

GetKmsConfigUuid returns the KmsConfigUuid field if non-nil, zero value otherwise.

### GetKmsConfigUuidOk

`func (o *EncryptionAtRestSpec) GetKmsConfigUuidOk() (*string, bool)`

GetKmsConfigUuidOk returns a tuple with the KmsConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUuid

`func (o *EncryptionAtRestSpec) SetKmsConfigUuid(v string)`

SetKmsConfigUuid sets KmsConfigUuid field to given value.

### HasKmsConfigUuid

`func (o *EncryptionAtRestSpec) HasKmsConfigUuid() bool`

HasKmsConfigUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


