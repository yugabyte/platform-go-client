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

// CertificateInfo SSL certificate used by the universe
type CertificateInfo struct {
	// Type of the certificate
	CertType *string `json:"certType,omitempty"`
	// Certificate path
	Certificate *string `json:"certificate,omitempty"`
	// The certificate file's checksum
	Checksum *string `json:"checksum,omitempty"`
	CustomCertPathParams CustomCertInfo `json:"customCertPathParams"`
	CustomHCPKICertInfo HashicorpVaultConfigParams `json:"customHCPKICertInfo"`
	CustomServerCertInfo CustomServerCertInfo `json:"customServerCertInfo"`
	// Customer UUID of the backup which it belongs to
	CustomerUUID *string `json:"customerUUID,omitempty"`
	// The certificate's expiry date
	ExpiryDateIso *time.Time `json:"expiryDateIso,omitempty"`
	// Indicates whether the certificate is in use. This value is `true` if the universe contains a reference to the certificate.
	InUse *bool `json:"inUse,omitempty"`
	// Certificate label
	Label *string `json:"label,omitempty"`
	// Private key path
	PrivateKey *string `json:"privateKey,omitempty"`
	// The certificate's creation date
	StartDateIso *time.Time `json:"startDateIso,omitempty"`
	UniverseDetailSubsets []UniverseDetailSubset `json:"universeDetailSubsets"`
	// Associated universe details for the certificate
	UniverseDetails *[]UniverseDetailSubset `json:"universeDetails,omitempty"`
	// Certificate UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewCertificateInfo instantiates a new CertificateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateInfo(customCertPathParams CustomCertInfo, customHCPKICertInfo HashicorpVaultConfigParams, customServerCertInfo CustomServerCertInfo, universeDetailSubsets []UniverseDetailSubset, ) *CertificateInfo {
	this := CertificateInfo{}
	this.CustomCertPathParams = customCertPathParams
	this.CustomHCPKICertInfo = customHCPKICertInfo
	this.CustomServerCertInfo = customServerCertInfo
	this.UniverseDetailSubsets = universeDetailSubsets
	return &this
}

// NewCertificateInfoWithDefaults instantiates a new CertificateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateInfoWithDefaults() *CertificateInfo {
	this := CertificateInfo{}
	return &this
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *CertificateInfo) GetCertType() string {
	if o == nil || o.CertType == nil {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetCertTypeOk() (*string, bool) {
	if o == nil || o.CertType == nil {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *CertificateInfo) HasCertType() bool {
	if o != nil && o.CertType != nil {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *CertificateInfo) SetCertType(v string) {
	o.CertType = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CertificateInfo) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CertificateInfo) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CertificateInfo) SetCertificate(v string) {
	o.Certificate = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *CertificateInfo) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *CertificateInfo) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *CertificateInfo) SetChecksum(v string) {
	o.Checksum = &v
}

// GetCustomCertPathParams returns the CustomCertPathParams field value
func (o *CertificateInfo) GetCustomCertPathParams() CustomCertInfo {
	if o == nil  {
		var ret CustomCertInfo
		return ret
	}

	return o.CustomCertPathParams
}

// GetCustomCertPathParamsOk returns a tuple with the CustomCertPathParams field value
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetCustomCertPathParamsOk() (*CustomCertInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomCertPathParams, true
}

// SetCustomCertPathParams sets field value
func (o *CertificateInfo) SetCustomCertPathParams(v CustomCertInfo) {
	o.CustomCertPathParams = v
}

// GetCustomHCPKICertInfo returns the CustomHCPKICertInfo field value
func (o *CertificateInfo) GetCustomHCPKICertInfo() HashicorpVaultConfigParams {
	if o == nil  {
		var ret HashicorpVaultConfigParams
		return ret
	}

	return o.CustomHCPKICertInfo
}

// GetCustomHCPKICertInfoOk returns a tuple with the CustomHCPKICertInfo field value
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetCustomHCPKICertInfoOk() (*HashicorpVaultConfigParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomHCPKICertInfo, true
}

// SetCustomHCPKICertInfo sets field value
func (o *CertificateInfo) SetCustomHCPKICertInfo(v HashicorpVaultConfigParams) {
	o.CustomHCPKICertInfo = v
}

// GetCustomServerCertInfo returns the CustomServerCertInfo field value
func (o *CertificateInfo) GetCustomServerCertInfo() CustomServerCertInfo {
	if o == nil  {
		var ret CustomServerCertInfo
		return ret
	}

	return o.CustomServerCertInfo
}

// GetCustomServerCertInfoOk returns a tuple with the CustomServerCertInfo field value
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetCustomServerCertInfoOk() (*CustomServerCertInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomServerCertInfo, true
}

// SetCustomServerCertInfo sets field value
func (o *CertificateInfo) SetCustomServerCertInfo(v CustomServerCertInfo) {
	o.CustomServerCertInfo = v
}

// GetCustomerUUID returns the CustomerUUID field value if set, zero value otherwise.
func (o *CertificateInfo) GetCustomerUUID() string {
	if o == nil || o.CustomerUUID == nil {
		var ret string
		return ret
	}
	return *o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetCustomerUUIDOk() (*string, bool) {
	if o == nil || o.CustomerUUID == nil {
		return nil, false
	}
	return o.CustomerUUID, true
}

// HasCustomerUUID returns a boolean if a field has been set.
func (o *CertificateInfo) HasCustomerUUID() bool {
	if o != nil && o.CustomerUUID != nil {
		return true
	}

	return false
}

// SetCustomerUUID gets a reference to the given string and assigns it to the CustomerUUID field.
func (o *CertificateInfo) SetCustomerUUID(v string) {
	o.CustomerUUID = &v
}

// GetExpiryDateIso returns the ExpiryDateIso field value if set, zero value otherwise.
func (o *CertificateInfo) GetExpiryDateIso() time.Time {
	if o == nil || o.ExpiryDateIso == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDateIso
}

// GetExpiryDateIsoOk returns a tuple with the ExpiryDateIso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetExpiryDateIsoOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDateIso == nil {
		return nil, false
	}
	return o.ExpiryDateIso, true
}

// HasExpiryDateIso returns a boolean if a field has been set.
func (o *CertificateInfo) HasExpiryDateIso() bool {
	if o != nil && o.ExpiryDateIso != nil {
		return true
	}

	return false
}

// SetExpiryDateIso gets a reference to the given time.Time and assigns it to the ExpiryDateIso field.
func (o *CertificateInfo) SetExpiryDateIso(v time.Time) {
	o.ExpiryDateIso = &v
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *CertificateInfo) GetInUse() bool {
	if o == nil || o.InUse == nil {
		var ret bool
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetInUseOk() (*bool, bool) {
	if o == nil || o.InUse == nil {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *CertificateInfo) HasInUse() bool {
	if o != nil && o.InUse != nil {
		return true
	}

	return false
}

// SetInUse gets a reference to the given bool and assigns it to the InUse field.
func (o *CertificateInfo) SetInUse(v bool) {
	o.InUse = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CertificateInfo) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CertificateInfo) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CertificateInfo) SetLabel(v string) {
	o.Label = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *CertificateInfo) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *CertificateInfo) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *CertificateInfo) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetStartDateIso returns the StartDateIso field value if set, zero value otherwise.
func (o *CertificateInfo) GetStartDateIso() time.Time {
	if o == nil || o.StartDateIso == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateIso
}

// GetStartDateIsoOk returns a tuple with the StartDateIso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetStartDateIsoOk() (*time.Time, bool) {
	if o == nil || o.StartDateIso == nil {
		return nil, false
	}
	return o.StartDateIso, true
}

// HasStartDateIso returns a boolean if a field has been set.
func (o *CertificateInfo) HasStartDateIso() bool {
	if o != nil && o.StartDateIso != nil {
		return true
	}

	return false
}

// SetStartDateIso gets a reference to the given time.Time and assigns it to the StartDateIso field.
func (o *CertificateInfo) SetStartDateIso(v time.Time) {
	o.StartDateIso = &v
}

// GetUniverseDetailSubsets returns the UniverseDetailSubsets field value
func (o *CertificateInfo) GetUniverseDetailSubsets() []UniverseDetailSubset {
	if o == nil  {
		var ret []UniverseDetailSubset
		return ret
	}

	return o.UniverseDetailSubsets
}

// GetUniverseDetailSubsetsOk returns a tuple with the UniverseDetailSubsets field value
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetUniverseDetailSubsetsOk() (*[]UniverseDetailSubset, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseDetailSubsets, true
}

// SetUniverseDetailSubsets sets field value
func (o *CertificateInfo) SetUniverseDetailSubsets(v []UniverseDetailSubset) {
	o.UniverseDetailSubsets = v
}

// GetUniverseDetails returns the UniverseDetails field value if set, zero value otherwise.
func (o *CertificateInfo) GetUniverseDetails() []UniverseDetailSubset {
	if o == nil || o.UniverseDetails == nil {
		var ret []UniverseDetailSubset
		return ret
	}
	return *o.UniverseDetails
}

// GetUniverseDetailsOk returns a tuple with the UniverseDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetUniverseDetailsOk() (*[]UniverseDetailSubset, bool) {
	if o == nil || o.UniverseDetails == nil {
		return nil, false
	}
	return o.UniverseDetails, true
}

// HasUniverseDetails returns a boolean if a field has been set.
func (o *CertificateInfo) HasUniverseDetails() bool {
	if o != nil && o.UniverseDetails != nil {
		return true
	}

	return false
}

// SetUniverseDetails gets a reference to the given []UniverseDetailSubset and assigns it to the UniverseDetails field.
func (o *CertificateInfo) SetUniverseDetails(v []UniverseDetailSubset) {
	o.UniverseDetails = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CertificateInfo) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CertificateInfo) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CertificateInfo) SetUuid(v string) {
	o.Uuid = &v
}

func (o CertificateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertType != nil {
		toSerialize["certType"] = o.CertType
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if true {
		toSerialize["customCertPathParams"] = o.CustomCertPathParams
	}
	if true {
		toSerialize["customHCPKICertInfo"] = o.CustomHCPKICertInfo
	}
	if true {
		toSerialize["customServerCertInfo"] = o.CustomServerCertInfo
	}
	if o.CustomerUUID != nil {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if o.ExpiryDateIso != nil {
		toSerialize["expiryDateIso"] = o.ExpiryDateIso
	}
	if o.InUse != nil {
		toSerialize["inUse"] = o.InUse
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.StartDateIso != nil {
		toSerialize["startDateIso"] = o.StartDateIso
	}
	if true {
		toSerialize["universeDetailSubsets"] = o.UniverseDetailSubsets
	}
	if o.UniverseDetails != nil {
		toSerialize["universeDetails"] = o.UniverseDetails
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateInfo struct {
	value *CertificateInfo
	isSet bool
}

func (v NullableCertificateInfo) Get() *CertificateInfo {
	return v.value
}

func (v *NullableCertificateInfo) Set(val *CertificateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateInfo(val *CertificateInfo) *NullableCertificateInfo {
	return &NullableCertificateInfo{value: val, isSet: true}
}

func (v NullableCertificateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


