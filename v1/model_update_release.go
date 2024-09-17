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

// UpdateRelease Release metadata required to create a new release
type UpdateRelease struct {
	Artifacts []Artifact `json:"artifacts"`
	ReleaseDate int64 `json:"release_date"`
	ReleaseNotes string `json:"release_notes"`
	ReleaseTag string `json:"release_tag"`
	State string `json:"state"`
}

// NewUpdateRelease instantiates a new UpdateRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRelease(artifacts []Artifact, releaseDate int64, releaseNotes string, releaseTag string, state string) *UpdateRelease {
	this := UpdateRelease{}
	this.Artifacts = artifacts
	this.ReleaseDate = releaseDate
	this.ReleaseNotes = releaseNotes
	this.ReleaseTag = releaseTag
	this.State = state
	return &this
}

// NewUpdateReleaseWithDefaults instantiates a new UpdateRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReleaseWithDefaults() *UpdateRelease {
	this := UpdateRelease{}
	return &this
}

// GetArtifacts returns the Artifacts field value
func (o *UpdateRelease) GetArtifacts() []Artifact {
	if o == nil {
		var ret []Artifact
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *UpdateRelease) GetArtifactsOk() (*[]Artifact, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Artifacts, true
}

// SetArtifacts sets field value
func (o *UpdateRelease) SetArtifacts(v []Artifact) {
	o.Artifacts = v
}

// GetReleaseDate returns the ReleaseDate field value
func (o *UpdateRelease) GetReleaseDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
func (o *UpdateRelease) GetReleaseDateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseDate, true
}

// SetReleaseDate sets field value
func (o *UpdateRelease) SetReleaseDate(v int64) {
	o.ReleaseDate = v
}

// GetReleaseNotes returns the ReleaseNotes field value
func (o *UpdateRelease) GetReleaseNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseNotes
}

// GetReleaseNotesOk returns a tuple with the ReleaseNotes field value
// and a boolean to check if the value has been set.
func (o *UpdateRelease) GetReleaseNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseNotes, true
}

// SetReleaseNotes sets field value
func (o *UpdateRelease) SetReleaseNotes(v string) {
	o.ReleaseNotes = v
}

// GetReleaseTag returns the ReleaseTag field value
func (o *UpdateRelease) GetReleaseTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseTag
}

// GetReleaseTagOk returns a tuple with the ReleaseTag field value
// and a boolean to check if the value has been set.
func (o *UpdateRelease) GetReleaseTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseTag, true
}

// SetReleaseTag sets field value
func (o *UpdateRelease) SetReleaseTag(v string) {
	o.ReleaseTag = v
}

// GetState returns the State field value
func (o *UpdateRelease) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UpdateRelease) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *UpdateRelease) SetState(v string) {
	o.State = v
}

func (o UpdateRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["artifacts"] = o.Artifacts
	}
	if true {
		toSerialize["release_date"] = o.ReleaseDate
	}
	if true {
		toSerialize["release_notes"] = o.ReleaseNotes
	}
	if true {
		toSerialize["release_tag"] = o.ReleaseTag
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRelease struct {
	value *UpdateRelease
	isSet bool
}

func (v NullableUpdateRelease) Get() *UpdateRelease {
	return v.value
}

func (v *NullableUpdateRelease) Set(val *UpdateRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRelease(val *UpdateRelease) *NullableUpdateRelease {
	return &NullableUpdateRelease{value: val, isSet: true}
}

func (v NullableUpdateRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


