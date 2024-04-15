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

// NodeInstance A single node instance, attached to a provider and availability zone
type NodeInstance struct {
	Details NodeInstanceData `json:"details"`
	// Node details (as a JSON object)
	DetailsJson *string `json:"detailsJson,omitempty"`
	InUse *bool `json:"inUse,omitempty"`
	// The node instance's name
	InstanceName *string `json:"instanceName,omitempty"`
	// The node's type code
	InstanceTypeCode *string `json:"instanceTypeCode,omitempty"`
	// The node's name
	NodeName *string `json:"nodeName,omitempty"`
	// The node's UUID
	NodeUuid *string `json:"nodeUuid,omitempty"`
	// State of on-prem node
	State *string `json:"state,omitempty"`
	// The availability zone's UUID
	ZoneUuid *string `json:"zoneUuid,omitempty"`
}

// NewNodeInstance instantiates a new NodeInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeInstance(details NodeInstanceData) *NodeInstance {
	this := NodeInstance{}
	this.Details = details
	return &this
}

// NewNodeInstanceWithDefaults instantiates a new NodeInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeInstanceWithDefaults() *NodeInstance {
	this := NodeInstance{}
	return &this
}

// GetDetails returns the Details field value
func (o *NodeInstance) GetDetails() NodeInstanceData {
	if o == nil {
		var ret NodeInstanceData
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetDetailsOk() (*NodeInstanceData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *NodeInstance) SetDetails(v NodeInstanceData) {
	o.Details = v
}

// GetDetailsJson returns the DetailsJson field value if set, zero value otherwise.
func (o *NodeInstance) GetDetailsJson() string {
	if o == nil || o.DetailsJson == nil {
		var ret string
		return ret
	}
	return *o.DetailsJson
}

// GetDetailsJsonOk returns a tuple with the DetailsJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetDetailsJsonOk() (*string, bool) {
	if o == nil || o.DetailsJson == nil {
		return nil, false
	}
	return o.DetailsJson, true
}

// HasDetailsJson returns a boolean if a field has been set.
func (o *NodeInstance) HasDetailsJson() bool {
	if o != nil && o.DetailsJson != nil {
		return true
	}

	return false
}

// SetDetailsJson gets a reference to the given string and assigns it to the DetailsJson field.
func (o *NodeInstance) SetDetailsJson(v string) {
	o.DetailsJson = &v
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *NodeInstance) GetInUse() bool {
	if o == nil || o.InUse == nil {
		var ret bool
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetInUseOk() (*bool, bool) {
	if o == nil || o.InUse == nil {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *NodeInstance) HasInUse() bool {
	if o != nil && o.InUse != nil {
		return true
	}

	return false
}

// SetInUse gets a reference to the given bool and assigns it to the InUse field.
func (o *NodeInstance) SetInUse(v bool) {
	o.InUse = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *NodeInstance) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *NodeInstance) HasInstanceName() bool {
	if o != nil && o.InstanceName != nil {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *NodeInstance) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetInstanceTypeCode returns the InstanceTypeCode field value if set, zero value otherwise.
func (o *NodeInstance) GetInstanceTypeCode() string {
	if o == nil || o.InstanceTypeCode == nil {
		var ret string
		return ret
	}
	return *o.InstanceTypeCode
}

// GetInstanceTypeCodeOk returns a tuple with the InstanceTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetInstanceTypeCodeOk() (*string, bool) {
	if o == nil || o.InstanceTypeCode == nil {
		return nil, false
	}
	return o.InstanceTypeCode, true
}

// HasInstanceTypeCode returns a boolean if a field has been set.
func (o *NodeInstance) HasInstanceTypeCode() bool {
	if o != nil && o.InstanceTypeCode != nil {
		return true
	}

	return false
}

// SetInstanceTypeCode gets a reference to the given string and assigns it to the InstanceTypeCode field.
func (o *NodeInstance) SetInstanceTypeCode(v string) {
	o.InstanceTypeCode = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *NodeInstance) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *NodeInstance) HasNodeName() bool {
	if o != nil && o.NodeName != nil {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *NodeInstance) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeUuid returns the NodeUuid field value if set, zero value otherwise.
func (o *NodeInstance) GetNodeUuid() string {
	if o == nil || o.NodeUuid == nil {
		var ret string
		return ret
	}
	return *o.NodeUuid
}

// GetNodeUuidOk returns a tuple with the NodeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetNodeUuidOk() (*string, bool) {
	if o == nil || o.NodeUuid == nil {
		return nil, false
	}
	return o.NodeUuid, true
}

// HasNodeUuid returns a boolean if a field has been set.
func (o *NodeInstance) HasNodeUuid() bool {
	if o != nil && o.NodeUuid != nil {
		return true
	}

	return false
}

// SetNodeUuid gets a reference to the given string and assigns it to the NodeUuid field.
func (o *NodeInstance) SetNodeUuid(v string) {
	o.NodeUuid = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NodeInstance) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NodeInstance) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NodeInstance) SetState(v string) {
	o.State = &v
}

// GetZoneUuid returns the ZoneUuid field value if set, zero value otherwise.
func (o *NodeInstance) GetZoneUuid() string {
	if o == nil || o.ZoneUuid == nil {
		var ret string
		return ret
	}
	return *o.ZoneUuid
}

// GetZoneUuidOk returns a tuple with the ZoneUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetZoneUuidOk() (*string, bool) {
	if o == nil || o.ZoneUuid == nil {
		return nil, false
	}
	return o.ZoneUuid, true
}

// HasZoneUuid returns a boolean if a field has been set.
func (o *NodeInstance) HasZoneUuid() bool {
	if o != nil && o.ZoneUuid != nil {
		return true
	}

	return false
}

// SetZoneUuid gets a reference to the given string and assigns it to the ZoneUuid field.
func (o *NodeInstance) SetZoneUuid(v string) {
	o.ZoneUuid = &v
}

func (o NodeInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["details"] = o.Details
	}
	if o.DetailsJson != nil {
		toSerialize["detailsJson"] = o.DetailsJson
	}
	if o.InUse != nil {
		toSerialize["inUse"] = o.InUse
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.InstanceTypeCode != nil {
		toSerialize["instanceTypeCode"] = o.InstanceTypeCode
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.NodeUuid != nil {
		toSerialize["nodeUuid"] = o.NodeUuid
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ZoneUuid != nil {
		toSerialize["zoneUuid"] = o.ZoneUuid
	}
	return json.Marshal(toSerialize)
}

type NullableNodeInstance struct {
	value *NodeInstance
	isSet bool
}

func (v NullableNodeInstance) Get() *NodeInstance {
	return v.value
}

func (v *NullableNodeInstance) Set(val *NodeInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeInstance(val *NodeInstance) *NullableNodeInstance {
	return &NullableNodeInstance{value: val, isSet: true}
}

func (v NullableNodeInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


