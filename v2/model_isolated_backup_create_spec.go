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

// IsolatedBackupCreateSpec IsolatedBackupCreateSpec  User provided details to trigger a one time backup of YBA to a particular local directory. 
type IsolatedBackupCreateSpec struct {
	// local directory to store backup
	LocalDir string `json:"local_dir"`
	// the components to include in YBA backup
	Components []YbaComponent `json:"components"`
}

// NewIsolatedBackupCreateSpec instantiates a new IsolatedBackupCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsolatedBackupCreateSpec(localDir string, components []YbaComponent) *IsolatedBackupCreateSpec {
	this := IsolatedBackupCreateSpec{}
	this.LocalDir = localDir
	this.Components = components
	return &this
}

// NewIsolatedBackupCreateSpecWithDefaults instantiates a new IsolatedBackupCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsolatedBackupCreateSpecWithDefaults() *IsolatedBackupCreateSpec {
	this := IsolatedBackupCreateSpec{}
	return &this
}

// GetLocalDir returns the LocalDir field value
func (o *IsolatedBackupCreateSpec) GetLocalDir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalDir
}

// GetLocalDirOk returns a tuple with the LocalDir field value
// and a boolean to check if the value has been set.
func (o *IsolatedBackupCreateSpec) GetLocalDirOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LocalDir, true
}

// SetLocalDir sets field value
func (o *IsolatedBackupCreateSpec) SetLocalDir(v string) {
	o.LocalDir = v
}

// GetComponents returns the Components field value
func (o *IsolatedBackupCreateSpec) GetComponents() []YbaComponent {
	if o == nil {
		var ret []YbaComponent
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *IsolatedBackupCreateSpec) GetComponentsOk() (*[]YbaComponent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value
func (o *IsolatedBackupCreateSpec) SetComponents(v []YbaComponent) {
	o.Components = v
}

func (o IsolatedBackupCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["local_dir"] = o.LocalDir
	}
	if true {
		toSerialize["components"] = o.Components
	}
	return json.Marshal(toSerialize)
}

type NullableIsolatedBackupCreateSpec struct {
	value *IsolatedBackupCreateSpec
	isSet bool
}

func (v NullableIsolatedBackupCreateSpec) Get() *IsolatedBackupCreateSpec {
	return v.value
}

func (v *NullableIsolatedBackupCreateSpec) Set(val *IsolatedBackupCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableIsolatedBackupCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableIsolatedBackupCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsolatedBackupCreateSpec(val *IsolatedBackupCreateSpec) *NullableIsolatedBackupCreateSpec {
	return &NullableIsolatedBackupCreateSpec{value: val, isSet: true}
}

func (v NullableIsolatedBackupCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsolatedBackupCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

