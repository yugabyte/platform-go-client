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

// TableInfoResp Table information response
type TableInfoResp struct {
	// Flag, indicating colocated table
	Colocated *bool `json:"colocated,omitempty"`
	// Colocation parent id
	ColocationParentId *string `json:"colocationParentId,omitempty"`
	// Keyspace
	KeySpace *string `json:"keySpace,omitempty"`
	// Main Table UUID of index tables
	MainTableUUID *string `json:"mainTableUUID,omitempty"`
	// Namespace or Schema
	NameSpace *string `json:"nameSpace,omitempty"`
	// Parent Table UUID
	ParentTableUUID *string `json:"parentTableUUID,omitempty"`
	// Postgres schema name of the table
	PgSchemaName *string `json:"pgSchemaName,omitempty"`
	// Relation type
	RelationType *string `json:"relationType,omitempty"`
	// SST size in bytes
	SizeBytes *float64 `json:"sizeBytes,omitempty"`
	// Table ID
	TableID *string `json:"tableID,omitempty"`
	// Table name
	TableName *string `json:"tableName,omitempty"`
	// Table space
	TableSpace *string `json:"tableSpace,omitempty"`
	// Table type
	TableType *string `json:"tableType,omitempty"`
	// Table UUID
	TableUUID *string `json:"tableUUID,omitempty"`
	// WAL size in bytes
	WalSizeBytes *float64 `json:"walSizeBytes,omitempty"`
}

// NewTableInfoResp instantiates a new TableInfoResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableInfoResp() *TableInfoResp {
	this := TableInfoResp{}
	return &this
}

// NewTableInfoRespWithDefaults instantiates a new TableInfoResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableInfoRespWithDefaults() *TableInfoResp {
	this := TableInfoResp{}
	return &this
}

// GetColocated returns the Colocated field value if set, zero value otherwise.
func (o *TableInfoResp) GetColocated() bool {
	if o == nil || o.Colocated == nil {
		var ret bool
		return ret
	}
	return *o.Colocated
}

// GetColocatedOk returns a tuple with the Colocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetColocatedOk() (*bool, bool) {
	if o == nil || o.Colocated == nil {
		return nil, false
	}
	return o.Colocated, true
}

// HasColocated returns a boolean if a field has been set.
func (o *TableInfoResp) HasColocated() bool {
	if o != nil && o.Colocated != nil {
		return true
	}

	return false
}

// SetColocated gets a reference to the given bool and assigns it to the Colocated field.
func (o *TableInfoResp) SetColocated(v bool) {
	o.Colocated = &v
}

// GetColocationParentId returns the ColocationParentId field value if set, zero value otherwise.
func (o *TableInfoResp) GetColocationParentId() string {
	if o == nil || o.ColocationParentId == nil {
		var ret string
		return ret
	}
	return *o.ColocationParentId
}

// GetColocationParentIdOk returns a tuple with the ColocationParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetColocationParentIdOk() (*string, bool) {
	if o == nil || o.ColocationParentId == nil {
		return nil, false
	}
	return o.ColocationParentId, true
}

// HasColocationParentId returns a boolean if a field has been set.
func (o *TableInfoResp) HasColocationParentId() bool {
	if o != nil && o.ColocationParentId != nil {
		return true
	}

	return false
}

// SetColocationParentId gets a reference to the given string and assigns it to the ColocationParentId field.
func (o *TableInfoResp) SetColocationParentId(v string) {
	o.ColocationParentId = &v
}

// GetKeySpace returns the KeySpace field value if set, zero value otherwise.
func (o *TableInfoResp) GetKeySpace() string {
	if o == nil || o.KeySpace == nil {
		var ret string
		return ret
	}
	return *o.KeySpace
}

// GetKeySpaceOk returns a tuple with the KeySpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetKeySpaceOk() (*string, bool) {
	if o == nil || o.KeySpace == nil {
		return nil, false
	}
	return o.KeySpace, true
}

// HasKeySpace returns a boolean if a field has been set.
func (o *TableInfoResp) HasKeySpace() bool {
	if o != nil && o.KeySpace != nil {
		return true
	}

	return false
}

// SetKeySpace gets a reference to the given string and assigns it to the KeySpace field.
func (o *TableInfoResp) SetKeySpace(v string) {
	o.KeySpace = &v
}

// GetMainTableUUID returns the MainTableUUID field value if set, zero value otherwise.
func (o *TableInfoResp) GetMainTableUUID() string {
	if o == nil || o.MainTableUUID == nil {
		var ret string
		return ret
	}
	return *o.MainTableUUID
}

// GetMainTableUUIDOk returns a tuple with the MainTableUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetMainTableUUIDOk() (*string, bool) {
	if o == nil || o.MainTableUUID == nil {
		return nil, false
	}
	return o.MainTableUUID, true
}

// HasMainTableUUID returns a boolean if a field has been set.
func (o *TableInfoResp) HasMainTableUUID() bool {
	if o != nil && o.MainTableUUID != nil {
		return true
	}

	return false
}

// SetMainTableUUID gets a reference to the given string and assigns it to the MainTableUUID field.
func (o *TableInfoResp) SetMainTableUUID(v string) {
	o.MainTableUUID = &v
}

// GetNameSpace returns the NameSpace field value if set, zero value otherwise.
func (o *TableInfoResp) GetNameSpace() string {
	if o == nil || o.NameSpace == nil {
		var ret string
		return ret
	}
	return *o.NameSpace
}

// GetNameSpaceOk returns a tuple with the NameSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetNameSpaceOk() (*string, bool) {
	if o == nil || o.NameSpace == nil {
		return nil, false
	}
	return o.NameSpace, true
}

// HasNameSpace returns a boolean if a field has been set.
func (o *TableInfoResp) HasNameSpace() bool {
	if o != nil && o.NameSpace != nil {
		return true
	}

	return false
}

// SetNameSpace gets a reference to the given string and assigns it to the NameSpace field.
func (o *TableInfoResp) SetNameSpace(v string) {
	o.NameSpace = &v
}

// GetParentTableUUID returns the ParentTableUUID field value if set, zero value otherwise.
func (o *TableInfoResp) GetParentTableUUID() string {
	if o == nil || o.ParentTableUUID == nil {
		var ret string
		return ret
	}
	return *o.ParentTableUUID
}

// GetParentTableUUIDOk returns a tuple with the ParentTableUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetParentTableUUIDOk() (*string, bool) {
	if o == nil || o.ParentTableUUID == nil {
		return nil, false
	}
	return o.ParentTableUUID, true
}

// HasParentTableUUID returns a boolean if a field has been set.
func (o *TableInfoResp) HasParentTableUUID() bool {
	if o != nil && o.ParentTableUUID != nil {
		return true
	}

	return false
}

// SetParentTableUUID gets a reference to the given string and assigns it to the ParentTableUUID field.
func (o *TableInfoResp) SetParentTableUUID(v string) {
	o.ParentTableUUID = &v
}

// GetPgSchemaName returns the PgSchemaName field value if set, zero value otherwise.
func (o *TableInfoResp) GetPgSchemaName() string {
	if o == nil || o.PgSchemaName == nil {
		var ret string
		return ret
	}
	return *o.PgSchemaName
}

// GetPgSchemaNameOk returns a tuple with the PgSchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetPgSchemaNameOk() (*string, bool) {
	if o == nil || o.PgSchemaName == nil {
		return nil, false
	}
	return o.PgSchemaName, true
}

// HasPgSchemaName returns a boolean if a field has been set.
func (o *TableInfoResp) HasPgSchemaName() bool {
	if o != nil && o.PgSchemaName != nil {
		return true
	}

	return false
}

// SetPgSchemaName gets a reference to the given string and assigns it to the PgSchemaName field.
func (o *TableInfoResp) SetPgSchemaName(v string) {
	o.PgSchemaName = &v
}

// GetRelationType returns the RelationType field value if set, zero value otherwise.
func (o *TableInfoResp) GetRelationType() string {
	if o == nil || o.RelationType == nil {
		var ret string
		return ret
	}
	return *o.RelationType
}

// GetRelationTypeOk returns a tuple with the RelationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetRelationTypeOk() (*string, bool) {
	if o == nil || o.RelationType == nil {
		return nil, false
	}
	return o.RelationType, true
}

// HasRelationType returns a boolean if a field has been set.
func (o *TableInfoResp) HasRelationType() bool {
	if o != nil && o.RelationType != nil {
		return true
	}

	return false
}

// SetRelationType gets a reference to the given string and assigns it to the RelationType field.
func (o *TableInfoResp) SetRelationType(v string) {
	o.RelationType = &v
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise.
func (o *TableInfoResp) GetSizeBytes() float64 {
	if o == nil || o.SizeBytes == nil {
		var ret float64
		return ret
	}
	return *o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetSizeBytesOk() (*float64, bool) {
	if o == nil || o.SizeBytes == nil {
		return nil, false
	}
	return o.SizeBytes, true
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *TableInfoResp) HasSizeBytes() bool {
	if o != nil && o.SizeBytes != nil {
		return true
	}

	return false
}

// SetSizeBytes gets a reference to the given float64 and assigns it to the SizeBytes field.
func (o *TableInfoResp) SetSizeBytes(v float64) {
	o.SizeBytes = &v
}

// GetTableID returns the TableID field value if set, zero value otherwise.
func (o *TableInfoResp) GetTableID() string {
	if o == nil || o.TableID == nil {
		var ret string
		return ret
	}
	return *o.TableID
}

// GetTableIDOk returns a tuple with the TableID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetTableIDOk() (*string, bool) {
	if o == nil || o.TableID == nil {
		return nil, false
	}
	return o.TableID, true
}

// HasTableID returns a boolean if a field has been set.
func (o *TableInfoResp) HasTableID() bool {
	if o != nil && o.TableID != nil {
		return true
	}

	return false
}

// SetTableID gets a reference to the given string and assigns it to the TableID field.
func (o *TableInfoResp) SetTableID(v string) {
	o.TableID = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *TableInfoResp) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *TableInfoResp) HasTableName() bool {
	if o != nil && o.TableName != nil {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *TableInfoResp) SetTableName(v string) {
	o.TableName = &v
}

// GetTableSpace returns the TableSpace field value if set, zero value otherwise.
func (o *TableInfoResp) GetTableSpace() string {
	if o == nil || o.TableSpace == nil {
		var ret string
		return ret
	}
	return *o.TableSpace
}

// GetTableSpaceOk returns a tuple with the TableSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetTableSpaceOk() (*string, bool) {
	if o == nil || o.TableSpace == nil {
		return nil, false
	}
	return o.TableSpace, true
}

// HasTableSpace returns a boolean if a field has been set.
func (o *TableInfoResp) HasTableSpace() bool {
	if o != nil && o.TableSpace != nil {
		return true
	}

	return false
}

// SetTableSpace gets a reference to the given string and assigns it to the TableSpace field.
func (o *TableInfoResp) SetTableSpace(v string) {
	o.TableSpace = &v
}

// GetTableType returns the TableType field value if set, zero value otherwise.
func (o *TableInfoResp) GetTableType() string {
	if o == nil || o.TableType == nil {
		var ret string
		return ret
	}
	return *o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetTableTypeOk() (*string, bool) {
	if o == nil || o.TableType == nil {
		return nil, false
	}
	return o.TableType, true
}

// HasTableType returns a boolean if a field has been set.
func (o *TableInfoResp) HasTableType() bool {
	if o != nil && o.TableType != nil {
		return true
	}

	return false
}

// SetTableType gets a reference to the given string and assigns it to the TableType field.
func (o *TableInfoResp) SetTableType(v string) {
	o.TableType = &v
}

// GetTableUUID returns the TableUUID field value if set, zero value otherwise.
func (o *TableInfoResp) GetTableUUID() string {
	if o == nil || o.TableUUID == nil {
		var ret string
		return ret
	}
	return *o.TableUUID
}

// GetTableUUIDOk returns a tuple with the TableUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetTableUUIDOk() (*string, bool) {
	if o == nil || o.TableUUID == nil {
		return nil, false
	}
	return o.TableUUID, true
}

// HasTableUUID returns a boolean if a field has been set.
func (o *TableInfoResp) HasTableUUID() bool {
	if o != nil && o.TableUUID != nil {
		return true
	}

	return false
}

// SetTableUUID gets a reference to the given string and assigns it to the TableUUID field.
func (o *TableInfoResp) SetTableUUID(v string) {
	o.TableUUID = &v
}

// GetWalSizeBytes returns the WalSizeBytes field value if set, zero value otherwise.
func (o *TableInfoResp) GetWalSizeBytes() float64 {
	if o == nil || o.WalSizeBytes == nil {
		var ret float64
		return ret
	}
	return *o.WalSizeBytes
}

// GetWalSizeBytesOk returns a tuple with the WalSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableInfoResp) GetWalSizeBytesOk() (*float64, bool) {
	if o == nil || o.WalSizeBytes == nil {
		return nil, false
	}
	return o.WalSizeBytes, true
}

// HasWalSizeBytes returns a boolean if a field has been set.
func (o *TableInfoResp) HasWalSizeBytes() bool {
	if o != nil && o.WalSizeBytes != nil {
		return true
	}

	return false
}

// SetWalSizeBytes gets a reference to the given float64 and assigns it to the WalSizeBytes field.
func (o *TableInfoResp) SetWalSizeBytes(v float64) {
	o.WalSizeBytes = &v
}

func (o TableInfoResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Colocated != nil {
		toSerialize["colocated"] = o.Colocated
	}
	if o.ColocationParentId != nil {
		toSerialize["colocationParentId"] = o.ColocationParentId
	}
	if o.KeySpace != nil {
		toSerialize["keySpace"] = o.KeySpace
	}
	if o.MainTableUUID != nil {
		toSerialize["mainTableUUID"] = o.MainTableUUID
	}
	if o.NameSpace != nil {
		toSerialize["nameSpace"] = o.NameSpace
	}
	if o.ParentTableUUID != nil {
		toSerialize["parentTableUUID"] = o.ParentTableUUID
	}
	if o.PgSchemaName != nil {
		toSerialize["pgSchemaName"] = o.PgSchemaName
	}
	if o.RelationType != nil {
		toSerialize["relationType"] = o.RelationType
	}
	if o.SizeBytes != nil {
		toSerialize["sizeBytes"] = o.SizeBytes
	}
	if o.TableID != nil {
		toSerialize["tableID"] = o.TableID
	}
	if o.TableName != nil {
		toSerialize["tableName"] = o.TableName
	}
	if o.TableSpace != nil {
		toSerialize["tableSpace"] = o.TableSpace
	}
	if o.TableType != nil {
		toSerialize["tableType"] = o.TableType
	}
	if o.TableUUID != nil {
		toSerialize["tableUUID"] = o.TableUUID
	}
	if o.WalSizeBytes != nil {
		toSerialize["walSizeBytes"] = o.WalSizeBytes
	}
	return json.Marshal(toSerialize)
}

type NullableTableInfoResp struct {
	value *TableInfoResp
	isSet bool
}

func (v NullableTableInfoResp) Get() *TableInfoResp {
	return v.value
}

func (v *NullableTableInfoResp) Set(val *TableInfoResp) {
	v.value = val
	v.isSet = true
}

func (v NullableTableInfoResp) IsSet() bool {
	return v.isSet
}

func (v *NullableTableInfoResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableInfoResp(val *TableInfoResp) *NullableTableInfoResp {
	return &NullableTableInfoResp{value: val, isSet: true}
}

func (v NullableTableInfoResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableInfoResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


