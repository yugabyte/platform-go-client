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

// UserRegistrationData User registration data. The API and UI use this to validate form data.
type UserRegistrationData struct {
	// Password confirmation
	ConfirmPassword *string `json:"confirmPassword,omitempty"`
	// Email address
	Email string `json:"email"`
	// User features
	Features *map[string]map[string]interface{} `json:"features,omitempty"`
	// Password
	Password *string `json:"password,omitempty"`
	// <b style=\"color:#ff0000\">Deprecated since YBA version 2.19.3.0.</b> Use field roleResourceDefinitions instead.
	Role *string `json:"role,omitempty"`
	// List of roles and resource groups defined for user.
	RoleResourceDefinitions *[]RoleResourceDefinition `json:"roleResourceDefinitions,omitempty"`
	// User timezone
	Timezone *string `json:"timezone,omitempty"`
}

// NewUserRegistrationData instantiates a new UserRegistrationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRegistrationData(email string) *UserRegistrationData {
	this := UserRegistrationData{}
	this.Email = email
	return &this
}

// NewUserRegistrationDataWithDefaults instantiates a new UserRegistrationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRegistrationDataWithDefaults() *UserRegistrationData {
	this := UserRegistrationData{}
	return &this
}

// GetConfirmPassword returns the ConfirmPassword field value if set, zero value otherwise.
func (o *UserRegistrationData) GetConfirmPassword() string {
	if o == nil || o.ConfirmPassword == nil {
		var ret string
		return ret
	}
	return *o.ConfirmPassword
}

// GetConfirmPasswordOk returns a tuple with the ConfirmPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistrationData) GetConfirmPasswordOk() (*string, bool) {
	if o == nil || o.ConfirmPassword == nil {
		return nil, false
	}
	return o.ConfirmPassword, true
}

// HasConfirmPassword returns a boolean if a field has been set.
func (o *UserRegistrationData) HasConfirmPassword() bool {
	if o != nil && o.ConfirmPassword != nil {
		return true
	}

	return false
}

// SetConfirmPassword gets a reference to the given string and assigns it to the ConfirmPassword field.
func (o *UserRegistrationData) SetConfirmPassword(v string) {
	o.ConfirmPassword = &v
}

// GetEmail returns the Email field value
func (o *UserRegistrationData) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserRegistrationData) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserRegistrationData) SetEmail(v string) {
	o.Email = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *UserRegistrationData) GetFeatures() map[string]map[string]interface{} {
	if o == nil || o.Features == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistrationData) GetFeaturesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *UserRegistrationData) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given map[string]map[string]interface{} and assigns it to the Features field.
func (o *UserRegistrationData) SetFeatures(v map[string]map[string]interface{}) {
	o.Features = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserRegistrationData) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistrationData) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserRegistrationData) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserRegistrationData) SetPassword(v string) {
	o.Password = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserRegistrationData) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistrationData) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserRegistrationData) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UserRegistrationData) SetRole(v string) {
	o.Role = &v
}

// GetRoleResourceDefinitions returns the RoleResourceDefinitions field value if set, zero value otherwise.
func (o *UserRegistrationData) GetRoleResourceDefinitions() []RoleResourceDefinition {
	if o == nil || o.RoleResourceDefinitions == nil {
		var ret []RoleResourceDefinition
		return ret
	}
	return *o.RoleResourceDefinitions
}

// GetRoleResourceDefinitionsOk returns a tuple with the RoleResourceDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistrationData) GetRoleResourceDefinitionsOk() (*[]RoleResourceDefinition, bool) {
	if o == nil || o.RoleResourceDefinitions == nil {
		return nil, false
	}
	return o.RoleResourceDefinitions, true
}

// HasRoleResourceDefinitions returns a boolean if a field has been set.
func (o *UserRegistrationData) HasRoleResourceDefinitions() bool {
	if o != nil && o.RoleResourceDefinitions != nil {
		return true
	}

	return false
}

// SetRoleResourceDefinitions gets a reference to the given []RoleResourceDefinition and assigns it to the RoleResourceDefinitions field.
func (o *UserRegistrationData) SetRoleResourceDefinitions(v []RoleResourceDefinition) {
	o.RoleResourceDefinitions = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UserRegistrationData) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistrationData) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UserRegistrationData) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *UserRegistrationData) SetTimezone(v string) {
	o.Timezone = &v
}

func (o UserRegistrationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfirmPassword != nil {
		toSerialize["confirmPassword"] = o.ConfirmPassword
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.RoleResourceDefinitions != nil {
		toSerialize["roleResourceDefinitions"] = o.RoleResourceDefinitions
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

type NullableUserRegistrationData struct {
	value *UserRegistrationData
	isSet bool
}

func (v NullableUserRegistrationData) Get() *UserRegistrationData {
	return v.value
}

func (v *NullableUserRegistrationData) Set(val *UserRegistrationData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRegistrationData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRegistrationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRegistrationData(val *UserRegistrationData) *NullableUserRegistrationData {
	return &NullableUserRegistrationData{value: val, isSet: true}
}

func (v NullableUserRegistrationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRegistrationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


