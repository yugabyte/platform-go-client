# NodeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSyncMasterAddrs** | Pointer to **bool** | WARNING: This is a preview API that could change. Used by auto master failover | [optional] [readonly] 
**AzUuid** | Pointer to **string** | The availability zone&#39;s UUID | [optional] 
**CloudInfo** | Pointer to [**CloudSpecificInfo**](CloudSpecificInfo.md) |  | [optional] 
**CronsActive** | Pointer to **bool** | True if cron jobs were properly configured for this node | [optional] 
**DedicatedTo** | Pointer to **string** | Used for configurations where each node can have only one process | [optional] 
**DisksAreMountedByUUID** | Pointer to **bool** | Disks are mounted by uuid | [optional] 
**InternalYsqlServerRpcPort** | Pointer to **int32** | Internal YSQL RPC port | [optional] 
**IsMaster** | Pointer to **bool** | True if this node is a master | [optional] 
**IsRedisServer** | Pointer to **bool** | True if this node is a REDIS server | [optional] 
**IsTserver** | Pointer to **bool** | True if this node is a Tablet server | [optional] 
**IsYqlServer** | Pointer to **bool** | True if this node is a YCQL server | [optional] 
**IsYsqlServer** | Pointer to **bool** | True if this node is a YSQL server | [optional] 
**LastVolumeUpdateTime** | Pointer to **time.Time** | Store last volume update time | [optional] [readonly] 
**MachineImage** | Pointer to **string** | Machine image name | [optional] 
**MasterHttpPort** | Pointer to **int32** | Master HTTP port | [optional] 
**MasterRpcPort** | Pointer to **int32** | Master RPC port | [optional] 
**MasterState** | Pointer to **string** | Master state | [optional] 
**NodeExporterPort** | Pointer to **int32** | Node exporter port | [optional] 
**NodeIdx** | Pointer to **int32** | Node ID | [optional] 
**NodeName** | Pointer to **string** | Node name | [optional] 
**NodeUuid** | Pointer to **string** | Node UUID | [optional] 
**OtelCollectorMetricsPort** | Pointer to **int32** | Otel collector metrics port | [optional] 
**PlacementUuid** | Pointer to **string** | UUID of the cluster to which this node belongs | [optional] 
**RedisServerHttpPort** | Pointer to **int32** | REDIS HTTP port | [optional] 
**RedisServerRpcPort** | Pointer to **int32** | REDIS RPC port | [optional] 
**SshPortOverride** | Pointer to **int32** | SSH port override for the AMI | [optional] 
**SshUserOverride** | Pointer to **string** | SSH user override for the AMI | [optional] 
**State** | Pointer to **string** | Node state | [optional] 
**TserverHttpPort** | Pointer to **int32** | Tablet server HTTP port | [optional] 
**TserverRpcPort** | Pointer to **int32** | Tablet server RPC port | [optional] 
**YbControllerHttpPort** | Pointer to **int32** | Yb controller HTTP port | [optional] 
**YbControllerRpcPort** | Pointer to **int32** | Yb controller RPC port | [optional] 
**YbPrebuiltAmi** | Pointer to **bool** | True if this a custom YB AMI | [optional] 
**YqlServerHttpPort** | Pointer to **int32** | YCQL HTTP port | [optional] 
**YqlServerRpcPort** | Pointer to **int32** | YCQL RPC port | [optional] 
**YsqlServerHttpPort** | Pointer to **int32** | YSQL HTTP port | [optional] 
**YsqlServerRpcPort** | Pointer to **int32** | YSQL RPC port | [optional] 

## Methods

### NewNodeDetails

`func NewNodeDetails() *NodeDetails`

NewNodeDetails instantiates a new NodeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDetailsWithDefaults

`func NewNodeDetailsWithDefaults() *NodeDetails`

NewNodeDetailsWithDefaults instantiates a new NodeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSyncMasterAddrs

`func (o *NodeDetails) GetAutoSyncMasterAddrs() bool`

GetAutoSyncMasterAddrs returns the AutoSyncMasterAddrs field if non-nil, zero value otherwise.

### GetAutoSyncMasterAddrsOk

`func (o *NodeDetails) GetAutoSyncMasterAddrsOk() (*bool, bool)`

GetAutoSyncMasterAddrsOk returns a tuple with the AutoSyncMasterAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncMasterAddrs

`func (o *NodeDetails) SetAutoSyncMasterAddrs(v bool)`

SetAutoSyncMasterAddrs sets AutoSyncMasterAddrs field to given value.

### HasAutoSyncMasterAddrs

`func (o *NodeDetails) HasAutoSyncMasterAddrs() bool`

HasAutoSyncMasterAddrs returns a boolean if a field has been set.

### GetAzUuid

`func (o *NodeDetails) GetAzUuid() string`

GetAzUuid returns the AzUuid field if non-nil, zero value otherwise.

### GetAzUuidOk

`func (o *NodeDetails) GetAzUuidOk() (*string, bool)`

GetAzUuidOk returns a tuple with the AzUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzUuid

`func (o *NodeDetails) SetAzUuid(v string)`

SetAzUuid sets AzUuid field to given value.

### HasAzUuid

`func (o *NodeDetails) HasAzUuid() bool`

HasAzUuid returns a boolean if a field has been set.

### GetCloudInfo

`func (o *NodeDetails) GetCloudInfo() CloudSpecificInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *NodeDetails) GetCloudInfoOk() (*CloudSpecificInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *NodeDetails) SetCloudInfo(v CloudSpecificInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *NodeDetails) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCronsActive

`func (o *NodeDetails) GetCronsActive() bool`

GetCronsActive returns the CronsActive field if non-nil, zero value otherwise.

### GetCronsActiveOk

`func (o *NodeDetails) GetCronsActiveOk() (*bool, bool)`

GetCronsActiveOk returns a tuple with the CronsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronsActive

`func (o *NodeDetails) SetCronsActive(v bool)`

SetCronsActive sets CronsActive field to given value.

### HasCronsActive

`func (o *NodeDetails) HasCronsActive() bool`

HasCronsActive returns a boolean if a field has been set.

### GetDedicatedTo

`func (o *NodeDetails) GetDedicatedTo() string`

GetDedicatedTo returns the DedicatedTo field if non-nil, zero value otherwise.

### GetDedicatedToOk

`func (o *NodeDetails) GetDedicatedToOk() (*string, bool)`

GetDedicatedToOk returns a tuple with the DedicatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTo

`func (o *NodeDetails) SetDedicatedTo(v string)`

SetDedicatedTo sets DedicatedTo field to given value.

### HasDedicatedTo

`func (o *NodeDetails) HasDedicatedTo() bool`

HasDedicatedTo returns a boolean if a field has been set.

### GetDisksAreMountedByUUID

`func (o *NodeDetails) GetDisksAreMountedByUUID() bool`

GetDisksAreMountedByUUID returns the DisksAreMountedByUUID field if non-nil, zero value otherwise.

### GetDisksAreMountedByUUIDOk

`func (o *NodeDetails) GetDisksAreMountedByUUIDOk() (*bool, bool)`

GetDisksAreMountedByUUIDOk returns a tuple with the DisksAreMountedByUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksAreMountedByUUID

`func (o *NodeDetails) SetDisksAreMountedByUUID(v bool)`

SetDisksAreMountedByUUID sets DisksAreMountedByUUID field to given value.

### HasDisksAreMountedByUUID

`func (o *NodeDetails) HasDisksAreMountedByUUID() bool`

HasDisksAreMountedByUUID returns a boolean if a field has been set.

### GetInternalYsqlServerRpcPort

`func (o *NodeDetails) GetInternalYsqlServerRpcPort() int32`

GetInternalYsqlServerRpcPort returns the InternalYsqlServerRpcPort field if non-nil, zero value otherwise.

### GetInternalYsqlServerRpcPortOk

`func (o *NodeDetails) GetInternalYsqlServerRpcPortOk() (*int32, bool)`

GetInternalYsqlServerRpcPortOk returns a tuple with the InternalYsqlServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalYsqlServerRpcPort

`func (o *NodeDetails) SetInternalYsqlServerRpcPort(v int32)`

SetInternalYsqlServerRpcPort sets InternalYsqlServerRpcPort field to given value.

### HasInternalYsqlServerRpcPort

`func (o *NodeDetails) HasInternalYsqlServerRpcPort() bool`

HasInternalYsqlServerRpcPort returns a boolean if a field has been set.

### GetIsMaster

`func (o *NodeDetails) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *NodeDetails) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *NodeDetails) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *NodeDetails) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### GetIsRedisServer

`func (o *NodeDetails) GetIsRedisServer() bool`

GetIsRedisServer returns the IsRedisServer field if non-nil, zero value otherwise.

### GetIsRedisServerOk

`func (o *NodeDetails) GetIsRedisServerOk() (*bool, bool)`

GetIsRedisServerOk returns a tuple with the IsRedisServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRedisServer

`func (o *NodeDetails) SetIsRedisServer(v bool)`

SetIsRedisServer sets IsRedisServer field to given value.

### HasIsRedisServer

`func (o *NodeDetails) HasIsRedisServer() bool`

HasIsRedisServer returns a boolean if a field has been set.

### GetIsTserver

`func (o *NodeDetails) GetIsTserver() bool`

GetIsTserver returns the IsTserver field if non-nil, zero value otherwise.

### GetIsTserverOk

`func (o *NodeDetails) GetIsTserverOk() (*bool, bool)`

GetIsTserverOk returns a tuple with the IsTserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTserver

`func (o *NodeDetails) SetIsTserver(v bool)`

SetIsTserver sets IsTserver field to given value.

### HasIsTserver

`func (o *NodeDetails) HasIsTserver() bool`

HasIsTserver returns a boolean if a field has been set.

### GetIsYqlServer

`func (o *NodeDetails) GetIsYqlServer() bool`

GetIsYqlServer returns the IsYqlServer field if non-nil, zero value otherwise.

### GetIsYqlServerOk

`func (o *NodeDetails) GetIsYqlServerOk() (*bool, bool)`

GetIsYqlServerOk returns a tuple with the IsYqlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYqlServer

`func (o *NodeDetails) SetIsYqlServer(v bool)`

SetIsYqlServer sets IsYqlServer field to given value.

### HasIsYqlServer

`func (o *NodeDetails) HasIsYqlServer() bool`

HasIsYqlServer returns a boolean if a field has been set.

### GetIsYsqlServer

`func (o *NodeDetails) GetIsYsqlServer() bool`

GetIsYsqlServer returns the IsYsqlServer field if non-nil, zero value otherwise.

### GetIsYsqlServerOk

`func (o *NodeDetails) GetIsYsqlServerOk() (*bool, bool)`

GetIsYsqlServerOk returns a tuple with the IsYsqlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYsqlServer

`func (o *NodeDetails) SetIsYsqlServer(v bool)`

SetIsYsqlServer sets IsYsqlServer field to given value.

### HasIsYsqlServer

`func (o *NodeDetails) HasIsYsqlServer() bool`

HasIsYsqlServer returns a boolean if a field has been set.

### GetLastVolumeUpdateTime

`func (o *NodeDetails) GetLastVolumeUpdateTime() time.Time`

GetLastVolumeUpdateTime returns the LastVolumeUpdateTime field if non-nil, zero value otherwise.

### GetLastVolumeUpdateTimeOk

`func (o *NodeDetails) GetLastVolumeUpdateTimeOk() (*time.Time, bool)`

GetLastVolumeUpdateTimeOk returns a tuple with the LastVolumeUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVolumeUpdateTime

`func (o *NodeDetails) SetLastVolumeUpdateTime(v time.Time)`

SetLastVolumeUpdateTime sets LastVolumeUpdateTime field to given value.

### HasLastVolumeUpdateTime

`func (o *NodeDetails) HasLastVolumeUpdateTime() bool`

HasLastVolumeUpdateTime returns a boolean if a field has been set.

### GetMachineImage

`func (o *NodeDetails) GetMachineImage() string`

GetMachineImage returns the MachineImage field if non-nil, zero value otherwise.

### GetMachineImageOk

`func (o *NodeDetails) GetMachineImageOk() (*string, bool)`

GetMachineImageOk returns a tuple with the MachineImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineImage

`func (o *NodeDetails) SetMachineImage(v string)`

SetMachineImage sets MachineImage field to given value.

### HasMachineImage

`func (o *NodeDetails) HasMachineImage() bool`

HasMachineImage returns a boolean if a field has been set.

### GetMasterHttpPort

`func (o *NodeDetails) GetMasterHttpPort() int32`

GetMasterHttpPort returns the MasterHttpPort field if non-nil, zero value otherwise.

### GetMasterHttpPortOk

`func (o *NodeDetails) GetMasterHttpPortOk() (*int32, bool)`

GetMasterHttpPortOk returns a tuple with the MasterHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterHttpPort

`func (o *NodeDetails) SetMasterHttpPort(v int32)`

SetMasterHttpPort sets MasterHttpPort field to given value.

### HasMasterHttpPort

`func (o *NodeDetails) HasMasterHttpPort() bool`

HasMasterHttpPort returns a boolean if a field has been set.

### GetMasterRpcPort

`func (o *NodeDetails) GetMasterRpcPort() int32`

GetMasterRpcPort returns the MasterRpcPort field if non-nil, zero value otherwise.

### GetMasterRpcPortOk

`func (o *NodeDetails) GetMasterRpcPortOk() (*int32, bool)`

GetMasterRpcPortOk returns a tuple with the MasterRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRpcPort

`func (o *NodeDetails) SetMasterRpcPort(v int32)`

SetMasterRpcPort sets MasterRpcPort field to given value.

### HasMasterRpcPort

`func (o *NodeDetails) HasMasterRpcPort() bool`

HasMasterRpcPort returns a boolean if a field has been set.

### GetMasterState

`func (o *NodeDetails) GetMasterState() string`

GetMasterState returns the MasterState field if non-nil, zero value otherwise.

### GetMasterStateOk

`func (o *NodeDetails) GetMasterStateOk() (*string, bool)`

GetMasterStateOk returns a tuple with the MasterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterState

`func (o *NodeDetails) SetMasterState(v string)`

SetMasterState sets MasterState field to given value.

### HasMasterState

`func (o *NodeDetails) HasMasterState() bool`

HasMasterState returns a boolean if a field has been set.

### GetNodeExporterPort

`func (o *NodeDetails) GetNodeExporterPort() int32`

GetNodeExporterPort returns the NodeExporterPort field if non-nil, zero value otherwise.

### GetNodeExporterPortOk

`func (o *NodeDetails) GetNodeExporterPortOk() (*int32, bool)`

GetNodeExporterPortOk returns a tuple with the NodeExporterPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterPort

`func (o *NodeDetails) SetNodeExporterPort(v int32)`

SetNodeExporterPort sets NodeExporterPort field to given value.

### HasNodeExporterPort

`func (o *NodeDetails) HasNodeExporterPort() bool`

HasNodeExporterPort returns a boolean if a field has been set.

### GetNodeIdx

`func (o *NodeDetails) GetNodeIdx() int32`

GetNodeIdx returns the NodeIdx field if non-nil, zero value otherwise.

### GetNodeIdxOk

`func (o *NodeDetails) GetNodeIdxOk() (*int32, bool)`

GetNodeIdxOk returns a tuple with the NodeIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIdx

`func (o *NodeDetails) SetNodeIdx(v int32)`

SetNodeIdx sets NodeIdx field to given value.

### HasNodeIdx

`func (o *NodeDetails) HasNodeIdx() bool`

HasNodeIdx returns a boolean if a field has been set.

### GetNodeName

`func (o *NodeDetails) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *NodeDetails) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *NodeDetails) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *NodeDetails) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeUuid

`func (o *NodeDetails) GetNodeUuid() string`

GetNodeUuid returns the NodeUuid field if non-nil, zero value otherwise.

### GetNodeUuidOk

`func (o *NodeDetails) GetNodeUuidOk() (*string, bool)`

GetNodeUuidOk returns a tuple with the NodeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuid

`func (o *NodeDetails) SetNodeUuid(v string)`

SetNodeUuid sets NodeUuid field to given value.

### HasNodeUuid

`func (o *NodeDetails) HasNodeUuid() bool`

HasNodeUuid returns a boolean if a field has been set.

### GetOtelCollectorMetricsPort

`func (o *NodeDetails) GetOtelCollectorMetricsPort() int32`

GetOtelCollectorMetricsPort returns the OtelCollectorMetricsPort field if non-nil, zero value otherwise.

### GetOtelCollectorMetricsPortOk

`func (o *NodeDetails) GetOtelCollectorMetricsPortOk() (*int32, bool)`

GetOtelCollectorMetricsPortOk returns a tuple with the OtelCollectorMetricsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelCollectorMetricsPort

`func (o *NodeDetails) SetOtelCollectorMetricsPort(v int32)`

SetOtelCollectorMetricsPort sets OtelCollectorMetricsPort field to given value.

### HasOtelCollectorMetricsPort

`func (o *NodeDetails) HasOtelCollectorMetricsPort() bool`

HasOtelCollectorMetricsPort returns a boolean if a field has been set.

### GetPlacementUuid

`func (o *NodeDetails) GetPlacementUuid() string`

GetPlacementUuid returns the PlacementUuid field if non-nil, zero value otherwise.

### GetPlacementUuidOk

`func (o *NodeDetails) GetPlacementUuidOk() (*string, bool)`

GetPlacementUuidOk returns a tuple with the PlacementUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementUuid

`func (o *NodeDetails) SetPlacementUuid(v string)`

SetPlacementUuid sets PlacementUuid field to given value.

### HasPlacementUuid

`func (o *NodeDetails) HasPlacementUuid() bool`

HasPlacementUuid returns a boolean if a field has been set.

### GetRedisServerHttpPort

`func (o *NodeDetails) GetRedisServerHttpPort() int32`

GetRedisServerHttpPort returns the RedisServerHttpPort field if non-nil, zero value otherwise.

### GetRedisServerHttpPortOk

`func (o *NodeDetails) GetRedisServerHttpPortOk() (*int32, bool)`

GetRedisServerHttpPortOk returns a tuple with the RedisServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisServerHttpPort

`func (o *NodeDetails) SetRedisServerHttpPort(v int32)`

SetRedisServerHttpPort sets RedisServerHttpPort field to given value.

### HasRedisServerHttpPort

`func (o *NodeDetails) HasRedisServerHttpPort() bool`

HasRedisServerHttpPort returns a boolean if a field has been set.

### GetRedisServerRpcPort

`func (o *NodeDetails) GetRedisServerRpcPort() int32`

GetRedisServerRpcPort returns the RedisServerRpcPort field if non-nil, zero value otherwise.

### GetRedisServerRpcPortOk

`func (o *NodeDetails) GetRedisServerRpcPortOk() (*int32, bool)`

GetRedisServerRpcPortOk returns a tuple with the RedisServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisServerRpcPort

`func (o *NodeDetails) SetRedisServerRpcPort(v int32)`

SetRedisServerRpcPort sets RedisServerRpcPort field to given value.

### HasRedisServerRpcPort

`func (o *NodeDetails) HasRedisServerRpcPort() bool`

HasRedisServerRpcPort returns a boolean if a field has been set.

### GetSshPortOverride

`func (o *NodeDetails) GetSshPortOverride() int32`

GetSshPortOverride returns the SshPortOverride field if non-nil, zero value otherwise.

### GetSshPortOverrideOk

`func (o *NodeDetails) GetSshPortOverrideOk() (*int32, bool)`

GetSshPortOverrideOk returns a tuple with the SshPortOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPortOverride

`func (o *NodeDetails) SetSshPortOverride(v int32)`

SetSshPortOverride sets SshPortOverride field to given value.

### HasSshPortOverride

`func (o *NodeDetails) HasSshPortOverride() bool`

HasSshPortOverride returns a boolean if a field has been set.

### GetSshUserOverride

`func (o *NodeDetails) GetSshUserOverride() string`

GetSshUserOverride returns the SshUserOverride field if non-nil, zero value otherwise.

### GetSshUserOverrideOk

`func (o *NodeDetails) GetSshUserOverrideOk() (*string, bool)`

GetSshUserOverrideOk returns a tuple with the SshUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverride

`func (o *NodeDetails) SetSshUserOverride(v string)`

SetSshUserOverride sets SshUserOverride field to given value.

### HasSshUserOverride

`func (o *NodeDetails) HasSshUserOverride() bool`

HasSshUserOverride returns a boolean if a field has been set.

### GetState

`func (o *NodeDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NodeDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NodeDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NodeDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTserverHttpPort

`func (o *NodeDetails) GetTserverHttpPort() int32`

GetTserverHttpPort returns the TserverHttpPort field if non-nil, zero value otherwise.

### GetTserverHttpPortOk

`func (o *NodeDetails) GetTserverHttpPortOk() (*int32, bool)`

GetTserverHttpPortOk returns a tuple with the TserverHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverHttpPort

`func (o *NodeDetails) SetTserverHttpPort(v int32)`

SetTserverHttpPort sets TserverHttpPort field to given value.

### HasTserverHttpPort

`func (o *NodeDetails) HasTserverHttpPort() bool`

HasTserverHttpPort returns a boolean if a field has been set.

### GetTserverRpcPort

`func (o *NodeDetails) GetTserverRpcPort() int32`

GetTserverRpcPort returns the TserverRpcPort field if non-nil, zero value otherwise.

### GetTserverRpcPortOk

`func (o *NodeDetails) GetTserverRpcPortOk() (*int32, bool)`

GetTserverRpcPortOk returns a tuple with the TserverRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverRpcPort

`func (o *NodeDetails) SetTserverRpcPort(v int32)`

SetTserverRpcPort sets TserverRpcPort field to given value.

### HasTserverRpcPort

`func (o *NodeDetails) HasTserverRpcPort() bool`

HasTserverRpcPort returns a boolean if a field has been set.

### GetYbControllerHttpPort

`func (o *NodeDetails) GetYbControllerHttpPort() int32`

GetYbControllerHttpPort returns the YbControllerHttpPort field if non-nil, zero value otherwise.

### GetYbControllerHttpPortOk

`func (o *NodeDetails) GetYbControllerHttpPortOk() (*int32, bool)`

GetYbControllerHttpPortOk returns a tuple with the YbControllerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbControllerHttpPort

`func (o *NodeDetails) SetYbControllerHttpPort(v int32)`

SetYbControllerHttpPort sets YbControllerHttpPort field to given value.

### HasYbControllerHttpPort

`func (o *NodeDetails) HasYbControllerHttpPort() bool`

HasYbControllerHttpPort returns a boolean if a field has been set.

### GetYbControllerRpcPort

`func (o *NodeDetails) GetYbControllerRpcPort() int32`

GetYbControllerRpcPort returns the YbControllerRpcPort field if non-nil, zero value otherwise.

### GetYbControllerRpcPortOk

`func (o *NodeDetails) GetYbControllerRpcPortOk() (*int32, bool)`

GetYbControllerRpcPortOk returns a tuple with the YbControllerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbControllerRpcPort

`func (o *NodeDetails) SetYbControllerRpcPort(v int32)`

SetYbControllerRpcPort sets YbControllerRpcPort field to given value.

### HasYbControllerRpcPort

`func (o *NodeDetails) HasYbControllerRpcPort() bool`

HasYbControllerRpcPort returns a boolean if a field has been set.

### GetYbPrebuiltAmi

`func (o *NodeDetails) GetYbPrebuiltAmi() bool`

GetYbPrebuiltAmi returns the YbPrebuiltAmi field if non-nil, zero value otherwise.

### GetYbPrebuiltAmiOk

`func (o *NodeDetails) GetYbPrebuiltAmiOk() (*bool, bool)`

GetYbPrebuiltAmiOk returns a tuple with the YbPrebuiltAmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrebuiltAmi

`func (o *NodeDetails) SetYbPrebuiltAmi(v bool)`

SetYbPrebuiltAmi sets YbPrebuiltAmi field to given value.

### HasYbPrebuiltAmi

`func (o *NodeDetails) HasYbPrebuiltAmi() bool`

HasYbPrebuiltAmi returns a boolean if a field has been set.

### GetYqlServerHttpPort

`func (o *NodeDetails) GetYqlServerHttpPort() int32`

GetYqlServerHttpPort returns the YqlServerHttpPort field if non-nil, zero value otherwise.

### GetYqlServerHttpPortOk

`func (o *NodeDetails) GetYqlServerHttpPortOk() (*int32, bool)`

GetYqlServerHttpPortOk returns a tuple with the YqlServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYqlServerHttpPort

`func (o *NodeDetails) SetYqlServerHttpPort(v int32)`

SetYqlServerHttpPort sets YqlServerHttpPort field to given value.

### HasYqlServerHttpPort

`func (o *NodeDetails) HasYqlServerHttpPort() bool`

HasYqlServerHttpPort returns a boolean if a field has been set.

### GetYqlServerRpcPort

`func (o *NodeDetails) GetYqlServerRpcPort() int32`

GetYqlServerRpcPort returns the YqlServerRpcPort field if non-nil, zero value otherwise.

### GetYqlServerRpcPortOk

`func (o *NodeDetails) GetYqlServerRpcPortOk() (*int32, bool)`

GetYqlServerRpcPortOk returns a tuple with the YqlServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYqlServerRpcPort

`func (o *NodeDetails) SetYqlServerRpcPort(v int32)`

SetYqlServerRpcPort sets YqlServerRpcPort field to given value.

### HasYqlServerRpcPort

`func (o *NodeDetails) HasYqlServerRpcPort() bool`

HasYqlServerRpcPort returns a boolean if a field has been set.

### GetYsqlServerHttpPort

`func (o *NodeDetails) GetYsqlServerHttpPort() int32`

GetYsqlServerHttpPort returns the YsqlServerHttpPort field if non-nil, zero value otherwise.

### GetYsqlServerHttpPortOk

`func (o *NodeDetails) GetYsqlServerHttpPortOk() (*int32, bool)`

GetYsqlServerHttpPortOk returns a tuple with the YsqlServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlServerHttpPort

`func (o *NodeDetails) SetYsqlServerHttpPort(v int32)`

SetYsqlServerHttpPort sets YsqlServerHttpPort field to given value.

### HasYsqlServerHttpPort

`func (o *NodeDetails) HasYsqlServerHttpPort() bool`

HasYsqlServerHttpPort returns a boolean if a field has been set.

### GetYsqlServerRpcPort

`func (o *NodeDetails) GetYsqlServerRpcPort() int32`

GetYsqlServerRpcPort returns the YsqlServerRpcPort field if non-nil, zero value otherwise.

### GetYsqlServerRpcPortOk

`func (o *NodeDetails) GetYsqlServerRpcPortOk() (*int32, bool)`

GetYsqlServerRpcPortOk returns a tuple with the YsqlServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlServerRpcPort

`func (o *NodeDetails) SetYsqlServerRpcPort(v int32)`

SetYsqlServerRpcPort sets YsqlServerRpcPort field to given value.

### HasYsqlServerRpcPort

`func (o *NodeDetails) HasYsqlServerRpcPort() bool`

HasYsqlServerRpcPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


