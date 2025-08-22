# QueryLogConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportActive** | Pointer to **bool** | Universe logs export active | [optional] 
**UniverseLogsExporterConfig** | [**[]UniverseQueryLogsExporterConfig**](UniverseQueryLogsExporterConfig.md) | Universe query logs exporter config | 
**YsqlQueryLogConfig** | Pointer to [**YSQLQueryLogConfig**](YSQLQueryLogConfig.md) |  | [optional] 

## Methods

### NewQueryLogConfig

`func NewQueryLogConfig(universeLogsExporterConfig []UniverseQueryLogsExporterConfig, ) *QueryLogConfig`

NewQueryLogConfig instantiates a new QueryLogConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLogConfigWithDefaults

`func NewQueryLogConfigWithDefaults() *QueryLogConfig`

NewQueryLogConfigWithDefaults instantiates a new QueryLogConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportActive

`func (o *QueryLogConfig) GetExportActive() bool`

GetExportActive returns the ExportActive field if non-nil, zero value otherwise.

### GetExportActiveOk

`func (o *QueryLogConfig) GetExportActiveOk() (*bool, bool)`

GetExportActiveOk returns a tuple with the ExportActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportActive

`func (o *QueryLogConfig) SetExportActive(v bool)`

SetExportActive sets ExportActive field to given value.

### HasExportActive

`func (o *QueryLogConfig) HasExportActive() bool`

HasExportActive returns a boolean if a field has been set.

### GetUniverseLogsExporterConfig

`func (o *QueryLogConfig) GetUniverseLogsExporterConfig() []UniverseQueryLogsExporterConfig`

GetUniverseLogsExporterConfig returns the UniverseLogsExporterConfig field if non-nil, zero value otherwise.

### GetUniverseLogsExporterConfigOk

`func (o *QueryLogConfig) GetUniverseLogsExporterConfigOk() (*[]UniverseQueryLogsExporterConfig, bool)`

GetUniverseLogsExporterConfigOk returns a tuple with the UniverseLogsExporterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseLogsExporterConfig

`func (o *QueryLogConfig) SetUniverseLogsExporterConfig(v []UniverseQueryLogsExporterConfig)`

SetUniverseLogsExporterConfig sets UniverseLogsExporterConfig field to given value.


### GetYsqlQueryLogConfig

`func (o *QueryLogConfig) GetYsqlQueryLogConfig() YSQLQueryLogConfig`

GetYsqlQueryLogConfig returns the YsqlQueryLogConfig field if non-nil, zero value otherwise.

### GetYsqlQueryLogConfigOk

`func (o *QueryLogConfig) GetYsqlQueryLogConfigOk() (*YSQLQueryLogConfig, bool)`

GetYsqlQueryLogConfigOk returns a tuple with the YsqlQueryLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlQueryLogConfig

`func (o *QueryLogConfig) SetYsqlQueryLogConfig(v YSQLQueryLogConfig)`

SetYsqlQueryLogConfig sets YsqlQueryLogConfig field to given value.

### HasYsqlQueryLogConfig

`func (o *QueryLogConfig) HasYsqlQueryLogConfig() bool`

HasYsqlQueryLogConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


