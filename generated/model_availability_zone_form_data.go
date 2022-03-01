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

// AvailabilityZoneFormData struct for AvailabilityZoneFormData
type AvailabilityZoneFormData struct {
	// List of availability zones
	AvailabilityZones []AvailabilityZoneData `json:"availabilityZones"`
}

// NewAvailabilityZoneFormData instantiates a new AvailabilityZoneFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityZoneFormData(availabilityZones []AvailabilityZoneData, ) *AvailabilityZoneFormData {
	this := AvailabilityZoneFormData{}
	this.AvailabilityZones = availabilityZones
	return &this
}

// NewAvailabilityZoneFormDataWithDefaults instantiates a new AvailabilityZoneFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityZoneFormDataWithDefaults() *AvailabilityZoneFormData {
	this := AvailabilityZoneFormData{}
	return &this
}

// GetAvailabilityZones returns the AvailabilityZones field value
func (o *AvailabilityZoneFormData) GetAvailabilityZones() []AvailabilityZoneData {
	if o == nil  {
		var ret []AvailabilityZoneData
		return ret
	}

	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value
// and a boolean to check if the value has been set.
func (o *AvailabilityZoneFormData) GetAvailabilityZonesOk() (*[]AvailabilityZoneData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// SetAvailabilityZones sets field value
func (o *AvailabilityZoneFormData) SetAvailabilityZones(v []AvailabilityZoneData) {
	o.AvailabilityZones = v
}

func (o AvailabilityZoneFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["availabilityZones"] = o.AvailabilityZones
	}
	return json.Marshal(toSerialize)
}

type NullableAvailabilityZoneFormData struct {
	value *AvailabilityZoneFormData
	isSet bool
}

func (v NullableAvailabilityZoneFormData) Get() *AvailabilityZoneFormData {
	return v.value
}

func (v *NullableAvailabilityZoneFormData) Set(val *AvailabilityZoneFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityZoneFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityZoneFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityZoneFormData(val *AvailabilityZoneFormData) *NullableAvailabilityZoneFormData {
	return &NullableAvailabilityZoneFormData{value: val, isSet: true}
}

func (v NullableAvailabilityZoneFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityZoneFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


