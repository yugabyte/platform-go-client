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

// ClusterEditSpec Edit Cluster Spec. Part of UniverseEditSpec.
type ClusterEditSpec struct {
	// The system generated cluster uuid to edit. This can be fetched from ClusterInfo.
	Uuid string `json:"uuid"`
	// Set the number of nodes (tservers) to provision in this cluster
	NumNodes *int32 `json:"num_nodes,omitempty"`
	NodeSpec *ClusterNodeSpec `json:"node_spec,omitempty"`
	ProviderSpec *ClusterProviderEditSpec `json:"provider_spec,omitempty"`
	PlacementSpec *ClusterPlacementSpec `json:"placement_spec,omitempty"`
	// A map of strings representing a set of Tags and Values to apply on nodes in the aws/gcp/azu cloud. See https://docs.yugabyte.com/preview/yugabyte-platform/manage-deployments/instance-tags/.
	InstanceTags *map[string]string `json:"instance_tags,omitempty"`
}

// NewClusterEditSpec instantiates a new ClusterEditSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterEditSpec(uuid string) *ClusterEditSpec {
	this := ClusterEditSpec{}
	this.Uuid = uuid
	return &this
}

// NewClusterEditSpecWithDefaults instantiates a new ClusterEditSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterEditSpecWithDefaults() *ClusterEditSpec {
	this := ClusterEditSpec{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *ClusterEditSpec) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ClusterEditSpec) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ClusterEditSpec) SetUuid(v string) {
	o.Uuid = v
}

// GetNumNodes returns the NumNodes field value if set, zero value otherwise.
func (o *ClusterEditSpec) GetNumNodes() int32 {
	if o == nil || o.NumNodes == nil {
		var ret int32
		return ret
	}
	return *o.NumNodes
}

// GetNumNodesOk returns a tuple with the NumNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEditSpec) GetNumNodesOk() (*int32, bool) {
	if o == nil || o.NumNodes == nil {
		return nil, false
	}
	return o.NumNodes, true
}

// HasNumNodes returns a boolean if a field has been set.
func (o *ClusterEditSpec) HasNumNodes() bool {
	if o != nil && o.NumNodes != nil {
		return true
	}

	return false
}

// SetNumNodes gets a reference to the given int32 and assigns it to the NumNodes field.
func (o *ClusterEditSpec) SetNumNodes(v int32) {
	o.NumNodes = &v
}

// GetNodeSpec returns the NodeSpec field value if set, zero value otherwise.
func (o *ClusterEditSpec) GetNodeSpec() ClusterNodeSpec {
	if o == nil || o.NodeSpec == nil {
		var ret ClusterNodeSpec
		return ret
	}
	return *o.NodeSpec
}

// GetNodeSpecOk returns a tuple with the NodeSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEditSpec) GetNodeSpecOk() (*ClusterNodeSpec, bool) {
	if o == nil || o.NodeSpec == nil {
		return nil, false
	}
	return o.NodeSpec, true
}

// HasNodeSpec returns a boolean if a field has been set.
func (o *ClusterEditSpec) HasNodeSpec() bool {
	if o != nil && o.NodeSpec != nil {
		return true
	}

	return false
}

// SetNodeSpec gets a reference to the given ClusterNodeSpec and assigns it to the NodeSpec field.
func (o *ClusterEditSpec) SetNodeSpec(v ClusterNodeSpec) {
	o.NodeSpec = &v
}

// GetProviderSpec returns the ProviderSpec field value if set, zero value otherwise.
func (o *ClusterEditSpec) GetProviderSpec() ClusterProviderEditSpec {
	if o == nil || o.ProviderSpec == nil {
		var ret ClusterProviderEditSpec
		return ret
	}
	return *o.ProviderSpec
}

// GetProviderSpecOk returns a tuple with the ProviderSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEditSpec) GetProviderSpecOk() (*ClusterProviderEditSpec, bool) {
	if o == nil || o.ProviderSpec == nil {
		return nil, false
	}
	return o.ProviderSpec, true
}

// HasProviderSpec returns a boolean if a field has been set.
func (o *ClusterEditSpec) HasProviderSpec() bool {
	if o != nil && o.ProviderSpec != nil {
		return true
	}

	return false
}

// SetProviderSpec gets a reference to the given ClusterProviderEditSpec and assigns it to the ProviderSpec field.
func (o *ClusterEditSpec) SetProviderSpec(v ClusterProviderEditSpec) {
	o.ProviderSpec = &v
}

// GetPlacementSpec returns the PlacementSpec field value if set, zero value otherwise.
func (o *ClusterEditSpec) GetPlacementSpec() ClusterPlacementSpec {
	if o == nil || o.PlacementSpec == nil {
		var ret ClusterPlacementSpec
		return ret
	}
	return *o.PlacementSpec
}

// GetPlacementSpecOk returns a tuple with the PlacementSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEditSpec) GetPlacementSpecOk() (*ClusterPlacementSpec, bool) {
	if o == nil || o.PlacementSpec == nil {
		return nil, false
	}
	return o.PlacementSpec, true
}

// HasPlacementSpec returns a boolean if a field has been set.
func (o *ClusterEditSpec) HasPlacementSpec() bool {
	if o != nil && o.PlacementSpec != nil {
		return true
	}

	return false
}

// SetPlacementSpec gets a reference to the given ClusterPlacementSpec and assigns it to the PlacementSpec field.
func (o *ClusterEditSpec) SetPlacementSpec(v ClusterPlacementSpec) {
	o.PlacementSpec = &v
}

// GetInstanceTags returns the InstanceTags field value if set, zero value otherwise.
func (o *ClusterEditSpec) GetInstanceTags() map[string]string {
	if o == nil || o.InstanceTags == nil {
		var ret map[string]string
		return ret
	}
	return *o.InstanceTags
}

// GetInstanceTagsOk returns a tuple with the InstanceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEditSpec) GetInstanceTagsOk() (*map[string]string, bool) {
	if o == nil || o.InstanceTags == nil {
		return nil, false
	}
	return o.InstanceTags, true
}

// HasInstanceTags returns a boolean if a field has been set.
func (o *ClusterEditSpec) HasInstanceTags() bool {
	if o != nil && o.InstanceTags != nil {
		return true
	}

	return false
}

// SetInstanceTags gets a reference to the given map[string]string and assigns it to the InstanceTags field.
func (o *ClusterEditSpec) SetInstanceTags(v map[string]string) {
	o.InstanceTags = &v
}

func (o ClusterEditSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if o.NumNodes != nil {
		toSerialize["num_nodes"] = o.NumNodes
	}
	if o.NodeSpec != nil {
		toSerialize["node_spec"] = o.NodeSpec
	}
	if o.ProviderSpec != nil {
		toSerialize["provider_spec"] = o.ProviderSpec
	}
	if o.PlacementSpec != nil {
		toSerialize["placement_spec"] = o.PlacementSpec
	}
	if o.InstanceTags != nil {
		toSerialize["instance_tags"] = o.InstanceTags
	}
	return json.Marshal(toSerialize)
}

type NullableClusterEditSpec struct {
	value *ClusterEditSpec
	isSet bool
}

func (v NullableClusterEditSpec) Get() *ClusterEditSpec {
	return v.value
}

func (v *NullableClusterEditSpec) Set(val *ClusterEditSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterEditSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterEditSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterEditSpec(val *ClusterEditSpec) *NullableClusterEditSpec {
	return &NullableClusterEditSpec{value: val, isSet: true}
}

func (v NullableClusterEditSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterEditSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


