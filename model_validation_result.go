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

// ValidationResult Validation result of a node config
type ValidationResult struct {
	Description string `json:"description"`
	IsRequired bool `json:"isRequired"`
	IsValid bool `json:"isValid"`
	Type string `json:"type"`
	Value string `json:"value"`
}

// NewValidationResult instantiates a new ValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationResult(description string, isRequired bool, isValid bool, type_ string, value string, ) *ValidationResult {
	this := ValidationResult{}
	this.Description = description
	this.IsRequired = isRequired
	this.IsValid = isValid
	this.Type = type_
	this.Value = value
	return &this
}

// NewValidationResultWithDefaults instantiates a new ValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationResultWithDefaults() *ValidationResult {
	this := ValidationResult{}
	return &this
}

// GetDescription returns the Description field value
func (o *ValidationResult) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ValidationResult) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ValidationResult) SetDescription(v string) {
	o.Description = v
}

// GetIsRequired returns the IsRequired field value
func (o *ValidationResult) GetIsRequired() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value
// and a boolean to check if the value has been set.
func (o *ValidationResult) GetIsRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsRequired, true
}

// SetIsRequired sets field value
func (o *ValidationResult) SetIsRequired(v bool) {
	o.IsRequired = v
}

// GetIsValid returns the IsValid field value
func (o *ValidationResult) GetIsValid() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *ValidationResult) GetIsValidOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value
func (o *ValidationResult) SetIsValid(v bool) {
	o.IsValid = v
}

// GetType returns the Type field value
func (o *ValidationResult) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ValidationResult) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ValidationResult) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ValidationResult) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ValidationResult) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ValidationResult) SetValue(v string) {
	o.Value = v
}

func (o ValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["isRequired"] = o.IsRequired
	}
	if true {
		toSerialize["isValid"] = o.IsValid
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableValidationResult struct {
	value *ValidationResult
	isSet bool
}

func (v NullableValidationResult) Get() *ValidationResult {
	return v.value
}

func (v *NullableValidationResult) Set(val *ValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationResult(val *ValidationResult) *NullableValidationResult {
	return &NullableValidationResult{value: val, isSet: true}
}

func (v NullableValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

