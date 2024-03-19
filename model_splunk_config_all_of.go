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

// SplunkConfigAllOf Splunk Config
type SplunkConfigAllOf struct {
	// End Point
	Endpoint *string `json:"endpoint,omitempty"`
	// Index
	Index *string `json:"index,omitempty"`
	// Source
	Source *string `json:"source,omitempty"`
	// Source Type
	SourceType *string `json:"sourceType,omitempty"`
	// Token
	Token *string `json:"token,omitempty"`
}

// NewSplunkConfigAllOf instantiates a new SplunkConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplunkConfigAllOf() *SplunkConfigAllOf {
	this := SplunkConfigAllOf{}
	return &this
}

// NewSplunkConfigAllOfWithDefaults instantiates a new SplunkConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplunkConfigAllOfWithDefaults() *SplunkConfigAllOf {
	this := SplunkConfigAllOf{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *SplunkConfigAllOf) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkConfigAllOf) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *SplunkConfigAllOf) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *SplunkConfigAllOf) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SplunkConfigAllOf) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkConfigAllOf) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SplunkConfigAllOf) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *SplunkConfigAllOf) SetIndex(v string) {
	o.Index = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SplunkConfigAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkConfigAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SplunkConfigAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SplunkConfigAllOf) SetSource(v string) {
	o.Source = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *SplunkConfigAllOf) GetSourceType() string {
	if o == nil || o.SourceType == nil {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkConfigAllOf) GetSourceTypeOk() (*string, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *SplunkConfigAllOf) HasSourceType() bool {
	if o != nil && o.SourceType != nil {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *SplunkConfigAllOf) SetSourceType(v string) {
	o.SourceType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SplunkConfigAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkConfigAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SplunkConfigAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SplunkConfigAllOf) SetToken(v string) {
	o.Token = &v
}

func (o SplunkConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.SourceType != nil {
		toSerialize["sourceType"] = o.SourceType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableSplunkConfigAllOf struct {
	value *SplunkConfigAllOf
	isSet bool
}

func (v NullableSplunkConfigAllOf) Get() *SplunkConfigAllOf {
	return v.value
}

func (v *NullableSplunkConfigAllOf) Set(val *SplunkConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSplunkConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSplunkConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplunkConfigAllOf(val *SplunkConfigAllOf) *NullableSplunkConfigAllOf {
	return &NullableSplunkConfigAllOf{value: val, isSet: true}
}

func (v NullableSplunkConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplunkConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

