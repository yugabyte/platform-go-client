# TelemetryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogs** | Pointer to [**AuditLogsTelemetrySpec**](AuditLogsTelemetrySpec.md) |  | [optional] 
**QueryLogs** | Pointer to [**QueryLogsTelemetrySpec**](QueryLogsTelemetrySpec.md) |  | [optional] 
**Metrics** | Pointer to [**MetricsTelemetrySpec**](MetricsTelemetrySpec.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


