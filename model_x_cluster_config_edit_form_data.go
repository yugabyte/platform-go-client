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

// XClusterConfigEditFormData xcluster edit form
type XClusterConfigEditFormData struct {
	BootstrapParams *BootstrapParams `json:"bootstrapParams,omitempty"`
	// Run the pre-checks without actually running the subtasks
	DryRun *bool `json:"dryRun,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
	// Source universe table IDs
	Tables *[]string `json:"tables,omitempty"`
}

// NewXClusterConfigEditFormData instantiates a new XClusterConfigEditFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXClusterConfigEditFormData() *XClusterConfigEditFormData {
	this := XClusterConfigEditFormData{}
	return &this
}

// NewXClusterConfigEditFormDataWithDefaults instantiates a new XClusterConfigEditFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXClusterConfigEditFormDataWithDefaults() *XClusterConfigEditFormData {
	this := XClusterConfigEditFormData{}
	return &this
}

// GetBootstrapParams returns the BootstrapParams field value if set, zero value otherwise.
func (o *XClusterConfigEditFormData) GetBootstrapParams() BootstrapParams {
	if o == nil || o.BootstrapParams == nil {
		var ret BootstrapParams
		return ret
	}
	return *o.BootstrapParams
}

// GetBootstrapParamsOk returns a tuple with the BootstrapParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigEditFormData) GetBootstrapParamsOk() (*BootstrapParams, bool) {
	if o == nil || o.BootstrapParams == nil {
		return nil, false
	}
	return o.BootstrapParams, true
}

// HasBootstrapParams returns a boolean if a field has been set.
func (o *XClusterConfigEditFormData) HasBootstrapParams() bool {
	if o != nil && o.BootstrapParams != nil {
		return true
	}

	return false
}

// SetBootstrapParams gets a reference to the given BootstrapParams and assigns it to the BootstrapParams field.
func (o *XClusterConfigEditFormData) SetBootstrapParams(v BootstrapParams) {
	o.BootstrapParams = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *XClusterConfigEditFormData) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigEditFormData) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *XClusterConfigEditFormData) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *XClusterConfigEditFormData) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XClusterConfigEditFormData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigEditFormData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XClusterConfigEditFormData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XClusterConfigEditFormData) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *XClusterConfigEditFormData) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigEditFormData) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *XClusterConfigEditFormData) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *XClusterConfigEditFormData) SetStatus(v string) {
	o.Status = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *XClusterConfigEditFormData) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return *o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XClusterConfigEditFormData) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *XClusterConfigEditFormData) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *XClusterConfigEditFormData) SetTables(v []string) {
	o.Tables = &v
}

func (o XClusterConfigEditFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootstrapParams != nil {
		toSerialize["bootstrapParams"] = o.BootstrapParams
	}
	if o.DryRun != nil {
		toSerialize["dryRun"] = o.DryRun
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	return json.Marshal(toSerialize)
}

type NullableXClusterConfigEditFormData struct {
	value *XClusterConfigEditFormData
	isSet bool
}

func (v NullableXClusterConfigEditFormData) Get() *XClusterConfigEditFormData {
	return v.value
}

func (v *NullableXClusterConfigEditFormData) Set(val *XClusterConfigEditFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableXClusterConfigEditFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableXClusterConfigEditFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXClusterConfigEditFormData(val *XClusterConfigEditFormData) *NullableXClusterConfigEditFormData {
	return &NullableXClusterConfigEditFormData{value: val, isSet: true}
}

func (v NullableXClusterConfigEditFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXClusterConfigEditFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


