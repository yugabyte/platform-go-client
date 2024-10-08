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

// UniverseRestartAllOf struct for UniverseRestartAllOf
type UniverseRestartAllOf struct {
	// Perform a rolling restart of the universe. Otherwise, all nodes will be restarted at the same time.
	RollingRestart *bool `json:"rolling_restart,omitempty"`
	// The method to reboot the node. This is not required for kubernetes universes, as the pods  will get restarted no matter what. \"HARD\" reboots are not supported today.  OS: Restarts the node via the operating system. SERVICE: Restart the YugabyteDB Process only (master, tserver, etc). 
	RestartType *string `json:"restart_type,omitempty"`
}

// NewUniverseRestartAllOf instantiates a new UniverseRestartAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseRestartAllOf() *UniverseRestartAllOf {
	this := UniverseRestartAllOf{}
	var rollingRestart bool = true
	this.RollingRestart = &rollingRestart
	var restartType string = "SERVICE"
	this.RestartType = &restartType
	return &this
}

// NewUniverseRestartAllOfWithDefaults instantiates a new UniverseRestartAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseRestartAllOfWithDefaults() *UniverseRestartAllOf {
	this := UniverseRestartAllOf{}
	var rollingRestart bool = true
	this.RollingRestart = &rollingRestart
	var restartType string = "SERVICE"
	this.RestartType = &restartType
	return &this
}

// GetRollingRestart returns the RollingRestart field value if set, zero value otherwise.
func (o *UniverseRestartAllOf) GetRollingRestart() bool {
	if o == nil || o.RollingRestart == nil {
		var ret bool
		return ret
	}
	return *o.RollingRestart
}

// GetRollingRestartOk returns a tuple with the RollingRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseRestartAllOf) GetRollingRestartOk() (*bool, bool) {
	if o == nil || o.RollingRestart == nil {
		return nil, false
	}
	return o.RollingRestart, true
}

// HasRollingRestart returns a boolean if a field has been set.
func (o *UniverseRestartAllOf) HasRollingRestart() bool {
	if o != nil && o.RollingRestart != nil {
		return true
	}

	return false
}

// SetRollingRestart gets a reference to the given bool and assigns it to the RollingRestart field.
func (o *UniverseRestartAllOf) SetRollingRestart(v bool) {
	o.RollingRestart = &v
}

// GetRestartType returns the RestartType field value if set, zero value otherwise.
func (o *UniverseRestartAllOf) GetRestartType() string {
	if o == nil || o.RestartType == nil {
		var ret string
		return ret
	}
	return *o.RestartType
}

// GetRestartTypeOk returns a tuple with the RestartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseRestartAllOf) GetRestartTypeOk() (*string, bool) {
	if o == nil || o.RestartType == nil {
		return nil, false
	}
	return o.RestartType, true
}

// HasRestartType returns a boolean if a field has been set.
func (o *UniverseRestartAllOf) HasRestartType() bool {
	if o != nil && o.RestartType != nil {
		return true
	}

	return false
}

// SetRestartType gets a reference to the given string and assigns it to the RestartType field.
func (o *UniverseRestartAllOf) SetRestartType(v string) {
	o.RestartType = &v
}

func (o UniverseRestartAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RollingRestart != nil {
		toSerialize["rolling_restart"] = o.RollingRestart
	}
	if o.RestartType != nil {
		toSerialize["restart_type"] = o.RestartType
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseRestartAllOf struct {
	value *UniverseRestartAllOf
	isSet bool
}

func (v NullableUniverseRestartAllOf) Get() *UniverseRestartAllOf {
	return v.value
}

func (v *NullableUniverseRestartAllOf) Set(val *UniverseRestartAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseRestartAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseRestartAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseRestartAllOf(val *UniverseRestartAllOf) *NullableUniverseRestartAllOf {
	return &NullableUniverseRestartAllOf{value: val, isSet: true}
}

func (v NullableUniverseRestartAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseRestartAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


