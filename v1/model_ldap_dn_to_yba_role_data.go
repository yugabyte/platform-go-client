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

// LdapDnToYbaRoleData A list of LDAP DN, YBA role pairs
type LdapDnToYbaRoleData struct {
	//  list of pairs of distinguishedName and ybaRole
	LdapDnToYbaRolePairs *[]LdapDnYbaRoleDataPair `json:"ldapDnToYbaRolePairs,omitempty"`
}

// NewLdapDnToYbaRoleData instantiates a new LdapDnToYbaRoleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapDnToYbaRoleData() *LdapDnToYbaRoleData {
	this := LdapDnToYbaRoleData{}
	return &this
}

// NewLdapDnToYbaRoleDataWithDefaults instantiates a new LdapDnToYbaRoleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapDnToYbaRoleDataWithDefaults() *LdapDnToYbaRoleData {
	this := LdapDnToYbaRoleData{}
	return &this
}

// GetLdapDnToYbaRolePairs returns the LdapDnToYbaRolePairs field value if set, zero value otherwise.
func (o *LdapDnToYbaRoleData) GetLdapDnToYbaRolePairs() []LdapDnYbaRoleDataPair {
	if o == nil || o.LdapDnToYbaRolePairs == nil {
		var ret []LdapDnYbaRoleDataPair
		return ret
	}
	return *o.LdapDnToYbaRolePairs
}

// GetLdapDnToYbaRolePairsOk returns a tuple with the LdapDnToYbaRolePairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDnToYbaRoleData) GetLdapDnToYbaRolePairsOk() (*[]LdapDnYbaRoleDataPair, bool) {
	if o == nil || o.LdapDnToYbaRolePairs == nil {
		return nil, false
	}
	return o.LdapDnToYbaRolePairs, true
}

// HasLdapDnToYbaRolePairs returns a boolean if a field has been set.
func (o *LdapDnToYbaRoleData) HasLdapDnToYbaRolePairs() bool {
	if o != nil && o.LdapDnToYbaRolePairs != nil {
		return true
	}

	return false
}

// SetLdapDnToYbaRolePairs gets a reference to the given []LdapDnYbaRoleDataPair and assigns it to the LdapDnToYbaRolePairs field.
func (o *LdapDnToYbaRoleData) SetLdapDnToYbaRolePairs(v []LdapDnYbaRoleDataPair) {
	o.LdapDnToYbaRolePairs = &v
}

func (o LdapDnToYbaRoleData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LdapDnToYbaRolePairs != nil {
		toSerialize["ldapDnToYbaRolePairs"] = o.LdapDnToYbaRolePairs
	}
	return json.Marshal(toSerialize)
}

type NullableLdapDnToYbaRoleData struct {
	value *LdapDnToYbaRoleData
	isSet bool
}

func (v NullableLdapDnToYbaRoleData) Get() *LdapDnToYbaRoleData {
	return v.value
}

func (v *NullableLdapDnToYbaRoleData) Set(val *LdapDnToYbaRoleData) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapDnToYbaRoleData) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapDnToYbaRoleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapDnToYbaRoleData(val *LdapDnToYbaRoleData) *NullableLdapDnToYbaRoleData {
	return &NullableLdapDnToYbaRoleData{value: val, isSet: true}
}

func (v NullableLdapDnToYbaRoleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapDnToYbaRoleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


