/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// Role struct for Role
type Role struct {
	// Role create time
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	// Customer UUID
	CustomerUUID *string `json:"customerUUID,omitempty"`
	// Role description
	Description *string `json:"description,omitempty"`
	// Role name
	Name *string `json:"name,omitempty"`
	PermissionDetails *PermissionDetails `json:"permissionDetails,omitempty"`
	// Type of the role
	RoleType *string `json:"roleType,omitempty"`
	// Role UUID
	RoleUUID *string `json:"roleUUID,omitempty"`
	// Role last updated time
	UpdatedOn *time.Time `json:"updatedOn,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole() *Role {
	this := Role{}
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Role) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Role) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Role) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetCustomerUUID returns the CustomerUUID field value if set, zero value otherwise.
func (o *Role) GetCustomerUUID() string {
	if o == nil || o.CustomerUUID == nil {
		var ret string
		return ret
	}
	return *o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetCustomerUUIDOk() (*string, bool) {
	if o == nil || o.CustomerUUID == nil {
		return nil, false
	}
	return o.CustomerUUID, true
}

// HasCustomerUUID returns a boolean if a field has been set.
func (o *Role) HasCustomerUUID() bool {
	if o != nil && o.CustomerUUID != nil {
		return true
	}

	return false
}

// SetCustomerUUID gets a reference to the given string and assigns it to the CustomerUUID field.
func (o *Role) SetCustomerUUID(v string) {
	o.CustomerUUID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Role) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Role) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Role) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Role) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Role) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Role) SetName(v string) {
	o.Name = &v
}

// GetPermissionDetails returns the PermissionDetails field value if set, zero value otherwise.
func (o *Role) GetPermissionDetails() PermissionDetails {
	if o == nil || o.PermissionDetails == nil {
		var ret PermissionDetails
		return ret
	}
	return *o.PermissionDetails
}

// GetPermissionDetailsOk returns a tuple with the PermissionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetPermissionDetailsOk() (*PermissionDetails, bool) {
	if o == nil || o.PermissionDetails == nil {
		return nil, false
	}
	return o.PermissionDetails, true
}

// HasPermissionDetails returns a boolean if a field has been set.
func (o *Role) HasPermissionDetails() bool {
	if o != nil && o.PermissionDetails != nil {
		return true
	}

	return false
}

// SetPermissionDetails gets a reference to the given PermissionDetails and assigns it to the PermissionDetails field.
func (o *Role) SetPermissionDetails(v PermissionDetails) {
	o.PermissionDetails = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *Role) GetRoleType() string {
	if o == nil || o.RoleType == nil {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleTypeOk() (*string, bool) {
	if o == nil || o.RoleType == nil {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *Role) HasRoleType() bool {
	if o != nil && o.RoleType != nil {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *Role) SetRoleType(v string) {
	o.RoleType = &v
}

// GetRoleUUID returns the RoleUUID field value if set, zero value otherwise.
func (o *Role) GetRoleUUID() string {
	if o == nil || o.RoleUUID == nil {
		var ret string
		return ret
	}
	return *o.RoleUUID
}

// GetRoleUUIDOk returns a tuple with the RoleUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleUUIDOk() (*string, bool) {
	if o == nil || o.RoleUUID == nil {
		return nil, false
	}
	return o.RoleUUID, true
}

// HasRoleUUID returns a boolean if a field has been set.
func (o *Role) HasRoleUUID() bool {
	if o != nil && o.RoleUUID != nil {
		return true
	}

	return false
}

// SetRoleUUID gets a reference to the given string and assigns it to the RoleUUID field.
func (o *Role) SetRoleUUID(v string) {
	o.RoleUUID = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *Role) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *Role) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *Role) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.CustomerUUID != nil {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PermissionDetails != nil {
		toSerialize["permissionDetails"] = o.PermissionDetails
	}
	if o.RoleType != nil {
		toSerialize["roleType"] = o.RoleType
	}
	if o.RoleUUID != nil {
		toSerialize["roleUUID"] = o.RoleUUID
	}
	if o.UpdatedOn != nil {
		toSerialize["updatedOn"] = o.UpdatedOn
	}
	return json.Marshal(toSerialize)
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


