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

// AlertTemplateSettingsFormData struct for AlertTemplateSettingsFormData
type AlertTemplateSettingsFormData struct {
	Settings []AlertTemplateSettings `json:"settings"`
}

// NewAlertTemplateSettingsFormData instantiates a new AlertTemplateSettingsFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertTemplateSettingsFormData(settings []AlertTemplateSettings, ) *AlertTemplateSettingsFormData {
	this := AlertTemplateSettingsFormData{}
	this.Settings = settings
	return &this
}

// NewAlertTemplateSettingsFormDataWithDefaults instantiates a new AlertTemplateSettingsFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertTemplateSettingsFormDataWithDefaults() *AlertTemplateSettingsFormData {
	this := AlertTemplateSettingsFormData{}
	return &this
}

// GetSettings returns the Settings field value
func (o *AlertTemplateSettingsFormData) GetSettings() []AlertTemplateSettings {
	if o == nil  {
		var ret []AlertTemplateSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *AlertTemplateSettingsFormData) GetSettingsOk() (*[]AlertTemplateSettings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *AlertTemplateSettingsFormData) SetSettings(v []AlertTemplateSettings) {
	o.Settings = v
}

func (o AlertTemplateSettingsFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableAlertTemplateSettingsFormData struct {
	value *AlertTemplateSettingsFormData
	isSet bool
}

func (v NullableAlertTemplateSettingsFormData) Get() *AlertTemplateSettingsFormData {
	return v.value
}

func (v *NullableAlertTemplateSettingsFormData) Set(val *AlertTemplateSettingsFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertTemplateSettingsFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertTemplateSettingsFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertTemplateSettingsFormData(val *AlertTemplateSettingsFormData) *NullableAlertTemplateSettingsFormData {
	return &NullableAlertTemplateSettingsFormData{value: val, isSet: true}
}

func (v NullableAlertTemplateSettingsFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertTemplateSettingsFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

