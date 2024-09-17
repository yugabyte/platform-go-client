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

// UniverseThirdPartySoftwareUpgradeStartAllOf struct for UniverseThirdPartySoftwareUpgradeStartAllOf
type UniverseThirdPartySoftwareUpgradeStartAllOf struct {
	// force all thirdparty softwares to be upgraded.
	ForceAll *bool `json:"force_all,omitempty"`
}

// NewUniverseThirdPartySoftwareUpgradeStartAllOf instantiates a new UniverseThirdPartySoftwareUpgradeStartAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseThirdPartySoftwareUpgradeStartAllOf() *UniverseThirdPartySoftwareUpgradeStartAllOf {
	this := UniverseThirdPartySoftwareUpgradeStartAllOf{}
	var forceAll bool = false
	this.ForceAll = &forceAll
	return &this
}

// NewUniverseThirdPartySoftwareUpgradeStartAllOfWithDefaults instantiates a new UniverseThirdPartySoftwareUpgradeStartAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseThirdPartySoftwareUpgradeStartAllOfWithDefaults() *UniverseThirdPartySoftwareUpgradeStartAllOf {
	this := UniverseThirdPartySoftwareUpgradeStartAllOf{}
	var forceAll bool = false
	this.ForceAll = &forceAll
	return &this
}

// GetForceAll returns the ForceAll field value if set, zero value otherwise.
func (o *UniverseThirdPartySoftwareUpgradeStartAllOf) GetForceAll() bool {
	if o == nil || o.ForceAll == nil {
		var ret bool
		return ret
	}
	return *o.ForceAll
}

// GetForceAllOk returns a tuple with the ForceAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseThirdPartySoftwareUpgradeStartAllOf) GetForceAllOk() (*bool, bool) {
	if o == nil || o.ForceAll == nil {
		return nil, false
	}
	return o.ForceAll, true
}

// HasForceAll returns a boolean if a field has been set.
func (o *UniverseThirdPartySoftwareUpgradeStartAllOf) HasForceAll() bool {
	if o != nil && o.ForceAll != nil {
		return true
	}

	return false
}

// SetForceAll gets a reference to the given bool and assigns it to the ForceAll field.
func (o *UniverseThirdPartySoftwareUpgradeStartAllOf) SetForceAll(v bool) {
	o.ForceAll = &v
}

func (o UniverseThirdPartySoftwareUpgradeStartAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceAll != nil {
		toSerialize["force_all"] = o.ForceAll
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseThirdPartySoftwareUpgradeStartAllOf struct {
	value *UniverseThirdPartySoftwareUpgradeStartAllOf
	isSet bool
}

func (v NullableUniverseThirdPartySoftwareUpgradeStartAllOf) Get() *UniverseThirdPartySoftwareUpgradeStartAllOf {
	return v.value
}

func (v *NullableUniverseThirdPartySoftwareUpgradeStartAllOf) Set(val *UniverseThirdPartySoftwareUpgradeStartAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseThirdPartySoftwareUpgradeStartAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseThirdPartySoftwareUpgradeStartAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseThirdPartySoftwareUpgradeStartAllOf(val *UniverseThirdPartySoftwareUpgradeStartAllOf) *NullableUniverseThirdPartySoftwareUpgradeStartAllOf {
	return &NullableUniverseThirdPartySoftwareUpgradeStartAllOf{value: val, isSet: true}
}

func (v NullableUniverseThirdPartySoftwareUpgradeStartAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseThirdPartySoftwareUpgradeStartAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


