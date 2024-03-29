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

// YbcThrottleParametersResponse YB-Controller throttle parameters response
type YbcThrottleParametersResponse struct {
	// Map of YBC throttle parameters
	ThrottleParamsMap *map[string]ThrottleParamValue `json:"throttleParamsMap,omitempty"`
}

// NewYbcThrottleParametersResponse instantiates a new YbcThrottleParametersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYbcThrottleParametersResponse() *YbcThrottleParametersResponse {
	this := YbcThrottleParametersResponse{}
	return &this
}

// NewYbcThrottleParametersResponseWithDefaults instantiates a new YbcThrottleParametersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYbcThrottleParametersResponseWithDefaults() *YbcThrottleParametersResponse {
	this := YbcThrottleParametersResponse{}
	return &this
}

// GetThrottleParamsMap returns the ThrottleParamsMap field value if set, zero value otherwise.
func (o *YbcThrottleParametersResponse) GetThrottleParamsMap() map[string]ThrottleParamValue {
	if o == nil || o.ThrottleParamsMap == nil {
		var ret map[string]ThrottleParamValue
		return ret
	}
	return *o.ThrottleParamsMap
}

// GetThrottleParamsMapOk returns a tuple with the ThrottleParamsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YbcThrottleParametersResponse) GetThrottleParamsMapOk() (*map[string]ThrottleParamValue, bool) {
	if o == nil || o.ThrottleParamsMap == nil {
		return nil, false
	}
	return o.ThrottleParamsMap, true
}

// HasThrottleParamsMap returns a boolean if a field has been set.
func (o *YbcThrottleParametersResponse) HasThrottleParamsMap() bool {
	if o != nil && o.ThrottleParamsMap != nil {
		return true
	}

	return false
}

// SetThrottleParamsMap gets a reference to the given map[string]ThrottleParamValue and assigns it to the ThrottleParamsMap field.
func (o *YbcThrottleParametersResponse) SetThrottleParamsMap(v map[string]ThrottleParamValue) {
	o.ThrottleParamsMap = &v
}

func (o YbcThrottleParametersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ThrottleParamsMap != nil {
		toSerialize["throttleParamsMap"] = o.ThrottleParamsMap
	}
	return json.Marshal(toSerialize)
}

type NullableYbcThrottleParametersResponse struct {
	value *YbcThrottleParametersResponse
	isSet bool
}

func (v NullableYbcThrottleParametersResponse) Get() *YbcThrottleParametersResponse {
	return v.value
}

func (v *NullableYbcThrottleParametersResponse) Set(val *YbcThrottleParametersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableYbcThrottleParametersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableYbcThrottleParametersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYbcThrottleParametersResponse(val *YbcThrottleParametersResponse) *NullableYbcThrottleParametersResponse {
	return &NullableYbcThrottleParametersResponse{value: val, isSet: true}
}

func (v NullableYbcThrottleParametersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYbcThrottleParametersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


