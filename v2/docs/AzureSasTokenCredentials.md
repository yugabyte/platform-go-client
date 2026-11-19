# AzureSasTokenCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | [**AzureCredentialType**](AzureCredentialType.md) |  | 
**SasToken** | Pointer to **string** | SAS token of the backup_location container. Required on create; never returned. | [optional] 

## Methods

### NewAzureSasTokenCredentials

`func NewAzureSasTokenCredentials(credentialType AzureCredentialType, ) *AzureSasTokenCredentials`

NewAzureSasTokenCredentials instantiates a new AzureSasTokenCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSasTokenCredentialsWithDefaults

`func NewAzureSasTokenCredentialsWithDefaults() *AzureSasTokenCredentials`

NewAzureSasTokenCredentialsWithDefaults instantiates a new AzureSasTokenCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *AzureSasTokenCredentials) GetCredentialType() AzureCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *AzureSasTokenCredentials) GetCredentialTypeOk() (*AzureCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *AzureSasTokenCredentials) SetCredentialType(v AzureCredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetSasToken

`func (o *AzureSasTokenCredentials) GetSasToken() string`

GetSasToken returns the SasToken field if non-nil, zero value otherwise.

### GetSasTokenOk

`func (o *AzureSasTokenCredentials) GetSasTokenOk() (*string, bool)`

GetSasTokenOk returns a tuple with the SasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasToken

`func (o *AzureSasTokenCredentials) SetSasToken(v string)`

SetSasToken sets SasToken field to given value.

### HasSasToken

`func (o *AzureSasTokenCredentials) HasSasToken() bool`

HasSasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


