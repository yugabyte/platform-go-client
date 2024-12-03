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

// YbSoftwareDetails YbSoftwareDetails  YugabyteDB Software Details object is used to represent the software version of YugabyteDB running on a Universe. 
type YbSoftwareDetails struct {
	// YugabyteDB Software version installed in DB nodes of this Universe
	YbSoftwareVersion *string `json:"yb_software_version,omitempty"`
	// Auto flag version installed in DB nodes of this Universe
	AutoFlagConfigVersion *int32 `json:"auto_flag_config_version,omitempty"`
}

// NewYbSoftwareDetails instantiates a new YbSoftwareDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYbSoftwareDetails() *YbSoftwareDetails {
	this := YbSoftwareDetails{}
	return &this
}

// NewYbSoftwareDetailsWithDefaults instantiates a new YbSoftwareDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYbSoftwareDetailsWithDefaults() *YbSoftwareDetails {
	this := YbSoftwareDetails{}
	return &this
}

// GetYbSoftwareVersion returns the YbSoftwareVersion field value if set, zero value otherwise.
func (o *YbSoftwareDetails) GetYbSoftwareVersion() string {
	if o == nil || o.YbSoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.YbSoftwareVersion
}

// GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YbSoftwareDetails) GetYbSoftwareVersionOk() (*string, bool) {
	if o == nil || o.YbSoftwareVersion == nil {
		return nil, false
	}
	return o.YbSoftwareVersion, true
}

// HasYbSoftwareVersion returns a boolean if a field has been set.
func (o *YbSoftwareDetails) HasYbSoftwareVersion() bool {
	if o != nil && o.YbSoftwareVersion != nil {
		return true
	}

	return false
}

// SetYbSoftwareVersion gets a reference to the given string and assigns it to the YbSoftwareVersion field.
func (o *YbSoftwareDetails) SetYbSoftwareVersion(v string) {
	o.YbSoftwareVersion = &v
}

// GetAutoFlagConfigVersion returns the AutoFlagConfigVersion field value if set, zero value otherwise.
func (o *YbSoftwareDetails) GetAutoFlagConfigVersion() int32 {
	if o == nil || o.AutoFlagConfigVersion == nil {
		var ret int32
		return ret
	}
	return *o.AutoFlagConfigVersion
}

// GetAutoFlagConfigVersionOk returns a tuple with the AutoFlagConfigVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YbSoftwareDetails) GetAutoFlagConfigVersionOk() (*int32, bool) {
	if o == nil || o.AutoFlagConfigVersion == nil {
		return nil, false
	}
	return o.AutoFlagConfigVersion, true
}

// HasAutoFlagConfigVersion returns a boolean if a field has been set.
func (o *YbSoftwareDetails) HasAutoFlagConfigVersion() bool {
	if o != nil && o.AutoFlagConfigVersion != nil {
		return true
	}

	return false
}

// SetAutoFlagConfigVersion gets a reference to the given int32 and assigns it to the AutoFlagConfigVersion field.
func (o *YbSoftwareDetails) SetAutoFlagConfigVersion(v int32) {
	o.AutoFlagConfigVersion = &v
}

func (o YbSoftwareDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.YbSoftwareVersion != nil {
		toSerialize["yb_software_version"] = o.YbSoftwareVersion
	}
	if o.AutoFlagConfigVersion != nil {
		toSerialize["auto_flag_config_version"] = o.AutoFlagConfigVersion
	}
	return json.Marshal(toSerialize)
}

type NullableYbSoftwareDetails struct {
	value *YbSoftwareDetails
	isSet bool
}

func (v NullableYbSoftwareDetails) Get() *YbSoftwareDetails {
	return v.value
}

func (v *NullableYbSoftwareDetails) Set(val *YbSoftwareDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableYbSoftwareDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableYbSoftwareDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYbSoftwareDetails(val *YbSoftwareDetails) *NullableYbSoftwareDetails {
	return &NullableYbSoftwareDetails{value: val, isSet: true}
}

func (v NullableYbSoftwareDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYbSoftwareDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

