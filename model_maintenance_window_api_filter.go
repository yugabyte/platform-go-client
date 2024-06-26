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

// MaintenanceWindowApiFilter struct for MaintenanceWindowApiFilter
type MaintenanceWindowApiFilter struct {
	States []string `json:"states"`
	Uuids []string `json:"uuids"`
}

// NewMaintenanceWindowApiFilter instantiates a new MaintenanceWindowApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindowApiFilter(states []string, uuids []string) *MaintenanceWindowApiFilter {
	this := MaintenanceWindowApiFilter{}
	this.States = states
	this.Uuids = uuids
	return &this
}

// NewMaintenanceWindowApiFilterWithDefaults instantiates a new MaintenanceWindowApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowApiFilterWithDefaults() *MaintenanceWindowApiFilter {
	this := MaintenanceWindowApiFilter{}
	return &this
}

// GetStates returns the States field value
func (o *MaintenanceWindowApiFilter) GetStates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.States
}

// GetStatesOk returns a tuple with the States field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowApiFilter) GetStatesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.States, true
}

// SetStates sets field value
func (o *MaintenanceWindowApiFilter) SetStates(v []string) {
	o.States = v
}

// GetUuids returns the Uuids field value
func (o *MaintenanceWindowApiFilter) GetUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Uuids
}

// GetUuidsOk returns a tuple with the Uuids field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowApiFilter) GetUuidsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuids, true
}

// SetUuids sets field value
func (o *MaintenanceWindowApiFilter) SetUuids(v []string) {
	o.Uuids = v
}

func (o MaintenanceWindowApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["states"] = o.States
	}
	if true {
		toSerialize["uuids"] = o.Uuids
	}
	return json.Marshal(toSerialize)
}

type NullableMaintenanceWindowApiFilter struct {
	value *MaintenanceWindowApiFilter
	isSet bool
}

func (v NullableMaintenanceWindowApiFilter) Get() *MaintenanceWindowApiFilter {
	return v.value
}

func (v *NullableMaintenanceWindowApiFilter) Set(val *MaintenanceWindowApiFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindowApiFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindowApiFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindowApiFilter(val *MaintenanceWindowApiFilter) *NullableMaintenanceWindowApiFilter {
	return &NullableMaintenanceWindowApiFilter{value: val, isSet: true}
}

func (v NullableMaintenanceWindowApiFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindowApiFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


