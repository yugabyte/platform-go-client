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

// UniverseResp Universe-creation response
type UniverseResp struct {
	// Universe creation date
	CreationDate *string `json:"creationDate,omitempty"`
	// DNS name
	DnsName *string `json:"dnsName,omitempty"`
	// Universe name
	Name *string `json:"name,omitempty"`
	// Price
	PricePerHour *float64 `json:"pricePerHour,omitempty"`
	Resources *UniverseResourceDetails `json:"resources,omitempty"`
	// Sample command
	SampleAppCommandTxt *string `json:"sampleAppCommandTxt,omitempty"`
	// Task UUID
	TaskUUID *string `json:"taskUUID,omitempty"`
	// Universe configuration
	UniverseConfig *map[string]string `json:"universeConfig,omitempty"`
	UniverseDetails *UniverseDefinitionTaskParamsResp `json:"universeDetails,omitempty"`
	// Universe UUID
	UniverseUUID *string `json:"universeUUID,omitempty"`
	// Universe version
	Version *int32 `json:"version,omitempty"`
}

// NewUniverseResp instantiates a new UniverseResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseResp() *UniverseResp {
	this := UniverseResp{}
	return &this
}

// NewUniverseRespWithDefaults instantiates a new UniverseResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseRespWithDefaults() *UniverseResp {
	this := UniverseResp{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *UniverseResp) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *UniverseResp) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *UniverseResp) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *UniverseResp) GetDnsName() string {
	if o == nil || o.DnsName == nil {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetDnsNameOk() (*string, bool) {
	if o == nil || o.DnsName == nil {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *UniverseResp) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *UniverseResp) SetDnsName(v string) {
	o.DnsName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UniverseResp) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UniverseResp) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UniverseResp) SetName(v string) {
	o.Name = &v
}

// GetPricePerHour returns the PricePerHour field value if set, zero value otherwise.
func (o *UniverseResp) GetPricePerHour() float64 {
	if o == nil || o.PricePerHour == nil {
		var ret float64
		return ret
	}
	return *o.PricePerHour
}

// GetPricePerHourOk returns a tuple with the PricePerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetPricePerHourOk() (*float64, bool) {
	if o == nil || o.PricePerHour == nil {
		return nil, false
	}
	return o.PricePerHour, true
}

// HasPricePerHour returns a boolean if a field has been set.
func (o *UniverseResp) HasPricePerHour() bool {
	if o != nil && o.PricePerHour != nil {
		return true
	}

	return false
}

// SetPricePerHour gets a reference to the given float64 and assigns it to the PricePerHour field.
func (o *UniverseResp) SetPricePerHour(v float64) {
	o.PricePerHour = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *UniverseResp) GetResources() UniverseResourceDetails {
	if o == nil || o.Resources == nil {
		var ret UniverseResourceDetails
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetResourcesOk() (*UniverseResourceDetails, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *UniverseResp) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given UniverseResourceDetails and assigns it to the Resources field.
func (o *UniverseResp) SetResources(v UniverseResourceDetails) {
	o.Resources = &v
}

// GetSampleAppCommandTxt returns the SampleAppCommandTxt field value if set, zero value otherwise.
func (o *UniverseResp) GetSampleAppCommandTxt() string {
	if o == nil || o.SampleAppCommandTxt == nil {
		var ret string
		return ret
	}
	return *o.SampleAppCommandTxt
}

// GetSampleAppCommandTxtOk returns a tuple with the SampleAppCommandTxt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetSampleAppCommandTxtOk() (*string, bool) {
	if o == nil || o.SampleAppCommandTxt == nil {
		return nil, false
	}
	return o.SampleAppCommandTxt, true
}

// HasSampleAppCommandTxt returns a boolean if a field has been set.
func (o *UniverseResp) HasSampleAppCommandTxt() bool {
	if o != nil && o.SampleAppCommandTxt != nil {
		return true
	}

	return false
}

// SetSampleAppCommandTxt gets a reference to the given string and assigns it to the SampleAppCommandTxt field.
func (o *UniverseResp) SetSampleAppCommandTxt(v string) {
	o.SampleAppCommandTxt = &v
}

// GetTaskUUID returns the TaskUUID field value if set, zero value otherwise.
func (o *UniverseResp) GetTaskUUID() string {
	if o == nil || o.TaskUUID == nil {
		var ret string
		return ret
	}
	return *o.TaskUUID
}

// GetTaskUUIDOk returns a tuple with the TaskUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetTaskUUIDOk() (*string, bool) {
	if o == nil || o.TaskUUID == nil {
		return nil, false
	}
	return o.TaskUUID, true
}

// HasTaskUUID returns a boolean if a field has been set.
func (o *UniverseResp) HasTaskUUID() bool {
	if o != nil && o.TaskUUID != nil {
		return true
	}

	return false
}

// SetTaskUUID gets a reference to the given string and assigns it to the TaskUUID field.
func (o *UniverseResp) SetTaskUUID(v string) {
	o.TaskUUID = &v
}

// GetUniverseConfig returns the UniverseConfig field value if set, zero value otherwise.
func (o *UniverseResp) GetUniverseConfig() map[string]string {
	if o == nil || o.UniverseConfig == nil {
		var ret map[string]string
		return ret
	}
	return *o.UniverseConfig
}

// GetUniverseConfigOk returns a tuple with the UniverseConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetUniverseConfigOk() (*map[string]string, bool) {
	if o == nil || o.UniverseConfig == nil {
		return nil, false
	}
	return o.UniverseConfig, true
}

// HasUniverseConfig returns a boolean if a field has been set.
func (o *UniverseResp) HasUniverseConfig() bool {
	if o != nil && o.UniverseConfig != nil {
		return true
	}

	return false
}

// SetUniverseConfig gets a reference to the given map[string]string and assigns it to the UniverseConfig field.
func (o *UniverseResp) SetUniverseConfig(v map[string]string) {
	o.UniverseConfig = &v
}

// GetUniverseDetails returns the UniverseDetails field value if set, zero value otherwise.
func (o *UniverseResp) GetUniverseDetails() UniverseDefinitionTaskParamsResp {
	if o == nil || o.UniverseDetails == nil {
		var ret UniverseDefinitionTaskParamsResp
		return ret
	}
	return *o.UniverseDetails
}

// GetUniverseDetailsOk returns a tuple with the UniverseDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetUniverseDetailsOk() (*UniverseDefinitionTaskParamsResp, bool) {
	if o == nil || o.UniverseDetails == nil {
		return nil, false
	}
	return o.UniverseDetails, true
}

// HasUniverseDetails returns a boolean if a field has been set.
func (o *UniverseResp) HasUniverseDetails() bool {
	if o != nil && o.UniverseDetails != nil {
		return true
	}

	return false
}

// SetUniverseDetails gets a reference to the given UniverseDefinitionTaskParamsResp and assigns it to the UniverseDetails field.
func (o *UniverseResp) SetUniverseDetails(v UniverseDefinitionTaskParamsResp) {
	o.UniverseDetails = &v
}

// GetUniverseUUID returns the UniverseUUID field value if set, zero value otherwise.
func (o *UniverseResp) GetUniverseUUID() string {
	if o == nil || o.UniverseUUID == nil {
		var ret string
		return ret
	}
	return *o.UniverseUUID
}

// GetUniverseUUIDOk returns a tuple with the UniverseUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetUniverseUUIDOk() (*string, bool) {
	if o == nil || o.UniverseUUID == nil {
		return nil, false
	}
	return o.UniverseUUID, true
}

// HasUniverseUUID returns a boolean if a field has been set.
func (o *UniverseResp) HasUniverseUUID() bool {
	if o != nil && o.UniverseUUID != nil {
		return true
	}

	return false
}

// SetUniverseUUID gets a reference to the given string and assigns it to the UniverseUUID field.
func (o *UniverseResp) SetUniverseUUID(v string) {
	o.UniverseUUID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UniverseResp) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResp) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UniverseResp) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *UniverseResp) SetVersion(v int32) {
	o.Version = &v
}

func (o UniverseResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.DnsName != nil {
		toSerialize["dnsName"] = o.DnsName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PricePerHour != nil {
		toSerialize["pricePerHour"] = o.PricePerHour
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.SampleAppCommandTxt != nil {
		toSerialize["sampleAppCommandTxt"] = o.SampleAppCommandTxt
	}
	if o.TaskUUID != nil {
		toSerialize["taskUUID"] = o.TaskUUID
	}
	if o.UniverseConfig != nil {
		toSerialize["universeConfig"] = o.UniverseConfig
	}
	if o.UniverseDetails != nil {
		toSerialize["universeDetails"] = o.UniverseDetails
	}
	if o.UniverseUUID != nil {
		toSerialize["universeUUID"] = o.UniverseUUID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseResp struct {
	value *UniverseResp
	isSet bool
}

func (v NullableUniverseResp) Get() *UniverseResp {
	return v.value
}

func (v *NullableUniverseResp) Set(val *UniverseResp) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseResp) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseResp(val *UniverseResp) *NullableUniverseResp {
	return &NullableUniverseResp{value: val, isSet: true}
}

func (v NullableUniverseResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


