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

// TokenAuthInformation struct for TokenAuthInformation
type TokenAuthInformation struct {
	HTTPAuthInformation
	// Header name
	TokenHeader *string `json:"tokenHeader,omitempty"`
	// Token value
	TokenValue *string `json:"tokenValue,omitempty"`
}

// NewTokenAuthInformation instantiates a new TokenAuthInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenAuthInformation() *TokenAuthInformation {
	this := TokenAuthInformation{}
	return &this
}

// NewTokenAuthInformationWithDefaults instantiates a new TokenAuthInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenAuthInformationWithDefaults() *TokenAuthInformation {
	this := TokenAuthInformation{}
	return &this
}

// GetTokenHeader returns the TokenHeader field value if set, zero value otherwise.
func (o *TokenAuthInformation) GetTokenHeader() string {
	if o == nil || o.TokenHeader == nil {
		var ret string
		return ret
	}
	return *o.TokenHeader
}

// GetTokenHeaderOk returns a tuple with the TokenHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthInformation) GetTokenHeaderOk() (*string, bool) {
	if o == nil || o.TokenHeader == nil {
		return nil, false
	}
	return o.TokenHeader, true
}

// HasTokenHeader returns a boolean if a field has been set.
func (o *TokenAuthInformation) HasTokenHeader() bool {
	if o != nil && o.TokenHeader != nil {
		return true
	}

	return false
}

// SetTokenHeader gets a reference to the given string and assigns it to the TokenHeader field.
func (o *TokenAuthInformation) SetTokenHeader(v string) {
	o.TokenHeader = &v
}

// GetTokenValue returns the TokenValue field value if set, zero value otherwise.
func (o *TokenAuthInformation) GetTokenValue() string {
	if o == nil || o.TokenValue == nil {
		var ret string
		return ret
	}
	return *o.TokenValue
}

// GetTokenValueOk returns a tuple with the TokenValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthInformation) GetTokenValueOk() (*string, bool) {
	if o == nil || o.TokenValue == nil {
		return nil, false
	}
	return o.TokenValue, true
}

// HasTokenValue returns a boolean if a field has been set.
func (o *TokenAuthInformation) HasTokenValue() bool {
	if o != nil && o.TokenValue != nil {
		return true
	}

	return false
}

// SetTokenValue gets a reference to the given string and assigns it to the TokenValue field.
func (o *TokenAuthInformation) SetTokenValue(v string) {
	o.TokenValue = &v
}

func (o TokenAuthInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedHTTPAuthInformation, errHTTPAuthInformation := json.Marshal(o.HTTPAuthInformation)
	if errHTTPAuthInformation != nil {
		return []byte{}, errHTTPAuthInformation
	}
	errHTTPAuthInformation = json.Unmarshal([]byte(serializedHTTPAuthInformation), &toSerialize)
	if errHTTPAuthInformation != nil {
		return []byte{}, errHTTPAuthInformation
	}
	if o.TokenHeader != nil {
		toSerialize["tokenHeader"] = o.TokenHeader
	}
	if o.TokenValue != nil {
		toSerialize["tokenValue"] = o.TokenValue
	}
	return json.Marshal(toSerialize)
}

type NullableTokenAuthInformation struct {
	value *TokenAuthInformation
	isSet bool
}

func (v NullableTokenAuthInformation) Get() *TokenAuthInformation {
	return v.value
}

func (v *NullableTokenAuthInformation) Set(val *TokenAuthInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenAuthInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenAuthInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenAuthInformation(val *TokenAuthInformation) *NullableTokenAuthInformation {
	return &NullableTokenAuthInformation{value: val, isSet: true}
}

func (v NullableTokenAuthInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenAuthInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


