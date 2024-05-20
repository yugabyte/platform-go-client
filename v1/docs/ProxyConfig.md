# ProxyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpProxy** | Pointer to **string** | The HTTP_PROXY to use | [optional] 
**HttpsProxy** | Pointer to **string** | The HTTPS_PROXY to use | [optional] 
**NoProxyList** | Pointer to **[]string** | The NO_PROXY settings. Should follow cURL no_proxy format | [optional] 

## Methods

### NewProxyConfig

`func NewProxyConfig() *ProxyConfig`

NewProxyConfig instantiates a new ProxyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyConfigWithDefaults

`func NewProxyConfigWithDefaults() *ProxyConfig`

NewProxyConfigWithDefaults instantiates a new ProxyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpProxy

`func (o *ProxyConfig) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *ProxyConfig) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *ProxyConfig) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *ProxyConfig) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### GetHttpsProxy

`func (o *ProxyConfig) GetHttpsProxy() string`

GetHttpsProxy returns the HttpsProxy field if non-nil, zero value otherwise.

### GetHttpsProxyOk

`func (o *ProxyConfig) GetHttpsProxyOk() (*string, bool)`

GetHttpsProxyOk returns a tuple with the HttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxy

`func (o *ProxyConfig) SetHttpsProxy(v string)`

SetHttpsProxy sets HttpsProxy field to given value.

### HasHttpsProxy

`func (o *ProxyConfig) HasHttpsProxy() bool`

HasHttpsProxy returns a boolean if a field has been set.

### GetNoProxyList

`func (o *ProxyConfig) GetNoProxyList() []string`

GetNoProxyList returns the NoProxyList field if non-nil, zero value otherwise.

### GetNoProxyListOk

`func (o *ProxyConfig) GetNoProxyListOk() (*[]string, bool)`

GetNoProxyListOk returns a tuple with the NoProxyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxyList

`func (o *ProxyConfig) SetNoProxyList(v []string)`

SetNoProxyList sets NoProxyList field to given value.

### HasNoProxyList

`func (o *ProxyConfig) HasNoProxyList() bool`

HasNoProxyList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


