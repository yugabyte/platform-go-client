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

// RestorePagedApiResponse struct for RestorePagedApiResponse
type RestorePagedApiResponse struct {
	Entities []RestoreResp `json:"entities"`
	HasNext bool `json:"hasNext"`
	HasPrev bool `json:"hasPrev"`
	TotalCount int32 `json:"totalCount"`
}

// NewRestorePagedApiResponse instantiates a new RestorePagedApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestorePagedApiResponse(entities []RestoreResp, hasNext bool, hasPrev bool, totalCount int32, ) *RestorePagedApiResponse {
	this := RestorePagedApiResponse{}
	this.Entities = entities
	this.HasNext = hasNext
	this.HasPrev = hasPrev
	this.TotalCount = totalCount
	return &this
}

// NewRestorePagedApiResponseWithDefaults instantiates a new RestorePagedApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestorePagedApiResponseWithDefaults() *RestorePagedApiResponse {
	this := RestorePagedApiResponse{}
	return &this
}

// GetEntities returns the Entities field value
func (o *RestorePagedApiResponse) GetEntities() []RestoreResp {
	if o == nil  {
		var ret []RestoreResp
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiResponse) GetEntitiesOk() (*[]RestoreResp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *RestorePagedApiResponse) SetEntities(v []RestoreResp) {
	o.Entities = v
}

// GetHasNext returns the HasNext field value
func (o *RestorePagedApiResponse) GetHasNext() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiResponse) GetHasNextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *RestorePagedApiResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrev returns the HasPrev field value
func (o *RestorePagedApiResponse) GetHasPrev() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiResponse) GetHasPrevOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasPrev, true
}

// SetHasPrev sets field value
func (o *RestorePagedApiResponse) SetHasPrev(v bool) {
	o.HasPrev = v
}

// GetTotalCount returns the TotalCount field value
func (o *RestorePagedApiResponse) GetTotalCount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *RestorePagedApiResponse) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o RestorePagedApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entities"] = o.Entities
	}
	if true {
		toSerialize["hasNext"] = o.HasNext
	}
	if true {
		toSerialize["hasPrev"] = o.HasPrev
	}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableRestorePagedApiResponse struct {
	value *RestorePagedApiResponse
	isSet bool
}

func (v NullableRestorePagedApiResponse) Get() *RestorePagedApiResponse {
	return v.value
}

func (v *NullableRestorePagedApiResponse) Set(val *RestorePagedApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRestorePagedApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRestorePagedApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestorePagedApiResponse(val *RestorePagedApiResponse) *NullableRestorePagedApiResponse {
	return &NullableRestorePagedApiResponse{value: val, isSet: true}
}

func (v NullableRestorePagedApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestorePagedApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

