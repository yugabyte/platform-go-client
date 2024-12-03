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

// UniverseResourceDetails Universe resource details. Part of UniverseInfo.
type UniverseResourceDetails struct {
	// Azs
	AzList []string `json:"az_list"`
	// EBS price per hour
	EbsPricePerHour *float64 `json:"ebs_price_per_hour,omitempty"`
	// gp3 free piops
	Gp3FreePiops *int32 `json:"gp3_free_piops,omitempty"`
	// gp3 free throughput
	Gp3FreeThroughput *int32 `json:"gp3_free_throughput,omitempty"`
	// Memory GB
	MemSizeGb *float64 `json:"mem_size_gb,omitempty"`
	// Numbers of cores
	NumCores *float64 `json:"num_cores,omitempty"`
	// Numbers of node
	NumNodes *int32 `json:"num_nodes,omitempty"`
	// Price per hour
	PricePerHour *float64 `json:"price_per_hour,omitempty"`
	// Known pricing info
	PricingKnown *bool `json:"pricing_known,omitempty"`
	// Volume count
	VolumeCount *int32 `json:"volume_count,omitempty"`
	// Volume in GB
	VolumeSizeGb *int32 `json:"volume_size_gb,omitempty"`
}

// NewUniverseResourceDetails instantiates a new UniverseResourceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseResourceDetails(azList []string) *UniverseResourceDetails {
	this := UniverseResourceDetails{}
	this.AzList = azList
	return &this
}

// NewUniverseResourceDetailsWithDefaults instantiates a new UniverseResourceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseResourceDetailsWithDefaults() *UniverseResourceDetails {
	this := UniverseResourceDetails{}
	return &this
}

// GetAzList returns the AzList field value
func (o *UniverseResourceDetails) GetAzList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AzList
}

// GetAzListOk returns a tuple with the AzList field value
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetAzListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AzList, true
}

// SetAzList sets field value
func (o *UniverseResourceDetails) SetAzList(v []string) {
	o.AzList = v
}

// GetEbsPricePerHour returns the EbsPricePerHour field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetEbsPricePerHour() float64 {
	if o == nil || o.EbsPricePerHour == nil {
		var ret float64
		return ret
	}
	return *o.EbsPricePerHour
}

// GetEbsPricePerHourOk returns a tuple with the EbsPricePerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetEbsPricePerHourOk() (*float64, bool) {
	if o == nil || o.EbsPricePerHour == nil {
		return nil, false
	}
	return o.EbsPricePerHour, true
}

// HasEbsPricePerHour returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasEbsPricePerHour() bool {
	if o != nil && o.EbsPricePerHour != nil {
		return true
	}

	return false
}

// SetEbsPricePerHour gets a reference to the given float64 and assigns it to the EbsPricePerHour field.
func (o *UniverseResourceDetails) SetEbsPricePerHour(v float64) {
	o.EbsPricePerHour = &v
}

// GetGp3FreePiops returns the Gp3FreePiops field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetGp3FreePiops() int32 {
	if o == nil || o.Gp3FreePiops == nil {
		var ret int32
		return ret
	}
	return *o.Gp3FreePiops
}

// GetGp3FreePiopsOk returns a tuple with the Gp3FreePiops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetGp3FreePiopsOk() (*int32, bool) {
	if o == nil || o.Gp3FreePiops == nil {
		return nil, false
	}
	return o.Gp3FreePiops, true
}

// HasGp3FreePiops returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasGp3FreePiops() bool {
	if o != nil && o.Gp3FreePiops != nil {
		return true
	}

	return false
}

// SetGp3FreePiops gets a reference to the given int32 and assigns it to the Gp3FreePiops field.
func (o *UniverseResourceDetails) SetGp3FreePiops(v int32) {
	o.Gp3FreePiops = &v
}

// GetGp3FreeThroughput returns the Gp3FreeThroughput field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetGp3FreeThroughput() int32 {
	if o == nil || o.Gp3FreeThroughput == nil {
		var ret int32
		return ret
	}
	return *o.Gp3FreeThroughput
}

// GetGp3FreeThroughputOk returns a tuple with the Gp3FreeThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetGp3FreeThroughputOk() (*int32, bool) {
	if o == nil || o.Gp3FreeThroughput == nil {
		return nil, false
	}
	return o.Gp3FreeThroughput, true
}

// HasGp3FreeThroughput returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasGp3FreeThroughput() bool {
	if o != nil && o.Gp3FreeThroughput != nil {
		return true
	}

	return false
}

// SetGp3FreeThroughput gets a reference to the given int32 and assigns it to the Gp3FreeThroughput field.
func (o *UniverseResourceDetails) SetGp3FreeThroughput(v int32) {
	o.Gp3FreeThroughput = &v
}

// GetMemSizeGb returns the MemSizeGb field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetMemSizeGb() float64 {
	if o == nil || o.MemSizeGb == nil {
		var ret float64
		return ret
	}
	return *o.MemSizeGb
}

// GetMemSizeGbOk returns a tuple with the MemSizeGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetMemSizeGbOk() (*float64, bool) {
	if o == nil || o.MemSizeGb == nil {
		return nil, false
	}
	return o.MemSizeGb, true
}

// HasMemSizeGb returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasMemSizeGb() bool {
	if o != nil && o.MemSizeGb != nil {
		return true
	}

	return false
}

// SetMemSizeGb gets a reference to the given float64 and assigns it to the MemSizeGb field.
func (o *UniverseResourceDetails) SetMemSizeGb(v float64) {
	o.MemSizeGb = &v
}

// GetNumCores returns the NumCores field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetNumCores() float64 {
	if o == nil || o.NumCores == nil {
		var ret float64
		return ret
	}
	return *o.NumCores
}

// GetNumCoresOk returns a tuple with the NumCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetNumCoresOk() (*float64, bool) {
	if o == nil || o.NumCores == nil {
		return nil, false
	}
	return o.NumCores, true
}

// HasNumCores returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasNumCores() bool {
	if o != nil && o.NumCores != nil {
		return true
	}

	return false
}

// SetNumCores gets a reference to the given float64 and assigns it to the NumCores field.
func (o *UniverseResourceDetails) SetNumCores(v float64) {
	o.NumCores = &v
}

// GetNumNodes returns the NumNodes field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetNumNodes() int32 {
	if o == nil || o.NumNodes == nil {
		var ret int32
		return ret
	}
	return *o.NumNodes
}

// GetNumNodesOk returns a tuple with the NumNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetNumNodesOk() (*int32, bool) {
	if o == nil || o.NumNodes == nil {
		return nil, false
	}
	return o.NumNodes, true
}

// HasNumNodes returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasNumNodes() bool {
	if o != nil && o.NumNodes != nil {
		return true
	}

	return false
}

// SetNumNodes gets a reference to the given int32 and assigns it to the NumNodes field.
func (o *UniverseResourceDetails) SetNumNodes(v int32) {
	o.NumNodes = &v
}

// GetPricePerHour returns the PricePerHour field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetPricePerHour() float64 {
	if o == nil || o.PricePerHour == nil {
		var ret float64
		return ret
	}
	return *o.PricePerHour
}

// GetPricePerHourOk returns a tuple with the PricePerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetPricePerHourOk() (*float64, bool) {
	if o == nil || o.PricePerHour == nil {
		return nil, false
	}
	return o.PricePerHour, true
}

// HasPricePerHour returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasPricePerHour() bool {
	if o != nil && o.PricePerHour != nil {
		return true
	}

	return false
}

// SetPricePerHour gets a reference to the given float64 and assigns it to the PricePerHour field.
func (o *UniverseResourceDetails) SetPricePerHour(v float64) {
	o.PricePerHour = &v
}

// GetPricingKnown returns the PricingKnown field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetPricingKnown() bool {
	if o == nil || o.PricingKnown == nil {
		var ret bool
		return ret
	}
	return *o.PricingKnown
}

// GetPricingKnownOk returns a tuple with the PricingKnown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetPricingKnownOk() (*bool, bool) {
	if o == nil || o.PricingKnown == nil {
		return nil, false
	}
	return o.PricingKnown, true
}

// HasPricingKnown returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasPricingKnown() bool {
	if o != nil && o.PricingKnown != nil {
		return true
	}

	return false
}

// SetPricingKnown gets a reference to the given bool and assigns it to the PricingKnown field.
func (o *UniverseResourceDetails) SetPricingKnown(v bool) {
	o.PricingKnown = &v
}

// GetVolumeCount returns the VolumeCount field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetVolumeCount() int32 {
	if o == nil || o.VolumeCount == nil {
		var ret int32
		return ret
	}
	return *o.VolumeCount
}

// GetVolumeCountOk returns a tuple with the VolumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetVolumeCountOk() (*int32, bool) {
	if o == nil || o.VolumeCount == nil {
		return nil, false
	}
	return o.VolumeCount, true
}

// HasVolumeCount returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasVolumeCount() bool {
	if o != nil && o.VolumeCount != nil {
		return true
	}

	return false
}

// SetVolumeCount gets a reference to the given int32 and assigns it to the VolumeCount field.
func (o *UniverseResourceDetails) SetVolumeCount(v int32) {
	o.VolumeCount = &v
}

// GetVolumeSizeGb returns the VolumeSizeGb field value if set, zero value otherwise.
func (o *UniverseResourceDetails) GetVolumeSizeGb() int32 {
	if o == nil || o.VolumeSizeGb == nil {
		var ret int32
		return ret
	}
	return *o.VolumeSizeGb
}

// GetVolumeSizeGbOk returns a tuple with the VolumeSizeGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseResourceDetails) GetVolumeSizeGbOk() (*int32, bool) {
	if o == nil || o.VolumeSizeGb == nil {
		return nil, false
	}
	return o.VolumeSizeGb, true
}

// HasVolumeSizeGb returns a boolean if a field has been set.
func (o *UniverseResourceDetails) HasVolumeSizeGb() bool {
	if o != nil && o.VolumeSizeGb != nil {
		return true
	}

	return false
}

// SetVolumeSizeGb gets a reference to the given int32 and assigns it to the VolumeSizeGb field.
func (o *UniverseResourceDetails) SetVolumeSizeGb(v int32) {
	o.VolumeSizeGb = &v
}

func (o UniverseResourceDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["az_list"] = o.AzList
	}
	if o.EbsPricePerHour != nil {
		toSerialize["ebs_price_per_hour"] = o.EbsPricePerHour
	}
	if o.Gp3FreePiops != nil {
		toSerialize["gp3_free_piops"] = o.Gp3FreePiops
	}
	if o.Gp3FreeThroughput != nil {
		toSerialize["gp3_free_throughput"] = o.Gp3FreeThroughput
	}
	if o.MemSizeGb != nil {
		toSerialize["mem_size_gb"] = o.MemSizeGb
	}
	if o.NumCores != nil {
		toSerialize["num_cores"] = o.NumCores
	}
	if o.NumNodes != nil {
		toSerialize["num_nodes"] = o.NumNodes
	}
	if o.PricePerHour != nil {
		toSerialize["price_per_hour"] = o.PricePerHour
	}
	if o.PricingKnown != nil {
		toSerialize["pricing_known"] = o.PricingKnown
	}
	if o.VolumeCount != nil {
		toSerialize["volume_count"] = o.VolumeCount
	}
	if o.VolumeSizeGb != nil {
		toSerialize["volume_size_gb"] = o.VolumeSizeGb
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseResourceDetails struct {
	value *UniverseResourceDetails
	isSet bool
}

func (v NullableUniverseResourceDetails) Get() *UniverseResourceDetails {
	return v.value
}

func (v *NullableUniverseResourceDetails) Set(val *UniverseResourceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseResourceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseResourceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseResourceDetails(val *UniverseResourceDetails) *NullableUniverseResourceDetails {
	return &NullableUniverseResourceDetails{value: val, isSet: true}
}

func (v NullableUniverseResourceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseResourceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

