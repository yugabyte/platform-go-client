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

// ImageBundle struct for ImageBundle
type ImageBundle struct {
	// Is the ImageBundle Active
	Active *bool `json:"active,omitempty"`
	Details *ImageBundleDetails `json:"details,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	// Image Bundle Name
	Name *string `json:"name,omitempty"`
	// Default Image Bundle. A provider can have two defaults, one per architecture
	UseAsDefault *bool `json:"useAsDefault,omitempty"`
	// Image Bundle UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewImageBundle instantiates a new ImageBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageBundle() *ImageBundle {
	this := ImageBundle{}
	return &this
}

// NewImageBundleWithDefaults instantiates a new ImageBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageBundleWithDefaults() *ImageBundle {
	this := ImageBundle{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ImageBundle) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBundle) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ImageBundle) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ImageBundle) SetActive(v bool) {
	o.Active = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ImageBundle) GetDetails() ImageBundleDetails {
	if o == nil || o.Details == nil {
		var ret ImageBundleDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBundle) GetDetailsOk() (*ImageBundleDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ImageBundle) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ImageBundleDetails and assigns it to the Details field.
func (o *ImageBundle) SetDetails(v ImageBundleDetails) {
	o.Details = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ImageBundle) GetMetadata() Metadata {
	if o == nil || o.Metadata == nil {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBundle) GetMetadataOk() (*Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ImageBundle) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *ImageBundle) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImageBundle) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBundle) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImageBundle) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImageBundle) SetName(v string) {
	o.Name = &v
}

// GetUseAsDefault returns the UseAsDefault field value if set, zero value otherwise.
func (o *ImageBundle) GetUseAsDefault() bool {
	if o == nil || o.UseAsDefault == nil {
		var ret bool
		return ret
	}
	return *o.UseAsDefault
}

// GetUseAsDefaultOk returns a tuple with the UseAsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBundle) GetUseAsDefaultOk() (*bool, bool) {
	if o == nil || o.UseAsDefault == nil {
		return nil, false
	}
	return o.UseAsDefault, true
}

// HasUseAsDefault returns a boolean if a field has been set.
func (o *ImageBundle) HasUseAsDefault() bool {
	if o != nil && o.UseAsDefault != nil {
		return true
	}

	return false
}

// SetUseAsDefault gets a reference to the given bool and assigns it to the UseAsDefault field.
func (o *ImageBundle) SetUseAsDefault(v bool) {
	o.UseAsDefault = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ImageBundle) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBundle) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ImageBundle) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ImageBundle) SetUuid(v string) {
	o.Uuid = &v
}

func (o ImageBundle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UseAsDefault != nil {
		toSerialize["useAsDefault"] = o.UseAsDefault
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableImageBundle struct {
	value *ImageBundle
	isSet bool
}

func (v NullableImageBundle) Get() *ImageBundle {
	return v.value
}

func (v *NullableImageBundle) Set(val *ImageBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBundle(val *ImageBundle) *NullableImageBundle {
	return &NullableImageBundle{value: val, isSet: true}
}

func (v NullableImageBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


