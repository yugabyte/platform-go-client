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

// GetWebhookResponse Webhook get response
type GetWebhookResponse struct {
	// Webhook url
	Url *string `json:"url,omitempty"`
	// Webhook UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewGetWebhookResponse instantiates a new GetWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhookResponse() *GetWebhookResponse {
	this := GetWebhookResponse{}
	return &this
}

// NewGetWebhookResponseWithDefaults instantiates a new GetWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhookResponseWithDefaults() *GetWebhookResponse {
	this := GetWebhookResponse{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetWebhookResponse) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhookResponse) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetWebhookResponse) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetWebhookResponse) SetUrl(v string) {
	o.Url = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetWebhookResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhookResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GetWebhookResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetWebhookResponse) SetUuid(v string) {
	o.Uuid = &v
}

func (o GetWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableGetWebhookResponse struct {
	value *GetWebhookResponse
	isSet bool
}

func (v NullableGetWebhookResponse) Get() *GetWebhookResponse {
	return v.value
}

func (v *NullableGetWebhookResponse) Set(val *GetWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhookResponse(val *GetWebhookResponse) *NullableGetWebhookResponse {
	return &NullableGetWebhookResponse{value: val, isSet: true}
}

func (v NullableGetWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


