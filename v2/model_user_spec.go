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

// UserSpec A user associated with a customer
type UserSpec struct {
	// User email address
	Email string `json:"email"`
	// User role
	Role *string `json:"role,omitempty"`
}

// NewUserSpec instantiates a new UserSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSpec(email string) *UserSpec {
	this := UserSpec{}
	this.Email = email
	return &this
}

// NewUserSpecWithDefaults instantiates a new UserSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSpecWithDefaults() *UserSpec {
	this := UserSpec{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserSpec) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserSpec) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserSpec) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserSpec) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSpec) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserSpec) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UserSpec) SetRole(v string) {
	o.Role = &v
}

func (o UserSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableUserSpec struct {
	value *UserSpec
	isSet bool
}

func (v NullableUserSpec) Get() *UserSpec {
	return v.value
}

func (v *NullableUserSpec) Set(val *UserSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSpec(val *UserSpec) *NullableUserSpec {
	return &NullableUserSpec{value: val, isSet: true}
}

func (v NullableUserSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


