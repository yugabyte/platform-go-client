# TelemetryProviderTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | Pointer to **string** | Provider type name | [optional] 
**IsAllowedForLogs** | Pointer to **bool** | Whether this provider supports logs export | [optional] 
**IsAllowedForMetrics** | Pointer to **bool** | Whether this provider supports metrics export | [optional] 

## Methods

### NewTelemetryProviderTypeInfo

`func NewTelemetryProviderTypeInfo() *TelemetryProviderTypeInfo`

NewTelemetryProviderTypeInfo instantiates a new TelemetryProviderTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryProviderTypeInfoWithDefaults

`func NewTelemetryProviderTypeInfoWithDefaults() *TelemetryProviderTypeInfo`

NewTelemetryProviderTypeInfoWithDefaults instantiates a new TelemetryProviderTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *TelemetryProviderTypeInfo) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *TelemetryProviderTypeInfo) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *TelemetryProviderTypeInfo) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *TelemetryProviderTypeInfo) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetIsAllowedForLogs

`func (o *TelemetryProviderTypeInfo) GetIsAllowedForLogs() bool`

GetIsAllowedForLogs returns the IsAllowedForLogs field if non-nil, zero value otherwise.

### GetIsAllowedForLogsOk

`func (o *TelemetryProviderTypeInfo) GetIsAllowedForLogsOk() (*bool, bool)`

GetIsAllowedForLogsOk returns a tuple with the IsAllowedForLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowedForLogs

`func (o *TelemetryProviderTypeInfo) SetIsAllowedForLogs(v bool)`

SetIsAllowedForLogs sets IsAllowedForLogs field to given value.

### HasIsAllowedForLogs

`func (o *TelemetryProviderTypeInfo) HasIsAllowedForLogs() bool`

HasIsAllowedForLogs returns a boolean if a field has been set.

### GetIsAllowedForMetrics

`func (o *TelemetryProviderTypeInfo) GetIsAllowedForMetrics() bool`

GetIsAllowedForMetrics returns the IsAllowedForMetrics field if non-nil, zero value otherwise.

### GetIsAllowedForMetricsOk

`func (o *TelemetryProviderTypeInfo) GetIsAllowedForMetricsOk() (*bool, bool)`

GetIsAllowedForMetricsOk returns a tuple with the IsAllowedForMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowedForMetrics

`func (o *TelemetryProviderTypeInfo) SetIsAllowedForMetrics(v bool)`

SetIsAllowedForMetrics sets IsAllowedForMetrics field to given value.

### HasIsAllowedForMetrics

`func (o *TelemetryProviderTypeInfo) HasIsAllowedForMetrics() bool`

HasIsAllowedForMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


