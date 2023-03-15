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

// AlertChannelEmailParams struct for AlertChannelEmailParams
type AlertChannelEmailParams struct {
	AlertChannelParams
	DefaultRecipients bool `json:"defaultRecipients"`
	DefaultSmtpSettings bool `json:"defaultSmtpSettings"`
	Recipients []string `json:"recipients"`
	SmtpData SmtpData `json:"smtpData"`
}

// NewAlertChannelEmailParams instantiates a new AlertChannelEmailParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertChannelEmailParams(defaultRecipients bool, defaultSmtpSettings bool, recipients []string, smtpData SmtpData, ) *AlertChannelEmailParams {
	this := AlertChannelEmailParams{}
	this.DefaultRecipients = defaultRecipients
	this.DefaultSmtpSettings = defaultSmtpSettings
	this.Recipients = recipients
	this.SmtpData = smtpData
	return &this
}

// NewAlertChannelEmailParamsWithDefaults instantiates a new AlertChannelEmailParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertChannelEmailParamsWithDefaults() *AlertChannelEmailParams {
	this := AlertChannelEmailParams{}
	return &this
}

// GetDefaultRecipients returns the DefaultRecipients field value
func (o *AlertChannelEmailParams) GetDefaultRecipients() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.DefaultRecipients
}

// GetDefaultRecipientsOk returns a tuple with the DefaultRecipients field value
// and a boolean to check if the value has been set.
func (o *AlertChannelEmailParams) GetDefaultRecipientsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultRecipients, true
}

// SetDefaultRecipients sets field value
func (o *AlertChannelEmailParams) SetDefaultRecipients(v bool) {
	o.DefaultRecipients = v
}

// GetDefaultSmtpSettings returns the DefaultSmtpSettings field value
func (o *AlertChannelEmailParams) GetDefaultSmtpSettings() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.DefaultSmtpSettings
}

// GetDefaultSmtpSettingsOk returns a tuple with the DefaultSmtpSettings field value
// and a boolean to check if the value has been set.
func (o *AlertChannelEmailParams) GetDefaultSmtpSettingsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultSmtpSettings, true
}

// SetDefaultSmtpSettings sets field value
func (o *AlertChannelEmailParams) SetDefaultSmtpSettings(v bool) {
	o.DefaultSmtpSettings = v
}

// GetRecipients returns the Recipients field value
func (o *AlertChannelEmailParams) GetRecipients() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *AlertChannelEmailParams) GetRecipientsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *AlertChannelEmailParams) SetRecipients(v []string) {
	o.Recipients = v
}

// GetSmtpData returns the SmtpData field value
func (o *AlertChannelEmailParams) GetSmtpData() SmtpData {
	if o == nil  {
		var ret SmtpData
		return ret
	}

	return o.SmtpData
}

// GetSmtpDataOk returns a tuple with the SmtpData field value
// and a boolean to check if the value has been set.
func (o *AlertChannelEmailParams) GetSmtpDataOk() (*SmtpData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SmtpData, true
}

// SetSmtpData sets field value
func (o *AlertChannelEmailParams) SetSmtpData(v SmtpData) {
	o.SmtpData = v
}

func (o AlertChannelEmailParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAlertChannelParams, errAlertChannelParams := json.Marshal(o.AlertChannelParams)
	if errAlertChannelParams != nil {
		return []byte{}, errAlertChannelParams
	}
	errAlertChannelParams = json.Unmarshal([]byte(serializedAlertChannelParams), &toSerialize)
	if errAlertChannelParams != nil {
		return []byte{}, errAlertChannelParams
	}
	if true {
		toSerialize["defaultRecipients"] = o.DefaultRecipients
	}
	if true {
		toSerialize["defaultSmtpSettings"] = o.DefaultSmtpSettings
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["smtpData"] = o.SmtpData
	}
	return json.Marshal(toSerialize)
}

type NullableAlertChannelEmailParams struct {
	value *AlertChannelEmailParams
	isSet bool
}

func (v NullableAlertChannelEmailParams) Get() *AlertChannelEmailParams {
	return v.value
}

func (v *NullableAlertChannelEmailParams) Set(val *AlertChannelEmailParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertChannelEmailParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertChannelEmailParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertChannelEmailParams(val *AlertChannelEmailParams) *NullableAlertChannelEmailParams {
	return &NullableAlertChannelEmailParams{value: val, isSet: true}
}

func (v NullableAlertChannelEmailParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertChannelEmailParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


