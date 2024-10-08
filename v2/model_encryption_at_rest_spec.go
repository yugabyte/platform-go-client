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

// EncryptionAtRestSpec Encryption At Rest specification for the Universe. Part of UniverseSpec.
type EncryptionAtRestSpec struct {
	// The KMS Configuration associated with the encryption keys being used on this Universe
	KmsConfigUuid *string `json:"kms_config_uuid,omitempty"`
}

// NewEncryptionAtRestSpec instantiates a new EncryptionAtRestSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionAtRestSpec() *EncryptionAtRestSpec {
	this := EncryptionAtRestSpec{}
	return &this
}

// NewEncryptionAtRestSpecWithDefaults instantiates a new EncryptionAtRestSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionAtRestSpecWithDefaults() *EncryptionAtRestSpec {
	this := EncryptionAtRestSpec{}
	return &this
}

// GetKmsConfigUuid returns the KmsConfigUuid field value if set, zero value otherwise.
func (o *EncryptionAtRestSpec) GetKmsConfigUuid() string {
	if o == nil || o.KmsConfigUuid == nil {
		var ret string
		return ret
	}
	return *o.KmsConfigUuid
}

// GetKmsConfigUuidOk returns a tuple with the KmsConfigUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionAtRestSpec) GetKmsConfigUuidOk() (*string, bool) {
	if o == nil || o.KmsConfigUuid == nil {
		return nil, false
	}
	return o.KmsConfigUuid, true
}

// HasKmsConfigUuid returns a boolean if a field has been set.
func (o *EncryptionAtRestSpec) HasKmsConfigUuid() bool {
	if o != nil && o.KmsConfigUuid != nil {
		return true
	}

	return false
}

// SetKmsConfigUuid gets a reference to the given string and assigns it to the KmsConfigUuid field.
func (o *EncryptionAtRestSpec) SetKmsConfigUuid(v string) {
	o.KmsConfigUuid = &v
}

func (o EncryptionAtRestSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KmsConfigUuid != nil {
		toSerialize["kms_config_uuid"] = o.KmsConfigUuid
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptionAtRestSpec struct {
	value *EncryptionAtRestSpec
	isSet bool
}

func (v NullableEncryptionAtRestSpec) Get() *EncryptionAtRestSpec {
	return v.value
}

func (v *NullableEncryptionAtRestSpec) Set(val *EncryptionAtRestSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionAtRestSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionAtRestSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionAtRestSpec(val *EncryptionAtRestSpec) *NullableEncryptionAtRestSpec {
	return &NullableEncryptionAtRestSpec{value: val, isSet: true}
}

func (v NullableEncryptionAtRestSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionAtRestSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


