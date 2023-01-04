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

// AlertChannelSlackParams struct for AlertChannelSlackParams
type AlertChannelSlackParams struct {
	AlertChannelParams
	// Slack icon URL
	IconUrl *string `json:"iconUrl,omitempty"`
	// Slack username
	Username string `json:"username"`
	// Slack webhook URL
	WebhookUrl string `json:"webhookUrl"`
}

// NewAlertChannelSlackParams instantiates a new AlertChannelSlackParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertChannelSlackParams(username string, webhookUrl string, ) *AlertChannelSlackParams {
	this := AlertChannelSlackParams{}
	this.Username = username
	this.WebhookUrl = webhookUrl
	return &this
}

// NewAlertChannelSlackParamsWithDefaults instantiates a new AlertChannelSlackParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertChannelSlackParamsWithDefaults() *AlertChannelSlackParams {
	this := AlertChannelSlackParams{}
	return &this
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *AlertChannelSlackParams) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertChannelSlackParams) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *AlertChannelSlackParams) HasIconUrl() bool {
	if o != nil && o.IconUrl != nil {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *AlertChannelSlackParams) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetUsername returns the Username field value
func (o *AlertChannelSlackParams) GetUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AlertChannelSlackParams) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AlertChannelSlackParams) SetUsername(v string) {
	o.Username = v
}

// GetWebhookUrl returns the WebhookUrl field value
func (o *AlertChannelSlackParams) GetWebhookUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value
// and a boolean to check if the value has been set.
func (o *AlertChannelSlackParams) GetWebhookUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookUrl, true
}

// SetWebhookUrl sets field value
func (o *AlertChannelSlackParams) SetWebhookUrl(v string) {
	o.WebhookUrl = v
}

func (o AlertChannelSlackParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAlertChannelParams, errAlertChannelParams := json.Marshal(o.AlertChannelParams)
	if errAlertChannelParams != nil {
		return []byte{}, errAlertChannelParams
	}
	errAlertChannelParams = json.Unmarshal([]byte(serializedAlertChannelParams), &toSerialize)
	if errAlertChannelParams != nil {
		return []byte{}, errAlertChannelParams
	}
	if o.IconUrl != nil {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return json.Marshal(toSerialize)
}

type NullableAlertChannelSlackParams struct {
	value *AlertChannelSlackParams
	isSet bool
}

func (v NullableAlertChannelSlackParams) Get() *AlertChannelSlackParams {
	return v.value
}

func (v *NullableAlertChannelSlackParams) Set(val *AlertChannelSlackParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertChannelSlackParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertChannelSlackParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertChannelSlackParams(val *AlertChannelSlackParams) *NullableAlertChannelSlackParams {
	return &NullableAlertChannelSlackParams{value: val, isSet: true}
}

func (v NullableAlertChannelSlackParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertChannelSlackParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


