# CommunicationPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterHttpPort** | Pointer to **int32** | Master table HTTP port | [optional] 
**MasterRpcPort** | Pointer to **int32** | Master table RCP port | [optional] 
**NodeExporterPort** | Pointer to **int32** | Node exporter port | [optional] 
**RedisServerHttpPort** | Pointer to **int32** | Redis HTTP port | [optional] 
**RedisServerRpcPort** | Pointer to **int32** | Redis RPC port | [optional] 
**TserverHttpPort** | Pointer to **int32** | Tablet server HTTP port | [optional] 
**TserverRpcPort** | Pointer to **int32** | Tablet server RPC port | [optional] 
**YbControllerHttpPort** | Pointer to **int32** | Yb controller HTTP port | [optional] 
**YbControllerrRpcPort** | Pointer to **int32** | Yb controller RPC port | [optional] 
**YqlServerHttpPort** | Pointer to **int32** | YQL HTTP port | [optional] 
**YqlServerRpcPort** | Pointer to **int32** | YQL RPC port | [optional] 
**YsqlServerHttpPort** | Pointer to **int32** | YSQL HTTP port | [optional] 
**YsqlServerRpcPort** | Pointer to **int32** | YSQL RPC port | [optional] 

## Methods

### NewCommunicationPorts

`func NewCommunicationPorts() *CommunicationPorts`

NewCommunicationPorts instantiates a new CommunicationPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationPortsWithDefaults

`func NewCommunicationPortsWithDefaults() *CommunicationPorts`

NewCommunicationPortsWithDefaults instantiates a new CommunicationPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterHttpPort

`func (o *CommunicationPorts) GetMasterHttpPort() int32`

GetMasterHttpPort returns the MasterHttpPort field if non-nil, zero value otherwise.

### GetMasterHttpPortOk

`func (o *CommunicationPorts) GetMasterHttpPortOk() (*int32, bool)`

GetMasterHttpPortOk returns a tuple with the MasterHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterHttpPort

`func (o *CommunicationPorts) SetMasterHttpPort(v int32)`

SetMasterHttpPort sets MasterHttpPort field to given value.

### HasMasterHttpPort

`func (o *CommunicationPorts) HasMasterHttpPort() bool`

HasMasterHttpPort returns a boolean if a field has been set.

### GetMasterRpcPort

`func (o *CommunicationPorts) GetMasterRpcPort() int32`

GetMasterRpcPort returns the MasterRpcPort field if non-nil, zero value otherwise.

### GetMasterRpcPortOk

`func (o *CommunicationPorts) GetMasterRpcPortOk() (*int32, bool)`

GetMasterRpcPortOk returns a tuple with the MasterRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRpcPort

`func (o *CommunicationPorts) SetMasterRpcPort(v int32)`

SetMasterRpcPort sets MasterRpcPort field to given value.

### HasMasterRpcPort

`func (o *CommunicationPorts) HasMasterRpcPort() bool`

HasMasterRpcPort returns a boolean if a field has been set.

### GetNodeExporterPort

`func (o *CommunicationPorts) GetNodeExporterPort() int32`

GetNodeExporterPort returns the NodeExporterPort field if non-nil, zero value otherwise.

### GetNodeExporterPortOk

`func (o *CommunicationPorts) GetNodeExporterPortOk() (*int32, bool)`

GetNodeExporterPortOk returns a tuple with the NodeExporterPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterPort

`func (o *CommunicationPorts) SetNodeExporterPort(v int32)`

SetNodeExporterPort sets NodeExporterPort field to given value.

### HasNodeExporterPort

`func (o *CommunicationPorts) HasNodeExporterPort() bool`

HasNodeExporterPort returns a boolean if a field has been set.

### GetRedisServerHttpPort

`func (o *CommunicationPorts) GetRedisServerHttpPort() int32`

GetRedisServerHttpPort returns the RedisServerHttpPort field if non-nil, zero value otherwise.

### GetRedisServerHttpPortOk

`func (o *CommunicationPorts) GetRedisServerHttpPortOk() (*int32, bool)`

GetRedisServerHttpPortOk returns a tuple with the RedisServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisServerHttpPort

`func (o *CommunicationPorts) SetRedisServerHttpPort(v int32)`

SetRedisServerHttpPort sets RedisServerHttpPort field to given value.

### HasRedisServerHttpPort

`func (o *CommunicationPorts) HasRedisServerHttpPort() bool`

HasRedisServerHttpPort returns a boolean if a field has been set.

### GetRedisServerRpcPort

`func (o *CommunicationPorts) GetRedisServerRpcPort() int32`

GetRedisServerRpcPort returns the RedisServerRpcPort field if non-nil, zero value otherwise.

### GetRedisServerRpcPortOk

`func (o *CommunicationPorts) GetRedisServerRpcPortOk() (*int32, bool)`

GetRedisServerRpcPortOk returns a tuple with the RedisServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisServerRpcPort

`func (o *CommunicationPorts) SetRedisServerRpcPort(v int32)`

SetRedisServerRpcPort sets RedisServerRpcPort field to given value.

### HasRedisServerRpcPort

`func (o *CommunicationPorts) HasRedisServerRpcPort() bool`

HasRedisServerRpcPort returns a boolean if a field has been set.

### GetTserverHttpPort

`func (o *CommunicationPorts) GetTserverHttpPort() int32`

GetTserverHttpPort returns the TserverHttpPort field if non-nil, zero value otherwise.

### GetTserverHttpPortOk

`func (o *CommunicationPorts) GetTserverHttpPortOk() (*int32, bool)`

GetTserverHttpPortOk returns a tuple with the TserverHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverHttpPort

`func (o *CommunicationPorts) SetTserverHttpPort(v int32)`

SetTserverHttpPort sets TserverHttpPort field to given value.

### HasTserverHttpPort

`func (o *CommunicationPorts) HasTserverHttpPort() bool`

HasTserverHttpPort returns a boolean if a field has been set.

### GetTserverRpcPort

`func (o *CommunicationPorts) GetTserverRpcPort() int32`

GetTserverRpcPort returns the TserverRpcPort field if non-nil, zero value otherwise.

### GetTserverRpcPortOk

`func (o *CommunicationPorts) GetTserverRpcPortOk() (*int32, bool)`

GetTserverRpcPortOk returns a tuple with the TserverRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverRpcPort

`func (o *CommunicationPorts) SetTserverRpcPort(v int32)`

SetTserverRpcPort sets TserverRpcPort field to given value.

### HasTserverRpcPort

`func (o *CommunicationPorts) HasTserverRpcPort() bool`

HasTserverRpcPort returns a boolean if a field has been set.

### GetYbControllerHttpPort

`func (o *CommunicationPorts) GetYbControllerHttpPort() int32`

GetYbControllerHttpPort returns the YbControllerHttpPort field if non-nil, zero value otherwise.

### GetYbControllerHttpPortOk

`func (o *CommunicationPorts) GetYbControllerHttpPortOk() (*int32, bool)`

GetYbControllerHttpPortOk returns a tuple with the YbControllerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbControllerHttpPort

`func (o *CommunicationPorts) SetYbControllerHttpPort(v int32)`

SetYbControllerHttpPort sets YbControllerHttpPort field to given value.

### HasYbControllerHttpPort

`func (o *CommunicationPorts) HasYbControllerHttpPort() bool`

HasYbControllerHttpPort returns a boolean if a field has been set.

### GetYbControllerrRpcPort

`func (o *CommunicationPorts) GetYbControllerrRpcPort() int32`

GetYbControllerrRpcPort returns the YbControllerrRpcPort field if non-nil, zero value otherwise.

### GetYbControllerrRpcPortOk

`func (o *CommunicationPorts) GetYbControllerrRpcPortOk() (*int32, bool)`

GetYbControllerrRpcPortOk returns a tuple with the YbControllerrRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbControllerrRpcPort

`func (o *CommunicationPorts) SetYbControllerrRpcPort(v int32)`

SetYbControllerrRpcPort sets YbControllerrRpcPort field to given value.

### HasYbControllerrRpcPort

`func (o *CommunicationPorts) HasYbControllerrRpcPort() bool`

HasYbControllerrRpcPort returns a boolean if a field has been set.

### GetYqlServerHttpPort

`func (o *CommunicationPorts) GetYqlServerHttpPort() int32`

GetYqlServerHttpPort returns the YqlServerHttpPort field if non-nil, zero value otherwise.

### GetYqlServerHttpPortOk

`func (o *CommunicationPorts) GetYqlServerHttpPortOk() (*int32, bool)`

GetYqlServerHttpPortOk returns a tuple with the YqlServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYqlServerHttpPort

`func (o *CommunicationPorts) SetYqlServerHttpPort(v int32)`

SetYqlServerHttpPort sets YqlServerHttpPort field to given value.

### HasYqlServerHttpPort

`func (o *CommunicationPorts) HasYqlServerHttpPort() bool`

HasYqlServerHttpPort returns a boolean if a field has been set.

### GetYqlServerRpcPort

`func (o *CommunicationPorts) GetYqlServerRpcPort() int32`

GetYqlServerRpcPort returns the YqlServerRpcPort field if non-nil, zero value otherwise.

### GetYqlServerRpcPortOk

`func (o *CommunicationPorts) GetYqlServerRpcPortOk() (*int32, bool)`

GetYqlServerRpcPortOk returns a tuple with the YqlServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYqlServerRpcPort

`func (o *CommunicationPorts) SetYqlServerRpcPort(v int32)`

SetYqlServerRpcPort sets YqlServerRpcPort field to given value.

### HasYqlServerRpcPort

`func (o *CommunicationPorts) HasYqlServerRpcPort() bool`

HasYqlServerRpcPort returns a boolean if a field has been set.

### GetYsqlServerHttpPort

`func (o *CommunicationPorts) GetYsqlServerHttpPort() int32`

GetYsqlServerHttpPort returns the YsqlServerHttpPort field if non-nil, zero value otherwise.

### GetYsqlServerHttpPortOk

`func (o *CommunicationPorts) GetYsqlServerHttpPortOk() (*int32, bool)`

GetYsqlServerHttpPortOk returns a tuple with the YsqlServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlServerHttpPort

`func (o *CommunicationPorts) SetYsqlServerHttpPort(v int32)`

SetYsqlServerHttpPort sets YsqlServerHttpPort field to given value.

### HasYsqlServerHttpPort

`func (o *CommunicationPorts) HasYsqlServerHttpPort() bool`

HasYsqlServerHttpPort returns a boolean if a field has been set.

### GetYsqlServerRpcPort

`func (o *CommunicationPorts) GetYsqlServerRpcPort() int32`

GetYsqlServerRpcPort returns the YsqlServerRpcPort field if non-nil, zero value otherwise.

### GetYsqlServerRpcPortOk

`func (o *CommunicationPorts) GetYsqlServerRpcPortOk() (*int32, bool)`

GetYsqlServerRpcPortOk returns a tuple with the YsqlServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlServerRpcPort

`func (o *CommunicationPorts) SetYsqlServerRpcPort(v int32)`

SetYsqlServerRpcPort sets YsqlServerRpcPort field to given value.

### HasYsqlServerRpcPort

`func (o *CommunicationPorts) HasYsqlServerRpcPort() bool`

HasYsqlServerRpcPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


