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

// AllowedUniverseTasksResp struct for AllowedUniverseTasksResp
type AllowedUniverseTasksResp struct {
	Restricted *bool `json:"restricted,omitempty"`
	TaskIds *[]string `json:"taskIds,omitempty"`
}

// NewAllowedUniverseTasksResp instantiates a new AllowedUniverseTasksResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedUniverseTasksResp() *AllowedUniverseTasksResp {
	this := AllowedUniverseTasksResp{}
	return &this
}

// NewAllowedUniverseTasksRespWithDefaults instantiates a new AllowedUniverseTasksResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedUniverseTasksRespWithDefaults() *AllowedUniverseTasksResp {
	this := AllowedUniverseTasksResp{}
	return &this
}

// GetRestricted returns the Restricted field value if set, zero value otherwise.
func (o *AllowedUniverseTasksResp) GetRestricted() bool {
	if o == nil || o.Restricted == nil {
		var ret bool
		return ret
	}
	return *o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedUniverseTasksResp) GetRestrictedOk() (*bool, bool) {
	if o == nil || o.Restricted == nil {
		return nil, false
	}
	return o.Restricted, true
}

// HasRestricted returns a boolean if a field has been set.
func (o *AllowedUniverseTasksResp) HasRestricted() bool {
	if o != nil && o.Restricted != nil {
		return true
	}

	return false
}

// SetRestricted gets a reference to the given bool and assigns it to the Restricted field.
func (o *AllowedUniverseTasksResp) SetRestricted(v bool) {
	o.Restricted = &v
}

// GetTaskIds returns the TaskIds field value if set, zero value otherwise.
func (o *AllowedUniverseTasksResp) GetTaskIds() []string {
	if o == nil || o.TaskIds == nil {
		var ret []string
		return ret
	}
	return *o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedUniverseTasksResp) GetTaskIdsOk() (*[]string, bool) {
	if o == nil || o.TaskIds == nil {
		return nil, false
	}
	return o.TaskIds, true
}

// HasTaskIds returns a boolean if a field has been set.
func (o *AllowedUniverseTasksResp) HasTaskIds() bool {
	if o != nil && o.TaskIds != nil {
		return true
	}

	return false
}

// SetTaskIds gets a reference to the given []string and assigns it to the TaskIds field.
func (o *AllowedUniverseTasksResp) SetTaskIds(v []string) {
	o.TaskIds = &v
}

func (o AllowedUniverseTasksResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Restricted != nil {
		toSerialize["restricted"] = o.Restricted
	}
	if o.TaskIds != nil {
		toSerialize["taskIds"] = o.TaskIds
	}
	return json.Marshal(toSerialize)
}

type NullableAllowedUniverseTasksResp struct {
	value *AllowedUniverseTasksResp
	isSet bool
}

func (v NullableAllowedUniverseTasksResp) Get() *AllowedUniverseTasksResp {
	return v.value
}

func (v *NullableAllowedUniverseTasksResp) Set(val *AllowedUniverseTasksResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedUniverseTasksResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedUniverseTasksResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedUniverseTasksResp(val *AllowedUniverseTasksResp) *NullableAllowedUniverseTasksResp {
	return &NullableAllowedUniverseTasksResp{value: val, isSet: true}
}

func (v NullableAllowedUniverseTasksResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedUniverseTasksResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

