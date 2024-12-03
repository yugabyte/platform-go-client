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

// JobInstancePagedResp Paged API query response for job instance.
type JobInstancePagedResp struct {
	Entities *[]JobInstanceInfo `json:"entities,omitempty"`
	// There are more next records if true.
	HasNext *bool `json:"has_next,omitempty"`
	// There are more previous records if true.
	HasPrev *bool `json:"has_prev,omitempty"`
	// Total number of records.
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewJobInstancePagedResp instantiates a new JobInstancePagedResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobInstancePagedResp() *JobInstancePagedResp {
	this := JobInstancePagedResp{}
	return &this
}

// NewJobInstancePagedRespWithDefaults instantiates a new JobInstancePagedResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobInstancePagedRespWithDefaults() *JobInstancePagedResp {
	this := JobInstancePagedResp{}
	return &this
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *JobInstancePagedResp) GetEntities() []JobInstanceInfo {
	if o == nil || o.Entities == nil {
		var ret []JobInstanceInfo
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstancePagedResp) GetEntitiesOk() (*[]JobInstanceInfo, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *JobInstancePagedResp) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []JobInstanceInfo and assigns it to the Entities field.
func (o *JobInstancePagedResp) SetEntities(v []JobInstanceInfo) {
	o.Entities = &v
}

// GetHasNext returns the HasNext field value if set, zero value otherwise.
func (o *JobInstancePagedResp) GetHasNext() bool {
	if o == nil || o.HasNext == nil {
		var ret bool
		return ret
	}
	return *o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstancePagedResp) GetHasNextOk() (*bool, bool) {
	if o == nil || o.HasNext == nil {
		return nil, false
	}
	return o.HasNext, true
}

// HasHasNext returns a boolean if a field has been set.
func (o *JobInstancePagedResp) HasHasNext() bool {
	if o != nil && o.HasNext != nil {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given bool and assigns it to the HasNext field.
func (o *JobInstancePagedResp) SetHasNext(v bool) {
	o.HasNext = &v
}

// GetHasPrev returns the HasPrev field value if set, zero value otherwise.
func (o *JobInstancePagedResp) GetHasPrev() bool {
	if o == nil || o.HasPrev == nil {
		var ret bool
		return ret
	}
	return *o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstancePagedResp) GetHasPrevOk() (*bool, bool) {
	if o == nil || o.HasPrev == nil {
		return nil, false
	}
	return o.HasPrev, true
}

// HasHasPrev returns a boolean if a field has been set.
func (o *JobInstancePagedResp) HasHasPrev() bool {
	if o != nil && o.HasPrev != nil {
		return true
	}

	return false
}

// SetHasPrev gets a reference to the given bool and assigns it to the HasPrev field.
func (o *JobInstancePagedResp) SetHasPrev(v bool) {
	o.HasPrev = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *JobInstancePagedResp) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInstancePagedResp) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *JobInstancePagedResp) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *JobInstancePagedResp) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o JobInstancePagedResp) MarshalJSON() ([]byte, error) {
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

type NullableJobInstancePagedResp struct {
	value *JobInstancePagedResp
	isSet bool
}

func (v NullableJobInstancePagedResp) Get() *JobInstancePagedResp {
	return v.value
}

func (v *NullableJobInstancePagedResp) Set(val *JobInstancePagedResp) {
	v.value = val
	v.isSet = true
}

func (v NullableJobInstancePagedResp) IsSet() bool {
	return v.isSet
}

func (v *NullableJobInstancePagedResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobInstancePagedResp(val *JobInstancePagedResp) *NullableJobInstancePagedResp {
	return &NullableJobInstancePagedResp{value: val, isSet: true}
}

func (v NullableJobInstancePagedResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobInstancePagedResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

