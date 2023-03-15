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

// AlertChannelParams struct for AlertChannelParams
type AlertChannelParams struct {
	TextTemplate string `json:"textTemplate"`
	TitleTemplate string `json:"titleTemplate"`
}

// NewAlertChannelParams instantiates a new AlertChannelParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertChannelParams(textTemplate string, titleTemplate string, ) *AlertChannelParams {
	this := AlertChannelParams{}
	this.TextTemplate = textTemplate
	this.TitleTemplate = titleTemplate
	return &this
}

// NewAlertChannelParamsWithDefaults instantiates a new AlertChannelParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertChannelParamsWithDefaults() *AlertChannelParams {
	this := AlertChannelParams{}
	return &this
}

// GetTextTemplate returns the TextTemplate field value
func (o *AlertChannelParams) GetTextTemplate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TextTemplate
}

// GetTextTemplateOk returns a tuple with the TextTemplate field value
// and a boolean to check if the value has been set.
func (o *AlertChannelParams) GetTextTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TextTemplate, true
}

// SetTextTemplate sets field value
func (o *AlertChannelParams) SetTextTemplate(v string) {
	o.TextTemplate = v
}

// GetTitleTemplate returns the TitleTemplate field value
func (o *AlertChannelParams) GetTitleTemplate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TitleTemplate
}

// GetTitleTemplateOk returns a tuple with the TitleTemplate field value
// and a boolean to check if the value has been set.
func (o *AlertChannelParams) GetTitleTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TitleTemplate, true
}

// SetTitleTemplate sets field value
func (o *AlertChannelParams) SetTitleTemplate(v string) {
	o.TitleTemplate = v
}

func (o AlertChannelParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["textTemplate"] = o.TextTemplate
	}
	if true {
		toSerialize["titleTemplate"] = o.TitleTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableAlertChannelParams struct {
	value *AlertChannelParams
	isSet bool
}

func (v NullableAlertChannelParams) Get() *AlertChannelParams {
	return v.value
}

func (v *NullableAlertChannelParams) Set(val *AlertChannelParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertChannelParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertChannelParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertChannelParams(val *AlertChannelParams) *NullableAlertChannelParams {
	return &NullableAlertChannelParams{value: val, isSet: true}
}

func (v NullableAlertChannelParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertChannelParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


