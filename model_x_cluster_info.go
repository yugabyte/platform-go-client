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

// XClusterInfo struct for XClusterInfo
type XClusterInfo struct {
	// The value of certs_for_cdc_dir gflag
	SourceRootCertDirPath *string `json:"sourceRootCertDirPath,omitempty"`
}

// NewXClusterInfo instantiates a new XClusterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXClusterInfo() *XClusterInfo {
	this := XClusterInfo{}
	return &this
}

// NewXClusterInfoWithDefaults instantiates a new XClusterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXClusterInfoWithDefaults() *XClusterInfo {
	this := XClusterInfo{}
	return &this
}

// GetSourceRootCertDirPath returns the SourceRootCertDirPath field value if set, zero value otherwise.
func (o *XClusterInfo) GetSourceRootCertDirPath() string {
	if o == nil || o.SourceRootCertDirPath == nil {
		var ret string
		return ret
	}
	return *o.SourceRootCertDirPath
}

// GetSourceRootCertDirPathOk returns a tuple with the SourceRootCertDirPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterInfo) GetSourceRootCertDirPathOk() (*string, bool) {
	if o == nil || o.SourceRootCertDirPath == nil {
		return nil, false
	}
	return o.SourceRootCertDirPath, true
}

// HasSourceRootCertDirPath returns a boolean if a field has been set.
func (o *XClusterInfo) HasSourceRootCertDirPath() bool {
	if o != nil && o.SourceRootCertDirPath != nil {
		return true
	}

	return false
}

// SetSourceRootCertDirPath gets a reference to the given string and assigns it to the SourceRootCertDirPath field.
func (o *XClusterInfo) SetSourceRootCertDirPath(v string) {
	o.SourceRootCertDirPath = &v
}

func (o XClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceRootCertDirPath != nil {
		toSerialize["sourceRootCertDirPath"] = o.SourceRootCertDirPath
	}
	return json.Marshal(toSerialize)
}

type NullableXClusterInfo struct {
	value *XClusterInfo
	isSet bool
}

func (v NullableXClusterInfo) Get() *XClusterInfo {
	return v.value
}

func (v *NullableXClusterInfo) Set(val *XClusterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableXClusterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableXClusterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXClusterInfo(val *XClusterInfo) *NullableXClusterInfo {
	return &NullableXClusterInfo{value: val, isSet: true}
}

func (v NullableXClusterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXClusterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


