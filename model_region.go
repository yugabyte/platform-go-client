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

// Region Region within a given provider. Typically, this maps to a single cloud provider region.
type Region struct {
	Active *bool `json:"active,omitempty"`
	// Cloud provider region code
	Code *string `json:"code,omitempty"`
	Config *map[string]string `json:"config,omitempty"`
	Details *RegionDetails `json:"details,omitempty"`
	// The region's latitude
	Latitude *float64 `json:"latitude,omitempty"`
	// The region's longitude
	Longitude *float64 `json:"longitude,omitempty"`
	// Cloud provider region name
	Name *string `json:"name,omitempty"`
	// <b style=\"color:#ff0000\">Deprecated since YBA version 2.17.2.0.</b> Moved to regionDetails.cloudInfo aws/azure securityGroupId property
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
	// Region UUID
	Uuid *string `json:"uuid,omitempty"`
	// <b style=\"color:#ff0000\">Deprecated since YBA version 2.17.2.0.</b> Moved to regionDetails.cloudInfo aws/azure vnet property
	VnetName *string `json:"vnetName,omitempty"`
	// <b style=\"color:#ff0000\">Deprecated since YBA version 2.17.2.0.</b> Moved to details.cloudInfo aws/gcp/azure ybImage property
	YbImage *string `json:"ybImage,omitempty"`
	Zones []AvailabilityZone `json:"zones"`
}

// NewRegion instantiates a new Region object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegion(zones []AvailabilityZone) *Region {
	this := Region{}
	this.Zones = zones
	return &this
}

// NewRegionWithDefaults instantiates a new Region object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionWithDefaults() *Region {
	this := Region{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Region) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Region) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Region) SetActive(v bool) {
	o.Active = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Region) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Region) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Region) SetCode(v string) {
	o.Code = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Region) GetConfig() map[string]string {
	if o == nil || o.Config == nil {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetConfigOk() (*map[string]string, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Region) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *Region) SetConfig(v map[string]string) {
	o.Config = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Region) GetDetails() RegionDetails {
	if o == nil || o.Details == nil {
		var ret RegionDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetDetailsOk() (*RegionDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Region) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given RegionDetails and assigns it to the Details field.
func (o *Region) SetDetails(v RegionDetails) {
	o.Details = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *Region) GetLatitude() float64 {
	if o == nil || o.Latitude == nil {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetLatitudeOk() (*float64, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *Region) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *Region) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *Region) GetLongitude() float64 {
	if o == nil || o.Longitude == nil {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetLongitudeOk() (*float64, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *Region) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *Region) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Region) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Region) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Region) SetName(v string) {
	o.Name = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *Region) GetSecurityGroupId() string {
	if o == nil || o.SecurityGroupId == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || o.SecurityGroupId == nil {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *Region) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId != nil {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *Region) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Region) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Region) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Region) SetUuid(v string) {
	o.Uuid = &v
}

// GetVnetName returns the VnetName field value if set, zero value otherwise.
func (o *Region) GetVnetName() string {
	if o == nil || o.VnetName == nil {
		var ret string
		return ret
	}
	return *o.VnetName
}

// GetVnetNameOk returns a tuple with the VnetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetVnetNameOk() (*string, bool) {
	if o == nil || o.VnetName == nil {
		return nil, false
	}
	return o.VnetName, true
}

// HasVnetName returns a boolean if a field has been set.
func (o *Region) HasVnetName() bool {
	if o != nil && o.VnetName != nil {
		return true
	}

	return false
}

// SetVnetName gets a reference to the given string and assigns it to the VnetName field.
func (o *Region) SetVnetName(v string) {
	o.VnetName = &v
}

// GetYbImage returns the YbImage field value if set, zero value otherwise.
func (o *Region) GetYbImage() string {
	if o == nil || o.YbImage == nil {
		var ret string
		return ret
	}
	return *o.YbImage
}

// GetYbImageOk returns a tuple with the YbImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetYbImageOk() (*string, bool) {
	if o == nil || o.YbImage == nil {
		return nil, false
	}
	return o.YbImage, true
}

// HasYbImage returns a boolean if a field has been set.
func (o *Region) HasYbImage() bool {
	if o != nil && o.YbImage != nil {
		return true
	}

	return false
}

// SetYbImage gets a reference to the given string and assigns it to the YbImage field.
func (o *Region) SetYbImage(v string) {
	o.YbImage = &v
}

// GetZones returns the Zones field value
func (o *Region) GetZones() []AvailabilityZone {
	if o == nil {
		var ret []AvailabilityZone
		return ret
	}

	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value
// and a boolean to check if the value has been set.
func (o *Region) GetZonesOk() (*[]AvailabilityZone, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Zones, true
}

// SetZones sets field value
func (o *Region) SetZones(v []AvailabilityZone) {
	o.Zones = v
}

func (o Region) MarshalJSON() ([]byte, error) {
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
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SecurityGroupId != nil {
		toSerialize["securityGroupId"] = o.SecurityGroupId
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.VnetName != nil {
		toSerialize["vnetName"] = o.VnetName
	}
	if o.YbImage != nil {
		toSerialize["ybImage"] = o.YbImage
	}
	if true {
		toSerialize["zones"] = o.Zones
	}
	return json.Marshal(toSerialize)
}

type NullableRegion struct {
	value *Region
	isSet bool
}

func (v NullableRegion) Get() *Region {
	return v.value
}

func (v *NullableRegion) Set(val *Region) {
	v.value = val
	v.isSet = true
}

func (v NullableRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegion(val *Region) *NullableRegion {
	return &NullableRegion{value: val, isSet: true}
}

func (v NullableRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


