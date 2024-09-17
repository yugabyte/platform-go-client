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

// RegionFormData struct for RegionFormData
type RegionFormData struct {
	Code string `json:"code"`
	DestVpcId string `json:"destVpcId"`
	HostVpcId string `json:"hostVpcId"`
	HostVpcRegion string `json:"hostVpcRegion"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name string `json:"name"`
	SecurityGroupId string `json:"securityGroupId"`
	VnetName string `json:"vnetName"`
	YbImage string `json:"ybImage"`
}

// NewRegionFormData instantiates a new RegionFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionFormData(code string, destVpcId string, hostVpcId string, hostVpcRegion string, latitude float64, longitude float64, name string, securityGroupId string, vnetName string, ybImage string) *RegionFormData {
	this := RegionFormData{}
	this.Code = code
	this.DestVpcId = destVpcId
	this.HostVpcId = hostVpcId
	this.HostVpcRegion = hostVpcRegion
	this.Latitude = latitude
	this.Longitude = longitude
	this.Name = name
	this.SecurityGroupId = securityGroupId
	this.VnetName = vnetName
	this.YbImage = ybImage
	return &this
}

// NewRegionFormDataWithDefaults instantiates a new RegionFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionFormDataWithDefaults() *RegionFormData {
	this := RegionFormData{}
	return &this
}

// GetCode returns the Code field value
func (o *RegionFormData) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *RegionFormData) SetCode(v string) {
	o.Code = v
}

// GetDestVpcId returns the DestVpcId field value
func (o *RegionFormData) GetDestVpcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestVpcId
}

// GetDestVpcIdOk returns a tuple with the DestVpcId field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetDestVpcIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestVpcId, true
}

// SetDestVpcId sets field value
func (o *RegionFormData) SetDestVpcId(v string) {
	o.DestVpcId = v
}

// GetHostVpcId returns the HostVpcId field value
func (o *RegionFormData) GetHostVpcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostVpcId
}

// GetHostVpcIdOk returns a tuple with the HostVpcId field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetHostVpcIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HostVpcId, true
}

// SetHostVpcId sets field value
func (o *RegionFormData) SetHostVpcId(v string) {
	o.HostVpcId = v
}

// GetHostVpcRegion returns the HostVpcRegion field value
func (o *RegionFormData) GetHostVpcRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostVpcRegion
}

// GetHostVpcRegionOk returns a tuple with the HostVpcRegion field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetHostVpcRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HostVpcRegion, true
}

// SetHostVpcRegion sets field value
func (o *RegionFormData) SetHostVpcRegion(v string) {
	o.HostVpcRegion = v
}

// GetLatitude returns the Latitude field value
func (o *RegionFormData) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetLatitudeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *RegionFormData) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *RegionFormData) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetLongitudeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *RegionFormData) SetLongitude(v float64) {
	o.Longitude = v
}

// GetName returns the Name field value
func (o *RegionFormData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RegionFormData) SetName(v string) {
	o.Name = v
}

// GetSecurityGroupId returns the SecurityGroupId field value
func (o *RegionFormData) GetSecurityGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecurityGroupId, true
}

// SetSecurityGroupId sets field value
func (o *RegionFormData) SetSecurityGroupId(v string) {
	o.SecurityGroupId = v
}

// GetVnetName returns the VnetName field value
func (o *RegionFormData) GetVnetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VnetName
}

// GetVnetNameOk returns a tuple with the VnetName field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetVnetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VnetName, true
}

// SetVnetName sets field value
func (o *RegionFormData) SetVnetName(v string) {
	o.VnetName = v
}

// GetYbImage returns the YbImage field value
func (o *RegionFormData) GetYbImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbImage
}

// GetYbImageOk returns a tuple with the YbImage field value
// and a boolean to check if the value has been set.
func (o *RegionFormData) GetYbImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbImage, true
}

// SetYbImage sets field value
func (o *RegionFormData) SetYbImage(v string) {
	o.YbImage = v
}

func (o RegionFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["destVpcId"] = o.DestVpcId
	}
	if true {
		toSerialize["hostVpcId"] = o.HostVpcId
	}
	if true {
		toSerialize["hostVpcRegion"] = o.HostVpcRegion
	}
	if true {
		toSerialize["latitude"] = o.Latitude
	}
	if true {
		toSerialize["longitude"] = o.Longitude
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["securityGroupId"] = o.SecurityGroupId
	}
	if true {
		toSerialize["vnetName"] = o.VnetName
	}
	if true {
		toSerialize["ybImage"] = o.YbImage
	}
	return json.Marshal(toSerialize)
}

type NullableRegionFormData struct {
	value *RegionFormData
	isSet bool
}

func (v NullableRegionFormData) Get() *RegionFormData {
	return v.value
}

func (v *NullableRegionFormData) Set(val *RegionFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionFormData(val *RegionFormData) *NullableRegionFormData {
	return &NullableRegionFormData{value: val, isSet: true}
}

func (v NullableRegionFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


