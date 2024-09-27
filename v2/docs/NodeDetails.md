# NodeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzUuid** | Pointer to **string** | The availability zone&#39;s UUID | [optional] 
**CloudInfo** | Pointer to [**CloudSpecificInfo**](CloudSpecificInfo.md) |  | [optional] 
**CronsActive** | Pointer to **bool** | True if cron jobs were properly configured for this node | [optional] 
**DedicatedTo** | Pointer to **string** | Used for configurations where each node can have only one process | [optional] 
**DisksAreMountedByUuid** | Pointer to **bool** | Disks are mounted by uuid | [optional] 
**IsMaster** | Pointer to **bool** | True if this node is a master | [optional] 
**IsRedisServer** | Pointer to **bool** | True if this node is a REDIS server | [optional] 
**IsTserver** | Pointer to **bool** | True if this node is a Tablet server | [optional] 
**IsYqlServer** | Pointer to **bool** | True if this node is a YCQL server | [optional] 
**IsYsqlServer** | Pointer to **bool** | True if this node is a YSQL server | [optional] 
**LastVolumeUpdateTime** | Pointer to **time.Time** | Store last volume update time | [optional] [readonly] 
**MachineImage** | Pointer to **string** | Machine image name | [optional] 
**MasterState** | Pointer to **string** | Master state | [optional] 
**NodeIdx** | Pointer to **int32** | Node ID | [optional] 
**NodeName** | Pointer to **string** | Node name | [optional] 
**NodeUuid** | Pointer to **string** | Node UUID | [optional] 
**PlacementUuid** | Pointer to **string** | UUID of the cluster to which this node belongs | [optional] 
**SshPortOverride** | Pointer to **int32** | SSH port override for the AMI | [optional] 
**SshUserOverride** | Pointer to **string** | SSH user override for the AMI | [optional] 
**State** | Pointer to **string** | Node state | [optional] 
**YbPrebuiltAmi** | Pointer to **bool** | True if this a custom YB AMI | [optional] 

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

### GetDisksAreMountedByUuid

`func (o *NodeDetails) GetDisksAreMountedByUuid() bool`

GetDisksAreMountedByUuid returns the DisksAreMountedByUuid field if non-nil, zero value otherwise.

### GetDisksAreMountedByUuidOk

`func (o *NodeDetails) GetDisksAreMountedByUuidOk() (*bool, bool)`

GetDisksAreMountedByUuidOk returns a tuple with the DisksAreMountedByUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksAreMountedByUuid

`func (o *NodeDetails) SetDisksAreMountedByUuid(v bool)`

SetDisksAreMountedByUuid sets DisksAreMountedByUuid field to given value.

### HasDisksAreMountedByUuid

`func (o *NodeDetails) HasDisksAreMountedByUuid() bool`

HasDisksAreMountedByUuid returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


