# NodeAgentResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchType** | Pointer to **string** | Node agent host machine arch | [optional] [readonly] 
**Config** | Pointer to [**Config**](Config.md) |  | [optional] 
**CustomerUuid** | Pointer to **string** | Customer UUID | [optional] [readonly] 
**Home** | Pointer to **string** | Node agent installation directory | [optional] [readonly] 
**Ip** | Pointer to **string** | Node agent server IP | [optional] [readonly] 
**Name** | Pointer to **string** | Node agent name | [optional] [readonly] 
**OsType** | Pointer to **string** | Node agent host OS | [optional] [readonly] 
**Port** | Pointer to **int32** | Node agent server port | [optional] [readonly] 
**Reachable** | Pointer to **bool** |  | [optional] [readonly] 
**State** | Pointer to **string** | Node agent state | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Updated time | [optional] [readonly] 
**Uuid** | Pointer to **string** | Node agent UUID | [optional] [readonly] 
**Version** | Pointer to **string** | Node agent installed version | [optional] [readonly] 
**VersionMatched** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewNodeAgentResp

`func NewNodeAgentResp() *NodeAgentResp`

NewNodeAgentResp instantiates a new NodeAgentResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAgentRespWithDefaults

`func NewNodeAgentRespWithDefaults() *NodeAgentResp`

NewNodeAgentRespWithDefaults instantiates a new NodeAgentResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchType

`func (o *NodeAgentResp) GetArchType() string`

GetArchType returns the ArchType field if non-nil, zero value otherwise.

### GetArchTypeOk

`func (o *NodeAgentResp) GetArchTypeOk() (*string, bool)`

GetArchTypeOk returns a tuple with the ArchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchType

`func (o *NodeAgentResp) SetArchType(v string)`

SetArchType sets ArchType field to given value.

### HasArchType

`func (o *NodeAgentResp) HasArchType() bool`

HasArchType returns a boolean if a field has been set.

### GetConfig

`func (o *NodeAgentResp) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NodeAgentResp) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NodeAgentResp) SetConfig(v Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NodeAgentResp) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *NodeAgentResp) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *NodeAgentResp) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *NodeAgentResp) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *NodeAgentResp) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetHome

`func (o *NodeAgentResp) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *NodeAgentResp) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *NodeAgentResp) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *NodeAgentResp) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIp

`func (o *NodeAgentResp) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NodeAgentResp) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NodeAgentResp) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NodeAgentResp) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *NodeAgentResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeAgentResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeAgentResp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeAgentResp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsType

`func (o *NodeAgentResp) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *NodeAgentResp) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *NodeAgentResp) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *NodeAgentResp) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetPort

`func (o *NodeAgentResp) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NodeAgentResp) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NodeAgentResp) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NodeAgentResp) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetReachable

`func (o *NodeAgentResp) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *NodeAgentResp) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *NodeAgentResp) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *NodeAgentResp) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetState

`func (o *NodeAgentResp) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NodeAgentResp) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NodeAgentResp) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NodeAgentResp) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NodeAgentResp) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NodeAgentResp) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NodeAgentResp) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NodeAgentResp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *NodeAgentResp) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NodeAgentResp) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NodeAgentResp) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NodeAgentResp) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *NodeAgentResp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeAgentResp) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeAgentResp) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeAgentResp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionMatched

`func (o *NodeAgentResp) GetVersionMatched() bool`

GetVersionMatched returns the VersionMatched field if non-nil, zero value otherwise.

### GetVersionMatchedOk

`func (o *NodeAgentResp) GetVersionMatchedOk() (*bool, bool)`

GetVersionMatchedOk returns a tuple with the VersionMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMatched

`func (o *NodeAgentResp) SetVersionMatched(v bool)`

SetVersionMatched sets VersionMatched field to given value.

### HasVersionMatched

`func (o *NodeAgentResp) HasVersionMatched() bool`

HasVersionMatched returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


