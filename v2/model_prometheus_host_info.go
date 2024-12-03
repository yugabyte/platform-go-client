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

// PrometheusHostInfo PrometheusHostInfo  Information about the the Prometheus instance for the current YBA. Part of PrometheusHostInfoResp. 
type PrometheusHostInfo struct {
	// URL for accessing Prometheus
	PrometheusUrl string `json:"prometheus_url"`
	// If the Prometheus link in the browser should use the browser's FQDN or use prometheus url directly.
	UseBrowserFqdn bool `json:"use_browser_fqdn"`
}

// NewPrometheusHostInfo instantiates a new PrometheusHostInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusHostInfo(prometheusUrl string, useBrowserFqdn bool) *PrometheusHostInfo {
	this := PrometheusHostInfo{}
	this.PrometheusUrl = prometheusUrl
	this.UseBrowserFqdn = useBrowserFqdn
	return &this
}

// NewPrometheusHostInfoWithDefaults instantiates a new PrometheusHostInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusHostInfoWithDefaults() *PrometheusHostInfo {
	this := PrometheusHostInfo{}
	return &this
}

// GetPrometheusUrl returns the PrometheusUrl field value
func (o *PrometheusHostInfo) GetPrometheusUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrometheusUrl
}

// GetPrometheusUrlOk returns a tuple with the PrometheusUrl field value
// and a boolean to check if the value has been set.
func (o *PrometheusHostInfo) GetPrometheusUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrometheusUrl, true
}

// SetPrometheusUrl sets field value
func (o *PrometheusHostInfo) SetPrometheusUrl(v string) {
	o.PrometheusUrl = v
}

// GetUseBrowserFqdn returns the UseBrowserFqdn field value
func (o *PrometheusHostInfo) GetUseBrowserFqdn() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseBrowserFqdn
}

// GetUseBrowserFqdnOk returns a tuple with the UseBrowserFqdn field value
// and a boolean to check if the value has been set.
func (o *PrometheusHostInfo) GetUseBrowserFqdnOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UseBrowserFqdn, true
}

// SetUseBrowserFqdn sets field value
func (o *PrometheusHostInfo) SetUseBrowserFqdn(v bool) {
	o.UseBrowserFqdn = v
}

func (o PrometheusHostInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["prometheus_url"] = o.PrometheusUrl
	}
	if true {
		toSerialize["use_browser_fqdn"] = o.UseBrowserFqdn
	}
	return json.Marshal(toSerialize)
}

type NullablePrometheusHostInfo struct {
	value *PrometheusHostInfo
	isSet bool
}

func (v NullablePrometheusHostInfo) Get() *PrometheusHostInfo {
	return v.value
}

func (v *NullablePrometheusHostInfo) Set(val *PrometheusHostInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusHostInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusHostInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusHostInfo(val *PrometheusHostInfo) *NullablePrometheusHostInfo {
	return &NullablePrometheusHostInfo{value: val, isSet: true}
}

func (v NullablePrometheusHostInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusHostInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

