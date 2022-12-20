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

// MaintenanceWindowPagedResponse struct for MaintenanceWindowPagedResponse
type MaintenanceWindowPagedResponse struct {
	Entities []MaintenanceWindow `json:"entities"`
	HasNext bool `json:"hasNext"`
	HasPrev bool `json:"hasPrev"`
	TotalCount int32 `json:"totalCount"`
}

// NewMaintenanceWindowPagedResponse instantiates a new MaintenanceWindowPagedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindowPagedResponse(entities []MaintenanceWindow, hasNext bool, hasPrev bool, totalCount int32, ) *MaintenanceWindowPagedResponse {
	this := MaintenanceWindowPagedResponse{}
	this.Entities = entities
	this.HasNext = hasNext
	this.HasPrev = hasPrev
	this.TotalCount = totalCount
	return &this
}

// NewMaintenanceWindowPagedResponseWithDefaults instantiates a new MaintenanceWindowPagedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowPagedResponseWithDefaults() *MaintenanceWindowPagedResponse {
	this := MaintenanceWindowPagedResponse{}
	return &this
}

// GetEntities returns the Entities field value
func (o *MaintenanceWindowPagedResponse) GetEntities() []MaintenanceWindow {
	if o == nil  {
		var ret []MaintenanceWindow
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowPagedResponse) GetEntitiesOk() (*[]MaintenanceWindow, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *MaintenanceWindowPagedResponse) SetEntities(v []MaintenanceWindow) {
	o.Entities = v
}

// GetHasNext returns the HasNext field value
func (o *MaintenanceWindowPagedResponse) GetHasNext() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowPagedResponse) GetHasNextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *MaintenanceWindowPagedResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrev returns the HasPrev field value
func (o *MaintenanceWindowPagedResponse) GetHasPrev() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowPagedResponse) GetHasPrevOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasPrev, true
}

// SetHasPrev sets field value
func (o *MaintenanceWindowPagedResponse) SetHasPrev(v bool) {
	o.HasPrev = v
}

// GetTotalCount returns the TotalCount field value
func (o *MaintenanceWindowPagedResponse) GetTotalCount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowPagedResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *MaintenanceWindowPagedResponse) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o MaintenanceWindowPagedResponse) MarshalJSON() ([]byte, error) {
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

type NullableMaintenanceWindowPagedResponse struct {
	value *MaintenanceWindowPagedResponse
	isSet bool
}

func (v NullableMaintenanceWindowPagedResponse) Get() *MaintenanceWindowPagedResponse {
	return v.value
}

func (v *NullableMaintenanceWindowPagedResponse) Set(val *MaintenanceWindowPagedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindowPagedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindowPagedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindowPagedResponse(val *MaintenanceWindowPagedResponse) *NullableMaintenanceWindowPagedResponse {
	return &NullableMaintenanceWindowPagedResponse{value: val, isSet: true}
}

func (v NullableMaintenanceWindowPagedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindowPagedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


