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

// StateChangeAuditInfoPagedResponse struct for StateChangeAuditInfoPagedResponse
type StateChangeAuditInfoPagedResponse struct {
	Entities []StateChangeAuditInfo `json:"entities"`
	HasNext bool `json:"hasNext"`
	HasPrev bool `json:"hasPrev"`
	TotalCount int32 `json:"totalCount"`
}

// NewStateChangeAuditInfoPagedResponse instantiates a new StateChangeAuditInfoPagedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateChangeAuditInfoPagedResponse(entities []StateChangeAuditInfo, hasNext bool, hasPrev bool, totalCount int32) *StateChangeAuditInfoPagedResponse {
	this := StateChangeAuditInfoPagedResponse{}
	this.Entities = entities
	this.HasNext = hasNext
	this.HasPrev = hasPrev
	this.TotalCount = totalCount
	return &this
}

// NewStateChangeAuditInfoPagedResponseWithDefaults instantiates a new StateChangeAuditInfoPagedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateChangeAuditInfoPagedResponseWithDefaults() *StateChangeAuditInfoPagedResponse {
	this := StateChangeAuditInfoPagedResponse{}
	return &this
}

// GetEntities returns the Entities field value
func (o *StateChangeAuditInfoPagedResponse) GetEntities() []StateChangeAuditInfo {
	if o == nil {
		var ret []StateChangeAuditInfo
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *StateChangeAuditInfoPagedResponse) GetEntitiesOk() (*[]StateChangeAuditInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *StateChangeAuditInfoPagedResponse) SetEntities(v []StateChangeAuditInfo) {
	o.Entities = v
}

// GetHasNext returns the HasNext field value
func (o *StateChangeAuditInfoPagedResponse) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *StateChangeAuditInfoPagedResponse) GetHasNextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *StateChangeAuditInfoPagedResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrev returns the HasPrev field value
func (o *StateChangeAuditInfoPagedResponse) GetHasPrev() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value
// and a boolean to check if the value has been set.
func (o *StateChangeAuditInfoPagedResponse) GetHasPrevOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasPrev, true
}

// SetHasPrev sets field value
func (o *StateChangeAuditInfoPagedResponse) SetHasPrev(v bool) {
	o.HasPrev = v
}

// GetTotalCount returns the TotalCount field value
func (o *StateChangeAuditInfoPagedResponse) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *StateChangeAuditInfoPagedResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *StateChangeAuditInfoPagedResponse) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o StateChangeAuditInfoPagedResponse) MarshalJSON() ([]byte, error) {
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

type NullableStateChangeAuditInfoPagedResponse struct {
	value *StateChangeAuditInfoPagedResponse
	isSet bool
}

func (v NullableStateChangeAuditInfoPagedResponse) Get() *StateChangeAuditInfoPagedResponse {
	return v.value
}

func (v *NullableStateChangeAuditInfoPagedResponse) Set(val *StateChangeAuditInfoPagedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStateChangeAuditInfoPagedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStateChangeAuditInfoPagedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateChangeAuditInfoPagedResponse(val *StateChangeAuditInfoPagedResponse) *NullableStateChangeAuditInfoPagedResponse {
	return &NullableStateChangeAuditInfoPagedResponse{value: val, isSet: true}
}

func (v NullableStateChangeAuditInfoPagedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateChangeAuditInfoPagedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


