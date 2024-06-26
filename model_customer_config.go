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

// CustomerConfig Customer configuration. Includes storage, alerts, password policy, and call-home level.
type CustomerConfig struct {
	// Config name
	ConfigName string `json:"configName"`
	// Config UUID
	ConfigUUID *string `json:"configUUID,omitempty"`
	// Customer UUID
	CustomerUUID string `json:"customerUUID"`
	// Configuration data
	Data map[string]interface{} `json:"data"`
	// Name
	Name string `json:"name"`
	// state of the customerConfig. Possible values are Active, QueuedForDeletion.
	State *string `json:"state,omitempty"`
	// Config type
	Type string `json:"type"`
}

// NewCustomerConfig instantiates a new CustomerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerConfig(configName string, customerUUID string, data map[string]interface{}, name string, type_ string) *CustomerConfig {
	this := CustomerConfig{}
	this.ConfigName = configName
	this.CustomerUUID = customerUUID
	this.Data = data
	this.Name = name
	this.Type = type_
	return &this
}

// NewCustomerConfigWithDefaults instantiates a new CustomerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerConfigWithDefaults() *CustomerConfig {
	this := CustomerConfig{}
	return &this
}

// GetConfigName returns the ConfigName field value
func (o *CustomerConfig) GetConfigName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigName
}

// GetConfigNameOk returns a tuple with the ConfigName field value
// and a boolean to check if the value has been set.
func (o *CustomerConfig) GetConfigNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigName, true
}

// SetConfigName sets field value
func (o *CustomerConfig) SetConfigName(v string) {
	o.ConfigName = v
}

// GetConfigUUID returns the ConfigUUID field value if set, zero value otherwise.
func (o *CustomerConfig) GetConfigUUID() string {
	if o == nil || o.ConfigUUID == nil {
		var ret string
		return ret
	}
	return *o.ConfigUUID
}

// GetConfigUUIDOk returns a tuple with the ConfigUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConfig) GetConfigUUIDOk() (*string, bool) {
	if o == nil || o.ConfigUUID == nil {
		return nil, false
	}
	return o.ConfigUUID, true
}

// HasConfigUUID returns a boolean if a field has been set.
func (o *CustomerConfig) HasConfigUUID() bool {
	if o != nil && o.ConfigUUID != nil {
		return true
	}

	return false
}

// SetConfigUUID gets a reference to the given string and assigns it to the ConfigUUID field.
func (o *CustomerConfig) SetConfigUUID(v string) {
	o.ConfigUUID = &v
}

// GetCustomerUUID returns the CustomerUUID field value
func (o *CustomerConfig) GetCustomerUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value
// and a boolean to check if the value has been set.
func (o *CustomerConfig) GetCustomerUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerUUID, true
}

// SetCustomerUUID sets field value
func (o *CustomerConfig) SetCustomerUUID(v string) {
	o.CustomerUUID = v
}

// GetData returns the Data field value
func (o *CustomerConfig) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerConfig) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerConfig) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetName returns the Name field value
func (o *CustomerConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerConfig) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerConfig) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CustomerConfig) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConfig) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CustomerConfig) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CustomerConfig) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value
func (o *CustomerConfig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerConfig) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerConfig) SetType(v string) {
	o.Type = v
}

func (o CustomerConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configName"] = o.ConfigName
	}
	if o.ConfigUUID != nil {
		toSerialize["configUUID"] = o.ConfigUUID
	}
	if true {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerConfig struct {
	value *CustomerConfig
	isSet bool
}

func (v NullableCustomerConfig) Get() *CustomerConfig {
	return v.value
}

func (v *NullableCustomerConfig) Set(val *CustomerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerConfig(val *CustomerConfig) *NullableCustomerConfig {
	return &NullableCustomerConfig{value: val, isSet: true}
}

func (v NullableCustomerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


