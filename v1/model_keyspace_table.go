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

// KeyspaceTable Keyspace and table info for backup
type KeyspaceTable struct {
	// keyspace
	Keyspace *string `json:"keyspace,omitempty"`
	// Tables
	TableNameList *[]string `json:"tableNameList,omitempty"`
	// Table UUIDs
	TableUUIDList *[]string `json:"tableUUIDList,omitempty"`
}

// NewKeyspaceTable instantiates a new KeyspaceTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyspaceTable() *KeyspaceTable {
	this := KeyspaceTable{}
	return &this
}

// NewKeyspaceTableWithDefaults instantiates a new KeyspaceTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyspaceTableWithDefaults() *KeyspaceTable {
	this := KeyspaceTable{}
	return &this
}

// GetKeyspace returns the Keyspace field value if set, zero value otherwise.
func (o *KeyspaceTable) GetKeyspace() string {
	if o == nil || o.Keyspace == nil {
		var ret string
		return ret
	}
	return *o.Keyspace
}

// GetKeyspaceOk returns a tuple with the Keyspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyspaceTable) GetKeyspaceOk() (*string, bool) {
	if o == nil || o.Keyspace == nil {
		return nil, false
	}
	return o.Keyspace, true
}

// HasKeyspace returns a boolean if a field has been set.
func (o *KeyspaceTable) HasKeyspace() bool {
	if o != nil && o.Keyspace != nil {
		return true
	}

	return false
}

// SetKeyspace gets a reference to the given string and assigns it to the Keyspace field.
func (o *KeyspaceTable) SetKeyspace(v string) {
	o.Keyspace = &v
}

// GetTableNameList returns the TableNameList field value if set, zero value otherwise.
func (o *KeyspaceTable) GetTableNameList() []string {
	if o == nil || o.TableNameList == nil {
		var ret []string
		return ret
	}
	return *o.TableNameList
}

// GetTableNameListOk returns a tuple with the TableNameList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyspaceTable) GetTableNameListOk() (*[]string, bool) {
	if o == nil || o.TableNameList == nil {
		return nil, false
	}
	return o.TableNameList, true
}

// HasTableNameList returns a boolean if a field has been set.
func (o *KeyspaceTable) HasTableNameList() bool {
	if o != nil && o.TableNameList != nil {
		return true
	}

	return false
}

// SetTableNameList gets a reference to the given []string and assigns it to the TableNameList field.
func (o *KeyspaceTable) SetTableNameList(v []string) {
	o.TableNameList = &v
}

// GetTableUUIDList returns the TableUUIDList field value if set, zero value otherwise.
func (o *KeyspaceTable) GetTableUUIDList() []string {
	if o == nil || o.TableUUIDList == nil {
		var ret []string
		return ret
	}
	return *o.TableUUIDList
}

// GetTableUUIDListOk returns a tuple with the TableUUIDList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyspaceTable) GetTableUUIDListOk() (*[]string, bool) {
	if o == nil || o.TableUUIDList == nil {
		return nil, false
	}
	return o.TableUUIDList, true
}

// HasTableUUIDList returns a boolean if a field has been set.
func (o *KeyspaceTable) HasTableUUIDList() bool {
	if o != nil && o.TableUUIDList != nil {
		return true
	}

	return false
}

// SetTableUUIDList gets a reference to the given []string and assigns it to the TableUUIDList field.
func (o *KeyspaceTable) SetTableUUIDList(v []string) {
	o.TableUUIDList = &v
}

func (o KeyspaceTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keyspace != nil {
		toSerialize["keyspace"] = o.Keyspace
	}
	if o.TableNameList != nil {
		toSerialize["tableNameList"] = o.TableNameList
	}
	if o.TableUUIDList != nil {
		toSerialize["tableUUIDList"] = o.TableUUIDList
	}
	return json.Marshal(toSerialize)
}

type NullableKeyspaceTable struct {
	value *KeyspaceTable
	isSet bool
}

func (v NullableKeyspaceTable) Get() *KeyspaceTable {
	return v.value
}

func (v *NullableKeyspaceTable) Set(val *KeyspaceTable) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyspaceTable) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyspaceTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyspaceTable(val *KeyspaceTable) *NullableKeyspaceTable {
	return &NullableKeyspaceTable{value: val, isSet: true}
}

func (v NullableKeyspaceTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyspaceTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


