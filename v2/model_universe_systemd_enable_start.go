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

// UniverseSystemdEnableStart UniverseSystemdEnableStart  Payload to move the universe to use systemd to manage the YugabyteDB services 
type UniverseSystemdEnableStart struct {
	// Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000.
	SleepAfterMasterRestartMillis *int32 `json:"sleep_after_master_restart_millis,omitempty"`
	// Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000.
	SleepAfterTserverRestartMillis *int32 `json:"sleep_after_tserver_restart_millis,omitempty"`
}

// NewUniverseSystemdEnableStart instantiates a new UniverseSystemdEnableStart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseSystemdEnableStart() *UniverseSystemdEnableStart {
	this := UniverseSystemdEnableStart{}
	var sleepAfterMasterRestartMillis int32 = 180000
	this.SleepAfterMasterRestartMillis = &sleepAfterMasterRestartMillis
	var sleepAfterTserverRestartMillis int32 = 180000
	this.SleepAfterTserverRestartMillis = &sleepAfterTserverRestartMillis
	return &this
}

// NewUniverseSystemdEnableStartWithDefaults instantiates a new UniverseSystemdEnableStart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseSystemdEnableStartWithDefaults() *UniverseSystemdEnableStart {
	this := UniverseSystemdEnableStart{}
	var sleepAfterMasterRestartMillis int32 = 180000
	this.SleepAfterMasterRestartMillis = &sleepAfterMasterRestartMillis
	var sleepAfterTserverRestartMillis int32 = 180000
	this.SleepAfterTserverRestartMillis = &sleepAfterTserverRestartMillis
	return &this
}

// GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field value if set, zero value otherwise.
func (o *UniverseSystemdEnableStart) GetSleepAfterMasterRestartMillis() int32 {
	if o == nil || o.SleepAfterMasterRestartMillis == nil {
		var ret int32
		return ret
	}
	return *o.SleepAfterMasterRestartMillis
}

// GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSystemdEnableStart) GetSleepAfterMasterRestartMillisOk() (*int32, bool) {
	if o == nil || o.SleepAfterMasterRestartMillis == nil {
		return nil, false
	}
	return o.SleepAfterMasterRestartMillis, true
}

// HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.
func (o *UniverseSystemdEnableStart) HasSleepAfterMasterRestartMillis() bool {
	if o != nil && o.SleepAfterMasterRestartMillis != nil {
		return true
	}

	return false
}

// SetSleepAfterMasterRestartMillis gets a reference to the given int32 and assigns it to the SleepAfterMasterRestartMillis field.
func (o *UniverseSystemdEnableStart) SetSleepAfterMasterRestartMillis(v int32) {
	o.SleepAfterMasterRestartMillis = &v
}

// GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field value if set, zero value otherwise.
func (o *UniverseSystemdEnableStart) GetSleepAfterTserverRestartMillis() int32 {
	if o == nil || o.SleepAfterTserverRestartMillis == nil {
		var ret int32
		return ret
	}
	return *o.SleepAfterTserverRestartMillis
}

// GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSystemdEnableStart) GetSleepAfterTserverRestartMillisOk() (*int32, bool) {
	if o == nil || o.SleepAfterTserverRestartMillis == nil {
		return nil, false
	}
	return o.SleepAfterTserverRestartMillis, true
}

// HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.
func (o *UniverseSystemdEnableStart) HasSleepAfterTserverRestartMillis() bool {
	if o != nil && o.SleepAfterTserverRestartMillis != nil {
		return true
	}

	return false
}

// SetSleepAfterTserverRestartMillis gets a reference to the given int32 and assigns it to the SleepAfterTserverRestartMillis field.
func (o *UniverseSystemdEnableStart) SetSleepAfterTserverRestartMillis(v int32) {
	o.SleepAfterTserverRestartMillis = &v
}

func (o UniverseSystemdEnableStart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SleepAfterMasterRestartMillis != nil {
		toSerialize["sleep_after_master_restart_millis"] = o.SleepAfterMasterRestartMillis
	}
	if o.SleepAfterTserverRestartMillis != nil {
		toSerialize["sleep_after_tserver_restart_millis"] = o.SleepAfterTserverRestartMillis
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseSystemdEnableStart struct {
	value *UniverseSystemdEnableStart
	isSet bool
}

func (v NullableUniverseSystemdEnableStart) Get() *UniverseSystemdEnableStart {
	return v.value
}

func (v *NullableUniverseSystemdEnableStart) Set(val *UniverseSystemdEnableStart) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseSystemdEnableStart) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseSystemdEnableStart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseSystemdEnableStart(val *UniverseSystemdEnableStart) *NullableUniverseSystemdEnableStart {
	return &NullableUniverseSystemdEnableStart{value: val, isSet: true}
}

func (v NullableUniverseSystemdEnableStart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseSystemdEnableStart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


