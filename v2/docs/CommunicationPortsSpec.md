# CommunicationPortsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterHttpPort** | Pointer to **int32** | Master table HTTP port | [optional] [default to 7000]
**MasterRpcPort** | Pointer to **int32** | Master table RCP port | [optional] [default to 7100]
**NodeExporterPort** | Pointer to **int32** | Node exporter port | [optional] [default to 9300]
**OtelCollectorMetricsPort** | Pointer to **int32** | Otel Collector metrics port | [optional] [default to 8889]
**RedisServerHttpPort** | Pointer to **int32** | Redis HTTP port | [optional] [default to 11000]
**RedisServerRpcPort** | Pointer to **int32** | Redis RPC port | [optional] [default to 6379]
**TserverHttpPort** | Pointer to **int32** | Tablet server HTTP port | [optional] [default to 9000]
**TserverRpcPort** | Pointer to **int32** | Tablet server RPC port | [optional] [default to 9100]
**YbControllerHttpPort** | Pointer to **int32** | Yb controller HTTP port | [optional] [default to 14000]
**YbControllerRpcPort** | Pointer to **int32** | Yb controller RPC port | [optional] [default to 18018]
**YqlServerHttpPort** | Pointer to **int32** | YQL HTTP port | [optional] [default to 12000]
**YqlServerRpcPort** | Pointer to **int32** | YQL RPC port | [optional] [default to 9042]
**YsqlServerHttpPort** | Pointer to **int32** | YSQL HTTP port | [optional] [default to 13000]
**YsqlServerRpcPort** | Pointer to **int32** | YSQL RPC port | [optional] [default to 5433]

## Methods

### NewCommunicationPortsSpec

`func NewCommunicationPortsSpec() *CommunicationPortsSpec`

NewCommunicationPortsSpec instantiates a new CommunicationPortsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationPortsSpecWithDefaults

`func NewCommunicationPortsSpecWithDefaults() *CommunicationPortsSpec`

NewCommunicationPortsSpecWithDefaults instantiates a new CommunicationPortsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterHttpPort

`func (o *CommunicationPortsSpec) GetMasterHttpPort() int32`

GetMasterHttpPort returns the MasterHttpPort field if non-nil, zero value otherwise.

### GetMasterHttpPortOk

`func (o *CommunicationPortsSpec) GetMasterHttpPortOk() (*int32, bool)`

GetMasterHttpPortOk returns a tuple with the MasterHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterHttpPort

`func (o *CommunicationPortsSpec) SetMasterHttpPort(v int32)`

SetMasterHttpPort sets MasterHttpPort field to given value.

### HasMasterHttpPort

`func (o *CommunicationPortsSpec) HasMasterHttpPort() bool`

HasMasterHttpPort returns a boolean if a field has been set.

### GetMasterRpcPort

`func (o *CommunicationPortsSpec) GetMasterRpcPort() int32`

GetMasterRpcPort returns the MasterRpcPort field if non-nil, zero value otherwise.

### GetMasterRpcPortOk

`func (o *CommunicationPortsSpec) GetMasterRpcPortOk() (*int32, bool)`

GetMasterRpcPortOk returns a tuple with the MasterRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRpcPort

`func (o *CommunicationPortsSpec) SetMasterRpcPort(v int32)`

SetMasterRpcPort sets MasterRpcPort field to given value.

### HasMasterRpcPort

`func (o *CommunicationPortsSpec) HasMasterRpcPort() bool`

HasMasterRpcPort returns a boolean if a field has been set.

### GetNodeExporterPort

`func (o *CommunicationPortsSpec) GetNodeExporterPort() int32`

GetNodeExporterPort returns the NodeExporterPort field if non-nil, zero value otherwise.

### GetNodeExporterPortOk

`func (o *CommunicationPortsSpec) GetNodeExporterPortOk() (*int32, bool)`

GetNodeExporterPortOk returns a tuple with the NodeExporterPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterPort

`func (o *CommunicationPortsSpec) SetNodeExporterPort(v int32)`

SetNodeExporterPort sets NodeExporterPort field to given value.

### HasNodeExporterPort

`func (o *CommunicationPortsSpec) HasNodeExporterPort() bool`

HasNodeExporterPort returns a boolean if a field has been set.

### GetOtelCollectorMetricsPort

`func (o *CommunicationPortsSpec) GetOtelCollectorMetricsPort() int32`

GetOtelCollectorMetricsPort returns the OtelCollectorMetricsPort field if non-nil, zero value otherwise.

### GetOtelCollectorMetricsPortOk

`func (o *CommunicationPortsSpec) GetOtelCollectorMetricsPortOk() (*int32, bool)`

GetOtelCollectorMetricsPortOk returns a tuple with the OtelCollectorMetricsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelCollectorMetricsPort

`func (o *CommunicationPortsSpec) SetOtelCollectorMetricsPort(v int32)`

SetOtelCollectorMetricsPort sets OtelCollectorMetricsPort field to given value.

### HasOtelCollectorMetricsPort

`func (o *CommunicationPortsSpec) HasOtelCollectorMetricsPort() bool`

HasOtelCollectorMetricsPort returns a boolean if a field has been set.

### GetRedisServerHttpPort

`func (o *CommunicationPortsSpec) GetRedisServerHttpPort() int32`

GetRedisServerHttpPort returns the RedisServerHttpPort field if non-nil, zero value otherwise.

### GetRedisServerHttpPortOk

`func (o *CommunicationPortsSpec) GetRedisServerHttpPortOk() (*int32, bool)`

GetRedisServerHttpPortOk returns a tuple with the RedisServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisServerHttpPort

`func (o *CommunicationPortsSpec) SetRedisServerHttpPort(v int32)`

SetRedisServerHttpPort sets RedisServerHttpPort field to given value.

### HasRedisServerHttpPort

`func (o *CommunicationPortsSpec) HasRedisServerHttpPort() bool`

HasRedisServerHttpPort returns a boolean if a field has been set.

### GetRedisServerRpcPort

`func (o *CommunicationPortsSpec) GetRedisServerRpcPort() int32`

GetRedisServerRpcPort returns the RedisServerRpcPort field if non-nil, zero value otherwise.

### GetRedisServerRpcPortOk

`func (o *CommunicationPortsSpec) GetRedisServerRpcPortOk() (*int32, bool)`

GetRedisServerRpcPortOk returns a tuple with the RedisServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisServerRpcPort

`func (o *CommunicationPortsSpec) SetRedisServerRpcPort(v int32)`

SetRedisServerRpcPort sets RedisServerRpcPort field to given value.

### HasRedisServerRpcPort

`func (o *CommunicationPortsSpec) HasRedisServerRpcPort() bool`

HasRedisServerRpcPort returns a boolean if a field has been set.

### GetTserverHttpPort

`func (o *CommunicationPortsSpec) GetTserverHttpPort() int32`

GetTserverHttpPort returns the TserverHttpPort field if non-nil, zero value otherwise.

### GetTserverHttpPortOk

`func (o *CommunicationPortsSpec) GetTserverHttpPortOk() (*int32, bool)`

GetTserverHttpPortOk returns a tuple with the TserverHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverHttpPort

`func (o *CommunicationPortsSpec) SetTserverHttpPort(v int32)`

SetTserverHttpPort sets TserverHttpPort field to given value.

### HasTserverHttpPort

`func (o *CommunicationPortsSpec) HasTserverHttpPort() bool`

HasTserverHttpPort returns a boolean if a field has been set.

### GetTserverRpcPort

`func (o *CommunicationPortsSpec) GetTserverRpcPort() int32`

GetTserverRpcPort returns the TserverRpcPort field if non-nil, zero value otherwise.

### GetTserverRpcPortOk

`func (o *CommunicationPortsSpec) GetTserverRpcPortOk() (*int32, bool)`

GetTserverRpcPortOk returns a tuple with the TserverRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverRpcPort

`func (o *CommunicationPortsSpec) SetTserverRpcPort(v int32)`

SetTserverRpcPort sets TserverRpcPort field to given value.

### HasTserverRpcPort

`func (o *CommunicationPortsSpec) HasTserverRpcPort() bool`

HasTserverRpcPort returns a boolean if a field has been set.

### GetYbControllerHttpPort

`func (o *CommunicationPortsSpec) GetYbControllerHttpPort() int32`

GetYbControllerHttpPort returns the YbControllerHttpPort field if non-nil, zero value otherwise.

### GetYbControllerHttpPortOk

`func (o *CommunicationPortsSpec) GetYbControllerHttpPortOk() (*int32, bool)`

GetYbControllerHttpPortOk returns a tuple with the YbControllerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbControllerHttpPort

`func (o *CommunicationPortsSpec) SetYbControllerHttpPort(v int32)`

SetYbControllerHttpPort sets YbControllerHttpPort field to given value.

### HasYbControllerHttpPort

`func (o *CommunicationPortsSpec) HasYbControllerHttpPort() bool`

HasYbControllerHttpPort returns a boolean if a field has been set.

### GetYbControllerRpcPort

`func (o *CommunicationPortsSpec) GetYbControllerRpcPort() int32`

GetYbControllerRpcPort returns the YbControllerRpcPort field if non-nil, zero value otherwise.

### GetYbControllerRpcPortOk

`func (o *CommunicationPortsSpec) GetYbControllerRpcPortOk() (*int32, bool)`

GetYbControllerRpcPortOk returns a tuple with the YbControllerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbControllerRpcPort

`func (o *CommunicationPortsSpec) SetYbControllerRpcPort(v int32)`

SetYbControllerRpcPort sets YbControllerRpcPort field to given value.

### HasYbControllerRpcPort

`func (o *CommunicationPortsSpec) HasYbControllerRpcPort() bool`

HasYbControllerRpcPort returns a boolean if a field has been set.

### GetYqlServerHttpPort

`func (o *CommunicationPortsSpec) GetYqlServerHttpPort() int32`

GetYqlServerHttpPort returns the YqlServerHttpPort field if non-nil, zero value otherwise.

### GetYqlServerHttpPortOk

`func (o *CommunicationPortsSpec) GetYqlServerHttpPortOk() (*int32, bool)`

GetYqlServerHttpPortOk returns a tuple with the YqlServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYqlServerHttpPort

`func (o *CommunicationPortsSpec) SetYqlServerHttpPort(v int32)`

SetYqlServerHttpPort sets YqlServerHttpPort field to given value.

### HasYqlServerHttpPort

`func (o *CommunicationPortsSpec) HasYqlServerHttpPort() bool`

HasYqlServerHttpPort returns a boolean if a field has been set.

### GetYqlServerRpcPort

`func (o *CommunicationPortsSpec) GetYqlServerRpcPort() int32`

GetYqlServerRpcPort returns the YqlServerRpcPort field if non-nil, zero value otherwise.

### GetYqlServerRpcPortOk

`func (o *CommunicationPortsSpec) GetYqlServerRpcPortOk() (*int32, bool)`

GetYqlServerRpcPortOk returns a tuple with the YqlServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYqlServerRpcPort

`func (o *CommunicationPortsSpec) SetYqlServerRpcPort(v int32)`

SetYqlServerRpcPort sets YqlServerRpcPort field to given value.

### HasYqlServerRpcPort

`func (o *CommunicationPortsSpec) HasYqlServerRpcPort() bool`

HasYqlServerRpcPort returns a boolean if a field has been set.

### GetYsqlServerHttpPort

`func (o *CommunicationPortsSpec) GetYsqlServerHttpPort() int32`

GetYsqlServerHttpPort returns the YsqlServerHttpPort field if non-nil, zero value otherwise.

### GetYsqlServerHttpPortOk

`func (o *CommunicationPortsSpec) GetYsqlServerHttpPortOk() (*int32, bool)`

GetYsqlServerHttpPortOk returns a tuple with the YsqlServerHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlServerHttpPort

`func (o *CommunicationPortsSpec) SetYsqlServerHttpPort(v int32)`

SetYsqlServerHttpPort sets YsqlServerHttpPort field to given value.

### HasYsqlServerHttpPort

`func (o *CommunicationPortsSpec) HasYsqlServerHttpPort() bool`

HasYsqlServerHttpPort returns a boolean if a field has been set.

### GetYsqlServerRpcPort

`func (o *CommunicationPortsSpec) GetYsqlServerRpcPort() int32`

GetYsqlServerRpcPort returns the YsqlServerRpcPort field if non-nil, zero value otherwise.

### GetYsqlServerRpcPortOk

`func (o *CommunicationPortsSpec) GetYsqlServerRpcPortOk() (*int32, bool)`

GetYsqlServerRpcPortOk returns a tuple with the YsqlServerRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlServerRpcPort

`func (o *CommunicationPortsSpec) SetYsqlServerRpcPort(v int32)`

SetYsqlServerRpcPort sets YsqlServerRpcPort field to given value.

### HasYsqlServerRpcPort

`func (o *CommunicationPortsSpec) HasYsqlServerRpcPort() bool`

HasYsqlServerRpcPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


