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

// ClusterPlacementSpec Data Placement Specification for the cluster. Part of ClusterSpec.  Note that this is optional to configure. YugabyteDB Anywhere will automatically place the data based on the available resources. If this data placement is configured, then YugabyteDB Anywhere will use this as a \"hint\" and the data will be placed based on this configuration on a best-effort basis. 
type ClusterPlacementSpec struct {
	CloudList []PlacementCloud `json:"cloud_list"`
}

// NewClusterPlacementSpec instantiates a new ClusterPlacementSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterPlacementSpec(cloudList []PlacementCloud) *ClusterPlacementSpec {
	this := ClusterPlacementSpec{}
	this.CloudList = cloudList
	return &this
}

// NewClusterPlacementSpecWithDefaults instantiates a new ClusterPlacementSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterPlacementSpecWithDefaults() *ClusterPlacementSpec {
	this := ClusterPlacementSpec{}
	return &this
}

// GetCloudList returns the CloudList field value
func (o *ClusterPlacementSpec) GetCloudList() []PlacementCloud {
	if o == nil {
		var ret []PlacementCloud
		return ret
	}

	return o.CloudList
}

// GetCloudListOk returns a tuple with the CloudList field value
// and a boolean to check if the value has been set.
func (o *ClusterPlacementSpec) GetCloudListOk() (*[]PlacementCloud, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudList, true
}

// SetCloudList sets field value
func (o *ClusterPlacementSpec) SetCloudList(v []PlacementCloud) {
	o.CloudList = v
}

func (o ClusterPlacementSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloud_list"] = o.CloudList
	}
	return json.Marshal(toSerialize)
}

type NullableClusterPlacementSpec struct {
	value *ClusterPlacementSpec
	isSet bool
}

func (v NullableClusterPlacementSpec) Get() *ClusterPlacementSpec {
	return v.value
}

func (v *NullableClusterPlacementSpec) Set(val *ClusterPlacementSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterPlacementSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterPlacementSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterPlacementSpec(val *ClusterPlacementSpec) *NullableClusterPlacementSpec {
	return &NullableClusterPlacementSpec{value: val, isSet: true}
}

func (v NullableClusterPlacementSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterPlacementSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


