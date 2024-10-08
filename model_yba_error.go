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

// YBAError struct for YBAError
type YBAError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

// NewYBAError instantiates a new YBAError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYBAError(code string, message string) *YBAError {
	this := YBAError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewYBAErrorWithDefaults instantiates a new YBAError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYBAErrorWithDefaults() *YBAError {
	this := YBAError{}
	return &this
}

// GetCode returns the Code field value
func (o *YBAError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *YBAError) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *YBAError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *YBAError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *YBAError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *YBAError) SetMessage(v string) {
	o.Message = v
}

func (o YBAError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableYBAError struct {
	value *YBAError
	isSet bool
}

func (v NullableYBAError) Get() *YBAError {
	return v.value
}

func (v *NullableYBAError) Set(val *YBAError) {
	v.value = val
	v.isSet = true
}

func (v NullableYBAError) IsSet() bool {
	return v.isSet
}

func (v *NullableYBAError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYBAError(val *YBAError) *NullableYBAError {
	return &NullableYBAError{value: val, isSet: true}
}

func (v NullableYBAError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYBAError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


