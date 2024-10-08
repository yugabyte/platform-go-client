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

// UniverseEditKubernetesOverrides UniverseEditKubernetesOverrides  Update all kubernetes overrides on the universe. Part of UniverseKubernetesOverridesReq  
type UniverseEditKubernetesOverrides struct {
	// Global kubernetes overrides to apply across the entire universe.
	Overrides *string `json:"overrides,omitempty"`
	// Granular kubernetes overrides per Availability Zone identified by AZ uuid.
	AzOverrides *map[string]string `json:"az_overrides,omitempty"`
}

// NewUniverseEditKubernetesOverrides instantiates a new UniverseEditKubernetesOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseEditKubernetesOverrides() *UniverseEditKubernetesOverrides {
	this := UniverseEditKubernetesOverrides{}
	return &this
}

// NewUniverseEditKubernetesOverridesWithDefaults instantiates a new UniverseEditKubernetesOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseEditKubernetesOverridesWithDefaults() *UniverseEditKubernetesOverrides {
	this := UniverseEditKubernetesOverrides{}
	return &this
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *UniverseEditKubernetesOverrides) GetOverrides() string {
	if o == nil || o.Overrides == nil {
		var ret string
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseEditKubernetesOverrides) GetOverridesOk() (*string, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *UniverseEditKubernetesOverrides) HasOverrides() bool {
	if o != nil && o.Overrides != nil {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given string and assigns it to the Overrides field.
func (o *UniverseEditKubernetesOverrides) SetOverrides(v string) {
	o.Overrides = &v
}

// GetAzOverrides returns the AzOverrides field value if set, zero value otherwise.
func (o *UniverseEditKubernetesOverrides) GetAzOverrides() map[string]string {
	if o == nil || o.AzOverrides == nil {
		var ret map[string]string
		return ret
	}
	return *o.AzOverrides
}

// GetAzOverridesOk returns a tuple with the AzOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseEditKubernetesOverrides) GetAzOverridesOk() (*map[string]string, bool) {
	if o == nil || o.AzOverrides == nil {
		return nil, false
	}
	return o.AzOverrides, true
}

// HasAzOverrides returns a boolean if a field has been set.
func (o *UniverseEditKubernetesOverrides) HasAzOverrides() bool {
	if o != nil && o.AzOverrides != nil {
		return true
	}

	return false
}

// SetAzOverrides gets a reference to the given map[string]string and assigns it to the AzOverrides field.
func (o *UniverseEditKubernetesOverrides) SetAzOverrides(v map[string]string) {
	o.AzOverrides = &v
}

func (o UniverseEditKubernetesOverrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	if o.AzOverrides != nil {
		toSerialize["az_overrides"] = o.AzOverrides
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseEditKubernetesOverrides struct {
	value *UniverseEditKubernetesOverrides
	isSet bool
}

func (v NullableUniverseEditKubernetesOverrides) Get() *UniverseEditKubernetesOverrides {
	return v.value
}

func (v *NullableUniverseEditKubernetesOverrides) Set(val *UniverseEditKubernetesOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseEditKubernetesOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseEditKubernetesOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseEditKubernetesOverrides(val *UniverseEditKubernetesOverrides) *NullableUniverseEditKubernetesOverrides {
	return &NullableUniverseEditKubernetesOverrides{value: val, isSet: true}
}

func (v NullableUniverseEditKubernetesOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseEditKubernetesOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


