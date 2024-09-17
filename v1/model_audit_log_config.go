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

// AuditLogConfig Audit Log Configuration
type AuditLogConfig struct {
	// Universe logs export active
	ExportActive *bool `json:"exportActive,omitempty"`
	// Universe logs exporter config
	UniverseLogsExporterConfig []UniverseLogsExporterConfig `json:"universeLogsExporterConfig"`
	YcqlAuditConfig *YCQLAuditConfig `json:"ycqlAuditConfig,omitempty"`
	YsqlAuditConfig *YSQLAuditConfig `json:"ysqlAuditConfig,omitempty"`
}

// NewAuditLogConfig instantiates a new AuditLogConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogConfig(universeLogsExporterConfig []UniverseLogsExporterConfig) *AuditLogConfig {
	this := AuditLogConfig{}
	this.UniverseLogsExporterConfig = universeLogsExporterConfig
	return &this
}

// NewAuditLogConfigWithDefaults instantiates a new AuditLogConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogConfigWithDefaults() *AuditLogConfig {
	this := AuditLogConfig{}
	return &this
}

// GetExportActive returns the ExportActive field value if set, zero value otherwise.
func (o *AuditLogConfig) GetExportActive() bool {
	if o == nil || o.ExportActive == nil {
		var ret bool
		return ret
	}
	return *o.ExportActive
}

// GetExportActiveOk returns a tuple with the ExportActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogConfig) GetExportActiveOk() (*bool, bool) {
	if o == nil || o.ExportActive == nil {
		return nil, false
	}
	return o.ExportActive, true
}

// HasExportActive returns a boolean if a field has been set.
func (o *AuditLogConfig) HasExportActive() bool {
	if o != nil && o.ExportActive != nil {
		return true
	}

	return false
}

// SetExportActive gets a reference to the given bool and assigns it to the ExportActive field.
func (o *AuditLogConfig) SetExportActive(v bool) {
	o.ExportActive = &v
}

// GetUniverseLogsExporterConfig returns the UniverseLogsExporterConfig field value
func (o *AuditLogConfig) GetUniverseLogsExporterConfig() []UniverseLogsExporterConfig {
	if o == nil {
		var ret []UniverseLogsExporterConfig
		return ret
	}

	return o.UniverseLogsExporterConfig
}

// GetUniverseLogsExporterConfigOk returns a tuple with the UniverseLogsExporterConfig field value
// and a boolean to check if the value has been set.
func (o *AuditLogConfig) GetUniverseLogsExporterConfigOk() (*[]UniverseLogsExporterConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseLogsExporterConfig, true
}

// SetUniverseLogsExporterConfig sets field value
func (o *AuditLogConfig) SetUniverseLogsExporterConfig(v []UniverseLogsExporterConfig) {
	o.UniverseLogsExporterConfig = v
}

// GetYcqlAuditConfig returns the YcqlAuditConfig field value if set, zero value otherwise.
func (o *AuditLogConfig) GetYcqlAuditConfig() YCQLAuditConfig {
	if o == nil || o.YcqlAuditConfig == nil {
		var ret YCQLAuditConfig
		return ret
	}
	return *o.YcqlAuditConfig
}

// GetYcqlAuditConfigOk returns a tuple with the YcqlAuditConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogConfig) GetYcqlAuditConfigOk() (*YCQLAuditConfig, bool) {
	if o == nil || o.YcqlAuditConfig == nil {
		return nil, false
	}
	return o.YcqlAuditConfig, true
}

// HasYcqlAuditConfig returns a boolean if a field has been set.
func (o *AuditLogConfig) HasYcqlAuditConfig() bool {
	if o != nil && o.YcqlAuditConfig != nil {
		return true
	}

	return false
}

// SetYcqlAuditConfig gets a reference to the given YCQLAuditConfig and assigns it to the YcqlAuditConfig field.
func (o *AuditLogConfig) SetYcqlAuditConfig(v YCQLAuditConfig) {
	o.YcqlAuditConfig = &v
}

// GetYsqlAuditConfig returns the YsqlAuditConfig field value if set, zero value otherwise.
func (o *AuditLogConfig) GetYsqlAuditConfig() YSQLAuditConfig {
	if o == nil || o.YsqlAuditConfig == nil {
		var ret YSQLAuditConfig
		return ret
	}
	return *o.YsqlAuditConfig
}

// GetYsqlAuditConfigOk returns a tuple with the YsqlAuditConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogConfig) GetYsqlAuditConfigOk() (*YSQLAuditConfig, bool) {
	if o == nil || o.YsqlAuditConfig == nil {
		return nil, false
	}
	return o.YsqlAuditConfig, true
}

// HasYsqlAuditConfig returns a boolean if a field has been set.
func (o *AuditLogConfig) HasYsqlAuditConfig() bool {
	if o != nil && o.YsqlAuditConfig != nil {
		return true
	}

	return false
}

// SetYsqlAuditConfig gets a reference to the given YSQLAuditConfig and assigns it to the YsqlAuditConfig field.
func (o *AuditLogConfig) SetYsqlAuditConfig(v YSQLAuditConfig) {
	o.YsqlAuditConfig = &v
}

func (o AuditLogConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExportActive != nil {
		toSerialize["exportActive"] = o.ExportActive
	}
	if true {
		toSerialize["universeLogsExporterConfig"] = o.UniverseLogsExporterConfig
	}
	if o.YcqlAuditConfig != nil {
		toSerialize["ycqlAuditConfig"] = o.YcqlAuditConfig
	}
	if o.YsqlAuditConfig != nil {
		toSerialize["ysqlAuditConfig"] = o.YsqlAuditConfig
	}
	return json.Marshal(toSerialize)
}

type NullableAuditLogConfig struct {
	value *AuditLogConfig
	isSet bool
}

func (v NullableAuditLogConfig) Get() *AuditLogConfig {
	return v.value
}

func (v *NullableAuditLogConfig) Set(val *AuditLogConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogConfig(val *AuditLogConfig) *NullableAuditLogConfig {
	return &NullableAuditLogConfig{value: val, isSet: true}
}

func (v NullableAuditLogConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


