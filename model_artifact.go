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

// Artifact struct for Artifact
type Artifact struct {
	Architecture string `json:"architecture"`
	PackageFileId string `json:"package_file_id"`
	PackageUrl string `json:"package_url"`
	Platform string `json:"platform"`
	Sha256 string `json:"sha256"`
}

// NewArtifact instantiates a new Artifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifact(architecture string, packageFileId string, packageUrl string, platform string, sha256 string) *Artifact {
	this := Artifact{}
	this.Architecture = architecture
	this.PackageFileId = packageFileId
	this.PackageUrl = packageUrl
	this.Platform = platform
	this.Sha256 = sha256
	return &this
}

// NewArtifactWithDefaults instantiates a new Artifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactWithDefaults() *Artifact {
	this := Artifact{}
	return &this
}

// GetArchitecture returns the Architecture field value
func (o *Artifact) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetArchitectureOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *Artifact) SetArchitecture(v string) {
	o.Architecture = v
}

// GetPackageFileId returns the PackageFileId field value
func (o *Artifact) GetPackageFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFileId
}

// GetPackageFileIdOk returns a tuple with the PackageFileId field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetPackageFileIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PackageFileId, true
}

// SetPackageFileId sets field value
func (o *Artifact) SetPackageFileId(v string) {
	o.PackageFileId = v
}

// GetPackageUrl returns the PackageUrl field value
func (o *Artifact) GetPackageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetPackageUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PackageUrl, true
}

// SetPackageUrl sets field value
func (o *Artifact) SetPackageUrl(v string) {
	o.PackageUrl = v
}

// GetPlatform returns the Platform field value
func (o *Artifact) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetPlatformOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *Artifact) SetPlatform(v string) {
	o.Platform = v
}

// GetSha256 returns the Sha256 field value
func (o *Artifact) GetSha256() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetSha256Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sha256, true
}

// SetSha256 sets field value
func (o *Artifact) SetSha256(v string) {
	o.Sha256 = v
}

func (o Artifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["architecture"] = o.Architecture
	}
	if true {
		toSerialize["package_file_id"] = o.PackageFileId
	}
	if true {
		toSerialize["package_url"] = o.PackageUrl
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["sha256"] = o.Sha256
	}
	return json.Marshal(toSerialize)
}

type NullableArtifact struct {
	value *Artifact
	isSet bool
}

func (v NullableArtifact) Get() *Artifact {
	return v.value
}

func (v *NullableArtifact) Set(val *Artifact) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifact(val *Artifact) *NullableArtifact {
	return &NullableArtifact{value: val, isSet: true}
}

func (v NullableArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

