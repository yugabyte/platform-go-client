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

// GCPCloudMonitoringConfigAllOf GCPCloudMonitoringConfig Config
type GCPCloudMonitoringConfigAllOf struct {
	// Project ID
	Project *string `json:"project,omitempty"`
}

// NewGCPCloudMonitoringConfigAllOf instantiates a new GCPCloudMonitoringConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPCloudMonitoringConfigAllOf() *GCPCloudMonitoringConfigAllOf {
	this := GCPCloudMonitoringConfigAllOf{}
	return &this
}

// NewGCPCloudMonitoringConfigAllOfWithDefaults instantiates a new GCPCloudMonitoringConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPCloudMonitoringConfigAllOfWithDefaults() *GCPCloudMonitoringConfigAllOf {
	this := GCPCloudMonitoringConfigAllOf{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *GCPCloudMonitoringConfigAllOf) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudMonitoringConfigAllOf) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *GCPCloudMonitoringConfigAllOf) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *GCPCloudMonitoringConfigAllOf) SetProject(v string) {
	o.Project = &v
}

func (o GCPCloudMonitoringConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableGCPCloudMonitoringConfigAllOf struct {
	value *GCPCloudMonitoringConfigAllOf
	isSet bool
}

func (v NullableGCPCloudMonitoringConfigAllOf) Get() *GCPCloudMonitoringConfigAllOf {
	return v.value
}

func (v *NullableGCPCloudMonitoringConfigAllOf) Set(val *GCPCloudMonitoringConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPCloudMonitoringConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPCloudMonitoringConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPCloudMonitoringConfigAllOf(val *GCPCloudMonitoringConfigAllOf) *NullableGCPCloudMonitoringConfigAllOf {
	return &NullableGCPCloudMonitoringConfigAllOf{value: val, isSet: true}
}

func (v NullableGCPCloudMonitoringConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPCloudMonitoringConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


