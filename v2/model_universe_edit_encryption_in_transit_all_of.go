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

// UniverseEditEncryptionInTransitAllOf struct for UniverseEditEncryptionInTransitAllOf
type UniverseEditEncryptionInTransitAllOf struct {
	// Control encryption in transit between YugabyteDB nodes
	NodeToNode *bool `json:"node_to_node,omitempty"`
	// Control encryption in transit between YugabyteDB nodes and clients
	ClientToNode *bool `json:"client_to_node,omitempty"`
	// Root CA cert for node to node encryption. Required if node_to_node is true
	RootCa *string `json:"root_ca,omitempty"`
	// Root CA used for node to client encryption. Required if client_to_node is true
	ClientRootCa *string `json:"client_root_ca,omitempty"`
}

// NewUniverseEditEncryptionInTransitAllOf instantiates a new UniverseEditEncryptionInTransitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseEditEncryptionInTransitAllOf() *UniverseEditEncryptionInTransitAllOf {
	this := UniverseEditEncryptionInTransitAllOf{}
	return &this
}

// NewUniverseEditEncryptionInTransitAllOfWithDefaults instantiates a new UniverseEditEncryptionInTransitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseEditEncryptionInTransitAllOfWithDefaults() *UniverseEditEncryptionInTransitAllOf {
	this := UniverseEditEncryptionInTransitAllOf{}
	return &this
}

// GetNodeToNode returns the NodeToNode field value if set, zero value otherwise.
func (o *UniverseEditEncryptionInTransitAllOf) GetNodeToNode() bool {
	if o == nil || o.NodeToNode == nil {
		var ret bool
		return ret
	}
	return *o.NodeToNode
}

// GetNodeToNodeOk returns a tuple with the NodeToNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseEditEncryptionInTransitAllOf) GetNodeToNodeOk() (*bool, bool) {
	if o == nil || o.NodeToNode == nil {
		return nil, false
	}
	return o.NodeToNode, true
}

// HasNodeToNode returns a boolean if a field has been set.
func (o *UniverseEditEncryptionInTransitAllOf) HasNodeToNode() bool {
	if o != nil && o.NodeToNode != nil {
		return true
	}

	return false
}

// SetNodeToNode gets a reference to the given bool and assigns it to the NodeToNode field.
func (o *UniverseEditEncryptionInTransitAllOf) SetNodeToNode(v bool) {
	o.NodeToNode = &v
}

// GetClientToNode returns the ClientToNode field value if set, zero value otherwise.
func (o *UniverseEditEncryptionInTransitAllOf) GetClientToNode() bool {
	if o == nil || o.ClientToNode == nil {
		var ret bool
		return ret
	}
	return *o.ClientToNode
}

// GetClientToNodeOk returns a tuple with the ClientToNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseEditEncryptionInTransitAllOf) GetClientToNodeOk() (*bool, bool) {
	if o == nil || o.ClientToNode == nil {
		return nil, false
	}
	return o.ClientToNode, true
}

// HasClientToNode returns a boolean if a field has been set.
func (o *UniverseEditEncryptionInTransitAllOf) HasClientToNode() bool {
	if o != nil && o.ClientToNode != nil {
		return true
	}

	return false
}

// SetClientToNode gets a reference to the given bool and assigns it to the ClientToNode field.
func (o *UniverseEditEncryptionInTransitAllOf) SetClientToNode(v bool) {
	o.ClientToNode = &v
}

// GetRootCa returns the RootCa field value if set, zero value otherwise.
func (o *UniverseEditEncryptionInTransitAllOf) GetRootCa() string {
	if o == nil || o.RootCa == nil {
		var ret string
		return ret
	}
	return *o.RootCa
}

// GetRootCaOk returns a tuple with the RootCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseEditEncryptionInTransitAllOf) GetRootCaOk() (*string, bool) {
	if o == nil || o.RootCa == nil {
		return nil, false
	}
	return o.RootCa, true
}

// HasRootCa returns a boolean if a field has been set.
func (o *UniverseEditEncryptionInTransitAllOf) HasRootCa() bool {
	if o != nil && o.RootCa != nil {
		return true
	}

	return false
}

// SetRootCa gets a reference to the given string and assigns it to the RootCa field.
func (o *UniverseEditEncryptionInTransitAllOf) SetRootCa(v string) {
	o.RootCa = &v
}

// GetClientRootCa returns the ClientRootCa field value if set, zero value otherwise.
func (o *UniverseEditEncryptionInTransitAllOf) GetClientRootCa() string {
	if o == nil || o.ClientRootCa == nil {
		var ret string
		return ret
	}
	return *o.ClientRootCa
}

// GetClientRootCaOk returns a tuple with the ClientRootCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseEditEncryptionInTransitAllOf) GetClientRootCaOk() (*string, bool) {
	if o == nil || o.ClientRootCa == nil {
		return nil, false
	}
	return o.ClientRootCa, true
}

// HasClientRootCa returns a boolean if a field has been set.
func (o *UniverseEditEncryptionInTransitAllOf) HasClientRootCa() bool {
	if o != nil && o.ClientRootCa != nil {
		return true
	}

	return false
}

// SetClientRootCa gets a reference to the given string and assigns it to the ClientRootCa field.
func (o *UniverseEditEncryptionInTransitAllOf) SetClientRootCa(v string) {
	o.ClientRootCa = &v
}

func (o UniverseEditEncryptionInTransitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeToNode != nil {
		toSerialize["node_to_node"] = o.NodeToNode
	}
	if o.ClientToNode != nil {
		toSerialize["client_to_node"] = o.ClientToNode
	}
	if o.RootCa != nil {
		toSerialize["root_ca"] = o.RootCa
	}
	if o.ClientRootCa != nil {
		toSerialize["client_root_ca"] = o.ClientRootCa
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseEditEncryptionInTransitAllOf struct {
	value *UniverseEditEncryptionInTransitAllOf
	isSet bool
}

func (v NullableUniverseEditEncryptionInTransitAllOf) Get() *UniverseEditEncryptionInTransitAllOf {
	return v.value
}

func (v *NullableUniverseEditEncryptionInTransitAllOf) Set(val *UniverseEditEncryptionInTransitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseEditEncryptionInTransitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseEditEncryptionInTransitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseEditEncryptionInTransitAllOf(val *UniverseEditEncryptionInTransitAllOf) *NullableUniverseEditEncryptionInTransitAllOf {
	return &NullableUniverseEditEncryptionInTransitAllOf{value: val, isSet: true}
}

func (v NullableUniverseEditEncryptionInTransitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseEditEncryptionInTransitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


