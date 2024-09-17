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

// NodeAgentPagedApiResponse struct for NodeAgentPagedApiResponse
type NodeAgentPagedApiResponse struct {
	Entities []NodeAgentResp `json:"entities"`
	HasNext bool `json:"hasNext"`
	HasPrev bool `json:"hasPrev"`
	TotalCount int32 `json:"totalCount"`
}

// NewNodeAgentPagedApiResponse instantiates a new NodeAgentPagedApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeAgentPagedApiResponse(entities []NodeAgentResp, hasNext bool, hasPrev bool, totalCount int32) *NodeAgentPagedApiResponse {
	this := NodeAgentPagedApiResponse{}
	this.Entities = entities
	this.HasNext = hasNext
	this.HasPrev = hasPrev
	this.TotalCount = totalCount
	return &this
}

// NewNodeAgentPagedApiResponseWithDefaults instantiates a new NodeAgentPagedApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeAgentPagedApiResponseWithDefaults() *NodeAgentPagedApiResponse {
	this := NodeAgentPagedApiResponse{}
	return &this
}

// GetEntities returns the Entities field value
func (o *NodeAgentPagedApiResponse) GetEntities() []NodeAgentResp {
	if o == nil {
		var ret []NodeAgentResp
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiResponse) GetEntitiesOk() (*[]NodeAgentResp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *NodeAgentPagedApiResponse) SetEntities(v []NodeAgentResp) {
	o.Entities = v
}

// GetHasNext returns the HasNext field value
func (o *NodeAgentPagedApiResponse) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiResponse) GetHasNextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *NodeAgentPagedApiResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrev returns the HasPrev field value
func (o *NodeAgentPagedApiResponse) GetHasPrev() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiResponse) GetHasPrevOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasPrev, true
}

// SetHasPrev sets field value
func (o *NodeAgentPagedApiResponse) SetHasPrev(v bool) {
	o.HasPrev = v
}

// GetTotalCount returns the TotalCount field value
func (o *NodeAgentPagedApiResponse) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *NodeAgentPagedApiResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *NodeAgentPagedApiResponse) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o NodeAgentPagedApiResponse) MarshalJSON() ([]byte, error) {
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

type NullableNodeAgentPagedApiResponse struct {
	value *NodeAgentPagedApiResponse
	isSet bool
}

func (v NullableNodeAgentPagedApiResponse) Get() *NodeAgentPagedApiResponse {
	return v.value
}

func (v *NullableNodeAgentPagedApiResponse) Set(val *NodeAgentPagedApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeAgentPagedApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeAgentPagedApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeAgentPagedApiResponse(val *NodeAgentPagedApiResponse) *NullableNodeAgentPagedApiResponse {
	return &NullableNodeAgentPagedApiResponse{value: val, isSet: true}
}

func (v NullableNodeAgentPagedApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeAgentPagedApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


