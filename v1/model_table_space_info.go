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

// TableSpaceInfo Tablespace information
type TableSpaceInfo struct {
	// Tablespace Name
	Name string `json:"name"`
	// numReplicas
	NumReplicas *int32 `json:"numReplicas,omitempty"`
	// placements
	PlacementBlocks []PlacementBlock `json:"placementBlocks"`
}

// NewTableSpaceInfo instantiates a new TableSpaceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableSpaceInfo(name string, placementBlocks []PlacementBlock) *TableSpaceInfo {
	this := TableSpaceInfo{}
	this.Name = name
	this.PlacementBlocks = placementBlocks
	return &this
}

// NewTableSpaceInfoWithDefaults instantiates a new TableSpaceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableSpaceInfoWithDefaults() *TableSpaceInfo {
	this := TableSpaceInfo{}
	return &this
}

// GetName returns the Name field value
func (o *TableSpaceInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TableSpaceInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TableSpaceInfo) SetName(v string) {
	o.Name = v
}

// GetNumReplicas returns the NumReplicas field value if set, zero value otherwise.
func (o *TableSpaceInfo) GetNumReplicas() int32 {
	if o == nil || o.NumReplicas == nil {
		var ret int32
		return ret
	}
	return *o.NumReplicas
}

// GetNumReplicasOk returns a tuple with the NumReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableSpaceInfo) GetNumReplicasOk() (*int32, bool) {
	if o == nil || o.NumReplicas == nil {
		return nil, false
	}
	return o.NumReplicas, true
}

// HasNumReplicas returns a boolean if a field has been set.
func (o *TableSpaceInfo) HasNumReplicas() bool {
	if o != nil && o.NumReplicas != nil {
		return true
	}

	return false
}

// SetNumReplicas gets a reference to the given int32 and assigns it to the NumReplicas field.
func (o *TableSpaceInfo) SetNumReplicas(v int32) {
	o.NumReplicas = &v
}

// GetPlacementBlocks returns the PlacementBlocks field value
func (o *TableSpaceInfo) GetPlacementBlocks() []PlacementBlock {
	if o == nil {
		var ret []PlacementBlock
		return ret
	}

	return o.PlacementBlocks
}

// GetPlacementBlocksOk returns a tuple with the PlacementBlocks field value
// and a boolean to check if the value has been set.
func (o *TableSpaceInfo) GetPlacementBlocksOk() (*[]PlacementBlock, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlacementBlocks, true
}

// SetPlacementBlocks sets field value
func (o *TableSpaceInfo) SetPlacementBlocks(v []PlacementBlock) {
	o.PlacementBlocks = v
}

func (o TableSpaceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NumReplicas != nil {
		toSerialize["numReplicas"] = o.NumReplicas
	}
	if true {
		toSerialize["placementBlocks"] = o.PlacementBlocks
	}
	return json.Marshal(toSerialize)
}

type NullableTableSpaceInfo struct {
	value *TableSpaceInfo
	isSet bool
}

func (v NullableTableSpaceInfo) Get() *TableSpaceInfo {
	return v.value
}

func (v *NullableTableSpaceInfo) Set(val *TableSpaceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTableSpaceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTableSpaceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableSpaceInfo(val *TableSpaceInfo) *NullableTableSpaceInfo {
	return &NullableTableSpaceInfo{value: val, isSet: true}
}

func (v NullableTableSpaceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableSpaceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


