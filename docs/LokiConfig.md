# LokiConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** | Auth Type | 
**BasicAuth** | Pointer to [**BasicAuthCredentials**](BasicAuthCredentials.md) |  | [optional] 
**Endpoint** | **string** | End Point | 
**OrganizationID** | Pointer to **string** | Organization/Tenant ID | [optional] 

## Methods

### NewLokiConfig

`func NewLokiConfig(authType string, endpoint string, ) *LokiConfig`

NewLokiConfig instantiates a new LokiConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLokiConfigWithDefaults

`func NewLokiConfigWithDefaults() *LokiConfig`

NewLokiConfigWithDefaults instantiates a new LokiConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *LokiConfig) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *LokiConfig) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *LokiConfig) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetBasicAuth

`func (o *LokiConfig) GetBasicAuth() BasicAuthCredentials`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *LokiConfig) GetBasicAuthOk() (*BasicAuthCredentials, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *LokiConfig) SetBasicAuth(v BasicAuthCredentials)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *LokiConfig) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetEndpoint

`func (o *LokiConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *LokiConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *LokiConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetOrganizationID

`func (o *LokiConfig) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *LokiConfig) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *LokiConfig) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *LokiConfig) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


