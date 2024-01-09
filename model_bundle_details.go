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

// BundleDetails struct for BundleDetails
type BundleDetails struct {
	Components []string `json:"components"`
	// Max size of the collected cores (if any)
	MaxCoreFileSize *int64 `json:"maxCoreFileSize,omitempty"`
	// Max number of most recent cores to collect (if any)
	MaxNumRecentCores *int32 `json:"maxNumRecentCores,omitempty"`
}

// NewBundleDetails instantiates a new BundleDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleDetails(components []string, ) *BundleDetails {
	this := BundleDetails{}
	this.Components = components
	return &this
}

// NewBundleDetailsWithDefaults instantiates a new BundleDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDetailsWithDefaults() *BundleDetails {
	this := BundleDetails{}
	return &this
}

// GetComponents returns the Components field value
func (o *BundleDetails) GetComponents() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *BundleDetails) GetComponentsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value
func (o *BundleDetails) SetComponents(v []string) {
	o.Components = v
}

// GetMaxCoreFileSize returns the MaxCoreFileSize field value if set, zero value otherwise.
func (o *BundleDetails) GetMaxCoreFileSize() int64 {
	if o == nil || o.MaxCoreFileSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxCoreFileSize
}

// GetMaxCoreFileSizeOk returns a tuple with the MaxCoreFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDetails) GetMaxCoreFileSizeOk() (*int64, bool) {
	if o == nil || o.MaxCoreFileSize == nil {
		return nil, false
	}
	return o.MaxCoreFileSize, true
}

// HasMaxCoreFileSize returns a boolean if a field has been set.
func (o *BundleDetails) HasMaxCoreFileSize() bool {
	if o != nil && o.MaxCoreFileSize != nil {
		return true
	}

	return false
}

// SetMaxCoreFileSize gets a reference to the given int64 and assigns it to the MaxCoreFileSize field.
func (o *BundleDetails) SetMaxCoreFileSize(v int64) {
	o.MaxCoreFileSize = &v
}

// GetMaxNumRecentCores returns the MaxNumRecentCores field value if set, zero value otherwise.
func (o *BundleDetails) GetMaxNumRecentCores() int32 {
	if o == nil || o.MaxNumRecentCores == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumRecentCores
}

// GetMaxNumRecentCoresOk returns a tuple with the MaxNumRecentCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDetails) GetMaxNumRecentCoresOk() (*int32, bool) {
	if o == nil || o.MaxNumRecentCores == nil {
		return nil, false
	}
	return o.MaxNumRecentCores, true
}

// HasMaxNumRecentCores returns a boolean if a field has been set.
func (o *BundleDetails) HasMaxNumRecentCores() bool {
	if o != nil && o.MaxNumRecentCores != nil {
		return true
	}

	return false
}

// SetMaxNumRecentCores gets a reference to the given int32 and assigns it to the MaxNumRecentCores field.
func (o *BundleDetails) SetMaxNumRecentCores(v int32) {
	o.MaxNumRecentCores = &v
}

func (o BundleDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["components"] = o.Components
	}
	if o.MaxCoreFileSize != nil {
		toSerialize["maxCoreFileSize"] = o.MaxCoreFileSize
	}
	if o.MaxNumRecentCores != nil {
		toSerialize["maxNumRecentCores"] = o.MaxNumRecentCores
	}
	return json.Marshal(toSerialize)
}

type NullableBundleDetails struct {
	value *BundleDetails
	isSet bool
}

func (v NullableBundleDetails) Get() *BundleDetails {
	return v.value
}

func (v *NullableBundleDetails) Set(val *BundleDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleDetails(val *BundleDetails) *NullableBundleDetails {
	return &NullableBundleDetails{value: val, isSet: true}
}

func (v NullableBundleDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


