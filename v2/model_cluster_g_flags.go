/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// ClusterGFlags GFlags for a single cluster of a YugabyteDB Universe. Used as part of ClusterSpec at Universe create time, and as part of UniverseEditGFlags to edit GFlags for a Universe.
type ClusterGFlags struct {
	// GFlags applied on TServer process
	Tserver *map[string]string `json:"tserver,omitempty"`
	// GFlags applied on Master process
	Master *map[string]string `json:"master,omitempty"`
	// GFlags per availability zone uuid
	AzGflags *map[string]AvailabilityZoneGFlags `json:"az_gflags,omitempty"`
}

// NewClusterGFlags instantiates a new ClusterGFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterGFlags() *ClusterGFlags {
	this := ClusterGFlags{}
	return &this
}

// NewClusterGFlagsWithDefaults instantiates a new ClusterGFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterGFlagsWithDefaults() *ClusterGFlags {
	this := ClusterGFlags{}
	return &this
}

// GetTserver returns the Tserver field value if set, zero value otherwise.
func (o *ClusterGFlags) GetTserver() map[string]string {
	if o == nil || o.Tserver == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tserver
}

// GetTserverOk returns a tuple with the Tserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterGFlags) GetTserverOk() (*map[string]string, bool) {
	if o == nil || o.Tserver == nil {
		return nil, false
	}
	return o.Tserver, true
}

// HasTserver returns a boolean if a field has been set.
func (o *ClusterGFlags) HasTserver() bool {
	if o != nil && o.Tserver != nil {
		return true
	}

	return false
}

// SetTserver gets a reference to the given map[string]string and assigns it to the Tserver field.
func (o *ClusterGFlags) SetTserver(v map[string]string) {
	o.Tserver = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *ClusterGFlags) GetMaster() map[string]string {
	if o == nil || o.Master == nil {
		var ret map[string]string
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterGFlags) GetMasterOk() (*map[string]string, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *ClusterGFlags) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given map[string]string and assigns it to the Master field.
func (o *ClusterGFlags) SetMaster(v map[string]string) {
	o.Master = &v
}

// GetAzGflags returns the AzGflags field value if set, zero value otherwise.
func (o *ClusterGFlags) GetAzGflags() map[string]AvailabilityZoneGFlags {
	if o == nil || o.AzGflags == nil {
		var ret map[string]AvailabilityZoneGFlags
		return ret
	}
	return *o.AzGflags
}

// GetAzGflagsOk returns a tuple with the AzGflags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterGFlags) GetAzGflagsOk() (*map[string]AvailabilityZoneGFlags, bool) {
	if o == nil || o.AzGflags == nil {
		return nil, false
	}
	return o.AzGflags, true
}

// HasAzGflags returns a boolean if a field has been set.
func (o *ClusterGFlags) HasAzGflags() bool {
	if o != nil && o.AzGflags != nil {
		return true
	}

	return false
}

// SetAzGflags gets a reference to the given map[string]AvailabilityZoneGFlags and assigns it to the AzGflags field.
func (o *ClusterGFlags) SetAzGflags(v map[string]AvailabilityZoneGFlags) {
	o.AzGflags = &v
}

func (o ClusterGFlags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tserver != nil {
		toSerialize["tserver"] = o.Tserver
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
	if o.AzGflags != nil {
		toSerialize["az_gflags"] = o.AzGflags
	}
	return json.Marshal(toSerialize)
}

type NullableClusterGFlags struct {
	value *ClusterGFlags
	isSet bool
}

func (v NullableClusterGFlags) Get() *ClusterGFlags {
	return v.value
}

func (v *NullableClusterGFlags) Set(val *ClusterGFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterGFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterGFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterGFlags(val *ClusterGFlags) *NullableClusterGFlags {
	return &NullableClusterGFlags{value: val, isSet: true}
}

func (v NullableClusterGFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterGFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


