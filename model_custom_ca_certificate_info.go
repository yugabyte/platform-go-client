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

// CustomCaCertificateInfo Custom CA certificate
type CustomCaCertificateInfo struct {
	Active bool `json:"active"`
	// Path to CA Certificate
	Contents *string `json:"contents,omitempty"`
	// Date when certificate was added.
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	CustomerId string `json:"customerId"`
	// End date of certificate validity.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	// Start date of certificate validity.
	StartDate *time.Time `json:"startDate,omitempty"`
}

// NewCustomCaCertificateInfo instantiates a new CustomCaCertificateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomCaCertificateInfo(active bool, customerId string, id string, name string, ) *CustomCaCertificateInfo {
	this := CustomCaCertificateInfo{}
	this.Active = active
	this.CustomerId = customerId
	this.Id = id
	this.Name = name
	return &this
}

// NewCustomCaCertificateInfoWithDefaults instantiates a new CustomCaCertificateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomCaCertificateInfoWithDefaults() *CustomCaCertificateInfo {
	this := CustomCaCertificateInfo{}
	return &this
}

// GetActive returns the Active field value
func (o *CustomCaCertificateInfo) GetActive() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CustomCaCertificateInfo) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *CustomCaCertificateInfo) SetActive(v bool) {
	o.Active = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *CustomCaCertificateInfo) GetContents() string {
	if o == nil || o.Contents == nil {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCaCertificateInfo) GetContentsOk() (*string, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *CustomCaCertificateInfo) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *CustomCaCertificateInfo) SetContents(v string) {
	o.Contents = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *CustomCaCertificateInfo) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCaCertificateInfo) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *CustomCaCertificateInfo) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *CustomCaCertificateInfo) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCustomerId returns the CustomerId field value
func (o *CustomCaCertificateInfo) GetCustomerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CustomCaCertificateInfo) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CustomCaCertificateInfo) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *CustomCaCertificateInfo) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCaCertificateInfo) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *CustomCaCertificateInfo) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *CustomCaCertificateInfo) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetId returns the Id field value
func (o *CustomCaCertificateInfo) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomCaCertificateInfo) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomCaCertificateInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CustomCaCertificateInfo) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomCaCertificateInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomCaCertificateInfo) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CustomCaCertificateInfo) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCaCertificateInfo) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CustomCaCertificateInfo) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CustomCaCertificateInfo) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o CustomCaCertificateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	if o.CreatedTime != nil {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if o.ExpiryDate != nil {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableCustomCaCertificateInfo struct {
	value *CustomCaCertificateInfo
	isSet bool
}

func (v NullableCustomCaCertificateInfo) Get() *CustomCaCertificateInfo {
	return v.value
}

func (v *NullableCustomCaCertificateInfo) Set(val *CustomCaCertificateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomCaCertificateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomCaCertificateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomCaCertificateInfo(val *CustomCaCertificateInfo) *NullableCustomCaCertificateInfo {
	return &NullableCustomCaCertificateInfo{value: val, isSet: true}
}

func (v NullableCustomCaCertificateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomCaCertificateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


