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

// SpecificGFlags GFlags for current cluster
type SpecificGFlags struct {
	InheritFromPrimary *bool `json:"inheritFromPrimary,omitempty"`
	// Overrides for gflags per availability zone
	PerAZ *map[string]PerProcessFlags `json:"perAZ,omitempty"`
	PerProcessFlags *PerProcessFlags `json:"perProcessFlags,omitempty"`
}

// NewSpecificGFlags instantiates a new SpecificGFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecificGFlags() *SpecificGFlags {
	this := SpecificGFlags{}
	return &this
}

// NewSpecificGFlagsWithDefaults instantiates a new SpecificGFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecificGFlagsWithDefaults() *SpecificGFlags {
	this := SpecificGFlags{}
	return &this
}

// GetInheritFromPrimary returns the InheritFromPrimary field value if set, zero value otherwise.
func (o *SpecificGFlags) GetInheritFromPrimary() bool {
	if o == nil || o.InheritFromPrimary == nil {
		var ret bool
		return ret
	}
	return *o.InheritFromPrimary
}

// GetInheritFromPrimaryOk returns a tuple with the InheritFromPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificGFlags) GetInheritFromPrimaryOk() (*bool, bool) {
	if o == nil || o.InheritFromPrimary == nil {
		return nil, false
	}
	return o.InheritFromPrimary, true
}

// HasInheritFromPrimary returns a boolean if a field has been set.
func (o *SpecificGFlags) HasInheritFromPrimary() bool {
	if o != nil && o.InheritFromPrimary != nil {
		return true
	}

	return false
}

// SetInheritFromPrimary gets a reference to the given bool and assigns it to the InheritFromPrimary field.
func (o *SpecificGFlags) SetInheritFromPrimary(v bool) {
	o.InheritFromPrimary = &v
}

// GetPerAZ returns the PerAZ field value if set, zero value otherwise.
func (o *SpecificGFlags) GetPerAZ() map[string]PerProcessFlags {
	if o == nil || o.PerAZ == nil {
		var ret map[string]PerProcessFlags
		return ret
	}
	return *o.PerAZ
}

// GetPerAZOk returns a tuple with the PerAZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificGFlags) GetPerAZOk() (*map[string]PerProcessFlags, bool) {
	if o == nil || o.PerAZ == nil {
		return nil, false
	}
	return o.PerAZ, true
}

// HasPerAZ returns a boolean if a field has been set.
func (o *SpecificGFlags) HasPerAZ() bool {
	if o != nil && o.PerAZ != nil {
		return true
	}

	return false
}

// SetPerAZ gets a reference to the given map[string]PerProcessFlags and assigns it to the PerAZ field.
func (o *SpecificGFlags) SetPerAZ(v map[string]PerProcessFlags) {
	o.PerAZ = &v
}

// GetPerProcessFlags returns the PerProcessFlags field value if set, zero value otherwise.
func (o *SpecificGFlags) GetPerProcessFlags() PerProcessFlags {
	if o == nil || o.PerProcessFlags == nil {
		var ret PerProcessFlags
		return ret
	}
	return *o.PerProcessFlags
}

// GetPerProcessFlagsOk returns a tuple with the PerProcessFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificGFlags) GetPerProcessFlagsOk() (*PerProcessFlags, bool) {
	if o == nil || o.PerProcessFlags == nil {
		return nil, false
	}
	return o.PerProcessFlags, true
}

// HasPerProcessFlags returns a boolean if a field has been set.
func (o *SpecificGFlags) HasPerProcessFlags() bool {
	if o != nil && o.PerProcessFlags != nil {
		return true
	}

	return false
}

// SetPerProcessFlags gets a reference to the given PerProcessFlags and assigns it to the PerProcessFlags field.
func (o *SpecificGFlags) SetPerProcessFlags(v PerProcessFlags) {
	o.PerProcessFlags = &v
}

func (o SpecificGFlags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InheritFromPrimary != nil {
		toSerialize["inheritFromPrimary"] = o.InheritFromPrimary
	}
	if o.PerAZ != nil {
		toSerialize["perAZ"] = o.PerAZ
	}
	if o.PerProcessFlags != nil {
		toSerialize["perProcessFlags"] = o.PerProcessFlags
	}
	return json.Marshal(toSerialize)
}

type NullableSpecificGFlags struct {
	value *SpecificGFlags
	isSet bool
}

func (v NullableSpecificGFlags) Get() *SpecificGFlags {
	return v.value
}

func (v *NullableSpecificGFlags) Set(val *SpecificGFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecificGFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecificGFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecificGFlags(val *SpecificGFlags) *NullableSpecificGFlags {
	return &NullableSpecificGFlags{value: val, isSet: true}
}

func (v NullableSpecificGFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecificGFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

