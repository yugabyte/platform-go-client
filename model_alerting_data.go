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

// AlertingData Alerting configuration
type AlertingData struct {
	// Period, which is used to send active alert notifications
	ActiveAlertNotificationIntervalMs *int64 `json:"activeAlertNotificationIntervalMs,omitempty"`
	// Alert email address
	AlertingEmail *string `json:"alertingEmail,omitempty"`
	// Alert interval, in milliseconds
	CheckIntervalMs *int64 `json:"checkIntervalMs,omitempty"`
	// Trigger an alert only for errors
	ReportOnlyErrors *bool `json:"reportOnlyErrors,omitempty"`
	// Send alerts to YB as well as to customer
	SendAlertsToYb *bool `json:"sendAlertsToYb,omitempty"`
	// Status update of alert interval, in milliseconds
	StatusUpdateIntervalMs *int64 `json:"statusUpdateIntervalMs,omitempty"`
}

// NewAlertingData instantiates a new AlertingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertingData() *AlertingData {
	this := AlertingData{}
	return &this
}

// NewAlertingDataWithDefaults instantiates a new AlertingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingDataWithDefaults() *AlertingData {
	this := AlertingData{}
	return &this
}

// GetActiveAlertNotificationIntervalMs returns the ActiveAlertNotificationIntervalMs field value if set, zero value otherwise.
func (o *AlertingData) GetActiveAlertNotificationIntervalMs() int64 {
	if o == nil || o.ActiveAlertNotificationIntervalMs == nil {
		var ret int64
		return ret
	}
	return *o.ActiveAlertNotificationIntervalMs
}

// GetActiveAlertNotificationIntervalMsOk returns a tuple with the ActiveAlertNotificationIntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingData) GetActiveAlertNotificationIntervalMsOk() (*int64, bool) {
	if o == nil || o.ActiveAlertNotificationIntervalMs == nil {
		return nil, false
	}
	return o.ActiveAlertNotificationIntervalMs, true
}

// HasActiveAlertNotificationIntervalMs returns a boolean if a field has been set.
func (o *AlertingData) HasActiveAlertNotificationIntervalMs() bool {
	if o != nil && o.ActiveAlertNotificationIntervalMs != nil {
		return true
	}

	return false
}

// SetActiveAlertNotificationIntervalMs gets a reference to the given int64 and assigns it to the ActiveAlertNotificationIntervalMs field.
func (o *AlertingData) SetActiveAlertNotificationIntervalMs(v int64) {
	o.ActiveAlertNotificationIntervalMs = &v
}

// GetAlertingEmail returns the AlertingEmail field value if set, zero value otherwise.
func (o *AlertingData) GetAlertingEmail() string {
	if o == nil || o.AlertingEmail == nil {
		var ret string
		return ret
	}
	return *o.AlertingEmail
}

// GetAlertingEmailOk returns a tuple with the AlertingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingData) GetAlertingEmailOk() (*string, bool) {
	if o == nil || o.AlertingEmail == nil {
		return nil, false
	}
	return o.AlertingEmail, true
}

// HasAlertingEmail returns a boolean if a field has been set.
func (o *AlertingData) HasAlertingEmail() bool {
	if o != nil && o.AlertingEmail != nil {
		return true
	}

	return false
}

// SetAlertingEmail gets a reference to the given string and assigns it to the AlertingEmail field.
func (o *AlertingData) SetAlertingEmail(v string) {
	o.AlertingEmail = &v
}

// GetCheckIntervalMs returns the CheckIntervalMs field value if set, zero value otherwise.
func (o *AlertingData) GetCheckIntervalMs() int64 {
	if o == nil || o.CheckIntervalMs == nil {
		var ret int64
		return ret
	}
	return *o.CheckIntervalMs
}

// GetCheckIntervalMsOk returns a tuple with the CheckIntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingData) GetCheckIntervalMsOk() (*int64, bool) {
	if o == nil || o.CheckIntervalMs == nil {
		return nil, false
	}
	return o.CheckIntervalMs, true
}

// HasCheckIntervalMs returns a boolean if a field has been set.
func (o *AlertingData) HasCheckIntervalMs() bool {
	if o != nil && o.CheckIntervalMs != nil {
		return true
	}

	return false
}

// SetCheckIntervalMs gets a reference to the given int64 and assigns it to the CheckIntervalMs field.
func (o *AlertingData) SetCheckIntervalMs(v int64) {
	o.CheckIntervalMs = &v
}

// GetReportOnlyErrors returns the ReportOnlyErrors field value if set, zero value otherwise.
func (o *AlertingData) GetReportOnlyErrors() bool {
	if o == nil || o.ReportOnlyErrors == nil {
		var ret bool
		return ret
	}
	return *o.ReportOnlyErrors
}

// GetReportOnlyErrorsOk returns a tuple with the ReportOnlyErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingData) GetReportOnlyErrorsOk() (*bool, bool) {
	if o == nil || o.ReportOnlyErrors == nil {
		return nil, false
	}
	return o.ReportOnlyErrors, true
}

// HasReportOnlyErrors returns a boolean if a field has been set.
func (o *AlertingData) HasReportOnlyErrors() bool {
	if o != nil && o.ReportOnlyErrors != nil {
		return true
	}

	return false
}

// SetReportOnlyErrors gets a reference to the given bool and assigns it to the ReportOnlyErrors field.
func (o *AlertingData) SetReportOnlyErrors(v bool) {
	o.ReportOnlyErrors = &v
}

// GetSendAlertsToYb returns the SendAlertsToYb field value if set, zero value otherwise.
func (o *AlertingData) GetSendAlertsToYb() bool {
	if o == nil || o.SendAlertsToYb == nil {
		var ret bool
		return ret
	}
	return *o.SendAlertsToYb
}

// GetSendAlertsToYbOk returns a tuple with the SendAlertsToYb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingData) GetSendAlertsToYbOk() (*bool, bool) {
	if o == nil || o.SendAlertsToYb == nil {
		return nil, false
	}
	return o.SendAlertsToYb, true
}

// HasSendAlertsToYb returns a boolean if a field has been set.
func (o *AlertingData) HasSendAlertsToYb() bool {
	if o != nil && o.SendAlertsToYb != nil {
		return true
	}

	return false
}

// SetSendAlertsToYb gets a reference to the given bool and assigns it to the SendAlertsToYb field.
func (o *AlertingData) SetSendAlertsToYb(v bool) {
	o.SendAlertsToYb = &v
}

// GetStatusUpdateIntervalMs returns the StatusUpdateIntervalMs field value if set, zero value otherwise.
func (o *AlertingData) GetStatusUpdateIntervalMs() int64 {
	if o == nil || o.StatusUpdateIntervalMs == nil {
		var ret int64
		return ret
	}
	return *o.StatusUpdateIntervalMs
}

// GetStatusUpdateIntervalMsOk returns a tuple with the StatusUpdateIntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingData) GetStatusUpdateIntervalMsOk() (*int64, bool) {
	if o == nil || o.StatusUpdateIntervalMs == nil {
		return nil, false
	}
	return o.StatusUpdateIntervalMs, true
}

// HasStatusUpdateIntervalMs returns a boolean if a field has been set.
func (o *AlertingData) HasStatusUpdateIntervalMs() bool {
	if o != nil && o.StatusUpdateIntervalMs != nil {
		return true
	}

	return false
}

// SetStatusUpdateIntervalMs gets a reference to the given int64 and assigns it to the StatusUpdateIntervalMs field.
func (o *AlertingData) SetStatusUpdateIntervalMs(v int64) {
	o.StatusUpdateIntervalMs = &v
}

func (o AlertingData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveAlertNotificationIntervalMs != nil {
		toSerialize["activeAlertNotificationIntervalMs"] = o.ActiveAlertNotificationIntervalMs
	}
	if o.AlertingEmail != nil {
		toSerialize["alertingEmail"] = o.AlertingEmail
	}
	if o.CheckIntervalMs != nil {
		toSerialize["checkIntervalMs"] = o.CheckIntervalMs
	}
	if o.ReportOnlyErrors != nil {
		toSerialize["reportOnlyErrors"] = o.ReportOnlyErrors
	}
	if o.SendAlertsToYb != nil {
		toSerialize["sendAlertsToYb"] = o.SendAlertsToYb
	}
	if o.StatusUpdateIntervalMs != nil {
		toSerialize["statusUpdateIntervalMs"] = o.StatusUpdateIntervalMs
	}
	return json.Marshal(toSerialize)
}

type NullableAlertingData struct {
	value *AlertingData
	isSet bool
}

func (v NullableAlertingData) Get() *AlertingData {
	return v.value
}

func (v *NullableAlertingData) Set(val *AlertingData) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingData) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingData(val *AlertingData) *NullableAlertingData {
	return &NullableAlertingData{value: val, isSet: true}
}

func (v NullableAlertingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


