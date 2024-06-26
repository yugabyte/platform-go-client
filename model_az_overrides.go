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

// AZOverrides WARNING: This is a preview API that could change. Availability zone level overrides
type AZOverrides struct {
	CgroupSize *int32 `json:"cgroupSize,omitempty"`
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
	PerProcess *map[string]PerProcessDetails `json:"perProcess,omitempty"`
	ProxyConfig *ProxyConfig `json:"proxyConfig,omitempty"`
}

// NewAZOverrides instantiates a new AZOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAZOverrides() *AZOverrides {
	this := AZOverrides{}
	return &this
}

// NewAZOverridesWithDefaults instantiates a new AZOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAZOverridesWithDefaults() *AZOverrides {
	this := AZOverrides{}
	return &this
}

// GetCgroupSize returns the CgroupSize field value if set, zero value otherwise.
func (o *AZOverrides) GetCgroupSize() int32 {
	if o == nil || o.CgroupSize == nil {
		var ret int32
		return ret
	}
	return *o.CgroupSize
}

// GetCgroupSizeOk returns a tuple with the CgroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AZOverrides) GetCgroupSizeOk() (*int32, bool) {
	if o == nil || o.CgroupSize == nil {
		return nil, false
	}
	return o.CgroupSize, true
}

// HasCgroupSize returns a boolean if a field has been set.
func (o *AZOverrides) HasCgroupSize() bool {
	if o != nil && o.CgroupSize != nil {
		return true
	}

	return false
}

// SetCgroupSize gets a reference to the given int32 and assigns it to the CgroupSize field.
func (o *AZOverrides) SetCgroupSize(v int32) {
	o.CgroupSize = &v
}

// GetDeviceInfo returns the DeviceInfo field value if set, zero value otherwise.
func (o *AZOverrides) GetDeviceInfo() DeviceInfo {
	if o == nil || o.DeviceInfo == nil {
		var ret DeviceInfo
		return ret
	}
	return *o.DeviceInfo
}

// GetDeviceInfoOk returns a tuple with the DeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AZOverrides) GetDeviceInfoOk() (*DeviceInfo, bool) {
	if o == nil || o.DeviceInfo == nil {
		return nil, false
	}
	return o.DeviceInfo, true
}

// HasDeviceInfo returns a boolean if a field has been set.
func (o *AZOverrides) HasDeviceInfo() bool {
	if o != nil && o.DeviceInfo != nil {
		return true
	}

	return false
}

// SetDeviceInfo gets a reference to the given DeviceInfo and assigns it to the DeviceInfo field.
func (o *AZOverrides) SetDeviceInfo(v DeviceInfo) {
	o.DeviceInfo = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *AZOverrides) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AZOverrides) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *AZOverrides) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *AZOverrides) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetPerProcess returns the PerProcess field value if set, zero value otherwise.
func (o *AZOverrides) GetPerProcess() map[string]PerProcessDetails {
	if o == nil || o.PerProcess == nil {
		var ret map[string]PerProcessDetails
		return ret
	}
	return *o.PerProcess
}

// GetPerProcessOk returns a tuple with the PerProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AZOverrides) GetPerProcessOk() (*map[string]PerProcessDetails, bool) {
	if o == nil || o.PerProcess == nil {
		return nil, false
	}
	return o.PerProcess, true
}

// HasPerProcess returns a boolean if a field has been set.
func (o *AZOverrides) HasPerProcess() bool {
	if o != nil && o.PerProcess != nil {
		return true
	}

	return false
}

// SetPerProcess gets a reference to the given map[string]PerProcessDetails and assigns it to the PerProcess field.
func (o *AZOverrides) SetPerProcess(v map[string]PerProcessDetails) {
	o.PerProcess = &v
}

// GetProxyConfig returns the ProxyConfig field value if set, zero value otherwise.
func (o *AZOverrides) GetProxyConfig() ProxyConfig {
	if o == nil || o.ProxyConfig == nil {
		var ret ProxyConfig
		return ret
	}
	return *o.ProxyConfig
}

// GetProxyConfigOk returns a tuple with the ProxyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AZOverrides) GetProxyConfigOk() (*ProxyConfig, bool) {
	if o == nil || o.ProxyConfig == nil {
		return nil, false
	}
	return o.ProxyConfig, true
}

// HasProxyConfig returns a boolean if a field has been set.
func (o *AZOverrides) HasProxyConfig() bool {
	if o != nil && o.ProxyConfig != nil {
		return true
	}

	return false
}

// SetProxyConfig gets a reference to the given ProxyConfig and assigns it to the ProxyConfig field.
func (o *AZOverrides) SetProxyConfig(v ProxyConfig) {
	o.ProxyConfig = &v
}

func (o AZOverrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CgroupSize != nil {
		toSerialize["cgroupSize"] = o.CgroupSize
	}
	if o.DeviceInfo != nil {
		toSerialize["deviceInfo"] = o.DeviceInfo
	}
	if o.InstanceType != nil {
		toSerialize["instanceType"] = o.InstanceType
	}
	if o.PerProcess != nil {
		toSerialize["perProcess"] = o.PerProcess
	}
	if o.ProxyConfig != nil {
		toSerialize["proxyConfig"] = o.ProxyConfig
	}
	return json.Marshal(toSerialize)
}

type NullableAZOverrides struct {
	value *AZOverrides
	isSet bool
}

func (v NullableAZOverrides) Get() *AZOverrides {
	return v.value
}

func (v *NullableAZOverrides) Set(val *AZOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullableAZOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullableAZOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAZOverrides(val *AZOverrides) *NullableAZOverrides {
	return &NullableAZOverrides{value: val, isSet: true}
}

func (v NullableAZOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAZOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


