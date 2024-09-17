/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// AlertApiFilter struct for AlertApiFilter
type AlertApiFilter struct {
	ConfigurationTypes []string `json:"configurationTypes"`
	ConfigurationUuid string `json:"configurationUuid"`
	Severities []string `json:"severities"`
	SourceName string `json:"sourceName"`
	SourceUUIDs []string `json:"sourceUUIDs"`
	States []string `json:"states"`
	Uuids []string `json:"uuids"`
}

// NewAlertApiFilter instantiates a new AlertApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertApiFilter(configurationTypes []string, configurationUuid string, severities []string, sourceName string, sourceUUIDs []string, states []string, uuids []string) *AlertApiFilter {
	this := AlertApiFilter{}
	this.ConfigurationTypes = configurationTypes
	this.ConfigurationUuid = configurationUuid
	this.Severities = severities
	this.SourceName = sourceName
	this.SourceUUIDs = sourceUUIDs
	this.States = states
	this.Uuids = uuids
	return &this
}

// NewAlertApiFilterWithDefaults instantiates a new AlertApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertApiFilterWithDefaults() *AlertApiFilter {
	this := AlertApiFilter{}
	return &this
}

// GetConfigurationTypes returns the ConfigurationTypes field value
func (o *AlertApiFilter) GetConfigurationTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfigurationTypes
}

// GetConfigurationTypesOk returns a tuple with the ConfigurationTypes field value
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetConfigurationTypesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigurationTypes, true
}

// SetConfigurationTypes sets field value
func (o *AlertApiFilter) SetConfigurationTypes(v []string) {
	o.ConfigurationTypes = v
}

// GetConfigurationUuid returns the ConfigurationUuid field value
func (o *AlertApiFilter) GetConfigurationUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationUuid
}

// GetConfigurationUuidOk returns a tuple with the ConfigurationUuid field value
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetConfigurationUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigurationUuid, true
}

// SetConfigurationUuid sets field value
func (o *AlertApiFilter) SetConfigurationUuid(v string) {
	o.ConfigurationUuid = v
}

// GetSeverities returns the Severities field value
func (o *AlertApiFilter) GetSeverities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Severities
}

// GetSeveritiesOk returns a tuple with the Severities field value
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetSeveritiesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Severities, true
}

// SetSeverities sets field value
func (o *AlertApiFilter) SetSeverities(v []string) {
	o.Severities = v
}

// GetSourceName returns the SourceName field value
func (o *AlertApiFilter) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetSourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *AlertApiFilter) SetSourceName(v string) {
	o.SourceName = v
}

// GetSourceUUIDs returns the SourceUUIDs field value
func (o *AlertApiFilter) GetSourceUUIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SourceUUIDs
}

// GetSourceUUIDsOk returns a tuple with the SourceUUIDs field value
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetSourceUUIDsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceUUIDs, true
}

// SetSourceUUIDs sets field value
func (o *AlertApiFilter) SetSourceUUIDs(v []string) {
	o.SourceUUIDs = v
}

// GetStates returns the States field value
func (o *AlertApiFilter) GetStates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.States
}

// GetStatesOk returns a tuple with the States field value
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetStatesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.States, true
}

// SetStates sets field value
func (o *AlertApiFilter) SetStates(v []string) {
	o.States = v
}

// GetUuids returns the Uuids field value
func (o *AlertApiFilter) GetUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Uuids
}

// GetUuidsOk returns a tuple with the Uuids field value
// and a boolean to check if the value has been set.
func (o *AlertApiFilter) GetUuidsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuids, true
}

// SetUuids sets field value
func (o *AlertApiFilter) SetUuids(v []string) {
	o.Uuids = v
}

func (o AlertApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configurationTypes"] = o.ConfigurationTypes
	}
	if true {
		toSerialize["configurationUuid"] = o.ConfigurationUuid
	}
	if true {
		toSerialize["severities"] = o.Severities
	}
	if true {
		toSerialize["sourceName"] = o.SourceName
	}
	if true {
		toSerialize["sourceUUIDs"] = o.SourceUUIDs
	}
	if true {
		toSerialize["states"] = o.States
	}
	if true {
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


