# NodeProxyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpProxy** | Pointer to **string** | The HTTP_PROXY to use | [optional] 
**HttpsProxy** | Pointer to **string** | The HTTPS_PROXY to use | [optional] 
**NoProxyList** | Pointer to **[]string** | The NO_PROXY settings. Should follow cURL no_proxy format | [optional] 

## Methods

### NewNodeProxyConfig

`func NewNodeProxyConfig() *NodeProxyConfig`

NewNodeProxyConfig instantiates a new NodeProxyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeProxyConfigWithDefaults

`func NewNodeProxyConfigWithDefaults() *NodeProxyConfig`

NewNodeProxyConfigWithDefaults instantiates a new NodeProxyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpProxy

`func (o *NodeProxyConfig) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *NodeProxyConfig) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *NodeProxyConfig) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *NodeProxyConfig) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### GetHttpsProxy

`func (o *NodeProxyConfig) GetHttpsProxy() string`

GetHttpsProxy returns the HttpsProxy field if non-nil, zero value otherwise.

### GetHttpsProxyOk

`func (o *NodeProxyConfig) GetHttpsProxyOk() (*string, bool)`

GetHttpsProxyOk returns a tuple with the HttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxy

`func (o *NodeProxyConfig) SetHttpsProxy(v string)`

SetHttpsProxy sets HttpsProxy field to given value.

### HasHttpsProxy

`func (o *NodeProxyConfig) HasHttpsProxy() bool`

HasHttpsProxy returns a boolean if a field has been set.

### GetNoProxyList

`func (o *NodeProxyConfig) GetNoProxyList() []string`

GetNoProxyList returns the NoProxyList field if non-nil, zero value otherwise.

### GetNoProxyListOk

`func (o *NodeProxyConfig) GetNoProxyListOk() (*[]string, bool)`

GetNoProxyListOk returns a tuple with the NoProxyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxyList

`func (o *NodeProxyConfig) SetNoProxyList(v []string)`

SetNoProxyList sets NoProxyList field to given value.

### HasNoProxyList

`func (o *NodeProxyConfig) HasNoProxyList() bool`

HasNoProxyList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


