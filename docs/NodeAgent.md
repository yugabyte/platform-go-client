# NodeAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchType** | Pointer to **string** | Node agent host machine arch | [optional] [readonly] 
**Config** | Pointer to [**Config**](Config.md) |  | [optional] 
**CustomerUuid** | Pointer to **string** | Customer UUID | [optional] [readonly] 
**Home** | Pointer to **string** | Node agent installation directory | [optional] [readonly] 
**Ip** | Pointer to **string** | Node agent server IP | [optional] [readonly] 
**LastError** | Pointer to [**YBAError**](YBAError.md) |  | [optional] 
**Name** | Pointer to **string** | Node agent name | [optional] [readonly] 
**OsType** | Pointer to **string** | Node agent host OS | [optional] [readonly] 
**Port** | Pointer to **int32** | Node agent server port | [optional] [readonly] 
**State** | Pointer to **string** | Node agent state | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Updated time | [optional] [readonly] 
**Uuid** | Pointer to **string** | Node agent UUID | [optional] [readonly] 
**Version** | Pointer to **string** | Node agent installed version | [optional] [readonly] 

## Methods

### NewNodeAgent

`func NewNodeAgent() *NodeAgent`

NewNodeAgent instantiates a new NodeAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAgentWithDefaults

`func NewNodeAgentWithDefaults() *NodeAgent`

NewNodeAgentWithDefaults instantiates a new NodeAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchType

`func (o *NodeAgent) GetArchType() string`

GetArchType returns the ArchType field if non-nil, zero value otherwise.

### GetArchTypeOk

`func (o *NodeAgent) GetArchTypeOk() (*string, bool)`

GetArchTypeOk returns a tuple with the ArchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchType

`func (o *NodeAgent) SetArchType(v string)`

SetArchType sets ArchType field to given value.

### HasArchType

`func (o *NodeAgent) HasArchType() bool`

HasArchType returns a boolean if a field has been set.

### GetConfig

`func (o *NodeAgent) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NodeAgent) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NodeAgent) SetConfig(v Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NodeAgent) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *NodeAgent) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *NodeAgent) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *NodeAgent) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *NodeAgent) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetHome

`func (o *NodeAgent) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *NodeAgent) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *NodeAgent) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *NodeAgent) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIp

`func (o *NodeAgent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NodeAgent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NodeAgent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NodeAgent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastError

`func (o *NodeAgent) GetLastError() YBAError`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *NodeAgent) GetLastErrorOk() (*YBAError, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *NodeAgent) SetLastError(v YBAError)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *NodeAgent) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetName

`func (o *NodeAgent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeAgent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeAgent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeAgent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsType

`func (o *NodeAgent) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *NodeAgent) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *NodeAgent) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *NodeAgent) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetPort

`func (o *NodeAgent) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NodeAgent) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NodeAgent) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NodeAgent) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetState

`func (o *NodeAgent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NodeAgent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NodeAgent) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NodeAgent) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NodeAgent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NodeAgent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NodeAgent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NodeAgent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *NodeAgent) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NodeAgent) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NodeAgent) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NodeAgent) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *NodeAgent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeAgent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeAgent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeAgent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


