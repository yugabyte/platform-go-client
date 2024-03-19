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

// AWSCloudInfo struct for AWSCloudInfo
type AWSCloudInfo struct {
	AwsAccessKeyID *string `json:"awsAccessKeyID,omitempty"`
	AwsAccessKeySecret *string `json:"awsAccessKeySecret,omitempty"`
	AwsHostedZoneId *string `json:"awsHostedZoneId,omitempty"`
	AwsHostedZoneName *string `json:"awsHostedZoneName,omitempty"`
	HostVpcId *string `json:"hostVpcId,omitempty"`
	HostVpcRegion *string `json:"hostVpcRegion,omitempty"`
	UseIMDSv2 *bool `json:"useIMDSv2,omitempty"`
	// New/Existing VPC for provider creation
	VpcType *string `json:"vpcType,omitempty"`
}

// NewAWSCloudInfo instantiates a new AWSCloudInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSCloudInfo() *AWSCloudInfo {
	this := AWSCloudInfo{}
	return &this
}

// NewAWSCloudInfoWithDefaults instantiates a new AWSCloudInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSCloudInfoWithDefaults() *AWSCloudInfo {
	this := AWSCloudInfo{}
	return &this
}

// GetAwsAccessKeyID returns the AwsAccessKeyID field value if set, zero value otherwise.
func (o *AWSCloudInfo) GetAwsAccessKeyID() string {
	if o == nil || o.AwsAccessKeyID == nil {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyID
}

// GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudInfo) GetAwsAccessKeyIDOk() (*string, bool) {
	if o == nil || o.AwsAccessKeyID == nil {
		return nil, false
	}
	return o.AwsAccessKeyID, true
}

// HasAwsAccessKeyID returns a boolean if a field has been set.
func (o *AWSCloudInfo) HasAwsAccessKeyID() bool {
	if o != nil && o.AwsAccessKeyID != nil {
		return true
	}

	return false
}

// SetAwsAccessKeyID gets a reference to the given string and assigns it to the AwsAccessKeyID field.
func (o *AWSCloudInfo) SetAwsAccessKeyID(v string) {
	o.AwsAccessKeyID = &v
}

// GetAwsAccessKeySecret returns the AwsAccessKeySecret field value if set, zero value otherwise.
func (o *AWSCloudInfo) GetAwsAccessKeySecret() string {
	if o == nil || o.AwsAccessKeySecret == nil {
		var ret string
		return ret
	}
	return *o.AwsAccessKeySecret
}

// GetAwsAccessKeySecretOk returns a tuple with the AwsAccessKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudInfo) GetAwsAccessKeySecretOk() (*string, bool) {
	if o == nil || o.AwsAccessKeySecret == nil {
		return nil, false
	}
	return o.AwsAccessKeySecret, true
}

// HasAwsAccessKeySecret returns a boolean if a field has been set.
func (o *AWSCloudInfo) HasAwsAccessKeySecret() bool {
	if o != nil && o.AwsAccessKeySecret != nil {
		return true
	}

	return false
}

// SetAwsAccessKeySecret gets a reference to the given string and assigns it to the AwsAccessKeySecret field.
func (o *AWSCloudInfo) SetAwsAccessKeySecret(v string) {
	o.AwsAccessKeySecret = &v
}

// GetAwsHostedZoneId returns the AwsHostedZoneId field value if set, zero value otherwise.
func (o *AWSCloudInfo) GetAwsHostedZoneId() string {
	if o == nil || o.AwsHostedZoneId == nil {
		var ret string
		return ret
	}
	return *o.AwsHostedZoneId
}

// GetAwsHostedZoneIdOk returns a tuple with the AwsHostedZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudInfo) GetAwsHostedZoneIdOk() (*string, bool) {
	if o == nil || o.AwsHostedZoneId == nil {
		return nil, false
	}
	return o.AwsHostedZoneId, true
}

// HasAwsHostedZoneId returns a boolean if a field has been set.
func (o *AWSCloudInfo) HasAwsHostedZoneId() bool {
	if o != nil && o.AwsHostedZoneId != nil {
		return true
	}

	return false
}

// SetAwsHostedZoneId gets a reference to the given string and assigns it to the AwsHostedZoneId field.
func (o *AWSCloudInfo) SetAwsHostedZoneId(v string) {
	o.AwsHostedZoneId = &v
}

// GetAwsHostedZoneName returns the AwsHostedZoneName field value if set, zero value otherwise.
func (o *AWSCloudInfo) GetAwsHostedZoneName() string {
	if o == nil || o.AwsHostedZoneName == nil {
		var ret string
		return ret
	}
	return *o.AwsHostedZoneName
}

// GetAwsHostedZoneNameOk returns a tuple with the AwsHostedZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudInfo) GetAwsHostedZoneNameOk() (*string, bool) {
	if o == nil || o.AwsHostedZoneName == nil {
		return nil, false
	}
	return o.AwsHostedZoneName, true
}

// HasAwsHostedZoneName returns a boolean if a field has been set.
func (o *AWSCloudInfo) HasAwsHostedZoneName() bool {
	if o != nil && o.AwsHostedZoneName != nil {
		return true
	}

	return false
}

// SetAwsHostedZoneName gets a reference to the given string and assigns it to the AwsHostedZoneName field.
func (o *AWSCloudInfo) SetAwsHostedZoneName(v string) {
	o.AwsHostedZoneName = &v
}

// GetHostVpcId returns the HostVpcId field value if set, zero value otherwise.
func (o *AWSCloudInfo) GetHostVpcId() string {
	if o == nil || o.HostVpcId == nil {
		var ret string
		return ret
	}
	return *o.HostVpcId
}

// GetHostVpcIdOk returns a tuple with the HostVpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudInfo) GetHostVpcIdOk() (*string, bool) {
	if o == nil || o.HostVpcId == nil {
		return nil, false
	}
	return o.HostVpcId, true
}

// HasHostVpcId returns a boolean if a field has been set.
func (o *AWSCloudInfo) HasHostVpcId() bool {
	if o != nil && o.HostVpcId != nil {
		return true
	}

	return false
}

// SetHostVpcId gets a reference to the given string and assigns it to the HostVpcId field.
func (o *AWSCloudInfo) SetHostVpcId(v string) {
	o.HostVpcId = &v
}

// GetHostVpcRegion returns the HostVpcRegion field value if set, zero value otherwise.
func (o *AWSCloudInfo) GetHostVpcRegion() string {
	if o == nil || o.HostVpcRegion == nil {
		var ret string
		return ret
	}
	return *o.HostVpcRegion
}

// GetHostVpcRegionOk returns a tuple with the HostVpcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudInfo) GetHostVpcRegionOk() (*string, bool) {
	if o == nil || o.HostVpcRegion == nil {
		return nil, false
	}
	return o.HostVpcRegion, true
}

// HasHostVpcRegion returns a boolean if a field has been set.
func (o *AWSCloudInfo) HasHostVpcRegion() bool {
	if o != nil && o.HostVpcRegion != nil {
		return true
	}

	return false
}

// SetHostVpcRegion gets a reference to the given string and assigns it to the HostVpcRegion field.
func (o *AWSCloudInfo) SetHostVpcRegion(v string) {
	o.HostVpcRegion = &v
}

// GetUseIMDSv2 returns the UseIMDSv2 field value if set, zero value otherwise.
func (o *AWSCloudInfo) GetUseIMDSv2() bool {
	if o == nil || o.UseIMDSv2 == nil {
		var ret bool
		return ret
	}
	return *o.UseIMDSv2
}

// GetUseIMDSv2Ok returns a tuple with the UseIMDSv2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudInfo) GetUseIMDSv2Ok() (*bool, bool) {
	if o == nil || o.UseIMDSv2 == nil {
		return nil, false
	}
	return o.UseIMDSv2, true
}

// HasUseIMDSv2 returns a boolean if a field has been set.
func (o *AWSCloudInfo) HasUseIMDSv2() bool {
	if o != nil && o.UseIMDSv2 != nil {
		return true
	}

	return false
}

// SetUseIMDSv2 gets a reference to the given bool and assigns it to the UseIMDSv2 field.
func (o *AWSCloudInfo) SetUseIMDSv2(v bool) {
	o.UseIMDSv2 = &v
}

// GetVpcType returns the VpcType field value if set, zero value otherwise.
func (o *AWSCloudInfo) GetVpcType() string {
	if o == nil || o.VpcType == nil {
		var ret string
		return ret
	}
	return *o.VpcType
}

// GetVpcTypeOk returns a tuple with the VpcType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudInfo) GetVpcTypeOk() (*string, bool) {
	if o == nil || o.VpcType == nil {
		return nil, false
	}
	return o.VpcType, true
}

// HasVpcType returns a boolean if a field has been set.
func (o *AWSCloudInfo) HasVpcType() bool {
	if o != nil && o.VpcType != nil {
		return true
	}

	return false
}

// SetVpcType gets a reference to the given string and assigns it to the VpcType field.
func (o *AWSCloudInfo) SetVpcType(v string) {
	o.VpcType = &v
}

func (o AWSCloudInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsAccessKeyID != nil {
		toSerialize["awsAccessKeyID"] = o.AwsAccessKeyID
	}
	if o.AwsAccessKeySecret != nil {
		toSerialize["awsAccessKeySecret"] = o.AwsAccessKeySecret
	}
	if o.AwsHostedZoneId != nil {
		toSerialize["awsHostedZoneId"] = o.AwsHostedZoneId
	}
	if o.AwsHostedZoneName != nil {
		toSerialize["awsHostedZoneName"] = o.AwsHostedZoneName
	}
	if o.HostVpcId != nil {
		toSerialize["hostVpcId"] = o.HostVpcId
	}
	if o.HostVpcRegion != nil {
		toSerialize["hostVpcRegion"] = o.HostVpcRegion
	}
	if o.UseIMDSv2 != nil {
		toSerialize["useIMDSv2"] = o.UseIMDSv2
	}
	if o.VpcType != nil {
		toSerialize["vpcType"] = o.VpcType
	}
	return json.Marshal(toSerialize)
}

type NullableAWSCloudInfo struct {
	value *AWSCloudInfo
	isSet bool
}

func (v NullableAWSCloudInfo) Get() *AWSCloudInfo {
	return v.value
}

func (v *NullableAWSCloudInfo) Set(val *AWSCloudInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSCloudInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSCloudInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSCloudInfo(val *AWSCloudInfo) *NullableAWSCloudInfo {
	return &NullableAWSCloudInfo{value: val, isSet: true}
}

func (v NullableAWSCloudInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSCloudInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


