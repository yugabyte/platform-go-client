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

// PerfAdvisorSettingsWithDefaults struct for PerfAdvisorSettingsWithDefaults
type PerfAdvisorSettingsWithDefaults struct {
	DefaultSettings *PerfAdvisorSettingsFormData `json:"defaultSettings,omitempty"`
	UniverseSettings *PerfAdvisorSettingsFormData `json:"universeSettings,omitempty"`
}

// NewPerfAdvisorSettingsWithDefaults instantiates a new PerfAdvisorSettingsWithDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfAdvisorSettingsWithDefaults() *PerfAdvisorSettingsWithDefaults {
	this := PerfAdvisorSettingsWithDefaults{}
	return &this
}

// NewPerfAdvisorSettingsWithDefaultsWithDefaults instantiates a new PerfAdvisorSettingsWithDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfAdvisorSettingsWithDefaultsWithDefaults() *PerfAdvisorSettingsWithDefaults {
	this := PerfAdvisorSettingsWithDefaults{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *PerfAdvisorSettingsWithDefaults) GetDefaultSettings() PerfAdvisorSettingsFormData {
	if o == nil || o.DefaultSettings == nil {
		var ret PerfAdvisorSettingsFormData
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfAdvisorSettingsWithDefaults) GetDefaultSettingsOk() (*PerfAdvisorSettingsFormData, bool) {
	if o == nil || o.DefaultSettings == nil {
		return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *PerfAdvisorSettingsWithDefaults) HasDefaultSettings() bool {
	if o != nil && o.DefaultSettings != nil {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given PerfAdvisorSettingsFormData and assigns it to the DefaultSettings field.
func (o *PerfAdvisorSettingsWithDefaults) SetDefaultSettings(v PerfAdvisorSettingsFormData) {
	o.DefaultSettings = &v
}

// GetUniverseSettings returns the UniverseSettings field value if set, zero value otherwise.
func (o *PerfAdvisorSettingsWithDefaults) GetUniverseSettings() PerfAdvisorSettingsFormData {
	if o == nil || o.UniverseSettings == nil {
		var ret PerfAdvisorSettingsFormData
		return ret
	}
	return *o.UniverseSettings
}

// GetUniverseSettingsOk returns a tuple with the UniverseSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfAdvisorSettingsWithDefaults) GetUniverseSettingsOk() (*PerfAdvisorSettingsFormData, bool) {
	if o == nil || o.UniverseSettings == nil {
		return nil, false
	}
	return o.UniverseSettings, true
}

// HasUniverseSettings returns a boolean if a field has been set.
func (o *PerfAdvisorSettingsWithDefaults) HasUniverseSettings() bool {
	if o != nil && o.UniverseSettings != nil {
		return true
	}

	return false
}

// SetUniverseSettings gets a reference to the given PerfAdvisorSettingsFormData and assigns it to the UniverseSettings field.
func (o *PerfAdvisorSettingsWithDefaults) SetUniverseSettings(v PerfAdvisorSettingsFormData) {
	o.UniverseSettings = &v
}

func (o PerfAdvisorSettingsWithDefaults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultSettings != nil {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if o.UniverseSettings != nil {
		toSerialize["universeSettings"] = o.UniverseSettings
	}
	return json.Marshal(toSerialize)
}

type NullablePerfAdvisorSettingsWithDefaults struct {
	value *PerfAdvisorSettingsWithDefaults
	isSet bool
}

func (v NullablePerfAdvisorSettingsWithDefaults) Get() *PerfAdvisorSettingsWithDefaults {
	return v.value
}

func (v *NullablePerfAdvisorSettingsWithDefaults) Set(val *PerfAdvisorSettingsWithDefaults) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfAdvisorSettingsWithDefaults) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfAdvisorSettingsWithDefaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfAdvisorSettingsWithDefaults(val *PerfAdvisorSettingsWithDefaults) *NullablePerfAdvisorSettingsWithDefaults {
	return &NullablePerfAdvisorSettingsWithDefaults{value: val, isSet: true}
}

func (v NullablePerfAdvisorSettingsWithDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfAdvisorSettingsWithDefaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


