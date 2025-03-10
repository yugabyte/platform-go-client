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

// CertificateParams Certificate Params is used to validate constraints for the custom certificate Data
type CertificateParams struct {
	CertContent string `json:"certContent"`
	CertExpiry int64 `json:"certExpiry"`
	CertStart int64 `json:"certStart"`
	CertType string `json:"certType"`
	CustomCertInfo *CustomCertInfo `json:"customCertInfo,omitempty"`
	CustomServerCertData *CustomServerCertData `json:"customServerCertData,omitempty"`
	HcVaultCertParams *HashicorpVaultConfigParams `json:"hcVaultCertParams,omitempty"`
	KeyContent *string `json:"keyContent,omitempty"`
	Label string `json:"label"`
}

// NewCertificateParams instantiates a new CertificateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateParams(certContent string, certExpiry int64, certStart int64, certType string, label string) *CertificateParams {
	this := CertificateParams{}
	this.CertContent = certContent
	this.CertExpiry = certExpiry
	this.CertStart = certStart
	this.CertType = certType
	this.Label = label
	return &this
}

// NewCertificateParamsWithDefaults instantiates a new CertificateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateParamsWithDefaults() *CertificateParams {
	this := CertificateParams{}
	return &this
}

// GetCertContent returns the CertContent field value
func (o *CertificateParams) GetCertContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertContent
}

// GetCertContentOk returns a tuple with the CertContent field value
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetCertContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertContent, true
}

// SetCertContent sets field value
func (o *CertificateParams) SetCertContent(v string) {
	o.CertContent = v
}

// GetCertExpiry returns the CertExpiry field value
func (o *CertificateParams) GetCertExpiry() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CertExpiry
}

// GetCertExpiryOk returns a tuple with the CertExpiry field value
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetCertExpiryOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertExpiry, true
}

// SetCertExpiry sets field value
func (o *CertificateParams) SetCertExpiry(v int64) {
	o.CertExpiry = v
}

// GetCertStart returns the CertStart field value
func (o *CertificateParams) GetCertStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CertStart
}

// GetCertStartOk returns a tuple with the CertStart field value
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetCertStartOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertStart, true
}

// SetCertStart sets field value
func (o *CertificateParams) SetCertStart(v int64) {
	o.CertStart = v
}

// GetCertType returns the CertType field value
func (o *CertificateParams) GetCertType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetCertTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertType, true
}

// SetCertType sets field value
func (o *CertificateParams) SetCertType(v string) {
	o.CertType = v
}

// GetCustomCertInfo returns the CustomCertInfo field value if set, zero value otherwise.
func (o *CertificateParams) GetCustomCertInfo() CustomCertInfo {
	if o == nil || o.CustomCertInfo == nil {
		var ret CustomCertInfo
		return ret
	}
	return *o.CustomCertInfo
}

// GetCustomCertInfoOk returns a tuple with the CustomCertInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetCustomCertInfoOk() (*CustomCertInfo, bool) {
	if o == nil || o.CustomCertInfo == nil {
		return nil, false
	}
	return o.CustomCertInfo, true
}

// HasCustomCertInfo returns a boolean if a field has been set.
func (o *CertificateParams) HasCustomCertInfo() bool {
	if o != nil && o.CustomCertInfo != nil {
		return true
	}

	return false
}

// SetCustomCertInfo gets a reference to the given CustomCertInfo and assigns it to the CustomCertInfo field.
func (o *CertificateParams) SetCustomCertInfo(v CustomCertInfo) {
	o.CustomCertInfo = &v
}

// GetCustomServerCertData returns the CustomServerCertData field value if set, zero value otherwise.
func (o *CertificateParams) GetCustomServerCertData() CustomServerCertData {
	if o == nil || o.CustomServerCertData == nil {
		var ret CustomServerCertData
		return ret
	}
	return *o.CustomServerCertData
}

// GetCustomServerCertDataOk returns a tuple with the CustomServerCertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetCustomServerCertDataOk() (*CustomServerCertData, bool) {
	if o == nil || o.CustomServerCertData == nil {
		return nil, false
	}
	return o.CustomServerCertData, true
}

// HasCustomServerCertData returns a boolean if a field has been set.
func (o *CertificateParams) HasCustomServerCertData() bool {
	if o != nil && o.CustomServerCertData != nil {
		return true
	}

	return false
}

// SetCustomServerCertData gets a reference to the given CustomServerCertData and assigns it to the CustomServerCertData field.
func (o *CertificateParams) SetCustomServerCertData(v CustomServerCertData) {
	o.CustomServerCertData = &v
}

// GetHcVaultCertParams returns the HcVaultCertParams field value if set, zero value otherwise.
func (o *CertificateParams) GetHcVaultCertParams() HashicorpVaultConfigParams {
	if o == nil || o.HcVaultCertParams == nil {
		var ret HashicorpVaultConfigParams
		return ret
	}
	return *o.HcVaultCertParams
}

// GetHcVaultCertParamsOk returns a tuple with the HcVaultCertParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetHcVaultCertParamsOk() (*HashicorpVaultConfigParams, bool) {
	if o == nil || o.HcVaultCertParams == nil {
		return nil, false
	}
	return o.HcVaultCertParams, true
}

// HasHcVaultCertParams returns a boolean if a field has been set.
func (o *CertificateParams) HasHcVaultCertParams() bool {
	if o != nil && o.HcVaultCertParams != nil {
		return true
	}

	return false
}

// SetHcVaultCertParams gets a reference to the given HashicorpVaultConfigParams and assigns it to the HcVaultCertParams field.
func (o *CertificateParams) SetHcVaultCertParams(v HashicorpVaultConfigParams) {
	o.HcVaultCertParams = &v
}

// GetKeyContent returns the KeyContent field value if set, zero value otherwise.
func (o *CertificateParams) GetKeyContent() string {
	if o == nil || o.KeyContent == nil {
		var ret string
		return ret
	}
	return *o.KeyContent
}

// GetKeyContentOk returns a tuple with the KeyContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetKeyContentOk() (*string, bool) {
	if o == nil || o.KeyContent == nil {
		return nil, false
	}
	return o.KeyContent, true
}

// HasKeyContent returns a boolean if a field has been set.
func (o *CertificateParams) HasKeyContent() bool {
	if o != nil && o.KeyContent != nil {
		return true
	}

	return false
}

// SetKeyContent gets a reference to the given string and assigns it to the KeyContent field.
func (o *CertificateParams) SetKeyContent(v string) {
	o.KeyContent = &v
}

// GetLabel returns the Label field value
func (o *CertificateParams) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CertificateParams) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CertificateParams) SetLabel(v string) {
	o.Label = v
}

func (o CertificateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["certContent"] = o.CertContent
	}
	if true {
		toSerialize["certExpiry"] = o.CertExpiry
	}
	if true {
		toSerialize["certStart"] = o.CertStart
	}
	if true {
		toSerialize["certType"] = o.CertType
	}
	if o.CustomCertInfo != nil {
		toSerialize["customCertInfo"] = o.CustomCertInfo
	}
	if o.CustomServerCertData != nil {
		toSerialize["customServerCertData"] = o.CustomServerCertData
	}
	if o.HcVaultCertParams != nil {
		toSerialize["hcVaultCertParams"] = o.HcVaultCertParams
	}
	if o.KeyContent != nil {
		toSerialize["keyContent"] = o.KeyContent
	}
	if true {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateParams struct {
	value *CertificateParams
	isSet bool
}

func (v NullableCertificateParams) Get() *CertificateParams {
	return v.value
}

func (v *NullableCertificateParams) Set(val *CertificateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateParams(val *CertificateParams) *NullableCertificateParams {
	return &NullableCertificateParams{value: val, isSet: true}
}

func (v NullableCertificateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


