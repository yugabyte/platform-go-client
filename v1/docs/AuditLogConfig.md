# AuditLogConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportActive** | Pointer to **bool** | Universe logs export active | [optional] 
**UniverseLogsExporterConfig** | [**[]UniverseLogsExporterConfig**](UniverseLogsExporterConfig.md) | Universe logs exporter config | 
**YcqlAuditConfig** | Pointer to [**YCQLAuditConfig**](YCQLAuditConfig.md) |  | [optional] 
**YsqlAuditConfig** | Pointer to [**YSQLAuditConfig**](YSQLAuditConfig.md) |  | [optional] 

## Methods

### NewAuditLogConfig

`func NewAuditLogConfig(universeLogsExporterConfig []UniverseLogsExporterConfig, ) *AuditLogConfig`

NewAuditLogConfig instantiates a new AuditLogConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogConfigWithDefaults

`func NewAuditLogConfigWithDefaults() *AuditLogConfig`

NewAuditLogConfigWithDefaults instantiates a new AuditLogConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportActive

`func (o *AuditLogConfig) GetExportActive() bool`

GetExportActive returns the ExportActive field if non-nil, zero value otherwise.

### GetExportActiveOk

`func (o *AuditLogConfig) GetExportActiveOk() (*bool, bool)`

GetExportActiveOk returns a tuple with the ExportActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportActive

`func (o *AuditLogConfig) SetExportActive(v bool)`

SetExportActive sets ExportActive field to given value.

### HasExportActive

`func (o *AuditLogConfig) HasExportActive() bool`

HasExportActive returns a boolean if a field has been set.

### GetUniverseLogsExporterConfig

`func (o *AuditLogConfig) GetUniverseLogsExporterConfig() []UniverseLogsExporterConfig`

GetUniverseLogsExporterConfig returns the UniverseLogsExporterConfig field if non-nil, zero value otherwise.

### GetUniverseLogsExporterConfigOk

`func (o *AuditLogConfig) GetUniverseLogsExporterConfigOk() (*[]UniverseLogsExporterConfig, bool)`

GetUniverseLogsExporterConfigOk returns a tuple with the UniverseLogsExporterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseLogsExporterConfig

`func (o *AuditLogConfig) SetUniverseLogsExporterConfig(v []UniverseLogsExporterConfig)`

SetUniverseLogsExporterConfig sets UniverseLogsExporterConfig field to given value.


### GetYcqlAuditConfig

`func (o *AuditLogConfig) GetYcqlAuditConfig() YCQLAuditConfig`

GetYcqlAuditConfig returns the YcqlAuditConfig field if non-nil, zero value otherwise.

### GetYcqlAuditConfigOk

`func (o *AuditLogConfig) GetYcqlAuditConfigOk() (*YCQLAuditConfig, bool)`

GetYcqlAuditConfigOk returns a tuple with the YcqlAuditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlAuditConfig

`func (o *AuditLogConfig) SetYcqlAuditConfig(v YCQLAuditConfig)`

SetYcqlAuditConfig sets YcqlAuditConfig field to given value.

### HasYcqlAuditConfig

`func (o *AuditLogConfig) HasYcqlAuditConfig() bool`

HasYcqlAuditConfig returns a boolean if a field has been set.

### GetYsqlAuditConfig

`func (o *AuditLogConfig) GetYsqlAuditConfig() YSQLAuditConfig`

GetYsqlAuditConfig returns the YsqlAuditConfig field if non-nil, zero value otherwise.

### GetYsqlAuditConfigOk

`func (o *AuditLogConfig) GetYsqlAuditConfigOk() (*YSQLAuditConfig, bool)`

GetYsqlAuditConfigOk returns a tuple with the YsqlAuditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlAuditConfig

`func (o *AuditLogConfig) SetYsqlAuditConfig(v YSQLAuditConfig)`

SetYsqlAuditConfig sets YsqlAuditConfig field to given value.

### HasYsqlAuditConfig

`func (o *AuditLogConfig) HasYsqlAuditConfig() bool`

HasYsqlAuditConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


