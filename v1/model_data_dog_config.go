/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// DataDogConfig struct for DataDogConfig
type DataDogConfig struct {
	TelemetryProviderConfig
	// API Key
	ApiKey *string `json:"apiKey,omitempty"`
	// Site
	Site *string `json:"site,omitempty"`
}

// NewDataDogConfig instantiates a new DataDogConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataDogConfig() *DataDogConfig {
	this := DataDogConfig{}
	return &this
}

// NewDataDogConfigWithDefaults instantiates a new DataDogConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataDogConfigWithDefaults() *DataDogConfig {
	this := DataDogConfig{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *DataDogConfig) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDogConfig) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *DataDogConfig) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *DataDogConfig) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *DataDogConfig) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDogConfig) GetSiteOk() (*string, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *DataDogConfig) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *DataDogConfig) SetSite(v string) {
	o.Site = &v
}

func (o DataDogConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTelemetryProviderConfig, errTelemetryProviderConfig := json.Marshal(o.TelemetryProviderConfig)
	if errTelemetryProviderConfig != nil {
		return []byte{}, errTelemetryProviderConfig
	}
	errTelemetryProviderConfig = json.Unmarshal([]byte(serializedTelemetryProviderConfig), &toSerialize)
	if errTelemetryProviderConfig != nil {
		return []byte{}, errTelemetryProviderConfig
	}
	if o.ApiKey != nil {
		toSerialize["apiKey"] = o.ApiKey
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	return json.Marshal(toSerialize)
}

type NullableDataDogConfig struct {
	value *DataDogConfig
	isSet bool
}

func (v NullableDataDogConfig) Get() *DataDogConfig {
	return v.value
}

func (v *NullableDataDogConfig) Set(val *DataDogConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDataDogConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDataDogConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataDogConfig(val *DataDogConfig) *NullableDataDogConfig {
	return &NullableDataDogConfig{value: val, isSet: true}
}

func (v NullableDataDogConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataDogConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


