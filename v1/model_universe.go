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
	"time"
)

// Universe struct for Universe
type Universe struct {
	// Universe creation date
	CreationDate *time.Time `json:"creation_date,omitempty"`
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

// NewUniverse instantiates a new Universe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverse(name string, uuid string) *Universe {
	this := Universe{}
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewUniverseWithDefaults instantiates a new Universe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseWithDefaults() *Universe {
	this := Universe{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Universe) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Universe) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Universe) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Universe) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetName returns the Name field value
func (o *Universe) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Universe) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Universe) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *Universe) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Universe) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Universe) SetUuid(v string) {
	o.Uuid = v
}

func (o Universe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableUniverse struct {
	value *Universe
	isSet bool
}

func (v NullableUniverse) Get() *Universe {
	return v.value
}

func (v *NullableUniverse) Set(val *Universe) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverse) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverse(val *Universe) *NullableUniverse {
	return &NullableUniverse{value: val, isSet: true}
}

func (v NullableUniverse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


