/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// GflagMetadata Metadata for Gflags of YB Controller.
type GflagMetadata struct {
	// The name of the flag.
	Name *string `json:"name,omitempty"`
	// A brief description of what the flag does.
	Meaning *string `json:"meaning,omitempty"`
	// The default value of the flag.
	Default *string `json:"default,omitempty"`
	// The data type of the flag.
	Type *string `json:"type,omitempty"`
}

// NewGflagMetadata instantiates a new GflagMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGflagMetadata() *GflagMetadata {
	this := GflagMetadata{}
	return &this
}

// NewGflagMetadataWithDefaults instantiates a new GflagMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGflagMetadataWithDefaults() *GflagMetadata {
	this := GflagMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GflagMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GflagMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GflagMetadata) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GflagMetadata) SetName(v string) {
	o.Name = &v
}

// GetMeaning returns the Meaning field value if set, zero value otherwise.
func (o *GflagMetadata) GetMeaning() string {
	if o == nil || o.Meaning == nil {
		var ret string
		return ret
	}
	return *o.Meaning
}

// GetMeaningOk returns a tuple with the Meaning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GflagMetadata) GetMeaningOk() (*string, bool) {
	if o == nil || o.Meaning == nil {
		return nil, false
	}
	return o.Meaning, true
}

// HasMeaning returns a boolean if a field has been set.
func (o *GflagMetadata) HasMeaning() bool {
	if o != nil && o.Meaning != nil {
		return true
	}

	return false
}

// SetMeaning gets a reference to the given string and assigns it to the Meaning field.
func (o *GflagMetadata) SetMeaning(v string) {
	o.Meaning = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *GflagMetadata) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GflagMetadata) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *GflagMetadata) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *GflagMetadata) SetDefault(v string) {
	o.Default = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GflagMetadata) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GflagMetadata) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GflagMetadata) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GflagMetadata) SetType(v string) {
	o.Type = &v
}

func (o GflagMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Meaning != nil {
		toSerialize["meaning"] = o.Meaning
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGflagMetadata struct {
	value *GflagMetadata
	isSet bool
}

func (v NullableGflagMetadata) Get() *GflagMetadata {
	return v.value
}

func (v *NullableGflagMetadata) Set(val *GflagMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableGflagMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableGflagMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGflagMetadata(val *GflagMetadata) *NullableGflagMetadata {
	return &NullableGflagMetadata{value: val, isSet: true}
}

func (v NullableGflagMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGflagMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


