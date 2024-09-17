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
)

// CustomServerCertInfo struct for CustomServerCertInfo
type CustomServerCertInfo struct {
	ServerCert string `json:"serverCert"`
	ServerKey string `json:"serverKey"`
}

// NewCustomServerCertInfo instantiates a new CustomServerCertInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomServerCertInfo(serverCert string, serverKey string) *CustomServerCertInfo {
	this := CustomServerCertInfo{}
	this.ServerCert = serverCert
	this.ServerKey = serverKey
	return &this
}

// NewCustomServerCertInfoWithDefaults instantiates a new CustomServerCertInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomServerCertInfoWithDefaults() *CustomServerCertInfo {
	this := CustomServerCertInfo{}
	return &this
}

// GetServerCert returns the ServerCert field value
func (o *CustomServerCertInfo) GetServerCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerCert
}

// GetServerCertOk returns a tuple with the ServerCert field value
// and a boolean to check if the value has been set.
func (o *CustomServerCertInfo) GetServerCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServerCert, true
}

// SetServerCert sets field value
func (o *CustomServerCertInfo) SetServerCert(v string) {
	o.ServerCert = v
}

// GetServerKey returns the ServerKey field value
func (o *CustomServerCertInfo) GetServerKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerKey
}

// GetServerKeyOk returns a tuple with the ServerKey field value
// and a boolean to check if the value has been set.
func (o *CustomServerCertInfo) GetServerKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServerKey, true
}

// SetServerKey sets field value
func (o *CustomServerCertInfo) SetServerKey(v string) {
	o.ServerKey = v
}

func (o CustomServerCertInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serverCert"] = o.ServerCert
	}
	if true {
		toSerialize["serverKey"] = o.ServerKey
	}
	return json.Marshal(toSerialize)
}

type NullableCustomServerCertInfo struct {
	value *CustomServerCertInfo
	isSet bool
}

func (v NullableCustomServerCertInfo) Get() *CustomServerCertInfo {
	return v.value
}

func (v *NullableCustomServerCertInfo) Set(val *CustomServerCertInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomServerCertInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomServerCertInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomServerCertInfo(val *CustomServerCertInfo) *NullableCustomServerCertInfo {
	return &NullableCustomServerCertInfo{value: val, isSet: true}
}

func (v NullableCustomServerCertInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomServerCertInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


