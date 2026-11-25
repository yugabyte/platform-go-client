# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertPath** | **string** |  | 
**Compressor** | Pointer to **string** |  | [optional] 
**ServerCert** | Pointer to **string** |  | [optional] 
**ServerCertExpirySecs** | **int64** |  | 
**ServerCertLocalPath** | Pointer to **string** |  | [optional] 
**ServerKey** | Pointer to **string** |  | [optional] 
**ServerKeyLocalPath** | Pointer to **string** |  | [optional] 
**SignerPrivateKey** | Pointer to **string** |  | [optional] 
**SignerPublicKey** | Pointer to **string** |  | [optional] 

## Methods

### NewConfig

`func NewConfig(certPath string, serverCertExpirySecs int64, ) *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertPath

`func (o *Config) GetCertPath() string`

GetCertPath returns the CertPath field if non-nil, zero value otherwise.

### GetCertPathOk

`func (o *Config) GetCertPathOk() (*string, bool)`

GetCertPathOk returns a tuple with the CertPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPath

`func (o *Config) SetCertPath(v string)`

SetCertPath sets CertPath field to given value.


### GetCompressor

`func (o *Config) GetCompressor() string`

GetCompressor returns the Compressor field if non-nil, zero value otherwise.

### GetCompressorOk

`func (o *Config) GetCompressorOk() (*string, bool)`

GetCompressorOk returns a tuple with the Compressor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressor

`func (o *Config) SetCompressor(v string)`

SetCompressor sets Compressor field to given value.

### HasCompressor

`func (o *Config) HasCompressor() bool`

HasCompressor returns a boolean if a field has been set.

### GetServerCert

`func (o *Config) GetServerCert() string`

GetServerCert returns the ServerCert field if non-nil, zero value otherwise.

### GetServerCertOk

`func (o *Config) GetServerCertOk() (*string, bool)`

GetServerCertOk returns a tuple with the ServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCert

`func (o *Config) SetServerCert(v string)`

SetServerCert sets ServerCert field to given value.

### HasServerCert

`func (o *Config) HasServerCert() bool`

HasServerCert returns a boolean if a field has been set.

### GetServerCertExpirySecs

`func (o *Config) GetServerCertExpirySecs() int64`

GetServerCertExpirySecs returns the ServerCertExpirySecs field if non-nil, zero value otherwise.

### GetServerCertExpirySecsOk

`func (o *Config) GetServerCertExpirySecsOk() (*int64, bool)`

GetServerCertExpirySecsOk returns a tuple with the ServerCertExpirySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertExpirySecs

`func (o *Config) SetServerCertExpirySecs(v int64)`

SetServerCertExpirySecs sets ServerCertExpirySecs field to given value.


### GetServerCertLocalPath

`func (o *Config) GetServerCertLocalPath() string`

GetServerCertLocalPath returns the ServerCertLocalPath field if non-nil, zero value otherwise.

### GetServerCertLocalPathOk

`func (o *Config) GetServerCertLocalPathOk() (*string, bool)`

GetServerCertLocalPathOk returns a tuple with the ServerCertLocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertLocalPath

`func (o *Config) SetServerCertLocalPath(v string)`

SetServerCertLocalPath sets ServerCertLocalPath field to given value.

### HasServerCertLocalPath

`func (o *Config) HasServerCertLocalPath() bool`

HasServerCertLocalPath returns a boolean if a field has been set.

### GetServerKey

`func (o *Config) GetServerKey() string`

GetServerKey returns the ServerKey field if non-nil, zero value otherwise.

### GetServerKeyOk

`func (o *Config) GetServerKeyOk() (*string, bool)`

GetServerKeyOk returns a tuple with the ServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKey

`func (o *Config) SetServerKey(v string)`

SetServerKey sets ServerKey field to given value.

### HasServerKey

`func (o *Config) HasServerKey() bool`

HasServerKey returns a boolean if a field has been set.

### GetServerKeyLocalPath

`func (o *Config) GetServerKeyLocalPath() string`

GetServerKeyLocalPath returns the ServerKeyLocalPath field if non-nil, zero value otherwise.

### GetServerKeyLocalPathOk

`func (o *Config) GetServerKeyLocalPathOk() (*string, bool)`

GetServerKeyLocalPathOk returns a tuple with the ServerKeyLocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeyLocalPath

`func (o *Config) SetServerKeyLocalPath(v string)`

SetServerKeyLocalPath sets ServerKeyLocalPath field to given value.

### HasServerKeyLocalPath

`func (o *Config) HasServerKeyLocalPath() bool`

HasServerKeyLocalPath returns a boolean if a field has been set.

### GetSignerPrivateKey

`func (o *Config) GetSignerPrivateKey() string`

GetSignerPrivateKey returns the SignerPrivateKey field if non-nil, zero value otherwise.

### GetSignerPrivateKeyOk

`func (o *Config) GetSignerPrivateKeyOk() (*string, bool)`

GetSignerPrivateKeyOk returns a tuple with the SignerPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerPrivateKey

`func (o *Config) SetSignerPrivateKey(v string)`

SetSignerPrivateKey sets SignerPrivateKey field to given value.

### HasSignerPrivateKey

`func (o *Config) HasSignerPrivateKey() bool`

HasSignerPrivateKey returns a boolean if a field has been set.

### GetSignerPublicKey

`func (o *Config) GetSignerPublicKey() string`

GetSignerPublicKey returns the SignerPublicKey field if non-nil, zero value otherwise.

### GetSignerPublicKeyOk

`func (o *Config) GetSignerPublicKeyOk() (*string, bool)`

GetSignerPublicKeyOk returns a tuple with the SignerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerPublicKey

`func (o *Config) SetSignerPublicKey(v string)`

SetSignerPublicKey sets SignerPublicKey field to given value.

### HasSignerPublicKey

`func (o *Config) HasSignerPublicKey() bool`

HasSignerPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


