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

// GFlagsValidationRequest struct for GFlagsValidationRequest
type GFlagsValidationRequest struct {
	// WARNING: This is a preview API that could change. Value of the gflag for master
	MASTER *string `json:"MASTER,omitempty"`
	// WARNING: This is a preview API that could change. Name of the gflag
	Name *string `json:"Name,omitempty"`
	// WARNING: This is a preview API that could change. Value of the gflag for tserver
	TSERVER *string `json:"TSERVER,omitempty"`
}

// NewGFlagsValidationRequest instantiates a new GFlagsValidationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGFlagsValidationRequest() *GFlagsValidationRequest {
	this := GFlagsValidationRequest{}
	return &this
}

// NewGFlagsValidationRequestWithDefaults instantiates a new GFlagsValidationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGFlagsValidationRequestWithDefaults() *GFlagsValidationRequest {
	this := GFlagsValidationRequest{}
	return &this
}

// GetMASTER returns the MASTER field value if set, zero value otherwise.
func (o *GFlagsValidationRequest) GetMASTER() string {
	if o == nil || o.MASTER == nil {
		var ret string
		return ret
	}
	return *o.MASTER
}

// GetMASTEROk returns a tuple with the MASTER field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagsValidationRequest) GetMASTEROk() (*string, bool) {
	if o == nil || o.MASTER == nil {
		return nil, false
	}
	return o.MASTER, true
}

// HasMASTER returns a boolean if a field has been set.
func (o *GFlagsValidationRequest) HasMASTER() bool {
	if o != nil && o.MASTER != nil {
		return true
	}

	return false
}

// SetMASTER gets a reference to the given string and assigns it to the MASTER field.
func (o *GFlagsValidationRequest) SetMASTER(v string) {
	o.MASTER = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GFlagsValidationRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagsValidationRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GFlagsValidationRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GFlagsValidationRequest) SetName(v string) {
	o.Name = &v
}

// GetTSERVER returns the TSERVER field value if set, zero value otherwise.
func (o *GFlagsValidationRequest) GetTSERVER() string {
	if o == nil || o.TSERVER == nil {
		var ret string
		return ret
	}
	return *o.TSERVER
}

// GetTSERVEROk returns a tuple with the TSERVER field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagsValidationRequest) GetTSERVEROk() (*string, bool) {
	if o == nil || o.TSERVER == nil {
		return nil, false
	}
	return o.TSERVER, true
}

// HasTSERVER returns a boolean if a field has been set.
func (o *GFlagsValidationRequest) HasTSERVER() bool {
	if o != nil && o.TSERVER != nil {
		return true
	}

	return false
}

// SetTSERVER gets a reference to the given string and assigns it to the TSERVER field.
func (o *GFlagsValidationRequest) SetTSERVER(v string) {
	o.TSERVER = &v
}

func (o GFlagsValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MASTER != nil {
		toSerialize["MASTER"] = o.MASTER
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.TSERVER != nil {
		toSerialize["TSERVER"] = o.TSERVER
	}
	return json.Marshal(toSerialize)
}

type NullableGFlagsValidationRequest struct {
	value *GFlagsValidationRequest
	isSet bool
}

func (v NullableGFlagsValidationRequest) Get() *GFlagsValidationRequest {
	return v.value
}

func (v *NullableGFlagsValidationRequest) Set(val *GFlagsValidationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGFlagsValidationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGFlagsValidationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGFlagsValidationRequest(val *GFlagsValidationRequest) *NullableGFlagsValidationRequest {
	return &NullableGFlagsValidationRequest{value: val, isSet: true}
}

func (v NullableGFlagsValidationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGFlagsValidationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

