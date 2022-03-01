/*
 * Yugabyte Platform APIs
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

// CertificateDetails struct for CertificateDetails
type CertificateDetails struct {
	YugabytedbCrt *string `json:"yugabytedb.crt,omitempty"`
	YugabytedbKey *string `json:"yugabytedb.key,omitempty"`
}

// NewCertificateDetails instantiates a new CertificateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateDetails() *CertificateDetails {
	this := CertificateDetails{}
	return &this
}

// NewCertificateDetailsWithDefaults instantiates a new CertificateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateDetailsWithDefaults() *CertificateDetails {
	this := CertificateDetails{}
	return &this
}

// GetYugabytedbCrt returns the YugabytedbCrt field value if set, zero value otherwise.
func (o *CertificateDetails) GetYugabytedbCrt() string {
	if o == nil || o.YugabytedbCrt == nil {
		var ret string
		return ret
	}
	return *o.YugabytedbCrt
}

// GetYugabytedbCrtOk returns a tuple with the YugabytedbCrt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetYugabytedbCrtOk() (*string, bool) {
	if o == nil || o.YugabytedbCrt == nil {
		return nil, false
	}
	return o.YugabytedbCrt, true
}

// HasYugabytedbCrt returns a boolean if a field has been set.
func (o *CertificateDetails) HasYugabytedbCrt() bool {
	if o != nil && o.YugabytedbCrt != nil {
		return true
	}

	return false
}

// SetYugabytedbCrt gets a reference to the given string and assigns it to the YugabytedbCrt field.
func (o *CertificateDetails) SetYugabytedbCrt(v string) {
	o.YugabytedbCrt = &v
}

// GetYugabytedbKey returns the YugabytedbKey field value if set, zero value otherwise.
func (o *CertificateDetails) GetYugabytedbKey() string {
	if o == nil || o.YugabytedbKey == nil {
		var ret string
		return ret
	}
	return *o.YugabytedbKey
}

// GetYugabytedbKeyOk returns a tuple with the YugabytedbKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetYugabytedbKeyOk() (*string, bool) {
	if o == nil || o.YugabytedbKey == nil {
		return nil, false
	}
	return o.YugabytedbKey, true
}

// HasYugabytedbKey returns a boolean if a field has been set.
func (o *CertificateDetails) HasYugabytedbKey() bool {
	if o != nil && o.YugabytedbKey != nil {
		return true
	}

	return false
}

// SetYugabytedbKey gets a reference to the given string and assigns it to the YugabytedbKey field.
func (o *CertificateDetails) SetYugabytedbKey(v string) {
	o.YugabytedbKey = &v
}

func (o CertificateDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.YugabytedbCrt != nil {
		toSerialize["yugabytedb.crt"] = o.YugabytedbCrt
	}
	if o.YugabytedbKey != nil {
		toSerialize["yugabytedb.key"] = o.YugabytedbKey
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateDetails struct {
	value *CertificateDetails
	isSet bool
}

func (v NullableCertificateDetails) Get() *CertificateDetails {
	return v.value
}

func (v *NullableCertificateDetails) Set(val *CertificateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateDetails(val *CertificateDetails) *NullableCertificateDetails {
	return &NullableCertificateDetails{value: val, isSet: true}
}

func (v NullableCertificateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


