/*
 * Yugabyte Platform APIs
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

// HttpLocation struct for HttpLocation
type HttpLocation struct {
	Paths *PackagePaths `json:"paths,omitempty"`
}

// NewHttpLocation instantiates a new HttpLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpLocation() *HttpLocation {
	this := HttpLocation{}
	return &this
}

// NewHttpLocationWithDefaults instantiates a new HttpLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpLocationWithDefaults() *HttpLocation {
	this := HttpLocation{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *HttpLocation) GetPaths() PackagePaths {
	if o == nil || o.Paths == nil {
		var ret PackagePaths
		return ret
	}
	return *o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpLocation) GetPathsOk() (*PackagePaths, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *HttpLocation) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given PackagePaths and assigns it to the Paths field.
func (o *HttpLocation) SetPaths(v PackagePaths) {
	o.Paths = &v
}

func (o HttpLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	return json.Marshal(toSerialize)
}

type NullableHttpLocation struct {
	value *HttpLocation
	isSet bool
}

func (v NullableHttpLocation) Get() *HttpLocation {
	return v.value
}

func (v *NullableHttpLocation) Set(val *HttpLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpLocation(val *HttpLocation) *NullableHttpLocation {
	return &NullableHttpLocation{value: val, isSet: true}
}

func (v NullableHttpLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


