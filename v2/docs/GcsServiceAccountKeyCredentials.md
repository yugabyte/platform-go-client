# GcsServiceAccountKeyCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | [**GcsCredentialType**](GcsCredentialType.md) |  | 
**CredentialsJson** | Pointer to **string** | The service account key file, as JSON. Required on create; on edit, omit it to keep the stored one. Never returned. | [optional] 

## Methods

### NewGcsServiceAccountKeyCredentials

`func NewGcsServiceAccountKeyCredentials(credentialType GcsCredentialType, ) *GcsServiceAccountKeyCredentials`

NewGcsServiceAccountKeyCredentials instantiates a new GcsServiceAccountKeyCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcsServiceAccountKeyCredentialsWithDefaults

`func NewGcsServiceAccountKeyCredentialsWithDefaults() *GcsServiceAccountKeyCredentials`

NewGcsServiceAccountKeyCredentialsWithDefaults instantiates a new GcsServiceAccountKeyCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *GcsServiceAccountKeyCredentials) GetCredentialType() GcsCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *GcsServiceAccountKeyCredentials) GetCredentialTypeOk() (*GcsCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *GcsServiceAccountKeyCredentials) SetCredentialType(v GcsCredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetCredentialsJson

`func (o *GcsServiceAccountKeyCredentials) GetCredentialsJson() string`

GetCredentialsJson returns the CredentialsJson field if non-nil, zero value otherwise.

### GetCredentialsJsonOk

`func (o *GcsServiceAccountKeyCredentials) GetCredentialsJsonOk() (*string, bool)`

GetCredentialsJsonOk returns a tuple with the CredentialsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsJson

`func (o *GcsServiceAccountKeyCredentials) SetCredentialsJson(v string)`

SetCredentialsJson sets CredentialsJson field to given value.

### HasCredentialsJson

`func (o *GcsServiceAccountKeyCredentials) HasCredentialsJson() bool`

HasCredentialsJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


