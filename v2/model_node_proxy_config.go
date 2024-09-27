/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// NodeProxyConfig The Proxy Settings for the nodes in the Universe. NodeProxyConfig is part of NodeSpec.
type NodeProxyConfig struct {
	// The HTTP_PROXY to use
	HttpProxy *string `json:"http_proxy,omitempty"`
	// The HTTPS_PROXY to use
	HttpsProxy *string `json:"https_proxy,omitempty"`
	// The NO_PROXY settings. Should follow cURL no_proxy format
	NoProxyList *[]string `json:"no_proxy_list,omitempty"`
}

// NewNodeProxyConfig instantiates a new NodeProxyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeProxyConfig() *NodeProxyConfig {
	this := NodeProxyConfig{}
	return &this
}

// NewNodeProxyConfigWithDefaults instantiates a new NodeProxyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeProxyConfigWithDefaults() *NodeProxyConfig {
	this := NodeProxyConfig{}
	return &this
}

// GetHttpProxy returns the HttpProxy field value if set, zero value otherwise.
func (o *NodeProxyConfig) GetHttpProxy() string {
	if o == nil || o.HttpProxy == nil {
		var ret string
		return ret
	}
	return *o.HttpProxy
}

// GetHttpProxyOk returns a tuple with the HttpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeProxyConfig) GetHttpProxyOk() (*string, bool) {
	if o == nil || o.HttpProxy == nil {
		return nil, false
	}
	return o.HttpProxy, true
}

// HasHttpProxy returns a boolean if a field has been set.
func (o *NodeProxyConfig) HasHttpProxy() bool {
	if o != nil && o.HttpProxy != nil {
		return true
	}

	return false
}

// SetHttpProxy gets a reference to the given string and assigns it to the HttpProxy field.
func (o *NodeProxyConfig) SetHttpProxy(v string) {
	o.HttpProxy = &v
}

// GetHttpsProxy returns the HttpsProxy field value if set, zero value otherwise.
func (o *NodeProxyConfig) GetHttpsProxy() string {
	if o == nil || o.HttpsProxy == nil {
		var ret string
		return ret
	}
	return *o.HttpsProxy
}

// GetHttpsProxyOk returns a tuple with the HttpsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeProxyConfig) GetHttpsProxyOk() (*string, bool) {
	if o == nil || o.HttpsProxy == nil {
		return nil, false
	}
	return o.HttpsProxy, true
}

// HasHttpsProxy returns a boolean if a field has been set.
func (o *NodeProxyConfig) HasHttpsProxy() bool {
	if o != nil && o.HttpsProxy != nil {
		return true
	}

	return false
}

// SetHttpsProxy gets a reference to the given string and assigns it to the HttpsProxy field.
func (o *NodeProxyConfig) SetHttpsProxy(v string) {
	o.HttpsProxy = &v
}

// GetNoProxyList returns the NoProxyList field value if set, zero value otherwise.
func (o *NodeProxyConfig) GetNoProxyList() []string {
	if o == nil || o.NoProxyList == nil {
		var ret []string
		return ret
	}
	return *o.NoProxyList
}

// GetNoProxyListOk returns a tuple with the NoProxyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeProxyConfig) GetNoProxyListOk() (*[]string, bool) {
	if o == nil || o.NoProxyList == nil {
		return nil, false
	}
	return o.NoProxyList, true
}

// HasNoProxyList returns a boolean if a field has been set.
func (o *NodeProxyConfig) HasNoProxyList() bool {
	if o != nil && o.NoProxyList != nil {
		return true
	}

	return false
}

// SetNoProxyList gets a reference to the given []string and assigns it to the NoProxyList field.
func (o *NodeProxyConfig) SetNoProxyList(v []string) {
	o.NoProxyList = &v
}

func (o NodeProxyConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpProxy != nil {
		toSerialize["http_proxy"] = o.HttpProxy
	}
	if o.HttpsProxy != nil {
		toSerialize["https_proxy"] = o.HttpsProxy
	}
	if o.NoProxyList != nil {
		toSerialize["no_proxy_list"] = o.NoProxyList
	}
	return json.Marshal(toSerialize)
}

type NullableNodeProxyConfig struct {
	value *NodeProxyConfig
	isSet bool
}

func (v NullableNodeProxyConfig) Get() *NodeProxyConfig {
	return v.value
}

func (v *NullableNodeProxyConfig) Set(val *NodeProxyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeProxyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeProxyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeProxyConfig(val *NodeProxyConfig) *NullableNodeProxyConfig {
	return &NullableNodeProxyConfig{value: val, isSet: true}
}

func (v NullableNodeProxyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeProxyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


