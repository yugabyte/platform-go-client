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

// ResponseRelease struct for ResponseRelease
type ResponseRelease struct {
	Artifacts []Artifact `json:"artifacts"`
	ReleaseDateMsecs int64 `json:"release_date_msecs"`
	ReleaseNotes string `json:"release_notes"`
	ReleaseTag string `json:"release_tag"`
	ReleaseType string `json:"release_type"`
	ReleaseUuid string `json:"release_uuid"`
	State string `json:"state"`
	Universes []Universe `json:"universes"`
	Version string `json:"version"`
	YbType string `json:"yb_type"`
}

// NewResponseRelease instantiates a new ResponseRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRelease(artifacts []Artifact, releaseDateMsecs int64, releaseNotes string, releaseTag string, releaseType string, releaseUuid string, state string, universes []Universe, version string, ybType string) *ResponseRelease {
	this := ResponseRelease{}
	this.Artifacts = artifacts
	this.ReleaseDateMsecs = releaseDateMsecs
	this.ReleaseNotes = releaseNotes
	this.ReleaseTag = releaseTag
	this.ReleaseType = releaseType
	this.ReleaseUuid = releaseUuid
	this.State = state
	this.Universes = universes
	this.Version = version
	this.YbType = ybType
	return &this
}

// NewResponseReleaseWithDefaults instantiates a new ResponseRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseReleaseWithDefaults() *ResponseRelease {
	this := ResponseRelease{}
	return &this
}

// GetArtifacts returns the Artifacts field value
func (o *ResponseRelease) GetArtifacts() []Artifact {
	if o == nil {
		var ret []Artifact
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetArtifactsOk() (*[]Artifact, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Artifacts, true
}

// SetArtifacts sets field value
func (o *ResponseRelease) SetArtifacts(v []Artifact) {
	o.Artifacts = v
}

// GetReleaseDateMsecs returns the ReleaseDateMsecs field value
func (o *ResponseRelease) GetReleaseDateMsecs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReleaseDateMsecs
}

// GetReleaseDateMsecsOk returns a tuple with the ReleaseDateMsecs field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetReleaseDateMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseDateMsecs, true
}

// SetReleaseDateMsecs sets field value
func (o *ResponseRelease) SetReleaseDateMsecs(v int64) {
	o.ReleaseDateMsecs = v
}

// GetReleaseNotes returns the ReleaseNotes field value
func (o *ResponseRelease) GetReleaseNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseNotes
}

// GetReleaseNotesOk returns a tuple with the ReleaseNotes field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetReleaseNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseNotes, true
}

// SetReleaseNotes sets field value
func (o *ResponseRelease) SetReleaseNotes(v string) {
	o.ReleaseNotes = v
}

// GetReleaseTag returns the ReleaseTag field value
func (o *ResponseRelease) GetReleaseTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseTag
}

// GetReleaseTagOk returns a tuple with the ReleaseTag field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetReleaseTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseTag, true
}

// SetReleaseTag sets field value
func (o *ResponseRelease) SetReleaseTag(v string) {
	o.ReleaseTag = v
}

// GetReleaseType returns the ReleaseType field value
func (o *ResponseRelease) GetReleaseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetReleaseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseType, true
}

// SetReleaseType sets field value
func (o *ResponseRelease) SetReleaseType(v string) {
	o.ReleaseType = v
}

// GetReleaseUuid returns the ReleaseUuid field value
func (o *ResponseRelease) GetReleaseUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseUuid
}

// GetReleaseUuidOk returns a tuple with the ReleaseUuid field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetReleaseUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseUuid, true
}

// SetReleaseUuid sets field value
func (o *ResponseRelease) SetReleaseUuid(v string) {
	o.ReleaseUuid = v
}

// GetState returns the State field value
func (o *ResponseRelease) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseRelease) SetState(v string) {
	o.State = v
}

// GetUniverses returns the Universes field value
func (o *ResponseRelease) GetUniverses() []Universe {
	if o == nil {
		var ret []Universe
		return ret
	}

	return o.Universes
}

// GetUniversesOk returns a tuple with the Universes field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetUniversesOk() (*[]Universe, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Universes, true
}

// SetUniverses sets field value
func (o *ResponseRelease) SetUniverses(v []Universe) {
	o.Universes = v
}

// GetVersion returns the Version field value
func (o *ResponseRelease) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ResponseRelease) SetVersion(v string) {
	o.Version = v
}

// GetYbType returns the YbType field value
func (o *ResponseRelease) GetYbType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbType
}

// GetYbTypeOk returns a tuple with the YbType field value
// and a boolean to check if the value has been set.
func (o *ResponseRelease) GetYbTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbType, true
}

// SetYbType sets field value
func (o *ResponseRelease) SetYbType(v string) {
	o.YbType = v
}

func (o ResponseRelease) MarshalJSON() ([]byte, error) {
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
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["universes"] = o.Universes
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["yb_type"] = o.YbType
	}
	return json.Marshal(toSerialize)
}

type NullableResponseRelease struct {
	value *ResponseRelease
	isSet bool
}

func (v NullableResponseRelease) Get() *ResponseRelease {
	return v.value
}

func (v *NullableResponseRelease) Set(val *ResponseRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRelease(val *ResponseRelease) *NullableResponseRelease {
	return &NullableResponseRelease{value: val, isSet: true}
}

func (v NullableResponseRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


