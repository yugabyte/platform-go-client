# NodeInstanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | **string** | Node instance name | 
**InstanceType** | **string** | Node instance type | 
**Ip** | **string** | IP address | 
**NodeConfigs** | Pointer to [**[]NodeConfig**](NodeConfig.md) | Node configurations | [optional] 
**NodeName** | Pointer to **string** | Node name in a universe | [optional] [readonly] 
**Region** | **string** | Region | 
**SshUser** | **string** | SSH user | 
**Zone** | **string** | Zone | 

## Methods

### NewNodeInstanceData

`func NewNodeInstanceData(instanceName string, instanceType string, ip string, region string, sshUser string, zone string, ) *NodeInstanceData`

NewNodeInstanceData instantiates a new NodeInstanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInstanceDataWithDefaults

`func NewNodeInstanceDataWithDefaults() *NodeInstanceData`

NewNodeInstanceDataWithDefaults instantiates a new NodeInstanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *NodeInstanceData) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *NodeInstanceData) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *NodeInstanceData) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetInstanceType

`func (o *NodeInstanceData) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *NodeInstanceData) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *NodeInstanceData) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetIp

`func (o *NodeInstanceData) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NodeInstanceData) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NodeInstanceData) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNodeConfigs

`func (o *NodeInstanceData) GetNodeConfigs() []NodeConfig`

GetNodeConfigs returns the NodeConfigs field if non-nil, zero value otherwise.

### GetNodeConfigsOk

`func (o *NodeInstanceData) GetNodeConfigsOk() (*[]NodeConfig, bool)`

GetNodeConfigsOk returns a tuple with the NodeConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConfigs

`func (o *NodeInstanceData) SetNodeConfigs(v []NodeConfig)`

SetNodeConfigs sets NodeConfigs field to given value.

### HasNodeConfigs

`func (o *NodeInstanceData) HasNodeConfigs() bool`

HasNodeConfigs returns a boolean if a field has been set.

### GetNodeName

`func (o *NodeInstanceData) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *NodeInstanceData) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *NodeInstanceData) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *NodeInstanceData) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetRegion

`func (o *NodeInstanceData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NodeInstanceData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NodeInstanceData) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSshUser

`func (o *NodeInstanceData) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *NodeInstanceData) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *NodeInstanceData) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.


### GetZone

`func (o *NodeInstanceData) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NodeInstanceData) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NodeInstanceData) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


