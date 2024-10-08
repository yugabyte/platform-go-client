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

// RoleResourceDefinition Defines the association of Role to Resource Groups. Part of AuthGroupToRolesMapping.
type RoleResourceDefinition struct {
	// UUID of the role to attach resource group to.
	RoleUuid string `json:"role_uuid"`
	ResourceGroup *ResourceGroup `json:"resource_group,omitempty"`
}

// NewRoleResourceDefinition instantiates a new RoleResourceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleResourceDefinition(roleUuid string) *RoleResourceDefinition {
	this := RoleResourceDefinition{}
	this.RoleUuid = roleUuid
	return &this
}

// NewRoleResourceDefinitionWithDefaults instantiates a new RoleResourceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleResourceDefinitionWithDefaults() *RoleResourceDefinition {
	this := RoleResourceDefinition{}
	return &this
}

// GetRoleUuid returns the RoleUuid field value
func (o *RoleResourceDefinition) GetRoleUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleUuid
}

// GetRoleUuidOk returns a tuple with the RoleUuid field value
// and a boolean to check if the value has been set.
func (o *RoleResourceDefinition) GetRoleUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleUuid, true
}

// SetRoleUuid sets field value
func (o *RoleResourceDefinition) SetRoleUuid(v string) {
	o.RoleUuid = v
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise.
func (o *RoleResourceDefinition) GetResourceGroup() ResourceGroup {
	if o == nil || o.ResourceGroup == nil {
		var ret ResourceGroup
		return ret
	}
	return *o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleResourceDefinition) GetResourceGroupOk() (*ResourceGroup, bool) {
	if o == nil || o.ResourceGroup == nil {
		return nil, false
	}
	return o.ResourceGroup, true
}

// HasResourceGroup returns a boolean if a field has been set.
func (o *RoleResourceDefinition) HasResourceGroup() bool {
	if o != nil && o.ResourceGroup != nil {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given ResourceGroup and assigns it to the ResourceGroup field.
func (o *RoleResourceDefinition) SetResourceGroup(v ResourceGroup) {
	o.ResourceGroup = &v
}

func (o RoleResourceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role_uuid"] = o.RoleUuid
	}
	if o.ResourceGroup != nil {
		toSerialize["resource_group"] = o.ResourceGroup
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


