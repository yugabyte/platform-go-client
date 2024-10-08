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

// XClusterInfo XCluster related states in this universe. Part of UniverseInfo.
type XClusterInfo struct {
	// The value of certs_for_cdc_dir gflag
	SourceRootCertDirPath *string `json:"source_root_cert_dir_path,omitempty"`
	// The source universe's xcluster replication relationships
	SourceXClusterConfigs *[]string `json:"source_x_cluster_configs,omitempty"`
	// The target universe's xcluster replication relationships
	TargetXClusterConfigs *[]string `json:"target_x_cluster_configs,omitempty"`
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

// GetSourceXClusterConfigs returns the SourceXClusterConfigs field value if set, zero value otherwise.
func (o *XClusterInfo) GetSourceXClusterConfigs() []string {
	if o == nil || o.SourceXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.SourceXClusterConfigs
}

// GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterInfo) GetSourceXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.SourceXClusterConfigs == nil {
		return nil, false
	}
	return o.SourceXClusterConfigs, true
}

// HasSourceXClusterConfigs returns a boolean if a field has been set.
func (o *XClusterInfo) HasSourceXClusterConfigs() bool {
	if o != nil && o.SourceXClusterConfigs != nil {
		return true
	}

	return false
}

// SetSourceXClusterConfigs gets a reference to the given []string and assigns it to the SourceXClusterConfigs field.
func (o *XClusterInfo) SetSourceXClusterConfigs(v []string) {
	o.SourceXClusterConfigs = &v
}

// GetTargetXClusterConfigs returns the TargetXClusterConfigs field value if set, zero value otherwise.
func (o *XClusterInfo) GetTargetXClusterConfigs() []string {
	if o == nil || o.TargetXClusterConfigs == nil {
		var ret []string
		return ret
	}
	return *o.TargetXClusterConfigs
}

// GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterInfo) GetTargetXClusterConfigsOk() (*[]string, bool) {
	if o == nil || o.TargetXClusterConfigs == nil {
		return nil, false
	}
	return o.TargetXClusterConfigs, true
}

// HasTargetXClusterConfigs returns a boolean if a field has been set.
func (o *XClusterInfo) HasTargetXClusterConfigs() bool {
	if o != nil && o.TargetXClusterConfigs != nil {
		return true
	}

	return false
}

// SetTargetXClusterConfigs gets a reference to the given []string and assigns it to the TargetXClusterConfigs field.
func (o *XClusterInfo) SetTargetXClusterConfigs(v []string) {
	o.TargetXClusterConfigs = &v
}

func (o XClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceRootCertDirPath != nil {
		toSerialize["source_root_cert_dir_path"] = o.SourceRootCertDirPath
	}
	if o.SourceXClusterConfigs != nil {
		toSerialize["source_x_cluster_configs"] = o.SourceXClusterConfigs
	}
	if o.TargetXClusterConfigs != nil {
		toSerialize["target_x_cluster_configs"] = o.TargetXClusterConfigs
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


