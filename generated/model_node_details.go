/*
 * Yugabyte Platform APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yb_platform_client

import (
	"encoding/json"
)

// NodeDetails Details of a cloud node
type NodeDetails struct {
	// The availability zone's UUID
	AzUuid *string `json:"azUuid,omitempty"`
	CloudInfo *CloudSpecificInfo `json:"cloudInfo,omitempty"`
	// True if cron jobs were properly configured for this node
	CronsActive *bool `json:"cronsActive,omitempty"`
	// True if this node is a master
	IsMaster *bool `json:"isMaster,omitempty"`
	// True if this node is a REDIS server
	IsRedisServer *bool `json:"isRedisServer,omitempty"`
	// True if this node is a Tablet server
	IsTserver *bool `json:"isTserver,omitempty"`
	// True if this node is a YCQL server
	IsYqlServer *bool `json:"isYqlServer,omitempty"`
	// True if this node is a YSQL server
	IsYsqlServer *bool `json:"isYsqlServer,omitempty"`
	// Machine image name
	MachineImage *string `json:"machineImage,omitempty"`
	// Master HTTP port
	MasterHttpPort *int32 `json:"masterHttpPort,omitempty"`
	// Master RCP port
	MasterRpcPort *int32 `json:"masterRpcPort,omitempty"`
	// Master state
	MasterState *string `json:"masterState,omitempty"`
	// Node exporter port
	NodeExporterPort *int32 `json:"nodeExporterPort,omitempty"`
	// Node ID
	NodeIdx *int32 `json:"nodeIdx,omitempty"`
	// Node name
	NodeName *string `json:"nodeName,omitempty"`
	NodeUUID string `json:"nodeUUID"`
	// Node UUID
	NodeUuid *string `json:"nodeUuid,omitempty"`
	// UUID of the cluster to which this node belongs
	PlacementUuid *string `json:"placementUuid,omitempty"`
	// REDIS HTTP port
	RedisServerHttpPort *int32 `json:"redisServerHttpPort,omitempty"`
	// REDIS RPC port
	RedisServerRpcPort *int32 `json:"redisServerRpcPort,omitempty"`
	// Node state
	State *string `json:"state,omitempty"`
	// Tablet server HTTP port
	TserverHttpPort *int32 `json:"tserverHttpPort,omitempty"`
	// Tablet server RPC port
	TserverRpcPort *int32 `json:"tserverRpcPort,omitempty"`
	// YCQL HTTP port
	YqlServerHttpPort *int32 `json:"yqlServerHttpPort,omitempty"`
	// YCQL RPC port
	YqlServerRpcPort *int32 `json:"yqlServerRpcPort,omitempty"`
	// YSQL HTTP port
	YsqlServerHttpPort *int32 `json:"ysqlServerHttpPort,omitempty"`
	// YSQL RPC port
	YsqlServerRpcPort *int32 `json:"ysqlServerRpcPort,omitempty"`
}

// NewNodeDetails instantiates a new NodeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeDetails(nodeUUID string, ) *NodeDetails {
	this := NodeDetails{}
	this.NodeUUID = nodeUUID
	return &this
}

// NewNodeDetailsWithDefaults instantiates a new NodeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeDetailsWithDefaults() *NodeDetails {
	this := NodeDetails{}
	return &this
}

// GetAzUuid returns the AzUuid field value if set, zero value otherwise.
func (o *NodeDetails) GetAzUuid() string {
	if o == nil || o.AzUuid == nil {
		var ret string
		return ret
	}
	return *o.AzUuid
}

// GetAzUuidOk returns a tuple with the AzUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetAzUuidOk() (*string, bool) {
	if o == nil || o.AzUuid == nil {
		return nil, false
	}
	return o.AzUuid, true
}

// HasAzUuid returns a boolean if a field has been set.
func (o *NodeDetails) HasAzUuid() bool {
	if o != nil && o.AzUuid != nil {
		return true
	}

	return false
}

// SetAzUuid gets a reference to the given string and assigns it to the AzUuid field.
func (o *NodeDetails) SetAzUuid(v string) {
	o.AzUuid = &v
}

// GetCloudInfo returns the CloudInfo field value if set, zero value otherwise.
func (o *NodeDetails) GetCloudInfo() CloudSpecificInfo {
	if o == nil || o.CloudInfo == nil {
		var ret CloudSpecificInfo
		return ret
	}
	return *o.CloudInfo
}

// GetCloudInfoOk returns a tuple with the CloudInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetCloudInfoOk() (*CloudSpecificInfo, bool) {
	if o == nil || o.CloudInfo == nil {
		return nil, false
	}
	return o.CloudInfo, true
}

// HasCloudInfo returns a boolean if a field has been set.
func (o *NodeDetails) HasCloudInfo() bool {
	if o != nil && o.CloudInfo != nil {
		return true
	}

	return false
}

// SetCloudInfo gets a reference to the given CloudSpecificInfo and assigns it to the CloudInfo field.
func (o *NodeDetails) SetCloudInfo(v CloudSpecificInfo) {
	o.CloudInfo = &v
}

// GetCronsActive returns the CronsActive field value if set, zero value otherwise.
func (o *NodeDetails) GetCronsActive() bool {
	if o == nil || o.CronsActive == nil {
		var ret bool
		return ret
	}
	return *o.CronsActive
}

// GetCronsActiveOk returns a tuple with the CronsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetCronsActiveOk() (*bool, bool) {
	if o == nil || o.CronsActive == nil {
		return nil, false
	}
	return o.CronsActive, true
}

// HasCronsActive returns a boolean if a field has been set.
func (o *NodeDetails) HasCronsActive() bool {
	if o != nil && o.CronsActive != nil {
		return true
	}

	return false
}

// SetCronsActive gets a reference to the given bool and assigns it to the CronsActive field.
func (o *NodeDetails) SetCronsActive(v bool) {
	o.CronsActive = &v
}

// GetIsMaster returns the IsMaster field value if set, zero value otherwise.
func (o *NodeDetails) GetIsMaster() bool {
	if o == nil || o.IsMaster == nil {
		var ret bool
		return ret
	}
	return *o.IsMaster
}

// GetIsMasterOk returns a tuple with the IsMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetIsMasterOk() (*bool, bool) {
	if o == nil || o.IsMaster == nil {
		return nil, false
	}
	return o.IsMaster, true
}

// HasIsMaster returns a boolean if a field has been set.
func (o *NodeDetails) HasIsMaster() bool {
	if o != nil && o.IsMaster != nil {
		return true
	}

	return false
}

// SetIsMaster gets a reference to the given bool and assigns it to the IsMaster field.
func (o *NodeDetails) SetIsMaster(v bool) {
	o.IsMaster = &v
}

// GetIsRedisServer returns the IsRedisServer field value if set, zero value otherwise.
func (o *NodeDetails) GetIsRedisServer() bool {
	if o == nil || o.IsRedisServer == nil {
		var ret bool
		return ret
	}
	return *o.IsRedisServer
}

// GetIsRedisServerOk returns a tuple with the IsRedisServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetIsRedisServerOk() (*bool, bool) {
	if o == nil || o.IsRedisServer == nil {
		return nil, false
	}
	return o.IsRedisServer, true
}

// HasIsRedisServer returns a boolean if a field has been set.
func (o *NodeDetails) HasIsRedisServer() bool {
	if o != nil && o.IsRedisServer != nil {
		return true
	}

	return false
}

// SetIsRedisServer gets a reference to the given bool and assigns it to the IsRedisServer field.
func (o *NodeDetails) SetIsRedisServer(v bool) {
	o.IsRedisServer = &v
}

// GetIsTserver returns the IsTserver field value if set, zero value otherwise.
func (o *NodeDetails) GetIsTserver() bool {
	if o == nil || o.IsTserver == nil {
		var ret bool
		return ret
	}
	return *o.IsTserver
}

// GetIsTserverOk returns a tuple with the IsTserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetIsTserverOk() (*bool, bool) {
	if o == nil || o.IsTserver == nil {
		return nil, false
	}
	return o.IsTserver, true
}

// HasIsTserver returns a boolean if a field has been set.
func (o *NodeDetails) HasIsTserver() bool {
	if o != nil && o.IsTserver != nil {
		return true
	}

	return false
}

// SetIsTserver gets a reference to the given bool and assigns it to the IsTserver field.
func (o *NodeDetails) SetIsTserver(v bool) {
	o.IsTserver = &v
}

// GetIsYqlServer returns the IsYqlServer field value if set, zero value otherwise.
func (o *NodeDetails) GetIsYqlServer() bool {
	if o == nil || o.IsYqlServer == nil {
		var ret bool
		return ret
	}
	return *o.IsYqlServer
}

// GetIsYqlServerOk returns a tuple with the IsYqlServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetIsYqlServerOk() (*bool, bool) {
	if o == nil || o.IsYqlServer == nil {
		return nil, false
	}
	return o.IsYqlServer, true
}

// HasIsYqlServer returns a boolean if a field has been set.
func (o *NodeDetails) HasIsYqlServer() bool {
	if o != nil && o.IsYqlServer != nil {
		return true
	}

	return false
}

// SetIsYqlServer gets a reference to the given bool and assigns it to the IsYqlServer field.
func (o *NodeDetails) SetIsYqlServer(v bool) {
	o.IsYqlServer = &v
}

// GetIsYsqlServer returns the IsYsqlServer field value if set, zero value otherwise.
func (o *NodeDetails) GetIsYsqlServer() bool {
	if o == nil || o.IsYsqlServer == nil {
		var ret bool
		return ret
	}
	return *o.IsYsqlServer
}

// GetIsYsqlServerOk returns a tuple with the IsYsqlServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetIsYsqlServerOk() (*bool, bool) {
	if o == nil || o.IsYsqlServer == nil {
		return nil, false
	}
	return o.IsYsqlServer, true
}

// HasIsYsqlServer returns a boolean if a field has been set.
func (o *NodeDetails) HasIsYsqlServer() bool {
	if o != nil && o.IsYsqlServer != nil {
		return true
	}

	return false
}

// SetIsYsqlServer gets a reference to the given bool and assigns it to the IsYsqlServer field.
func (o *NodeDetails) SetIsYsqlServer(v bool) {
	o.IsYsqlServer = &v
}

// GetMachineImage returns the MachineImage field value if set, zero value otherwise.
func (o *NodeDetails) GetMachineImage() string {
	if o == nil || o.MachineImage == nil {
		var ret string
		return ret
	}
	return *o.MachineImage
}

// GetMachineImageOk returns a tuple with the MachineImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetMachineImageOk() (*string, bool) {
	if o == nil || o.MachineImage == nil {
		return nil, false
	}
	return o.MachineImage, true
}

// HasMachineImage returns a boolean if a field has been set.
func (o *NodeDetails) HasMachineImage() bool {
	if o != nil && o.MachineImage != nil {
		return true
	}

	return false
}

// SetMachineImage gets a reference to the given string and assigns it to the MachineImage field.
func (o *NodeDetails) SetMachineImage(v string) {
	o.MachineImage = &v
}

// GetMasterHttpPort returns the MasterHttpPort field value if set, zero value otherwise.
func (o *NodeDetails) GetMasterHttpPort() int32 {
	if o == nil || o.MasterHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.MasterHttpPort
}

// GetMasterHttpPortOk returns a tuple with the MasterHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetMasterHttpPortOk() (*int32, bool) {
	if o == nil || o.MasterHttpPort == nil {
		return nil, false
	}
	return o.MasterHttpPort, true
}

// HasMasterHttpPort returns a boolean if a field has been set.
func (o *NodeDetails) HasMasterHttpPort() bool {
	if o != nil && o.MasterHttpPort != nil {
		return true
	}

	return false
}

// SetMasterHttpPort gets a reference to the given int32 and assigns it to the MasterHttpPort field.
func (o *NodeDetails) SetMasterHttpPort(v int32) {
	o.MasterHttpPort = &v
}

// GetMasterRpcPort returns the MasterRpcPort field value if set, zero value otherwise.
func (o *NodeDetails) GetMasterRpcPort() int32 {
	if o == nil || o.MasterRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.MasterRpcPort
}

// GetMasterRpcPortOk returns a tuple with the MasterRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetMasterRpcPortOk() (*int32, bool) {
	if o == nil || o.MasterRpcPort == nil {
		return nil, false
	}
	return o.MasterRpcPort, true
}

// HasMasterRpcPort returns a boolean if a field has been set.
func (o *NodeDetails) HasMasterRpcPort() bool {
	if o != nil && o.MasterRpcPort != nil {
		return true
	}

	return false
}

// SetMasterRpcPort gets a reference to the given int32 and assigns it to the MasterRpcPort field.
func (o *NodeDetails) SetMasterRpcPort(v int32) {
	o.MasterRpcPort = &v
}

// GetMasterState returns the MasterState field value if set, zero value otherwise.
func (o *NodeDetails) GetMasterState() string {
	if o == nil || o.MasterState == nil {
		var ret string
		return ret
	}
	return *o.MasterState
}

// GetMasterStateOk returns a tuple with the MasterState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetMasterStateOk() (*string, bool) {
	if o == nil || o.MasterState == nil {
		return nil, false
	}
	return o.MasterState, true
}

// HasMasterState returns a boolean if a field has been set.
func (o *NodeDetails) HasMasterState() bool {
	if o != nil && o.MasterState != nil {
		return true
	}

	return false
}

// SetMasterState gets a reference to the given string and assigns it to the MasterState field.
func (o *NodeDetails) SetMasterState(v string) {
	o.MasterState = &v
}

// GetNodeExporterPort returns the NodeExporterPort field value if set, zero value otherwise.
func (o *NodeDetails) GetNodeExporterPort() int32 {
	if o == nil || o.NodeExporterPort == nil {
		var ret int32
		return ret
	}
	return *o.NodeExporterPort
}

// GetNodeExporterPortOk returns a tuple with the NodeExporterPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetNodeExporterPortOk() (*int32, bool) {
	if o == nil || o.NodeExporterPort == nil {
		return nil, false
	}
	return o.NodeExporterPort, true
}

// HasNodeExporterPort returns a boolean if a field has been set.
func (o *NodeDetails) HasNodeExporterPort() bool {
	if o != nil && o.NodeExporterPort != nil {
		return true
	}

	return false
}

// SetNodeExporterPort gets a reference to the given int32 and assigns it to the NodeExporterPort field.
func (o *NodeDetails) SetNodeExporterPort(v int32) {
	o.NodeExporterPort = &v
}

// GetNodeIdx returns the NodeIdx field value if set, zero value otherwise.
func (o *NodeDetails) GetNodeIdx() int32 {
	if o == nil || o.NodeIdx == nil {
		var ret int32
		return ret
	}
	return *o.NodeIdx
}

// GetNodeIdxOk returns a tuple with the NodeIdx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetNodeIdxOk() (*int32, bool) {
	if o == nil || o.NodeIdx == nil {
		return nil, false
	}
	return o.NodeIdx, true
}

// HasNodeIdx returns a boolean if a field has been set.
func (o *NodeDetails) HasNodeIdx() bool {
	if o != nil && o.NodeIdx != nil {
		return true
	}

	return false
}

// SetNodeIdx gets a reference to the given int32 and assigns it to the NodeIdx field.
func (o *NodeDetails) SetNodeIdx(v int32) {
	o.NodeIdx = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *NodeDetails) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *NodeDetails) HasNodeName() bool {
	if o != nil && o.NodeName != nil {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *NodeDetails) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeUUID returns the NodeUUID field value
func (o *NodeDetails) GetNodeUUID() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.NodeUUID
}

// GetNodeUUIDOk returns a tuple with the NodeUUID field value
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetNodeUUIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeUUID, true
}

// SetNodeUUID sets field value
func (o *NodeDetails) SetNodeUUID(v string) {
	o.NodeUUID = v
}

// GetNodeUuid returns the NodeUuid field value if set, zero value otherwise.
func (o *NodeDetails) GetNodeUuid() string {
	if o == nil || o.NodeUuid == nil {
		var ret string
		return ret
	}
	return *o.NodeUuid
}

// GetNodeUuidOk returns a tuple with the NodeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetNodeUuidOk() (*string, bool) {
	if o == nil || o.NodeUuid == nil {
		return nil, false
	}
	return o.NodeUuid, true
}

// HasNodeUuid returns a boolean if a field has been set.
func (o *NodeDetails) HasNodeUuid() bool {
	if o != nil && o.NodeUuid != nil {
		return true
	}

	return false
}

// SetNodeUuid gets a reference to the given string and assigns it to the NodeUuid field.
func (o *NodeDetails) SetNodeUuid(v string) {
	o.NodeUuid = &v
}

// GetPlacementUuid returns the PlacementUuid field value if set, zero value otherwise.
func (o *NodeDetails) GetPlacementUuid() string {
	if o == nil || o.PlacementUuid == nil {
		var ret string
		return ret
	}
	return *o.PlacementUuid
}

// GetPlacementUuidOk returns a tuple with the PlacementUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetPlacementUuidOk() (*string, bool) {
	if o == nil || o.PlacementUuid == nil {
		return nil, false
	}
	return o.PlacementUuid, true
}

// HasPlacementUuid returns a boolean if a field has been set.
func (o *NodeDetails) HasPlacementUuid() bool {
	if o != nil && o.PlacementUuid != nil {
		return true
	}

	return false
}

// SetPlacementUuid gets a reference to the given string and assigns it to the PlacementUuid field.
func (o *NodeDetails) SetPlacementUuid(v string) {
	o.PlacementUuid = &v
}

// GetRedisServerHttpPort returns the RedisServerHttpPort field value if set, zero value otherwise.
func (o *NodeDetails) GetRedisServerHttpPort() int32 {
	if o == nil || o.RedisServerHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.RedisServerHttpPort
}

// GetRedisServerHttpPortOk returns a tuple with the RedisServerHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetRedisServerHttpPortOk() (*int32, bool) {
	if o == nil || o.RedisServerHttpPort == nil {
		return nil, false
	}
	return o.RedisServerHttpPort, true
}

// HasRedisServerHttpPort returns a boolean if a field has been set.
func (o *NodeDetails) HasRedisServerHttpPort() bool {
	if o != nil && o.RedisServerHttpPort != nil {
		return true
	}

	return false
}

// SetRedisServerHttpPort gets a reference to the given int32 and assigns it to the RedisServerHttpPort field.
func (o *NodeDetails) SetRedisServerHttpPort(v int32) {
	o.RedisServerHttpPort = &v
}

// GetRedisServerRpcPort returns the RedisServerRpcPort field value if set, zero value otherwise.
func (o *NodeDetails) GetRedisServerRpcPort() int32 {
	if o == nil || o.RedisServerRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.RedisServerRpcPort
}

// GetRedisServerRpcPortOk returns a tuple with the RedisServerRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetRedisServerRpcPortOk() (*int32, bool) {
	if o == nil || o.RedisServerRpcPort == nil {
		return nil, false
	}
	return o.RedisServerRpcPort, true
}

// HasRedisServerRpcPort returns a boolean if a field has been set.
func (o *NodeDetails) HasRedisServerRpcPort() bool {
	if o != nil && o.RedisServerRpcPort != nil {
		return true
	}

	return false
}

// SetRedisServerRpcPort gets a reference to the given int32 and assigns it to the RedisServerRpcPort field.
func (o *NodeDetails) SetRedisServerRpcPort(v int32) {
	o.RedisServerRpcPort = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NodeDetails) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NodeDetails) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NodeDetails) SetState(v string) {
	o.State = &v
}

// GetTserverHttpPort returns the TserverHttpPort field value if set, zero value otherwise.
func (o *NodeDetails) GetTserverHttpPort() int32 {
	if o == nil || o.TserverHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.TserverHttpPort
}

// GetTserverHttpPortOk returns a tuple with the TserverHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetTserverHttpPortOk() (*int32, bool) {
	if o == nil || o.TserverHttpPort == nil {
		return nil, false
	}
	return o.TserverHttpPort, true
}

// HasTserverHttpPort returns a boolean if a field has been set.
func (o *NodeDetails) HasTserverHttpPort() bool {
	if o != nil && o.TserverHttpPort != nil {
		return true
	}

	return false
}

// SetTserverHttpPort gets a reference to the given int32 and assigns it to the TserverHttpPort field.
func (o *NodeDetails) SetTserverHttpPort(v int32) {
	o.TserverHttpPort = &v
}

// GetTserverRpcPort returns the TserverRpcPort field value if set, zero value otherwise.
func (o *NodeDetails) GetTserverRpcPort() int32 {
	if o == nil || o.TserverRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.TserverRpcPort
}

// GetTserverRpcPortOk returns a tuple with the TserverRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetTserverRpcPortOk() (*int32, bool) {
	if o == nil || o.TserverRpcPort == nil {
		return nil, false
	}
	return o.TserverRpcPort, true
}

// HasTserverRpcPort returns a boolean if a field has been set.
func (o *NodeDetails) HasTserverRpcPort() bool {
	if o != nil && o.TserverRpcPort != nil {
		return true
	}

	return false
}

// SetTserverRpcPort gets a reference to the given int32 and assigns it to the TserverRpcPort field.
func (o *NodeDetails) SetTserverRpcPort(v int32) {
	o.TserverRpcPort = &v
}

// GetYqlServerHttpPort returns the YqlServerHttpPort field value if set, zero value otherwise.
func (o *NodeDetails) GetYqlServerHttpPort() int32 {
	if o == nil || o.YqlServerHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.YqlServerHttpPort
}

// GetYqlServerHttpPortOk returns a tuple with the YqlServerHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetYqlServerHttpPortOk() (*int32, bool) {
	if o == nil || o.YqlServerHttpPort == nil {
		return nil, false
	}
	return o.YqlServerHttpPort, true
}

// HasYqlServerHttpPort returns a boolean if a field has been set.
func (o *NodeDetails) HasYqlServerHttpPort() bool {
	if o != nil && o.YqlServerHttpPort != nil {
		return true
	}

	return false
}

// SetYqlServerHttpPort gets a reference to the given int32 and assigns it to the YqlServerHttpPort field.
func (o *NodeDetails) SetYqlServerHttpPort(v int32) {
	o.YqlServerHttpPort = &v
}

// GetYqlServerRpcPort returns the YqlServerRpcPort field value if set, zero value otherwise.
func (o *NodeDetails) GetYqlServerRpcPort() int32 {
	if o == nil || o.YqlServerRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.YqlServerRpcPort
}

// GetYqlServerRpcPortOk returns a tuple with the YqlServerRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetYqlServerRpcPortOk() (*int32, bool) {
	if o == nil || o.YqlServerRpcPort == nil {
		return nil, false
	}
	return o.YqlServerRpcPort, true
}

// HasYqlServerRpcPort returns a boolean if a field has been set.
func (o *NodeDetails) HasYqlServerRpcPort() bool {
	if o != nil && o.YqlServerRpcPort != nil {
		return true
	}

	return false
}

// SetYqlServerRpcPort gets a reference to the given int32 and assigns it to the YqlServerRpcPort field.
func (o *NodeDetails) SetYqlServerRpcPort(v int32) {
	o.YqlServerRpcPort = &v
}

// GetYsqlServerHttpPort returns the YsqlServerHttpPort field value if set, zero value otherwise.
func (o *NodeDetails) GetYsqlServerHttpPort() int32 {
	if o == nil || o.YsqlServerHttpPort == nil {
		var ret int32
		return ret
	}
	return *o.YsqlServerHttpPort
}

// GetYsqlServerHttpPortOk returns a tuple with the YsqlServerHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetYsqlServerHttpPortOk() (*int32, bool) {
	if o == nil || o.YsqlServerHttpPort == nil {
		return nil, false
	}
	return o.YsqlServerHttpPort, true
}

// HasYsqlServerHttpPort returns a boolean if a field has been set.
func (o *NodeDetails) HasYsqlServerHttpPort() bool {
	if o != nil && o.YsqlServerHttpPort != nil {
		return true
	}

	return false
}

// SetYsqlServerHttpPort gets a reference to the given int32 and assigns it to the YsqlServerHttpPort field.
func (o *NodeDetails) SetYsqlServerHttpPort(v int32) {
	o.YsqlServerHttpPort = &v
}

// GetYsqlServerRpcPort returns the YsqlServerRpcPort field value if set, zero value otherwise.
func (o *NodeDetails) GetYsqlServerRpcPort() int32 {
	if o == nil || o.YsqlServerRpcPort == nil {
		var ret int32
		return ret
	}
	return *o.YsqlServerRpcPort
}

// GetYsqlServerRpcPortOk returns a tuple with the YsqlServerRpcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeDetails) GetYsqlServerRpcPortOk() (*int32, bool) {
	if o == nil || o.YsqlServerRpcPort == nil {
		return nil, false
	}
	return o.YsqlServerRpcPort, true
}

// HasYsqlServerRpcPort returns a boolean if a field has been set.
func (o *NodeDetails) HasYsqlServerRpcPort() bool {
	if o != nil && o.YsqlServerRpcPort != nil {
		return true
	}

	return false
}

// SetYsqlServerRpcPort gets a reference to the given int32 and assigns it to the YsqlServerRpcPort field.
func (o *NodeDetails) SetYsqlServerRpcPort(v int32) {
	o.YsqlServerRpcPort = &v
}

func (o NodeDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzUuid != nil {
		toSerialize["azUuid"] = o.AzUuid
	}
	if o.CloudInfo != nil {
		toSerialize["cloudInfo"] = o.CloudInfo
	}
	if o.CronsActive != nil {
		toSerialize["cronsActive"] = o.CronsActive
	}
	if o.IsMaster != nil {
		toSerialize["isMaster"] = o.IsMaster
	}
	if o.IsRedisServer != nil {
		toSerialize["isRedisServer"] = o.IsRedisServer
	}
	if o.IsTserver != nil {
		toSerialize["isTserver"] = o.IsTserver
	}
	if o.IsYqlServer != nil {
		toSerialize["isYqlServer"] = o.IsYqlServer
	}
	if o.IsYsqlServer != nil {
		toSerialize["isYsqlServer"] = o.IsYsqlServer
	}
	if o.MachineImage != nil {
		toSerialize["machineImage"] = o.MachineImage
	}
	if o.MasterHttpPort != nil {
		toSerialize["masterHttpPort"] = o.MasterHttpPort
	}
	if o.MasterRpcPort != nil {
		toSerialize["masterRpcPort"] = o.MasterRpcPort
	}
	if o.MasterState != nil {
		toSerialize["masterState"] = o.MasterState
	}
	if o.NodeExporterPort != nil {
		toSerialize["nodeExporterPort"] = o.NodeExporterPort
	}
	if o.NodeIdx != nil {
		toSerialize["nodeIdx"] = o.NodeIdx
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if true {
		toSerialize["nodeUUID"] = o.NodeUUID
	}
	if o.NodeUuid != nil {
		toSerialize["nodeUuid"] = o.NodeUuid
	}
	if o.PlacementUuid != nil {
		toSerialize["placementUuid"] = o.PlacementUuid
	}
	if o.RedisServerHttpPort != nil {
		toSerialize["redisServerHttpPort"] = o.RedisServerHttpPort
	}
	if o.RedisServerRpcPort != nil {
		toSerialize["redisServerRpcPort"] = o.RedisServerRpcPort
	}
	if o.State != nil {
		toSerialize["state"] = o.State
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

type NullableNodeDetails struct {
	value *NodeDetails
	isSet bool
}

func (v NullableNodeDetails) Get() *NodeDetails {
	return v.value
}

func (v *NullableNodeDetails) Set(val *NodeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeDetails(val *NodeDetails) *NullableNodeDetails {
	return &NullableNodeDetails{value: val, isSet: true}
}

func (v NullableNodeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


