# TelemetryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogs** | Pointer to [**AuditLogsTelemetrySpec**](AuditLogsTelemetrySpec.md) |  | [optional] 
**QueryLogs** | Pointer to [**QueryLogsTelemetrySpec**](QueryLogsTelemetrySpec.md) |  | [optional] 
**Metrics** | Pointer to [**MetricsTelemetrySpec**](MetricsTelemetrySpec.md) |  | [optional] 
**MasterLogs** | Pointer to [**MasterLogsTelemetrySpec**](MasterLogsTelemetrySpec.md) |  | [optional] 
**TserverLogs** | Pointer to [**TServerLogsTelemetrySpec**](TServerLogsTelemetrySpec.md) |  | [optional] 
**YsqlConnMgrLogs** | Pointer to [**YsqlConnMgrLogsTelemetrySpec**](YsqlConnMgrLogsTelemetrySpec.md) |  | [optional] 
**NodeAgentLogs** | Pointer to [**NodeAgentLogsTelemetrySpec**](NodeAgentLogsTelemetrySpec.md) |  | [optional] 
**YnpLogs** | Pointer to [**YnpLogsTelemetrySpec**](YnpLogsTelemetrySpec.md) |  | [optional] 
**ControllerLogs** | Pointer to [**ControllerLogsTelemetrySpec**](ControllerLogsTelemetrySpec.md) |  | [optional] 

## Methods

### NewTelemetryConfig

`func NewTelemetryConfig() *TelemetryConfig`

NewTelemetryConfig instantiates a new TelemetryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryConfigWithDefaults

`func NewTelemetryConfigWithDefaults() *TelemetryConfig`

NewTelemetryConfigWithDefaults instantiates a new TelemetryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogs

`func (o *TelemetryConfig) GetAuditLogs() AuditLogsTelemetrySpec`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *TelemetryConfig) GetAuditLogsOk() (*AuditLogsTelemetrySpec, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *TelemetryConfig) SetAuditLogs(v AuditLogsTelemetrySpec)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *TelemetryConfig) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### GetQueryLogs

`func (o *TelemetryConfig) GetQueryLogs() QueryLogsTelemetrySpec`

GetQueryLogs returns the QueryLogs field if non-nil, zero value otherwise.

### GetQueryLogsOk

`func (o *TelemetryConfig) GetQueryLogsOk() (*QueryLogsTelemetrySpec, bool)`

GetQueryLogsOk returns a tuple with the QueryLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryLogs

`func (o *TelemetryConfig) SetQueryLogs(v QueryLogsTelemetrySpec)`

SetQueryLogs sets QueryLogs field to given value.

### HasQueryLogs

`func (o *TelemetryConfig) HasQueryLogs() bool`

HasQueryLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *TelemetryConfig) GetMetrics() MetricsTelemetrySpec`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TelemetryConfig) GetMetricsOk() (*MetricsTelemetrySpec, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TelemetryConfig) SetMetrics(v MetricsTelemetrySpec)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *TelemetryConfig) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMasterLogs

`func (o *TelemetryConfig) GetMasterLogs() MasterLogsTelemetrySpec`

GetMasterLogs returns the MasterLogs field if non-nil, zero value otherwise.

### GetMasterLogsOk

`func (o *TelemetryConfig) GetMasterLogsOk() (*MasterLogsTelemetrySpec, bool)`

GetMasterLogsOk returns a tuple with the MasterLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterLogs

`func (o *TelemetryConfig) SetMasterLogs(v MasterLogsTelemetrySpec)`

SetMasterLogs sets MasterLogs field to given value.

### HasMasterLogs

`func (o *TelemetryConfig) HasMasterLogs() bool`

HasMasterLogs returns a boolean if a field has been set.

### GetTserverLogs

`func (o *TelemetryConfig) GetTserverLogs() TServerLogsTelemetrySpec`

GetTserverLogs returns the TserverLogs field if non-nil, zero value otherwise.

### GetTserverLogsOk

`func (o *TelemetryConfig) GetTserverLogsOk() (*TServerLogsTelemetrySpec, bool)`

GetTserverLogsOk returns a tuple with the TserverLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverLogs

`func (o *TelemetryConfig) SetTserverLogs(v TServerLogsTelemetrySpec)`

SetTserverLogs sets TserverLogs field to given value.

### HasTserverLogs

`func (o *TelemetryConfig) HasTserverLogs() bool`

HasTserverLogs returns a boolean if a field has been set.

### GetYsqlConnMgrLogs

`func (o *TelemetryConfig) GetYsqlConnMgrLogs() YsqlConnMgrLogsTelemetrySpec`

GetYsqlConnMgrLogs returns the YsqlConnMgrLogs field if non-nil, zero value otherwise.

### GetYsqlConnMgrLogsOk

`func (o *TelemetryConfig) GetYsqlConnMgrLogsOk() (*YsqlConnMgrLogsTelemetrySpec, bool)`

GetYsqlConnMgrLogsOk returns a tuple with the YsqlConnMgrLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlConnMgrLogs

`func (o *TelemetryConfig) SetYsqlConnMgrLogs(v YsqlConnMgrLogsTelemetrySpec)`

SetYsqlConnMgrLogs sets YsqlConnMgrLogs field to given value.

### HasYsqlConnMgrLogs

`func (o *TelemetryConfig) HasYsqlConnMgrLogs() bool`

HasYsqlConnMgrLogs returns a boolean if a field has been set.

### GetNodeAgentLogs

`func (o *TelemetryConfig) GetNodeAgentLogs() NodeAgentLogsTelemetrySpec`

GetNodeAgentLogs returns the NodeAgentLogs field if non-nil, zero value otherwise.

### GetNodeAgentLogsOk

`func (o *TelemetryConfig) GetNodeAgentLogsOk() (*NodeAgentLogsTelemetrySpec, bool)`

GetNodeAgentLogsOk returns a tuple with the NodeAgentLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAgentLogs

`func (o *TelemetryConfig) SetNodeAgentLogs(v NodeAgentLogsTelemetrySpec)`

SetNodeAgentLogs sets NodeAgentLogs field to given value.

### HasNodeAgentLogs

`func (o *TelemetryConfig) HasNodeAgentLogs() bool`

HasNodeAgentLogs returns a boolean if a field has been set.

### GetYnpLogs

`func (o *TelemetryConfig) GetYnpLogs() YnpLogsTelemetrySpec`

GetYnpLogs returns the YnpLogs field if non-nil, zero value otherwise.

### GetYnpLogsOk

`func (o *TelemetryConfig) GetYnpLogsOk() (*YnpLogsTelemetrySpec, bool)`

GetYnpLogsOk returns a tuple with the YnpLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYnpLogs

`func (o *TelemetryConfig) SetYnpLogs(v YnpLogsTelemetrySpec)`

SetYnpLogs sets YnpLogs field to given value.

### HasYnpLogs

`func (o *TelemetryConfig) HasYnpLogs() bool`

HasYnpLogs returns a boolean if a field has been set.

### GetControllerLogs

`func (o *TelemetryConfig) GetControllerLogs() ControllerLogsTelemetrySpec`

GetControllerLogs returns the ControllerLogs field if non-nil, zero value otherwise.

### GetControllerLogsOk

`func (o *TelemetryConfig) GetControllerLogsOk() (*ControllerLogsTelemetrySpec, bool)`

GetControllerLogsOk returns a tuple with the ControllerLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerLogs

`func (o *TelemetryConfig) SetControllerLogs(v ControllerLogsTelemetrySpec)`

SetControllerLogs sets ControllerLogs field to given value.

### HasControllerLogs

`func (o *TelemetryConfig) HasControllerLogs() bool`

HasControllerLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


