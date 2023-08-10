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

// RoleResourceDefinition Role and resource group definition.
type RoleResourceDefinition struct {
	ResourceGroup ResourceGroup `json:"resourceGroup"`
	// UUID of the role to attach resource group to.
	RoleUUID string `json:"roleUUID"`
}

// NewRoleResourceDefinition instantiates a new RoleResourceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleResourceDefinition(resourceGroup ResourceGroup, roleUUID string, ) *RoleResourceDefinition {
	this := RoleResourceDefinition{}
	this.ResourceGroup = resourceGroup
	this.RoleUUID = roleUUID
	return &this
}

// NewRoleResourceDefinitionWithDefaults instantiates a new RoleResourceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleResourceDefinitionWithDefaults() *RoleResourceDefinition {
	this := RoleResourceDefinition{}
	return &this
}

// GetResourceGroup returns the ResourceGroup field value
func (o *RoleResourceDefinition) GetResourceGroup() ResourceGroup {
	if o == nil  {
		var ret ResourceGroup
		return ret
	}

	return o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value
// and a boolean to check if the value has been set.
func (o *RoleResourceDefinition) GetResourceGroupOk() (*ResourceGroup, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceGroup, true
}

// SetResourceGroup sets field value
func (o *RoleResourceDefinition) SetResourceGroup(v ResourceGroup) {
	o.ResourceGroup = v
}

// GetRoleUUID returns the RoleUUID field value
func (o *RoleResourceDefinition) GetRoleUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RoleUUID
}

// GetRoleUUIDOk returns a tuple with the RoleUUID field value
// and a boolean to check if the value has been set.
func (o *RoleResourceDefinition) GetRoleUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleUUID, true
}

// SetRoleUUID sets field value
func (o *RoleResourceDefinition) SetRoleUUID(v string) {
	o.RoleUUID = v
}

func (o RoleResourceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceGroup"] = o.ResourceGroup
	}
	if true {
		toSerialize["roleUUID"] = o.RoleUUID
	}
	return json.Marshal(toSerialize)
}

type NullableRoleResourceDefinition struct {
	value *RoleResourceDefinition
	isSet bool
}

func (v NullableRoleResourceDefinition) Get() *RoleResourceDefinition {
	return v.value
}

func (v *NullableRoleResourceDefinition) Set(val *RoleResourceDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleResourceDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleResourceDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleResourceDefinition(val *RoleResourceDefinition) *NullableRoleResourceDefinition {
	return &NullableRoleResourceDefinition{value: val, isSet: true}
}

func (v NullableRoleResourceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleResourceDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


