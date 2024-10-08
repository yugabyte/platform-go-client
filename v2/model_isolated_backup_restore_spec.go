/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// IsolatedBackupRestoreSpec IsolatedBackupRestoreSpec  User provided details to trigger a restore of YBA from a local filepath. 
type IsolatedBackupRestoreSpec struct {
	// full path to YBA backup file on host
	LocalPath string `json:"local_path"`
}

// NewIsolatedBackupRestoreSpec instantiates a new IsolatedBackupRestoreSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsolatedBackupRestoreSpec(localPath string) *IsolatedBackupRestoreSpec {
	this := IsolatedBackupRestoreSpec{}
	this.LocalPath = localPath
	return &this
}

// NewIsolatedBackupRestoreSpecWithDefaults instantiates a new IsolatedBackupRestoreSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsolatedBackupRestoreSpecWithDefaults() *IsolatedBackupRestoreSpec {
	this := IsolatedBackupRestoreSpec{}
	return &this
}

// GetLocalPath returns the LocalPath field value
func (o *IsolatedBackupRestoreSpec) GetLocalPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalPath
}

// GetLocalPathOk returns a tuple with the LocalPath field value
// and a boolean to check if the value has been set.
func (o *IsolatedBackupRestoreSpec) GetLocalPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LocalPath, true
}

// SetLocalPath sets field value
func (o *IsolatedBackupRestoreSpec) SetLocalPath(v string) {
	o.LocalPath = v
}

func (o IsolatedBackupRestoreSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["local_path"] = o.LocalPath
	}
	return json.Marshal(toSerialize)
}

type NullableIsolatedBackupRestoreSpec struct {
	value *IsolatedBackupRestoreSpec
	isSet bool
}

func (v NullableIsolatedBackupRestoreSpec) Get() *IsolatedBackupRestoreSpec {
	return v.value
}

func (v *NullableIsolatedBackupRestoreSpec) Set(val *IsolatedBackupRestoreSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableIsolatedBackupRestoreSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableIsolatedBackupRestoreSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsolatedBackupRestoreSpec(val *IsolatedBackupRestoreSpec) *NullableIsolatedBackupRestoreSpec {
	return &NullableIsolatedBackupRestoreSpec{value: val, isSet: true}
}

func (v NullableIsolatedBackupRestoreSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsolatedBackupRestoreSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


