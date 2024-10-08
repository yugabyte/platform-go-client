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

// YBPCreateSuccess struct for YBPCreateSuccess
type YBPCreateSuccess struct {
	// UUID of the successfully created resource
	ResourceUUID *string `json:"resourceUUID,omitempty"`
}

// NewYBPCreateSuccess instantiates a new YBPCreateSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYBPCreateSuccess() *YBPCreateSuccess {
	this := YBPCreateSuccess{}
	return &this
}

// NewYBPCreateSuccessWithDefaults instantiates a new YBPCreateSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYBPCreateSuccessWithDefaults() *YBPCreateSuccess {
	this := YBPCreateSuccess{}
	return &this
}

// GetResourceUUID returns the ResourceUUID field value if set, zero value otherwise.
func (o *YBPCreateSuccess) GetResourceUUID() string {
	if o == nil || o.ResourceUUID == nil {
		var ret string
		return ret
	}
	return *o.ResourceUUID
}

// GetResourceUUIDOk returns a tuple with the ResourceUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YBPCreateSuccess) GetResourceUUIDOk() (*string, bool) {
	if o == nil || o.ResourceUUID == nil {
		return nil, false
	}
	return o.ResourceUUID, true
}

// HasResourceUUID returns a boolean if a field has been set.
func (o *YBPCreateSuccess) HasResourceUUID() bool {
	if o != nil && o.ResourceUUID != nil {
		return true
	}

	return false
}

// SetResourceUUID gets a reference to the given string and assigns it to the ResourceUUID field.
func (o *YBPCreateSuccess) SetResourceUUID(v string) {
	o.ResourceUUID = &v
}

func (o YBPCreateSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceUUID != nil {
		toSerialize["resourceUUID"] = o.ResourceUUID
	}
	return json.Marshal(toSerialize)
}

type NullableYBPCreateSuccess struct {
	value *YBPCreateSuccess
	isSet bool
}

func (v NullableYBPCreateSuccess) Get() *YBPCreateSuccess {
	return v.value
}

func (v *NullableYBPCreateSuccess) Set(val *YBPCreateSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableYBPCreateSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableYBPCreateSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYBPCreateSuccess(val *YBPCreateSuccess) *NullableYBPCreateSuccess {
	return &NullableYBPCreateSuccess{value: val, isSet: true}
}

func (v NullableYBPCreateSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYBPCreateSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


