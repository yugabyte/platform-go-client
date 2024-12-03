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

// Principal struct for Principal
type Principal struct {
	GroupUUID string `json:"groupUUID"`
	Type string `json:"type"`
	UserUUID string `json:"userUUID"`
	Uuid string `json:"uuid"`
}

// NewPrincipal instantiates a new Principal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipal(groupUUID string, type_ string, userUUID string, uuid string) *Principal {
	this := Principal{}
	this.GroupUUID = groupUUID
	this.Type = type_
	this.UserUUID = userUUID
	this.Uuid = uuid
	return &this
}

// NewPrincipalWithDefaults instantiates a new Principal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalWithDefaults() *Principal {
	this := Principal{}
	return &this
}

// GetGroupUUID returns the GroupUUID field value
func (o *Principal) GetGroupUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupUUID
}

// GetGroupUUIDOk returns a tuple with the GroupUUID field value
// and a boolean to check if the value has been set.
func (o *Principal) GetGroupUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GroupUUID, true
}

// SetGroupUUID sets field value
func (o *Principal) SetGroupUUID(v string) {
	o.GroupUUID = v
}

// GetType returns the Type field value
func (o *Principal) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Principal) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Principal) SetType(v string) {
	o.Type = v
}

// GetUserUUID returns the UserUUID field value
func (o *Principal) GetUserUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserUUID
}

// GetUserUUIDOk returns a tuple with the UserUUID field value
// and a boolean to check if the value has been set.
func (o *Principal) GetUserUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserUUID, true
}

// SetUserUUID sets field value
func (o *Principal) SetUserUUID(v string) {
	o.UserUUID = v
}

// GetUuid returns the Uuid field value
func (o *Principal) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Principal) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Principal) SetUuid(v string) {
	o.Uuid = v
}

func (o Principal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["groupUUID"] = o.GroupUUID
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["userUUID"] = o.UserUUID
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullablePrincipal struct {
	value *Principal
	isSet bool
}

func (v NullablePrincipal) Get() *Principal {
	return v.value
}

func (v *NullablePrincipal) Set(val *Principal) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipal(val *Principal) *NullablePrincipal {
	return &NullablePrincipal{value: val, isSet: true}
}

func (v NullablePrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

