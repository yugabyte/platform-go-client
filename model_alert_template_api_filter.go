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

// AlertTemplateApiFilter struct for AlertTemplateApiFilter
type AlertTemplateApiFilter struct {
	// The name of the alert template.
	Name *string `json:"name,omitempty"`
	// The target type of the alert template.
	TargetType *string `json:"targetType,omitempty"`
}

// NewAlertTemplateApiFilter instantiates a new AlertTemplateApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertTemplateApiFilter() *AlertTemplateApiFilter {
	this := AlertTemplateApiFilter{}
	return &this
}

// NewAlertTemplateApiFilterWithDefaults instantiates a new AlertTemplateApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertTemplateApiFilterWithDefaults() *AlertTemplateApiFilter {
	this := AlertTemplateApiFilter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertTemplateApiFilter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplateApiFilter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertTemplateApiFilter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertTemplateApiFilter) SetName(v string) {
	o.Name = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AlertTemplateApiFilter) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplateApiFilter) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AlertTemplateApiFilter) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AlertTemplateApiFilter) SetTargetType(v string) {
	o.TargetType = &v
}

func (o AlertTemplateApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TargetType != nil {
		toSerialize["targetType"] = o.TargetType
	}
	return json.Marshal(toSerialize)
}

type NullableAlertTemplateApiFilter struct {
	value *AlertTemplateApiFilter
	isSet bool
}

func (v NullableAlertTemplateApiFilter) Get() *AlertTemplateApiFilter {
	return v.value
}

func (v *NullableAlertTemplateApiFilter) Set(val *AlertTemplateApiFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertTemplateApiFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertTemplateApiFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertTemplateApiFilter(val *AlertTemplateApiFilter) *NullableAlertTemplateApiFilter {
	return &NullableAlertTemplateApiFilter{value: val, isSet: true}
}

func (v NullableAlertTemplateApiFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertTemplateApiFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


