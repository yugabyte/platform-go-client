# QueryLogsTelemetrySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**YsqlQueryLogConfig** | Pointer to [**YSQLQueryLogConfig**](YSQLQueryLogConfig.md) |  | [optional] 
**Exporters** | Pointer to [**[]TelemetryExporterEntry**](TelemetryExporterEntry.md) | List of exporters. Empty &#x3D; no export. | [optional] 

## Methods

### NewQueryLogsTelemetrySpec

`func NewQueryLogsTelemetrySpec() *QueryLogsTelemetrySpec`

NewQueryLogsTelemetrySpec instantiates a new QueryLogsTelemetrySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLogsTelemetrySpecWithDefaults

`func NewQueryLogsTelemetrySpecWithDefaults() *QueryLogsTelemetrySpec`

NewQueryLogsTelemetrySpecWithDefaults instantiates a new QueryLogsTelemetrySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYsqlQueryLogConfig

`func (o *QueryLogsTelemetrySpec) GetYsqlQueryLogConfig() YSQLQueryLogConfig`

GetYsqlQueryLogConfig returns the YsqlQueryLogConfig field if non-nil, zero value otherwise.

### GetYsqlQueryLogConfigOk

`func (o *QueryLogsTelemetrySpec) GetYsqlQueryLogConfigOk() (*YSQLQueryLogConfig, bool)`

GetYsqlQueryLogConfigOk returns a tuple with the YsqlQueryLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlQueryLogConfig

`func (o *QueryLogsTelemetrySpec) SetYsqlQueryLogConfig(v YSQLQueryLogConfig)`

SetYsqlQueryLogConfig sets YsqlQueryLogConfig field to given value.

### HasYsqlQueryLogConfig

`func (o *QueryLogsTelemetrySpec) HasYsqlQueryLogConfig() bool`

HasYsqlQueryLogConfig returns a boolean if a field has been set.

### GetExporters

`func (o *QueryLogsTelemetrySpec) GetExporters() []TelemetryExporterEntry`

GetExporters returns the Exporters field if non-nil, zero value otherwise.

### GetExportersOk

`func (o *QueryLogsTelemetrySpec) GetExportersOk() (*[]TelemetryExporterEntry, bool)`

GetExportersOk returns a tuple with the Exporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporters

`func (o *QueryLogsTelemetrySpec) SetExporters(v []TelemetryExporterEntry)`

SetExporters sets Exporters field to given value.

### HasExporters

`func (o *QueryLogsTelemetrySpec) HasExporters() bool`

HasExporters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


