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

// AlertConfiguration Alert configuration
type AlertConfiguration struct {
	// Is configured alerts raised or not
	Active bool `json:"active"`
	AlertCount float64 `json:"alertCount"`
	// Creation time
	CreateTime time.Time `json:"createTime"`
	// Customer UUID
	CustomerUUID string `json:"customerUUID"`
	// Is default destination used for this config
	DefaultDestination bool `json:"defaultDestination"`
	// Description
	Description string `json:"description"`
	// Alert destination UUID
	DestinationUUID *string `json:"destinationUUID,omitempty"`
	// Duration in seconds, while condition is met to raise an alert
	DurationSec int32 `json:"durationSec"`
	// Maintenance window UUIDs, applied to this alert config
	MaintenanceWindowUuids *[]string `json:"maintenanceWindowUuids,omitempty"`
	// Name
	Name string `json:"name"`
	Target AlertConfigurationTarget `json:"target"`
	// Target type
	TargetType string `json:"targetType"`
	// Template name
	Template string `json:"template"`
	// Threshold unit
	ThresholdUnit string `json:"thresholdUnit"`
	// Thresholds
	Thresholds map[string]AlertConfigurationThreshold `json:"thresholds"`
	// Configuration UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewAlertConfiguration instantiates a new AlertConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertConfiguration(active bool, alertCount float64, createTime time.Time, customerUUID string, defaultDestination bool, description string, durationSec int32, name string, target AlertConfigurationTarget, targetType string, template string, thresholdUnit string, thresholds map[string]AlertConfigurationThreshold, ) *AlertConfiguration {
	this := AlertConfiguration{}
	this.Active = active
	this.AlertCount = alertCount
	this.CreateTime = createTime
	this.CustomerUUID = customerUUID
	this.DefaultDestination = defaultDestination
	this.Description = description
	this.DurationSec = durationSec
	this.Name = name
	this.Target = target
	this.TargetType = targetType
	this.Template = template
	this.ThresholdUnit = thresholdUnit
	this.Thresholds = thresholds
	return &this
}

// NewAlertConfigurationWithDefaults instantiates a new AlertConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertConfigurationWithDefaults() *AlertConfiguration {
	this := AlertConfiguration{}
	return &this
}

// GetActive returns the Active field value
func (o *AlertConfiguration) GetActive() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *AlertConfiguration) SetActive(v bool) {
	o.Active = v
}

// GetAlertCount returns the AlertCount field value
func (o *AlertConfiguration) GetAlertCount() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.AlertCount
}

// GetAlertCountOk returns a tuple with the AlertCount field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetAlertCountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AlertCount, true
}

// SetAlertCount sets field value
func (o *AlertConfiguration) SetAlertCount(v float64) {
	o.AlertCount = v
}

// GetCreateTime returns the CreateTime field value
func (o *AlertConfiguration) GetCreateTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *AlertConfiguration) SetCreateTime(v time.Time) {
	o.CreateTime = v
}

// GetCustomerUUID returns the CustomerUUID field value
func (o *AlertConfiguration) GetCustomerUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetCustomerUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerUUID, true
}

// SetCustomerUUID sets field value
func (o *AlertConfiguration) SetCustomerUUID(v string) {
	o.CustomerUUID = v
}

// GetDefaultDestination returns the DefaultDestination field value
func (o *AlertConfiguration) GetDefaultDestination() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.DefaultDestination
}

// GetDefaultDestinationOk returns a tuple with the DefaultDestination field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetDefaultDestinationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultDestination, true
}

// SetDefaultDestination sets field value
func (o *AlertConfiguration) SetDefaultDestination(v bool) {
	o.DefaultDestination = v
}

// GetDescription returns the Description field value
func (o *AlertConfiguration) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AlertConfiguration) SetDescription(v string) {
	o.Description = v
}

// GetDestinationUUID returns the DestinationUUID field value if set, zero value otherwise.
func (o *AlertConfiguration) GetDestinationUUID() string {
	if o == nil || o.DestinationUUID == nil {
		var ret string
		return ret
	}
	return *o.DestinationUUID
}

// GetDestinationUUIDOk returns a tuple with the DestinationUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetDestinationUUIDOk() (*string, bool) {
	if o == nil || o.DestinationUUID == nil {
		return nil, false
	}
	return o.DestinationUUID, true
}

// HasDestinationUUID returns a boolean if a field has been set.
func (o *AlertConfiguration) HasDestinationUUID() bool {
	if o != nil && o.DestinationUUID != nil {
		return true
	}

	return false
}

// SetDestinationUUID gets a reference to the given string and assigns it to the DestinationUUID field.
func (o *AlertConfiguration) SetDestinationUUID(v string) {
	o.DestinationUUID = &v
}

// GetDurationSec returns the DurationSec field value
func (o *AlertConfiguration) GetDurationSec() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.DurationSec
}

// GetDurationSecOk returns a tuple with the DurationSec field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetDurationSecOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DurationSec, true
}

// SetDurationSec sets field value
func (o *AlertConfiguration) SetDurationSec(v int32) {
	o.DurationSec = v
}

// GetMaintenanceWindowUuids returns the MaintenanceWindowUuids field value if set, zero value otherwise.
func (o *AlertConfiguration) GetMaintenanceWindowUuids() []string {
	if o == nil || o.MaintenanceWindowUuids == nil {
		var ret []string
		return ret
	}
	return *o.MaintenanceWindowUuids
}

// GetMaintenanceWindowUuidsOk returns a tuple with the MaintenanceWindowUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetMaintenanceWindowUuidsOk() (*[]string, bool) {
	if o == nil || o.MaintenanceWindowUuids == nil {
		return nil, false
	}
	return o.MaintenanceWindowUuids, true
}

// HasMaintenanceWindowUuids returns a boolean if a field has been set.
func (o *AlertConfiguration) HasMaintenanceWindowUuids() bool {
	if o != nil && o.MaintenanceWindowUuids != nil {
		return true
	}

	return false
}

// SetMaintenanceWindowUuids gets a reference to the given []string and assigns it to the MaintenanceWindowUuids field.
func (o *AlertConfiguration) SetMaintenanceWindowUuids(v []string) {
	o.MaintenanceWindowUuids = &v
}

// GetName returns the Name field value
func (o *AlertConfiguration) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlertConfiguration) SetName(v string) {
	o.Name = v
}

// GetTarget returns the Target field value
func (o *AlertConfiguration) GetTarget() AlertConfigurationTarget {
	if o == nil  {
		var ret AlertConfigurationTarget
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetTargetOk() (*AlertConfigurationTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *AlertConfiguration) SetTarget(v AlertConfigurationTarget) {
	o.Target = v
}

// GetTargetType returns the TargetType field value
func (o *AlertConfiguration) GetTargetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetTargetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetType, true
}

// SetTargetType sets field value
func (o *AlertConfiguration) SetTargetType(v string) {
	o.TargetType = v
}

// GetTemplate returns the Template field value
func (o *AlertConfiguration) GetTemplate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *AlertConfiguration) SetTemplate(v string) {
	o.Template = v
}

// GetThresholdUnit returns the ThresholdUnit field value
func (o *AlertConfiguration) GetThresholdUnit() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ThresholdUnit
}

// GetThresholdUnitOk returns a tuple with the ThresholdUnit field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetThresholdUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ThresholdUnit, true
}

// SetThresholdUnit sets field value
func (o *AlertConfiguration) SetThresholdUnit(v string) {
	o.ThresholdUnit = v
}

// GetThresholds returns the Thresholds field value
func (o *AlertConfiguration) GetThresholds() map[string]AlertConfigurationThreshold {
	if o == nil  {
		var ret map[string]AlertConfigurationThreshold
		return ret
	}

	return o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetThresholdsOk() (*map[string]AlertConfigurationThreshold, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Thresholds, true
}

// SetThresholds sets field value
func (o *AlertConfiguration) SetThresholds(v map[string]AlertConfigurationThreshold) {
	o.Thresholds = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AlertConfiguration) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfiguration) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AlertConfiguration) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AlertConfiguration) SetUuid(v string) {
	o.Uuid = &v
}

func (o AlertConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["alertCount"] = o.AlertCount
	}
	if true {
		toSerialize["createTime"] = o.CreateTime
	}
	if true {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if true {
		toSerialize["defaultDestination"] = o.DefaultDestination
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.DestinationUUID != nil {
		toSerialize["destinationUUID"] = o.DestinationUUID
	}
	if true {
		toSerialize["durationSec"] = o.DurationSec
	}
	if o.MaintenanceWindowUuids != nil {
		toSerialize["maintenanceWindowUuids"] = o.MaintenanceWindowUuids
	}
	if true {
		toSerialize["name"] = o.Name
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
		toSerialize["thresholdUnit"] = o.ThresholdUnit
	}
	if true {
		toSerialize["thresholds"] = o.Thresholds
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlertConfiguration struct {
	value *AlertConfiguration
	isSet bool
}

func (v NullableAlertConfiguration) Get() *AlertConfiguration {
	return v.value
}

func (v *NullableAlertConfiguration) Set(val *AlertConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertConfiguration(val *AlertConfiguration) *NullableAlertConfiguration {
	return &NullableAlertConfiguration{value: val, isSet: true}
}

func (v NullableAlertConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


