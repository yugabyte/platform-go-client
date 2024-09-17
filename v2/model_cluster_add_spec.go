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

// ClusterAddSpec ClusterAddSpec  Request payload to add a new cluster to an existing YugabyteDB Universe. 
type ClusterAddSpec struct {
	// Cluster type can be one of READ_REPLICA, ADDON
	ClusterType string `json:"cluster_type"`
	// Set the number of nodes (tservers) to provision in this cluster
	NumNodes int32 `json:"num_nodes"`
	NodeSpec ClusterNodeSpec `json:"node_spec"`
	ProviderSpec ClusterProviderEditSpec `json:"provider_spec"`
	PlacementSpec *ClusterPlacementSpec `json:"placement_spec,omitempty"`
	// A map of strings representing a set of Tags and Values to apply on nodes in the aws/gcp/azu cloud. See https://docs.yugabyte.com/preview/yugabyte-platform/manage-deployments/instance-tags/.
	InstanceTags *map[string]string `json:"instance_tags,omitempty"`
	Gflags *ClusterGFlags `json:"gflags,omitempty"`
}

// NewClusterAddSpec instantiates a new ClusterAddSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAddSpec(clusterType string, numNodes int32, nodeSpec ClusterNodeSpec, providerSpec ClusterProviderEditSpec) *ClusterAddSpec {
	this := ClusterAddSpec{}
	this.ClusterType = clusterType
	this.NumNodes = numNodes
	this.NodeSpec = nodeSpec
	this.ProviderSpec = providerSpec
	return &this
}

// NewClusterAddSpecWithDefaults instantiates a new ClusterAddSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAddSpecWithDefaults() *ClusterAddSpec {
	this := ClusterAddSpec{}
	return &this
}

// GetClusterType returns the ClusterType field value
func (o *ClusterAddSpec) GetClusterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value
// and a boolean to check if the value has been set.
func (o *ClusterAddSpec) GetClusterTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterType, true
}

// SetClusterType sets field value
func (o *ClusterAddSpec) SetClusterType(v string) {
	o.ClusterType = v
}

// GetNumNodes returns the NumNodes field value
func (o *ClusterAddSpec) GetNumNodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumNodes
}

// GetNumNodesOk returns a tuple with the NumNodes field value
// and a boolean to check if the value has been set.
func (o *ClusterAddSpec) GetNumNodesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumNodes, true
}

// SetNumNodes sets field value
func (o *ClusterAddSpec) SetNumNodes(v int32) {
	o.NumNodes = v
}

// GetNodeSpec returns the NodeSpec field value
func (o *ClusterAddSpec) GetNodeSpec() ClusterNodeSpec {
	if o == nil {
		var ret ClusterNodeSpec
		return ret
	}

	return o.NodeSpec
}

// GetNodeSpecOk returns a tuple with the NodeSpec field value
// and a boolean to check if the value has been set.
func (o *ClusterAddSpec) GetNodeSpecOk() (*ClusterNodeSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeSpec, true
}

// SetNodeSpec sets field value
func (o *ClusterAddSpec) SetNodeSpec(v ClusterNodeSpec) {
	o.NodeSpec = v
}

// GetProviderSpec returns the ProviderSpec field value
func (o *ClusterAddSpec) GetProviderSpec() ClusterProviderEditSpec {
	if o == nil {
		var ret ClusterProviderEditSpec
		return ret
	}

	return o.ProviderSpec
}

// GetProviderSpecOk returns a tuple with the ProviderSpec field value
// and a boolean to check if the value has been set.
func (o *ClusterAddSpec) GetProviderSpecOk() (*ClusterProviderEditSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProviderSpec, true
}

// SetProviderSpec sets field value
func (o *ClusterAddSpec) SetProviderSpec(v ClusterProviderEditSpec) {
	o.ProviderSpec = v
}

// GetPlacementSpec returns the PlacementSpec field value if set, zero value otherwise.
func (o *ClusterAddSpec) GetPlacementSpec() ClusterPlacementSpec {
	if o == nil || o.PlacementSpec == nil {
		var ret ClusterPlacementSpec
		return ret
	}
	return *o.PlacementSpec
}

// GetPlacementSpecOk returns a tuple with the PlacementSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAddSpec) GetPlacementSpecOk() (*ClusterPlacementSpec, bool) {
	if o == nil || o.PlacementSpec == nil {
		return nil, false
	}
	return o.PlacementSpec, true
}

// HasPlacementSpec returns a boolean if a field has been set.
func (o *ClusterAddSpec) HasPlacementSpec() bool {
	if o != nil && o.PlacementSpec != nil {
		return true
	}

	return false
}

// SetPlacementSpec gets a reference to the given ClusterPlacementSpec and assigns it to the PlacementSpec field.
func (o *ClusterAddSpec) SetPlacementSpec(v ClusterPlacementSpec) {
	o.PlacementSpec = &v
}

// GetInstanceTags returns the InstanceTags field value if set, zero value otherwise.
func (o *ClusterAddSpec) GetInstanceTags() map[string]string {
	if o == nil || o.InstanceTags == nil {
		var ret map[string]string
		return ret
	}
	return *o.InstanceTags
}

// GetInstanceTagsOk returns a tuple with the InstanceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAddSpec) GetInstanceTagsOk() (*map[string]string, bool) {
	if o == nil || o.InstanceTags == nil {
		return nil, false
	}
	return o.InstanceTags, true
}

// HasInstanceTags returns a boolean if a field has been set.
func (o *ClusterAddSpec) HasInstanceTags() bool {
	if o != nil && o.InstanceTags != nil {
		return true
	}

	return false
}

// SetInstanceTags gets a reference to the given map[string]string and assigns it to the InstanceTags field.
func (o *ClusterAddSpec) SetInstanceTags(v map[string]string) {
	o.InstanceTags = &v
}

// GetGflags returns the Gflags field value if set, zero value otherwise.
func (o *ClusterAddSpec) GetGflags() ClusterGFlags {
	if o == nil || o.Gflags == nil {
		var ret ClusterGFlags
		return ret
	}
	return *o.Gflags
}

// GetGflagsOk returns a tuple with the Gflags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAddSpec) GetGflagsOk() (*ClusterGFlags, bool) {
	if o == nil || o.Gflags == nil {
		return nil, false
	}
	return o.Gflags, true
}

// HasGflags returns a boolean if a field has been set.
func (o *ClusterAddSpec) HasGflags() bool {
	if o != nil && o.Gflags != nil {
		return true
	}

	return false
}

// SetGflags gets a reference to the given ClusterGFlags and assigns it to the Gflags field.
func (o *ClusterAddSpec) SetGflags(v ClusterGFlags) {
	o.Gflags = &v
}

func (o ClusterAddSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_type"] = o.ClusterType
	}
	if true {
		toSerialize["num_nodes"] = o.NumNodes
	}
	if true {
		toSerialize["node_spec"] = o.NodeSpec
	}
	if true {
		toSerialize["provider_spec"] = o.ProviderSpec
	}
	if o.PlacementSpec != nil {
		toSerialize["placement_spec"] = o.PlacementSpec
	}
	if o.InstanceTags != nil {
		toSerialize["instance_tags"] = o.InstanceTags
	}
	if o.Gflags != nil {
		toSerialize["gflags"] = o.Gflags
	}
	return json.Marshal(toSerialize)
}

type NullableClusterAddSpec struct {
	value *ClusterAddSpec
	isSet bool
}

func (v NullableClusterAddSpec) Get() *ClusterAddSpec {
	return v.value
}

func (v *NullableClusterAddSpec) Set(val *ClusterAddSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAddSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAddSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAddSpec(val *ClusterAddSpec) *NullableClusterAddSpec {
	return &NullableClusterAddSpec{value: val, isSet: true}
}

func (v NullableClusterAddSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAddSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


