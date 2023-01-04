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

// TableDetails Table details
type TableDetails struct {
	// Details of all columns in the table
	Columns *[]ColumnDetails `json:"columns,omitempty"`
	// Keyspace to which this table belongs
	Keyspace *string `json:"keyspace,omitempty"`
	// Primary key values used to split table into tablets
	SplitValues *[]string `json:"splitValues,omitempty"`
	// Table name
	TableName *string `json:"tableName,omitempty"`
	// The default table-level time to live, in seconds. A value of `-1` represents an infinite TTL.
	TtlInSeconds *int64 `json:"ttlInSeconds,omitempty"`
}

// NewTableDetails instantiates a new TableDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableDetails() *TableDetails {
	this := TableDetails{}
	return &this
}

// NewTableDetailsWithDefaults instantiates a new TableDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableDetailsWithDefaults() *TableDetails {
	this := TableDetails{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TableDetails) GetColumns() []ColumnDetails {
	if o == nil || o.Columns == nil {
		var ret []ColumnDetails
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDetails) GetColumnsOk() (*[]ColumnDetails, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TableDetails) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ColumnDetails and assigns it to the Columns field.
func (o *TableDetails) SetColumns(v []ColumnDetails) {
	o.Columns = &v
}

// GetKeyspace returns the Keyspace field value if set, zero value otherwise.
func (o *TableDetails) GetKeyspace() string {
	if o == nil || o.Keyspace == nil {
		var ret string
		return ret
	}
	return *o.Keyspace
}

// GetKeyspaceOk returns a tuple with the Keyspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDetails) GetKeyspaceOk() (*string, bool) {
	if o == nil || o.Keyspace == nil {
		return nil, false
	}
	return o.Keyspace, true
}

// HasKeyspace returns a boolean if a field has been set.
func (o *TableDetails) HasKeyspace() bool {
	if o != nil && o.Keyspace != nil {
		return true
	}

	return false
}

// SetKeyspace gets a reference to the given string and assigns it to the Keyspace field.
func (o *TableDetails) SetKeyspace(v string) {
	o.Keyspace = &v
}

// GetSplitValues returns the SplitValues field value if set, zero value otherwise.
func (o *TableDetails) GetSplitValues() []string {
	if o == nil || o.SplitValues == nil {
		var ret []string
		return ret
	}
	return *o.SplitValues
}

// GetSplitValuesOk returns a tuple with the SplitValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDetails) GetSplitValuesOk() (*[]string, bool) {
	if o == nil || o.SplitValues == nil {
		return nil, false
	}
	return o.SplitValues, true
}

// HasSplitValues returns a boolean if a field has been set.
func (o *TableDetails) HasSplitValues() bool {
	if o != nil && o.SplitValues != nil {
		return true
	}

	return false
}

// SetSplitValues gets a reference to the given []string and assigns it to the SplitValues field.
func (o *TableDetails) SetSplitValues(v []string) {
	o.SplitValues = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *TableDetails) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDetails) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *TableDetails) HasTableName() bool {
	if o != nil && o.TableName != nil {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *TableDetails) SetTableName(v string) {
	o.TableName = &v
}

// GetTtlInSeconds returns the TtlInSeconds field value if set, zero value otherwise.
func (o *TableDetails) GetTtlInSeconds() int64 {
	if o == nil || o.TtlInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.TtlInSeconds
}

// GetTtlInSecondsOk returns a tuple with the TtlInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDetails) GetTtlInSecondsOk() (*int64, bool) {
	if o == nil || o.TtlInSeconds == nil {
		return nil, false
	}
	return o.TtlInSeconds, true
}

// HasTtlInSeconds returns a boolean if a field has been set.
func (o *TableDetails) HasTtlInSeconds() bool {
	if o != nil && o.TtlInSeconds != nil {
		return true
	}

	return false
}

// SetTtlInSeconds gets a reference to the given int64 and assigns it to the TtlInSeconds field.
func (o *TableDetails) SetTtlInSeconds(v int64) {
	o.TtlInSeconds = &v
}

func (o TableDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Keyspace != nil {
		toSerialize["keyspace"] = o.Keyspace
	}
	if o.SplitValues != nil {
		toSerialize["splitValues"] = o.SplitValues
	}
	if o.TableName != nil {
		toSerialize["tableName"] = o.TableName
	}
	if o.TtlInSeconds != nil {
		toSerialize["ttlInSeconds"] = o.TtlInSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableTableDetails struct {
	value *TableDetails
	isSet bool
}

func (v NullableTableDetails) Get() *TableDetails {
	return v.value
}

func (v *NullableTableDetails) Set(val *TableDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTableDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTableDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableDetails(val *TableDetails) *NullableTableDetails {
	return &NullableTableDetails{value: val, isSet: true}
}

func (v NullableTableDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


