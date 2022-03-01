/*
 * Yugabyte Platform APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ywclient

import (
	"encoding/json"
	"time"
)

// SupportBundle struct for SupportBundle
type SupportBundle struct {
	BundleDetails BundleDetails `json:"bundleDetails"`
	BundleUUID string `json:"bundleUUID"`
	EndDate time.Time `json:"endDate"`
	Path string `json:"path"`
	ScopeUUID string `json:"scopeUUID"`
	StartDate time.Time `json:"startDate"`
	Status string `json:"status"`
}

// NewSupportBundle instantiates a new SupportBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportBundle(bundleDetails BundleDetails, bundleUUID string, endDate time.Time, path string, scopeUUID string, startDate time.Time, status string, ) *SupportBundle {
	this := SupportBundle{}
	this.BundleDetails = bundleDetails
	this.BundleUUID = bundleUUID
	this.EndDate = endDate
	this.Path = path
	this.ScopeUUID = scopeUUID
	this.StartDate = startDate
	this.Status = status
	return &this
}

// NewSupportBundleWithDefaults instantiates a new SupportBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportBundleWithDefaults() *SupportBundle {
	this := SupportBundle{}
	return &this
}

// GetBundleDetails returns the BundleDetails field value
func (o *SupportBundle) GetBundleDetails() BundleDetails {
	if o == nil  {
		var ret BundleDetails
		return ret
	}

	return o.BundleDetails
}

// GetBundleDetailsOk returns a tuple with the BundleDetails field value
// and a boolean to check if the value has been set.
func (o *SupportBundle) GetBundleDetailsOk() (*BundleDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BundleDetails, true
}

// SetBundleDetails sets field value
func (o *SupportBundle) SetBundleDetails(v BundleDetails) {
	o.BundleDetails = v
}

// GetBundleUUID returns the BundleUUID field value
func (o *SupportBundle) GetBundleUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BundleUUID
}

// GetBundleUUIDOk returns a tuple with the BundleUUID field value
// and a boolean to check if the value has been set.
func (o *SupportBundle) GetBundleUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BundleUUID, true
}

// SetBundleUUID sets field value
func (o *SupportBundle) SetBundleUUID(v string) {
	o.BundleUUID = v
}

// GetEndDate returns the EndDate field value
func (o *SupportBundle) GetEndDate() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *SupportBundle) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *SupportBundle) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetPath returns the Path field value
func (o *SupportBundle) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *SupportBundle) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *SupportBundle) SetPath(v string) {
	o.Path = v
}

// GetScopeUUID returns the ScopeUUID field value
func (o *SupportBundle) GetScopeUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ScopeUUID
}

// GetScopeUUIDOk returns a tuple with the ScopeUUID field value
// and a boolean to check if the value has been set.
func (o *SupportBundle) GetScopeUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScopeUUID, true
}

// SetScopeUUID sets field value
func (o *SupportBundle) SetScopeUUID(v string) {
	o.ScopeUUID = v
}

// GetStartDate returns the StartDate field value
func (o *SupportBundle) GetStartDate() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *SupportBundle) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *SupportBundle) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetStatus returns the Status field value
func (o *SupportBundle) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SupportBundle) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SupportBundle) SetStatus(v string) {
	o.Status = v
}

func (o SupportBundle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bundleDetails"] = o.BundleDetails
	}
	if true {
		toSerialize["bundleUUID"] = o.BundleUUID
	}
	if true {
		toSerialize["endDate"] = o.EndDate
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["scopeUUID"] = o.ScopeUUID
	}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSupportBundle struct {
	value *SupportBundle
	isSet bool
}

func (v NullableSupportBundle) Get() *SupportBundle {
	return v.value
}

func (v *NullableSupportBundle) Set(val *SupportBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportBundle(val *SupportBundle) *NullableSupportBundle {
	return &NullableSupportBundle{value: val, isSet: true}
}

func (v NullableSupportBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


