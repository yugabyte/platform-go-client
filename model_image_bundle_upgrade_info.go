/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ywclient

import (
	"encoding/json"
)

// ImageBundleUpgradeInfo struct for ImageBundleUpgradeInfo
type ImageBundleUpgradeInfo struct {
	ClusterUuid string `json:"clusterUuid"`
	ImageBundleUuid string `json:"imageBundleUuid"`
}

// NewImageBundleUpgradeInfo instantiates a new ImageBundleUpgradeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageBundleUpgradeInfo(clusterUuid string, imageBundleUuid string) *ImageBundleUpgradeInfo {
	this := ImageBundleUpgradeInfo{}
	this.ClusterUuid = clusterUuid
	this.ImageBundleUuid = imageBundleUuid
	return &this
}

// NewImageBundleUpgradeInfoWithDefaults instantiates a new ImageBundleUpgradeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageBundleUpgradeInfoWithDefaults() *ImageBundleUpgradeInfo {
	this := ImageBundleUpgradeInfo{}
	return &this
}

// GetClusterUuid returns the ClusterUuid field value
func (o *ImageBundleUpgradeInfo) GetClusterUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value
// and a boolean to check if the value has been set.
func (o *ImageBundleUpgradeInfo) GetClusterUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterUuid, true
}

// SetClusterUuid sets field value
func (o *ImageBundleUpgradeInfo) SetClusterUuid(v string) {
	o.ClusterUuid = v
}

// GetImageBundleUuid returns the ImageBundleUuid field value
func (o *ImageBundleUpgradeInfo) GetImageBundleUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageBundleUuid
}

// GetImageBundleUuidOk returns a tuple with the ImageBundleUuid field value
// and a boolean to check if the value has been set.
func (o *ImageBundleUpgradeInfo) GetImageBundleUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageBundleUuid, true
}

// SetImageBundleUuid sets field value
func (o *ImageBundleUpgradeInfo) SetImageBundleUuid(v string) {
	o.ImageBundleUuid = v
}

func (o ImageBundleUpgradeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clusterUuid"] = o.ClusterUuid
	}
	if true {
		toSerialize["imageBundleUuid"] = o.ImageBundleUuid
	}
	return json.Marshal(toSerialize)
}

type NullableImageBundleUpgradeInfo struct {
	value *ImageBundleUpgradeInfo
	isSet bool
}

func (v NullableImageBundleUpgradeInfo) Get() *ImageBundleUpgradeInfo {
	return v.value
}

func (v *NullableImageBundleUpgradeInfo) Set(val *ImageBundleUpgradeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBundleUpgradeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBundleUpgradeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBundleUpgradeInfo(val *ImageBundleUpgradeInfo) *NullableImageBundleUpgradeInfo {
	return &NullableImageBundleUpgradeInfo{value: val, isSet: true}
}

func (v NullableImageBundleUpgradeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBundleUpgradeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

