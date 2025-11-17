# OTLPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** | Auth Type | 
**BasicAuth** | Pointer to [**BasicAuthCredentials**](BasicAuthCredentials.md) |  | [optional] 
**BearerToken** | Pointer to [**BearerToken**](BearerToken.md) |  | [optional] 
**Compression** | Pointer to **string** | Compression | [optional] 
**Endpoint** | **string** | End Point. For HTTP protcol logs export \&quot;/v1/logs\&quot; will be appended. | 
**Headers** | Pointer to **map[string]string** | Headers | [optional] 
**LogsEndpoint** | Pointer to **string** | Logs endpoint. The target URL to send log data to (e.g.: https://example.com:4318/v1/logs). If this setting is present the endpoint setting is ignored logs. Allowed only for HTTP protocol | [optional] 
**Protocol** | Pointer to **string** | Protocol | [optional] 
**TimeoutSeconds** | Pointer to **int32** | Timeout in seconds | [optional] 

## Methods

### NewOTLPConfig

`func NewOTLPConfig(authType string, endpoint string, ) *OTLPConfig`

NewOTLPConfig instantiates a new OTLPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTLPConfigWithDefaults

`func NewOTLPConfigWithDefaults() *OTLPConfig`

NewOTLPConfigWithDefaults instantiates a new OTLPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *OTLPConfig) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *OTLPConfig) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *OTLPConfig) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetBasicAuth

`func (o *OTLPConfig) GetBasicAuth() BasicAuthCredentials`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *OTLPConfig) GetBasicAuthOk() (*BasicAuthCredentials, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *OTLPConfig) SetBasicAuth(v BasicAuthCredentials)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *OTLPConfig) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBearerToken

`func (o *OTLPConfig) GetBearerToken() BearerToken`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *OTLPConfig) GetBearerTokenOk() (*BearerToken, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *OTLPConfig) SetBearerToken(v BearerToken)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *OTLPConfig) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetCompression

`func (o *OTLPConfig) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *OTLPConfig) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *OTLPConfig) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *OTLPConfig) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetEndpoint

`func (o *OTLPConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *OTLPConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *OTLPConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetHeaders

`func (o *OTLPConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OTLPConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OTLPConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *OTLPConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetLogsEndpoint

`func (o *OTLPConfig) GetLogsEndpoint() string`

GetLogsEndpoint returns the LogsEndpoint field if non-nil, zero value otherwise.

### GetLogsEndpointOk

`func (o *OTLPConfig) GetLogsEndpointOk() (*string, bool)`

GetLogsEndpointOk returns a tuple with the LogsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsEndpoint

`func (o *OTLPConfig) SetLogsEndpoint(v string)`

SetLogsEndpoint sets LogsEndpoint field to given value.

### HasLogsEndpoint

`func (o *OTLPConfig) HasLogsEndpoint() bool`

HasLogsEndpoint returns a boolean if a field has been set.

### GetProtocol

`func (o *OTLPConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OTLPConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OTLPConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *OTLPConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *OTLPConfig) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *OTLPConfig) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *OTLPConfig) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *OTLPConfig) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


