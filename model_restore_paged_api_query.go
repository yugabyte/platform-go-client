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

// RestorePagedApiQuery struct for RestorePagedApiQuery
type RestorePagedApiQuery struct {
	Direction string `json:"direction"`
	Filter RestoreApiFilter `json:"filter"`
	Limit int32 `json:"limit"`
	NeedTotalCount bool `json:"needTotalCount"`
	Offset int32 `json:"offset"`
	SortBy string `json:"sortBy"`
}

// NewRestorePagedApiQuery instantiates a new RestorePagedApiQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestorePagedApiQuery(direction string, filter RestoreApiFilter, limit int32, needTotalCount bool, offset int32, sortBy string) *RestorePagedApiQuery {
	this := RestorePagedApiQuery{}
	this.Direction = direction
	this.Filter = filter
	this.Limit = limit
	this.NeedTotalCount = needTotalCount
	this.Offset = offset
	this.SortBy = sortBy
	return &this
}

// NewRestorePagedApiQueryWithDefaults instantiates a new RestorePagedApiQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestorePagedApiQueryWithDefaults() *RestorePagedApiQuery {
	this := RestorePagedApiQuery{}
	return &this
}

// GetDirection returns the Direction field value
func (o *RestorePagedApiQuery) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiQuery) GetDirectionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *RestorePagedApiQuery) SetDirection(v string) {
	o.Direction = v
}

// GetFilter returns the Filter field value
func (o *RestorePagedApiQuery) GetFilter() RestoreApiFilter {
	if o == nil {
		var ret RestoreApiFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiQuery) GetFilterOk() (*RestoreApiFilter, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *RestorePagedApiQuery) SetFilter(v RestoreApiFilter) {
	o.Filter = v
}

// GetLimit returns the Limit field value
func (o *RestorePagedApiQuery) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiQuery) GetLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *RestorePagedApiQuery) SetLimit(v int32) {
	o.Limit = v
}

// GetNeedTotalCount returns the NeedTotalCount field value
func (o *RestorePagedApiQuery) GetNeedTotalCount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeedTotalCount
}

// GetNeedTotalCountOk returns a tuple with the NeedTotalCount field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiQuery) GetNeedTotalCountOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NeedTotalCount, true
}

// SetNeedTotalCount sets field value
func (o *RestorePagedApiQuery) SetNeedTotalCount(v bool) {
	o.NeedTotalCount = v
}

// GetOffset returns the Offset field value
func (o *RestorePagedApiQuery) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiQuery) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *RestorePagedApiQuery) SetOffset(v int32) {
	o.Offset = v
}

// GetSortBy returns the SortBy field value
func (o *RestorePagedApiQuery) GetSortBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value
// and a boolean to check if the value has been set.
func (o *RestorePagedApiQuery) GetSortByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SortBy, true
}

// SetSortBy sets field value
func (o *RestorePagedApiQuery) SetSortBy(v string) {
	o.SortBy = v
}

func (o RestorePagedApiQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["direction"] = o.Direction
	}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["needTotalCount"] = o.NeedTotalCount
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["sortBy"] = o.SortBy
	}
	return json.Marshal(toSerialize)
}

type NullableRestorePagedApiQuery struct {
	value *RestorePagedApiQuery
	isSet bool
}

func (v NullableRestorePagedApiQuery) Get() *RestorePagedApiQuery {
	return v.value
}

func (v *NullableRestorePagedApiQuery) Set(val *RestorePagedApiQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRestorePagedApiQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRestorePagedApiQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestorePagedApiQuery(val *RestorePagedApiQuery) *NullableRestorePagedApiQuery {
	return &NullableRestorePagedApiQuery{value: val, isSet: true}
}

func (v NullableRestorePagedApiQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestorePagedApiQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


