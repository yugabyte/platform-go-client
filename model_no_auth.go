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

// NoAuth struct for NoAuth
type NoAuth struct {
	HTTPAuthInformation
}

// NewNoAuth instantiates a new NoAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoAuth() *NoAuth {
	this := NoAuth{}
	return &this
}

// NewNoAuthWithDefaults instantiates a new NoAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoAuthWithDefaults() *NoAuth {
	this := NoAuth{}
	return &this
}

func (o NoAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedHTTPAuthInformation, errHTTPAuthInformation := json.Marshal(o.HTTPAuthInformation)
	if errHTTPAuthInformation != nil {
		return []byte{}, errHTTPAuthInformation
	}
	errHTTPAuthInformation = json.Unmarshal([]byte(serializedHTTPAuthInformation), &toSerialize)
	if errHTTPAuthInformation != nil {
		return []byte{}, errHTTPAuthInformation
	}
	return json.Marshal(toSerialize)
}

type NullableNoAuth struct {
	value *NoAuth
	isSet bool
}

func (v NullableNoAuth) Get() *NoAuth {
	return v.value
}

func (v *NullableNoAuth) Set(val *NoAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableNoAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableNoAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoAuth(val *NoAuth) *NullableNoAuth {
	return &NullableNoAuth{value: val, isSet: true}
}

func (v NullableNoAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


