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

// AlertApiFilter struct for AlertApiFilter
type AlertApiFilter struct {
	// Alert Configuration Target Types
	ConfigurationTypes *[]string `json:"configurationTypes,omitempty"`
	// The uuid of the alert configuration.
	ConfigurationUuid *string `json:"configurationUuid,omitempty"`
	// The severity of the alerts.
	Severities *[]string `json:"severities,omitempty"`
	// The source name of the alerts.
	SourceName *string `json:"sourceName,omitempty"`
	// The source uuids of the alerts.
	SourceUUIDs *[]string `json:"sourceUUIDs,omitempty"`
	// The state of the alerts.
	States *[]string `json:"states,omitempty"`
	// The uuids of the alerts.
	Uuids *[]string `json:"uuids,omitempty"`
}

// NewAlertApiFilter instantiates a new AlertApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertApiFilter() *AlertApiFilter {
	this := AlertApiFilter{}
	return &this
}

// NewAlertApiFilterWithDefaults instantiates a new AlertApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertApiFilterWithDefaults() *AlertApiFilter {
	this := AlertApiFilter{}
	return &this
}

// GetConfigurationTypes returns the ConfigurationTypes field value if set, zero value otherwise.
func (o *AlertApiFilter) GetConfigurationTypes() []string {
	if o == nil || o.ConfigurationTypes == nil {
		var ret []string
		return ret
	}
	return *o.ConfigurationTypes
}

// GetConfigurationTypesOk returns a tuple with the ConfigurationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetConfigurationTypesOk() (*[]string, bool) {
	if o == nil || o.ConfigurationTypes == nil {
		return nil, false
	}
	return o.ConfigurationTypes, true
}

// HasConfigurationTypes returns a boolean if a field has been set.
func (o *AlertApiFilter) HasConfigurationTypes() bool {
	if o != nil && o.ConfigurationTypes != nil {
		return true
	}

	return false
}

// SetConfigurationTypes gets a reference to the given []string and assigns it to the ConfigurationTypes field.
func (o *AlertApiFilter) SetConfigurationTypes(v []string) {
	o.ConfigurationTypes = &v
}

// GetConfigurationUuid returns the ConfigurationUuid field value if set, zero value otherwise.
func (o *AlertApiFilter) GetConfigurationUuid() string {
	if o == nil || o.ConfigurationUuid == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationUuid
}

// GetConfigurationUuidOk returns a tuple with the ConfigurationUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetConfigurationUuidOk() (*string, bool) {
	if o == nil || o.ConfigurationUuid == nil {
		return nil, false
	}
	return o.ConfigurationUuid, true
}

// HasConfigurationUuid returns a boolean if a field has been set.
func (o *AlertApiFilter) HasConfigurationUuid() bool {
	if o != nil && o.ConfigurationUuid != nil {
		return true
	}

	return false
}

// SetConfigurationUuid gets a reference to the given string and assigns it to the ConfigurationUuid field.
func (o *AlertApiFilter) SetConfigurationUuid(v string) {
	o.ConfigurationUuid = &v
}

// GetSeverities returns the Severities field value if set, zero value otherwise.
func (o *AlertApiFilter) GetSeverities() []string {
	if o == nil || o.Severities == nil {
		var ret []string
		return ret
	}
	return *o.Severities
}

// GetSeveritiesOk returns a tuple with the Severities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetSeveritiesOk() (*[]string, bool) {
	if o == nil || o.Severities == nil {
		return nil, false
	}
	return o.Severities, true
}

// HasSeverities returns a boolean if a field has been set.
func (o *AlertApiFilter) HasSeverities() bool {
	if o != nil && o.Severities != nil {
		return true
	}

	return false
}

// SetSeverities gets a reference to the given []string and assigns it to the Severities field.
func (o *AlertApiFilter) SetSeverities(v []string) {
	o.Severities = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AlertApiFilter) GetSourceName() string {
	if o == nil || o.SourceName == nil {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetSourceNameOk() (*string, bool) {
	if o == nil || o.SourceName == nil {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AlertApiFilter) HasSourceName() bool {
	if o != nil && o.SourceName != nil {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AlertApiFilter) SetSourceName(v string) {
	o.SourceName = &v
}

// GetSourceUUIDs returns the SourceUUIDs field value if set, zero value otherwise.
func (o *AlertApiFilter) GetSourceUUIDs() []string {
	if o == nil || o.SourceUUIDs == nil {
		var ret []string
		return ret
	}
	return *o.SourceUUIDs
}

// GetSourceUUIDsOk returns a tuple with the SourceUUIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetSourceUUIDsOk() (*[]string, bool) {
	if o == nil || o.SourceUUIDs == nil {
		return nil, false
	}
	return o.SourceUUIDs, true
}

// HasSourceUUIDs returns a boolean if a field has been set.
func (o *AlertApiFilter) HasSourceUUIDs() bool {
	if o != nil && o.SourceUUIDs != nil {
		return true
	}

	return false
}

// SetSourceUUIDs gets a reference to the given []string and assigns it to the SourceUUIDs field.
func (o *AlertApiFilter) SetSourceUUIDs(v []string) {
	o.SourceUUIDs = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *AlertApiFilter) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetStatesOk() (*[]string, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *AlertApiFilter) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *AlertApiFilter) SetStates(v []string) {
	o.States = &v
}

// GetUuids returns the Uuids field value if set, zero value otherwise.
func (o *AlertApiFilter) GetUuids() []string {
	if o == nil || o.Uuids == nil {
		var ret []string
		return ret
	}
	return *o.Uuids
}

// GetUuidsOk returns a tuple with the Uuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetUuidsOk() (*[]string, bool) {
	if o == nil || o.Uuids == nil {
		return nil, false
	}
	return o.Uuids, true
}

// HasUuids returns a boolean if a field has been set.
func (o *AlertApiFilter) HasUuids() bool {
	if o != nil && o.Uuids != nil {
		return true
	}

	return false
}

// SetUuids gets a reference to the given []string and assigns it to the Uuids field.
func (o *AlertApiFilter) SetUuids(v []string) {
	o.Uuids = &v
}

func (o AlertApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigurationTypes != nil {
		toSerialize["configurationTypes"] = o.ConfigurationTypes
	}
	if o.ConfigurationUuid != nil {
		toSerialize["configurationUuid"] = o.ConfigurationUuid
	}
	if o.Severities != nil {
		toSerialize["severities"] = o.Severities
	}
	if o.SourceName != nil {
		toSerialize["sourceName"] = o.SourceName
	}
	if o.SourceUUIDs != nil {
		toSerialize["sourceUUIDs"] = o.SourceUUIDs
	}
	if o.States != nil {
		toSerialize["states"] = o.States
	}
	if o.Uuids != nil {
		toSerialize["uuids"] = o.Uuids
	}
	return json.Marshal(toSerialize)
}

type NullableAlertApiFilter struct {
	value *AlertApiFilter
	isSet bool
}

func (v NullableAlertApiFilter) Get() *AlertApiFilter {
	return v.value
}

func (v *NullableAlertApiFilter) Set(val *AlertApiFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertApiFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertApiFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertApiFilter(val *AlertApiFilter) *NullableAlertApiFilter {
	return &NullableAlertApiFilter{value: val, isSet: true}
}

func (v NullableAlertApiFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertApiFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


