# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertPath** | **string** |  | 
**Compressor** | **string** |  | 
**Offloadable** | **bool** |  | 
**ServerCert** | **string** |  | 
**ServerKey** | **string** |  | 

## Methods

### NewConfig

`func NewConfig(certPath string, compressor string, offloadable bool, serverCert string, serverKey string, ) *Config`

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


### GetOffloadable

`func (o *Config) GetOffloadable() bool`

GetOffloadable returns the Offloadable field if non-nil, zero value otherwise.

### GetOffloadableOk

`func (o *Config) GetOffloadableOk() (*bool, bool)`

GetOffloadableOk returns a tuple with the Offloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloadable

`func (o *Config) SetOffloadable(v bool)`

SetOffloadable sets Offloadable field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


