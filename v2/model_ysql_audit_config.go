/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// YSQLAuditConfig YSQL Audit Logging Configuration. Part of AuditLogConfig.
type YSQLAuditConfig struct {
	// YSQL statement classes
	Classes []string `json:"classes"`
	// Enabled
	Enabled bool `json:"enabled"`
	// Log catalog
	LogCatalog bool `json:"log_catalog"`
	// Log client
	LogClient bool `json:"log_client"`
	// Log level
	LogLevel string `json:"log_level"`
	// Log parameter
	LogParameter bool `json:"log_parameter"`
	// Log parameter max size
	LogParameterMaxSize int32 `json:"log_parameter_max_size"`
	// Log relation
	LogRelation bool `json:"log_relation"`
	// Log row
	LogRows bool `json:"log_rows"`
	// Log statement
	LogStatement bool `json:"log_statement"`
	// Log statement once
	LogStatementOnce bool `json:"log_statement_once"`
}

// NewYSQLAuditConfig instantiates a new YSQLAuditConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYSQLAuditConfig(classes []string, enabled bool, logCatalog bool, logClient bool, logLevel string, logParameter bool, logParameterMaxSize int32, logRelation bool, logRows bool, logStatement bool, logStatementOnce bool) *YSQLAuditConfig {
	this := YSQLAuditConfig{}
	this.Classes = classes
	this.Enabled = enabled
	this.LogCatalog = logCatalog
	this.LogClient = logClient
	this.LogLevel = logLevel
	this.LogParameter = logParameter
	this.LogParameterMaxSize = logParameterMaxSize
	this.LogRelation = logRelation
	this.LogRows = logRows
	this.LogStatement = logStatement
	this.LogStatementOnce = logStatementOnce
	return &this
}

// NewYSQLAuditConfigWithDefaults instantiates a new YSQLAuditConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYSQLAuditConfigWithDefaults() *YSQLAuditConfig {
	this := YSQLAuditConfig{}
	return &this
}

// GetClasses returns the Classes field value
func (o *YSQLAuditConfig) GetClasses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Classes
}

// GetClassesOk returns a tuple with the Classes field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetClassesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Classes, true
}

// SetClasses sets field value
func (o *YSQLAuditConfig) SetClasses(v []string) {
	o.Classes = v
}

// GetEnabled returns the Enabled field value
func (o *YSQLAuditConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *YSQLAuditConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLogCatalog returns the LogCatalog field value
func (o *YSQLAuditConfig) GetLogCatalog() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LogCatalog
}

// GetLogCatalogOk returns a tuple with the LogCatalog field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogCatalogOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogCatalog, true
}

// SetLogCatalog sets field value
func (o *YSQLAuditConfig) SetLogCatalog(v bool) {
	o.LogCatalog = v
}

// GetLogClient returns the LogClient field value
func (o *YSQLAuditConfig) GetLogClient() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LogClient
}

// GetLogClientOk returns a tuple with the LogClient field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogClientOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogClient, true
}

// SetLogClient sets field value
func (o *YSQLAuditConfig) SetLogClient(v bool) {
	o.LogClient = v
}

// GetLogLevel returns the LogLevel field value
func (o *YSQLAuditConfig) GetLogLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogLevelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogLevel, true
}

// SetLogLevel sets field value
func (o *YSQLAuditConfig) SetLogLevel(v string) {
	o.LogLevel = v
}

// GetLogParameter returns the LogParameter field value
func (o *YSQLAuditConfig) GetLogParameter() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LogParameter
}

// GetLogParameterOk returns a tuple with the LogParameter field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogParameterOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogParameter, true
}

// SetLogParameter sets field value
func (o *YSQLAuditConfig) SetLogParameter(v bool) {
	o.LogParameter = v
}

// GetLogParameterMaxSize returns the LogParameterMaxSize field value
func (o *YSQLAuditConfig) GetLogParameterMaxSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LogParameterMaxSize
}

// GetLogParameterMaxSizeOk returns a tuple with the LogParameterMaxSize field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogParameterMaxSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogParameterMaxSize, true
}

// SetLogParameterMaxSize sets field value
func (o *YSQLAuditConfig) SetLogParameterMaxSize(v int32) {
	o.LogParameterMaxSize = v
}

// GetLogRelation returns the LogRelation field value
func (o *YSQLAuditConfig) GetLogRelation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LogRelation
}

// GetLogRelationOk returns a tuple with the LogRelation field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogRelationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogRelation, true
}

// SetLogRelation sets field value
func (o *YSQLAuditConfig) SetLogRelation(v bool) {
	o.LogRelation = v
}

// GetLogRows returns the LogRows field value
func (o *YSQLAuditConfig) GetLogRows() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LogRows
}

// GetLogRowsOk returns a tuple with the LogRows field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogRowsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogRows, true
}

// SetLogRows sets field value
func (o *YSQLAuditConfig) SetLogRows(v bool) {
	o.LogRows = v
}

// GetLogStatement returns the LogStatement field value
func (o *YSQLAuditConfig) GetLogStatement() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LogStatement
}

// GetLogStatementOk returns a tuple with the LogStatement field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogStatementOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogStatement, true
}

// SetLogStatement sets field value
func (o *YSQLAuditConfig) SetLogStatement(v bool) {
	o.LogStatement = v
}

// GetLogStatementOnce returns the LogStatementOnce field value
func (o *YSQLAuditConfig) GetLogStatementOnce() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LogStatementOnce
}

// GetLogStatementOnceOk returns a tuple with the LogStatementOnce field value
// and a boolean to check if the value has been set.
func (o *YSQLAuditConfig) GetLogStatementOnceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogStatementOnce, true
}

// SetLogStatementOnce sets field value
func (o *YSQLAuditConfig) SetLogStatementOnce(v bool) {
	o.LogStatementOnce = v
}

func (o YSQLAuditConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["classes"] = o.Classes
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["log_catalog"] = o.LogCatalog
	}
	if true {
		toSerialize["log_client"] = o.LogClient
	}
	if true {
		toSerialize["log_level"] = o.LogLevel
	}
	if true {
		toSerialize["log_parameter"] = o.LogParameter
	}
	if true {
		toSerialize["log_parameter_max_size"] = o.LogParameterMaxSize
	}
	if true {
		toSerialize["log_relation"] = o.LogRelation
	}
	if true {
		toSerialize["log_rows"] = o.LogRows
	}
	if true {
		toSerialize["log_statement"] = o.LogStatement
	}
	if true {
		toSerialize["log_statement_once"] = o.LogStatementOnce
	}
	return json.Marshal(toSerialize)
}

type NullableYSQLAuditConfig struct {
	value *YSQLAuditConfig
	isSet bool
}

func (v NullableYSQLAuditConfig) Get() *YSQLAuditConfig {
	return v.value
}

func (v *NullableYSQLAuditConfig) Set(val *YSQLAuditConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableYSQLAuditConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableYSQLAuditConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYSQLAuditConfig(val *YSQLAuditConfig) *NullableYSQLAuditConfig {
	return &NullableYSQLAuditConfig{value: val, isSet: true}
}

func (v NullableYSQLAuditConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYSQLAuditConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

