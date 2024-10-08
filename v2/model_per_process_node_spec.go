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

// PerProcessNodeSpec Instance settings for each node in the cluster. The instances can be onprem nodes, VMs in GCP/AWS/Azure, or pods in k8s. Part of AvailabilityZoneNodeSpec and ClusterNodeSpec.
type PerProcessNodeSpec struct {
	// Instance type for tserver/master nodes of cluster that determines the cpu and memory resources.
	InstanceType *string `json:"instance_type,omitempty"`
	StorageSpec *ClusterStorageSpec `json:"storage_spec,omitempty"`
}

// NewPerProcessNodeSpec instantiates a new PerProcessNodeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerProcessNodeSpec() *PerProcessNodeSpec {
	this := PerProcessNodeSpec{}
	return &this
}

// NewPerProcessNodeSpecWithDefaults instantiates a new PerProcessNodeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerProcessNodeSpecWithDefaults() *PerProcessNodeSpec {
	this := PerProcessNodeSpec{}
	return &this
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *PerProcessNodeSpec) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerProcessNodeSpec) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *PerProcessNodeSpec) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *PerProcessNodeSpec) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetStorageSpec returns the StorageSpec field value if set, zero value otherwise.
func (o *PerProcessNodeSpec) GetStorageSpec() ClusterStorageSpec {
	if o == nil || o.StorageSpec == nil {
		var ret ClusterStorageSpec
		return ret
	}
	return *o.StorageSpec
}

// GetStorageSpecOk returns a tuple with the StorageSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerProcessNodeSpec) GetStorageSpecOk() (*ClusterStorageSpec, bool) {
	if o == nil || o.StorageSpec == nil {
		return nil, false
	}
	return o.StorageSpec, true
}

// HasStorageSpec returns a boolean if a field has been set.
func (o *PerProcessNodeSpec) HasStorageSpec() bool {
	if o != nil && o.StorageSpec != nil {
		return true
	}

	return false
}

// SetStorageSpec gets a reference to the given ClusterStorageSpec and assigns it to the StorageSpec field.
func (o *PerProcessNodeSpec) SetStorageSpec(v ClusterStorageSpec) {
	o.StorageSpec = &v
}

func (o PerProcessNodeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceType != nil {
		toSerialize["instance_type"] = o.InstanceType
	}
	if o.StorageSpec != nil {
		toSerialize["storage_spec"] = o.StorageSpec
	}
	return json.Marshal(toSerialize)
}

type NullablePerProcessNodeSpec struct {
	value *PerProcessNodeSpec
	isSet bool
}

func (v NullablePerProcessNodeSpec) Get() *PerProcessNodeSpec {
	return v.value
}

func (v *NullablePerProcessNodeSpec) Set(val *PerProcessNodeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePerProcessNodeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePerProcessNodeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerProcessNodeSpec(val *PerProcessNodeSpec) *NullablePerProcessNodeSpec {
	return &NullablePerProcessNodeSpec{value: val, isSet: true}
}

func (v NullablePerProcessNodeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerProcessNodeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


