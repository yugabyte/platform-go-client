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

// ClusterNodeSpec Node settings (like CPU / memory) for each node in the cluster. The node settings configured at top-level are uniform settings for both tserver and master nodes. Granular settings for tserver and master (honoured if dedicated_nodes is true or this is k8s cluster) are available for certain node properties. Granular settings can also be overridden per Availability Zone. This is part of ClusterSpec.
type ClusterNodeSpec struct {
	// Instance type for tserver/master nodes of cluster that determines the cpu and memory resources.
	InstanceType *string `json:"instance_type,omitempty"`
	StorageSpec *ClusterStorageSpec `json:"storage_spec,omitempty"`
	// Amount of memory in MB to limit the postgres process using the ysql cgroup. The value should be greater than 0. When set to 0 it results in no cgroup limits. For a read replica cluster, setting this value to null or -1 would inherit this value from the primary cluster. Applicable only for nodes running as Linux VMs on AWS/GCP/Azure Cloud Provider. Only used internally by YBM.
	CgroupSize *int32 `json:"cgroup_size,omitempty"`
	Tserver *PerProcessNodeSpec `json:"tserver,omitempty"`
	Master *PerProcessNodeSpec `json:"master,omitempty"`
	// Whether to run tserver and master processes in dedicated nodes in this cluster. Defaults to false where master and tserver processes share the same node.
	DedicatedNodes *bool `json:"dedicated_nodes,omitempty"`
	K8sMasterResourceSpec *K8SNodeResourceSpec `json:"k8s_master_resource_spec,omitempty"`
	K8sTserverResourceSpec *K8SNodeResourceSpec `json:"k8s_tserver_resource_spec,omitempty"`
	// Granular node settings overridden per Availability Zone identified by AZ uuid.
	AzNodeSpec *map[string]AvailabilityZoneNodeSpec `json:"az_node_spec,omitempty"`
}

// NewClusterNodeSpec instantiates a new ClusterNodeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNodeSpec() *ClusterNodeSpec {
	this := ClusterNodeSpec{}
	var dedicatedNodes bool = false
	this.DedicatedNodes = &dedicatedNodes
	return &this
}

// NewClusterNodeSpecWithDefaults instantiates a new ClusterNodeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNodeSpecWithDefaults() *ClusterNodeSpec {
	this := ClusterNodeSpec{}
	var dedicatedNodes bool = false
	this.DedicatedNodes = &dedicatedNodes
	return &this
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *ClusterNodeSpec) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetStorageSpec returns the StorageSpec field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetStorageSpec() ClusterStorageSpec {
	if o == nil || o.StorageSpec == nil {
		var ret ClusterStorageSpec
		return ret
	}
	return *o.StorageSpec
}

// GetStorageSpecOk returns a tuple with the StorageSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetStorageSpecOk() (*ClusterStorageSpec, bool) {
	if o == nil || o.StorageSpec == nil {
		return nil, false
	}
	return o.StorageSpec, true
}

// HasStorageSpec returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasStorageSpec() bool {
	if o != nil && o.StorageSpec != nil {
		return true
	}

	return false
}

// SetStorageSpec gets a reference to the given ClusterStorageSpec and assigns it to the StorageSpec field.
func (o *ClusterNodeSpec) SetStorageSpec(v ClusterStorageSpec) {
	o.StorageSpec = &v
}

// GetCgroupSize returns the CgroupSize field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetCgroupSize() int32 {
	if o == nil || o.CgroupSize == nil {
		var ret int32
		return ret
	}
	return *o.CgroupSize
}

// GetCgroupSizeOk returns a tuple with the CgroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetCgroupSizeOk() (*int32, bool) {
	if o == nil || o.CgroupSize == nil {
		return nil, false
	}
	return o.CgroupSize, true
}

// HasCgroupSize returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasCgroupSize() bool {
	if o != nil && o.CgroupSize != nil {
		return true
	}

	return false
}

// SetCgroupSize gets a reference to the given int32 and assigns it to the CgroupSize field.
func (o *ClusterNodeSpec) SetCgroupSize(v int32) {
	o.CgroupSize = &v
}

// GetTserver returns the Tserver field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetTserver() PerProcessNodeSpec {
	if o == nil || o.Tserver == nil {
		var ret PerProcessNodeSpec
		return ret
	}
	return *o.Tserver
}

// GetTserverOk returns a tuple with the Tserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetTserverOk() (*PerProcessNodeSpec, bool) {
	if o == nil || o.Tserver == nil {
		return nil, false
	}
	return o.Tserver, true
}

// HasTserver returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasTserver() bool {
	if o != nil && o.Tserver != nil {
		return true
	}

	return false
}

// SetTserver gets a reference to the given PerProcessNodeSpec and assigns it to the Tserver field.
func (o *ClusterNodeSpec) SetTserver(v PerProcessNodeSpec) {
	o.Tserver = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetMaster() PerProcessNodeSpec {
	if o == nil || o.Master == nil {
		var ret PerProcessNodeSpec
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetMasterOk() (*PerProcessNodeSpec, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given PerProcessNodeSpec and assigns it to the Master field.
func (o *ClusterNodeSpec) SetMaster(v PerProcessNodeSpec) {
	o.Master = &v
}

// GetDedicatedNodes returns the DedicatedNodes field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetDedicatedNodes() bool {
	if o == nil || o.DedicatedNodes == nil {
		var ret bool
		return ret
	}
	return *o.DedicatedNodes
}

// GetDedicatedNodesOk returns a tuple with the DedicatedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetDedicatedNodesOk() (*bool, bool) {
	if o == nil || o.DedicatedNodes == nil {
		return nil, false
	}
	return o.DedicatedNodes, true
}

// HasDedicatedNodes returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasDedicatedNodes() bool {
	if o != nil && o.DedicatedNodes != nil {
		return true
	}

	return false
}

// SetDedicatedNodes gets a reference to the given bool and assigns it to the DedicatedNodes field.
func (o *ClusterNodeSpec) SetDedicatedNodes(v bool) {
	o.DedicatedNodes = &v
}

// GetK8sMasterResourceSpec returns the K8sMasterResourceSpec field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetK8sMasterResourceSpec() K8SNodeResourceSpec {
	if o == nil || o.K8sMasterResourceSpec == nil {
		var ret K8SNodeResourceSpec
		return ret
	}
	return *o.K8sMasterResourceSpec
}

// GetK8sMasterResourceSpecOk returns a tuple with the K8sMasterResourceSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetK8sMasterResourceSpecOk() (*K8SNodeResourceSpec, bool) {
	if o == nil || o.K8sMasterResourceSpec == nil {
		return nil, false
	}
	return o.K8sMasterResourceSpec, true
}

// HasK8sMasterResourceSpec returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasK8sMasterResourceSpec() bool {
	if o != nil && o.K8sMasterResourceSpec != nil {
		return true
	}

	return false
}

// SetK8sMasterResourceSpec gets a reference to the given K8SNodeResourceSpec and assigns it to the K8sMasterResourceSpec field.
func (o *ClusterNodeSpec) SetK8sMasterResourceSpec(v K8SNodeResourceSpec) {
	o.K8sMasterResourceSpec = &v
}

// GetK8sTserverResourceSpec returns the K8sTserverResourceSpec field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetK8sTserverResourceSpec() K8SNodeResourceSpec {
	if o == nil || o.K8sTserverResourceSpec == nil {
		var ret K8SNodeResourceSpec
		return ret
	}
	return *o.K8sTserverResourceSpec
}

// GetK8sTserverResourceSpecOk returns a tuple with the K8sTserverResourceSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetK8sTserverResourceSpecOk() (*K8SNodeResourceSpec, bool) {
	if o == nil || o.K8sTserverResourceSpec == nil {
		return nil, false
	}
	return o.K8sTserverResourceSpec, true
}

// HasK8sTserverResourceSpec returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasK8sTserverResourceSpec() bool {
	if o != nil && o.K8sTserverResourceSpec != nil {
		return true
	}

	return false
}

// SetK8sTserverResourceSpec gets a reference to the given K8SNodeResourceSpec and assigns it to the K8sTserverResourceSpec field.
func (o *ClusterNodeSpec) SetK8sTserverResourceSpec(v K8SNodeResourceSpec) {
	o.K8sTserverResourceSpec = &v
}

// GetAzNodeSpec returns the AzNodeSpec field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetAzNodeSpec() map[string]AvailabilityZoneNodeSpec {
	if o == nil || o.AzNodeSpec == nil {
		var ret map[string]AvailabilityZoneNodeSpec
		return ret
	}
	return *o.AzNodeSpec
}

// GetAzNodeSpecOk returns a tuple with the AzNodeSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetAzNodeSpecOk() (*map[string]AvailabilityZoneNodeSpec, bool) {
	if o == nil || o.AzNodeSpec == nil {
		return nil, false
	}
	return o.AzNodeSpec, true
}

// HasAzNodeSpec returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasAzNodeSpec() bool {
	if o != nil && o.AzNodeSpec != nil {
		return true
	}

	return false
}

// SetAzNodeSpec gets a reference to the given map[string]AvailabilityZoneNodeSpec and assigns it to the AzNodeSpec field.
func (o *ClusterNodeSpec) SetAzNodeSpec(v map[string]AvailabilityZoneNodeSpec) {
	o.AzNodeSpec = &v
}

func (o ClusterNodeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceType != nil {
		toSerialize["instance_type"] = o.InstanceType
	}
	if o.StorageSpec != nil {
		toSerialize["storage_spec"] = o.StorageSpec
	}
	if o.CgroupSize != nil {
		toSerialize["cgroup_size"] = o.CgroupSize
	}
	if o.Tserver != nil {
		toSerialize["tserver"] = o.Tserver
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
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

type NullableClusterNodeSpec struct {
	value *ClusterNodeSpec
	isSet bool
}

func (v NullableClusterNodeSpec) Get() *ClusterNodeSpec {
	return v.value
}

func (v *NullableClusterNodeSpec) Set(val *ClusterNodeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNodeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNodeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNodeSpec(val *ClusterNodeSpec) *NullableClusterNodeSpec {
	return &NullableClusterNodeSpec{value: val, isSet: true}
}

func (v NullableClusterNodeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNodeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


