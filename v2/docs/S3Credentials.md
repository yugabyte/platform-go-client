# S3Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | [**S3CredentialType**](S3CredentialType.md) |  | 
**AccessKeyId** | Pointer to **string** | Access key ID. Required on create; never returned. | [optional] 
**SecretAccessKey** | Pointer to **string** | Secret access key. Required on create; never returned. | [optional] 
**CredentialSource** | Pointer to [**S3IamCredentialSource**](S3IamCredentialSource.md) |  | [optional] [default to S3IamSourceDefault]
**IamUserProfile** | Pointer to **string** | AWS profile to use with credential_source IAM_USER. Defaults to \&quot;default\&quot;. | [optional] 
**SessionDurationSecs** | Pointer to **int32** | Duration of assumed-role sessions, in seconds, when the role&#39;s maximum can&#39;t be read (no iam:GetRole permission). Defaults to 3600. | [optional] 
**RegionalSts** | Pointer to **bool** | Whether to use the regional STS endpoint rather than the global one. Defaults to true. | [optional] 
**StsRegion** | Pointer to **string** | Region of the STS endpoint, with regional_sts | [optional] 

## Methods

### NewS3Credentials

`func NewS3Credentials(credentialType S3CredentialType, ) *S3Credentials`

NewS3Credentials instantiates a new S3Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CredentialsWithDefaults

`func NewS3CredentialsWithDefaults() *S3Credentials`

NewS3CredentialsWithDefaults instantiates a new S3Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *S3Credentials) GetCredentialType() S3CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *S3Credentials) GetCredentialTypeOk() (*S3CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *S3Credentials) SetCredentialType(v S3CredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetAccessKeyId

`func (o *S3Credentials) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *S3Credentials) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *S3Credentials) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *S3Credentials) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *S3Credentials) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *S3Credentials) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *S3Credentials) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *S3Credentials) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetCredentialSource

`func (o *S3Credentials) GetCredentialSource() S3IamCredentialSource`

GetCredentialSource returns the CredentialSource field if non-nil, zero value otherwise.

### GetCredentialSourceOk

`func (o *S3Credentials) GetCredentialSourceOk() (*S3IamCredentialSource, bool)`

GetCredentialSourceOk returns a tuple with the CredentialSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSource

`func (o *S3Credentials) SetCredentialSource(v S3IamCredentialSource)`

SetCredentialSource sets CredentialSource field to given value.

### HasCredentialSource

`func (o *S3Credentials) HasCredentialSource() bool`

HasCredentialSource returns a boolean if a field has been set.

### GetIamUserProfile

`func (o *S3Credentials) GetIamUserProfile() string`

GetIamUserProfile returns the IamUserProfile field if non-nil, zero value otherwise.

### GetIamUserProfileOk

`func (o *S3Credentials) GetIamUserProfileOk() (*string, bool)`

GetIamUserProfileOk returns a tuple with the IamUserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamUserProfile

`func (o *S3Credentials) SetIamUserProfile(v string)`

SetIamUserProfile sets IamUserProfile field to given value.

### HasIamUserProfile

`func (o *S3Credentials) HasIamUserProfile() bool`

HasIamUserProfile returns a boolean if a field has been set.

### GetSessionDurationSecs

`func (o *S3Credentials) GetSessionDurationSecs() int32`

GetSessionDurationSecs returns the SessionDurationSecs field if non-nil, zero value otherwise.

### GetSessionDurationSecsOk

`func (o *S3Credentials) GetSessionDurationSecsOk() (*int32, bool)`

GetSessionDurationSecsOk returns a tuple with the SessionDurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDurationSecs

`func (o *S3Credentials) SetSessionDurationSecs(v int32)`

SetSessionDurationSecs sets SessionDurationSecs field to given value.

### HasSessionDurationSecs

`func (o *S3Credentials) HasSessionDurationSecs() bool`

HasSessionDurationSecs returns a boolean if a field has been set.

### GetRegionalSts

`func (o *S3Credentials) GetRegionalSts() bool`

GetRegionalSts returns the RegionalSts field if non-nil, zero value otherwise.

### GetRegionalStsOk

`func (o *S3Credentials) GetRegionalStsOk() (*bool, bool)`

GetRegionalStsOk returns a tuple with the RegionalSts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalSts

`func (o *S3Credentials) SetRegionalSts(v bool)`

SetRegionalSts sets RegionalSts field to given value.

### HasRegionalSts

`func (o *S3Credentials) HasRegionalSts() bool`

HasRegionalSts returns a boolean if a field has been set.

### GetStsRegion

`func (o *S3Credentials) GetStsRegion() string`

GetStsRegion returns the StsRegion field if non-nil, zero value otherwise.

### GetStsRegionOk

`func (o *S3Credentials) GetStsRegionOk() (*string, bool)`

GetStsRegionOk returns a tuple with the StsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsRegion

`func (o *S3Credentials) SetStsRegion(v string)`

SetStsRegion sets StsRegion field to given value.

### HasStsRegion

`func (o *S3Credentials) HasStsRegion() bool`

HasStsRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


