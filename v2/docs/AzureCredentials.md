# AzureCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | [**AzureCredentialType**](AzureCredentialType.md) |  | 
**SasToken** | Pointer to **string** | SAS token of the backup_location container. Required on create; never returned. | [optional] 
**ClientId** | Pointer to **string** | Client ID of the user-assigned managed identity to use. Omit to use the system-assigned identity. | [optional] 

## Methods

### NewAzureCredentials

`func NewAzureCredentials(credentialType AzureCredentialType, ) *AzureCredentials`

NewAzureCredentials instantiates a new AzureCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCredentialsWithDefaults

`func NewAzureCredentialsWithDefaults() *AzureCredentials`

NewAzureCredentialsWithDefaults instantiates a new AzureCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *AzureCredentials) GetCredentialType() AzureCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *AzureCredentials) GetCredentialTypeOk() (*AzureCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *AzureCredentials) SetCredentialType(v AzureCredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetSasToken

`func (o *AzureCredentials) GetSasToken() string`

GetSasToken returns the SasToken field if non-nil, zero value otherwise.

### GetSasTokenOk

`func (o *AzureCredentials) GetSasTokenOk() (*string, bool)`

GetSasTokenOk returns a tuple with the SasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasToken

`func (o *AzureCredentials) SetSasToken(v string)`

SetSasToken sets SasToken field to given value.

### HasSasToken

`func (o *AzureCredentials) HasSasToken() bool`

HasSasToken returns a boolean if a field has been set.

### GetClientId

`func (o *AzureCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AzureCredentials) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


