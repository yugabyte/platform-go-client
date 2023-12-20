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

// Tablespace struct for Tablespace
type Tablespace struct {
	ReplicaPlacement ReplicaPlacement `json:"replicaPlacement"`
	TablespaceName string `json:"tablespaceName"`
}

// NewTablespace instantiates a new Tablespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTablespace(replicaPlacement ReplicaPlacement, tablespaceName string, ) *Tablespace {
	this := Tablespace{}
	this.ReplicaPlacement = replicaPlacement
	this.TablespaceName = tablespaceName
	return &this
}

// NewTablespaceWithDefaults instantiates a new Tablespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTablespaceWithDefaults() *Tablespace {
	this := Tablespace{}
	return &this
}

// GetReplicaPlacement returns the ReplicaPlacement field value
func (o *Tablespace) GetReplicaPlacement() ReplicaPlacement {
	if o == nil  {
		var ret ReplicaPlacement
		return ret
	}

	return o.ReplicaPlacement
}

// GetReplicaPlacementOk returns a tuple with the ReplicaPlacement field value
// and a boolean to check if the value has been set.
func (o *Tablespace) GetReplicaPlacementOk() (*ReplicaPlacement, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReplicaPlacement, true
}

// SetReplicaPlacement sets field value
func (o *Tablespace) SetReplicaPlacement(v ReplicaPlacement) {
	o.ReplicaPlacement = v
}

// GetTablespaceName returns the TablespaceName field value
func (o *Tablespace) GetTablespaceName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TablespaceName
}

// GetTablespaceNameOk returns a tuple with the TablespaceName field value
// and a boolean to check if the value has been set.
func (o *Tablespace) GetTablespaceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TablespaceName, true
}

// SetTablespaceName sets field value
func (o *Tablespace) SetTablespaceName(v string) {
	o.TablespaceName = v
}

func (o Tablespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["replicaPlacement"] = o.ReplicaPlacement
	}
	if true {
		toSerialize["tablespaceName"] = o.TablespaceName
	}
	return json.Marshal(toSerialize)
}

type NullableTablespace struct {
	value *Tablespace
	isSet bool
}

func (v NullableTablespace) Get() *Tablespace {
	return v.value
}

func (v *NullableTablespace) Set(val *Tablespace) {
	v.value = val
	v.isSet = true
}

func (v NullableTablespace) IsSet() bool {
	return v.isSet
}

func (v *NullableTablespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTablespace(val *Tablespace) *NullableTablespace {
	return &NullableTablespace{value: val, isSet: true}
}

func (v NullableTablespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTablespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


