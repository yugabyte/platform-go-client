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

// CommunicationPorts Communication ports
type CommunicationPorts struct {
	// Master table HTTP port
	MasterHttpPort *int32 `json:"masterHttpPort,omitempty"`
	// Master table RCP port
	MasterRpcPort *int32 `json:"masterRpcPort,omitempty"`
	// Node exporter port
	NodeExporterPort *int32 `json:"nodeExporterPort,omitempty"`
	// Redis HTTP port
	RedisServerHttpPort *int32 `json:"redisServerHttpPort,omitempty"`
	// Redis RPC port
	RedisServerRpcPort *int32 `json:"redisServerRpcPort,omitempty"`
	// Tablet server HTTP port
	TserverHttpPort *int32 `json:"tserverHttpPort,omitempty"`
	// Tablet server RPC port
	TserverRpcPort *int32 `json:"tserverRpcPort,omitempty"`
	// YQL HTTP port
	YqlServerHttpPort *int32 `json:"yqlServerHttpPort,omitempty"`
	// YQL RPC port
	YqlServerRpcPort *int32 `json:"yqlServerRpcPort,omitempty"`
	// YSQL HTTP port
	YsqlServerHttpPort *int32 `json:"ysqlServerHttpPort,omitempty"`
	// YSQL RPC port
	YsqlServerRpcPort *int32 `json:"ysqlServerRpcPort,omitempty"`
}

// NewCommunicationPorts instantiates a new CommunicationPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationPorts() *CommunicationPorts {
	this := CommunicationPorts{}
	return &this
}

// NewCommunicationPortsWithDefaults instantiates a new CommunicationPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationPortsWithDefaults() *CommunicationPorts {
	this := CommunicationPorts{}
	return &this
}

// GetMasterHttpPort returns the MasterHttpPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetMasterHttpPort() int32 {
	if o == nil || o.MasterHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.MasterHttpPort
}

// GetMasterHttpPortOk returns a tuple with the MasterHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetMasterHttpPortOk() (*int32, bool) {
	if o == nil || o.MasterHttpPort == nil {
		return nil, false
	}
	return o.MasterHttpPort, true
}

// HasMasterHttpPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasMasterHttpPort() bool {
	if o != nil && o.MasterHttpPort != nil {
		return true
	}

	return false
}

// SetMasterHttpPort gets a reference to the given int32 and assigns it to the MasterHttpPort field.
func (o *CommunicationPorts) SetMasterHttpPort(v int32) {
	o.MasterHttpPort = &v
}

// GetMasterRpcPort returns the MasterRpcPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetMasterRpcPort() int32 {
	if o == nil || o.MasterRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.MasterRpcPort
}

// GetMasterRpcPortOk returns a tuple with the MasterRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetMasterRpcPortOk() (*int32, bool) {
	if o == nil || o.MasterRpcPort == nil {
		return nil, false
	}
	return o.MasterRpcPort, true
}

// HasMasterRpcPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasMasterRpcPort() bool {
	if o != nil && o.MasterRpcPort != nil {
		return true
	}

	return false
}

// SetMasterRpcPort gets a reference to the given int32 and assigns it to the MasterRpcPort field.
func (o *CommunicationPorts) SetMasterRpcPort(v int32) {
	o.MasterRpcPort = &v
}

// GetNodeExporterPort returns the NodeExporterPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetNodeExporterPort() int32 {
	if o == nil || o.NodeExporterPort == nil {
		var ret int32
		return ret
	}
	return *o.NodeExporterPort
}

// GetNodeExporterPortOk returns a tuple with the NodeExporterPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetNodeExporterPortOk() (*int32, bool) {
	if o == nil || o.NodeExporterPort == nil {
		return nil, false
	}
	return o.NodeExporterPort, true
}

// HasNodeExporterPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasNodeExporterPort() bool {
	if o != nil && o.NodeExporterPort != nil {
		return true
	}

	return false
}

// SetNodeExporterPort gets a reference to the given int32 and assigns it to the NodeExporterPort field.
func (o *CommunicationPorts) SetNodeExporterPort(v int32) {
	o.NodeExporterPort = &v
}

// GetRedisServerHttpPort returns the RedisServerHttpPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetRedisServerHttpPort() int32 {
	if o == nil || o.RedisServerHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.RedisServerHttpPort
}

// GetRedisServerHttpPortOk returns a tuple with the RedisServerHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetRedisServerHttpPortOk() (*int32, bool) {
	if o == nil || o.RedisServerHttpPort == nil {
		return nil, false
	}
	return o.RedisServerHttpPort, true
}

// HasRedisServerHttpPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasRedisServerHttpPort() bool {
	if o != nil && o.RedisServerHttpPort != nil {
		return true
	}

	return false
}

// SetRedisServerHttpPort gets a reference to the given int32 and assigns it to the RedisServerHttpPort field.
func (o *CommunicationPorts) SetRedisServerHttpPort(v int32) {
	o.RedisServerHttpPort = &v
}

// GetRedisServerRpcPort returns the RedisServerRpcPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetRedisServerRpcPort() int32 {
	if o == nil || o.RedisServerRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.RedisServerRpcPort
}

// GetRedisServerRpcPortOk returns a tuple with the RedisServerRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetRedisServerRpcPortOk() (*int32, bool) {
	if o == nil || o.RedisServerRpcPort == nil {
		return nil, false
	}
	return o.RedisServerRpcPort, true
}

// HasRedisServerRpcPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasRedisServerRpcPort() bool {
	if o != nil && o.RedisServerRpcPort != nil {
		return true
	}

	return false
}

// SetRedisServerRpcPort gets a reference to the given int32 and assigns it to the RedisServerRpcPort field.
func (o *CommunicationPorts) SetRedisServerRpcPort(v int32) {
	o.RedisServerRpcPort = &v
}

// GetTserverHttpPort returns the TserverHttpPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetTserverHttpPort() int32 {
	if o == nil || o.TserverHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.TserverHttpPort
}

// GetTserverHttpPortOk returns a tuple with the TserverHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetTserverHttpPortOk() (*int32, bool) {
	if o == nil || o.TserverHttpPort == nil {
		return nil, false
	}
	return o.TserverHttpPort, true
}

// HasTserverHttpPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasTserverHttpPort() bool {
	if o != nil && o.TserverHttpPort != nil {
		return true
	}

	return false
}

// SetTserverHttpPort gets a reference to the given int32 and assigns it to the TserverHttpPort field.
func (o *CommunicationPorts) SetTserverHttpPort(v int32) {
	o.TserverHttpPort = &v
}

// GetTserverRpcPort returns the TserverRpcPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetTserverRpcPort() int32 {
	if o == nil || o.TserverRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.TserverRpcPort
}

// GetTserverRpcPortOk returns a tuple with the TserverRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetTserverRpcPortOk() (*int32, bool) {
	if o == nil || o.TserverRpcPort == nil {
		return nil, false
	}
	return o.TserverRpcPort, true
}

// HasTserverRpcPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasTserverRpcPort() bool {
	if o != nil && o.TserverRpcPort != nil {
		return true
	}

	return false
}

// SetTserverRpcPort gets a reference to the given int32 and assigns it to the TserverRpcPort field.
func (o *CommunicationPorts) SetTserverRpcPort(v int32) {
	o.TserverRpcPort = &v
}

// GetYqlServerHttpPort returns the YqlServerHttpPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetYqlServerHttpPort() int32 {
	if o == nil || o.YqlServerHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.YqlServerHttpPort
}

// GetYqlServerHttpPortOk returns a tuple with the YqlServerHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetYqlServerHttpPortOk() (*int32, bool) {
	if o == nil || o.YqlServerHttpPort == nil {
		return nil, false
	}
	return o.YqlServerHttpPort, true
}

// HasYqlServerHttpPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasYqlServerHttpPort() bool {
	if o != nil && o.YqlServerHttpPort != nil {
		return true
	}

	return false
}

// SetYqlServerHttpPort gets a reference to the given int32 and assigns it to the YqlServerHttpPort field.
func (o *CommunicationPorts) SetYqlServerHttpPort(v int32) {
	o.YqlServerHttpPort = &v
}

// GetYqlServerRpcPort returns the YqlServerRpcPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetYqlServerRpcPort() int32 {
	if o == nil || o.YqlServerRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.YqlServerRpcPort
}

// GetYqlServerRpcPortOk returns a tuple with the YqlServerRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetYqlServerRpcPortOk() (*int32, bool) {
	if o == nil || o.YqlServerRpcPort == nil {
		return nil, false
	}
	return o.YqlServerRpcPort, true
}

// HasYqlServerRpcPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasYqlServerRpcPort() bool {
	if o != nil && o.YqlServerRpcPort != nil {
		return true
	}

	return false
}

// SetYqlServerRpcPort gets a reference to the given int32 and assigns it to the YqlServerRpcPort field.
func (o *CommunicationPorts) SetYqlServerRpcPort(v int32) {
	o.YqlServerRpcPort = &v
}

// GetYsqlServerHttpPort returns the YsqlServerHttpPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetYsqlServerHttpPort() int32 {
	if o == nil || o.YsqlServerHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.YsqlServerHttpPort
}

// GetYsqlServerHttpPortOk returns a tuple with the YsqlServerHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetYsqlServerHttpPortOk() (*int32, bool) {
	if o == nil || o.YsqlServerHttpPort == nil {
		return nil, false
	}
	return o.YsqlServerHttpPort, true
}

// HasYsqlServerHttpPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasYsqlServerHttpPort() bool {
	if o != nil && o.YsqlServerHttpPort != nil {
		return true
	}

	return false
}

// SetYsqlServerHttpPort gets a reference to the given int32 and assigns it to the YsqlServerHttpPort field.
func (o *CommunicationPorts) SetYsqlServerHttpPort(v int32) {
	o.YsqlServerHttpPort = &v
}

// GetYsqlServerRpcPort returns the YsqlServerRpcPort field value if set, zero value otherwise.
func (o *CommunicationPorts) GetYsqlServerRpcPort() int32 {
	if o == nil || o.YsqlServerRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.YsqlServerRpcPort
}

// GetYsqlServerRpcPortOk returns a tuple with the YsqlServerRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationPorts) GetYsqlServerRpcPortOk() (*int32, bool) {
	if o == nil || o.YsqlServerRpcPort == nil {
		return nil, false
	}
	return o.YsqlServerRpcPort, true
}

// HasYsqlServerRpcPort returns a boolean if a field has been set.
func (o *CommunicationPorts) HasYsqlServerRpcPort() bool {
	if o != nil && o.YsqlServerRpcPort != nil {
		return true
	}

	return false
}

// SetYsqlServerRpcPort gets a reference to the given int32 and assigns it to the YsqlServerRpcPort field.
func (o *CommunicationPorts) SetYsqlServerRpcPort(v int32) {
	o.YsqlServerRpcPort = &v
}

func (o CommunicationPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MasterHttpPort != nil {
		toSerialize["masterHttpPort"] = o.MasterHttpPort
	}
	if o.MasterRpcPort != nil {
		toSerialize["masterRpcPort"] = o.MasterRpcPort
	}
	if o.NodeExporterPort != nil {
		toSerialize["nodeExporterPort"] = o.NodeExporterPort
	}
	if o.RedisServerHttpPort != nil {
		toSerialize["redisServerHttpPort"] = o.RedisServerHttpPort
	}
	if o.RedisServerRpcPort != nil {
		toSerialize["redisServerRpcPort"] = o.RedisServerRpcPort
	}
	if o.TserverHttpPort != nil {
		toSerialize["tserverHttpPort"] = o.TserverHttpPort
	}
	if o.TserverRpcPort != nil {
		toSerialize["tserverRpcPort"] = o.TserverRpcPort
	}
	if o.YqlServerHttpPort != nil {
		toSerialize["yqlServerHttpPort"] = o.YqlServerHttpPort
	}
	if o.YqlServerRpcPort != nil {
		toSerialize["yqlServerRpcPort"] = o.YqlServerRpcPort
	}
	if o.YsqlServerHttpPort != nil {
		toSerialize["ysqlServerHttpPort"] = o.YsqlServerHttpPort
	}
	if o.YsqlServerRpcPort != nil {
		toSerialize["ysqlServerRpcPort"] = o.YsqlServerRpcPort
	}
	return json.Marshal(toSerialize)
}

type NullableCommunicationPorts struct {
	value *CommunicationPorts
	isSet bool
}

func (v NullableCommunicationPorts) Get() *CommunicationPorts {
	return v.value
}

func (v *NullableCommunicationPorts) Set(val *CommunicationPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationPorts(val *CommunicationPorts) *NullableCommunicationPorts {
	return &NullableCommunicationPorts{value: val, isSet: true}
}

func (v NullableCommunicationPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


