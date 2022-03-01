/*
 * Yugabyte Platform APIs
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

// AlertDestination struct for AlertDestination
type AlertDestination struct {
	Channels []string `json:"channels"`
	CustomerUUID string `json:"customerUUID"`
	DefaultDestination bool `json:"defaultDestination"`
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

// NewAlertDestination instantiates a new AlertDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertDestination(channels []string, customerUUID string, defaultDestination bool, name string, uuid string, ) *AlertDestination {
	this := AlertDestination{}
	this.Channels = channels
	this.CustomerUUID = customerUUID
	this.DefaultDestination = defaultDestination
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewAlertDestinationWithDefaults instantiates a new AlertDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertDestinationWithDefaults() *AlertDestination {
	this := AlertDestination{}
	return &this
}

// GetChannels returns the Channels field value
func (o *AlertDestination) GetChannels() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *AlertDestination) GetChannelsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Channels, true
}

// SetChannels sets field value
func (o *AlertDestination) SetChannels(v []string) {
	o.Channels = v
}

// GetCustomerUUID returns the CustomerUUID field value
func (o *AlertDestination) GetCustomerUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value
// and a boolean to check if the value has been set.
func (o *AlertDestination) GetCustomerUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerUUID, true
}

// SetCustomerUUID sets field value
func (o *AlertDestination) SetCustomerUUID(v string) {
	o.CustomerUUID = v
}

// GetDefaultDestination returns the DefaultDestination field value
func (o *AlertDestination) GetDefaultDestination() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.DefaultDestination
}

// GetDefaultDestinationOk returns a tuple with the DefaultDestination field value
// and a boolean to check if the value has been set.
func (o *AlertDestination) GetDefaultDestinationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultDestination, true
}

// SetDefaultDestination sets field value
func (o *AlertDestination) SetDefaultDestination(v bool) {
	o.DefaultDestination = v
}

// GetName returns the Name field value
func (o *AlertDestination) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertDestination) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlertDestination) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *AlertDestination) GetUuid() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *AlertDestination) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *AlertDestination) SetUuid(v string) {
	o.Uuid = v
}

func (o AlertDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channels"] = o.Channels
	}
	if true {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if true {
		toSerialize["defaultDestination"] = o.DefaultDestination
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlertDestination struct {
	value *AlertDestination
	isSet bool
}

func (v NullableAlertDestination) Get() *AlertDestination {
	return v.value
}

func (v *NullableAlertDestination) Set(val *AlertDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertDestination(val *AlertDestination) *NullableAlertDestination {
	return &NullableAlertDestination{value: val, isSet: true}
}

func (v NullableAlertDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


