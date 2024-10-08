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

// UniverseSoftwareUpgradePrecheckResp UniverseSoftwareUpgradePrecheckResp  Response to a YugabyteDB software upgrade precheck on a Universe. Returns if a finalize is  required. Part of UniverseSoftwareUpgradePrecheckResponse. 
type UniverseSoftwareUpgradePrecheckResp struct {
	// If the upgrade requires a finalize step. If true, the user must call the finalize  endpoint to complete the upgrade. 
	FinalizeRequired bool `json:"finalize_required"`
}

// NewUniverseSoftwareUpgradePrecheckResp instantiates a new UniverseSoftwareUpgradePrecheckResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseSoftwareUpgradePrecheckResp(finalizeRequired bool) *UniverseSoftwareUpgradePrecheckResp {
	this := UniverseSoftwareUpgradePrecheckResp{}
	this.FinalizeRequired = finalizeRequired
	return &this
}

// NewUniverseSoftwareUpgradePrecheckRespWithDefaults instantiates a new UniverseSoftwareUpgradePrecheckResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseSoftwareUpgradePrecheckRespWithDefaults() *UniverseSoftwareUpgradePrecheckResp {
	this := UniverseSoftwareUpgradePrecheckResp{}
	return &this
}

// GetFinalizeRequired returns the FinalizeRequired field value
func (o *UniverseSoftwareUpgradePrecheckResp) GetFinalizeRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FinalizeRequired
}

// GetFinalizeRequiredOk returns a tuple with the FinalizeRequired field value
// and a boolean to check if the value has been set.
func (o *UniverseSoftwareUpgradePrecheckResp) GetFinalizeRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FinalizeRequired, true
}

// SetFinalizeRequired sets field value
func (o *UniverseSoftwareUpgradePrecheckResp) SetFinalizeRequired(v bool) {
	o.FinalizeRequired = v
}

func (o UniverseSoftwareUpgradePrecheckResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["finalize_required"] = o.FinalizeRequired
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseSoftwareUpgradePrecheckResp struct {
	value *UniverseSoftwareUpgradePrecheckResp
	isSet bool
}

func (v NullableUniverseSoftwareUpgradePrecheckResp) Get() *UniverseSoftwareUpgradePrecheckResp {
	return v.value
}

func (v *NullableUniverseSoftwareUpgradePrecheckResp) Set(val *UniverseSoftwareUpgradePrecheckResp) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseSoftwareUpgradePrecheckResp) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseSoftwareUpgradePrecheckResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseSoftwareUpgradePrecheckResp(val *UniverseSoftwareUpgradePrecheckResp) *NullableUniverseSoftwareUpgradePrecheckResp {
	return &NullableUniverseSoftwareUpgradePrecheckResp{value: val, isSet: true}
}

func (v NullableUniverseSoftwareUpgradePrecheckResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseSoftwareUpgradePrecheckResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


