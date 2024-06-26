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

// PackagesRequestParams Packages request parameters
type PackagesRequestParams struct {
	// Architecture Type
	ArchitectureType *string `json:"architectureType,omitempty"`
	// Archive Type
	ArchiveType *string `json:"archiveType,omitempty"`
	// Build number
	BuildNumber string `json:"buildNumber"`
	// OS Type
	OsType *string `json:"osType,omitempty"`
	// Package name
	PackageName *string `json:"packageName,omitempty"`
}

// NewPackagesRequestParams instantiates a new PackagesRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesRequestParams(buildNumber string) *PackagesRequestParams {
	this := PackagesRequestParams{}
	this.BuildNumber = buildNumber
	return &this
}

// NewPackagesRequestParamsWithDefaults instantiates a new PackagesRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesRequestParamsWithDefaults() *PackagesRequestParams {
	this := PackagesRequestParams{}
	return &this
}

// GetArchitectureType returns the ArchitectureType field value if set, zero value otherwise.
func (o *PackagesRequestParams) GetArchitectureType() string {
	if o == nil || o.ArchitectureType == nil {
		var ret string
		return ret
	}
	return *o.ArchitectureType
}

// GetArchitectureTypeOk returns a tuple with the ArchitectureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesRequestParams) GetArchitectureTypeOk() (*string, bool) {
	if o == nil || o.ArchitectureType == nil {
		return nil, false
	}
	return o.ArchitectureType, true
}

// HasArchitectureType returns a boolean if a field has been set.
func (o *PackagesRequestParams) HasArchitectureType() bool {
	if o != nil && o.ArchitectureType != nil {
		return true
	}

	return false
}

// SetArchitectureType gets a reference to the given string and assigns it to the ArchitectureType field.
func (o *PackagesRequestParams) SetArchitectureType(v string) {
	o.ArchitectureType = &v
}

// GetArchiveType returns the ArchiveType field value if set, zero value otherwise.
func (o *PackagesRequestParams) GetArchiveType() string {
	if o == nil || o.ArchiveType == nil {
		var ret string
		return ret
	}
	return *o.ArchiveType
}

// GetArchiveTypeOk returns a tuple with the ArchiveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesRequestParams) GetArchiveTypeOk() (*string, bool) {
	if o == nil || o.ArchiveType == nil {
		return nil, false
	}
	return o.ArchiveType, true
}

// HasArchiveType returns a boolean if a field has been set.
func (o *PackagesRequestParams) HasArchiveType() bool {
	if o != nil && o.ArchiveType != nil {
		return true
	}

	return false
}

// SetArchiveType gets a reference to the given string and assigns it to the ArchiveType field.
func (o *PackagesRequestParams) SetArchiveType(v string) {
	o.ArchiveType = &v
}

// GetBuildNumber returns the BuildNumber field value
func (o *PackagesRequestParams) GetBuildNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value
// and a boolean to check if the value has been set.
func (o *PackagesRequestParams) GetBuildNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BuildNumber, true
}

// SetBuildNumber sets field value
func (o *PackagesRequestParams) SetBuildNumber(v string) {
	o.BuildNumber = v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *PackagesRequestParams) GetOsType() string {
	if o == nil || o.OsType == nil {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesRequestParams) GetOsTypeOk() (*string, bool) {
	if o == nil || o.OsType == nil {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *PackagesRequestParams) HasOsType() bool {
	if o != nil && o.OsType != nil {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *PackagesRequestParams) SetOsType(v string) {
	o.OsType = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *PackagesRequestParams) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesRequestParams) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *PackagesRequestParams) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *PackagesRequestParams) SetPackageName(v string) {
	o.PackageName = &v
}

func (o PackagesRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchitectureType != nil {
		toSerialize["architectureType"] = o.ArchitectureType
	}
	if o.ArchiveType != nil {
		toSerialize["archiveType"] = o.ArchiveType
	}
	if true {
		toSerialize["buildNumber"] = o.BuildNumber
	}
	if o.OsType != nil {
		toSerialize["osType"] = o.OsType
	}
	if o.PackageName != nil {
		toSerialize["packageName"] = o.PackageName
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesRequestParams struct {
	value *PackagesRequestParams
	isSet bool
}

func (v NullablePackagesRequestParams) Get() *PackagesRequestParams {
	return v.value
}

func (v *NullablePackagesRequestParams) Set(val *PackagesRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesRequestParams(val *PackagesRequestParams) *NullablePackagesRequestParams {
	return &NullablePackagesRequestParams{value: val, isSet: true}
}

func (v NullablePackagesRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


