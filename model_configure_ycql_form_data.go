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

// ConfigureYCQLFormData YCQL properties
type ConfigureYCQLFormData struct {
	CommunicationPorts *CommunicationPorts `json:"communicationPorts,omitempty"`
	// Enable YCQL Api for the universe
	EnableYCQL *bool `json:"enableYCQL,omitempty"`
	// Enable YCQL Auth for the universe
	EnableYCQLAuth *bool `json:"enableYCQLAuth,omitempty"`
	// YCQL Auth password
	YcqlPassword *string `json:"ycqlPassword,omitempty"`
}

// NewConfigureYCQLFormData instantiates a new ConfigureYCQLFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigureYCQLFormData() *ConfigureYCQLFormData {
	this := ConfigureYCQLFormData{}
	return &this
}

// NewConfigureYCQLFormDataWithDefaults instantiates a new ConfigureYCQLFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigureYCQLFormDataWithDefaults() *ConfigureYCQLFormData {
	this := ConfigureYCQLFormData{}
	return &this
}

// GetCommunicationPorts returns the CommunicationPorts field value if set, zero value otherwise.
func (o *ConfigureYCQLFormData) GetCommunicationPorts() CommunicationPorts {
	if o == nil || o.CommunicationPorts == nil {
		var ret CommunicationPorts
		return ret
	}
	return *o.CommunicationPorts
}

// GetCommunicationPortsOk returns a tuple with the CommunicationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureYCQLFormData) GetCommunicationPortsOk() (*CommunicationPorts, bool) {
	if o == nil || o.CommunicationPorts == nil {
		return nil, false
	}
	return o.CommunicationPorts, true
}

// HasCommunicationPorts returns a boolean if a field has been set.
func (o *ConfigureYCQLFormData) HasCommunicationPorts() bool {
	if o != nil && o.CommunicationPorts != nil {
		return true
	}

	return false
}

// SetCommunicationPorts gets a reference to the given CommunicationPorts and assigns it to the CommunicationPorts field.
func (o *ConfigureYCQLFormData) SetCommunicationPorts(v CommunicationPorts) {
	o.CommunicationPorts = &v
}

// GetEnableYCQL returns the EnableYCQL field value if set, zero value otherwise.
func (o *ConfigureYCQLFormData) GetEnableYCQL() bool {
	if o == nil || o.EnableYCQL == nil {
		var ret bool
		return ret
	}
	return *o.EnableYCQL
}

// GetEnableYCQLOk returns a tuple with the EnableYCQL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureYCQLFormData) GetEnableYCQLOk() (*bool, bool) {
	if o == nil || o.EnableYCQL == nil {
		return nil, false
	}
	return o.EnableYCQL, true
}

// HasEnableYCQL returns a boolean if a field has been set.
func (o *ConfigureYCQLFormData) HasEnableYCQL() bool {
	if o != nil && o.EnableYCQL != nil {
		return true
	}

	return false
}

// SetEnableYCQL gets a reference to the given bool and assigns it to the EnableYCQL field.
func (o *ConfigureYCQLFormData) SetEnableYCQL(v bool) {
	o.EnableYCQL = &v
}

// GetEnableYCQLAuth returns the EnableYCQLAuth field value if set, zero value otherwise.
func (o *ConfigureYCQLFormData) GetEnableYCQLAuth() bool {
	if o == nil || o.EnableYCQLAuth == nil {
		var ret bool
		return ret
	}
	return *o.EnableYCQLAuth
}

// GetEnableYCQLAuthOk returns a tuple with the EnableYCQLAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureYCQLFormData) GetEnableYCQLAuthOk() (*bool, bool) {
	if o == nil || o.EnableYCQLAuth == nil {
		return nil, false
	}
	return o.EnableYCQLAuth, true
}

// HasEnableYCQLAuth returns a boolean if a field has been set.
func (o *ConfigureYCQLFormData) HasEnableYCQLAuth() bool {
	if o != nil && o.EnableYCQLAuth != nil {
		return true
	}

	return false
}

// SetEnableYCQLAuth gets a reference to the given bool and assigns it to the EnableYCQLAuth field.
func (o *ConfigureYCQLFormData) SetEnableYCQLAuth(v bool) {
	o.EnableYCQLAuth = &v
}

// GetYcqlPassword returns the YcqlPassword field value if set, zero value otherwise.
func (o *ConfigureYCQLFormData) GetYcqlPassword() string {
	if o == nil || o.YcqlPassword == nil {
		var ret string
		return ret
	}
	return *o.YcqlPassword
}

// GetYcqlPasswordOk returns a tuple with the YcqlPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureYCQLFormData) GetYcqlPasswordOk() (*string, bool) {
	if o == nil || o.YcqlPassword == nil {
		return nil, false
	}
	return o.YcqlPassword, true
}

// HasYcqlPassword returns a boolean if a field has been set.
func (o *ConfigureYCQLFormData) HasYcqlPassword() bool {
	if o != nil && o.YcqlPassword != nil {
		return true
	}

	return false
}

// SetYcqlPassword gets a reference to the given string and assigns it to the YcqlPassword field.
func (o *ConfigureYCQLFormData) SetYcqlPassword(v string) {
	o.YcqlPassword = &v
}

func (o ConfigureYCQLFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommunicationPorts != nil {
		toSerialize["communicationPorts"] = o.CommunicationPorts
	}
	if o.EnableYCQL != nil {
		toSerialize["enableYCQL"] = o.EnableYCQL
	}
	if o.EnableYCQLAuth != nil {
		toSerialize["enableYCQLAuth"] = o.EnableYCQLAuth
	}
	if o.YcqlPassword != nil {
		toSerialize["ycqlPassword"] = o.YcqlPassword
	}
	return json.Marshal(toSerialize)
}

type NullableConfigureYCQLFormData struct {
	value *ConfigureYCQLFormData
	isSet bool
}

func (v NullableConfigureYCQLFormData) Get() *ConfigureYCQLFormData {
	return v.value
}

func (v *NullableConfigureYCQLFormData) Set(val *ConfigureYCQLFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigureYCQLFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigureYCQLFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigureYCQLFormData(val *ConfigureYCQLFormData) *NullableConfigureYCQLFormData {
	return &NullableConfigureYCQLFormData{value: val, isSet: true}
}

func (v NullableConfigureYCQLFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigureYCQLFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


