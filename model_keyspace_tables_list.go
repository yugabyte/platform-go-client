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

// KeyspaceTablesList struct for KeyspaceTablesList
type KeyspaceTablesList struct {
	AllTables bool `json:"allTables"`
	BackupSizeInBytes int64 `json:"backupSizeInBytes"`
	DefaultLocation string `json:"defaultLocation"`
	Keyspace string `json:"keyspace"`
	PerRegionLocations []RegionLocations `json:"perRegionLocations"`
	TableUUIDList []string `json:"tableUUIDList"`
	TablesList []string `json:"tablesList"`
}

// NewKeyspaceTablesList instantiates a new KeyspaceTablesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyspaceTablesList(allTables bool, backupSizeInBytes int64, defaultLocation string, keyspace string, perRegionLocations []RegionLocations, tableUUIDList []string, tablesList []string, ) *KeyspaceTablesList {
	this := KeyspaceTablesList{}
	this.AllTables = allTables
	this.BackupSizeInBytes = backupSizeInBytes
	this.DefaultLocation = defaultLocation
	this.Keyspace = keyspace
	this.PerRegionLocations = perRegionLocations
	this.TableUUIDList = tableUUIDList
	this.TablesList = tablesList
	return &this
}

// NewKeyspaceTablesListWithDefaults instantiates a new KeyspaceTablesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyspaceTablesListWithDefaults() *KeyspaceTablesList {
	this := KeyspaceTablesList{}
	return &this
}

// GetAllTables returns the AllTables field value
func (o *KeyspaceTablesList) GetAllTables() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AllTables
}

// GetAllTablesOk returns a tuple with the AllTables field value
// and a boolean to check if the value has been set.
func (o *KeyspaceTablesList) GetAllTablesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AllTables, true
}

// SetAllTables sets field value
func (o *KeyspaceTablesList) SetAllTables(v bool) {
	o.AllTables = v
}

// GetBackupSizeInBytes returns the BackupSizeInBytes field value
func (o *KeyspaceTablesList) GetBackupSizeInBytes() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.BackupSizeInBytes
}

// GetBackupSizeInBytesOk returns a tuple with the BackupSizeInBytes field value
// and a boolean to check if the value has been set.
func (o *KeyspaceTablesList) GetBackupSizeInBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupSizeInBytes, true
}

// SetBackupSizeInBytes sets field value
func (o *KeyspaceTablesList) SetBackupSizeInBytes(v int64) {
	o.BackupSizeInBytes = v
}

// GetDefaultLocation returns the DefaultLocation field value
func (o *KeyspaceTablesList) GetDefaultLocation() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DefaultLocation
}

// GetDefaultLocationOk returns a tuple with the DefaultLocation field value
// and a boolean to check if the value has been set.
func (o *KeyspaceTablesList) GetDefaultLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultLocation, true
}

// SetDefaultLocation sets field value
func (o *KeyspaceTablesList) SetDefaultLocation(v string) {
	o.DefaultLocation = v
}

// GetKeyspace returns the Keyspace field value
func (o *KeyspaceTablesList) GetKeyspace() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Keyspace
}

// GetKeyspaceOk returns a tuple with the Keyspace field value
// and a boolean to check if the value has been set.
func (o *KeyspaceTablesList) GetKeyspaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Keyspace, true
}

// SetKeyspace sets field value
func (o *KeyspaceTablesList) SetKeyspace(v string) {
	o.Keyspace = v
}

// GetPerRegionLocations returns the PerRegionLocations field value
func (o *KeyspaceTablesList) GetPerRegionLocations() []RegionLocations {
	if o == nil  {
		var ret []RegionLocations
		return ret
	}

	return o.PerRegionLocations
}

// GetPerRegionLocationsOk returns a tuple with the PerRegionLocations field value
// and a boolean to check if the value has been set.
func (o *KeyspaceTablesList) GetPerRegionLocationsOk() (*[]RegionLocations, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PerRegionLocations, true
}

// SetPerRegionLocations sets field value
func (o *KeyspaceTablesList) SetPerRegionLocations(v []RegionLocations) {
	o.PerRegionLocations = v
}

// GetTableUUIDList returns the TableUUIDList field value
func (o *KeyspaceTablesList) GetTableUUIDList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.TableUUIDList
}

// GetTableUUIDListOk returns a tuple with the TableUUIDList field value
// and a boolean to check if the value has been set.
func (o *KeyspaceTablesList) GetTableUUIDListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TableUUIDList, true
}

// SetTableUUIDList sets field value
func (o *KeyspaceTablesList) SetTableUUIDList(v []string) {
	o.TableUUIDList = v
}

// GetTablesList returns the TablesList field value
func (o *KeyspaceTablesList) GetTablesList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.TablesList
}

// GetTablesListOk returns a tuple with the TablesList field value
// and a boolean to check if the value has been set.
func (o *KeyspaceTablesList) GetTablesListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TablesList, true
}

// SetTablesList sets field value
func (o *KeyspaceTablesList) SetTablesList(v []string) {
	o.TablesList = v
}

func (o KeyspaceTablesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allTables"] = o.AllTables
	}
	if true {
		toSerialize["backupSizeInBytes"] = o.BackupSizeInBytes
	}
	if true {
		toSerialize["defaultLocation"] = o.DefaultLocation
	}
	if true {
		toSerialize["keyspace"] = o.Keyspace
	}
	if true {
		toSerialize["perRegionLocations"] = o.PerRegionLocations
	}
	if true {
		toSerialize["tableUUIDList"] = o.TableUUIDList
	}
	if true {
		toSerialize["tablesList"] = o.TablesList
	}
	return json.Marshal(toSerialize)
}

type NullableKeyspaceTablesList struct {
	value *KeyspaceTablesList
	isSet bool
}

func (v NullableKeyspaceTablesList) Get() *KeyspaceTablesList {
	return v.value
}

func (v *NullableKeyspaceTablesList) Set(val *KeyspaceTablesList) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyspaceTablesList) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyspaceTablesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyspaceTablesList(val *KeyspaceTablesList) *NullableKeyspaceTablesList {
	return &NullableKeyspaceTablesList{value: val, isSet: true}
}

func (v NullableKeyspaceTablesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyspaceTablesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


