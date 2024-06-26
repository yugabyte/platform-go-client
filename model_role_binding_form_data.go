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

// RoleBindingFormData Role Binding form metadata
type RoleBindingFormData struct {
	// List of roles and resource groups defined.
	RoleResourceDefinitions []RoleResourceDefinition `json:"roleResourceDefinitions"`
}

// NewRoleBindingFormData instantiates a new RoleBindingFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBindingFormData(roleResourceDefinitions []RoleResourceDefinition) *RoleBindingFormData {
	this := RoleBindingFormData{}
	this.RoleResourceDefinitions = roleResourceDefinitions
	return &this
}

// NewRoleBindingFormDataWithDefaults instantiates a new RoleBindingFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBindingFormDataWithDefaults() *RoleBindingFormData {
	this := RoleBindingFormData{}
	return &this
}

// GetRoleResourceDefinitions returns the RoleResourceDefinitions field value
func (o *RoleBindingFormData) GetRoleResourceDefinitions() []RoleResourceDefinition {
	if o == nil {
		var ret []RoleResourceDefinition
		return ret
	}

	return o.RoleResourceDefinitions
}

// GetRoleResourceDefinitionsOk returns a tuple with the RoleResourceDefinitions field value
// and a boolean to check if the value has been set.
func (o *RoleBindingFormData) GetRoleResourceDefinitionsOk() (*[]RoleResourceDefinition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleResourceDefinitions, true
}

// SetRoleResourceDefinitions sets field value
func (o *RoleBindingFormData) SetRoleResourceDefinitions(v []RoleResourceDefinition) {
	o.RoleResourceDefinitions = v
}

func (o RoleBindingFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["roleResourceDefinitions"] = o.RoleResourceDefinitions
	}
	return json.Marshal(toSerialize)
}

type NullableRoleBindingFormData struct {
	value *RoleBindingFormData
	isSet bool
}

func (v NullableRoleBindingFormData) Get() *RoleBindingFormData {
	return v.value
}

func (v *NullableRoleBindingFormData) Set(val *RoleBindingFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBindingFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBindingFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBindingFormData(val *RoleBindingFormData) *NullableRoleBindingFormData {
	return &NullableRoleBindingFormData{value: val, isSet: true}
}

func (v NullableRoleBindingFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBindingFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


