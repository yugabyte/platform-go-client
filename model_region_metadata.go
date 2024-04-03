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

// RegionMetadata struct for RegionMetadata
type RegionMetadata struct {
	RegionMetadata map[string]RegionMetadataInfo `json:"regionMetadata"`
}

// NewRegionMetadata instantiates a new RegionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionMetadata(regionMetadata map[string]RegionMetadataInfo) *RegionMetadata {
	this := RegionMetadata{}
	this.RegionMetadata = regionMetadata
	return &this
}

// NewRegionMetadataWithDefaults instantiates a new RegionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionMetadataWithDefaults() *RegionMetadata {
	this := RegionMetadata{}
	return &this
}

// GetRegionMetadata returns the RegionMetadata field value
func (o *RegionMetadata) GetRegionMetadata() map[string]RegionMetadataInfo {
	if o == nil {
		var ret map[string]RegionMetadataInfo
		return ret
	}

	return o.RegionMetadata
}

// GetRegionMetadataOk returns a tuple with the RegionMetadata field value
// and a boolean to check if the value has been set.
func (o *RegionMetadata) GetRegionMetadataOk() (*map[string]RegionMetadataInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionMetadata, true
}

// SetRegionMetadata sets field value
func (o *RegionMetadata) SetRegionMetadata(v map[string]RegionMetadataInfo) {
	o.RegionMetadata = v
}

func (o RegionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regionMetadata"] = o.RegionMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableRegionMetadata struct {
	value *RegionMetadata
	isSet bool
}

func (v NullableRegionMetadata) Get() *RegionMetadata {
	return v.value
}

func (v *NullableRegionMetadata) Set(val *RegionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionMetadata(val *RegionMetadata) *NullableRegionMetadata {
	return &NullableRegionMetadata{value: val, isSet: true}
}

func (v NullableRegionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


