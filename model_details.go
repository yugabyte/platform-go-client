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
	"time"
)

// Details struct for Details
type Details struct {
	Data []NodeData `json:"data"`
	HasError bool `json:"has_error"`
	HasWarning bool `json:"has_warning"`
	TimestampIso *time.Time `json:"timestamp_iso,omitempty"`
	YbVersion string `json:"yb_version"`
}

// NewDetails instantiates a new Details object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetails(data []NodeData, hasError bool, hasWarning bool, ybVersion string) *Details {
	this := Details{}
	this.Data = data
	this.HasError = hasError
	this.HasWarning = hasWarning
	this.YbVersion = ybVersion
	return &this
}

// NewDetailsWithDefaults instantiates a new Details object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailsWithDefaults() *Details {
	this := Details{}
	return &this
}

// GetData returns the Data field value
func (o *Details) GetData() []NodeData {
	if o == nil {
		var ret []NodeData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Details) GetDataOk() (*[]NodeData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Details) SetData(v []NodeData) {
	o.Data = v
}

// GetHasError returns the HasError field value
func (o *Details) GetHasError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasError
}

// GetHasErrorOk returns a tuple with the HasError field value
// and a boolean to check if the value has been set.
func (o *Details) GetHasErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasError, true
}

// SetHasError sets field value
func (o *Details) SetHasError(v bool) {
	o.HasError = v
}

// GetHasWarning returns the HasWarning field value
func (o *Details) GetHasWarning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasWarning
}

// GetHasWarningOk returns a tuple with the HasWarning field value
// and a boolean to check if the value has been set.
func (o *Details) GetHasWarningOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasWarning, true
}

// SetHasWarning sets field value
func (o *Details) SetHasWarning(v bool) {
	o.HasWarning = v
}

// GetTimestampIso returns the TimestampIso field value if set, zero value otherwise.
func (o *Details) GetTimestampIso() time.Time {
	if o == nil || o.TimestampIso == nil {
		var ret time.Time
		return ret
	}
	return *o.TimestampIso
}

// GetTimestampIsoOk returns a tuple with the TimestampIso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Details) GetTimestampIsoOk() (*time.Time, bool) {
	if o == nil || o.TimestampIso == nil {
		return nil, false
	}
	return o.TimestampIso, true
}

// HasTimestampIso returns a boolean if a field has been set.
func (o *Details) HasTimestampIso() bool {
	if o != nil && o.TimestampIso != nil {
		return true
	}

	return false
}

// SetTimestampIso gets a reference to the given time.Time and assigns it to the TimestampIso field.
func (o *Details) SetTimestampIso(v time.Time) {
	o.TimestampIso = &v
}

// GetYbVersion returns the YbVersion field value
func (o *Details) GetYbVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbVersion
}

// GetYbVersionOk returns a tuple with the YbVersion field value
// and a boolean to check if the value has been set.
func (o *Details) GetYbVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbVersion, true
}

// SetYbVersion sets field value
func (o *Details) SetYbVersion(v string) {
	o.YbVersion = v
}

func (o Details) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["has_error"] = o.HasError
	}
	if true {
		toSerialize["has_warning"] = o.HasWarning
	}
	if o.TimestampIso != nil {
		toSerialize["timestamp_iso"] = o.TimestampIso
	}
	if true {
		toSerialize["yb_version"] = o.YbVersion
	}
	return json.Marshal(toSerialize)
}

type NullableDetails struct {
	value *Details
	isSet bool
}

func (v NullableDetails) Get() *Details {
	return v.value
}

func (v *NullableDetails) Set(val *Details) {
	v.value = val
	v.isSet = true
}

func (v NullableDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetails(val *Details) *NullableDetails {
	return &NullableDetails{value: val, isSet: true}
}

func (v NullableDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


