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

// AlertTemplateVariablesFormData struct for AlertTemplateVariablesFormData
type AlertTemplateVariablesFormData struct {
	Variables []AlertTemplateVariable `json:"variables"`
}

// NewAlertTemplateVariablesFormData instantiates a new AlertTemplateVariablesFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertTemplateVariablesFormData(variables []AlertTemplateVariable) *AlertTemplateVariablesFormData {
	this := AlertTemplateVariablesFormData{}
	this.Variables = variables
	return &this
}

// NewAlertTemplateVariablesFormDataWithDefaults instantiates a new AlertTemplateVariablesFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertTemplateVariablesFormDataWithDefaults() *AlertTemplateVariablesFormData {
	this := AlertTemplateVariablesFormData{}
	return &this
}

// GetVariables returns the Variables field value
func (o *AlertTemplateVariablesFormData) GetVariables() []AlertTemplateVariable {
	if o == nil {
		var ret []AlertTemplateVariable
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *AlertTemplateVariablesFormData) GetVariablesOk() (*[]AlertTemplateVariable, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value
func (o *AlertTemplateVariablesFormData) SetVariables(v []AlertTemplateVariable) {
	o.Variables = v
}

func (o AlertTemplateVariablesFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableAlertTemplateVariablesFormData struct {
	value *AlertTemplateVariablesFormData
	isSet bool
}

func (v NullableAlertTemplateVariablesFormData) Get() *AlertTemplateVariablesFormData {
	return v.value
}

func (v *NullableAlertTemplateVariablesFormData) Set(val *AlertTemplateVariablesFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertTemplateVariablesFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertTemplateVariablesFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertTemplateVariablesFormData(val *AlertTemplateVariablesFormData) *NullableAlertTemplateVariablesFormData {
	return &NullableAlertTemplateVariablesFormData{value: val, isSet: true}
}

func (v NullableAlertTemplateVariablesFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertTemplateVariablesFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


