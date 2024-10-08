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

// ClusterNodeSpecAllOf struct for ClusterNodeSpecAllOf
type ClusterNodeSpecAllOf struct {
	// Whether to run tserver and master processes in dedicated nodes in this cluster. Defaults to false where master and tserver processes share the same node.
	DedicatedNodes *bool `json:"dedicated_nodes,omitempty"`
	K8sMasterResourceSpec *K8SNodeResourceSpec `json:"k8s_master_resource_spec,omitempty"`
	K8sTserverResourceSpec *K8SNodeResourceSpec `json:"k8s_tserver_resource_spec,omitempty"`
	// Granular node settings overridden per Availability Zone identified by AZ uuid.
	AzNodeSpec *map[string]AvailabilityZoneNodeSpec `json:"az_node_spec,omitempty"`
}

// NewClusterNodeSpecAllOf instantiates a new ClusterNodeSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNodeSpecAllOf() *ClusterNodeSpecAllOf {
	this := ClusterNodeSpecAllOf{}
	var dedicatedNodes bool = false
	this.DedicatedNodes = &dedicatedNodes
	return &this
}

// NewClusterNodeSpecAllOfWithDefaults instantiates a new ClusterNodeSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNodeSpecAllOfWithDefaults() *ClusterNodeSpecAllOf {
	this := ClusterNodeSpecAllOf{}
	var dedicatedNodes bool = false
	this.DedicatedNodes = &dedicatedNodes
	return &this
}

// GetDedicatedNodes returns the DedicatedNodes field value if set, zero value otherwise.
func (o *ClusterNodeSpecAllOf) GetDedicatedNodes() bool {
	if o == nil || o.DedicatedNodes == nil {
		var ret bool
		return ret
	}
	return *o.DedicatedNodes
}

// GetDedicatedNodesOk returns a tuple with the DedicatedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpecAllOf) GetDedicatedNodesOk() (*bool, bool) {
	if o == nil || o.DedicatedNodes == nil {
		return nil, false
	}
	return o.DedicatedNodes, true
}

// HasDedicatedNodes returns a boolean if a field has been set.
func (o *ClusterNodeSpecAllOf) HasDedicatedNodes() bool {
	if o != nil && o.DedicatedNodes != nil {
		return true
	}

	return false
}

// SetDedicatedNodes gets a reference to the given bool and assigns it to the DedicatedNodes field.
func (o *ClusterNodeSpecAllOf) SetDedicatedNodes(v bool) {
	o.DedicatedNodes = &v
}

// GetK8sMasterResourceSpec returns the K8sMasterResourceSpec field value if set, zero value otherwise.
func (o *ClusterNodeSpecAllOf) GetK8sMasterResourceSpec() K8SNodeResourceSpec {
	if o == nil || o.K8sMasterResourceSpec == nil {
		var ret K8SNodeResourceSpec
		return ret
	}
	return *o.K8sMasterResourceSpec
}

// GetK8sMasterResourceSpecOk returns a tuple with the K8sMasterResourceSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpecAllOf) GetK8sMasterResourceSpecOk() (*K8SNodeResourceSpec, bool) {
	if o == nil || o.K8sMasterResourceSpec == nil {
		return nil, false
	}
	return o.K8sMasterResourceSpec, true
}

// HasK8sMasterResourceSpec returns a boolean if a field has been set.
func (o *ClusterNodeSpecAllOf) HasK8sMasterResourceSpec() bool {
	if o != nil && o.K8sMasterResourceSpec != nil {
		return true
	}

	return false
}

// SetK8sMasterResourceSpec gets a reference to the given K8SNodeResourceSpec and assigns it to the K8sMasterResourceSpec field.
func (o *ClusterNodeSpecAllOf) SetK8sMasterResourceSpec(v K8SNodeResourceSpec) {
	o.K8sMasterResourceSpec = &v
}

// GetK8sTserverResourceSpec returns the K8sTserverResourceSpec field value if set, zero value otherwise.
func (o *ClusterNodeSpecAllOf) GetK8sTserverResourceSpec() K8SNodeResourceSpec {
	if o == nil || o.K8sTserverResourceSpec == nil {
		var ret K8SNodeResourceSpec
		return ret
	}
	return *o.K8sTserverResourceSpec
}

// GetK8sTserverResourceSpecOk returns a tuple with the K8sTserverResourceSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpecAllOf) GetK8sTserverResourceSpecOk() (*K8SNodeResourceSpec, bool) {
	if o == nil || o.K8sTserverResourceSpec == nil {
		return nil, false
	}
	return o.K8sTserverResourceSpec, true
}

// HasK8sTserverResourceSpec returns a boolean if a field has been set.
func (o *ClusterNodeSpecAllOf) HasK8sTserverResourceSpec() bool {
	if o != nil && o.K8sTserverResourceSpec != nil {
		return true
	}

	return false
}

// SetK8sTserverResourceSpec gets a reference to the given K8SNodeResourceSpec and assigns it to the K8sTserverResourceSpec field.
func (o *ClusterNodeSpecAllOf) SetK8sTserverResourceSpec(v K8SNodeResourceSpec) {
	o.K8sTserverResourceSpec = &v
}

// GetAzNodeSpec returns the AzNodeSpec field value if set, zero value otherwise.
func (o *ClusterNodeSpecAllOf) GetAzNodeSpec() map[string]AvailabilityZoneNodeSpec {
	if o == nil || o.AzNodeSpec == nil {
		var ret map[string]AvailabilityZoneNodeSpec
		return ret
	}
	return *o.AzNodeSpec
}

// GetAzNodeSpecOk returns a tuple with the AzNodeSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpecAllOf) GetAzNodeSpecOk() (*map[string]AvailabilityZoneNodeSpec, bool) {
	if o == nil || o.AzNodeSpec == nil {
		return nil, false
	}
	return o.AzNodeSpec, true
}

// HasAzNodeSpec returns a boolean if a field has been set.
func (o *ClusterNodeSpecAllOf) HasAzNodeSpec() bool {
	if o != nil && o.AzNodeSpec != nil {
		return true
	}

	return false
}

// SetAzNodeSpec gets a reference to the given map[string]AvailabilityZoneNodeSpec and assigns it to the AzNodeSpec field.
func (o *ClusterNodeSpecAllOf) SetAzNodeSpec(v map[string]AvailabilityZoneNodeSpec) {
	o.AzNodeSpec = &v
}

func (o ClusterNodeSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DedicatedNodes != nil {
		toSerialize["dedicated_nodes"] = o.DedicatedNodes
	}
	if o.K8sMasterResourceSpec != nil {
		toSerialize["k8s_master_resource_spec"] = o.K8sMasterResourceSpec
	}
	if o.K8sTserverResourceSpec != nil {
		toSerialize["k8s_tserver_resource_spec"] = o.K8sTserverResourceSpec
	}
	if o.AzNodeSpec != nil {
		toSerialize["az_node_spec"] = o.AzNodeSpec
	}
	return json.Marshal(toSerialize)
}

type NullableClusterNodeSpecAllOf struct {
	value *ClusterNodeSpecAllOf
	isSet bool
}

func (v NullableClusterNodeSpecAllOf) Get() *ClusterNodeSpecAllOf {
	return v.value
}

func (v *NullableClusterNodeSpecAllOf) Set(val *ClusterNodeSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNodeSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNodeSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNodeSpecAllOf(val *ClusterNodeSpecAllOf) *NullableClusterNodeSpecAllOf {
	return &NullableClusterNodeSpecAllOf{value: val, isSet: true}
}

func (v NullableClusterNodeSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNodeSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


