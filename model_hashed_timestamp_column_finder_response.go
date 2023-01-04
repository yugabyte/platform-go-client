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

// HashedTimestampColumnFinderResponse struct for HashedTimestampColumnFinderResponse
type HashedTimestampColumnFinderResponse struct {
	CurrentDatabase *string `json:"current_database,omitempty"`
	Description *string `json:"description,omitempty"`
	IndexCommand *string `json:"index_command,omitempty"`
	IndexName *string `json:"index_name,omitempty"`
	TableName *string `json:"table_name,omitempty"`
}

// NewHashedTimestampColumnFinderResponse instantiates a new HashedTimestampColumnFinderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashedTimestampColumnFinderResponse() *HashedTimestampColumnFinderResponse {
	this := HashedTimestampColumnFinderResponse{}
	return &this
}

// NewHashedTimestampColumnFinderResponseWithDefaults instantiates a new HashedTimestampColumnFinderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashedTimestampColumnFinderResponseWithDefaults() *HashedTimestampColumnFinderResponse {
	this := HashedTimestampColumnFinderResponse{}
	return &this
}

// GetCurrentDatabase returns the CurrentDatabase field value if set, zero value otherwise.
func (o *HashedTimestampColumnFinderResponse) GetCurrentDatabase() string {
	if o == nil || o.CurrentDatabase == nil {
		var ret string
		return ret
	}
	return *o.CurrentDatabase
}

// GetCurrentDatabaseOk returns a tuple with the CurrentDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashedTimestampColumnFinderResponse) GetCurrentDatabaseOk() (*string, bool) {
	if o == nil || o.CurrentDatabase == nil {
		return nil, false
	}
	return o.CurrentDatabase, true
}

// HasCurrentDatabase returns a boolean if a field has been set.
func (o *HashedTimestampColumnFinderResponse) HasCurrentDatabase() bool {
	if o != nil && o.CurrentDatabase != nil {
		return true
	}

	return false
}

// SetCurrentDatabase gets a reference to the given string and assigns it to the CurrentDatabase field.
func (o *HashedTimestampColumnFinderResponse) SetCurrentDatabase(v string) {
	o.CurrentDatabase = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HashedTimestampColumnFinderResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashedTimestampColumnFinderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HashedTimestampColumnFinderResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HashedTimestampColumnFinderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIndexCommand returns the IndexCommand field value if set, zero value otherwise.
func (o *HashedTimestampColumnFinderResponse) GetIndexCommand() string {
	if o == nil || o.IndexCommand == nil {
		var ret string
		return ret
	}
	return *o.IndexCommand
}

// GetIndexCommandOk returns a tuple with the IndexCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashedTimestampColumnFinderResponse) GetIndexCommandOk() (*string, bool) {
	if o == nil || o.IndexCommand == nil {
		return nil, false
	}
	return o.IndexCommand, true
}

// HasIndexCommand returns a boolean if a field has been set.
func (o *HashedTimestampColumnFinderResponse) HasIndexCommand() bool {
	if o != nil && o.IndexCommand != nil {
		return true
	}

	return false
}

// SetIndexCommand gets a reference to the given string and assigns it to the IndexCommand field.
func (o *HashedTimestampColumnFinderResponse) SetIndexCommand(v string) {
	o.IndexCommand = &v
}

// GetIndexName returns the IndexName field value if set, zero value otherwise.
func (o *HashedTimestampColumnFinderResponse) GetIndexName() string {
	if o == nil || o.IndexName == nil {
		var ret string
		return ret
	}
	return *o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashedTimestampColumnFinderResponse) GetIndexNameOk() (*string, bool) {
	if o == nil || o.IndexName == nil {
		return nil, false
	}
	return o.IndexName, true
}

// HasIndexName returns a boolean if a field has been set.
func (o *HashedTimestampColumnFinderResponse) HasIndexName() bool {
	if o != nil && o.IndexName != nil {
		return true
	}

	return false
}

// SetIndexName gets a reference to the given string and assigns it to the IndexName field.
func (o *HashedTimestampColumnFinderResponse) SetIndexName(v string) {
	o.IndexName = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *HashedTimestampColumnFinderResponse) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashedTimestampColumnFinderResponse) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *HashedTimestampColumnFinderResponse) HasTableName() bool {
	if o != nil && o.TableName != nil {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *HashedTimestampColumnFinderResponse) SetTableName(v string) {
	o.TableName = &v
}

func (o HashedTimestampColumnFinderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentDatabase != nil {
		toSerialize["current_database"] = o.CurrentDatabase
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IndexCommand != nil {
		toSerialize["index_command"] = o.IndexCommand
	}
	if o.IndexName != nil {
		toSerialize["index_name"] = o.IndexName
	}
	if o.TableName != nil {
		toSerialize["table_name"] = o.TableName
	}
	return json.Marshal(toSerialize)
}

type NullableHashedTimestampColumnFinderResponse struct {
	value *HashedTimestampColumnFinderResponse
	isSet bool
}

func (v NullableHashedTimestampColumnFinderResponse) Get() *HashedTimestampColumnFinderResponse {
	return v.value
}

func (v *NullableHashedTimestampColumnFinderResponse) Set(val *HashedTimestampColumnFinderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHashedTimestampColumnFinderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHashedTimestampColumnFinderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashedTimestampColumnFinderResponse(val *HashedTimestampColumnFinderResponse) *NullableHashedTimestampColumnFinderResponse {
	return &NullableHashedTimestampColumnFinderResponse{value: val, isSet: true}
}

func (v NullableHashedTimestampColumnFinderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashedTimestampColumnFinderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


