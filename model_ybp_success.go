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

// YBPSuccess struct for YBPSuccess
type YBPSuccess struct {
	// API response message.
	Message *string `json:"message,omitempty"`
	// API operation status. A value of true indicates the operation was successful.
	Success *bool `json:"success,omitempty"`
}

// NewYBPSuccess instantiates a new YBPSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYBPSuccess() *YBPSuccess {
	this := YBPSuccess{}
	return &this
}

// NewYBPSuccessWithDefaults instantiates a new YBPSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYBPSuccessWithDefaults() *YBPSuccess {
	this := YBPSuccess{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *YBPSuccess) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YBPSuccess) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *YBPSuccess) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *YBPSuccess) SetMessage(v string) {
	o.Message = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *YBPSuccess) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YBPSuccess) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *YBPSuccess) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *YBPSuccess) SetSuccess(v bool) {
	o.Success = &v
}

func (o YBPSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableYBPSuccess struct {
	value *YBPSuccess
	isSet bool
}

func (v NullableYBPSuccess) Get() *YBPSuccess {
	return v.value
}

func (v *NullableYBPSuccess) Set(val *YBPSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableYBPSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableYBPSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYBPSuccess(val *YBPSuccess) *NullableYBPSuccess {
	return &NullableYBPSuccess{value: val, isSet: true}
}

func (v NullableYBPSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYBPSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


