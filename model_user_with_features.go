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
	"time"
)

// UserWithFeatures A user with set of features, associated with a customer
type UserWithFeatures struct {
	// UI session token creation date
	AuthTokenIssueDate *time.Time `json:"authTokenIssueDate,omitempty"`
	// User creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Customer UUID
	CustomerUUID *string `json:"customerUUID,omitempty"`
	// User email address
	Email string `json:"email"`
	// True if the user is the primary user
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// LDAP Specified Role
	LdapSpecifiedRole *bool `json:"ldapSpecifiedRole,omitempty"`
	// User role
	Role *string `json:"role,omitempty"`
	// User timezone
	Timezone *string `json:"timezone,omitempty"`
	// User Type
	UserType *string `json:"userType,omitempty"`
	// User UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewUserWithFeatures instantiates a new UserWithFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWithFeatures(email string, ) *UserWithFeatures {
	this := UserWithFeatures{}
	this.Email = email
	return &this
}

// NewUserWithFeaturesWithDefaults instantiates a new UserWithFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithFeaturesWithDefaults() *UserWithFeatures {
	this := UserWithFeatures{}
	return &this
}

// GetAuthTokenIssueDate returns the AuthTokenIssueDate field value if set, zero value otherwise.
func (o *UserWithFeatures) GetAuthTokenIssueDate() time.Time {
	if o == nil || o.AuthTokenIssueDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AuthTokenIssueDate
}

// GetAuthTokenIssueDateOk returns a tuple with the AuthTokenIssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetAuthTokenIssueDateOk() (*time.Time, bool) {
	if o == nil || o.AuthTokenIssueDate == nil {
		return nil, false
	}
	return o.AuthTokenIssueDate, true
}

// HasAuthTokenIssueDate returns a boolean if a field has been set.
func (o *UserWithFeatures) HasAuthTokenIssueDate() bool {
	if o != nil && o.AuthTokenIssueDate != nil {
		return true
	}

	return false
}

// SetAuthTokenIssueDate gets a reference to the given time.Time and assigns it to the AuthTokenIssueDate field.
func (o *UserWithFeatures) SetAuthTokenIssueDate(v time.Time) {
	o.AuthTokenIssueDate = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *UserWithFeatures) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *UserWithFeatures) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *UserWithFeatures) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetCustomerUUID returns the CustomerUUID field value if set, zero value otherwise.
func (o *UserWithFeatures) GetCustomerUUID() string {
	if o == nil || o.CustomerUUID == nil {
		var ret string
		return ret
	}
	return *o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetCustomerUUIDOk() (*string, bool) {
	if o == nil || o.CustomerUUID == nil {
		return nil, false
	}
	return o.CustomerUUID, true
}

// HasCustomerUUID returns a boolean if a field has been set.
func (o *UserWithFeatures) HasCustomerUUID() bool {
	if o != nil && o.CustomerUUID != nil {
		return true
	}

	return false
}

// SetCustomerUUID gets a reference to the given string and assigns it to the CustomerUUID field.
func (o *UserWithFeatures) SetCustomerUUID(v string) {
	o.CustomerUUID = &v
}

// GetEmail returns the Email field value
func (o *UserWithFeatures) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserWithFeatures) SetEmail(v string) {
	o.Email = v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *UserWithFeatures) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *UserWithFeatures) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *UserWithFeatures) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetLdapSpecifiedRole returns the LdapSpecifiedRole field value if set, zero value otherwise.
func (o *UserWithFeatures) GetLdapSpecifiedRole() bool {
	if o == nil || o.LdapSpecifiedRole == nil {
		var ret bool
		return ret
	}
	return *o.LdapSpecifiedRole
}

// GetLdapSpecifiedRoleOk returns a tuple with the LdapSpecifiedRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetLdapSpecifiedRoleOk() (*bool, bool) {
	if o == nil || o.LdapSpecifiedRole == nil {
		return nil, false
	}
	return o.LdapSpecifiedRole, true
}

// HasLdapSpecifiedRole returns a boolean if a field has been set.
func (o *UserWithFeatures) HasLdapSpecifiedRole() bool {
	if o != nil && o.LdapSpecifiedRole != nil {
		return true
	}

	return false
}

// SetLdapSpecifiedRole gets a reference to the given bool and assigns it to the LdapSpecifiedRole field.
func (o *UserWithFeatures) SetLdapSpecifiedRole(v bool) {
	o.LdapSpecifiedRole = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserWithFeatures) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserWithFeatures) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UserWithFeatures) SetRole(v string) {
	o.Role = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UserWithFeatures) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UserWithFeatures) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *UserWithFeatures) SetTimezone(v string) {
	o.Timezone = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *UserWithFeatures) GetUserType() string {
	if o == nil || o.UserType == nil {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetUserTypeOk() (*string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserWithFeatures) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *UserWithFeatures) SetUserType(v string) {
	o.UserType = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UserWithFeatures) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithFeatures) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UserWithFeatures) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UserWithFeatures) SetUuid(v string) {
	o.Uuid = &v
}

func (o UserWithFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthTokenIssueDate != nil {
		toSerialize["authTokenIssueDate"] = o.AuthTokenIssueDate
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.CustomerUUID != nil {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.IsPrimary != nil {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if o.LdapSpecifiedRole != nil {
		toSerialize["ldapSpecifiedRole"] = o.LdapSpecifiedRole
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.UserType != nil {
		toSerialize["userType"] = o.UserType
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableUserWithFeatures struct {
	value *UserWithFeatures
	isSet bool
}

func (v NullableUserWithFeatures) Get() *UserWithFeatures {
	return v.value
}

func (v *NullableUserWithFeatures) Set(val *UserWithFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWithFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWithFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWithFeatures(val *UserWithFeatures) *NullableUserWithFeatures {
	return &NullableUserWithFeatures{value: val, isSet: true}
}

func (v NullableUserWithFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWithFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


