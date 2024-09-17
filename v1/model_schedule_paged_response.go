/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// SchedulePagedResponse struct for SchedulePagedResponse
type SchedulePagedResponse struct {
	Entities []Schedule `json:"entities"`
	HasNext bool `json:"hasNext"`
	HasPrev bool `json:"hasPrev"`
	TotalCount int32 `json:"totalCount"`
}

// NewSchedulePagedResponse instantiates a new SchedulePagedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePagedResponse(entities []Schedule, hasNext bool, hasPrev bool, totalCount int32) *SchedulePagedResponse {
	this := SchedulePagedResponse{}
	this.Entities = entities
	this.HasNext = hasNext
	this.HasPrev = hasPrev
	this.TotalCount = totalCount
	return &this
}

// NewSchedulePagedResponseWithDefaults instantiates a new SchedulePagedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePagedResponseWithDefaults() *SchedulePagedResponse {
	this := SchedulePagedResponse{}
	return &this
}

// GetEntities returns the Entities field value
func (o *SchedulePagedResponse) GetEntities() []Schedule {
	if o == nil {
		var ret []Schedule
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *SchedulePagedResponse) GetEntitiesOk() (*[]Schedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *SchedulePagedResponse) SetEntities(v []Schedule) {
	o.Entities = v
}

// GetHasNext returns the HasNext field value
func (o *SchedulePagedResponse) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *SchedulePagedResponse) GetHasNextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *SchedulePagedResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrev returns the HasPrev field value
func (o *SchedulePagedResponse) GetHasPrev() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value
// and a boolean to check if the value has been set.
func (o *SchedulePagedResponse) GetHasPrevOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasPrev, true
}

// SetHasPrev sets field value
func (o *SchedulePagedResponse) SetHasPrev(v bool) {
	o.HasPrev = v
}

// GetTotalCount returns the TotalCount field value
func (o *SchedulePagedResponse) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SchedulePagedResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SchedulePagedResponse) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o SchedulePagedResponse) MarshalJSON() ([]byte, error) {
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

type NullableSchedulePagedResponse struct {
	value *SchedulePagedResponse
	isSet bool
}

func (v NullableSchedulePagedResponse) Get() *SchedulePagedResponse {
	return v.value
}

func (v *NullableSchedulePagedResponse) Set(val *SchedulePagedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePagedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePagedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePagedResponse(val *SchedulePagedResponse) *NullableSchedulePagedResponse {
	return &NullableSchedulePagedResponse{value: val, isSet: true}
}

func (v NullableSchedulePagedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePagedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


