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

// DrConfigFailoverForm drConfig failover form
type DrConfigFailoverForm struct {
	// New dr replica universe UUID
	DrReplicaUniverseUuid *string `json:"drReplicaUniverseUuid,omitempty"`
	// A map from database ID to its safetime since epoch in micro-seconds to use during unplanned failover
	NamespaceIdSafetimeEpochUsMap *map[string]int64 `json:"namespaceIdSafetimeEpochUsMap,omitempty"`
	// New primary universe UUID
	PrimaryUniverseUuid *string `json:"primaryUniverseUuid,omitempty"`
}

// NewDrConfigFailoverForm instantiates a new DrConfigFailoverForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrConfigFailoverForm() *DrConfigFailoverForm {
	this := DrConfigFailoverForm{}
	return &this
}

// NewDrConfigFailoverFormWithDefaults instantiates a new DrConfigFailoverForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrConfigFailoverFormWithDefaults() *DrConfigFailoverForm {
	this := DrConfigFailoverForm{}
	return &this
}

// GetDrReplicaUniverseUuid returns the DrReplicaUniverseUuid field value if set, zero value otherwise.
func (o *DrConfigFailoverForm) GetDrReplicaUniverseUuid() string {
	if o == nil || o.DrReplicaUniverseUuid == nil {
		var ret string
		return ret
	}
	return *o.DrReplicaUniverseUuid
}

// GetDrReplicaUniverseUuidOk returns a tuple with the DrReplicaUniverseUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigFailoverForm) GetDrReplicaUniverseUuidOk() (*string, bool) {
	if o == nil || o.DrReplicaUniverseUuid == nil {
		return nil, false
	}
	return o.DrReplicaUniverseUuid, true
}

// HasDrReplicaUniverseUuid returns a boolean if a field has been set.
func (o *DrConfigFailoverForm) HasDrReplicaUniverseUuid() bool {
	if o != nil && o.DrReplicaUniverseUuid != nil {
		return true
	}

	return false
}

// SetDrReplicaUniverseUuid gets a reference to the given string and assigns it to the DrReplicaUniverseUuid field.
func (o *DrConfigFailoverForm) SetDrReplicaUniverseUuid(v string) {
	o.DrReplicaUniverseUuid = &v
}

// GetNamespaceIdSafetimeEpochUsMap returns the NamespaceIdSafetimeEpochUsMap field value if set, zero value otherwise.
func (o *DrConfigFailoverForm) GetNamespaceIdSafetimeEpochUsMap() map[string]int64 {
	if o == nil || o.NamespaceIdSafetimeEpochUsMap == nil {
		var ret map[string]int64
		return ret
	}
	return *o.NamespaceIdSafetimeEpochUsMap
}

// GetNamespaceIdSafetimeEpochUsMapOk returns a tuple with the NamespaceIdSafetimeEpochUsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigFailoverForm) GetNamespaceIdSafetimeEpochUsMapOk() (*map[string]int64, bool) {
	if o == nil || o.NamespaceIdSafetimeEpochUsMap == nil {
		return nil, false
	}
	return o.NamespaceIdSafetimeEpochUsMap, true
}

// HasNamespaceIdSafetimeEpochUsMap returns a boolean if a field has been set.
func (o *DrConfigFailoverForm) HasNamespaceIdSafetimeEpochUsMap() bool {
	if o != nil && o.NamespaceIdSafetimeEpochUsMap != nil {
		return true
	}

	return false
}

// SetNamespaceIdSafetimeEpochUsMap gets a reference to the given map[string]int64 and assigns it to the NamespaceIdSafetimeEpochUsMap field.
func (o *DrConfigFailoverForm) SetNamespaceIdSafetimeEpochUsMap(v map[string]int64) {
	o.NamespaceIdSafetimeEpochUsMap = &v
}

// GetPrimaryUniverseUuid returns the PrimaryUniverseUuid field value if set, zero value otherwise.
func (o *DrConfigFailoverForm) GetPrimaryUniverseUuid() string {
	if o == nil || o.PrimaryUniverseUuid == nil {
		var ret string
		return ret
	}
	return *o.PrimaryUniverseUuid
}

// GetPrimaryUniverseUuidOk returns a tuple with the PrimaryUniverseUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrConfigFailoverForm) GetPrimaryUniverseUuidOk() (*string, bool) {
	if o == nil || o.PrimaryUniverseUuid == nil {
		return nil, false
	}
	return o.PrimaryUniverseUuid, true
}

// HasPrimaryUniverseUuid returns a boolean if a field has been set.
func (o *DrConfigFailoverForm) HasPrimaryUniverseUuid() bool {
	if o != nil && o.PrimaryUniverseUuid != nil {
		return true
	}

	return false
}

// SetPrimaryUniverseUuid gets a reference to the given string and assigns it to the PrimaryUniverseUuid field.
func (o *DrConfigFailoverForm) SetPrimaryUniverseUuid(v string) {
	o.PrimaryUniverseUuid = &v
}

func (o DrConfigFailoverForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DrReplicaUniverseUuid != nil {
		toSerialize["drReplicaUniverseUuid"] = o.DrReplicaUniverseUuid
	}
	if o.NamespaceIdSafetimeEpochUsMap != nil {
		toSerialize["namespaceIdSafetimeEpochUsMap"] = o.NamespaceIdSafetimeEpochUsMap
	}
	if o.PrimaryUniverseUuid != nil {
		toSerialize["primaryUniverseUuid"] = o.PrimaryUniverseUuid
	}
	return json.Marshal(toSerialize)
}

type NullableDrConfigFailoverForm struct {
	value *DrConfigFailoverForm
	isSet bool
}

func (v NullableDrConfigFailoverForm) Get() *DrConfigFailoverForm {
	return v.value
}

func (v *NullableDrConfigFailoverForm) Set(val *DrConfigFailoverForm) {
	v.value = val
	v.isSet = true
}

func (v NullableDrConfigFailoverForm) IsSet() bool {
	return v.isSet
}

func (v *NullableDrConfigFailoverForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrConfigFailoverForm(val *DrConfigFailoverForm) *NullableDrConfigFailoverForm {
	return &NullableDrConfigFailoverForm{value: val, isSet: true}
}

func (v NullableDrConfigFailoverForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrConfigFailoverForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


