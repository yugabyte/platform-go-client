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

// RoleFormData Role form metadata
type RoleFormData struct {
	// Description of the role to be created
	Description string `json:"description"`
	// Name of the role to be created
	Name string `json:"name"`
	// List of permissions given to the role
	PermissionList []Permission `json:"permissionList"`
}

// NewRoleFormData instantiates a new RoleFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleFormData(description string, name string, permissionList []Permission) *RoleFormData {
	this := RoleFormData{}
	this.Description = description
	this.Name = name
	this.PermissionList = permissionList
	return &this
}

// NewRoleFormDataWithDefaults instantiates a new RoleFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleFormDataWithDefaults() *RoleFormData {
	this := RoleFormData{}
	return &this
}

// GetDescription returns the Description field value
func (o *RoleFormData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RoleFormData) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RoleFormData) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *RoleFormData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleFormData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleFormData) SetName(v string) {
	o.Name = v
}

// GetPermissionList returns the PermissionList field value
func (o *RoleFormData) GetPermissionList() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}

	return o.PermissionList
}

// GetPermissionListOk returns a tuple with the PermissionList field value
// and a boolean to check if the value has been set.
func (o *RoleFormData) GetPermissionListOk() (*[]Permission, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PermissionList, true
}

// SetPermissionList sets field value
func (o *RoleFormData) SetPermissionList(v []Permission) {
	o.PermissionList = v
}

func (o RoleFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["permissionList"] = o.PermissionList
	}
	return json.Marshal(toSerialize)
}

type NullableRoleFormData struct {
	value *RoleFormData
	isSet bool
}

func (v NullableRoleFormData) Get() *RoleFormData {
	return v.value
}

func (v *NullableRoleFormData) Set(val *RoleFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleFormData(val *RoleFormData) *NullableRoleFormData {
	return &NullableRoleFormData{value: val, isSet: true}
}

func (v NullableRoleFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


