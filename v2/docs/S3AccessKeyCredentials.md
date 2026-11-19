# S3AccessKeyCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | [**S3CredentialType**](S3CredentialType.md) |  | 
**AccessKeyId** | Pointer to **string** | Access key ID. Required on create; never returned. | [optional] 
**SecretAccessKey** | Pointer to **string** | Secret access key. Required on create; never returned. | [optional] 

## Methods

### NewS3AccessKeyCredentials

`func NewS3AccessKeyCredentials(credentialType S3CredentialType, ) *S3AccessKeyCredentials`

NewS3AccessKeyCredentials instantiates a new S3AccessKeyCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3AccessKeyCredentialsWithDefaults

`func NewS3AccessKeyCredentialsWithDefaults() *S3AccessKeyCredentials`

NewS3AccessKeyCredentialsWithDefaults instantiates a new S3AccessKeyCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *S3AccessKeyCredentials) GetCredentialType() S3CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *S3AccessKeyCredentials) GetCredentialTypeOk() (*S3CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *S3AccessKeyCredentials) SetCredentialType(v S3CredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetAccessKeyId

`func (o *S3AccessKeyCredentials) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *S3AccessKeyCredentials) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *S3AccessKeyCredentials) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *S3AccessKeyCredentials) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *S3AccessKeyCredentials) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *S3AccessKeyCredentials) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *S3AccessKeyCredentials) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *S3AccessKeyCredentials) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


