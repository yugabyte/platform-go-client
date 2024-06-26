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

// CustomerLicense Customer Licenses. This helps customer to upload licenses for the thirdparrty softwares if required.
type CustomerLicense struct {
	// Creation date of license
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Customer UUID that owns this license
	CustomerUUID *string `json:"customerUUID,omitempty"`
	// License File Path
	License string `json:"license"`
	// Type of the license
	LicenseType string `json:"licenseType"`
	// License UUID
	LicenseUUID *string `json:"licenseUUID,omitempty"`
}

// NewCustomerLicense instantiates a new CustomerLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerLicense(license string, licenseType string) *CustomerLicense {
	this := CustomerLicense{}
	this.License = license
	this.LicenseType = licenseType
	return &this
}

// NewCustomerLicenseWithDefaults instantiates a new CustomerLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerLicenseWithDefaults() *CustomerLicense {
	this := CustomerLicense{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *CustomerLicense) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLicense) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *CustomerLicense) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *CustomerLicense) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetCustomerUUID returns the CustomerUUID field value if set, zero value otherwise.
func (o *CustomerLicense) GetCustomerUUID() string {
	if o == nil || o.CustomerUUID == nil {
		var ret string
		return ret
	}
	return *o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLicense) GetCustomerUUIDOk() (*string, bool) {
	if o == nil || o.CustomerUUID == nil {
		return nil, false
	}
	return o.CustomerUUID, true
}

// HasCustomerUUID returns a boolean if a field has been set.
func (o *CustomerLicense) HasCustomerUUID() bool {
	if o != nil && o.CustomerUUID != nil {
		return true
	}

	return false
}

// SetCustomerUUID gets a reference to the given string and assigns it to the CustomerUUID field.
func (o *CustomerLicense) SetCustomerUUID(v string) {
	o.CustomerUUID = &v
}

// GetLicense returns the License field value
func (o *CustomerLicense) GetLicense() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *CustomerLicense) GetLicenseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value
func (o *CustomerLicense) SetLicense(v string) {
	o.License = v
}

// GetLicenseType returns the LicenseType field value
func (o *CustomerLicense) GetLicenseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value
// and a boolean to check if the value has been set.
func (o *CustomerLicense) GetLicenseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LicenseType, true
}

// SetLicenseType sets field value
func (o *CustomerLicense) SetLicenseType(v string) {
	o.LicenseType = v
}

// GetLicenseUUID returns the LicenseUUID field value if set, zero value otherwise.
func (o *CustomerLicense) GetLicenseUUID() string {
	if o == nil || o.LicenseUUID == nil {
		var ret string
		return ret
	}
	return *o.LicenseUUID
}

// GetLicenseUUIDOk returns a tuple with the LicenseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLicense) GetLicenseUUIDOk() (*string, bool) {
	if o == nil || o.LicenseUUID == nil {
		return nil, false
	}
	return o.LicenseUUID, true
}

// HasLicenseUUID returns a boolean if a field has been set.
func (o *CustomerLicense) HasLicenseUUID() bool {
	if o != nil && o.LicenseUUID != nil {
		return true
	}

	return false
}

// SetLicenseUUID gets a reference to the given string and assigns it to the LicenseUUID field.
func (o *CustomerLicense) SetLicenseUUID(v string) {
	o.LicenseUUID = &v
}

func (o CustomerLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.CustomerUUID != nil {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if true {
		toSerialize["license"] = o.License
	}
	if true {
		toSerialize["licenseType"] = o.LicenseType
	}
	if o.LicenseUUID != nil {
		toSerialize["licenseUUID"] = o.LicenseUUID
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerLicense struct {
	value *CustomerLicense
	isSet bool
}

func (v NullableCustomerLicense) Get() *CustomerLicense {
	return v.value
}

func (v *NullableCustomerLicense) Set(val *CustomerLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerLicense(val *CustomerLicense) *NullableCustomerLicense {
	return &NullableCustomerLicense{value: val, isSet: true}
}

func (v NullableCustomerLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


