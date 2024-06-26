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

// AlertChannelPagerDutyParamsAllOf struct for AlertChannelPagerDutyParamsAllOf
type AlertChannelPagerDutyParamsAllOf struct {
	// API key
	ApiKey string `json:"apiKey"`
	// Routing key
	RoutingKey string `json:"routingKey"`
}

// NewAlertChannelPagerDutyParamsAllOf instantiates a new AlertChannelPagerDutyParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertChannelPagerDutyParamsAllOf(apiKey string, routingKey string) *AlertChannelPagerDutyParamsAllOf {
	this := AlertChannelPagerDutyParamsAllOf{}
	this.ApiKey = apiKey
	this.RoutingKey = routingKey
	return &this
}

// NewAlertChannelPagerDutyParamsAllOfWithDefaults instantiates a new AlertChannelPagerDutyParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertChannelPagerDutyParamsAllOfWithDefaults() *AlertChannelPagerDutyParamsAllOf {
	this := AlertChannelPagerDutyParamsAllOf{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *AlertChannelPagerDutyParamsAllOf) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *AlertChannelPagerDutyParamsAllOf) GetApiKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *AlertChannelPagerDutyParamsAllOf) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRoutingKey returns the RoutingKey field value
func (o *AlertChannelPagerDutyParamsAllOf) GetRoutingKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingKey
}

// GetRoutingKeyOk returns a tuple with the RoutingKey field value
// and a boolean to check if the value has been set.
func (o *AlertChannelPagerDutyParamsAllOf) GetRoutingKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoutingKey, true
}

// SetRoutingKey sets field value
func (o *AlertChannelPagerDutyParamsAllOf) SetRoutingKey(v string) {
	o.RoutingKey = v
}

func (o AlertChannelPagerDutyParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiKey"] = o.ApiKey
	}
	if true {
		toSerialize["routingKey"] = o.RoutingKey
	}
	return json.Marshal(toSerialize)
}

type NullableAlertChannelPagerDutyParamsAllOf struct {
	value *AlertChannelPagerDutyParamsAllOf
	isSet bool
}

func (v NullableAlertChannelPagerDutyParamsAllOf) Get() *AlertChannelPagerDutyParamsAllOf {
	return v.value
}

func (v *NullableAlertChannelPagerDutyParamsAllOf) Set(val *AlertChannelPagerDutyParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertChannelPagerDutyParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertChannelPagerDutyParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertChannelPagerDutyParamsAllOf(val *AlertChannelPagerDutyParamsAllOf) *NullableAlertChannelPagerDutyParamsAllOf {
	return &NullableAlertChannelPagerDutyParamsAllOf{value: val, isSet: true}
}

func (v NullableAlertChannelPagerDutyParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertChannelPagerDutyParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


