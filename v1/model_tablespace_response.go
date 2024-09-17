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

// TablespaceResponse struct for TablespaceResponse
type TablespaceResponse struct {
	ConflictingTablespaces []string `json:"conflictingTablespaces"`
	ContainsTablespaces bool `json:"containsTablespaces"`
	UnsupportedTablespaces []string `json:"unsupportedTablespaces"`
}

// NewTablespaceResponse instantiates a new TablespaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTablespaceResponse(conflictingTablespaces []string, containsTablespaces bool, unsupportedTablespaces []string) *TablespaceResponse {
	this := TablespaceResponse{}
	this.ConflictingTablespaces = conflictingTablespaces
	this.ContainsTablespaces = containsTablespaces
	this.UnsupportedTablespaces = unsupportedTablespaces
	return &this
}

// NewTablespaceResponseWithDefaults instantiates a new TablespaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTablespaceResponseWithDefaults() *TablespaceResponse {
	this := TablespaceResponse{}
	return &this
}

// GetConflictingTablespaces returns the ConflictingTablespaces field value
func (o *TablespaceResponse) GetConflictingTablespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConflictingTablespaces
}

// GetConflictingTablespacesOk returns a tuple with the ConflictingTablespaces field value
// and a boolean to check if the value has been set.
func (o *TablespaceResponse) GetConflictingTablespacesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConflictingTablespaces, true
}

// SetConflictingTablespaces sets field value
func (o *TablespaceResponse) SetConflictingTablespaces(v []string) {
	o.ConflictingTablespaces = v
}

// GetContainsTablespaces returns the ContainsTablespaces field value
func (o *TablespaceResponse) GetContainsTablespaces() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ContainsTablespaces
}

// GetContainsTablespacesOk returns a tuple with the ContainsTablespaces field value
// and a boolean to check if the value has been set.
func (o *TablespaceResponse) GetContainsTablespacesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContainsTablespaces, true
}

// SetContainsTablespaces sets field value
func (o *TablespaceResponse) SetContainsTablespaces(v bool) {
	o.ContainsTablespaces = v
}

// GetUnsupportedTablespaces returns the UnsupportedTablespaces field value
func (o *TablespaceResponse) GetUnsupportedTablespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UnsupportedTablespaces
}

// GetUnsupportedTablespacesOk returns a tuple with the UnsupportedTablespaces field value
// and a boolean to check if the value has been set.
func (o *TablespaceResponse) GetUnsupportedTablespacesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnsupportedTablespaces, true
}

// SetUnsupportedTablespaces sets field value
func (o *TablespaceResponse) SetUnsupportedTablespaces(v []string) {
	o.UnsupportedTablespaces = v
}

func (o TablespaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["conflictingTablespaces"] = o.ConflictingTablespaces
	}
	if true {
		toSerialize["containsTablespaces"] = o.ContainsTablespaces
	}
	if true {
		toSerialize["unsupportedTablespaces"] = o.UnsupportedTablespaces
	}
	return json.Marshal(toSerialize)
}

type NullableTablespaceResponse struct {
	value *TablespaceResponse
	isSet bool
}

func (v NullableTablespaceResponse) Get() *TablespaceResponse {
	return v.value
}

func (v *NullableTablespaceResponse) Set(val *TablespaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTablespaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTablespaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTablespaceResponse(val *TablespaceResponse) *NullableTablespaceResponse {
	return &NullableTablespaceResponse{value: val, isSet: true}
}

func (v NullableTablespaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTablespaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


