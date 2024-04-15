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

// CustomerAlertData Format of an alert, used by the API and UI to validate data against input constraints
type CustomerAlertData struct {
	AlertingData AlertingData `json:"alertingData"`
	CallhomeLevel string `json:"callhomeLevel"`
	// Alert code
	Code *string `json:"code,omitempty"`
	// Email password confirmation
	ConfirmPassword *string `json:"confirmPassword,omitempty"`
	// Alert email address
	Email *string `json:"email,omitempty"`
	// Features
	Features *map[string]map[string]interface{} `json:"features,omitempty"`
	// Alert name
	Name *string `json:"name,omitempty"`
	// Email password
	Password *string `json:"password,omitempty"`
	SmtpData SmtpData `json:"smtpData"`
}

// NewCustomerAlertData instantiates a new CustomerAlertData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAlertData(alertingData AlertingData, callhomeLevel string, smtpData SmtpData) *CustomerAlertData {
	this := CustomerAlertData{}
	this.AlertingData = alertingData
	this.CallhomeLevel = callhomeLevel
	this.SmtpData = smtpData
	return &this
}

// NewCustomerAlertDataWithDefaults instantiates a new CustomerAlertData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAlertDataWithDefaults() *CustomerAlertData {
	this := CustomerAlertData{}
	return &this
}

// GetAlertingData returns the AlertingData field value
func (o *CustomerAlertData) GetAlertingData() AlertingData {
	if o == nil {
		var ret AlertingData
		return ret
	}

	return o.AlertingData
}

// GetAlertingDataOk returns a tuple with the AlertingData field value
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetAlertingDataOk() (*AlertingData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AlertingData, true
}

// SetAlertingData sets field value
func (o *CustomerAlertData) SetAlertingData(v AlertingData) {
	o.AlertingData = v
}

// GetCallhomeLevel returns the CallhomeLevel field value
func (o *CustomerAlertData) GetCallhomeLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallhomeLevel
}

// GetCallhomeLevelOk returns a tuple with the CallhomeLevel field value
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetCallhomeLevelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallhomeLevel, true
}

// SetCallhomeLevel sets field value
func (o *CustomerAlertData) SetCallhomeLevel(v string) {
	o.CallhomeLevel = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CustomerAlertData) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CustomerAlertData) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CustomerAlertData) SetCode(v string) {
	o.Code = &v
}

// GetConfirmPassword returns the ConfirmPassword field value if set, zero value otherwise.
func (o *CustomerAlertData) GetConfirmPassword() string {
	if o == nil || o.ConfirmPassword == nil {
		var ret string
		return ret
	}
	return *o.ConfirmPassword
}

// GetConfirmPasswordOk returns a tuple with the ConfirmPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetConfirmPasswordOk() (*string, bool) {
	if o == nil || o.ConfirmPassword == nil {
		return nil, false
	}
	return o.ConfirmPassword, true
}

// HasConfirmPassword returns a boolean if a field has been set.
func (o *CustomerAlertData) HasConfirmPassword() bool {
	if o != nil && o.ConfirmPassword != nil {
		return true
	}

	return false
}

// SetConfirmPassword gets a reference to the given string and assigns it to the ConfirmPassword field.
func (o *CustomerAlertData) SetConfirmPassword(v string) {
	o.ConfirmPassword = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerAlertData) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerAlertData) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerAlertData) SetEmail(v string) {
	o.Email = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *CustomerAlertData) GetFeatures() map[string]map[string]interface{} {
	if o == nil || o.Features == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetFeaturesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *CustomerAlertData) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given map[string]map[string]interface{} and assigns it to the Features field.
func (o *CustomerAlertData) SetFeatures(v map[string]map[string]interface{}) {
	o.Features = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerAlertData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerAlertData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerAlertData) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CustomerAlertData) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CustomerAlertData) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CustomerAlertData) SetPassword(v string) {
	o.Password = &v
}

// GetSmtpData returns the SmtpData field value
func (o *CustomerAlertData) GetSmtpData() SmtpData {
	if o == nil {
		var ret SmtpData
		return ret
	}

	return o.SmtpData
}

// GetSmtpDataOk returns a tuple with the SmtpData field value
// and a boolean to check if the value has been set.
func (o *CustomerAlertData) GetSmtpDataOk() (*SmtpData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SmtpData, true
}

// SetSmtpData sets field value
func (o *CustomerAlertData) SetSmtpData(v SmtpData) {
	o.SmtpData = v
}

func (o CustomerAlertData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alertingData"] = o.AlertingData
	}
	if true {
		toSerialize["callhomeLevel"] = o.CallhomeLevel
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.ConfirmPassword != nil {
		toSerialize["confirmPassword"] = o.ConfirmPassword
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["smtpData"] = o.SmtpData
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAlertData struct {
	value *CustomerAlertData
	isSet bool
}

func (v NullableCustomerAlertData) Get() *CustomerAlertData {
	return v.value
}

func (v *NullableCustomerAlertData) Set(val *CustomerAlertData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAlertData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAlertData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAlertData(val *CustomerAlertData) *NullableCustomerAlertData {
	return &NullableCustomerAlertData{value: val, isSet: true}
}

func (v NullableCustomerAlertData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAlertData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


