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

// S3Location struct for S3Location
type S3Location struct {
	Paths *PackagePaths `json:"paths,omitempty"`
}

// NewS3Location instantiates a new S3Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3Location() *S3Location {
	this := S3Location{}
	return &this
}

// NewS3LocationWithDefaults instantiates a new S3Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3LocationWithDefaults() *S3Location {
	this := S3Location{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *S3Location) GetPaths() PackagePaths {
	if o == nil || o.Paths == nil {
		var ret PackagePaths
		return ret
	}
	return *o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Location) GetPathsOk() (*PackagePaths, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *S3Location) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given PackagePaths and assigns it to the Paths field.
func (o *S3Location) SetPaths(v PackagePaths) {
	o.Paths = &v
}

func (o S3Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	return json.Marshal(toSerialize)
}

type NullableS3Location struct {
	value *S3Location
	isSet bool
}

func (v NullableS3Location) Get() *S3Location {
	return v.value
}

func (v *NullableS3Location) Set(val *S3Location) {
	v.value = val
	v.isSet = true
}

func (v NullableS3Location) IsSet() bool {
	return v.isSet
}

func (v *NullableS3Location) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3Location(val *S3Location) *NullableS3Location {
	return &NullableS3Location{value: val, isSet: true}
}

func (v NullableS3Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3Location) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


