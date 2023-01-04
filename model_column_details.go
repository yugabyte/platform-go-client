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

// ColumnDetails Details of a CQL database column
type ColumnDetails struct {
	// Relative position (column order) for this column, in the table and in CQL commands
	ColumnOrder *int32 `json:"columnOrder,omitempty"`
	// True if this column is a clustering key
	IsClusteringKey *bool `json:"isClusteringKey,omitempty"`
	// True if this column is a partition key
	IsPartitionKey *bool `json:"isPartitionKey,omitempty"`
	// Column key type
	KeyType *string `json:"keyType,omitempty"`
	// Column name
	Name *string `json:"name,omitempty"`
	// Sort order for this column. Valid only for clustering columns.
	SortOrder *string `json:"sortOrder,omitempty"`
	// The column's data type
	Type *string `json:"type,omitempty"`
	// Column value name
	ValueType *string `json:"valueType,omitempty"`
}

// NewColumnDetails instantiates a new ColumnDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumnDetails() *ColumnDetails {
	this := ColumnDetails{}
	return &this
}

// NewColumnDetailsWithDefaults instantiates a new ColumnDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnDetailsWithDefaults() *ColumnDetails {
	this := ColumnDetails{}
	return &this
}

// GetColumnOrder returns the ColumnOrder field value if set, zero value otherwise.
func (o *ColumnDetails) GetColumnOrder() int32 {
	if o == nil || o.ColumnOrder == nil {
		var ret int32
		return ret
	}
	return *o.ColumnOrder
}

// GetColumnOrderOk returns a tuple with the ColumnOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnDetails) GetColumnOrderOk() (*int32, bool) {
	if o == nil || o.ColumnOrder == nil {
		return nil, false
	}
	return o.ColumnOrder, true
}

// HasColumnOrder returns a boolean if a field has been set.
func (o *ColumnDetails) HasColumnOrder() bool {
	if o != nil && o.ColumnOrder != nil {
		return true
	}

	return false
}

// SetColumnOrder gets a reference to the given int32 and assigns it to the ColumnOrder field.
func (o *ColumnDetails) SetColumnOrder(v int32) {
	o.ColumnOrder = &v
}

// GetIsClusteringKey returns the IsClusteringKey field value if set, zero value otherwise.
func (o *ColumnDetails) GetIsClusteringKey() bool {
	if o == nil || o.IsClusteringKey == nil {
		var ret bool
		return ret
	}
	return *o.IsClusteringKey
}

// GetIsClusteringKeyOk returns a tuple with the IsClusteringKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnDetails) GetIsClusteringKeyOk() (*bool, bool) {
	if o == nil || o.IsClusteringKey == nil {
		return nil, false
	}
	return o.IsClusteringKey, true
}

// HasIsClusteringKey returns a boolean if a field has been set.
func (o *ColumnDetails) HasIsClusteringKey() bool {
	if o != nil && o.IsClusteringKey != nil {
		return true
	}

	return false
}

// SetIsClusteringKey gets a reference to the given bool and assigns it to the IsClusteringKey field.
func (o *ColumnDetails) SetIsClusteringKey(v bool) {
	o.IsClusteringKey = &v
}

// GetIsPartitionKey returns the IsPartitionKey field value if set, zero value otherwise.
func (o *ColumnDetails) GetIsPartitionKey() bool {
	if o == nil || o.IsPartitionKey == nil {
		var ret bool
		return ret
	}
	return *o.IsPartitionKey
}

// GetIsPartitionKeyOk returns a tuple with the IsPartitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnDetails) GetIsPartitionKeyOk() (*bool, bool) {
	if o == nil || o.IsPartitionKey == nil {
		return nil, false
	}
	return o.IsPartitionKey, true
}

// HasIsPartitionKey returns a boolean if a field has been set.
func (o *ColumnDetails) HasIsPartitionKey() bool {
	if o != nil && o.IsPartitionKey != nil {
		return true
	}

	return false
}

// SetIsPartitionKey gets a reference to the given bool and assigns it to the IsPartitionKey field.
func (o *ColumnDetails) SetIsPartitionKey(v bool) {
	o.IsPartitionKey = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *ColumnDetails) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnDetails) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *ColumnDetails) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *ColumnDetails) SetKeyType(v string) {
	o.KeyType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ColumnDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ColumnDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ColumnDetails) SetName(v string) {
	o.Name = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *ColumnDetails) GetSortOrder() string {
	if o == nil || o.SortOrder == nil {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnDetails) GetSortOrderOk() (*string, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *ColumnDetails) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *ColumnDetails) SetSortOrder(v string) {
	o.SortOrder = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ColumnDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnDetails) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ColumnDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ColumnDetails) SetType(v string) {
	o.Type = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ColumnDetails) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnDetails) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ColumnDetails) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *ColumnDetails) SetValueType(v string) {
	o.ValueType = &v
}

func (o ColumnDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnOrder != nil {
		toSerialize["columnOrder"] = o.ColumnOrder
	}
	if o.IsClusteringKey != nil {
		toSerialize["isClusteringKey"] = o.IsClusteringKey
	}
	if o.IsPartitionKey != nil {
		toSerialize["isPartitionKey"] = o.IsPartitionKey
	}
	if o.KeyType != nil {
		toSerialize["keyType"] = o.KeyType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

type NullableColumnDetails struct {
	value *ColumnDetails
	isSet bool
}

func (v NullableColumnDetails) Get() *ColumnDetails {
	return v.value
}

func (v *NullableColumnDetails) Set(val *ColumnDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableColumnDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableColumnDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumnDetails(val *ColumnDetails) *NullableColumnDetails {
	return &NullableColumnDetails{value: val, isSet: true}
}

func (v NullableColumnDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumnDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


