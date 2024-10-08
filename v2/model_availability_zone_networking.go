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

// AvailabilityZoneNetworking Networking properties for each node in the cluster. The settings can be configured at top-level for uniform node settings for both tserver and master nodes. Granular settings for tserver and master will be honored if provided (and dedicated_nodes is true or this is k8s cluster). Part of ClusterNetworkingSpec.
type AvailabilityZoneNetworking struct {
	// (Place holder for) Network settings that can be overridden per tserver or master process for nodes in the cluster. The node instances can be onprem nodes, VMs in GCP/AWS/Azure, or pods in k8s. Part of AvailabilityZoneNetworking.
	Tserver *map[string]interface{} `json:"tserver,omitempty"`
	// (Place holder for) Network settings that can be overridden per tserver or master process for nodes in the cluster. The node instances can be onprem nodes, VMs in GCP/AWS/Azure, or pods in k8s. Part of AvailabilityZoneNetworking.
	Master *map[string]interface{} `json:"master,omitempty"`
	ProxyConfig *NodeProxyConfig `json:"proxy_config,omitempty"`
}

// NewAvailabilityZoneNetworking instantiates a new AvailabilityZoneNetworking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityZoneNetworking() *AvailabilityZoneNetworking {
	this := AvailabilityZoneNetworking{}
	return &this
}

// NewAvailabilityZoneNetworkingWithDefaults instantiates a new AvailabilityZoneNetworking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityZoneNetworkingWithDefaults() *AvailabilityZoneNetworking {
	this := AvailabilityZoneNetworking{}
	return &this
}

// GetTserver returns the Tserver field value if set, zero value otherwise.
func (o *AvailabilityZoneNetworking) GetTserver() map[string]interface{} {
	if o == nil || o.Tserver == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tserver
}

// GetTserverOk returns a tuple with the Tserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZoneNetworking) GetTserverOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tserver == nil {
		return nil, false
	}
	return o.Tserver, true
}

// HasTserver returns a boolean if a field has been set.
func (o *AvailabilityZoneNetworking) HasTserver() bool {
	if o != nil && o.Tserver != nil {
		return true
	}

	return false
}

// SetTserver gets a reference to the given map[string]interface{} and assigns it to the Tserver field.
func (o *AvailabilityZoneNetworking) SetTserver(v map[string]interface{}) {
	o.Tserver = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *AvailabilityZoneNetworking) GetMaster() map[string]interface{} {
	if o == nil || o.Master == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZoneNetworking) GetMasterOk() (*map[string]interface{}, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *AvailabilityZoneNetworking) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given map[string]interface{} and assigns it to the Master field.
func (o *AvailabilityZoneNetworking) SetMaster(v map[string]interface{}) {
	o.Master = &v
}

// GetProxyConfig returns the ProxyConfig field value if set, zero value otherwise.
func (o *AvailabilityZoneNetworking) GetProxyConfig() NodeProxyConfig {
	if o == nil || o.ProxyConfig == nil {
		var ret NodeProxyConfig
		return ret
	}
	return *o.ProxyConfig
}

// GetProxyConfigOk returns a tuple with the ProxyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZoneNetworking) GetProxyConfigOk() (*NodeProxyConfig, bool) {
	if o == nil || o.ProxyConfig == nil {
		return nil, false
	}
	return o.ProxyConfig, true
}

// HasProxyConfig returns a boolean if a field has been set.
func (o *AvailabilityZoneNetworking) HasProxyConfig() bool {
	if o != nil && o.ProxyConfig != nil {
		return true
	}

	return false
}

// SetProxyConfig gets a reference to the given NodeProxyConfig and assigns it to the ProxyConfig field.
func (o *AvailabilityZoneNetworking) SetProxyConfig(v NodeProxyConfig) {
	o.ProxyConfig = &v
}

func (o AvailabilityZoneNetworking) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tserver != nil {
		toSerialize["tserver"] = o.Tserver
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
	if o.ProxyConfig != nil {
		toSerialize["proxy_config"] = o.ProxyConfig
	}
	return json.Marshal(toSerialize)
}

type NullableAvailabilityZoneNetworking struct {
	value *AvailabilityZoneNetworking
	isSet bool
}

func (v NullableAvailabilityZoneNetworking) Get() *AvailabilityZoneNetworking {
	return v.value
}

func (v *NullableAvailabilityZoneNetworking) Set(val *AvailabilityZoneNetworking) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityZoneNetworking) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityZoneNetworking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityZoneNetworking(val *AvailabilityZoneNetworking) *NullableAvailabilityZoneNetworking {
	return &NullableAvailabilityZoneNetworking{value: val, isSet: true}
}

func (v NullableAvailabilityZoneNetworking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityZoneNetworking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


