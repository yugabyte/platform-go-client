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

// ReinstallNodeAgentForm struct for ReinstallNodeAgentForm
type ReinstallNodeAgentForm struct {
	// Node names
	NodeNames *[]string `json:"nodeNames,omitempty"`
}

// NewReinstallNodeAgentForm instantiates a new ReinstallNodeAgentForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReinstallNodeAgentForm() *ReinstallNodeAgentForm {
	this := ReinstallNodeAgentForm{}
	return &this
}

// NewReinstallNodeAgentFormWithDefaults instantiates a new ReinstallNodeAgentForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReinstallNodeAgentFormWithDefaults() *ReinstallNodeAgentForm {
	this := ReinstallNodeAgentForm{}
	return &this
}

// GetNodeNames returns the NodeNames field value if set, zero value otherwise.
func (o *ReinstallNodeAgentForm) GetNodeNames() []string {
	if o == nil || o.NodeNames == nil {
		var ret []string
		return ret
	}
	return *o.NodeNames
}

// GetNodeNamesOk returns a tuple with the NodeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReinstallNodeAgentForm) GetNodeNamesOk() (*[]string, bool) {
	if o == nil || o.NodeNames == nil {
		return nil, false
	}
	return o.NodeNames, true
}

// HasNodeNames returns a boolean if a field has been set.
func (o *ReinstallNodeAgentForm) HasNodeNames() bool {
	if o != nil && o.NodeNames != nil {
		return true
	}

	return false
}

// SetNodeNames gets a reference to the given []string and assigns it to the NodeNames field.
func (o *ReinstallNodeAgentForm) SetNodeNames(v []string) {
	o.NodeNames = &v
}

func (o ReinstallNodeAgentForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeNames != nil {
		toSerialize["nodeNames"] = o.NodeNames
	}
	return json.Marshal(toSerialize)
}

type NullableReinstallNodeAgentForm struct {
	value *ReinstallNodeAgentForm
	isSet bool
}

func (v NullableReinstallNodeAgentForm) Get() *ReinstallNodeAgentForm {
	return v.value
}

func (v *NullableReinstallNodeAgentForm) Set(val *ReinstallNodeAgentForm) {
	v.value = val
	v.isSet = true
}

func (v NullableReinstallNodeAgentForm) IsSet() bool {
	return v.isSet
}

func (v *NullableReinstallNodeAgentForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReinstallNodeAgentForm(val *ReinstallNodeAgentForm) *NullableReinstallNodeAgentForm {
	return &NullableReinstallNodeAgentForm{value: val, isSet: true}
}

func (v NullableReinstallNodeAgentForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReinstallNodeAgentForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


