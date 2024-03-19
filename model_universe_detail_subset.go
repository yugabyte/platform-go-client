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

// UniverseDetailSubset A small subset of universe information
type UniverseDetailSubset struct {
	CreationDate int64 `json:"creationDate"`
	Name string `json:"name"`
	UniversePaused bool `json:"universePaused"`
	UpdateInProgress bool `json:"updateInProgress"`
	UpdateSucceeded bool `json:"updateSucceeded"`
	Uuid string `json:"uuid"`
}

// NewUniverseDetailSubset instantiates a new UniverseDetailSubset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseDetailSubset(creationDate int64, name string, universePaused bool, updateInProgress bool, updateSucceeded bool, uuid string) *UniverseDetailSubset {
	this := UniverseDetailSubset{}
	this.CreationDate = creationDate
	this.Name = name
	this.UniversePaused = universePaused
	this.UpdateInProgress = updateInProgress
	this.UpdateSucceeded = updateSucceeded
	this.Uuid = uuid
	return &this
}

// NewUniverseDetailSubsetWithDefaults instantiates a new UniverseDetailSubset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseDetailSubsetWithDefaults() *UniverseDetailSubset {
	this := UniverseDetailSubset{}
	return &this
}

// GetCreationDate returns the CreationDate field value
func (o *UniverseDetailSubset) GetCreationDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value
// and a boolean to check if the value has been set.
func (o *UniverseDetailSubset) GetCreationDateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreationDate, true
}

// SetCreationDate sets field value
func (o *UniverseDetailSubset) SetCreationDate(v int64) {
	o.CreationDate = v
}

// GetName returns the Name field value
func (o *UniverseDetailSubset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UniverseDetailSubset) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UniverseDetailSubset) SetName(v string) {
	o.Name = v
}

// GetUniversePaused returns the UniversePaused field value
func (o *UniverseDetailSubset) GetUniversePaused() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UniversePaused
}

// GetUniversePausedOk returns a tuple with the UniversePaused field value
// and a boolean to check if the value has been set.
func (o *UniverseDetailSubset) GetUniversePausedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniversePaused, true
}

// SetUniversePaused sets field value
func (o *UniverseDetailSubset) SetUniversePaused(v bool) {
	o.UniversePaused = v
}

// GetUpdateInProgress returns the UpdateInProgress field value
func (o *UniverseDetailSubset) GetUpdateInProgress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UpdateInProgress
}

// GetUpdateInProgressOk returns a tuple with the UpdateInProgress field value
// and a boolean to check if the value has been set.
func (o *UniverseDetailSubset) GetUpdateInProgressOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateInProgress, true
}

// SetUpdateInProgress sets field value
func (o *UniverseDetailSubset) SetUpdateInProgress(v bool) {
	o.UpdateInProgress = v
}

// GetUpdateSucceeded returns the UpdateSucceeded field value
func (o *UniverseDetailSubset) GetUpdateSucceeded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UpdateSucceeded
}

// GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field value
// and a boolean to check if the value has been set.
func (o *UniverseDetailSubset) GetUpdateSucceededOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateSucceeded, true
}

// SetUpdateSucceeded sets field value
func (o *UniverseDetailSubset) SetUpdateSucceeded(v bool) {
	o.UpdateSucceeded = v
}

// GetUuid returns the Uuid field value
func (o *UniverseDetailSubset) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *UniverseDetailSubset) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *UniverseDetailSubset) SetUuid(v string) {
	o.Uuid = v
}

func (o UniverseDetailSubset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["creationDate"] = o.CreationDate
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["universePaused"] = o.UniversePaused
	}
	if true {
		toSerialize["updateInProgress"] = o.UpdateInProgress
	}
	if true {
		toSerialize["updateSucceeded"] = o.UpdateSucceeded
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseDetailSubset struct {
	value *UniverseDetailSubset
	isSet bool
}

func (v NullableUniverseDetailSubset) Get() *UniverseDetailSubset {
	return v.value
}

func (v *NullableUniverseDetailSubset) Set(val *UniverseDetailSubset) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseDetailSubset) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseDetailSubset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseDetailSubset(val *UniverseDetailSubset) *NullableUniverseDetailSubset {
	return &NullableUniverseDetailSubset{value: val, isSet: true}
}

func (v NullableUniverseDetailSubset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseDetailSubset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


