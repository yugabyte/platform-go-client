/*
 * Yugabyte Platform APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yb_platform_client

import (
	"encoding/json"
)

// ConfigEntry Configuration entry
type ConfigEntry struct {
	// Is this configuration inherited?
	Inherited *bool `json:"inherited,omitempty"`
	// Configuration key
	Key *string `json:"key,omitempty"`
	// Configuration value
	Value *string `json:"value,omitempty"`
}

// NewConfigEntry instantiates a new ConfigEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigEntry() *ConfigEntry {
	this := ConfigEntry{}
	return &this
}

// NewConfigEntryWithDefaults instantiates a new ConfigEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigEntryWithDefaults() *ConfigEntry {
	this := ConfigEntry{}
	return &this
}

// GetInherited returns the Inherited field value if set, zero value otherwise.
func (o *ConfigEntry) GetInherited() bool {
	if o == nil || o.Inherited == nil {
		var ret bool
		return ret
	}
	return *o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetInheritedOk() (*bool, bool) {
	if o == nil || o.Inherited == nil {
		return nil, false
	}
	return o.Inherited, true
}

// HasInherited returns a boolean if a field has been set.
func (o *ConfigEntry) HasInherited() bool {
	if o != nil && o.Inherited != nil {
		return true
	}

	return false
}

// SetInherited gets a reference to the given bool and assigns it to the Inherited field.
func (o *ConfigEntry) SetInherited(v bool) {
	o.Inherited = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ConfigEntry) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ConfigEntry) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ConfigEntry) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConfigEntry) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigEntry) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigEntry) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConfigEntry) SetValue(v string) {
	o.Value = &v
}

func (o ConfigEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inherited != nil {
		toSerialize["inherited"] = o.Inherited
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableConfigEntry struct {
	value *ConfigEntry
	isSet bool
}

func (v NullableConfigEntry) Get() *ConfigEntry {
	return v.value
}

func (v *NullableConfigEntry) Set(val *ConfigEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigEntry(val *ConfigEntry) *NullableConfigEntry {
	return &NullableConfigEntry{value: val, isSet: true}
}

func (v NullableConfigEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


