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

// GCSLocation struct for GCSLocation
type GCSLocation struct {
	Paths *PackagePaths `json:"paths,omitempty"`
}

// NewGCSLocation instantiates a new GCSLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCSLocation() *GCSLocation {
	this := GCSLocation{}
	return &this
}

// NewGCSLocationWithDefaults instantiates a new GCSLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCSLocationWithDefaults() *GCSLocation {
	this := GCSLocation{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *GCSLocation) GetPaths() PackagePaths {
	if o == nil || o.Paths == nil {
		var ret PackagePaths
		return ret
	}
	return *o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSLocation) GetPathsOk() (*PackagePaths, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *GCSLocation) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given PackagePaths and assigns it to the Paths field.
func (o *GCSLocation) SetPaths(v PackagePaths) {
	o.Paths = &v
}

func (o GCSLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	return json.Marshal(toSerialize)
}

type NullableGCSLocation struct {
	value *GCSLocation
	isSet bool
}

func (v NullableGCSLocation) Get() *GCSLocation {
	return v.value
}

func (v *NullableGCSLocation) Set(val *GCSLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGCSLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGCSLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCSLocation(val *GCSLocation) *NullableGCSLocation {
	return &NullableGCSLocation{value: val, isSet: true}
}

func (v NullableGCSLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCSLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


