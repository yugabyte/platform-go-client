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
	"time"
)

// AccessKey Access key for the cloud provider. This helps to authenticate the user and get access to the provider.
type AccessKey struct {
	// Creation date of key
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Expiration date of key
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	IdKey AccessKeyId `json:"idKey"`
	KeyInfo KeyInfo `json:"keyInfo"`
}

// NewAccessKey instantiates a new AccessKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKey(idKey AccessKeyId, keyInfo KeyInfo) *AccessKey {
	this := AccessKey{}
	this.IdKey = idKey
	this.KeyInfo = keyInfo
	return &this
}

// NewAccessKeyWithDefaults instantiates a new AccessKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyWithDefaults() *AccessKey {
	this := AccessKey{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *AccessKey) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *AccessKey) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *AccessKey) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *AccessKey) GetExpirationDate() time.Time {
	if o == nil || o.ExpirationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *AccessKey) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *AccessKey) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetIdKey returns the IdKey field value
func (o *AccessKey) GetIdKey() AccessKeyId {
	if o == nil {
		var ret AccessKeyId
		return ret
	}

	return o.IdKey
}

// GetIdKeyOk returns a tuple with the IdKey field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetIdKeyOk() (*AccessKeyId, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdKey, true
}

// SetIdKey sets field value
func (o *AccessKey) SetIdKey(v AccessKeyId) {
	o.IdKey = v
}

// GetKeyInfo returns the KeyInfo field value
func (o *AccessKey) GetKeyInfo() KeyInfo {
	if o == nil {
		var ret KeyInfo
		return ret
	}

	return o.KeyInfo
}

// GetKeyInfoOk returns a tuple with the KeyInfo field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetKeyInfoOk() (*KeyInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyInfo, true
}

// SetKeyInfo sets field value
func (o *AccessKey) SetKeyInfo(v KeyInfo) {
	o.KeyInfo = v
}

func (o AccessKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if true {
		toSerialize["idKey"] = o.IdKey
	}
	if true {
		toSerialize["keyInfo"] = o.KeyInfo
	}
	return json.Marshal(toSerialize)
}

type NullableAccessKey struct {
	value *AccessKey
	isSet bool
}

func (v NullableAccessKey) Get() *AccessKey {
	return v.value
}

func (v *NullableAccessKey) Set(val *AccessKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKey(val *AccessKey) *NullableAccessKey {
	return &NullableAccessKey{value: val, isSet: true}
}

func (v NullableAccessKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


