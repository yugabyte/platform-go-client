/*
 * Yugabyte Platform APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yb_platform_client

import (
	"encoding/json"
)

// AvailabilityZone Availability zone (AZ) for a region
type AvailabilityZone struct {
	// AZ status. This value is `true` for an active AZ.
	Active *bool `json:"active,omitempty"`
	// AZ code
	Code *string `json:"code,omitempty"`
	// AZ configuration values
	Config *map[string]string `json:"config,omitempty"`
	// Path to Kubernetes configuration file
	KubeconfigPath *string `json:"kubeconfigPath,omitempty"`
	// AZ name
	Name string `json:"name"`
	// AZ secondary subnet
	SecondarySubnet *string `json:"secondarySubnet,omitempty"`
	// AZ subnet
	Subnet *string `json:"subnet,omitempty"`
	// AZ UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewAvailabilityZone instantiates a new AvailabilityZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityZone(name string, ) *AvailabilityZone {
	this := AvailabilityZone{}
	this.Name = name
	return &this
}

// NewAvailabilityZoneWithDefaults instantiates a new AvailabilityZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityZoneWithDefaults() *AvailabilityZone {
	this := AvailabilityZone{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AvailabilityZone) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AvailabilityZone) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AvailabilityZone) SetActive(v bool) {
	o.Active = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AvailabilityZone) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AvailabilityZone) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AvailabilityZone) SetCode(v string) {
	o.Code = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AvailabilityZone) GetConfig() map[string]string {
	if o == nil || o.Config == nil {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetConfigOk() (*map[string]string, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *AvailabilityZone) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *AvailabilityZone) SetConfig(v map[string]string) {
	o.Config = &v
}

// GetKubeconfigPath returns the KubeconfigPath field value if set, zero value otherwise.
func (o *AvailabilityZone) GetKubeconfigPath() string {
	if o == nil || o.KubeconfigPath == nil {
		var ret string
		return ret
	}
	return *o.KubeconfigPath
}

// GetKubeconfigPathOk returns a tuple with the KubeconfigPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetKubeconfigPathOk() (*string, bool) {
	if o == nil || o.KubeconfigPath == nil {
		return nil, false
	}
	return o.KubeconfigPath, true
}

// HasKubeconfigPath returns a boolean if a field has been set.
func (o *AvailabilityZone) HasKubeconfigPath() bool {
	if o != nil && o.KubeconfigPath != nil {
		return true
	}

	return false
}

// SetKubeconfigPath gets a reference to the given string and assigns it to the KubeconfigPath field.
func (o *AvailabilityZone) SetKubeconfigPath(v string) {
	o.KubeconfigPath = &v
}

// GetName returns the Name field value
func (o *AvailabilityZone) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AvailabilityZone) SetName(v string) {
	o.Name = v
}

// GetSecondarySubnet returns the SecondarySubnet field value if set, zero value otherwise.
func (o *AvailabilityZone) GetSecondarySubnet() string {
	if o == nil || o.SecondarySubnet == nil {
		var ret string
		return ret
	}
	return *o.SecondarySubnet
}

// GetSecondarySubnetOk returns a tuple with the SecondarySubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetSecondarySubnetOk() (*string, bool) {
	if o == nil || o.SecondarySubnet == nil {
		return nil, false
	}
	return o.SecondarySubnet, true
}

// HasSecondarySubnet returns a boolean if a field has been set.
func (o *AvailabilityZone) HasSecondarySubnet() bool {
	if o != nil && o.SecondarySubnet != nil {
		return true
	}

	return false
}

// SetSecondarySubnet gets a reference to the given string and assigns it to the SecondarySubnet field.
func (o *AvailabilityZone) SetSecondarySubnet(v string) {
	o.SecondarySubnet = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *AvailabilityZone) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *AvailabilityZone) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *AvailabilityZone) SetSubnet(v string) {
	o.Subnet = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AvailabilityZone) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AvailabilityZone) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AvailabilityZone) SetUuid(v string) {
	o.Uuid = &v
}

func (o AvailabilityZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.KubeconfigPath != nil {
		toSerialize["kubeconfigPath"] = o.KubeconfigPath
	}
	if true {
		toSerialize["name"] = o.Name
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

type NullableAvailabilityZone struct {
	value *AvailabilityZone
	isSet bool
}

func (v NullableAvailabilityZone) Get() *AvailabilityZone {
	return v.value
}

func (v *NullableAvailabilityZone) Set(val *AvailabilityZone) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityZone) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityZone(val *AvailabilityZone) *NullableAvailabilityZone {
	return &NullableAvailabilityZone{value: val, isSet: true}
}

func (v NullableAvailabilityZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


