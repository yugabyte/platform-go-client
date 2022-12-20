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

// BackupPagedApiResponse struct for BackupPagedApiResponse
type BackupPagedApiResponse struct {
	Entities []BackupResp `json:"entities"`
	HasNext bool `json:"hasNext"`
	HasPrev bool `json:"hasPrev"`
	TotalCount int32 `json:"totalCount"`
}

// NewBackupPagedApiResponse instantiates a new BackupPagedApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupPagedApiResponse(entities []BackupResp, hasNext bool, hasPrev bool, totalCount int32, ) *BackupPagedApiResponse {
	this := BackupPagedApiResponse{}
	this.Entities = entities
	this.HasNext = hasNext
	this.HasPrev = hasPrev
	this.TotalCount = totalCount
	return &this
}

// NewBackupPagedApiResponseWithDefaults instantiates a new BackupPagedApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPagedApiResponseWithDefaults() *BackupPagedApiResponse {
	this := BackupPagedApiResponse{}
	return &this
}

// GetEntities returns the Entities field value
func (o *BackupPagedApiResponse) GetEntities() []BackupResp {
	if o == nil  {
		var ret []BackupResp
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *BackupPagedApiResponse) GetEntitiesOk() (*[]BackupResp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *BackupPagedApiResponse) SetEntities(v []BackupResp) {
	o.Entities = v
}

// GetHasNext returns the HasNext field value
func (o *BackupPagedApiResponse) GetHasNext() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *BackupPagedApiResponse) GetHasNextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *BackupPagedApiResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrev returns the HasPrev field value
func (o *BackupPagedApiResponse) GetHasPrev() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value
// and a boolean to check if the value has been set.
func (o *BackupPagedApiResponse) GetHasPrevOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasPrev, true
}

// SetHasPrev sets field value
func (o *BackupPagedApiResponse) SetHasPrev(v bool) {
	o.HasPrev = v
}

// GetTotalCount returns the TotalCount field value
func (o *BackupPagedApiResponse) GetTotalCount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *BackupPagedApiResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *BackupPagedApiResponse) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o BackupPagedApiResponse) MarshalJSON() ([]byte, error) {
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

type NullableBackupPagedApiResponse struct {
	value *BackupPagedApiResponse
	isSet bool
}

func (v NullableBackupPagedApiResponse) Get() *BackupPagedApiResponse {
	return v.value
}

func (v *NullableBackupPagedApiResponse) Set(val *BackupPagedApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupPagedApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupPagedApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupPagedApiResponse(val *BackupPagedApiResponse) *NullableBackupPagedApiResponse {
	return &NullableBackupPagedApiResponse{value: val, isSet: true}
}

func (v NullableBackupPagedApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupPagedApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


