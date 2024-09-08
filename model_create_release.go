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

// CreateRelease Release metadata required to create a new release
type CreateRelease struct {
	Artifacts []Artifact `json:"artifacts"`
	ReleaseDateMsecs int64 `json:"release_date_msecs"`
	ReleaseNotes string `json:"release_notes"`
	ReleaseTag string `json:"release_tag"`
	ReleaseType string `json:"release_type"`
	ReleaseUuid string `json:"release_uuid"`
	Version string `json:"version"`
	YbType string `json:"yb_type"`
}

// NewCreateRelease instantiates a new CreateRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRelease(artifacts []Artifact, releaseDateMsecs int64, releaseNotes string, releaseTag string, releaseType string, releaseUuid string, version string, ybType string) *CreateRelease {
	this := CreateRelease{}
	this.Artifacts = artifacts
	this.ReleaseDateMsecs = releaseDateMsecs
	this.ReleaseNotes = releaseNotes
	this.ReleaseTag = releaseTag
	this.ReleaseType = releaseType
	this.ReleaseUuid = releaseUuid
	this.Version = version
	this.YbType = ybType
	return &this
}

// NewCreateReleaseWithDefaults instantiates a new CreateRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReleaseWithDefaults() *CreateRelease {
	this := CreateRelease{}
	return &this
}

// GetArtifacts returns the Artifacts field value
func (o *CreateRelease) GetArtifacts() []Artifact {
	if o == nil {
		var ret []Artifact
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *CreateRelease) GetArtifactsOk() (*[]Artifact, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Artifacts, true
}

// SetArtifacts sets field value
func (o *CreateRelease) SetArtifacts(v []Artifact) {
	o.Artifacts = v
}

// GetReleaseDateMsecs returns the ReleaseDateMsecs field value
func (o *CreateRelease) GetReleaseDateMsecs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReleaseDateMsecs
}

// GetReleaseDateMsecsOk returns a tuple with the ReleaseDateMsecs field value
// and a boolean to check if the value has been set.
func (o *CreateRelease) GetReleaseDateMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseDateMsecs, true
}

// SetReleaseDateMsecs sets field value
func (o *CreateRelease) SetReleaseDateMsecs(v int64) {
	o.ReleaseDateMsecs = v
}

// GetReleaseNotes returns the ReleaseNotes field value
func (o *CreateRelease) GetReleaseNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseNotes
}

// GetReleaseNotesOk returns a tuple with the ReleaseNotes field value
// and a boolean to check if the value has been set.
func (o *CreateRelease) GetReleaseNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseNotes, true
}

// SetReleaseNotes sets field value
func (o *CreateRelease) SetReleaseNotes(v string) {
	o.ReleaseNotes = v
}

// GetReleaseTag returns the ReleaseTag field value
func (o *CreateRelease) GetReleaseTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseTag
}

// GetReleaseTagOk returns a tuple with the ReleaseTag field value
// and a boolean to check if the value has been set.
func (o *CreateRelease) GetReleaseTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseTag, true
}

// SetReleaseTag sets field value
func (o *CreateRelease) SetReleaseTag(v string) {
	o.ReleaseTag = v
}

// GetReleaseType returns the ReleaseType field value
func (o *CreateRelease) GetReleaseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value
// and a boolean to check if the value has been set.
func (o *CreateRelease) GetReleaseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseType, true
}

// SetReleaseType sets field value
func (o *CreateRelease) SetReleaseType(v string) {
	o.ReleaseType = v
}

// GetReleaseUuid returns the ReleaseUuid field value
func (o *CreateRelease) GetReleaseUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseUuid
}

// GetReleaseUuidOk returns a tuple with the ReleaseUuid field value
// and a boolean to check if the value has been set.
func (o *CreateRelease) GetReleaseUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseUuid, true
}

// SetReleaseUuid sets field value
func (o *CreateRelease) SetReleaseUuid(v string) {
	o.ReleaseUuid = v
}

// GetVersion returns the Version field value
func (o *CreateRelease) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CreateRelease) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CreateRelease) SetVersion(v string) {
	o.Version = v
}

// GetYbType returns the YbType field value
func (o *CreateRelease) GetYbType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbType
}

// GetYbTypeOk returns a tuple with the YbType field value
// and a boolean to check if the value has been set.
func (o *CreateRelease) GetYbTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbType, true
}

// SetYbType sets field value
func (o *CreateRelease) SetYbType(v string) {
	o.YbType = v
}

func (o CreateRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["artifacts"] = o.Artifacts
	}
	if true {
		toSerialize["release_date_msecs"] = o.ReleaseDateMsecs
	}
	if true {
		toSerialize["release_notes"] = o.ReleaseNotes
	}
	if true {
		toSerialize["release_tag"] = o.ReleaseTag
	}
	if true {
		toSerialize["release_type"] = o.ReleaseType
	}
	if true {
		toSerialize["release_uuid"] = o.ReleaseUuid
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["yb_type"] = o.YbType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRelease struct {
	value *CreateRelease
	isSet bool
}

func (v NullableCreateRelease) Get() *CreateRelease {
	return v.value
}

func (v *NullableCreateRelease) Set(val *CreateRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRelease(val *CreateRelease) *NullableCreateRelease {
	return &NullableCreateRelease{value: val, isSet: true}
}

func (v NullableCreateRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


