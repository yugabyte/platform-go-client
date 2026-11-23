# OciCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | [**OciCredentialType**](OciCredentialType.md) |  | 
**AccessKeyId** | Pointer to **string** | Access key of the customer secret key. Required on create; never returned. | [optional] 
**SecretAccessKey** | Pointer to **string** | Secret of the customer secret key. Required on create; never returned. | [optional] 
**HostBase** | **string** | Host of the S3-compatible endpoint | 

## Methods

### NewOciCredentials

`func NewOciCredentials(credentialType OciCredentialType, hostBase string, ) *OciCredentials`

NewOciCredentials instantiates a new OciCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciCredentialsWithDefaults

`func NewOciCredentialsWithDefaults() *OciCredentials`

NewOciCredentialsWithDefaults instantiates a new OciCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *OciCredentials) GetCredentialType() OciCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *OciCredentials) GetCredentialTypeOk() (*OciCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *OciCredentials) SetCredentialType(v OciCredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetAccessKeyId

`func (o *OciCredentials) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *OciCredentials) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *OciCredentials) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *OciCredentials) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *OciCredentials) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *OciCredentials) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *OciCredentials) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *OciCredentials) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetHostBase

`func (o *OciCredentials) GetHostBase() string`

GetHostBase returns the HostBase field if non-nil, zero value otherwise.

### GetHostBaseOk

`func (o *OciCredentials) GetHostBaseOk() (*string, bool)`

GetHostBaseOk returns a tuple with the HostBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBase

`func (o *OciCredentials) SetHostBase(v string)`

SetHostBase sets HostBase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


