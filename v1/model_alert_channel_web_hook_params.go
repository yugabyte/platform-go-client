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

// AlertChannelWebHookParams struct for AlertChannelWebHookParams
type AlertChannelWebHookParams struct {
	AlertChannelParams
	HttpAuth *HTTPAuthInformation `json:"httpAuth,omitempty"`
	// WARNING: This is a preview API that could change. Send resolved alert notification
	SendResolved *bool `json:"sendResolved,omitempty"`
	// Webhook URL
	WebhookUrl string `json:"webhookUrl"`
}

// NewAlertChannelWebHookParams instantiates a new AlertChannelWebHookParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertChannelWebHookParams(webhookUrl string) *AlertChannelWebHookParams {
	this := AlertChannelWebHookParams{}
	this.WebhookUrl = webhookUrl
	return &this
}

// NewAlertChannelWebHookParamsWithDefaults instantiates a new AlertChannelWebHookParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertChannelWebHookParamsWithDefaults() *AlertChannelWebHookParams {
	this := AlertChannelWebHookParams{}
	return &this
}

// GetHttpAuth returns the HttpAuth field value if set, zero value otherwise.
func (o *AlertChannelWebHookParams) GetHttpAuth() HTTPAuthInformation {
	if o == nil || o.HttpAuth == nil {
		var ret HTTPAuthInformation
		return ret
	}
	return *o.HttpAuth
}

// GetHttpAuthOk returns a tuple with the HttpAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertChannelWebHookParams) GetHttpAuthOk() (*HTTPAuthInformation, bool) {
	if o == nil || o.HttpAuth == nil {
		return nil, false
	}
	return o.HttpAuth, true
}

// HasHttpAuth returns a boolean if a field has been set.
func (o *AlertChannelWebHookParams) HasHttpAuth() bool {
	if o != nil && o.HttpAuth != nil {
		return true
	}

	return false
}

// SetHttpAuth gets a reference to the given HTTPAuthInformation and assigns it to the HttpAuth field.
func (o *AlertChannelWebHookParams) SetHttpAuth(v HTTPAuthInformation) {
	o.HttpAuth = &v
}

// GetSendResolved returns the SendResolved field value if set, zero value otherwise.
func (o *AlertChannelWebHookParams) GetSendResolved() bool {
	if o == nil || o.SendResolved == nil {
		var ret bool
		return ret
	}
	return *o.SendResolved
}

// GetSendResolvedOk returns a tuple with the SendResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertChannelWebHookParams) GetSendResolvedOk() (*bool, bool) {
	if o == nil || o.SendResolved == nil {
		return nil, false
	}
	return o.SendResolved, true
}

// HasSendResolved returns a boolean if a field has been set.
func (o *AlertChannelWebHookParams) HasSendResolved() bool {
	if o != nil && o.SendResolved != nil {
		return true
	}

	return false
}

// SetSendResolved gets a reference to the given bool and assigns it to the SendResolved field.
func (o *AlertChannelWebHookParams) SetSendResolved(v bool) {
	o.SendResolved = &v
}

// GetWebhookUrl returns the WebhookUrl field value
func (o *AlertChannelWebHookParams) GetWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value
// and a boolean to check if the value has been set.
func (o *AlertChannelWebHookParams) GetWebhookUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookUrl, true
}

// SetWebhookUrl sets field value
func (o *AlertChannelWebHookParams) SetWebhookUrl(v string) {
	o.WebhookUrl = v
}

func (o AlertChannelWebHookParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAlertChannelParams, errAlertChannelParams := json.Marshal(o.AlertChannelParams)
	if errAlertChannelParams != nil {
		return []byte{}, errAlertChannelParams
	}
	errAlertChannelParams = json.Unmarshal([]byte(serializedAlertChannelParams), &toSerialize)
	if errAlertChannelParams != nil {
		return []byte{}, errAlertChannelParams
	}
	if o.HttpAuth != nil {
		toSerialize["httpAuth"] = o.HttpAuth
	}
	if o.SendResolved != nil {
		toSerialize["sendResolved"] = o.SendResolved
	}
	if true {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return json.Marshal(toSerialize)
}

type NullableAlertChannelWebHookParams struct {
	value *AlertChannelWebHookParams
	isSet bool
}

func (v NullableAlertChannelWebHookParams) Get() *AlertChannelWebHookParams {
	return v.value
}

func (v *NullableAlertChannelWebHookParams) Set(val *AlertChannelWebHookParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertChannelWebHookParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertChannelWebHookParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertChannelWebHookParams(val *AlertChannelWebHookParams) *NullableAlertChannelWebHookParams {
	return &NullableAlertChannelWebHookParams{value: val, isSet: true}
}

func (v NullableAlertChannelWebHookParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertChannelWebHookParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


