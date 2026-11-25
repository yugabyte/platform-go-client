# OciS3CompatibleCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | [**OciCredentialType**](OciCredentialType.md) |  | 
**AccessKeyId** | Pointer to **string** | Access key of the customer secret key. Required on create; on edit, omit it to keep the stored one. Never returned. | [optional] 
**SecretAccessKey** | Pointer to **string** | Secret of the customer secret key. Required on create; on edit, omit it to keep the stored one. Never returned. | [optional] 
**HostBase** | **string** | Host of the S3-compatible endpoint | 

## Methods

### NewOciS3CompatibleCredentials

`func NewOciS3CompatibleCredentials(credentialType OciCredentialType, hostBase string, ) *OciS3CompatibleCredentials`

NewOciS3CompatibleCredentials instantiates a new OciS3CompatibleCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciS3CompatibleCredentialsWithDefaults

`func NewOciS3CompatibleCredentialsWithDefaults() *OciS3CompatibleCredentials`

NewOciS3CompatibleCredentialsWithDefaults instantiates a new OciS3CompatibleCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *OciS3CompatibleCredentials) GetCredentialType() OciCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *OciS3CompatibleCredentials) GetCredentialTypeOk() (*OciCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *OciS3CompatibleCredentials) SetCredentialType(v OciCredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetAccessKeyId

`func (o *OciS3CompatibleCredentials) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *OciS3CompatibleCredentials) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *OciS3CompatibleCredentials) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *OciS3CompatibleCredentials) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *OciS3CompatibleCredentials) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *OciS3CompatibleCredentials) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *OciS3CompatibleCredentials) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *OciS3CompatibleCredentials) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetHostBase

`func (o *OciS3CompatibleCredentials) GetHostBase() string`

GetHostBase returns the HostBase field if non-nil, zero value otherwise.

### GetHostBaseOk

`func (o *OciS3CompatibleCredentials) GetHostBaseOk() (*string, bool)`

GetHostBaseOk returns a tuple with the HostBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBase

`func (o *OciS3CompatibleCredentials) SetHostBase(v string)`

SetHostBase sets HostBase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


