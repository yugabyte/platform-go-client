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

// TroubleshootingPlatform Troubleshooting Platform Model
type TroubleshootingPlatform struct {
	// YBA API Token
	ApiToken string `json:"apiToken"`
	// Customer UUID
	CustomerUUID string `json:"customerUUID"`
	// Metrics Scrape Period Seconds
	MetricsScrapePeriodSecs int64 `json:"metricsScrapePeriodSecs"`
	// Metrics URL
	MetricsUrl string `json:"metricsUrl"`
	// TP API Token
	TpApiToken *string `json:"tpApiToken,omitempty"`
	// Troubleshooting Platform URL
	TpUrl string `json:"tpUrl"`
	// Troubleshooting Platform UUID
	Uuid string `json:"uuid"`
	// YBA URL
	YbaUrl string `json:"ybaUrl"`
}

// NewTroubleshootingPlatform instantiates a new TroubleshootingPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTroubleshootingPlatform(apiToken string, customerUUID string, metricsScrapePeriodSecs int64, metricsUrl string, tpUrl string, uuid string, ybaUrl string) *TroubleshootingPlatform {
	this := TroubleshootingPlatform{}
	this.ApiToken = apiToken
	this.CustomerUUID = customerUUID
	this.MetricsScrapePeriodSecs = metricsScrapePeriodSecs
	this.MetricsUrl = metricsUrl
	this.TpUrl = tpUrl
	this.Uuid = uuid
	this.YbaUrl = ybaUrl
	return &this
}

// NewTroubleshootingPlatformWithDefaults instantiates a new TroubleshootingPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTroubleshootingPlatformWithDefaults() *TroubleshootingPlatform {
	this := TroubleshootingPlatform{}
	return &this
}

// GetApiToken returns the ApiToken field value
func (o *TroubleshootingPlatform) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *TroubleshootingPlatform) GetApiTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value
func (o *TroubleshootingPlatform) SetApiToken(v string) {
	o.ApiToken = v
}

// GetCustomerUUID returns the CustomerUUID field value
func (o *TroubleshootingPlatform) GetCustomerUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value
// and a boolean to check if the value has been set.
func (o *TroubleshootingPlatform) GetCustomerUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerUUID, true
}

// SetCustomerUUID sets field value
func (o *TroubleshootingPlatform) SetCustomerUUID(v string) {
	o.CustomerUUID = v
}

// GetMetricsScrapePeriodSecs returns the MetricsScrapePeriodSecs field value
func (o *TroubleshootingPlatform) GetMetricsScrapePeriodSecs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricsScrapePeriodSecs
}

// GetMetricsScrapePeriodSecsOk returns a tuple with the MetricsScrapePeriodSecs field value
// and a boolean to check if the value has been set.
func (o *TroubleshootingPlatform) GetMetricsScrapePeriodSecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricsScrapePeriodSecs, true
}

// SetMetricsScrapePeriodSecs sets field value
func (o *TroubleshootingPlatform) SetMetricsScrapePeriodSecs(v int64) {
	o.MetricsScrapePeriodSecs = v
}

// GetMetricsUrl returns the MetricsUrl field value
func (o *TroubleshootingPlatform) GetMetricsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricsUrl
}

// GetMetricsUrlOk returns a tuple with the MetricsUrl field value
// and a boolean to check if the value has been set.
func (o *TroubleshootingPlatform) GetMetricsUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricsUrl, true
}

// SetMetricsUrl sets field value
func (o *TroubleshootingPlatform) SetMetricsUrl(v string) {
	o.MetricsUrl = v
}

// GetTpApiToken returns the TpApiToken field value if set, zero value otherwise.
func (o *TroubleshootingPlatform) GetTpApiToken() string {
	if o == nil || o.TpApiToken == nil {
		var ret string
		return ret
	}
	return *o.TpApiToken
}

// GetTpApiTokenOk returns a tuple with the TpApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootingPlatform) GetTpApiTokenOk() (*string, bool) {
	if o == nil || o.TpApiToken == nil {
		return nil, false
	}
	return o.TpApiToken, true
}

// HasTpApiToken returns a boolean if a field has been set.
func (o *TroubleshootingPlatform) HasTpApiToken() bool {
	if o != nil && o.TpApiToken != nil {
		return true
	}

	return false
}

// SetTpApiToken gets a reference to the given string and assigns it to the TpApiToken field.
func (o *TroubleshootingPlatform) SetTpApiToken(v string) {
	o.TpApiToken = &v
}

// GetTpUrl returns the TpUrl field value
func (o *TroubleshootingPlatform) GetTpUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TpUrl
}

// GetTpUrlOk returns a tuple with the TpUrl field value
// and a boolean to check if the value has been set.
func (o *TroubleshootingPlatform) GetTpUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TpUrl, true
}

// SetTpUrl sets field value
func (o *TroubleshootingPlatform) SetTpUrl(v string) {
	o.TpUrl = v
}

// GetUuid returns the Uuid field value
func (o *TroubleshootingPlatform) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *TroubleshootingPlatform) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *TroubleshootingPlatform) SetUuid(v string) {
	o.Uuid = v
}

// GetYbaUrl returns the YbaUrl field value
func (o *TroubleshootingPlatform) GetYbaUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbaUrl
}

// GetYbaUrlOk returns a tuple with the YbaUrl field value
// and a boolean to check if the value has been set.
func (o *TroubleshootingPlatform) GetYbaUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbaUrl, true
}

// SetYbaUrl sets field value
func (o *TroubleshootingPlatform) SetYbaUrl(v string) {
	o.YbaUrl = v
}

func (o TroubleshootingPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiToken"] = o.ApiToken
	}
	if true {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if true {
		toSerialize["metricsScrapePeriodSecs"] = o.MetricsScrapePeriodSecs
	}
	if true {
		toSerialize["metricsUrl"] = o.MetricsUrl
	}
	if o.TpApiToken != nil {
		toSerialize["tpApiToken"] = o.TpApiToken
	}
	if true {
		toSerialize["tpUrl"] = o.TpUrl
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["ybaUrl"] = o.YbaUrl
	}
	return json.Marshal(toSerialize)
}

type NullableTroubleshootingPlatform struct {
	value *TroubleshootingPlatform
	isSet bool
}

func (v NullableTroubleshootingPlatform) Get() *TroubleshootingPlatform {
	return v.value
}

func (v *NullableTroubleshootingPlatform) Set(val *TroubleshootingPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableTroubleshootingPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableTroubleshootingPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTroubleshootingPlatform(val *TroubleshootingPlatform) *NullableTroubleshootingPlatform {
	return &NullableTroubleshootingPlatform{value: val, isSet: true}
}

func (v NullableTroubleshootingPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTroubleshootingPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

