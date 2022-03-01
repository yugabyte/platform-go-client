/*
 * Yugabyte Platform APIs
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

// ImportUniverseFormData struct for ImportUniverseFormData
type ImportUniverseFormData struct {
	CloudName string `json:"cloudName"`
	CloudProviderType string `json:"cloudProviderType"`
	CurrentState string `json:"currentState"`
	InstanceType string `json:"instanceType"`
	MasterAddresses string `json:"masterAddresses"`
	ProviderType string `json:"providerType"`
	RegionCode string `json:"regionCode"`
	RegionName string `json:"regionName"`
	ReplicationFactor int32 `json:"replicationFactor"`
	SingleStep bool `json:"singleStep"`
	UniverseName string `json:"universeName"`
	UniverseUUID string `json:"universeUUID"`
	ZoneCode string `json:"zoneCode"`
	ZoneName string `json:"zoneName"`
}

// NewImportUniverseFormData instantiates a new ImportUniverseFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportUniverseFormData(cloudName string, cloudProviderType string, currentState string, instanceType string, masterAddresses string, providerType string, regionCode string, regionName string, replicationFactor int32, singleStep bool, universeName string, universeUUID string, zoneCode string, zoneName string, ) *ImportUniverseFormData {
	this := ImportUniverseFormData{}
	this.CloudName = cloudName
	this.CloudProviderType = cloudProviderType
	this.CurrentState = currentState
	this.InstanceType = instanceType
	this.MasterAddresses = masterAddresses
	this.ProviderType = providerType
	this.RegionCode = regionCode
	this.RegionName = regionName
	this.ReplicationFactor = replicationFactor
	this.SingleStep = singleStep
	this.UniverseName = universeName
	this.UniverseUUID = universeUUID
	this.ZoneCode = zoneCode
	this.ZoneName = zoneName
	return &this
}

// NewImportUniverseFormDataWithDefaults instantiates a new ImportUniverseFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportUniverseFormDataWithDefaults() *ImportUniverseFormData {
	this := ImportUniverseFormData{}
	return &this
}

// GetCloudName returns the CloudName field value
func (o *ImportUniverseFormData) GetCloudName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CloudName
}

// GetCloudNameOk returns a tuple with the CloudName field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetCloudNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudName, true
}

// SetCloudName sets field value
func (o *ImportUniverseFormData) SetCloudName(v string) {
	o.CloudName = v
}

// GetCloudProviderType returns the CloudProviderType field value
func (o *ImportUniverseFormData) GetCloudProviderType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CloudProviderType
}

// GetCloudProviderTypeOk returns a tuple with the CloudProviderType field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetCloudProviderTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudProviderType, true
}

// SetCloudProviderType sets field value
func (o *ImportUniverseFormData) SetCloudProviderType(v string) {
	o.CloudProviderType = v
}

// GetCurrentState returns the CurrentState field value
func (o *ImportUniverseFormData) GetCurrentState() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetCurrentStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentState, true
}

// SetCurrentState sets field value
func (o *ImportUniverseFormData) SetCurrentState(v string) {
	o.CurrentState = v
}

// GetInstanceType returns the InstanceType field value
func (o *ImportUniverseFormData) GetInstanceType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetInstanceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *ImportUniverseFormData) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetMasterAddresses returns the MasterAddresses field value
func (o *ImportUniverseFormData) GetMasterAddresses() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MasterAddresses
}

// GetMasterAddressesOk returns a tuple with the MasterAddresses field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetMasterAddressesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MasterAddresses, true
}

// SetMasterAddresses sets field value
func (o *ImportUniverseFormData) SetMasterAddresses(v string) {
	o.MasterAddresses = v
}

// GetProviderType returns the ProviderType field value
func (o *ImportUniverseFormData) GetProviderType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetProviderTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *ImportUniverseFormData) SetProviderType(v string) {
	o.ProviderType = v
}

// GetRegionCode returns the RegionCode field value
func (o *ImportUniverseFormData) GetRegionCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetRegionCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionCode, true
}

// SetRegionCode sets field value
func (o *ImportUniverseFormData) SetRegionCode(v string) {
	o.RegionCode = v
}

// GetRegionName returns the RegionName field value
func (o *ImportUniverseFormData) GetRegionName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetRegionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *ImportUniverseFormData) SetRegionName(v string) {
	o.RegionName = v
}

// GetReplicationFactor returns the ReplicationFactor field value
func (o *ImportUniverseFormData) GetReplicationFactor() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetReplicationFactorOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReplicationFactor, true
}

// SetReplicationFactor sets field value
func (o *ImportUniverseFormData) SetReplicationFactor(v int32) {
	o.ReplicationFactor = v
}

// GetSingleStep returns the SingleStep field value
func (o *ImportUniverseFormData) GetSingleStep() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.SingleStep
}

// GetSingleStepOk returns a tuple with the SingleStep field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetSingleStepOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SingleStep, true
}

// SetSingleStep sets field value
func (o *ImportUniverseFormData) SetSingleStep(v bool) {
	o.SingleStep = v
}

// GetUniverseName returns the UniverseName field value
func (o *ImportUniverseFormData) GetUniverseName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniverseName
}

// GetUniverseNameOk returns a tuple with the UniverseName field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetUniverseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseName, true
}

// SetUniverseName sets field value
func (o *ImportUniverseFormData) SetUniverseName(v string) {
	o.UniverseName = v
}

// GetUniverseUUID returns the UniverseUUID field value
func (o *ImportUniverseFormData) GetUniverseUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniverseUUID
}

// GetUniverseUUIDOk returns a tuple with the UniverseUUID field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetUniverseUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniverseUUID, true
}

// SetUniverseUUID sets field value
func (o *ImportUniverseFormData) SetUniverseUUID(v string) {
	o.UniverseUUID = v
}

// GetZoneCode returns the ZoneCode field value
func (o *ImportUniverseFormData) GetZoneCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ZoneCode
}

// GetZoneCodeOk returns a tuple with the ZoneCode field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetZoneCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ZoneCode, true
}

// SetZoneCode sets field value
func (o *ImportUniverseFormData) SetZoneCode(v string) {
	o.ZoneCode = v
}

// GetZoneName returns the ZoneName field value
func (o *ImportUniverseFormData) GetZoneName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value
// and a boolean to check if the value has been set.
func (o *ImportUniverseFormData) GetZoneNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ZoneName, true
}

// SetZoneName sets field value
func (o *ImportUniverseFormData) SetZoneName(v string) {
	o.ZoneName = v
}

func (o ImportUniverseFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloudName"] = o.CloudName
	}
	if true {
		toSerialize["cloudProviderType"] = o.CloudProviderType
	}
	if true {
		toSerialize["currentState"] = o.CurrentState
	}
	if true {
		toSerialize["instanceType"] = o.InstanceType
	}
	if true {
		toSerialize["masterAddresses"] = o.MasterAddresses
	}
	if true {
		toSerialize["providerType"] = o.ProviderType
	}
	if true {
		toSerialize["regionCode"] = o.RegionCode
	}
	if true {
		toSerialize["regionName"] = o.RegionName
	}
	if true {
		toSerialize["replicationFactor"] = o.ReplicationFactor
	}
	if true {
		toSerialize["singleStep"] = o.SingleStep
	}
	if true {
		toSerialize["universeName"] = o.UniverseName
	}
	if true {
		toSerialize["universeUUID"] = o.UniverseUUID
	}
	if true {
		toSerialize["zoneCode"] = o.ZoneCode
	}
	if true {
		toSerialize["zoneName"] = o.ZoneName
	}
	return json.Marshal(toSerialize)
}

type NullableImportUniverseFormData struct {
	value *ImportUniverseFormData
	isSet bool
}

func (v NullableImportUniverseFormData) Get() *ImportUniverseFormData {
	return v.value
}

func (v *NullableImportUniverseFormData) Set(val *ImportUniverseFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableImportUniverseFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableImportUniverseFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportUniverseFormData(val *ImportUniverseFormData) *NullableImportUniverseFormData {
	return &NullableImportUniverseFormData{value: val, isSet: true}
}

func (v NullableImportUniverseFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportUniverseFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


