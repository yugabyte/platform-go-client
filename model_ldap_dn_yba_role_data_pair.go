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

// LdapDnYbaRoleDataPair struct for LdapDnYbaRoleDataPair
type LdapDnYbaRoleDataPair struct {
	DistinguishedName string `json:"distinguishedName"`
	YbaRole string `json:"ybaRole"`
}

// NewLdapDnYbaRoleDataPair instantiates a new LdapDnYbaRoleDataPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapDnYbaRoleDataPair(distinguishedName string, ybaRole string) *LdapDnYbaRoleDataPair {
	this := LdapDnYbaRoleDataPair{}
	this.DistinguishedName = distinguishedName
	this.YbaRole = ybaRole
	return &this
}

// NewLdapDnYbaRoleDataPairWithDefaults instantiates a new LdapDnYbaRoleDataPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapDnYbaRoleDataPairWithDefaults() *LdapDnYbaRoleDataPair {
	this := LdapDnYbaRoleDataPair{}
	return &this
}

// GetDistinguishedName returns the DistinguishedName field value
func (o *LdapDnYbaRoleDataPair) GetDistinguishedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value
// and a boolean to check if the value has been set.
func (o *LdapDnYbaRoleDataPair) GetDistinguishedNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DistinguishedName, true
}

// SetDistinguishedName sets field value
func (o *LdapDnYbaRoleDataPair) SetDistinguishedName(v string) {
	o.DistinguishedName = v
}

// GetYbaRole returns the YbaRole field value
func (o *LdapDnYbaRoleDataPair) GetYbaRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbaRole
}

// GetYbaRoleOk returns a tuple with the YbaRole field value
// and a boolean to check if the value has been set.
func (o *LdapDnYbaRoleDataPair) GetYbaRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbaRole, true
}

// SetYbaRole sets field value
func (o *LdapDnYbaRoleDataPair) SetYbaRole(v string) {
	o.YbaRole = v
}

func (o LdapDnYbaRoleDataPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["distinguishedName"] = o.DistinguishedName
	}
	if true {
		toSerialize["ybaRole"] = o.YbaRole
	}
	return json.Marshal(toSerialize)
}

type NullableLdapDnYbaRoleDataPair struct {
	value *LdapDnYbaRoleDataPair
	isSet bool
}

func (v NullableLdapDnYbaRoleDataPair) Get() *LdapDnYbaRoleDataPair {
	return v.value
}

func (v *NullableLdapDnYbaRoleDataPair) Set(val *LdapDnYbaRoleDataPair) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapDnYbaRoleDataPair) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapDnYbaRoleDataPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapDnYbaRoleDataPair(val *LdapDnYbaRoleDataPair) *NullableLdapDnYbaRoleDataPair {
	return &NullableLdapDnYbaRoleDataPair{value: val, isSet: true}
}

func (v NullableLdapDnYbaRoleDataPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapDnYbaRoleDataPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


