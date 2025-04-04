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

// ClusterStorageSpec Storage volume specification that is used for tserver nodes in this cluster. Part of ClusterSpec, ClusterAddSpec and ClusterEditSpec.
type ClusterStorageSpec struct {
	// The size of each volume in each instance. Could be modified in payload for /resize_node API call
	VolumeSize int32 `json:"volume_size"`
	// Number of volumes to be mounted on this instance at the default path
	NumVolumes int32 `json:"num_volumes"`
	// Comma-separated list of mount points for the volumes in each instance. Required for an onprem cluster.
	MountPoints *string `json:"mount_points,omitempty"`
	// Name of the storage class, if this is a kubernetes cluster
	StorageClass *string `json:"storage_class,omitempty"`
	// Storage type used for this instance, if this is a aws (IO1, GP2, GP3), gcp (Scratch, Persistent, Hyperdisk_Balanced, Hyperdisk_Extreme) or azu (StandardSSD_LRS, Premium_LRS, PremiumV2_LRS, UltraSSD_LRS) cluster.
	StorageType *string `json:"storage_type,omitempty"`
	// Desired IOPS for the volumes mounted on this aws, gcp or azu instance
	DiskIops *int32 `json:"disk_iops,omitempty"`
	// Desired throughput for the volumes mounted on this aws, gcp or azu instance
	Throughput *int32 `json:"throughput,omitempty"`
}

// NewClusterStorageSpec instantiates a new ClusterStorageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStorageSpec(volumeSize int32, numVolumes int32) *ClusterStorageSpec {
	this := ClusterStorageSpec{}
	this.VolumeSize = volumeSize
	this.NumVolumes = numVolumes
	return &this
}

// NewClusterStorageSpecWithDefaults instantiates a new ClusterStorageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStorageSpecWithDefaults() *ClusterStorageSpec {
	this := ClusterStorageSpec{}
	return &this
}

// GetVolumeSize returns the VolumeSize field value
func (o *ClusterStorageSpec) GetVolumeSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageSpec) GetVolumeSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VolumeSize, true
}

// SetVolumeSize sets field value
func (o *ClusterStorageSpec) SetVolumeSize(v int32) {
	o.VolumeSize = v
}

// GetNumVolumes returns the NumVolumes field value
func (o *ClusterStorageSpec) GetNumVolumes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumVolumes
}

// GetNumVolumesOk returns a tuple with the NumVolumes field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageSpec) GetNumVolumesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumVolumes, true
}

// SetNumVolumes sets field value
func (o *ClusterStorageSpec) SetNumVolumes(v int32) {
	o.NumVolumes = v
}

// GetMountPoints returns the MountPoints field value if set, zero value otherwise.
func (o *ClusterStorageSpec) GetMountPoints() string {
	if o == nil || o.MountPoints == nil {
		var ret string
		return ret
	}
	return *o.MountPoints
}

// GetMountPointsOk returns a tuple with the MountPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageSpec) GetMountPointsOk() (*string, bool) {
	if o == nil || o.MountPoints == nil {
		return nil, false
	}
	return o.MountPoints, true
}

// HasMountPoints returns a boolean if a field has been set.
func (o *ClusterStorageSpec) HasMountPoints() bool {
	if o != nil && o.MountPoints != nil {
		return true
	}

	return false
}

// SetMountPoints gets a reference to the given string and assigns it to the MountPoints field.
func (o *ClusterStorageSpec) SetMountPoints(v string) {
	o.MountPoints = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *ClusterStorageSpec) GetStorageClass() string {
	if o == nil || o.StorageClass == nil {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageSpec) GetStorageClassOk() (*string, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *ClusterStorageSpec) HasStorageClass() bool {
	if o != nil && o.StorageClass != nil {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *ClusterStorageSpec) SetStorageClass(v string) {
	o.StorageClass = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *ClusterStorageSpec) GetStorageType() string {
	if o == nil || o.StorageType == nil {
		var ret string
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageSpec) GetStorageTypeOk() (*string, bool) {
	if o == nil || o.StorageType == nil {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *ClusterStorageSpec) HasStorageType() bool {
	if o != nil && o.StorageType != nil {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given string and assigns it to the StorageType field.
func (o *ClusterStorageSpec) SetStorageType(v string) {
	o.StorageType = &v
}

// GetDiskIops returns the DiskIops field value if set, zero value otherwise.
func (o *ClusterStorageSpec) GetDiskIops() int32 {
	if o == nil || o.DiskIops == nil {
		var ret int32
		return ret
	}
	return *o.DiskIops
}

// GetDiskIopsOk returns a tuple with the DiskIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageSpec) GetDiskIopsOk() (*int32, bool) {
	if o == nil || o.DiskIops == nil {
		return nil, false
	}
	return o.DiskIops, true
}

// HasDiskIops returns a boolean if a field has been set.
func (o *ClusterStorageSpec) HasDiskIops() bool {
	if o != nil && o.DiskIops != nil {
		return true
	}

	return false
}

// SetDiskIops gets a reference to the given int32 and assigns it to the DiskIops field.
func (o *ClusterStorageSpec) SetDiskIops(v int32) {
	o.DiskIops = &v
}

// GetThroughput returns the Throughput field value if set, zero value otherwise.
func (o *ClusterStorageSpec) GetThroughput() int32 {
	if o == nil || o.Throughput == nil {
		var ret int32
		return ret
	}
	return *o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageSpec) GetThroughputOk() (*int32, bool) {
	if o == nil || o.Throughput == nil {
		return nil, false
	}
	return o.Throughput, true
}

// HasThroughput returns a boolean if a field has been set.
func (o *ClusterStorageSpec) HasThroughput() bool {
	if o != nil && o.Throughput != nil {
		return true
	}

	return false
}

// SetThroughput gets a reference to the given int32 and assigns it to the Throughput field.
func (o *ClusterStorageSpec) SetThroughput(v int32) {
	o.Throughput = &v
}

func (o ClusterStorageSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["volume_size"] = o.VolumeSize
	}
	if true {
		toSerialize["num_volumes"] = o.NumVolumes
	}
	if o.MountPoints != nil {
		toSerialize["mount_points"] = o.MountPoints
	}
	if o.StorageClass != nil {
		toSerialize["storage_class"] = o.StorageClass
	}
	if o.StorageType != nil {
		toSerialize["storage_type"] = o.StorageType
	}
	if o.DiskIops != nil {
		toSerialize["disk_iops"] = o.DiskIops
	}
	if o.Throughput != nil {
		toSerialize["throughput"] = o.Throughput
	}
	return json.Marshal(toSerialize)
}

type NullableClusterStorageSpec struct {
	value *ClusterStorageSpec
	isSet bool
}

func (v NullableClusterStorageSpec) Get() *ClusterStorageSpec {
	return v.value
}

func (v *NullableClusterStorageSpec) Set(val *ClusterStorageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStorageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStorageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStorageSpec(val *ClusterStorageSpec) *NullableClusterStorageSpec {
	return &NullableClusterStorageSpec{value: val, isSet: true}
}

func (v NullableClusterStorageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStorageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


