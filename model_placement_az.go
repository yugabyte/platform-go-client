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

// PlacementAZ struct for PlacementAZ
type PlacementAZ struct {
	IsAffinitized *bool `json:"isAffinitized,omitempty"`
	LbName *string `json:"lbName,omitempty"`
	Name *string `json:"name,omitempty"`
	NumNodesInAZ *int32 `json:"numNodesInAZ,omitempty"`
	ReplicationFactor *int32 `json:"replicationFactor,omitempty"`
	SecondarySubnet *string `json:"secondarySubnet,omitempty"`
	Subnet *string `json:"subnet,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewPlacementAZ instantiates a new PlacementAZ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlacementAZ() *PlacementAZ {
	this := PlacementAZ{}
	return &this
}

// NewPlacementAZWithDefaults instantiates a new PlacementAZ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlacementAZWithDefaults() *PlacementAZ {
	this := PlacementAZ{}
	return &this
}

// GetIsAffinitized returns the IsAffinitized field value if set, zero value otherwise.
func (o *PlacementAZ) GetIsAffinitized() bool {
	if o == nil || o.IsAffinitized == nil {
		var ret bool
		return ret
	}
	return *o.IsAffinitized
}

// GetIsAffinitizedOk returns a tuple with the IsAffinitized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAZ) GetIsAffinitizedOk() (*bool, bool) {
	if o == nil || o.IsAffinitized == nil {
		return nil, false
	}
	return o.IsAffinitized, true
}

// HasIsAffinitized returns a boolean if a field has been set.
func (o *PlacementAZ) HasIsAffinitized() bool {
	if o != nil && o.IsAffinitized != nil {
		return true
	}

	return false
}

// SetIsAffinitized gets a reference to the given bool and assigns it to the IsAffinitized field.
func (o *PlacementAZ) SetIsAffinitized(v bool) {
	o.IsAffinitized = &v
}

// GetLbName returns the LbName field value if set, zero value otherwise.
func (o *PlacementAZ) GetLbName() string {
	if o == nil || o.LbName == nil {
		var ret string
		return ret
	}
	return *o.LbName
}

// GetLbNameOk returns a tuple with the LbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAZ) GetLbNameOk() (*string, bool) {
	if o == nil || o.LbName == nil {
		return nil, false
	}
	return o.LbName, true
}

// HasLbName returns a boolean if a field has been set.
func (o *PlacementAZ) HasLbName() bool {
	if o != nil && o.LbName != nil {
		return true
	}

	return false
}

// SetLbName gets a reference to the given string and assigns it to the LbName field.
func (o *PlacementAZ) SetLbName(v string) {
	o.LbName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlacementAZ) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAZ) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlacementAZ) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlacementAZ) SetName(v string) {
	o.Name = &v
}

// GetNumNodesInAZ returns the NumNodesInAZ field value if set, zero value otherwise.
func (o *PlacementAZ) GetNumNodesInAZ() int32 {
	if o == nil || o.NumNodesInAZ == nil {
		var ret int32
		return ret
	}
	return *o.NumNodesInAZ
}

// GetNumNodesInAZOk returns a tuple with the NumNodesInAZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAZ) GetNumNodesInAZOk() (*int32, bool) {
	if o == nil || o.NumNodesInAZ == nil {
		return nil, false
	}
	return o.NumNodesInAZ, true
}

// HasNumNodesInAZ returns a boolean if a field has been set.
func (o *PlacementAZ) HasNumNodesInAZ() bool {
	if o != nil && o.NumNodesInAZ != nil {
		return true
	}

	return false
}

// SetNumNodesInAZ gets a reference to the given int32 and assigns it to the NumNodesInAZ field.
func (o *PlacementAZ) SetNumNodesInAZ(v int32) {
	o.NumNodesInAZ = &v
}

// GetReplicationFactor returns the ReplicationFactor field value if set, zero value otherwise.
func (o *PlacementAZ) GetReplicationFactor() int32 {
	if o == nil || o.ReplicationFactor == nil {
		var ret int32
		return ret
	}
	return *o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAZ) GetReplicationFactorOk() (*int32, bool) {
	if o == nil || o.ReplicationFactor == nil {
		return nil, false
	}
	return o.ReplicationFactor, true
}

// HasReplicationFactor returns a boolean if a field has been set.
func (o *PlacementAZ) HasReplicationFactor() bool {
	if o != nil && o.ReplicationFactor != nil {
		return true
	}

	return false
}

// SetReplicationFactor gets a reference to the given int32 and assigns it to the ReplicationFactor field.
func (o *PlacementAZ) SetReplicationFactor(v int32) {
	o.ReplicationFactor = &v
}

// GetSecondarySubnet returns the SecondarySubnet field value if set, zero value otherwise.
func (o *PlacementAZ) GetSecondarySubnet() string {
	if o == nil || o.SecondarySubnet == nil {
		var ret string
		return ret
	}
	return *o.SecondarySubnet
}

// GetSecondarySubnetOk returns a tuple with the SecondarySubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAZ) GetSecondarySubnetOk() (*string, bool) {
	if o == nil || o.SecondarySubnet == nil {
		return nil, false
	}
	return o.SecondarySubnet, true
}

// HasSecondarySubnet returns a boolean if a field has been set.
func (o *PlacementAZ) HasSecondarySubnet() bool {
	if o != nil && o.SecondarySubnet != nil {
		return true
	}

	return false
}

// SetSecondarySubnet gets a reference to the given string and assigns it to the SecondarySubnet field.
func (o *PlacementAZ) SetSecondarySubnet(v string) {
	o.SecondarySubnet = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *PlacementAZ) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAZ) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *PlacementAZ) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *PlacementAZ) SetSubnet(v string) {
	o.Subnet = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *PlacementAZ) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAZ) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *PlacementAZ) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *PlacementAZ) SetUuid(v string) {
	o.Uuid = &v
}

func (o PlacementAZ) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsAffinitized != nil {
		toSerialize["isAffinitized"] = o.IsAffinitized
	}
	if o.LbName != nil {
		toSerialize["lbName"] = o.LbName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NumNodesInAZ != nil {
		toSerialize["numNodesInAZ"] = o.NumNodesInAZ
	}
	if o.ReplicationFactor != nil {
		toSerialize["replicationFactor"] = o.ReplicationFactor
	}
	if o.SecondarySubnet != nil {
		toSerialize["secondarySubnet"] = o.SecondarySubnet
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullablePlacementAZ struct {
	value *PlacementAZ
	isSet bool
}

func (v NullablePlacementAZ) Get() *PlacementAZ {
	return v.value
}

func (v *NullablePlacementAZ) Set(val *PlacementAZ) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacementAZ) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacementAZ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacementAZ(val *PlacementAZ) *NullablePlacementAZ {
	return &NullablePlacementAZ{value: val, isSet: true}
}

func (v NullablePlacementAZ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlacementAZ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


