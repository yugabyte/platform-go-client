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

// NodeAgentPagedApiQuery struct for NodeAgentPagedApiQuery
type NodeAgentPagedApiQuery struct {
	Direction string `json:"direction"`
	Filter NodeAgentApiFilter `json:"filter"`
	Limit int32 `json:"limit"`
	NeedTotalCount bool `json:"needTotalCount"`
	Offset int32 `json:"offset"`
	SortBy string `json:"sortBy"`
}

// NewNodeAgentPagedApiQuery instantiates a new NodeAgentPagedApiQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeAgentPagedApiQuery(direction string, filter NodeAgentApiFilter, limit int32, needTotalCount bool, offset int32, sortBy string, ) *NodeAgentPagedApiQuery {
	this := NodeAgentPagedApiQuery{}
	this.Direction = direction
	this.Filter = filter
	this.Limit = limit
	this.NeedTotalCount = needTotalCount
	this.Offset = offset
	this.SortBy = sortBy
	return &this
}

// NewNodeAgentPagedApiQueryWithDefaults instantiates a new NodeAgentPagedApiQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeAgentPagedApiQueryWithDefaults() *NodeAgentPagedApiQuery {
	this := NodeAgentPagedApiQuery{}
	return &this
}

// GetDirection returns the Direction field value
func (o *NodeAgentPagedApiQuery) GetDirection() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiQuery) GetDirectionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *NodeAgentPagedApiQuery) SetDirection(v string) {
	o.Direction = v
}

// GetFilter returns the Filter field value
func (o *NodeAgentPagedApiQuery) GetFilter() NodeAgentApiFilter {
	if o == nil  {
		var ret NodeAgentApiFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiQuery) GetFilterOk() (*NodeAgentApiFilter, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *NodeAgentPagedApiQuery) SetFilter(v NodeAgentApiFilter) {
	o.Filter = v
}

// GetLimit returns the Limit field value
func (o *NodeAgentPagedApiQuery) GetLimit() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiQuery) GetLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *NodeAgentPagedApiQuery) SetLimit(v int32) {
	o.Limit = v
}

// GetNeedTotalCount returns the NeedTotalCount field value
func (o *NodeAgentPagedApiQuery) GetNeedTotalCount() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.NeedTotalCount
}

// GetNeedTotalCountOk returns a tuple with the NeedTotalCount field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiQuery) GetNeedTotalCountOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NeedTotalCount, true
}

// SetNeedTotalCount sets field value
func (o *NodeAgentPagedApiQuery) SetNeedTotalCount(v bool) {
	o.NeedTotalCount = v
}

// GetOffset returns the Offset field value
func (o *NodeAgentPagedApiQuery) GetOffset() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiQuery) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *NodeAgentPagedApiQuery) SetOffset(v int32) {
	o.Offset = v
}

// GetSortBy returns the SortBy field value
func (o *NodeAgentPagedApiQuery) GetSortBy() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiQuery) GetSortByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SortBy, true
}

// SetSortBy sets field value
func (o *NodeAgentPagedApiQuery) SetSortBy(v string) {
	o.SortBy = v
}

func (o NodeAgentPagedApiQuery) MarshalJSON() ([]byte, error) {
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

type NullableNodeAgentPagedApiQuery struct {
	value *NodeAgentPagedApiQuery
	isSet bool
}

func (v NullableNodeAgentPagedApiQuery) Get() *NodeAgentPagedApiQuery {
	return v.value
}

func (v *NullableNodeAgentPagedApiQuery) Set(val *NodeAgentPagedApiQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeAgentPagedApiQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeAgentPagedApiQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeAgentPagedApiQuery(val *NodeAgentPagedApiQuery) *NullableNodeAgentPagedApiQuery {
	return &NullableNodeAgentPagedApiQuery{value: val, isSet: true}
}

func (v NullableNodeAgentPagedApiQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeAgentPagedApiQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

