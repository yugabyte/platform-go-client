# CloudVolumeEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableVolumeEncryption** | Pointer to **bool** | To enable volume encryption or not | [optional] [default to false]
**KmsConfigUuid** | Pointer to **string** | KMS config UUID to be used for volume encryption | [optional] 

## Methods

### NewCloudVolumeEncryption

`func NewCloudVolumeEncryption() *CloudVolumeEncryption`

NewCloudVolumeEncryption instantiates a new CloudVolumeEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeEncryptionWithDefaults

`func NewCloudVolumeEncryptionWithDefaults() *CloudVolumeEncryption`

NewCloudVolumeEncryptionWithDefaults instantiates a new CloudVolumeEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableVolumeEncryption

`func (o *CloudVolumeEncryption) GetEnableVolumeEncryption() bool`

GetEnableVolumeEncryption returns the EnableVolumeEncryption field if non-nil, zero value otherwise.

### GetEnableVolumeEncryptionOk

`func (o *CloudVolumeEncryption) GetEnableVolumeEncryptionOk() (*bool, bool)`

GetEnableVolumeEncryptionOk returns a tuple with the EnableVolumeEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVolumeEncryption

`func (o *CloudVolumeEncryption) SetEnableVolumeEncryption(v bool)`

SetEnableVolumeEncryption sets EnableVolumeEncryption field to given value.

### HasEnableVolumeEncryption

`func (o *CloudVolumeEncryption) HasEnableVolumeEncryption() bool`

HasEnableVolumeEncryption returns a boolean if a field has been set.

### GetKmsConfigUuid

`func (o *CloudVolumeEncryption) GetKmsConfigUuid() string`

GetKmsConfigUuid returns the KmsConfigUuid field if non-nil, zero value otherwise.

### GetKmsConfigUuidOk

`func (o *CloudVolumeEncryption) GetKmsConfigUuidOk() (*string, bool)`

GetKmsConfigUuidOk returns a tuple with the KmsConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsConfigUuid

`func (o *CloudVolumeEncryption) SetKmsConfigUuid(v string)`

SetKmsConfigUuid sets KmsConfigUuid field to given value.

### HasKmsConfigUuid

`func (o *CloudVolumeEncryption) HasKmsConfigUuid() bool`

HasKmsConfigUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


