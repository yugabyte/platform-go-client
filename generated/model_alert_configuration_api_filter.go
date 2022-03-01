/*
 * Yugabyte Platform APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yb_platform_client

import (
	"encoding/json"
)

// AlertConfigurationApiFilter struct for AlertConfigurationApiFilter
type AlertConfigurationApiFilter struct {
	Active bool `json:"active"`
	DestinationType string `json:"destinationType"`
	DestinationUuid string `json:"destinationUuid"`
	Name string `json:"name"`
	Severity string `json:"severity"`
	Target AlertConfigurationTarget `json:"target"`
	TargetType string `json:"targetType"`
	Template string `json:"template"`
	Uuids []string `json:"uuids"`
}

// NewAlertConfigurationApiFilter instantiates a new AlertConfigurationApiFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertConfigurationApiFilter(active bool, destinationType string, destinationUuid string, name string, severity string, target AlertConfigurationTarget, targetType string, template string, uuids []string, ) *AlertConfigurationApiFilter {
	this := AlertConfigurationApiFilter{}
	this.Active = active
	this.DestinationType = destinationType
	this.DestinationUuid = destinationUuid
	this.Name = name
	this.Severity = severity
	this.Target = target
	this.TargetType = targetType
	this.Template = template
	this.Uuids = uuids
	return &this
}

// NewAlertConfigurationApiFilterWithDefaults instantiates a new AlertConfigurationApiFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertConfigurationApiFilterWithDefaults() *AlertConfigurationApiFilter {
	this := AlertConfigurationApiFilter{}
	return &this
}

// GetActive returns the Active field value
func (o *AlertConfigurationApiFilter) GetActive() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *AlertConfigurationApiFilter) SetActive(v bool) {
	o.Active = v
}

// GetDestinationType returns the DestinationType field value
func (o *AlertConfigurationApiFilter) GetDestinationType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetDestinationTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *AlertConfigurationApiFilter) SetDestinationType(v string) {
	o.DestinationType = v
}

// GetDestinationUuid returns the DestinationUuid field value
func (o *AlertConfigurationApiFilter) GetDestinationUuid() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DestinationUuid
}

// GetDestinationUuidOk returns a tuple with the DestinationUuid field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetDestinationUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationUuid, true
}

// SetDestinationUuid sets field value
func (o *AlertConfigurationApiFilter) SetDestinationUuid(v string) {
	o.DestinationUuid = v
}

// GetName returns the Name field value
func (o *AlertConfigurationApiFilter) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlertConfigurationApiFilter) SetName(v string) {
	o.Name = v
}

// GetSeverity returns the Severity field value
func (o *AlertConfigurationApiFilter) GetSeverity() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetSeverityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *AlertConfigurationApiFilter) SetSeverity(v string) {
	o.Severity = v
}

// GetTarget returns the Target field value
func (o *AlertConfigurationApiFilter) GetTarget() AlertConfigurationTarget {
	if o == nil  {
		var ret AlertConfigurationTarget
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetTargetOk() (*AlertConfigurationTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *AlertConfigurationApiFilter) SetTarget(v AlertConfigurationTarget) {
	o.Target = v
}

// GetTargetType returns the TargetType field value
func (o *AlertConfigurationApiFilter) GetTargetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetTargetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetType, true
}

// SetTargetType sets field value
func (o *AlertConfigurationApiFilter) SetTargetType(v string) {
	o.TargetType = v
}

// GetTemplate returns the Template field value
func (o *AlertConfigurationApiFilter) GetTemplate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *AlertConfigurationApiFilter) SetTemplate(v string) {
	o.Template = v
}

// GetUuids returns the Uuids field value
func (o *AlertConfigurationApiFilter) GetUuids() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Uuids
}

// GetUuidsOk returns a tuple with the Uuids field value
// and a boolean to check if the value has been set.
func (o *AlertConfigurationApiFilter) GetUuidsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuids, true
}

// SetUuids sets field value
func (o *AlertConfigurationApiFilter) SetUuids(v []string) {
	o.Uuids = v
}

func (o AlertConfigurationApiFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["destinationType"] = o.DestinationType
	}
	if true {
		toSerialize["destinationUuid"] = o.DestinationUuid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["targetType"] = o.TargetType
	}
	if true {
		toSerialize["template"] = o.Template
	}
	if true {
		toSerialize["uuids"] = o.Uuids
	}
	return json.Marshal(toSerialize)
}

type NullableAlertConfigurationApiFilter struct {
	value *AlertConfigurationApiFilter
	isSet bool
}

func (v NullableAlertConfigurationApiFilter) Get() *AlertConfigurationApiFilter {
	return v.value
}

func (v *NullableAlertConfigurationApiFilter) Set(val *AlertConfigurationApiFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertConfigurationApiFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertConfigurationApiFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertConfigurationApiFilter(val *AlertConfigurationApiFilter) *NullableAlertConfigurationApiFilter {
	return &NullableAlertConfigurationApiFilter{value: val, isSet: true}
}

func (v NullableAlertConfigurationApiFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertConfigurationApiFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


