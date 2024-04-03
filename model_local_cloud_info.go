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

// LocalCloudInfo struct for LocalCloudInfo
type LocalCloudInfo struct {
	DataHomeDir string `json:"dataHomeDir"`
	EnvVars map[string]string `json:"envVars"`
	YbcBinDir string `json:"ybcBinDir"`
	YugabyteBinDir string `json:"yugabyteBinDir"`
}

// NewLocalCloudInfo instantiates a new LocalCloudInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalCloudInfo(dataHomeDir string, envVars map[string]string, ybcBinDir string, yugabyteBinDir string) *LocalCloudInfo {
	this := LocalCloudInfo{}
	this.DataHomeDir = dataHomeDir
	this.EnvVars = envVars
	this.YbcBinDir = ybcBinDir
	this.YugabyteBinDir = yugabyteBinDir
	return &this
}

// NewLocalCloudInfoWithDefaults instantiates a new LocalCloudInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalCloudInfoWithDefaults() *LocalCloudInfo {
	this := LocalCloudInfo{}
	return &this
}

// GetDataHomeDir returns the DataHomeDir field value
func (o *LocalCloudInfo) GetDataHomeDir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataHomeDir
}

// GetDataHomeDirOk returns a tuple with the DataHomeDir field value
// and a boolean to check if the value has been set.
func (o *LocalCloudInfo) GetDataHomeDirOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataHomeDir, true
}

// SetDataHomeDir sets field value
func (o *LocalCloudInfo) SetDataHomeDir(v string) {
	o.DataHomeDir = v
}

// GetEnvVars returns the EnvVars field value
func (o *LocalCloudInfo) GetEnvVars() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value
// and a boolean to check if the value has been set.
func (o *LocalCloudInfo) GetEnvVarsOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnvVars, true
}

// SetEnvVars sets field value
func (o *LocalCloudInfo) SetEnvVars(v map[string]string) {
	o.EnvVars = v
}

// GetYbcBinDir returns the YbcBinDir field value
func (o *LocalCloudInfo) GetYbcBinDir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbcBinDir
}

// GetYbcBinDirOk returns a tuple with the YbcBinDir field value
// and a boolean to check if the value has been set.
func (o *LocalCloudInfo) GetYbcBinDirOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbcBinDir, true
}

// SetYbcBinDir sets field value
func (o *LocalCloudInfo) SetYbcBinDir(v string) {
	o.YbcBinDir = v
}

// GetYugabyteBinDir returns the YugabyteBinDir field value
func (o *LocalCloudInfo) GetYugabyteBinDir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YugabyteBinDir
}

// GetYugabyteBinDirOk returns a tuple with the YugabyteBinDir field value
// and a boolean to check if the value has been set.
func (o *LocalCloudInfo) GetYugabyteBinDirOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YugabyteBinDir, true
}

// SetYugabyteBinDir sets field value
func (o *LocalCloudInfo) SetYugabyteBinDir(v string) {
	o.YugabyteBinDir = v
}

func (o LocalCloudInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataHomeDir"] = o.DataHomeDir
	}
	if true {
		toSerialize["envVars"] = o.EnvVars
	}
	if true {
		toSerialize["ybcBinDir"] = o.YbcBinDir
	}
	if true {
		toSerialize["yugabyteBinDir"] = o.YugabyteBinDir
	}
	return json.Marshal(toSerialize)
}

type NullableLocalCloudInfo struct {
	value *LocalCloudInfo
	isSet bool
}

func (v NullableLocalCloudInfo) Get() *LocalCloudInfo {
	return v.value
}

func (v *NullableLocalCloudInfo) Set(val *LocalCloudInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalCloudInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalCloudInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalCloudInfo(val *LocalCloudInfo) *NullableLocalCloudInfo {
	return &NullableLocalCloudInfo{value: val, isSet: true}
}

func (v NullableLocalCloudInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalCloudInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


