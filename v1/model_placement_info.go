/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// PlacementInfo struct for PlacementInfo
type PlacementInfo struct {
	CloudList []PlacementCloud `json:"cloudList"`
}

// NewPlacementInfo instantiates a new PlacementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlacementInfo(cloudList []PlacementCloud) *PlacementInfo {
	this := PlacementInfo{}
	this.CloudList = cloudList
	return &this
}

// NewPlacementInfoWithDefaults instantiates a new PlacementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlacementInfoWithDefaults() *PlacementInfo {
	this := PlacementInfo{}
	return &this
}

// GetCloudList returns the CloudList field value
func (o *PlacementInfo) GetCloudList() []PlacementCloud {
	if o == nil {
		var ret []PlacementCloud
		return ret
	}

	return o.CloudList
}

// GetCloudListOk returns a tuple with the CloudList field value
// and a boolean to check if the value has been set.
func (o *PlacementInfo) GetCloudListOk() (*[]PlacementCloud, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudList, true
}

// SetCloudList sets field value
func (o *PlacementInfo) SetCloudList(v []PlacementCloud) {
	o.CloudList = v
}

func (o PlacementInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloudList"] = o.CloudList
	}
	return json.Marshal(toSerialize)
}

type NullablePlacementInfo struct {
	value *PlacementInfo
	isSet bool
}

func (v NullablePlacementInfo) Get() *PlacementInfo {
	return v.value
}

func (v *NullablePlacementInfo) Set(val *PlacementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacementInfo(val *PlacementInfo) *NullablePlacementInfo {
	return &NullablePlacementInfo{value: val, isSet: true}
}

func (v NullablePlacementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlacementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


