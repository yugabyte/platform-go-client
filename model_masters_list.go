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

// MastersList struct for MastersList
type MastersList struct {
	Masters []MasterNode `json:"masters"`
}

// NewMastersList instantiates a new MastersList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMastersList(masters []MasterNode) *MastersList {
	this := MastersList{}
	this.Masters = masters
	return &this
}

// NewMastersListWithDefaults instantiates a new MastersList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMastersListWithDefaults() *MastersList {
	this := MastersList{}
	return &this
}

// GetMasters returns the Masters field value
func (o *MastersList) GetMasters() []MasterNode {
	if o == nil {
		var ret []MasterNode
		return ret
	}

	return o.Masters
}

// GetMastersOk returns a tuple with the Masters field value
// and a boolean to check if the value has been set.
func (o *MastersList) GetMastersOk() (*[]MasterNode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Masters, true
}

// SetMasters sets field value
func (o *MastersList) SetMasters(v []MasterNode) {
	o.Masters = v
}

func (o MastersList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["masters"] = o.Masters
	}
	return json.Marshal(toSerialize)
}

type NullableMastersList struct {
	value *MastersList
	isSet bool
}

func (v NullableMastersList) Get() *MastersList {
	return v.value
}

func (v *NullableMastersList) Set(val *MastersList) {
	v.value = val
	v.isSet = true
}

func (v NullableMastersList) IsSet() bool {
	return v.isSet
}

func (v *NullableMastersList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMastersList(val *MastersList) *NullableMastersList {
	return &NullableMastersList{value: val, isSet: true}
}

func (v NullableMastersList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMastersList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


