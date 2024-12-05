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

// GFlagDetails struct for GFlagDetails
type GFlagDetails struct {
	// WARNING: This is a preview API that could change. Current value of the gflag
	Current *string `json:"current,omitempty"`
	// WARNING: This is a preview API that could change. Default value of the gflag
	Default *string `json:"default,omitempty"`
	// WARNING: This is a preview API that could change. File where the gflag is defined
	File *string `json:"file,omitempty"`
	// WARNING: This is a preview API that could change. Initial value of the gflag
	Initial *string `json:"initial,omitempty"`
	// WARNING: This is a preview API that could change. Meaning of the gflag
	Meaning *string `json:"meaning,omitempty"`
	// WARNING: This is a preview API that could change. Name of the gflag
	Name *string `json:"name,omitempty"`
	// WARNING: This is a preview API that could change. Tags of the gflag
	Tags *string `json:"tags,omitempty"`
	// WARNING: This is a preview API that could change. Target of the gflag
	Target *string `json:"target,omitempty"`
	// WARNING: This is a preview API that could change. Type of the gflag
	Type *string `json:"type,omitempty"`
}

// NewGFlagDetails instantiates a new GFlagDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGFlagDetails() *GFlagDetails {
	this := GFlagDetails{}
	return &this
}

// NewGFlagDetailsWithDefaults instantiates a new GFlagDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGFlagDetailsWithDefaults() *GFlagDetails {
	this := GFlagDetails{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *GFlagDetails) GetCurrent() string {
	if o == nil || o.Current == nil {
		var ret string
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetCurrentOk() (*string, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *GFlagDetails) HasCurrent() bool {
	if o != nil && o.Current != nil {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given string and assigns it to the Current field.
func (o *GFlagDetails) SetCurrent(v string) {
	o.Current = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *GFlagDetails) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *GFlagDetails) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *GFlagDetails) SetDefault(v string) {
	o.Default = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *GFlagDetails) GetFile() string {
	if o == nil || o.File == nil {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetFileOk() (*string, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *GFlagDetails) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *GFlagDetails) SetFile(v string) {
	o.File = &v
}

// GetInitial returns the Initial field value if set, zero value otherwise.
func (o *GFlagDetails) GetInitial() string {
	if o == nil || o.Initial == nil {
		var ret string
		return ret
	}
	return *o.Initial
}

// GetInitialOk returns a tuple with the Initial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetInitialOk() (*string, bool) {
	if o == nil || o.Initial == nil {
		return nil, false
	}
	return o.Initial, true
}

// HasInitial returns a boolean if a field has been set.
func (o *GFlagDetails) HasInitial() bool {
	if o != nil && o.Initial != nil {
		return true
	}

	return false
}

// SetInitial gets a reference to the given string and assigns it to the Initial field.
func (o *GFlagDetails) SetInitial(v string) {
	o.Initial = &v
}

// GetMeaning returns the Meaning field value if set, zero value otherwise.
func (o *GFlagDetails) GetMeaning() string {
	if o == nil || o.Meaning == nil {
		var ret string
		return ret
	}
	return *o.Meaning
}

// GetMeaningOk returns a tuple with the Meaning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetMeaningOk() (*string, bool) {
	if o == nil || o.Meaning == nil {
		return nil, false
	}
	return o.Meaning, true
}

// HasMeaning returns a boolean if a field has been set.
func (o *GFlagDetails) HasMeaning() bool {
	if o != nil && o.Meaning != nil {
		return true
	}

	return false
}

// SetMeaning gets a reference to the given string and assigns it to the Meaning field.
func (o *GFlagDetails) SetMeaning(v string) {
	o.Meaning = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GFlagDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GFlagDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GFlagDetails) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GFlagDetails) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GFlagDetails) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *GFlagDetails) SetTags(v string) {
	o.Tags = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *GFlagDetails) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *GFlagDetails) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *GFlagDetails) SetTarget(v string) {
	o.Target = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GFlagDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFlagDetails) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GFlagDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GFlagDetails) SetType(v string) {
	o.Type = &v
}

func (o GFlagDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Initial != nil {
		toSerialize["initial"] = o.Initial
	}
	if o.Meaning != nil {
		toSerialize["meaning"] = o.Meaning
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGFlagDetails struct {
	value *GFlagDetails
	isSet bool
}

func (v NullableGFlagDetails) Get() *GFlagDetails {
	return v.value
}

func (v *NullableGFlagDetails) Set(val *GFlagDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGFlagDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGFlagDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGFlagDetails(val *GFlagDetails) *NullableGFlagDetails {
	return &NullableGFlagDetails{value: val, isSet: true}
}

func (v NullableGFlagDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGFlagDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

