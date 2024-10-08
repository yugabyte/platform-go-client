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

// UniverseThirdPartySoftwareUpgradeStart UniverseThirdPartyUpgradeStart  Payload to start a third party software upgrade on a Universe. Part of  UniverseThirdPartyUpgradeReq 
type UniverseThirdPartySoftwareUpgradeStart struct {
	// Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000.
	SleepAfterMasterRestartMillis *int32 `json:"sleep_after_master_restart_millis,omitempty"`
	// Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000.
	SleepAfterTserverRestartMillis *int32 `json:"sleep_after_tserver_restart_millis,omitempty"`
	// force all thirdparty softwares to be upgraded.
	ForceAll *bool `json:"force_all,omitempty"`
}

// NewUniverseThirdPartySoftwareUpgradeStart instantiates a new UniverseThirdPartySoftwareUpgradeStart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseThirdPartySoftwareUpgradeStart() *UniverseThirdPartySoftwareUpgradeStart {
	this := UniverseThirdPartySoftwareUpgradeStart{}
	var sleepAfterMasterRestartMillis int32 = 180000
	this.SleepAfterMasterRestartMillis = &sleepAfterMasterRestartMillis
	var sleepAfterTserverRestartMillis int32 = 180000
	this.SleepAfterTserverRestartMillis = &sleepAfterTserverRestartMillis
	var forceAll bool = false
	this.ForceAll = &forceAll
	return &this
}

// NewUniverseThirdPartySoftwareUpgradeStartWithDefaults instantiates a new UniverseThirdPartySoftwareUpgradeStart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseThirdPartySoftwareUpgradeStartWithDefaults() *UniverseThirdPartySoftwareUpgradeStart {
	this := UniverseThirdPartySoftwareUpgradeStart{}
	var sleepAfterMasterRestartMillis int32 = 180000
	this.SleepAfterMasterRestartMillis = &sleepAfterMasterRestartMillis
	var sleepAfterTserverRestartMillis int32 = 180000
	this.SleepAfterTserverRestartMillis = &sleepAfterTserverRestartMillis
	var forceAll bool = false
	this.ForceAll = &forceAll
	return &this
}

// GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field value if set, zero value otherwise.
func (o *UniverseThirdPartySoftwareUpgradeStart) GetSleepAfterMasterRestartMillis() int32 {
	if o == nil || o.SleepAfterMasterRestartMillis == nil {
		var ret int32
		return ret
	}
	return *o.SleepAfterMasterRestartMillis
}

// GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseThirdPartySoftwareUpgradeStart) GetSleepAfterMasterRestartMillisOk() (*int32, bool) {
	if o == nil || o.SleepAfterMasterRestartMillis == nil {
		return nil, false
	}
	return o.SleepAfterMasterRestartMillis, true
}

// HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.
func (o *UniverseThirdPartySoftwareUpgradeStart) HasSleepAfterMasterRestartMillis() bool {
	if o != nil && o.SleepAfterMasterRestartMillis != nil {
		return true
	}

	return false
}

// SetSleepAfterMasterRestartMillis gets a reference to the given int32 and assigns it to the SleepAfterMasterRestartMillis field.
func (o *UniverseThirdPartySoftwareUpgradeStart) SetSleepAfterMasterRestartMillis(v int32) {
	o.SleepAfterMasterRestartMillis = &v
}

// GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field value if set, zero value otherwise.
func (o *UniverseThirdPartySoftwareUpgradeStart) GetSleepAfterTserverRestartMillis() int32 {
	if o == nil || o.SleepAfterTserverRestartMillis == nil {
		var ret int32
		return ret
	}
	return *o.SleepAfterTserverRestartMillis
}

// GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseThirdPartySoftwareUpgradeStart) GetSleepAfterTserverRestartMillisOk() (*int32, bool) {
	if o == nil || o.SleepAfterTserverRestartMillis == nil {
		return nil, false
	}
	return o.SleepAfterTserverRestartMillis, true
}

// HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.
func (o *UniverseThirdPartySoftwareUpgradeStart) HasSleepAfterTserverRestartMillis() bool {
	if o != nil && o.SleepAfterTserverRestartMillis != nil {
		return true
	}

	return false
}

// SetSleepAfterTserverRestartMillis gets a reference to the given int32 and assigns it to the SleepAfterTserverRestartMillis field.
func (o *UniverseThirdPartySoftwareUpgradeStart) SetSleepAfterTserverRestartMillis(v int32) {
	o.SleepAfterTserverRestartMillis = &v
}

// GetForceAll returns the ForceAll field value if set, zero value otherwise.
func (o *UniverseThirdPartySoftwareUpgradeStart) GetForceAll() bool {
	if o == nil || o.ForceAll == nil {
		var ret bool
		return ret
	}
	return *o.ForceAll
}

// GetForceAllOk returns a tuple with the ForceAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseThirdPartySoftwareUpgradeStart) GetForceAllOk() (*bool, bool) {
	if o == nil || o.ForceAll == nil {
		return nil, false
	}
	return o.ForceAll, true
}

// HasForceAll returns a boolean if a field has been set.
func (o *UniverseThirdPartySoftwareUpgradeStart) HasForceAll() bool {
	if o != nil && o.ForceAll != nil {
		return true
	}

	return false
}

// SetForceAll gets a reference to the given bool and assigns it to the ForceAll field.
func (o *UniverseThirdPartySoftwareUpgradeStart) SetForceAll(v bool) {
	o.ForceAll = &v
}

func (o UniverseThirdPartySoftwareUpgradeStart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SleepAfterMasterRestartMillis != nil {
		toSerialize["sleep_after_master_restart_millis"] = o.SleepAfterMasterRestartMillis
	}
	if o.SleepAfterTserverRestartMillis != nil {
		toSerialize["sleep_after_tserver_restart_millis"] = o.SleepAfterTserverRestartMillis
	}
	if o.ForceAll != nil {
		toSerialize["force_all"] = o.ForceAll
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseThirdPartySoftwareUpgradeStart struct {
	value *UniverseThirdPartySoftwareUpgradeStart
	isSet bool
}

func (v NullableUniverseThirdPartySoftwareUpgradeStart) Get() *UniverseThirdPartySoftwareUpgradeStart {
	return v.value
}

func (v *NullableUniverseThirdPartySoftwareUpgradeStart) Set(val *UniverseThirdPartySoftwareUpgradeStart) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseThirdPartySoftwareUpgradeStart) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseThirdPartySoftwareUpgradeStart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseThirdPartySoftwareUpgradeStart(val *UniverseThirdPartySoftwareUpgradeStart) *NullableUniverseThirdPartySoftwareUpgradeStart {
	return &NullableUniverseThirdPartySoftwareUpgradeStart{value: val, isSet: true}
}

func (v NullableUniverseThirdPartySoftwareUpgradeStart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseThirdPartySoftwareUpgradeStart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


