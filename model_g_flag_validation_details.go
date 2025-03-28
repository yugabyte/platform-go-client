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

// GFlagValidationDetails struct for GFlagValidationDetails
type GFlagValidationDetails struct {
	// WARNING: This is a preview API that could change. Error message if gflag is invalid
	Error *string `json:"error,omitempty"`
	// WARNING: This is a preview API that could change. Flag to indicate if gflag exists
	Exist *bool `json:"exist,omitempty"`
}

// NewGFlagValidationDetails instantiates a new GFlagValidationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGFlagValidationDetails() *GFlagValidationDetails {
	this := GFlagValidationDetails{}
	return &this
}

// NewGFlagValidationDetailsWithDefaults instantiates a new GFlagValidationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGFlagValidationDetailsWithDefaults() *GFlagValidationDetails {
	this := GFlagValidationDetails{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GFlagValidationDetails) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagValidationDetails) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GFlagValidationDetails) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GFlagValidationDetails) SetError(v string) {
	o.Error = &v
}

// GetExist returns the Exist field value if set, zero value otherwise.
func (o *GFlagValidationDetails) GetExist() bool {
	if o == nil || o.Exist == nil {
		var ret bool
		return ret
	}
	return *o.Exist
}

// GetExistOk returns a tuple with the Exist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagValidationDetails) GetExistOk() (*bool, bool) {
	if o == nil || o.Exist == nil {
		return nil, false
	}
	return o.Exist, true
}

// HasExist returns a boolean if a field has been set.
func (o *GFlagValidationDetails) HasExist() bool {
	if o != nil && o.Exist != nil {
		return true
	}

	return false
}

// SetExist gets a reference to the given bool and assigns it to the Exist field.
func (o *GFlagValidationDetails) SetExist(v bool) {
	o.Exist = &v
}

func (o GFlagValidationDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Exist != nil {
		toSerialize["exist"] = o.Exist
	}
	return json.Marshal(toSerialize)
}

type NullableGFlagValidationDetails struct {
	value *GFlagValidationDetails
	isSet bool
}

func (v NullableGFlagValidationDetails) Get() *GFlagValidationDetails {
	return v.value
}

func (v *NullableGFlagValidationDetails) Set(val *GFlagValidationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGFlagValidationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGFlagValidationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGFlagValidationDetails(val *GFlagValidationDetails) *NullableGFlagValidationDetails {
	return &NullableGFlagValidationDetails{value: val, isSet: true}
}

func (v NullableGFlagValidationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGFlagValidationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


