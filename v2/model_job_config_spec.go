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

// JobConfigSpec Job config spec for a schedule that is a part of JobScheduleSpec
type JobConfigSpec struct {
	// Class name of the job config.
	Classname string `json:"classname"`
	// Additional custom configuration.
	Config *map[string]map[string]interface{} `json:"config,omitempty"`
}

// NewJobConfigSpec instantiates a new JobConfigSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobConfigSpec(classname string) *JobConfigSpec {
	this := JobConfigSpec{}
	this.Classname = classname
	return &this
}

// NewJobConfigSpecWithDefaults instantiates a new JobConfigSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobConfigSpecWithDefaults() *JobConfigSpec {
	this := JobConfigSpec{}
	return &this
}

// GetClassname returns the Classname field value
func (o *JobConfigSpec) GetClassname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Classname
}

// GetClassnameOk returns a tuple with the Classname field value
// and a boolean to check if the value has been set.
func (o *JobConfigSpec) GetClassnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Classname, true
}

// SetClassname sets field value
func (o *JobConfigSpec) SetClassname(v string) {
	o.Classname = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *JobConfigSpec) GetConfig() map[string]map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobConfigSpec) GetConfigOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *JobConfigSpec) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]map[string]interface{} and assigns it to the Config field.
func (o *JobConfigSpec) SetConfig(v map[string]map[string]interface{}) {
	o.Config = &v
}

func (o JobConfigSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["classname"] = o.Classname
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableJobConfigSpec struct {
	value *JobConfigSpec
	isSet bool
}

func (v NullableJobConfigSpec) Get() *JobConfigSpec {
	return v.value
}

func (v *NullableJobConfigSpec) Set(val *JobConfigSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableJobConfigSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableJobConfigSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobConfigSpec(val *JobConfigSpec) *NullableJobConfigSpec {
	return &NullableJobConfigSpec{value: val, isSet: true}
}

func (v NullableJobConfigSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobConfigSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


