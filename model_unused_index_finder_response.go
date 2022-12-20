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

// UnusedIndexFinderResponse struct for UnusedIndexFinderResponse
type UnusedIndexFinderResponse struct {
	CurrentDatabase *string `json:"current_database,omitempty"`
	Description *string `json:"description,omitempty"`
	IndexCommand *string `json:"index_command,omitempty"`
	IndexName *string `json:"index_name,omitempty"`
	TableName *string `json:"table_name,omitempty"`
}

// NewUnusedIndexFinderResponse instantiates a new UnusedIndexFinderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnusedIndexFinderResponse() *UnusedIndexFinderResponse {
	this := UnusedIndexFinderResponse{}
	return &this
}

// NewUnusedIndexFinderResponseWithDefaults instantiates a new UnusedIndexFinderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnusedIndexFinderResponseWithDefaults() *UnusedIndexFinderResponse {
	this := UnusedIndexFinderResponse{}
	return &this
}

// GetCurrentDatabase returns the CurrentDatabase field value if set, zero value otherwise.
func (o *UnusedIndexFinderResponse) GetCurrentDatabase() string {
	if o == nil || o.CurrentDatabase == nil {
		var ret string
		return ret
	}
	return *o.CurrentDatabase
}

// GetCurrentDatabaseOk returns a tuple with the CurrentDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnusedIndexFinderResponse) GetCurrentDatabaseOk() (*string, bool) {
	if o == nil || o.CurrentDatabase == nil {
		return nil, false
	}
	return o.CurrentDatabase, true
}

// HasCurrentDatabase returns a boolean if a field has been set.
func (o *UnusedIndexFinderResponse) HasCurrentDatabase() bool {
	if o != nil && o.CurrentDatabase != nil {
		return true
	}

	return false
}

// SetCurrentDatabase gets a reference to the given string and assigns it to the CurrentDatabase field.
func (o *UnusedIndexFinderResponse) SetCurrentDatabase(v string) {
	o.CurrentDatabase = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnusedIndexFinderResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnusedIndexFinderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnusedIndexFinderResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnusedIndexFinderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIndexCommand returns the IndexCommand field value if set, zero value otherwise.
func (o *UnusedIndexFinderResponse) GetIndexCommand() string {
	if o == nil || o.IndexCommand == nil {
		var ret string
		return ret
	}
	return *o.IndexCommand
}

// GetIndexCommandOk returns a tuple with the IndexCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnusedIndexFinderResponse) GetIndexCommandOk() (*string, bool) {
	if o == nil || o.IndexCommand == nil {
		return nil, false
	}
	return o.IndexCommand, true
}

// HasIndexCommand returns a boolean if a field has been set.
func (o *UnusedIndexFinderResponse) HasIndexCommand() bool {
	if o != nil && o.IndexCommand != nil {
		return true
	}

	return false
}

// SetIndexCommand gets a reference to the given string and assigns it to the IndexCommand field.
func (o *UnusedIndexFinderResponse) SetIndexCommand(v string) {
	o.IndexCommand = &v
}

// GetIndexName returns the IndexName field value if set, zero value otherwise.
func (o *UnusedIndexFinderResponse) GetIndexName() string {
	if o == nil || o.IndexName == nil {
		var ret string
		return ret
	}
	return *o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnusedIndexFinderResponse) GetIndexNameOk() (*string, bool) {
	if o == nil || o.IndexName == nil {
		return nil, false
	}
	return o.IndexName, true
}

// HasIndexName returns a boolean if a field has been set.
func (o *UnusedIndexFinderResponse) HasIndexName() bool {
	if o != nil && o.IndexName != nil {
		return true
	}

	return false
}

// SetIndexName gets a reference to the given string and assigns it to the IndexName field.
func (o *UnusedIndexFinderResponse) SetIndexName(v string) {
	o.IndexName = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *UnusedIndexFinderResponse) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnusedIndexFinderResponse) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *UnusedIndexFinderResponse) HasTableName() bool {
	if o != nil && o.TableName != nil {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *UnusedIndexFinderResponse) SetTableName(v string) {
	o.TableName = &v
}

func (o UnusedIndexFinderResponse) MarshalJSON() ([]byte, error) {
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

type NullableUnusedIndexFinderResponse struct {
	value *UnusedIndexFinderResponse
	isSet bool
}

func (v NullableUnusedIndexFinderResponse) Get() *UnusedIndexFinderResponse {
	return v.value
}

func (v *NullableUnusedIndexFinderResponse) Set(val *UnusedIndexFinderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnusedIndexFinderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnusedIndexFinderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnusedIndexFinderResponse(val *UnusedIndexFinderResponse) *NullableUnusedIndexFinderResponse {
	return &NullableUnusedIndexFinderResponse{value: val, isSet: true}
}

func (v NullableUnusedIndexFinderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnusedIndexFinderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


