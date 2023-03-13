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

// AdminNotification struct for AdminNotification
type AdminNotification struct {
	// Notification code
	Code *string `json:"code,omitempty"`
	// Notification message with HTML markup
	HtmlMessage *string `json:"htmlMessage,omitempty"`
}

// NewAdminNotification instantiates a new AdminNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminNotification() *AdminNotification {
	this := AdminNotification{}
	return &this
}

// NewAdminNotificationWithDefaults instantiates a new AdminNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminNotificationWithDefaults() *AdminNotification {
	this := AdminNotification{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AdminNotification) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminNotification) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AdminNotification) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AdminNotification) SetCode(v string) {
	o.Code = &v
}

// GetHtmlMessage returns the HtmlMessage field value if set, zero value otherwise.
func (o *AdminNotification) GetHtmlMessage() string {
	if o == nil || o.HtmlMessage == nil {
		var ret string
		return ret
	}
	return *o.HtmlMessage
}

// GetHtmlMessageOk returns a tuple with the HtmlMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminNotification) GetHtmlMessageOk() (*string, bool) {
	if o == nil || o.HtmlMessage == nil {
		return nil, false
	}
	return o.HtmlMessage, true
}

// HasHtmlMessage returns a boolean if a field has been set.
func (o *AdminNotification) HasHtmlMessage() bool {
	if o != nil && o.HtmlMessage != nil {
		return true
	}

	return false
}

// SetHtmlMessage gets a reference to the given string and assigns it to the HtmlMessage field.
func (o *AdminNotification) SetHtmlMessage(v string) {
	o.HtmlMessage = &v
}

func (o AdminNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.HtmlMessage != nil {
		toSerialize["htmlMessage"] = o.HtmlMessage
	}
	return json.Marshal(toSerialize)
}

type NullableAdminNotification struct {
	value *AdminNotification
	isSet bool
}

func (v NullableAdminNotification) Get() *AdminNotification {
	return v.value
}

func (v *NullableAdminNotification) Set(val *AdminNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminNotification(val *AdminNotification) *NullableAdminNotification {
	return &NullableAdminNotification{value: val, isSet: true}
}

func (v NullableAdminNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

