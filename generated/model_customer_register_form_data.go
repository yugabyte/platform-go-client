/*
 * Yugabyte Platform APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yb_platform_client

import (
	"encoding/json"
)

// CustomerRegisterFormData struct for CustomerRegisterFormData
type CustomerRegisterFormData struct {
	Code string `json:"code"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

// NewCustomerRegisterFormData instantiates a new CustomerRegisterFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerRegisterFormData(code string, email string, name string, password string, ) *CustomerRegisterFormData {
	this := CustomerRegisterFormData{}
	this.Code = code
	this.Email = email
	this.Name = name
	this.Password = password
	return &this
}

// NewCustomerRegisterFormDataWithDefaults instantiates a new CustomerRegisterFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerRegisterFormDataWithDefaults() *CustomerRegisterFormData {
	this := CustomerRegisterFormData{}
	return &this
}

// GetCode returns the Code field value
func (o *CustomerRegisterFormData) GetCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CustomerRegisterFormData) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CustomerRegisterFormData) SetCode(v string) {
	o.Code = v
}

// GetEmail returns the Email field value
func (o *CustomerRegisterFormData) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CustomerRegisterFormData) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CustomerRegisterFormData) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *CustomerRegisterFormData) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerRegisterFormData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerRegisterFormData) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value
func (o *CustomerRegisterFormData) GetPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CustomerRegisterFormData) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CustomerRegisterFormData) SetPassword(v string) {
	o.Password = v
}

func (o CustomerRegisterFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerRegisterFormData struct {
	value *CustomerRegisterFormData
	isSet bool
}

func (v NullableCustomerRegisterFormData) Get() *CustomerRegisterFormData {
	return v.value
}

func (v *NullableCustomerRegisterFormData) Set(val *CustomerRegisterFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerRegisterFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerRegisterFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerRegisterFormData(val *CustomerRegisterFormData) *NullableCustomerRegisterFormData {
	return &NullableCustomerRegisterFormData{value: val, isSet: true}
}

func (v NullableCustomerRegisterFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerRegisterFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


