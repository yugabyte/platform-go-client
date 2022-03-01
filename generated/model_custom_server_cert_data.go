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

// CustomServerCertData struct for CustomServerCertData
type CustomServerCertData struct {
	ServerCertContent string `json:"serverCertContent"`
	ServerKeyContent string `json:"serverKeyContent"`
}

// NewCustomServerCertData instantiates a new CustomServerCertData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomServerCertData(serverCertContent string, serverKeyContent string, ) *CustomServerCertData {
	this := CustomServerCertData{}
	this.ServerCertContent = serverCertContent
	this.ServerKeyContent = serverKeyContent
	return &this
}

// NewCustomServerCertDataWithDefaults instantiates a new CustomServerCertData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomServerCertDataWithDefaults() *CustomServerCertData {
	this := CustomServerCertData{}
	return &this
}

// GetServerCertContent returns the ServerCertContent field value
func (o *CustomServerCertData) GetServerCertContent() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServerCertContent
}

// GetServerCertContentOk returns a tuple with the ServerCertContent field value
// and a boolean to check if the value has been set.
func (o *CustomServerCertData) GetServerCertContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServerCertContent, true
}

// SetServerCertContent sets field value
func (o *CustomServerCertData) SetServerCertContent(v string) {
	o.ServerCertContent = v
}

// GetServerKeyContent returns the ServerKeyContent field value
func (o *CustomServerCertData) GetServerKeyContent() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServerKeyContent
}

// GetServerKeyContentOk returns a tuple with the ServerKeyContent field value
// and a boolean to check if the value has been set.
func (o *CustomServerCertData) GetServerKeyContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServerKeyContent, true
}

// SetServerKeyContent sets field value
func (o *CustomServerCertData) SetServerKeyContent(v string) {
	o.ServerKeyContent = v
}

func (o CustomServerCertData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serverCertContent"] = o.ServerCertContent
	}
	if true {
		toSerialize["serverKeyContent"] = o.ServerKeyContent
	}
	return json.Marshal(toSerialize)
}

type NullableCustomServerCertData struct {
	value *CustomServerCertData
	isSet bool
}

func (v NullableCustomServerCertData) Get() *CustomServerCertData {
	return v.value
}

func (v *NullableCustomServerCertData) Set(val *CustomServerCertData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomServerCertData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomServerCertData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomServerCertData(val *CustomServerCertData) *NullableCustomServerCertData {
	return &NullableCustomServerCertData{value: val, isSet: true}
}

func (v NullableCustomServerCertData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomServerCertData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


