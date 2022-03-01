/*
 * Yugabyte Platform APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yb_platform_client

import (
	"encoding/json"
)

// AlertChannelPagerDutyParams struct for AlertChannelPagerDutyParams
type AlertChannelPagerDutyParams struct {
	AlertChannelParams
	ApiKey string `json:"apiKey"`
	RoutingKey string `json:"routingKey"`
}

// NewAlertChannelPagerDutyParams instantiates a new AlertChannelPagerDutyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertChannelPagerDutyParams(apiKey string, routingKey string, ) *AlertChannelPagerDutyParams {
	this := AlertChannelPagerDutyParams{}
	this.ApiKey = apiKey
	this.RoutingKey = routingKey
	return &this
}

// NewAlertChannelPagerDutyParamsWithDefaults instantiates a new AlertChannelPagerDutyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertChannelPagerDutyParamsWithDefaults() *AlertChannelPagerDutyParams {
	this := AlertChannelPagerDutyParams{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *AlertChannelPagerDutyParams) GetApiKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *AlertChannelPagerDutyParams) GetApiKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *AlertChannelPagerDutyParams) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRoutingKey returns the RoutingKey field value
func (o *AlertChannelPagerDutyParams) GetRoutingKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RoutingKey
}

// GetRoutingKeyOk returns a tuple with the RoutingKey field value
// and a boolean to check if the value has been set.
func (o *AlertChannelPagerDutyParams) GetRoutingKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoutingKey, true
}

// SetRoutingKey sets field value
func (o *AlertChannelPagerDutyParams) SetRoutingKey(v string) {
	o.RoutingKey = v
}

func (o AlertChannelPagerDutyParams) MarshalJSON() ([]byte, error) {
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
		toSerialize["apiKey"] = o.ApiKey
	}
	if true {
		toSerialize["routingKey"] = o.RoutingKey
	}
	return json.Marshal(toSerialize)
}

type NullableAlertChannelPagerDutyParams struct {
	value *AlertChannelPagerDutyParams
	isSet bool
}

func (v NullableAlertChannelPagerDutyParams) Get() *AlertChannelPagerDutyParams {
	return v.value
}

func (v *NullableAlertChannelPagerDutyParams) Set(val *AlertChannelPagerDutyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertChannelPagerDutyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertChannelPagerDutyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertChannelPagerDutyParams(val *AlertChannelPagerDutyParams) *NullableAlertChannelPagerDutyParams {
	return &NullableAlertChannelPagerDutyParams{value: val, isSet: true}
}

func (v NullableAlertChannelPagerDutyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertChannelPagerDutyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


