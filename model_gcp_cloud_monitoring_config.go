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

// GCPCloudMonitoringConfig struct for GCPCloudMonitoringConfig
type GCPCloudMonitoringConfig struct {
	TelemetryProviderConfig
	// Project
	Project *string `json:"project,omitempty"`
}

// NewGCPCloudMonitoringConfig instantiates a new GCPCloudMonitoringConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPCloudMonitoringConfig() *GCPCloudMonitoringConfig {
	this := GCPCloudMonitoringConfig{}
	return &this
}

// NewGCPCloudMonitoringConfigWithDefaults instantiates a new GCPCloudMonitoringConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPCloudMonitoringConfigWithDefaults() *GCPCloudMonitoringConfig {
	this := GCPCloudMonitoringConfig{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *GCPCloudMonitoringConfig) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudMonitoringConfig) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *GCPCloudMonitoringConfig) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *GCPCloudMonitoringConfig) SetProject(v string) {
	o.Project = &v
}

func (o GCPCloudMonitoringConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTelemetryProviderConfig, errTelemetryProviderConfig := json.Marshal(o.TelemetryProviderConfig)
	if errTelemetryProviderConfig != nil {
		return []byte{}, errTelemetryProviderConfig
	}
	errTelemetryProviderConfig = json.Unmarshal([]byte(serializedTelemetryProviderConfig), &toSerialize)
	if errTelemetryProviderConfig != nil {
		return []byte{}, errTelemetryProviderConfig
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableGCPCloudMonitoringConfig struct {
	value *GCPCloudMonitoringConfig
	isSet bool
}

func (v NullableGCPCloudMonitoringConfig) Get() *GCPCloudMonitoringConfig {
	return v.value
}

func (v *NullableGCPCloudMonitoringConfig) Set(val *GCPCloudMonitoringConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPCloudMonitoringConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPCloudMonitoringConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPCloudMonitoringConfig(val *GCPCloudMonitoringConfig) *NullableGCPCloudMonitoringConfig {
	return &NullableGCPCloudMonitoringConfig{value: val, isSet: true}
}

func (v NullableGCPCloudMonitoringConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPCloudMonitoringConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


