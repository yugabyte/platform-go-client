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

// PerformanceRecommendationFilter struct for PerformanceRecommendationFilter
type PerformanceRecommendationFilter struct {
	CreatedInstantBefore int64 `json:"createdInstantBefore"`
	CustomerId string `json:"customerId"`
	Ids []string `json:"ids"`
	IsStale bool `json:"isStale"`
	Priorities []string `json:"priorities"`
	States []string `json:"states"`
	Types []string `json:"types"`
	UniverseId string `json:"universeId"`
}

// NewPerformanceRecommendationFilter instantiates a new PerformanceRecommendationFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceRecommendationFilter(createdInstantBefore int64, customerId string, ids []string, isStale bool, priorities []string, states []string, types []string, universeId string) *PerformanceRecommendationFilter {
	this := PerformanceRecommendationFilter{}
	this.CreatedInstantBefore = createdInstantBefore
	this.CustomerId = customerId
	this.Ids = ids
	this.IsStale = isStale
	this.Priorities = priorities
	this.States = states
	this.Types = types
	this.UniverseId = universeId
	return &this
}

// NewPerformanceRecommendationFilterWithDefaults instantiates a new PerformanceRecommendationFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceRecommendationFilterWithDefaults() *PerformanceRecommendationFilter {
	this := PerformanceRecommendationFilter{}
	return &this
}

// GetCreatedInstantBefore returns the CreatedInstantBefore field value
func (o *PerformanceRecommendationFilter) GetCreatedInstantBefore() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedInstantBefore
}

// GetCreatedInstantBeforeOk returns a tuple with the CreatedInstantBefore field value
// and a boolean to check if the value has been set.
func (o *PerformanceRecommendationFilter) GetCreatedInstantBeforeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedInstantBefore, true
}

// SetCreatedInstantBefore sets field value
func (o *PerformanceRecommendationFilter) SetCreatedInstantBefore(v int64) {
	o.CreatedInstantBefore = v
}

// GetCustomerId returns the CustomerId field value
func (o *PerformanceRecommendationFilter) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *PerformanceRecommendationFilter) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *PerformanceRecommendationFilter) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetIds returns the Ids field value
func (o *PerformanceRecommendationFilter) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *PerformanceRecommendationFilter) GetIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value
func (o *PerformanceRecommendationFilter) SetIds(v []string) {
	o.Ids = v
}

// GetIsStale returns the IsStale field value
func (o *PerformanceRecommendationFilter) GetIsStale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsStale
}

// GetIsStaleOk returns a tuple with the IsStale field value
// and a boolean to check if the value has been set.
func (o *PerformanceRecommendationFilter) GetIsStaleOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsStale, true
}

// SetIsStale sets field value
func (o *PerformanceRecommendationFilter) SetIsStale(v bool) {
	o.IsStale = v
}

// GetPriorities returns the Priorities field value
func (o *PerformanceRecommendationFilter) GetPriorities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value
// and a boolean to check if the value has been set.
func (o *PerformanceRecommendationFilter) GetPrioritiesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Priorities, true
}

// SetPriorities sets field value
func (o *PerformanceRecommendationFilter) SetPriorities(v []string) {
	o.Priorities = v
}

// GetStates returns the States field value
func (o *PerformanceRecommendationFilter) GetStates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.States
}

// GetStatesOk returns a tuple with the States field value
// and a boolean to check if the value has been set.
func (o *PerformanceRecommendationFilter) GetStatesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.States, true
}

// SetStates sets field value
func (o *PerformanceRecommendationFilter) SetStates(v []string) {
	o.States = v
}

// GetTypes returns the Types field value
func (o *PerformanceRecommendationFilter) GetTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *PerformanceRecommendationFilter) GetTypesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Types, true
}

// SetTypes sets field value
func (o *PerformanceRecommendationFilter) SetTypes(v []string) {
	o.Types = v
}

// GetUniverseId returns the UniverseId field value
func (o *PerformanceRecommendationFilter) GetUniverseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniverseId
}

// GetUniverseIdOk returns a tuple with the UniverseId field value
// and a boolean to check if the value has been set.
func (o *PerformanceRecommendationFilter) GetUniverseIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseId, true
}

// SetUniverseId sets field value
func (o *PerformanceRecommendationFilter) SetUniverseId(v string) {
	o.UniverseId = v
}

func (o PerformanceRecommendationFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["createdInstantBefore"] = o.CreatedInstantBefore
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["ids"] = o.Ids
	}
	if true {
		toSerialize["isStale"] = o.IsStale
	}
	if true {
		toSerialize["priorities"] = o.Priorities
	}
	if true {
		toSerialize["states"] = o.States
	}
	if true {
		toSerialize["types"] = o.Types
	}
	if true {
		toSerialize["universeId"] = o.UniverseId
	}
	return json.Marshal(toSerialize)
}

type NullablePerformanceRecommendationFilter struct {
	value *PerformanceRecommendationFilter
	isSet bool
}

func (v NullablePerformanceRecommendationFilter) Get() *PerformanceRecommendationFilter {
	return v.value
}

func (v *NullablePerformanceRecommendationFilter) Set(val *PerformanceRecommendationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceRecommendationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceRecommendationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceRecommendationFilter(val *PerformanceRecommendationFilter) *NullablePerformanceRecommendationFilter {
	return &NullablePerformanceRecommendationFilter{value: val, isSet: true}
}

func (v NullablePerformanceRecommendationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceRecommendationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


