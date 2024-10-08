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

// ResponseExtractMetadata struct for ResponseExtractMetadata
type ResponseExtractMetadata struct {
	Architecture string `json:"architecture"`
	MetadataUuid string `json:"metadata_uuid"`
	Platform string `json:"platform"`
	ReleaseDateMsecs int64 `json:"release_date_msecs"`
	ReleaseNotes string `json:"release_notes"`
	ReleaseType string `json:"release_type"`
	Sha256 string `json:"sha256"`
	Status string `json:"status"`
	Version string `json:"version"`
	YbType string `json:"yb_type"`
}

// NewResponseExtractMetadata instantiates a new ResponseExtractMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseExtractMetadata(architecture string, metadataUuid string, platform string, releaseDateMsecs int64, releaseNotes string, releaseType string, sha256 string, status string, version string, ybType string) *ResponseExtractMetadata {
	this := ResponseExtractMetadata{}
	this.Architecture = architecture
	this.MetadataUuid = metadataUuid
	this.Platform = platform
	this.ReleaseDateMsecs = releaseDateMsecs
	this.ReleaseNotes = releaseNotes
	this.ReleaseType = releaseType
	this.Sha256 = sha256
	this.Status = status
	this.Version = version
	this.YbType = ybType
	return &this
}

// NewResponseExtractMetadataWithDefaults instantiates a new ResponseExtractMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseExtractMetadataWithDefaults() *ResponseExtractMetadata {
	this := ResponseExtractMetadata{}
	return &this
}

// GetArchitecture returns the Architecture field value
func (o *ResponseExtractMetadata) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetArchitectureOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *ResponseExtractMetadata) SetArchitecture(v string) {
	o.Architecture = v
}

// GetMetadataUuid returns the MetadataUuid field value
func (o *ResponseExtractMetadata) GetMetadataUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataUuid
}

// GetMetadataUuidOk returns a tuple with the MetadataUuid field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetMetadataUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetadataUuid, true
}

// SetMetadataUuid sets field value
func (o *ResponseExtractMetadata) SetMetadataUuid(v string) {
	o.MetadataUuid = v
}

// GetPlatform returns the Platform field value
func (o *ResponseExtractMetadata) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetPlatformOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *ResponseExtractMetadata) SetPlatform(v string) {
	o.Platform = v
}

// GetReleaseDateMsecs returns the ReleaseDateMsecs field value
func (o *ResponseExtractMetadata) GetReleaseDateMsecs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReleaseDateMsecs
}

// GetReleaseDateMsecsOk returns a tuple with the ReleaseDateMsecs field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetReleaseDateMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseDateMsecs, true
}

// SetReleaseDateMsecs sets field value
func (o *ResponseExtractMetadata) SetReleaseDateMsecs(v int64) {
	o.ReleaseDateMsecs = v
}

// GetReleaseNotes returns the ReleaseNotes field value
func (o *ResponseExtractMetadata) GetReleaseNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseNotes
}

// GetReleaseNotesOk returns a tuple with the ReleaseNotes field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetReleaseNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseNotes, true
}

// SetReleaseNotes sets field value
func (o *ResponseExtractMetadata) SetReleaseNotes(v string) {
	o.ReleaseNotes = v
}

// GetReleaseType returns the ReleaseType field value
func (o *ResponseExtractMetadata) GetReleaseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetReleaseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseType, true
}

// SetReleaseType sets field value
func (o *ResponseExtractMetadata) SetReleaseType(v string) {
	o.ReleaseType = v
}

// GetSha256 returns the Sha256 field value
func (o *ResponseExtractMetadata) GetSha256() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetSha256Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sha256, true
}

// SetSha256 sets field value
func (o *ResponseExtractMetadata) SetSha256(v string) {
	o.Sha256 = v
}

// GetStatus returns the Status field value
func (o *ResponseExtractMetadata) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResponseExtractMetadata) SetStatus(v string) {
	o.Status = v
}

// GetVersion returns the Version field value
func (o *ResponseExtractMetadata) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ResponseExtractMetadata) SetVersion(v string) {
	o.Version = v
}

// GetYbType returns the YbType field value
func (o *ResponseExtractMetadata) GetYbType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbType
}

// GetYbTypeOk returns a tuple with the YbType field value
// and a boolean to check if the value has been set.
func (o *ResponseExtractMetadata) GetYbTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbType, true
}

// SetYbType sets field value
func (o *ResponseExtractMetadata) SetYbType(v string) {
	o.YbType = v
}

func (o ResponseExtractMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["architecture"] = o.Architecture
	}
	if true {
		toSerialize["metadata_uuid"] = o.MetadataUuid
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["release_date_msecs"] = o.ReleaseDateMsecs
	}
	if true {
		toSerialize["release_notes"] = o.ReleaseNotes
	}
	if true {
		toSerialize["release_type"] = o.ReleaseType
	}
	if true {
		toSerialize["sha256"] = o.Sha256
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["yb_type"] = o.YbType
	}
	return json.Marshal(toSerialize)
}

type NullableResponseExtractMetadata struct {
	value *ResponseExtractMetadata
	isSet bool
}

func (v NullableResponseExtractMetadata) Get() *ResponseExtractMetadata {
	return v.value
}

func (v *NullableResponseExtractMetadata) Set(val *ResponseExtractMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseExtractMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseExtractMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseExtractMetadata(val *ResponseExtractMetadata) *NullableResponseExtractMetadata {
	return &NullableResponseExtractMetadata{value: val, isSet: true}
}

func (v NullableResponseExtractMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseExtractMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


