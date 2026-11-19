# S3IamCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | [**S3CredentialType**](S3CredentialType.md) |  | 
**CredentialSource** | Pointer to [**S3IamCredentialSource**](S3IamCredentialSource.md) |  | [optional] [default to S3IamSourceDefault]
**IamUserProfile** | Pointer to **string** | AWS profile to use with credential_source IAM_USER. Defaults to \&quot;default\&quot;. | [optional] 
**SessionDurationSecs** | Pointer to **int32** | Duration of assumed-role sessions, in seconds, when the role&#39;s maximum can&#39;t be read (no iam:GetRole permission). Defaults to 3600. | [optional] 
**RegionalSts** | Pointer to **bool** | Whether to use the regional STS endpoint rather than the global one. Defaults to true. | [optional] 
**StsRegion** | Pointer to **string** | Region of the STS endpoint, with regional_sts | [optional] 

## Methods

### NewS3IamCredentials

`func NewS3IamCredentials(credentialType S3CredentialType, ) *S3IamCredentials`

NewS3IamCredentials instantiates a new S3IamCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3IamCredentialsWithDefaults

`func NewS3IamCredentialsWithDefaults() *S3IamCredentials`

NewS3IamCredentialsWithDefaults instantiates a new S3IamCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *S3IamCredentials) GetCredentialType() S3CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *S3IamCredentials) GetCredentialTypeOk() (*S3CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *S3IamCredentials) SetCredentialType(v S3CredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetCredentialSource

`func (o *S3IamCredentials) GetCredentialSource() S3IamCredentialSource`

GetCredentialSource returns the CredentialSource field if non-nil, zero value otherwise.

### GetCredentialSourceOk

`func (o *S3IamCredentials) GetCredentialSourceOk() (*S3IamCredentialSource, bool)`

GetCredentialSourceOk returns a tuple with the CredentialSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSource

`func (o *S3IamCredentials) SetCredentialSource(v S3IamCredentialSource)`

SetCredentialSource sets CredentialSource field to given value.

### HasCredentialSource

`func (o *S3IamCredentials) HasCredentialSource() bool`

HasCredentialSource returns a boolean if a field has been set.

### GetIamUserProfile

`func (o *S3IamCredentials) GetIamUserProfile() string`

GetIamUserProfile returns the IamUserProfile field if non-nil, zero value otherwise.

### GetIamUserProfileOk

`func (o *S3IamCredentials) GetIamUserProfileOk() (*string, bool)`

GetIamUserProfileOk returns a tuple with the IamUserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamUserProfile

`func (o *S3IamCredentials) SetIamUserProfile(v string)`

SetIamUserProfile sets IamUserProfile field to given value.

### HasIamUserProfile

`func (o *S3IamCredentials) HasIamUserProfile() bool`

HasIamUserProfile returns a boolean if a field has been set.

### GetSessionDurationSecs

`func (o *S3IamCredentials) GetSessionDurationSecs() int32`

GetSessionDurationSecs returns the SessionDurationSecs field if non-nil, zero value otherwise.

### GetSessionDurationSecsOk

`func (o *S3IamCredentials) GetSessionDurationSecsOk() (*int32, bool)`

GetSessionDurationSecsOk returns a tuple with the SessionDurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDurationSecs

`func (o *S3IamCredentials) SetSessionDurationSecs(v int32)`

SetSessionDurationSecs sets SessionDurationSecs field to given value.

### HasSessionDurationSecs

`func (o *S3IamCredentials) HasSessionDurationSecs() bool`

HasSessionDurationSecs returns a boolean if a field has been set.

### GetRegionalSts

`func (o *S3IamCredentials) GetRegionalSts() bool`

GetRegionalSts returns the RegionalSts field if non-nil, zero value otherwise.

### GetRegionalStsOk

`func (o *S3IamCredentials) GetRegionalStsOk() (*bool, bool)`

GetRegionalStsOk returns a tuple with the RegionalSts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalSts

`func (o *S3IamCredentials) SetRegionalSts(v bool)`

SetRegionalSts sets RegionalSts field to given value.

### HasRegionalSts

`func (o *S3IamCredentials) HasRegionalSts() bool`

HasRegionalSts returns a boolean if a field has been set.

### GetStsRegion

`func (o *S3IamCredentials) GetStsRegion() string`

GetStsRegion returns the StsRegion field if non-nil, zero value otherwise.

### GetStsRegionOk

`func (o *S3IamCredentials) GetStsRegionOk() (*string, bool)`

GetStsRegionOk returns a tuple with the StsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsRegion

`func (o *S3IamCredentials) SetStsRegion(v string)`

SetStsRegion sets StsRegion field to given value.

### HasStsRegion

`func (o *S3IamCredentials) HasStsRegion() bool`

HasStsRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


