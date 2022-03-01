/*
 * Yugabyte Platform APIs
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

// ReleaseMetadata Yugabyte release metadata
type ReleaseMetadata struct {
	// Helm chart path
	ChartPath *string `json:"chartPath,omitempty"`
	// Release file path
	FilePath *string `json:"filePath,omitempty"`
	Gcs *GCSLocation `json:"gcs,omitempty"`
	Http *HttpLocation `json:"http,omitempty"`
	// Release image tag
	ImageTag *string `json:"imageTag,omitempty"`
	// Release notes
	Notes *[]string `json:"notes,omitempty"`
	S3 *S3Location `json:"s3,omitempty"`
	// Release state
	State *string `json:"state,omitempty"`
}

// NewReleaseMetadata instantiates a new ReleaseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseMetadata() *ReleaseMetadata {
	this := ReleaseMetadata{}
	return &this
}

// NewReleaseMetadataWithDefaults instantiates a new ReleaseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseMetadataWithDefaults() *ReleaseMetadata {
	this := ReleaseMetadata{}
	return &this
}

// GetChartPath returns the ChartPath field value if set, zero value otherwise.
func (o *ReleaseMetadata) GetChartPath() string {
	if o == nil || o.ChartPath == nil {
		var ret string
		return ret
	}
	return *o.ChartPath
}

// GetChartPathOk returns a tuple with the ChartPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseMetadata) GetChartPathOk() (*string, bool) {
	if o == nil || o.ChartPath == nil {
		return nil, false
	}
	return o.ChartPath, true
}

// HasChartPath returns a boolean if a field has been set.
func (o *ReleaseMetadata) HasChartPath() bool {
	if o != nil && o.ChartPath != nil {
		return true
	}

	return false
}

// SetChartPath gets a reference to the given string and assigns it to the ChartPath field.
func (o *ReleaseMetadata) SetChartPath(v string) {
	o.ChartPath = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *ReleaseMetadata) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseMetadata) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *ReleaseMetadata) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *ReleaseMetadata) SetFilePath(v string) {
	o.FilePath = &v
}

// GetGcs returns the Gcs field value if set, zero value otherwise.
func (o *ReleaseMetadata) GetGcs() GCSLocation {
	if o == nil || o.Gcs == nil {
		var ret GCSLocation
		return ret
	}
	return *o.Gcs
}

// GetGcsOk returns a tuple with the Gcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseMetadata) GetGcsOk() (*GCSLocation, bool) {
	if o == nil || o.Gcs == nil {
		return nil, false
	}
	return o.Gcs, true
}

// HasGcs returns a boolean if a field has been set.
func (o *ReleaseMetadata) HasGcs() bool {
	if o != nil && o.Gcs != nil {
		return true
	}

	return false
}

// SetGcs gets a reference to the given GCSLocation and assigns it to the Gcs field.
func (o *ReleaseMetadata) SetGcs(v GCSLocation) {
	o.Gcs = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *ReleaseMetadata) GetHttp() HttpLocation {
	if o == nil || o.Http == nil {
		var ret HttpLocation
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseMetadata) GetHttpOk() (*HttpLocation, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *ReleaseMetadata) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given HttpLocation and assigns it to the Http field.
func (o *ReleaseMetadata) SetHttp(v HttpLocation) {
	o.Http = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *ReleaseMetadata) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseMetadata) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *ReleaseMetadata) HasImageTag() bool {
	if o != nil && o.ImageTag != nil {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *ReleaseMetadata) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ReleaseMetadata) GetNotes() []string {
	if o == nil || o.Notes == nil {
		var ret []string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseMetadata) GetNotesOk() (*[]string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ReleaseMetadata) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *ReleaseMetadata) SetNotes(v []string) {
	o.Notes = &v
}

// GetS3 returns the S3 field value if set, zero value otherwise.
func (o *ReleaseMetadata) GetS3() S3Location {
	if o == nil || o.S3 == nil {
		var ret S3Location
		return ret
	}
	return *o.S3
}

// GetS3Ok returns a tuple with the S3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseMetadata) GetS3Ok() (*S3Location, bool) {
	if o == nil || o.S3 == nil {
		return nil, false
	}
	return o.S3, true
}

// HasS3 returns a boolean if a field has been set.
func (o *ReleaseMetadata) HasS3() bool {
	if o != nil && o.S3 != nil {
		return true
	}

	return false
}

// SetS3 gets a reference to the given S3Location and assigns it to the S3 field.
func (o *ReleaseMetadata) SetS3(v S3Location) {
	o.S3 = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ReleaseMetadata) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseMetadata) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ReleaseMetadata) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ReleaseMetadata) SetState(v string) {
	o.State = &v
}

func (o ReleaseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChartPath != nil {
		toSerialize["chartPath"] = o.ChartPath
	}
	if o.FilePath != nil {
		toSerialize["filePath"] = o.FilePath
	}
	if o.Gcs != nil {
		toSerialize["gcs"] = o.Gcs
	}
	if o.Http != nil {
		toSerialize["http"] = o.Http
	}
	if o.ImageTag != nil {
		toSerialize["imageTag"] = o.ImageTag
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.S3 != nil {
		toSerialize["s3"] = o.S3
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableReleaseMetadata struct {
	value *ReleaseMetadata
	isSet bool
}

func (v NullableReleaseMetadata) Get() *ReleaseMetadata {
	return v.value
}

func (v *NullableReleaseMetadata) Set(val *ReleaseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseMetadata(val *ReleaseMetadata) *NullableReleaseMetadata {
	return &NullableReleaseMetadata{value: val, isSet: true}
}

func (v NullableReleaseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


