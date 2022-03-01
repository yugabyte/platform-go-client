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

// ScopedConfig Scoped configuration
type ScopedConfig struct {
	// List of configurations
	ConfigEntries *[]ConfigEntry `json:"configEntries,omitempty"`
	// Mutability of the scope. Only super admin users can change global scope; other scopes are customer-mutable.
	MutableScope *bool `json:"mutableScope,omitempty"`
	// Scope type
	Type *string `json:"type,omitempty"`
	// Scope UIID
	Uuid *string `json:"uuid,omitempty"`
}

// NewScopedConfig instantiates a new ScopedConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopedConfig() *ScopedConfig {
	this := ScopedConfig{}
	return &this
}

// NewScopedConfigWithDefaults instantiates a new ScopedConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopedConfigWithDefaults() *ScopedConfig {
	this := ScopedConfig{}
	return &this
}

// GetConfigEntries returns the ConfigEntries field value if set, zero value otherwise.
func (o *ScopedConfig) GetConfigEntries() []ConfigEntry {
	if o == nil || o.ConfigEntries == nil {
		var ret []ConfigEntry
		return ret
	}
	return *o.ConfigEntries
}

// GetConfigEntriesOk returns a tuple with the ConfigEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopedConfig) GetConfigEntriesOk() (*[]ConfigEntry, bool) {
	if o == nil || o.ConfigEntries == nil {
		return nil, false
	}
	return o.ConfigEntries, true
}

// HasConfigEntries returns a boolean if a field has been set.
func (o *ScopedConfig) HasConfigEntries() bool {
	if o != nil && o.ConfigEntries != nil {
		return true
	}

	return false
}

// SetConfigEntries gets a reference to the given []ConfigEntry and assigns it to the ConfigEntries field.
func (o *ScopedConfig) SetConfigEntries(v []ConfigEntry) {
	o.ConfigEntries = &v
}

// GetMutableScope returns the MutableScope field value if set, zero value otherwise.
func (o *ScopedConfig) GetMutableScope() bool {
	if o == nil || o.MutableScope == nil {
		var ret bool
		return ret
	}
	return *o.MutableScope
}

// GetMutableScopeOk returns a tuple with the MutableScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopedConfig) GetMutableScopeOk() (*bool, bool) {
	if o == nil || o.MutableScope == nil {
		return nil, false
	}
	return o.MutableScope, true
}

// HasMutableScope returns a boolean if a field has been set.
func (o *ScopedConfig) HasMutableScope() bool {
	if o != nil && o.MutableScope != nil {
		return true
	}

	return false
}

// SetMutableScope gets a reference to the given bool and assigns it to the MutableScope field.
func (o *ScopedConfig) SetMutableScope(v bool) {
	o.MutableScope = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScopedConfig) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopedConfig) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScopedConfig) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScopedConfig) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ScopedConfig) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopedConfig) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ScopedConfig) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ScopedConfig) SetUuid(v string) {
	o.Uuid = &v
}

func (o ScopedConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigEntries != nil {
		toSerialize["configEntries"] = o.ConfigEntries
	}
	if o.MutableScope != nil {
		toSerialize["mutableScope"] = o.MutableScope
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableScopedConfig struct {
	value *ScopedConfig
	isSet bool
}

func (v NullableScopedConfig) Get() *ScopedConfig {
	return v.value
}

func (v *NullableScopedConfig) Set(val *ScopedConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableScopedConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableScopedConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopedConfig(val *ScopedConfig) *NullableScopedConfig {
	return &NullableScopedConfig{value: val, isSet: true}
}

func (v NullableScopedConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopedConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


