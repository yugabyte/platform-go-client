/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// JobSchedulePagedResp Paged API query response for job schedule.
type JobSchedulePagedResp struct {
	Entities *[]JobSchedule `json:"entities,omitempty"`
	// There are more next records if true.
	HasNext *bool `json:"has_next,omitempty"`
	// There are more previous records if true.
	HasPrev *bool `json:"has_prev,omitempty"`
	// Total number of records.
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewJobSchedulePagedResp instantiates a new JobSchedulePagedResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobSchedulePagedResp() *JobSchedulePagedResp {
	this := JobSchedulePagedResp{}
	return &this
}

// NewJobSchedulePagedRespWithDefaults instantiates a new JobSchedulePagedResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobSchedulePagedRespWithDefaults() *JobSchedulePagedResp {
	this := JobSchedulePagedResp{}
	return &this
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *JobSchedulePagedResp) GetEntities() []JobSchedule {
	if o == nil || o.Entities == nil {
		var ret []JobSchedule
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobSchedulePagedResp) GetEntitiesOk() (*[]JobSchedule, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *JobSchedulePagedResp) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []JobSchedule and assigns it to the Entities field.
func (o *JobSchedulePagedResp) SetEntities(v []JobSchedule) {
	o.Entities = &v
}

// GetHasNext returns the HasNext field value if set, zero value otherwise.
func (o *JobSchedulePagedResp) GetHasNext() bool {
	if o == nil || o.HasNext == nil {
		var ret bool
		return ret
	}
	return *o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobSchedulePagedResp) GetHasNextOk() (*bool, bool) {
	if o == nil || o.HasNext == nil {
		return nil, false
	}
	return o.HasNext, true
}

// HasHasNext returns a boolean if a field has been set.
func (o *JobSchedulePagedResp) HasHasNext() bool {
	if o != nil && o.HasNext != nil {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given bool and assigns it to the HasNext field.
func (o *JobSchedulePagedResp) SetHasNext(v bool) {
	o.HasNext = &v
}

// GetHasPrev returns the HasPrev field value if set, zero value otherwise.
func (o *JobSchedulePagedResp) GetHasPrev() bool {
	if o == nil || o.HasPrev == nil {
		var ret bool
		return ret
	}
	return *o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobSchedulePagedResp) GetHasPrevOk() (*bool, bool) {
	if o == nil || o.HasPrev == nil {
		return nil, false
	}
	return o.HasPrev, true
}

// HasHasPrev returns a boolean if a field has been set.
func (o *JobSchedulePagedResp) HasHasPrev() bool {
	if o != nil && o.HasPrev != nil {
		return true
	}

	return false
}

// SetHasPrev gets a reference to the given bool and assigns it to the HasPrev field.
func (o *JobSchedulePagedResp) SetHasPrev(v bool) {
	o.HasPrev = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *JobSchedulePagedResp) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobSchedulePagedResp) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *JobSchedulePagedResp) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *JobSchedulePagedResp) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o JobSchedulePagedResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.HasNext != nil {
		toSerialize["has_next"] = o.HasNext
	}
	if o.HasPrev != nil {
		toSerialize["has_prev"] = o.HasPrev
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableJobSchedulePagedResp struct {
	value *JobSchedulePagedResp
	isSet bool
}

func (v NullableJobSchedulePagedResp) Get() *JobSchedulePagedResp {
	return v.value
}

func (v *NullableJobSchedulePagedResp) Set(val *JobSchedulePagedResp) {
	v.value = val
	v.isSet = true
}

func (v NullableJobSchedulePagedResp) IsSet() bool {
	return v.isSet
}

func (v *NullableJobSchedulePagedResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobSchedulePagedResp(val *JobSchedulePagedResp) *NullableJobSchedulePagedResp {
	return &NullableJobSchedulePagedResp{value: val, isSet: true}
}

func (v NullableJobSchedulePagedResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobSchedulePagedResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


